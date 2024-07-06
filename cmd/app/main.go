package main

import (
	"log"

	"github.com/Ndraaa15/IQuest/internal/app/bootstrap"
	"github.com/Ndraaa15/IQuest/pkg/env"
)

func main() {
	if err := env.LoadEnv(".../../.env"); err != nil {
		log.Fatalf("[IQuest] Failed to load env: %v", err)
	}

	app, err := bootstrap.NewBootstrap()
	if err != nil {
		log.Fatalf("[IQuest] Failed to initialize app: %v", err)
	}

	app.RegisterHandler()
	if err := app.Run(); err != nil {
		log.Fatalf("[IQuest] Failed to run app: %v", err)
	}
}
