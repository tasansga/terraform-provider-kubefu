DIST_DIR=dist
RESOURCEGEN_BIN=$(DIST_DIR)/resourcegen
PROVIDER_BIN=$(DIST_DIR)/terraform-provider-kubefu

default: all

all: lint unittest build-provider

$(DIST_DIR):
	mkdir -p $(DIST_DIR)

build-resourcegen: | $(DIST_DIR)
	go build -o "$(RESOURCEGEN_BIN)" cmd/resourcegen/main.go

build-provider: | $(DIST_DIR)
	go build -o "$(PROVIDER_BIN)" cmd/terraform-provider-kubefu/main.go

build: build-resourcegen build-provider

generate: build-resourcegen
	"$(RESOURCEGEN_BIN)" generate schemas

schema-download: build-resourcegen
	"$(RESOURCEGEN_BIN)" download-schema

lint: generate
	#tfproviderlint -R001=false ./...
	golangci-lint run kubefu
	golangci-lint run kubefu/generated
	golangci-lint run cmd/resourcegen
	golangci-lint run cmd/terraform-provider-kubefu

modupdate:
	go get -u ./...
	go mod tidy

tidy:
	go mod tidy

docs: generate
	rm -Rf ./docs ./cmd/terraform-provider-kubefu/docs
	tfplugindocs generate --provider-dir ./cmd/terraform-provider-kubefu
	rm -Rf ./docs
	mv ./cmd/terraform-provider-kubefu/docs .

test: unittest inttest

unittest: unittest-kubefu unittest-resourcegen

unittest-kubefu:
	go test -v ./kubefu

unittest-resourcegen:
	go test -v ./resourcegen

inttest:
	./inttest/test.sh apply
	./inttest/test.sh destroy

debug:
	DEBUG_LOCAL=yes go run cmd/terraform-provider-kubefu/main.go

clean:
	go clean
	rm -f "$(PROVIDER_BIN)" "$(RESOURCEGEN_BIN)"

.PHONY: all build build-resourcegen build-provider lint modupdate tidy docs test inttest unittest unittest-kubefu unittest-resourcegen debug clean generate schema-download
