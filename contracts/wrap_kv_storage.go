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

// WrapKvStorageABI is the input ABI used to generate the binding from.
const WrapKvStorageABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// WrapKvStorageBin is the compiled bytecode used for deploying new contracts.
const WrapKvStorageBin = `608060405234801561001057600080fd5b5060405160208061033c83398101806040528101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506102b9806100836000396000f30060806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063273f4940146100515780638eaa6ac01461008c575b600080fd5b34801561005d57600080fd5b5061008a6004803603810190808035600019169060200190929190803590602001909291905050506100d1565b005b34801561009857600080fd5b506100bb60048036038101908080356000191690602001909291905050506101b1565b6040518082815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663273f494083836040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180836000191660001916815260200182815260200192505050602060405180830381600087803b15801561017157600080fd5b505af1158015610185573d6000803e3d6000fd5b505050506040513d602081101561019b57600080fd5b8101908080519060200190929190505050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638eaa6ac0836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b15801561024b57600080fd5b505af115801561025f573d6000803e3d6000fd5b505050506040513d602081101561027557600080fd5b810190808051906020019092919050505090509190505600a165627a7a723058205ee12dc0f1245b5e9849200d8dcaa4fcd82bfce6cb2f66bca2764cdce47b4d090029`

// DeployWrapKvStorage deploys a new Ethereum contract, binding an instance of WrapKvStorage to it.
func DeployWrapKvStorage(goSDK *sdk.GoSDK, auth sdk.AccountBase, _addr common.Address) (common.Hash, *WrapKvStorage, error) {

	cc := sdk.ContractCreate{
		AccountBase: auth,
		Code:        WrapKvStorageBin,
		ABI:         WrapKvStorageABI,
		Params:      []interface{}{_addr},
	}
	ret, err := goSDK.ContractCreate(&cc)
	if err != nil {
		return common.Hash{}, nil, err
	}
	txHash := common.HexToHash(ret["tx"].(string))
	address := common.HexToAddress(ret["contract"].(string))
	return txHash, &WrapKvStorage{address: address, cli: goSDK}, nil
}

// WrapKvStorage is an auto generated Go binding around an Ethereum contract.
type WrapKvStorage struct {
	address common.Address
	cli     *sdk.GoSDK
}

// NewWrapKvStorage creates a new instance of WrapKvStorage, bound to a specific deployed contract.
func NewWrapKvStorage(goSdk *sdk.GoSDK, address common.Address) *WrapKvStorage {
	return &WrapKvStorage{address: address, cli: goSdk}
}

func (_WrapKvStorage *WrapKvStorage) GetAddress() common.Address {
	return _WrapKvStorage.address
}

func (_WrapKvStorage *WrapKvStorage) Get(auth sdk.AccountBase, key [32]byte) (*big.Int, error) {

	var (
		ret0 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _WrapKvStorage.address.Hex(),
		ABI:         WrapKvStorageABI,
		Method:      "get",
		Params:      []interface{}{key},
	}

	ret, err := _WrapKvStorage.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)

	return ret0, err
}

func (_WrapKvStorage *WrapKvStorage) GetByHeight(auth sdk.AccountBase, key [32]byte, height uint64) (*big.Int, error) {
	var (
		ret0 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _WrapKvStorage.address.Hex(),
		ABI:         WrapKvStorageABI,
		Method:      "get",
		Params:      []interface{}{key},
	}

	ret, err := _WrapKvStorage.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)

	return ret0, err
}

func (_WrapKvStorage *WrapKvStorage) Set(auth sdk.AccountBase, key [32]byte, value *big.Int) (txHash string, err error) {

	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _WrapKvStorage.address.Hex(),
		ABI:         WrapKvStorageABI,
		Method:      "set",
		Params:      []interface{}{key, value},
	}

	return _WrapKvStorage.cli.ContractAsync(&m)
}

func (_WrapKvStorage *WrapKvStorage) SetByHeight(auth sdk.AccountBase, key [32]byte, value *big.Int, height uint64) error {
	var ()
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _WrapKvStorage.address.Hex(),
		ABI:         WrapKvStorageABI,
		Method:      "set",
		Params:      []interface{}{key, value},
	}

	ret, err := _WrapKvStorage.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})

	_ = arr
	return err
}
