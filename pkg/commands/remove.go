package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/spacebarlabs/lean_vault/pkg/api"
	"github.com/spacebarlabs/lean_vault/pkg/vault"
)

// Remove handles the removal of an API key
func Remove(keyName string, force bool) error {
	v := vault.New()

	// Get the key ID first to check if it exists
	keyID, err := v.GetSecretID(keyName)
	if err != nil {
		return fmt.Errorf("failed to get key ID: %w", err)
	}

	if !force {
		// Get the main provisioning key from the vault
		provisioningKey, err := v.GetMainProvisioningKey()
		if err != nil {
			return fmt.Errorf("failed to get provisioning key: %w", err)
		}

		// Create OpenRouter API client
		client := api.NewClient(provisioningKey)

		// Enable debug mode if environment variable is set
		if debug := strings.ToLower(os.Getenv("LEAN_VAULT_DEBUG")); debug == "1" || debug == "true" {
			client.SetDebug(true)
			fmt.Fprintln(os.Stderr, "Debug mode enabled")
		}

		fmt.Fprintf(os.Stderr, "Attempting to revoke API key '%s'...\n", keyName)

		// Attempt to revoke the key via OpenRouter API
		err = client.RevokeKey(keyID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Failed to revoke key on OpenRouter: %v\n", err)
			fmt.Fprintln(os.Stderr, "If the key is already inactive or you want to remove it anyway, use --force:")
			fmt.Fprintf(os.Stderr, "  lean_vault remove %s --force\n", keyName)
			return fmt.Errorf("key revocation failed")
		}
		fmt.Fprintf(os.Stderr, "✓ API key '%s' revoked successfully\n", keyName)
	}

	// Remove the key from the vault
	err = v.RemoveSecret(keyName)
	if err != nil {
		return fmt.Errorf("failed to remove key from vault: %w", err)
	}

	if force {
		fmt.Fprintf(os.Stderr, "✓ API key '%s' removed from vault (revocation skipped)\n", keyName)
	} else {
		fmt.Fprintf(os.Stderr, "✓ API key '%s' removed from vault\n", keyName)
	}
	return nil
}
