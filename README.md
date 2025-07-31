# isauthz-processor

A multi-language implementation of an authorization processor that reads JSON from stdin, processes authorization requests, and outputs "authorized" or "unauthorized".

## Project Structure

```
isauthz-processor/
├── Makefile                 # Build system for all implementations
├── README.md               # This file
├── bin/                    # Output directory for built binaries
├── go/                     # Go implementation
│   ├── main.go
│   └── go.mod
├── python/                 # Python implementation
│   └── main.py
├── javascript/             # JavaScript (Node.js) implementation
│   ├── main.js
│   └── package.json
├── rust/                   # Rust implementation
│   ├── Cargo.toml
│   └── src/
│       └── main.rs
└── c/                      # C implementation
    └── main.c
```

## Authorization Logic

All implementations follow the same authorization logic:

1. **Authentication Check**: User must have `azure.authenticated` set to "true"
2. **Admin Override**: Users with `azure.role` = "admin" are always authorized
3. **Engineering Developers**: Users in Engineering department with "developers" in their groups
4. **Example.com Users**: Users with @example.com email and "user" role
5. **Default Deny**: All other cases result in "unauthorized"

## Building

### Prerequisites

- **Go**: Go 1.20+ 
- **Python**: Python 3.6+
- **Node.js**: Node.js 14+
- **Rust**: Rust 1.60+
- **C**: GCC or compatible C compiler

### Build All Implementations

```bash
make all
```

### Build Individual Implementations

```bash
make go          # Build Go version
make python      # Build Python version
make javascript  # Build JavaScript version
make rust        # Build Rust version
make c           # Build C version
```

### Install Dependencies

```bash
make deps        # Install Rust and JavaScript dependencies
```

## Usage

All implementations read JSON from stdin line by line and output authorization results:

```bash
# Example with Go implementation
echo '{"azure.authenticated": "true", "azure.role": "admin"}' | ./bin/isauthz-processor-go

# Example with multiple inputs
cat <<EOF | ./bin/isauthz-processor-python
{"azure.authenticated": "true", "azure.role": "admin"}
{"azure.authenticated": "true", "azure.role": "user", "azure.email": "john@example.com"}
{"azure.authenticated": "false", "azure.role": "admin"}
EOF
```

## Testing

### Test All Implementations

```bash
make test
```

### Test Individual Implementations

```bash
make test-go
make test-python
make test-javascript
make test-rust
make test-c
```

## Input Format

Each line should be a JSON object with string key-value pairs:

```json
{
  "azure.authenticated": "true",
  "azure.role": "admin",
  "azure.department": "Engineering",
  "azure.groups": "developers,users",
  "azure.email": "user@example.com"
}
```

## Output

Each implementation outputs one of:
- `authorized`
- `unauthorized`

## Example Test Cases

```bash
# Admin user (should be authorized)
echo '{"azure.authenticated": "true", "azure.role": "admin"}' | ./bin/isauthz-processor-go

# Engineering developer (should be authorized)
echo '{"azure.authenticated": "true", "azure.department": "Engineering", "azure.groups": "developers"}' | ./bin/isauthz-processor-go

# Example.com user (should be authorized)
echo '{"azure.authenticated": "true", "azure.role": "user", "azure.email": "john@example.com"}' | ./bin/isauthz-processor-go

# Unauthenticated user (should be unauthorized)
echo '{"azure.authenticated": "false", "azure.role": "admin"}' | ./bin/isauthz-processor-go

# Invalid JSON (should be unauthorized)
echo 'invalid json' | ./bin/isauthz-processor-go
```

## Cleaning

```bash
make clean        # Clean all build artifacts
make clean-go     # Clean Go artifacts
make clean-rust   # Clean Rust artifacts
# etc.
```
