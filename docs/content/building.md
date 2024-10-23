---
title: "Building"
date: 2022-07-20T00:00:00+00:00
anchor: "building"
weight: 20
---

As this project is built with Go you need to install Go first. The installation
of Go is out of the scope of this document, please follow the
[official documentation][golang]. After the installation of Go you need to get
the sources:

{{< highlight txt >}}
git clone https://github.com/promhippie/hetzner_exporter.git
cd hetzner_exporter/
{{< / highlight >}}

All required tool besides Go itself are bundled, all you need is part of the
`Makefile`:

{{< highlight txt >}}
make generate build
{{< / highlight >}}

Finally you should have the binary within the `bin/` folder now, give it a try
with `./bin/hetzner_exporter -h` to see all available options.

[golang]: https://golang.org/doc/install
