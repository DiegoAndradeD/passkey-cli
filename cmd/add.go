package cmd

import (
	"log"

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
passkey-cli add --name <service_name> --passkey <vault_passkey>

Examples:
passkey-cli add --name github --passkey password
passkey-cli add -n "my-database -p password"

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
	if err := vault.AddService(vaultPath, name, passkey); err != nil {
		log.Fatalf("failed to add service: %v", err)
	}

	log.Printf("Service '%s' added successfully.", name)
}
