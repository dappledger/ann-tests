#!/usr/bin/env bash

docker run --rm -v `pwd`:/contracts ethereum/solc:0.5.0 \
    --evm-version constantinople \
    --bin --abi -o /contracts/abi_bin \
    --overwrite /contracts/admin_op.sol /contracts/token.sol /contracts/kv_storage.sol /contracts/vari_and_fun.sol

go generate gen.go
