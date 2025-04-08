package commands

import (
	"fmt"

	"github.com/spacebarlabs/lean_vault/pkg/vault"
)

// List displays all stored keys
func List() error {
	v := vault.New()
	secrets, err := v.ListSecrets()
	if err != nil {
		return fmt.Errorf("failed to list secrets: %w", err)
	}

	// Check if provisioning key exists
	_, err = v.GetMainProvisioningKey()
	hasProvisioningKey := err == nil

	if len(secrets) == 0 && !hasProvisioningKey {
		fmt.Println("No API keys found.")
		fmt.Println("Use 'lean_vault add <key-name>' to add a new key.")
		return nil
	}

	// Show provisioning key status
	fmt.Println("System Status:")
	if hasProvisioningKey {
		fmt.Println("  ✓ OpenRouter Provisioning Key")
	} else {
		fmt.Println("  ✗ OpenRouter Provisioning Key (not configured)")
	}

	if len(secrets) > 0 {
		fmt.Println("\nStored API keys:")
		for _, name := range secrets {
			fmt.Printf("  - %s\n", name)
		}
	}
	return nil
}
