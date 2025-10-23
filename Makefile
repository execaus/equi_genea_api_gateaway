# Переменные
API_CONTRACTS_MODULE := github.com/execaus/equi_genea_api_contracts
PROTO_OUT_DIR := internal/pb
PROTOC := protoc
PROTOC_PLUGINS := google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc

install_protoc_plugins:
	@echo "Installing protoc Go plugins..."
	@for plugin in $(PROTOC_PLUGINS); do \
		go install $$plugin@latest; \
	done

generate_protobuf:
	$(eval CONTRACTS_DIR := $(shell go list -m -f '{{.Dir}}' $(API_CONTRACTS_MODULE)))
	@echo "Using contracts from: $(CONTRACTS_DIR)"
	@echo "Creating output directory: $(PROTO_OUT_DIR)"
	@mkdir -p $(PROTO_OUT_DIR)
	@echo "Generating Go protobuf and gRPC code..."
	@find "$(CONTRACTS_DIR)" -type f -name '*.proto' | xargs $(PROTOC) \
		-I "$(CONTRACTS_DIR)" \
		--go_out="$(PROTO_OUT_DIR)" \
		--go_opt=paths=source_relative \
		--go-grpc_out="$(PROTO_OUT_DIR)" \
		--go-grpc_opt=paths=source_relative
	@echo "Generation completed! Files are in $(PROTO_OUT_DIR)"

update_api_contracts:
	@( \
		CURRENT=$$(go list -m -f '{{.Version}}' $(API_CONTRACTS_MODULE) 2>/dev/null || echo "none"); \
		echo "Current API contracts version: $$CURRENT"; \
		echo "Fetching latest tags from GitHub..."; \
		LATEST=$$(git ls-remote --tags https://$(API_CONTRACTS_MODULE).git | grep -o 'refs/tags/v[0-9]\+\.[0-9]\+\.[0-9]\+$$' | sort -V | tail -n1 | sed 's/refs\/tags\///'); \
		if [ -z "$$LATEST" ]; then \
			echo "No tags found on GitHub"; \
		elif [ "$$CURRENT" = "$$LATEST" ]; then \
			echo "Already up to date."; \
		else \
			echo "Updating to $$LATEST..."; \
			go get $(API_CONTRACTS_MODULE)@$$LATEST; \
			echo "Updated API contracts to $$LATEST"; \
		fi \
	)
