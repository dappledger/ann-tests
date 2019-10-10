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

package admin

import (
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	sdk "github.com/dappledger/ann-go-sdk"
	crypto "github.com/dappledger/ann-go-sdk/crypto"
	"github.com/dappledger/ann-go-sdk/types"
	"github.com/dappledger/ann-tests"
	"github.com/dappledger/ann-tests/cluster"
)

type ValidatorInfo struct {
	RPC     string
	Address string
	PrivKey string
	PubKey  string
	Role    NodeRole
	Running bool
}

type NodeRole int

const (
	NodeRoleNone NodeRole = iota
	NodeRolePeer
	NodeRoleValidator
)

const (
	allSign = -1
)

func (sv *ValidatorInfo) getPubHex() string {
	if len(sv.PubKey) == 0 {
		data, err := hex.DecodeString(sv.PrivKey)
		Expect(err).To(BeNil())
		privKey := crypto.SetNodePrivKey(tests.TestCryptoType, data)
		sv.PubKey = privKey.PubKey().KeyString()
	}

	return sv.PubKey
}

var (
	clusters             cluster.Cluster
	actualValidatorsInfo []*ValidatorInfo
)

func TestTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AdminOP Suite")
}

var _ = BeforeSuite(func() {
	var err error
	clusters, err = tests.NewClusterFromEnv()
	Expect(err).To(BeNil())
	err = clusters.Up()
	Expect(err).To(BeNil())
	time.Sleep(1 * time.Second)
	//
	for i := 0; i < clusters.ValidatorNum(); i++ {
		node, err := clusters.GetValidatorInfo(i)
		Expect(err).To(BeNil())
		rt := &ValidatorInfo{
			RPC:     node.RPC,
			Address: node.Address,
			PrivKey: node.PrivKey,
			Role:    NodeRoleValidator,
			Running: true,
		}
		rt.getPubHex()
		actualValidatorsInfo = append(actualValidatorsInfo, rt)
	}
	goSDK := ExpectNewSDK(actualValidatorsInfo[0].RPC)
	ncount := 0
	for {
		ok := ExpectHealth(goSDK)
		if ok {
			break
		}
		time.Sleep(1 * time.Second)
		ncount++
		Expect(ncount < 20).To(BeTrue())
	}
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

func lastNode() int {
	return len(actualValidatorsInfo) - 1
}

var _ = Describe("Tests AdminOP", func() {
	BeforeEach(beforeAfterEachCheck)
	AfterEach(beforeAfterEachCheck)
	Context(fmt.Sprintf("With %d validators", tests.TestValidatorNum), func() {
		It("test validator update", updateTest)

		It("test validator add/remove", addRemoveTest)

		It("test golang-sdk validator add/remove", addRemoveGOSDKTest)

		It("test validator replay", replayTest)

		It("test validator travel(multi tx)", multiTXTest)
	})
})

func updateTest() {
	node0 := actualValidatorsInfo[0]
	goSDK := ExpectNewSDK(actualValidatorsInfo[lastNode()].RPC)
	checkValidators(goSDK)
	ExpectUpdateNode(node0, 0, goSDK, allSign)
	node0.Role = NodeRolePeer
	checkValidators(goSDK)
	ExpectUpdateNode(node0, 100, goSDK, allSign)
	node0.Role = NodeRoleValidator
	checkValidators(goSDK)
}

func addRemoveGOSDKTest() {
	node0 := actualValidatorsInfo[0]
	canodes := actualValidatorsInfo[1:]
	client := ExpectNewSDK(actualValidatorsInfo[3].RPC)
	acc, err := client.AccountCreate()
	Expect(err).To(BeNil())
	nonce, err := client.Nonce(acc.Address)
	Expect(err).To(BeNil())

	accbase := sdk.AccountBase{
		PrivKey: acc.Privkey,
		Nonce:   nonce,
	}

	ExpectSDKChangeNode(node0, canodes, 0, client, types.ValidatorCmdRemoveNode, accbase)
	node0.Role = NodeRoleNone
	checkValidators(client)
	accbase.Nonce++

	ExpectSDKChangeNode(node0, canodes, 0, client, types.ValidatorCmdAddPeer, accbase)
	node0.Role = NodeRolePeer
	checkValidators(client)
	accbase.Nonce++

	ExpectSDKChangeNode(node0, canodes, 100, client, types.ValidatorCmdUpdateNode, accbase)
	node0.Role = NodeRoleValidator
	waitSyncBlock(client)
	checkValidators(client)
	accbase.Nonce++
}

func addRemoveTest() {
	node0 := actualValidatorsInfo[0]
	goSDK := ExpectNewSDK(actualValidatorsInfo[3].RPC)

	ExpectRemoveNode(node0, goSDK, 1)
	checkValidators(goSDK)
	ExpectRemoveNode(node0, goSDK, allSign)
	node0.Role = NodeRoleNone
	checkValidators(goSDK)

	ExpectAddPeer(node0, goSDK, 1)
	checkValidators(goSDK)
	ExpectAddPeer(node0, goSDK, allSign)
	node0.Role = NodeRolePeer
	checkValidators(goSDK)

	ExpectUpdateNode(node0, 100, goSDK, 1)
	checkValidators(goSDK)
	ExpectUpdateNode(node0, 100, goSDK, allSign)
	node0.Role = NodeRoleValidator
	checkValidators(goSDK)
	waitSyncBlock(goSDK)
}

func multiTXTest() {
	nd := actualValidatorsInfo[0]
	goSDK := ExpectNewSDK(actualValidatorsInfo[3].RPC)
	var buf = make(map[string]bool, 10)
	au := NewAdminOPUser(goSDK)
	for i := 0; i < 2; i++ {
		err := au.MakeUpdateNodeMsg(nd, 0, allSign)
		Expect(err).To(BeNil())
		txHash, err := asyncChangeValidator(au, goSDK)
		Expect(err).To(BeNil())
		buf[txHash] = true
		au.Nonce++

		err = au.MakeUpdateNodeMsg(nd, 100, allSign)
		Expect(err).To(BeNil())
		txHash, err = asyncChangeValidator(au, goSDK)
		Expect(err).To(BeNil())
		buf[txHash] = true
		au.Nonce++
	}

	waitSyncBlock(goSDK)
	waitSyncBlock(ExpectNewSDK(nd.RPC))

	var bufInt = map[uint64]bool{}
	sameBlock := false
	for hash, _ := range buf {
		_, err := tests.WaitTxMined(goSDK, hash)
		Expect(err).To(BeNil())
		if !sameBlock {
			tra1 := ExpectGetBlockByHash(hash, goSDK)
			_, ok := bufInt[tra1.BlockHeight]
			if !ok {
				bufInt[tra1.BlockHeight] = true
			} else {
				sameBlock = true
				break
			}
		}
	}
	Expect(sameBlock).To(BeTrue())
	checkValidators(goSDK)
}

func replayTest() {
	node0 := actualValidatorsInfo[0]
	goSDK := ExpectNewSDK(actualValidatorsInfo[3].RPC)

	ExpectRemoveNode(node0, goSDK, allSign)
	node0.Role = NodeRoleNone
	checkValidators(goSDK)

	au := NewAdminOPUser(goSDK)
	err := au.MakeAddPeerMsg(node0, allSign)
	Expect(err).To(BeNil())
	ExpectChangeValidator(au, goSDK)
	node0.Role = NodeRolePeer
	checkValidators(goSDK)

	ExpectRemoveNode(node0, goSDK, allSign)
	node0.Role = NodeRoleNone
	checkValidators(goSDK)

	ExpectChangeValidator(au, goSDK)
	checkValidators(goSDK)

	newAU := NewAdminOPUser(goSDK)
	newAU.cmd = au.cmd
	ExpectChangeValidator(newAU, goSDK)
	checkValidators(goSDK)

	ExpectAddPeer(node0, goSDK, allSign)
	node0.Role = NodeRolePeer
	waitSyncBlock(goSDK)
	waitSyncBlock(ExpectNewSDK(node0.RPC))
	checkValidators(goSDK)

	ExpectUpdateNode(node0, 100, goSDK, allSign)
	node0.Role = NodeRoleValidator
	checkValidators(goSDK)
}

func reloadValidatorInfo(i int) {
	v, err := clusters.GetValidatorInfo(i)
	Expect(err).To(BeNil())
	actualValidatorsInfo[i].RPC = v.RPC
}

func beforeAfterEachCheck() {
	for _, v := range actualValidatorsInfo {
		Expect(v.Role).To(Equal(NodeRoleValidator))
	}
	checkValidatorStatus()
}
