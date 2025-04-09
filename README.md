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

## Language Support

Currently, Lean Vault provides a Ruby integration example that demonstrates how to use the CLI tool in a Ruby application. This serves as a reference implementation for other languages.

### Ruby Integration

The Ruby integration provides a simple interface for managing API keys in your Ruby applications:

```ruby
require 'lean_vault'

# Initialize the vault client
vault = LeanVault::Client.new

# Get an API key
api_key = vault.get_key('my-production-key')

# Use the key in your application
client = OpenRouter::Client.new(api_key: api_key)
```

For more details and examples, see the [Ruby Integration Guide](examples/ruby/README.md).

### Future Language Support

We plan to add official language integrations in the future, including:
- TypeScript/JavaScript
- Python
- Go

These will be released as separate packages with proper versioning and documentation. For now, you can use the CLI tool directly in any language by executing the binary and parsing its output.

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

Lean Vault is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

The MIT License is a permissive license that:
- Allows commercial use
- Allows modification
- Allows distribution
- Allows private use
- Includes a limitation of liability
- Includes a warranty shield

## Contributing

We welcome contributions to Lean Vault! Here's how you can help:

1. **Reporting Issues**
   - Check existing issues before creating a new one
   - Include steps to reproduce, expected behavior, and actual behavior
   - Include system information (OS, Go version, etc.)

2. **Submitting Pull Requests**
   - Fork the repository
   - Create a feature branch (`git checkout -b feature/amazing-feature`)
   - Commit your changes (`git commit -m 'Add amazing feature'`)
   - Push to the branch (`git push origin feature/amazing-feature`)
   - Open a Pull Request

3. **Development Guidelines**
   - Follow Go best practices and idioms
   - Write tests for new features
   - Update documentation for significant changes
   - Keep commits focused and atomic
   - Use meaningful commit messages

4. **Code Style**
   - Run `gofmt` before committing
   - Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
   - Use `golint` and `go vet` to check your code

5. **Testing**
   - Write unit tests for new features
   - Ensure all tests pass (`go test ./...`)
   - Add integration tests for significant changes

6. **Documentation**
   - Update README.md for user-facing changes
   - Add or update docstrings for new functions
   - Keep the Command Lifecycle section up to date

7. **Security**
   - Do not commit sensitive information
   - Report security vulnerabilities privately to security@spacebarlabs.com
   - Follow secure coding practices

8. **Community**
   - Be respectful and inclusive
   - Help review pull requests
   - Answer questions in issues
   - Share your use cases and experiences

For more detailed information, see our [Contributing Guide](CONTRIBUTING.md).