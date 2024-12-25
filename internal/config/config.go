package config

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/agniBit/cryptonian/model/cfg"
	"github.com/spf13/viper"
)

var config *cfg.Config
var initOnce sync.Once

func LoadConfig() *cfg.Config {
	if config == nil {
		initOnce.Do(load)
	}
	return config
}

func GetConfig() *cfg.Config {
	return LoadConfig()
}

func load() {
	c := &cfg.Config{}
	if len(os.Getenv("ENCRYPTION_SECRET_KEY")) != 32 {
		panic("invalid enrcypion secret key")
	}

	configFileName := os.Getenv("CONFIG_FILE")
	if configFileName == "" {
		configFileName = "cmd/config/local.yaml"
	}
	fmt.Println("loading config from file: ", configFileName)
	viper.SetConfigFile(configFileName)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		panic(err)
	}
	if err := viper.Unmarshal(c); err != nil {
		fmt.Printf("Error unmarshing file %s", err)
		panic(err)
	}

	config = c
}
