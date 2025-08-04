package main

import (
	"log"

	"vladislove-gRPC/internal/infrastructure/bootstrap"
)

func main() {
	cfg, bootstrapErr := bootstrap.ConfigFromEnv()
	if bootstrapErr != nil {
		log.Fatalf("произошла ошибка при прочтении конфигурации .env: %v", bootstrapErr)
	}

	if initErr := initApp(cfg); initErr != nil {
		log.Fatalf("произошла ошибка при инциализации приложения: %v", initErr)
		return
	}
}

func initApp(cfg *bootstrap.Config) error {
	newApp := bootstrap.NewApp(cfg)

	err := newApp.Run()
	if err != nil {
		return err
	}

	return nil
}
