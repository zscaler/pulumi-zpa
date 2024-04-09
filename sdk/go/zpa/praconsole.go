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

// * [Official documentation](https://help.zscaler.com/zpa/about-privileged-consoles)
// * [API documentation](https://help.zscaler.com/zpa/configuring-privileged-consoles-using-api)
//
// The **zpa_pra_console_controller** resource creates a privileged remote access console in the Zscaler Private Access cloud. This resource can then be referenced in an privileged access policy resource and a privileged access portal.
//
// ## Import
//
// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
//
// # Visit
//
// **pra_credential_controller** can be imported by using `<CONSOLE ID>` or `<CONSOLE NAME>` as the import ID.
//
// For example:
//
// ```sh
// $ pulumi import zpa:index/pRAConsole:PRAConsole this <console_id>
// ```
//
// or
//
// ```sh
// $ pulumi import zpa:index/pRAConsole:PRAConsole this <console_name>
// ```
type PRAConsole struct {
	pulumi.CustomResourceState

	// The description of the privileged console
	Description pulumi.StringOutput `pulumi:"description"`
	// Whether or not the privileged console is enabled
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The privileged console icon. The icon image is converted to base64 encoded text format
	IconText pulumi.StringOutput `pulumi:"iconText"`
	// The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
	// microtenantId as 0 when making requests to retrieve data from the Default Microtenant.
	MicrotenantId pulumi.StringOutput `pulumi:"microtenantId"`
	// The name of the privileged console
	Name           pulumi.StringOutput            `pulumi:"name"`
	PraApplication PRAConsolePraApplicationOutput `pulumi:"praApplication"`
	PraPortals     PRAConsolePraPortalArrayOutput `pulumi:"praPortals"`
}

// NewPRAConsole registers a new resource with the given unique name, arguments, and options.
func NewPRAConsole(ctx *pulumi.Context,
	name string, args *PRAConsoleArgs, opts ...pulumi.ResourceOption) (*PRAConsole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PraApplication == nil {
		return nil, errors.New("invalid value for required argument 'PraApplication'")
	}
	if args.PraPortals == nil {
		return nil, errors.New("invalid value for required argument 'PraPortals'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("zpa:index/praConsoleController:PraConsoleController"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PRAConsole
	err := ctx.RegisterResource("zpa:index/pRAConsole:PRAConsole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPRAConsole gets an existing PRAConsole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPRAConsole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PRAConsoleState, opts ...pulumi.ResourceOption) (*PRAConsole, error) {
	var resource PRAConsole
	err := ctx.ReadResource("zpa:index/pRAConsole:PRAConsole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PRAConsole resources.
type praconsoleState struct {
	// The description of the privileged console
	Description *string `pulumi:"description"`
	// Whether or not the privileged console is enabled
	Enabled *bool `pulumi:"enabled"`
	// The privileged console icon. The icon image is converted to base64 encoded text format
	IconText *string `pulumi:"iconText"`
	// The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
	// microtenantId as 0 when making requests to retrieve data from the Default Microtenant.
	MicrotenantId *string `pulumi:"microtenantId"`
	// The name of the privileged console
	Name           *string                   `pulumi:"name"`
	PraApplication *PRAConsolePraApplication `pulumi:"praApplication"`
	PraPortals     []PRAConsolePraPortal     `pulumi:"praPortals"`
}

type PRAConsoleState struct {
	// The description of the privileged console
	Description pulumi.StringPtrInput
	// Whether or not the privileged console is enabled
	Enabled pulumi.BoolPtrInput
	// The privileged console icon. The icon image is converted to base64 encoded text format
	IconText pulumi.StringPtrInput
	// The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
	// microtenantId as 0 when making requests to retrieve data from the Default Microtenant.
	MicrotenantId pulumi.StringPtrInput
	// The name of the privileged console
	Name           pulumi.StringPtrInput
	PraApplication PRAConsolePraApplicationPtrInput
	PraPortals     PRAConsolePraPortalArrayInput
}

func (PRAConsoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*praconsoleState)(nil)).Elem()
}

type praconsoleArgs struct {
	// The description of the privileged console
	Description *string `pulumi:"description"`
	// Whether or not the privileged console is enabled
	Enabled *bool `pulumi:"enabled"`
	// The privileged console icon. The icon image is converted to base64 encoded text format
	IconText *string `pulumi:"iconText"`
	// The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
	// microtenantId as 0 when making requests to retrieve data from the Default Microtenant.
	MicrotenantId *string `pulumi:"microtenantId"`
	// The name of the privileged console
	Name           *string                  `pulumi:"name"`
	PraApplication PRAConsolePraApplication `pulumi:"praApplication"`
	PraPortals     []PRAConsolePraPortal    `pulumi:"praPortals"`
}

// The set of arguments for constructing a PRAConsole resource.
type PRAConsoleArgs struct {
	// The description of the privileged console
	Description pulumi.StringPtrInput
	// Whether or not the privileged console is enabled
	Enabled pulumi.BoolPtrInput
	// The privileged console icon. The icon image is converted to base64 encoded text format
	IconText pulumi.StringPtrInput
	// The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
	// microtenantId as 0 when making requests to retrieve data from the Default Microtenant.
	MicrotenantId pulumi.StringPtrInput
	// The name of the privileged console
	Name           pulumi.StringPtrInput
	PraApplication PRAConsolePraApplicationInput
	PraPortals     PRAConsolePraPortalArrayInput
}

func (PRAConsoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*praconsoleArgs)(nil)).Elem()
}

type PRAConsoleInput interface {
	pulumi.Input

	ToPRAConsoleOutput() PRAConsoleOutput
	ToPRAConsoleOutputWithContext(ctx context.Context) PRAConsoleOutput
}

func (*PRAConsole) ElementType() reflect.Type {
	return reflect.TypeOf((**PRAConsole)(nil)).Elem()
}

func (i *PRAConsole) ToPRAConsoleOutput() PRAConsoleOutput {
	return i.ToPRAConsoleOutputWithContext(context.Background())
}

func (i *PRAConsole) ToPRAConsoleOutputWithContext(ctx context.Context) PRAConsoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PRAConsoleOutput)
}

// PRAConsoleArrayInput is an input type that accepts PRAConsoleArray and PRAConsoleArrayOutput values.
// You can construct a concrete instance of `PRAConsoleArrayInput` via:
//
//	PRAConsoleArray{ PRAConsoleArgs{...} }
type PRAConsoleArrayInput interface {
	pulumi.Input

	ToPRAConsoleArrayOutput() PRAConsoleArrayOutput
	ToPRAConsoleArrayOutputWithContext(context.Context) PRAConsoleArrayOutput
}

type PRAConsoleArray []PRAConsoleInput

func (PRAConsoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PRAConsole)(nil)).Elem()
}

func (i PRAConsoleArray) ToPRAConsoleArrayOutput() PRAConsoleArrayOutput {
	return i.ToPRAConsoleArrayOutputWithContext(context.Background())
}

func (i PRAConsoleArray) ToPRAConsoleArrayOutputWithContext(ctx context.Context) PRAConsoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PRAConsoleArrayOutput)
}

// PRAConsoleMapInput is an input type that accepts PRAConsoleMap and PRAConsoleMapOutput values.
// You can construct a concrete instance of `PRAConsoleMapInput` via:
//
//	PRAConsoleMap{ "key": PRAConsoleArgs{...} }
type PRAConsoleMapInput interface {
	pulumi.Input

	ToPRAConsoleMapOutput() PRAConsoleMapOutput
	ToPRAConsoleMapOutputWithContext(context.Context) PRAConsoleMapOutput
}

type PRAConsoleMap map[string]PRAConsoleInput

func (PRAConsoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PRAConsole)(nil)).Elem()
}

func (i PRAConsoleMap) ToPRAConsoleMapOutput() PRAConsoleMapOutput {
	return i.ToPRAConsoleMapOutputWithContext(context.Background())
}

func (i PRAConsoleMap) ToPRAConsoleMapOutputWithContext(ctx context.Context) PRAConsoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PRAConsoleMapOutput)
}

type PRAConsoleOutput struct{ *pulumi.OutputState }

func (PRAConsoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PRAConsole)(nil)).Elem()
}

func (o PRAConsoleOutput) ToPRAConsoleOutput() PRAConsoleOutput {
	return o
}

func (o PRAConsoleOutput) ToPRAConsoleOutputWithContext(ctx context.Context) PRAConsoleOutput {
	return o
}

// The description of the privileged console
func (o PRAConsoleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *PRAConsole) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Whether or not the privileged console is enabled
func (o PRAConsoleOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *PRAConsole) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// The privileged console icon. The icon image is converted to base64 encoded text format
func (o PRAConsoleOutput) IconText() pulumi.StringOutput {
	return o.ApplyT(func(v *PRAConsole) pulumi.StringOutput { return v.IconText }).(pulumi.StringOutput)
}

// The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
// microtenantId as 0 when making requests to retrieve data from the Default Microtenant.
func (o PRAConsoleOutput) MicrotenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *PRAConsole) pulumi.StringOutput { return v.MicrotenantId }).(pulumi.StringOutput)
}

// The name of the privileged console
func (o PRAConsoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PRAConsole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PRAConsoleOutput) PraApplication() PRAConsolePraApplicationOutput {
	return o.ApplyT(func(v *PRAConsole) PRAConsolePraApplicationOutput { return v.PraApplication }).(PRAConsolePraApplicationOutput)
}

func (o PRAConsoleOutput) PraPortals() PRAConsolePraPortalArrayOutput {
	return o.ApplyT(func(v *PRAConsole) PRAConsolePraPortalArrayOutput { return v.PraPortals }).(PRAConsolePraPortalArrayOutput)
}

type PRAConsoleArrayOutput struct{ *pulumi.OutputState }

func (PRAConsoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PRAConsole)(nil)).Elem()
}

func (o PRAConsoleArrayOutput) ToPRAConsoleArrayOutput() PRAConsoleArrayOutput {
	return o
}

func (o PRAConsoleArrayOutput) ToPRAConsoleArrayOutputWithContext(ctx context.Context) PRAConsoleArrayOutput {
	return o
}

func (o PRAConsoleArrayOutput) Index(i pulumi.IntInput) PRAConsoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PRAConsole {
		return vs[0].([]*PRAConsole)[vs[1].(int)]
	}).(PRAConsoleOutput)
}

type PRAConsoleMapOutput struct{ *pulumi.OutputState }

func (PRAConsoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PRAConsole)(nil)).Elem()
}

func (o PRAConsoleMapOutput) ToPRAConsoleMapOutput() PRAConsoleMapOutput {
	return o
}

func (o PRAConsoleMapOutput) ToPRAConsoleMapOutputWithContext(ctx context.Context) PRAConsoleMapOutput {
	return o
}

func (o PRAConsoleMapOutput) MapIndex(k pulumi.StringInput) PRAConsoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PRAConsole {
		return vs[0].(map[string]*PRAConsole)[vs[1].(string)]
	}).(PRAConsoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PRAConsoleInput)(nil)).Elem(), &PRAConsole{})
	pulumi.RegisterInputType(reflect.TypeOf((*PRAConsoleArrayInput)(nil)).Elem(), PRAConsoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PRAConsoleMapInput)(nil)).Elem(), PRAConsoleMap{})
	pulumi.RegisterOutputType(PRAConsoleOutput{})
	pulumi.RegisterOutputType(PRAConsoleArrayOutput{})
	pulumi.RegisterOutputType(PRAConsoleMapOutput{})
}
