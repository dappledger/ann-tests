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
const TokenBin = `6080604052620f424060005534801561001757600080fd5b5033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060005460026000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610d70806100d06000396000f3fe6080604052600436106100875760003560e01c63ffffffff16806306fdde031461008c578063095ea7b31461011c57806318160ddd1461018f57806323b872dd146101ba578063313ce5671461024d57806370a082311461027e5780638da5cb5b146102e357806395d89b411461033a578063a9059cbb146103ca578063dd62ed3e1461043d575b600080fd5b34801561009857600080fd5b506100a16104c2565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100e15780820151818401526020810190506100c6565b50505050905090810190601f16801561010e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561012857600080fd5b506101756004803603604081101561013f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506104fb565b604051808215151515815260200191505060405180910390f35b34801561019b57600080fd5b506101a46105ed565b6040518082815260200191505060405180910390f35b3480156101c657600080fd5b50610233600480360360608110156101dd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506105f6565b604051808215151515815260200191505060405180910390f35b34801561025957600080fd5b506102626109b5565b604051808260ff1660ff16815260200191505060405180910390f35b34801561028a57600080fd5b506102cd600480360360208110156102a157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109ba565b6040518082815260200191505060405180910390f35b3480156102ef57600080fd5b506102f8610a03565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561034657600080fd5b5061034f610a29565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561038f578082015181840152602081019050610374565b50505050905090810190601f1680156103bc5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156103d657600080fd5b50610423600480360360408110156103ed57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a62565b604051808215151515815260200191505060405180910390f35b34801561044957600080fd5b506104ac6004803603604081101561046057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c86565b6040518082815260200191505060405180910390f35b6040805190810160405280600581526020017f546f6b656e00000000000000000000000000000000000000000000000000000081525081565b600081600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b60008054905090565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561063357600080fd5b600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054821115151561068157600080fd5b600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054821115151561070c57600080fd5b61075e82600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610d0d90919063ffffffff16565b600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506107f382600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610d2690919063ffffffff16565b600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506108c582600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610d0d90919063ffffffff16565b600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190509392505050565b601281565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040805190810160405280600181526020017f540000000000000000000000000000000000000000000000000000000000000081525081565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515610a9f57600080fd5b600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548211151515610aed57600080fd5b610b3f82600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610d0d90919063ffffffff16565b600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610bd482600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610d2690919063ffffffff16565b600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a36001905092915050565b6000600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000828211151515610d1b57fe5b818303905092915050565b6000808284019050838110151515610d3a57fe5b809150509291505056fea165627a7a723058205d88b75fae95eb7ee350a262b72b3afc55b7d82fa4c9730963bd13cf089862470029`

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
