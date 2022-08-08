
mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(patsubst %/,%,$(dir $(mkfile_path)))

.PHONY: dev-up
dev-up:
	docker-compose up -d

.PHONY: dev-infra-up
dev-infra-up:
	docker-compose up -d redis localstack

.PHONY: dev-infra-stop
dev-infra-stop:
	docker-compose stop
