package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"math/big"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	passwordLength = 18
	passwordChars  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*"
)

func GetVaultPath() string {
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

func HashPasskey(password string) (string, []byte) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		log.Fatal(err)
	}
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	encoded := base64.RawStdEncoding.EncodeToString(append(salt, hash...))
	return encoded, salt
}

func VerifyPassword(encoded, password string) bool {
	data, _ := base64.RawStdEncoding.DecodeString(encoded)
	salt := data[:16]
	hash := data[16:]

	newHash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	return string(hash) == string(newHash)
}
