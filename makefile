.PHONY: build test run-dot

build:
	cd cmd/gomponere/; go build -race -o ../../dist/gomponere
test:
	ginkgo -cover -race --randomizeAllSpecs --failOnPending ./...
run-dot: build
	cd dist; ./gomponere | dot -Tpng  > ../test/output/test.png && open ../test/output/test.png