package main

import (
	"javacode-test-task/app/internal/app"
	"javacode-test-task/app/internal/config"
	"javacode-test-task/app/pkg/logger"
	"time"
)

func main() {
	time.Sleep(10 * time.Second)
	cfg := config.MustLoad()
	log := logger.SetupLogger()
	a, err := app.NewApp(cfg, log)
	if err != nil {
		log.Error("failed to initialize app", "error", err)
		panic(err)
	}

	a.StartHTTP()
}
