# Hetzner Exporter

[![Build Status](http://github.dronehippie.de/api/badges/promhippie/hetzner_exporter/status.svg)](http://github.dronehippie.de/promhippie/hetzner_exporter)
[![Stories in Ready](https://badge.waffle.io/promhippie/hetzner_exporter.svg?label=ready&title=Ready)](http://waffle.io/promhippie/hetzner_exporter)
[![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/f26fafdffe134732b196de6c5e2f16b8)](https://www.codacy.com/app/promhippie/hetzner_exporter?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=promhippie/hetzner_exporter&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/promhippie/hetzner_exporter?status.svg)](http://godoc.org/github.com/promhippie/hetzner_exporter)
[![Go Report](http://goreportcard.com/badge/github.com/promhippie/hetzner_exporter)](http://goreportcard.com/report/github.com/promhippie/hetzner_exporter)
[![](https://images.microbadger.com/badges/image/promhippie/hetzner-exporter.svg)](http://microbadger.com/images/promhippie/hetzner-exporter "Get your own image badge on microbadger.com")

An exporter for [Prometheus](https://prometheus.io/) that collects metrics from [Hetzner](http://robot.your-server.de).

## Install

You can download prebuilt binaries from our [GitHub releases](https://github.com/promhippie/hetzner_exporter/releases), or you can use our Docker images published on [Docker Hub](https://hub.docker.com/r/promhippie/hetzner-exporter/tags/). If you need further guidance how to install this take a look at our [documentation](https://promhippie.github.io/hetzner_exporter/#getting-started).

## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.8.

```bash
go get -d github.com/promhippie/hetzner_exporter
cd $GOPATH/src/github.com/promhippie/hetzner_exporter

# install retool
make retool

# sync dependencies
make sync

# generate code
make generate

# build binary
make build

./bin/hetzner_exporter -h
```

## Security

If you find a security issue please contact thomas@webhippie.de first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

* [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```
