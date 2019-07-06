build:
	go build -o dist/gomponere
run-dot: build
	./dist/gomponere | dot -Tpng  > support/output/test.png && open support/output/test.png