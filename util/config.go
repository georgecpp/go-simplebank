package util

import "github.com/spf13/viper"

// Config stores all configurations of the app
// the values are read by viper from a config file or env variables
type Config struct {
	DBDriver		string
	DBSource		string
	ServerAddress	string
}

// LoadConfig reads config from file or env variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}