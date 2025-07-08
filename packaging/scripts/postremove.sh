#!/bin/sh
set -e

if [ ! -d /var/lib/hetzner-exporter ]; then
    userdel hetzner-exporter 2>/dev/null || true
    groupdel hetzner-exporter 2>/dev/null || true
fi
