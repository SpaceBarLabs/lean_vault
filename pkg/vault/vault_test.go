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
