# Lean Vault Examples

This directory contains example applications demonstrating how to use Lean Vault in various programming languages. Each example shows how to integrate Lean Vault into your application for secure OpenRouter API key management.

## Available Examples

### Ruby (`/ruby`)
A simple Ruby application demonstrating:
- Retrieving keys from Lean Vault
- Making API calls to OpenRouter
- Error handling and logging
- Best practices for key management

## Common Requirements

All examples assume you have:
1. Lean Vault installed and configured
2. At least one API key stored in your vault
3. Basic familiarity with the language of the example

## Example Structure

Each language example follows a similar pattern:
1. Project setup and dependency management
2. Key retrieval from Lean Vault
3. Simple OpenRouter API integration
4. Error handling and logging
5. Documentation and usage instructions

## Running Examples

Each example directory contains its own README with specific setup and running instructions. Generally, you'll need to:

1. Install language-specific dependencies
2. Ensure Lean Vault is in your PATH
3. Have a valid key stored in your vault
4. Follow the example-specific README instructions

## Contributing New Examples

We welcome contributions for additional language examples! Please ensure your example:

1. Follows the language's best practices and conventions
2. Includes clear documentation and setup instructions
3. Demonstrates proper error handling
4. Uses modern dependency management
5. Includes a README.md with:
   - Prerequisites
   - Installation steps
   - Usage instructions
   - Troubleshooting guide

## Security Notes

- Never commit API keys or sensitive information
- Always use environment variables or Lean Vault for key management
- Follow security best practices for your chosen language
- Handle errors appropriately to avoid exposing sensitive information

## Support

If you encounter issues with any example:
1. Check the example's README for troubleshooting steps
2. Ensure Lean Vault is properly configured
3. Verify your API key is valid
4. Check the main Lean Vault documentation for additional help 