package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/promhippie/hetzner_exporter/pkg/command"
)

func main() {
	if env := os.Getenv("HETZNER_EXPORTER_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}
