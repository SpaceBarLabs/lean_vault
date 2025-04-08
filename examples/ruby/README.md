# Ruby Example - Lean Vault Integration

This example demonstrates how to integrate Lean Vault into a Ruby application for secure OpenRouter API key management.

## Prerequisites

- Ruby 3.0 or higher
- Bundler
- Lean Vault CLI tool (see below)
- At least one API key stored in your vault

### Installing Lean Vault CLI

Before running this example, ensure you have the Lean Vault CLI tool installed and accessible in your PATH. You have three options:

1. **Build and use locally** (from project root):
```bash
go build -o lean_vault ./cmd/cli
# Then either:
export PATH="$PATH:$(pwd)"  # Temporary: only for current session
# Or move it to a PATH directory:
sudo mv lean_vault /usr/local/bin/
```

2. **Install directly to PATH**:
```bash
go build -o /usr/local/bin/lean_vault ./cmd/cli
```

3. **Install to $GOPATH/bin** (if in your PATH):
```bash
go install ./cmd/cli
```

Verify the installation:
```bash
lean_vault --version  # Should show version information
```

## Setup

1. Install Ruby dependencies:
```bash
bundle install
```

2. Create a `.env` file (optional):
```bash
LEAN_VAULT_KEY_NAME=my-key  # The name of your stored key
DEBUG=true                  # Enable debug output
```

3. Initialize Lean Vault and add a test key:
```bash
lean_vault init
lean_vault add ruby-app-key
```

## Usage

Basic usage:
```bash
ruby test_key.rb
```

With specific key name:
```bash
ruby test_key.rb --key-name=my-custom-key
```

With debug output:
```bash
DEBUG=true ruby test_key.rb
```

## Example Output

Success case:
```
✓ Successfully retrieved key from Lean Vault
✓ OpenRouter API connection test successful
Response: "Hello! I'm Claude, an AI assistant. How can I help you today?"
```

Error case:
```
✗ Error retrieving key: Key 'invalid-key' not found
Please check:
- Key exists in vault (lean_vault list)
- Lean Vault is properly configured
- You have necessary permissions
```

## Implementation Details

The example demonstrates:
1. Secure key retrieval from Lean Vault
2. OpenRouter API integration
3. Error handling and logging
4. Configuration management

Key features:
- Uses environment variables for configuration
- Implements proper error handling
- Includes debug logging
- Shows best practices for key management

## Troubleshooting

1. **lean_vault: command not found**
   - Follow the "Installing Lean Vault CLI" section above
   - Check if lean_vault is in your PATH: `which lean_vault`
   - Try running with the full path to the binary

2. **Key not found**
   - Check if key exists: `lean_vault list`
   - Verify key name in `.env` or command line argument

3. **API Connection Failed**
   - Verify internet connection
   - Check if key is valid
   - Enable debug mode for more information

## Security Notes

- Never commit `.env` files
- Don't log or display full API keys
- Handle errors without exposing sensitive information
- Use environment variables for configuration

## Next Steps

1. Modify the example for your use case
2. Integrate into your application
3. Add additional error handling as needed
4. Implement key rotation if required 