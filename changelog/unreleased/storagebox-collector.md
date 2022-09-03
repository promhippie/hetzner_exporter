Change: Add collector for storageboxes

We implemented a new collector to gather metrics for the Hetzner storageboxes.
You should increase the scrape time as you could reach the API rate limit. This
collector will be disabled by default, so you got to enable it via flag or
environment variable.

https://github.com/promhippie/hetzner_exporter/issues/61
