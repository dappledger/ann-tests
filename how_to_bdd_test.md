## 简述
BDD测试的基本思想是，根据指定版本部署一个多节点的区块链集群，可以是单机的，也可是多机分布式的。
一般的，本地开发环境可以使用单机，而CI环境，尤其在做TPS性能测试时，应该使用分布式集群环境。

## 通过环境变量设置测试环境

```
TEST_VALIDATOR_NUM=4 （验证节点数量）
TEST_PLATFORM=binary 
TEST_CRYPTO_TYPE=ZA （加密方法）
TEST_GENESIS_IMAGE=annchain/genesis:latest (待测试的genesis镜像)
```