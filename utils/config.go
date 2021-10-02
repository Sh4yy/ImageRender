package utils

import (
	"github.com/spf13/viper"
	"strings"
)

type ServerConfig struct {
	Host string
	Port int
}

type RenderConfig struct {
	Width int
	Height int
	Quality int
	Format string
}

type DirectoryConfig struct {
	Static string
	Templates string
}

type Configuration struct {
	Server       ServerConfig
	Debug        bool
	Render		 RenderConfig
	Directory 	 DirectoryConfig
}

var Config *Configuration

func NewConfig() *Configuration {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../")
	viper.SetConfigType("yml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var config Configuration
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return &config

}

func init() {
	Config = NewConfig()
}