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
// <!--Start PulumiCodeChooser -->
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
// <!--End PulumiCodeChooser -->
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
// $ pulumi import zpa:index/pRAPortal:PRAPortal this <portal_id>
// ```
//
// or
//
// ```sh
// $ pulumi import zpa:index/pRAPortal:PRAPortal this <portal_name>
// ```
type PRAPortal struct {
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

// NewPRAPortal registers a new resource with the given unique name, arguments, and options.
func NewPRAPortal(ctx *pulumi.Context,
	name string, args *PRAPortalArgs, opts ...pulumi.ResourceOption) (*PRAPortal, error) {
	if args == nil {
		args = &PRAPortalArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("zpa:index/praPortalController:PraPortalController"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PRAPortal
	err := ctx.RegisterResource("zpa:index/pRAPortal:PRAPortal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPRAPortal gets an existing PRAPortal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPRAPortal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PRAPortalState, opts ...pulumi.ResourceOption) (*PRAPortal, error) {
	var resource PRAPortal
	err := ctx.ReadResource("zpa:index/pRAPortal:PRAPortal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PRAPortal resources.
type praportalState struct {
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

type PRAPortalState struct {
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

func (PRAPortalState) ElementType() reflect.Type {
	return reflect.TypeOf((*praportalState)(nil)).Elem()
}

type praportalArgs struct {
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

// The set of arguments for constructing a PRAPortal resource.
type PRAPortalArgs struct {
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

func (PRAPortalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*praportalArgs)(nil)).Elem()
}

type PRAPortalInput interface {
	pulumi.Input

	ToPRAPortalOutput() PRAPortalOutput
	ToPRAPortalOutputWithContext(ctx context.Context) PRAPortalOutput
}

func (*PRAPortal) ElementType() reflect.Type {
	return reflect.TypeOf((**PRAPortal)(nil)).Elem()
}

func (i *PRAPortal) ToPRAPortalOutput() PRAPortalOutput {
	return i.ToPRAPortalOutputWithContext(context.Background())
}

func (i *PRAPortal) ToPRAPortalOutputWithContext(ctx context.Context) PRAPortalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PRAPortalOutput)
}

// PRAPortalArrayInput is an input type that accepts PRAPortalArray and PRAPortalArrayOutput values.
// You can construct a concrete instance of `PRAPortalArrayInput` via:
//
//	PRAPortalArray{ PRAPortalArgs{...} }
type PRAPortalArrayInput interface {
	pulumi.Input

	ToPRAPortalArrayOutput() PRAPortalArrayOutput
	ToPRAPortalArrayOutputWithContext(context.Context) PRAPortalArrayOutput
}

type PRAPortalArray []PRAPortalInput

func (PRAPortalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PRAPortal)(nil)).Elem()
}

func (i PRAPortalArray) ToPRAPortalArrayOutput() PRAPortalArrayOutput {
	return i.ToPRAPortalArrayOutputWithContext(context.Background())
}

func (i PRAPortalArray) ToPRAPortalArrayOutputWithContext(ctx context.Context) PRAPortalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PRAPortalArrayOutput)
}

// PRAPortalMapInput is an input type that accepts PRAPortalMap and PRAPortalMapOutput values.
// You can construct a concrete instance of `PRAPortalMapInput` via:
//
//	PRAPortalMap{ "key": PRAPortalArgs{...} }
type PRAPortalMapInput interface {
	pulumi.Input

	ToPRAPortalMapOutput() PRAPortalMapOutput
	ToPRAPortalMapOutputWithContext(context.Context) PRAPortalMapOutput
}

type PRAPortalMap map[string]PRAPortalInput

func (PRAPortalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PRAPortal)(nil)).Elem()
}

func (i PRAPortalMap) ToPRAPortalMapOutput() PRAPortalMapOutput {
	return i.ToPRAPortalMapOutputWithContext(context.Background())
}

func (i PRAPortalMap) ToPRAPortalMapOutputWithContext(ctx context.Context) PRAPortalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PRAPortalMapOutput)
}

type PRAPortalOutput struct{ *pulumi.OutputState }

func (PRAPortalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PRAPortal)(nil)).Elem()
}

func (o PRAPortalOutput) ToPRAPortalOutput() PRAPortalOutput {
	return o
}

func (o PRAPortalOutput) ToPRAPortalOutputWithContext(ctx context.Context) PRAPortalOutput {
	return o
}

// The unique identifier of the certificate
func (o PRAPortalOutput) CertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PRAPortal) pulumi.StringPtrOutput { return v.CertificateId }).(pulumi.StringPtrOutput)
}

// The description of the privileged portal
func (o PRAPortalOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PRAPortal) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The domain of the privileged portal
func (o PRAPortalOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PRAPortal) pulumi.StringPtrOutput { return v.Domain }).(pulumi.StringPtrOutput)
}

// Whether or not the privileged portal is enabled
func (o PRAPortalOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PRAPortal) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
// microtenantId as 0 when making requests to retrieve data from the Default Microtenant. Pass microtenantId as null to
// retrieve data from all customers associated with the tenant.
func (o PRAPortalOutput) MicrotenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *PRAPortal) pulumi.StringOutput { return v.MicrotenantId }).(pulumi.StringOutput)
}

// The name of the privileged portal
func (o PRAPortalOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PRAPortal) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The notification message displayed in the banner of the privileged portallink, if enabled
func (o PRAPortalOutput) UserNotification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PRAPortal) pulumi.StringPtrOutput { return v.UserNotification }).(pulumi.StringPtrOutput)
}

// Indicates if the Notification Banner is enabled (true) or disabled (false)
func (o PRAPortalOutput) UserNotificationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PRAPortal) pulumi.BoolPtrOutput { return v.UserNotificationEnabled }).(pulumi.BoolPtrOutput)
}

type PRAPortalArrayOutput struct{ *pulumi.OutputState }

func (PRAPortalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PRAPortal)(nil)).Elem()
}

func (o PRAPortalArrayOutput) ToPRAPortalArrayOutput() PRAPortalArrayOutput {
	return o
}

func (o PRAPortalArrayOutput) ToPRAPortalArrayOutputWithContext(ctx context.Context) PRAPortalArrayOutput {
	return o
}

func (o PRAPortalArrayOutput) Index(i pulumi.IntInput) PRAPortalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PRAPortal {
		return vs[0].([]*PRAPortal)[vs[1].(int)]
	}).(PRAPortalOutput)
}

type PRAPortalMapOutput struct{ *pulumi.OutputState }

func (PRAPortalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PRAPortal)(nil)).Elem()
}

func (o PRAPortalMapOutput) ToPRAPortalMapOutput() PRAPortalMapOutput {
	return o
}

func (o PRAPortalMapOutput) ToPRAPortalMapOutputWithContext(ctx context.Context) PRAPortalMapOutput {
	return o
}

func (o PRAPortalMapOutput) MapIndex(k pulumi.StringInput) PRAPortalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PRAPortal {
		return vs[0].(map[string]*PRAPortal)[vs[1].(string)]
	}).(PRAPortalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PRAPortalInput)(nil)).Elem(), &PRAPortal{})
	pulumi.RegisterInputType(reflect.TypeOf((*PRAPortalArrayInput)(nil)).Elem(), PRAPortalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PRAPortalMapInput)(nil)).Elem(), PRAPortalMap{})
	pulumi.RegisterOutputType(PRAPortalOutput{})
	pulumi.RegisterOutputType(PRAPortalArrayOutput{})
	pulumi.RegisterOutputType(PRAPortalMapOutput{})
}