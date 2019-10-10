#!/usr/bin/env bash

docker run --rm -v `pwd`:/contracts ethereum/solc:0.4.25 \
    --bin --abi -o /contracts/abi_bin \
    --overwrite /contracts/token.sol /contracts/kv_storage.sol /contracts/vari_and_fun.sol /contracts/limit.sol

docker run --rm -v `pwd`:/contracts ethereum/solc:0.5.1 \
    --bin --abi -o /contracts/abi_bin \
    --overwrite /contracts/admin_op.sol

go generate gen.go

