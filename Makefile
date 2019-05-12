test:
	@rm -rf outter/outter_gen.go
	@rm -rf tests/test_gen.go
	@rm -rf commons/*
	@go build .
	
	@./stream -builtin -dir commons
	@./stream -dir outter 
	@./stream -dir tests -fs
	@go test -v ./tests

