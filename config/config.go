package config

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func LoadConfig() error {
	dotEnvPath := ".env"
	_, err := os.Stat(dotEnvPath)
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(dotEnvPath)
	if err != nil {
		exampleDotEnvPath := ".env.example"
		_, exampleDotEnvPathErr := os.Stat(exampleDotEnvPath)
		if exampleDotEnvPathErr != nil {
			log.Fatal(exampleDotEnvPathErr)
		}

		exampleEnvLoadErr := godotenv.Load(exampleDotEnvPath)
		if exampleEnvLoadErr != nil {
			log.Fatal(errors.Join(exampleEnvLoadErr, err))
		}
	}

	appPrefix := os.Getenv("DCV_APP_PREFIX")
	if appPrefix == "" {
		log.Fatalln("DCV_APP_PREFIX environment variable not set")
	}

	viper.SetEnvPrefix(appPrefix)
	viper.AutomaticEnv()

	viper.SetDefault("DCV_ENV", "local")

	return nil
}
