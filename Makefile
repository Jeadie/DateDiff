default: build

build:
	go build -o datediff .

test:
	gofmt -d -l `find ./** | grep "\.go"`
	go test ./...

release: test build

CI: test