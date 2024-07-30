// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// * [Official documentation](https://help.zscaler.com/zpa/about-privileged-portals)
// * [API documentation](https://help.zscaler.com/zpa/configuring-privileged-portals-using-api)
//
// The **zpa_pra_portal_controller** resource creates a privileged remote access portal in the Zscaler Private Access cloud. This resource can then be referenced in an privileged remote access console resource.
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
//			thisBaCertificate, err := zpa.GetBaCertificate(ctx, &zpa.GetBaCertificateArgs{
//				Name: pulumi.StringRef("portal.acme.com"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.NewPRAPortal(ctx, "thisPRAPortal", &zpa.PRAPortalArgs{
//				Description:             pulumi.String("portal.acme.com"),
//				Enabled:                 pulumi.Bool(true),
//				Domain:                  pulumi.String("portal.acme.com"),
//				CertificateId:           pulumi.String(thisBaCertificate.Id),
//				UserNotification:        pulumi.String("Created with Terraform"),
//				UserNotificationEnabled: pulumi.Bool(true),
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
// **pra_portal_controller** can be imported by using `<PORTAL ID>` or `<PORTAL NAME>` as the import ID.
//
// For example:
//
// ```sh
// $ pulumi import zpa:index/praPortalController:PraPortalController this <portal_id>
// ```
//
// or
//
// ```sh
// $ pulumi import zpa:index/praPortalController:PraPortalController this <portal_name>
// ```
//
// Deprecated: zpa.index/praportalcontroller.PraPortalController has been deprecated in favor of zpa.index/praportal.PRAPortal
type PraPortalController struct {
	pulumi.CustomResourceState

	// The unique identifier of the certificate
	CertificateId pulumi.StringPtrOutput `pulumi:"certificateId"`
	// The description of the privileged portal
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The domain of the privileged portal
	Domain pulumi.StringPtrOutput `pulumi:"domain"`
	// Whether or not the privileged portal is enabled
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
	// microtenantId as 0 when making requests to retrieve data from the Default Microtenant. Pass microtenantId as null to
	// retrieve data from all customers associated with the tenant.
	MicrotenantId pulumi.StringOutput `pulumi:"microtenantId"`
	// The name of the privileged portal
	Name pulumi.StringOutput `pulumi:"name"`
	// The notification message displayed in the banner of the privileged portallink, if enabled
	UserNotification pulumi.StringPtrOutput `pulumi:"userNotification"`
	// Indicates if the Notification Banner is enabled (true) or disabled (false)
	UserNotificationEnabled pulumi.BoolPtrOutput `pulumi:"userNotificationEnabled"`
}

// NewPraPortalController registers a new resource with the given unique name, arguments, and options.
func NewPraPortalController(ctx *pulumi.Context,
	name string, args *PraPortalControllerArgs, opts ...pulumi.ResourceOption) (*PraPortalController, error) {
	if args == nil {
		args = &PraPortalControllerArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PraPortalController
	err := ctx.RegisterResource("zpa:index/praPortalController:PraPortalController", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPraPortalController gets an existing PraPortalController resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPraPortalController(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PraPortalControllerState, opts ...pulumi.ResourceOption) (*PraPortalController, error) {
	var resource PraPortalController
	err := ctx.ReadResource("zpa:index/praPortalController:PraPortalController", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PraPortalController resources.
type praPortalControllerState struct {
	// The unique identifier of the certificate
	CertificateId *string `pulumi:"certificateId"`
	// The description of the privileged portal
	Description *string `pulumi:"description"`
	// The domain of the privileged portal
	Domain *string `pulumi:"domain"`
	// Whether or not the privileged portal is enabled
	Enabled *bool `pulumi:"enabled"`
	// The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
	// microtenantId as 0 when making requests to retrieve data from the Default Microtenant. Pass microtenantId as null to
	// retrieve data from all customers associated with the tenant.
	MicrotenantId *string `pulumi:"microtenantId"`
	// The name of the privileged portal
	Name *string `pulumi:"name"`
	// The notification message displayed in the banner of the privileged portallink, if enabled
	UserNotification *string `pulumi:"userNotification"`
	// Indicates if the Notification Banner is enabled (true) or disabled (false)
	UserNotificationEnabled *bool `pulumi:"userNotificationEnabled"`
}

type PraPortalControllerState struct {
	// The unique identifier of the certificate
	CertificateId pulumi.StringPtrInput
	// The description of the privileged portal
	Description pulumi.StringPtrInput
	// The domain of the privileged portal
	Domain pulumi.StringPtrInput
	// Whether or not the privileged portal is enabled
	Enabled pulumi.BoolPtrInput
	// The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
	// microtenantId as 0 when making requests to retrieve data from the Default Microtenant. Pass microtenantId as null to
	// retrieve data from all customers associated with the tenant.
	MicrotenantId pulumi.StringPtrInput
	// The name of the privileged portal
	Name pulumi.StringPtrInput
	// The notification message displayed in the banner of the privileged portallink, if enabled
	UserNotification pulumi.StringPtrInput
	// Indicates if the Notification Banner is enabled (true) or disabled (false)
	UserNotificationEnabled pulumi.BoolPtrInput
}

func (PraPortalControllerState) ElementType() reflect.Type {
	return reflect.TypeOf((*praPortalControllerState)(nil)).Elem()
}

type praPortalControllerArgs struct {
	// The unique identifier of the certificate
	CertificateId *string `pulumi:"certificateId"`
	// The description of the privileged portal
	Description *string `pulumi:"description"`
	// The domain of the privileged portal
	Domain *string `pulumi:"domain"`
	// Whether or not the privileged portal is enabled
	Enabled *bool `pulumi:"enabled"`
	// The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
	// microtenantId as 0 when making requests to retrieve data from the Default Microtenant. Pass microtenantId as null to
	// retrieve data from all customers associated with the tenant.
	MicrotenantId *string `pulumi:"microtenantId"`
	// The name of the privileged portal
	Name *string `pulumi:"name"`
	// The notification message displayed in the banner of the privileged portallink, if enabled
	UserNotification *string `pulumi:"userNotification"`
	// Indicates if the Notification Banner is enabled (true) or disabled (false)
	UserNotificationEnabled *bool `pulumi:"userNotificationEnabled"`
}

// The set of arguments for constructing a PraPortalController resource.
type PraPortalControllerArgs struct {
	// The unique identifier of the certificate
	CertificateId pulumi.StringPtrInput
	// The description of the privileged portal
	Description pulumi.StringPtrInput
	// The domain of the privileged portal
	Domain pulumi.StringPtrInput
	// Whether or not the privileged portal is enabled
	Enabled pulumi.BoolPtrInput
	// The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
	// microtenantId as 0 when making requests to retrieve data from the Default Microtenant. Pass microtenantId as null to
	// retrieve data from all customers associated with the tenant.
	MicrotenantId pulumi.StringPtrInput
	// The name of the privileged portal
	Name pulumi.StringPtrInput
	// The notification message displayed in the banner of the privileged portallink, if enabled
	UserNotification pulumi.StringPtrInput
	// Indicates if the Notification Banner is enabled (true) or disabled (false)
	UserNotificationEnabled pulumi.BoolPtrInput
}

func (PraPortalControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*praPortalControllerArgs)(nil)).Elem()
}

type PraPortalControllerInput interface {
	pulumi.Input

	ToPraPortalControllerOutput() PraPortalControllerOutput
	ToPraPortalControllerOutputWithContext(ctx context.Context) PraPortalControllerOutput
}

func (*PraPortalController) ElementType() reflect.Type {
	return reflect.TypeOf((**PraPortalController)(nil)).Elem()
}

func (i *PraPortalController) ToPraPortalControllerOutput() PraPortalControllerOutput {
	return i.ToPraPortalControllerOutputWithContext(context.Background())
}

func (i *PraPortalController) ToPraPortalControllerOutputWithContext(ctx context.Context) PraPortalControllerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PraPortalControllerOutput)
}

// PraPortalControllerArrayInput is an input type that accepts PraPortalControllerArray and PraPortalControllerArrayOutput values.
// You can construct a concrete instance of `PraPortalControllerArrayInput` via:
//
//	PraPortalControllerArray{ PraPortalControllerArgs{...} }
type PraPortalControllerArrayInput interface {
	pulumi.Input

	ToPraPortalControllerArrayOutput() PraPortalControllerArrayOutput
	ToPraPortalControllerArrayOutputWithContext(context.Context) PraPortalControllerArrayOutput
}

type PraPortalControllerArray []PraPortalControllerInput

func (PraPortalControllerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PraPortalController)(nil)).Elem()
}

func (i PraPortalControllerArray) ToPraPortalControllerArrayOutput() PraPortalControllerArrayOutput {
	return i.ToPraPortalControllerArrayOutputWithContext(context.Background())
}

func (i PraPortalControllerArray) ToPraPortalControllerArrayOutputWithContext(ctx context.Context) PraPortalControllerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PraPortalControllerArrayOutput)
}

// PraPortalControllerMapInput is an input type that accepts PraPortalControllerMap and PraPortalControllerMapOutput values.
// You can construct a concrete instance of `PraPortalControllerMapInput` via:
//
//	PraPortalControllerMap{ "key": PraPortalControllerArgs{...} }
type PraPortalControllerMapInput interface {
	pulumi.Input

	ToPraPortalControllerMapOutput() PraPortalControllerMapOutput
	ToPraPortalControllerMapOutputWithContext(context.Context) PraPortalControllerMapOutput
}

type PraPortalControllerMap map[string]PraPortalControllerInput

func (PraPortalControllerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PraPortalController)(nil)).Elem()
}

func (i PraPortalControllerMap) ToPraPortalControllerMapOutput() PraPortalControllerMapOutput {
	return i.ToPraPortalControllerMapOutputWithContext(context.Background())
}

func (i PraPortalControllerMap) ToPraPortalControllerMapOutputWithContext(ctx context.Context) PraPortalControllerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PraPortalControllerMapOutput)
}

type PraPortalControllerOutput struct{ *pulumi.OutputState }

func (PraPortalControllerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PraPortalController)(nil)).Elem()
}

func (o PraPortalControllerOutput) ToPraPortalControllerOutput() PraPortalControllerOutput {
	return o
}

func (o PraPortalControllerOutput) ToPraPortalControllerOutputWithContext(ctx context.Context) PraPortalControllerOutput {
	return o
}

// The unique identifier of the certificate
func (o PraPortalControllerOutput) CertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PraPortalController) pulumi.StringPtrOutput { return v.CertificateId }).(pulumi.StringPtrOutput)
}

// The description of the privileged portal
func (o PraPortalControllerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PraPortalController) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The domain of the privileged portal
func (o PraPortalControllerOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PraPortalController) pulumi.StringPtrOutput { return v.Domain }).(pulumi.StringPtrOutput)
}

// Whether or not the privileged portal is enabled
func (o PraPortalControllerOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PraPortalController) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
// microtenantId as 0 when making requests to retrieve data from the Default Microtenant. Pass microtenantId as null to
// retrieve data from all customers associated with the tenant.
func (o PraPortalControllerOutput) MicrotenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *PraPortalController) pulumi.StringOutput { return v.MicrotenantId }).(pulumi.StringOutput)
}

// The name of the privileged portal
func (o PraPortalControllerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PraPortalController) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The notification message displayed in the banner of the privileged portallink, if enabled
func (o PraPortalControllerOutput) UserNotification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PraPortalController) pulumi.StringPtrOutput { return v.UserNotification }).(pulumi.StringPtrOutput)
}

// Indicates if the Notification Banner is enabled (true) or disabled (false)
func (o PraPortalControllerOutput) UserNotificationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PraPortalController) pulumi.BoolPtrOutput { return v.UserNotificationEnabled }).(pulumi.BoolPtrOutput)
}

type PraPortalControllerArrayOutput struct{ *pulumi.OutputState }

func (PraPortalControllerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PraPortalController)(nil)).Elem()
}

func (o PraPortalControllerArrayOutput) ToPraPortalControllerArrayOutput() PraPortalControllerArrayOutput {
	return o
}

func (o PraPortalControllerArrayOutput) ToPraPortalControllerArrayOutputWithContext(ctx context.Context) PraPortalControllerArrayOutput {
	return o
}

func (o PraPortalControllerArrayOutput) Index(i pulumi.IntInput) PraPortalControllerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PraPortalController {
		return vs[0].([]*PraPortalController)[vs[1].(int)]
	}).(PraPortalControllerOutput)
}

type PraPortalControllerMapOutput struct{ *pulumi.OutputState }

func (PraPortalControllerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PraPortalController)(nil)).Elem()
}

func (o PraPortalControllerMapOutput) ToPraPortalControllerMapOutput() PraPortalControllerMapOutput {
	return o
}

func (o PraPortalControllerMapOutput) ToPraPortalControllerMapOutputWithContext(ctx context.Context) PraPortalControllerMapOutput {
	return o
}

func (o PraPortalControllerMapOutput) MapIndex(k pulumi.StringInput) PraPortalControllerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PraPortalController {
		return vs[0].(map[string]*PraPortalController)[vs[1].(string)]
	}).(PraPortalControllerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PraPortalControllerInput)(nil)).Elem(), &PraPortalController{})
	pulumi.RegisterInputType(reflect.TypeOf((*PraPortalControllerArrayInput)(nil)).Elem(), PraPortalControllerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PraPortalControllerMapInput)(nil)).Elem(), PraPortalControllerMap{})
	pulumi.RegisterOutputType(PraPortalControllerOutput{})
	pulumi.RegisterOutputType(PraPortalControllerArrayOutput{})
	pulumi.RegisterOutputType(PraPortalControllerMapOutput{})
}
