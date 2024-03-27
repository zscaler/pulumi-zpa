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

type InspectionCustomControls struct {
	pulumi.CustomResourceState

	// The performed action
	Action      pulumi.StringOutput `pulumi:"action"`
	ActionValue pulumi.StringOutput `pulumi:"actionValue"`
	// Name of the inspection profile
	AssociatedInspectionProfileNames InspectionCustomControlsAssociatedInspectionProfileNameArrayOutput `pulumi:"associatedInspectionProfileNames"`
	ControlNumber                    pulumi.StringOutput                                                `pulumi:"controlNumber"`
	// The control rule in JSON format that has the conditions and type of control for the inspection control
	ControlRuleJson pulumi.StringOutput `pulumi:"controlRuleJson"`
	ControlType     pulumi.StringOutput `pulumi:"controlType"`
	// The performed action
	DefaultAction pulumi.StringOutput `pulumi:"defaultAction"`
	// This is used to provide the redirect URL if the default action is set to REDIRECT
	DefaultActionValue pulumi.StringOutput `pulumi:"defaultActionValue"`
	// Description of the custom control
	Description pulumi.StringOutput `pulumi:"description"`
	Name        pulumi.StringOutput `pulumi:"name"`
	// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
	ParanoiaLevel pulumi.StringOutput `pulumi:"paranoiaLevel"`
	ProtocolType  pulumi.StringOutput `pulumi:"protocolType"`
	// Rules of the custom controls applied as conditions (JSON)
	Rules InspectionCustomControlsRuleArrayOutput `pulumi:"rules"`
	// Severity of the control number
	Severity pulumi.StringOutput `pulumi:"severity"`
	// Rules to be applied to the request or response type
	Type    pulumi.StringOutput `pulumi:"type"`
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewInspectionCustomControls registers a new resource with the given unique name, arguments, and options.
func NewInspectionCustomControls(ctx *pulumi.Context,
	name string, args *InspectionCustomControlsArgs, opts ...pulumi.ResourceOption) (*InspectionCustomControls, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultAction == nil {
		return nil, errors.New("invalid value for required argument 'DefaultAction'")
	}
	if args.Severity == nil {
		return nil, errors.New("invalid value for required argument 'Severity'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InspectionCustomControls
	err := ctx.RegisterResource("zpa:index/inspectionCustomControls:InspectionCustomControls", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInspectionCustomControls gets an existing InspectionCustomControls resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInspectionCustomControls(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InspectionCustomControlsState, opts ...pulumi.ResourceOption) (*InspectionCustomControls, error) {
	var resource InspectionCustomControls
	err := ctx.ReadResource("zpa:index/inspectionCustomControls:InspectionCustomControls", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InspectionCustomControls resources.
type inspectionCustomControlsState struct {
	// The performed action
	Action      *string `pulumi:"action"`
	ActionValue *string `pulumi:"actionValue"`
	// Name of the inspection profile
	AssociatedInspectionProfileNames []InspectionCustomControlsAssociatedInspectionProfileName `pulumi:"associatedInspectionProfileNames"`
	ControlNumber                    *string                                                   `pulumi:"controlNumber"`
	// The control rule in JSON format that has the conditions and type of control for the inspection control
	ControlRuleJson *string `pulumi:"controlRuleJson"`
	ControlType     *string `pulumi:"controlType"`
	// The performed action
	DefaultAction *string `pulumi:"defaultAction"`
	// This is used to provide the redirect URL if the default action is set to REDIRECT
	DefaultActionValue *string `pulumi:"defaultActionValue"`
	// Description of the custom control
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
	// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
	ParanoiaLevel *string `pulumi:"paranoiaLevel"`
	ProtocolType  *string `pulumi:"protocolType"`
	// Rules of the custom controls applied as conditions (JSON)
	Rules []InspectionCustomControlsRule `pulumi:"rules"`
	// Severity of the control number
	Severity *string `pulumi:"severity"`
	// Rules to be applied to the request or response type
	Type    *string `pulumi:"type"`
	Version *string `pulumi:"version"`
}

type InspectionCustomControlsState struct {
	// The performed action
	Action      pulumi.StringPtrInput
	ActionValue pulumi.StringPtrInput
	// Name of the inspection profile
	AssociatedInspectionProfileNames InspectionCustomControlsAssociatedInspectionProfileNameArrayInput
	ControlNumber                    pulumi.StringPtrInput
	// The control rule in JSON format that has the conditions and type of control for the inspection control
	ControlRuleJson pulumi.StringPtrInput
	ControlType     pulumi.StringPtrInput
	// The performed action
	DefaultAction pulumi.StringPtrInput
	// This is used to provide the redirect URL if the default action is set to REDIRECT
	DefaultActionValue pulumi.StringPtrInput
	// Description of the custom control
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
	ParanoiaLevel pulumi.StringPtrInput
	ProtocolType  pulumi.StringPtrInput
	// Rules of the custom controls applied as conditions (JSON)
	Rules InspectionCustomControlsRuleArrayInput
	// Severity of the control number
	Severity pulumi.StringPtrInput
	// Rules to be applied to the request or response type
	Type    pulumi.StringPtrInput
	Version pulumi.StringPtrInput
}

func (InspectionCustomControlsState) ElementType() reflect.Type {
	return reflect.TypeOf((*inspectionCustomControlsState)(nil)).Elem()
}

type inspectionCustomControlsArgs struct {
	// The performed action
	Action      *string `pulumi:"action"`
	ActionValue *string `pulumi:"actionValue"`
	// Name of the inspection profile
	AssociatedInspectionProfileNames []InspectionCustomControlsAssociatedInspectionProfileName `pulumi:"associatedInspectionProfileNames"`
	ControlNumber                    *string                                                   `pulumi:"controlNumber"`
	// The control rule in JSON format that has the conditions and type of control for the inspection control
	ControlRuleJson *string `pulumi:"controlRuleJson"`
	ControlType     *string `pulumi:"controlType"`
	// The performed action
	DefaultAction string `pulumi:"defaultAction"`
	// This is used to provide the redirect URL if the default action is set to REDIRECT
	DefaultActionValue *string `pulumi:"defaultActionValue"`
	// Description of the custom control
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
	// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
	ParanoiaLevel *string `pulumi:"paranoiaLevel"`
	ProtocolType  *string `pulumi:"protocolType"`
	// Rules of the custom controls applied as conditions (JSON)
	Rules []InspectionCustomControlsRule `pulumi:"rules"`
	// Severity of the control number
	Severity string `pulumi:"severity"`
	// Rules to be applied to the request or response type
	Type    string  `pulumi:"type"`
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a InspectionCustomControls resource.
type InspectionCustomControlsArgs struct {
	// The performed action
	Action      pulumi.StringPtrInput
	ActionValue pulumi.StringPtrInput
	// Name of the inspection profile
	AssociatedInspectionProfileNames InspectionCustomControlsAssociatedInspectionProfileNameArrayInput
	ControlNumber                    pulumi.StringPtrInput
	// The control rule in JSON format that has the conditions and type of control for the inspection control
	ControlRuleJson pulumi.StringPtrInput
	ControlType     pulumi.StringPtrInput
	// The performed action
	DefaultAction pulumi.StringInput
	// This is used to provide the redirect URL if the default action is set to REDIRECT
	DefaultActionValue pulumi.StringPtrInput
	// Description of the custom control
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
	ParanoiaLevel pulumi.StringPtrInput
	ProtocolType  pulumi.StringPtrInput
	// Rules of the custom controls applied as conditions (JSON)
	Rules InspectionCustomControlsRuleArrayInput
	// Severity of the control number
	Severity pulumi.StringInput
	// Rules to be applied to the request or response type
	Type    pulumi.StringInput
	Version pulumi.StringPtrInput
}

func (InspectionCustomControlsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inspectionCustomControlsArgs)(nil)).Elem()
}

type InspectionCustomControlsInput interface {
	pulumi.Input

	ToInspectionCustomControlsOutput() InspectionCustomControlsOutput
	ToInspectionCustomControlsOutputWithContext(ctx context.Context) InspectionCustomControlsOutput
}

func (*InspectionCustomControls) ElementType() reflect.Type {
	return reflect.TypeOf((**InspectionCustomControls)(nil)).Elem()
}

func (i *InspectionCustomControls) ToInspectionCustomControlsOutput() InspectionCustomControlsOutput {
	return i.ToInspectionCustomControlsOutputWithContext(context.Background())
}

func (i *InspectionCustomControls) ToInspectionCustomControlsOutputWithContext(ctx context.Context) InspectionCustomControlsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InspectionCustomControlsOutput)
}

// InspectionCustomControlsArrayInput is an input type that accepts InspectionCustomControlsArray and InspectionCustomControlsArrayOutput values.
// You can construct a concrete instance of `InspectionCustomControlsArrayInput` via:
//
//	InspectionCustomControlsArray{ InspectionCustomControlsArgs{...} }
type InspectionCustomControlsArrayInput interface {
	pulumi.Input

	ToInspectionCustomControlsArrayOutput() InspectionCustomControlsArrayOutput
	ToInspectionCustomControlsArrayOutputWithContext(context.Context) InspectionCustomControlsArrayOutput
}

type InspectionCustomControlsArray []InspectionCustomControlsInput

func (InspectionCustomControlsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InspectionCustomControls)(nil)).Elem()
}

func (i InspectionCustomControlsArray) ToInspectionCustomControlsArrayOutput() InspectionCustomControlsArrayOutput {
	return i.ToInspectionCustomControlsArrayOutputWithContext(context.Background())
}

func (i InspectionCustomControlsArray) ToInspectionCustomControlsArrayOutputWithContext(ctx context.Context) InspectionCustomControlsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InspectionCustomControlsArrayOutput)
}

// InspectionCustomControlsMapInput is an input type that accepts InspectionCustomControlsMap and InspectionCustomControlsMapOutput values.
// You can construct a concrete instance of `InspectionCustomControlsMapInput` via:
//
//	InspectionCustomControlsMap{ "key": InspectionCustomControlsArgs{...} }
type InspectionCustomControlsMapInput interface {
	pulumi.Input

	ToInspectionCustomControlsMapOutput() InspectionCustomControlsMapOutput
	ToInspectionCustomControlsMapOutputWithContext(context.Context) InspectionCustomControlsMapOutput
}

type InspectionCustomControlsMap map[string]InspectionCustomControlsInput

func (InspectionCustomControlsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InspectionCustomControls)(nil)).Elem()
}

func (i InspectionCustomControlsMap) ToInspectionCustomControlsMapOutput() InspectionCustomControlsMapOutput {
	return i.ToInspectionCustomControlsMapOutputWithContext(context.Background())
}

func (i InspectionCustomControlsMap) ToInspectionCustomControlsMapOutputWithContext(ctx context.Context) InspectionCustomControlsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InspectionCustomControlsMapOutput)
}

type InspectionCustomControlsOutput struct{ *pulumi.OutputState }

func (InspectionCustomControlsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InspectionCustomControls)(nil)).Elem()
}

func (o InspectionCustomControlsOutput) ToInspectionCustomControlsOutput() InspectionCustomControlsOutput {
	return o
}

func (o InspectionCustomControlsOutput) ToInspectionCustomControlsOutputWithContext(ctx context.Context) InspectionCustomControlsOutput {
	return o
}

// The performed action
func (o InspectionCustomControlsOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *InspectionCustomControls) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

func (o InspectionCustomControlsOutput) ActionValue() pulumi.StringOutput {
	return o.ApplyT(func(v *InspectionCustomControls) pulumi.StringOutput { return v.ActionValue }).(pulumi.StringOutput)
}

// Name of the inspection profile
func (o InspectionCustomControlsOutput) AssociatedInspectionProfileNames() InspectionCustomControlsAssociatedInspectionProfileNameArrayOutput {
	return o.ApplyT(func(v *InspectionCustomControls) InspectionCustomControlsAssociatedInspectionProfileNameArrayOutput {
		return v.AssociatedInspectionProfileNames
	}).(InspectionCustomControlsAssociatedInspectionProfileNameArrayOutput)
}

func (o InspectionCustomControlsOutput) ControlNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *InspectionCustomControls) pulumi.StringOutput { return v.ControlNumber }).(pulumi.StringOutput)
}

// The control rule in JSON format that has the conditions and type of control for the inspection control
func (o InspectionCustomControlsOutput) ControlRuleJson() pulumi.StringOutput {
	return o.ApplyT(func(v *InspectionCustomControls) pulumi.StringOutput { return v.ControlRuleJson }).(pulumi.StringOutput)
}

func (o InspectionCustomControlsOutput) ControlType() pulumi.StringOutput {
	return o.ApplyT(func(v *InspectionCustomControls) pulumi.StringOutput { return v.ControlType }).(pulumi.StringOutput)
}

// The performed action
func (o InspectionCustomControlsOutput) DefaultAction() pulumi.StringOutput {
	return o.ApplyT(func(v *InspectionCustomControls) pulumi.StringOutput { return v.DefaultAction }).(pulumi.StringOutput)
}

// This is used to provide the redirect URL if the default action is set to REDIRECT
func (o InspectionCustomControlsOutput) DefaultActionValue() pulumi.StringOutput {
	return o.ApplyT(func(v *InspectionCustomControls) pulumi.StringOutput { return v.DefaultActionValue }).(pulumi.StringOutput)
}

// Description of the custom control
func (o InspectionCustomControlsOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *InspectionCustomControls) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o InspectionCustomControlsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InspectionCustomControls) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
func (o InspectionCustomControlsOutput) ParanoiaLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *InspectionCustomControls) pulumi.StringOutput { return v.ParanoiaLevel }).(pulumi.StringOutput)
}

func (o InspectionCustomControlsOutput) ProtocolType() pulumi.StringOutput {
	return o.ApplyT(func(v *InspectionCustomControls) pulumi.StringOutput { return v.ProtocolType }).(pulumi.StringOutput)
}

// Rules of the custom controls applied as conditions (JSON)
func (o InspectionCustomControlsOutput) Rules() InspectionCustomControlsRuleArrayOutput {
	return o.ApplyT(func(v *InspectionCustomControls) InspectionCustomControlsRuleArrayOutput { return v.Rules }).(InspectionCustomControlsRuleArrayOutput)
}

// Severity of the control number
func (o InspectionCustomControlsOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *InspectionCustomControls) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

// Rules to be applied to the request or response type
func (o InspectionCustomControlsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InspectionCustomControls) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o InspectionCustomControlsOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *InspectionCustomControls) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type InspectionCustomControlsArrayOutput struct{ *pulumi.OutputState }

func (InspectionCustomControlsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InspectionCustomControls)(nil)).Elem()
}

func (o InspectionCustomControlsArrayOutput) ToInspectionCustomControlsArrayOutput() InspectionCustomControlsArrayOutput {
	return o
}

func (o InspectionCustomControlsArrayOutput) ToInspectionCustomControlsArrayOutputWithContext(ctx context.Context) InspectionCustomControlsArrayOutput {
	return o
}

func (o InspectionCustomControlsArrayOutput) Index(i pulumi.IntInput) InspectionCustomControlsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InspectionCustomControls {
		return vs[0].([]*InspectionCustomControls)[vs[1].(int)]
	}).(InspectionCustomControlsOutput)
}

type InspectionCustomControlsMapOutput struct{ *pulumi.OutputState }

func (InspectionCustomControlsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InspectionCustomControls)(nil)).Elem()
}

func (o InspectionCustomControlsMapOutput) ToInspectionCustomControlsMapOutput() InspectionCustomControlsMapOutput {
	return o
}

func (o InspectionCustomControlsMapOutput) ToInspectionCustomControlsMapOutputWithContext(ctx context.Context) InspectionCustomControlsMapOutput {
	return o
}

func (o InspectionCustomControlsMapOutput) MapIndex(k pulumi.StringInput) InspectionCustomControlsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InspectionCustomControls {
		return vs[0].(map[string]*InspectionCustomControls)[vs[1].(string)]
	}).(InspectionCustomControlsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InspectionCustomControlsInput)(nil)).Elem(), &InspectionCustomControls{})
	pulumi.RegisterInputType(reflect.TypeOf((*InspectionCustomControlsArrayInput)(nil)).Elem(), InspectionCustomControlsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InspectionCustomControlsMapInput)(nil)).Elem(), InspectionCustomControlsMap{})
	pulumi.RegisterOutputType(InspectionCustomControlsOutput{})
	pulumi.RegisterOutputType(InspectionCustomControlsArrayOutput{})
	pulumi.RegisterOutputType(InspectionCustomControlsMapOutput{})
}