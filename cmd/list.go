package cmd

import (
	"errors"
	"log"
	"time"

	"github.com/DiegoAndradeD/passkey-cli/utils"
	"github.com/DiegoAndradeD/passkey-cli/vault"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all stored services or details for a specific service",
	Long: `The "list" command retrieves and displays services stored in the secure vault.
By default, it lists all stored services, showing their names and basic information.

You can also specify a service name to display only the details of that specific service.

Usage:
  passkey-cli list                # Lists all services
  passkey-cli list --name github  # Shows details only for the 'github' service

Examples:
  passkey-cli list --passkey password
  passkey-cli list --name "my-database -p password"

This command is useful to quickly check which services are stored in your vault
and retrieve details about individual services without modifying any data.`,
	Run: func(cmd *cobra.Command, args []string) {
		service, err := cmd.Flags().GetString("service")
		if err != nil {
			log.Fatal(err)
		}
		passkey, err := cmd.Flags().GetString("passkey")
		if err != nil || passkey == "" {
			log.Fatal("you must specify the passkey with --passkey flag")
		}
		if service == "" {
			listAllServices(passkey)
		} else {
			listServiceByName(service, passkey)
		}

	},
}

func listAllServices(passkey string) {
	vaultPath := utils.GetVaultPath()

	services, err := vault.GetServices(vaultPath, passkey)
	if err != nil {
		log.Fatalf("failed to get services: %v", err)
	}

	if len(services) == 0 {
		log.Println("No services found in the vault.")
		return
	}

	log.Println("Services stored in vault:")
	for _, s := range services {
		log.Printf("- %s (created at: %s)", s.Name, s.CreatedAt.Format(time.RFC1123))
	}
}

func listServiceByName(name, passkey string) {
	vaultPath := utils.GetVaultPath()

	service, err := vault.GetService(vaultPath, name, passkey)
	if err != nil {
		if errors.Is(err, vault.ErrServiceNotFound) {
			log.Printf("Service '%s' not found in the vault.", name)
			return
		}
		log.Fatalf("failed to get service: %v", err)
	}

	log.Printf("Service: %s\nPassword: %s\nCreated At: %s\n", service.Name, service.Password, service.CreatedAt.Format(time.RFC1123))
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("service", "s", "", "Service name")
	listCmd.Flags().StringP("passkey", "p", "", "Vault passkey")
}
