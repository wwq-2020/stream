test:
	@rm -rf tests/test_gen.go
	@go run main.go -dir tests
	@go test -v ./tests


	