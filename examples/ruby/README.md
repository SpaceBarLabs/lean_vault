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

The example demonstrates several ways to use Lean Vault in your Ruby applications:

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

### 3. Key Rotation (Security Best Practice)

The example includes a key rotation demonstration that shows how to safely rotate API keys in a production environment:

```bash
# Basic usage
ruby key_rotation_example.rb

# With debug output
ruby key_rotation_example.rb --debug

# Rotate a specific key
ruby key_rotation_example.rb --key my-custom-key

# Customize retry attempts
ruby key_rotation_example.rb --retries 5
```

The key rotation example demonstrates:
- Safe key rotation with zero downtime
- Verification of both old and new keys
- Error handling and retry logic
- Best practices for key management

## Running the Examples

1. Basic API test:
```bash
ruby test_key.rb
```

2. Key rotation example:
```bash
ruby key_rotation_example.rb
```

Debug mode available for both:
```bash
ruby test_key.rb --debug
ruby key_rotation_example.rb --debug
```

## Example Output

Success case (basic test):
```
✓ Successfully loaded key from Lean Vault
✓ OpenRouter API connection test successful
Response: "Hello! I'm Claude, an AI assistant. How can I help you today?"
```

Success case (key rotation):
```
DEBUG: Starting key rotation example...
DEBUG: Testing current key...
✓ API key test successful
DEBUG: Rotating key...
Rotating API key 'ruby-app-key'...
✓ Key rotated successfully
DEBUG: Testing new key...
✓ API key test successful
✓ Key rotation completed successfully!
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
5. Safe key rotation strategies

Key features:
- Constants-based secret management (similar to dotenv)
- Clean error handling
- Debug logging
- Thread-safe key loading
- Zero-downtime key rotation

## Key Rotation Best Practices

When rotating keys in a production environment:

1. **Verify Current Key**
   - Always test the current key before rotation
   - Ensure your application has proper access

2. **Safe Rotation**
   - Use Lean Vault's built-in rotation command
   - The command handles both creating new key and revoking old
   - Automatically updates vault storage

3. **Verify New Key**
   - Test the new key immediately after rotation
   - Ensure all required permissions are preserved

4. **Error Handling**
   - Implement proper retry logic
   - Log rotation events for audit purposes
   - Have a rollback strategy

5. **Application Updates**
   - Clear any cached keys in your application
   - Reload keys from the vault after rotation
   - Consider implementing automatic key refresh

## Troubleshooting

If you encounter issues:

1. **Key Access Issues**
   - Check key exists: `lean_vault list`
   - Verify permissions: `ls -l ~/.lean_vault/`
   - Try manual key retrieval: `lean_vault get <key-name>`

2. **Rotation Issues**
   - Enable debug mode: `--debug`
   - Check OpenRouter API status
   - Verify provisioning key permissions

3. **Integration Issues**
   - Ensure Lean Vault is in PATH
   - Check Ruby version compatibility
   - Verify all gems are installed

For more help, see the main [Lean Vault documentation](../../README.md).

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