package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	Address           string `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey string `mapstructure:"JWT_SECRET_KEY"`
	Duration          string `mapstructure:"JWT_EXPIRY"`
}

func LoadConfig(path string) (config Config, err error) {
	// viper.AddConfigPath(path)
	// viper.SetConfigFile("app")
	// viper.SetConfigType("env")
	viper.SetConfigFile(path)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

}
