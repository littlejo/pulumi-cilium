---
title: Cilium Installation & Configuration
meta_desc: Information on how to install the Cilium provider.
layout: installation
---

## Installation

The Pulumi Cilium provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/cilium`](https://www.npmjs.com/package/@pulumiverse/cilium)
* Python: [`pulumiverse_cilium`](https://pypi.org/project/pulumiverse_cilium/)
* Go: [`github.com/littlejo/pulumi-cilium/sdk/go/cilium`](https://pkg.go.dev/github.com/littlejo/pulumi-cilium/sdk/go/cilium)
* .NET: [`Pulumiverse.Cilium`](https://www.nuget.org/packages/Pulumiverse.Cilium)


## Configuration

> Note:  
> Replace the following **sample content**, with the configuration options
> of the wrapped Terraform provider and remove this note.

The following configuration points are available for the `cilium` provider:

- `cilium:apiKey` (environment: `cilium_API_KEY`) - the API key for `cilium`
- `cilium:region` (environment: `cilium_REGION`) - the region in which to deploy resources

### Provider Binary

The Cilium provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource cilium <version>
```

Replace the version string `<version>` with your desired version.
