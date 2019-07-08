.PHONY: build test run-dot

build:
	cd cmd/gomponere/; go build -race -o ../../dist/gomponere
test:
	ginkgo ./... -cover -race --randomizeAllSpecs --failOnPending
run-dot: build
	./dist/gomponere | dot -Tpng  > support/output/test.png && open support/output/test.png