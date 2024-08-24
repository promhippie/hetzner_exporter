# Changelog for 1.2.0

The following sections list the changes for 1.2.0.

## Summary

 * Chg #142: Read secrets form files
 * Chg #142: Integrate standard web config
 * Enh #142: Integrate option pprof profiling

## Details

 * Change #142: Read secrets form files

   We have added proper support to load secrets like the password from files or
   from base64-encoded strings. Just provide the flags or environment variables for
   token or private key with a DSN formatted string like `file://path/to/file` or
   `base64://Zm9vYmFy`.

   https://github.com/promhippie/hetzner_exporter/pull/142

 * Change #142: Integrate standard web config

   We integrated the new web config from the Prometheus toolkit which provides a
   configuration for TLS support and also some basic builtin authentication. For
   the detailed configuration you can check out the documentation.

   https://github.com/promhippie/hetzner_exporter/pull/142

 * Enhancement #142: Integrate option pprof profiling

   We have added an option to enable a pprof endpoint for proper profiling support
   with the help of tools like Parca. The endpoint `/debug/pprof` can now
   optionally be enabled to get the profiling details for catching potential memory
   leaks.

   https://github.com/promhippie/hetzner_exporter/pull/142


# Changelog for 1.1.0

The following sections list the changes for 1.1.0.

## Summary

 * Chg #61: Replace archived client library
 * Chg #61: Add collector for storageboxes

## Details

 * Change #61: Replace archived client library

   We replaced the Hetzner client library by a custom internal package as the
   upstream library have been archived/deprecated.

   https://github.com/promhippie/hetzner_exporter/issues/61

 * Change #61: Add collector for storageboxes

   We implemented a new collector to gather metrics for the Hetzner storageboxes.
   You should increase the scrape time as you could reach the API rate limit. This
   collector will be disabled by default, so you got to enable it via flag or
   environment variable.

   https://github.com/promhippie/hetzner_exporter/issues/61


# Changelog for 1.0.0

The following sections list the changes for 1.0.0.

## Summary

 * Chg #16: Refactor build tools and project structure
 * Chg #18: Drop darwin/386 release builds

## Details

 * Change #16: Refactor build tools and project structure

   To have a unified project structure and build tooling we have integrated the
   same structure we already got within our GitHub exporter.

   https://github.com/promhippie/hetzner_exporter/issues/16

 * Change #18: Drop darwin/386 release builds

   We dropped the build of 386 builds on Darwin as this architecture is not
   supported by current Go versions anymore.

   https://github.com/promhippie/hetzner_exporter/issues/18


# Changelog for 0.1.0

The following sections list the changes for 0.1.0.

## Summary

 * Chg #15: Initial release of basic version

## Details

 * Change #15: Initial release of basic version

   Just prepared an initial basic version which could be released to the public.

   https://github.com/promhippie/hetzner_exporter/issues/15


