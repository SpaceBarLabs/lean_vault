# Lean Vault Implementation Plan

This document outlines the detailed implementation plan for high-priority features identified in the roadmap. Each task includes specific technical requirements, implementation steps, and testing criteria.

## 1. Key Removal and Revocation Improvements

### Phase 1: Analysis and Design
- [ ] Review current implementation
  - Analyze existing code for key removal
  - Document current workflow
  - Identify specific failure points
  - Map OpenRouter API interactions

- [ ] Design improved workflow
  ```
  1. User initiates key removal
  2. Verify key exists in vault
  3. Verify key status with OpenRouter
  4. Attempt revocation with OpenRouter
  5. Update vault state
  6. Provide user feedback
  ```

### Phase 2: Implementation
- [ ] Enhance OpenRouter API Integration
  ```go
  // Add new API methods
  type OpenRouterClient interface {
    // ... existing methods ...
    VerifyKeyStatus(keyID string) (*KeyStatus, error)
    RevokeKey(keyID string) error
    GetKeyDetails(keyID string) (*KeyDetails, error)
  }
  ```

- [ ] Implement Key Status Verification
  - Add status check before removal
  - Handle cases where key is already revoked
  - Implement retry logic for API calls

- [ ] Improve Error Handling
  ```go
  // New error types
  type KeyRevocationError struct {
    KeyID string
    Reason string
    IsRetryable bool
  }

  // Error scenarios to handle
  - Key not found in vault
  - Key already revoked
  - Network failures
  - API rate limiting
  - Partial failure states
  ```

- [ ] Add Confirmation Flow
  - Implement interactive confirmation
  - Support --force flag for scripts
  - Add dry-run option

### Phase 3: Testing
- [ ] Unit Tests
  ```go
  func TestKeyRemoval(t *testing.T) {
    // Test scenarios
    - Happy path (successful removal)
    - Key not found
    - Already revoked key
    - Network failures
    - Retry scenarios
    - Force flag behavior
  }
  ```

- [ ] Integration Tests
  - End-to-end removal workflow
  - API interaction verification
  - Vault state consistency
  - Error recovery

## 2. Spend Limits Implementation

### Phase 1: Design
- [ ] Define Spend Limit Schema
  ```yaml
  key_name:
    value: "<key_value>"
    id: "<openrouter_key_id>"
    limits:
      monthly_spend: 100.00
      daily_spend: 10.00
      alert_threshold: 0.80  # 80% of limit
    usage:
      current_month: 45.20
      current_day: 5.30
      last_updated: "2024-03-20T15:04:05Z"
  ```

- [ ] Design CLI Interface
  ```bash
  # New commands
  lean_vault add my-key --monthly-limit 100 --daily-limit 10
  lean_vault update-limits my-key --monthly-limit 200
  lean_vault get-usage my-key
  ```

### Phase 2: Implementation
- [ ] Add Limit Management
  ```go
  type SpendLimit struct {
    Monthly float64
    Daily   float64
    AlertThreshold float64
  }

  // Methods to implement
  - SetSpendLimits(keyName string, limits SpendLimit) error
  - UpdateSpendLimits(keyName string, limits SpendLimit) error
  - GetCurrentUsage(keyName string) (*Usage, error)
  ```

- [ ] Implement Usage Tracking
  - Periodic usage fetching
  - Usage data storage
  - Alert generation

- [ ] Add CLI Commands
  - Limit setting/updating
  - Usage reporting
  - Alert configuration

### Phase 3: Testing
- [ ] Unit Tests
  - Limit validation
  - Usage calculation
  - Alert triggering

- [ ] Integration Tests
  - End-to-end limit setting
  - Usage tracking accuracy
  - Alert delivery

## Implementation Timeline

### Week 1: Key Removal Improvements
- Days 1-2: Analysis and design
- Days 3-4: Core implementation
- Day 5: Testing and refinement

### Week 2: Spend Limits (Part 1)
- Days 1-2: Schema and CLI design
- Days 3-4: Basic limit implementation
- Day 5: Initial testing

### Week 3: Spend Limits (Part 2)
- Days 1-2: Usage tracking
- Days 3-4: Alert system
- Day 5: Integration testing

## Success Criteria

### Key Removal
- [ ] 100% success rate for valid removals
- [ ] Proper error handling for all failure cases
- [ ] Vault remains consistent after failed operations
- [ ] Clear user feedback for all operations

### Spend Limits
- [ ] Accurate limit enforcement
- [ ] Real-time usage tracking
- [ ] Reliable alert delivery
- [ ] Clear usage reporting

## Dependencies

### External APIs
- OpenRouter key management endpoints
- OpenRouter usage tracking endpoints
- (Optional) Alert delivery service

### Internal Changes
- Vault schema updates
- CLI command additions
- Configuration file updates

## Risks and Mitigations

### Key Removal
- **Risk**: API failures during removal
  - **Mitigation**: Implement robust retry logic
  - **Mitigation**: Add rollback capability

- **Risk**: Inconsistent vault state
  - **Mitigation**: Implement atomic operations
  - **Mitigation**: Add state verification

### Spend Limits
- **Risk**: Usage data lag
  - **Mitigation**: Clear update timestamps
  - **Mitigation**: Conservative limit enforcement

- **Risk**: Alert spam
  - **Mitigation**: Configurable alert thresholds
  - **Mitigation**: Alert coalescing

## Next Steps

1. Begin with key removal improvements
   - Start analysis phase
   - Schedule design review
   - Set up test environment

2. Parallel planning for spend limits
   - Confirm API capabilities
   - Design schema
   - Plan migration strategy 