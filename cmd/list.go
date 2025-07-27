package cmd

import (
	"encoding/json"
	"fmt"
	"log"

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
  passkey-cli list
  passkey-cli list --name "my-database"

This command is useful to quickly check which services are stored in your vault
and retrieve details about individual services without modifying any data.`,
	Run: func(cmd *cobra.Command, args []string) {
		service, err := cmd.Flags().GetString("service")
		if err != nil {
			log.Fatal(err)
		}
		if service == "" {
			listAllServices()
		} else {
			listServiceByName(service)
		}

	},
}

func listAllServices() {
	vaultPath := utils.GetVaultPath()
	v := vault.NewVault(vaultPath)
	services, err := v.GetServices()
	if err != nil {
		log.Fatal("Failed to load vault services", err)
	}
	output, err := json.MarshalIndent(services, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}

func listServiceByName(name string) {
	vaultPath := utils.GetVaultPath()
	v := vault.NewVault(vaultPath)
	service, err := v.GetService(name)
	if err != nil {
		log.Fatal("Failed to load service", err)
	}
	output, err := json.MarshalIndent(service, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("service", "s", "", "Service name")
}
