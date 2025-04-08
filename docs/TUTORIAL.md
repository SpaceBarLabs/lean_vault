# Lean Vault Tutorial

Welcome to Lean Vault! This tutorial will guide you through using this secure CLI tool for managing your API keys, with a focus on OpenRouter API keys.

## Table of Contents
- [Installation](#installation)
- [Getting Started](#getting-started)
- [Basic Commands](#basic-commands)
- [Advanced Usage](#advanced-usage)
- [Best Practices](#best-practices)

## Installation

*Coming soon - The tool will be available as a static binary for easy installation.*

## Getting Started

### Initial Setup

To start using Lean Vault, you'll need to initialize it with your OpenRouter provisioning key:

```bash
lean_vault init
```

This will:
1. Create a secure vault in your home directory (`~/.lean_vault/`)
2. Prompt you for your OpenRouter provisioning key
3. Set up encryption for your secrets

### Basic Commands

#### 1. Adding a New API Key

To provision and store a new OpenRouter API key:

```bash
lean_vault add my-new-key
```

This will:
- Create a new API key through OpenRouter
- Store it securely in your vault
- Label it as "my-new-key" for future reference

#### 2. Retrieving a Key

To get the value of a stored key:

```bash
lean_vault get my-new-key
```

This command outputs only the key value, making it perfect for use in scripts:

```bash
export OPENROUTER_API_KEY=$(lean_vault get my-new-key)
```

#### 3. Listing Your Keys

To see all stored keys:

```bash
lean_vault list
```

This will show all your stored keys (except the main provisioning key).

#### 4. Removing a Key

To remove and revoke a key:

```bash
lean_vault remove my-new-key
```

This will:
- Revoke the key on OpenRouter's side
- Remove it from your local vault

## Advanced Usage

### Key Rotation

To rotate an existing key (create new + revoke old):

```bash
lean_vault rotate my-new-key
```

This is useful for:
- Regular security maintenance
- Responding to potential key exposure
- Updating key permissions or limits

### Usage Tracking

To monitor your API key usage:

```bash
lean_vault usage
```

This provides a detailed view of:
- Current usage amounts
- Spending limits
- Rate limits
- Key status

Example output:
```
Key Name      | Usage ($) | Limit ($) | Limit Remaining ($) | Rate Limit | Status
--------------|-----------|-----------|---------------------|------------|--------
MY_KEY_1      | 3.25      | 10.00     | 6.75                | 60 req/min | OK
ANOTHER_KEY   | 0.50      | 5.00      | 4.50                | 60 req/min | OK
```

## Best Practices

1. **Vault Security**
   - Keep your `~/.lean_vault/.secret_vault.key` file secure
   - Add `.lean_vault/` to your global gitignore
   - Never share or expose your master key file

2. **Key Management**
   - Use descriptive names for your keys
   - Regularly rotate keys for security
   - Monitor usage to prevent unexpected charges

3. **Integration Tips**
   - Use `lean_vault get` in scripts and CI/CD pipelines
   - Consider rotating keys after major deployments
   - Set up regular usage monitoring

## Troubleshooting

Common error scenarios and solutions:

1. **Initialization Issues**
   ```
   Error: Vault already exists
   ```
   - Solution: Remove existing vault or use a different location

2. **Access Problems**
   ```
   Error: Cannot read vault
   ```
   - Check file permissions
   - Verify `.secret_vault.key` exists
   - Run `lean_vault init` if starting fresh

3. **API Errors**
   ```
   Error: Failed to provision key
   ```
   - Verify your main provisioning key
   - Check OpenRouter service status
   - Ensure you haven't hit account limits

## Need Help?

- Check the error messages (they're designed to be helpful!)
- Ensure you're using the latest version
- Review this tutorial's relevant section

Remember: Lean Vault is designed to be simple and secure. If something seems wrong, trust the error messages and don't force operations that might compromise security. 