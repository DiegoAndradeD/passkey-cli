package cmd

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/DiegoAndradeD/passkey-cli/utils"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "A brief description of your command",
	Long: `The "setup" command initializes the local vault used to securely store
service credentials and generated passwords.

When you run this command, it will:
  • Check if a vault already exists in your user configuration directory.
  • If no vault exists, it creates the necessary directories and an empty vault file.
  • If a vault already exists, it leaves it unchanged.

Usage:
  passkey-cli setup

Examples:
  passkey-cli setup
    Initializes the vault if it doesn't already exist.

After running this command successfully, your CLI will be ready to store
services and generated passwords using other commands such as "add".`,
	Run: func(cmd *cobra.Command, args []string) {
		getOrCreateVault()
	},
}

func isVaultCreated(vaultPath string) bool {
	_, err := os.Stat(vaultPath)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	log.Fatal("Failed to verify vault path ", err)
	return false
}

func getOrCreateVault() {
	vaultPath := utils.GetVaultPath()

	vaultAlreadyExists := isVaultCreated(vaultPath)

	if vaultAlreadyExists {
		log.Println("Vault already exists")
		return
	}

	vaultDir := filepath.Dir(vaultPath)
	err := os.MkdirAll(vaultDir, 0700)
	if err != nil {
		log.Fatal("Failed to create directories: ", err)
	}

	vault, err := os.Create(vaultPath)
	if err != nil {
		log.Fatal("Failed to create vault file: ", err)
	}
	defer vault.Close()
	_, err = vault.WriteString("[]")
	utils.HandleError(err)
	log.Println("Vault has been created successfully")
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
