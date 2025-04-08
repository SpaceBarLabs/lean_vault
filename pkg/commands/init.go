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

	// Print instructions
	fmt.Fprintln(os.Stderr, "Initialize your Lean Vault")
	fmt.Fprintln(os.Stderr, "------------------------")
	fmt.Fprintln(os.Stderr, "This will create a secure vault in:", v.VaultDir())
	fmt.Fprintln(os.Stderr, "Please enter your OpenRouter provisioning key.")
	fmt.Fprintln(os.Stderr, "This is the master key used to provision new API keys.")
	fmt.Fprintln(os.Stderr, "You can find this in your OpenRouter dashboard.")
	fmt.Fprintln(os.Stderr)

	// Prompt for the OpenRouter provisioning key
	fmt.Print("OpenRouter Provisioning Key (input hidden): ")
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

	fmt.Fprintln(os.Stderr, "\nInitializing vault...")

	// Initialize the vault
	if err := v.Init(keyStr); err != nil {
		return fmt.Errorf("failed to initialize vault: %w", err)
	}

	fmt.Fprintln(os.Stderr, "\n✓ Vault initialized successfully!")
	fmt.Fprintln(os.Stderr, "✓ Your vault is located at:", v.VaultDir())
	fmt.Fprintln(os.Stderr, "\nYou can now use 'lean_vault add <key-name>' to create new API keys.")
	return nil
}
