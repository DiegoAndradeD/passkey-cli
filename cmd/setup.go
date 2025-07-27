/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/DiegoAndradeD/passkey-cli/utils"
	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
