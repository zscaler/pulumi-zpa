// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
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
//			// Create a App Connector Group
//			exampleConnectorGroup, err := zpa.NewConnectorGroup(ctx, "exampleConnectorGroup", &zpa.ConnectorGroupArgs{
//				Description:            pulumi.String("Example"),
//				Enabled:                pulumi.Bool(true),
//				CityCountry:            pulumi.String("San Jose, CA"),
//				CountryCode:            pulumi.String("US"),
//				Latitude:               pulumi.String("37.338"),
//				Longitude:              pulumi.String("-121.8863"),
//				Location:               pulumi.String("San Jose, CA, US"),
//				UpgradeDay:             pulumi.String("SUNDAY"),
//				UpgradeTimeInSecs:      pulumi.String("66600"),
//				OverrideVersionProfile: pulumi.Bool(true),
//				VersionProfileId:       pulumi.String("0"),
//				DnsQueryType:           pulumi.String("IPV4"),
//			})
//			if err != nil {
//				return err
//			}
//			// Create a Server Group resource with Dynamic Discovery Enabled
//			_, err = zpa.NewServerGroup(ctx, "exampleServerGroup", &zpa.ServerGroupArgs{
//				Description:      pulumi.String("Example"),
//				Enabled:          pulumi.Bool(true),
//				DynamicDiscovery: pulumi.Bool(true),
//				AppConnectorGroups: zpa.ServerGroupAppConnectorGroupArray{
//					&zpa.ServerGroupAppConnectorGroupArgs{
//						Ids: pulumi.StringArray{
//							exampleConnectorGroup.ID(),
//						},
//					},
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				exampleConnectorGroup,
//			}))
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
//			// Create an application server
//			exampleApplicationServer, err := zpa.NewApplicationServer(ctx, "exampleApplicationServer", &zpa.ApplicationServerArgs{
//				Description: pulumi.String("Example"),
//				Address:     pulumi.String("server.example.com"),
//				Enabled:     pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			// Create a App Connector Group
//			exampleConnectorGroup, err := zpa.NewConnectorGroup(ctx, "exampleConnectorGroup", &zpa.ConnectorGroupArgs{
//				Description:            pulumi.String("Example"),
//				Enabled:                pulumi.Bool(true),
//				CityCountry:            pulumi.String("San Jose, CA"),
//				CountryCode:            pulumi.String("US"),
//				Latitude:               pulumi.String("37.338"),
//				Longitude:              pulumi.String("-121.8863"),
//				Location:               pulumi.String("San Jose, CA, US"),
//				UpgradeDay:             pulumi.String("SUNDAY"),
//				UpgradeTimeInSecs:      pulumi.String("66600"),
//				OverrideVersionProfile: pulumi.Bool(true),
//				VersionProfileId:       pulumi.String("0"),
//				DnsQueryType:           pulumi.String("IPV4"),
//			})
//			if err != nil {
//				return err
//			}
//			// ZPA Server Group resource with Dynamic Discovery Disabled
//			_, err = zpa.NewServerGroup(ctx, "exampleServerGroup", &zpa.ServerGroupArgs{
//				Description:      pulumi.String("Example"),
//				Enabled:          pulumi.Bool(true),
//				DynamicDiscovery: pulumi.Bool(false),
//				Servers: zpa.ServerGroupServerArray{
//					&zpa.ServerGroupServerArgs{
//						Ids: pulumi.StringArray{
//							exampleApplicationServer.ID(),
//						},
//					},
//				},
//				AppConnectorGroups: zpa.ServerGroupAppConnectorGroupArray{
//					&zpa.ServerGroupAppConnectorGroupArgs{
//						Ids: pulumi.StringArray{
//							exampleConnectorGroup.ID(),
//						},
//					},
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				exampleConnectorGroup,
//				zpa_application_server.Server,
//			}))
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
// ### Required
//
// * `name` - (Required) This field defines the name of the server group.
// * `appConnectorGroups` - (Required)
//   - `id` - (Required) The ID of this resource.
//
// ## Import
//
// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
//
// # Visit
//
// Server Groups can be imported; use `<SERVER GROUP ID>` or `<SERVER GROUP NAME>` as the import ID.
//
// For example:
//
// ```sh
// $ pulumi import zpa:index/serverGroup:ServerGroup example <server_group_id>
// ```
//
// or
//
// ```sh
// $ pulumi import zpa:index/serverGroup:ServerGroup example <server_group_name>
// ```
type ServerGroup struct {
	pulumi.CustomResourceState

	// List of app-connector IDs.
	AppConnectorGroups ServerGroupAppConnectorGroupArrayOutput `pulumi:"appConnectorGroups"`
	// This field is a json array of app-connector-id only.
	Applications ServerGroupApplicationArrayOutput `pulumi:"applications"`
	ConfigSpace  pulumi.StringPtrOutput            `pulumi:"configSpace"`
	// (Optional) This field is the description of the server group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// (Optional) This field controls dynamic discovery of the servers.
	DynamicDiscovery pulumi.BoolPtrOutput `pulumi:"dynamicDiscovery"`
	// (Optional) This field defines if the server group is enabled or disabled.
	Enabled    pulumi.BoolPtrOutput `pulumi:"enabled"`
	IpAnchored pulumi.BoolPtrOutput `pulumi:"ipAnchored"`
	// (Optional) The ID of the microtenant the resource is to be associated with.
	MicrotenantId pulumi.StringOutput `pulumi:"microtenantId"`
	// This field defines the name of the server group.
	Name pulumi.StringOutput `pulumi:"name"`
	// (Block List) This field is a list of servers that are applicable only when dynamic discovery is disabled. Server name is required only in cases where the new servers need to be created in this API.
	Servers ServerGroupServerArrayOutput `pulumi:"servers"`
}

// NewServerGroup registers a new resource with the given unique name, arguments, and options.
func NewServerGroup(ctx *pulumi.Context,
	name string, args *ServerGroupArgs, opts ...pulumi.ResourceOption) (*ServerGroup, error) {
	if args == nil {
		args = &ServerGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServerGroup
	err := ctx.RegisterResource("zpa:index/serverGroup:ServerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerGroup gets an existing ServerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerGroupState, opts ...pulumi.ResourceOption) (*ServerGroup, error) {
	var resource ServerGroup
	err := ctx.ReadResource("zpa:index/serverGroup:ServerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerGroup resources.
type serverGroupState struct {
	// List of app-connector IDs.
	AppConnectorGroups []ServerGroupAppConnectorGroup `pulumi:"appConnectorGroups"`
	// This field is a json array of app-connector-id only.
	Applications []ServerGroupApplication `pulumi:"applications"`
	ConfigSpace  *string                  `pulumi:"configSpace"`
	// (Optional) This field is the description of the server group.
	Description *string `pulumi:"description"`
	// (Optional) This field controls dynamic discovery of the servers.
	DynamicDiscovery *bool `pulumi:"dynamicDiscovery"`
	// (Optional) This field defines if the server group is enabled or disabled.
	Enabled    *bool `pulumi:"enabled"`
	IpAnchored *bool `pulumi:"ipAnchored"`
	// (Optional) The ID of the microtenant the resource is to be associated with.
	MicrotenantId *string `pulumi:"microtenantId"`
	// This field defines the name of the server group.
	Name *string `pulumi:"name"`
	// (Block List) This field is a list of servers that are applicable only when dynamic discovery is disabled. Server name is required only in cases where the new servers need to be created in this API.
	Servers []ServerGroupServer `pulumi:"servers"`
}

type ServerGroupState struct {
	// List of app-connector IDs.
	AppConnectorGroups ServerGroupAppConnectorGroupArrayInput
	// This field is a json array of app-connector-id only.
	Applications ServerGroupApplicationArrayInput
	ConfigSpace  pulumi.StringPtrInput
	// (Optional) This field is the description of the server group.
	Description pulumi.StringPtrInput
	// (Optional) This field controls dynamic discovery of the servers.
	DynamicDiscovery pulumi.BoolPtrInput
	// (Optional) This field defines if the server group is enabled or disabled.
	Enabled    pulumi.BoolPtrInput
	IpAnchored pulumi.BoolPtrInput
	// (Optional) The ID of the microtenant the resource is to be associated with.
	MicrotenantId pulumi.StringPtrInput
	// This field defines the name of the server group.
	Name pulumi.StringPtrInput
	// (Block List) This field is a list of servers that are applicable only when dynamic discovery is disabled. Server name is required only in cases where the new servers need to be created in this API.
	Servers ServerGroupServerArrayInput
}

func (ServerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverGroupState)(nil)).Elem()
}

type serverGroupArgs struct {
	// List of app-connector IDs.
	AppConnectorGroups []ServerGroupAppConnectorGroup `pulumi:"appConnectorGroups"`
	// This field is a json array of app-connector-id only.
	Applications []ServerGroupApplication `pulumi:"applications"`
	ConfigSpace  *string                  `pulumi:"configSpace"`
	// (Optional) This field is the description of the server group.
	Description *string `pulumi:"description"`
	// (Optional) This field controls dynamic discovery of the servers.
	DynamicDiscovery *bool `pulumi:"dynamicDiscovery"`
	// (Optional) This field defines if the server group is enabled or disabled.
	Enabled    *bool `pulumi:"enabled"`
	IpAnchored *bool `pulumi:"ipAnchored"`
	// (Optional) The ID of the microtenant the resource is to be associated with.
	MicrotenantId *string `pulumi:"microtenantId"`
	// This field defines the name of the server group.
	Name *string `pulumi:"name"`
	// (Block List) This field is a list of servers that are applicable only when dynamic discovery is disabled. Server name is required only in cases where the new servers need to be created in this API.
	Servers []ServerGroupServer `pulumi:"servers"`
}

// The set of arguments for constructing a ServerGroup resource.
type ServerGroupArgs struct {
	// List of app-connector IDs.
	AppConnectorGroups ServerGroupAppConnectorGroupArrayInput
	// This field is a json array of app-connector-id only.
	Applications ServerGroupApplicationArrayInput
	ConfigSpace  pulumi.StringPtrInput
	// (Optional) This field is the description of the server group.
	Description pulumi.StringPtrInput
	// (Optional) This field controls dynamic discovery of the servers.
	DynamicDiscovery pulumi.BoolPtrInput
	// (Optional) This field defines if the server group is enabled or disabled.
	Enabled    pulumi.BoolPtrInput
	IpAnchored pulumi.BoolPtrInput
	// (Optional) The ID of the microtenant the resource is to be associated with.
	MicrotenantId pulumi.StringPtrInput
	// This field defines the name of the server group.
	Name pulumi.StringPtrInput
	// (Block List) This field is a list of servers that are applicable only when dynamic discovery is disabled. Server name is required only in cases where the new servers need to be created in this API.
	Servers ServerGroupServerArrayInput
}

func (ServerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverGroupArgs)(nil)).Elem()
}

type ServerGroupInput interface {
	pulumi.Input

	ToServerGroupOutput() ServerGroupOutput
	ToServerGroupOutputWithContext(ctx context.Context) ServerGroupOutput
}

func (*ServerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroup)(nil)).Elem()
}

func (i *ServerGroup) ToServerGroupOutput() ServerGroupOutput {
	return i.ToServerGroupOutputWithContext(context.Background())
}

func (i *ServerGroup) ToServerGroupOutputWithContext(ctx context.Context) ServerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupOutput)
}

// ServerGroupArrayInput is an input type that accepts ServerGroupArray and ServerGroupArrayOutput values.
// You can construct a concrete instance of `ServerGroupArrayInput` via:
//
//	ServerGroupArray{ ServerGroupArgs{...} }
type ServerGroupArrayInput interface {
	pulumi.Input

	ToServerGroupArrayOutput() ServerGroupArrayOutput
	ToServerGroupArrayOutputWithContext(context.Context) ServerGroupArrayOutput
}

type ServerGroupArray []ServerGroupInput

func (ServerGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerGroup)(nil)).Elem()
}

func (i ServerGroupArray) ToServerGroupArrayOutput() ServerGroupArrayOutput {
	return i.ToServerGroupArrayOutputWithContext(context.Background())
}

func (i ServerGroupArray) ToServerGroupArrayOutputWithContext(ctx context.Context) ServerGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupArrayOutput)
}

// ServerGroupMapInput is an input type that accepts ServerGroupMap and ServerGroupMapOutput values.
// You can construct a concrete instance of `ServerGroupMapInput` via:
//
//	ServerGroupMap{ "key": ServerGroupArgs{...} }
type ServerGroupMapInput interface {
	pulumi.Input

	ToServerGroupMapOutput() ServerGroupMapOutput
	ToServerGroupMapOutputWithContext(context.Context) ServerGroupMapOutput
}

type ServerGroupMap map[string]ServerGroupInput

func (ServerGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerGroup)(nil)).Elem()
}

func (i ServerGroupMap) ToServerGroupMapOutput() ServerGroupMapOutput {
	return i.ToServerGroupMapOutputWithContext(context.Background())
}

func (i ServerGroupMap) ToServerGroupMapOutputWithContext(ctx context.Context) ServerGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupMapOutput)
}

type ServerGroupOutput struct{ *pulumi.OutputState }

func (ServerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroup)(nil)).Elem()
}

func (o ServerGroupOutput) ToServerGroupOutput() ServerGroupOutput {
	return o
}

func (o ServerGroupOutput) ToServerGroupOutputWithContext(ctx context.Context) ServerGroupOutput {
	return o
}

// List of app-connector IDs.
func (o ServerGroupOutput) AppConnectorGroups() ServerGroupAppConnectorGroupArrayOutput {
	return o.ApplyT(func(v *ServerGroup) ServerGroupAppConnectorGroupArrayOutput { return v.AppConnectorGroups }).(ServerGroupAppConnectorGroupArrayOutput)
}

// This field is a json array of app-connector-id only.
func (o ServerGroupOutput) Applications() ServerGroupApplicationArrayOutput {
	return o.ApplyT(func(v *ServerGroup) ServerGroupApplicationArrayOutput { return v.Applications }).(ServerGroupApplicationArrayOutput)
}

func (o ServerGroupOutput) ConfigSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringPtrOutput { return v.ConfigSpace }).(pulumi.StringPtrOutput)
}

// (Optional) This field is the description of the server group.
func (o ServerGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// (Optional) This field controls dynamic discovery of the servers.
func (o ServerGroupOutput) DynamicDiscovery() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.BoolPtrOutput { return v.DynamicDiscovery }).(pulumi.BoolPtrOutput)
}

// (Optional) This field defines if the server group is enabled or disabled.
func (o ServerGroupOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ServerGroupOutput) IpAnchored() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.BoolPtrOutput { return v.IpAnchored }).(pulumi.BoolPtrOutput)
}

// (Optional) The ID of the microtenant the resource is to be associated with.
func (o ServerGroupOutput) MicrotenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.MicrotenantId }).(pulumi.StringOutput)
}

// This field defines the name of the server group.
func (o ServerGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// (Block List) This field is a list of servers that are applicable only when dynamic discovery is disabled. Server name is required only in cases where the new servers need to be created in this API.
func (o ServerGroupOutput) Servers() ServerGroupServerArrayOutput {
	return o.ApplyT(func(v *ServerGroup) ServerGroupServerArrayOutput { return v.Servers }).(ServerGroupServerArrayOutput)
}

type ServerGroupArrayOutput struct{ *pulumi.OutputState }

func (ServerGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerGroup)(nil)).Elem()
}

func (o ServerGroupArrayOutput) ToServerGroupArrayOutput() ServerGroupArrayOutput {
	return o
}

func (o ServerGroupArrayOutput) ToServerGroupArrayOutputWithContext(ctx context.Context) ServerGroupArrayOutput {
	return o
}

func (o ServerGroupArrayOutput) Index(i pulumi.IntInput) ServerGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerGroup {
		return vs[0].([]*ServerGroup)[vs[1].(int)]
	}).(ServerGroupOutput)
}

type ServerGroupMapOutput struct{ *pulumi.OutputState }

func (ServerGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerGroup)(nil)).Elem()
}

func (o ServerGroupMapOutput) ToServerGroupMapOutput() ServerGroupMapOutput {
	return o
}

func (o ServerGroupMapOutput) ToServerGroupMapOutputWithContext(ctx context.Context) ServerGroupMapOutput {
	return o
}

func (o ServerGroupMapOutput) MapIndex(k pulumi.StringInput) ServerGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerGroup {
		return vs[0].(map[string]*ServerGroup)[vs[1].(string)]
	}).(ServerGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupInput)(nil)).Elem(), &ServerGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupArrayInput)(nil)).Elem(), ServerGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupMapInput)(nil)).Elem(), ServerGroupMap{})
	pulumi.RegisterOutputType(ServerGroupOutput{})
	pulumi.RegisterOutputType(ServerGroupArrayOutput{})
	pulumi.RegisterOutputType(ServerGroupMapOutput{})
}
