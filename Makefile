SHELL         := /bin/bash
PACK          := cilium
PROJECT       := github.com/littlejo/pulumi-$(PACK)
PROVIDER_PATH := provider
TFGEN         := pulumi-tfgen-$(PACK)
PROVIDER      := pulumi-resource-$(PACK)
WORKING_DIR   := $(shell pwd)

# Version stamped into the provider and the SDKs. The SDKs committed to the repository are generated
# with this development version so that regenerating them is reproducible; the release workflow
# passes the version of the tag instead.
PROVIDER_VERSION ?= 0.0.0-alpha.0+dev
LDFLAGS          := -X $(PROJECT)/$(PROVIDER_PATH)/pkg/version.Version=$(PROVIDER_VERSION)

# The Go SDK requires the Pulumi SDK version the code generator was built with.
PULUMI_SDK_VERSION = $(shell cd $(PROVIDER_PATH) && go list -m -f '{{.Version}}' github.com/pulumi/pulumi/sdk/v3)

TESTPARALLELISM ?= 1

.PHONY: help build tfgen schema schema_embed provider build_sdks build_nodejs build_python build_go \
	lint fmt test_provider test clean

help: ## Show this help
	@grep -E '^[a-z_]+:.*## ' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*## "}; {printf "\033[36m%-14s\033[0m %s\n", $$1, $$2}'

build: provider build_sdks ## Build the provider and all the SDKs

tfgen: ## Build the code generator
	cd $(PROVIDER_PATH) && go build -o $(WORKING_DIR)/bin/$(TFGEN) -ldflags "$(LDFLAGS)" ./cmd/$(TFGEN)

schema: tfgen ## Regenerate the Pulumi schema from the Terraform provider
	$(WORKING_DIR)/bin/$(TFGEN) schema --out $(PROVIDER_PATH)/cmd/$(PROVIDER)

schema_embed: ## Stamp the version into the schema embedded in the provider binary
	cd $(PROVIDER_PATH)/cmd/$(PROVIDER) && VERSION=$(PROVIDER_VERSION) go generate .

provider: schema_embed ## Build the provider binary
	cd $(PROVIDER_PATH) && go build -o $(WORKING_DIR)/bin/$(PROVIDER) -ldflags "$(LDFLAGS)" ./cmd/$(PROVIDER)

build_sdks: build_nodejs build_python build_go ## Generate and build all the SDKs

build_nodejs: tfgen ## Generate and build the Node.js SDK
	rm -rf sdk/nodejs
	$(WORKING_DIR)/bin/$(TFGEN) nodejs --out sdk/nodejs/
	cd sdk/nodejs && \
		npm install --no-package-lock && \
		npx tsc && \
		cp ../../README.md ../../LICENSE package.json ./bin/

build_python: tfgen ## Generate and build the Python SDK
	rm -rf sdk/python
	$(WORKING_DIR)/bin/$(TFGEN) python --out sdk/python/
	cd sdk/python && \
		mkdir bin && \
		cp -R littlejo_cilium pyproject.toml ../../README.md ./bin/ && \
		python3 -m build --outdir ./bin/dist ./bin

build_go: tfgen ## Generate and build the Go SDK
	rm -rf sdk/go
	$(WORKING_DIR)/bin/$(TFGEN) go --out sdk/go/
	cd sdk && \
		go get github.com/pulumi/pulumi/sdk/v3@$(PULUMI_SDK_VERSION) && \
		go mod tidy && \
		go build ./...

lint: schema_embed ## Lint the provider code
	cd $(PROVIDER_PATH) && golangci-lint run -c ../.golangci.yml

fmt: ## Format the provider code
	cd $(PROVIDER_PATH) && golangci-lint fmt -c ../.golangci.yml

test_provider: ## Run the provider unit tests
	cd $(PROVIDER_PATH) && go test -v -count=1 ./...

test: provider ## Run the examples against the current Kubernetes context (needs a cluster without CNI)
	cd examples && go test -v -count=1 -tags=all -parallel $(TESTPARALLELISM) -timeout 1h ./...

clean: ## Remove the build outputs
	rm -rf bin sdk/nodejs/bin sdk/nodejs/node_modules sdk/python/bin
	rm -f $(PROVIDER_PATH)/cmd/$(PROVIDER)/schema-embed.json
