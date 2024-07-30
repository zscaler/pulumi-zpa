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
	// - (String) This field defines the id of the application server.
	Id *string `pulumi:"id"`
	// - (String) This field defines the name of the server.
	Name         *string                                          `pulumi:"name"`
	TcpPortRange []GetApplicationSegmentBrowserAccessTcpPortRange `pulumi:"tcpPortRange"`
	UdpPortRange []GetApplicationSegmentBrowserAccessUdpPortRange `pulumi:"udpPortRange"`
}

// A collection of values returned by getApplicationSegmentBrowserAccess.
type LookupApplicationSegmentBrowserAccessResult struct {
	BypassType      string                                            `pulumi:"bypassType"`
	ClientlessApps  []GetApplicationSegmentBrowserAccessClientlessApp `pulumi:"clientlessApps"`
	ConfigSpace     string                                            `pulumi:"configSpace"`
	Description     string                                            `pulumi:"description"`
	DomainNames     []string                                          `pulumi:"domainNames"`
	DoubleEncrypt   bool                                              `pulumi:"doubleEncrypt"`
	Enabled         bool                                              `pulumi:"enabled"`
	HealthCheckType string                                            `pulumi:"healthCheckType"`
	HealthReporting string                                            `pulumi:"healthReporting"`
	// - (String) This field defines the id of the application server.
	Id             *string `pulumi:"id"`
	IpAnchored     bool    `pulumi:"ipAnchored"`
	IsCnameEnabled bool    `pulumi:"isCnameEnabled"`
	// - (String) This field defines the name of the server.
	Name                 *string                                          `pulumi:"name"`
	PassiveHealthEnabled bool                                             `pulumi:"passiveHealthEnabled"`
	SegmentGroupId       string                                           `pulumi:"segmentGroupId"`
	SegmentGroupName     string                                           `pulumi:"segmentGroupName"`
	ServerGroups         []GetApplicationSegmentBrowserAccessServerGroup  `pulumi:"serverGroups"`
	TcpPortRange         []GetApplicationSegmentBrowserAccessTcpPortRange `pulumi:"tcpPortRange"`
	TcpPortRanges        []string                                         `pulumi:"tcpPortRanges"`
	UdpPortRange         []GetApplicationSegmentBrowserAccessUdpPortRange `pulumi:"udpPortRange"`
	UdpPortRanges        []string                                         `pulumi:"udpPortRanges"`
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
	// - (String) This field defines the id of the application server.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// - (String) This field defines the name of the server.
	Name         pulumi.StringPtrInput                                    `pulumi:"name"`
	TcpPortRange GetApplicationSegmentBrowserAccessTcpPortRangeArrayInput `pulumi:"tcpPortRange"`
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

func (o LookupApplicationSegmentBrowserAccessResultOutput) BypassType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) string { return v.BypassType }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) ClientlessApps() GetApplicationSegmentBrowserAccessClientlessAppArrayOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) []GetApplicationSegmentBrowserAccessClientlessApp {
		return v.ClientlessApps
	}).(GetApplicationSegmentBrowserAccessClientlessAppArrayOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) ConfigSpace() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) string { return v.ConfigSpace }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) DomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) []string { return v.DomainNames }).(pulumi.StringArrayOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) DoubleEncrypt() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) bool { return v.DoubleEncrypt }).(pulumi.BoolOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) HealthCheckType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) string { return v.HealthCheckType }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) HealthReporting() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) string { return v.HealthReporting }).(pulumi.StringOutput)
}

// - (String) This field defines the id of the application server.
func (o LookupApplicationSegmentBrowserAccessResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) IpAnchored() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) bool { return v.IpAnchored }).(pulumi.BoolOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) IsCnameEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) bool { return v.IsCnameEnabled }).(pulumi.BoolOutput)
}

// - (String) This field defines the name of the server.
func (o LookupApplicationSegmentBrowserAccessResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) PassiveHealthEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) bool { return v.PassiveHealthEnabled }).(pulumi.BoolOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) SegmentGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) string { return v.SegmentGroupId }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) SegmentGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) string { return v.SegmentGroupName }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) ServerGroups() GetApplicationSegmentBrowserAccessServerGroupArrayOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) []GetApplicationSegmentBrowserAccessServerGroup {
		return v.ServerGroups
	}).(GetApplicationSegmentBrowserAccessServerGroupArrayOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) TcpPortRange() GetApplicationSegmentBrowserAccessTcpPortRangeArrayOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) []GetApplicationSegmentBrowserAccessTcpPortRange {
		return v.TcpPortRange
	}).(GetApplicationSegmentBrowserAccessTcpPortRangeArrayOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) TcpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) []string { return v.TcpPortRanges }).(pulumi.StringArrayOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) UdpPortRange() GetApplicationSegmentBrowserAccessUdpPortRangeArrayOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) []GetApplicationSegmentBrowserAccessUdpPortRange {
		return v.UdpPortRange
	}).(GetApplicationSegmentBrowserAccessUdpPortRangeArrayOutput)
}

func (o LookupApplicationSegmentBrowserAccessResultOutput) UdpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationSegmentBrowserAccessResult) []string { return v.UdpPortRanges }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationSegmentBrowserAccessResultOutput{})
}
