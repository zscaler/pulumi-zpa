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

// The **zpa_service_edge_group** resource creates a service edge group in the Zscaler Private Access cloud. This resource can then be referenced in a service edge connector.
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
//			// ZPA Service Edge Group resource - Trusted Network
//			_, err := zpa.NewServiceEdgeGroup(ctx, "serviceEdgeGroupSjc", &zpa.ServiceEdgeGroupArgs{
//				Description:        pulumi.String("Service Edge Group in San Jose"),
//				Enabled:            pulumi.Bool(true),
//				IsPublic:           pulumi.Bool(true),
//				UpgradeDay:         pulumi.String("SUNDAY"),
//				UpgradeTimeInSecs:  pulumi.String("66600"),
//				Latitude:           pulumi.String("37.3382082"),
//				Longitude:          pulumi.String("-121.8863286"),
//				Location:           pulumi.String("San Jose, CA, USA"),
//				VersionProfileName: pulumi.String("New Release"),
//				TrustedNetworks: zpa.ServiceEdgeGroupTrustedNetworkArray{
//					&zpa.ServiceEdgeGroupTrustedNetworkArgs{
//						Ids: pulumi.StringArray{
//							data.Zpa_trusted_network.Example.Id,
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
//			// ZPA Service Edge Group resource - No Trusted Network
//			_, err := zpa.NewServiceEdgeGroup(ctx, "serviceEdgeGroupNyc", &zpa.ServiceEdgeGroupArgs{
//				Description:        pulumi.String("Service Edge Group in New York"),
//				Enabled:            pulumi.Bool(true),
//				IsPublic:           pulumi.Bool(true),
//				Latitude:           pulumi.String("40.7128"),
//				Location:           pulumi.String("New York, NY, USA"),
//				Longitude:          pulumi.String("-73.935242"),
//				UpgradeDay:         pulumi.String("SUNDAY"),
//				UpgradeTimeInSecs:  pulumi.String("66600"),
//				VersionProfileName: pulumi.String("New Release"),
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
// Service Edge Group can be imported; use `<SERVER EDGE GROUP ID>` or `<SERVER EDGE GROUP NAME>` as the import ID.
//
// For example:
//
// ```sh
// $ pulumi import zpa:index/serviceEdgeGroup:ServiceEdgeGroup example <service_edge_group_id>
// ```
//
// or
//
// ```sh
// $ pulumi import zpa:index/serviceEdgeGroup:ServiceEdgeGroup example <service_edge_group_name>
// ```
type ServiceEdgeGroup struct {
	pulumi.CustomResourceState

	// This field controls dynamic discovery of the servers.
	CityCountry pulumi.StringOutput `pulumi:"cityCountry"`
	// This field is an array of app-connector-id only.
	CountryCode pulumi.StringOutput `pulumi:"countryCode"`
	// Description of the Service Edge Group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether this Service Edge Group is enabled or not. Default value: `true` Supported values: `true`, `false`
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Enable or disable public access for the Service Edge Group. Default value: `false` Supported values: `true`, `false`
	IsPublic pulumi.BoolPtrOutput `pulumi:"isPublic"`
	// Latitude for the Service Edge Group. Integer or decimal with values in the range of `-90` to `90`
	Latitude pulumi.StringOutput `pulumi:"latitude"`
	// Location for the Service Edge Group.
	Location pulumi.StringOutput `pulumi:"location"`
	// Longitude for the Service Edge Group. Integer or decimal with values in the range of `-180` to `180`
	Longitude pulumi.StringOutput `pulumi:"longitude"`
	// The ID of the microtenant the resource is to be associated with.
	//
	// ⚠️ **WARNING:**: The attribute ``microtenantId`` is optional and requires the microtenant license and feature flag enabled for the respective tenant. The provider also supports the microtenant ID configuration via the environment variable `ZPA_MICROTENANT_ID` which is the recommended method.
	MicrotenantId pulumi.StringOutput `pulumi:"microtenantId"`
	// Name of the Service Edge Group.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
	OverrideVersionProfile pulumi.BoolPtrOutput                   `pulumi:"overrideVersionProfile"`
	ServiceEdges           ServiceEdgeGroupServiceEdgeArrayOutput `pulumi:"serviceEdges"`
	// Trusted networks for this Service Edge Group. List of trusted network objects
	TrustedNetworks ServiceEdgeGroupTrustedNetworkArrayOutput `pulumi:"trustedNetworks"`
	// Service Edges in this group will attempt to update to a newer version of the software during this specified day. Default value: `SUNDAY` List of valid days (i.e., Sunday, Monday)
	UpgradeDay pulumi.StringPtrOutput `pulumi:"upgradeDay"`
	// Service Edges in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600` Integer in seconds (i..e, 66600). The integer must be greater than or equal to 0 and less than `86400`, in `15` minute intervals
	UpgradeTimeInSecs pulumi.StringPtrOutput `pulumi:"upgradeTimeInSecs"`
	// ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
	VersionProfileId pulumi.StringOutput `pulumi:"versionProfileId"`
	// ID of the version profile.
	VersionProfileName pulumi.StringOutput `pulumi:"versionProfileName"`
	// ID of the version profile.
	VersionProfileVisibilityScope pulumi.StringOutput `pulumi:"versionProfileVisibilityScope"`
}

// NewServiceEdgeGroup registers a new resource with the given unique name, arguments, and options.
func NewServiceEdgeGroup(ctx *pulumi.Context,
	name string, args *ServiceEdgeGroupArgs, opts ...pulumi.ResourceOption) (*ServiceEdgeGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Latitude == nil {
		return nil, errors.New("invalid value for required argument 'Latitude'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Longitude == nil {
		return nil, errors.New("invalid value for required argument 'Longitude'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceEdgeGroup
	err := ctx.RegisterResource("zpa:index/serviceEdgeGroup:ServiceEdgeGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEdgeGroup gets an existing ServiceEdgeGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEdgeGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEdgeGroupState, opts ...pulumi.ResourceOption) (*ServiceEdgeGroup, error) {
	var resource ServiceEdgeGroup
	err := ctx.ReadResource("zpa:index/serviceEdgeGroup:ServiceEdgeGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEdgeGroup resources.
type serviceEdgeGroupState struct {
	// This field controls dynamic discovery of the servers.
	CityCountry *string `pulumi:"cityCountry"`
	// This field is an array of app-connector-id only.
	CountryCode *string `pulumi:"countryCode"`
	// Description of the Service Edge Group.
	Description *string `pulumi:"description"`
	// Whether this Service Edge Group is enabled or not. Default value: `true` Supported values: `true`, `false`
	Enabled *bool `pulumi:"enabled"`
	// Enable or disable public access for the Service Edge Group. Default value: `false` Supported values: `true`, `false`
	IsPublic *bool `pulumi:"isPublic"`
	// Latitude for the Service Edge Group. Integer or decimal with values in the range of `-90` to `90`
	Latitude *string `pulumi:"latitude"`
	// Location for the Service Edge Group.
	Location *string `pulumi:"location"`
	// Longitude for the Service Edge Group. Integer or decimal with values in the range of `-180` to `180`
	Longitude *string `pulumi:"longitude"`
	// The ID of the microtenant the resource is to be associated with.
	//
	// ⚠️ **WARNING:**: The attribute ``microtenantId`` is optional and requires the microtenant license and feature flag enabled for the respective tenant. The provider also supports the microtenant ID configuration via the environment variable `ZPA_MICROTENANT_ID` which is the recommended method.
	MicrotenantId *string `pulumi:"microtenantId"`
	// Name of the Service Edge Group.
	Name *string `pulumi:"name"`
	// Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
	OverrideVersionProfile *bool                         `pulumi:"overrideVersionProfile"`
	ServiceEdges           []ServiceEdgeGroupServiceEdge `pulumi:"serviceEdges"`
	// Trusted networks for this Service Edge Group. List of trusted network objects
	TrustedNetworks []ServiceEdgeGroupTrustedNetwork `pulumi:"trustedNetworks"`
	// Service Edges in this group will attempt to update to a newer version of the software during this specified day. Default value: `SUNDAY` List of valid days (i.e., Sunday, Monday)
	UpgradeDay *string `pulumi:"upgradeDay"`
	// Service Edges in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600` Integer in seconds (i..e, 66600). The integer must be greater than or equal to 0 and less than `86400`, in `15` minute intervals
	UpgradeTimeInSecs *string `pulumi:"upgradeTimeInSecs"`
	// ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
	VersionProfileId *string `pulumi:"versionProfileId"`
	// ID of the version profile.
	VersionProfileName *string `pulumi:"versionProfileName"`
	// ID of the version profile.
	VersionProfileVisibilityScope *string `pulumi:"versionProfileVisibilityScope"`
}

type ServiceEdgeGroupState struct {
	// This field controls dynamic discovery of the servers.
	CityCountry pulumi.StringPtrInput
	// This field is an array of app-connector-id only.
	CountryCode pulumi.StringPtrInput
	// Description of the Service Edge Group.
	Description pulumi.StringPtrInput
	// Whether this Service Edge Group is enabled or not. Default value: `true` Supported values: `true`, `false`
	Enabled pulumi.BoolPtrInput
	// Enable or disable public access for the Service Edge Group. Default value: `false` Supported values: `true`, `false`
	IsPublic pulumi.BoolPtrInput
	// Latitude for the Service Edge Group. Integer or decimal with values in the range of `-90` to `90`
	Latitude pulumi.StringPtrInput
	// Location for the Service Edge Group.
	Location pulumi.StringPtrInput
	// Longitude for the Service Edge Group. Integer or decimal with values in the range of `-180` to `180`
	Longitude pulumi.StringPtrInput
	// The ID of the microtenant the resource is to be associated with.
	//
	// ⚠️ **WARNING:**: The attribute ``microtenantId`` is optional and requires the microtenant license and feature flag enabled for the respective tenant. The provider also supports the microtenant ID configuration via the environment variable `ZPA_MICROTENANT_ID` which is the recommended method.
	MicrotenantId pulumi.StringPtrInput
	// Name of the Service Edge Group.
	Name pulumi.StringPtrInput
	// Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
	OverrideVersionProfile pulumi.BoolPtrInput
	ServiceEdges           ServiceEdgeGroupServiceEdgeArrayInput
	// Trusted networks for this Service Edge Group. List of trusted network objects
	TrustedNetworks ServiceEdgeGroupTrustedNetworkArrayInput
	// Service Edges in this group will attempt to update to a newer version of the software during this specified day. Default value: `SUNDAY` List of valid days (i.e., Sunday, Monday)
	UpgradeDay pulumi.StringPtrInput
	// Service Edges in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600` Integer in seconds (i..e, 66600). The integer must be greater than or equal to 0 and less than `86400`, in `15` minute intervals
	UpgradeTimeInSecs pulumi.StringPtrInput
	// ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
	VersionProfileId pulumi.StringPtrInput
	// ID of the version profile.
	VersionProfileName pulumi.StringPtrInput
	// ID of the version profile.
	VersionProfileVisibilityScope pulumi.StringPtrInput
}

func (ServiceEdgeGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEdgeGroupState)(nil)).Elem()
}

type serviceEdgeGroupArgs struct {
	// This field controls dynamic discovery of the servers.
	CityCountry *string `pulumi:"cityCountry"`
	// This field is an array of app-connector-id only.
	CountryCode *string `pulumi:"countryCode"`
	// Description of the Service Edge Group.
	Description *string `pulumi:"description"`
	// Whether this Service Edge Group is enabled or not. Default value: `true` Supported values: `true`, `false`
	Enabled *bool `pulumi:"enabled"`
	// Enable or disable public access for the Service Edge Group. Default value: `false` Supported values: `true`, `false`
	IsPublic *bool `pulumi:"isPublic"`
	// Latitude for the Service Edge Group. Integer or decimal with values in the range of `-90` to `90`
	Latitude string `pulumi:"latitude"`
	// Location for the Service Edge Group.
	Location string `pulumi:"location"`
	// Longitude for the Service Edge Group. Integer or decimal with values in the range of `-180` to `180`
	Longitude string `pulumi:"longitude"`
	// The ID of the microtenant the resource is to be associated with.
	//
	// ⚠️ **WARNING:**: The attribute ``microtenantId`` is optional and requires the microtenant license and feature flag enabled for the respective tenant. The provider also supports the microtenant ID configuration via the environment variable `ZPA_MICROTENANT_ID` which is the recommended method.
	MicrotenantId *string `pulumi:"microtenantId"`
	// Name of the Service Edge Group.
	Name *string `pulumi:"name"`
	// Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
	OverrideVersionProfile *bool                         `pulumi:"overrideVersionProfile"`
	ServiceEdges           []ServiceEdgeGroupServiceEdge `pulumi:"serviceEdges"`
	// Trusted networks for this Service Edge Group. List of trusted network objects
	TrustedNetworks []ServiceEdgeGroupTrustedNetwork `pulumi:"trustedNetworks"`
	// Service Edges in this group will attempt to update to a newer version of the software during this specified day. Default value: `SUNDAY` List of valid days (i.e., Sunday, Monday)
	UpgradeDay *string `pulumi:"upgradeDay"`
	// Service Edges in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600` Integer in seconds (i..e, 66600). The integer must be greater than or equal to 0 and less than `86400`, in `15` minute intervals
	UpgradeTimeInSecs *string `pulumi:"upgradeTimeInSecs"`
	// ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
	VersionProfileId *string `pulumi:"versionProfileId"`
	// ID of the version profile.
	VersionProfileName *string `pulumi:"versionProfileName"`
}

// The set of arguments for constructing a ServiceEdgeGroup resource.
type ServiceEdgeGroupArgs struct {
	// This field controls dynamic discovery of the servers.
	CityCountry pulumi.StringPtrInput
	// This field is an array of app-connector-id only.
	CountryCode pulumi.StringPtrInput
	// Description of the Service Edge Group.
	Description pulumi.StringPtrInput
	// Whether this Service Edge Group is enabled or not. Default value: `true` Supported values: `true`, `false`
	Enabled pulumi.BoolPtrInput
	// Enable or disable public access for the Service Edge Group. Default value: `false` Supported values: `true`, `false`
	IsPublic pulumi.BoolPtrInput
	// Latitude for the Service Edge Group. Integer or decimal with values in the range of `-90` to `90`
	Latitude pulumi.StringInput
	// Location for the Service Edge Group.
	Location pulumi.StringInput
	// Longitude for the Service Edge Group. Integer or decimal with values in the range of `-180` to `180`
	Longitude pulumi.StringInput
	// The ID of the microtenant the resource is to be associated with.
	//
	// ⚠️ **WARNING:**: The attribute ``microtenantId`` is optional and requires the microtenant license and feature flag enabled for the respective tenant. The provider also supports the microtenant ID configuration via the environment variable `ZPA_MICROTENANT_ID` which is the recommended method.
	MicrotenantId pulumi.StringPtrInput
	// Name of the Service Edge Group.
	Name pulumi.StringPtrInput
	// Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
	OverrideVersionProfile pulumi.BoolPtrInput
	ServiceEdges           ServiceEdgeGroupServiceEdgeArrayInput
	// Trusted networks for this Service Edge Group. List of trusted network objects
	TrustedNetworks ServiceEdgeGroupTrustedNetworkArrayInput
	// Service Edges in this group will attempt to update to a newer version of the software during this specified day. Default value: `SUNDAY` List of valid days (i.e., Sunday, Monday)
	UpgradeDay pulumi.StringPtrInput
	// Service Edges in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600` Integer in seconds (i..e, 66600). The integer must be greater than or equal to 0 and less than `86400`, in `15` minute intervals
	UpgradeTimeInSecs pulumi.StringPtrInput
	// ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
	VersionProfileId pulumi.StringPtrInput
	// ID of the version profile.
	VersionProfileName pulumi.StringPtrInput
}

func (ServiceEdgeGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEdgeGroupArgs)(nil)).Elem()
}

type ServiceEdgeGroupInput interface {
	pulumi.Input

	ToServiceEdgeGroupOutput() ServiceEdgeGroupOutput
	ToServiceEdgeGroupOutputWithContext(ctx context.Context) ServiceEdgeGroupOutput
}

func (*ServiceEdgeGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEdgeGroup)(nil)).Elem()
}

func (i *ServiceEdgeGroup) ToServiceEdgeGroupOutput() ServiceEdgeGroupOutput {
	return i.ToServiceEdgeGroupOutputWithContext(context.Background())
}

func (i *ServiceEdgeGroup) ToServiceEdgeGroupOutputWithContext(ctx context.Context) ServiceEdgeGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEdgeGroupOutput)
}

// ServiceEdgeGroupArrayInput is an input type that accepts ServiceEdgeGroupArray and ServiceEdgeGroupArrayOutput values.
// You can construct a concrete instance of `ServiceEdgeGroupArrayInput` via:
//
//	ServiceEdgeGroupArray{ ServiceEdgeGroupArgs{...} }
type ServiceEdgeGroupArrayInput interface {
	pulumi.Input

	ToServiceEdgeGroupArrayOutput() ServiceEdgeGroupArrayOutput
	ToServiceEdgeGroupArrayOutputWithContext(context.Context) ServiceEdgeGroupArrayOutput
}

type ServiceEdgeGroupArray []ServiceEdgeGroupInput

func (ServiceEdgeGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEdgeGroup)(nil)).Elem()
}

func (i ServiceEdgeGroupArray) ToServiceEdgeGroupArrayOutput() ServiceEdgeGroupArrayOutput {
	return i.ToServiceEdgeGroupArrayOutputWithContext(context.Background())
}

func (i ServiceEdgeGroupArray) ToServiceEdgeGroupArrayOutputWithContext(ctx context.Context) ServiceEdgeGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEdgeGroupArrayOutput)
}

// ServiceEdgeGroupMapInput is an input type that accepts ServiceEdgeGroupMap and ServiceEdgeGroupMapOutput values.
// You can construct a concrete instance of `ServiceEdgeGroupMapInput` via:
//
//	ServiceEdgeGroupMap{ "key": ServiceEdgeGroupArgs{...} }
type ServiceEdgeGroupMapInput interface {
	pulumi.Input

	ToServiceEdgeGroupMapOutput() ServiceEdgeGroupMapOutput
	ToServiceEdgeGroupMapOutputWithContext(context.Context) ServiceEdgeGroupMapOutput
}

type ServiceEdgeGroupMap map[string]ServiceEdgeGroupInput

func (ServiceEdgeGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEdgeGroup)(nil)).Elem()
}

func (i ServiceEdgeGroupMap) ToServiceEdgeGroupMapOutput() ServiceEdgeGroupMapOutput {
	return i.ToServiceEdgeGroupMapOutputWithContext(context.Background())
}

func (i ServiceEdgeGroupMap) ToServiceEdgeGroupMapOutputWithContext(ctx context.Context) ServiceEdgeGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEdgeGroupMapOutput)
}

type ServiceEdgeGroupOutput struct{ *pulumi.OutputState }

func (ServiceEdgeGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEdgeGroup)(nil)).Elem()
}

func (o ServiceEdgeGroupOutput) ToServiceEdgeGroupOutput() ServiceEdgeGroupOutput {
	return o
}

func (o ServiceEdgeGroupOutput) ToServiceEdgeGroupOutputWithContext(ctx context.Context) ServiceEdgeGroupOutput {
	return o
}

// This field controls dynamic discovery of the servers.
func (o ServiceEdgeGroupOutput) CityCountry() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringOutput { return v.CityCountry }).(pulumi.StringOutput)
}

// This field is an array of app-connector-id only.
func (o ServiceEdgeGroupOutput) CountryCode() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringOutput { return v.CountryCode }).(pulumi.StringOutput)
}

// Description of the Service Edge Group.
func (o ServiceEdgeGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether this Service Edge Group is enabled or not. Default value: `true` Supported values: `true`, `false`
func (o ServiceEdgeGroupOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Enable or disable public access for the Service Edge Group. Default value: `false` Supported values: `true`, `false`
func (o ServiceEdgeGroupOutput) IsPublic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.BoolPtrOutput { return v.IsPublic }).(pulumi.BoolPtrOutput)
}

// Latitude for the Service Edge Group. Integer or decimal with values in the range of `-90` to `90`
func (o ServiceEdgeGroupOutput) Latitude() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringOutput { return v.Latitude }).(pulumi.StringOutput)
}

// Location for the Service Edge Group.
func (o ServiceEdgeGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Longitude for the Service Edge Group. Integer or decimal with values in the range of `-180` to `180`
func (o ServiceEdgeGroupOutput) Longitude() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringOutput { return v.Longitude }).(pulumi.StringOutput)
}

// The ID of the microtenant the resource is to be associated with.
//
// ⚠️ **WARNING:**: The attribute “microtenantId“ is optional and requires the microtenant license and feature flag enabled for the respective tenant. The provider also supports the microtenant ID configuration via the environment variable `ZPA_MICROTENANT_ID` which is the recommended method.
func (o ServiceEdgeGroupOutput) MicrotenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringOutput { return v.MicrotenantId }).(pulumi.StringOutput)
}

// Name of the Service Edge Group.
func (o ServiceEdgeGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
func (o ServiceEdgeGroupOutput) OverrideVersionProfile() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.BoolPtrOutput { return v.OverrideVersionProfile }).(pulumi.BoolPtrOutput)
}

func (o ServiceEdgeGroupOutput) ServiceEdges() ServiceEdgeGroupServiceEdgeArrayOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) ServiceEdgeGroupServiceEdgeArrayOutput { return v.ServiceEdges }).(ServiceEdgeGroupServiceEdgeArrayOutput)
}

// Trusted networks for this Service Edge Group. List of trusted network objects
func (o ServiceEdgeGroupOutput) TrustedNetworks() ServiceEdgeGroupTrustedNetworkArrayOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) ServiceEdgeGroupTrustedNetworkArrayOutput { return v.TrustedNetworks }).(ServiceEdgeGroupTrustedNetworkArrayOutput)
}

// Service Edges in this group will attempt to update to a newer version of the software during this specified day. Default value: `SUNDAY` List of valid days (i.e., Sunday, Monday)
func (o ServiceEdgeGroupOutput) UpgradeDay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringPtrOutput { return v.UpgradeDay }).(pulumi.StringPtrOutput)
}

// Service Edges in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600` Integer in seconds (i..e, 66600). The integer must be greater than or equal to 0 and less than `86400`, in `15` minute intervals
func (o ServiceEdgeGroupOutput) UpgradeTimeInSecs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringPtrOutput { return v.UpgradeTimeInSecs }).(pulumi.StringPtrOutput)
}

// ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
func (o ServiceEdgeGroupOutput) VersionProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringOutput { return v.VersionProfileId }).(pulumi.StringOutput)
}

// ID of the version profile.
func (o ServiceEdgeGroupOutput) VersionProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringOutput { return v.VersionProfileName }).(pulumi.StringOutput)
}

// ID of the version profile.
func (o ServiceEdgeGroupOutput) VersionProfileVisibilityScope() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringOutput { return v.VersionProfileVisibilityScope }).(pulumi.StringOutput)
}

type ServiceEdgeGroupArrayOutput struct{ *pulumi.OutputState }

func (ServiceEdgeGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEdgeGroup)(nil)).Elem()
}

func (o ServiceEdgeGroupArrayOutput) ToServiceEdgeGroupArrayOutput() ServiceEdgeGroupArrayOutput {
	return o
}

func (o ServiceEdgeGroupArrayOutput) ToServiceEdgeGroupArrayOutputWithContext(ctx context.Context) ServiceEdgeGroupArrayOutput {
	return o
}

func (o ServiceEdgeGroupArrayOutput) Index(i pulumi.IntInput) ServiceEdgeGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceEdgeGroup {
		return vs[0].([]*ServiceEdgeGroup)[vs[1].(int)]
	}).(ServiceEdgeGroupOutput)
}

type ServiceEdgeGroupMapOutput struct{ *pulumi.OutputState }

func (ServiceEdgeGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEdgeGroup)(nil)).Elem()
}

func (o ServiceEdgeGroupMapOutput) ToServiceEdgeGroupMapOutput() ServiceEdgeGroupMapOutput {
	return o
}

func (o ServiceEdgeGroupMapOutput) ToServiceEdgeGroupMapOutputWithContext(ctx context.Context) ServiceEdgeGroupMapOutput {
	return o
}

func (o ServiceEdgeGroupMapOutput) MapIndex(k pulumi.StringInput) ServiceEdgeGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceEdgeGroup {
		return vs[0].(map[string]*ServiceEdgeGroup)[vs[1].(string)]
	}).(ServiceEdgeGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEdgeGroupInput)(nil)).Elem(), &ServiceEdgeGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEdgeGroupArrayInput)(nil)).Elem(), ServiceEdgeGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEdgeGroupMapInput)(nil)).Elem(), ServiceEdgeGroupMap{})
	pulumi.RegisterOutputType(ServiceEdgeGroupOutput{})
	pulumi.RegisterOutputType(ServiceEdgeGroupArrayOutput{})
	pulumi.RegisterOutputType(ServiceEdgeGroupMapOutput{})
}
