default: build

build:
	go build .

test:
	test -z `go fmt ./...`
	go test ./...

builddockerbinary:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o datediff_docker .
	chmod +x datediff_docker

release: test build

CI: test

