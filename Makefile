ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

lint:
	which golangci-lint || (go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.0)
	golangci-lint run --config=$(ROOT)/.golangci.yml $(ROOT)/...

test:
	go test ./...

logs:
	docker-compose logs

format:
	@which gofumpt || (go install mvdan.cc/gofumpt@latest)
	@gofumpt -l -w $(ROOT)
	@which gci || (go install github.com/daixiang0/gci@latest)
	@gci write $(ROOT)
	@which golangci-lint || (go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.0)
	@golangci-lint run --fix

start-infrastructure:
	docker compose -p shop-infrastructure --env-file ./.env -f ./serve/services/infrastructure/api-gateway/gateway.yml -f ./serve/services/infrastructure/database/postgress.yml -f ./serve/services/infrastructure/message-broker/broker.yml up -d

start-product-service:
	cd source/product-service
	air -c ./.air.win.toml

start-order-service:
	cd source/order-service
	air -c ./.air.win.toml

start-order-service:
	cd source/customer-service
	air -c ./.air.win.toml
