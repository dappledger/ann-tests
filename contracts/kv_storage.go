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
const KvStorageBin = `608060405234801561001057600080fd5b50610527806100206000396000f3fe6080604052600436106100505760003560e01c63ffffffff168063273f4940146100555780633b3e4b27146100ae57806362b12feb146101c85780638eaa6ac014610321578063b36d219d14610370575b600080fd5b34801561006157600080fd5b506100986004803603604081101561007857600080fd5b81019080803590602001909291908035906020019092919050505061039b565b6040518082815260200191505060405180910390f35b3480156100ba57600080fd5b50610171600480360360208110156100d157600080fd5b81019080803590602001906401000000008111156100ee57600080fd5b82018360208201111561010057600080fd5b8035906020019184602083028401116401000000008311171561012257600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192905050506103d1565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156101b4578082015181840152602081019050610199565b505050509050019250505060405180910390f35b3480156101d457600080fd5b5061031f600480360360408110156101eb57600080fd5b810190808035906020019064010000000081111561020857600080fd5b82018360208201111561021a57600080fd5b8035906020019184602083028401116401000000008311171561023c57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561029c57600080fd5b8201836020820111156102ae57600080fd5b803590602001918460208302840111640100000000831117156102d057600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290505050610473565b005b34801561032d57600080fd5b5061035a6004803603602081101561034457600080fd5b81019080803590602001909291905050506104d9565b6040518082815260200191505060405180910390f35b34801561037c57600080fd5b506103856104f5565b6040518082815260200191505060405180910390f35b60008160008085815260200190815260200160002081905550600160008154809291906001019190505550600154905092915050565b60608082516040519080825280602002602001820160405280156104045781602001602082028038833980820191505090505b50905060008090505b835181101561046957600080858381518110151561042757fe5b90602001906020020151815260200190815260200160002054828281518110151561044e57fe5b9060200190602002018181525050808060010191505061040d565b5080915050919050565b60008090505b82518110156104d457818181518110151561049057fe5b9060200190602002015160008085848151811015156104ab57fe5b906020019060200201518152602001908152602001600020819055508080600101915050610479565b505050565b6000806000838152602001908152602001600020549050919050565b6001548156fea165627a7a72305820dba9d950b241c27136f0e79b5c96e28ae78b9455519cae5a9a6a4f7a79cb178e0029`

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
