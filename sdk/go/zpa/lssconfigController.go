// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

type LSSConfigController struct {
	pulumi.CustomResourceState

	Config LSSConfigControllerConfigPtrOutput `pulumi:"config"`
	// App Connector Group(s) to be added to the LSS configuration
	ConnectorGroups    LSSConfigControllerConnectorGroupArrayOutput   `pulumi:"connectorGroups"`
	PolicyRuleId       pulumi.StringOutput                            `pulumi:"policyRuleId"`
	PolicyRuleResource LSSConfigControllerPolicyRuleResourcePtrOutput `pulumi:"policyRuleResource"`
}

// NewLSSConfigController registers a new resource with the given unique name, arguments, and options.
func NewLSSConfigController(ctx *pulumi.Context,
	name string, args *LSSConfigControllerArgs, opts ...pulumi.ResourceOption) (*LSSConfigController, error) {
	if args == nil {
		args = &LSSConfigControllerArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LSSConfigController
	err := ctx.RegisterResource("zpa:index/lSSConfigController:LSSConfigController", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLSSConfigController gets an existing LSSConfigController resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLSSConfigController(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LSSConfigControllerState, opts ...pulumi.ResourceOption) (*LSSConfigController, error) {
	var resource LSSConfigController
	err := ctx.ReadResource("zpa:index/lSSConfigController:LSSConfigController", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LSSConfigController resources.
type lssconfigControllerState struct {
	Config *LSSConfigControllerConfig `pulumi:"config"`
	// App Connector Group(s) to be added to the LSS configuration
	ConnectorGroups    []LSSConfigControllerConnectorGroup    `pulumi:"connectorGroups"`
	PolicyRuleId       *string                                `pulumi:"policyRuleId"`
	PolicyRuleResource *LSSConfigControllerPolicyRuleResource `pulumi:"policyRuleResource"`
}

type LSSConfigControllerState struct {
	Config LSSConfigControllerConfigPtrInput
	// App Connector Group(s) to be added to the LSS configuration
	ConnectorGroups    LSSConfigControllerConnectorGroupArrayInput
	PolicyRuleId       pulumi.StringPtrInput
	PolicyRuleResource LSSConfigControllerPolicyRuleResourcePtrInput
}

func (LSSConfigControllerState) ElementType() reflect.Type {
	return reflect.TypeOf((*lssconfigControllerState)(nil)).Elem()
}

type lssconfigControllerArgs struct {
	Config *LSSConfigControllerConfig `pulumi:"config"`
	// App Connector Group(s) to be added to the LSS configuration
	ConnectorGroups    []LSSConfigControllerConnectorGroup    `pulumi:"connectorGroups"`
	PolicyRuleResource *LSSConfigControllerPolicyRuleResource `pulumi:"policyRuleResource"`
}

// The set of arguments for constructing a LSSConfigController resource.
type LSSConfigControllerArgs struct {
	Config LSSConfigControllerConfigPtrInput
	// App Connector Group(s) to be added to the LSS configuration
	ConnectorGroups    LSSConfigControllerConnectorGroupArrayInput
	PolicyRuleResource LSSConfigControllerPolicyRuleResourcePtrInput
}

func (LSSConfigControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lssconfigControllerArgs)(nil)).Elem()
}

type LSSConfigControllerInput interface {
	pulumi.Input

	ToLSSConfigControllerOutput() LSSConfigControllerOutput
	ToLSSConfigControllerOutputWithContext(ctx context.Context) LSSConfigControllerOutput
}

func (*LSSConfigController) ElementType() reflect.Type {
	return reflect.TypeOf((**LSSConfigController)(nil)).Elem()
}

func (i *LSSConfigController) ToLSSConfigControllerOutput() LSSConfigControllerOutput {
	return i.ToLSSConfigControllerOutputWithContext(context.Background())
}

func (i *LSSConfigController) ToLSSConfigControllerOutputWithContext(ctx context.Context) LSSConfigControllerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LSSConfigControllerOutput)
}

// LSSConfigControllerArrayInput is an input type that accepts LSSConfigControllerArray and LSSConfigControllerArrayOutput values.
// You can construct a concrete instance of `LSSConfigControllerArrayInput` via:
//
//	LSSConfigControllerArray{ LSSConfigControllerArgs{...} }
type LSSConfigControllerArrayInput interface {
	pulumi.Input

	ToLSSConfigControllerArrayOutput() LSSConfigControllerArrayOutput
	ToLSSConfigControllerArrayOutputWithContext(context.Context) LSSConfigControllerArrayOutput
}

type LSSConfigControllerArray []LSSConfigControllerInput

func (LSSConfigControllerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LSSConfigController)(nil)).Elem()
}

func (i LSSConfigControllerArray) ToLSSConfigControllerArrayOutput() LSSConfigControllerArrayOutput {
	return i.ToLSSConfigControllerArrayOutputWithContext(context.Background())
}

func (i LSSConfigControllerArray) ToLSSConfigControllerArrayOutputWithContext(ctx context.Context) LSSConfigControllerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LSSConfigControllerArrayOutput)
}

// LSSConfigControllerMapInput is an input type that accepts LSSConfigControllerMap and LSSConfigControllerMapOutput values.
// You can construct a concrete instance of `LSSConfigControllerMapInput` via:
//
//	LSSConfigControllerMap{ "key": LSSConfigControllerArgs{...} }
type LSSConfigControllerMapInput interface {
	pulumi.Input

	ToLSSConfigControllerMapOutput() LSSConfigControllerMapOutput
	ToLSSConfigControllerMapOutputWithContext(context.Context) LSSConfigControllerMapOutput
}

type LSSConfigControllerMap map[string]LSSConfigControllerInput

func (LSSConfigControllerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LSSConfigController)(nil)).Elem()
}

func (i LSSConfigControllerMap) ToLSSConfigControllerMapOutput() LSSConfigControllerMapOutput {
	return i.ToLSSConfigControllerMapOutputWithContext(context.Background())
}

func (i LSSConfigControllerMap) ToLSSConfigControllerMapOutputWithContext(ctx context.Context) LSSConfigControllerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LSSConfigControllerMapOutput)
}

type LSSConfigControllerOutput struct{ *pulumi.OutputState }

func (LSSConfigControllerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LSSConfigController)(nil)).Elem()
}

func (o LSSConfigControllerOutput) ToLSSConfigControllerOutput() LSSConfigControllerOutput {
	return o
}

func (o LSSConfigControllerOutput) ToLSSConfigControllerOutputWithContext(ctx context.Context) LSSConfigControllerOutput {
	return o
}

func (o LSSConfigControllerOutput) Config() LSSConfigControllerConfigPtrOutput {
	return o.ApplyT(func(v *LSSConfigController) LSSConfigControllerConfigPtrOutput { return v.Config }).(LSSConfigControllerConfigPtrOutput)
}

// App Connector Group(s) to be added to the LSS configuration
func (o LSSConfigControllerOutput) ConnectorGroups() LSSConfigControllerConnectorGroupArrayOutput {
	return o.ApplyT(func(v *LSSConfigController) LSSConfigControllerConnectorGroupArrayOutput { return v.ConnectorGroups }).(LSSConfigControllerConnectorGroupArrayOutput)
}

func (o LSSConfigControllerOutput) PolicyRuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *LSSConfigController) pulumi.StringOutput { return v.PolicyRuleId }).(pulumi.StringOutput)
}

func (o LSSConfigControllerOutput) PolicyRuleResource() LSSConfigControllerPolicyRuleResourcePtrOutput {
	return o.ApplyT(func(v *LSSConfigController) LSSConfigControllerPolicyRuleResourcePtrOutput {
		return v.PolicyRuleResource
	}).(LSSConfigControllerPolicyRuleResourcePtrOutput)
}

type LSSConfigControllerArrayOutput struct{ *pulumi.OutputState }

func (LSSConfigControllerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LSSConfigController)(nil)).Elem()
}

func (o LSSConfigControllerArrayOutput) ToLSSConfigControllerArrayOutput() LSSConfigControllerArrayOutput {
	return o
}

func (o LSSConfigControllerArrayOutput) ToLSSConfigControllerArrayOutputWithContext(ctx context.Context) LSSConfigControllerArrayOutput {
	return o
}

func (o LSSConfigControllerArrayOutput) Index(i pulumi.IntInput) LSSConfigControllerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LSSConfigController {
		return vs[0].([]*LSSConfigController)[vs[1].(int)]
	}).(LSSConfigControllerOutput)
}

type LSSConfigControllerMapOutput struct{ *pulumi.OutputState }

func (LSSConfigControllerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LSSConfigController)(nil)).Elem()
}

func (o LSSConfigControllerMapOutput) ToLSSConfigControllerMapOutput() LSSConfigControllerMapOutput {
	return o
}

func (o LSSConfigControllerMapOutput) ToLSSConfigControllerMapOutputWithContext(ctx context.Context) LSSConfigControllerMapOutput {
	return o
}

func (o LSSConfigControllerMapOutput) MapIndex(k pulumi.StringInput) LSSConfigControllerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LSSConfigController {
		return vs[0].(map[string]*LSSConfigController)[vs[1].(string)]
	}).(LSSConfigControllerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LSSConfigControllerInput)(nil)).Elem(), &LSSConfigController{})
	pulumi.RegisterInputType(reflect.TypeOf((*LSSConfigControllerArrayInput)(nil)).Elem(), LSSConfigControllerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LSSConfigControllerMapInput)(nil)).Elem(), LSSConfigControllerMap{})
	pulumi.RegisterOutputType(LSSConfigControllerOutput{})
	pulumi.RegisterOutputType(LSSConfigControllerArrayOutput{})
	pulumi.RegisterOutputType(LSSConfigControllerMapOutput{})
}