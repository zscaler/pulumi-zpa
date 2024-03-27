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

// ## Import
//
// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
//
// # Visit
//
// Application Segment can be imported by using `<APPLICATION SEGMENT ID>` or `<APPLICATION SEGMENT NAME>` as the import ID.
//
// ```sh
// $ pulumi import zpa:index/applicationSegment:ApplicationSegment example <application_segment_id>
// ```
//
// or
//
// ```sh
// $ pulumi import zpa:index/applicationSegment:ApplicationSegment example <application_segment_name>
// ```
type ApplicationSegment struct {
	pulumi.CustomResourceState

	// (Optional) Indicates whether users can bypass ZPA to access applications. Supported values: `ALWAYS`, `NEVER`, `ON_NET`.
	BypassType pulumi.StringOutput `pulumi:"bypassType"`
	// (Optional) Supported values: `DEFAULT`, `SIEM`.
	ConfigSpace pulumi.StringPtrOutput `pulumi:"configSpace"`
	// (Optional) Description of the application.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of domains and IPs.
	DomainNames pulumi.StringArrayOutput `pulumi:"domainNames"`
	// (Optional) Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt pulumi.BoolPtrOutput `pulumi:"doubleEncrypt"`
	// (Optional) Whether this application is enabled or not.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// (Optional) Supported values: `DEFAULT`, `NONE`.
	HealthCheckType pulumi.StringPtrOutput `pulumi:"healthCheckType"`
	// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
	HealthReporting pulumi.StringPtrOutput `pulumi:"healthReporting"`
	// (Optional) Supported values: `NONE`, `PING_TRACEROUTING`, `PING`.
	IcmpAccessType pulumi.StringPtrOutput `pulumi:"icmpAccessType"`
	// (Optional) Supported values: `true`, `false`
	IpAnchored pulumi.BoolPtrOutput `pulumi:"ipAnchored"`
	// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
	IsCnameEnabled pulumi.BoolOutput `pulumi:"isCnameEnabled"`
	// (Optional) Supported values: `true`, `false`
	IsIncompleteDrConfig pulumi.BoolPtrOutput `pulumi:"isIncompleteDrConfig"`
	// (Optional) The ID of the microtenant the resource is to be associated with.
	MicrotenantId pulumi.StringPtrOutput `pulumi:"microtenantId"`
	// Name. The name of the App Connector Group to be exported.
	Name pulumi.StringOutput `pulumi:"name"`
	// (Optional) Supported values: `true`, `false`
	PassiveHealthEnabled pulumi.BoolOutput `pulumi:"passiveHealthEnabled"`
	// List of Segment Group IDs
	SegmentGroupId            pulumi.StringOutput  `pulumi:"segmentGroupId"`
	SegmentGroupName          pulumi.StringOutput  `pulumi:"segmentGroupName"`
	SelectConnectorCloseToApp pulumi.BoolPtrOutput `pulumi:"selectConnectorCloseToApp"`
	// List of Server Group IDs
	ServerGroups ApplicationSegmentServerGroupArrayOutput `pulumi:"serverGroups"`
	// (Optional) Supported values: ``1`` for Enabled and ``0`` for Disabled
	TcpKeepAlive pulumi.StringOutput `pulumi:"tcpKeepAlive"`
	// TCP port ranges used to access the app.
	// * `from:`
	// * `to:`
	TcpPortRange ApplicationSegmentTcpPortRangeArrayOutput `pulumi:"tcpPortRange"`
	// TCP port ranges used to access the app.
	TcpPortRanges pulumi.StringArrayOutput `pulumi:"tcpPortRanges"`
	// UDP port ranges used to access the app.
	// * `from:`
	// * `to:`
	//
	// > **NOTE:** Application segments must have unique ports and cannot have overlapping domain names using the same tcp/udp ports across multiple application segments.
	UdpPortRange ApplicationSegmentUdpPortRangeArrayOutput `pulumi:"udpPortRange"`
	// UDP port ranges used to access the app.
	//
	// > **NOTE:**  TCP and UDP ports can also be defined using the following model:
	// **NOTE:** When removing TCP and/or UDP ports, parameter must be defined but set as empty due to current API behavior.
	UdpPortRanges pulumi.StringArrayOutput `pulumi:"udpPortRanges"`
	// (Optional) Supported values: `true`, `false`
	UseInDrMode pulumi.BoolPtrOutput `pulumi:"useInDrMode"`
}

// NewApplicationSegment registers a new resource with the given unique name, arguments, and options.
func NewApplicationSegment(ctx *pulumi.Context,
	name string, args *ApplicationSegmentArgs, opts ...pulumi.ResourceOption) (*ApplicationSegment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainNames == nil {
		return nil, errors.New("invalid value for required argument 'DomainNames'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationSegment
	err := ctx.RegisterResource("zpa:index/applicationSegment:ApplicationSegment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationSegment gets an existing ApplicationSegment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationSegment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationSegmentState, opts ...pulumi.ResourceOption) (*ApplicationSegment, error) {
	var resource ApplicationSegment
	err := ctx.ReadResource("zpa:index/applicationSegment:ApplicationSegment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationSegment resources.
type applicationSegmentState struct {
	// (Optional) Indicates whether users can bypass ZPA to access applications. Supported values: `ALWAYS`, `NEVER`, `ON_NET`.
	BypassType *string `pulumi:"bypassType"`
	// (Optional) Supported values: `DEFAULT`, `SIEM`.
	ConfigSpace *string `pulumi:"configSpace"`
	// (Optional) Description of the application.
	Description *string `pulumi:"description"`
	// List of domains and IPs.
	DomainNames []string `pulumi:"domainNames"`
	// (Optional) Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt *bool `pulumi:"doubleEncrypt"`
	// (Optional) Whether this application is enabled or not.
	Enabled *bool `pulumi:"enabled"`
	// (Optional) Supported values: `DEFAULT`, `NONE`.
	HealthCheckType *string `pulumi:"healthCheckType"`
	// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
	HealthReporting *string `pulumi:"healthReporting"`
	// (Optional) Supported values: `NONE`, `PING_TRACEROUTING`, `PING`.
	IcmpAccessType *string `pulumi:"icmpAccessType"`
	// (Optional) Supported values: `true`, `false`
	IpAnchored *bool `pulumi:"ipAnchored"`
	// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
	IsCnameEnabled *bool `pulumi:"isCnameEnabled"`
	// (Optional) Supported values: `true`, `false`
	IsIncompleteDrConfig *bool `pulumi:"isIncompleteDrConfig"`
	// (Optional) The ID of the microtenant the resource is to be associated with.
	MicrotenantId *string `pulumi:"microtenantId"`
	// Name. The name of the App Connector Group to be exported.
	Name *string `pulumi:"name"`
	// (Optional) Supported values: `true`, `false`
	PassiveHealthEnabled *bool `pulumi:"passiveHealthEnabled"`
	// List of Segment Group IDs
	SegmentGroupId            *string `pulumi:"segmentGroupId"`
	SegmentGroupName          *string `pulumi:"segmentGroupName"`
	SelectConnectorCloseToApp *bool   `pulumi:"selectConnectorCloseToApp"`
	// List of Server Group IDs
	ServerGroups []ApplicationSegmentServerGroup `pulumi:"serverGroups"`
	// (Optional) Supported values: ``1`` for Enabled and ``0`` for Disabled
	TcpKeepAlive *string `pulumi:"tcpKeepAlive"`
	// TCP port ranges used to access the app.
	// * `from:`
	// * `to:`
	TcpPortRange []ApplicationSegmentTcpPortRange `pulumi:"tcpPortRange"`
	// TCP port ranges used to access the app.
	TcpPortRanges []string `pulumi:"tcpPortRanges"`
	// UDP port ranges used to access the app.
	// * `from:`
	// * `to:`
	//
	// > **NOTE:** Application segments must have unique ports and cannot have overlapping domain names using the same tcp/udp ports across multiple application segments.
	UdpPortRange []ApplicationSegmentUdpPortRange `pulumi:"udpPortRange"`
	// UDP port ranges used to access the app.
	//
	// > **NOTE:**  TCP and UDP ports can also be defined using the following model:
	// **NOTE:** When removing TCP and/or UDP ports, parameter must be defined but set as empty due to current API behavior.
	UdpPortRanges []string `pulumi:"udpPortRanges"`
	// (Optional) Supported values: `true`, `false`
	UseInDrMode *bool `pulumi:"useInDrMode"`
}

type ApplicationSegmentState struct {
	// (Optional) Indicates whether users can bypass ZPA to access applications. Supported values: `ALWAYS`, `NEVER`, `ON_NET`.
	BypassType pulumi.StringPtrInput
	// (Optional) Supported values: `DEFAULT`, `SIEM`.
	ConfigSpace pulumi.StringPtrInput
	// (Optional) Description of the application.
	Description pulumi.StringPtrInput
	// List of domains and IPs.
	DomainNames pulumi.StringArrayInput
	// (Optional) Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt pulumi.BoolPtrInput
	// (Optional) Whether this application is enabled or not.
	Enabled pulumi.BoolPtrInput
	// (Optional) Supported values: `DEFAULT`, `NONE`.
	HealthCheckType pulumi.StringPtrInput
	// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
	HealthReporting pulumi.StringPtrInput
	// (Optional) Supported values: `NONE`, `PING_TRACEROUTING`, `PING`.
	IcmpAccessType pulumi.StringPtrInput
	// (Optional) Supported values: `true`, `false`
	IpAnchored pulumi.BoolPtrInput
	// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
	IsCnameEnabled pulumi.BoolPtrInput
	// (Optional) Supported values: `true`, `false`
	IsIncompleteDrConfig pulumi.BoolPtrInput
	// (Optional) The ID of the microtenant the resource is to be associated with.
	MicrotenantId pulumi.StringPtrInput
	// Name. The name of the App Connector Group to be exported.
	Name pulumi.StringPtrInput
	// (Optional) Supported values: `true`, `false`
	PassiveHealthEnabled pulumi.BoolPtrInput
	// List of Segment Group IDs
	SegmentGroupId            pulumi.StringPtrInput
	SegmentGroupName          pulumi.StringPtrInput
	SelectConnectorCloseToApp pulumi.BoolPtrInput
	// List of Server Group IDs
	ServerGroups ApplicationSegmentServerGroupArrayInput
	// (Optional) Supported values: ``1`` for Enabled and ``0`` for Disabled
	TcpKeepAlive pulumi.StringPtrInput
	// TCP port ranges used to access the app.
	// * `from:`
	// * `to:`
	TcpPortRange ApplicationSegmentTcpPortRangeArrayInput
	// TCP port ranges used to access the app.
	TcpPortRanges pulumi.StringArrayInput
	// UDP port ranges used to access the app.
	// * `from:`
	// * `to:`
	//
	// > **NOTE:** Application segments must have unique ports and cannot have overlapping domain names using the same tcp/udp ports across multiple application segments.
	UdpPortRange ApplicationSegmentUdpPortRangeArrayInput
	// UDP port ranges used to access the app.
	//
	// > **NOTE:**  TCP and UDP ports can also be defined using the following model:
	// **NOTE:** When removing TCP and/or UDP ports, parameter must be defined but set as empty due to current API behavior.
	UdpPortRanges pulumi.StringArrayInput
	// (Optional) Supported values: `true`, `false`
	UseInDrMode pulumi.BoolPtrInput
}

func (ApplicationSegmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationSegmentState)(nil)).Elem()
}

type applicationSegmentArgs struct {
	// (Optional) Indicates whether users can bypass ZPA to access applications. Supported values: `ALWAYS`, `NEVER`, `ON_NET`.
	BypassType *string `pulumi:"bypassType"`
	// (Optional) Supported values: `DEFAULT`, `SIEM`.
	ConfigSpace *string `pulumi:"configSpace"`
	// (Optional) Description of the application.
	Description *string `pulumi:"description"`
	// List of domains and IPs.
	DomainNames []string `pulumi:"domainNames"`
	// (Optional) Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt *bool `pulumi:"doubleEncrypt"`
	// (Optional) Whether this application is enabled or not.
	Enabled *bool `pulumi:"enabled"`
	// (Optional) Supported values: `DEFAULT`, `NONE`.
	HealthCheckType *string `pulumi:"healthCheckType"`
	// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
	HealthReporting *string `pulumi:"healthReporting"`
	// (Optional) Supported values: `NONE`, `PING_TRACEROUTING`, `PING`.
	IcmpAccessType *string `pulumi:"icmpAccessType"`
	// (Optional) Supported values: `true`, `false`
	IpAnchored *bool `pulumi:"ipAnchored"`
	// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
	IsCnameEnabled *bool `pulumi:"isCnameEnabled"`
	// (Optional) Supported values: `true`, `false`
	IsIncompleteDrConfig *bool `pulumi:"isIncompleteDrConfig"`
	// (Optional) The ID of the microtenant the resource is to be associated with.
	MicrotenantId *string `pulumi:"microtenantId"`
	// Name. The name of the App Connector Group to be exported.
	Name *string `pulumi:"name"`
	// (Optional) Supported values: `true`, `false`
	PassiveHealthEnabled *bool `pulumi:"passiveHealthEnabled"`
	// List of Segment Group IDs
	SegmentGroupId            *string `pulumi:"segmentGroupId"`
	SegmentGroupName          *string `pulumi:"segmentGroupName"`
	SelectConnectorCloseToApp *bool   `pulumi:"selectConnectorCloseToApp"`
	// List of Server Group IDs
	ServerGroups []ApplicationSegmentServerGroup `pulumi:"serverGroups"`
	// (Optional) Supported values: ``1`` for Enabled and ``0`` for Disabled
	TcpKeepAlive *string `pulumi:"tcpKeepAlive"`
	// TCP port ranges used to access the app.
	// * `from:`
	// * `to:`
	TcpPortRange []ApplicationSegmentTcpPortRange `pulumi:"tcpPortRange"`
	// TCP port ranges used to access the app.
	TcpPortRanges []string `pulumi:"tcpPortRanges"`
	// UDP port ranges used to access the app.
	// * `from:`
	// * `to:`
	//
	// > **NOTE:** Application segments must have unique ports and cannot have overlapping domain names using the same tcp/udp ports across multiple application segments.
	UdpPortRange []ApplicationSegmentUdpPortRange `pulumi:"udpPortRange"`
	// UDP port ranges used to access the app.
	//
	// > **NOTE:**  TCP and UDP ports can also be defined using the following model:
	// **NOTE:** When removing TCP and/or UDP ports, parameter must be defined but set as empty due to current API behavior.
	UdpPortRanges []string `pulumi:"udpPortRanges"`
	// (Optional) Supported values: `true`, `false`
	UseInDrMode *bool `pulumi:"useInDrMode"`
}

// The set of arguments for constructing a ApplicationSegment resource.
type ApplicationSegmentArgs struct {
	// (Optional) Indicates whether users can bypass ZPA to access applications. Supported values: `ALWAYS`, `NEVER`, `ON_NET`.
	BypassType pulumi.StringPtrInput
	// (Optional) Supported values: `DEFAULT`, `SIEM`.
	ConfigSpace pulumi.StringPtrInput
	// (Optional) Description of the application.
	Description pulumi.StringPtrInput
	// List of domains and IPs.
	DomainNames pulumi.StringArrayInput
	// (Optional) Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt pulumi.BoolPtrInput
	// (Optional) Whether this application is enabled or not.
	Enabled pulumi.BoolPtrInput
	// (Optional) Supported values: `DEFAULT`, `NONE`.
	HealthCheckType pulumi.StringPtrInput
	// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
	HealthReporting pulumi.StringPtrInput
	// (Optional) Supported values: `NONE`, `PING_TRACEROUTING`, `PING`.
	IcmpAccessType pulumi.StringPtrInput
	// (Optional) Supported values: `true`, `false`
	IpAnchored pulumi.BoolPtrInput
	// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
	IsCnameEnabled pulumi.BoolPtrInput
	// (Optional) Supported values: `true`, `false`
	IsIncompleteDrConfig pulumi.BoolPtrInput
	// (Optional) The ID of the microtenant the resource is to be associated with.
	MicrotenantId pulumi.StringPtrInput
	// Name. The name of the App Connector Group to be exported.
	Name pulumi.StringPtrInput
	// (Optional) Supported values: `true`, `false`
	PassiveHealthEnabled pulumi.BoolPtrInput
	// List of Segment Group IDs
	SegmentGroupId            pulumi.StringPtrInput
	SegmentGroupName          pulumi.StringPtrInput
	SelectConnectorCloseToApp pulumi.BoolPtrInput
	// List of Server Group IDs
	ServerGroups ApplicationSegmentServerGroupArrayInput
	// (Optional) Supported values: ``1`` for Enabled and ``0`` for Disabled
	TcpKeepAlive pulumi.StringPtrInput
	// TCP port ranges used to access the app.
	// * `from:`
	// * `to:`
	TcpPortRange ApplicationSegmentTcpPortRangeArrayInput
	// TCP port ranges used to access the app.
	TcpPortRanges pulumi.StringArrayInput
	// UDP port ranges used to access the app.
	// * `from:`
	// * `to:`
	//
	// > **NOTE:** Application segments must have unique ports and cannot have overlapping domain names using the same tcp/udp ports across multiple application segments.
	UdpPortRange ApplicationSegmentUdpPortRangeArrayInput
	// UDP port ranges used to access the app.
	//
	// > **NOTE:**  TCP and UDP ports can also be defined using the following model:
	// **NOTE:** When removing TCP and/or UDP ports, parameter must be defined but set as empty due to current API behavior.
	UdpPortRanges pulumi.StringArrayInput
	// (Optional) Supported values: `true`, `false`
	UseInDrMode pulumi.BoolPtrInput
}

func (ApplicationSegmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationSegmentArgs)(nil)).Elem()
}

type ApplicationSegmentInput interface {
	pulumi.Input

	ToApplicationSegmentOutput() ApplicationSegmentOutput
	ToApplicationSegmentOutputWithContext(ctx context.Context) ApplicationSegmentOutput
}

func (*ApplicationSegment) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationSegment)(nil)).Elem()
}

func (i *ApplicationSegment) ToApplicationSegmentOutput() ApplicationSegmentOutput {
	return i.ToApplicationSegmentOutputWithContext(context.Background())
}

func (i *ApplicationSegment) ToApplicationSegmentOutputWithContext(ctx context.Context) ApplicationSegmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSegmentOutput)
}

// ApplicationSegmentArrayInput is an input type that accepts ApplicationSegmentArray and ApplicationSegmentArrayOutput values.
// You can construct a concrete instance of `ApplicationSegmentArrayInput` via:
//
//	ApplicationSegmentArray{ ApplicationSegmentArgs{...} }
type ApplicationSegmentArrayInput interface {
	pulumi.Input

	ToApplicationSegmentArrayOutput() ApplicationSegmentArrayOutput
	ToApplicationSegmentArrayOutputWithContext(context.Context) ApplicationSegmentArrayOutput
}

type ApplicationSegmentArray []ApplicationSegmentInput

func (ApplicationSegmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationSegment)(nil)).Elem()
}

func (i ApplicationSegmentArray) ToApplicationSegmentArrayOutput() ApplicationSegmentArrayOutput {
	return i.ToApplicationSegmentArrayOutputWithContext(context.Background())
}

func (i ApplicationSegmentArray) ToApplicationSegmentArrayOutputWithContext(ctx context.Context) ApplicationSegmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSegmentArrayOutput)
}

// ApplicationSegmentMapInput is an input type that accepts ApplicationSegmentMap and ApplicationSegmentMapOutput values.
// You can construct a concrete instance of `ApplicationSegmentMapInput` via:
//
//	ApplicationSegmentMap{ "key": ApplicationSegmentArgs{...} }
type ApplicationSegmentMapInput interface {
	pulumi.Input

	ToApplicationSegmentMapOutput() ApplicationSegmentMapOutput
	ToApplicationSegmentMapOutputWithContext(context.Context) ApplicationSegmentMapOutput
}

type ApplicationSegmentMap map[string]ApplicationSegmentInput

func (ApplicationSegmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationSegment)(nil)).Elem()
}

func (i ApplicationSegmentMap) ToApplicationSegmentMapOutput() ApplicationSegmentMapOutput {
	return i.ToApplicationSegmentMapOutputWithContext(context.Background())
}

func (i ApplicationSegmentMap) ToApplicationSegmentMapOutputWithContext(ctx context.Context) ApplicationSegmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSegmentMapOutput)
}

type ApplicationSegmentOutput struct{ *pulumi.OutputState }

func (ApplicationSegmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationSegment)(nil)).Elem()
}

func (o ApplicationSegmentOutput) ToApplicationSegmentOutput() ApplicationSegmentOutput {
	return o
}

func (o ApplicationSegmentOutput) ToApplicationSegmentOutputWithContext(ctx context.Context) ApplicationSegmentOutput {
	return o
}

// (Optional) Indicates whether users can bypass ZPA to access applications. Supported values: `ALWAYS`, `NEVER`, `ON_NET`.
func (o ApplicationSegmentOutput) BypassType() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.StringOutput { return v.BypassType }).(pulumi.StringOutput)
}

// (Optional) Supported values: `DEFAULT`, `SIEM`.
func (o ApplicationSegmentOutput) ConfigSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.StringPtrOutput { return v.ConfigSpace }).(pulumi.StringPtrOutput)
}

// (Optional) Description of the application.
func (o ApplicationSegmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of domains and IPs.
func (o ApplicationSegmentOutput) DomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.StringArrayOutput { return v.DomainNames }).(pulumi.StringArrayOutput)
}

// (Optional) Whether Double Encryption is enabled or disabled for the app.
func (o ApplicationSegmentOutput) DoubleEncrypt() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.BoolPtrOutput { return v.DoubleEncrypt }).(pulumi.BoolPtrOutput)
}

// (Optional) Whether this application is enabled or not.
func (o ApplicationSegmentOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// (Optional) Supported values: `DEFAULT`, `NONE`.
func (o ApplicationSegmentOutput) HealthCheckType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.StringPtrOutput { return v.HealthCheckType }).(pulumi.StringPtrOutput)
}

// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
func (o ApplicationSegmentOutput) HealthReporting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.StringPtrOutput { return v.HealthReporting }).(pulumi.StringPtrOutput)
}

// (Optional) Supported values: `NONE`, `PING_TRACEROUTING`, `PING`.
func (o ApplicationSegmentOutput) IcmpAccessType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.StringPtrOutput { return v.IcmpAccessType }).(pulumi.StringPtrOutput)
}

// (Optional) Supported values: `true`, `false`
func (o ApplicationSegmentOutput) IpAnchored() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.BoolPtrOutput { return v.IpAnchored }).(pulumi.BoolPtrOutput)
}

// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
func (o ApplicationSegmentOutput) IsCnameEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.BoolOutput { return v.IsCnameEnabled }).(pulumi.BoolOutput)
}

// (Optional) Supported values: `true`, `false`
func (o ApplicationSegmentOutput) IsIncompleteDrConfig() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.BoolPtrOutput { return v.IsIncompleteDrConfig }).(pulumi.BoolPtrOutput)
}

// (Optional) The ID of the microtenant the resource is to be associated with.
func (o ApplicationSegmentOutput) MicrotenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.StringPtrOutput { return v.MicrotenantId }).(pulumi.StringPtrOutput)
}

// Name. The name of the App Connector Group to be exported.
func (o ApplicationSegmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// (Optional) Supported values: `true`, `false`
func (o ApplicationSegmentOutput) PassiveHealthEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.BoolOutput { return v.PassiveHealthEnabled }).(pulumi.BoolOutput)
}

// List of Segment Group IDs
func (o ApplicationSegmentOutput) SegmentGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.StringOutput { return v.SegmentGroupId }).(pulumi.StringOutput)
}

func (o ApplicationSegmentOutput) SegmentGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.StringOutput { return v.SegmentGroupName }).(pulumi.StringOutput)
}

func (o ApplicationSegmentOutput) SelectConnectorCloseToApp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.BoolPtrOutput { return v.SelectConnectorCloseToApp }).(pulumi.BoolPtrOutput)
}

// List of Server Group IDs
func (o ApplicationSegmentOutput) ServerGroups() ApplicationSegmentServerGroupArrayOutput {
	return o.ApplyT(func(v *ApplicationSegment) ApplicationSegmentServerGroupArrayOutput { return v.ServerGroups }).(ApplicationSegmentServerGroupArrayOutput)
}

// (Optional) Supported values: “1“ for Enabled and “0“ for Disabled
func (o ApplicationSegmentOutput) TcpKeepAlive() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.StringOutput { return v.TcpKeepAlive }).(pulumi.StringOutput)
}

// TCP port ranges used to access the app.
// * `from:`
// * `to:`
func (o ApplicationSegmentOutput) TcpPortRange() ApplicationSegmentTcpPortRangeArrayOutput {
	return o.ApplyT(func(v *ApplicationSegment) ApplicationSegmentTcpPortRangeArrayOutput { return v.TcpPortRange }).(ApplicationSegmentTcpPortRangeArrayOutput)
}

// TCP port ranges used to access the app.
func (o ApplicationSegmentOutput) TcpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.StringArrayOutput { return v.TcpPortRanges }).(pulumi.StringArrayOutput)
}

// UDP port ranges used to access the app.
// * `from:`
// * `to:`
//
// > **NOTE:** Application segments must have unique ports and cannot have overlapping domain names using the same tcp/udp ports across multiple application segments.
func (o ApplicationSegmentOutput) UdpPortRange() ApplicationSegmentUdpPortRangeArrayOutput {
	return o.ApplyT(func(v *ApplicationSegment) ApplicationSegmentUdpPortRangeArrayOutput { return v.UdpPortRange }).(ApplicationSegmentUdpPortRangeArrayOutput)
}

// UDP port ranges used to access the app.
//
// > **NOTE:**  TCP and UDP ports can also be defined using the following model:
// **NOTE:** When removing TCP and/or UDP ports, parameter must be defined but set as empty due to current API behavior.
func (o ApplicationSegmentOutput) UdpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.StringArrayOutput { return v.UdpPortRanges }).(pulumi.StringArrayOutput)
}

// (Optional) Supported values: `true`, `false`
func (o ApplicationSegmentOutput) UseInDrMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationSegment) pulumi.BoolPtrOutput { return v.UseInDrMode }).(pulumi.BoolPtrOutput)
}

type ApplicationSegmentArrayOutput struct{ *pulumi.OutputState }

func (ApplicationSegmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationSegment)(nil)).Elem()
}

func (o ApplicationSegmentArrayOutput) ToApplicationSegmentArrayOutput() ApplicationSegmentArrayOutput {
	return o
}

func (o ApplicationSegmentArrayOutput) ToApplicationSegmentArrayOutputWithContext(ctx context.Context) ApplicationSegmentArrayOutput {
	return o
}

func (o ApplicationSegmentArrayOutput) Index(i pulumi.IntInput) ApplicationSegmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationSegment {
		return vs[0].([]*ApplicationSegment)[vs[1].(int)]
	}).(ApplicationSegmentOutput)
}

type ApplicationSegmentMapOutput struct{ *pulumi.OutputState }

func (ApplicationSegmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationSegment)(nil)).Elem()
}

func (o ApplicationSegmentMapOutput) ToApplicationSegmentMapOutput() ApplicationSegmentMapOutput {
	return o
}

func (o ApplicationSegmentMapOutput) ToApplicationSegmentMapOutputWithContext(ctx context.Context) ApplicationSegmentMapOutput {
	return o
}

func (o ApplicationSegmentMapOutput) MapIndex(k pulumi.StringInput) ApplicationSegmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationSegment {
		return vs[0].(map[string]*ApplicationSegment)[vs[1].(string)]
	}).(ApplicationSegmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationSegmentInput)(nil)).Elem(), &ApplicationSegment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationSegmentArrayInput)(nil)).Elem(), ApplicationSegmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationSegmentMapInput)(nil)).Elem(), ApplicationSegmentMap{})
	pulumi.RegisterOutputType(ApplicationSegmentOutput{})
	pulumi.RegisterOutputType(ApplicationSegmentArrayOutput{})
	pulumi.RegisterOutputType(ApplicationSegmentMapOutput{})
}