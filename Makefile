build:
	go build -o bin/brainf cmd/brainf/main.go

test:
	go test ./...

fmt:
	go fmt ./...

install:
	cd cmd/brainf && go install

uninstall:
	rm -f $(shell which brainf)
