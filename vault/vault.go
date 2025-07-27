package vault

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

var ErrServiceNotFound = errors.New("service not found")

type Service struct {
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type Vault struct {
	Path     string
	Services []Service
}

func NewVault(path string) *Vault {
	return &Vault{
		Path: path,
	}
}

func (v *Vault) Load() error {
	data, err := os.ReadFile(v.Path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			v.Services = []Service{}
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &v.Services)
}

func (v *Vault) Save() error {
	data, err := json.MarshalIndent(v.Services, "", " ")
	if err != nil {
		return err
	}
	dir := filepath.Dir(v.Path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	return os.WriteFile(v.Path, data, 0600)
}

func (v *Vault) AddService(service Service) error {
	v.Services = append(v.Services, service)
	return v.Save()
}

func (v *Vault) GetServices() ([]Service, error) {
	if err := v.Load(); err != nil {
		return nil, err
	}
	return v.Services, nil
}

func (v *Vault) GetService(name string) (Service, error) {
	if err := v.Load(); err != nil {
		return Service{}, err
	}
	for i := range v.Services {
		if v.Services[i].Name == name {
			return v.Services[i], nil
		}
	}
	return Service{}, ErrServiceNotFound
}
