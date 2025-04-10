package config

import "github.com/spf13/viper"

type Config struct {
	Port                string `mapstructure:"PORT"`
	DBDriver            string `mapstructure:"DB_DRIVER"`
	DBSourceDevelopment string `mapstructure:"DB_SOURCE_DEVELOPMENT"`
	DBSourceProduction  string `mapstructure:"DB_SOURCE_PRODUCTION"`
	File_Svc_Url        string `mapstructure:"FILE_SVC_URL"`
	File_Svc_Port       string `mapstructure:"FILE_SVC_PORT"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
