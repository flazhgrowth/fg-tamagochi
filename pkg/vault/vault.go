package vault

import (
	"encoding/json"
	"os"

	"github.com/flazhgrowth/fg-tamagotchi/pkg/config"
	"github.com/spf13/viper"
)

var (
	vault *Vault
)

type Vault struct {
	Database DatabaseVault `json:"database"`
	app      *viper.Viper  `json:"-"`
}

type (
	DatabaseVault struct {
		Driver    string `json:"driver"`
		WriterDSN string `json:"writer_dsn"`
		ReaderDSN string `json:"reader_dsn"`
	}
)

func New() error {
	vaultPath := "./etc/vault/vault.json"
	f, err := os.Open(vaultPath)
	if err != nil {
		return err
	}
	if config.GetConfig().IsEnvProduction() {
		// fetch from secret manager or something similar
	}

	vault = &Vault{}
	if err = json.NewDecoder(f).Decode(vault); err != nil {
		return err
	}

	v := viper.New()
	v.SetConfigFile(vaultPath)
	if err = v.ReadInConfig(); err != nil {
		return err
	}
	vault.app = v

	return nil
}

func GetVault() *Vault {
	return vault
}
