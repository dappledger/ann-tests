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

package payload

import (
	"math/big"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	sdk "github.com/dappledger/ann-go-sdk"
	"github.com/dappledger/ann-tests"
	"github.com/dappledger/ann-tests/cluster"
)

func TestTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tests Payload Suite")
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

var _ = Describe("Tests Payload", func() {

	Context("evm payload", func() {
		var (
			goSDK *sdk.GoSDK
			nonce uint64
			acc   sdk.Account
		)

		BeforeEach(func() {
			var err error
			node0, err := clusters.GetValidatorInfo(0)
			Expect(err).To(BeNil())
			goSDK = sdk.New(node0.RPC, sdk.CyrptoType(tests.TestCryptoType))
			acc, err = goSDK.AccountCreate()
			Expect(err).To(BeNil())
		})

		It("set and get", func() {
			nonce = 0
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: acc.Privkey}
			sendTx := &sdk.Tx{
				AccountBase: auth,
				To:          "",
				Payload:     "test",
				Value:       big.NewInt(0),
			}

			tx, err := goSDK.Transaction(sendTx)
			Expect(err).To(BeNil())
			_, err = tests.WaitTxMined(goSDK, tx)
			Expect(err).To(BeNil())

			//query
			time.Sleep(3 * time.Second)
			payload, err := goSDK.TransactionPayLoad(tx)
			Expect(err).To(BeNil())
			Expect(payload).To(Equal("test"))
		})
	})

	Context("kv", func() {
		var (
			goSDK *sdk.GoSDK
			nonce uint64
			acc   sdk.Account
		)

		BeforeEach(func() {
			var err error
			node0, err := clusters.GetValidatorInfo(0)
			Expect(err).To(BeNil())
			goSDK = sdk.New(node0.RPC, sdk.CyrptoType(tests.TestCryptoType))
			acc, err = goSDK.AccountCreate()
			Expect(err).To(BeNil())
		})

		It("set and get", func() {
			nonce = 0
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: acc.Privkey}
			kvTx := &sdk.KVTx{
				AccountBase: auth,
				Key:         []byte("key1"),
				Value:       []byte("value1"),
			}

			_, err := goSDK.Put(kvTx)
			Expect(err).To(BeNil())

			//query
			time.Sleep(2 * time.Second)
			value, err := goSDK.Get([]byte("key1"))
			Expect(err).To(BeNil())
			Expect(value).To(Equal([]byte("value1")))

			//check update
			nonce++
			kvTx.Nonce = nonce
			_, err = goSDK.Put(kvTx)
			Expect(err).To(BeNil())
			//Expect(err.Error()).Should(ContainSubstring("duplicate key"))
		})

		It("get with prefix", func() {
			nonce = 0
			auth := sdk.AccountBase{Nonce: nonce, PrivKey: acc.Privkey}
			kvTx := &sdk.KVTx{
				AccountBase: auth,
				Key:         []byte("key2"),
				Value:       []byte("value2"),
			}

			_, err := goSDK.Put(kvTx)
			Expect(err).To(BeNil())

			nonce++
			kvTx.Nonce = nonce
			kvTx.Key = []byte("key3")
			kvTx.Value = []byte("value3")
			_, err = goSDK.Put(kvTx)
			Expect(err).To(BeNil())

			time.Sleep(2 * time.Second)
			kvs, err := goSDK.GetWithPrefix([]byte("k"), []byte("key1"), 2)
			Expect(err).To(BeNil())
			Expect(len(kvs)).To(Equal(2))
			Expect(kvs[0]).To(Equal(&sdk.KVResult{Key: []byte("key2"), Value: []byte("value2")}))
			Expect(kvs[1]).To(Equal(&sdk.KVResult{Key: []byte("key3"), Value: []byte("value3")}))
		})
	})
})
