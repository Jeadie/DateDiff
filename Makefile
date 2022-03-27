default: build

build:
	go build .

test:
	gofmt -d -l `find ./** | grep "\.go"`
	go test ./...

builddockerbinary:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o datediff_docker .
	chmod +x datediff_docker

release: build test

CI: test

