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
//			_, err := zpa.LookupApplicationSegmentBrowserAccess(ctx, &zpa.LookupApplicationSegmentBrowserAccessArgs{
//				Name: pulumi.StringRef("example"),
//			}, nil)
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
//			_, err := zpa.LookupApplicationSegmentBrowserAccess(ctx, &zpa.LookupApplicationSegmentBrowserAccessArgs{
//				Id: pulumi.StringRef("123456789"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupApplicationSegmentBrowserAccess(ctx *pulumi.Context, args *LookupApplicationSegmentBrowserAccessArgs, opts ...pulumi.InvokeOption) (*LookupApplicationSegmentBrowserAccessResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupApplicationSegmentBrowserAccessResult
	err := ctx.Invoke("zpa:index/getApplicationSegmentBrowserAccess:getApplicationSegmentBrowserAccess", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplicationSegmentBrowserAccess.
type LookupApplicationSegmentBrowserAccessArgs struct {
	// This field defines the id of the application server.
	Id *string `pulumi:"id"`
	// This field defines the name of the server.
	Name *string `pulumi:"name"`
	// (string) TCP port ranges used to access the app.
	// * `from:`
	// * `to:`
	TcpPortRange []GetApplicationSegmentBrowserAccessTcpPortRange `pulumi:"tcpPortRange"`
	// (string) UDP port ranges used to access the app.
	// * `from:`
	// * `to:`
	UdpPortRange []GetApplicationSegmentBrowserAccessUdpPortRange `pulumi:"udpPortRange"`
}

// A collection of values returned by getApplicationSegmentBrowserAccess.
type LookupApplicationSegmentBrowserAccessResult struct {
	// (string) Indicates whether users can bypass ZPA to access applications. Default: `NEVER`. Supported values: `ALWAYS`, `NEVER`, `ON_NET`. The value `NEVER` indicates the use of the client forwarding policy.
	BypassType     string                                            `pulumi:"bypassType"`
	ClientlessApps []GetApplicationSegmentBrowserAccessClientlessApp `pulumi:"clientlessApps"`
	// (string)
	ConfigSpace string `pulumi:"configSpace"`
	// (string)
	Description string `pulumi:"description"`
	// List of domains and IPs.
	DomainNames []string `pulumi:"domainNames"`
	// (string) Whether Double Encryption is enabled or disabled for the app. Default: false. Boolean: `true`, `false`.
	DoubleEncrypt bool `pulumi:"doubleEncrypt"`
	// (bool)
	Enabled         bool   `pulumi:"enabled"`
	HealthCheckType string `pulumi:"healthCheckType"`
	// (string)
	HealthReporting string  `pulumi:"healthReporting"`
	Id              *string `pulumi:"id"`
	// (bool)
	IpAnchored bool `pulumi:"ipAnchored"`
	// (bool) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors. Default: true. Boolean: `true`, `false`.
	IsCnameEnabled bool `pulumi:"isCnameEnabled"`
	// (string)
	Name *string `pulumi:"name"`
	// (bool)
	PassiveHealthEnabled bool `pulumi:"passiveHealthEnabled"`
	// (string)
	SegmentGroupId string `pulumi:"segmentGroupId"`
	// (string)
	SegmentGroupName string                                          `pulumi:"segmentGroupName"`
	ServerGroups     []GetApplicationSegmentBrowserAccessServerGroup `pulumi:"serverGroups"`
	// (string) TCP port ranges used to access the app.
	// * `from:`
	// * `to:`
	TcpPortRange []GetApplicationSegmentBrowserAccessTcpPortRange `pulumi:"tcpPortRange"`
	// (string) TCP port ranges used to access the app.
	TcpPortRanges []string `pulumi:"tcpPortRanges"`
	// (string) UDP port ranges used to access the app.
	// * `from:`
	// * `to:`
	UdpPortRange []GetApplicationSegmentBrowserAccessUdpPortRange `pulumi:"udpPortRange"`
	// (string) UDP port ranges used to access the app.
	UdpPortRanges []string `pulumi:"udpPortRanges"`
}

func LookupApplicationSegmentBrowserAccessOutput(ctx *pulumi.Context, args LookupApplicationSegmentBrowserAccessOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationSegmentBrowserAccessResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationSegmentBrowserAccessResult, error) {
			args := v.(LookupApplicationSegmentBrowserAccessArgs)
			r, err := LookupApplicationSegmentBrowserAccess(ctx, &args, opts...)
			var s LookupApplicationSegmentBrowserAccessResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationSegmentBrowserAccessResultOutput)
}

// A collection of arguments for invoking getApplicationSegmentBrowserAccess.
type LookupApplicationSegmentBrowserAccessOutputArgs struct {
	// This field defines the id of the application server.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// This field defines the name of the server.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// (string) TCP port ranges used to access the app.
	// * `from:`
	// * `to:`
	TcpPortRange GetApplicationSegmentBrowserAccessTcpPortRangeArrayInput `pulumi:"tcpPortRange"`
	// (string) UDP port ranges used to access the app.
	// * `from:`
	// * `to:`
	UdpPortRange GetApplicationSegmentBrowserAccessUdpPortRangeArrayInput `pulumi:"udpPortRange"`
}

func (LookupApplicationSegmentBrowserAccessOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationSegmentBrowserAccessArgs)(nil)).Elem()
}

// A collection of values returned by getApplicationSegmentBrowserAccess.
type LookupApplicationSegmentBrowserAccessResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationSegmentBrowserAccessResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationSegmentBrowserAccessResult)(nil)).Elem()
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) ToLookupApplicationSegmentBrowserAccessResultOutput() LookupApplicationSegmentBrowserAccessResultOutput {
	return o
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) ToLookupApplicationSegmentBrowserAccessResultOutputWithContext(ctx context.Context) LookupApplicationSegmentBrowserAccessResultOutput {
	return o
}

// (string) Indicates whether users can bypass ZPA to access applications. Default: `NEVER`. Supported values: `ALWAYS`, `NEVER`, `ON_NET`. The value `NEVER` indicates the use of the client forwarding policy.
func (o LookupApplicationSegmentBrowserAccessResultOutput) BypassType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) string { return v.BypassType }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) ClientlessApps() GetApplicationSegmentBrowserAccessClientlessAppArrayOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) []GetApplicationSegmentBrowserAccessClientlessApp {
		return v.ClientlessApps
	}).(GetApplicationSegmentBrowserAccessClientlessAppArrayOutput)
}

// (string)
func (o LookupApplicationSegmentBrowserAccessResultOutput) ConfigSpace() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) string { return v.ConfigSpace }).(pulumi.StringOutput)
}

// (string)
func (o LookupApplicationSegmentBrowserAccessResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) string { return v.Description }).(pulumi.StringOutput)
}

// List of domains and IPs.
func (o LookupApplicationSegmentBrowserAccessResultOutput) DomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) []string { return v.DomainNames }).(pulumi.StringArrayOutput)
}

// (string) Whether Double Encryption is enabled or disabled for the app. Default: false. Boolean: `true`, `false`.
func (o LookupApplicationSegmentBrowserAccessResultOutput) DoubleEncrypt() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) bool { return v.DoubleEncrypt }).(pulumi.BoolOutput)
}

// (bool)
func (o LookupApplicationSegmentBrowserAccessResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) HealthCheckType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) string { return v.HealthCheckType }).(pulumi.StringOutput)
}

// (string)
func (o LookupApplicationSegmentBrowserAccessResultOutput) HealthReporting() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) string { return v.HealthReporting }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// (bool)
func (o LookupApplicationSegmentBrowserAccessResultOutput) IpAnchored() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) bool { return v.IpAnchored }).(pulumi.BoolOutput)
}

// (bool) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors. Default: true. Boolean: `true`, `false`.
func (o LookupApplicationSegmentBrowserAccessResultOutput) IsCnameEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) bool { return v.IsCnameEnabled }).(pulumi.BoolOutput)
}

// (string)
func (o LookupApplicationSegmentBrowserAccessResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// (bool)
func (o LookupApplicationSegmentBrowserAccessResultOutput) PassiveHealthEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) bool { return v.PassiveHealthEnabled }).(pulumi.BoolOutput)
}

// (string)
func (o LookupApplicationSegmentBrowserAccessResultOutput) SegmentGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) string { return v.SegmentGroupId }).(pulumi.StringOutput)
}

// (string)
func (o LookupApplicationSegmentBrowserAccessResultOutput) SegmentGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) string { return v.SegmentGroupName }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) ServerGroups() GetApplicationSegmentBrowserAccessServerGroupArrayOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) []GetApplicationSegmentBrowserAccessServerGroup {
		return v.ServerGroups
	}).(GetApplicationSegmentBrowserAccessServerGroupArrayOutput)
}

// (string) TCP port ranges used to access the app.
// * `from:`
// * `to:`
func (o LookupApplicationSegmentBrowserAccessResultOutput) TcpPortRange() GetApplicationSegmentBrowserAccessTcpPortRangeArrayOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) []GetApplicationSegmentBrowserAccessTcpPortRange {
		return v.TcpPortRange
	}).(GetApplicationSegmentBrowserAccessTcpPortRangeArrayOutput)
}

// (string) TCP port ranges used to access the app.
func (o LookupApplicationSegmentBrowserAccessResultOutput) TcpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) []string { return v.TcpPortRanges }).(pulumi.StringArrayOutput)
}

// (string) UDP port ranges used to access the app.
// * `from:`
// * `to:`
func (o LookupApplicationSegmentBrowserAccessResultOutput) UdpPortRange() GetApplicationSegmentBrowserAccessUdpPortRangeArrayOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) []GetApplicationSegmentBrowserAccessUdpPortRange {
		return v.UdpPortRange
	}).(GetApplicationSegmentBrowserAccessUdpPortRangeArrayOutput)
}

// (string) UDP port ranges used to access the app.
func (o LookupApplicationSegmentBrowserAccessResultOutput) UdpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) []string { return v.UdpPortRanges }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationSegmentBrowserAccessResultOutput{})
}
