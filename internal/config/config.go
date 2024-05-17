package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type AppConfig struct {
	Environment string   `mapstructure:"environment"`
	ServerPort  int      `mapstructure:"server_port"`
	Logging     Logging  `mapstructure:"logging"`
	Db          DbConfig `mapstructure:"db"`
}

type Logging struct {
	FileName string `mapstructure:"filename"`
	Level    string `mapstructure:"level"`
}

type DbConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

func GetConfig(configFilePath string) (AppConfig, error) {
	log.Printf("Config File Path: %v", configFilePath)
	conf := viper.New()
	conf.SetConfigFile(configFilePath)
	replacer := strings.NewReplacer("_", ":")
	conf.SetEnvKeyReplacer(replacer)
	conf.AutomaticEnv()

	err := conf.ReadInConfig()
	if err != nil {
		fmt.Printf("error reading config from path: %v", err)
	}

	var cfg AppConfig

	err = conf.Unmarshal(&cfg)
	if err != nil {
		fmt.Printf("error unmarshalling config: %v", err)
	}

	return cfg, nil
}
