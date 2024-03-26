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
//			// Create a App Connector Group
//			_, err := zpa.NewConnectorGroup(ctx, "example", &zpa.ConnectorGroupArgs{
//				CityCountry:            pulumi.String("San Jose, CA"),
//				CountryCode:            pulumi.String("US"),
//				Description:            pulumi.String("Example"),
//				DnsQueryType:           pulumi.String("IPV4_IPV6"),
//				Enabled:                pulumi.Bool(true),
//				Latitude:               pulumi.String("37.338"),
//				Location:               pulumi.String("San Jose, CA, US"),
//				Longitude:              pulumi.String("-121.8863"),
//				OverrideVersionProfile: pulumi.Bool(true),
//				UpgradeDay:             pulumi.String("SUNDAY"),
//				UpgradeTimeInSecs:      pulumi.String("66600"),
//				UseInDrMode:            pulumi.Bool(true),
//				VersionProfileName:     pulumi.String("New Release"),
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
// App Connector Group can be imported by using `<APP CONNECTOR GROUP ID>` or `<APP CONNECTOR GROUP NAME>`as the import ID.
//
// ```sh
// $ pulumi import zpa:index/connectorGroup:ConnectorGroup example <app_connector_group_id>
// ```
//
// or
//
// ```sh
// $ pulumi import zpa:index/connectorGroup:ConnectorGroup example <app_connector_group_name>
// ```
type ConnectorGroup struct {
	pulumi.CustomResourceState

	// Whether Double Encryption is enabled or disabled for the app. i.e ``"San Jose, US"``
	CityCountry pulumi.StringOutput `pulumi:"cityCountry"`
	// i.e ``"US"``, ``"CA"``
	CountryCode pulumi.StringOutput `pulumi:"countryCode"`
	// Description of the App Connector Group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Supported values are:
	DnsQueryType pulumi.StringPtrOutput `pulumi:"dnsQueryType"`
	// Whether this App Connector Group is enabled or not. Default value: `true`. Supported values: `true`, `false`
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Latitude of the App Connector Group. Integer or decimal. With values in the range of `-90` to `90`
	Latitude pulumi.StringOutput `pulumi:"latitude"`
	// Location of the App Connector Group. i.e ``"San Jose, CA, USA"``
	Location pulumi.StringOutput `pulumi:"location"`
	// Longitude of the App Connector Group. Integer or decimal. With values in the range of `-180` to `180`
	Longitude            pulumi.StringOutput `pulumi:"longitude"`
	LssAppConnectorGroup pulumi.BoolOutput   `pulumi:"lssAppConnectorGroup"`
	// The ID of the microtenant the resource is to be associated with.
	//
	// ⚠️ **WARNING:**: The attribute ``microtenantId`` is optional and requires the microtenant license and feature flag enabled for the respective tenant. The provider also supports the microtenant ID configuration via the environment variable `ZPA_MICROTENANT_ID` which is the recommended method.
	MicrotenantId pulumi.StringOutput `pulumi:"microtenantId"`
	// Name of the App Connector Group.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
	OverrideVersionProfile pulumi.BoolOutput `pulumi:"overrideVersionProfile"`
	// Supported values: `true`, `false`
	PraEnabled pulumi.BoolOutput `pulumi:"praEnabled"`
	// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
	TcpQuickAckApp pulumi.BoolOutput `pulumi:"tcpQuickAckApp"`
	// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
	TcpQuickAckAssistant pulumi.BoolOutput `pulumi:"tcpQuickAckAssistant"`
	// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
	TcpQuickAckReadAssistant pulumi.BoolOutput `pulumi:"tcpQuickAckReadAssistant"`
	// App Connectors in this group will attempt to update to a newer version of the software during this specified day i.e ``SUNDAY``
	UpgradeDay pulumi.StringPtrOutput `pulumi:"upgradeDay"`
	// App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600`. Integer in seconds (i.e., `-66600`). The integer should be greater than or equal to `0` and less than `86400`, in `15` minute intervals
	UpgradeTimeInSecs pulumi.StringPtrOutput `pulumi:"upgradeTimeInSecs"`
	// Supported values: `true`, `false`
	UseInDrMode pulumi.BoolOutput `pulumi:"useInDrMode"`
	// ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
	VersionProfileId pulumi.StringOutput `pulumi:"versionProfileId"`
	// Name of the version profile. To learn more, see Version Profile Use Cases. This value is required, if the value for
	// overrideVersionProfile is set to true
	VersionProfileName pulumi.StringOutput `pulumi:"versionProfileName"`
	// Supported values: `true`, `false`
	WafDisabled pulumi.BoolOutput `pulumi:"wafDisabled"`
}

// NewConnectorGroup registers a new resource with the given unique name, arguments, and options.
func NewConnectorGroup(ctx *pulumi.Context,
	name string, args *ConnectorGroupArgs, opts ...pulumi.ResourceOption) (*ConnectorGroup, error) {
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
	var resource ConnectorGroup
	err := ctx.RegisterResource("zpa:index/connectorGroup:ConnectorGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectorGroup gets an existing ConnectorGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectorGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectorGroupState, opts ...pulumi.ResourceOption) (*ConnectorGroup, error) {
	var resource ConnectorGroup
	err := ctx.ReadResource("zpa:index/connectorGroup:ConnectorGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectorGroup resources.
type connectorGroupState struct {
	// Whether Double Encryption is enabled or disabled for the app. i.e ``"San Jose, US"``
	CityCountry *string `pulumi:"cityCountry"`
	// i.e ``"US"``, ``"CA"``
	CountryCode *string `pulumi:"countryCode"`
	// Description of the App Connector Group.
	Description *string `pulumi:"description"`
	// Supported values are:
	DnsQueryType *string `pulumi:"dnsQueryType"`
	// Whether this App Connector Group is enabled or not. Default value: `true`. Supported values: `true`, `false`
	Enabled *bool `pulumi:"enabled"`
	// Latitude of the App Connector Group. Integer or decimal. With values in the range of `-90` to `90`
	Latitude *string `pulumi:"latitude"`
	// Location of the App Connector Group. i.e ``"San Jose, CA, USA"``
	Location *string `pulumi:"location"`
	// Longitude of the App Connector Group. Integer or decimal. With values in the range of `-180` to `180`
	Longitude            *string `pulumi:"longitude"`
	LssAppConnectorGroup *bool   `pulumi:"lssAppConnectorGroup"`
	// The ID of the microtenant the resource is to be associated with.
	//
	// ⚠️ **WARNING:**: The attribute ``microtenantId`` is optional and requires the microtenant license and feature flag enabled for the respective tenant. The provider also supports the microtenant ID configuration via the environment variable `ZPA_MICROTENANT_ID` which is the recommended method.
	MicrotenantId *string `pulumi:"microtenantId"`
	// Name of the App Connector Group.
	Name *string `pulumi:"name"`
	// Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
	OverrideVersionProfile *bool `pulumi:"overrideVersionProfile"`
	// Supported values: `true`, `false`
	PraEnabled *bool `pulumi:"praEnabled"`
	// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
	TcpQuickAckApp *bool `pulumi:"tcpQuickAckApp"`
	// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
	TcpQuickAckAssistant *bool `pulumi:"tcpQuickAckAssistant"`
	// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
	TcpQuickAckReadAssistant *bool `pulumi:"tcpQuickAckReadAssistant"`
	// App Connectors in this group will attempt to update to a newer version of the software during this specified day i.e ``SUNDAY``
	UpgradeDay *string `pulumi:"upgradeDay"`
	// App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600`. Integer in seconds (i.e., `-66600`). The integer should be greater than or equal to `0` and less than `86400`, in `15` minute intervals
	UpgradeTimeInSecs *string `pulumi:"upgradeTimeInSecs"`
	// Supported values: `true`, `false`
	UseInDrMode *bool `pulumi:"useInDrMode"`
	// ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
	VersionProfileId *string `pulumi:"versionProfileId"`
	// Name of the version profile. To learn more, see Version Profile Use Cases. This value is required, if the value for
	// overrideVersionProfile is set to true
	VersionProfileName *string `pulumi:"versionProfileName"`
	// Supported values: `true`, `false`
	WafDisabled *bool `pulumi:"wafDisabled"`
}

type ConnectorGroupState struct {
	// Whether Double Encryption is enabled or disabled for the app. i.e ``"San Jose, US"``
	CityCountry pulumi.StringPtrInput
	// i.e ``"US"``, ``"CA"``
	CountryCode pulumi.StringPtrInput
	// Description of the App Connector Group.
	Description pulumi.StringPtrInput
	// Supported values are:
	DnsQueryType pulumi.StringPtrInput
	// Whether this App Connector Group is enabled or not. Default value: `true`. Supported values: `true`, `false`
	Enabled pulumi.BoolPtrInput
	// Latitude of the App Connector Group. Integer or decimal. With values in the range of `-90` to `90`
	Latitude pulumi.StringPtrInput
	// Location of the App Connector Group. i.e ``"San Jose, CA, USA"``
	Location pulumi.StringPtrInput
	// Longitude of the App Connector Group. Integer or decimal. With values in the range of `-180` to `180`
	Longitude            pulumi.StringPtrInput
	LssAppConnectorGroup pulumi.BoolPtrInput
	// The ID of the microtenant the resource is to be associated with.
	//
	// ⚠️ **WARNING:**: The attribute ``microtenantId`` is optional and requires the microtenant license and feature flag enabled for the respective tenant. The provider also supports the microtenant ID configuration via the environment variable `ZPA_MICROTENANT_ID` which is the recommended method.
	MicrotenantId pulumi.StringPtrInput
	// Name of the App Connector Group.
	Name pulumi.StringPtrInput
	// Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
	OverrideVersionProfile pulumi.BoolPtrInput
	// Supported values: `true`, `false`
	PraEnabled pulumi.BoolPtrInput
	// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
	TcpQuickAckApp pulumi.BoolPtrInput
	// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
	TcpQuickAckAssistant pulumi.BoolPtrInput
	// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
	TcpQuickAckReadAssistant pulumi.BoolPtrInput
	// App Connectors in this group will attempt to update to a newer version of the software during this specified day i.e ``SUNDAY``
	UpgradeDay pulumi.StringPtrInput
	// App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600`. Integer in seconds (i.e., `-66600`). The integer should be greater than or equal to `0` and less than `86400`, in `15` minute intervals
	UpgradeTimeInSecs pulumi.StringPtrInput
	// Supported values: `true`, `false`
	UseInDrMode pulumi.BoolPtrInput
	// ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
	VersionProfileId pulumi.StringPtrInput
	// Name of the version profile. To learn more, see Version Profile Use Cases. This value is required, if the value for
	// overrideVersionProfile is set to true
	VersionProfileName pulumi.StringPtrInput
	// Supported values: `true`, `false`
	WafDisabled pulumi.BoolPtrInput
}

func (ConnectorGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorGroupState)(nil)).Elem()
}

type connectorGroupArgs struct {
	// Whether Double Encryption is enabled or disabled for the app. i.e ``"San Jose, US"``
	CityCountry *string `pulumi:"cityCountry"`
	// i.e ``"US"``, ``"CA"``
	CountryCode *string `pulumi:"countryCode"`
	// Description of the App Connector Group.
	Description *string `pulumi:"description"`
	// Supported values are:
	DnsQueryType *string `pulumi:"dnsQueryType"`
	// Whether this App Connector Group is enabled or not. Default value: `true`. Supported values: `true`, `false`
	Enabled *bool `pulumi:"enabled"`
	// Latitude of the App Connector Group. Integer or decimal. With values in the range of `-90` to `90`
	Latitude string `pulumi:"latitude"`
	// Location of the App Connector Group. i.e ``"San Jose, CA, USA"``
	Location string `pulumi:"location"`
	// Longitude of the App Connector Group. Integer or decimal. With values in the range of `-180` to `180`
	Longitude            string `pulumi:"longitude"`
	LssAppConnectorGroup *bool  `pulumi:"lssAppConnectorGroup"`
	// The ID of the microtenant the resource is to be associated with.
	//
	// ⚠️ **WARNING:**: The attribute ``microtenantId`` is optional and requires the microtenant license and feature flag enabled for the respective tenant. The provider also supports the microtenant ID configuration via the environment variable `ZPA_MICROTENANT_ID` which is the recommended method.
	MicrotenantId *string `pulumi:"microtenantId"`
	// Name of the App Connector Group.
	Name *string `pulumi:"name"`
	// Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
	OverrideVersionProfile *bool `pulumi:"overrideVersionProfile"`
	// Supported values: `true`, `false`
	PraEnabled *bool `pulumi:"praEnabled"`
	// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
	TcpQuickAckApp *bool `pulumi:"tcpQuickAckApp"`
	// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
	TcpQuickAckAssistant *bool `pulumi:"tcpQuickAckAssistant"`
	// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
	TcpQuickAckReadAssistant *bool `pulumi:"tcpQuickAckReadAssistant"`
	// App Connectors in this group will attempt to update to a newer version of the software during this specified day i.e ``SUNDAY``
	UpgradeDay *string `pulumi:"upgradeDay"`
	// App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600`. Integer in seconds (i.e., `-66600`). The integer should be greater than or equal to `0` and less than `86400`, in `15` minute intervals
	UpgradeTimeInSecs *string `pulumi:"upgradeTimeInSecs"`
	// Supported values: `true`, `false`
	UseInDrMode *bool `pulumi:"useInDrMode"`
	// ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
	VersionProfileId *string `pulumi:"versionProfileId"`
	// Name of the version profile. To learn more, see Version Profile Use Cases. This value is required, if the value for
	// overrideVersionProfile is set to true
	VersionProfileName *string `pulumi:"versionProfileName"`
	// Supported values: `true`, `false`
	WafDisabled *bool `pulumi:"wafDisabled"`
}

// The set of arguments for constructing a ConnectorGroup resource.
type ConnectorGroupArgs struct {
	// Whether Double Encryption is enabled or disabled for the app. i.e ``"San Jose, US"``
	CityCountry pulumi.StringPtrInput
	// i.e ``"US"``, ``"CA"``
	CountryCode pulumi.StringPtrInput
	// Description of the App Connector Group.
	Description pulumi.StringPtrInput
	// Supported values are:
	DnsQueryType pulumi.StringPtrInput
	// Whether this App Connector Group is enabled or not. Default value: `true`. Supported values: `true`, `false`
	Enabled pulumi.BoolPtrInput
	// Latitude of the App Connector Group. Integer or decimal. With values in the range of `-90` to `90`
	Latitude pulumi.StringInput
	// Location of the App Connector Group. i.e ``"San Jose, CA, USA"``
	Location pulumi.StringInput
	// Longitude of the App Connector Group. Integer or decimal. With values in the range of `-180` to `180`
	Longitude            pulumi.StringInput
	LssAppConnectorGroup pulumi.BoolPtrInput
	// The ID of the microtenant the resource is to be associated with.
	//
	// ⚠️ **WARNING:**: The attribute ``microtenantId`` is optional and requires the microtenant license and feature flag enabled for the respective tenant. The provider also supports the microtenant ID configuration via the environment variable `ZPA_MICROTENANT_ID` which is the recommended method.
	MicrotenantId pulumi.StringPtrInput
	// Name of the App Connector Group.
	Name pulumi.StringPtrInput
	// Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
	OverrideVersionProfile pulumi.BoolPtrInput
	// Supported values: `true`, `false`
	PraEnabled pulumi.BoolPtrInput
	// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
	TcpQuickAckApp pulumi.BoolPtrInput
	// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
	TcpQuickAckAssistant pulumi.BoolPtrInput
	// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
	TcpQuickAckReadAssistant pulumi.BoolPtrInput
	// App Connectors in this group will attempt to update to a newer version of the software during this specified day i.e ``SUNDAY``
	UpgradeDay pulumi.StringPtrInput
	// App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600`. Integer in seconds (i.e., `-66600`). The integer should be greater than or equal to `0` and less than `86400`, in `15` minute intervals
	UpgradeTimeInSecs pulumi.StringPtrInput
	// Supported values: `true`, `false`
	UseInDrMode pulumi.BoolPtrInput
	// ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
	VersionProfileId pulumi.StringPtrInput
	// Name of the version profile. To learn more, see Version Profile Use Cases. This value is required, if the value for
	// overrideVersionProfile is set to true
	VersionProfileName pulumi.StringPtrInput
	// Supported values: `true`, `false`
	WafDisabled pulumi.BoolPtrInput
}

func (ConnectorGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorGroupArgs)(nil)).Elem()
}

type ConnectorGroupInput interface {
	pulumi.Input

	ToConnectorGroupOutput() ConnectorGroupOutput
	ToConnectorGroupOutputWithContext(ctx context.Context) ConnectorGroupOutput
}

func (*ConnectorGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorGroup)(nil)).Elem()
}

func (i *ConnectorGroup) ToConnectorGroupOutput() ConnectorGroupOutput {
	return i.ToConnectorGroupOutputWithContext(context.Background())
}

func (i *ConnectorGroup) ToConnectorGroupOutputWithContext(ctx context.Context) ConnectorGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorGroupOutput)
}

// ConnectorGroupArrayInput is an input type that accepts ConnectorGroupArray and ConnectorGroupArrayOutput values.
// You can construct a concrete instance of `ConnectorGroupArrayInput` via:
//
//	ConnectorGroupArray{ ConnectorGroupArgs{...} }
type ConnectorGroupArrayInput interface {
	pulumi.Input

	ToConnectorGroupArrayOutput() ConnectorGroupArrayOutput
	ToConnectorGroupArrayOutputWithContext(context.Context) ConnectorGroupArrayOutput
}

type ConnectorGroupArray []ConnectorGroupInput

func (ConnectorGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConnectorGroup)(nil)).Elem()
}

func (i ConnectorGroupArray) ToConnectorGroupArrayOutput() ConnectorGroupArrayOutput {
	return i.ToConnectorGroupArrayOutputWithContext(context.Background())
}

func (i ConnectorGroupArray) ToConnectorGroupArrayOutputWithContext(ctx context.Context) ConnectorGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorGroupArrayOutput)
}

// ConnectorGroupMapInput is an input type that accepts ConnectorGroupMap and ConnectorGroupMapOutput values.
// You can construct a concrete instance of `ConnectorGroupMapInput` via:
//
//	ConnectorGroupMap{ "key": ConnectorGroupArgs{...} }
type ConnectorGroupMapInput interface {
	pulumi.Input

	ToConnectorGroupMapOutput() ConnectorGroupMapOutput
	ToConnectorGroupMapOutputWithContext(context.Context) ConnectorGroupMapOutput
}

type ConnectorGroupMap map[string]ConnectorGroupInput

func (ConnectorGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConnectorGroup)(nil)).Elem()
}

func (i ConnectorGroupMap) ToConnectorGroupMapOutput() ConnectorGroupMapOutput {
	return i.ToConnectorGroupMapOutputWithContext(context.Background())
}

func (i ConnectorGroupMap) ToConnectorGroupMapOutputWithContext(ctx context.Context) ConnectorGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorGroupMapOutput)
}

type ConnectorGroupOutput struct{ *pulumi.OutputState }

func (ConnectorGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorGroup)(nil)).Elem()
}

func (o ConnectorGroupOutput) ToConnectorGroupOutput() ConnectorGroupOutput {
	return o
}

func (o ConnectorGroupOutput) ToConnectorGroupOutputWithContext(ctx context.Context) ConnectorGroupOutput {
	return o
}

// Whether Double Encryption is enabled or disabled for the app. i.e “"San Jose, US"“
func (o ConnectorGroupOutput) CityCountry() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.StringOutput { return v.CityCountry }).(pulumi.StringOutput)
}

// i.e “"US"“, “"CA"“
func (o ConnectorGroupOutput) CountryCode() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.StringOutput { return v.CountryCode }).(pulumi.StringOutput)
}

// Description of the App Connector Group.
func (o ConnectorGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Supported values are:
func (o ConnectorGroupOutput) DnsQueryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.StringPtrOutput { return v.DnsQueryType }).(pulumi.StringPtrOutput)
}

// Whether this App Connector Group is enabled or not. Default value: `true`. Supported values: `true`, `false`
func (o ConnectorGroupOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Latitude of the App Connector Group. Integer or decimal. With values in the range of `-90` to `90`
func (o ConnectorGroupOutput) Latitude() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.StringOutput { return v.Latitude }).(pulumi.StringOutput)
}

// Location of the App Connector Group. i.e “"San Jose, CA, USA"“
func (o ConnectorGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Longitude of the App Connector Group. Integer or decimal. With values in the range of `-180` to `180`
func (o ConnectorGroupOutput) Longitude() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.StringOutput { return v.Longitude }).(pulumi.StringOutput)
}

func (o ConnectorGroupOutput) LssAppConnectorGroup() pulumi.BoolOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.BoolOutput { return v.LssAppConnectorGroup }).(pulumi.BoolOutput)
}

// The ID of the microtenant the resource is to be associated with.
//
// ⚠️ **WARNING:**: The attribute “microtenantId“ is optional and requires the microtenant license and feature flag enabled for the respective tenant. The provider also supports the microtenant ID configuration via the environment variable `ZPA_MICROTENANT_ID` which is the recommended method.
func (o ConnectorGroupOutput) MicrotenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.StringOutput { return v.MicrotenantId }).(pulumi.StringOutput)
}

// Name of the App Connector Group.
func (o ConnectorGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
func (o ConnectorGroupOutput) OverrideVersionProfile() pulumi.BoolOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.BoolOutput { return v.OverrideVersionProfile }).(pulumi.BoolOutput)
}

// Supported values: `true`, `false`
func (o ConnectorGroupOutput) PraEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.BoolOutput { return v.PraEnabled }).(pulumi.BoolOutput)
}

// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
func (o ConnectorGroupOutput) TcpQuickAckApp() pulumi.BoolOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.BoolOutput { return v.TcpQuickAckApp }).(pulumi.BoolOutput)
}

// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
func (o ConnectorGroupOutput) TcpQuickAckAssistant() pulumi.BoolOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.BoolOutput { return v.TcpQuickAckAssistant }).(pulumi.BoolOutput)
}

// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
func (o ConnectorGroupOutput) TcpQuickAckReadAssistant() pulumi.BoolOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.BoolOutput { return v.TcpQuickAckReadAssistant }).(pulumi.BoolOutput)
}

// App Connectors in this group will attempt to update to a newer version of the software during this specified day i.e “SUNDAY“
func (o ConnectorGroupOutput) UpgradeDay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.StringPtrOutput { return v.UpgradeDay }).(pulumi.StringPtrOutput)
}

// App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600`. Integer in seconds (i.e., `-66600`). The integer should be greater than or equal to `0` and less than `86400`, in `15` minute intervals
func (o ConnectorGroupOutput) UpgradeTimeInSecs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.StringPtrOutput { return v.UpgradeTimeInSecs }).(pulumi.StringPtrOutput)
}

// Supported values: `true`, `false`
func (o ConnectorGroupOutput) UseInDrMode() pulumi.BoolOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.BoolOutput { return v.UseInDrMode }).(pulumi.BoolOutput)
}

// ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
func (o ConnectorGroupOutput) VersionProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.StringOutput { return v.VersionProfileId }).(pulumi.StringOutput)
}

// Name of the version profile. To learn more, see Version Profile Use Cases. This value is required, if the value for
// overrideVersionProfile is set to true
func (o ConnectorGroupOutput) VersionProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.StringOutput { return v.VersionProfileName }).(pulumi.StringOutput)
}

// Supported values: `true`, `false`
func (o ConnectorGroupOutput) WafDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ConnectorGroup) pulumi.BoolOutput { return v.WafDisabled }).(pulumi.BoolOutput)
}

type ConnectorGroupArrayOutput struct{ *pulumi.OutputState }

func (ConnectorGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConnectorGroup)(nil)).Elem()
}

func (o ConnectorGroupArrayOutput) ToConnectorGroupArrayOutput() ConnectorGroupArrayOutput {
	return o
}

func (o ConnectorGroupArrayOutput) ToConnectorGroupArrayOutputWithContext(ctx context.Context) ConnectorGroupArrayOutput {
	return o
}

func (o ConnectorGroupArrayOutput) Index(i pulumi.IntInput) ConnectorGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConnectorGroup {
		return vs[0].([]*ConnectorGroup)[vs[1].(int)]
	}).(ConnectorGroupOutput)
}

type ConnectorGroupMapOutput struct{ *pulumi.OutputState }

func (ConnectorGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConnectorGroup)(nil)).Elem()
}

func (o ConnectorGroupMapOutput) ToConnectorGroupMapOutput() ConnectorGroupMapOutput {
	return o
}

func (o ConnectorGroupMapOutput) ToConnectorGroupMapOutputWithContext(ctx context.Context) ConnectorGroupMapOutput {
	return o
}

func (o ConnectorGroupMapOutput) MapIndex(k pulumi.StringInput) ConnectorGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConnectorGroup {
		return vs[0].(map[string]*ConnectorGroup)[vs[1].(string)]
	}).(ConnectorGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectorGroupInput)(nil)).Elem(), &ConnectorGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectorGroupArrayInput)(nil)).Elem(), ConnectorGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectorGroupMapInput)(nil)).Elem(), ConnectorGroupMap{})
	pulumi.RegisterOutputType(ConnectorGroupOutput{})
	pulumi.RegisterOutputType(ConnectorGroupArrayOutput{})
	pulumi.RegisterOutputType(ConnectorGroupMapOutput{})
}
