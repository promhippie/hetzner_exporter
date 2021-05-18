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

hetzner_server_throttled{id, name, datacenter}
: If 1 the server is in a throttled state, 0 otherwise

hetzner_server_traffic_bytes{id, name, datacenter}
: Amount of included traffic for the server

hetzner_ssh_key{name, type, size, fingerprint}
: Information about SSH keys in your Hetzner robot
