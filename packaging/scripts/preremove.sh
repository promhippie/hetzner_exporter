#!/bin/sh
set -e

systemctl stop hetzner-exporter.service || true
systemctl disable hetzner-exporter.service || true
