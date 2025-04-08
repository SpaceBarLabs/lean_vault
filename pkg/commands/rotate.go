package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/spacebarlabs/lean_vault/pkg/api"
	"github.com/spacebarlabs/lean_vault/pkg/vault"
)

// Rotate handles the rotation of an API key
func Rotate(keyName string) error {
	v := vault.New()

	// 1. Get the current key's ID and verify it exists
	oldKeyID, err := v.GetSecretID(keyName)
	if err != nil {
		return fmt.Errorf("failed to get current key ID: %w", err)
	}

	// 2. Get the main provisioning key
	provisioningKey, err := v.GetMainProvisioningKey()
	if err != nil {
		return fmt.Errorf("failed to get provisioning key: %w", err)
	}

	// 3. Create OpenRouter API client
	client := api.NewClient(provisioningKey)

	// Enable debug mode if set
	if debug := strings.ToLower(os.Getenv("LEAN_VAULT_DEBUG")); debug == "1" || debug == "true" {
		client.SetDebug(true)
		fmt.Fprintln(os.Stderr, "Debug mode enabled")
	}

	fmt.Fprintf(os.Stderr, "Rotating API key '%s'...\n", keyName)

	// 4. Create new key
	fmt.Fprintf(os.Stderr, "Creating new key...\n")
	resp, err := client.CreateKey(keyName)
	if err != nil {
		return fmt.Errorf("failed to create new API key: %w", err)
	}

	// 5. Update vault with new key
	fmt.Fprintf(os.Stderr, "Updating vault with new key...\n")
	err = v.UpdateSecret(keyName, resp.Key, resp.Data.Hash)
	if err != nil {
		return fmt.Errorf("failed to store new API key: %w", err)
	}

	// 6. Revoke old key
	fmt.Fprintf(os.Stderr, "Revoking old key...\n")
	err = client.RevokeKey(oldKeyID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "⚠️  Warning: Failed to revoke old key: %v\n", err)
		fmt.Fprintf(os.Stderr, "The new key has been stored successfully, but the old key may still be active.\n")
		fmt.Fprintf(os.Stderr, "You may want to try revoking it manually or contact OpenRouter support.\n")
		return fmt.Errorf("key rotation partially succeeded but revocation failed")
	}

	fmt.Fprintf(os.Stderr, "✓ API key '%s' rotated successfully!\n", keyName)
	return nil
}
