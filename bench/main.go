package main

import (
	"flag"
	"fmt"
	//"math/big"
	"strings"

	sdk "github.com/dappledger/ann-go-sdk"
)

var (
	byteCode   = `608060405234801561001057600080fd5b5061020c806100206000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806326f34b8b146100725780633c6bb436146100b35780633d4197f0146100de578063862c242e1461010b578063e1cb0e5214610142575b600080fd5b34801561007e57600080fd5b5061009d6004803603810190808035906020019092919050505061016d565b6040518082815260200191505060405180910390f35b3480156100bf57600080fd5b506100c861018d565b6040518082815260200191505060405180910390f35b3480156100ea57600080fd5b5061010960048036038101908080359060200190929190505050610193565b005b34801561011757600080fd5b50610140600480360381019080803590602001909291908035906020019092919050505061019d565b005b34801561014e57600080fd5b506101576101d7565b6040518082815260200191505060405180910390f35b600060016000838152602001908152602001600020600101549050919050565b60005481565b8060008190555050565b8160016000848152602001908152602001600020600001819055508060016000848152602001908152602001600020600101819055505050565b600080549050905600a165627a7a7230582090454576ac53e48f93db33e3502a6c1c9bff38697b8035a76500dc8ab84056b50029`
	abi        = `[{"constant": true,"inputs": [{"name": "_no","type": "uint256"}],"name": "getBatchVal","outputs": [{"name": "","type": "uint256"}],"payable": false,"stateMutability": "view","type": "function"},{"constant": true,"inputs": [],"name": "val","outputs": [{"name": "","type": "uint256"}],"payable": false,"stateMutability": "view","type": "function"},{"constant": false,"inputs": [{"name": "_val","type": "uint256"}],"name": "setVal","outputs": [],"payable": false,"stateMutability": "nonpayable","type": "function"},{"constant": false,"inputs": [{"name": "_no","type": "uint256"},{"name": "_val","type": "uint256"}],"name": "setBatchVal","outputs": [],"payable": false,"stateMutability": "nonpayable","type": "function"},{"constant": true,"inputs": [],"name": "getVal","outputs": [{"name": "","type": "uint256"}],"payable": false,"stateMutability": "view","type": "function"}]`
	contractId string
)

func genAccBase(cli *sdk.GoSDK) sdk.AccountBase {
	acc, _ := cli.AccountCreate()
	return sdk.AccountBase{
		acc.Privkey,
		0,
	}
}

func doTask_setValue(dosync bool, cli *sdk.GoSDK) {
	req := sdk.Tx{
		AccountBase: genAccBase(cli),
		To:          "4F453058FAACAFC4D9AC17E2B5BEE161EE66145D",
		Payload:     "1",
	}
	var err error
	if dosync {
		_, err = cli.Transaction(&req)
	} else {
		_, err = cli.TransactionAsync(&req)
	}
	if err != nil {
		fmt.Println("setValue:", err)
	}
}

func doTask_getValue(hash string, cli *sdk.GoSDK) {
	_, err := cli.TransactionPayLoad(hash)
	if err != nil {
		fmt.Printf("hash(%s)err:%s\n", hash, err.Error())
	}
}

func doTask_contractWrite(dosync bool, cli *sdk.GoSDK) error {
	req := sdk.ContractMethod{
		AccountBase: genAccBase(cli),
		ABI:         abi,
		Contract:    contractId,
		Method:      "setVal",
		Params:      []interface{}{169},
	}
	var err error
	if dosync {
		_, err = cli.ContractCall(&req)
	} else {
		_, err = cli.ContractAsync(&req)
	}
	return err
}

func doTask_contractRead(addr string, cli *sdk.GoSDK) {
	req := sdk.ContractMethod{
		AccountBase: genAccBase(cli),
		ABI:         abi,
		Contract:    addr,
		Method:      "getVal",
		Params:      nil,
	}
	_, err := cli.ContractRead(&req)
	if err != nil {
		fmt.Printf("getVal err:%s\n", err.Error())
	} else {
		// res := resp.([]interface{})
		// rv := res[0].(*big.Int)
		// if 0 != bigintv.Cmp(rv) {
		// 	fmt.Println("resp=", resp)
		// }
	}
}

var (
	tps     = flag.Int("tps", 100, "tps for client<n msg in 1 sec>")
	nsec    = flag.Int("cost", 100, "run n second")
	nthread = flag.Int("nt", 10, "thread of number")
	rpcarg  = flag.String("rpcaddrs", "172.28.152.153:46657,172.28.152.154:46657,172.28.152.155:46657,172.28.152.156:46657", "address of chain;seperate by ',' ") //default value = "172.28.152.153:46657,172.28.152.154:46657,172.28.152.155:46657,172.28.152.156:46657"
	dosync  = flag.Bool("sync", false, "if call contract-write/setvalue,should sync or async;default is 'async'")
	dotype  = flag.String("dotype", "", "witch to do;default is 'set' for set value;\n\t'cwrite' : contract write;")
	//dotype  = flag.String("dotype", "", "witch to do;default is 'set' for set value;\n\t'set' : set value;\n\t'get' : get value by hash;\n\t'cwrite' : contract write;\n\t'cread' : contract read;")
	//gethash = flag.String("hash", "", "if get value;set hash")
	//caddr   = flag.String("caddr", "", "address of contract for contract read")
	debug = flag.Bool("debug", false, "debug for this benchmark")
)

var (
	rpcClis []*sdk.GoSDK
)

const (
	dt_set    = "set"
	dt_get    = "get"
	dt_cwrite = "cwrite"
	dt_cread  = "cread"
)

func Init() {
	var rpcaddrs []string
	arr := strings.Split(*rpcarg, ",")
	if len(arr) > 0 && arr[0] != "" {
		rpcaddrs = arr
	}
	if len(rpcaddrs) == 0 {
		panic("no rpcs to connect")
	}
	for _, addr := range rpcaddrs {
		cli := sdk.New(addr, sdk.ZaCryptoType)
		rpcClis = append(rpcClis, cli)
	}
	//
	var output = "\n"
	if *debug {
		output += "\n\n-----------------> current is debug module <-------------------\n\n"
	}
	output = fmt.Sprintf("tps=%d;nsec=%d;nthread=%d;\n", *tps, *nsec, *nthread)
	output += fmt.Sprintf("rpcaddr=%s;\n", *rpcarg)
	switch *dotype {
	case dt_set:
		if *dosync {
			output += fmt.Sprintf("set value(Transaction)")
		} else {
			output += fmt.Sprintf("set value(TransactionAsync)")
		}

	// case dt_get:
	// 	if *gethash == "" {
	// 		fmt.Println("please provide hash for get value")
	// 		panic("set hash ")
	// 	} else {
	// 		output += fmt.Sprintf("get value(TransactionPayLoad) read from ", *gethash)
	// 	}
	case dt_cwrite:
		//创建合约;
		var arg = sdk.ContractCreate{
			AccountBase: genAccBase(rpcClis[0]),
			ABI:         abi,
			Code:        byteCode,
		}
		result, err := rpcClis[0].ContractCreate(&arg)
		if err != nil {
			fmt.Println("ContractCreate err:", err)
			panic(err)
		}
		contractId = result["contract"].(string)
		if *dosync {
			output += fmt.Sprintln("contract write(sync) contractId=", contractId)
		} else {
			output += fmt.Sprintln("contract write(Async) contractId=", contractId)
		}
	// case dt_cread:
	// 	if *caddr == "" {
	// 		fmt.Println("please provide caddr(contract-addr) for contract read")
	// 		panic("contract read ")
	// 	} else {
	// 		output += fmt.Sprintf("contract read from ", *caddr)
	// 	}
	default:
		panic("pleas set 'dotype' first")
	}
	fmt.Println(output)
}

func getClient(idx int) *sdk.GoSDK {
	idx = idx % len(rpcClis)
	return rpcClis[idx]
}

func dotask(b *Bench) {
	if *tps > 1000 {
		str := fmt.Sprintf("tps(%d) is too big;please small than 1000;", *tps)
		panic(str)
	}
	switch *dotype {
	case dt_set:
		b.Do = func(work int) BCRet {
			doTask_setValue(*dosync, getClient(work))
			return BCSuccess
		}
	// case dt_get:
	// 	hash := *gethash
	// 	b.Do = func(work int) BCRet {
	// 		doTask_getValue(hash, getClient(work))
	// 		return BCSuccess
	// 	}
	case dt_cwrite:
		//合约写;
		b.Do = func(work int) BCRet {
			doTask_contractWrite(*dosync, getClient(work))
			return BCSuccess
		}
		// case dt_cread:

		// 	b.Do = func(work int) BCRet {
		// 		doTask_contractRead(*caddr, getClient(work))
		// 		return BCSuccess
		// 	}
	}

	fmt.Println("---->benchmark running...\n")
	b.Running()
	fmt.Println("\n---->benchmark exit.")
}

func main() {
	flag.Parse()
	Init()
	bench := Bench{
		BenchBase: BenchBase{
			Nthread:      *nthread,
			Ntps:         *tps,
			total1Thread: (*nsec) * (*tps),
			Do: func(idx int) BCRet {
				return BCSuccess
			},
		},
		stop: 0,
		dbg:  *debug,
	}
	dotask(&bench)
}
