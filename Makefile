.PHONY: dev
dev:
	docker-compose up -d
	docker-compose exec app bash

.PHONY: down
down:
	docker-compose down --remove-orphans

.PHONY: gen
gen:
	buf generate

.PHONY: lint
lint:
	buf lint

.PHONY: run
run:
	go run cmd/shortener/main.go