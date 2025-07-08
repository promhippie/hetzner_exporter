#!/bin/sh
set -e

if ! getent group hetzner-exporter >/dev/null 2>&1; then
    groupadd --system hetzner-exporter
fi

if ! getent passwd hetzner-exporter >/dev/null 2>&1; then
    useradd --system --create-home --home-dir /var/lib/hetzner-exporter --shell /bin/bash -g hetzner-exporter hetzner-exporter
fi
