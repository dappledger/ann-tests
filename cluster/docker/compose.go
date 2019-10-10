package docker

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/dappledger/AnnChain/gemmill/go-crypto"
	"github.com/dappledger/AnnChain/gemmill/go-wire"
	"github.com/dappledger/AnnChain/gemmill/types"
	"github.com/dappledger/ann-tests/cluster"
	dockerTypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

const (
	composeTplStr = `version: '3'
services:
  {{- range .Services }}
  {{ . }}
  {{- end }}
networks:
  app_net:
    driver: bridge
    ipam:
      driver: default
      config:
      - subnet: {{ .IPPrefix }}.0/24`

	mongoDBTplStr = `mgo:
    image: {{.Image}}
    labels:
      app: db
      cluster: {{.ClusterName}}
    ports:
      - '{{.BindPort}}:{{.BindPort}}'
    command: mongod
    restart: always
    networks:
      app_net:
        ipv4_address: {{.IP}}`

	blockBrowserServiceTplStr = `block-browser:
    image: {{.Image}}
    ports:
      - '{{.BindPort}}:{{.BindPort}}'
    labels:
      app: browser
      cluster: {{.ClusterName}}
    entrypoint:
      - /bin/sh
      - -c
      - |
        echo 'appname = hubble
              httpport = {{.BindPort}}
              runmode ="dev"

              # blockchain's api ,not app's api
              api_addr = "{{.ValidatorRPC}}"
              chain_id = "1024"

              mogo_addr = "{{.MgoDB}}"
              #mogo_user = hzc
              #mogo_pwd = 123456

              mogo_db = "ann-browser"
              sync_job = 1
              sync_from_height = 0' > conf/app.conf
        block-browser
    restart: always
    networks:
      app_net:
        ipv4_address: {{.IP}}`

	genesisAPIServiceTplStr = `genesis-api:
    image: {{.Image}}
    ports:
      - '{{.BindPort}}:{{.BindPort}}'
    labels:
      app: api
      cluster: {{.ClusterName}}
    entrypoint:
      - /bin/sh
      - -c
      - |
        echo '{"ListenAddress": ":{{.BindPort}}", "RPCAddress": "tcp://{{.ValidatorRPC}}", "GasLimit" : 10000000000 }' > api.conf
        genesis-api -f api.conf
    restart: always
    networks:
      app_net:
        ipv4_address: {{.IP}}`

	validatorServiceTplStr = `{{.Name}}:
    hostname: {{.HostName}}
    image: {{.Image}}
    labels:
      app: validator
      app_id: {{.ID}}
      cluster: {{.ClusterName}}
    ports:
      - '{{.P2PPort}}:{{.P2PPort}}'
      - '{{.RPCPort}}:{{.RPCPort}}'
    entrypoint:
      - /bin/sh
      - -c
      - |
        mkdir -p logs
        mkdir -p /genesis
        genesis init --runtime="/genesis" --log_path="logs/genesis.log" --app="evm" --chainid={{.ChainID}}
        echo '{{.GenesisDoc}}' > /genesis/genesis.json
        echo 'app_name = "evm"
              auth_by_ca = false
              block_size = 5000
              crypto_type = "{{.CryptoType}}"
              db_backend = "leveldb"
              db_conn_str = "sqlite3"
              db_type = "sqlite3"
              environment = "production"
              fast_sync = true
              log_path = ""
              moniker = "anonymous"
              non_validator_auth_by_ca = false
              non_validator_node_auth = false
              p2p_laddr = "tcp://0.0.0.0:{{.P2PPort}}"
              rpc_laddr = "tcp://0.0.0.0:{{.RPCPort}}"
              seeds = "{{.Seeds}}"
              signbyca = "{{.SignByCA}}"
              skip_upnp = true
              threshold_blocks = 0
              tracerouter_msg_ttl = 5
              network_rate_limit = 1024' > /genesis/config.toml
        echo '{{.PrivValidator}}' > /genesis/priv_validator.json
        genesis run --runtime="/genesis"
    networks:
      app_net:
        ipv4_address: {{.IP}}
    restart: always`
)

type mgo struct {
	BindPort    string
	Image       string
	IP          string
	ClusterName string
	dockerTypes.Container
}

type blockBrowser struct {
	BindPort     string
	Image        string
	IP           string
	MgoDB        string
	ValidatorRPC string
	ClusterName  string
	dockerTypes.Container
}

type genesisAPI struct {
	BindPort     string
	Image        string
	ValidatorRPC string
	IP           string
	ClusterName  string
	dockerTypes.Container
}

type validator struct {
	ID            int
	Name          string
	Image         string
	HostName      string
	ChainID       string
	P2PPort       string
	RPCPort       string
	Seeds         string
	SignByCA      string
	Address       string
	PrivValidator string
	GenesisDoc    string
	PrivKey       string
	PubKey        string
	IP            string
	CryptoType    string
	ClusterName   string
	dockerTypes.Container
}

type DockerCompose struct {
	Browser    *blockBrowser
	API        *genesisAPI
	DB         *mgo
	Validators []validator
	runtimeDir string
	Opt        *Option
}

type Option struct {
	ValidatorNum      int
	IPPrefix          string
	HasBrowser        bool
	HasAPI            bool
	GenesisImage      string
	GenesisAPIImage   string
	BlockBrowserImage string
	CryptoType        string
}

var DefaultOpt = Option{
	ValidatorNum:      1,
	IPPrefix:          "192.168.32",
	HasBrowser:        true,
	HasAPI:            true,
	GenesisImage:      "annchain/genesis:latest",
	GenesisAPIImage:   "genesis-api:latest",
	BlockBrowserImage: "block-browser:latest",
	CryptoType:        crypto.CryptoTypeZhongAn,
}

func NewDockerCompose(opt *Option) *DockerCompose {

	if opt == nil {
		opt = &DefaultOpt
	}

	runtimeDir := filepath.Join(os.TempDir(), fmt.Sprintf("ann-docker-compose-%d", time.Now().UnixNano()))
	clusterName := strings.Replace(runtimeDir, "-", "", -1)
	validators := setupValidators(opt.ValidatorNum, opt.CryptoType, opt.IPPrefix, opt.GenesisImage, clusterName)
	compose := &DockerCompose{
		Validators: validators,
		runtimeDir: runtimeDir,
		Opt:        opt,
	}

	if compose.Opt.HasBrowser {

		m := mgo{
			BindPort:    "27017",
			Image:       "mongo:3.2",
			IP:          compose.Opt.IPPrefix + ".7",
			ClusterName: clusterName,
		}
		browser := blockBrowser{
			BindPort:     "9090",
			MgoDB:        m.IP + ":" + m.BindPort,
			Image:        opt.BlockBrowserImage,
			IP:           compose.Opt.IPPrefix + ".8",
			ValidatorRPC: validators[0].HostName + ":" + validators[0].RPCPort,
			ClusterName:  clusterName,
		}
		compose.DB = &m
		compose.Browser = &browser
	}

	if compose.Opt.HasAPI {
		api := genesisAPI{
			BindPort:     "8800",
			Image:        opt.GenesisAPIImage,
			ValidatorRPC: validators[0].HostName + ":" + validators[0].RPCPort,
			IP:           opt.IPPrefix + ".9",
			ClusterName:  clusterName,
		}
		compose.API = &api
	}
	return compose
}

func (d *DockerCompose) PrintInfo() {

	if d.Opt.HasBrowser {
		fmt.Printf("block-browser: http://127.0.0.1:%s\tcontainerId:%s\n", d.Browser.BindPort, d.Browser.Container.ID)
		fmt.Printf("mongoDB: http://127.0.0.1:%s\tcontainerId:%s\n", d.DB.BindPort, d.DB.Container.ID)
	}
	if d.Opt.HasAPI {
		fmt.Printf("genesis-api: http://127.0.0.1:%s\tcontainerId:%s\n", d.API.BindPort, d.API.Container.ID)
	}
	fmt.Println("validators:")
	for _, v := range d.Validators {
		fmt.Printf("\t%s\tcontainerId:%s\n", v.Name, v.Container.ID)
	}
	fmt.Println("workDir:", d.runtimeDir)
}

func (d *DockerCompose) ValidatorNum() int {
	return len(d.Validators)
}

func (d *DockerCompose) Up() error {

	if err := os.MkdirAll(d.runtimeDir, 0777); err != nil {
		return err
	}

	cfg, err := d.GenerateConfig()
	if err != nil {
		return err
	}

	ymlFile := filepath.Join(d.runtimeDir, "docker-compose.yml")
	if err := ioutil.WriteFile(ymlFile, []byte(cfg), 0666); err != nil {
		return err
	}

	errBuff := bytes.Buffer{}

	cmd := exec.Command("docker-compose", "up")
	cmd.Dir = d.runtimeDir
	cmd.Stderr = &errBuff

	errChan := make(chan error, 1)
	go func() {
		err = cmd.Run()
		if err != nil {

			cmd := exec.Command("docker", "ps")
			cmd.Stdout = os.Stdout
			cmd.Run()
			err = fmt.Errorf("%v, %s %s", err, errBuff.String(), d.runtimeDir)
			errChan <- err
		}
	}()

	cli, err := client.NewClient(client.DefaultDockerHost, "1.39", nil, nil)
	if err != nil {
		return err
	}

	wait := func(cli *client.Client) (error, bool) {

		args := filters.NewArgs()
		args.Add("label", "cluster="+d.Validators[0].ClusterName)

		allContainers, err := cli.ContainerList(context.Background(), dockerTypes.ContainerListOptions{Filters: args})
		if err != nil {
			return err, false
		}

		expectContainersNum := len(d.Validators)

		if d.Opt.HasAPI {
			expectContainersNum += 1
		}

		if d.Opt.HasBrowser {
			expectContainersNum += 2
		}

		if len(allContainers) != expectContainersNum {
			return nil, false
		}

		for i, v := range allContainers {
			if !strings.HasPrefix(v.Status, "Up ") {
				return nil, false
			}
			switch v.Labels["app"] {
			case "api":
				d.API.Container = allContainers[i]
			case "db":
				d.DB.Container = allContainers[i]
			case "browser":
				d.Browser.Container = allContainers[i]
			case "validator":
				index, _ := strconv.Atoi(v.Labels["app_id"])
				d.Validators[index].Container = allContainers[i]
			}
		}
		return nil, true
	}

	for {
		select {
		case err := <-errChan:
			return err
		default:
			if err, done := wait(cli); err != nil {
				return err
			} else if done {
				return nil
			}
		}
		time.Sleep(time.Second)
	}
}

func (d *DockerCompose) Down() error {

	var err error
	defer func() {
		err = os.RemoveAll(d.runtimeDir)
	}()

	cmd := exec.Command("docker-compose", "down")
	cmd.Dir = d.runtimeDir
	errBuff := bytes.Buffer{}
	cmd.Stderr = &errBuff
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%v:%s", err, errBuff.String())
	}
	return err
}

func (d *DockerCompose) GetValidatorInfo(index int) (*cluster.ValidatorInfo, error) {

	return &cluster.ValidatorInfo{
		RPC:     "127.0.0.1:" + d.Validators[index].RPCPort,
		Address: d.Validators[index].Address,
		PrivKey: d.Validators[index].PrivKey,
	}, nil
}

func (d *DockerCompose) StopValidator(index int) error {

	if len(d.Validators) < index {
		return errors.New("invalid index")
	}

	cli, err := client.NewClient(client.DefaultDockerHost, "1.39", nil, nil)
	if err != nil {
		return err
	}
	timeout := time.Second
	err = cli.ContainerStop(context.Background(), d.Validators[index].Container.ID, &timeout)
	if err != nil {
		return err
	}

	args := filters.NewArgs()
	args.Add("id", d.Validators[index].Container.ID)
	for i := 0; i < 100; i++ {
		containers, err := cli.ContainerList(context.Background(), dockerTypes.ContainerListOptions{Filters: args})
		if err != nil {
			return err
		}
		if len(containers) == 0 {
			return nil
		}
		time.Sleep(time.Second * 2)
	}
	return errors.New("wait stop validator timeout")
}

func (d *DockerCompose) StartValidator(index int) error {

	if len(d.Validators) < index {
		return errors.New("invalid index")
	}

	cli, err := client.NewClient(client.DefaultDockerHost, "1.39", nil, nil)
	if err != nil {
		return err
	}

	err = cli.ContainerStart(context.Background(), d.Validators[index].Container.ID, dockerTypes.ContainerStartOptions{})
	if err != nil {
		return err
	}

	args := filters.NewArgs()
	args.Add("id", d.Validators[index].Container.ID)
	for i := 0; i < 100; i++ {
		containers, err := cli.ContainerList(context.Background(), dockerTypes.ContainerListOptions{Filters: args})
		if err != nil {
			return err
		}
		if len(containers) == 1 {
			return nil
		}
		time.Sleep(time.Second * 2)
	}
	return errors.New("wait start validator timeout")
}

func (d *DockerCompose) GenerateConfig() (string, error) {

	servicesTpl := make([]string, 0, len(d.Validators)+2)

	if d.Opt.HasBrowser {

		mgoTpl, err := executeTpl(d.DB, "mgo", mongoDBTplStr)
		if err != nil {
			return "", err
		}
		servicesTpl = append(servicesTpl, mgoTpl)

		browserTpl, err := executeTpl(d.Browser, "browser", blockBrowserServiceTplStr)
		if err != nil {
			return "", err
		}
		servicesTpl = append(servicesTpl, browserTpl)
	}

	if d.Opt.HasAPI {

		apiTpl, err := executeTpl(d.API, "api", genesisAPIServiceTplStr)
		if err != nil {
			return "", err
		}
		servicesTpl = append(servicesTpl, apiTpl)
	}

	for _, v := range d.Validators {
		tpl, err := executeTpl(v, "genesis", validatorServiceTplStr)
		if err != nil {
			return "", err
		}
		servicesTpl = append(servicesTpl, tpl)
	}

	compose := map[string]interface{}{
		"Services": servicesTpl,
		"IPPrefix": d.Opt.IPPrefix,
	}

	composeTpl, err := executeTpl(compose, "compose", composeTplStr)
	if err != nil {
		return "", err
	}
	return composeTpl, nil
}

func makeValidatorsIp(ipPrefix string, num int) []string {
	ips := make([]string, 0, num)
	for i := 0; i < num; i++ {
		ips = append(ips, ipPrefix+fmt.Sprintf(".%d", 10+i))
	}
	return ips
}

func makeValidatorsPorts(num int) (p2p []string, rpc []string) {

	for i := 0; i < num; i++ {
		p2p = append(p2p, fmt.Sprintf("%d", 46000+i))
		rpc = append(rpc, fmt.Sprintf("%d", 47000+i))
	}
	return
}

func makeSeeds(ips, ports []string) string {
	seeds := make([]string, 0, len(ips))

	for k, _ := range ips {
		seeds = append(seeds, ips[k]+":"+ports[k])
	}
	return strings.Join(seeds, ",")
}

func setupValidators(num int, cryptoType string, ipPrefix string, image, clusterName string) []validator {
	crypto.NodeInit(cryptoType)
	ips := makeValidatorsIp(ipPrefix, num)
	p2pPorts, rpcPorts := makeValidatorsPorts(num)
	seeds := makeSeeds(ips, p2pPorts)

	chainID := "9102"

	validators := make([]validator, 0, num)

	genDoc := &types.GenesisDoc{
		ChainID:    chainID,
		Plugins:    "adminOp,querycache",
		Validators: make([]types.GenesisValidator, 0, num),
	}

	for i := 0; i < num; i++ {

		privVal, err := types.GenPrivValidator(cryptoType, nil)
		if err != nil {
			panic(err)
		}
		key := privVal.PrivKey
		jsonBytes := wire.JSONBytes(privVal)
		pubkey := key.PubKey()
		v := validator{
			ID:            i,
			Name:          fmt.Sprintf("validator-%d", i),
			Image:         image,
			HostName:      fmt.Sprintf("validator-%d", i),
			ChainID:       chainID,
			P2PPort:       p2pPorts[i],
			RPCPort:       rpcPorts[i],
			Seeds:         seeds,
			PrivValidator: string(jsonBytes),
			Address:       fmt.Sprintf("%X", pubkey.Address()),
			PrivKey:       key.KeyString(),
			PubKey:        pubkey.KeyString(),
			IP:            ips[i],
			CryptoType:    cryptoType,
			ClusterName:   clusterName,
		}
		validators = append(validators, v)
		genDoc.Validators = append(genDoc.Validators, types.GenesisValidator{
			PubKey: privVal.PubKey,
			Amount: 100,
			IsCA:   true,
		})
	}

	genesisDoc := wire.JSONBytes(genDoc)
	for i, _ := range validators {
		validators[i].GenesisDoc = string(genesisDoc)
	}
	return validators
}

func executeTpl(obj interface{}, tplName string, tplString string) (string, error) {

	tpl, err := template.New(tplName).Parse(tplString)
	if err != nil {
		return "", err
	}
	cfg := new(bytes.Buffer)
	if err := tpl.Execute(cfg, obj); err != nil {
		return "", err
	}
	return cfg.String(), nil
}
