---
title: Cilium Installation & Configuration
meta_desc: Information on how to install the Cilium provider.
layout: package
---

## Installation

The Pulumi Cilium provider is available as a package in these Pulumi languages:

* JavaScript/TypeScript: [`@littlejo/cilium`](https://www.npmjs.com/package/@littlejo/cilium)
* Python: [`littlejo-cilium`](https://pypi.org/project/littlejo-cilium/)
* Go: [`github.com/littlejo/pulumi-cilium/sdk/go/cilium`](https://pkg.go.dev/github.com/littlejo/pulumi-cilium/sdk/go/cilium)
* YAML: no package to install

### Provider Binary

The Cilium provider binary is a third party binary. Pulumi downloads it from the GitHub releases of the provider;
it can also be installed with the `pulumi plugin` command:

```bash
pulumi plugin install resource cilium <version> --server github://api.github.com/littlejo
```

Replace the version string `<version>` with your desired version.

## Configuration

The following configuration settings are available for the Cilium provider:

* `cilium:configPath`: The path of your kubeconfig (Default: `~/.kube/config`)
* `cilium:configContent`: The content of your kubeconfig, instead of `cilium:configPath`
* `cilium:context`: The context in your kubeconfig (Default: the current context)
* `cilium:namespace`: The namespace of Cilium (Default: `kube-system`)
* `cilium:helmRelease`: The name of the Helm release of Cilium (Default: `cilium`)

For example:

```bash
pulumi config set cilium:configPath kubeconfig
```
