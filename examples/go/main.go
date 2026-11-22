package main

import (
	"github.com/littlejo/pulumi-cilium/sdk/go/cilium"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		args := &cilium.InstallArgs{
			Sets: pulumi.StringArray{
				pulumi.String("ipam.mode=kubernetes"),
				pulumi.String("operator.replicas=1"),
			},
		}
		// Without a version, the provider installs its default version of Cilium.
		if version := config.Get(ctx, "ciliumVersion"); version != "" {
			args.Version = pulumi.String(version)
		}

		install, err := cilium.NewInstall(ctx, "cilium", args)
		if err != nil {
			return err
		}

		ctx.Export("version", install.Version)
		return nil
	})
}
