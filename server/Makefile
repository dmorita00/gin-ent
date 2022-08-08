
mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(patsubst %/,%,$(dir $(mkfile_path)))
exec_file := ./tmp/main

export GOBIN=$(current_dir)/bin

.PHONY: gen
gen: clean ent

.PHONY: ent
ent:
	go generate ./...
	go run entgo.io/ent/cmd/ent@v0.11.1 describe ./ent/schema > schema.txt

.PHONY: build
build:
	go build -o $(exec_file)

.PHONY: start
start: build
	$(exec_file)

.PHONY: hotreload
hotreload:
	./bin/air

.PHONY: gensql
gensql: build
	$(exec_file) gensql

.PHONY: migrate-up
migrate-up:
	./bin/sql-migrate up

.PHONY: migrate-down
migrate-down:
	./bin/sql-migrate down

.PHONY: lint
lint:
	./bin/golangci-lint run

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	git clean -fdX ent

.PHONY: seed
seed: build
	$(exec_file) seed
