test:
	@rm -rf tests/test_gen.go
	@go run main.go -dir tests
	@go test -v ./tests

	@rm -rf commons
	@mkdir commons
	@go run main.go -builtin -dir commons
	