---
title: Cilium Installation & Configuration
meta_desc: Information on how to install the Cilium provider.
layout: installation
---

## Installation

The Pulumi Cilium provider is available as a package in all Pulumi languages:

* Go: [`github.com/littlejo/pulumi-cilium/sdk/go/cilium`](https://pkg.go.dev/github.com/littlejo/pulumi-cilium/sdk/go/cilium)

## Configuration

The following configuration points are available for the `cilium` provider:
* `cilium:configPath`
* `cilium:context`
* `cilium:namespace`

### Provider Binary

The Cilium provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource cilium <version>
```

Replace the version string `<version>` with your desired version.
