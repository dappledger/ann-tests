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

// AdminOpABI is the input ABI used to generate the binding from.
const AdminOpABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"txdata\",\"type\":\"bytes\"}],\"name\":\"changenode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AdminOpBin is the compiled bytecode used for deploying new contracts.
const AdminOpBin = `608060405234801561001057600080fd5b50610210806100206000396000f3fe60806040526004361061003b576000357c010000000000000000000000000000000000000000000000000000000090048063ba9c716e14610040575b600080fd5b34801561004c57600080fd5b506101066004803603602081101561006357600080fd5b810190808035906020019064010000000081111561008057600080fd5b82018360208201111561009257600080fd5b803590602001918460018302840111640100000000831117156100b457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610108565b005b60603382604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c0100000000000000000000000002815260140182805190602001908083835b6020831015156101855780518252602082019150602081019050602083039250610160565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405290506000602082510190506000808284600060fe600019f18015156101de57600080fd5b5050505056fea165627a7a72305820d046b966dcd008f1acd08fc94962c671b9f01af09f4fde0b5c24988aa41153610029`

// DeployAdminOp deploys a new Ethereum contract, binding an instance of AdminOp to it.
func DeployAdminOp(goSDK *sdk.GoSDK, auth sdk.AccountBase) (common.Hash, *AdminOp, error) {

	cc := sdk.ContractCreate{
		AccountBase: auth,
		Code:        AdminOpBin,
		ABI:         AdminOpABI,
		Params:      []interface{}{},
	}
	ret, err := goSDK.ContractCreate(&cc)
	if err != nil {
		return common.Hash{}, nil, err
	}
	txHash := common.HexToHash(ret["tx"].(string))
	address := common.HexToAddress(ret["contract"].(string))
	return txHash, &AdminOp{address: address, cli: goSDK}, nil
}

// AdminOp is an auto generated Go binding around an Ethereum contract.
type AdminOp struct {
	address common.Address
	cli     *sdk.GoSDK
}

// NewAdminOp creates a new instance of AdminOp, bound to a specific deployed contract.
func NewAdminOp(goSdk *sdk.GoSDK, address common.Address) *AdminOp {
	return &AdminOp{address: address, cli: goSdk}
}

func (_AdminOp *AdminOp) GetAddress() common.Address {
	return _AdminOp.address
}

func (_AdminOp *AdminOp) Changenode(auth sdk.AccountBase, txdata []byte) (txHash string, err error) {

	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _AdminOp.address.Hex(),
		ABI:         AdminOpABI,
		Method:      "changenode",
		Params:      []interface{}{txdata},
	}

	return _AdminOp.cli.ContractAsync(&m)
}

func (_AdminOp *AdminOp) ChangenodeByHeight(auth sdk.AccountBase, txdata []byte, height uint64) error {
	var ()
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _AdminOp.address.Hex(),
		ABI:         AdminOpABI,
		Method:      "changenode",
		Params:      []interface{}{txdata},
	}

	ret, err := _AdminOp.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})

	_ = arr
	return err
}
