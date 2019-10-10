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

package tests

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dappledger/AnnChain/gemmill/go-crypto"
	gcrypto "github.com/dappledger/ann-go-sdk/crypto"
	. "github.com/dappledger/ann-tests/cluster"
	"github.com/dappledger/ann-tests/cluster/binary"
	"github.com/dappledger/ann-tests/cluster/docker"
)

var (
	TestValidatorNum  = lookupIntEnv("TEST_VALIDATOR_NUM", 4)
	TestPlatform      = lookupStringEnv("TEST_PLATFORM", "binary")
	TestCryptoType    = lookupStringEnv("TEST_CRYPTO_TYPE", "ZA")
	TestConsensusType = lookupStringEnv("TEST_CONSENSUS_TYPE", "pbft")
)

var (
	TestGenesisImage = lookupStringEnv("TEST_GENESIS_IMAGE", "annchain/genesis:latest")
)

var (
	TestKeepCluster = lookupBoolEnv("TEST_KEEP_CLUSTER", false)
	TestDebug       = lookupBoolEnv("TEST_DEBUG", false)
)

func init() {
	fmt.Printf(`Testing Env:
	image: %s
	crypto_type: %s
	validators num: %d, 
`, TestGenesisImage, TestCryptoType, TestValidatorNum)
}

func lookupIntEnv(key string, defaultValue int) int {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}
	return i
}

func lookupStringEnv(key, defaultValue string) string {

	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return v
}

func lookupBoolEnv(key string, defaultValue bool) bool {

	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return strings.ToLower(v) == "true"
}

func NewClusterFromEnv() (Cluster, error) {

	blockBrowserImage := "annchain/browser"

	crypto.NodeInit(TestCryptoType)
	gcrypto.NodeInit(TestCryptoType)

	if TestPlatform == "binary" {
		c, err := binary.New(&binary.Option{
			ValidatorNum: TestValidatorNum,
			CryptoType:   TestCryptoType,
			BinaryPath:   TestGenesisImage,
			Consensus:    TestConsensusType,
		})

		return c, err
	}

	return docker.NewDockerCompose(&docker.Option{
		ValidatorNum:      TestValidatorNum,
		IPPrefix:          "192.168.32",
		HasBrowser:        false,
		GenesisImage:      TestGenesisImage,
		BlockBrowserImage: blockBrowserImage + ":latest",
		CryptoType:        TestCryptoType,
	}), nil
}
