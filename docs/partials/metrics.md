hetzner_request_duration_seconds{collector}
: Histogram of latencies for requests to the api per collector

hetzner_request_failures_total{collector}
: Total number of failed requests to the api per collector

hetzner_server_cancelled{id, name, datacenter}
: If 1 the server have been cancelled, 0 otherwise

hetzner_server_flatrate{id, name, datacenter}
: If 1 the server got a flatrate enabled, 0 otherwise

hetzner_server_paid_timestamp{id, name, datacenter}
: Timestamp of the date until server is paid

hetzner_server_running{id, name, datacenter}
: If 1 the server is running, 0 otherwise

hetzner_server_traffic_bytes{id, name, datacenter}
: Amount of included traffic for the server

hetzner_ssh_key{name, type, size, fingerprint}
: Information about SSH keys in your Hetzner robot

hetzner_storagebox_cancelled{id, name, location, login}
: If 1 the storagebox have been cancelled, 0 otherwise

hetzner_storagebox_data{id, name, location, login}
: Used storage by files for the storagebox in MB

hetzner_storagebox_external{id, name, location, login}
: If 1 the storagebox can be accessed from external, 0 otherwise

hetzner_storagebox_locked{id, name, location, login}
: If 1 the storagebox have been locked, 0 otherwise

hetzner_storagebox_paid{id, name, location, login}
: Timestamp of the date until storagebox is paid

hetzner_storagebox_quota{id, name, location, login}
: Available storage for the storagebox in MB

hetzner_storagebox_samba{id, name, location, login}
: If 1 the storagebox can be accessed via samba, 0 otherwise

hetzner_storagebox_snapshots{id, name, location, login}
: Used storage by snapshots for the storagebox in MB

hetzner_storagebox_ssh{id, name, location, login}
: If 1 the storagebox can be accessed via ssh, 0 otherwise

hetzner_storagebox_usage{id, name, location, login}
: Used storage for the storagebox in MB

hetzner_storagebox_webdav{id, name, location, login}
: If 1 the storagebox can be accessed via webdav, 0 otherwise

hetzner_storagebox_zfs{id, name, location, login}
: If 1 the zfs directory is visible, 0 otherwise
