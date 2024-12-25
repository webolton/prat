get:
	go get -u ./...

run:
	go run cmd/prat/main.go

init:
	go run cmd/prat/main.go init

specs:
	ginkgo ./...
