.PHONY: build-binary
build:
	go build -o bin/qp_laravel cmd/*.go

.PHONY: run-binary
run:
	go run cmd/main.go

.PHONY: clean-binary
clean-build:
	rm -f bin/main

.PHONY: move-binary
move:
	@echo "This target requires sudo privileges"
	sudo mv bin/qp_laravel /usr/local/bin/