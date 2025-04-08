# Rotate Command Implementation Plan

## Overview

The `rotate` command provides a secure way to rotate OpenRouter API keys by creating a new key and revoking the old one. This document outlines the implementation details and considerations.

## Command Specification

```bash
lean_vault rotate <key-name>
```

### Purpose
- Create a new API key to replace an existing one
- Securely update the vault with the new key
- Revoke the old key for security
- Handle partial success scenarios gracefully

## Implementation Components

### 1. Vault Package Updates

Added new method to handle key updates:

```go
// UpdateSecret updates an existing secret in the vault
func (v *Vault) UpdateSecret(name, value, id string) error {
    vaultData, masterKey, err := v.load()
    if err != nil {
        return err
    }

    if _, exists := vaultData.Secrets[name]; !exists {
        return fmt.Errorf("secret %s not found", name)
    }

    vaultData.Secrets[name] = SecretEntry{
        Value: value,
        ID:    id,
    }

    return v.save(vaultData, masterKey)
}
```

### 2. Rotate Command Implementation

The rotation process follows these steps:

1. **Validation**
   - Verify the key exists in the vault
   - Get the current key's ID
   - Retrieve the main provisioning key

2. **Key Creation**
   - Create new key via OpenRouter API
   - Store response containing new key and ID

3. **Vault Update**
   - Update vault with new key details
   - Ensure atomic update operation

4. **Key Revocation**
   - Revoke old key via OpenRouter API
   - Handle revocation failures gracefully

### 3. Error Handling

The command handles several error scenarios:

1. **Pre-rotation Errors** (Fatal)
   - Key not found in vault
   - Failed to get provisioning key
   - Failed to access vault

2. **Creation Errors** (Fatal)
   - API errors during key creation
   - Network issues
   - Rate limiting/quota issues

3. **Update Errors** (Fatal)
   - Vault access/encryption errors
   - File system errors

4. **Revocation Errors** (Warning)
   - Failed to revoke old key
   - Network issues during revocation
   - API errors during revocation

### 4. User Feedback

The command provides clear feedback through stderr:

```
Rotating API key 'my-key'...
Creating new key...
Updating vault with new key...
Revoking old key...
✓ API key 'my-key' rotated successfully!
```

In case of revocation failure:
```
⚠️  Warning: Failed to revoke old key: <error>
The new key has been stored successfully, but the old key may still be active.
You may want to try revoking it manually or contact OpenRouter support.
```

## Testing Strategy

1. **Unit Tests**
   - Test UpdateSecret method
   - Test key rotation workflow
   - Test error handling scenarios

2. **Integration Tests**
   - Test with OpenRouter API
   - Test vault file operations
   - Test partial success scenarios

3. **Error Cases**
   - Test network failures
   - Test API errors
   - Test file system errors

## Security Considerations

1. **Atomic Operations**
   - Ensure vault updates are atomic
   - Prevent partial updates
   - Maintain vault integrity

2. **Key Management**
   - Secure handling of both old and new keys
   - Proper cleanup of old keys
   - Clear error messaging for security-related issues

3. **File Operations**
   - Maintain proper file permissions
   - Secure handling of sensitive data
   - Clean up on failure

## Usage Examples

Basic usage:
```bash
lean_vault rotate my-api-key
```

With debug output:
```bash
LEAN_VAULT_DEBUG=1 lean_vault rotate my-api-key
```

## Future Enhancements

1. **Additional Features**
   - Add `--force` flag to skip revocation
   - Add `--dry-run` option
   - Support for key rotation schedules

2. **Improvements**
   - Add retry logic for revocation
   - Add confirmation prompt for production keys
   - Support for bulk key rotation

## References

- [Lean Vault Specification](lean_vault_spec.md)
- [OpenRouter API Documentation](https://openrouter.ai/docs) 