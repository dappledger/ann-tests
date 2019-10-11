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
const VariAndFunOriginABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getOrigin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"testSelfdestruct\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VariAndFunOriginBin is the compiled bytecode used for deploying new contracts.
const VariAndFunOriginBin = `608060405234801561001057600080fd5b506101b0806100206000396000f3fe60806040526004361061002f5760003560e01c63ffffffff168063cb1149d314610034578063e8fd23cb146100c5575b600080fd5b34801561004057600080fd5b506100836004803603602081101561005757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506100dc565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156100d157600080fd5b506100da61016b565b005b6000808290508073ffffffffffffffffffffffffffffffffffffffff1663df1f29ee6040518163ffffffff1660e01b815260040160206040518083038186803b15801561012857600080fd5b505afa15801561013c573d6000803e3d6000fd5b505050506040513d602081101561015257600080fd5b8101908080519060200190929190505050915050919050565b3373ffffffffffffffffffffffffffffffffffffffff16fffea165627a7a723058208b25bb190cb93764e77e1b2b3616cf63986a1977057bfc161f0154ec057d45cc0029`

// DeployVariAndFunOrigin deploys a new Ethereum contract, binding an instance of VariAndFunOrigin to it.
func DeployVariAndFunOrigin(goSDK *sdk.GoSDK, auth sdk.AccountBase) (common.Hash, *VariAndFunOrigin, error) {

	cc := sdk.ContractCreate{
		AccountBase: auth,
		Code:        VariAndFunOriginBin,
		ABI:         VariAndFunOriginABI,
		Params:      []interface{}{},
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

func (_VariAndFunOrigin *VariAndFunOrigin) GetOrigin(auth sdk.AccountBase, _addr common.Address) (common.Address, error) {

	var (
		ret0 common.Address
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFunOrigin.address.Hex(),
		ABI:         VariAndFunOriginABI,
		Method:      "getOrigin",
		Params:      []interface{}{_addr},
	}

	ret, err := _VariAndFunOrigin.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(common.Address)

	return ret0, err
}

func (_VariAndFunOrigin *VariAndFunOrigin) GetOriginByHeight(auth sdk.AccountBase, _addr common.Address, height uint64) (common.Address, error) {
	var (
		ret0 common.Address
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _VariAndFunOrigin.address.Hex(),
		ABI:         VariAndFunOriginABI,
		Method:      "getOrigin",
		Params:      []interface{}{_addr},
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
