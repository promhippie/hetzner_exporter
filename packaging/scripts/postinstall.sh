#!/bin/sh
set -e

chown -R hetzner-exporter:hetzner-exporter /var/lib/hetzner-exporter
chmod 750 /var/lib/hetzner-exporter

if [ -d /run/systemd/system ]; then
    systemctl daemon-reload

    if systemctl is-enabled --quiet hetzner-exporter.service; then
        systemctl restart hetzner-exporter.service
    fi
fi
