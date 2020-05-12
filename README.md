# Example
TEST_PLATFORM=binary TEST_GENESIS_IMAGE=$GOPATH/src/github.com/dappledger/AnnChain/build/genesis TEST_CONSENSUS_TYPE=pbft  make test-bdd

# benchmark
    部署链环境，运行程序直接跑就可以了.参数说明如下：  
- nthread : 跑N个协程(默认10).
- tps : 每个协程 每秒 N次请求(默认100).
- nsec : 在上述条件下，跑N秒钟(默认100).
- rpcarg : 链的rpc地址，可以多个链地址，用','分隔.（例如：172.28.152.153:46657,172.28.152.154:46657,172.28.152.155:46657,172.28.152.156:46657）
- dotype : 操作类型，现在支持`set(存证上链)`,`cwrite(合约上链)`（默认'set'）.
- dosync : 是`同步(true)`还是`异步(false)`操作.(默认是异步操作).
- debug : 是否是debug模式——如果debug模式下，会打印更多的日志，显示更多信息.(默认非debug模式)

> 从上面参数说明来看，`dotype`是必须要指定的，其他都可以用默认值.

