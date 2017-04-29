# Hetzner Exporter

[![Build Status](http://github.dronehippie.de/api/badges/webhippie/hetzner_exporter/status.svg)](http://github.dronehippie.de/webhippie/hetzner_exporter)
[![Go Doc](https://godoc.org/github.com/webhippie/hetzner_exporter?status.svg)](http://godoc.org/github.com/webhippie/hetzner_exporter)
[![Go Report](http://goreportcard.com/badge/github.com/webhippie/hetzner_exporter)](http://goreportcard.com/report/github.com/webhippie/hetzner_exporter)
[![](https://images.microbadger.com/badges/image/tboerger/hetzner-exporter.svg)](http://microbadger.com/images/tboerger/hetzner-exporter "Get your own image badge on microbadger.com")
[![Join the chat at https://gitter.im/webhippie/general](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/webhippie/general)

A [Prometheus](https://prometheus.io/) exporter that collects Hetzner statistics.


## Installation

If you are missing something just write us on our nice [Gitter](https://gitter.im/webhippie/general) chat. If you find a security issue please contact thomas@webhippie.de first. Currently we are providing only a Docker image at `tboerger/hetzner-exporter`.


### Usage

```bash
# docker run -ti --rm tboerger/hetzner-exporter -h
Usage of /bin/hetzner_exporter:
  -hetzner.password string
      Password to authenticate on the API
  -hetzner.username string
      Username to authenticate on the API
  -log.format value
      Set the log target and format. Example: "logger:syslog?appname=bob&local=7" or "logger:stdout?json=true" (default "logger:stderr")
  -log.level value
      Only log messages with the given severity or above. Valid levels: [debug, info, warn, error, fatal] (default "info")
  -version
      Print version information
  -web.listen-address string
      Address to listen on for web interface and telemetry (default ":9107")
  -web.telemetry-path string
      Path to expose metrics of the exporter (default "/metrics")
```


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). It is also possible to just simply execute the `go get github.com/webhippie/hetzner_exporter` command, but we prefer to use our `Makefile`:

```bash
go get -d github.com/webhippie/hetzner_exporter
cd $GOPATH/src/github.com/webhippie/hetzner_exporter
make test build

./hetzner_exporter -h
```


## Metrics

```





```


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2017 Thomas Boerger <http://www.webhippie.de>
```
