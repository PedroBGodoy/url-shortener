.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down --remove-orphans

.PHONY: buf\:gen
buf\:gen:
	docker-compose up -d
	docker-compose exec app buf generate

.PHONY: buf\:lint
buf\:lint:
	docker-compose up -d
	docker-compose exec app buf lint

.PHONY: buf\:update
buf\:update:
	docker-compose up -d
	docker-compose exec app buf mod update

.PHONY: run
run:
	docker-compose up -d
	docker-compose exec app go run main.go

.PHONY: run\:dev
run\:dev:
	docker-compose up -d
	docker-compose exec app fresh

.PHONY: evans
evans:
	docker-compose up -d
	docker-compose exec app evans -r

.PHONY: tidy
tidy:
	docker-compose up -d
	docker-compose exec app go mod tidy

.PHONY: vet
vet:
	docker-compose up -d
	docker-compose exec app go vet ./...