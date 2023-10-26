---
title: "Usage"
date: 2022-07-22T00:00:00+00:00
anchor: "getting-started"
weight: 10
---

## Installation

We won't cover further details how to properly setup [Prometheus][prometheus]
itself, we will only cover some basic setup based on [docker-compose][compose].
But if you want to run this exporter without [docker-compose][compose] you
should be able to adopt that to your needs.

First of all we need to prepare a configuration for [Prometheus][prometheus]
that includes the exporter based on a static configuration with the container
name as a hostname:

{{< highlight yaml >}}
global:
  scrape_interval: 1m
  scrape_timeout: 10s
  evaluation_interval: 1m

scrape_configs:
- job_name: hetzner
  static_configs:
  - targets:
    - hetzner_exporter:9502
{{< / highlight >}}

After preparing the configuration we need to create the `docker-compose.yml`
within the same folder, this `docker-compose.yml` starts a simple
[Prometheus][prometheus] instance together with the exporter. Don't forget to
update the envrionment variables with the required credentials.

{{< highlight yaml >}}
version: '2'

volumes:
  prometheus:

services:
  prometheus:
    image: prom/prometheus:latest
    restart: always
    ports:
      - 9090:9090
    volumes:
      - prometheus:/prometheus
      - ./prometheus.yml:/etc/prometheus/prometheus.yml

  hetzner_exporter:
    image: promhippie/hetzner-exporter:latest
    restart: always
    environment:
      - HETZNER_EXPORTER_USERNAME=#ws+qOeMD4UP
      - HETZNER_EXPORTER_PASSWORD=CNFPCgivAAqWu613
      - HETZNER_EXPORTER_LOG_PRETTY=true
{{< / highlight >}}

Since our `latest` tag always refers to the `master` branch of the Git
repository you should always use some fixed version. You can see all available
tags at [DockerHub][dockerhub] or [Quay][quayio], there you will see that we
also provide a manifest, you can easily start the exporter on various
architectures without any change to the image name. You should apply a change
like this to the `docker-compose.yml` file:

{{< highlight diff >}}
  hetzner_exporter:
-   image: promhippie/hetzner-exporter:latest
+   image: promhippie/hetzner-exporter:1.0.0
    restart: always
    environment:
      - HETZNER_EXPORTER_USERNAME=#ws+qOeMD4UP
      - HETZNER_EXPORTER_PASSWORD=CNFPCgivAAqWu613
      - HETZNER_EXPORTER_LOG_PRETTY=true
{{< / highlight >}}

If you want to access the exporter directly you should bind it to a local port,
otherwise only [Prometheus][prometheus] will have access to the exporter. For
debugging purpose or just to discover all available metrics directly you can
apply this change to your `docker-compose.yml`, after that you can access it
directly at [http://localhost:9502/metrics](http://localhost:9502/metrics):

{{< highlight diff >}}
  hetzner-exporter:
    image: promhippie/hetzner-exporter:latest
    restart: always
+   ports:
+     - 127.0.0.1:9502:9502
    environment:
      - HETZNER_EXPORTER_USERNAME=#ws+qOeMD4UP
      - HETZNER_EXPORTER_PASSWORD=CNFPCgivAAqWu613
      - HETZNER_EXPORTER_LOG_PRETTY=true
{{< / highlight >}}

If you want to secure the access to the exporter you can provide a web config.
You just need to provide a path to the config file in order to enable the
support for it, for details about the config format look at the
[documentation](#web-configuration) section:

{{< highlight diff >}}
  hetzner_exporter:
    image: promhippie/hetzner-exporter:latest
    restart: always
    environment:
+     - HETZNER_EXPORTER_WEB_CONFIG=path/to/web-config.json
      - HETZNER_EXPORTER_USERNAME=#ws+qOeMD4UP
      - HETZNER_EXPORTER_PASSWORD=CNFPCgivAAqWu613
      - HETZNER_EXPORTER_LOG_PRETTY=true
{{< / highlight >}}

If you want to provide the required secrets from a file it's also possible. For
this use case you can write the secret to a file on any path and reference it
with the following format:

{{< highlight diff >}}
  hetzner_exporter:
    image: promhippie/hetzner-exporter:latest
    restart: always
    environment:
-     - HETZNER_EXPORTER_USERNAME=#ws+qOeMD4UP
-     - HETZNER_EXPORTER_PASSWORD=CNFPCgivAAqWu613
+     - HETZNER_EXPORTER_USERNAME=file://path/to/secret/file/with/username
+     - HETZNER_EXPORTER_PASSWORD=file://path/to/secret/file/with/password
      - HETZNER_EXPORTER_LOG_PRETTY=true
{{< / highlight >}}

Finally the exporter should be configured fine, let's start this stack with
[docker-compose][compose], you just need to execute `docker-compose up` within
the directory where you have stored the `prometheus.yml` and
`docker-compose.yml`.

That's all, the exporter should be up and running. Have fun with it and
hopefully you will gather interesting metrics and never run into issues. You can
access the exporter at
[http://localhost:9502/metrics](http://localhost:9502/metrics) and
[Prometheus][prometheus] at [http://localhost:9090](http://localhost:9090).

## Configuration

{{< partial "envvars.md" >}}

### Web Configuration

If you want to secure the service by TLS or by some basic authentication you can
provide a `YAML` configuration file whch follows the [Prometheus][prometheus]
toolkit format. You can see a full configration example within the
[toolkit documentation][toolkit].

## Metrics

You can a rough list of available metrics below, additionally to these metrics
you will always get the standard metrics exported by the Golang client of
[Prometheus][prometheus]. If you want to know more about these standard metrics
take a look at the [process collector][proccollector] and the
[Go collector][gocollector].

{{< partial "metrics.md" >}}

[prometheus]: https://prometheus.io
[compose]: https://docs.docker.com/compose/
[dockerhub]: https://hub.docker.com/r/promhippie/hetzner-exporter/tags/
[quayio]: https://quay.io/repository/promhippie/hetzner-exporter?tab=tags
[toolkit]: https://github.com/prometheus/exporter-toolkit/blob/master/docs/web-configuration.md
[proccollector]: https://github.com/prometheus/client_golang/blob/master/prometheus/process_collector.go
[gocollector]: https://github.com/prometheus/client_golang/blob/master/prometheus/go_collector.go
