// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// * [Official documentation](https://help.zscaler.com/zpa/about-zpa-private-service-edge-groups)
// * [API documentation](https://help.zscaler.com/zpa/configuring-zpa-private-service-edge-groups-using-api)
//
// The **zpa_service_edge_group** resource creates a service edge group in the Zscaler Private Access cloud. This resource can then be referenced in a service edge connector.
//
// ## Example Usage
//
// ### Using Version Profile Name
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
//			// ZPA Service Edge Group resource - No Trusted Network
//			_, err := zpa.NewServiceEdgeGroup(ctx, "serviceEdgeGroupNyc", &zpa.ServiceEdgeGroupArgs{
//				Description:       pulumi.String("Service Edge Group in New York"),
//				Enabled:           pulumi.Bool(true),
//				IsPublic:          pulumi.Bool(true),
//				UpgradeDay:        pulumi.String("SUNDAY"),
//				UpgradeTimeInSecs: pulumi.String("66600"),
//				Latitude:          pulumi.String("40.7128"),
//				Longitude:         pulumi.String("-73.935242"),
//				Location:          pulumi.String("New York, NY, USA"),
//				VersionProfileId:  pulumi.Any(data.Zpa_customer_version_profile.This.Id),
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
// ### Using Version Profile ID
//
//	data "getCustomerVersionProfile" "this" {
//	  name = "New Release"
//	}
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

	CityCountry pulumi.StringOutput `pulumi:"cityCountry"`
	CountryCode pulumi.StringOutput `pulumi:"countryCode"`
	// Description of the Service Edge Group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether this Service Edge Group is enabled or not.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// If enabled, allows ZPA Private Service Edge Groups within the specified distance to be prioritized over a closer ZPA
	// Public Service Edge.
	GraceDistanceEnabled pulumi.BoolOutput `pulumi:"graceDistanceEnabled"`
	// Indicates the maximum distance in miles or kilometers to ZPA Private Service Edge groups that would override a ZPA
	// Public Service Edge
	GraceDistanceValue pulumi.StringOutput `pulumi:"graceDistanceValue"`
	// Indicates the grace distance unit of measure in miles or kilometers. This value is only required if graceDistanceValue
	// is set to true
	GraceDistanceValueUnit pulumi.StringOutput `pulumi:"graceDistanceValueUnit"`
	// Enable or disable public access for the Service Edge Group.
	IsPublic pulumi.BoolPtrOutput `pulumi:"isPublic"`
	// Latitude for the Service Edge Group.
	Latitude pulumi.StringOutput `pulumi:"latitude"`
	// Location for the Service Edge Group.
	Location pulumi.StringOutput `pulumi:"location"`
	// Longitude for the Service Edge Group.
	Longitude     pulumi.StringOutput `pulumi:"longitude"`
	MicrotenantId pulumi.StringOutput `pulumi:"microtenantId"`
	// Name of the Service Edge Group.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether the default version profile of the App Connector Group is applied or overridden.
	OverrideVersionProfile pulumi.BoolPtrOutput `pulumi:"overrideVersionProfile"`
	// WARNING: Service edge membership is managed externally. Specifying this field will enforce exact membership
	// matching.This field will be deprecated in future versions
	ServiceEdges ServiceEdgeGroupServiceEdgesOutput `pulumi:"serviceEdges"`
	// List of trusted network IDs.
	TrustedNetworks ServiceEdgeGroupTrustedNetworkArrayOutput `pulumi:"trustedNetworks"`
	// Service Edges in this group will attempt to update to a newer version of the software during this specified day.
	UpgradeDay pulumi.StringPtrOutput `pulumi:"upgradeDay"`
	// Service Edges in this group will attempt to update to a newer version of the software during this specified time.
	UpgradeTimeInSecs pulumi.StringPtrOutput `pulumi:"upgradeTimeInSecs"`
	UseInDrMode       pulumi.BoolOutput      `pulumi:"useInDrMode"`
	// ID of the version profile.
	VersionProfileId pulumi.StringOutput `pulumi:"versionProfileId"`
	// Name of the version profile. To learn more, see Version Profile Use Cases. This value is required, if the value for
	// overrideVersionProfile is set to true
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
	CityCountry *string `pulumi:"cityCountry"`
	CountryCode *string `pulumi:"countryCode"`
	// Description of the Service Edge Group.
	Description *string `pulumi:"description"`
	// Whether this Service Edge Group is enabled or not.
	Enabled *bool `pulumi:"enabled"`
	// If enabled, allows ZPA Private Service Edge Groups within the specified distance to be prioritized over a closer ZPA
	// Public Service Edge.
	GraceDistanceEnabled *bool `pulumi:"graceDistanceEnabled"`
	// Indicates the maximum distance in miles or kilometers to ZPA Private Service Edge groups that would override a ZPA
	// Public Service Edge
	GraceDistanceValue *string `pulumi:"graceDistanceValue"`
	// Indicates the grace distance unit of measure in miles or kilometers. This value is only required if graceDistanceValue
	// is set to true
	GraceDistanceValueUnit *string `pulumi:"graceDistanceValueUnit"`
	// Enable or disable public access for the Service Edge Group.
	IsPublic *bool `pulumi:"isPublic"`
	// Latitude for the Service Edge Group.
	Latitude *string `pulumi:"latitude"`
	// Location for the Service Edge Group.
	Location *string `pulumi:"location"`
	// Longitude for the Service Edge Group.
	Longitude     *string `pulumi:"longitude"`
	MicrotenantId *string `pulumi:"microtenantId"`
	// Name of the Service Edge Group.
	Name *string `pulumi:"name"`
	// Whether the default version profile of the App Connector Group is applied or overridden.
	OverrideVersionProfile *bool `pulumi:"overrideVersionProfile"`
	// WARNING: Service edge membership is managed externally. Specifying this field will enforce exact membership
	// matching.This field will be deprecated in future versions
	ServiceEdges *ServiceEdgeGroupServiceEdges `pulumi:"serviceEdges"`
	// List of trusted network IDs.
	TrustedNetworks []ServiceEdgeGroupTrustedNetwork `pulumi:"trustedNetworks"`
	// Service Edges in this group will attempt to update to a newer version of the software during this specified day.
	UpgradeDay *string `pulumi:"upgradeDay"`
	// Service Edges in this group will attempt to update to a newer version of the software during this specified time.
	UpgradeTimeInSecs *string `pulumi:"upgradeTimeInSecs"`
	UseInDrMode       *bool   `pulumi:"useInDrMode"`
	// ID of the version profile.
	VersionProfileId *string `pulumi:"versionProfileId"`
	// Name of the version profile. To learn more, see Version Profile Use Cases. This value is required, if the value for
	// overrideVersionProfile is set to true
	VersionProfileName *string `pulumi:"versionProfileName"`
	// ID of the version profile.
	VersionProfileVisibilityScope *string `pulumi:"versionProfileVisibilityScope"`
}

type ServiceEdgeGroupState struct {
	CityCountry pulumi.StringPtrInput
	CountryCode pulumi.StringPtrInput
	// Description of the Service Edge Group.
	Description pulumi.StringPtrInput
	// Whether this Service Edge Group is enabled or not.
	Enabled pulumi.BoolPtrInput
	// If enabled, allows ZPA Private Service Edge Groups within the specified distance to be prioritized over a closer ZPA
	// Public Service Edge.
	GraceDistanceEnabled pulumi.BoolPtrInput
	// Indicates the maximum distance in miles or kilometers to ZPA Private Service Edge groups that would override a ZPA
	// Public Service Edge
	GraceDistanceValue pulumi.StringPtrInput
	// Indicates the grace distance unit of measure in miles or kilometers. This value is only required if graceDistanceValue
	// is set to true
	GraceDistanceValueUnit pulumi.StringPtrInput
	// Enable or disable public access for the Service Edge Group.
	IsPublic pulumi.BoolPtrInput
	// Latitude for the Service Edge Group.
	Latitude pulumi.StringPtrInput
	// Location for the Service Edge Group.
	Location pulumi.StringPtrInput
	// Longitude for the Service Edge Group.
	Longitude     pulumi.StringPtrInput
	MicrotenantId pulumi.StringPtrInput
	// Name of the Service Edge Group.
	Name pulumi.StringPtrInput
	// Whether the default version profile of the App Connector Group is applied or overridden.
	OverrideVersionProfile pulumi.BoolPtrInput
	// WARNING: Service edge membership is managed externally. Specifying this field will enforce exact membership
	// matching.This field will be deprecated in future versions
	ServiceEdges ServiceEdgeGroupServiceEdgesPtrInput
	// List of trusted network IDs.
	TrustedNetworks ServiceEdgeGroupTrustedNetworkArrayInput
	// Service Edges in this group will attempt to update to a newer version of the software during this specified day.
	UpgradeDay pulumi.StringPtrInput
	// Service Edges in this group will attempt to update to a newer version of the software during this specified time.
	UpgradeTimeInSecs pulumi.StringPtrInput
	UseInDrMode       pulumi.BoolPtrInput
	// ID of the version profile.
	VersionProfileId pulumi.StringPtrInput
	// Name of the version profile. To learn more, see Version Profile Use Cases. This value is required, if the value for
	// overrideVersionProfile is set to true
	VersionProfileName pulumi.StringPtrInput
	// ID of the version profile.
	VersionProfileVisibilityScope pulumi.StringPtrInput
}

func (ServiceEdgeGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEdgeGroupState)(nil)).Elem()
}

type serviceEdgeGroupArgs struct {
	CityCountry *string `pulumi:"cityCountry"`
	CountryCode *string `pulumi:"countryCode"`
	// Description of the Service Edge Group.
	Description *string `pulumi:"description"`
	// Whether this Service Edge Group is enabled or not.
	Enabled *bool `pulumi:"enabled"`
	// If enabled, allows ZPA Private Service Edge Groups within the specified distance to be prioritized over a closer ZPA
	// Public Service Edge.
	GraceDistanceEnabled *bool `pulumi:"graceDistanceEnabled"`
	// Indicates the maximum distance in miles or kilometers to ZPA Private Service Edge groups that would override a ZPA
	// Public Service Edge
	GraceDistanceValue *string `pulumi:"graceDistanceValue"`
	// Indicates the grace distance unit of measure in miles or kilometers. This value is only required if graceDistanceValue
	// is set to true
	GraceDistanceValueUnit *string `pulumi:"graceDistanceValueUnit"`
	// Enable or disable public access for the Service Edge Group.
	IsPublic *bool `pulumi:"isPublic"`
	// Latitude for the Service Edge Group.
	Latitude string `pulumi:"latitude"`
	// Location for the Service Edge Group.
	Location string `pulumi:"location"`
	// Longitude for the Service Edge Group.
	Longitude     string  `pulumi:"longitude"`
	MicrotenantId *string `pulumi:"microtenantId"`
	// Name of the Service Edge Group.
	Name *string `pulumi:"name"`
	// Whether the default version profile of the App Connector Group is applied or overridden.
	OverrideVersionProfile *bool `pulumi:"overrideVersionProfile"`
	// WARNING: Service edge membership is managed externally. Specifying this field will enforce exact membership
	// matching.This field will be deprecated in future versions
	ServiceEdges *ServiceEdgeGroupServiceEdges `pulumi:"serviceEdges"`
	// List of trusted network IDs.
	TrustedNetworks []ServiceEdgeGroupTrustedNetwork `pulumi:"trustedNetworks"`
	// Service Edges in this group will attempt to update to a newer version of the software during this specified day.
	UpgradeDay *string `pulumi:"upgradeDay"`
	// Service Edges in this group will attempt to update to a newer version of the software during this specified time.
	UpgradeTimeInSecs *string `pulumi:"upgradeTimeInSecs"`
	UseInDrMode       *bool   `pulumi:"useInDrMode"`
	// ID of the version profile.
	VersionProfileId *string `pulumi:"versionProfileId"`
	// Name of the version profile. To learn more, see Version Profile Use Cases. This value is required, if the value for
	// overrideVersionProfile is set to true
	VersionProfileName *string `pulumi:"versionProfileName"`
	// ID of the version profile.
	VersionProfileVisibilityScope *string `pulumi:"versionProfileVisibilityScope"`
}

// The set of arguments for constructing a ServiceEdgeGroup resource.
type ServiceEdgeGroupArgs struct {
	CityCountry pulumi.StringPtrInput
	CountryCode pulumi.StringPtrInput
	// Description of the Service Edge Group.
	Description pulumi.StringPtrInput
	// Whether this Service Edge Group is enabled or not.
	Enabled pulumi.BoolPtrInput
	// If enabled, allows ZPA Private Service Edge Groups within the specified distance to be prioritized over a closer ZPA
	// Public Service Edge.
	GraceDistanceEnabled pulumi.BoolPtrInput
	// Indicates the maximum distance in miles or kilometers to ZPA Private Service Edge groups that would override a ZPA
	// Public Service Edge
	GraceDistanceValue pulumi.StringPtrInput
	// Indicates the grace distance unit of measure in miles or kilometers. This value is only required if graceDistanceValue
	// is set to true
	GraceDistanceValueUnit pulumi.StringPtrInput
	// Enable or disable public access for the Service Edge Group.
	IsPublic pulumi.BoolPtrInput
	// Latitude for the Service Edge Group.
	Latitude pulumi.StringInput
	// Location for the Service Edge Group.
	Location pulumi.StringInput
	// Longitude for the Service Edge Group.
	Longitude     pulumi.StringInput
	MicrotenantId pulumi.StringPtrInput
	// Name of the Service Edge Group.
	Name pulumi.StringPtrInput
	// Whether the default version profile of the App Connector Group is applied or overridden.
	OverrideVersionProfile pulumi.BoolPtrInput
	// WARNING: Service edge membership is managed externally. Specifying this field will enforce exact membership
	// matching.This field will be deprecated in future versions
	ServiceEdges ServiceEdgeGroupServiceEdgesPtrInput
	// List of trusted network IDs.
	TrustedNetworks ServiceEdgeGroupTrustedNetworkArrayInput
	// Service Edges in this group will attempt to update to a newer version of the software during this specified day.
	UpgradeDay pulumi.StringPtrInput
	// Service Edges in this group will attempt to update to a newer version of the software during this specified time.
	UpgradeTimeInSecs pulumi.StringPtrInput
	UseInDrMode       pulumi.BoolPtrInput
	// ID of the version profile.
	VersionProfileId pulumi.StringPtrInput
	// Name of the version profile. To learn more, see Version Profile Use Cases. This value is required, if the value for
	// overrideVersionProfile is set to true
	VersionProfileName pulumi.StringPtrInput
	// ID of the version profile.
	VersionProfileVisibilityScope pulumi.StringPtrInput
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

func (o ServiceEdgeGroupOutput) CityCountry() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringOutput { return v.CityCountry }).(pulumi.StringOutput)
}

func (o ServiceEdgeGroupOutput) CountryCode() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringOutput { return v.CountryCode }).(pulumi.StringOutput)
}

// Description of the Service Edge Group.
func (o ServiceEdgeGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether this Service Edge Group is enabled or not.
func (o ServiceEdgeGroupOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// If enabled, allows ZPA Private Service Edge Groups within the specified distance to be prioritized over a closer ZPA
// Public Service Edge.
func (o ServiceEdgeGroupOutput) GraceDistanceEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.BoolOutput { return v.GraceDistanceEnabled }).(pulumi.BoolOutput)
}

// Indicates the maximum distance in miles or kilometers to ZPA Private Service Edge groups that would override a ZPA
// Public Service Edge
func (o ServiceEdgeGroupOutput) GraceDistanceValue() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringOutput { return v.GraceDistanceValue }).(pulumi.StringOutput)
}

// Indicates the grace distance unit of measure in miles or kilometers. This value is only required if graceDistanceValue
// is set to true
func (o ServiceEdgeGroupOutput) GraceDistanceValueUnit() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringOutput { return v.GraceDistanceValueUnit }).(pulumi.StringOutput)
}

// Enable or disable public access for the Service Edge Group.
func (o ServiceEdgeGroupOutput) IsPublic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.BoolPtrOutput { return v.IsPublic }).(pulumi.BoolPtrOutput)
}

// Latitude for the Service Edge Group.
func (o ServiceEdgeGroupOutput) Latitude() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringOutput { return v.Latitude }).(pulumi.StringOutput)
}

// Location for the Service Edge Group.
func (o ServiceEdgeGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Longitude for the Service Edge Group.
func (o ServiceEdgeGroupOutput) Longitude() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringOutput { return v.Longitude }).(pulumi.StringOutput)
}

func (o ServiceEdgeGroupOutput) MicrotenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringOutput { return v.MicrotenantId }).(pulumi.StringOutput)
}

// Name of the Service Edge Group.
func (o ServiceEdgeGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Whether the default version profile of the App Connector Group is applied or overridden.
func (o ServiceEdgeGroupOutput) OverrideVersionProfile() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.BoolPtrOutput { return v.OverrideVersionProfile }).(pulumi.BoolPtrOutput)
}

// WARNING: Service edge membership is managed externally. Specifying this field will enforce exact membership
// matching.This field will be deprecated in future versions
func (o ServiceEdgeGroupOutput) ServiceEdges() ServiceEdgeGroupServiceEdgesOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) ServiceEdgeGroupServiceEdgesOutput { return v.ServiceEdges }).(ServiceEdgeGroupServiceEdgesOutput)
}

// List of trusted network IDs.
func (o ServiceEdgeGroupOutput) TrustedNetworks() ServiceEdgeGroupTrustedNetworkArrayOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) ServiceEdgeGroupTrustedNetworkArrayOutput { return v.TrustedNetworks }).(ServiceEdgeGroupTrustedNetworkArrayOutput)
}

// Service Edges in this group will attempt to update to a newer version of the software during this specified day.
func (o ServiceEdgeGroupOutput) UpgradeDay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringPtrOutput { return v.UpgradeDay }).(pulumi.StringPtrOutput)
}

// Service Edges in this group will attempt to update to a newer version of the software during this specified time.
func (o ServiceEdgeGroupOutput) UpgradeTimeInSecs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringPtrOutput { return v.UpgradeTimeInSecs }).(pulumi.StringPtrOutput)
}

func (o ServiceEdgeGroupOutput) UseInDrMode() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.BoolOutput { return v.UseInDrMode }).(pulumi.BoolOutput)
}

// ID of the version profile.
func (o ServiceEdgeGroupOutput) VersionProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeGroup) pulumi.StringOutput { return v.VersionProfileId }).(pulumi.StringOutput)
}

// Name of the version profile. To learn more, see Version Profile Use Cases. This value is required, if the value for
// overrideVersionProfile is set to true
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
