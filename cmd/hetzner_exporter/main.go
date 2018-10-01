package main

import (
	"errors"
	"os"
	"time"

	"github.com/go-kit/kit/log/level"
	"github.com/joho/godotenv"
	"github.com/promhippie/hetzner_exporter/pkg/action"
	"github.com/promhippie/hetzner_exporter/pkg/config"
	"github.com/promhippie/hetzner_exporter/pkg/version"
	"gopkg.in/urfave/cli.v2"
)

var (
	// ErrMissingHetznerUsername defines the error if hetzner.username is empty.
	ErrMissingHetznerUsername = errors.New("Missing required hetzner.username")

	// ErrMissingHetznerPassword defines the error if hetzner.password is empty.
	ErrMissingHetznerPassword = errors.New("Missing required hetzner.password")
)

func main() {
	cfg := config.Load()

	if env := os.Getenv("HETZNER_EXPORTER_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := &cli.App{
		Name:    "hetzner_exporter",
		Version: version.Version,
		Usage:   "Hetzner Exporter",
		Authors: []*cli.Author{
			{
				Name:  "Thomas Boerger",
				Email: "thomas@webhippie.de",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "log.level",
				Value:       "info",
				Usage:       "Only log messages with given severity",
				EnvVars:     []string{"HETZNER_EXPORTER_LOG_LEVEL"},
				Destination: &cfg.Logs.Level,
			},
			&cli.BoolFlag{
				Name:        "log.pretty",
				Value:       false,
				Usage:       "Enable pretty messages for logging",
				EnvVars:     []string{"HETZNER_EXPORTER_LOG_PRETTY"},
				Destination: &cfg.Logs.Pretty,
			},
			&cli.StringFlag{
				Name:        "web.address",
				Aliases:     []string{"web.listen-address"},
				Value:       "0.0.0.0:9107",
				Usage:       "Address to bind the metrics server",
				EnvVars:     []string{"HETZNER_EXPORTER_WEB_ADDRESS"},
				Destination: &cfg.Server.Addr,
			},
			&cli.StringFlag{
				Name:        "web.path",
				Aliases:     []string{"web.telemetry-path"},
				Value:       "/metrics",
				Usage:       "Path to bind the metrics server",
				EnvVars:     []string{"HETZNER_EXPORTER_WEB_PATH"},
				Destination: &cfg.Server.Path,
			},
			&cli.DurationFlag{
				Name:        "request.timeout",
				Value:       5 * time.Second,
				Usage:       "Request timeout as duration",
				EnvVars:     []string{"HETZNER_EXPORTER_REQUEST_TIMEOUT"},
				Destination: &cfg.Target.Timeout,
			},
			&cli.StringFlag{
				Name:        "hetzner.username",
				Value:       "",
				Usage:       "Username for the Hetzner API",
				EnvVars:     []string{"HETZNER_EXPORTER_USERNAME"},
				Destination: &cfg.Target.Username,
			},
			&cli.StringFlag{
				Name:        "hetzner.password",
				Value:       "",
				Usage:       "Password for the Hetzner API",
				EnvVars:     []string{"HETZNER_EXPORTER_PASSWORD"},
				Destination: &cfg.Target.Password,
			},
		},
		Action: func(c *cli.Context) error {
			logger := setupLogger(cfg)

			if cfg.Target.Username == "" {
				level.Error(logger).Log(
					"msg", ErrMissingHetznerUsername,
				)

				return ErrMissingHetznerUsername
			}

			if cfg.Target.Password == "" {
				level.Error(logger).Log(
					"msg", ErrMissingHetznerPassword,
				)

				return ErrMissingHetznerPassword
			}

			return action.Server(cfg, logger)
		},
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Show the help, so what you see now",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print the current version of that tool",
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
