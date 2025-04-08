package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

const (
	// KeySize is the size of the AES-256 key in bytes
	KeySize = 32
	// SaltSize is the size of the salt used in PBKDF2
	SaltSize = 32
	// NonceSize is the size of the nonce used in GCM
	NonceSize = 12
	// Iterations is the number of iterations for PBKDF2
	Iterations = 100000
)

// GenerateMasterKey generates a new random master key
func GenerateMasterKey() ([]byte, error) {
	key := make([]byte, KeySize)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, fmt.Errorf("failed to generate master key: %w", err)
	}
	return key, nil
}

// DeriveKey derives an encryption key from a master key using PBKDF2
func DeriveKey(masterKey []byte, salt []byte) []byte {
	return pbkdf2.Key(masterKey, salt, Iterations, KeySize, sha256.New)
}

// Encrypt encrypts data using AES-256-GCM
func Encrypt(key []byte, plaintext []byte) (string, error) {
	// Generate a new salt for key derivation
	salt := make([]byte, SaltSize)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}

	// Derive the encryption key
	derivedKey := DeriveKey(key, salt)

	// Create cipher
	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %w", err)
	}

	// Generate nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("failed to generate nonce: %w", err)
	}

	// Encrypt and seal
	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)

	// Combine salt + nonce + ciphertext and encode
	combined := make([]byte, 0, len(salt)+len(nonce)+len(ciphertext))
	combined = append(combined, salt...)
	combined = append(combined, nonce...)
	combined = append(combined, ciphertext...)

	return base64.StdEncoding.EncodeToString(combined), nil
}

// Decrypt decrypts data using AES-256-GCM
func Decrypt(key []byte, encryptedData string) ([]byte, error) {
	// Decode the combined data
	combined, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode data: %w", err)
	}

	// Extract salt, nonce, and ciphertext
	if len(combined) < SaltSize+NonceSize {
		return nil, fmt.Errorf("encrypted data is too short")
	}

	salt := combined[:SaltSize]
	nonce := combined[SaltSize : SaltSize+NonceSize]
	ciphertext := combined[SaltSize+NonceSize:]

	// Derive the decryption key
	derivedKey := DeriveKey(key, salt)

	// Create cipher
	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	// Decrypt
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt: %w", err)
	}

	return plaintext, nil
}
