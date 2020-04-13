.PHONY: run test

PORT := 3001

run:
	GO111MODULE=on go build -o victoria && ./victoria -port=${PORT}

test:
	GO111MODULE=on go test ./tests/...
