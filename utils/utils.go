package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
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

func HashPasskey(password string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	combined := append(salt, hash...)
	encoded := base64.RawStdEncoding.EncodeToString(combined)
	return encoded, nil
}

func VerifyPassword(encoded, password string) (bool, error) {
	data, err := base64.RawStdEncoding.DecodeString(encoded)
	if err != nil {
		return false, errors.New("invalid stored hash format")
	}

	if len(data) < 17 {
		return false, errors.New("invalid stored hash length")
	}

	salt := data[:16]
	storedHash := data[16:]

	newHash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	if subtle.ConstantTimeCompare(storedHash, newHash) == 1 {
		return true, nil
	}

	return false, nil
}
