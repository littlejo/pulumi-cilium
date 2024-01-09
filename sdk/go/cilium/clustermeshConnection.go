// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cilium

import (
	"context"
	"reflect"

	"github.com/littlejo/pulumi-cilium/sdk/go/cilium/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Cluster Mesh connection resource. This is equivalent to cilium cli: `cilium clustermesh connect` and `cilium clustermesh disconnect`: It manages the connections between two Kubernetes clusters.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/littlejo/pulumi-cilium/sdk/go/cilium"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cilium.NewClustermeshConnection(ctx, "example", &cilium.ClustermeshConnectionArgs{
//				DestinationContext: pulumi.String("context-2"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ClustermeshConnection struct {
	pulumi.CustomResourceState

	// Kubernetes configuration context of destination cluster
	DestinationContext pulumi.StringOutput `pulumi:"destinationContext"`
	// Namespace in which to install (Default: `kube-system`).
	Namespace pulumi.StringOutput `pulumi:"namespace"`
}

// NewClustermeshConnection registers a new resource with the given unique name, arguments, and options.
func NewClustermeshConnection(ctx *pulumi.Context,
	name string, args *ClustermeshConnectionArgs, opts ...pulumi.ResourceOption) (*ClustermeshConnection, error) {
	if args == nil {
		args = &ClustermeshConnectionArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClustermeshConnection
	err := ctx.RegisterResource("cilium:index/clustermeshConnection:ClustermeshConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClustermeshConnection gets an existing ClustermeshConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClustermeshConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClustermeshConnectionState, opts ...pulumi.ResourceOption) (*ClustermeshConnection, error) {
	var resource ClustermeshConnection
	err := ctx.ReadResource("cilium:index/clustermeshConnection:ClustermeshConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClustermeshConnection resources.
type clustermeshConnectionState struct {
	// Kubernetes configuration context of destination cluster
	DestinationContext *string `pulumi:"destinationContext"`
	// Namespace in which to install (Default: `kube-system`).
	Namespace *string `pulumi:"namespace"`
}

type ClustermeshConnectionState struct {
	// Kubernetes configuration context of destination cluster
	DestinationContext pulumi.StringPtrInput
	// Namespace in which to install (Default: `kube-system`).
	Namespace pulumi.StringPtrInput
}

func (ClustermeshConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*clustermeshConnectionState)(nil)).Elem()
}

type clustermeshConnectionArgs struct {
	// Kubernetes configuration context of destination cluster
	DestinationContext *string `pulumi:"destinationContext"`
	// Namespace in which to install (Default: `kube-system`).
	Namespace *string `pulumi:"namespace"`
}

// The set of arguments for constructing a ClustermeshConnection resource.
type ClustermeshConnectionArgs struct {
	// Kubernetes configuration context of destination cluster
	DestinationContext pulumi.StringPtrInput
	// Namespace in which to install (Default: `kube-system`).
	Namespace pulumi.StringPtrInput
}

func (ClustermeshConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clustermeshConnectionArgs)(nil)).Elem()
}

type ClustermeshConnectionInput interface {
	pulumi.Input

	ToClustermeshConnectionOutput() ClustermeshConnectionOutput
	ToClustermeshConnectionOutputWithContext(ctx context.Context) ClustermeshConnectionOutput
}

func (*ClustermeshConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**ClustermeshConnection)(nil)).Elem()
}

func (i *ClustermeshConnection) ToClustermeshConnectionOutput() ClustermeshConnectionOutput {
	return i.ToClustermeshConnectionOutputWithContext(context.Background())
}

func (i *ClustermeshConnection) ToClustermeshConnectionOutputWithContext(ctx context.Context) ClustermeshConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClustermeshConnectionOutput)
}

func (i *ClustermeshConnection) ToOutput(ctx context.Context) pulumix.Output[*ClustermeshConnection] {
	return pulumix.Output[*ClustermeshConnection]{
		OutputState: i.ToClustermeshConnectionOutputWithContext(ctx).OutputState,
	}
}

// ClustermeshConnectionArrayInput is an input type that accepts ClustermeshConnectionArray and ClustermeshConnectionArrayOutput values.
// You can construct a concrete instance of `ClustermeshConnectionArrayInput` via:
//
//	ClustermeshConnectionArray{ ClustermeshConnectionArgs{...} }
type ClustermeshConnectionArrayInput interface {
	pulumi.Input

	ToClustermeshConnectionArrayOutput() ClustermeshConnectionArrayOutput
	ToClustermeshConnectionArrayOutputWithContext(context.Context) ClustermeshConnectionArrayOutput
}

type ClustermeshConnectionArray []ClustermeshConnectionInput

func (ClustermeshConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClustermeshConnection)(nil)).Elem()
}

func (i ClustermeshConnectionArray) ToClustermeshConnectionArrayOutput() ClustermeshConnectionArrayOutput {
	return i.ToClustermeshConnectionArrayOutputWithContext(context.Background())
}

func (i ClustermeshConnectionArray) ToClustermeshConnectionArrayOutputWithContext(ctx context.Context) ClustermeshConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClustermeshConnectionArrayOutput)
}

func (i ClustermeshConnectionArray) ToOutput(ctx context.Context) pulumix.Output[[]*ClustermeshConnection] {
	return pulumix.Output[[]*ClustermeshConnection]{
		OutputState: i.ToClustermeshConnectionArrayOutputWithContext(ctx).OutputState,
	}
}

// ClustermeshConnectionMapInput is an input type that accepts ClustermeshConnectionMap and ClustermeshConnectionMapOutput values.
// You can construct a concrete instance of `ClustermeshConnectionMapInput` via:
//
//	ClustermeshConnectionMap{ "key": ClustermeshConnectionArgs{...} }
type ClustermeshConnectionMapInput interface {
	pulumi.Input

	ToClustermeshConnectionMapOutput() ClustermeshConnectionMapOutput
	ToClustermeshConnectionMapOutputWithContext(context.Context) ClustermeshConnectionMapOutput
}

type ClustermeshConnectionMap map[string]ClustermeshConnectionInput

func (ClustermeshConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClustermeshConnection)(nil)).Elem()
}

func (i ClustermeshConnectionMap) ToClustermeshConnectionMapOutput() ClustermeshConnectionMapOutput {
	return i.ToClustermeshConnectionMapOutputWithContext(context.Background())
}

func (i ClustermeshConnectionMap) ToClustermeshConnectionMapOutputWithContext(ctx context.Context) ClustermeshConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClustermeshConnectionMapOutput)
}

func (i ClustermeshConnectionMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ClustermeshConnection] {
	return pulumix.Output[map[string]*ClustermeshConnection]{
		OutputState: i.ToClustermeshConnectionMapOutputWithContext(ctx).OutputState,
	}
}

type ClustermeshConnectionOutput struct{ *pulumi.OutputState }

func (ClustermeshConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClustermeshConnection)(nil)).Elem()
}

func (o ClustermeshConnectionOutput) ToClustermeshConnectionOutput() ClustermeshConnectionOutput {
	return o
}

func (o ClustermeshConnectionOutput) ToClustermeshConnectionOutputWithContext(ctx context.Context) ClustermeshConnectionOutput {
	return o
}

func (o ClustermeshConnectionOutput) ToOutput(ctx context.Context) pulumix.Output[*ClustermeshConnection] {
	return pulumix.Output[*ClustermeshConnection]{
		OutputState: o.OutputState,
	}
}

// Kubernetes configuration context of destination cluster
func (o ClustermeshConnectionOutput) DestinationContext() pulumi.StringOutput {
	return o.ApplyT(func(v *ClustermeshConnection) pulumi.StringOutput { return v.DestinationContext }).(pulumi.StringOutput)
}

// Namespace in which to install (Default: `kube-system`).
func (o ClustermeshConnectionOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v *ClustermeshConnection) pulumi.StringOutput { return v.Namespace }).(pulumi.StringOutput)
}

type ClustermeshConnectionArrayOutput struct{ *pulumi.OutputState }

func (ClustermeshConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClustermeshConnection)(nil)).Elem()
}

func (o ClustermeshConnectionArrayOutput) ToClustermeshConnectionArrayOutput() ClustermeshConnectionArrayOutput {
	return o
}

func (o ClustermeshConnectionArrayOutput) ToClustermeshConnectionArrayOutputWithContext(ctx context.Context) ClustermeshConnectionArrayOutput {
	return o
}

func (o ClustermeshConnectionArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ClustermeshConnection] {
	return pulumix.Output[[]*ClustermeshConnection]{
		OutputState: o.OutputState,
	}
}

func (o ClustermeshConnectionArrayOutput) Index(i pulumi.IntInput) ClustermeshConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClustermeshConnection {
		return vs[0].([]*ClustermeshConnection)[vs[1].(int)]
	}).(ClustermeshConnectionOutput)
}

type ClustermeshConnectionMapOutput struct{ *pulumi.OutputState }

func (ClustermeshConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClustermeshConnection)(nil)).Elem()
}

func (o ClustermeshConnectionMapOutput) ToClustermeshConnectionMapOutput() ClustermeshConnectionMapOutput {
	return o
}

func (o ClustermeshConnectionMapOutput) ToClustermeshConnectionMapOutputWithContext(ctx context.Context) ClustermeshConnectionMapOutput {
	return o
}

func (o ClustermeshConnectionMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ClustermeshConnection] {
	return pulumix.Output[map[string]*ClustermeshConnection]{
		OutputState: o.OutputState,
	}
}

func (o ClustermeshConnectionMapOutput) MapIndex(k pulumi.StringInput) ClustermeshConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClustermeshConnection {
		return vs[0].(map[string]*ClustermeshConnection)[vs[1].(string)]
	}).(ClustermeshConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClustermeshConnectionInput)(nil)).Elem(), &ClustermeshConnection{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClustermeshConnectionArrayInput)(nil)).Elem(), ClustermeshConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClustermeshConnectionMapInput)(nil)).Elem(), ClustermeshConnectionMap{})
	pulumi.RegisterOutputType(ClustermeshConnectionOutput{})
	pulumi.RegisterOutputType(ClustermeshConnectionArrayOutput{})
	pulumi.RegisterOutputType(ClustermeshConnectionMapOutput{})
}