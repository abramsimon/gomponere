run-dot:
	go run main.go | dot -Tpng  > support/test.png && open support/test.png