.PHONY: build build-dev test

build-dev: --copy-env --go-fmt build

build:
	echo "Building..."
	go build -o ./bin/quantumblogserver ./quantumblogserver.go; \
	echo "Built to ./bin/quantumblogserver"; \
	echo "\n\n";

--go-fmt:
	go fmt ./...;

--copy-env:
	@if [ ! -f conifg/.env ]; then \
		cp config/.env.example config/.env; \
		echo "Copied config/.env.example to config/.env"; \
	fi
