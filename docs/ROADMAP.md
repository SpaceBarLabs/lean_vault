# Lean Vault Development Roadmap

## Current State (as of March 2024)

### Completed Core Features
- ✅ Basic vault initialization and encryption
- ✅ Adding keys
- ✅ Getting keys
- ✅ Listing keys
- ✅ Basic removal functionality (needs improvement)
- ✅ Core security infrastructure (AES-256-GCM, PBKDF2)
- ✅ Basic file operations and permissions

### Pending Features and Improvements

Below are the prioritized features and improvements, ordered by importance and dependency relationships.

## 1. Key Removal and Revocation Improvements (HIGH PRIORITY)
**Status**: 🔴 Needs Immediate Attention

**Why Critical**:
- Core security feature
- Currently not working correctly
- Essential for proper key lifecycle management

**Required Changes**:
- Review and fix current removal implementation
- Add verification of key status with OpenRouter
- Implement proper error handling
- Add confirmation steps
- Ensure vault consistency during operations

## 2. Spend Limits Implementation (HIGH PRIORITY)
**Status**: 🔴 Not Started

**Why Important**:
- Critical for cost control
- Prevents unexpected charges
- Essential for production usage

**Required Features**:
- Set spend limits during key creation
- Update limits for existing keys
- Monitor usage against limits
- Automatic alerts/actions when limits approached

## 3. Key Metadata Tracking (MEDIUM PRIORITY)
**Status**: 🔴 Not Started

**Why Important**:
- Provides audit trail
- Helps with key rotation decisions
- Relatively simple to implement

**Required Changes**:
- Add created_at timestamp
- Add updated_at timestamp
- Add key status tracking
- Update vault schema to support metadata

## 4. Key Rotation Implementation (MEDIUM PRIORITY)
**Status**: 🔴 Not Started

**Dependencies**:
- Requires working key removal/revocation
- Needs metadata tracking

**Required Features**:
- Implement rotation command
- Ensure atomic operations
- Handle partial failures
- Update documentation

## 5. CLI Help System Improvements (MEDIUM PRIORITY)
**Status**: 🔴 Not Started

**Why Important**:
- Improves user experience
- Ensures consistent interface
- Helps with adoption

**Required Changes**:
- Implement consistent --help across all commands
- Add examples to help text
- Include error scenarios in help
- Document all command options

## 6. Find or Create Functionality (LOW PRIORITY)
**Status**: 🔴 Not Started

**Why Important**:
- Improves user experience
- Reduces accidental key duplication
- Simplifies automation

**Required Features**:
- Implement find or create command
- Add idempotency support
- Update documentation

## Additional Considerations

### Security Improvements
- Verify key revocation status with OpenRouter
- Implement key expiration
- Add audit logging

### Technical Debt
- Improve error handling consistency
- Add more comprehensive tests
- Review and update documentation

### Future Considerations
- Support for additional LLM providers
- Automatic key rotation
- Team sharing features
- Configuration file support

## Next Steps

1. Begin work on key removal/revocation improvements:
   - Review current implementation
   - Document specific issues
   - Design improved workflow
   - Implement and test fixes

2. After key removal is fixed:
   - Implement spend limits
   - Add key metadata tracking
   - Begin key rotation implementation

## Success Metrics
- All commands work reliably
- Keys are properly revoked when removed
- Spend limits are enforced
- Usage tracking is accurate
- Documentation is complete and accurate
- Help system is consistent and helpful

This roadmap will be updated as features are completed and priorities change. 