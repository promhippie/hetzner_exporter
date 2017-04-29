package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	"github.com/webhippie/hetzner_exporter/exporter"

	_ "net/http/pprof"
)

var (
	// showVersion is a flag to display the current version.
	showVersion = flag.Bool("version", false, "Print version information")

	// listenAddress defines the local address binding for the server.
	listenAddress = flag.String("web.listen-address", ":9107", "Address to listen on for web interface and telemetry")

	// metricsPath defines the path to access the metrics.
	metricsPath = flag.String("web.telemetry-path", "/metrics", "Path to expose metrics of the exporter")

	// username defines the username for the Hetzner API.
	username = flag.String("hetzner.username", "", "Username to authenticate on the API")

	// password defines the password for the Hetzner API.
	password = flag.String("hetzner.password", "", "Password to authenticate on the API")
)

// init registers the collector version.
func init() {
	prometheus.MustRegister(version.NewCollector("hetzner_exporter"))
}

// main simply initializes this tool.
func main() {
	flag.Parse()

	if *showVersion {
		fmt.Fprintln(os.Stdout, version.Print("hetzner_exporter"))
		os.Exit(0)
	}

	if *username == "" {
		fmt.Fprintln(os.Stderr, "Please provide a username for authentication")
		os.Exit(1)
	}

	if *password == "" {
		fmt.Fprintln(os.Stderr, "Please provide a password for authentication")
		os.Exit(1)
	}

	log.Infoln("Starting Hetzner exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	e := exporter.NewExporter(*username, *password)

	prometheus.MustRegister(e)
	prometheus.Unregister(prometheus.NewGoCollector())
	prometheus.Unregister(prometheus.NewProcessCollector(os.Getpid(), ""))

	http.Handle(*metricsPath, promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, *metricsPath, http.StatusMovedPermanently)
	})

	log.Infof("Listening on %s", *listenAddress)

	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		log.Fatal(err)
	}
}
