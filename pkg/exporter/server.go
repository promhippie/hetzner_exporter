package exporter

import (
	"context"
	"log/slog"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/promhippie/hetzner_exporter/pkg/config"
	"github.com/promhippie/hetzner_exporter/pkg/internal/hetzner"
)

// ServerCollector collects metrics about the account in general.
type ServerCollector struct {
	client   *hetzner.Client
	logger   *slog.Logger
	failures *prometheus.CounterVec
	duration *prometheus.HistogramVec
	config   config.Target

	Up        *prometheus.Desc
	Traffic   *prometheus.Desc
	Flatrate  *prometheus.Desc
	Cancelled *prometheus.Desc
	Paid      *prometheus.Desc
}

// NewServerCollector returns a new ServerCollector.
func NewServerCollector(logger *slog.Logger, client *hetzner.Client, failures *prometheus.CounterVec, duration *prometheus.HistogramVec, cfg config.Target) *ServerCollector {
	if failures != nil {
		failures.WithLabelValues("account").Add(0)
	}

	labels := []string{"id", "name", "datacenter"}
	return &ServerCollector{
		client:   client,
		logger:   logger.With("collector", "server"),
		failures: failures,
		duration: duration,
		config:   cfg,

		Up: prometheus.NewDesc(
			"hetzner_server_running",
			"If 1 the server is running, 0 otherwise",
			labels,
			nil,
		),
		Traffic: prometheus.NewDesc(
			"hetzner_server_traffic_bytes",
			"Amount of included traffic for the server",
			labels,
			nil,
		),
		Flatrate: prometheus.NewDesc(
			"hetzner_server_flatrate",
			"If 1 the server got a flatrate enabled, 0 otherwise",
			labels,
			nil,
		),
		Cancelled: prometheus.NewDesc(
			"hetzner_server_cancelled",
			"If 1 the server have been cancelled, 0 otherwise",
			labels,
			nil,
		),
		Paid: prometheus.NewDesc(
			"hetzner_server_paid_timestamp",
			"Timestamp of the date until server is paid",
			labels,
			nil,
		),
	}
}

// Metrics simply returns the list metric descriptors for generating a documentation.
func (c *ServerCollector) Metrics() []*prometheus.Desc {
	return []*prometheus.Desc{
		c.Up,
		c.Traffic,
		c.Flatrate,
		c.Cancelled,
		c.Paid,
	}
}

// Describe sends the super-set of all possible descriptors of metrics collected by this Collector.
func (c *ServerCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.Up
	ch <- c.Traffic
	ch <- c.Flatrate
	ch <- c.Cancelled
	ch <- c.Paid
}

// Collect is called by the Prometheus registry when collecting metrics.
func (c *ServerCollector) Collect(ch chan<- prometheus.Metric) {
	ctx, cancel := context.WithTimeout(context.Background(), c.config.Timeout)
	defer cancel()

	now := time.Now()
	servers, err := c.client.Server.All(ctx)
	c.duration.WithLabelValues("server").Observe(time.Since(now).Seconds())

	if err != nil {
		c.logger.Error("Failed to fetch servers",
			"err", err,
		)

		c.failures.WithLabelValues("server").Inc()
		return
	}

	c.logger.Debug("Fetched servers",
		"count", len(servers),
	)

	for _, server := range servers {
		labels := []string{
			server.Number,
			server.Name,
			server.Datacenter,
		}

		ch <- prometheus.MustNewConstMetric(
			c.Up,
			prometheus.GaugeValue,
			server.Status,
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Traffic,
			prometheus.GaugeValue,
			server.Traffic,
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Flatrate,
			prometheus.GaugeValue,
			server.Flatrate,
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Cancelled,
			prometheus.GaugeValue,
			server.Cancelled,
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Paid,
			prometheus.GaugeValue,
			server.Paid,
			labels...,
		)
	}
}
