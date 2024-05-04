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

// * [Official documentation](https://help.zscaler.com/zpa/about-appprotection-applications)
// * [API documentation](https://help.zscaler.com/zpa/configuring-application-segments-using-api)
//
// The **zpa_application_segment_inspection** resource creates an inspection application segment in the Zscaler Private Access cloud. This resource can then be referenced in an access policy inspection rule. This resource supports Inspection for both `HTTP` and `HTTPS`.
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
//			jenkins, err := zpa.GetBaCertificate(ctx, &zpa.GetBaCertificateArgs{
//				Name: pulumi.StringRef("jenkins.example.com"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.NewApplicationSegmentInspection(ctx, "this", &zpa.ApplicationSegmentInspectionArgs{
//				Description:     pulumi.String("ZPA_Inspection_Example"),
//				Enabled:         pulumi.Bool(true),
//				HealthReporting: pulumi.String("ON_ACCESS"),
//				BypassType:      pulumi.String("NEVER"),
//				IsCnameEnabled:  pulumi.Bool(true),
//				TcpPortRanges: pulumi.StringArray{
//					pulumi.String("443"),
//					pulumi.String("443"),
//				},
//				DomainNames: pulumi.StringArray{
//					pulumi.String("jenkins.example.com"),
//				},
//				SegmentGroupId: pulumi.Any(zpa_segment_group.This.Id),
//				ServerGroups: zpa.ApplicationSegmentInspectionServerGroupArray{
//					&zpa.ApplicationSegmentInspectionServerGroupArgs{
//						Ids: pulumi.StringArray{
//							zpa_server_group.This.Id,
//						},
//					},
//				},
//				CommonAppsDto: &zpa.ApplicationSegmentInspectionCommonAppsDtoArgs{
//					AppsConfigs: zpa.ApplicationSegmentInspectionCommonAppsDtoAppsConfigArray{
//						&zpa.ApplicationSegmentInspectionCommonAppsDtoAppsConfigArgs{
//							Name:                pulumi.String("jenkins.example.com"),
//							Domain:              pulumi.String("jenkins.example.com"),
//							ApplicationProtocol: pulumi.String("HTTPS"),
//							ApplicationPort:     pulumi.String("443"),
//							CertificateId:       pulumi.String(jenkins.Id),
//							Enabled:             pulumi.Bool(true),
//							AppTypes: pulumi.StringArray{
//								pulumi.String("INSPECT"),
//							},
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
// ## Import
//
// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
//
// # Visit
//
// Inspection Application Segment can be imported by using `<APPLICATION SEGMENT ID>` or `<APPLICATION SEGMENT NAME>` as the import ID.
//
// ```sh
// $ pulumi import zpa:index/applicationSegmentInspection:ApplicationSegmentInspection example <application_segment_id>
// ```
//
// or
//
// ```sh
// $ pulumi import zpa:index/applicationSegmentInspection:ApplicationSegmentInspection example <application_segment_name>
// ```
type ApplicationSegmentInspection struct {
	pulumi.CustomResourceState

	// Indicates whether users can bypass ZPA to access applications. Default: NEVER. Supported values: ALWAYS, NEVER, ON_NET.
	// The value NEVER indicates the use of the client forwarding policy.
	BypassType    pulumi.StringOutput                                `pulumi:"bypassType"`
	CommonAppsDto ApplicationSegmentInspectionCommonAppsDtoPtrOutput `pulumi:"commonAppsDto"`
	ConfigSpace   pulumi.StringPtrOutput                             `pulumi:"configSpace"`
	// Description of the application.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of domains and IPs.
	DomainNames pulumi.StringArrayOutput `pulumi:"domainNames"`
	// Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt   pulumi.BoolOutput      `pulumi:"doubleEncrypt"`
	Enabled         pulumi.BoolOutput      `pulumi:"enabled"`
	HealthCheckType pulumi.StringPtrOutput `pulumi:"healthCheckType"`
	// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
	HealthReporting pulumi.StringPtrOutput `pulumi:"healthReporting"`
	IcmpAccessType  pulumi.StringOutput    `pulumi:"icmpAccessType"`
	IpAnchored      pulumi.BoolOutput      `pulumi:"ipAnchored"`
	// Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the
	// connectors.
	IsCnameEnabled       pulumi.BoolOutput    `pulumi:"isCnameEnabled"`
	IsIncompleteDrConfig pulumi.BoolPtrOutput `pulumi:"isIncompleteDrConfig"`
	MatchStyle           pulumi.StringOutput  `pulumi:"matchStyle"`
	// Name of the application.
	Name                      pulumi.StringOutput  `pulumi:"name"`
	PassiveHealthEnabled      pulumi.BoolOutput    `pulumi:"passiveHealthEnabled"`
	SegmentGroupId            pulumi.StringOutput  `pulumi:"segmentGroupId"`
	SegmentGroupName          pulumi.StringOutput  `pulumi:"segmentGroupName"`
	SelectConnectorCloseToApp pulumi.BoolPtrOutput `pulumi:"selectConnectorCloseToApp"`
	// List of the server group IDs.
	ServerGroups ApplicationSegmentInspectionServerGroupArrayOutput `pulumi:"serverGroups"`
	TcpKeepAlive pulumi.StringOutput                                `pulumi:"tcpKeepAlive"`
	// tcp port range
	TcpPortRange ApplicationSegmentInspectionTcpPortRangeArrayOutput `pulumi:"tcpPortRange"`
	// TCP port ranges used to access the app.
	TcpPortRanges pulumi.StringArrayOutput `pulumi:"tcpPortRanges"`
	// udp port range
	UdpPortRange ApplicationSegmentInspectionUdpPortRangeArrayOutput `pulumi:"udpPortRange"`
	// UDP port ranges used to access the app.
	UdpPortRanges pulumi.StringArrayOutput `pulumi:"udpPortRanges"`
	UseInDrMode   pulumi.BoolPtrOutput     `pulumi:"useInDrMode"`
}

// NewApplicationSegmentInspection registers a new resource with the given unique name, arguments, and options.
func NewApplicationSegmentInspection(ctx *pulumi.Context,
	name string, args *ApplicationSegmentInspectionArgs, opts ...pulumi.ResourceOption) (*ApplicationSegmentInspection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SegmentGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SegmentGroupId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationSegmentInspection
	err := ctx.RegisterResource("zpa:index/applicationSegmentInspection:ApplicationSegmentInspection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationSegmentInspection gets an existing ApplicationSegmentInspection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationSegmentInspection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationSegmentInspectionState, opts ...pulumi.ResourceOption) (*ApplicationSegmentInspection, error) {
	var resource ApplicationSegmentInspection
	err := ctx.ReadResource("zpa:index/applicationSegmentInspection:ApplicationSegmentInspection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationSegmentInspection resources.
type applicationSegmentInspectionState struct {
	// Indicates whether users can bypass ZPA to access applications. Default: NEVER. Supported values: ALWAYS, NEVER, ON_NET.
	// The value NEVER indicates the use of the client forwarding policy.
	BypassType    *string                                    `pulumi:"bypassType"`
	CommonAppsDto *ApplicationSegmentInspectionCommonAppsDto `pulumi:"commonAppsDto"`
	ConfigSpace   *string                                    `pulumi:"configSpace"`
	// Description of the application.
	Description *string `pulumi:"description"`
	// List of domains and IPs.
	DomainNames []string `pulumi:"domainNames"`
	// Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt   *bool   `pulumi:"doubleEncrypt"`
	Enabled         *bool   `pulumi:"enabled"`
	HealthCheckType *string `pulumi:"healthCheckType"`
	// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
	HealthReporting *string `pulumi:"healthReporting"`
	IcmpAccessType  *string `pulumi:"icmpAccessType"`
	IpAnchored      *bool   `pulumi:"ipAnchored"`
	// Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the
	// connectors.
	IsCnameEnabled       *bool   `pulumi:"isCnameEnabled"`
	IsIncompleteDrConfig *bool   `pulumi:"isIncompleteDrConfig"`
	MatchStyle           *string `pulumi:"matchStyle"`
	// Name of the application.
	Name                      *string `pulumi:"name"`
	PassiveHealthEnabled      *bool   `pulumi:"passiveHealthEnabled"`
	SegmentGroupId            *string `pulumi:"segmentGroupId"`
	SegmentGroupName          *string `pulumi:"segmentGroupName"`
	SelectConnectorCloseToApp *bool   `pulumi:"selectConnectorCloseToApp"`
	// List of the server group IDs.
	ServerGroups []ApplicationSegmentInspectionServerGroup `pulumi:"serverGroups"`
	TcpKeepAlive *string                                   `pulumi:"tcpKeepAlive"`
	// tcp port range
	TcpPortRange []ApplicationSegmentInspectionTcpPortRange `pulumi:"tcpPortRange"`
	// TCP port ranges used to access the app.
	TcpPortRanges []string `pulumi:"tcpPortRanges"`
	// udp port range
	UdpPortRange []ApplicationSegmentInspectionUdpPortRange `pulumi:"udpPortRange"`
	// UDP port ranges used to access the app.
	UdpPortRanges []string `pulumi:"udpPortRanges"`
	UseInDrMode   *bool    `pulumi:"useInDrMode"`
}

type ApplicationSegmentInspectionState struct {
	// Indicates whether users can bypass ZPA to access applications. Default: NEVER. Supported values: ALWAYS, NEVER, ON_NET.
	// The value NEVER indicates the use of the client forwarding policy.
	BypassType    pulumi.StringPtrInput
	CommonAppsDto ApplicationSegmentInspectionCommonAppsDtoPtrInput
	ConfigSpace   pulumi.StringPtrInput
	// Description of the application.
	Description pulumi.StringPtrInput
	// List of domains and IPs.
	DomainNames pulumi.StringArrayInput
	// Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt   pulumi.BoolPtrInput
	Enabled         pulumi.BoolPtrInput
	HealthCheckType pulumi.StringPtrInput
	// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
	HealthReporting pulumi.StringPtrInput
	IcmpAccessType  pulumi.StringPtrInput
	IpAnchored      pulumi.BoolPtrInput
	// Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the
	// connectors.
	IsCnameEnabled       pulumi.BoolPtrInput
	IsIncompleteDrConfig pulumi.BoolPtrInput
	MatchStyle           pulumi.StringPtrInput
	// Name of the application.
	Name                      pulumi.StringPtrInput
	PassiveHealthEnabled      pulumi.BoolPtrInput
	SegmentGroupId            pulumi.StringPtrInput
	SegmentGroupName          pulumi.StringPtrInput
	SelectConnectorCloseToApp pulumi.BoolPtrInput
	// List of the server group IDs.
	ServerGroups ApplicationSegmentInspectionServerGroupArrayInput
	TcpKeepAlive pulumi.StringPtrInput
	// tcp port range
	TcpPortRange ApplicationSegmentInspectionTcpPortRangeArrayInput
	// TCP port ranges used to access the app.
	TcpPortRanges pulumi.StringArrayInput
	// udp port range
	UdpPortRange ApplicationSegmentInspectionUdpPortRangeArrayInput
	// UDP port ranges used to access the app.
	UdpPortRanges pulumi.StringArrayInput
	UseInDrMode   pulumi.BoolPtrInput
}

func (ApplicationSegmentInspectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationSegmentInspectionState)(nil)).Elem()
}

type applicationSegmentInspectionArgs struct {
	// Indicates whether users can bypass ZPA to access applications. Default: NEVER. Supported values: ALWAYS, NEVER, ON_NET.
	// The value NEVER indicates the use of the client forwarding policy.
	BypassType    *string                                    `pulumi:"bypassType"`
	CommonAppsDto *ApplicationSegmentInspectionCommonAppsDto `pulumi:"commonAppsDto"`
	ConfigSpace   *string                                    `pulumi:"configSpace"`
	// Description of the application.
	Description *string `pulumi:"description"`
	// List of domains and IPs.
	DomainNames []string `pulumi:"domainNames"`
	// Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt   *bool   `pulumi:"doubleEncrypt"`
	Enabled         *bool   `pulumi:"enabled"`
	HealthCheckType *string `pulumi:"healthCheckType"`
	// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
	HealthReporting *string `pulumi:"healthReporting"`
	IcmpAccessType  *string `pulumi:"icmpAccessType"`
	IpAnchored      *bool   `pulumi:"ipAnchored"`
	// Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the
	// connectors.
	IsCnameEnabled       *bool   `pulumi:"isCnameEnabled"`
	IsIncompleteDrConfig *bool   `pulumi:"isIncompleteDrConfig"`
	MatchStyle           *string `pulumi:"matchStyle"`
	// Name of the application.
	Name                      *string `pulumi:"name"`
	PassiveHealthEnabled      *bool   `pulumi:"passiveHealthEnabled"`
	SegmentGroupId            string  `pulumi:"segmentGroupId"`
	SegmentGroupName          *string `pulumi:"segmentGroupName"`
	SelectConnectorCloseToApp *bool   `pulumi:"selectConnectorCloseToApp"`
	// List of the server group IDs.
	ServerGroups []ApplicationSegmentInspectionServerGroup `pulumi:"serverGroups"`
	TcpKeepAlive *string                                   `pulumi:"tcpKeepAlive"`
	// tcp port range
	TcpPortRange []ApplicationSegmentInspectionTcpPortRange `pulumi:"tcpPortRange"`
	// TCP port ranges used to access the app.
	TcpPortRanges []string `pulumi:"tcpPortRanges"`
	// udp port range
	UdpPortRange []ApplicationSegmentInspectionUdpPortRange `pulumi:"udpPortRange"`
	// UDP port ranges used to access the app.
	UdpPortRanges []string `pulumi:"udpPortRanges"`
	UseInDrMode   *bool    `pulumi:"useInDrMode"`
}

// The set of arguments for constructing a ApplicationSegmentInspection resource.
type ApplicationSegmentInspectionArgs struct {
	// Indicates whether users can bypass ZPA to access applications. Default: NEVER. Supported values: ALWAYS, NEVER, ON_NET.
	// The value NEVER indicates the use of the client forwarding policy.
	BypassType    pulumi.StringPtrInput
	CommonAppsDto ApplicationSegmentInspectionCommonAppsDtoPtrInput
	ConfigSpace   pulumi.StringPtrInput
	// Description of the application.
	Description pulumi.StringPtrInput
	// List of domains and IPs.
	DomainNames pulumi.StringArrayInput
	// Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt   pulumi.BoolPtrInput
	Enabled         pulumi.BoolPtrInput
	HealthCheckType pulumi.StringPtrInput
	// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
	HealthReporting pulumi.StringPtrInput
	IcmpAccessType  pulumi.StringPtrInput
	IpAnchored      pulumi.BoolPtrInput
	// Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the
	// connectors.
	IsCnameEnabled       pulumi.BoolPtrInput
	IsIncompleteDrConfig pulumi.BoolPtrInput
	MatchStyle           pulumi.StringPtrInput
	// Name of the application.
	Name                      pulumi.StringPtrInput
	PassiveHealthEnabled      pulumi.BoolPtrInput
	SegmentGroupId            pulumi.StringInput
	SegmentGroupName          pulumi.StringPtrInput
	SelectConnectorCloseToApp pulumi.BoolPtrInput
	// List of the server group IDs.
	ServerGroups ApplicationSegmentInspectionServerGroupArrayInput
	TcpKeepAlive pulumi.StringPtrInput
	// tcp port range
	TcpPortRange ApplicationSegmentInspectionTcpPortRangeArrayInput
	// TCP port ranges used to access the app.
	TcpPortRanges pulumi.StringArrayInput
	// udp port range
	UdpPortRange ApplicationSegmentInspectionUdpPortRangeArrayInput
	// UDP port ranges used to access the app.
	UdpPortRanges pulumi.StringArrayInput
	UseInDrMode   pulumi.BoolPtrInput
}

func (ApplicationSegmentInspectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationSegmentInspectionArgs)(nil)).Elem()
}

type ApplicationSegmentInspectionInput interface {
	pulumi.Input

	ToApplicationSegmentInspectionOutput() ApplicationSegmentInspectionOutput
	ToApplicationSegmentInspectionOutputWithContext(ctx context.Context) ApplicationSegmentInspectionOutput
}

func (*ApplicationSegmentInspection) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationSegmentInspection)(nil)).Elem()
}

func (i *ApplicationSegmentInspection) ToApplicationSegmentInspectionOutput() ApplicationSegmentInspectionOutput {
	return i.ToApplicationSegmentInspectionOutputWithContext(context.Background())
}

func (i *ApplicationSegmentInspection) ToApplicationSegmentInspectionOutputWithContext(ctx context.Context) ApplicationSegmentInspectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSegmentInspectionOutput)
}

// ApplicationSegmentInspectionArrayInput is an input type that accepts ApplicationSegmentInspectionArray and ApplicationSegmentInspectionArrayOutput values.
// You can construct a concrete instance of `ApplicationSegmentInspectionArrayInput` via:
//
//	ApplicationSegmentInspectionArray{ ApplicationSegmentInspectionArgs{...} }
type ApplicationSegmentInspectionArrayInput interface {
	pulumi.Input

	ToApplicationSegmentInspectionArrayOutput() ApplicationSegmentInspectionArrayOutput
	ToApplicationSegmentInspectionArrayOutputWithContext(context.Context) ApplicationSegmentInspectionArrayOutput
}

type ApplicationSegmentInspectionArray []ApplicationSegmentInspectionInput

func (ApplicationSegmentInspectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationSegmentInspection)(nil)).Elem()
}

func (i ApplicationSegmentInspectionArray) ToApplicationSegmentInspectionArrayOutput() ApplicationSegmentInspectionArrayOutput {
	return i.ToApplicationSegmentInspectionArrayOutputWithContext(context.Background())
}

func (i ApplicationSegmentInspectionArray) ToApplicationSegmentInspectionArrayOutputWithContext(ctx context.Context) ApplicationSegmentInspectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSegmentInspectionArrayOutput)
}

// ApplicationSegmentInspectionMapInput is an input type that accepts ApplicationSegmentInspectionMap and ApplicationSegmentInspectionMapOutput values.
// You can construct a concrete instance of `ApplicationSegmentInspectionMapInput` via:
//
//	ApplicationSegmentInspectionMap{ "key": ApplicationSegmentInspectionArgs{...} }
type ApplicationSegmentInspectionMapInput interface {
	pulumi.Input

	ToApplicationSegmentInspectionMapOutput() ApplicationSegmentInspectionMapOutput
	ToApplicationSegmentInspectionMapOutputWithContext(context.Context) ApplicationSegmentInspectionMapOutput
}

type ApplicationSegmentInspectionMap map[string]ApplicationSegmentInspectionInput

func (ApplicationSegmentInspectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationSegmentInspection)(nil)).Elem()
}

func (i ApplicationSegmentInspectionMap) ToApplicationSegmentInspectionMapOutput() ApplicationSegmentInspectionMapOutput {
	return i.ToApplicationSegmentInspectionMapOutputWithContext(context.Background())
}

func (i ApplicationSegmentInspectionMap) ToApplicationSegmentInspectionMapOutputWithContext(ctx context.Context) ApplicationSegmentInspectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSegmentInspectionMapOutput)
}

type ApplicationSegmentInspectionOutput struct{ *pulumi.OutputState }

func (ApplicationSegmentInspectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationSegmentInspection)(nil)).Elem()
}

func (o ApplicationSegmentInspectionOutput) ToApplicationSegmentInspectionOutput() ApplicationSegmentInspectionOutput {
	return o
}

func (o ApplicationSegmentInspectionOutput) ToApplicationSegmentInspectionOutputWithContext(ctx context.Context) ApplicationSegmentInspectionOutput {
	return o
}

// Indicates whether users can bypass ZPA to access applications. Default: NEVER. Supported values: ALWAYS, NEVER, ON_NET.
// The value NEVER indicates the use of the client forwarding policy.
func (o ApplicationSegmentInspectionOutput) BypassType() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.StringOutput { return v.BypassType }).(pulumi.StringOutput)
}

func (o ApplicationSegmentInspectionOutput) CommonAppsDto() ApplicationSegmentInspectionCommonAppsDtoPtrOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) ApplicationSegmentInspectionCommonAppsDtoPtrOutput {
		return v.CommonAppsDto
	}).(ApplicationSegmentInspectionCommonAppsDtoPtrOutput)
}

func (o ApplicationSegmentInspectionOutput) ConfigSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.StringPtrOutput { return v.ConfigSpace }).(pulumi.StringPtrOutput)
}

// Description of the application.
func (o ApplicationSegmentInspectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of domains and IPs.
func (o ApplicationSegmentInspectionOutput) DomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.StringArrayOutput { return v.DomainNames }).(pulumi.StringArrayOutput)
}

// Whether Double Encryption is enabled or disabled for the app.
func (o ApplicationSegmentInspectionOutput) DoubleEncrypt() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.BoolOutput { return v.DoubleEncrypt }).(pulumi.BoolOutput)
}

func (o ApplicationSegmentInspectionOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o ApplicationSegmentInspectionOutput) HealthCheckType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.StringPtrOutput { return v.HealthCheckType }).(pulumi.StringPtrOutput)
}

// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
func (o ApplicationSegmentInspectionOutput) HealthReporting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.StringPtrOutput { return v.HealthReporting }).(pulumi.StringPtrOutput)
}

func (o ApplicationSegmentInspectionOutput) IcmpAccessType() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.StringOutput { return v.IcmpAccessType }).(pulumi.StringOutput)
}

func (o ApplicationSegmentInspectionOutput) IpAnchored() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.BoolOutput { return v.IpAnchored }).(pulumi.BoolOutput)
}

// Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the
// connectors.
func (o ApplicationSegmentInspectionOutput) IsCnameEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.BoolOutput { return v.IsCnameEnabled }).(pulumi.BoolOutput)
}

func (o ApplicationSegmentInspectionOutput) IsIncompleteDrConfig() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.BoolPtrOutput { return v.IsIncompleteDrConfig }).(pulumi.BoolPtrOutput)
}

func (o ApplicationSegmentInspectionOutput) MatchStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.StringOutput { return v.MatchStyle }).(pulumi.StringOutput)
}

// Name of the application.
func (o ApplicationSegmentInspectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationSegmentInspectionOutput) PassiveHealthEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.BoolOutput { return v.PassiveHealthEnabled }).(pulumi.BoolOutput)
}

func (o ApplicationSegmentInspectionOutput) SegmentGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.StringOutput { return v.SegmentGroupId }).(pulumi.StringOutput)
}

func (o ApplicationSegmentInspectionOutput) SegmentGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.StringOutput { return v.SegmentGroupName }).(pulumi.StringOutput)
}

func (o ApplicationSegmentInspectionOutput) SelectConnectorCloseToApp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.BoolPtrOutput { return v.SelectConnectorCloseToApp }).(pulumi.BoolPtrOutput)
}

// List of the server group IDs.
func (o ApplicationSegmentInspectionOutput) ServerGroups() ApplicationSegmentInspectionServerGroupArrayOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) ApplicationSegmentInspectionServerGroupArrayOutput {
		return v.ServerGroups
	}).(ApplicationSegmentInspectionServerGroupArrayOutput)
}

func (o ApplicationSegmentInspectionOutput) TcpKeepAlive() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.StringOutput { return v.TcpKeepAlive }).(pulumi.StringOutput)
}

// tcp port range
func (o ApplicationSegmentInspectionOutput) TcpPortRange() ApplicationSegmentInspectionTcpPortRangeArrayOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) ApplicationSegmentInspectionTcpPortRangeArrayOutput {
		return v.TcpPortRange
	}).(ApplicationSegmentInspectionTcpPortRangeArrayOutput)
}

// TCP port ranges used to access the app.
func (o ApplicationSegmentInspectionOutput) TcpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.StringArrayOutput { return v.TcpPortRanges }).(pulumi.StringArrayOutput)
}

// udp port range
func (o ApplicationSegmentInspectionOutput) UdpPortRange() ApplicationSegmentInspectionUdpPortRangeArrayOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) ApplicationSegmentInspectionUdpPortRangeArrayOutput {
		return v.UdpPortRange
	}).(ApplicationSegmentInspectionUdpPortRangeArrayOutput)
}

// UDP port ranges used to access the app.
func (o ApplicationSegmentInspectionOutput) UdpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.StringArrayOutput { return v.UdpPortRanges }).(pulumi.StringArrayOutput)
}

func (o ApplicationSegmentInspectionOutput) UseInDrMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationSegmentInspection) pulumi.BoolPtrOutput { return v.UseInDrMode }).(pulumi.BoolPtrOutput)
}

type ApplicationSegmentInspectionArrayOutput struct{ *pulumi.OutputState }

func (ApplicationSegmentInspectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationSegmentInspection)(nil)).Elem()
}

func (o ApplicationSegmentInspectionArrayOutput) ToApplicationSegmentInspectionArrayOutput() ApplicationSegmentInspectionArrayOutput {
	return o
}

func (o ApplicationSegmentInspectionArrayOutput) ToApplicationSegmentInspectionArrayOutputWithContext(ctx context.Context) ApplicationSegmentInspectionArrayOutput {
	return o
}

func (o ApplicationSegmentInspectionArrayOutput) Index(i pulumi.IntInput) ApplicationSegmentInspectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationSegmentInspection {
		return vs[0].([]*ApplicationSegmentInspection)[vs[1].(int)]
	}).(ApplicationSegmentInspectionOutput)
}

type ApplicationSegmentInspectionMapOutput struct{ *pulumi.OutputState }

func (ApplicationSegmentInspectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationSegmentInspection)(nil)).Elem()
}

func (o ApplicationSegmentInspectionMapOutput) ToApplicationSegmentInspectionMapOutput() ApplicationSegmentInspectionMapOutput {
	return o
}

func (o ApplicationSegmentInspectionMapOutput) ToApplicationSegmentInspectionMapOutputWithContext(ctx context.Context) ApplicationSegmentInspectionMapOutput {
	return o
}

func (o ApplicationSegmentInspectionMapOutput) MapIndex(k pulumi.StringInput) ApplicationSegmentInspectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationSegmentInspection {
		return vs[0].(map[string]*ApplicationSegmentInspection)[vs[1].(string)]
	}).(ApplicationSegmentInspectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationSegmentInspectionInput)(nil)).Elem(), &ApplicationSegmentInspection{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationSegmentInspectionArrayInput)(nil)).Elem(), ApplicationSegmentInspectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationSegmentInspectionMapInput)(nil)).Elem(), ApplicationSegmentInspectionMap{})
	pulumi.RegisterOutputType(ApplicationSegmentInspectionOutput{})
	pulumi.RegisterOutputType(ApplicationSegmentInspectionArrayOutput{})
	pulumi.RegisterOutputType(ApplicationSegmentInspectionMapOutput{})
}
