package config

import (
	"log"

	"example.com/svelte-go/database"
	"github.com/spf13/viper"
)

type Config struct {
	Host string
	Port string
	CertFile string
	KeyFile string
	DBConfig database.Config
}
// LoadConfig loads the configuration from the config file
func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 443)
	viper.SetDefault("certFile", "cert.pem")
	viper.SetDefault("keyFile", "key.pem")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	return Config{
		Host: viper.GetString("host"),
		Port: viper.GetString("port"),
		CertFile: viper.GetString("certFile"),
		KeyFile: viper.GetString("keyFile"),
		DBConfig: database.Config{
			Host: viper.GetString("database.host"),
			Port: viper.GetInt("database.port"),
			User: viper.GetString("database.user"),
			Password: viper.GetString("database.password"),
			Name: viper.GetString("database.name"),
		},
	}
}