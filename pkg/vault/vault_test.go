package vault

import (
	"os"
	"path/filepath"
	"testing"
)

func setupTestVault(t *testing.T) (*Vault, func()) {
	// Create a temporary directory for the test vault
	tmpDir, err := os.MkdirTemp("", "vault_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	// Create a test vault instance
	v := &Vault{
		vaultDir:  filepath.Join(tmpDir, DefaultVaultDir),
		vaultFile: filepath.Join(tmpDir, DefaultVaultDir, DefaultVaultFile),
		keyFile:   filepath.Join(tmpDir, DefaultVaultDir, DefaultKeyFile),
	}

	// Return the vault and a cleanup function
	cleanup := func() {
		os.RemoveAll(tmpDir)
	}

	return v, cleanup
}

func TestVaultInit(t *testing.T) {
	v, cleanup := setupTestVault(t)
	defer cleanup()

	// Test initialization
	err := v.Init("test-provisioning-key")
	if err != nil {
		t.Fatalf("Failed to initialize vault: %v", err)
	}

	// Check if files were created
	if _, err := os.Stat(v.vaultDir); os.IsNotExist(err) {
		t.Error("Vault directory was not created")
	}
	if _, err := os.Stat(v.vaultFile); os.IsNotExist(err) {
		t.Error("Vault file was not created")
	}
	if _, err := os.Stat(v.keyFile); os.IsNotExist(err) {
		t.Error("Key file was not created")
	}

	// Check file permissions
	keyInfo, err := os.Stat(v.keyFile)
	if err != nil {
		t.Fatalf("Failed to stat key file: %v", err)
	}
	if keyInfo.Mode().Perm() != DefaultFileMode {
		t.Errorf("Key file has wrong permissions: got %v, want %v", keyInfo.Mode().Perm(), DefaultFileMode)
	}

	// Test double initialization
	err = v.Init("another-key")
	if err == nil {
		t.Error("Second initialization should fail")
	}
}

func TestVaultOperations(t *testing.T) {
	v, cleanup := setupTestVault(t)
	defer cleanup()

	// Initialize vault
	if err := v.Init("test-provisioning-key"); err != nil {
		t.Fatalf("Failed to initialize vault: %v", err)
	}

	// Test adding a secret
	err := v.AddSecret("test-key", "test-value", "test-id")
	if err != nil {
		t.Fatalf("Failed to add secret: %v", err)
	}

	// Test getting the secret
	value, err := v.GetSecret("test-key")
	if err != nil {
		t.Fatalf("Failed to get secret: %v", err)
	}
	if value != "test-value" {
		t.Errorf("Got wrong secret value: got %v, want %v", value, "test-value")
	}

	// Test getting the secret ID
	id, err := v.GetSecretID("test-key")
	if err != nil {
		t.Fatalf("Failed to get secret ID: %v", err)
	}
	if id != "test-id" {
		t.Errorf("Got wrong secret ID: got %v, want %v", id, "test-id")
	}

	// Test listing secrets
	secrets, err := v.ListSecrets()
	if err != nil {
		t.Fatalf("Failed to list secrets: %v", err)
	}
	if len(secrets) != 1 || secrets[0] != "test-key" {
		t.Errorf("Got wrong secrets list: %v", secrets)
	}

	// Test getting main provisioning key
	provKey, err := v.GetMainProvisioningKey()
	if err != nil {
		t.Fatalf("Failed to get main provisioning key: %v", err)
	}
	if provKey != "test-provisioning-key" {
		t.Errorf("Got wrong provisioning key: got %v, want %v", provKey, "test-provisioning-key")
	}

	// Test removing a secret
	err = v.RemoveSecret("test-key")
	if err != nil {
		t.Fatalf("Failed to remove secret: %v", err)
	}

	// Verify secret was removed
	_, err = v.GetSecret("test-key")
	if err == nil {
		t.Error("Secret should have been removed")
	}

	// Test removing main provisioning key (should fail)
	err = v.RemoveSecret(MainProvisioningKeyName)
	if err == nil {
		t.Error("Removing main provisioning key should fail")
	}
}

func TestVaultErrors(t *testing.T) {
	v, cleanup := setupTestVault(t)
	defer cleanup()

	// Test operations before initialization
	_, err := v.GetSecret("test-key")
	if err == nil {
		t.Error("Getting secret before initialization should fail")
	}

	err = v.AddSecret("test-key", "test-value", "test-id")
	if err == nil {
		t.Error("Adding secret before initialization should fail")
	}

	// Initialize vault
	if err := v.Init("test-provisioning-key"); err != nil {
		t.Fatalf("Failed to initialize vault: %v", err)
	}

	// Test duplicate secret
	err = v.AddSecret("test-key", "test-value", "test-id")
	if err != nil {
		t.Fatalf("Failed to add first secret: %v", err)
	}

	err = v.AddSecret("test-key", "another-value", "another-id")
	if err == nil {
		t.Error("Adding duplicate secret should fail")
	}

	// Test getting non-existent secret
	_, err = v.GetSecret("non-existent")
	if err == nil {
		t.Error("Getting non-existent secret should fail")
	}

	// Test removing non-existent secret
	err = v.RemoveSecret("non-existent")
	if err == nil {
		t.Error("Removing non-existent secret should fail")
	}
}

func TestKeyRotation(t *testing.T) {
	v, cleanup := setupTestVault(t)
	defer cleanup()

	// Initialize vault with test data
	if err := v.Init("test-provisioning-key"); err != nil {
		t.Fatalf("Failed to initialize vault: %v", err)
	}

	// Add some test secrets
	testSecrets := map[string]struct {
		value string
		id    string
	}{
		"secret1": {"value1", "id1"},
		"secret2": {"value2", "id2"},
	}

	for name, secret := range testSecrets {
		if err := v.AddSecret(name, secret.value, secret.id); err != nil {
			t.Fatalf("Failed to add test secret %s: %v", name, err)
		}
	}

	// Get initial key version
	initialKey, err := v.GetCurrentKeyVersion()
	if err != nil {
		t.Fatalf("Failed to get initial key version: %v", err)
	}

	// Rotate the master key
	if err := v.RotateMasterKey(); err != nil {
		t.Fatalf("Failed to rotate master key: %v", err)
	}

	// Get new key version
	newKey, err := v.GetCurrentKeyVersion()
	if err != nil {
		t.Fatalf("Failed to get new key version: %v", err)
	}

	// Verify key rotation
	if newKey.ID == initialKey.ID {
		t.Error("Key rotation did not generate new key ID")
	}
	if newKey.CreatedAt.Before(initialKey.CreatedAt) {
		t.Error("New key version has incorrect timestamp")
	}

	// Verify all secrets are still accessible
	for name, secret := range testSecrets {
		value, err := v.GetSecret(name)
		if err != nil {
			t.Errorf("Failed to get secret %s after rotation: %v", name, err)
		}
		if value != secret.value {
			t.Errorf("Secret %s has wrong value after rotation: got %v, want %v", name, value, secret.value)
		}
	}

	// Verify main provisioning key is still accessible
	provKey, err := v.GetMainProvisioningKey()
	if err != nil {
		t.Fatalf("Failed to get main provisioning key after rotation: %v", err)
	}
	if provKey != "test-provisioning-key" {
		t.Errorf("Main provisioning key has wrong value after rotation: got %v, want %v", provKey, "test-provisioning-key")
	}
}

func TestUpdateSecret(t *testing.T) {
	v, cleanup := setupTestVault(t)
	defer cleanup()

	// Initialize vault
	if err := v.Init("test-provisioning-key"); err != nil {
		t.Fatalf("Failed to initialize vault: %v", err)
	}

	// Add initial secret
	if err := v.AddSecret("test-key", "initial-value", "initial-id"); err != nil {
		t.Fatalf("Failed to add initial secret: %v", err)
	}

	// Update the secret
	if err := v.UpdateSecret("test-key", "updated-value", "updated-id"); err != nil {
		t.Fatalf("Failed to update secret: %v", err)
	}

	// Verify updated value
	value, err := v.GetSecret("test-key")
	if err != nil {
		t.Fatalf("Failed to get updated secret: %v", err)
	}
	if value != "updated-value" {
		t.Errorf("Secret has wrong value after update: got %v, want %v", value, "updated-value")
	}

	// Verify updated ID
	id, err := v.GetSecretID("test-key")
	if err != nil {
		t.Fatalf("Failed to get updated secret ID: %v", err)
	}
	if id != "updated-id" {
		t.Errorf("Secret has wrong ID after update: got %v, want %v", id, "updated-id")
	}

	// Test updating non-existent secret
	err = v.UpdateSecret("non-existent", "value", "id")
	if err == nil {
		t.Error("Updating non-existent secret should fail")
	}
}

func TestCorruptedVault(t *testing.T) {
	v, cleanup := setupTestVault(t)
	defer cleanup()

	// Initialize vault
	if err := v.Init("test-provisioning-key"); err != nil {
		t.Fatalf("Failed to initialize vault: %v", err)
	}

	// Corrupt the vault file
	if err := os.WriteFile(v.vaultFile, []byte("corrupted data"), DefaultFileMode); err != nil {
		t.Fatalf("Failed to corrupt vault file: %v", err)
	}

	// Verify operations fail gracefully
	_, err := v.GetSecret("any-key")
	if err == nil {
		t.Error("Operation on corrupted vault should fail")
	}

	// Corrupt the key file
	if err := os.WriteFile(v.keyFile, []byte("corrupted key"), DefaultFileMode); err != nil {
		t.Fatalf("Failed to corrupt key file: %v", err)
	}

	// Verify operations fail gracefully
	_, err = v.GetSecret("any-key")
	if err == nil {
		t.Error("Operation with corrupted key should fail")
	}
}

func TestPermissionHandling(t *testing.T) {
	v, cleanup := setupTestVault(t)
	defer cleanup()

	// Initialize vault
	if err := v.Init("test-provisioning-key"); err != nil {
		t.Fatalf("Failed to initialize vault: %v", err)
	}

	// Test readonly vault file
	if err := os.Chmod(v.vaultFile, 0400); err != nil {
		t.Fatalf("Failed to change vault file permissions: %v", err)
	}

	// Verify write operations fail gracefully
	err := v.AddSecret("test-key", "test-value", "test-id")
	if err == nil {
		t.Error("Write operation on readonly vault should fail")
	}

	// Reset permissions
	if err := os.Chmod(v.vaultFile, DefaultFileMode); err != nil {
		t.Fatalf("Failed to reset vault file permissions: %v", err)
	}

	// Test readonly key file
	if err := os.Chmod(v.keyFile, 0400); err != nil {
		t.Fatalf("Failed to change key file permissions: %v", err)
	}

	// Verify key rotation fails gracefully
	err = v.RotateMasterKey()
	if err == nil {
		t.Error("Key rotation with readonly key file should fail")
	}
}
