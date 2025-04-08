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
go install github.com/spacebarlabs/lean_vault@latest
```

## Quick Start

1. Initialize the vault:
```bash
lean_vault init
```

2. Add a new API key:
```bash
lean_vault add my-key
```

3. Use the key in your applications:
```bash
export OPENROUTER_API_KEY=$(lean_vault get my-key)
```

## Security

- Uses AES-256-GCM for encryption
- PBKDF2 key derivation
- Secure file permissions
- No plaintext storage of secrets

## Documentation

For detailed usage instructions, see [TUTORIAL.md](TUTORIAL.md).
For implementation details, see [lean_vault_spec.md](lean_vault_spec.md).

## Development

### Project Structure

```
lean_vault/
├── cmd/
│   └── lean_vault/     # Main CLI application
├── pkg/
│   ├── crypto/        # Encryption utilities
│   ├── vault/         # Vault management
│   └── api/           # OpenRouter API client
```

### Building

```bash
go build ./cmd/lean_vault
```

## License

*TBD*

## Contributing

*Coming soon*
