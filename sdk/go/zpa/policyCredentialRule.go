// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

type PolicyCredentialRule struct {
	pulumi.CustomResourceState

	// This is for providing the rule action.
	Action pulumi.StringPtrOutput `pulumi:"action"`
	// This is for proviidng the set of conditions for the policy.
	Conditions  PolicyCredentialRuleConditionArrayOutput  `pulumi:"conditions"`
	Credentials PolicyCredentialRuleCredentialArrayOutput `pulumi:"credentials"`
	// This is the description of the access policy.
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	MicrotenantId pulumi.StringOutput    `pulumi:"microtenantId"`
	// This is the name of the policy.
	Name        pulumi.StringOutput `pulumi:"name"`
	PolicySetId pulumi.StringOutput `pulumi:"policySetId"`
}

// NewPolicyCredentialRule registers a new resource with the given unique name, arguments, and options.
func NewPolicyCredentialRule(ctx *pulumi.Context,
	name string, args *PolicyCredentialRuleArgs, opts ...pulumi.ResourceOption) (*PolicyCredentialRule, error) {
	if args == nil {
		args = &PolicyCredentialRuleArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PolicyCredentialRule
	err := ctx.RegisterResource("zpa:index/policyCredentialRule:PolicyCredentialRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyCredentialRule gets an existing PolicyCredentialRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyCredentialRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyCredentialRuleState, opts ...pulumi.ResourceOption) (*PolicyCredentialRule, error) {
	var resource PolicyCredentialRule
	err := ctx.ReadResource("zpa:index/policyCredentialRule:PolicyCredentialRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyCredentialRule resources.
type policyCredentialRuleState struct {
	// This is for providing the rule action.
	Action *string `pulumi:"action"`
	// This is for proviidng the set of conditions for the policy.
	Conditions  []PolicyCredentialRuleCondition  `pulumi:"conditions"`
	Credentials []PolicyCredentialRuleCredential `pulumi:"credentials"`
	// This is the description of the access policy.
	Description   *string `pulumi:"description"`
	MicrotenantId *string `pulumi:"microtenantId"`
	// This is the name of the policy.
	Name        *string `pulumi:"name"`
	PolicySetId *string `pulumi:"policySetId"`
}

type PolicyCredentialRuleState struct {
	// This is for providing the rule action.
	Action pulumi.StringPtrInput
	// This is for proviidng the set of conditions for the policy.
	Conditions  PolicyCredentialRuleConditionArrayInput
	Credentials PolicyCredentialRuleCredentialArrayInput
	// This is the description of the access policy.
	Description   pulumi.StringPtrInput
	MicrotenantId pulumi.StringPtrInput
	// This is the name of the policy.
	Name        pulumi.StringPtrInput
	PolicySetId pulumi.StringPtrInput
}

func (PolicyCredentialRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyCredentialRuleState)(nil)).Elem()
}

type policyCredentialRuleArgs struct {
	// This is for providing the rule action.
	Action *string `pulumi:"action"`
	// This is for proviidng the set of conditions for the policy.
	Conditions  []PolicyCredentialRuleCondition  `pulumi:"conditions"`
	Credentials []PolicyCredentialRuleCredential `pulumi:"credentials"`
	// This is the description of the access policy.
	Description   *string `pulumi:"description"`
	MicrotenantId *string `pulumi:"microtenantId"`
	// This is the name of the policy.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a PolicyCredentialRule resource.
type PolicyCredentialRuleArgs struct {
	// This is for providing the rule action.
	Action pulumi.StringPtrInput
	// This is for proviidng the set of conditions for the policy.
	Conditions  PolicyCredentialRuleConditionArrayInput
	Credentials PolicyCredentialRuleCredentialArrayInput
	// This is the description of the access policy.
	Description   pulumi.StringPtrInput
	MicrotenantId pulumi.StringPtrInput
	// This is the name of the policy.
	Name pulumi.StringPtrInput
}

func (PolicyCredentialRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyCredentialRuleArgs)(nil)).Elem()
}

type PolicyCredentialRuleInput interface {
	pulumi.Input

	ToPolicyCredentialRuleOutput() PolicyCredentialRuleOutput
	ToPolicyCredentialRuleOutputWithContext(ctx context.Context) PolicyCredentialRuleOutput
}

func (*PolicyCredentialRule) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyCredentialRule)(nil)).Elem()
}

func (i *PolicyCredentialRule) ToPolicyCredentialRuleOutput() PolicyCredentialRuleOutput {
	return i.ToPolicyCredentialRuleOutputWithContext(context.Background())
}

func (i *PolicyCredentialRule) ToPolicyCredentialRuleOutputWithContext(ctx context.Context) PolicyCredentialRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyCredentialRuleOutput)
}

// PolicyCredentialRuleArrayInput is an input type that accepts PolicyCredentialRuleArray and PolicyCredentialRuleArrayOutput values.
// You can construct a concrete instance of `PolicyCredentialRuleArrayInput` via:
//
//	PolicyCredentialRuleArray{ PolicyCredentialRuleArgs{...} }
type PolicyCredentialRuleArrayInput interface {
	pulumi.Input

	ToPolicyCredentialRuleArrayOutput() PolicyCredentialRuleArrayOutput
	ToPolicyCredentialRuleArrayOutputWithContext(context.Context) PolicyCredentialRuleArrayOutput
}

type PolicyCredentialRuleArray []PolicyCredentialRuleInput

func (PolicyCredentialRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyCredentialRule)(nil)).Elem()
}

func (i PolicyCredentialRuleArray) ToPolicyCredentialRuleArrayOutput() PolicyCredentialRuleArrayOutput {
	return i.ToPolicyCredentialRuleArrayOutputWithContext(context.Background())
}

func (i PolicyCredentialRuleArray) ToPolicyCredentialRuleArrayOutputWithContext(ctx context.Context) PolicyCredentialRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyCredentialRuleArrayOutput)
}

// PolicyCredentialRuleMapInput is an input type that accepts PolicyCredentialRuleMap and PolicyCredentialRuleMapOutput values.
// You can construct a concrete instance of `PolicyCredentialRuleMapInput` via:
//
//	PolicyCredentialRuleMap{ "key": PolicyCredentialRuleArgs{...} }
type PolicyCredentialRuleMapInput interface {
	pulumi.Input

	ToPolicyCredentialRuleMapOutput() PolicyCredentialRuleMapOutput
	ToPolicyCredentialRuleMapOutputWithContext(context.Context) PolicyCredentialRuleMapOutput
}

type PolicyCredentialRuleMap map[string]PolicyCredentialRuleInput

func (PolicyCredentialRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyCredentialRule)(nil)).Elem()
}

func (i PolicyCredentialRuleMap) ToPolicyCredentialRuleMapOutput() PolicyCredentialRuleMapOutput {
	return i.ToPolicyCredentialRuleMapOutputWithContext(context.Background())
}

func (i PolicyCredentialRuleMap) ToPolicyCredentialRuleMapOutputWithContext(ctx context.Context) PolicyCredentialRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyCredentialRuleMapOutput)
}

type PolicyCredentialRuleOutput struct{ *pulumi.OutputState }

func (PolicyCredentialRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyCredentialRule)(nil)).Elem()
}

func (o PolicyCredentialRuleOutput) ToPolicyCredentialRuleOutput() PolicyCredentialRuleOutput {
	return o
}

func (o PolicyCredentialRuleOutput) ToPolicyCredentialRuleOutputWithContext(ctx context.Context) PolicyCredentialRuleOutput {
	return o
}

// This is for providing the rule action.
func (o PolicyCredentialRuleOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyCredentialRule) pulumi.StringPtrOutput { return v.Action }).(pulumi.StringPtrOutput)
}

// This is for proviidng the set of conditions for the policy.
func (o PolicyCredentialRuleOutput) Conditions() PolicyCredentialRuleConditionArrayOutput {
	return o.ApplyT(func(v *PolicyCredentialRule) PolicyCredentialRuleConditionArrayOutput { return v.Conditions }).(PolicyCredentialRuleConditionArrayOutput)
}

func (o PolicyCredentialRuleOutput) Credentials() PolicyCredentialRuleCredentialArrayOutput {
	return o.ApplyT(func(v *PolicyCredentialRule) PolicyCredentialRuleCredentialArrayOutput { return v.Credentials }).(PolicyCredentialRuleCredentialArrayOutput)
}

// This is the description of the access policy.
func (o PolicyCredentialRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyCredentialRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PolicyCredentialRuleOutput) MicrotenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyCredentialRule) pulumi.StringOutput { return v.MicrotenantId }).(pulumi.StringOutput)
}

// This is the name of the policy.
func (o PolicyCredentialRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyCredentialRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PolicyCredentialRuleOutput) PolicySetId() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyCredentialRule) pulumi.StringOutput { return v.PolicySetId }).(pulumi.StringOutput)
}

type PolicyCredentialRuleArrayOutput struct{ *pulumi.OutputState }

func (PolicyCredentialRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyCredentialRule)(nil)).Elem()
}

func (o PolicyCredentialRuleArrayOutput) ToPolicyCredentialRuleArrayOutput() PolicyCredentialRuleArrayOutput {
	return o
}

func (o PolicyCredentialRuleArrayOutput) ToPolicyCredentialRuleArrayOutputWithContext(ctx context.Context) PolicyCredentialRuleArrayOutput {
	return o
}

func (o PolicyCredentialRuleArrayOutput) Index(i pulumi.IntInput) PolicyCredentialRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PolicyCredentialRule {
		return vs[0].([]*PolicyCredentialRule)[vs[1].(int)]
	}).(PolicyCredentialRuleOutput)
}

type PolicyCredentialRuleMapOutput struct{ *pulumi.OutputState }

func (PolicyCredentialRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyCredentialRule)(nil)).Elem()
}

func (o PolicyCredentialRuleMapOutput) ToPolicyCredentialRuleMapOutput() PolicyCredentialRuleMapOutput {
	return o
}

func (o PolicyCredentialRuleMapOutput) ToPolicyCredentialRuleMapOutputWithContext(ctx context.Context) PolicyCredentialRuleMapOutput {
	return o
}

func (o PolicyCredentialRuleMapOutput) MapIndex(k pulumi.StringInput) PolicyCredentialRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PolicyCredentialRule {
		return vs[0].(map[string]*PolicyCredentialRule)[vs[1].(string)]
	}).(PolicyCredentialRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyCredentialRuleInput)(nil)).Elem(), &PolicyCredentialRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyCredentialRuleArrayInput)(nil)).Elem(), PolicyCredentialRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyCredentialRuleMapInput)(nil)).Elem(), PolicyCredentialRuleMap{})
	pulumi.RegisterOutputType(PolicyCredentialRuleOutput{})
	pulumi.RegisterOutputType(PolicyCredentialRuleArrayOutput{})
	pulumi.RegisterOutputType(PolicyCredentialRuleMapOutput{})
}
