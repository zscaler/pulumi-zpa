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
//			// ZPA Application Server resource (IP Address)
//			_, err := zpa.NewApplicationServer(ctx, "testAppServer", &zpa.ApplicationServerArgs{
//				Address:     pulumi.String("192.168.1.1"),
//				Description: pulumi.String("test1-app-server"),
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
// <!--End PulumiCodeChooser -->
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
//			// ZPA Application Server resource (FQDN Address)
//			_, err := zpa.NewApplicationServer(ctx, "testAppServer", &zpa.ApplicationServerArgs{
//				Address:     pulumi.String("server1.acme.com"),
//				Description: pulumi.String("test1-app-server"),
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
//
// # Visit
//
// Application Server can be imported by using `<APPLICATION SERVER ID>` or `<APPLICATION SERVER NAME>` as the import ID
//
// For example:
//
// ```sh
// $ pulumi import zpa:index/applicationServer:ApplicationServer example <application_server_id>
// ```
//
// or
//
// ```sh
// $ pulumi import zpa:index/applicationServer:ApplicationServer example <application_server_name>
// ```
type ApplicationServer struct {
	pulumi.CustomResourceState

	// This field defines the domain or IP address of the server.
	Address pulumi.StringOutput `pulumi:"address"`
	// This field defines the list of server groups IDs.
	AppServerGroupIds pulumi.StringArrayOutput `pulumi:"appServerGroupIds"`
	ConfigSpace       pulumi.StringPtrOutput   `pulumi:"configSpace"`
	// This field defines the description of the server.
	Description pulumi.StringOutput `pulumi:"description"`
	// This field defines the status of the server.
	Enabled       pulumi.BoolOutput   `pulumi:"enabled"`
	MicrotenantId pulumi.StringOutput `pulumi:"microtenantId"`
	// This field defines the name of the server.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewApplicationServer registers a new resource with the given unique name, arguments, and options.
func NewApplicationServer(ctx *pulumi.Context,
	name string, args *ApplicationServerArgs, opts ...pulumi.ResourceOption) (*ApplicationServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationServer
	err := ctx.RegisterResource("zpa:index/applicationServer:ApplicationServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationServer gets an existing ApplicationServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationServerState, opts ...pulumi.ResourceOption) (*ApplicationServer, error) {
	var resource ApplicationServer
	err := ctx.ReadResource("zpa:index/applicationServer:ApplicationServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationServer resources.
type applicationServerState struct {
	// This field defines the domain or IP address of the server.
	Address *string `pulumi:"address"`
	// This field defines the list of server groups IDs.
	AppServerGroupIds []string `pulumi:"appServerGroupIds"`
	ConfigSpace       *string  `pulumi:"configSpace"`
	// This field defines the description of the server.
	Description *string `pulumi:"description"`
	// This field defines the status of the server.
	Enabled       *bool   `pulumi:"enabled"`
	MicrotenantId *string `pulumi:"microtenantId"`
	// This field defines the name of the server.
	Name *string `pulumi:"name"`
}

type ApplicationServerState struct {
	// This field defines the domain or IP address of the server.
	Address pulumi.StringPtrInput
	// This field defines the list of server groups IDs.
	AppServerGroupIds pulumi.StringArrayInput
	ConfigSpace       pulumi.StringPtrInput
	// This field defines the description of the server.
	Description pulumi.StringPtrInput
	// This field defines the status of the server.
	Enabled       pulumi.BoolPtrInput
	MicrotenantId pulumi.StringPtrInput
	// This field defines the name of the server.
	Name pulumi.StringPtrInput
}

func (ApplicationServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationServerState)(nil)).Elem()
}

type applicationServerArgs struct {
	// This field defines the domain or IP address of the server.
	Address string `pulumi:"address"`
	// This field defines the list of server groups IDs.
	AppServerGroupIds []string `pulumi:"appServerGroupIds"`
	ConfigSpace       *string  `pulumi:"configSpace"`
	// This field defines the description of the server.
	Description *string `pulumi:"description"`
	// This field defines the status of the server.
	Enabled       *bool   `pulumi:"enabled"`
	MicrotenantId *string `pulumi:"microtenantId"`
	// This field defines the name of the server.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ApplicationServer resource.
type ApplicationServerArgs struct {
	// This field defines the domain or IP address of the server.
	Address pulumi.StringInput
	// This field defines the list of server groups IDs.
	AppServerGroupIds pulumi.StringArrayInput
	ConfigSpace       pulumi.StringPtrInput
	// This field defines the description of the server.
	Description pulumi.StringPtrInput
	// This field defines the status of the server.
	Enabled       pulumi.BoolPtrInput
	MicrotenantId pulumi.StringPtrInput
	// This field defines the name of the server.
	Name pulumi.StringPtrInput
}

func (ApplicationServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationServerArgs)(nil)).Elem()
}

type ApplicationServerInput interface {
	pulumi.Input

	ToApplicationServerOutput() ApplicationServerOutput
	ToApplicationServerOutputWithContext(ctx context.Context) ApplicationServerOutput
}

func (*ApplicationServer) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationServer)(nil)).Elem()
}

func (i *ApplicationServer) ToApplicationServerOutput() ApplicationServerOutput {
	return i.ToApplicationServerOutputWithContext(context.Background())
}

func (i *ApplicationServer) ToApplicationServerOutputWithContext(ctx context.Context) ApplicationServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationServerOutput)
}

// ApplicationServerArrayInput is an input type that accepts ApplicationServerArray and ApplicationServerArrayOutput values.
// You can construct a concrete instance of `ApplicationServerArrayInput` via:
//
//	ApplicationServerArray{ ApplicationServerArgs{...} }
type ApplicationServerArrayInput interface {
	pulumi.Input

	ToApplicationServerArrayOutput() ApplicationServerArrayOutput
	ToApplicationServerArrayOutputWithContext(context.Context) ApplicationServerArrayOutput
}

type ApplicationServerArray []ApplicationServerInput

func (ApplicationServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationServer)(nil)).Elem()
}

func (i ApplicationServerArray) ToApplicationServerArrayOutput() ApplicationServerArrayOutput {
	return i.ToApplicationServerArrayOutputWithContext(context.Background())
}

func (i ApplicationServerArray) ToApplicationServerArrayOutputWithContext(ctx context.Context) ApplicationServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationServerArrayOutput)
}

// ApplicationServerMapInput is an input type that accepts ApplicationServerMap and ApplicationServerMapOutput values.
// You can construct a concrete instance of `ApplicationServerMapInput` via:
//
//	ApplicationServerMap{ "key": ApplicationServerArgs{...} }
type ApplicationServerMapInput interface {
	pulumi.Input

	ToApplicationServerMapOutput() ApplicationServerMapOutput
	ToApplicationServerMapOutputWithContext(context.Context) ApplicationServerMapOutput
}

type ApplicationServerMap map[string]ApplicationServerInput

func (ApplicationServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationServer)(nil)).Elem()
}

func (i ApplicationServerMap) ToApplicationServerMapOutput() ApplicationServerMapOutput {
	return i.ToApplicationServerMapOutputWithContext(context.Background())
}

func (i ApplicationServerMap) ToApplicationServerMapOutputWithContext(ctx context.Context) ApplicationServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationServerMapOutput)
}

type ApplicationServerOutput struct{ *pulumi.OutputState }

func (ApplicationServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationServer)(nil)).Elem()
}

func (o ApplicationServerOutput) ToApplicationServerOutput() ApplicationServerOutput {
	return o
}

func (o ApplicationServerOutput) ToApplicationServerOutputWithContext(ctx context.Context) ApplicationServerOutput {
	return o
}

// This field defines the domain or IP address of the server.
func (o ApplicationServerOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationServer) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// This field defines the list of server groups IDs.
func (o ApplicationServerOutput) AppServerGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationServer) pulumi.StringArrayOutput { return v.AppServerGroupIds }).(pulumi.StringArrayOutput)
}

func (o ApplicationServerOutput) ConfigSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationServer) pulumi.StringPtrOutput { return v.ConfigSpace }).(pulumi.StringPtrOutput)
}

// This field defines the description of the server.
func (o ApplicationServerOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationServer) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// This field defines the status of the server.
func (o ApplicationServerOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApplicationServer) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o ApplicationServerOutput) MicrotenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationServer) pulumi.StringOutput { return v.MicrotenantId }).(pulumi.StringOutput)
}

// This field defines the name of the server.
func (o ApplicationServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type ApplicationServerArrayOutput struct{ *pulumi.OutputState }

func (ApplicationServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationServer)(nil)).Elem()
}

func (o ApplicationServerArrayOutput) ToApplicationServerArrayOutput() ApplicationServerArrayOutput {
	return o
}

func (o ApplicationServerArrayOutput) ToApplicationServerArrayOutputWithContext(ctx context.Context) ApplicationServerArrayOutput {
	return o
}

func (o ApplicationServerArrayOutput) Index(i pulumi.IntInput) ApplicationServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationServer {
		return vs[0].([]*ApplicationServer)[vs[1].(int)]
	}).(ApplicationServerOutput)
}

type ApplicationServerMapOutput struct{ *pulumi.OutputState }

func (ApplicationServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationServer)(nil)).Elem()
}

func (o ApplicationServerMapOutput) ToApplicationServerMapOutput() ApplicationServerMapOutput {
	return o
}

func (o ApplicationServerMapOutput) ToApplicationServerMapOutputWithContext(ctx context.Context) ApplicationServerMapOutput {
	return o
}

func (o ApplicationServerMapOutput) MapIndex(k pulumi.StringInput) ApplicationServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationServer {
		return vs[0].(map[string]*ApplicationServer)[vs[1].(string)]
	}).(ApplicationServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationServerInput)(nil)).Elem(), &ApplicationServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationServerArrayInput)(nil)).Elem(), ApplicationServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationServerMapInput)(nil)).Elem(), ApplicationServerMap{})
	pulumi.RegisterOutputType(ApplicationServerOutput{})
	pulumi.RegisterOutputType(ApplicationServerArrayOutput{})
	pulumi.RegisterOutputType(ApplicationServerMapOutput{})
}
