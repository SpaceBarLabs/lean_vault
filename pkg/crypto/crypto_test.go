package crypto

import (
	"bytes"
	"testing"
)

func TestGenerateMasterKey(t *testing.T) {
	key, err := GenerateMasterKey()
	if err != nil {
		t.Fatalf("Failed to generate master key: %v", err)
	}
	if len(key) != KeySize {
		t.Errorf("Expected key size %d, got %d", KeySize, len(key))
	}
}

func TestEncryptDecrypt(t *testing.T) {
	tests := []struct {
		name      string
		plaintext []byte
	}{
		{
			name:      "Empty string",
			plaintext: []byte(""),
		},
		{
			name:      "Simple string",
			plaintext: []byte("hello world"),
		},
		{
			name:      "Long string",
			plaintext: []byte("this is a much longer string that we'll use to test encryption and decryption of larger data"),
		},
		{
			name:      "Binary data",
			plaintext: []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Generate a master key
			key, err := GenerateMasterKey()
			if err != nil {
				t.Fatalf("Failed to generate master key: %v", err)
			}

			// Encrypt the data
			encrypted, err := Encrypt(key, tt.plaintext)
			if err != nil {
				t.Fatalf("Failed to encrypt: %v", err)
			}

			// Decrypt the data
			decrypted, err := Decrypt(key, encrypted)
			if err != nil {
				t.Fatalf("Failed to decrypt: %v", err)
			}

			// Compare the result
			if !bytes.Equal(tt.plaintext, decrypted) {
				t.Errorf("Decrypted data doesn't match original.\nExpected: %v\nGot: %v", tt.plaintext, decrypted)
			}
		})
	}
}

func TestDecryptWithWrongKey(t *testing.T) {
	// Generate a key and encrypt some data
	key1, _ := GenerateMasterKey()
	plaintext := []byte("secret message")
	encrypted, err := Encrypt(key1, plaintext)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	// Try to decrypt with a different key
	key2, _ := GenerateMasterKey()
	_, err = Decrypt(key2, encrypted)
	if err == nil {
		t.Error("Expected decryption with wrong key to fail")
	}
}

func TestDeriveKey(t *testing.T) {
	masterKey := []byte("master key")
	salt1 := []byte("salt1")
	salt2 := []byte("salt2")

	// Same master key and salt should produce same derived key
	key1 := DeriveKey(masterKey, salt1)
	key2 := DeriveKey(masterKey, salt1)
	if !bytes.Equal(key1, key2) {
		t.Error("Derived keys should be equal with same input")
	}

	// Same master key but different salt should produce different keys
	key3 := DeriveKey(masterKey, salt2)
	if bytes.Equal(key1, key3) {
		t.Error("Derived keys should be different with different salts")
	}
}
