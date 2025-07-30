package cmd

import (
	"fmt"

	"github.com/DiegoAndradeD/passkey-cli/utils"
	"github.com/DiegoAndradeD/passkey-cli/vault"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing service in the vault",
	Long: `Update an existing service stored in your password vault.
You can change the service name and optionally regenerate its password.
If the new service name already exists, the command will fail.

Examples:
	Update only the service name:
	passkey-cli update --passkey myPass --old github --new github-work
	passkey-cli update -p myPass -o github -n github-work

	Update the name and regenerate the password:
	passkey-cli update --passkey myPass --old github --new github-work --regen
	passkey-cli update -p myPass -o github -n github-work -r
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		vaultPath := utils.GetVaultPath()
		if err := vault.UpdateService(vaultPath, oldName, newName, passkey, regenPass); err != nil {
			return err
		}
		fmt.Println("Service updated successfully")
		return nil
	},
}

var (
	passkey   string
	oldName   string
	newName   string
	regenPass bool
)

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&passkey, "passkey", "p", "", "Vault passkey (required)")
	updateCmd.Flags().StringVarP(&oldName, "old", "o", "", "Current service name (required)")
	updateCmd.Flags().StringVarP(&newName, "new", "n", "", "New service name (required)")
	updateCmd.Flags().BoolVarP(&regenPass, "regen", "r", false, "Regenerate password as well")

	updateCmd.MarkFlagRequired("passkey")
	updateCmd.MarkFlagRequired("old")
	updateCmd.MarkFlagRequired("new")
}
