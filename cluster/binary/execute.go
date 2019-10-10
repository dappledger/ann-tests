package binary

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/dappledger/AnnChain/gemmill/consensus/raft"
	"github.com/dappledger/AnnChain/gemmill/go-wire"
	"github.com/dappledger/AnnChain/gemmill/types"
	"github.com/dappledger/ann-tests/cluster"
)

type Cluster struct {
	Validators []*validator
	Opt        *Option
}

type validator struct {
	Name          string
	HostName      string
	ChainID       string
	P2PPort       string
	RPCPort       string
	BindPort      string
	Seeds         string
	SignByCA      string
	Address       string
	PrivValidator string
	GenesisDoc    string
	PrivKey       string
	PubKey        string
	IP            string
	CryptoType    string
	Home          string
	ClusterConf   string
	Consensus     string
	stdlogFile    *os.File
	errlogFile    *os.File
	cmd           *exec.Cmd
}

const (
	configToml = `app_name = "evm"
auth_by_ca = false
block_size = 5000
crypto_type = "%s"
db_backend = "leveldb"
db_conn_str = "sqlite3"
db_type = "sqlite3"
environment = "production"
fast_sync = true
log_path = "%s"
moniker = "anonymous"
non_validator_auth_by_ca = false
non_validator_node_auth = false
p2p_laddr = "tcp://0.0.0.0:%s"
rpc_laddr = "tcp://0.0.0.0:%s"
seeds = "%s"
signbyca = ""
skip_upnp = true
threshold_blocks = 0
tracerouter_msg_ttl = 5
network_rate_limit = 1024
consensus = "%s"`
)

func (v *validator) init() error {

	if err := os.MkdirAll(filepath.Join(v.Home, "data"), 0777); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(v.Home, "log"), 0777); err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath.Join(v.Home, "genesis.json"), []byte(v.GenesisDoc), 0666); err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath.Join(v.Home, "priv_validator.json"), []byte(v.PrivValidator), 0666); err != nil {
		return err
	}

	if v.Consensus == "raft" {
		if err := ioutil.WriteFile(filepath.Join(v.Home, "raft-cluster.json"), []byte(v.ClusterConf), 0666); err != nil {
			return err
		}
	}

	config := fmt.Sprintf(configToml, v.CryptoType, "genesis.log", v.P2PPort, v.RPCPort, v.Seeds, v.Consensus)
	if err := ioutil.WriteFile(filepath.Join(v.Home, "config.toml"), []byte(config), 0666); err != nil {
		return err
	}
	return nil
}

type Option struct {
	ValidatorNum int
	CryptoType   string
	Workdir      string
	BinaryPath   string
	Consensus    string
}

func makeValidatorsPorts(num int) (p2p []string, rpc []string, bind []string) {

	for i := 0; i < num; i++ {
		p2p = append(p2p, fmt.Sprintf("%d", 46100+i))
		rpc = append(rpc, fmt.Sprintf("%d", 47100+i))
		bind = append(bind, fmt.Sprintf("%d", 23000+i))
	}
	return
}

func makeValidatorsIp(num int) []string {
	ips := make([]string, 0, num)
	for i := 0; i < num; i++ {
		ips = append(ips, "0.0.0.0")
	}
	return ips
}

func concatSeeds(ips, ports []string) string {

	seeds := make([]string, 0, len(ips))
	for i, _ := range ips {
		seeds = append(seeds, fmt.Sprintf("%s:%s", ips[i], ports[i]))
	}
	return strings.Join(seeds, ",")
}

func New(opt *Option) (*Cluster, error) {

	checkOpt(opt)
	ips := makeValidatorsIp(opt.ValidatorNum)
	p2pPorts, rpcPorts, bindPorts := makeValidatorsPorts(opt.ValidatorNum)
	chainID := "9102"
	seeds := concatSeeds(ips, p2pPorts)

	validators := make([]*validator, 0, opt.ValidatorNum)

	genDoc := &types.GenesisDoc{
		ChainID:     chainID,
		Plugins:     "adminOp,querycache",
		GenesisTime: time.Now(),
		Validators:  make([]types.GenesisValidator, 0, opt.ValidatorNum),
	}

	clusterConf := &raft.ClusterConfig{}

	for i := 0; i < opt.ValidatorNum; i++ {

		privVal, err := types.GenPrivValidator(opt.CryptoType, nil)
		if err != nil {
			panic(err)
		}
		key := privVal.PrivKey
		jsonBytes := wire.JSONBytes(privVal)
		pubkey := key.PubKey()

		v := validator{
			Name:          fmt.Sprintf("validator-%d", i),
			HostName:      fmt.Sprintf("validator-%d", i),
			ChainID:       chainID,
			P2PPort:       p2pPorts[i],
			RPCPort:       rpcPorts[i],
			BindPort:      bindPorts[i],
			Seeds:         seeds,
			PrivValidator: string(jsonBytes),
			Address:       fmt.Sprintf("%X", pubkey.Address()),
			PrivKey:       key.KeyString(),
			PubKey:        pubkey.KeyString(),
			IP:            ips[i],
			CryptoType:    opt.CryptoType,
			Home:          filepath.Join(opt.Workdir, fmt.Sprintf("%d", i)),
			Consensus:     opt.Consensus,
		}
		validators = append(validators, &v)

		genDoc.Validators = append(genDoc.Validators, types.GenesisValidator{
			PubKey: key.PubKey(),
			Amount: 100,
		})

		if opt.Consensus == "raft" {
			clusterConf.Peers = append(clusterConf.Peers, raft.Peer{
				PubKey: pubkey,
				RPC:    fmt.Sprintf("127.0.0.1:%s", rpcPorts[i]),
				Bind:   fmt.Sprintf("127.0.0.1:%s", bindPorts[i]),
			})
		}
	}

	genesisDoc := wire.JSONBytes(genDoc)
	for i, _ := range validators {
		validators[i].GenesisDoc = string(genesisDoc)

		if opt.Consensus == "raft" {
			peer := clusterConf.Peers[i]

			clusterConf.Advertise = peer.Bind
			clusterConf.Local.Bind = peer.Bind
			clusterConf.Local.RPC = peer.RPC
			clusterConf.Local.PubKey = peer.PubKey

			clusterConfDoc := wire.JSONBytes(clusterConf)
			validators[i].ClusterConf = string(clusterConfDoc)
		}
	}

	c := Cluster{
		Validators: validators,
		Opt:        opt,
	}

	if _, err := os.Stat(opt.Workdir); err == nil {
		return nil, errors.New(opt.Workdir + "not empty")
	}
	if err := os.MkdirAll(filepath.Join(opt.Workdir), 0777); err != nil {
		return nil, err
	}

	for _, v := range validators {
		if err := v.init(); err != nil {
			return nil, err
		}
	}
	return &c, nil
}

func checkOpt(opt *Option) {
	if opt.Workdir == "" {
		opt.Workdir = filepath.Join(os.TempDir(), fmt.Sprintf("annchain-cluster-%d", time.Now().Unix()))
	}

	if opt.BinaryPath == "" {
		opt.BinaryPath = "genesis"
	}

	if opt.ValidatorNum == 0 {
		opt.ValidatorNum = 4
	}
}

func (c *Cluster) Up() error {

	for i, _ := range c.Validators {
		if err := c.StartValidator(i); err != nil {
			return err
		}
	}
	time.Sleep(time.Second * 3)
	return nil
}

func (c *Cluster) Down() error {
	for _, v := range c.Validators {
		if v.cmd != nil {
			if err := v.cmd.Process.Kill(); err != nil {
				return err
			}
		}
	}
	return os.RemoveAll(c.Opt.Workdir)
}

func (c *Cluster) ValidatorNum() int {
	return len(c.Validators)
}

func (c *Cluster) GetValidatorInfo(index int) (*cluster.ValidatorInfo, error) {
	return &cluster.ValidatorInfo{
		RPC:     "127.0.0.1:" + c.Validators[index].RPCPort,
		Address: c.Validators[index].Address,
		PrivKey: c.Validators[index].PrivKey,
	}, nil
}

func (c *Cluster) StartValidator(index int) error {

	v := c.Validators[index]
	if v.cmd != nil {
		return nil
	}
	cmd := exec.Command(c.Opt.BinaryPath, "run", "--runtime", v.Home)
	cmd.Dir = v.Home
	if err := cmd.Start(); err != nil {
		return err
	}
	v.cmd = cmd
	return nil
}

func (c *Cluster) StopValidator(index int) error {
	v := c.Validators[index]
	if v.cmd == nil {
		return nil
	}
	if err := v.cmd.Process.Kill(); err != nil {
		return err
	}
	v.cmd = nil
	return nil
}
