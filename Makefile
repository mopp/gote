gote:
	cd cmd/gote && go build

.PHONY: run
run: gote
	./cmd/gote/gote

.PHONY: lint
lint:
	golangci-lint run
