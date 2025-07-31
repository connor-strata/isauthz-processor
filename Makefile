# Makefile for building all isauthz-processor implementations

# Output directory for binaries
BINDIR = bin

# Binary names
GO_BIN = $(BINDIR)/isauthz-processor-go
PYTHON_BIN = $(BINDIR)/isauthz-processor-python
JS_BIN = $(BINDIR)/isauthz-processor-js
RUST_BIN = $(BINDIR)/isauthz-processor-rust
C_BIN = $(BINDIR)/isauthz-processor-c

# Default target
.PHONY: all
all: $(GO_BIN) $(PYTHON_BIN) $(JS_BIN) $(RUST_BIN) $(C_BIN)

# Create bin directory
$(BINDIR):
	mkdir -p $(BINDIR)

# Go build
$(GO_BIN): go/main.go go/go.mod | $(BINDIR)
	cd go && go build -o ../$(GO_BIN) .

# Python "build" (create executable wrapper)
$(PYTHON_BIN): python/main.py | $(BINDIR)
	echo '#!/bin/bash' > $(PYTHON_BIN)
	echo 'exec python3 "$(shell pwd)/python/main.py" "$$@"' >> $(PYTHON_BIN)
	chmod +x $(PYTHON_BIN)

# JavaScript "build" (create executable wrapper)
$(JS_BIN): javascript/main.js javascript/package.json | $(BINDIR)
	echo '#!/bin/bash' > $(JS_BIN)
	echo 'exec node "$(shell pwd)/javascript/main.js" "$$@"' >> $(JS_BIN)
	chmod +x $(JS_BIN)

# Rust build
$(RUST_BIN): rust/src/main.rs rust/Cargo.toml | $(BINDIR)
	cd rust && cargo build --release
	cp rust/target/release/isauthz-processor-rust $(RUST_BIN)

# C build
$(C_BIN): c/main.c | $(BINDIR)
	gcc -o $(C_BIN) c/main.c -std=c99

# Individual build targets
.PHONY: go python javascript rust c
go: $(GO_BIN)
python: $(PYTHON_BIN)
javascript: $(JS_BIN)
rust: $(RUST_BIN)
c: $(C_BIN)

# Test targets
.PHONY: test test-go test-python test-javascript test-rust test-c
test: test-go test-python test-javascript test-rust test-c

test-go: $(GO_BIN)
	@echo "Testing Go implementation..."
	@echo '{"azure.authenticated": "true", "azure.role": "admin"}' | $(GO_BIN)
	@echo '{"azure.authenticated": "false", "azure.role": "admin"}' | $(GO_BIN)

test-python: $(PYTHON_BIN)
	@echo "Testing Python implementation..."
	@echo '{"azure.authenticated": "true", "azure.role": "admin"}' | $(PYTHON_BIN)
	@echo '{"azure.authenticated": "false", "azure.role": "admin"}' | $(PYTHON_BIN)

test-javascript: $(JS_BIN)
	@echo "Testing JavaScript implementation..."
	@echo '{"azure.authenticated": "true", "azure.role": "admin"}' | $(JS_BIN)
	@echo '{"azure.authenticated": "false", "azure.role": "admin"}' | $(JS_BIN)

test-rust: $(RUST_BIN)
	@echo "Testing Rust implementation..."
	@echo '{"azure.authenticated": "true", "azure.role": "admin"}' | $(RUST_BIN)
	@echo '{"azure.authenticated": "false", "azure.role": "admin"}' | $(RUST_BIN)

test-c: $(C_BIN)
	@echo "Testing C implementation..."
	@echo '{"azure.authenticated": "true", "azure.role": "admin"}' | $(C_BIN)
	@echo '{"azure.authenticated": "false", "azure.role": "admin"}' | $(C_BIN)

# Clean targets
.PHONY: clean clean-go clean-python clean-javascript clean-rust clean-c
clean: clean-go clean-python clean-javascript clean-rust clean-c
	rm -rf $(BINDIR)

clean-go:
	cd go && go clean
	rm -f $(GO_BIN)

clean-python:
	rm -f $(PYTHON_BIN)
	find python -name "*.pyc" -delete
	find python -name "__pycache__" -type d -exec rm -rf {} + 2>/dev/null || true

clean-javascript:
	rm -f $(JS_BIN)
	cd javascript && rm -rf node_modules

clean-rust:
	cd rust && cargo clean
	rm -f $(RUST_BIN)

clean-c:
	rm -f $(C_BIN)

# Install dependencies
.PHONY: deps deps-rust deps-javascript
deps: deps-rust deps-javascript

deps-rust:
	cd rust && cargo fetch

deps-javascript:
	cd javascript && npm install

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all          - Build all implementations"
	@echo "  go           - Build Go implementation"
	@echo "  python       - Build Python implementation"
	@echo "  javascript   - Build JavaScript implementation"
	@echo "  rust         - Build Rust implementation"
	@echo "  c            - Build C implementation"
	@echo "  test         - Test all implementations"
	@echo "  test-<lang>  - Test specific implementation"
	@echo "  clean        - Clean all build artifacts"
	@echo "  clean-<lang> - Clean specific implementation"
	@echo "  deps         - Install dependencies"
	@echo "  help         - Show this help message"
