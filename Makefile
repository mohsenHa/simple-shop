ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

lint:
	which golangci-lint || (go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.0)
	cd ./source/customer-sercice
	golangci-lint run --config=$(ROOT)/../../.golangci.yml $(ROOT)/...
	cd ./source/order-sercice
	golangci-lint run --config=$(ROOT)/../../.golangci.yml $(ROOT)/...
	cd ./source/product-sercice
	golangci-lint run --config=$(ROOT)/../../.golangci.yml $(ROOT)/...


test:
	cd ./source/customer-sercice
	go test ./...
	cd ./source/order-sercice
	go test ./...
	cd ./source/product-sercice
	go test ./...

format:
	@which gofumpt || (go install mvdan.cc/gofumpt@latest)
	@gofumpt -l -w $(ROOT)
	@which gci || (go install github.com/daixiang0/gci@latest)
	@gci write $(ROOT)
	@which golangci-lint || (go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.0)
	@golangci-lint run --fix