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
	"time"

	"github.com/pkg/errors"
	"github.com/dappledger/ann-go-sdk"
	"github.com/dappledger/ann-go-sdk/types"
)

var ErrMinedTimeout = errors.New("timeout for wait tx to be mined")

func WaitTxMined(goSDK *sdk.GoSDK, tx string) (*types.ReceiptForDisplay, error) {
	t, err := goSDK.Receipt(tx)
	if err == nil {
		return t, err
	}
	wait := time.After(time.Second * 30)
	ticker := time.NewTicker(time.Millisecond * 300)
	defer ticker.Stop()

	for {
		select {
		case <-wait:
			return nil, ErrMinedTimeout
		case <-ticker.C:
			t, err := goSDK.Receipt(tx)
			if err == nil {
				return t, err
			}
		}
	}

	return nil, nil
}
