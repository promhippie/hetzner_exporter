---
title: "Getting Started"
date: 2018-05-02T00:00:00+00:00
anchor: "getting-started"
weight: 10
---

## Installation

We won't cover further details how to properly setup [Prometheus](https://prometheus.io) itself, we will only cover some basic setup based on [docker-compose](https://docs.docker.com/compose/). But if you want to run this exporter without [docker-compose](https://docs.docker.com/compose/) you should be able to adopt that to your needs.

First of all we need to prepare a configuration for [Prometheus](https://prometheus.io) that includes the exporter as a target based on a static host mapping which is just the [docker-compose](https://docs.docker.com/compose/) container name, e.g. `hetzner-exporter`.

{{< gist tboerger 55dbda948f64b68eaee68858ccac32d1 "prometheus.yml" >}}

After preparing the configuration we need to create the `docker-compose.yml` within the same folder, this `docker-compose.yml` starts a simple [Prometheus](https://prometheus.io) instance together with the exporter. Don't forget to update the exporter envrionment variables with the required credentials.

{{< gist tboerger 55dbda948f64b68eaee68858ccac32d1 "docker-compose.yml" >}}

Since our `latest` Docker tag always refers to the `master` branch of the Git repository you should always use some fixed version. You can see all available tags at our [DockerHub repository](https://hub.docker.com/r/promhippie/hetzner-exporter/tags/), there you will see that we also provide a manifest, you can easily start the exporter on various architectures without any change to the image name. You should apply a change like this to the `docker-compose.yml`:

{{< gist tboerger 55dbda948f64b68eaee68858ccac32d1 "tag.diff" >}}

If you want to access the exporter directly you should bind it to a local port, otherwise only [Prometheus](https://prometheus.io) will have access to the exporter. For debugging purpose or just to discover all available metrics directly you can apply this change to your `docker-compose.yml`, after that you can access it directly at [http://localhost:9502/metrics](http://localhost:9502/metrics):

{{< gist tboerger 55dbda948f64b68eaee68858ccac32d1 "port.diff" >}}

Finally the exporter should be configured fine, let's start this stack with [docker-compose](https://docs.docker.com/compose/), you just need to execute `docker-compose up` within the directory where you have stored the `prometheus.yml` and `docker-compose.yml`.

{{< gist tboerger 55dbda948f64b68eaee68858ccac32d1 "output.log" >}}

That's all, the exporter should be up and running. Have fun with it and hopefully you will gather interesting metrics and never run into issues. You can access the exporter at [http://localhost:9502/metrics](http://localhost:9502/metrics) and [Prometheus](https://prometheus.io) at [http://localhost:9090](http://localhost:9090). Below you can see the [Prometheus](https://prometheus.io) targets and some simple query:

{{< figure src="targets.png" title="Prometheus targets including Hetzner exporter" >}}

{{< figure src="query.png" title="Prometheus query based on Hetzner exporter" >}}

## Kubernetes

Currently we have not prepared a deployment for Kubernetes, but this is something we will provide for sure. Most interesting will be the integration into the [Prometheus Operator](https://coreos.com/operators/prometheus/docs/latest/), so stay tuned.

## Configuration

HETZNER_EXPORTER_USERNAME
: Username for the Hetzner API, required for authentication

HETZNER_EXPORTER_PASSWORD
: Password for the Hetzner API, required for authentication

HETZNER_EXPORTER_LOG_LEVEL
: Only log messages with given severity, defaults to `info`

HETZNER_EXPORTER_LOG_PRETTY
: Enable pretty messages for logging, defaults to `false`

HETZNER_EXPORTER_WEB_ADDRESS
: Address to bind the metrics server, defaults to `0.0.0.0:9502`

HETZNER_EXPORTER_WEB_PATH
: Path to bind the metrics server, defaults to `/metrics`

HETZNER_EXPORTER_REQUEST_TIMEOUT
: Request timeout as duration, defaults to `5s`

HETZNER_EXPORTER_COLLECTOR_SERVERS
: Enable collector for servers, defaults to `true`

HETZNER_EXPORTER_COLLECTOR_SSH_KEYS
: Enable collector for SSH keys, defaults to `true`

## Metrics

hetzner_request_duration_seconds
: Histogram of latencies for requests to the Hetzner API per collector

hetzner_request_failures_total
: Total number of failed requests to the Hetzner API per collector

hetzner_server_running
: If 1 the server is running, 0 otherwise

hetzner_server_traffic_bytes
: Amount of included traffic for the server

hetzner_server_paid_timestamp
: Timestamp of the date until server is paid

hetzner_server_flatrate
: If 1 the server got a flatrate enabled, 0 otherwise

hetzner_server_throttled
: If 1 the server is in a throttled state, 0 otherwise

hetzner_server_cancelled
: If 1 the server have been cancelled, 0 otherwise

hetzner_ssh_key
: Information about SSH keys in your Hetzner robot
