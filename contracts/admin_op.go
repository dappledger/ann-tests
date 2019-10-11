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
const AdminOpBin = `608060405234801561001057600080fd5b506101ed806100206000396000f3fe6080604052600436106100245760003560e01c63ffffffff168063ba9c716e14610029575b600080fd5b34801561003557600080fd5b506100ef6004803603602081101561004c57600080fd5b810190808035906020019064010000000081111561006957600080fd5b82018360208201111561007b57600080fd5b8035906020019184600183028401116401000000008311171561009d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506100f1565b005b60603382604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b815260140182805190602001908083835b602083101515610162578051825260208201915060208101905060208303925061013d565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405290506000602082510190506000808284600060fe600019f18015156101bb57600080fd5b5050505056fea165627a7a72305820df2f1aae1d5d37323c73fbb446601f0a091768b31fd9fdf0d9b15639e963bd6e0029`

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
