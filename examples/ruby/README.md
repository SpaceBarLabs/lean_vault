# Ruby Example - Lean Vault Integration

This example demonstrates how to integrate Lean Vault into a Ruby application for secure OpenRouter API key management, using a pattern similar to how `dotenv` works with environment variables.

## Prerequisites

- Ruby 3.0 or higher
- Bundler
- Lean Vault CLI tool installed and in your PATH (see [root README](../../README.md) for installation instructions)
- At least one API key stored in your vault

## Setup

1. Install Ruby dependencies:
```bash
bundle install
```

2. Initialize Lean Vault and add a test key:
```bash
lean_vault init
lean_vault add ruby-app-key
```

## Usage

The example demonstrates two ways to use Lean Vault in your Ruby applications:

### 1. Using the LeanVault Module (Recommended)

Similar to how `dotenv` loads environment variables, the `LeanVault` module loads secrets into constants:

```ruby
require_relative 'lean_vault'

# Load specific keys into constants
LeanVault.load('ruby-app-key', 'another-key')

# Use the constants in your code
puts "Using API key: #{LeanVault::RUBY_APP_KEY}"
```

### 2. Direct Command Execution

For simpler use cases, you can execute the lean_vault command directly:

```ruby
api_key = `lean_vault get ruby-app-key`.strip
raise "Failed to get key" unless $?.success?
```

## Running the Example

Basic usage:
```bash
ruby test_key.rb
```

With debug output:
```bash
ruby test_key.rb --debug
```

## Example Output

Success case:
```
✓ Successfully loaded key from Lean Vault
✓ OpenRouter API connection test successful
Response: "Hello! I'm Claude, an AI assistant. How can I help you today?"
```

Error case:
```
✗ Error: Failed to load key 'invalid-key': Key not found
Please check:
- Key exists in vault (lean_vault list)
- Lean Vault is properly configured
- You have necessary permissions
```

## Implementation Details

The example demonstrates:
1. A reusable `LeanVault` module for loading secrets as constants
2. OpenRouter API integration
3. Error handling and logging
4. Best practices for key management

Key features:
- Constants-based secret management (similar to dotenv)
- Clean error handling
- Debug logging
- Thread-safe key loading

## Troubleshooting

1. **lean_vault: command not found**
   - Ensure lean_vault is installed: `which lean_vault`
   - Check if it's in your PATH
   - Try running with the full path to the binary

2. **Key not found**
   - Check if key exists: `lean_vault list`
   - Verify the key name matches exactly

3. **Constant already defined warning**
   - This can happen if you load the same key multiple times
   - Consider loading keys only once at application startup

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