# Lean Vault - MVP Specification

This document outlines the Minimum Viable Product (MVP) specification for the `lean_vault` CLI tool.

## Goal

To provide a secure and convenient way for developers to manage API keys, initially focusing on OpenRouter API keys, preventing accidental exposure (e.g., during livestreams or in version control). The tool aims to handle provisioning and secure retrieval of keys.

## Target Platform

Command-line interface (CLI) tool.

## Implementation Language

Go (Golang), chosen for its ease of distribution via static binaries.

## Core Functionality (CLI Commands)

The tool will be invoked via the `lean_vault` command.

1.  **`lean_vault init`**
    *   **Purpose:** Initialize the vault system for a user.
    *   **Actions:**
        *   Checks if `~/.lean_vault/.secret_vault.key` or `~/.lean_vault/secrets.vault` exist. If either exists, print an error and exit (refuse to overwrite).
        *   Prompts the user securely for their main OpenRouter provisioning API key.
        *   Generates a new master key and saves it to `~/.lean_vault/.secret_vault.key`.
        *   Creates the encrypted vault file `~/.lean_vault/secrets.vault`.
        *   Stores the user-provided main OpenRouter provisioning key within the vault under the reserved name `_MAIN_OPENROUTER_PROVISIONING_KEY_`.
        *   Ensures `.secret_vault.key` has restrictive file permissions.
    *   **Error Handling:** Clear errors if files already exist or if file operations fail.

2.  **`lean_vault add <key_name>`**
    *   **Purpose:** Provision a new OpenRouter API key and store it securely.
    *   **Actions:**
        *   Decrypts `secrets.vault` using `.secret_vault.key`.
        *   Retrieves the `_MAIN_OPENROUTER_PROVISIONING_KEY_`.
        *   Calls the OpenRouter API (using the assumed endpoint `https://openrouter.ai/api/v1`) to provision a new API key, potentially using `<key_name>` as a label.
        *   Stores the *newly generated* key value and its corresponding OpenRouter Key ID (required for revocation) under `<key_name>` in the decrypted vault data.
        *   Re-encrypts and saves `secrets.vault`.
    *   **Error Handling:** Report verbose errors to stderr if the OpenRouter API call fails (e.g., network error, invalid main key, quota exceeded) and exit with non-zero status. Do not modify the vault if provisioning fails. Report errors if the vault cannot be read/decrypted.

3.  **`lean_vault get <key_name>`**
    *   **Purpose:** Retrieve a stored secret value.
    *   **Actions:**
        *   Decrypts `secrets.vault` using `.secret_vault.key`.
        *   Looks up `<key_name>`.
        *   If found, prints *only* the raw secret value to standard output (stdout) with no trailing newline. Exits with status 0.
        *   If not found, prints nothing to stdout, prints an error message to standard error (stderr), and exits with a non-zero status.
    *   **Error Handling:** Report errors if the vault cannot be read/decrypted or if the key is not found.

4.  **`lean_vault list`**
    *   **Purpose:** List the names of all stored secrets.
    *   **Actions:**
        *   Decrypts `secrets.vault` using `.secret_vault.key`.
        *   Prints the names of all stored keys (excluding the reserved main key name, e.g., `_MAIN_OPENROUTER_PROVISIONING_KEY_`), one per line, to stdout.
    *   **Error Handling:** Report errors if the vault cannot be read/decrypted.

5.  **`lean_vault remove <key_name>`**
    *   **Purpose:** Remove a secret locally and attempt to revoke it on OpenRouter.
    *   **Actions:**
        *   Decrypts `secrets.vault` using `.secret_vault.key`.
        *   Retrieves the `_MAIN_OPENROUTER_PROVISIONING_KEY_` and the OpenRouter Key ID associated with `<key_name>`.
        *   Calls the OpenRouter API to revoke the key identified by the stored ID, using the main provisioning key.
        *   **If revocation succeeds:** Remove the entry for `<key_name>` from the decrypted vault data. Re-encrypt and save `secrets.vault`. Print a success message to stderr.
        *   **If revocation fails:** Print a verbose error message to stderr explaining the failure. Do *not* remove the key from the local vault. Exit with a non-zero status.
    *   **Error Handling:** Report errors if the vault cannot be read/decrypted, if the key name is not found, or if the OpenRouter API call fails.

6.  **`lean_vault rotate <key_name>`** (Manual Rotation - Added Post-MVP for Value Prop)
    *   **Purpose:** Provision a new key for the given name and revoke the old one.
    *   **Actions:**
        *   Decrypts `secrets.vault` using `.secret_vault.key`.
        *   Retrieves the `_MAIN_OPENROUTER_PROVISIONING_KEY_` and the OpenRouter Key ID of the *current* key associated with `<key_name>`.
        *   Calls the OpenRouter API to provision a *new* API key (potentially reusing `<key_name>` as the label).
        *   **If provisioning the new key fails:** Report a verbose error to stderr and exit with non-zero status. The vault remains unchanged.
        *   **If provisioning the new key succeeds:**
            *   Update the vault entry for `<key_name>` with the *new* key's value and its corresponding new OpenRouter Key ID.
            *   Re-encrypt and save `secrets.vault`.
            *   Call the OpenRouter API using the main provisioning key to revoke the *old* key (using the ID retrieved initially).
            *   **If revocation of the old key succeeds:** Report successful rotation to stderr. Exit with status 0.
            *   **If revocation of the old key fails:** Report the successful provisioning and vault update, but *also* report the failure to revoke the old key to stderr. Exit with a non-zero status (indicating partial success/failure).
    *   **Error Handling:** Report errors if the vault cannot be read/decrypted, if the key name is not found, or if either OpenRouter API call (provision new, revoke old) fails, following the logic above.

7.  **`lean_vault usage`** (Usage Tracking - Added Post-MVP for Value Prop)
    *   **Purpose:** Display usage and limit information for managed OpenRouter keys.
    *   **Assumptions:** Requires an OpenRouter API endpoint to fetch usage/limit details for a specific child API key (identified by its ID) when authenticated with the main provisioning key.
    *   **Actions:**
        *   Decrypts `secrets.vault` using `.secret_vault.key`.
        *   Retrieves the `_MAIN_OPENROUTER_PROVISIONING_KEY_`.
        *   Initializes a structure to hold usage data for each key.
        *   Iterates through all user-generated keys stored in the vault (identified by having a `value` and `id` field).
        *   For each key:
            *   Retrieve its stored OpenRouter Key ID.
            *   Call the assumed OpenRouter API endpoint to fetch usage/limit details for this specific key ID, authenticating with the main provisioning key.
            *   If the API call succeeds, parse and store the relevant usage details (e.g., usage amount, total limit, remaining limit, rate limits) associated with the key name.
            *   If the API call fails, record an error status for this key.
        *   Format the collected information into a user-friendly table (or similar structured output) printed to standard output. The table should include the Key Name, relevant usage/limit figures, and a Status (e.g., "OK" or an error message).
        *   Example Output Format:
          ```
          Key Name      | Usage ($) | Limit ($) | Limit Remaining ($) | Rate Limit | Status
          --------------|-----------|-----------|---------------------|------------|--------
          MY_KEY_1      | 3.25      | 10.00     | 6.75                | 60 req/min | OK
          ANOTHER_KEY   | 0.50      | 5.00      | 4.50                | 60 req/min | OK
          OLD_KEY       | -         | -         | -                   | -          | Error: Failed to fetch
          ```
    *   **Error Handling:** Report errors if the vault cannot be read/decrypted. Report specific key errors within the output table. Exit with status 0 if all key usages were fetched successfully, otherwise exit with a non-zero status.

## Storage and Encryption

*   **Master Key File:** `~/.lean_vault/.secret_vault.key`
    *   Contains the key used (after derivation) to encrypt/decrypt the vault.
    *   Should be gitignored globally or locally.
    *   File permissions should be restricted (e.g., `600`).
*   **Vault File:** `~/.lean_vault/secrets.vault`
    *   Contains the encrypted secrets.
*   **Internal Vault Format:** YAML (before encryption, after decryption).
    *   Structure example:
      ```yaml
      _MAIN_OPENROUTER_PROVISIONING_KEY_: <main_key_value>
      MY_KEY_1:
        value: <secret_value_1>
        id: <openrouter_key_id_1>
      ANOTHER_KEY:
        value: <secret_value_2>
        id: <openrouter_key_id_2>
      ```
*   **Encryption Mechanism:**
    *   Use the standard `openssl` command-line tool available on macOS and Ubuntu.
    *   Algorithm: **AES-256-GCM** (Authenticated Encryption).
    *   Key Derivation: Use a strong Key Derivation Function (KDF) like **PBKDF2** (available via `openssl`) to derive the actual encryption key from the contents of `.secret_vault.key`. This adds protection against brute-force attacks on the master key file.

## General Error Handling

*   All commands should print clear, user-friendly error messages to standard error (stderr) upon failure.
*   Use non-zero exit codes to indicate errors for scripting purposes.
*   Specifically handle cases where `.secret_vault.key` is missing or unreadable, suggesting `lean_vault init` or checking permissions.

## Assumptions (MVP)

*   The OpenRouter API endpoint is fixed at `https://openrouter.ai/api/v1`.
*   The OpenRouter API supports provisioning new keys via an API call using a provisioning key and provides a Key ID suitable for later revocation.
*   The OpenRouter API supports fetching usage/limit details for a specific child API key using its ID and authenticating with the parent provisioning key.
*   Users manage the security and backup of their `~/.lean_vault/` directory.

## Future Considerations (Post-MVP)

*   Support for other LLM providers/secret types.
*   Automatic key rotation.
*   Integration with OS keychains (macOS Keychain, etc.) as an alternative to the master key file.
*   Team features (sharing vaults).
*   Audit logging.
*   Configuration file for settings (e.g., API endpoints).
*   More sophisticated error handling or retries for API calls. 