package commands

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/spacebarlabs/lean_vault/pkg/vault"
	"golang.org/x/term"
)

// Init handles the initialization of the vault
func Init() error {
	// Create a new vault instance
	v := vault.New()

	// Prompt for the OpenRouter provisioning key
	fmt.Print("Enter your OpenRouter provisioning key: ")
	key, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return fmt.Errorf("failed to read key: %w", err)
	}
	fmt.Println() // Add newline after password input

	// Trim any whitespace from the key
	keyStr := strings.TrimSpace(string(key))
	if keyStr == "" {
		return fmt.Errorf("provisioning key cannot be empty")
	}

	// Initialize the vault
	if err := v.Init(keyStr); err != nil {
		return fmt.Errorf("failed to initialize vault: %w", err)
	}

	fmt.Fprintln(os.Stderr, "Vault initialized successfully!")
	fmt.Fprintln(os.Stderr, "Your vault is located at:", v.VaultDir())
	return nil
}
