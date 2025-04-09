# Language Examples Guide

This document outlines the implementation of Lean Vault examples in multiple programming languages. Currently, only the Ruby implementation is available, with other language integrations planned for future releases.

## Current Status

- ✅ **Ruby**: Fully implemented and ready to use
- 🚧 **TypeScript**: Coming soon
- 🚧 **Python**: Coming soon
- 🚧 **Go**: Coming soon

## Common Features

All language examples will implement these core features:

1. **Key Management**
   - Loading keys from vault
   - Environment variable-like access
   - Key existence checking
   - Error handling

2. **Key Rotation**
   - Safe key rotation with zero downtime
   - Pre and post-rotation validation
   - Retry mechanisms
   - Error recovery

3. **OpenRouter API Integration**
   - API key validation
   - Basic chat completion example
   - Error handling

4. **Debug Support**
   - Verbose logging options
   - Command-line flags
   - Error tracing

## Language-Specific Implementations

### Ruby (Current Implementation)

The Ruby implementation is our first and currently only available integration:

```ruby
# Key loading (dotenv-style)
LeanVault.load('my-key')
puts ENV['MY_KEY']

# Key rotation
KeyRotationExample.new(debug: true).run
```

Key features:
- Constants-based secret management
- Thread-safe key loading
- Zero-downtime rotation
- Comprehensive error handling

For detailed documentation and examples, see the [Ruby Integration Guide](../examples/ruby/README.md).

### TypeScript (Coming Soon)

```typescript
// Key management
import { LeanVault } from './LeanVault';
LeanVault.load('my-key');
console.log(LeanVault.get('my-key'));

// Key rotation
new KeyRotationExample({ debug: true }).run();
```

Planned TypeScript-specific enhancements:
- Strong typing for all operations
- Async/await for API operations
- Interface definitions
- Error type hierarchies

### Python (Coming Soon)

```python
# Key management
from lean_vault import LeanVault
LeanVault.load('my-key')
print(LeanVault.get('my-key'))

# Key rotation
KeyRotationExample(debug=True).run()
```

Planned Python-specific enhancements:
- Context managers for resource handling
- Dataclasses for structured data
- Type hints throughout
- Async support (optional)

### Go (Coming Soon)

```go
// Key management
vault := leanvault.New()
vault.Load("my-key")
fmt.Println(vault.Get("my-key"))

// Key rotation
rotation := keyrotation.New(keyrotation.Options{
    Debug: true,
})
rotation.Run()
```

Planned Go-specific enhancements:
- Concurrency-safe operations
- Channel-based communication
- Strong error handling patterns
- Interface-driven design

## Project Structure

Current structure (Ruby only):
```
examples/
└── ruby/
    ├── lib/
    │   └── lean_vault.rb      # Core vault integration
    ├── examples/
    │   ├── key_rotation.rb    # Key rotation example
    │   └── basic_usage.rb     # Basic usage example
    └── README.md
```

Future structure (planned):
```
examples/
├── ruby/          # ✅ Implemented
├── typescript/    # 🚧 Coming soon
├── python/        # 🚧 Coming soon
└── golang/        # 🚧 Coming soon
```

## Implementation Guidelines

### 1. Core Integration (LeanVault Class/Module)

Each implementation will provide:
- Key loading functionality
- Constant/environment-like access
- Error handling
- Debug logging
- Thread/concurrency safety

Example interface:
```typescript
interface VaultInterface {
    load(...keys: string[]): void;
    get(key: string): string;
    isLoaded(key: string): boolean;
}
```

### 2. Key Rotation Example

Each implementation will demonstrate:
- Safe key rotation process
- Pre-rotation validation
- Post-rotation validation
- Error handling and retries
- Debug logging

Common options:
```typescript
interface RotationOptions {
    debug?: boolean;
    keyName?: string;
    retries?: number;
}
```

### 3. Basic Usage Example

Each implementation will include:
- Simple key loading
- OpenRouter API integration
- Error handling examples
- Debug output examples

## Testing

Each implementation will include tests for:
1. Basic key operations
2. Key rotation scenarios
3. Error handling
4. Edge cases

Example test structure:
```typescript
describe('LeanVault', () => {
    it('loads keys successfully')
    it('handles missing keys')
    it('rotates keys safely')
    it('handles API errors')
})
```

## Documentation

Each language example will include:
1. README with setup instructions
2. Code comments explaining key concepts
3. Example output and error scenarios
4. Troubleshooting guide

## Best Practices

1. **Error Handling**
   - Clear error messages
   - Proper error propagation
   - Recovery strategies
   - Debug information

2. **Security**
   - No key logging in debug mode
   - Proper secret handling
   - Secure API integration
   - Safe key rotation

3. **Performance**
   - Efficient key caching
   - Minimal system calls
   - Resource cleanup
   - Memory management

4. **Maintainability**
   - Clear code structure
   - Consistent naming
   - Comprehensive documentation
   - Example use cases

## Getting Started

To implement a new example:

1. Create directory structure
2. Implement core LeanVault class
3. Add key rotation example
4. Add basic usage example
5. Write tests
6. Add documentation

## Contributing

When adding a new language example:

1. Follow existing patterns
2. Maintain consistent interfaces
3. Include all core features
4. Add comprehensive tests
5. Document language-specific features
6. Update this guide 