FROM quay.io/prometheus/busybox:latest
MAINTAINER Thomas Boerger <thomas@webhippie.de>

COPY hetzner_exporter /bin/hetzner_exporter

EXPOSE 9107
ENTRYPOINT ["/bin/hetzner_exporter"]
