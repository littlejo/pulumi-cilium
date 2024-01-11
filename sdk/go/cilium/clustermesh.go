// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cilium

import (
	"context"
	"reflect"

	"github.com/littlejo/pulumi-cilium/sdk/go/cilium/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Cluster Mesh resource. This is equivalent to cilium cli: `cilium clustermesh enable` and `cilium clustermesh disable`: It manages the activation of Cluster Mesh on one Kubernetes cluster.
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
//			exampleInstall, err := cilium.NewInstall(ctx, "exampleInstall", &cilium.InstallArgs{
//				Sets: pulumi.StringArray{
//					pulumi.String("cluster.name=clustermesh1"),
//					pulumi.String("cluster.id=1"),
//					pulumi.String("ipam.mode=kubernetes"),
//				},
//				Version: pulumi.String("1.14.5"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cilium.NewClustermesh(ctx, "exampleClustermesh", &cilium.ClustermeshArgs{
//				ServiceType: pulumi.String("LoadBalancer"),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				exampleInstall,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Clustermesh struct {
	pulumi.CustomResourceState

	// Enable support for external workloads, such as VMs (Default: `false`).
	EnableExternalWorkloads pulumi.BoolOutput `pulumi:"enableExternalWorkloads"`
	// Enable kvstoremesh, an extension which caches remote cluster information in the local kvstore (Cilium >=1.14 only) (Default: `false`).
	EnableKvStoreMesh pulumi.BoolOutput `pulumi:"enableKvStoreMesh"`
	// Namespace in which to install (Default: `kube-system`).
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// Type of Kubernetes service to expose control plane { LoadBalancer | NodePort | ClusterIP } (Default: `autodetected`).
	ServiceType pulumi.StringOutput `pulumi:"serviceType"`
	// Wait Cluster Mesh status is ok (Default: `true`).
	Wait pulumi.BoolOutput `pulumi:"wait"`
}

// NewClustermesh registers a new resource with the given unique name, arguments, and options.
func NewClustermesh(ctx *pulumi.Context,
	name string, args *ClustermeshArgs, opts ...pulumi.ResourceOption) (*Clustermesh, error) {
	if args == nil {
		args = &ClustermeshArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Clustermesh
	err := ctx.RegisterResource("cilium:index/clustermesh:Clustermesh", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClustermesh gets an existing Clustermesh resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClustermesh(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClustermeshState, opts ...pulumi.ResourceOption) (*Clustermesh, error) {
	var resource Clustermesh
	err := ctx.ReadResource("cilium:index/clustermesh:Clustermesh", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Clustermesh resources.
type clustermeshState struct {
	// Enable support for external workloads, such as VMs (Default: `false`).
	EnableExternalWorkloads *bool `pulumi:"enableExternalWorkloads"`
	// Enable kvstoremesh, an extension which caches remote cluster information in the local kvstore (Cilium >=1.14 only) (Default: `false`).
	EnableKvStoreMesh *bool `pulumi:"enableKvStoreMesh"`
	// Namespace in which to install (Default: `kube-system`).
	Namespace *string `pulumi:"namespace"`
	// Type of Kubernetes service to expose control plane { LoadBalancer | NodePort | ClusterIP } (Default: `autodetected`).
	ServiceType *string `pulumi:"serviceType"`
	// Wait Cluster Mesh status is ok (Default: `true`).
	Wait *bool `pulumi:"wait"`
}

type ClustermeshState struct {
	// Enable support for external workloads, such as VMs (Default: `false`).
	EnableExternalWorkloads pulumi.BoolPtrInput
	// Enable kvstoremesh, an extension which caches remote cluster information in the local kvstore (Cilium >=1.14 only) (Default: `false`).
	EnableKvStoreMesh pulumi.BoolPtrInput
	// Namespace in which to install (Default: `kube-system`).
	Namespace pulumi.StringPtrInput
	// Type of Kubernetes service to expose control plane { LoadBalancer | NodePort | ClusterIP } (Default: `autodetected`).
	ServiceType pulumi.StringPtrInput
	// Wait Cluster Mesh status is ok (Default: `true`).
	Wait pulumi.BoolPtrInput
}

func (ClustermeshState) ElementType() reflect.Type {
	return reflect.TypeOf((*clustermeshState)(nil)).Elem()
}

type clustermeshArgs struct {
	// Enable support for external workloads, such as VMs (Default: `false`).
	EnableExternalWorkloads *bool `pulumi:"enableExternalWorkloads"`
	// Enable kvstoremesh, an extension which caches remote cluster information in the local kvstore (Cilium >=1.14 only) (Default: `false`).
	EnableKvStoreMesh *bool `pulumi:"enableKvStoreMesh"`
	// Namespace in which to install (Default: `kube-system`).
	Namespace *string `pulumi:"namespace"`
	// Type of Kubernetes service to expose control plane { LoadBalancer | NodePort | ClusterIP } (Default: `autodetected`).
	ServiceType *string `pulumi:"serviceType"`
	// Wait Cluster Mesh status is ok (Default: `true`).
	Wait *bool `pulumi:"wait"`
}

// The set of arguments for constructing a Clustermesh resource.
type ClustermeshArgs struct {
	// Enable support for external workloads, such as VMs (Default: `false`).
	EnableExternalWorkloads pulumi.BoolPtrInput
	// Enable kvstoremesh, an extension which caches remote cluster information in the local kvstore (Cilium >=1.14 only) (Default: `false`).
	EnableKvStoreMesh pulumi.BoolPtrInput
	// Namespace in which to install (Default: `kube-system`).
	Namespace pulumi.StringPtrInput
	// Type of Kubernetes service to expose control plane { LoadBalancer | NodePort | ClusterIP } (Default: `autodetected`).
	ServiceType pulumi.StringPtrInput
	// Wait Cluster Mesh status is ok (Default: `true`).
	Wait pulumi.BoolPtrInput
}

func (ClustermeshArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clustermeshArgs)(nil)).Elem()
}

type ClustermeshInput interface {
	pulumi.Input

	ToClustermeshOutput() ClustermeshOutput
	ToClustermeshOutputWithContext(ctx context.Context) ClustermeshOutput
}

func (*Clustermesh) ElementType() reflect.Type {
	return reflect.TypeOf((**Clustermesh)(nil)).Elem()
}

func (i *Clustermesh) ToClustermeshOutput() ClustermeshOutput {
	return i.ToClustermeshOutputWithContext(context.Background())
}

func (i *Clustermesh) ToClustermeshOutputWithContext(ctx context.Context) ClustermeshOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClustermeshOutput)
}

// ClustermeshArrayInput is an input type that accepts ClustermeshArray and ClustermeshArrayOutput values.
// You can construct a concrete instance of `ClustermeshArrayInput` via:
//
//	ClustermeshArray{ ClustermeshArgs{...} }
type ClustermeshArrayInput interface {
	pulumi.Input

	ToClustermeshArrayOutput() ClustermeshArrayOutput
	ToClustermeshArrayOutputWithContext(context.Context) ClustermeshArrayOutput
}

type ClustermeshArray []ClustermeshInput

func (ClustermeshArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Clustermesh)(nil)).Elem()
}

func (i ClustermeshArray) ToClustermeshArrayOutput() ClustermeshArrayOutput {
	return i.ToClustermeshArrayOutputWithContext(context.Background())
}

func (i ClustermeshArray) ToClustermeshArrayOutputWithContext(ctx context.Context) ClustermeshArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClustermeshArrayOutput)
}

// ClustermeshMapInput is an input type that accepts ClustermeshMap and ClustermeshMapOutput values.
// You can construct a concrete instance of `ClustermeshMapInput` via:
//
//	ClustermeshMap{ "key": ClustermeshArgs{...} }
type ClustermeshMapInput interface {
	pulumi.Input

	ToClustermeshMapOutput() ClustermeshMapOutput
	ToClustermeshMapOutputWithContext(context.Context) ClustermeshMapOutput
}

type ClustermeshMap map[string]ClustermeshInput

func (ClustermeshMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Clustermesh)(nil)).Elem()
}

func (i ClustermeshMap) ToClustermeshMapOutput() ClustermeshMapOutput {
	return i.ToClustermeshMapOutputWithContext(context.Background())
}

func (i ClustermeshMap) ToClustermeshMapOutputWithContext(ctx context.Context) ClustermeshMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClustermeshMapOutput)
}

type ClustermeshOutput struct{ *pulumi.OutputState }

func (ClustermeshOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Clustermesh)(nil)).Elem()
}

func (o ClustermeshOutput) ToClustermeshOutput() ClustermeshOutput {
	return o
}

func (o ClustermeshOutput) ToClustermeshOutputWithContext(ctx context.Context) ClustermeshOutput {
	return o
}

// Enable support for external workloads, such as VMs (Default: `false`).
func (o ClustermeshOutput) EnableExternalWorkloads() pulumi.BoolOutput {
	return o.ApplyT(func(v *Clustermesh) pulumi.BoolOutput { return v.EnableExternalWorkloads }).(pulumi.BoolOutput)
}

// Enable kvstoremesh, an extension which caches remote cluster information in the local kvstore (Cilium >=1.14 only) (Default: `false`).
func (o ClustermeshOutput) EnableKvStoreMesh() pulumi.BoolOutput {
	return o.ApplyT(func(v *Clustermesh) pulumi.BoolOutput { return v.EnableKvStoreMesh }).(pulumi.BoolOutput)
}

// Namespace in which to install (Default: `kube-system`).
func (o ClustermeshOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v *Clustermesh) pulumi.StringOutput { return v.Namespace }).(pulumi.StringOutput)
}

// Type of Kubernetes service to expose control plane { LoadBalancer | NodePort | ClusterIP } (Default: `autodetected`).
func (o ClustermeshOutput) ServiceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Clustermesh) pulumi.StringOutput { return v.ServiceType }).(pulumi.StringOutput)
}

// Wait Cluster Mesh status is ok (Default: `true`).
func (o ClustermeshOutput) Wait() pulumi.BoolOutput {
	return o.ApplyT(func(v *Clustermesh) pulumi.BoolOutput { return v.Wait }).(pulumi.BoolOutput)
}

type ClustermeshArrayOutput struct{ *pulumi.OutputState }

func (ClustermeshArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Clustermesh)(nil)).Elem()
}

func (o ClustermeshArrayOutput) ToClustermeshArrayOutput() ClustermeshArrayOutput {
	return o
}

func (o ClustermeshArrayOutput) ToClustermeshArrayOutputWithContext(ctx context.Context) ClustermeshArrayOutput {
	return o
}

func (o ClustermeshArrayOutput) Index(i pulumi.IntInput) ClustermeshOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Clustermesh {
		return vs[0].([]*Clustermesh)[vs[1].(int)]
	}).(ClustermeshOutput)
}

type ClustermeshMapOutput struct{ *pulumi.OutputState }

func (ClustermeshMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Clustermesh)(nil)).Elem()
}

func (o ClustermeshMapOutput) ToClustermeshMapOutput() ClustermeshMapOutput {
	return o
}

func (o ClustermeshMapOutput) ToClustermeshMapOutputWithContext(ctx context.Context) ClustermeshMapOutput {
	return o
}

func (o ClustermeshMapOutput) MapIndex(k pulumi.StringInput) ClustermeshOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Clustermesh {
		return vs[0].(map[string]*Clustermesh)[vs[1].(string)]
	}).(ClustermeshOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClustermeshInput)(nil)).Elem(), &Clustermesh{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClustermeshArrayInput)(nil)).Elem(), ClustermeshArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClustermeshMapInput)(nil)).Elem(), ClustermeshMap{})
	pulumi.RegisterOutputType(ClustermeshOutput{})
	pulumi.RegisterOutputType(ClustermeshArrayOutput{})
	pulumi.RegisterOutputType(ClustermeshMapOutput{})
}
