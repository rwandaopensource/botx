vet:
	go vet .

fix:
	go fix .

fmt:
	gofmt -w . || go fmt -w .

test:
	go test -v -race -coverprofile=coverage.coverprofile -covermode=atomic -tags integration ./...

build:
	go build -o bin/main main.go

run:
	go run main.go

start_db: #this may only work on UNIX
	mongod --fork --logpath ~/mongod/mongod.log --dbpath ~/mongod

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/linux/arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/linux/arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/freebsd/386 main.go

lint: vet fix fmt