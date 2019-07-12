.PHONY: build test run-dot

build:
	cd cmd/gomponere/; CGO_ENABLED=0 go build -o ../../dist/gomponere
test:
	ginkgo -cover -race --randomizeAllSpecs --failOnPending ./...
run-dot: build
	cd dist; ./gomponere -i=../test/input | dot -Tpng  > ../test/output/test.png && open ../test/output/test.png