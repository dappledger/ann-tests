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

package contracts

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/ripemd160"

	"github.com/dappledger/ann-go-sdk"
	"github.com/dappledger/ann-go-sdk/abi"
	"github.com/dappledger/ann-go-sdk/common"
	"github.com/dappledger/ann-go-sdk/crypto"
	"github.com/dappledger/ann-go-sdk/types"
	"github.com/dappledger/ann-tests"
	"github.com/dappledger/ann-tests/cluster"
	"github.com/dappledger/ann-tests/contracts"
)

func TestTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tests Contracts Suite")
}

var clusters cluster.Cluster

var _ = BeforeSuite(func() {
	var err error
	clusters, err = tests.NewClusterFromEnv()
	Expect(err).To(BeNil())
	err = clusters.Up()
	Expect(err).To(BeNil())
})

var _ = AfterSuite(func() {
	if clusters == nil {
		return
	}
	if tests.TestKeepCluster {
		return
	}
	err := clusters.Down()
	Expect(err).To(BeNil())
})

var _ = Describe("Tests Contracts", func() {

	Context("KvStorage", func() {

		var (
			goSDK     *sdk.GoSDK
			nonce     uint64
			acc       sdk.Account
			kvStorage *contracts.KvStorage
		)

		// deploy kvstorage contract
		BeforeEach(func() {
			var err error
			node0, err := clusters.GetValidatorInfo(0)
			Expect(err).To(BeNil())

			goSDK = sdk.New(node0.RPC, sdk.CyrptoType(tests.TestCryptoType))
			acc, err = goSDK.AccountCreate()
			Expect(err).To(BeNil())

			nonce = 0
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: acc.Privkey}
			var tx common.Hash
			tx, kvStorage, err = contracts.DeployKvStorage(goSDK, auth)
			Expect(err).To(BeNil())
			_, err = tests.WaitTxMined(goSDK, tx.Hex())
			Expect(err).To(BeNil())
		})

		It("invoke set method", func() {
			nonce++
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: acc.Privkey}

			key := crypto.Keccak256Hash([]byte("test key"))
			value := time.Now().UnixNano()
			bvalue := big.NewInt(int64(value))

			tx, err := kvStorage.Set(auth, key, bvalue)

			Expect(err).To(BeNil())
			_, err = tests.WaitTxMined(goSDK, tx)
			Expect(err).To(BeNil())

			time.Sleep(5 * time.Second)

			actual, err := kvStorage.Get(auth, key)
			Expect(err).To(BeNil())
			Expect(actual).To(Equal(bvalue))
		})

		It("call set method by height", func() {

			num := 4
			trans := make([]*sdk.RPCTransaction, 0, num)
			for i := 0; i < num; i++ {
				nonce++
				auth := sdk.AccountBase{Nonce: nonce, PrivKey: acc.Privkey}

				key := crypto.Keccak256Hash([]byte(fmt.Sprintf("test key-%d", i)))
				value := time.Now().UnixNano()
				bvalue := big.NewInt(int64(value))

				tx, err := kvStorage.Set(auth, key, bvalue)

				Expect(err).To(BeNil())
				_, err = tests.WaitTxMined(goSDK, tx)
				Expect(err).To(BeNil())

				actual, err := kvStorage.Get(auth, key)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(bvalue))

				t, err := goSDK.GetTransactionByHash(tx)
				Expect(err).To(BeNil())
				trans = append(trans, t)
			}

			// exclude last transaction because of it hasn't be mined
			for i, t := range trans[:num-1] {
				auth := sdk.AccountBase{Nonce: nonce, PrivKey: acc.Privkey}

				key := crypto.Keccak256Hash([]byte(fmt.Sprintf("test key-%d", i)))
				value := time.Now().UnixNano()
				bvalue := big.NewInt(int64(value))

				calledCount, err := kvStorage.SetByHeight(auth, key, bvalue, t.BlockHeight-1)
				Expect(err).To(BeNil())
				Expect(calledCount.Int64()).To(Equal(int64(i + 1)))
			}
		})

		It("invoke batch set method", func() {

			nonce++
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: acc.Privkey}

			batchSize := 10
			keys := make([][32]byte, 0, batchSize)
			expected := make([]*big.Int, 0, batchSize)
			for i := 0; i < batchSize; i++ {
				keys = append(keys, crypto.Keccak256Hash([]byte(fmt.Sprintf("key-%d", i))))
				expected = append(expected, big.NewInt(int64(i)))
			}

			tx, err := kvStorage.BatchSet(auth, keys, expected)
			Expect(err).To(BeNil(), "kvStorage.BatchSet")
			_, err = tests.WaitTxMined(goSDK, tx)
			Expect(err).To(BeNil())

			actual, err := kvStorage.BatchGet(auth, keys)
			Expect(err).To(BeNil())

			for i, v := range expected {
				Expect(v.Int64()).To(Equal(actual[i].Int64()))
			}
		})

		It("get transaction by hash", func() {
			txNum := 10
			allNonce := make([]uint64, 0, txNum)
			txs := make([]string, 0, txNum)
			for i := 0; i < txNum; i++ {
				nonce++
				allNonce = append(allNonce, nonce)
				key := crypto.Keccak256Hash([]byte(fmt.Sprintf("test key%d", i)))
				value := time.Now().UnixNano()
				bvalue := big.NewInt(int64(value))
				auth := sdk.AccountBase{Nonce: nonce, PrivKey: acc.Privkey}
				tx, err := kvStorage.Set(auth, key, bvalue)
				Expect(err).To(BeNil())
				txs = append(txs, tx)
			}

			for _, tx := range txs {
				_, err := tests.WaitTxMined(goSDK, tx)
				Expect(err).To(BeNil())
			}

			for i, tx := range txs {
				t, err := goSDK.GetTransactionByHash(tx)
				Expect(err).To(BeNil())
				accAddrStr := fmt.Sprintf("%X", common.FromHex(acc.Address))

				Expect(err).To(BeNil())
				fromAddrStr := fmt.Sprintf("%X", t.From)
				Expect(fromAddrStr).To(Equal(accAddrStr))
				Expect(t.To.String()).To(Equal(kvStorage.GetAddress().Hex()))
				Expect(t.Nonce).To(Equal(allNonce[i]))
				b, err := goSDK.Block(int64(t.BlockHeight))
				Expect(err).To(BeNil())
				Expect(t.BlockHash).To(Equal(b.BlockMeta.Hash))
				txStr := b.Block.Data.Txs[int(t.TransactionIndex)]
				Expect(fmt.Sprintf("0x%x", types.Tx(txStr).Hash())).To(Equal(tx))
			}
		})
	})

	Context("WrapKvStorage", func() {

		var (
			goSDK         *sdk.GoSDK
			nonce         uint64
			acc           sdk.Account
			wrapKvStorage *contracts.WrapKvStorage
			kvStorage     *contracts.KvStorage
		)

		// deploy kvstorage contract
		BeforeEach(func() {
			var err error
			node0, err := clusters.GetValidatorInfo(0)
			Expect(err).To(BeNil())

			goSDK = sdk.New(node0.RPC, sdk.CyrptoType(tests.TestCryptoType))
			acc, err = goSDK.AccountCreate()
			Expect(err).To(BeNil())

			nonce = 0
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: acc.Privkey}
			var tx common.Hash
			tx, kvStorage, err = contracts.DeployKvStorage(goSDK, auth)
			Expect(err).To(BeNil())
			_, err = tests.WaitTxMined(goSDK, tx.Hex())
			Expect(err).To(BeNil())

			nonce++
			auth.Nonce = nonce
			tx, wrapKvStorage, err = contracts.DeployWrapKvStorage(goSDK, auth, kvStorage.GetAddress())
			Expect(err).To(BeNil())
			_, err = tests.WaitTxMined(goSDK, tx.Hex())
			Expect(err).To(BeNil())
		})

		It("call set via wrapped contracts", func() {
			nonce++
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: acc.Privkey}

			expected := time.Now().UnixNano()
			key := crypto.Keccak256Hash([]byte("test key"))
			tx, err := wrapKvStorage.Set(auth, key, big.NewInt(expected))
			Expect(err).To(BeNil())
			_, err = tests.WaitTxMined(goSDK, tx)
			Expect(err).To(BeNil())

			actual, err := wrapKvStorage.Get(auth, key)
			Expect(err).To(BeNil())
			Expect(actual.Int64()).To(Equal(expected))

			actual2, err := kvStorage.Get(auth, key)
			Expect(err).To(BeNil())
			Expect(actual2.Int64()).To(Equal(expected))
		})
	})

	Context("ERC20", func() {

		var (
			goSDK *sdk.GoSDK
			token *contracts.Token

			owner        sdk.Account
			ownerNonce   uint64
			ownerAddress common.Address

			receiver        sdk.Account
			receiverNonce   uint64
			receiverAddress common.Address
		)

		const (
			TotalSupply = int64(1000000)
		)

		// issue token
		BeforeEach(func() {
			var err error
			node0, err := clusters.GetValidatorInfo(0)
			Expect(err).To(BeNil())

			goSDK = sdk.New(node0.RPC, sdk.CyrptoType(tests.TestCryptoType))
			owner, err = goSDK.AccountCreate()
			Expect(err).To(BeNil())
			ownerNonce = 0
			ownerAddress = common.HexToAddress(owner.Address)

			receiver, err = goSDK.AccountCreate()
			Expect(err).To(BeNil())
			receiverNonce = 0
			receiverAddress = common.HexToAddress(receiver.Address)

			var txHash common.Hash
			txHash, token, err = contracts.DeployToken(goSDK, sdk.AccountBase{Nonce: ownerNonce, PrivKey: owner.Privkey})
			_, err = tests.WaitTxMined(goSDK, txHash.Hex())
			ownerNonce++
			Expect(err).To(BeNil())
		})

		It("token info", func() {

			auth := sdk.AccountBase{Nonce: ownerNonce, PrivKey: owner.Privkey}

			totalSupply, err := token.TotalSupply(auth)
			Expect(err).To(BeNil())
			Expect(totalSupply.Int64()).To(Equal(TotalSupply))

			symbol, err := token.Symbol(auth)
			Expect(err).To(BeNil())
			Expect(symbol).To(Equal("T"))

			name, err := token.Name(auth)
			Expect(err).To(BeNil())
			Expect(name).To(Equal("Token"))

			decimals, err := token.Decimals(auth)
			Expect(err).To(BeNil())
			Expect(decimals).To(Equal(uint8(18)))

			balanceOfOwner, err := token.BalanceOf(auth, ownerAddress)
			Expect(err).To(BeNil())
			Expect(balanceOfOwner.Int64()).To(Equal(TotalSupply))
		})

		It("transfer token", func() {

			ownerAuth := sdk.AccountBase{Nonce: ownerNonce, PrivKey: owner.Privkey}

			const TransferValue = int64(100)
			value := big.NewInt(TransferValue)

			tx, err := token.Transfer(ownerAuth, receiverAddress, value)
			Expect(err).To(BeNil())
			_, err = tests.WaitTxMined(goSDK, tx)
			ownerNonce++
			Expect(err).To(BeNil())

			ownerAuth.Nonce = ownerNonce
			balanceOfReceiver, err := token.BalanceOf(ownerAuth, receiverAddress)
			Expect(err).To(BeNil())
			Expect(balanceOfReceiver.Int64()).To(Equal(TransferValue))

			balanceOfOwner, err := token.BalanceOf(ownerAuth, ownerAddress)
			Expect(err).To(BeNil())
			Expect(balanceOfOwner.Int64()).To(Equal(TotalSupply - TransferValue))
		})

		It("transfer token after approve", func() {

			ownerAuth := sdk.AccountBase{Nonce: ownerNonce, PrivKey: owner.Privkey}

			const TransferValue = int64(100)
			value := big.NewInt(TransferValue)

			// approve
			tx, err := token.Approve(ownerAuth, receiverAddress, value)
			Expect(err).To(BeNil())
			_, err = tests.WaitTxMined(goSDK, tx)
			Expect(err).To(BeNil())
			ownerNonce++

			// check allowance
			ownerAuth = sdk.AccountBase{Nonce: ownerNonce, PrivKey: owner.Privkey}
			allowance, err := token.Allowance(ownerAuth, ownerAddress, receiverAddress)
			Expect(err).To(BeNil())
			Expect(allowance.Int64()).To(Equal(TransferValue))

			// transferFrom
			receiverAuth := sdk.AccountBase{Nonce: receiverNonce, PrivKey: receiver.Privkey}
			to, _ := goSDK.AccountCreate()
			toAddress := common.HexToAddress(to.Address)
			tx, err = token.TransferFrom(receiverAuth, ownerAddress, toAddress, value)
			Expect(err).To(BeNil())
			_, err = tests.WaitTxMined(goSDK, tx)
			Expect(err).To(BeNil())
			receiverNonce++

			// check balance
			allowance, err = token.Allowance(ownerAuth, ownerAddress, receiverAddress)
			Expect(err).To(BeNil())
			Expect(allowance.Int64()).To(Equal(int64(0)))

			balance, err := token.BalanceOf(ownerAuth, ownerAddress)
			Expect(err).To(BeNil())
			Expect(balance.Int64()).To(Equal(TotalSupply - TransferValue))

			balance, err = token.BalanceOf(ownerAuth, toAddress)
			Expect(err).To(BeNil())
			Expect(balance.Int64()).To(Equal(TransferValue))
		})
	})

	Context("VariAndFun", func() {

		var (
			goSDK *sdk.GoSDK
			nonce uint64
			admin sdk.Account
			user  sdk.Account

			variAndFun *contracts.VariAndFun
			dataBytes  = common.Hex2Bytes("test")
		)

		const (
			contractABI = contracts.VariAndFunABI
		)

		// deploy variAndFun contract
		BeforeEach(func() {
			var err error
			node0, err := clusters.GetValidatorInfo(0)
			Expect(err).To(BeNil())

			goSDK = sdk.New(node0.RPC, sdk.CyrptoType(tests.TestCryptoType))
			admin, err = goSDK.AccountCreate()
			Expect(err).To(BeNil())
			user, err = goSDK.AccountCreate()
			Expect(err).To(BeNil())

			nonce = 0
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: admin.Privkey}
			var tx common.Hash
			tx, variAndFun, err = contracts.DeployVariAndFun(goSDK, auth)
			Expect(err).To(BeNil())
			_, err = tests.WaitTxMined(goSDK, tx.Hex())
			Expect(err).To(BeNil())
		})

		It("block height", func() {
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: admin.Privkey}

			blockNum, err := variAndFun.GetBlockNum(auth)
			Expect(err).To(BeNil())
			expectHeight, err := goSDK.LastHeight()
			Expect(err).To(BeNil())
			Expect(blockNum.Int64()).To(Equal(expectHeight))
		})

		It("block hash", func() {
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: admin.Privkey}

			blockHash, err := variAndFun.GetBlockHash(auth)
			Expect(err).To(BeNil())

			lastHeight, err := goSDK.LastHeight()
			Expect(err).To(BeNil())
			lastBlock, err := goSDK.Block(lastHeight - 1)
			Expect(err).To(BeNil())
			expectBlockHash := lastBlock.BlockMeta.Hash
			Expect(blockHash[12:32]).To(Equal(expectBlockHash))
		})

		It("block time and now", func() {
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: admin.Privkey}

			blockTime, nowTime, blockHeight, err := variAndFun.GetBlockTime(auth)
			Expect(err).To(BeNil())
			Expect(blockTime).To(Equal(nowTime))

			blockInfo, err := goSDK.Block(blockHeight.Int64())
			Expect(err).To(BeNil())
			expectTime := blockInfo.Block.Header.Time.Unix()
			Expect(blockTime.Int64()).To(Equal(expectTime))
		})

		It("msg data", func() {
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: admin.Privkey}

			msgData, err := variAndFun.GetMsgData(auth)
			Expect(err).To(BeNil())

			abiJson, err := abi.JSON(strings.NewReader(contractABI))
			Expect(err).To(BeNil())
			expectData, err := sdk.PackCalldata(&abiJson, "getMsgData", []interface{}{})
			Expect(err).To(BeNil())
			Expect(msgData).To(Equal(expectData))
		})

		It("msg sender", func() {
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: admin.Privkey}

			msgSender, err := variAndFun.GetMsgSender(auth)
			Expect(err).To(BeNil())

			msgSenderStr := strings.ToUpper(msgSender.String())
			if strings.Index(msgSenderStr, "0X") == 0 {
				msgSenderStr = msgSenderStr[2:]
			}

			expectSender := strings.ToUpper(admin.Address)
			if strings.Index(expectSender, "0X") == 0 {
				expectSender = expectSender[2:]
			}

			Expect(msgSenderStr).To(Equal(expectSender))
		})

		It("msg sig", func() {
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: admin.Privkey}

			msgSig, err := variAndFun.GetMsgSig(auth)
			Expect(err).To(BeNil())

			abiJson, err := abi.JSON(strings.NewReader(contractABI))
			Expect(err).To(BeNil())
			expectSig, err := sdk.PackCalldata(&abiJson, "getMsgSig", []interface{}{})
			Expect(err).To(BeNil())
			Expect(msgSig[:]).To(Equal(expectSig))
		})

		It("assert", func() {
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: admin.Privkey}
			userAuth := sdk.AccountBase{Nonce: nonce, PrivKey: user.Privkey}

			adminResp, err := variAndFun.TestAssert(auth)
			Expect(err).To(BeNil())
			Expect(adminResp).To(Equal(true))

			userResp, err := variAndFun.TestAssert((userAuth))
			Expect(userResp).To(Equal(false))
		})

		It("require", func() {
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: admin.Privkey}
			userAuth := sdk.AccountBase{Nonce: nonce, PrivKey: user.Privkey}

			adminResp, err := variAndFun.TestRequire(auth)
			Expect(err).To(BeNil())
			Expect(adminResp).To(Equal(true))

			userResp, err := variAndFun.TestRequire((userAuth))
			Expect(userResp).To(Equal(false))
		})

		It("revert", func() {
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: admin.Privkey}
			userAuth := sdk.AccountBase{Nonce: nonce, PrivKey: user.Privkey}

			adminResp, err := variAndFun.TestRevert(auth)
			Expect(err).To(BeNil())
			Expect(adminResp).To(Equal(true))

			userResp, err := variAndFun.TestRevert((userAuth))
			Expect(userResp).To(Equal(false))
		})

		It("keccak256", func() {
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: admin.Privkey}

			resp, err := variAndFun.TestKeccak256(auth, dataBytes)
			Expect(err).To(BeNil())
			expectResp := crypto.Keccak256(dataBytes)
			Expect(resp[:]).To(Equal(expectResp))
		})

		It("sha256", func() {
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: admin.Privkey}

			resp, err := variAndFun.TestSHA256(auth, dataBytes)
			Expect(err).To(BeNil())
			expectResp := sha256.Sum256(dataBytes)
			Expect(resp).To(Equal(expectResp))
		})

		It("ripemd160", func() {
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: admin.Privkey}

			resp, err := variAndFun.TestRipemd160(auth, dataBytes)
			Expect(err).To(BeNil())

			ripemd := ripemd160.New()
			ripemd.Write(dataBytes)
			expectResp := common.LeftPadBytes(ripemd.Sum(nil), 20)
			Expect(resp[:]).To(Equal(expectResp))
		})

		It("ecrecover", func() {
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: admin.Privkey}

			priv, err := crypto.ToECDSA(common.Hex2Bytes(admin.Privkey))
			Expect(err).To(BeNil())
			opHash := crypto.Keccak256(dataBytes)
			sig, err := crypto.Sign(opHash, priv)
			Expect(err).To(BeNil())

			opHashBytes32 := [32]byte{}
			copy(opHashBytes32[:], opHash)
			resp, err := variAndFun.TestEcrecover(auth, opHashBytes32, sig)
			Expect(err).To(BeNil())
			Expect(resp.String()).To(Equal(admin.Address))
		})

		It("addmod", func() {
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: admin.Privkey}

			resp, err := variAndFun.TestAddMod(auth, big.NewInt(3), big.NewInt(6), big.NewInt(4))
			Expect(err).To(BeNil())
			Expect(resp).To(Equal(big.NewInt(1)))
		})

		It("mulmod", func() {
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: admin.Privkey}

			resp, err := variAndFun.TestMulMod(auth, big.NewInt(3), big.NewInt(6), big.NewInt(4))
			Expect(err).To(BeNil())
			Expect(resp).To(Equal(big.NewInt(2)))
		})

		It("this", func() {
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: admin.Privkey}

			addr, err := variAndFun.GetThis(auth)
			Expect(err).To(BeNil())
			Expect(addr).To(Equal(variAndFun.GetAddress()))
		})
	})

	Context("VariAndFunOrigin", func() {

		var (
			goSDK            *sdk.GoSDK
			nonce            uint64
			acc              sdk.Account
			variAndFunOrigin *contracts.VariAndFunOrigin
			variAndFun       *contracts.VariAndFun
		)

		// deploy variAndFun and variAndFunOrigin contract
		BeforeEach(func() {
			var err error
			node0, err := clusters.GetValidatorInfo(0)
			Expect(err).To(BeNil())

			goSDK = sdk.New(node0.RPC, sdk.CyrptoType(tests.TestCryptoType))
			acc, err = goSDK.AccountCreate()
			Expect(err).To(BeNil())

			nonce = 0
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: acc.Privkey}
			var tx common.Hash
			tx, variAndFun, err = contracts.DeployVariAndFun(goSDK, auth)
			Expect(err).To(BeNil())
			_, err = tests.WaitTxMined(goSDK, tx.Hex())
			Expect(err).To(BeNil())

			nonce++
			auth.Nonce = nonce
			tx, variAndFunOrigin, err = contracts.DeployVariAndFunOrigin(goSDK, auth, variAndFun.GetAddress())
			Expect(err).To(BeNil())
			_, err = tests.WaitTxMined(goSDK, tx.Hex())
			Expect(err).To(BeNil())
		})

		It("tx origin", func() {
			nonce++
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: acc.Privkey}

			addr, err := variAndFunOrigin.GetOrigin(auth)
			addrStr := strings.ToUpper(addr.String())
			if strings.Index(addrStr, "0X") == 0 {
				addrStr = addrStr[2:]
			}

			expectAddr := strings.ToUpper(acc.Address)
			if strings.Index(expectAddr, "0X") == 0 {
				expectAddr = expectAddr[2:]
			}

			Expect(err).To(BeNil())
			Expect(addrStr).To(Equal(expectAddr))
		})

		It("selfdestruct", func() {
			nonce++
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: acc.Privkey}

			_, err := variAndFunOrigin.TestSelfdestruct(auth)
			Expect(err).To(BeNil())

			time.Sleep(2 * time.Second)
			addr, err := variAndFunOrigin.GetOrigin(auth)
			Expect(err).To(BeNil())
			Expect(addr).To(Equal(common.Address{}))
		})
	})
})
