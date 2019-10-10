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

var (
	_ = ""
)

//go:generate rtool abigen --abi abi_bin/KvStorage.abi --bin abi_bin/KvStorage.bin --type KvStorage --pkg contracts --out kv_storage.go
//go:generate rtool abigen --abi abi_bin/WrapKvStorage.abi --bin abi_bin/WrapKvStorage.bin --type WrapKvStorage --pkg contracts --out wrap_kv_storage.go
//go:generate rtool abigen --abi abi_bin/Token.abi --bin abi_bin/Token.bin --type Token --pkg contracts --out token.go
//go:generate rtool abigen --abi abi_bin/VariAndFun.abi --bin abi_bin/VariAndFun.bin --type VariAndFun --pkg contracts --out vari_and_fun.go
//go:generate rtool abigen --abi abi_bin/VariAndFunOrigin.abi --bin abi_bin/VariAndFunOrigin.bin --type VariAndFunOrigin --pkg contracts --out vari_and_fun_origin.go
//go:generate rtool abigen --abi abi_bin/Admin.abi --bin abi_bin/Admin.bin --type AdminOp --pkg contracts --out admin_op.go
