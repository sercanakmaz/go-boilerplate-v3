package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func Load(config interface{}) (err error) {

	filePath := fmt.Sprintf("%s/%s.json", "./configs", os.Getenv("ENV"))
	viper.SetConfigFile(filePath)

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
