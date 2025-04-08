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

	// Check if vault already exists before showing any prompts
	if _, err := os.Stat(v.VaultDir()); err == nil {
		fmt.Fprintln(os.Stderr, "\n⚠️  Vault already exists!")
		fmt.Fprintln(os.Stderr, "Location:", v.VaultDir())
		fmt.Fprintln(os.Stderr, "\nTo start fresh:")
		fmt.Fprintf(os.Stderr, "1. Remove the existing vault: rm -rf %s\n", v.VaultDir())
		fmt.Fprintln(os.Stderr, "2. Run 'lean_vault init' again")
		fmt.Fprintln(os.Stderr, "\n⚠️  Warning: Removing the vault will delete all stored API keys!")
		return fmt.Errorf("vault already initialized")
	}

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
		// This should rarely happen since we checked earlier, but handle it just in case
		if strings.Contains(err.Error(), "already exists") {
			fmt.Fprintln(os.Stderr, "\n⚠️  Another process may have initialized the vault!")
			fmt.Fprintln(os.Stderr, "Location:", v.VaultDir())
			fmt.Fprintln(os.Stderr, "\nTo start fresh:")
			fmt.Fprintf(os.Stderr, "1. Remove the existing vault: rm -rf %s\n", v.VaultDir())
			fmt.Fprintln(os.Stderr, "2. Run 'lean_vault init' again")
			fmt.Fprintln(os.Stderr, "\n⚠️  Warning: Removing the vault will delete all stored API keys!")
			return fmt.Errorf("vault already initialized")
		}
		return fmt.Errorf("failed to initialize vault: %w", err)
	}

	fmt.Fprintln(os.Stderr, "\n✓ Vault initialized successfully!")
	fmt.Fprintln(os.Stderr, "✓ Your vault is located at:", v.VaultDir())
	fmt.Fprintln(os.Stderr, "\nYou can now use 'lean_vault add <key-name>' to create new API keys.")
	return nil
}
