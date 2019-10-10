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

package pbft

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/dappledger/AnnChain/gemmill/types"
	"github.com/dappledger/ann-go-sdk"
	"github.com/dappledger/ann-go-sdk/common"
	types2 "github.com/dappledger/ann-go-sdk/types"
	"github.com/dappledger/ann-tests"
	"github.com/dappledger/ann-tests/cluster"
)

func TestTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PBFT Suite")
}

var _ = Describe("Tests PBFT", func() {

	Context(fmt.Sprintf("With %d validators", tests.TestValidatorNum), func() {

		var clusters cluster.Cluster
		var vals []*types.Validator
		BeforeEach(func() {
			var err error
			clusters, err = tests.NewClusterFromEnv()
			Expect(err).To(BeNil())
			err = clusters.Up()
			Expect(err).To(BeNil())

			vals = make([]*types.Validator, 0, clusters.ValidatorNum())
			for i := 0; i < clusters.ValidatorNum(); i++ {
				node, err := clusters.GetValidatorInfo(i)
				Expect(err).To(BeNil())
				vals = append(vals, &types.Validator{Address: common.Hex2Bytes(node.Address), VotingPower: 100})
			}
		})
		AfterEach(func() {
			if clusters == nil {
				return
			}
			if tests.TestKeepCluster {
				return
			}
			err := clusters.Down()
			Expect(err).To(BeNil())
		})

		It("should be create blocks by each validator", func() {

			valSet := types.NewValidatorSet(vals)
			node, err := clusters.GetValidatorInfo(0)
			Expect(err).To(BeNil())

			cli := sdk.New(node.RPC, sdk.CyrptoType(tests.TestCryptoType))
			err = cli.IterBlock(1, int64(tests.TestValidatorNum*3), func(prev, cur *types2.ResultBlock) error {
				Expect(cur.BlockMeta.Header.ProposerAddress).To(Equal(valSet.Proposer().Address))
				valSet.IncrementAccum(1)
				return nil
			})
			Expect(err).To(BeNil())
		})

		It("with one validator stopped, should be create blocks by left", func() {

			valSet := types.NewValidatorSet(vals)

			stoppedValidator, err := clusters.GetValidatorInfo(1)
			Expect(err).To(BeNil())
			stopCli := sdk.New(stoppedValidator.RPC, sdk.CyrptoType(tests.TestCryptoType))

			err = stopCli.IterBlock(1, 1+int64(tests.TestValidatorNum), func(prev, cur *types2.ResultBlock) error { return nil })
			Expect(err).To(BeNil())

			err = clusters.StopValidator(1)
			Expect(err).To(BeNil())

			node, err := clusters.GetValidatorInfo(0)
			Expect(err).To(BeNil())

			cli := sdk.New(node.RPC, sdk.CyrptoType(tests.TestCryptoType))
			from, err := cli.LastHeight()
			Expect(err).To(BeNil())
			to := from + int64(tests.TestValidatorNum*2)

			valSet = types.NewValidatorSet(vals)
			for i := 0; i < int(from)-1; i++ {
				valSet.IncrementAccum(1)
			}

			err = cli.IterBlock(from, to, func(prev, cur *types2.ResultBlock) error {
				if prev == nil {
					return nil
				}
				Expect(prev.BlockMeta.Header.ProposerAddress).ToNot(Equal(stoppedValidator.Address))
				var originValSet = valSet.Copy()
				for _, v := range cur.Block.LastCommit.Precommits {
					if v != nil {
						if v.ValidatorAddress != nil {
							originValSet.IncrementAccum(int64(v.Round))
							break
						}
					}
				}
				p := originValSet.Proposer().Address
				Expect(prev.BlockMeta.Header.ProposerAddress).To(Equal(p))
				valSet.IncrementAccum(1)
				return nil
			})
			Expect(err).To(BeNil())

			err = clusters.StartValidator(1)
			Expect(err).To(BeNil())
			time.Sleep(time.Second * 3)

			stoppedValidator, err = clusters.GetValidatorInfo(1)
			Expect(err).To(BeNil())
			stopCli = sdk.New(stoppedValidator.RPC, sdk.CyrptoType(tests.TestCryptoType))

			from, err = stopCli.LastHeight()
			Expect(err).To(BeNil())
			// find stopped validator
			err = stopCli.IterBlock(from, from+int64(tests.TestValidatorNum*10), func(prev, cur *types2.ResultBlock) error {
				if hex.EncodeToString(cur.BlockMeta.Header.ProposerAddress) == strings.ToLower(stoppedValidator.Address) {
					from = int64(cur.BlockMeta.Header.Height)
					return errors.New("found stopped validator")
				}
				return nil
			})
			Expect(err).NotTo(BeNil())

			valSet = types.NewValidatorSet(vals)
			for i := 0; i < int(from)-1; i++ {
				valSet.IncrementAccum(1)
			}
			to = from + int64(tests.TestValidatorNum*2)
			err = cli.IterBlock(from, to, func(prev, cur *types2.ResultBlock) error {
				Expect(cur.BlockMeta.Header.ProposerAddress).To(Equal(valSet.Proposer().Address))
				valSet.IncrementAccum(1)
				return nil
			})
		})
	})
})
