package config

import (
	"digital-cash-vault/db"
	applicationUsers "digital-cash-vault/internal/applications/users"
	"digital-cash-vault/internal/infrastructure/users/repositories"
	"digital-cash-vault/internal/interfaces/http"
	httpUsers "digital-cash-vault/internal/interfaces/http/users"
	"digital-cash-vault/pkg/jwt"
	"digital-cash-vault/pkg/static"
	"digital-cash-vault/router"
	"errors"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"syscall"
)

type AppConfig struct {
	DB         db.DB
	Router     router.Router
	ServerPort string
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
	appPrefix := os.Getenv(static.EnvAppPrefix)
	if appPrefix == "" {
		log.Fatalf("%s environment variable not set", static.EnvAppPrefix)
	}
	viper.SetEnvPrefix(appPrefix)
	viper.AutomaticEnv()
	viper.SetDefault(static.EnvServiceEnvironment, static.ServiceEnvironmentLocal)

	serverPort := viper.GetString(static.EnvServerPort)
	if serverPort == "" {
		serverPort = "8080"
	}

	jwtSecret := viper.GetString(static.EnvJwtSecret)
	jwt.Secret = []byte(jwtSecret)

	// database
	dbConfig := db.ConfigFromEnv()
	dbConn, err := db.Open(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	// dependencies
	userRepository := repositories.NewUserRepository(dbConn)
	userService := applicationUsers.NewUserService(userRepository)
	userHandler := httpUsers.NewUserHandler(userService)

	// router
	handlers := map[string]any{
		static.HandlerUser: userHandler,
	}
	rootRouter, err := http.InitRouter(os.Getenv(static.EnvRouter), handlers)
	if err != nil {
		log.Fatal(err)
	}

	return &AppConfig{
		DB:         dbConn,
		Router:     rootRouter,
		ServerPort: serverPort,
	}, nil
}
