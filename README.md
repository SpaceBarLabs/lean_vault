# Lean Vault

A secure CLI tool for managing OpenRouter API keys.

## Features

- Secure storage of API keys using AES-256-GCM encryption
- OpenRouter API key provisioning and management
- Usage tracking and monitoring
- Key rotation capabilities
- Simple CLI interface

## Installation

*Coming soon - The tool will be distributed as a static binary.*

For now, to build from source:

```bash
# Clone the repository
git clone https://github.com/spacebarlabs/lean_vault.git
cd lean_vault

# Build the binary
go build -o lean_vault ./cmd/cli

# Optional: Install to your $GOPATH/bin
go install ./cmd/cli
```

## Quick Start

1. Initialize the vault:
```bash
./lean_vault init
```

2. Add a new API key:
```bash
./lean_vault add my-key
```

3. Use the key in your applications:
```bash
export OPENROUTER_API_KEY=$(./lean_vault get my-key)
```

## Security

- Uses AES-256-GCM for encryption
- PBKDF2 key derivation
- Secure file permissions
- No plaintext storage of secrets

## Debugging

If you encounter issues, you can enable debug mode by setting the `LEAN_VAULT_DEBUG` environment variable:

```bash
LEAN_VAULT_DEBUG=1 ./lean_vault add my-key
```

This will show:
- API request details
- Response status and body
- Detailed error messages

## Documentation

For detailed usage instructions, see [TUTORIAL.md](TUTORIAL.md).
For implementation details, see [lean_vault_spec.md](lean_vault_spec.md).

## Development

### Project Structure

```
lean_vault/
├── cmd/
│   └── cli/          # Main CLI application
├── pkg/
│   ├── crypto/       # Encryption utilities
│   ├── vault/        # Vault management
│   └── api/          # OpenRouter API client
```

### Building from Source

```bash
# Development build
go build -o lean_vault ./cmd/cli

# Production build (static binary)
CGO_ENABLED=0 go build -ldflags="-s -w" -o lean_vault ./cmd/cli
```

### Running Tests

```bash
go test ./...
```

## License

*TBD*

## Contributing

*Coming soon*
