package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/spacebarlabs/lean_vault/pkg/api"
	"github.com/spacebarlabs/lean_vault/pkg/vault"
)

// Add handles the addition of a new API key
func Add(keyName string) error {
	v := vault.New()

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

	fmt.Fprintf(os.Stderr, "Provisioning new API key '%s'...\n", keyName)

	// Create new key via OpenRouter API
	resp, err := client.CreateKey(keyName)
	if err != nil {
		return fmt.Errorf("failed to create API key: %w", err)
	}

	// Store the new key in the vault
	err = v.AddSecret(keyName, resp.Data.Key, resp.Data.Hash)
	if err != nil {
		return fmt.Errorf("failed to store API key: %w", err)
	}

	fmt.Fprintf(os.Stderr, "✓ API key '%s' created and stored successfully!\n", keyName)
	return nil
}
