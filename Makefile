.PHONY: ann api all

test-bdd:
	go test -count=1 -v ./suites/contracts/...
ifneq ($(TEST_CONSENSUS_TYPE), raft)
	go test -count=1 -v ./suites/admin/...
	go test -count=1 -v ./suites/pbft/...
	go test -count=1 -v ./suites/payload/...
endif