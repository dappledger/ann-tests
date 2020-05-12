// Copyright © 2017 ZhongAn Technology
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"

	sdk "github.com/dappledger/ann-go-sdk"
	"github.com/dappledger/ann-go-sdk/common"
)

// VariAndFunABI is the input ABI used to generate the binding from.
const VariAndFunABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_x\",\"type\":\"uint256\"},{\"name\":\"_y\",\"type\":\"uint256\"},{\"name\":\"_k\",\"type\":\"uint256\"}],\"name\":\"testMulMod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operationHash\",\"type\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"testEcrecover\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_x\",\"type\":\"uint256\"},{\"name\":\"_y\",\"type\":\"uint256\"},{\"name\":\"_k\",\"type\":\"uint256\"}],\"name\":\"testAddMod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"testAssert\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"testRequire\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMsgSig\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"testSHA256\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMsgSender\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlockNum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlockTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_code\",\"type\":\"bytes\"},{\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"create2Deploy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlockHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getContractAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"testRevert\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThis\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"testExist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMsgData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"testKeccak256\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOrigin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"testRipemd160\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes20\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// VariAndFunBin is the compiled bytecode used for deploying new contracts.
const VariAndFunBin = `608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610f91806100606000396000f3fe6080604052600436106101005760003560e01c63ffffffff168063014066f31461010557806315334417146101685780631ebfbced1461027a5780632b813bc0146102dd578063357815c41461030c5780635a8f6d711461033b57806363138d4f146103a45780637a6ce2e1146104805780637f6c6f10146104d757806387ceff091461050257806394f0f9d81461053b5780639663f88f1461060d5780639e45593914610638578063a26388bb1461068f578063b8368615146106be578063bef4091114610715578063c8e7ca2e1461077e578063c8ec99f11461080e578063df1f29ee146108ea578063f6b0bbf714610941578063f851a44014610a3b575b600080fd5b34801561011157600080fd5b506101526004803603606081101561012857600080fd5b81019080803590602001909291908035906020019092919080359060200190929190505050610a92565b6040518082815260200191505060405180910390f35b34801561017457600080fd5b506102386004803603604081101561018b57600080fd5b8101908080359060200190929190803590602001906401000000008111156101b257600080fd5b8201836020820111156101c457600080fd5b803590602001918460018302840111640100000000831117156101e657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610aaa565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561028657600080fd5b506102c76004803603606081101561029d57600080fd5b81019080803590602001909291908035906020019092919080359060200190929190505050610b5e565b6040518082815260200191505060405180910390f35b3480156102e957600080fd5b506102f2610b76565b604051808215151515815260200191505060405180910390f35b34801561031857600080fd5b50610321610bd7565b604051808215151515815260200191505060405180910390f35b34801561034757600080fd5b50610350610c3b565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b3480156103b057600080fd5b5061046a600480360360208110156103c757600080fd5b81019080803590602001906401000000008111156103e457600080fd5b8201836020820111156103f657600080fd5b8035906020019184600183028401116401000000008311171561041857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610c66565b6040518082815260200191505060405180910390f35b34801561048c57600080fd5b50610495610d0f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104e357600080fd5b506104ec610d17565b6040518082815260200191505060405180910390f35b34801561050e57600080fd5b50610517610d1f565b60405180848152602001838152602001828152602001935050505060405180910390f35b34801561054757600080fd5b5061060b6004803603604081101561055e57600080fd5b810190808035906020019064010000000081111561057b57600080fd5b82018360208201111561058d57600080fd5b803590602001918460018302840111640100000000831117156105af57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190505050610d32565b005b34801561061957600080fd5b50610622610d93565b6040518082815260200191505060405180910390f35b34801561064457600080fd5b5061064d610da4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561069b57600080fd5b506106a4610dce565b604051808215151515815260200191505060405180910390f35b3480156106ca57600080fd5b506106d3610e32565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561072157600080fd5b506107646004803603602081101561073857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e3a565b604051808215151515815260200191505060405180910390f35b34801561078a57600080fd5b50610793610e4d565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156107d35780820151818401526020810190506107b8565b50505050905090810190601f1680156108005780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561081a57600080fd5b506108d46004803603602081101561083157600080fd5b810190808035906020019064010000000081111561084e57600080fd5b82018360208201111561086057600080fd5b8035906020019184600183028401116401000000008311171561088257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610e9a565b6040518082815260200191505060405180910390f35b3480156108f657600080fd5b506108ff610eab565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561094d57600080fd5b50610a076004803603602081101561096457600080fd5b810190808035906020019064010000000081111561098157600080fd5b82018360208201111561099357600080fd5b803590602001918460018302840111640100000000831117156109b557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610eb3565b60405180826bffffffffffffffffffffffff19166bffffffffffffffffffffffff1916815260200191505060405180910390f35b348015610a4757600080fd5b50610a50610f40565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600081801515610a9e57fe5b83850990509392505050565b600060418251141515610abc57600080fd5b6000806000602085015192506040850151915060ff6041860151169050601b8160ff161015610aec57601b810190505b60018682858560405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610b49573d6000803e3d6000fd5b50505060206040510351935050505092915050565b600081801515610b6a57fe5b83850890509392505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610bd057fe5b6001905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610c3457600080fd5b6001905090565b600080357fffffffff0000000000000000000000000000000000000000000000000000000016905090565b60006002826040518082805190602001908083835b602083101515610ca05780518252602082019150602081019050602083039250610c7b565b6001836020036101000a038019825116818451168082178552505050505050905001915050602060405180830381855afa158015610ce2573d6000803e3d6000fd5b5050506040513d6020811015610cf757600080fd5b81019080805190602001909291905050509050919050565b600033905090565b600043905090565b6000806000424243925092509250909192565b6000818351602085016000f59050803b1515610d4d57600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b600080600143034090508091505090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610e2b57600080fd5b6001905090565b600030905090565b600080823b905060008111915050919050565b60606000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050905090565b600081805190602001209050919050565b600032905090565b60006003826040518082805190602001908083835b602083101515610eed5780518252602082019150602081019050602083039250610ec8565b6001836020036101000a038019825116818451168082178552505050505050905001915050602060405180830381855afa158015610f2f573d6000803e3d6000fd5b5050506040515160601b9050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff168156fea165627a7a723058205d701cc530f8ce41425b7c261e03e9dacbad2c4bc037dda9a434df915c7c21730029`

// DeployVariAndFun deploys a new Ethereum contract, binding an instance of VariAndFun to it.
func DeployVariAndFun(goSDK *sdk.GoSDK, auth sdk.AccountBase) (common.Hash, *VariAndFun, error) {

	cc := sdk.ContractCreate{
		AccountBase: auth,
		Code:        VariAndFunBin,
		ABI:         VariAndFunABI,
		Params:      []interface{}{},
	}
	ret, err := goSDK.ContractCreate(&cc)
	if err != nil {
		return common.Hash{}, nil, err
	}
	txHash := common.HexToHash(ret["tx"].(string))
	address := common.HexToAddress(ret["contract"].(string))
	return txHash, &VariAndFun{address: address, cli: goSDK}, nil
}

// VariAndFun is an auto generated Go binding around an Ethereum contract.
type VariAndFun struct {
	address common.Address
	cli     *sdk.GoSDK
}

// NewVariAndFun creates a new instance of VariAndFun, bound to a specific deployed contract.
func NewVariAndFun(goSdk *sdk.GoSDK, address common.Address) *VariAndFun {
	return &VariAndFun{address: address, cli: goSdk}
}

func (_VariAndFun *VariAndFun) GetAddress() common.Address {
	return _VariAndFun.address
}

func (_VariAndFun *VariAndFun) Admin(auth sdk.AccountBase) (common.Address, error) {

	var (
		ret0 common.Address
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "admin",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(common.Address)

	return ret0, err
}

func (_VariAndFun *VariAndFun) AdminByHeight(auth sdk.AccountBase, height uint64) (common.Address, error) {
	var (
		ret0 common.Address
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "admin",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(common.Address)

	return ret0, err
}

func (_VariAndFun *VariAndFun) GetBlockHash(auth sdk.AccountBase) ([32]byte, error) {

	var (
		ret0 [32]byte
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "getBlockHash",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].([32]byte)

	return ret0, err
}

func (_VariAndFun *VariAndFun) GetBlockHashByHeight(auth sdk.AccountBase, height uint64) ([32]byte, error) {
	var (
		ret0 [32]byte
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "getBlockHash",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].([32]byte)

	return ret0, err
}

func (_VariAndFun *VariAndFun) GetBlockNum(auth sdk.AccountBase) (*big.Int, error) {

	var (
		ret0 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "getBlockNum",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)

	return ret0, err
}

func (_VariAndFun *VariAndFun) GetBlockNumByHeight(auth sdk.AccountBase, height uint64) (*big.Int, error) {
	var (
		ret0 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "getBlockNum",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)

	return ret0, err
}

func (_VariAndFun *VariAndFun) GetBlockTime(auth sdk.AccountBase) (*big.Int, *big.Int, *big.Int, error) {

	var (
		ret0 *big.Int
		ret1 *big.Int
		ret2 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "getBlockTime",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)
	ret1 = arr[1].(*big.Int)
	ret2 = arr[2].(*big.Int)

	return ret0, ret1, ret2, err
}

func (_VariAndFun *VariAndFun) GetBlockTimeByHeight(auth sdk.AccountBase, height uint64) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 *big.Int
		ret1 *big.Int
		ret2 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "getBlockTime",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)
	ret1 = arr[1].(*big.Int)
	ret2 = arr[2].(*big.Int)

	return ret0, ret1, ret2, err
}

func (_VariAndFun *VariAndFun) GetContractAddr(auth sdk.AccountBase) (common.Address, error) {

	var (
		ret0 common.Address
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "getContractAddr",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(common.Address)

	return ret0, err
}

func (_VariAndFun *VariAndFun) GetContractAddrByHeight(auth sdk.AccountBase, height uint64) (common.Address, error) {
	var (
		ret0 common.Address
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "getContractAddr",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(common.Address)

	return ret0, err
}

func (_VariAndFun *VariAndFun) GetMsgData(auth sdk.AccountBase) ([]byte, error) {

	var (
		ret0 []byte
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "getMsgData",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].([]byte)

	return ret0, err
}

func (_VariAndFun *VariAndFun) GetMsgDataByHeight(auth sdk.AccountBase, height uint64) ([]byte, error) {
	var (
		ret0 []byte
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "getMsgData",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].([]byte)

	return ret0, err
}

func (_VariAndFun *VariAndFun) GetMsgSender(auth sdk.AccountBase) (common.Address, error) {

	var (
		ret0 common.Address
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "getMsgSender",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(common.Address)

	return ret0, err
}

func (_VariAndFun *VariAndFun) GetMsgSenderByHeight(auth sdk.AccountBase, height uint64) (common.Address, error) {
	var (
		ret0 common.Address
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "getMsgSender",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(common.Address)

	return ret0, err
}

func (_VariAndFun *VariAndFun) GetMsgSig(auth sdk.AccountBase) ([4]byte, error) {

	var (
		ret0 [4]byte
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "getMsgSig",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].([4]byte)

	return ret0, err
}

func (_VariAndFun *VariAndFun) GetMsgSigByHeight(auth sdk.AccountBase, height uint64) ([4]byte, error) {
	var (
		ret0 [4]byte
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "getMsgSig",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].([4]byte)

	return ret0, err
}

func (_VariAndFun *VariAndFun) GetOrigin(auth sdk.AccountBase) (common.Address, error) {

	var (
		ret0 common.Address
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "getOrigin",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(common.Address)

	return ret0, err
}

func (_VariAndFun *VariAndFun) GetOriginByHeight(auth sdk.AccountBase, height uint64) (common.Address, error) {
	var (
		ret0 common.Address
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "getOrigin",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(common.Address)

	return ret0, err
}

func (_VariAndFun *VariAndFun) GetThis(auth sdk.AccountBase) (common.Address, error) {

	var (
		ret0 common.Address
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "getThis",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(common.Address)

	return ret0, err
}

func (_VariAndFun *VariAndFun) GetThisByHeight(auth sdk.AccountBase, height uint64) (common.Address, error) {
	var (
		ret0 common.Address
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "getThis",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(common.Address)

	return ret0, err
}

func (_VariAndFun *VariAndFun) TestAddMod(auth sdk.AccountBase, _x *big.Int, _y *big.Int, _k *big.Int) (*big.Int, error) {

	var (
		ret0 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "testAddMod",
		Params:      []interface{}{_x, _y, _k},
	}

	ret, err := _VariAndFun.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)

	return ret0, err
}

func (_VariAndFun *VariAndFun) TestAddModByHeight(auth sdk.AccountBase, _x *big.Int, _y *big.Int, _k *big.Int, height uint64) (*big.Int, error) {
	var (
		ret0 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "testAddMod",
		Params:      []interface{}{_x, _y, _k},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)

	return ret0, err
}

func (_VariAndFun *VariAndFun) TestAssert(auth sdk.AccountBase) (bool, error) {

	var (
		ret0 bool
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "testAssert",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(bool)

	return ret0, err
}

func (_VariAndFun *VariAndFun) TestAssertByHeight(auth sdk.AccountBase, height uint64) (bool, error) {
	var (
		ret0 bool
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "testAssert",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(bool)

	return ret0, err
}

func (_VariAndFun *VariAndFun) TestEcrecover(auth sdk.AccountBase, _operationHash [32]byte, _signature []byte) (common.Address, error) {

	var (
		ret0 common.Address
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "testEcrecover",
		Params:      []interface{}{_operationHash, _signature},
	}

	ret, err := _VariAndFun.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(common.Address)

	return ret0, err
}

func (_VariAndFun *VariAndFun) TestEcrecoverByHeight(auth sdk.AccountBase, _operationHash [32]byte, _signature []byte, height uint64) (common.Address, error) {
	var (
		ret0 common.Address
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "testEcrecover",
		Params:      []interface{}{_operationHash, _signature},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(common.Address)

	return ret0, err
}

func (_VariAndFun *VariAndFun) TestExist(auth sdk.AccountBase, _addr common.Address) (bool, error) {

	var (
		ret0 bool
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "testExist",
		Params:      []interface{}{_addr},
	}

	ret, err := _VariAndFun.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(bool)

	return ret0, err
}

func (_VariAndFun *VariAndFun) TestExistByHeight(auth sdk.AccountBase, _addr common.Address, height uint64) (bool, error) {
	var (
		ret0 bool
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "testExist",
		Params:      []interface{}{_addr},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(bool)

	return ret0, err
}

func (_VariAndFun *VariAndFun) TestKeccak256(auth sdk.AccountBase, _data []byte) ([32]byte, error) {

	var (
		ret0 [32]byte
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "testKeccak256",
		Params:      []interface{}{_data},
	}

	ret, err := _VariAndFun.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].([32]byte)

	return ret0, err
}

func (_VariAndFun *VariAndFun) TestKeccak256ByHeight(auth sdk.AccountBase, _data []byte, height uint64) ([32]byte, error) {
	var (
		ret0 [32]byte
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "testKeccak256",
		Params:      []interface{}{_data},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].([32]byte)

	return ret0, err
}

func (_VariAndFun *VariAndFun) TestMulMod(auth sdk.AccountBase, _x *big.Int, _y *big.Int, _k *big.Int) (*big.Int, error) {

	var (
		ret0 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "testMulMod",
		Params:      []interface{}{_x, _y, _k},
	}

	ret, err := _VariAndFun.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)

	return ret0, err
}

func (_VariAndFun *VariAndFun) TestMulModByHeight(auth sdk.AccountBase, _x *big.Int, _y *big.Int, _k *big.Int, height uint64) (*big.Int, error) {
	var (
		ret0 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "testMulMod",
		Params:      []interface{}{_x, _y, _k},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)

	return ret0, err
}

func (_VariAndFun *VariAndFun) TestRequire(auth sdk.AccountBase) (bool, error) {

	var (
		ret0 bool
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "testRequire",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(bool)

	return ret0, err
}

func (_VariAndFun *VariAndFun) TestRequireByHeight(auth sdk.AccountBase, height uint64) (bool, error) {
	var (
		ret0 bool
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "testRequire",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(bool)

	return ret0, err
}

func (_VariAndFun *VariAndFun) TestRevert(auth sdk.AccountBase) (bool, error) {

	var (
		ret0 bool
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "testRevert",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(bool)

	return ret0, err
}

func (_VariAndFun *VariAndFun) TestRevertByHeight(auth sdk.AccountBase, height uint64) (bool, error) {
	var (
		ret0 bool
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "testRevert",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(bool)

	return ret0, err
}

func (_VariAndFun *VariAndFun) TestRipemd160(auth sdk.AccountBase, _data []byte) ([20]byte, error) {

	var (
		ret0 [20]byte
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "testRipemd160",
		Params:      []interface{}{_data},
	}

	ret, err := _VariAndFun.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].([20]byte)

	return ret0, err
}

func (_VariAndFun *VariAndFun) TestRipemd160ByHeight(auth sdk.AccountBase, _data []byte, height uint64) ([20]byte, error) {
	var (
		ret0 [20]byte
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "testRipemd160",
		Params:      []interface{}{_data},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].([20]byte)

	return ret0, err
}

func (_VariAndFun *VariAndFun) TestSHA256(auth sdk.AccountBase, _data []byte) ([32]byte, error) {

	var (
		ret0 [32]byte
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "testSHA256",
		Params:      []interface{}{_data},
	}

	ret, err := _VariAndFun.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].([32]byte)

	return ret0, err
}

func (_VariAndFun *VariAndFun) TestSHA256ByHeight(auth sdk.AccountBase, _data []byte, height uint64) ([32]byte, error) {
	var (
		ret0 [32]byte
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "testSHA256",
		Params:      []interface{}{_data},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].([32]byte)

	return ret0, err
}

func (_VariAndFun *VariAndFun) Create2Deploy(auth sdk.AccountBase, _code []byte, _salt *big.Int) (txHash string, err error) {

	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "create2Deploy",
		Params:      []interface{}{_code, _salt},
	}

	return _VariAndFun.cli.ContractAsync(&m)
}

func (_VariAndFun *VariAndFun) Create2DeployByHeight(auth sdk.AccountBase, _code []byte, _salt *big.Int, height uint64) error {
	var ()
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFun.address.Hex(),
		ABI:         VariAndFunABI,
		Method:      "create2Deploy",
		Params:      []interface{}{_code, _salt},
	}

	ret, err := _VariAndFun.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})

	_ = arr
	return err
}
