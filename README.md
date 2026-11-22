# Cilium Resource Provider

[![CI](https://github.com/littlejo/pulumi-cilium/actions/workflows/ci.yml/badge.svg)](https://github.com/littlejo/pulumi-cilium/actions/workflows/ci.yml)
[![PyPI](https://img.shields.io/pypi/v/littlejo-cilium)](https://pypi.org/project/littlejo-cilium/)
[![npm](https://img.shields.io/npm/v/@littlejo/cilium)](https://www.npmjs.com/package/@littlejo/cilium)

The Cilium resource provider for Pulumi installs and manages [Cilium](https://cilium.io) on Kubernetes clusters.
It does what the `cilium` CLI does: install and upgrade Cilium, configure it, enable Hubble, Cluster Mesh and
kube-proxy replacement.

It is built from the [`cilium` Terraform provider](https://github.com/littlejo/terraform-provider-cilium) and supports
Cilium 1.17 and newer.

## Installing

The provider is available in these languages:

| Language   | Package                                                                                              | Install                                                  |
|------------|------------------------------------------------------------------------------------------------------|----------------------------------------------------------|
| Python     | [`littlejo-cilium`](https://pypi.org/project/littlejo-cilium/)                                       | `pip install littlejo-cilium`                            |
| TypeScript | [`@littlejo/cilium`](https://www.npmjs.com/package/@littlejo/cilium)                                 | `npm install @littlejo/cilium`                           |
| Go         | [`github.com/littlejo/pulumi-cilium/sdk`](https://pkg.go.dev/github.com/littlejo/pulumi-cilium/sdk) | `go get github.com/littlejo/pulumi-cilium/sdk@latest`    |
| YAML       | -                                                                                                    | -                                                        |

Pulumi downloads the provider plugin from the
[GitHub releases](https://github.com/littlejo/pulumi-cilium/releases) of this repository. To install it by hand:

```bash
pulumi plugin install resource cilium <version> --server github://api.github.com/littlejo
```

## Example

```typescript
import * as cilium from "@littlejo/cilium";

const install = new cilium.Install("cilium", {
    version: "1.20.5",
    sets: [
        "ipam.mode=kubernetes",
        "operator.replicas=1",
    ],
});

new cilium.Hubble("hubble", { ui: true }, { dependsOn: install });
```

```python
import littlejo_cilium as cilium

install = cilium.Install(
    "cilium",
    version="1.20.5",
    sets=[
        "ipam.mode=kubernetes",
        "operator.replicas=1",
    ],
)
```

The [examples](examples) directory has a complete program for each language.

## Resources

| Resource                | `cilium` CLI equivalent                          |
|-------------------------|--------------------------------------------------|
| `Install`               | `cilium install`, `cilium upgrade`, `cilium uninstall` |
| `Config`                | `cilium config`                                  |
| `Hubble`                | `cilium hubble enable`, `cilium hubble disable`  |
| `Clustermesh`           | `cilium clustermesh enable`, `cilium clustermesh disable` |
| `ClustermeshConnection` | `cilium clustermesh connect`, `cilium clustermesh disconnect` |
| `KubeproxyFree`         | Disables the kube-proxy DaemonSet                |
| `getHelmValues`         | Reads the Helm values of the Cilium release      |

## Configuration

| Option                 | Description                                                        |
|------------------------|--------------------------------------------------------------------|
| `cilium:configPath`    | Path of the kubeconfig file (default: `~/.kube/config`)            |
| `cilium:configContent` | Content of the kubeconfig file, instead of `configPath`            |
| `cilium:context`       | Context of the kubeconfig file (default: the current context)      |
| `cilium:namespace`     | Namespace where Cilium is installed (default: `kube-system`)       |
| `cilium:helmRelease`   | Name of the Helm release of Cilium (default: `cilium`)             |

```bash
pulumi config set cilium:context kind-test
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) to build and test the provider.
