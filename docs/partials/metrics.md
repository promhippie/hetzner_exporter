hetzner_request_duration_seconds{collector}
: Histogram of latencies for requests to the api per collector

hetzner_request_failures_total{collector}
: Total number of failed requests to the api per collector

hetzner_server_cancelled{<prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>}
: If 1 the server have been cancelled, 0 otherwise

hetzner_server_flatrate{<prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>}
: If 1 the server got a flatrate enabled, 0 otherwise

hetzner_server_paid_timestamp{<prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>}
: Timestamp of the date until server is paid

hetzner_server_running{<prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>}
: If 1 the server is running, 0 otherwise

hetzner_server_traffic_bytes{<prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>}
: Amount of included traffic for the server

hetzner_ssh_key{<prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>}
: Information about SSH keys in your Hetzner robot

hetzner_storagebox_cancelled{<prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>}
: If 1 the storagebox have been cancelled, 0 otherwise

hetzner_storagebox_data{<prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>}
: Used storage by files for the storagebox in MB

hetzner_storagebox_external{<prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>}
: If 1 the storagebox can be accessed from external, 0 otherwise

hetzner_storagebox_locked{<prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>}
: If 1 the storagebox have been locked, 0 otherwise

hetzner_storagebox_paid{<prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>}
: Timestamp of the date until storagebox is paid

hetzner_storagebox_quota{<prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>}
: Available storage for the storagebox in MB

hetzner_storagebox_samba{<prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>}
: If 1 the storagebox can be accessed via samba, 0 otherwise

hetzner_storagebox_snapshots{<prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>}
: Used storage by snapshots for the storagebox in MB

hetzner_storagebox_ssh{<prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>}
: If 1 the storagebox can be accessed via ssh, 0 otherwise

hetzner_storagebox_usage{<prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>}
: Used storage for the storagebox in MB

hetzner_storagebox_webdav{<prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>}
: If 1 the storagebox can be accessed via webdav, 0 otherwise

hetzner_storagebox_zfs{<prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>, <prometheus.ConstrainedLabel Value>}
: If 1 the zfs directory is visible, 0 otherwise
