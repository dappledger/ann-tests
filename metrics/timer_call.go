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

package metrics

import (
	"fmt"
	"time"
)

func Run(dur time.Duration, callPerSecond int, callFunc func()) {

	durTicker := time.After(dur)
	callTicker := time.NewTicker(time.Duration(float64(time.Second) / float64(callPerSecond)))
	defer callTicker.Stop()
L1:
	for {
		select {
		case <-durTicker:
			break L1
		case <-callTicker.C:
			start := time.Now()
			callFunc()
			end := time.Now()
			if end.Sub(start) > time.Millisecond*200 {
				fmt.Println("callFunc takes", time.Now().Sub(start))
			}
		}
	}
	return
}
