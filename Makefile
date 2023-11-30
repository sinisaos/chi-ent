server:
	go run cmd/main.go

generate:
	go generate ./ent

tests:
	go test -v -cover ./pkg/service/...