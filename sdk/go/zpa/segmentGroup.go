// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/zscaler/pulumi-zpa/sdk/go/zpa"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// ZPA Segment Group resource
//			_, err := zpa.NewSegmentGroup(ctx, "testSegmentGroup", &zpa.SegmentGroupArgs{
//				Description: pulumi.String("test1-segment-group"),
//				Enabled:     pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
//
// # Visit
//
// **segment_group** can be imported by using `<SEGMENT GROUP ID>` or `<SEGMENT GROUP NAME>` as the import ID.
//
// For example:
//
// ```sh
// $ pulumi import zpa:index/segmentGroup:SegmentGroup example <segment_group_id>
// ```
//
// or
//
// ```sh
// $ pulumi import zpa:index/segmentGroup:SegmentGroup example <segment_group_name>
// ```
type SegmentGroup struct {
	pulumi.CustomResourceState

	Applications SegmentGroupApplicationArrayOutput `pulumi:"applications"`
	// Description of the app group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether this app group is enabled or not.
	Enabled       pulumi.BoolPtrOutput `pulumi:"enabled"`
	MicrotenantId pulumi.StringOutput  `pulumi:"microtenantId"`
	// Name of the app group.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewSegmentGroup registers a new resource with the given unique name, arguments, and options.
func NewSegmentGroup(ctx *pulumi.Context,
	name string, args *SegmentGroupArgs, opts ...pulumi.ResourceOption) (*SegmentGroup, error) {
	if args == nil {
		args = &SegmentGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SegmentGroup
	err := ctx.RegisterResource("zpa:index/segmentGroup:SegmentGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSegmentGroup gets an existing SegmentGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSegmentGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SegmentGroupState, opts ...pulumi.ResourceOption) (*SegmentGroup, error) {
	var resource SegmentGroup
	err := ctx.ReadResource("zpa:index/segmentGroup:SegmentGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SegmentGroup resources.
type segmentGroupState struct {
	Applications []SegmentGroupApplication `pulumi:"applications"`
	// Description of the app group.
	Description *string `pulumi:"description"`
	// Whether this app group is enabled or not.
	Enabled       *bool   `pulumi:"enabled"`
	MicrotenantId *string `pulumi:"microtenantId"`
	// Name of the app group.
	Name *string `pulumi:"name"`
}

type SegmentGroupState struct {
	Applications SegmentGroupApplicationArrayInput
	// Description of the app group.
	Description pulumi.StringPtrInput
	// Whether this app group is enabled or not.
	Enabled       pulumi.BoolPtrInput
	MicrotenantId pulumi.StringPtrInput
	// Name of the app group.
	Name pulumi.StringPtrInput
}

func (SegmentGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*segmentGroupState)(nil)).Elem()
}

type segmentGroupArgs struct {
	Applications []SegmentGroupApplication `pulumi:"applications"`
	// Description of the app group.
	Description *string `pulumi:"description"`
	// Whether this app group is enabled or not.
	Enabled       *bool   `pulumi:"enabled"`
	MicrotenantId *string `pulumi:"microtenantId"`
	// Name of the app group.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a SegmentGroup resource.
type SegmentGroupArgs struct {
	Applications SegmentGroupApplicationArrayInput
	// Description of the app group.
	Description pulumi.StringPtrInput
	// Whether this app group is enabled or not.
	Enabled       pulumi.BoolPtrInput
	MicrotenantId pulumi.StringPtrInput
	// Name of the app group.
	Name pulumi.StringPtrInput
}

func (SegmentGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*segmentGroupArgs)(nil)).Elem()
}

type SegmentGroupInput interface {
	pulumi.Input

	ToSegmentGroupOutput() SegmentGroupOutput
	ToSegmentGroupOutputWithContext(ctx context.Context) SegmentGroupOutput
}

func (*SegmentGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**SegmentGroup)(nil)).Elem()
}

func (i *SegmentGroup) ToSegmentGroupOutput() SegmentGroupOutput {
	return i.ToSegmentGroupOutputWithContext(context.Background())
}

func (i *SegmentGroup) ToSegmentGroupOutputWithContext(ctx context.Context) SegmentGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SegmentGroupOutput)
}

// SegmentGroupArrayInput is an input type that accepts SegmentGroupArray and SegmentGroupArrayOutput values.
// You can construct a concrete instance of `SegmentGroupArrayInput` via:
//
//	SegmentGroupArray{ SegmentGroupArgs{...} }
type SegmentGroupArrayInput interface {
	pulumi.Input

	ToSegmentGroupArrayOutput() SegmentGroupArrayOutput
	ToSegmentGroupArrayOutputWithContext(context.Context) SegmentGroupArrayOutput
}

type SegmentGroupArray []SegmentGroupInput

func (SegmentGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SegmentGroup)(nil)).Elem()
}

func (i SegmentGroupArray) ToSegmentGroupArrayOutput() SegmentGroupArrayOutput {
	return i.ToSegmentGroupArrayOutputWithContext(context.Background())
}

func (i SegmentGroupArray) ToSegmentGroupArrayOutputWithContext(ctx context.Context) SegmentGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SegmentGroupArrayOutput)
}

// SegmentGroupMapInput is an input type that accepts SegmentGroupMap and SegmentGroupMapOutput values.
// You can construct a concrete instance of `SegmentGroupMapInput` via:
//
//	SegmentGroupMap{ "key": SegmentGroupArgs{...} }
type SegmentGroupMapInput interface {
	pulumi.Input

	ToSegmentGroupMapOutput() SegmentGroupMapOutput
	ToSegmentGroupMapOutputWithContext(context.Context) SegmentGroupMapOutput
}

type SegmentGroupMap map[string]SegmentGroupInput

func (SegmentGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SegmentGroup)(nil)).Elem()
}

func (i SegmentGroupMap) ToSegmentGroupMapOutput() SegmentGroupMapOutput {
	return i.ToSegmentGroupMapOutputWithContext(context.Background())
}

func (i SegmentGroupMap) ToSegmentGroupMapOutputWithContext(ctx context.Context) SegmentGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SegmentGroupMapOutput)
}

type SegmentGroupOutput struct{ *pulumi.OutputState }

func (SegmentGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SegmentGroup)(nil)).Elem()
}

func (o SegmentGroupOutput) ToSegmentGroupOutput() SegmentGroupOutput {
	return o
}

func (o SegmentGroupOutput) ToSegmentGroupOutputWithContext(ctx context.Context) SegmentGroupOutput {
	return o
}

func (o SegmentGroupOutput) Applications() SegmentGroupApplicationArrayOutput {
	return o.ApplyT(func(v *SegmentGroup) SegmentGroupApplicationArrayOutput { return v.Applications }).(SegmentGroupApplicationArrayOutput)
}

// Description of the app group.
func (o SegmentGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SegmentGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether this app group is enabled or not.
func (o SegmentGroupOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SegmentGroup) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o SegmentGroupOutput) MicrotenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *SegmentGroup) pulumi.StringOutput { return v.MicrotenantId }).(pulumi.StringOutput)
}

// Name of the app group.
func (o SegmentGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SegmentGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type SegmentGroupArrayOutput struct{ *pulumi.OutputState }

func (SegmentGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SegmentGroup)(nil)).Elem()
}

func (o SegmentGroupArrayOutput) ToSegmentGroupArrayOutput() SegmentGroupArrayOutput {
	return o
}

func (o SegmentGroupArrayOutput) ToSegmentGroupArrayOutputWithContext(ctx context.Context) SegmentGroupArrayOutput {
	return o
}

func (o SegmentGroupArrayOutput) Index(i pulumi.IntInput) SegmentGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SegmentGroup {
		return vs[0].([]*SegmentGroup)[vs[1].(int)]
	}).(SegmentGroupOutput)
}

type SegmentGroupMapOutput struct{ *pulumi.OutputState }

func (SegmentGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SegmentGroup)(nil)).Elem()
}

func (o SegmentGroupMapOutput) ToSegmentGroupMapOutput() SegmentGroupMapOutput {
	return o
}

func (o SegmentGroupMapOutput) ToSegmentGroupMapOutputWithContext(ctx context.Context) SegmentGroupMapOutput {
	return o
}

func (o SegmentGroupMapOutput) MapIndex(k pulumi.StringInput) SegmentGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SegmentGroup {
		return vs[0].(map[string]*SegmentGroup)[vs[1].(string)]
	}).(SegmentGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SegmentGroupInput)(nil)).Elem(), &SegmentGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*SegmentGroupArrayInput)(nil)).Elem(), SegmentGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SegmentGroupMapInput)(nil)).Elem(), SegmentGroupMap{})
	pulumi.RegisterOutputType(SegmentGroupOutput{})
	pulumi.RegisterOutputType(SegmentGroupArrayOutput{})
	pulumi.RegisterOutputType(SegmentGroupMapOutput{})
}
