// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cilium

import (
	"context"
	"reflect"

	"github.com/littlejo/pulumi-cilium/sdk/go/cilium/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Hubble resource for Cilium. This is equivalent to cilium cli: `cilium hubble`: It manages cilium hubble
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
//			_, err := cilium.NewHubble(ctx, "example", &cilium.HubbleArgs{
//				Ui: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Hubble struct {
	pulumi.CustomResourceState

	// Deploy Hubble Relay (Default: `true`).
	Relay pulumi.BoolOutput `pulumi:"relay"`
	// Enable Hubble UI (Default: `false`).
	Ui pulumi.BoolOutput `pulumi:"ui"`
}

// NewHubble registers a new resource with the given unique name, arguments, and options.
func NewHubble(ctx *pulumi.Context,
	name string, args *HubbleArgs, opts ...pulumi.ResourceOption) (*Hubble, error) {
	if args == nil {
		args = &HubbleArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Hubble
	err := ctx.RegisterResource("cilium:index/hubble:Hubble", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHubble gets an existing Hubble resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHubble(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HubbleState, opts ...pulumi.ResourceOption) (*Hubble, error) {
	var resource Hubble
	err := ctx.ReadResource("cilium:index/hubble:Hubble", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Hubble resources.
type hubbleState struct {
	// Deploy Hubble Relay (Default: `true`).
	Relay *bool `pulumi:"relay"`
	// Enable Hubble UI (Default: `false`).
	Ui *bool `pulumi:"ui"`
}

type HubbleState struct {
	// Deploy Hubble Relay (Default: `true`).
	Relay pulumi.BoolPtrInput
	// Enable Hubble UI (Default: `false`).
	Ui pulumi.BoolPtrInput
}

func (HubbleState) ElementType() reflect.Type {
	return reflect.TypeOf((*hubbleState)(nil)).Elem()
}

type hubbleArgs struct {
	// Deploy Hubble Relay (Default: `true`).
	Relay *bool `pulumi:"relay"`
	// Enable Hubble UI (Default: `false`).
	Ui *bool `pulumi:"ui"`
}

// The set of arguments for constructing a Hubble resource.
type HubbleArgs struct {
	// Deploy Hubble Relay (Default: `true`).
	Relay pulumi.BoolPtrInput
	// Enable Hubble UI (Default: `false`).
	Ui pulumi.BoolPtrInput
}

func (HubbleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hubbleArgs)(nil)).Elem()
}

type HubbleInput interface {
	pulumi.Input

	ToHubbleOutput() HubbleOutput
	ToHubbleOutputWithContext(ctx context.Context) HubbleOutput
}

func (*Hubble) ElementType() reflect.Type {
	return reflect.TypeOf((**Hubble)(nil)).Elem()
}

func (i *Hubble) ToHubbleOutput() HubbleOutput {
	return i.ToHubbleOutputWithContext(context.Background())
}

func (i *Hubble) ToHubbleOutputWithContext(ctx context.Context) HubbleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HubbleOutput)
}

// HubbleArrayInput is an input type that accepts HubbleArray and HubbleArrayOutput values.
// You can construct a concrete instance of `HubbleArrayInput` via:
//
//	HubbleArray{ HubbleArgs{...} }
type HubbleArrayInput interface {
	pulumi.Input

	ToHubbleArrayOutput() HubbleArrayOutput
	ToHubbleArrayOutputWithContext(context.Context) HubbleArrayOutput
}

type HubbleArray []HubbleInput

func (HubbleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Hubble)(nil)).Elem()
}

func (i HubbleArray) ToHubbleArrayOutput() HubbleArrayOutput {
	return i.ToHubbleArrayOutputWithContext(context.Background())
}

func (i HubbleArray) ToHubbleArrayOutputWithContext(ctx context.Context) HubbleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HubbleArrayOutput)
}

// HubbleMapInput is an input type that accepts HubbleMap and HubbleMapOutput values.
// You can construct a concrete instance of `HubbleMapInput` via:
//
//	HubbleMap{ "key": HubbleArgs{...} }
type HubbleMapInput interface {
	pulumi.Input

	ToHubbleMapOutput() HubbleMapOutput
	ToHubbleMapOutputWithContext(context.Context) HubbleMapOutput
}

type HubbleMap map[string]HubbleInput

func (HubbleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Hubble)(nil)).Elem()
}

func (i HubbleMap) ToHubbleMapOutput() HubbleMapOutput {
	return i.ToHubbleMapOutputWithContext(context.Background())
}

func (i HubbleMap) ToHubbleMapOutputWithContext(ctx context.Context) HubbleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HubbleMapOutput)
}

type HubbleOutput struct{ *pulumi.OutputState }

func (HubbleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Hubble)(nil)).Elem()
}

func (o HubbleOutput) ToHubbleOutput() HubbleOutput {
	return o
}

func (o HubbleOutput) ToHubbleOutputWithContext(ctx context.Context) HubbleOutput {
	return o
}

// Deploy Hubble Relay (Default: `true`).
func (o HubbleOutput) Relay() pulumi.BoolOutput {
	return o.ApplyT(func(v *Hubble) pulumi.BoolOutput { return v.Relay }).(pulumi.BoolOutput)
}

// Enable Hubble UI (Default: `false`).
func (o HubbleOutput) Ui() pulumi.BoolOutput {
	return o.ApplyT(func(v *Hubble) pulumi.BoolOutput { return v.Ui }).(pulumi.BoolOutput)
}

type HubbleArrayOutput struct{ *pulumi.OutputState }

func (HubbleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Hubble)(nil)).Elem()
}

func (o HubbleArrayOutput) ToHubbleArrayOutput() HubbleArrayOutput {
	return o
}

func (o HubbleArrayOutput) ToHubbleArrayOutputWithContext(ctx context.Context) HubbleArrayOutput {
	return o
}

func (o HubbleArrayOutput) Index(i pulumi.IntInput) HubbleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Hubble {
		return vs[0].([]*Hubble)[vs[1].(int)]
	}).(HubbleOutput)
}

type HubbleMapOutput struct{ *pulumi.OutputState }

func (HubbleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Hubble)(nil)).Elem()
}

func (o HubbleMapOutput) ToHubbleMapOutput() HubbleMapOutput {
	return o
}

func (o HubbleMapOutput) ToHubbleMapOutputWithContext(ctx context.Context) HubbleMapOutput {
	return o
}

func (o HubbleMapOutput) MapIndex(k pulumi.StringInput) HubbleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Hubble {
		return vs[0].(map[string]*Hubble)[vs[1].(string)]
	}).(HubbleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HubbleInput)(nil)).Elem(), &Hubble{})
	pulumi.RegisterInputType(reflect.TypeOf((*HubbleArrayInput)(nil)).Elem(), HubbleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HubbleMapInput)(nil)).Elem(), HubbleMap{})
	pulumi.RegisterOutputType(HubbleOutput{})
	pulumi.RegisterOutputType(HubbleArrayOutput{})
	pulumi.RegisterOutputType(HubbleMapOutput{})
}
