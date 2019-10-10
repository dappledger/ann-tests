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
const VariAndFunABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_x\",\"type\":\"uint256\"},{\"name\":\"_y\",\"type\":\"uint256\"},{\"name\":\"_k\",\"type\":\"uint256\"}],\"name\":\"testMulMod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operationHash\",\"type\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"testEcrecover\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_x\",\"type\":\"uint256\"},{\"name\":\"_y\",\"type\":\"uint256\"},{\"name\":\"_k\",\"type\":\"uint256\"}],\"name\":\"testAddMod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"testAssert\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"testRequire\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMsgSig\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"testSHA256\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMsgSender\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlockNum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlockTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlockHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"testRevert\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThis\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMsgData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"testKeccak256\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOrigin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"testRipemd160\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes20\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// VariAndFunBin is the compiled bytecode used for deploying new contracts.
const VariAndFunBin = `608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610c54806100606000396000f3006080604052600436106100fc576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063014066f31461010157806315334417146101565780631ebfbced1461020d5780632b813bc014610262578063357815c4146102915780635a8f6d71146102c057806363138d4f146103295780637a6ce2e1146103ae5780637f6c6f101461040557806387ceff09146104305780639663f88f14610469578063a26388bb1461049c578063b8368615146104cb578063c8e7ca2e14610522578063c8ec99f1146105b2578063df1f29ee14610637578063f6b0bbf71461068e578063f851a44014610729575b600080fd5b34801561010d57600080fd5b50610140600480360381019080803590602001909291908035906020019092919080359060200190929190505050610780565b6040518082815260200191505060405180910390f35b34801561016257600080fd5b506101cb6004803603810190808035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610798565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561021957600080fd5b5061024c600480360381019080803590602001909291908035906020019092919080359060200190929190505050610866565b6040518082815260200191505060405180910390f35b34801561026e57600080fd5b5061027761087e565b604051808215151515815260200191505060405180910390f35b34801561029d57600080fd5b506102a66108df565b604051808215151515815260200191505060405180910390f35b3480156102cc57600080fd5b506102d5610943565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b34801561033557600080fd5b50610390600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061096e565b60405180826000191660001916815260200191505060405180910390f35b3480156103ba57600080fd5b506103c3610a19565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561041157600080fd5b5061041a610a21565b6040518082815260200191505060405180910390f35b34801561043c57600080fd5b50610445610a29565b60405180848152602001838152602001828152602001935050505060405180910390f35b34801561047557600080fd5b5061047e610a3c565b60405180826000191660001916815260200191505060405180910390f35b3480156104a857600080fd5b506104b1610a4d565b604051808215151515815260200191505060405180910390f35b3480156104d757600080fd5b506104e0610ab1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561052e57600080fd5b50610537610ab9565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561057757808201518184015260208101905061055c565b50505050905090810190601f1680156105a45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156105be57600080fd5b50610619600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610af4565b60405180826000191660001916815260200191505060405180910390f35b34801561064357600080fd5b5061064c610b60565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561069a57600080fd5b506106f5600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610b68565b60405180826bffffffffffffffffffffffff19166bffffffffffffffffffffffff1916815260200191505060405180910390f35b34801561073557600080fd5b5061073e610c03565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60008180151561078c57fe5b83850990509392505050565b600080600080604185511415156107ae5761085d565b602085015192506040850151915060ff6041860151169050601b8160ff1610156107d957601b810190505b600186828585604051600081526020016040526040518085600019166000191681526020018460ff1660ff1681526020018360001916600019168152602001826000191660001916815260200194505050505060206040516020810390808403906000865af1158015610850573d6000803e3d6000fd5b5050506020604051035193505b50505092915050565b60008180151561087257fe5b83850890509392505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108d857fe5b6001905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561093c57600080fd5b6001905090565b600080357fffffffff0000000000000000000000000000000000000000000000000000000016905090565b60006002826040518082805190602001908083835b6020831015156109a85780518252602082019150602081019050602083039250610983565b6001836020036101000a0380198251168184511680821785525050505050509050019150506020604051808303816000865af11580156109ec573d6000803e3d6000fd5b5050506040513d6020811015610a0157600080fd5b81019080805190602001909291905050509050919050565b600033905090565b600043905090565b6000806000424243925092509250909192565b600080600143034090508091505090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610aaa57600080fd5b6001905090565b600030905090565b60606000368080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050905090565b6000816040518082805190602001908083835b602083101515610b2c5780518252602082019150602081019050602083039250610b07565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209050919050565b600032905090565b60006003826040518082805190602001908083835b602083101515610ba25780518252602082019150602081019050602083039250610b7d565b6001836020036101000a0380198251168184511680821785525050505050509050019150506020604051808303816000865af1158015610be6573d6000803e3d6000fd5b505050604051516c01000000000000000000000000029050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820428f02601d1e132b6a3ff8ae090715e4e9a903d3f6530ae82a580e6b160164c70029`

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
