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
	sdk "github.com/dappledger/ann-go-sdk"
	"github.com/dappledger/ann-go-sdk/common"
)

// VariAndFunOriginABI is the input ABI used to generate the binding from.
const VariAndFunOriginABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getOrigin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"testSelfdestruct\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// VariAndFunOriginBin is the compiled bytecode used for deploying new contracts.
const VariAndFunOriginBin = `608060405234801561001057600080fd5b506040516020806102b183398101806040528101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506101ed806100c46000396000f30060806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063df1f29ee14610051578063e8fd23cb146100a8575b600080fd5b34801561005d57600080fd5b506100666100bf565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156100b457600080fd5b506100bd610186565b005b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663df1f29ee6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561014657600080fd5b505af115801561015a573d6000803e3d6000fd5b505050506040513d602081101561017057600080fd5b8101908080519060200190929190505050905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff00a165627a7a7230582043a9116a1ee1e0630c42ac124d9603930452a1d96796282995365976d60ef82c0029`

// DeployVariAndFunOrigin deploys a new Ethereum contract, binding an instance of VariAndFunOrigin to it.
func DeployVariAndFunOrigin(goSDK *sdk.GoSDK, auth sdk.AccountBase, _addr common.Address) (common.Hash, *VariAndFunOrigin, error) {

	cc := sdk.ContractCreate{
		AccountBase: auth,
		Code:        VariAndFunOriginBin,
		ABI:         VariAndFunOriginABI,
		Params:      []interface{}{_addr},
	}
	ret, err := goSDK.ContractCreate(&cc)
	if err != nil {
		return common.Hash{}, nil, err
	}
	txHash := common.HexToHash(ret["tx"].(string))
	address := common.HexToAddress(ret["contract"].(string))
	return txHash, &VariAndFunOrigin{address: address, cli: goSDK}, nil
}

// VariAndFunOrigin is an auto generated Go binding around an Ethereum contract.
type VariAndFunOrigin struct {
	address common.Address
	cli     *sdk.GoSDK
}

// NewVariAndFunOrigin creates a new instance of VariAndFunOrigin, bound to a specific deployed contract.
func NewVariAndFunOrigin(goSdk *sdk.GoSDK, address common.Address) *VariAndFunOrigin {
	return &VariAndFunOrigin{address: address, cli: goSdk}
}

func (_VariAndFunOrigin *VariAndFunOrigin) GetAddress() common.Address {
	return _VariAndFunOrigin.address
}

func (_VariAndFunOrigin *VariAndFunOrigin) GetOrigin(auth sdk.AccountBase) (common.Address, error) {

	var (
		ret0 common.Address
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFunOrigin.address.Hex(),
		ABI:         VariAndFunOriginABI,
		Method:      "getOrigin",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFunOrigin.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(common.Address)

	return ret0, err
}

func (_VariAndFunOrigin *VariAndFunOrigin) GetOriginByHeight(auth sdk.AccountBase, height uint64) (common.Address, error) {
	var (
		ret0 common.Address
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFunOrigin.address.Hex(),
		ABI:         VariAndFunOriginABI,
		Method:      "getOrigin",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFunOrigin.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(common.Address)

	return ret0, err
}

func (_VariAndFunOrigin *VariAndFunOrigin) TestSelfdestruct(auth sdk.AccountBase) (txHash string, err error) {

	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFunOrigin.address.Hex(),
		ABI:         VariAndFunOriginABI,
		Method:      "testSelfdestruct",
		Params:      []interface{}{},
	}

	return _VariAndFunOrigin.cli.ContractAsync(&m)
}

func (_VariAndFunOrigin *VariAndFunOrigin) TestSelfdestructByHeight(auth sdk.AccountBase, height uint64) error {
	var ()
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFunOrigin.address.Hex(),
		ABI:         VariAndFunOriginABI,
		Method:      "testSelfdestruct",
		Params:      []interface{}{},
	}

	ret, err := _VariAndFunOrigin.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})

	_ = arr
	return err
}
