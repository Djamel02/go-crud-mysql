package dbconfig

import (
	"log"

	"github.com/spf13/viper"
)

func GetEnvironmentVars(propName string) string {
	viper.SetConfigFile("app_dev.env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(propName).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}
