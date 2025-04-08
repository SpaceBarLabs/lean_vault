package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/spacebarlabs/lean_vault/pkg/vault"
)

// Get retrieves a secret from the vault
func Get(keyName string) error {
	v := vault.New()

	// Get the secret value
	value, err := v.GetSecret(keyName)
	if err != nil {
		return fmt.Errorf("failed to get secret: %w", err)
	}

	// Clean the value and print only the value to stdout (no newline)
	value = strings.TrimSpace(value)
	fmt.Fprint(os.Stdout, value)
	return nil
}
