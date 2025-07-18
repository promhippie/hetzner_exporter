package command

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/promhippie/hetzner_exporter/pkg/action"
	"github.com/promhippie/hetzner_exporter/pkg/config"
	"github.com/promhippie/hetzner_exporter/pkg/version"
	"github.com/urfave/cli/v3"
)

// Run parses the command line arguments and executes the program.
func Run() error {
	cfg := config.Load()

	app := &cli.Command{
		Name:    "hetzner_exporter",
		Version: version.String,
		Usage:   "Hetzner Exporter",
		Authors: []any{
			"Thomas Boerger <thomas@webhippie.de>",
		},
		Flags: RootFlags(cfg),
		Commands: []*cli.Command{
			Health(cfg),
		},
		Action: func(_ context.Context, _ *cli.Command) error {
			logger := setupLogger(cfg)

			if cfg.Target.Username == "" {
				logger.Error("Missing required hetzner.username")
				return fmt.Errorf("missing required hetzner.username")
			}

			if cfg.Target.Password == "" {
				logger.Error("Missing required hetzner.password")
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

	return app.Run(context.Background(), os.Args)
}

// RootFlags defines the available root flags.
func RootFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "log.level",
			Value:       "info",
			Usage:       "Only log messages with given severity",
			Sources:     cli.EnvVars("HETZNER_EXPORTER_LOG_LEVEL"),
			Destination: &cfg.Logs.Level,
		},
		&cli.BoolFlag{
			Name:        "log.pretty",
			Value:       false,
			Usage:       "Enable pretty messages for logging",
			Sources:     cli.EnvVars("HETZNER_EXPORTER_LOG_PRETTY"),
			Destination: &cfg.Logs.Pretty,
		},
		&cli.StringFlag{
			Name:        "web.address",
			Aliases:     []string{"web.listen-address"},
			Value:       "0.0.0.0:9502",
			Usage:       "Address to bind the metrics server",
			Sources:     cli.EnvVars("HETZNER_EXPORTER_WEB_ADDRESS"),
			Destination: &cfg.Server.Addr,
		},
		&cli.StringFlag{
			Name:        "web.path",
			Aliases:     []string{"web.telemetry-path"},
			Value:       "/metrics",
			Usage:       "Path to bind the metrics server",
			Sources:     cli.EnvVars("HETZNER_EXPORTER_WEB_PATH"),
			Destination: &cfg.Server.Path,
		},
		&cli.BoolFlag{
			Name:        "web.debug",
			Value:       false,
			Usage:       "Enable pprof debugging for server",
			Sources:     cli.EnvVars("HETZNER_EXPORTER_WEB_PPROF"),
			Destination: &cfg.Server.Pprof,
		},
		&cli.DurationFlag{
			Name:        "web.timeout",
			Value:       10 * time.Second,
			Usage:       "Server metrics endpoint timeout",
			Sources:     cli.EnvVars("HETZNER_EXPORTER_WEB_TIMEOUT"),
			Destination: &cfg.Server.Timeout,
		},
		&cli.StringFlag{
			Name:        "web.config",
			Value:       "",
			Usage:       "Path to web-config file",
			Sources:     cli.EnvVars("HETZNER_EXPORTER_WEB_CONFIG"),
			Destination: &cfg.Server.Web,
		},
		&cli.DurationFlag{
			Name:        "request.timeout",
			Value:       5 * time.Second,
			Usage:       "Request timeout as duration",
			Sources:     cli.EnvVars("HETZNER_EXPORTER_REQUEST_TIMEOUT"),
			Destination: &cfg.Target.Timeout,
		},
		&cli.StringFlag{
			Name:        "hetzner.username",
			Value:       "",
			Usage:       "Username for the Hetzner API",
			Sources:     cli.EnvVars("HETZNER_EXPORTER_USERNAME"),
			Destination: &cfg.Target.Username,
		},
		&cli.StringFlag{
			Name:        "hetzner.password",
			Value:       "",
			Usage:       "Password for the Hetzner API",
			Sources:     cli.EnvVars("HETZNER_EXPORTER_PASSWORD"),
			Destination: &cfg.Target.Password,
		},
		&cli.BoolFlag{
			Name:        "collector.servers",
			Value:       true,
			Usage:       "Enable collector for servers",
			Sources:     cli.EnvVars("HETZNER_EXPORTER_COLLECTOR_SERVERS"),
			Destination: &cfg.Collector.Servers,
		},
		&cli.BoolFlag{
			Name:        "collector.ssh-keys",
			Value:       true,
			Usage:       "Enable collector for SSH keys",
			Sources:     cli.EnvVars("HETZNER_EXPORTER_COLLECTOR_SSH_KEYS"),
			Destination: &cfg.Collector.SSHKeys,
		},
	}
}
