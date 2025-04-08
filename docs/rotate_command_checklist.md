# Rotate Command Implementation Checklist

## Core Implementation

### Vault Package Updates
- [ ] Add `UpdateSecret` method to vault package
  - [ ] Implement load/decrypt functionality
  - [ ] Add existence check for secret
  - [ ] Implement update logic
  - [ ] Add save/encrypt functionality
- [ ] Add unit tests for `UpdateSecret`
  - [ ] Test successful update
  - [ ] Test non-existent key
  - [ ] Test encryption/decryption during update
  - [ ] Test file permission preservation

### Rotate Command Implementation
- [ ] Create `pkg/commands/rotate.go`
  - [ ] Implement main `Rotate` function
  - [ ] Add debug mode support
  - [ ] Implement progress feedback
  - [ ] Add error handling
- [ ] Update `cmd/lean_vault/main.go`
  - [ ] Add rotate command case
  - [ ] Add usage information
  - [ ] Add example commands
  - [ ] Wire up error handling

### API Integration
- [ ] Verify OpenRouter API endpoints
  - [ ] Test key creation endpoint
  - [ ] Test key revocation endpoint
  - [ ] Document rate limits
- [ ] Implement API calls
  - [ ] Add new key creation
  - [ ] Add old key revocation
  - [ ] Add error handling for API responses

## Testing

### Unit Tests
- [ ] Vault Package Tests
  - [ ] Test `UpdateSecret` success cases
  - [ ] Test `UpdateSecret` error cases
  - [ ] Test file permission preservation
  - [ ] Test atomic updates
- [ ] Command Tests
  - [ ] Test rotation workflow
  - [ ] Test error handling
  - [ ] Test user feedback
  - [ ] Test debug mode

### Integration Tests
- [ ] API Integration Tests
  - [ ] Mock OpenRouter API responses
  - [ ] Test successful rotation
  - [ ] Test failed key creation
  - [ ] Test failed key revocation
- [ ] File System Tests
  - [ ] Test vault file updates
  - [ ] Test permission preservation
  - [ ] Test concurrent access
  - [ ] Test cleanup on failure

### Error Case Tests
- [ ] Network Error Tests
  - [ ] Test timeout handling
  - [ ] Test connection failures
  - [ ] Test partial responses
- [ ] API Error Tests
  - [ ] Test rate limiting
  - [ ] Test invalid responses
  - [ ] Test authorization failures
- [ ] File System Error Tests
  - [ ] Test permission denied
  - [ ] Test disk full
  - [ ] Test concurrent access conflicts

## Security Verification

### Atomic Operations
- [ ] Verify vault update atomicity
  - [ ] Test partial updates
  - [ ] Test concurrent access
  - [ ] Test recovery from failures
- [ ] Verify file integrity
  - [ ] Test corruption detection
  - [ ] Test backup/recovery
  - [ ] Test permission preservation

### Key Management
- [ ] Verify secure key handling
  - [ ] Test key encryption
  - [ ] Test key storage
  - [ ] Test key cleanup
- [ ] Verify revocation process
  - [ ] Test successful revocation
  - [ ] Test failed revocation handling
  - [ ] Test partial success scenarios

### File Operations
- [ ] Verify file permissions
  - [ ] Test permission settings
  - [ ] Test permission inheritance
  - [ ] Test permission recovery
- [ ] Verify secure cleanup
  - [ ] Test temporary file cleanup
  - [ ] Test error state cleanup
  - [ ] Test partial operation cleanup

## Documentation

### Code Documentation
- [ ] Add function documentation
  - [ ] Document `UpdateSecret`
  - [ ] Document `Rotate`
  - [ ] Document error types
  - [ ] Document debug flags
- [ ] Add example usage
  - [ ] Basic usage examples
  - [ ] Error handling examples
  - [ ] Debug mode examples

### User Documentation
- [ ] Update README.md
  - [ ] Add rotate command
  - [ ] Add usage examples
  - [ ] Add error solutions
- [ ] Update TUTORIAL.md
  - [ ] Add rotation guide
  - [ ] Add troubleshooting
  - [ ] Add best practices

## Future Enhancements Preparation

### Feature Foundations
- [ ] Prepare for `--force` flag
  - [ ] Design flag handling
  - [ ] Plan implementation
  - [ ] Document usage
- [ ] Prepare for `--dry-run`
  - [ ] Design simulation logic
  - [ ] Plan implementation
  - [ ] Document usage
- [ ] Prepare for scheduled rotation
  - [ ] Design scheduling system
  - [ ] Plan implementation
  - [ ] Document usage

### Improvement Foundations
- [ ] Design retry logic
  - [ ] Define retry strategies
  - [ ] Plan implementation
  - [ ] Document behavior
- [ ] Design confirmation prompts
  - [ ] Define prompt triggers
  - [ ] Plan implementation
  - [ ] Document behavior
- [ ] Design bulk rotation
  - [ ] Define batch processing
  - [ ] Plan implementation
  - [ ] Document usage

## Final Verification

### Quality Assurance
- [ ] Run all tests
  - [ ] Unit tests
  - [ ] Integration tests
  - [ ] Security tests
- [ ] Verify documentation
  - [ ] Check accuracy
  - [ ] Test examples
  - [ ] Verify links
- [ ] Performance testing
  - [ ] Test with large vaults
  - [ ] Test concurrent operations
  - [ ] Test error recovery

### Security Review
- [ ] Audit code
  - [ ] Check for vulnerabilities
  - [ ] Verify secure practices
  - [ ] Review error handling
- [ ] Test edge cases
  - [ ] Test boundary conditions
  - [ ] Test error scenarios
  - [ ] Test recovery procedures

### User Experience
- [ ] Verify error messages
  - [ ] Check clarity
  - [ ] Test helpfulness
  - [ ] Verify accuracy
- [ ] Test command flow
  - [ ] Check feedback
  - [ ] Verify progress indication
  - [ ] Test interruption handling 