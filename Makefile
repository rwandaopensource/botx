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

lint: vet fix fmt