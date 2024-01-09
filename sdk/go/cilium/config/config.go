// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/littlejo/pulumi-cilium/sdk/go/cilium/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// A path to a kube config file (Default: `~/.kube/config`).
func GetConfigPath(ctx *pulumi.Context) string {
	return config.Get(ctx, "cilium:configPath")
}

// Context of kubeconfig file (Default: `default context`).
func GetContext(ctx *pulumi.Context) string {
	return config.Get(ctx, "cilium:context")
}

// Namespace to install cilium (Default: `kube-system`).
func GetNamespace(ctx *pulumi.Context) string {
	return config.Get(ctx, "cilium:namespace")
}