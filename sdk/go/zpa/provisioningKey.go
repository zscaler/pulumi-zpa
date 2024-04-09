// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// ## Import
//
// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
//
// # Visit
//
// Provisioning key can be imported by using `<PROVISIONING KEY ID>` or `<PROVISIONING KEY NAME>` as the import ID.
//
// For example:
//
// ```sh
// $ pulumi import zpa:index/provisioningKey:ProvisioningKey example <provisioning_key_id>
// ```
//
// or
//
// ```sh
// $ pulumi import zpa:index/provisioningKey:ProvisioningKey example <provisioning_key_name>
// ```
type ProvisioningKey struct {
	pulumi.CustomResourceState

	// read only field. Ignored in PUT/POST calls.
	ProvisioningKeyValue pulumi.StringOutput    `pulumi:"ProvisioningKeyValue"`
	AppConnectorGroupId  pulumi.StringPtrOutput `pulumi:"appConnectorGroupId"`
	// Read only property. Applicable only for GET calls, ignored in PUT/POST calls.
	AppConnectorGroupName pulumi.StringOutput `pulumi:"appConnectorGroupName"`
	// Specifies the provisioning key type for App Connectors or ZPA Private Service Edges. The supported values are
	// CONNECTOR_GRP and SERVICE_EDGE_GRP.
	AssociationType pulumi.StringOutput `pulumi:"associationType"`
	// Whether the provisioning key is enabled or not. Supported values: true, false
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// ID of the enrollment certificate that can be used for this provisioning key.
	EnrollmentCertId pulumi.StringOutput      `pulumi:"enrollmentCertId"`
	IpAcls           pulumi.StringArrayOutput `pulumi:"ipAcls"`
	// The maximum number of instances where this provisioning key can be used for enrolling an App Connector or Service Edge.
	MaxUsage      pulumi.StringOutput    `pulumi:"maxUsage"`
	MicrotenantId pulumi.StringPtrOutput `pulumi:"microtenantId"`
	// Name of the provisioning key.
	Name     pulumi.StringOutput    `pulumi:"name"`
	UiConfig pulumi.StringPtrOutput `pulumi:"uiConfig"`
	// The provisioning key utilization count.
	UsageCount pulumi.StringOutput `pulumi:"usageCount"`
	// ID of the existing App Connector or Service Edge Group.
	ZcomponentId pulumi.StringOutput `pulumi:"zcomponentId"`
	// Read only property. Applicable only for GET calls, ignored in PUT/POST calls.
	ZcomponentName pulumi.StringOutput `pulumi:"zcomponentName"`
}

// NewProvisioningKey registers a new resource with the given unique name, arguments, and options.
func NewProvisioningKey(ctx *pulumi.Context,
	name string, args *ProvisioningKeyArgs, opts ...pulumi.ResourceOption) (*ProvisioningKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssociationType == nil {
		return nil, errors.New("invalid value for required argument 'AssociationType'")
	}
	if args.EnrollmentCertId == nil {
		return nil, errors.New("invalid value for required argument 'EnrollmentCertId'")
	}
	if args.MaxUsage == nil {
		return nil, errors.New("invalid value for required argument 'MaxUsage'")
	}
	if args.ZcomponentId == nil {
		return nil, errors.New("invalid value for required argument 'ZcomponentId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"ProvisioningKeyValue",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProvisioningKey
	err := ctx.RegisterResource("zpa:index/provisioningKey:ProvisioningKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProvisioningKey gets an existing ProvisioningKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProvisioningKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProvisioningKeyState, opts ...pulumi.ResourceOption) (*ProvisioningKey, error) {
	var resource ProvisioningKey
	err := ctx.ReadResource("zpa:index/provisioningKey:ProvisioningKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProvisioningKey resources.
type provisioningKeyState struct {
	// read only field. Ignored in PUT/POST calls.
	ProvisioningKeyValue *string `pulumi:"ProvisioningKeyValue"`
	AppConnectorGroupId  *string `pulumi:"appConnectorGroupId"`
	// Read only property. Applicable only for GET calls, ignored in PUT/POST calls.
	AppConnectorGroupName *string `pulumi:"appConnectorGroupName"`
	// Specifies the provisioning key type for App Connectors or ZPA Private Service Edges. The supported values are
	// CONNECTOR_GRP and SERVICE_EDGE_GRP.
	AssociationType *string `pulumi:"associationType"`
	// Whether the provisioning key is enabled or not. Supported values: true, false
	Enabled *bool `pulumi:"enabled"`
	// ID of the enrollment certificate that can be used for this provisioning key.
	EnrollmentCertId *string  `pulumi:"enrollmentCertId"`
	IpAcls           []string `pulumi:"ipAcls"`
	// The maximum number of instances where this provisioning key can be used for enrolling an App Connector or Service Edge.
	MaxUsage      *string `pulumi:"maxUsage"`
	MicrotenantId *string `pulumi:"microtenantId"`
	// Name of the provisioning key.
	Name     *string `pulumi:"name"`
	UiConfig *string `pulumi:"uiConfig"`
	// The provisioning key utilization count.
	UsageCount *string `pulumi:"usageCount"`
	// ID of the existing App Connector or Service Edge Group.
	ZcomponentId *string `pulumi:"zcomponentId"`
	// Read only property. Applicable only for GET calls, ignored in PUT/POST calls.
	ZcomponentName *string `pulumi:"zcomponentName"`
}

type ProvisioningKeyState struct {
	// read only field. Ignored in PUT/POST calls.
	ProvisioningKeyValue pulumi.StringPtrInput
	AppConnectorGroupId  pulumi.StringPtrInput
	// Read only property. Applicable only for GET calls, ignored in PUT/POST calls.
	AppConnectorGroupName pulumi.StringPtrInput
	// Specifies the provisioning key type for App Connectors or ZPA Private Service Edges. The supported values are
	// CONNECTOR_GRP and SERVICE_EDGE_GRP.
	AssociationType pulumi.StringPtrInput
	// Whether the provisioning key is enabled or not. Supported values: true, false
	Enabled pulumi.BoolPtrInput
	// ID of the enrollment certificate that can be used for this provisioning key.
	EnrollmentCertId pulumi.StringPtrInput
	IpAcls           pulumi.StringArrayInput
	// The maximum number of instances where this provisioning key can be used for enrolling an App Connector or Service Edge.
	MaxUsage      pulumi.StringPtrInput
	MicrotenantId pulumi.StringPtrInput
	// Name of the provisioning key.
	Name     pulumi.StringPtrInput
	UiConfig pulumi.StringPtrInput
	// The provisioning key utilization count.
	UsageCount pulumi.StringPtrInput
	// ID of the existing App Connector or Service Edge Group.
	ZcomponentId pulumi.StringPtrInput
	// Read only property. Applicable only for GET calls, ignored in PUT/POST calls.
	ZcomponentName pulumi.StringPtrInput
}

func (ProvisioningKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*provisioningKeyState)(nil)).Elem()
}

type provisioningKeyArgs struct {
	AppConnectorGroupId *string `pulumi:"appConnectorGroupId"`
	// Specifies the provisioning key type for App Connectors or ZPA Private Service Edges. The supported values are
	// CONNECTOR_GRP and SERVICE_EDGE_GRP.
	AssociationType string `pulumi:"associationType"`
	// Whether the provisioning key is enabled or not. Supported values: true, false
	Enabled *bool `pulumi:"enabled"`
	// ID of the enrollment certificate that can be used for this provisioning key.
	EnrollmentCertId string   `pulumi:"enrollmentCertId"`
	IpAcls           []string `pulumi:"ipAcls"`
	// The maximum number of instances where this provisioning key can be used for enrolling an App Connector or Service Edge.
	MaxUsage      string  `pulumi:"maxUsage"`
	MicrotenantId *string `pulumi:"microtenantId"`
	// Name of the provisioning key.
	Name     *string `pulumi:"name"`
	UiConfig *string `pulumi:"uiConfig"`
	// The provisioning key utilization count.
	UsageCount *string `pulumi:"usageCount"`
	// ID of the existing App Connector or Service Edge Group.
	ZcomponentId string `pulumi:"zcomponentId"`
	// Read only property. Applicable only for GET calls, ignored in PUT/POST calls.
	ZcomponentName *string `pulumi:"zcomponentName"`
}

// The set of arguments for constructing a ProvisioningKey resource.
type ProvisioningKeyArgs struct {
	AppConnectorGroupId pulumi.StringPtrInput
	// Specifies the provisioning key type for App Connectors or ZPA Private Service Edges. The supported values are
	// CONNECTOR_GRP and SERVICE_EDGE_GRP.
	AssociationType pulumi.StringInput
	// Whether the provisioning key is enabled or not. Supported values: true, false
	Enabled pulumi.BoolPtrInput
	// ID of the enrollment certificate that can be used for this provisioning key.
	EnrollmentCertId pulumi.StringInput
	IpAcls           pulumi.StringArrayInput
	// The maximum number of instances where this provisioning key can be used for enrolling an App Connector or Service Edge.
	MaxUsage      pulumi.StringInput
	MicrotenantId pulumi.StringPtrInput
	// Name of the provisioning key.
	Name     pulumi.StringPtrInput
	UiConfig pulumi.StringPtrInput
	// The provisioning key utilization count.
	UsageCount pulumi.StringPtrInput
	// ID of the existing App Connector or Service Edge Group.
	ZcomponentId pulumi.StringInput
	// Read only property. Applicable only for GET calls, ignored in PUT/POST calls.
	ZcomponentName pulumi.StringPtrInput
}

func (ProvisioningKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*provisioningKeyArgs)(nil)).Elem()
}

type ProvisioningKeyInput interface {
	pulumi.Input

	ToProvisioningKeyOutput() ProvisioningKeyOutput
	ToProvisioningKeyOutputWithContext(ctx context.Context) ProvisioningKeyOutput
}

func (*ProvisioningKey) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisioningKey)(nil)).Elem()
}

func (i *ProvisioningKey) ToProvisioningKeyOutput() ProvisioningKeyOutput {
	return i.ToProvisioningKeyOutputWithContext(context.Background())
}

func (i *ProvisioningKey) ToProvisioningKeyOutputWithContext(ctx context.Context) ProvisioningKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisioningKeyOutput)
}

// ProvisioningKeyArrayInput is an input type that accepts ProvisioningKeyArray and ProvisioningKeyArrayOutput values.
// You can construct a concrete instance of `ProvisioningKeyArrayInput` via:
//
//	ProvisioningKeyArray{ ProvisioningKeyArgs{...} }
type ProvisioningKeyArrayInput interface {
	pulumi.Input

	ToProvisioningKeyArrayOutput() ProvisioningKeyArrayOutput
	ToProvisioningKeyArrayOutputWithContext(context.Context) ProvisioningKeyArrayOutput
}

type ProvisioningKeyArray []ProvisioningKeyInput

func (ProvisioningKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProvisioningKey)(nil)).Elem()
}

func (i ProvisioningKeyArray) ToProvisioningKeyArrayOutput() ProvisioningKeyArrayOutput {
	return i.ToProvisioningKeyArrayOutputWithContext(context.Background())
}

func (i ProvisioningKeyArray) ToProvisioningKeyArrayOutputWithContext(ctx context.Context) ProvisioningKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisioningKeyArrayOutput)
}

// ProvisioningKeyMapInput is an input type that accepts ProvisioningKeyMap and ProvisioningKeyMapOutput values.
// You can construct a concrete instance of `ProvisioningKeyMapInput` via:
//
//	ProvisioningKeyMap{ "key": ProvisioningKeyArgs{...} }
type ProvisioningKeyMapInput interface {
	pulumi.Input

	ToProvisioningKeyMapOutput() ProvisioningKeyMapOutput
	ToProvisioningKeyMapOutputWithContext(context.Context) ProvisioningKeyMapOutput
}

type ProvisioningKeyMap map[string]ProvisioningKeyInput

func (ProvisioningKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProvisioningKey)(nil)).Elem()
}

func (i ProvisioningKeyMap) ToProvisioningKeyMapOutput() ProvisioningKeyMapOutput {
	return i.ToProvisioningKeyMapOutputWithContext(context.Background())
}

func (i ProvisioningKeyMap) ToProvisioningKeyMapOutputWithContext(ctx context.Context) ProvisioningKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisioningKeyMapOutput)
}

type ProvisioningKeyOutput struct{ *pulumi.OutputState }

func (ProvisioningKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisioningKey)(nil)).Elem()
}

func (o ProvisioningKeyOutput) ToProvisioningKeyOutput() ProvisioningKeyOutput {
	return o
}

func (o ProvisioningKeyOutput) ToProvisioningKeyOutputWithContext(ctx context.Context) ProvisioningKeyOutput {
	return o
}

// read only field. Ignored in PUT/POST calls.
func (o ProvisioningKeyOutput) ProvisioningKeyValue() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisioningKey) pulumi.StringOutput { return v.ProvisioningKeyValue }).(pulumi.StringOutput)
}

func (o ProvisioningKeyOutput) AppConnectorGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisioningKey) pulumi.StringPtrOutput { return v.AppConnectorGroupId }).(pulumi.StringPtrOutput)
}

// Read only property. Applicable only for GET calls, ignored in PUT/POST calls.
func (o ProvisioningKeyOutput) AppConnectorGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisioningKey) pulumi.StringOutput { return v.AppConnectorGroupName }).(pulumi.StringOutput)
}

// Specifies the provisioning key type for App Connectors or ZPA Private Service Edges. The supported values are
// CONNECTOR_GRP and SERVICE_EDGE_GRP.
func (o ProvisioningKeyOutput) AssociationType() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisioningKey) pulumi.StringOutput { return v.AssociationType }).(pulumi.StringOutput)
}

// Whether the provisioning key is enabled or not. Supported values: true, false
func (o ProvisioningKeyOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProvisioningKey) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// ID of the enrollment certificate that can be used for this provisioning key.
func (o ProvisioningKeyOutput) EnrollmentCertId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisioningKey) pulumi.StringOutput { return v.EnrollmentCertId }).(pulumi.StringOutput)
}

func (o ProvisioningKeyOutput) IpAcls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProvisioningKey) pulumi.StringArrayOutput { return v.IpAcls }).(pulumi.StringArrayOutput)
}

// The maximum number of instances where this provisioning key can be used for enrolling an App Connector or Service Edge.
func (o ProvisioningKeyOutput) MaxUsage() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisioningKey) pulumi.StringOutput { return v.MaxUsage }).(pulumi.StringOutput)
}

func (o ProvisioningKeyOutput) MicrotenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisioningKey) pulumi.StringPtrOutput { return v.MicrotenantId }).(pulumi.StringPtrOutput)
}

// Name of the provisioning key.
func (o ProvisioningKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisioningKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProvisioningKeyOutput) UiConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisioningKey) pulumi.StringPtrOutput { return v.UiConfig }).(pulumi.StringPtrOutput)
}

// The provisioning key utilization count.
func (o ProvisioningKeyOutput) UsageCount() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisioningKey) pulumi.StringOutput { return v.UsageCount }).(pulumi.StringOutput)
}

// ID of the existing App Connector or Service Edge Group.
func (o ProvisioningKeyOutput) ZcomponentId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisioningKey) pulumi.StringOutput { return v.ZcomponentId }).(pulumi.StringOutput)
}

// Read only property. Applicable only for GET calls, ignored in PUT/POST calls.
func (o ProvisioningKeyOutput) ZcomponentName() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisioningKey) pulumi.StringOutput { return v.ZcomponentName }).(pulumi.StringOutput)
}

type ProvisioningKeyArrayOutput struct{ *pulumi.OutputState }

func (ProvisioningKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProvisioningKey)(nil)).Elem()
}

func (o ProvisioningKeyArrayOutput) ToProvisioningKeyArrayOutput() ProvisioningKeyArrayOutput {
	return o
}

func (o ProvisioningKeyArrayOutput) ToProvisioningKeyArrayOutputWithContext(ctx context.Context) ProvisioningKeyArrayOutput {
	return o
}

func (o ProvisioningKeyArrayOutput) Index(i pulumi.IntInput) ProvisioningKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProvisioningKey {
		return vs[0].([]*ProvisioningKey)[vs[1].(int)]
	}).(ProvisioningKeyOutput)
}

type ProvisioningKeyMapOutput struct{ *pulumi.OutputState }

func (ProvisioningKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProvisioningKey)(nil)).Elem()
}

func (o ProvisioningKeyMapOutput) ToProvisioningKeyMapOutput() ProvisioningKeyMapOutput {
	return o
}

func (o ProvisioningKeyMapOutput) ToProvisioningKeyMapOutputWithContext(ctx context.Context) ProvisioningKeyMapOutput {
	return o
}

func (o ProvisioningKeyMapOutput) MapIndex(k pulumi.StringInput) ProvisioningKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProvisioningKey {
		return vs[0].(map[string]*ProvisioningKey)[vs[1].(string)]
	}).(ProvisioningKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProvisioningKeyInput)(nil)).Elem(), &ProvisioningKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProvisioningKeyArrayInput)(nil)).Elem(), ProvisioningKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProvisioningKeyMapInput)(nil)).Elem(), ProvisioningKeyMap{})
	pulumi.RegisterOutputType(ProvisioningKeyOutput{})
	pulumi.RegisterOutputType(ProvisioningKeyArrayOutput{})
	pulumi.RegisterOutputType(ProvisioningKeyMapOutput{})
}
