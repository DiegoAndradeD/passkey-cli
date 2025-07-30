package cmd

import (
	"log"

	"github.com/DiegoAndradeD/passkey-cli/utils"
	"github.com/DiegoAndradeD/passkey-cli/vault"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes one/many services by its name",
	Long: `The "delete" command allows you to delete one/many services in your secure vault.
Usage:
passkey-cli delete --name <service_name> --passkey <vault_passkey>

Examples:
passkey-cli delete --name github --passkey password
passkey-cli delete -n "my-database -p password"

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
		deleteService(name, passkey)
	},
}

func deleteService(name, passkey string) {
	vaultPath := utils.GetVaultPath()

	if err := vault.DeleteService(vaultPath, name, passkey); err != nil {
		log.Fatalf("failed to delete service: %v", err)
	}
	log.Printf("Service '%s' deleted successfully.", name)
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("name", "n", "", "Service name")
	deleteCmd.Flags().StringP("passkey", "p", "", "Vault passkey")
}
