.PHONY: build-binary
build:
	go build -o bin/main cmd/main.go

.PHONY: run-binary
run:
	go run cmd/main.go

.PHONY: clean-binary
clean-build:
	rm -f bin/main

.PHONY: move-binary
move:
	mv bin/main /usr/local/bin/