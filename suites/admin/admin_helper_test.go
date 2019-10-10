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
	"encoding/json"
	"strings"
	"time"

	. "github.com/onsi/gomega"

	crypto "github.com/dappledger/AnnChain/gemmill/go-crypto"
	"github.com/dappledger/AnnChain/gemmill/types"
	sdk "github.com/dappledger/ann-go-sdk"
	"github.com/dappledger/ann-go-sdk/common"
	types2 "github.com/dappledger/ann-go-sdk/types"
	"github.com/dappledger/ann-tests"
	"github.com/dappledger/ann-tests/contracts"
)

var AdminTo = common.HexToAddress("0x02000000")

func ExpectSDKChangeNode(node *ValidatorInfo, cas []*ValidatorInfo, power int64, goSDK *sdk.GoSDK, opcmd types2.ValidatorCmd, accbase sdk.AccountBase) {
	data, err := goSDK.MakeNodeOpMsg(node.getPubHex(), power, accbase, opcmd)
	Expect(err).To(BeNil())
	sinfo := ExpectMakeNodeSign(node.PrivKey, data)
	if opcmd != types2.ValidatorCmdAddPeer {
		sinfo.Signature = nil
	}
	var caSinfo []types2.SigInfo
	for _, node := range cas {
		signPair := ExpectMakeNodeSign(node.PrivKey, data)

		newCaSinfo := types2.SigInfo{
			PubKey:    signPair.PublicKey,
			Signature: signPair.Signature,
		}

		caSinfo = append(caSinfo, newCaSinfo)
	}
	req, err := goSDK.MakeNodeContractRequest(data, sinfo.Signature, caSinfo, accbase)
	Expect(err).To(BeNil())
	_, err = goSDK.ContractCall(req)
	Expect(err).To(BeNil())
}

func ExpectUpdateNode(node *ValidatorInfo, power int64, goSDK *sdk.GoSDK, signNodesNum int) {
	au := NewAdminOPUser(goSDK)
	err := au.MakeUpdateNodeMsg(node, power, signNodesNum)
	Expect(err).To(BeNil())

	ExpectChangeValidator(au, goSDK)
}

func ExpectRemoveNode(node *ValidatorInfo, goSDK *sdk.GoSDK, signNodesNum int) {
	au := NewAdminOPUser(goSDK)
	err := au.MakeRemoveNodeMsg(node, signNodesNum)
	Expect(err).To(BeNil())

	ExpectChangeValidator(au, goSDK)
}

func ExpectAddPeer(node *ValidatorInfo, goSDK *sdk.GoSDK, signNodesNum int) {
	au := NewAdminOPUser(goSDK)
	err := au.MakeAddPeerMsg(node, signNodesNum)
	Expect(err).To(BeNil())
	ExpectChangeValidator(au, goSDK)
}

func ExpectChangeValidator(au *AdminOPUser, goSDK *sdk.GoSDK) {
	txHash, err := asyncChangeValidator(au, goSDK)
	Expect(err).To(BeNil())
	_, err = tests.WaitTxMined(goSDK, txHash)
	Expect(err).To(BeNil())
}

func asyncChangeValidator(au *AdminOPUser, goSDK *sdk.GoSDK) (string, error) {
	adminOp := contracts.NewAdminOp(goSDK, AdminTo)
	cmdBytes, err := json.Marshal(au.cmd)
	Expect(err).To(BeNil())
	account := sdk.AccountBase{PrivKey: au.Priv, Nonce: au.Nonce}
	return adminOp.Changenode(account, types2.TagAdminOPTx(cmdBytes))
}

type AdminOPUser struct {
	Priv          string
	Addr          string
	Nonce         uint64
	validatorAttr *types2.ValidatorAttr
	cmd           *types2.AdminOPCmd
}

func NewAdminOPUser(goSDK *sdk.GoSDK) *AdminOPUser {
	acc, err := goSDK.AccountCreate()
	Expect(err).To(BeNil())
	if strings.HasPrefix(acc.Address, "0x") || strings.HasPrefix(acc.Address, "0X") {
		acc.Address = acc.Address[2:]
	}
	user := &AdminOPUser{
		Priv: acc.Privkey,
		Addr: acc.Address,
	}
	user.Nonce = ExpectNonce(acc.Address, goSDK)
	return user
}

func (au *AdminOPUser) MakeAddPeerMsg(node *ValidatorInfo, signNodesNum int) error {
	vAttr, err := au.makeValidatorAttr(node.getPubHex(), 0, types2.ValidatorCmdAddPeer)
	if err != nil {
		return err
	}
	cmd, err := ContractsAdminCmd(vAttr)
	if err != nil {
		return err
	}
	pair := ExpectMakeNodeSign(node.PrivKey, cmd.Msg)
	if signNodesNum > 0 {
		cmd.SInfos = cmd.SInfos[:signNodesNum]
	}

	cmd.SelfSign = pair.Signature
	au.validatorAttr = vAttr
	au.cmd = cmd
	return nil
}

func (au *AdminOPUser) MakeUpdateNodeMsg(node *ValidatorInfo, power int64, signNodesNum int) error {
	vAttr, err := au.makeValidatorAttr(node.getPubHex(), power, types2.ValidatorCmdUpdateNode)
	if err != nil {
		return err
	}
	cmd, err := ContractsAdminCmd(vAttr)
	if signNodesNum > 0 {
		cmd.SInfos = cmd.SInfos[:signNodesNum]
	}
	au.validatorAttr = vAttr
	au.cmd = cmd
	return nil
}

func (au *AdminOPUser) MakeRemoveNodeMsg(node *ValidatorInfo, signNodesNum int) error {
	vAttr, err := au.makeValidatorAttr(node.getPubHex(), 0, types2.ValidatorCmdRemoveNode)
	if err != nil {
		return err
	}
	cmd, err := ContractsAdminCmd(vAttr)
	if signNodesNum >= 0 {
		cmd.SInfos = cmd.SInfos[:signNodesNum]
	}
	au.validatorAttr = vAttr
	au.cmd = cmd
	return nil
}

func (au *AdminOPUser) makeValidatorAttr(nodepub string, power int64, cmd types2.ValidatorCmd) (*types2.ValidatorAttr, error) {
	if strings.HasPrefix(nodepub, "0x") || strings.HasPrefix(nodepub, "0X") {
		nodepub = nodepub[2:]
	}
	pub, err := hex.DecodeString(nodepub)
	if err != nil {
		return nil, err
	}
	addrBytes, err := hex.DecodeString(au.Addr)
	if err != nil {
		return nil, err
	}
	return &types2.ValidatorAttr{
		PubKey: pub,
		Cmd:    cmd,
		Power:  power,
		Nonce:  au.Nonce,
		Addr:   addrBytes,
	}, nil
}

func ContractsAdminCmd(vAttr *types2.ValidatorAttr) (*types2.AdminOPCmd, error) {
	data, err := json.Marshal(vAttr)
	if err != nil {
		return nil, err
	}
	collectedSigns := mockCollectNodesSign(data)
	cmd := &types2.AdminOPCmd{}
	cmd.CmdType = types2.AdminOpChangeValidator
	cmd.Time = time.Now()
	cmd.Msg = data
	for _, s := range collectedSigns {
		cmd.SInfos = append(cmd.SInfos, types2.SigInfo{PubKey: s.PublicKey, Signature: s.Signature})
	}
	return cmd, nil
}

type SignPair struct {
	Signature []byte
	PublicKey []byte
}

func ExpectMakeNodeSign(privkey string, databytes []byte) *SignPair {
	privBytes := ExpectHex2Bytes(privkey)
	pk := crypto.SetNodePrivKey(privBytes)
	sigbytes := crypto.GetNodeSigBytes(pk.Sign(databytes))
	sp := SignPair{
		Signature: sigbytes,
		PublicKey: crypto.GetNodePubkeyBytes(pk.PubKey()),
	}

	return &sp
}

func ExpectHex2Bytes(hexstr string) []byte {
	if strings.HasPrefix(hexstr, "0x") || strings.HasPrefix(hexstr, "0X") {
		hexstr = hexstr[2:]
	}
	d, err := hex.DecodeString(hexstr)
	Expect(err).To(BeNil())
	return d
}

func ExpectNewSDK(rpc string) *sdk.GoSDK {
	goSDK := sdk.New(rpc, sdk.CyrptoType(tests.TestCryptoType))
	Expect(goSDK).NotTo(BeNil())
	return goSDK
}

func ExpectValidatorSet(goSDK *sdk.GoSDK) (int64, *types.ValidatorSet) {
	ret, err := goSDK.Validators()
	Expect(err).To(BeNil())
	var validators []*types.Validator
	for _, v := range ret.Validators {
		pub := ExpectHex2Bytes(v.PubKey)
		pubKey := crypto.SetNodePubkey(pub)
		validators = append(validators, &types.Validator{
			Address:     v.Address,
			PubKey:      pubKey,
			VotingPower: v.VotingPower,
			Accum:       v.Accum,
			IsCA:        v.VotingPower > 0,
		})
	}
	return ret.BlockHeight, types.NewValidatorSet(validators)
}

func ExpectNonce(addr string, goSDK *sdk.GoSDK) uint64 {
	nonce, err := goSDK.Nonce(addr)
	Expect(err).To(BeNil())
	return nonce
}

func ExpectHealth(hostSdk *sdk.GoSDK) bool {
	ok, err := hostSdk.CheckHealth()
	Expect(err).To(BeNil())
	return ok
}

func ExpectStopNode(idx int) {
	err := clusters.StopValidator(idx)
	Expect(err).To(BeNil())
	nd := actualValidatorsInfo[idx]
	nd.Running = false
}

func ExpectStartNode(idx int) {
	err := clusters.StartValidator(idx)
	Expect(err).To(BeNil())
	nd := actualValidatorsInfo[idx]
	nd.Running = true
	reloadValidatorInfo(idx)
}

func IsValidator(val *types.Validator) bool {
	return val.VotingPower > 0
}

func checkHeight(goSDK *sdk.GoSDK) bool {
	var h int64
	prevh, err := goSDK.LastHeight()
	Expect(err).To(BeNil())
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		h, err = goSDK.LastHeight()
		Expect(err).To(BeNil())
		if h > prevh {
			return true
		}
	}
	return false
}

func checkValidators(goSDK *sdk.GoSDK) {
	time.Sleep(time.Second)
	_, valSet := ExpectValidatorSet(goSDK)
	for _, validatorInfo := range actualValidatorsInfo {
		addrBytes, err := hex.DecodeString(validatorInfo.Address)
		Expect(err).To(BeNil())
		_, val := valSet.GetByAddress(addrBytes)
		switch validatorInfo.Role {
		case NodeRoleNone:
			Expect(val).To(BeNil())
		case NodeRolePeer:
			Expect(val).NotTo(BeNil())
			Expect(IsValidator(val)).To(BeFalse())
		case NodeRoleValidator:
			Expect(val).NotTo(BeNil())
			Expect(IsValidator(val)).To(BeTrue())
		}
	}
}

func mockCollectNodesSign(data []byte) (ret []*SignPair) {
	for _, val := range actualValidatorsInfo {
		if val.Role == NodeRoleValidator {
			pair := ExpectMakeNodeSign(val.PrivKey, data)
			ret = append(ret, pair)
		}
	}
	return ret
}

func checkMajor23() bool {
	total := 0
	major := 0
	for _, val := range actualValidatorsInfo {
		if val.Role == NodeRoleValidator {
			total += 100
			if val.Running {
				major += 100
			}
		}
	}
	total = total * 2 / 3
	return major > total
}

func checkValidatorStatus() {
	major23 := checkMajor23()
	for _, sval := range actualValidatorsInfo {
		val := sval
		if !sval.Running {
			goSDK := ExpectNewSDK(val.RPC)
			_, err := goSDK.CheckHealth()
			Expect(err).NotTo(BeNil())
		} else {
			switch sval.Role {
			case NodeRoleNone:
				goSDK := ExpectNewSDK(val.RPC)
				ok := ExpectHealth(goSDK)
				Expect(ok).To(BeTrue())
				ok = checkHeight(goSDK)
				Expect(ok).To(BeFalse())
			case NodeRolePeer, NodeRoleValidator:
				goSDK := ExpectNewSDK(val.RPC)
				ok := ExpectHealth(goSDK)
				Expect(ok).To(BeTrue())
				ok = checkHeight(goSDK)
				if major23 {
					Expect(ok).To(BeTrue())
				} else {
					Expect(ok).To(BeFalse())
				}
			}
		}
	}
}

func waitSyncBlock(goSDK *sdk.GoSDK) {
	ok := false
	for i := 0; i < 50; i++ {
		ok = checkHeight(goSDK)
		if !ok {
			time.Sleep(time.Second)
		} else {
			break
		}
	}
	Expect(ok).To(BeTrue())
}

func ExpectGetBlockByHash(txHash string, goSDK *sdk.GoSDK) *sdk.RPCTransaction {
	ret, err := goSDK.GetTransactionByHash(txHash)
	Expect(err).To(BeNil())
	return ret
}
