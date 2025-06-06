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

// * [Official documentation](https://help.zscaler.com/zpa/about-privileged-remote-access-applications)
// * [API documentation](https://help.zscaler.com/zpa/configuring-application-segments-using-api)
//
// The **zpa_application_segment_pra** resource creates an application segment for Privileged Remote Access in the Zscaler Private Access cloud. This resource can then be referenced in an access policy rule, access policy timeout rule, access policy client forwarding rule and inspection policy. This resource supports Privileged Remote Access for both `RDP` and `SSH`.
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
//			_, err := zpa.NewApplicationSegmentPRA(ctx, "this", &zpa.ApplicationSegmentPRAArgs{
//				Description:     pulumi.String("PRA_Example"),
//				Enabled:         pulumi.Bool(true),
//				HealthReporting: pulumi.String("ON_ACCESS"),
//				BypassType:      pulumi.String("NEVER"),
//				IsCnameEnabled:  pulumi.Bool(true),
//				TcpPortRanges: pulumi.StringArray{
//					pulumi.String("22"),
//					pulumi.String("22"),
//					pulumi.String("3389"),
//					pulumi.String("3389"),
//				},
//				DomainNames: pulumi.StringArray{
//					pulumi.String("ssh_pra.example.com"),
//					pulumi.String("rdp_pra.example.com"),
//				},
//				SegmentGroupId: pulumi.Any(zpa_segment_group.This.Id),
//				ServerGroups: zpa.ApplicationSegmentPRAServerGroupArray{
//					&zpa.ApplicationSegmentPRAServerGroupArgs{
//						Ids: pulumi.StringArray{
//							zpa_server_group.This.Id,
//						},
//					},
//				},
//				CommonAppsDtos: zpa.ApplicationSegmentPRACommonAppsDtoArray{
//					&zpa.ApplicationSegmentPRACommonAppsDtoArgs{
//						AppsConfigs: zpa.ApplicationSegmentPRACommonAppsDtoAppsConfigArray{
//							&zpa.ApplicationSegmentPRACommonAppsDtoAppsConfigArgs{
//								Domain:              pulumi.String("ssh_pra.example.com"),
//								ApplicationProtocol: pulumi.String("SSH"),
//								ApplicationPort:     pulumi.String("22"),
//								Enabled:             pulumi.Bool(true),
//								AppTypes: pulumi.StringArray{
//									pulumi.String("SECURE_REMOTE_ACCESS"),
//								},
//							},
//							&zpa.ApplicationSegmentPRACommonAppsDtoAppsConfigArgs{
//								Domain:              pulumi.String("rdp_pra.example.com"),
//								ApplicationProtocol: pulumi.String("RDP"),
//								ConnectionSecurity:  pulumi.String("ANY"),
//								ApplicationPort:     pulumi.String("3389"),
//								Enabled:             pulumi.Bool(true),
//								AppTypes: pulumi.StringArray{
//									pulumi.String("SECURE_REMOTE_ACCESS"),
//								},
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
// Application Segment can be imported by using `<APPLICATION SEGMENT ID>` or `<APPLICATION SEGMENT NAME>` as the import ID.
//
// ```sh
// $ pulumi import zpa:index/applicationSegmentPRA:ApplicationSegmentPRA example <application_segment_id>
// ```
//
// or
//
// ```sh
// $ pulumi import zpa:index/applicationSegmentPRA:ApplicationSegmentPRA example <application_segment_name>
// ```
type ApplicationSegmentPRA struct {
	pulumi.CustomResourceState

	BypassOnReauth pulumi.BoolOutput `pulumi:"bypassOnReauth"`
	// Indicates whether users can bypass ZPA to access applications. Default: NEVER. Supported values: ALWAYS, NEVER, ON_NET.
	// The value NEVER indicates the use of the client forwarding policy.
	BypassType     pulumi.StringOutput                           `pulumi:"bypassType"`
	CommonAppsDtos ApplicationSegmentPRACommonAppsDtoArrayOutput `pulumi:"commonAppsDtos"`
	ConfigSpace    pulumi.StringPtrOutput                        `pulumi:"configSpace"`
	// Description of the application.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of domains and IPs.
	DomainNames pulumi.StringArrayOutput `pulumi:"domainNames"`
	// Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt   pulumi.BoolOutput      `pulumi:"doubleEncrypt"`
	Enabled         pulumi.BoolOutput      `pulumi:"enabled"`
	FqdnDnsCheck    pulumi.BoolPtrOutput   `pulumi:"fqdnDnsCheck"`
	HealthCheckType pulumi.StringPtrOutput `pulumi:"healthCheckType"`
	// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
	HealthReporting pulumi.StringPtrOutput `pulumi:"healthReporting"`
	IcmpAccessType  pulumi.StringOutput    `pulumi:"icmpAccessType"`
	IpAnchored      pulumi.BoolPtrOutput   `pulumi:"ipAnchored"`
	// Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the
	// connectors.
	IsCnameEnabled       pulumi.BoolOutput   `pulumi:"isCnameEnabled"`
	IsIncompleteDrConfig pulumi.BoolOutput   `pulumi:"isIncompleteDrConfig"`
	MicrotenantId        pulumi.StringOutput `pulumi:"microtenantId"`
	// Name of the application.
	Name                      pulumi.StringOutput                         `pulumi:"name"`
	PassiveHealthEnabled      pulumi.BoolOutput                           `pulumi:"passiveHealthEnabled"`
	SegmentGroupId            pulumi.StringOutput                         `pulumi:"segmentGroupId"`
	SelectConnectorCloseToApp pulumi.BoolPtrOutput                        `pulumi:"selectConnectorCloseToApp"`
	ServerGroups              ApplicationSegmentPRAServerGroupArrayOutput `pulumi:"serverGroups"`
	TcpKeepAlive              pulumi.StringOutput                         `pulumi:"tcpKeepAlive"`
	// tcp port range
	TcpPortRange ApplicationSegmentPRATcpPortRangeArrayOutput `pulumi:"tcpPortRange"`
	// TCP port ranges used to access the app.
	TcpPortRanges pulumi.StringArrayOutput `pulumi:"tcpPortRanges"`
	// udp port range
	UdpPortRange ApplicationSegmentPRAUdpPortRangeArrayOutput `pulumi:"udpPortRange"`
	// UDP port ranges used to access the app.
	UdpPortRanges pulumi.StringArrayOutput `pulumi:"udpPortRanges"`
	UseInDrMode   pulumi.BoolOutput        `pulumi:"useInDrMode"`
}

// NewApplicationSegmentPRA registers a new resource with the given unique name, arguments, and options.
func NewApplicationSegmentPRA(ctx *pulumi.Context,
	name string, args *ApplicationSegmentPRAArgs, opts ...pulumi.ResourceOption) (*ApplicationSegmentPRA, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainNames == nil {
		return nil, errors.New("invalid value for required argument 'DomainNames'")
	}
	if args.SegmentGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SegmentGroupId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationSegmentPRA
	err := ctx.RegisterResource("zpa:index/applicationSegmentPRA:ApplicationSegmentPRA", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationSegmentPRA gets an existing ApplicationSegmentPRA resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationSegmentPRA(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationSegmentPRAState, opts ...pulumi.ResourceOption) (*ApplicationSegmentPRA, error) {
	var resource ApplicationSegmentPRA
	err := ctx.ReadResource("zpa:index/applicationSegmentPRA:ApplicationSegmentPRA", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationSegmentPRA resources.
type applicationSegmentPRAState struct {
	BypassOnReauth *bool `pulumi:"bypassOnReauth"`
	// Indicates whether users can bypass ZPA to access applications. Default: NEVER. Supported values: ALWAYS, NEVER, ON_NET.
	// The value NEVER indicates the use of the client forwarding policy.
	BypassType     *string                              `pulumi:"bypassType"`
	CommonAppsDtos []ApplicationSegmentPRACommonAppsDto `pulumi:"commonAppsDtos"`
	ConfigSpace    *string                              `pulumi:"configSpace"`
	// Description of the application.
	Description *string `pulumi:"description"`
	// List of domains and IPs.
	DomainNames []string `pulumi:"domainNames"`
	// Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt   *bool   `pulumi:"doubleEncrypt"`
	Enabled         *bool   `pulumi:"enabled"`
	FqdnDnsCheck    *bool   `pulumi:"fqdnDnsCheck"`
	HealthCheckType *string `pulumi:"healthCheckType"`
	// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
	HealthReporting *string `pulumi:"healthReporting"`
	IcmpAccessType  *string `pulumi:"icmpAccessType"`
	IpAnchored      *bool   `pulumi:"ipAnchored"`
	// Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the
	// connectors.
	IsCnameEnabled       *bool   `pulumi:"isCnameEnabled"`
	IsIncompleteDrConfig *bool   `pulumi:"isIncompleteDrConfig"`
	MicrotenantId        *string `pulumi:"microtenantId"`
	// Name of the application.
	Name                      *string                            `pulumi:"name"`
	PassiveHealthEnabled      *bool                              `pulumi:"passiveHealthEnabled"`
	SegmentGroupId            *string                            `pulumi:"segmentGroupId"`
	SelectConnectorCloseToApp *bool                              `pulumi:"selectConnectorCloseToApp"`
	ServerGroups              []ApplicationSegmentPRAServerGroup `pulumi:"serverGroups"`
	TcpKeepAlive              *string                            `pulumi:"tcpKeepAlive"`
	// tcp port range
	TcpPortRange []ApplicationSegmentPRATcpPortRange `pulumi:"tcpPortRange"`
	// TCP port ranges used to access the app.
	TcpPortRanges []string `pulumi:"tcpPortRanges"`
	// udp port range
	UdpPortRange []ApplicationSegmentPRAUdpPortRange `pulumi:"udpPortRange"`
	// UDP port ranges used to access the app.
	UdpPortRanges []string `pulumi:"udpPortRanges"`
	UseInDrMode   *bool    `pulumi:"useInDrMode"`
}

type ApplicationSegmentPRAState struct {
	BypassOnReauth pulumi.BoolPtrInput
	// Indicates whether users can bypass ZPA to access applications. Default: NEVER. Supported values: ALWAYS, NEVER, ON_NET.
	// The value NEVER indicates the use of the client forwarding policy.
	BypassType     pulumi.StringPtrInput
	CommonAppsDtos ApplicationSegmentPRACommonAppsDtoArrayInput
	ConfigSpace    pulumi.StringPtrInput
	// Description of the application.
	Description pulumi.StringPtrInput
	// List of domains and IPs.
	DomainNames pulumi.StringArrayInput
	// Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt   pulumi.BoolPtrInput
	Enabled         pulumi.BoolPtrInput
	FqdnDnsCheck    pulumi.BoolPtrInput
	HealthCheckType pulumi.StringPtrInput
	// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
	HealthReporting pulumi.StringPtrInput
	IcmpAccessType  pulumi.StringPtrInput
	IpAnchored      pulumi.BoolPtrInput
	// Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the
	// connectors.
	IsCnameEnabled       pulumi.BoolPtrInput
	IsIncompleteDrConfig pulumi.BoolPtrInput
	MicrotenantId        pulumi.StringPtrInput
	// Name of the application.
	Name                      pulumi.StringPtrInput
	PassiveHealthEnabled      pulumi.BoolPtrInput
	SegmentGroupId            pulumi.StringPtrInput
	SelectConnectorCloseToApp pulumi.BoolPtrInput
	ServerGroups              ApplicationSegmentPRAServerGroupArrayInput
	TcpKeepAlive              pulumi.StringPtrInput
	// tcp port range
	TcpPortRange ApplicationSegmentPRATcpPortRangeArrayInput
	// TCP port ranges used to access the app.
	TcpPortRanges pulumi.StringArrayInput
	// udp port range
	UdpPortRange ApplicationSegmentPRAUdpPortRangeArrayInput
	// UDP port ranges used to access the app.
	UdpPortRanges pulumi.StringArrayInput
	UseInDrMode   pulumi.BoolPtrInput
}

func (ApplicationSegmentPRAState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationSegmentPRAState)(nil)).Elem()
}

type applicationSegmentPRAArgs struct {
	BypassOnReauth *bool `pulumi:"bypassOnReauth"`
	// Indicates whether users can bypass ZPA to access applications. Default: NEVER. Supported values: ALWAYS, NEVER, ON_NET.
	// The value NEVER indicates the use of the client forwarding policy.
	BypassType     *string                              `pulumi:"bypassType"`
	CommonAppsDtos []ApplicationSegmentPRACommonAppsDto `pulumi:"commonAppsDtos"`
	ConfigSpace    *string                              `pulumi:"configSpace"`
	// Description of the application.
	Description *string `pulumi:"description"`
	// List of domains and IPs.
	DomainNames []string `pulumi:"domainNames"`
	// Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt   *bool   `pulumi:"doubleEncrypt"`
	Enabled         *bool   `pulumi:"enabled"`
	FqdnDnsCheck    *bool   `pulumi:"fqdnDnsCheck"`
	HealthCheckType *string `pulumi:"healthCheckType"`
	// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
	HealthReporting *string `pulumi:"healthReporting"`
	IcmpAccessType  *string `pulumi:"icmpAccessType"`
	IpAnchored      *bool   `pulumi:"ipAnchored"`
	// Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the
	// connectors.
	IsCnameEnabled       *bool   `pulumi:"isCnameEnabled"`
	IsIncompleteDrConfig *bool   `pulumi:"isIncompleteDrConfig"`
	MicrotenantId        *string `pulumi:"microtenantId"`
	// Name of the application.
	Name                      *string                            `pulumi:"name"`
	PassiveHealthEnabled      *bool                              `pulumi:"passiveHealthEnabled"`
	SegmentGroupId            string                             `pulumi:"segmentGroupId"`
	SelectConnectorCloseToApp *bool                              `pulumi:"selectConnectorCloseToApp"`
	ServerGroups              []ApplicationSegmentPRAServerGroup `pulumi:"serverGroups"`
	TcpKeepAlive              *string                            `pulumi:"tcpKeepAlive"`
	// tcp port range
	TcpPortRange []ApplicationSegmentPRATcpPortRange `pulumi:"tcpPortRange"`
	// TCP port ranges used to access the app.
	TcpPortRanges []string `pulumi:"tcpPortRanges"`
	// udp port range
	UdpPortRange []ApplicationSegmentPRAUdpPortRange `pulumi:"udpPortRange"`
	// UDP port ranges used to access the app.
	UdpPortRanges []string `pulumi:"udpPortRanges"`
	UseInDrMode   *bool    `pulumi:"useInDrMode"`
}

// The set of arguments for constructing a ApplicationSegmentPRA resource.
type ApplicationSegmentPRAArgs struct {
	BypassOnReauth pulumi.BoolPtrInput
	// Indicates whether users can bypass ZPA to access applications. Default: NEVER. Supported values: ALWAYS, NEVER, ON_NET.
	// The value NEVER indicates the use of the client forwarding policy.
	BypassType     pulumi.StringPtrInput
	CommonAppsDtos ApplicationSegmentPRACommonAppsDtoArrayInput
	ConfigSpace    pulumi.StringPtrInput
	// Description of the application.
	Description pulumi.StringPtrInput
	// List of domains and IPs.
	DomainNames pulumi.StringArrayInput
	// Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt   pulumi.BoolPtrInput
	Enabled         pulumi.BoolPtrInput
	FqdnDnsCheck    pulumi.BoolPtrInput
	HealthCheckType pulumi.StringPtrInput
	// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
	HealthReporting pulumi.StringPtrInput
	IcmpAccessType  pulumi.StringPtrInput
	IpAnchored      pulumi.BoolPtrInput
	// Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the
	// connectors.
	IsCnameEnabled       pulumi.BoolPtrInput
	IsIncompleteDrConfig pulumi.BoolPtrInput
	MicrotenantId        pulumi.StringPtrInput
	// Name of the application.
	Name                      pulumi.StringPtrInput
	PassiveHealthEnabled      pulumi.BoolPtrInput
	SegmentGroupId            pulumi.StringInput
	SelectConnectorCloseToApp pulumi.BoolPtrInput
	ServerGroups              ApplicationSegmentPRAServerGroupArrayInput
	TcpKeepAlive              pulumi.StringPtrInput
	// tcp port range
	TcpPortRange ApplicationSegmentPRATcpPortRangeArrayInput
	// TCP port ranges used to access the app.
	TcpPortRanges pulumi.StringArrayInput
	// udp port range
	UdpPortRange ApplicationSegmentPRAUdpPortRangeArrayInput
	// UDP port ranges used to access the app.
	UdpPortRanges pulumi.StringArrayInput
	UseInDrMode   pulumi.BoolPtrInput
}

func (ApplicationSegmentPRAArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationSegmentPRAArgs)(nil)).Elem()
}

type ApplicationSegmentPRAInput interface {
	pulumi.Input

	ToApplicationSegmentPRAOutput() ApplicationSegmentPRAOutput
	ToApplicationSegmentPRAOutputWithContext(ctx context.Context) ApplicationSegmentPRAOutput
}

func (*ApplicationSegmentPRA) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationSegmentPRA)(nil)).Elem()
}

func (i *ApplicationSegmentPRA) ToApplicationSegmentPRAOutput() ApplicationSegmentPRAOutput {
	return i.ToApplicationSegmentPRAOutputWithContext(context.Background())
}

func (i *ApplicationSegmentPRA) ToApplicationSegmentPRAOutputWithContext(ctx context.Context) ApplicationSegmentPRAOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSegmentPRAOutput)
}

// ApplicationSegmentPRAArrayInput is an input type that accepts ApplicationSegmentPRAArray and ApplicationSegmentPRAArrayOutput values.
// You can construct a concrete instance of `ApplicationSegmentPRAArrayInput` via:
//
//	ApplicationSegmentPRAArray{ ApplicationSegmentPRAArgs{...} }
type ApplicationSegmentPRAArrayInput interface {
	pulumi.Input

	ToApplicationSegmentPRAArrayOutput() ApplicationSegmentPRAArrayOutput
	ToApplicationSegmentPRAArrayOutputWithContext(context.Context) ApplicationSegmentPRAArrayOutput
}

type ApplicationSegmentPRAArray []ApplicationSegmentPRAInput

func (ApplicationSegmentPRAArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationSegmentPRA)(nil)).Elem()
}

func (i ApplicationSegmentPRAArray) ToApplicationSegmentPRAArrayOutput() ApplicationSegmentPRAArrayOutput {
	return i.ToApplicationSegmentPRAArrayOutputWithContext(context.Background())
}

func (i ApplicationSegmentPRAArray) ToApplicationSegmentPRAArrayOutputWithContext(ctx context.Context) ApplicationSegmentPRAArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSegmentPRAArrayOutput)
}

// ApplicationSegmentPRAMapInput is an input type that accepts ApplicationSegmentPRAMap and ApplicationSegmentPRAMapOutput values.
// You can construct a concrete instance of `ApplicationSegmentPRAMapInput` via:
//
//	ApplicationSegmentPRAMap{ "key": ApplicationSegmentPRAArgs{...} }
type ApplicationSegmentPRAMapInput interface {
	pulumi.Input

	ToApplicationSegmentPRAMapOutput() ApplicationSegmentPRAMapOutput
	ToApplicationSegmentPRAMapOutputWithContext(context.Context) ApplicationSegmentPRAMapOutput
}

type ApplicationSegmentPRAMap map[string]ApplicationSegmentPRAInput

func (ApplicationSegmentPRAMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationSegmentPRA)(nil)).Elem()
}

func (i ApplicationSegmentPRAMap) ToApplicationSegmentPRAMapOutput() ApplicationSegmentPRAMapOutput {
	return i.ToApplicationSegmentPRAMapOutputWithContext(context.Background())
}

func (i ApplicationSegmentPRAMap) ToApplicationSegmentPRAMapOutputWithContext(ctx context.Context) ApplicationSegmentPRAMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSegmentPRAMapOutput)
}

type ApplicationSegmentPRAOutput struct{ *pulumi.OutputState }

func (ApplicationSegmentPRAOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationSegmentPRA)(nil)).Elem()
}

func (o ApplicationSegmentPRAOutput) ToApplicationSegmentPRAOutput() ApplicationSegmentPRAOutput {
	return o
}

func (o ApplicationSegmentPRAOutput) ToApplicationSegmentPRAOutputWithContext(ctx context.Context) ApplicationSegmentPRAOutput {
	return o
}

func (o ApplicationSegmentPRAOutput) BypassOnReauth() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.BoolOutput { return v.BypassOnReauth }).(pulumi.BoolOutput)
}

// Indicates whether users can bypass ZPA to access applications. Default: NEVER. Supported values: ALWAYS, NEVER, ON_NET.
// The value NEVER indicates the use of the client forwarding policy.
func (o ApplicationSegmentPRAOutput) BypassType() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.StringOutput { return v.BypassType }).(pulumi.StringOutput)
}

func (o ApplicationSegmentPRAOutput) CommonAppsDtos() ApplicationSegmentPRACommonAppsDtoArrayOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) ApplicationSegmentPRACommonAppsDtoArrayOutput { return v.CommonAppsDtos }).(ApplicationSegmentPRACommonAppsDtoArrayOutput)
}

func (o ApplicationSegmentPRAOutput) ConfigSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.StringPtrOutput { return v.ConfigSpace }).(pulumi.StringPtrOutput)
}

// Description of the application.
func (o ApplicationSegmentPRAOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of domains and IPs.
func (o ApplicationSegmentPRAOutput) DomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.StringArrayOutput { return v.DomainNames }).(pulumi.StringArrayOutput)
}

// Whether Double Encryption is enabled or disabled for the app.
func (o ApplicationSegmentPRAOutput) DoubleEncrypt() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.BoolOutput { return v.DoubleEncrypt }).(pulumi.BoolOutput)
}

func (o ApplicationSegmentPRAOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o ApplicationSegmentPRAOutput) FqdnDnsCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.BoolPtrOutput { return v.FqdnDnsCheck }).(pulumi.BoolPtrOutput)
}

func (o ApplicationSegmentPRAOutput) HealthCheckType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.StringPtrOutput { return v.HealthCheckType }).(pulumi.StringPtrOutput)
}

// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
func (o ApplicationSegmentPRAOutput) HealthReporting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.StringPtrOutput { return v.HealthReporting }).(pulumi.StringPtrOutput)
}

func (o ApplicationSegmentPRAOutput) IcmpAccessType() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.StringOutput { return v.IcmpAccessType }).(pulumi.StringOutput)
}

func (o ApplicationSegmentPRAOutput) IpAnchored() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.BoolPtrOutput { return v.IpAnchored }).(pulumi.BoolPtrOutput)
}

// Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the
// connectors.
func (o ApplicationSegmentPRAOutput) IsCnameEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.BoolOutput { return v.IsCnameEnabled }).(pulumi.BoolOutput)
}

func (o ApplicationSegmentPRAOutput) IsIncompleteDrConfig() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.BoolOutput { return v.IsIncompleteDrConfig }).(pulumi.BoolOutput)
}

func (o ApplicationSegmentPRAOutput) MicrotenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.StringOutput { return v.MicrotenantId }).(pulumi.StringOutput)
}

// Name of the application.
func (o ApplicationSegmentPRAOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationSegmentPRAOutput) PassiveHealthEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.BoolOutput { return v.PassiveHealthEnabled }).(pulumi.BoolOutput)
}

func (o ApplicationSegmentPRAOutput) SegmentGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.StringOutput { return v.SegmentGroupId }).(pulumi.StringOutput)
}

func (o ApplicationSegmentPRAOutput) SelectConnectorCloseToApp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.BoolPtrOutput { return v.SelectConnectorCloseToApp }).(pulumi.BoolPtrOutput)
}

func (o ApplicationSegmentPRAOutput) ServerGroups() ApplicationSegmentPRAServerGroupArrayOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) ApplicationSegmentPRAServerGroupArrayOutput { return v.ServerGroups }).(ApplicationSegmentPRAServerGroupArrayOutput)
}

func (o ApplicationSegmentPRAOutput) TcpKeepAlive() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.StringOutput { return v.TcpKeepAlive }).(pulumi.StringOutput)
}

// tcp port range
func (o ApplicationSegmentPRAOutput) TcpPortRange() ApplicationSegmentPRATcpPortRangeArrayOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) ApplicationSegmentPRATcpPortRangeArrayOutput { return v.TcpPortRange }).(ApplicationSegmentPRATcpPortRangeArrayOutput)
}

// TCP port ranges used to access the app.
func (o ApplicationSegmentPRAOutput) TcpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.StringArrayOutput { return v.TcpPortRanges }).(pulumi.StringArrayOutput)
}

// udp port range
func (o ApplicationSegmentPRAOutput) UdpPortRange() ApplicationSegmentPRAUdpPortRangeArrayOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) ApplicationSegmentPRAUdpPortRangeArrayOutput { return v.UdpPortRange }).(ApplicationSegmentPRAUdpPortRangeArrayOutput)
}

// UDP port ranges used to access the app.
func (o ApplicationSegmentPRAOutput) UdpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.StringArrayOutput { return v.UdpPortRanges }).(pulumi.StringArrayOutput)
}

func (o ApplicationSegmentPRAOutput) UseInDrMode() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApplicationSegmentPRA) pulumi.BoolOutput { return v.UseInDrMode }).(pulumi.BoolOutput)
}

type ApplicationSegmentPRAArrayOutput struct{ *pulumi.OutputState }

func (ApplicationSegmentPRAArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationSegmentPRA)(nil)).Elem()
}

func (o ApplicationSegmentPRAArrayOutput) ToApplicationSegmentPRAArrayOutput() ApplicationSegmentPRAArrayOutput {
	return o
}

func (o ApplicationSegmentPRAArrayOutput) ToApplicationSegmentPRAArrayOutputWithContext(ctx context.Context) ApplicationSegmentPRAArrayOutput {
	return o
}

func (o ApplicationSegmentPRAArrayOutput) Index(i pulumi.IntInput) ApplicationSegmentPRAOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationSegmentPRA {
		return vs[0].([]*ApplicationSegmentPRA)[vs[1].(int)]
	}).(ApplicationSegmentPRAOutput)
}

type ApplicationSegmentPRAMapOutput struct{ *pulumi.OutputState }

func (ApplicationSegmentPRAMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationSegmentPRA)(nil)).Elem()
}

func (o ApplicationSegmentPRAMapOutput) ToApplicationSegmentPRAMapOutput() ApplicationSegmentPRAMapOutput {
	return o
}

func (o ApplicationSegmentPRAMapOutput) ToApplicationSegmentPRAMapOutputWithContext(ctx context.Context) ApplicationSegmentPRAMapOutput {
	return o
}

func (o ApplicationSegmentPRAMapOutput) MapIndex(k pulumi.StringInput) ApplicationSegmentPRAOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationSegmentPRA {
		return vs[0].(map[string]*ApplicationSegmentPRA)[vs[1].(string)]
	}).(ApplicationSegmentPRAOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationSegmentPRAInput)(nil)).Elem(), &ApplicationSegmentPRA{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationSegmentPRAArrayInput)(nil)).Elem(), ApplicationSegmentPRAArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationSegmentPRAMapInput)(nil)).Elem(), ApplicationSegmentPRAMap{})
	pulumi.RegisterOutputType(ApplicationSegmentPRAOutput{})
	pulumi.RegisterOutputType(ApplicationSegmentPRAArrayOutput{})
	pulumi.RegisterOutputType(ApplicationSegmentPRAMapOutput{})
}
