"""Install Cilium."""

import pulumi
import littlejo_cilium as cilium

config = pulumi.Config()

install = cilium.Install(
    "cilium",
    # Without a version, the provider installs its default version of Cilium.
    version=config.get("ciliumVersion"),
    sets=[
        "ipam.mode=kubernetes",
        "operator.replicas=1",
    ],
)

pulumi.export("version", install.version)
