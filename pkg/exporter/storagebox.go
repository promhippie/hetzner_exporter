package exporter

import (
	"context"
	"log/slog"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/promhippie/hetzner_exporter/pkg/config"
	"github.com/promhippie/hetzner_exporter/pkg/internal/hetzner"
)

// StorageboxCollector collects metrics about the SSH keys.
type StorageboxCollector struct {
	client   *hetzner.Client
	logger   *slog.Logger
	failures *prometheus.CounterVec
	duration *prometheus.HistogramVec
	config   config.Target

	Cancelled *prometheus.Desc
	Locked    *prometheus.Desc
	Paid      *prometheus.Desc
	Quota     *prometheus.Desc
	Usage     *prometheus.Desc
	Data      *prometheus.Desc
	Snapshots *prometheus.Desc
	Webdav    *prometheus.Desc
	Samba     *prometheus.Desc
	SSH       *prometheus.Desc
	External  *prometheus.Desc
	ZFS       *prometheus.Desc
}

// NewStorageboxCollector returns a new StorageboxCollector.
func NewStorageboxCollector(logger *slog.Logger, client *hetzner.Client, failures *prometheus.CounterVec, duration *prometheus.HistogramVec, cfg config.Target) *StorageboxCollector {
	if failures != nil {
		failures.WithLabelValues("storagebox").Add(0)
	}

	labels := []string{"id", "name", "location", "login"}
	return &StorageboxCollector{
		client:   client,
		logger:   logger.With("collector", "storagebox"),
		failures: failures,
		duration: duration,
		config:   cfg,

		Cancelled: prometheus.NewDesc(
			"hetzner_storagebox_cancelled",
			"If 1 the storagebox have been cancelled, 0 otherwise",
			labels,
			nil,
		),
		Locked: prometheus.NewDesc(
			"hetzner_storagebox_locked",
			"If 1 the storagebox have been locked, 0 otherwise",
			labels,
			nil,
		),
		Paid: prometheus.NewDesc(
			"hetzner_storagebox_paid",
			"Timestamp of the date until storagebox is paid",
			labels,
			nil,
		),
		Quota: prometheus.NewDesc(
			"hetzner_storagebox_quota",
			"Available storage for the storagebox in MB",
			labels,
			nil,
		),
		Usage: prometheus.NewDesc(
			"hetzner_storagebox_usage",
			"Used storage for the storagebox in MB",
			labels,
			nil,
		),
		Data: prometheus.NewDesc(
			"hetzner_storagebox_data",
			"Used storage by files for the storagebox in MB",
			labels,
			nil,
		),
		Snapshots: prometheus.NewDesc(
			"hetzner_storagebox_snapshots",
			"Used storage by snapshots for the storagebox in MB",
			labels,
			nil,
		),
		Webdav: prometheus.NewDesc(
			"hetzner_storagebox_webdav",
			"If 1 the storagebox can be accessed via webdav, 0 otherwise",
			labels,
			nil,
		),
		Samba: prometheus.NewDesc(
			"hetzner_storagebox_samba",
			"If 1 the storagebox can be accessed via samba, 0 otherwise",
			labels,
			nil,
		),
		SSH: prometheus.NewDesc(
			"hetzner_storagebox_ssh",
			"If 1 the storagebox can be accessed via ssh, 0 otherwise",
			labels,
			nil,
		),
		External: prometheus.NewDesc(
			"hetzner_storagebox_external",
			"If 1 the storagebox can be accessed from external, 0 otherwise",
			labels,
			nil,
		),
		ZFS: prometheus.NewDesc(
			"hetzner_storagebox_zfs",
			"If 1 the zfs directory is visible, 0 otherwise",
			labels,
			nil,
		),
	}
}

// Metrics simply returns the list metric descriptors for generating a documentation.
func (c *StorageboxCollector) Metrics() []*prometheus.Desc {
	return []*prometheus.Desc{
		c.Cancelled,
		c.Locked,
		c.Paid,
		c.Quota,
		c.Usage,
		c.Data,
		c.Snapshots,
		c.Webdav,
		c.Samba,
		c.SSH,
		c.External,
		c.ZFS,
	}
}

// Describe sends the super-set of all possible descriptors of metrics collected by this Collector.
func (c *StorageboxCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.Cancelled
	ch <- c.Locked
	ch <- c.Paid
	ch <- c.Quota
	ch <- c.Usage
	ch <- c.Data
	ch <- c.Snapshots
	ch <- c.Webdav
	ch <- c.Samba
	ch <- c.SSH
	ch <- c.External
	ch <- c.ZFS
}

// Collect is called by the Prometheus registry when collecting metrics.
func (c *StorageboxCollector) Collect(ch chan<- prometheus.Metric) {
	ctx, cancel := context.WithTimeout(context.Background(), c.config.Timeout)
	defer cancel()

	now := time.Now()
	storageboxes, err := c.client.Storagebox.All(ctx)
	c.duration.WithLabelValues("storagebox").Observe(time.Since(now).Seconds())

	if err != nil {
		c.logger.Error("Failed to fetch storageboxes",
			"err", err,
		)

		c.failures.WithLabelValues("storagebox").Inc()
		return
	}

	c.logger.Debug("Fetched storageboxes",
		"count", len(storageboxes),
	)

	for _, storagebox := range storageboxes {
		ctx, cancel := context.WithTimeout(context.Background(), c.config.Timeout)
		defer cancel()

		now := time.Now()
		record, err := c.client.Storagebox.Get(ctx, storagebox.Number)
		c.duration.WithLabelValues("storagebox").Observe(time.Since(now).Seconds())

		if err != nil {
			c.logger.Error("Failed to fetch storagebox",
				"number", storagebox.Number,
				"err", err,
			)

			c.failures.WithLabelValues("storagebox").Inc()
			return
		}

		labels := []string{
			record.Number,
			record.Name,
			record.Location,
			record.Login,
		}

		ch <- prometheus.MustNewConstMetric(
			c.Cancelled,
			prometheus.GaugeValue,
			record.Cancelled,
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Locked,
			prometheus.GaugeValue,
			record.Locked,
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Paid,
			prometheus.GaugeValue,
			record.Paid,
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Quota,
			prometheus.GaugeValue,
			record.Quota,
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Usage,
			prometheus.GaugeValue,
			record.Usage,
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Data,
			prometheus.GaugeValue,
			record.Data,
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Snapshots,
			prometheus.GaugeValue,
			record.Snapshots,
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Webdav,
			prometheus.GaugeValue,
			record.Webdav,
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Samba,
			prometheus.GaugeValue,
			record.Samba,
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.SSH,
			prometheus.GaugeValue,
			record.SSH,
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.External,
			prometheus.GaugeValue,
			record.External,
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			c.ZFS,
			prometheus.GaugeValue,
			record.ZFS,
			labels...,
		)
	}
}
