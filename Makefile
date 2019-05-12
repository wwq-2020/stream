test:
	@rm -rf outter/outter_gen.go
	@rm -rf tests/test_gen.go
	@go run main.go -dir outter
	@go run main.go -dir tests -fs
	@go test -v ./tests

	@rm -rf commons
	@mkdir commons
	@go run main.go -builtin -dir commons