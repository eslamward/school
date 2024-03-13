package config

import (
	"log"

	"github.com/spf13/viper"
)

func SetConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Can't Read Config File")
	}

	err := viper.Unmarshal(&appConfigration)

	if err != nil {
		log.Fatal("The Configration Didn't Parsed")
	}

}
