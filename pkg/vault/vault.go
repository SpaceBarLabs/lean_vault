package vault

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spacebarlabs/lean_vault/pkg/crypto"
	"gopkg.in/yaml.v3"
)

const (
	// DefaultVaultDir is the default directory for vault files
	DefaultVaultDir = ".lean_vault"
	// DefaultVaultFile is the default name for the vault file
	DefaultVaultFile = "secrets.vault"
	// DefaultKeyFile is the default name for the master key file
	DefaultKeyFile = ".secret_vault.key"
	// DefaultFileMode is the default file permissions for sensitive files
	DefaultFileMode = 0600
	// DefaultDirMode is the default directory permissions
	DefaultDirMode = 0700
	// MainProvisioningKeyName is the reserved name for the main OpenRouter provisioning key
	MainProvisioningKeyName = "_MAIN_OPENROUTER_PROVISIONING_KEY_"
)

// VaultData represents the structure of the vault file
type VaultData struct {
	Secrets map[string]SecretEntry `yaml:"secrets"`
}

// SecretEntry represents a single secret entry in the vault
type SecretEntry struct {
	Value string `yaml:"value"`
	ID    string `yaml:"id,omitempty"`
}

// Vault represents the vault manager
type Vault struct {
	vaultDir  string
	vaultFile string
	keyFile   string
}

// New creates a new vault manager
func New() *Vault {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}

	return &Vault{
		vaultDir:  filepath.Join(homeDir, DefaultVaultDir),
		vaultFile: filepath.Join(homeDir, DefaultVaultDir, DefaultVaultFile),
		keyFile:   filepath.Join(homeDir, DefaultVaultDir, DefaultKeyFile),
	}
}

// Init initializes a new vault
func (v *Vault) Init(mainProvisioningKey string) error {
	// Check if vault already exists
	if _, err := os.Stat(v.vaultFile); err == nil {
		return fmt.Errorf("vault file already exists at %s", v.vaultFile)
	}
	if _, err := os.Stat(v.keyFile); err == nil {
		return fmt.Errorf("key file already exists at %s", v.keyFile)
	}

	// Create vault directory with restricted permissions
	if err := os.MkdirAll(v.vaultDir, DefaultDirMode); err != nil {
		return fmt.Errorf("failed to create vault directory: %w", err)
	}

	// Generate master key
	masterKey, err := crypto.GenerateMasterKey()
	if err != nil {
		return fmt.Errorf("failed to generate master key: %w", err)
	}

	// Save master key
	if err := os.WriteFile(v.keyFile, masterKey, DefaultFileMode); err != nil {
		return fmt.Errorf("failed to save master key: %w", err)
	}

	// Create initial vault data
	vaultData := VaultData{
		Secrets: map[string]SecretEntry{
			MainProvisioningKeyName: {
				Value: mainProvisioningKey,
			},
		},
	}

	// Marshal vault data
	data, err := yaml.Marshal(vaultData)
	if err != nil {
		return fmt.Errorf("failed to marshal vault data: %w", err)
	}

	// Encrypt vault data
	encrypted, err := crypto.Encrypt(masterKey, data)
	if err != nil {
		return fmt.Errorf("failed to encrypt vault data: %w", err)
	}

	// Save encrypted vault
	if err := os.WriteFile(v.vaultFile, []byte(encrypted), DefaultFileMode); err != nil {
		return fmt.Errorf("failed to save vault file: %w", err)
	}

	return nil
}

// load reads and decrypts the vault
func (v *Vault) load() (*VaultData, []byte, error) {
	// Read master key
	masterKey, err := os.ReadFile(v.keyFile)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read master key: %w", err)
	}

	// Read encrypted vault
	encrypted, err := os.ReadFile(v.vaultFile)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read vault file: %w", err)
	}

	// Decrypt vault
	decrypted, err := crypto.Decrypt(masterKey, string(encrypted))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decrypt vault: %w", err)
	}

	// Unmarshal vault data
	var vaultData VaultData
	if err := yaml.Unmarshal(decrypted, &vaultData); err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal vault data: %w", err)
	}

	return &vaultData, masterKey, nil
}

// save encrypts and saves the vault
func (v *Vault) save(vaultData *VaultData, masterKey []byte) error {
	// Marshal vault data
	data, err := yaml.Marshal(vaultData)
	if err != nil {
		return fmt.Errorf("failed to marshal vault data: %w", err)
	}

	// Encrypt vault data
	encrypted, err := crypto.Encrypt(masterKey, data)
	if err != nil {
		return fmt.Errorf("failed to encrypt vault data: %w", err)
	}

	// Save encrypted vault
	if err := os.WriteFile(v.vaultFile, []byte(encrypted), DefaultFileMode); err != nil {
		return fmt.Errorf("failed to save vault file: %w", err)
	}

	return nil
}

// AddSecret adds a new secret to the vault
func (v *Vault) AddSecret(name, value, id string) error {
	vaultData, masterKey, err := v.load()
	if err != nil {
		return err
	}

	if _, exists := vaultData.Secrets[name]; exists {
		return fmt.Errorf("secret %s already exists", name)
	}

	vaultData.Secrets[name] = SecretEntry{
		Value: value,
		ID:    id,
	}

	return v.save(vaultData, masterKey)
}

// GetSecret retrieves a secret from the vault
func (v *Vault) GetSecret(name string) (string, error) {
	vaultData, _, err := v.load()
	if err != nil {
		return "", err
	}

	secret, exists := vaultData.Secrets[name]
	if !exists {
		return "", fmt.Errorf("secret %s not found", name)
	}

	return secret.Value, nil
}

// ListSecrets returns a list of all secret names (excluding the main provisioning key)
func (v *Vault) ListSecrets() ([]string, error) {
	vaultData, _, err := v.load()
	if err != nil {
		return nil, err
	}

	secrets := make([]string, 0, len(vaultData.Secrets)-1)
	for name := range vaultData.Secrets {
		if name != MainProvisioningKeyName {
			secrets = append(secrets, name)
		}
	}

	return secrets, nil
}

// RemoveSecret removes a secret from the vault
func (v *Vault) RemoveSecret(name string) error {
	vaultData, masterKey, err := v.load()
	if err != nil {
		return err
	}

	if name == MainProvisioningKeyName {
		return fmt.Errorf("cannot remove main provisioning key")
	}

	if _, exists := vaultData.Secrets[name]; !exists {
		return fmt.Errorf("secret %s not found", name)
	}

	delete(vaultData.Secrets, name)
	return v.save(vaultData, masterKey)
}

// GetSecretID retrieves the ID associated with a secret
func (v *Vault) GetSecretID(name string) (string, error) {
	vaultData, _, err := v.load()
	if err != nil {
		return "", err
	}

	secret, exists := vaultData.Secrets[name]
	if !exists {
		return "", fmt.Errorf("secret %s not found", name)
	}

	return secret.ID, nil
}

// GetMainProvisioningKey retrieves the main OpenRouter provisioning key
func (v *Vault) GetMainProvisioningKey() (string, error) {
	return v.GetSecret(MainProvisioningKeyName)
}

// VaultDir returns the path to the vault directory
func (v *Vault) VaultDir() string {
	return v.vaultDir
}
