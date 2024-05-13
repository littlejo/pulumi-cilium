// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cilium

import (
	"context"
	"reflect"

	"github.com/littlejo/pulumi-cilium/sdk/go/cilium/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Deploy struct {
	pulumi.CustomResourceState

	// Datapath mode to use { tunnel | native | aws-eni | gke | azure | aks-byocni } (Default: `autodetected`).
	DataPath pulumi.StringOutput `pulumi:"dataPath"`
	// Helm values (`helm get values -n kube-system cilium`)
	HelmValues pulumi.StringOutput `pulumi:"helmValues"`
	// Helm chart repository to download Cilium charts from (Default: `https://helm.cilium.io`).
	Repository pulumi.StringOutput `pulumi:"repository"`
	// When upgrading, reset the helm values to the ones built into the chart (Default: `false`).
	Reset pulumi.BoolOutput `pulumi:"reset"`
	// When upgrading, reuse the helm values from the latest release unless any overrides from are set from other flags. This
	// option takes precedence over HelmResetValues (Default: `true`).
	Reuse pulumi.BoolOutput `pulumi:"reuse"`
	// Set helm values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2 (Default:
	// `[]`).
	Sets pulumi.StringArrayOutput `pulumi:"sets"`
	// values in raw yaml to pass to helm. (Default: `empty`).
	Values pulumi.StringOutput `pulumi:"values"`
	// Version of Cilium (Default: `v1.15.4`).
	Version pulumi.StringOutput `pulumi:"version"`
	// Wait for Cilium status is ok (Default: `true`).
	Wait pulumi.BoolOutput `pulumi:"wait"`
}

// NewDeploy registers a new resource with the given unique name, arguments, and options.
func NewDeploy(ctx *pulumi.Context,
	name string, args *DeployArgs, opts ...pulumi.ResourceOption) (*Deploy, error) {
	if args == nil {
		args = &DeployArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Deploy
	err := ctx.RegisterResource("cilium:index/deploy:Deploy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeploy gets an existing Deploy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeploy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeployState, opts ...pulumi.ResourceOption) (*Deploy, error) {
	var resource Deploy
	err := ctx.ReadResource("cilium:index/deploy:Deploy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Deploy resources.
type deployState struct {
	// Datapath mode to use { tunnel | native | aws-eni | gke | azure | aks-byocni } (Default: `autodetected`).
	DataPath *string `pulumi:"dataPath"`
	// Helm values (`helm get values -n kube-system cilium`)
	HelmValues *string `pulumi:"helmValues"`
	// Helm chart repository to download Cilium charts from (Default: `https://helm.cilium.io`).
	Repository *string `pulumi:"repository"`
	// When upgrading, reset the helm values to the ones built into the chart (Default: `false`).
	Reset *bool `pulumi:"reset"`
	// When upgrading, reuse the helm values from the latest release unless any overrides from are set from other flags. This
	// option takes precedence over HelmResetValues (Default: `true`).
	Reuse *bool `pulumi:"reuse"`
	// Set helm values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2 (Default:
	// `[]`).
	Sets []string `pulumi:"sets"`
	// values in raw yaml to pass to helm. (Default: `empty`).
	Values *string `pulumi:"values"`
	// Version of Cilium (Default: `v1.15.4`).
	Version *string `pulumi:"version"`
	// Wait for Cilium status is ok (Default: `true`).
	Wait *bool `pulumi:"wait"`
}

type DeployState struct {
	// Datapath mode to use { tunnel | native | aws-eni | gke | azure | aks-byocni } (Default: `autodetected`).
	DataPath pulumi.StringPtrInput
	// Helm values (`helm get values -n kube-system cilium`)
	HelmValues pulumi.StringPtrInput
	// Helm chart repository to download Cilium charts from (Default: `https://helm.cilium.io`).
	Repository pulumi.StringPtrInput
	// When upgrading, reset the helm values to the ones built into the chart (Default: `false`).
	Reset pulumi.BoolPtrInput
	// When upgrading, reuse the helm values from the latest release unless any overrides from are set from other flags. This
	// option takes precedence over HelmResetValues (Default: `true`).
	Reuse pulumi.BoolPtrInput
	// Set helm values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2 (Default:
	// `[]`).
	Sets pulumi.StringArrayInput
	// values in raw yaml to pass to helm. (Default: `empty`).
	Values pulumi.StringPtrInput
	// Version of Cilium (Default: `v1.15.4`).
	Version pulumi.StringPtrInput
	// Wait for Cilium status is ok (Default: `true`).
	Wait pulumi.BoolPtrInput
}

func (DeployState) ElementType() reflect.Type {
	return reflect.TypeOf((*deployState)(nil)).Elem()
}

type deployArgs struct {
	// Datapath mode to use { tunnel | native | aws-eni | gke | azure | aks-byocni } (Default: `autodetected`).
	DataPath *string `pulumi:"dataPath"`
	// Helm chart repository to download Cilium charts from (Default: `https://helm.cilium.io`).
	Repository *string `pulumi:"repository"`
	// When upgrading, reset the helm values to the ones built into the chart (Default: `false`).
	Reset *bool `pulumi:"reset"`
	// When upgrading, reuse the helm values from the latest release unless any overrides from are set from other flags. This
	// option takes precedence over HelmResetValues (Default: `true`).
	Reuse *bool `pulumi:"reuse"`
	// Set helm values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2 (Default:
	// `[]`).
	Sets []string `pulumi:"sets"`
	// values in raw yaml to pass to helm. (Default: `empty`).
	Values *string `pulumi:"values"`
	// Version of Cilium (Default: `v1.15.4`).
	Version *string `pulumi:"version"`
	// Wait for Cilium status is ok (Default: `true`).
	Wait *bool `pulumi:"wait"`
}

// The set of arguments for constructing a Deploy resource.
type DeployArgs struct {
	// Datapath mode to use { tunnel | native | aws-eni | gke | azure | aks-byocni } (Default: `autodetected`).
	DataPath pulumi.StringPtrInput
	// Helm chart repository to download Cilium charts from (Default: `https://helm.cilium.io`).
	Repository pulumi.StringPtrInput
	// When upgrading, reset the helm values to the ones built into the chart (Default: `false`).
	Reset pulumi.BoolPtrInput
	// When upgrading, reuse the helm values from the latest release unless any overrides from are set from other flags. This
	// option takes precedence over HelmResetValues (Default: `true`).
	Reuse pulumi.BoolPtrInput
	// Set helm values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2 (Default:
	// `[]`).
	Sets pulumi.StringArrayInput
	// values in raw yaml to pass to helm. (Default: `empty`).
	Values pulumi.StringPtrInput
	// Version of Cilium (Default: `v1.15.4`).
	Version pulumi.StringPtrInput
	// Wait for Cilium status is ok (Default: `true`).
	Wait pulumi.BoolPtrInput
}

func (DeployArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deployArgs)(nil)).Elem()
}

type DeployInput interface {
	pulumi.Input

	ToDeployOutput() DeployOutput
	ToDeployOutputWithContext(ctx context.Context) DeployOutput
}

func (*Deploy) ElementType() reflect.Type {
	return reflect.TypeOf((**Deploy)(nil)).Elem()
}

func (i *Deploy) ToDeployOutput() DeployOutput {
	return i.ToDeployOutputWithContext(context.Background())
}

func (i *Deploy) ToDeployOutputWithContext(ctx context.Context) DeployOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployOutput)
}

// DeployArrayInput is an input type that accepts DeployArray and DeployArrayOutput values.
// You can construct a concrete instance of `DeployArrayInput` via:
//
//	DeployArray{ DeployArgs{...} }
type DeployArrayInput interface {
	pulumi.Input

	ToDeployArrayOutput() DeployArrayOutput
	ToDeployArrayOutputWithContext(context.Context) DeployArrayOutput
}

type DeployArray []DeployInput

func (DeployArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Deploy)(nil)).Elem()
}

func (i DeployArray) ToDeployArrayOutput() DeployArrayOutput {
	return i.ToDeployArrayOutputWithContext(context.Background())
}

func (i DeployArray) ToDeployArrayOutputWithContext(ctx context.Context) DeployArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployArrayOutput)
}

// DeployMapInput is an input type that accepts DeployMap and DeployMapOutput values.
// You can construct a concrete instance of `DeployMapInput` via:
//
//	DeployMap{ "key": DeployArgs{...} }
type DeployMapInput interface {
	pulumi.Input

	ToDeployMapOutput() DeployMapOutput
	ToDeployMapOutputWithContext(context.Context) DeployMapOutput
}

type DeployMap map[string]DeployInput

func (DeployMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Deploy)(nil)).Elem()
}

func (i DeployMap) ToDeployMapOutput() DeployMapOutput {
	return i.ToDeployMapOutputWithContext(context.Background())
}

func (i DeployMap) ToDeployMapOutputWithContext(ctx context.Context) DeployMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployMapOutput)
}

type DeployOutput struct{ *pulumi.OutputState }

func (DeployOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Deploy)(nil)).Elem()
}

func (o DeployOutput) ToDeployOutput() DeployOutput {
	return o
}

func (o DeployOutput) ToDeployOutputWithContext(ctx context.Context) DeployOutput {
	return o
}

// Datapath mode to use { tunnel | native | aws-eni | gke | azure | aks-byocni } (Default: `autodetected`).
func (o DeployOutput) DataPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Deploy) pulumi.StringOutput { return v.DataPath }).(pulumi.StringOutput)
}

// Helm values (`helm get values -n kube-system cilium`)
func (o DeployOutput) HelmValues() pulumi.StringOutput {
	return o.ApplyT(func(v *Deploy) pulumi.StringOutput { return v.HelmValues }).(pulumi.StringOutput)
}

// Helm chart repository to download Cilium charts from (Default: `https://helm.cilium.io`).
func (o DeployOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *Deploy) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

// When upgrading, reset the helm values to the ones built into the chart (Default: `false`).
func (o DeployOutput) Reset() pulumi.BoolOutput {
	return o.ApplyT(func(v *Deploy) pulumi.BoolOutput { return v.Reset }).(pulumi.BoolOutput)
}

// When upgrading, reuse the helm values from the latest release unless any overrides from are set from other flags. This
// option takes precedence over HelmResetValues (Default: `true`).
func (o DeployOutput) Reuse() pulumi.BoolOutput {
	return o.ApplyT(func(v *Deploy) pulumi.BoolOutput { return v.Reuse }).(pulumi.BoolOutput)
}

// Set helm values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2 (Default:
// `[]`).
func (o DeployOutput) Sets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Deploy) pulumi.StringArrayOutput { return v.Sets }).(pulumi.StringArrayOutput)
}

// values in raw yaml to pass to helm. (Default: `empty`).
func (o DeployOutput) Values() pulumi.StringOutput {
	return o.ApplyT(func(v *Deploy) pulumi.StringOutput { return v.Values }).(pulumi.StringOutput)
}

// Version of Cilium (Default: `v1.15.4`).
func (o DeployOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Deploy) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

// Wait for Cilium status is ok (Default: `true`).
func (o DeployOutput) Wait() pulumi.BoolOutput {
	return o.ApplyT(func(v *Deploy) pulumi.BoolOutput { return v.Wait }).(pulumi.BoolOutput)
}

type DeployArrayOutput struct{ *pulumi.OutputState }

func (DeployArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Deploy)(nil)).Elem()
}

func (o DeployArrayOutput) ToDeployArrayOutput() DeployArrayOutput {
	return o
}

func (o DeployArrayOutput) ToDeployArrayOutputWithContext(ctx context.Context) DeployArrayOutput {
	return o
}

func (o DeployArrayOutput) Index(i pulumi.IntInput) DeployOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Deploy {
		return vs[0].([]*Deploy)[vs[1].(int)]
	}).(DeployOutput)
}

type DeployMapOutput struct{ *pulumi.OutputState }

func (DeployMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Deploy)(nil)).Elem()
}

func (o DeployMapOutput) ToDeployMapOutput() DeployMapOutput {
	return o
}

func (o DeployMapOutput) ToDeployMapOutputWithContext(ctx context.Context) DeployMapOutput {
	return o
}

func (o DeployMapOutput) MapIndex(k pulumi.StringInput) DeployOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Deploy {
		return vs[0].(map[string]*Deploy)[vs[1].(string)]
	}).(DeployOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeployInput)(nil)).Elem(), &Deploy{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeployArrayInput)(nil)).Elem(), DeployArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeployMapInput)(nil)).Elem(), DeployMap{})
	pulumi.RegisterOutputType(DeployOutput{})
	pulumi.RegisterOutputType(DeployArrayOutput{})
	pulumi.RegisterOutputType(DeployMapOutput{})
}