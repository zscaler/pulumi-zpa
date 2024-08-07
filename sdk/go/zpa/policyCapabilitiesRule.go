// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// * [Official documentation](https://help.zscaler.com/zpa/about-privileged-capabilities-policy)
// * [API documentation](https://help.zscaler.com/zpa/configuring-privileged-policies-using-api)
//
// The **zpa_policy_capabilities_rule** resource creates a policy capabilities rule in the Zscaler Private Access cloud.
//
//	⚠️ **WARNING:**: The attribute ``ruleOrder`` is now deprecated in favor of the new resource  ``policyAccessRuleReorder``
//
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
//			thisIdPController, err := zpa.GetIdPController(ctx, &zpa.GetIdPControllerArgs{
//				Name: pulumi.StringRef("IdP_Users"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			emailUserSso, err := zpa.GetSAMLAttribute(ctx, &zpa.GetSAMLAttributeArgs{
//				Name:    pulumi.StringRef("Email_IdP_Users"),
//				IdpName: pulumi.StringRef("IdP_Users"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			groupUser, err := zpa.GetSAMLAttribute(ctx, &zpa.GetSAMLAttributeArgs{
//				Name:    pulumi.StringRef("GroupName_IdP_Users"),
//				IdpName: pulumi.StringRef("IdP_Users"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			a000, err := zpa.GetSCIMGroups(ctx, &zpa.GetSCIMGroupsArgs{
//				Name:    pulumi.StringRef("A000"),
//				IdpName: pulumi.StringRef("IdP_Users"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			b000, err := zpa.GetSCIMGroups(ctx, &zpa.GetSCIMGroupsArgs{
//				Name:    pulumi.StringRef("B000"),
//				IdpName: pulumi.StringRef("IdP_Users"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.NewPolicyCapabilitiesRule(ctx, "thisPolicyCapabilitiesRule", &zpa.PolicyCapabilitiesRuleArgs{
//				Description: pulumi.String("Example"),
//				Action:      pulumi.String("CHECK_CAPABILITIES"),
//				PrivilegedCapabilities: &zpa.PolicyCapabilitiesRulePrivilegedCapabilitiesArgs{
//					FileUpload:        pulumi.Bool(true),
//					FileDownload:      pulumi.Bool(true),
//					InspectFileUpload: pulumi.Bool(true),
//					ClipboardCopy:     pulumi.Bool(true),
//					ClipboardPaste:    pulumi.Bool(true),
//					RecordSession:     pulumi.Bool(true),
//				},
//				Conditions: zpa.PolicyCapabilitiesRuleConditionArray{
//					&zpa.PolicyCapabilitiesRuleConditionArgs{
//						Operator: pulumi.String("OR"),
//						Operands: zpa.PolicyCapabilitiesRuleConditionOperandArray{
//							&zpa.PolicyCapabilitiesRuleConditionOperandArgs{
//								ObjectType: pulumi.String("SAML"),
//								EntryValues: zpa.PolicyCapabilitiesRuleConditionOperandEntryValueArray{
//									&zpa.PolicyCapabilitiesRuleConditionOperandEntryValueArgs{
//										Rhs: pulumi.String("user1@example.com"),
//										Lhs: pulumi.String(emailUserSso.Id),
//									},
//									&zpa.PolicyCapabilitiesRuleConditionOperandEntryValueArgs{
//										Rhs: pulumi.String("A000"),
//										Lhs: pulumi.String(groupUser.Id),
//									},
//								},
//							},
//							&zpa.PolicyCapabilitiesRuleConditionOperandArgs{
//								ObjectType: pulumi.String("SCIM_GROUP"),
//								EntryValues: zpa.PolicyCapabilitiesRuleConditionOperandEntryValueArray{
//									&zpa.PolicyCapabilitiesRuleConditionOperandEntryValueArgs{
//										Rhs: pulumi.String(a000.Id),
//										Lhs: pulumi.String(thisIdPController.Id),
//									},
//									&zpa.PolicyCapabilitiesRuleConditionOperandEntryValueArgs{
//										Rhs: pulumi.String(b000.Id),
//										Lhs: pulumi.String(thisIdPController.Id),
//									},
//								},
//							},
//						},
//					},
//				},
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
// ## LHS and RHS Values
//
// | Object Type | LHS| RHS| VALUES
// |----------|-----------|----------|----------
// | APP |   |  | “applicationSegmentId“
// | APP_GROUP |   |  | “segmentGroupId“
// | SAML | “samlAttributeId“  | “attributeValueToMatch“ |
// | SCIM | “scimAttributeId“  | “attributeValueToMatch“  |
// | SCIM_GROUP | “scimGroupAttributeId“  | “attributeValueToMatch“  |
//
// ## Import
//
// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
//
// # Visit
//
// Policy access capability can be imported by using `<RULE ID>` as the import ID.
//
// For example:
//
// ```sh
// $ pulumi import zpa:index/policyCapabilitiesRule:PolicyCapabilitiesRule example <rule_id>
// ```
type PolicyCapabilitiesRule struct {
	pulumi.CustomResourceState

	// This is for providing the rule action. Supported value: `CHECK_CAPABILITIES`
	Action pulumi.StringPtrOutput `pulumi:"action"`
	// This is for proviidng the set of conditions for the policy.
	Conditions PolicyCapabilitiesRuleConditionArrayOutput `pulumi:"conditions"`
	// This is the description of the access policy.
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	MicrotenantId pulumi.StringOutput    `pulumi:"microtenantId"`
	// This is the name of the policy rule.
	Name                   pulumi.StringOutput                                `pulumi:"name"`
	PolicySetId            pulumi.StringOutput                                `pulumi:"policySetId"`
	PrivilegedCapabilities PolicyCapabilitiesRulePrivilegedCapabilitiesOutput `pulumi:"privilegedCapabilities"`
}

// NewPolicyCapabilitiesRule registers a new resource with the given unique name, arguments, and options.
func NewPolicyCapabilitiesRule(ctx *pulumi.Context,
	name string, args *PolicyCapabilitiesRuleArgs, opts ...pulumi.ResourceOption) (*PolicyCapabilitiesRule, error) {
	if args == nil {
		args = &PolicyCapabilitiesRuleArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PolicyCapabilitiesRule
	err := ctx.RegisterResource("zpa:index/policyCapabilitiesRule:PolicyCapabilitiesRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyCapabilitiesRule gets an existing PolicyCapabilitiesRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyCapabilitiesRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyCapabilitiesRuleState, opts ...pulumi.ResourceOption) (*PolicyCapabilitiesRule, error) {
	var resource PolicyCapabilitiesRule
	err := ctx.ReadResource("zpa:index/policyCapabilitiesRule:PolicyCapabilitiesRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyCapabilitiesRule resources.
type policyCapabilitiesRuleState struct {
	// This is for providing the rule action. Supported value: `CHECK_CAPABILITIES`
	Action *string `pulumi:"action"`
	// This is for proviidng the set of conditions for the policy.
	Conditions []PolicyCapabilitiesRuleCondition `pulumi:"conditions"`
	// This is the description of the access policy.
	Description   *string `pulumi:"description"`
	MicrotenantId *string `pulumi:"microtenantId"`
	// This is the name of the policy rule.
	Name                   *string                                       `pulumi:"name"`
	PolicySetId            *string                                       `pulumi:"policySetId"`
	PrivilegedCapabilities *PolicyCapabilitiesRulePrivilegedCapabilities `pulumi:"privilegedCapabilities"`
}

type PolicyCapabilitiesRuleState struct {
	// This is for providing the rule action. Supported value: `CHECK_CAPABILITIES`
	Action pulumi.StringPtrInput
	// This is for proviidng the set of conditions for the policy.
	Conditions PolicyCapabilitiesRuleConditionArrayInput
	// This is the description of the access policy.
	Description   pulumi.StringPtrInput
	MicrotenantId pulumi.StringPtrInput
	// This is the name of the policy rule.
	Name                   pulumi.StringPtrInput
	PolicySetId            pulumi.StringPtrInput
	PrivilegedCapabilities PolicyCapabilitiesRulePrivilegedCapabilitiesPtrInput
}

func (PolicyCapabilitiesRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyCapabilitiesRuleState)(nil)).Elem()
}

type policyCapabilitiesRuleArgs struct {
	// This is for providing the rule action. Supported value: `CHECK_CAPABILITIES`
	Action *string `pulumi:"action"`
	// This is for proviidng the set of conditions for the policy.
	Conditions []PolicyCapabilitiesRuleCondition `pulumi:"conditions"`
	// This is the description of the access policy.
	Description   *string `pulumi:"description"`
	MicrotenantId *string `pulumi:"microtenantId"`
	// This is the name of the policy rule.
	Name                   *string                                       `pulumi:"name"`
	PrivilegedCapabilities *PolicyCapabilitiesRulePrivilegedCapabilities `pulumi:"privilegedCapabilities"`
}

// The set of arguments for constructing a PolicyCapabilitiesRule resource.
type PolicyCapabilitiesRuleArgs struct {
	// This is for providing the rule action. Supported value: `CHECK_CAPABILITIES`
	Action pulumi.StringPtrInput
	// This is for proviidng the set of conditions for the policy.
	Conditions PolicyCapabilitiesRuleConditionArrayInput
	// This is the description of the access policy.
	Description   pulumi.StringPtrInput
	MicrotenantId pulumi.StringPtrInput
	// This is the name of the policy rule.
	Name                   pulumi.StringPtrInput
	PrivilegedCapabilities PolicyCapabilitiesRulePrivilegedCapabilitiesPtrInput
}

func (PolicyCapabilitiesRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyCapabilitiesRuleArgs)(nil)).Elem()
}

type PolicyCapabilitiesRuleInput interface {
	pulumi.Input

	ToPolicyCapabilitiesRuleOutput() PolicyCapabilitiesRuleOutput
	ToPolicyCapabilitiesRuleOutputWithContext(ctx context.Context) PolicyCapabilitiesRuleOutput
}

func (*PolicyCapabilitiesRule) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyCapabilitiesRule)(nil)).Elem()
}

func (i *PolicyCapabilitiesRule) ToPolicyCapabilitiesRuleOutput() PolicyCapabilitiesRuleOutput {
	return i.ToPolicyCapabilitiesRuleOutputWithContext(context.Background())
}

func (i *PolicyCapabilitiesRule) ToPolicyCapabilitiesRuleOutputWithContext(ctx context.Context) PolicyCapabilitiesRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyCapabilitiesRuleOutput)
}

// PolicyCapabilitiesRuleArrayInput is an input type that accepts PolicyCapabilitiesRuleArray and PolicyCapabilitiesRuleArrayOutput values.
// You can construct a concrete instance of `PolicyCapabilitiesRuleArrayInput` via:
//
//	PolicyCapabilitiesRuleArray{ PolicyCapabilitiesRuleArgs{...} }
type PolicyCapabilitiesRuleArrayInput interface {
	pulumi.Input

	ToPolicyCapabilitiesRuleArrayOutput() PolicyCapabilitiesRuleArrayOutput
	ToPolicyCapabilitiesRuleArrayOutputWithContext(context.Context) PolicyCapabilitiesRuleArrayOutput
}

type PolicyCapabilitiesRuleArray []PolicyCapabilitiesRuleInput

func (PolicyCapabilitiesRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyCapabilitiesRule)(nil)).Elem()
}

func (i PolicyCapabilitiesRuleArray) ToPolicyCapabilitiesRuleArrayOutput() PolicyCapabilitiesRuleArrayOutput {
	return i.ToPolicyCapabilitiesRuleArrayOutputWithContext(context.Background())
}

func (i PolicyCapabilitiesRuleArray) ToPolicyCapabilitiesRuleArrayOutputWithContext(ctx context.Context) PolicyCapabilitiesRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyCapabilitiesRuleArrayOutput)
}

// PolicyCapabilitiesRuleMapInput is an input type that accepts PolicyCapabilitiesRuleMap and PolicyCapabilitiesRuleMapOutput values.
// You can construct a concrete instance of `PolicyCapabilitiesRuleMapInput` via:
//
//	PolicyCapabilitiesRuleMap{ "key": PolicyCapabilitiesRuleArgs{...} }
type PolicyCapabilitiesRuleMapInput interface {
	pulumi.Input

	ToPolicyCapabilitiesRuleMapOutput() PolicyCapabilitiesRuleMapOutput
	ToPolicyCapabilitiesRuleMapOutputWithContext(context.Context) PolicyCapabilitiesRuleMapOutput
}

type PolicyCapabilitiesRuleMap map[string]PolicyCapabilitiesRuleInput

func (PolicyCapabilitiesRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyCapabilitiesRule)(nil)).Elem()
}

func (i PolicyCapabilitiesRuleMap) ToPolicyCapabilitiesRuleMapOutput() PolicyCapabilitiesRuleMapOutput {
	return i.ToPolicyCapabilitiesRuleMapOutputWithContext(context.Background())
}

func (i PolicyCapabilitiesRuleMap) ToPolicyCapabilitiesRuleMapOutputWithContext(ctx context.Context) PolicyCapabilitiesRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyCapabilitiesRuleMapOutput)
}

type PolicyCapabilitiesRuleOutput struct{ *pulumi.OutputState }

func (PolicyCapabilitiesRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyCapabilitiesRule)(nil)).Elem()
}

func (o PolicyCapabilitiesRuleOutput) ToPolicyCapabilitiesRuleOutput() PolicyCapabilitiesRuleOutput {
	return o
}

func (o PolicyCapabilitiesRuleOutput) ToPolicyCapabilitiesRuleOutputWithContext(ctx context.Context) PolicyCapabilitiesRuleOutput {
	return o
}

// This is for providing the rule action. Supported value: `CHECK_CAPABILITIES`
func (o PolicyCapabilitiesRuleOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyCapabilitiesRule) pulumi.StringPtrOutput { return v.Action }).(pulumi.StringPtrOutput)
}

// This is for proviidng the set of conditions for the policy.
func (o PolicyCapabilitiesRuleOutput) Conditions() PolicyCapabilitiesRuleConditionArrayOutput {
	return o.ApplyT(func(v *PolicyCapabilitiesRule) PolicyCapabilitiesRuleConditionArrayOutput { return v.Conditions }).(PolicyCapabilitiesRuleConditionArrayOutput)
}

// This is the description of the access policy.
func (o PolicyCapabilitiesRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyCapabilitiesRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PolicyCapabilitiesRuleOutput) MicrotenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyCapabilitiesRule) pulumi.StringOutput { return v.MicrotenantId }).(pulumi.StringOutput)
}

// This is the name of the policy rule.
func (o PolicyCapabilitiesRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyCapabilitiesRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PolicyCapabilitiesRuleOutput) PolicySetId() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyCapabilitiesRule) pulumi.StringOutput { return v.PolicySetId }).(pulumi.StringOutput)
}

func (o PolicyCapabilitiesRuleOutput) PrivilegedCapabilities() PolicyCapabilitiesRulePrivilegedCapabilitiesOutput {
	return o.ApplyT(func(v *PolicyCapabilitiesRule) PolicyCapabilitiesRulePrivilegedCapabilitiesOutput {
		return v.PrivilegedCapabilities
	}).(PolicyCapabilitiesRulePrivilegedCapabilitiesOutput)
}

type PolicyCapabilitiesRuleArrayOutput struct{ *pulumi.OutputState }

func (PolicyCapabilitiesRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyCapabilitiesRule)(nil)).Elem()
}

func (o PolicyCapabilitiesRuleArrayOutput) ToPolicyCapabilitiesRuleArrayOutput() PolicyCapabilitiesRuleArrayOutput {
	return o
}

func (o PolicyCapabilitiesRuleArrayOutput) ToPolicyCapabilitiesRuleArrayOutputWithContext(ctx context.Context) PolicyCapabilitiesRuleArrayOutput {
	return o
}

func (o PolicyCapabilitiesRuleArrayOutput) Index(i pulumi.IntInput) PolicyCapabilitiesRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PolicyCapabilitiesRule {
		return vs[0].([]*PolicyCapabilitiesRule)[vs[1].(int)]
	}).(PolicyCapabilitiesRuleOutput)
}

type PolicyCapabilitiesRuleMapOutput struct{ *pulumi.OutputState }

func (PolicyCapabilitiesRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyCapabilitiesRule)(nil)).Elem()
}

func (o PolicyCapabilitiesRuleMapOutput) ToPolicyCapabilitiesRuleMapOutput() PolicyCapabilitiesRuleMapOutput {
	return o
}

func (o PolicyCapabilitiesRuleMapOutput) ToPolicyCapabilitiesRuleMapOutputWithContext(ctx context.Context) PolicyCapabilitiesRuleMapOutput {
	return o
}

func (o PolicyCapabilitiesRuleMapOutput) MapIndex(k pulumi.StringInput) PolicyCapabilitiesRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PolicyCapabilitiesRule {
		return vs[0].(map[string]*PolicyCapabilitiesRule)[vs[1].(string)]
	}).(PolicyCapabilitiesRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyCapabilitiesRuleInput)(nil)).Elem(), &PolicyCapabilitiesRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyCapabilitiesRuleArrayInput)(nil)).Elem(), PolicyCapabilitiesRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyCapabilitiesRuleMapInput)(nil)).Elem(), PolicyCapabilitiesRuleMap{})
	pulumi.RegisterOutputType(PolicyCapabilitiesRuleOutput{})
	pulumi.RegisterOutputType(PolicyCapabilitiesRuleArrayOutput{})
	pulumi.RegisterOutputType(PolicyCapabilitiesRuleMapOutput{})
}
