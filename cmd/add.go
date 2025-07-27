/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"time"

	"github.com/DiegoAndradeD/passkey-cli/utils"
	"github.com/DiegoAndradeD/passkey-cli/vault"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new service with automatically generated password",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil || name == "" {
			log.Fatal("the service name must be specified with --name", err)
		}
		createService(name)
	},
}

func createService(name string) {
	vaultPath := utils.GetVaultPath()
	newVault := vault.NewVault(vaultPath)
	if err := newVault.Load(); err != nil {
		log.Fatal("Failed to load vault", err)
	}

	password, err := utils.GeneratePassword()
	if err != nil {
		log.Fatal("Failed to generate password", err)
	}

	service := vault.Service{
		Name:      name,
		Password:  password,
		CreatedAt: time.Now(),
	}
	if err := newVault.AddService(service); err != nil {
		log.Fatal("Failed to save vault", err)
	}
	log.Printf("Service %q added successfully!\n", name)
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("name", "n", "", "Service name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
