# Language Examples Guide

This document outlines the implementation of Lean Vault examples in multiple programming languages. Each example demonstrates best practices for integrating Lean Vault into your applications, with a focus on key management and rotation.

## Common Features

All language examples implement these core features:

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

### Ruby (Reference Implementation)

The Ruby implementation serves as our reference, demonstrating core patterns:

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

### TypeScript Implementation

```typescript
// Key management
import { LeanVault } from './LeanVault';
LeanVault.load('my-key');
console.log(LeanVault.get('my-key'));

// Key rotation
new KeyRotationExample({ debug: true }).run();
```

TypeScript-specific enhancements:
- Strong typing for all operations
- Async/await for API operations
- Interface definitions
- Error type hierarchies

### Python Implementation

```python
# Key management
from lean_vault import LeanVault
LeanVault.load('my-key')
print(LeanVault.get('my-key'))

# Key rotation
KeyRotationExample(debug=True).run()
```

Python-specific enhancements:
- Context managers for resource handling
- Dataclasses for structured data
- Type hints throughout
- Async support (optional)

### Go Implementation

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

Go-specific enhancements:
- Concurrency-safe operations
- Channel-based communication
- Strong error handling patterns
- Interface-driven design

## Project Structure

```
examples/
├── typescript/
│   ├── package.json
│   ├── tsconfig.json
│   ├── src/
│   │   ├── LeanVault.ts        # Core vault integration
│   │   ├── keyRotation.ts      # Key rotation example
│   │   └── test.ts            # Basic usage example
│   └── README.md
├── python/
│   ├── requirements.txt
│   ├── lean_vault.py          # Core vault integration
│   ├── key_rotation.py        # Key rotation example
│   ├── test.py               # Basic usage example
│   └── README.md
└── golang/
    ├── go.mod
    ├── go.sum
    ├── leanvault/
    │   └── leanvault.go       # Core vault integration
    ├── cmd/
    │   ├── keyrotation/
    │   │   └── main.go        # Key rotation example
    │   └── test/
    │       └── main.go        # Basic usage example
    └── README.md
```

## Implementation Guidelines

### 1. Core Integration (LeanVault Class/Module)

Each implementation should provide:
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

Each implementation should demonstrate:
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

Each implementation should include:
- Simple key loading
- OpenRouter API integration
- Error handling examples
- Debug output examples

## Testing

Each implementation should include tests for:
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

Each language example includes:
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