# Lean Vault

A secure CLI tool for managing OpenRouter API keys.

## Features

- Secure storage of API keys using AES-256-GCM encryption
- OpenRouter API key provisioning and management
- Usage tracking and monitoring
- Key rotation capabilities
- Simple CLI interface
- Multi-language integrations (Ruby, TypeScript, Python, Go)

## Installation

*Coming soon - The tool will be distributed as a static binary.*

For now, to build from source:

```bash
# Clone the repository
git clone https://github.com/spacebarlabs/lean_vault.git
cd lean_vault

# Build the binary (choose one of these options):

# Option 1: Build and use locally
go build -o bin/lean_vault ./cmd/lean_vault
# The binary will be created as bin/lean_vault in the project directory
# You can now use it with bin/lean_vault <command>

# Option 2: Install to your PATH
sudo cp bin/lean_vault /usr/local/bin/lean_vault
# Now you can use it from anywhere with: lean_vault <command>

# Option 3: Install to $GOPATH/bin (if $GOPATH/bin is in your PATH)
go install ./cmd/lean_vault
# Now you can use it from anywhere with: lean_vault <command>
```

## Quick Start

1. Initialize the vault:
```bash
# If installed to PATH:
lean_vault init

# If using locally:
bin/lean_vault init
```

2. Add a new API key:
```bash
# If installed to PATH:
lean_vault add my-key

# If using locally:
bin/lean_vault add my-key
```

3. Use the key in your applications:
```bash
# If installed to PATH:
export OPENROUTER_API_KEY=$(lean_vault get my-key)

# If using locally:
export OPENROUTER_API_KEY=$(bin/lean_vault get my-key)
```

## Language Integrations

Lean Vault provides example integrations for multiple programming languages:

- **Ruby**: Constants-based secret management with dotenv-like interface
- **TypeScript**: Strongly-typed integration with async/await support
- **Python**: Context manager-based integration with type hints
- **Go**: Concurrency-safe integration with channel support

See [Language Examples Guide](docs/language_examples.md) for detailed implementation guides and best practices.

## Security

- Uses AES-256-GCM for encryption
- PBKDF2 key derivation
- Secure file permissions
- No plaintext storage of secrets

## Debugging

If you encounter issues, you can enable debug mode by setting the `LEAN_VAULT_DEBUG` environment variable:

```bash
LEAN_VAULT_DEBUG=1 bin/lean_vault add my-key
```

This will show:
- API request details
- Response status and body
- Detailed error messages

## Documentation

- [Tutorial](docs/TUTORIAL.md) - Detailed usage instructions
- [Specification](docs/lean_vault_spec.md) - Implementation details
- [Language Examples](docs/language_examples.md) - Multi-language integration guides

## Development

### Project Structure

```
lean_vault/
├── bin/             # Build outputs
├── cmd/
│   └── lean_vault/  # Main CLI application
├── pkg/
│   ├── crypto/      # Encryption utilities
│   ├── vault/       # Vault management
│   └── api/         # OpenRouter API client
├── examples/        # Language integration examples
│   ├── ruby/
│   ├── typescript/
│   ├── python/
│   └── golang/
└── docs/           # Documentation
```

### Building from Source

```bash
# Development build
go build -o bin/lean_vault ./cmd/lean_vault

# Production build (static binary)
CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/lean_vault ./cmd/lean_vault
```

### Running Tests

```bash
go test ./...
```

## License

*TBD*

## Contributing

*Coming soon*