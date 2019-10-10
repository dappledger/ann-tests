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

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `6080604052620f424060005534801561001757600080fd5b5033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060005460026000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610d47806100d06000396000f3006080604052600436106100a4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde03146100a9578063095ea7b31461013957806318160ddd1461019e57806323b872dd146101c9578063313ce5671461024e57806370a082311461027f5780638da5cb5b146102d657806395d89b411461032d578063a9059cbb146103bd578063dd62ed3e14610422575b600080fd5b3480156100b557600080fd5b506100be610499565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100fe5780820151818401526020810190506100e3565b50505050905090810190601f16801561012b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561014557600080fd5b50610184600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506104d2565b604051808215151515815260200191505060405180910390f35b3480156101aa57600080fd5b506101b36105c4565b6040518082815260200191505060405180910390f35b3480156101d557600080fd5b50610234600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506105cd565b604051808215151515815260200191505060405180910390f35b34801561025a57600080fd5b5061026361098c565b604051808260ff1660ff16815260200191505060405180910390f35b34801561028b57600080fd5b506102c0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610991565b6040518082815260200191505060405180910390f35b3480156102e257600080fd5b506102eb6109da565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561033957600080fd5b50610342610a00565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610382578082015181840152602081019050610367565b50505050905090810190601f1680156103af5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156103c957600080fd5b50610408600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a39565b604051808215151515815260200191505060405180910390f35b34801561042e57600080fd5b50610483600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c5d565b6040518082815260200191505060405180910390f35b6040805190810160405280600581526020017f546f6b656e00000000000000000000000000000000000000000000000000000081525081565b600081600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b60008054905090565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561060a57600080fd5b600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054821115151561065857600080fd5b600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482111515156106e357600080fd5b61073582600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610ce490919063ffffffff16565b600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506107ca82600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610cfd90919063ffffffff16565b600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061089c82600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610ce490919063ffffffff16565b600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190509392505050565b601281565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040805190810160405280600181526020017f540000000000000000000000000000000000000000000000000000000000000081525081565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515610a7657600080fd5b600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548211151515610ac457600080fd5b610b1682600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610ce490919063ffffffff16565b600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610bab82600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610cfd90919063ffffffff16565b600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a36001905092915050565b6000600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000828211151515610cf257fe5b818303905092915050565b6000808284019050838110151515610d1157fe5b80915050929150505600a165627a7a723058206bd0dd6f826b805b164402ee52ae4162735853b57b27ad5f3b0caa6f955186a50029`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(goSDK *sdk.GoSDK, auth sdk.AccountBase) (common.Hash, *Token, error) {

	cc := sdk.ContractCreate{
		AccountBase: auth,
		Code:        TokenBin,
		ABI:         TokenABI,
		Params:      []interface{}{},
	}
	ret, err := goSDK.ContractCreate(&cc)
	if err != nil {
		return common.Hash{}, nil, err
	}
	txHash := common.HexToHash(ret["tx"].(string))
	address := common.HexToAddress(ret["contract"].(string))
	return txHash, &Token{address: address, cli: goSDK}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	address common.Address
	cli     *sdk.GoSDK
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(goSdk *sdk.GoSDK, address common.Address) *Token {
	return &Token{address: address, cli: goSdk}
}

func (_Token *Token) GetAddress() common.Address {
	return _Token.address
}

func (_Token *Token) Allowance(auth sdk.AccountBase, _owner common.Address, _spender common.Address) (*big.Int, error) {

	var (
		ret0 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _Token.address.Hex(),
		ABI:         TokenABI,
		Method:      "allowance",
		Params:      []interface{}{_owner, _spender},
	}

	ret, err := _Token.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)

	return ret0, err
}

func (_Token *Token) AllowanceByHeight(auth sdk.AccountBase, _owner common.Address, _spender common.Address, height uint64) (*big.Int, error) {
	var (
		ret0 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _Token.address.Hex(),
		ABI:         TokenABI,
		Method:      "allowance",
		Params:      []interface{}{_owner, _spender},
	}

	ret, err := _Token.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)

	return ret0, err
}

func (_Token *Token) BalanceOf(auth sdk.AccountBase, _owner common.Address) (*big.Int, error) {

	var (
		ret0 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _Token.address.Hex(),
		ABI:         TokenABI,
		Method:      "balanceOf",
		Params:      []interface{}{_owner},
	}

	ret, err := _Token.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)

	return ret0, err
}

func (_Token *Token) BalanceOfByHeight(auth sdk.AccountBase, _owner common.Address, height uint64) (*big.Int, error) {
	var (
		ret0 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _Token.address.Hex(),
		ABI:         TokenABI,
		Method:      "balanceOf",
		Params:      []interface{}{_owner},
	}

	ret, err := _Token.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)

	return ret0, err
}

func (_Token *Token) Decimals(auth sdk.AccountBase) (uint8, error) {

	var (
		ret0 uint8
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _Token.address.Hex(),
		ABI:         TokenABI,
		Method:      "decimals",
		Params:      []interface{}{},
	}

	ret, err := _Token.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(uint8)

	return ret0, err
}

func (_Token *Token) DecimalsByHeight(auth sdk.AccountBase, height uint64) (uint8, error) {
	var (
		ret0 uint8
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _Token.address.Hex(),
		ABI:         TokenABI,
		Method:      "decimals",
		Params:      []interface{}{},
	}

	ret, err := _Token.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(uint8)

	return ret0, err
}

func (_Token *Token) Name(auth sdk.AccountBase) (string, error) {

	var (
		ret0 string
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _Token.address.Hex(),
		ABI:         TokenABI,
		Method:      "name",
		Params:      []interface{}{},
	}

	ret, err := _Token.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(string)

	return ret0, err
}

func (_Token *Token) NameByHeight(auth sdk.AccountBase, height uint64) (string, error) {
	var (
		ret0 string
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _Token.address.Hex(),
		ABI:         TokenABI,
		Method:      "name",
		Params:      []interface{}{},
	}

	ret, err := _Token.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(string)

	return ret0, err
}

func (_Token *Token) Owner(auth sdk.AccountBase) (common.Address, error) {

	var (
		ret0 common.Address
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _Token.address.Hex(),
		ABI:         TokenABI,
		Method:      "owner",
		Params:      []interface{}{},
	}

	ret, err := _Token.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(common.Address)

	return ret0, err
}

func (_Token *Token) OwnerByHeight(auth sdk.AccountBase, height uint64) (common.Address, error) {
	var (
		ret0 common.Address
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _Token.address.Hex(),
		ABI:         TokenABI,
		Method:      "owner",
		Params:      []interface{}{},
	}

	ret, err := _Token.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(common.Address)

	return ret0, err
}

func (_Token *Token) Symbol(auth sdk.AccountBase) (string, error) {

	var (
		ret0 string
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _Token.address.Hex(),
		ABI:         TokenABI,
		Method:      "symbol",
		Params:      []interface{}{},
	}

	ret, err := _Token.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(string)

	return ret0, err
}

func (_Token *Token) SymbolByHeight(auth sdk.AccountBase, height uint64) (string, error) {
	var (
		ret0 string
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _Token.address.Hex(),
		ABI:         TokenABI,
		Method:      "symbol",
		Params:      []interface{}{},
	}

	ret, err := _Token.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(string)

	return ret0, err
}

func (_Token *Token) TotalSupply(auth sdk.AccountBase) (*big.Int, error) {

	var (
		ret0 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _Token.address.Hex(),
		ABI:         TokenABI,
		Method:      "totalSupply",
		Params:      []interface{}{},
	}

	ret, err := _Token.cli.ContractRead(&m)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)

	return ret0, err
}

func (_Token *Token) TotalSupplyByHeight(auth sdk.AccountBase, height uint64) (*big.Int, error) {
	var (
		ret0 *big.Int
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _Token.address.Hex(),
		ABI:         TokenABI,
		Method:      "totalSupply",
		Params:      []interface{}{},
	}

	ret, err := _Token.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(*big.Int)

	return ret0, err
}

func (_Token *Token) Approve(auth sdk.AccountBase, _spender common.Address, _value *big.Int) (txHash string, err error) {

	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _Token.address.Hex(),
		ABI:         TokenABI,
		Method:      "approve",
		Params:      []interface{}{_spender, _value},
	}

	return _Token.cli.ContractAsync(&m)
}

func (_Token *Token) ApproveByHeight(auth sdk.AccountBase, _spender common.Address, _value *big.Int, height uint64) (bool, error) {
	var (
		ret0 bool
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _Token.address.Hex(),
		ABI:         TokenABI,
		Method:      "approve",
		Params:      []interface{}{_spender, _value},
	}

	ret, err := _Token.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(bool)

	_ = arr
	return ret0, err
}

func (_Token *Token) Transfer(auth sdk.AccountBase, _to common.Address, _value *big.Int) (txHash string, err error) {

	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _Token.address.Hex(),
		ABI:         TokenABI,
		Method:      "transfer",
		Params:      []interface{}{_to, _value},
	}

	return _Token.cli.ContractAsync(&m)
}

func (_Token *Token) TransferByHeight(auth sdk.AccountBase, _to common.Address, _value *big.Int, height uint64) (bool, error) {
	var (
		ret0 bool
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _Token.address.Hex(),
		ABI:         TokenABI,
		Method:      "transfer",
		Params:      []interface{}{_to, _value},
	}

	ret, err := _Token.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(bool)

	_ = arr
	return ret0, err
}

func (_Token *Token) TransferFrom(auth sdk.AccountBase, _from common.Address, _to common.Address, _value *big.Int) (txHash string, err error) {

	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _Token.address.Hex(),
		ABI:         TokenABI,
		Method:      "transferFrom",
		Params:      []interface{}{_from, _to, _value},
	}

	return _Token.cli.ContractAsync(&m)
}

func (_Token *Token) TransferFromByHeight(auth sdk.AccountBase, _from common.Address, _to common.Address, _value *big.Int, height uint64) (bool, error) {
	var (
		ret0 bool
	)
	m := sdk.ContractMethod{
		AccountBase: auth,
		Contract:    _Token.address.Hex(),
		ABI:         TokenABI,
		Method:      "transferFrom",
		Params:      []interface{}{_from, _to, _value},
	}

	ret, err := _Token.cli.ContractReadByHeight(&m, height)
	arr := ret.([]interface{})
	ret0 = arr[0].(bool)

	_ = arr
	return ret0, err
}
