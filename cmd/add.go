package cmd

import (
	"log"
	"time"

	"github.com/DiegoAndradeD/passkey-cli/utils"
	"github.com/DiegoAndradeD/passkey-cli/vault"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("name", "n", "", "Service name")
	addCmd.Flags().StringP("passkey", "p", "", "Vault passkey")

}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new service with automatically generated password",
	Long: `The "add" command allows you to register a new service in your secure vault.
It automatically generates a strong random password, associates it with the
specified service name, and stores it in the encrypted vault file.

Usage:
passkey-cli add --name <service_name>

Examples:
passkey-cli add --name github
passkey-cli add -n "my-database"

After running this command, the service and its generated password will be
saved securely in your vault file. You can later retrieve or manage it using
other commands provided by the CLI.`,

	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil || name == "" {
			log.Fatal("the service name must be specified with --name", err)
		}
		passkey, err := cmd.Flags().GetString("passkey")
		if err != nil || passkey == "" {
			log.Fatal("you must specify the passkey with --passkey flag")
		}
		createService(name, passkey)
	},
}

func createService(name, passkey string) {

	vaultPath := utils.GetVaultPath()

	v, err := vault.LoadVault(vaultPath, passkey)
	if err != nil {
		log.Fatalf("failed to load vault: %v", err)
	}

	password, err := utils.GeneratePassword()
	if err != nil {
		log.Fatalf("failed to generate password: %v", err)
	}

	service := vault.Service{
		Name:      name,
		Password:  password,
		CreatedAt: time.Now(),
	}

	v.Services = append(v.Services, service)

	err = vault.SaveVault(vaultPath, v)
	if err != nil {
		log.Fatalf("failed to save vault: %v", err)
	}

	log.Printf("Service '%s' added successfully.", name)
}
