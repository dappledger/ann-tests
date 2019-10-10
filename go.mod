module github.com/dappledger/ann-tests

go 1.12

require (
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/dappledger/AnnChain v0.0.0
	github.com/dappledger/ann-go-sdk v0.0.0
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/onsi/ginkgo v1.8.0
	github.com/onsi/gomega v1.5.0
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/pkg/errors v0.8.1
	github.com/rcrowley/go-metrics v0.0.0-20190826022208-cac0b30c2563
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
)

replace (
	github.com/dappledger/AnnChain => ../AnnChain
	github.com/dappledger/ann-go-sdk => ../ann-go-sdk
	github.com/dappledger/ann-tests => ../ann-tests
)
