package utils

import (
	"crypto/rand"
	"log"
	"math/big"
	"strings"
)

const (
	passwordLength = 18
	passwordChars  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*"
)

func GetVaultPath() string {
	// home, err := os.UserHomeDir()
	// if err != nil {
	// 	log.Fatal("Could not get user's home directory: ", err)
	// }
	// return filepath.Join(home, ".config", "passkey-cli", "vault.json")

	// Test Path
	return "vault.json"
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GeneratePassword() (string, error) {
	var password strings.Builder
	for range passwordLength {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(passwordChars))))
		if err != nil {
			return "", err
		}
		password.WriteByte(passwordChars[index.Int64()])
	}
	return password.String(), nil
}
