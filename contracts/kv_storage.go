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

// KvStorageABI is the input ABI used to generate the binding from.
const KvStorageABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"keys\",\"type\":\"bytes32[]\"}],\"name\":\"batchGet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"keys\",\"type\":\"bytes32[]\"},{\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"batchSet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"calledCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// KvStorageBin is the compiled bytecode used for deploying new contracts.
const KvStorageBin = `608060405234801561001057600080fd5b50610442806100206000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063273f4940146100725780633b3e4b27146100c157806362b12feb1461017c5780638eaa6ac014610225578063b36d219d1461026a575b600080fd5b34801561007e57600080fd5b506100ab600480360381019080803560001916906020019092919080359060200190929190505050610295565b6040518082815260200191505060405180910390f35b3480156100cd57600080fd5b50610125600480360381019080803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091929192905050506102d3565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561016857808201518184015260208101905061014d565b505050509050019250505060405180910390f35b34801561018857600080fd5b50610223600480360381019080803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091929192908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919291929050505061037e565b005b34801561023157600080fd5b5061025460048036038101908080356000191690602001909291905050506103ec565b6040518082815260200191505060405180910390f35b34801561027657600080fd5b5061027f610410565b6040518082815260200191505060405180910390f35b600081600080856000191660001916815260200190815260200160002081905550600160008154809291906001019190505550600154905092915050565b606080600083516040519080825280602002602001820160405280156103085781602001602082028038833980820191505090505b509150600090505b835181101561037457600080858381518110151561032a57fe5b906020019060200201516000191660001916815260200190815260200160002054828281518110151561035957fe5b90602001906020020181815250508080600101915050610310565b8192505050919050565b60008090505b82518110156103e757818181518110151561039b57fe5b9060200190602002015160008085848151811015156103b657fe5b9060200190602002015160001916600019168152602001908152602001600020819055508080600101915050610384565b505050565b60008060008360001916600019168152602001908152602001600020549050919050565b600154815600a165627a7a72305820840a016c3f268aef37afb66a8ac0532c0fd7798752d5bd9ab8c71b40d652b9280029`

// DeployKvStorage deploys a new Ethereum contract, binding an instance of KvStorage to it.
func DeployKvStorage(goSDK *sdk.GoSDK, auth sdk.AccountBase) (common.Hash, *KvStorage, error) {

	cc := sdk.ContractCreate{
		AccountBase: auth,
		Code:        KvStorageBin,
		ABI:         KvStorageABI,
		Params:      []interface{}{},
	}
	ret, err := goSDK.ContractCreate(&cc)
	if err != nil {
		return common.Hash{}, nil, err
	}
	txHash := common.HexToHash(ret["tx"].(string))
	address := common.HexToAddress(ret["contract"].(string))
	return txHash, &KvStorage{address: address, cli: goSDK}, nil
}

// KvStorage is an auto generated Go binding around an Ethereum contract.
type KvStorage struct {
	address common.Address
	cli     *sdk.GoSDK
}

// NewKvStorage creates a new instance of KvStorage, bound to a specific deployed contract.
func NewKvStorage(goSdk *sdk.GoSDK, address common.Address) *KvStorage {
	return &KvStorage{address: address, cli: goSdk}
}

func (_KvStorage *KvStorage) GetAddress() common.Address {
	return _KvStorage.address
}

func (_KvStorage *KvStorage) BatchGet(auth sdk.AccountBase, keys [][32]byte) ([]*big.Int, error) {

	var (
		ret0 []*big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _KvStorage.address.Hex(),
		ABI:         KvStorageABI,
		Method:      "batchGet",
		Params:      []interface{}{keys},
	}

	ret, err := _KvStorage.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].([]*big.Int)

	return ret0, err
}

func (_KvStorage *KvStorage) BatchGetByHeight(auth sdk.AccountBase, keys [][32]byte, height uint64) ([]*big.Int, error) {
	var (
		ret0 []*big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _KvStorage.address.Hex(),
		ABI:         KvStorageABI,
		Method:      "batchGet",
		Params:      []interface{}{keys},
	}

	ret, err := _KvStorage.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].([]*big.Int)

	return ret0, err
}

func (_KvStorage *KvStorage) CalledCount(auth sdk.AccountBase) (*big.Int, error) {

	var (
		ret0 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _KvStorage.address.Hex(),
		ABI:         KvStorageABI,
		Method:      "calledCount",
		Params:      []interface{}{},
	}

	ret, err := _KvStorage.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)

	return ret0, err
}

func (_KvStorage *KvStorage) CalledCountByHeight(auth sdk.AccountBase, height uint64) (*big.Int, error) {
	var (
		ret0 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _KvStorage.address.Hex(),
		ABI:         KvStorageABI,
		Method:      "calledCount",
		Params:      []interface{}{},
	}

	ret, err := _KvStorage.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)

	return ret0, err
}

func (_KvStorage *KvStorage) Get(auth sdk.AccountBase, key [32]byte) (*big.Int, error) {

	var (
		ret0 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _KvStorage.address.Hex(),
		ABI:         KvStorageABI,
		Method:      "get",
		Params:      []interface{}{key},
	}

	ret, err := _KvStorage.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)

	return ret0, err
}

func (_KvStorage *KvStorage) GetByHeight(auth sdk.AccountBase, key [32]byte, height uint64) (*big.Int, error) {
	var (
		ret0 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _KvStorage.address.Hex(),
		ABI:         KvStorageABI,
		Method:      "get",
		Params:      []interface{}{key},
	}

	ret, err := _KvStorage.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)

	return ret0, err
}

func (_KvStorage *KvStorage) BatchSet(auth sdk.AccountBase, keys [][32]byte, values []*big.Int) (txHash string, err error) {

	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _KvStorage.address.Hex(),
		ABI:         KvStorageABI,
		Method:      "batchSet",
		Params:      []interface{}{keys, values},
	}

	return _KvStorage.cli.ContractAsync(&m)
}

func (_KvStorage *KvStorage) BatchSetByHeight(auth sdk.AccountBase, keys [][32]byte, values []*big.Int, height uint64) error {
	var ()
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _KvStorage.address.Hex(),
		ABI:         KvStorageABI,
		Method:      "batchSet",
		Params:      []interface{}{keys, values},
	}

	ret, err := _KvStorage.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})

	_ = arr
	return err
}

func (_KvStorage *KvStorage) Set(auth sdk.AccountBase, key [32]byte, value *big.Int) (txHash string, err error) {

	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _KvStorage.address.Hex(),
		ABI:         KvStorageABI,
		Method:      "set",
		Params:      []interface{}{key, value},
	}

	return _KvStorage.cli.ContractAsync(&m)
}

func (_KvStorage *KvStorage) SetByHeight(auth sdk.AccountBase, key [32]byte, value *big.Int, height uint64) (*big.Int, error) {
	var (
		ret0 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _KvStorage.address.Hex(),
		ABI:         KvStorageABI,
		Method:      "set",
		Params:      []interface{}{key, value},
	}

	ret, err := _KvStorage.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)

	_ = arr
	return ret0, err
}
