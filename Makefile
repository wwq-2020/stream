test:
	@rm -rf outter/outter_gen.go
	@rm -rf tests/test_gen.go
	@rm -rf commons/*

	@./stream -builtin -dir commons
	@go build .
	@./stream -dir outter 
	@./stream -dir tests -fs
	@go test -v ./tests

