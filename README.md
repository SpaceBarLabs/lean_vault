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

# Install to $GOPATH/bin (recommended)
go install ./cmd/lean_vault

# Verify $GOPATH/bin is in your PATH
echo $PATH | grep -q "$GOPATH/bin" || echo "Warning: $GOPATH/bin is not in your PATH"
# If not in PATH, add it to your shell configuration:
# For bash: echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
# For zsh:  echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc

# Verify installation
lean_vault version
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

## Command Lifecycle

Here's a walkthrough of how to use Lean Vault in your daily workflow:

1. **Initialize Your Vault** (First-time setup)
   ```bash
   lean_vault init
   ```
   This creates your encrypted vault and sets up the necessary configuration.

2. **Add Your First API Key**
   ```bash
   lean_vault add my-production-key
   ```
   This will prompt you to enter your OpenRouter API key securely.

3. **List Your Keys**
   ```bash
   lean_vault list
   ```
   View all your stored API keys and their status.

4. **Use a Key in Your Application**
   ```bash
   export OPENROUTER_API_KEY=$(lean_vault get my-production-key)
   ```
   Retrieve a key for use in your application or scripts.

5. **Rotate a Key** (When needed)
   ```bash
   lean_vault rotate my-production-key
   ```
   This creates a new key and automatically revokes the old one.

6. **Remove a Key**
   ```bash
   lean_vault remove my-production-key
   ```
   This removes the key and attempts to revoke it on OpenRouter.
   
   If you need to force remove without revocation:
   ```bash
   lean_vault remove my-production-key --force
   ```

7. **Check Version**
   ```bash
   lean_vault version
   ```
   Verify which version of Lean Vault you're running.

8. **Monitor Usage** (Coming Soon)
   ```bash
   lean_vault usage
   ```
   Track API usage across all your keys.

## Available Commands

- `init` - Initialize the vault
- `add <key-name>` - Add a new OpenRouter API key
- `get <key-name>` - Retrieve a stored key
- `list` - List all stored keys
- `remove <key-name> [--force]` - Remove and revoke a key (use --force to skip revocation)
- `rotate <key-name>` - Rotate a key (create new + revoke old)
- `usage` - Display usage information for all keys (coming soon)
- `version` - Show version information

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

For development and testing, you can build the binary locally:

```bash
# Development build (with debug information)
go build -o bin/lean_vault ./cmd/lean_vault

# Production build (static binary, smaller size)
CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/lean_vault ./cmd/lean_vault
```

The binary will be created in the `bin/` directory. You can then run it directly:
```bash
./bin/lean_vault <command>
```

### Running Tests

```bash
go test ./...
```

## License

*TBD*

## Contributing

*Coming soon*