package vault

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"time"

	"github.com/DiegoAndradeD/passkey-cli/utils"
)

var ErrServiceNotFound = errors.New("service not found")
var ErrVaultAlreadyExists = errors.New("vault already exists")
var ErrCopyToClipboardFailed = errors.New("failed to copy to clipboard")
var ErrServiceAlreadyExists = errors.New("service already is registered within vault")

type Service struct {
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type Vault struct {
	PasskeyHash string    `json:"passkey_hash"`
	Services    []Service `json:"services"`
}

func NewVault(passkeyHash string) *Vault {
	return &Vault{
		PasskeyHash: passkeyHash,
		Services:    []Service{},
	}
}

func LoadVault(path string, passkey string) (*Vault, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, errors.New("vault does not exist, please setup")
		}
		return nil, err
	}

	var vault Vault
	if err := json.Unmarshal(data, &vault); err != nil {
		return nil, err
	}

	if vault.PasskeyHash == "" {
		return nil, errors.New("vault is not initialized")
	}
	valid, err := utils.VerifyPassword(vault.PasskeyHash, passkey)

	if err != nil {
		return nil, fmt.Errorf("could not verify passkey: %w", err)
	}
	if !valid {
		return nil, errors.New("invalid passkey")
	}

	return &vault, nil
}

func SaveVault(path string, vault *Vault) error {
	data, err := json.MarshalIndent(vault, "", "  ")
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	return os.WriteFile(path, data, 0600)
}

func AddService(path string, name string, passkey string) error {
	vault, err := LoadVault(path, passkey)
	if err != nil {
		return fmt.Errorf("failed to load vault: %w", err)
	}

	password, err := utils.GeneratePassword()
	if err != nil {
		return fmt.Errorf("failed to generate password: %w", err)
	}

	service := Service{
		Name:      name,
		Password:  password,
		CreatedAt: time.Now(),
	}

	vault.Services, err = appendUniqueService(vault.Services, service)
	if err != nil {
		return err
	}

	if err := SaveVault(path, vault); err != nil {
		return fmt.Errorf("failed to save vault: %w", err)
	}

	return nil
}

func DeleteService(path, serviceName, passkey string) error {
	vault, err := LoadVault(path, passkey)
	if err != nil {
		return err
	}
	initialCount := len(vault.Services)
	vault.Services = slices.DeleteFunc(vault.Services, func(s Service) bool {
		return s.Name == serviceName
	})

	if len(vault.Services) == initialCount {
		return fmt.Errorf("service '%s' not found", serviceName)
	}

	if err := SaveVault(path, vault); err != nil {
		return err
	}

	return SaveVault(path, vault)
}

func UpdateService(path, oldName, newName, passkey string, regeneratePassword bool) error {
	vault, err := LoadVault(path, passkey)
	if err != nil {
		return fmt.Errorf("failed to load vault: %w", err)
	}

	index := -1
	for i, s := range vault.Services {
		if s.Name == oldName {
			index = i
			break
		}
	}
	if index == -1 {
		return ErrServiceNotFound
	}

	if oldName != newName {
		for _, s := range vault.Services {
			if s.Name == newName {
				return ErrServiceAlreadyExists
			}
		}
		vault.Services[index].Name = newName
	}

	if regeneratePassword {
		newPass, err := utils.GeneratePassword()
		if err != nil {
			return fmt.Errorf("failed to generate password: %w", err)
		}
		vault.Services[index].Password = newPass
	}

	if err := SaveVault(path, vault); err != nil {
		return fmt.Errorf("failed to save vault: %w", err)
	}

	return nil
}

func GetServices(path, passkey string) ([]Service, error) {
	vault, err := LoadVault(path, passkey)
	if err != nil {
		return nil, err
	}
	return vault.Services, nil
}

func GetService(path, name, passkey string) (Service, error) {
	vault, err := LoadVault(path, passkey)
	if err != nil {
		return Service{}, err
	}
	for _, s := range vault.Services {
		if s.Name == name {
			return s, nil
		}
	}
	return Service{}, ErrServiceNotFound
}

func CopyServicePassword(path, name, passkey string) error {
	service, err := GetService(path, name, passkey)
	if err != nil {
		return err
	}
	if service == (Service{}) {
		return fmt.Errorf("service %q not found", name)
	}

	if err := utils.CopyToClipboard(service.Password); err != nil {
		return fmt.Errorf("copy to clipboard failed: %w", err)
	}

	fmt.Println("Copied to clipboard!")
	return nil
}

func appendUniqueService(services []Service, newService Service) ([]Service, error) {
	for _, s := range services {
		if s.Name == newService.Name {
			return services, ErrServiceAlreadyExists
		}
	}
	return append(services, newService), nil
}
