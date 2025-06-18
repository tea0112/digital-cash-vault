package config

import (
	"digital-cash-vault/internal/db"
	userHandlers "digital-cash-vault/internal/infrastructure/users/handlers"
	"digital-cash-vault/internal/infrastructure/users/repositories"
	"digital-cash-vault/pkg/static"
	"errors"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"syscall"
)

type AppConfig struct {
	DB db.DB
}

func LoadConfig() (*AppConfig, error) {
	dotEnvPath := ".env"
	_, err := os.Stat(dotEnvPath)
	if err != nil && errors.Is(err, syscall.ENOENT) {
		dotEnvPath = ".env.example"
		_, exampleDotEnvPathErr := os.Stat(dotEnvPath)
		if exampleDotEnvPathErr != nil {
			log.Fatal(exampleDotEnvPathErr)
		}
	} else if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(dotEnvPath)
	if err != nil {
		log.Fatal(err)
	}

	// config env with viper
	appPrefix := os.Getenv("DCV_APP_PREFIX")
	if appPrefix == "" {
		log.Fatalln("DCV_APP_PREFIX environment variable not set")
	}
	viper.SetEnvPrefix(appPrefix)
	viper.AutomaticEnv()
	viper.SetDefault("DCV_ENV", "local")

	// database
	dbConfig := db.ConfigFromEnv()
	dbConn, err := db.Open(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	// dependencies
	_ = repositories.NewUserRepo(dbConn)

	// router
	_, err = userHandlers.InitRouter(os.Getenv(static.EnvDcvRouter))
	if err != nil {
		log.Fatal(err)
	}

	return &AppConfig{
		DB: dbConn,
	}, nil
}
