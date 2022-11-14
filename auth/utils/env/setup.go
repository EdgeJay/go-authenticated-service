package env

import "github.com/spf13/viper"

func setupDefaults() {
	viper.SetDefault("PORT", "3000")
}

func LoadEnv() {
	setupDefaults()
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}
