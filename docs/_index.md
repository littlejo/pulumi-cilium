---
title: Cilium
meta_desc: Provides an overview of the Cilium Provider for Pulumi.
layout: package
---

The Cilium provider for Pulumi installs and manages [Cilium](https://cilium.io) on Kubernetes clusters: installation
and upgrades, configuration, Hubble, Cluster Mesh and kube-proxy replacement. It supports Cilium 1.17 and newer.

The provider uses a kubeconfig file to connect to the cluster: see [Installation & Configuration](./installation-configuration).

## Example

{{< chooser language "typescript,python,go,yaml" />}}

{{% choosable language typescript %}}

```typescript
import * as cilium from "@littlejo/cilium";

const install = new cilium.Install("cilium", {
    version: "1.20.5",
    sets: [
        "ipam.mode=kubernetes",
        "operator.replicas=1",
    ],
});
```

{{% /choosable %}}

{{% choosable language python %}}

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

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/littlejo/pulumi-cilium/sdk/go/cilium"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cilium.NewInstall(ctx, "cilium", &cilium.InstallArgs{
			Version: pulumi.String("1.20.5"),
			Sets: pulumi.StringArray{
				pulumi.String("ipam.mode=kubernetes"),
				pulumi.String("operator.replicas=1"),
			},
		})
		return err
	})
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
resources:
  cilium:
    type: cilium:Install
    properties:
      version: "1.20.5"
      sets:
        - ipam.mode=kubernetes
        - operator.replicas=1
```

{{% /choosable %}}
