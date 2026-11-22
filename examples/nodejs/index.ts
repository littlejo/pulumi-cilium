import * as pulumi from "@pulumi/pulumi";
import * as cilium from "@littlejo/cilium";

const config = new pulumi.Config();

const install = new cilium.Install("cilium", {
    // Without a version, the provider installs its default version of Cilium.
    version: config.get("ciliumVersion"),
    sets: [
        "ipam.mode=kubernetes",
        "operator.replicas=1",
    ],
});

export const version = install.version;
