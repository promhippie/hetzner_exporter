package exporter

import (
	"context"
	"log/slog"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/promhippie/hetzner_exporter/pkg/config"
	"github.com/promhippie/hetzner_exporter/pkg/internal/hetzner"
)

// SSHKeyCollector collects metrics about the SSH keys.
type SSHKeyCollector struct {
	client   *hetzner.Client
	logger   *slog.Logger
	failures *prometheus.CounterVec
	duration *prometheus.HistogramVec
	config   config.Target

	Key *prometheus.Desc
}

// NewSSHKeyCollector returns a new SSHKeyCollector.
func NewSSHKeyCollector(logger *slog.Logger, client *hetzner.Client, failures *prometheus.CounterVec, duration *prometheus.HistogramVec, cfg config.Target) *SSHKeyCollector {
	if failures != nil {
		failures.WithLabelValues("ssh_key").Add(0)
	}

	labels := []string{"name", "type", "size", "fingerprint"}
	return &SSHKeyCollector{
		client:   client,
		logger:   logger.With("collector", "ssh_key"),
		failures: failures,
		duration: duration,
		config:   cfg,

		Key: prometheus.NewDesc(
			"hetzner_ssh_key",
			"Information about SSH keys in your Hetzner robot",
			labels,
			nil,
		),
	}
}

// Metrics simply returns the list metric descriptors for generating a documentation.
func (c *SSHKeyCollector) Metrics() []*prometheus.Desc {
	return []*prometheus.Desc{
		c.Key,
	}
}

// Describe sends the super-set of all possible descriptors of metrics collected by this Collector.
func (c *SSHKeyCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.Key
}

// Collect is called by the Prometheus registry when collecting metrics.
func (c *SSHKeyCollector) Collect(ch chan<- prometheus.Metric) {
	ctx, cancel := context.WithTimeout(context.Background(), c.config.Timeout)
	defer cancel()

	now := time.Now()
	keys, err := c.client.SSHKey.All(ctx)
	c.duration.WithLabelValues("ssh_key").Observe(time.Since(now).Seconds())

	if err != nil {
		c.logger.Error("Failed to fetch SSH keys",
			"err", err,
		)

		c.failures.WithLabelValues("ssh_key").Inc()
		return
	}

	c.logger.Debug("Fetched SSH keys",
		"count", len(keys),
	)

	for _, key := range keys {
		labels := []string{
			key.Name,
			key.Type,
			key.Size,
			key.Fingerprint,
		}

		ch <- prometheus.MustNewConstMetric(
			c.Key,
			prometheus.GaugeValue,
			1.0,
			labels...,
		)
	}
}
