package exporter

import (
	"strconv"
	"strings"
	"time"

	"github.com/appscode/go-hetzner"
	"github.com/dustin/go-humanize"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/promhippie/hetzner_exporter/pkg/config"
)

// ServerCollector collects metrics about the account in general.
type ServerCollector struct {
	client   *hetzner.Client
	logger   log.Logger
	failures *prometheus.CounterVec
	duration *prometheus.HistogramVec
	config   config.Target

	Up        *prometheus.Desc
	Traffic   *prometheus.Desc
	Paid      *prometheus.Desc
	Flatrate  *prometheus.Desc
	Throttled *prometheus.Desc
	Cancelled *prometheus.Desc
}

// NewServerCollector returns a new ServerCollector.
func NewServerCollector(logger log.Logger, client *hetzner.Client, failures *prometheus.CounterVec, duration *prometheus.HistogramVec, cfg config.Target) *ServerCollector {
	if failures != nil {
		failures.WithLabelValues("account").Add(0)
	}

	labels := []string{"id", "name", "datacenter"}
	return &ServerCollector{
		client:   client,
		logger:   log.With(logger, "collector", "server"),
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
		Paid: prometheus.NewDesc(
			"hetzner_server_paid_timestamp",
			"Timestamp of the date until server is paid",
			labels,
			nil,
		),
		Flatrate: prometheus.NewDesc(
			"hetzner_server_flatrate",
			"If 1 the server got a flatrate enabled, 0 otherwise",
			labels,
			nil,
		),
		Throttled: prometheus.NewDesc(
			"hetzner_server_throttled",
			"If 1 the server is in a throttled state, 0 otherwise",
			labels,
			nil,
		),
		Cancelled: prometheus.NewDesc(
			"hetzner_server_cancelled",
			"If 1 the server have been cancelled, 0 otherwise",
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
		c.Paid,
		c.Flatrate,
		c.Throttled,
		c.Cancelled,
	}
}

// Describe sends the super-set of all possible descriptors of metrics collected by this Collector.
func (c *ServerCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.Up
	ch <- c.Traffic
	ch <- c.Paid
	ch <- c.Flatrate
	ch <- c.Throttled
	ch <- c.Cancelled
}

// Collect is called by the Prometheus registry when collecting metrics.
func (c *ServerCollector) Collect(ch chan<- prometheus.Metric) {
	now := time.Now()
	servers, _, err := c.client.WithTimeout(c.config.Timeout).Server.ListServers()
	c.duration.WithLabelValues("server").Observe(time.Since(now).Seconds())

	if err != nil {
		level.Error(c.logger).Log(
			"msg", "Failed to fetch servers",
			"err", err,
		)

		c.failures.WithLabelValues("server").Inc()
		return
	}

	level.Debug(c.logger).Log(
		"msg", "Fetched servers",
		"count", len(servers),
	)

	for _, server := range servers {
		var (
			up        float64
			traffic   float64
			paid      float64
			flatrate  float64
			throttled float64
			cancelled float64
		)

		labels := []string{
			strconv.Itoa(server.ServerNumber),
			server.ServerName,
			strings.ToLower(server.Dc),
		}

		if server.Status == "ready" {
			up = 1.0
		}

		ch <- prometheus.MustNewConstMetric(
			c.Up,
			prometheus.GaugeValue,
			up,
			labels...,
		)

		if num, err := humanize.ParseBytes(server.Traffic); err == nil {
			traffic = float64(num)
		}

		ch <- prometheus.MustNewConstMetric(
			c.Traffic,
			prometheus.GaugeValue,
			traffic,
			labels...,
		)

		if num, err := time.Parse("2006-01-02", server.PaidUntil); err == nil {
			paid = float64(num.Unix())
		}

		ch <- prometheus.MustNewConstMetric(
			c.Paid,
			prometheus.GaugeValue,
			paid,
			labels...,
		)

		if server.Flatrate {
			flatrate = 1.0
		}

		ch <- prometheus.MustNewConstMetric(
			c.Flatrate,
			prometheus.GaugeValue,
			flatrate,
			labels...,
		)

		if server.Throttled {
			throttled = 1.0
		}

		ch <- prometheus.MustNewConstMetric(
			c.Throttled,
			prometheus.GaugeValue,
			throttled,
			labels...,
		)

		if server.Cancelled {
			cancelled = 1.0
		}

		ch <- prometheus.MustNewConstMetric(
			c.Cancelled,
			prometheus.GaugeValue,
			cancelled,
			labels...,
		)
	}
}
