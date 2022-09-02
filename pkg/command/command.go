package command

import (
	"fmt"
	"os"
	"time"

	"github.com/go-kit/log/level"
	"github.com/promhippie/hetzner_exporter/pkg/action"
	"github.com/promhippie/hetzner_exporter/pkg/config"
	"github.com/promhippie/hetzner_exporter/pkg/version"
	"github.com/urfave/cli/v2"
)

// Run parses the command line arguments and executes the program.
func Run() error {
	cfg := config.Load()

	app := &cli.App{
		Name:    "hetzner_exporter",
		Version: version.String,
		Usage:   "Hetzner Exporter",
		Authors: []*cli.Author{
			{
				Name:  "Thomas Boerger",
				Email: "thomas@webhippie.de",
			},
		},
		Flags: RootFlags(cfg),
		Commands: []*cli.Command{
			Health(cfg),
		},
		Action: func(c *cli.Context) error {
			logger := setupLogger(cfg)

			if cfg.Target.Username == "" {
				level.Error(logger).Log(
					"msg", "Missing required hetzner.username",
				)

				return fmt.Errorf("missing required hetzner.username")
			}

			if cfg.Target.Password == "" {
				level.Error(logger).Log(
					"msg", "Missing required hetzner.password",
				)

				return fmt.Errorf("missing required hetzner.password")
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

	return app.Run(os.Args)
}

// RootFlags defines the available root flags.
func RootFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
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
			Value:       "0.0.0.0:9502",
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
		&cli.BoolFlag{
			Name:        "collector.servers",
			Value:       true,
			Usage:       "Enable collector for servers",
			EnvVars:     []string{"HETZNER_EXPORTER_COLLECTOR_SERVERS"},
			Destination: &cfg.Collector.Servers,
		},
		&cli.BoolFlag{
			Name:        "collector.ssh-keys",
			Value:       true,
			Usage:       "Enable collector for SSH keys",
			EnvVars:     []string{"HETZNER_EXPORTER_COLLECTOR_SSH_KEYS"},
			Destination: &cfg.Collector.SSHKeys,
		},
		&cli.BoolFlag{
			Name:        "collector.storageboxes",
			Value:       false,
			Usage:       "Enable collector for Storageboxes",
			EnvVars:     []string{"HETZNER_EXPORTER_COLLECTOR_STORAGEBOXES"},
			Destination: &cfg.Collector.Storageboxes,
		},
	}
}
