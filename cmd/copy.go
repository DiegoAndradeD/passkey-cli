package cmd

import (
	"fmt"
	"log"

	"github.com/DiegoAndradeD/passkey-cli/utils"
	"github.com/DiegoAndradeD/passkey-cli/vault"
	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy a service password to the system clipboard",
	Long: `The copy command finds the password for a given service in your vault
and copies it directly to your clipboard for easy pasting.

You must specify the service name using the --name (-n) flag and provide
your vault passkey using the --passkey (-p) flag.

Example usage:

  myapp copy --name myservice --passkey mysecret

This will load your vault, find the password for 'myservice', and copy it
to your clipboard securely. Make sure your clipboard utility is properly configured.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil || name == "" {
			log.Fatal("the service name must be specified with --name", err)
		}
		passkey, err := cmd.Flags().GetString("passkey")
		if err != nil || passkey == "" {
			log.Fatal("you must specify the passkey with --passkey flag")
		}
		vaultPath := utils.GetVaultPath()
		if err = vault.CopyServicePassword(vaultPath, name, passkey); err != nil {
			fmt.Print(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.Flags().StringP("name", "n", "", "Service name")
	copyCmd.Flags().StringP("passkey", "p", "", "Vault passkey")

}
