GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/extract cmd/extract/main.go

docker:
	docker buildx build --platform=linux/amd64 --no-cache=true -f Dockerfile -t placeholder-custom-db .
