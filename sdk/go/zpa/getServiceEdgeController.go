// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// * [Official documentation](https://help.zscaler.com/zpa/about-zpa-private-service-edges)
// * [API documentation](https://help.zscaler.com/zpa/managing-zpa-private-service-edges-using-api)
//
// Use the **zpa_service_edge_controller** data source to get information about a service edge controller in the Zscaler Private Access cloud. This data source can then be referenced in a Service Edge Group and Provisioning Key.
//
// **NOTE:** To ensure consistent search results across data sources, please avoid using multiple spaces or special characters in your search queries.
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
//			_, err := zpa.GetServiceEdgeController(ctx, &zpa.GetServiceEdgeControllerArgs{
//				Name: pulumi.StringRef("On-Prem-PSE"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetServiceEdgeController(ctx *pulumi.Context, args *GetServiceEdgeControllerArgs, opts ...pulumi.InvokeOption) (*GetServiceEdgeControllerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServiceEdgeControllerResult
	err := ctx.Invoke("zpa:index/getServiceEdgeController:getServiceEdgeController", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceEdgeController.
type GetServiceEdgeControllerArgs struct {
	MicrotenantId   *string `pulumi:"microtenantId"`
	MicrotenantName *string `pulumi:"microtenantName"`
	Name            *string `pulumi:"name"`
}

// A collection of values returned by getServiceEdgeController.
type GetServiceEdgeControllerResult struct {
	ApplicationStartTime             string                                         `pulumi:"applicationStartTime"`
	ControlChannelStatus             string                                         `pulumi:"controlChannelStatus"`
	CreationTime                     string                                         `pulumi:"creationTime"`
	CtrlBrokerName                   string                                         `pulumi:"ctrlBrokerName"`
	CurrentVersion                   string                                         `pulumi:"currentVersion"`
	Description                      string                                         `pulumi:"description"`
	Enabled                          bool                                           `pulumi:"enabled"`
	EnrollmentCert                   map[string]string                              `pulumi:"enrollmentCert"`
	ExpectedUpgradeTime              string                                         `pulumi:"expectedUpgradeTime"`
	ExpectedVersion                  string                                         `pulumi:"expectedVersion"`
	Fingerprint                      string                                         `pulumi:"fingerprint"`
	Id                               string                                         `pulumi:"id"`
	IpAcl                            string                                         `pulumi:"ipAcl"`
	IssuedCertId                     string                                         `pulumi:"issuedCertId"`
	LastBrokerConnectTime            string                                         `pulumi:"lastBrokerConnectTime"`
	LastBrokerConnectTimeDuration    string                                         `pulumi:"lastBrokerConnectTimeDuration"`
	LastBrokerDisconnectTime         string                                         `pulumi:"lastBrokerDisconnectTime"`
	LastBrokerDisconnectTimeDuration string                                         `pulumi:"lastBrokerDisconnectTimeDuration"`
	LastUpgradeTime                  string                                         `pulumi:"lastUpgradeTime"`
	Latitude                         string                                         `pulumi:"latitude"`
	ListenIps                        []string                                       `pulumi:"listenIps"`
	Location                         string                                         `pulumi:"location"`
	Longitude                        string                                         `pulumi:"longitude"`
	MicrotenantId                    *string                                        `pulumi:"microtenantId"`
	MicrotenantName                  *string                                        `pulumi:"microtenantName"`
	ModifiedBy                       string                                         `pulumi:"modifiedBy"`
	ModifiedTime                     string                                         `pulumi:"modifiedTime"`
	Name                             *string                                        `pulumi:"name"`
	Platform                         string                                         `pulumi:"platform"`
	PreviousVersion                  string                                         `pulumi:"previousVersion"`
	PrivateBrokerVersions            []GetServiceEdgeControllerPrivateBrokerVersion `pulumi:"privateBrokerVersions"`
	PrivateIp                        string                                         `pulumi:"privateIp"`
	ProvisioningKeyId                string                                         `pulumi:"provisioningKeyId"`
	ProvisioningKeyName              string                                         `pulumi:"provisioningKeyName"`
	PublicIp                         string                                         `pulumi:"publicIp"`
	PublishIps                       []string                                       `pulumi:"publishIps"`
	PublishIpv6                      bool                                           `pulumi:"publishIpv6"`
	RuntimeOs                        string                                         `pulumi:"runtimeOs"`
	SargeVersion                     string                                         `pulumi:"sargeVersion"`
	ServiceEdgeGroupId               string                                         `pulumi:"serviceEdgeGroupId"`
	ServiceEdgeGroupName             string                                         `pulumi:"serviceEdgeGroupName"`
	UpgradeAttempt                   string                                         `pulumi:"upgradeAttempt"`
	UpgradeStatus                    string                                         `pulumi:"upgradeStatus"`
}

func GetServiceEdgeControllerOutput(ctx *pulumi.Context, args GetServiceEdgeControllerOutputArgs, opts ...pulumi.InvokeOption) GetServiceEdgeControllerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetServiceEdgeControllerResultOutput, error) {
			args := v.(GetServiceEdgeControllerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("zpa:index/getServiceEdgeController:getServiceEdgeController", args, GetServiceEdgeControllerResultOutput{}, options).(GetServiceEdgeControllerResultOutput), nil
		}).(GetServiceEdgeControllerResultOutput)
}

// A collection of arguments for invoking getServiceEdgeController.
type GetServiceEdgeControllerOutputArgs struct {
	MicrotenantId   pulumi.StringPtrInput `pulumi:"microtenantId"`
	MicrotenantName pulumi.StringPtrInput `pulumi:"microtenantName"`
	Name            pulumi.StringPtrInput `pulumi:"name"`
}

func (GetServiceEdgeControllerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceEdgeControllerArgs)(nil)).Elem()
}

// A collection of values returned by getServiceEdgeController.
type GetServiceEdgeControllerResultOutput struct{ *pulumi.OutputState }

func (GetServiceEdgeControllerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceEdgeControllerResult)(nil)).Elem()
}

func (o GetServiceEdgeControllerResultOutput) ToGetServiceEdgeControllerResultOutput() GetServiceEdgeControllerResultOutput {
	return o
}

func (o GetServiceEdgeControllerResultOutput) ToGetServiceEdgeControllerResultOutputWithContext(ctx context.Context) GetServiceEdgeControllerResultOutput {
	return o
}

func (o GetServiceEdgeControllerResultOutput) ApplicationStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.ApplicationStartTime }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) ControlChannelStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.ControlChannelStatus }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) CtrlBrokerName() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.CtrlBrokerName }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) CurrentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.CurrentVersion }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o GetServiceEdgeControllerResultOutput) EnrollmentCert() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) map[string]string { return v.EnrollmentCert }).(pulumi.StringMapOutput)
}

func (o GetServiceEdgeControllerResultOutput) ExpectedUpgradeTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.ExpectedUpgradeTime }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) ExpectedVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.ExpectedVersion }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) IpAcl() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.IpAcl }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) IssuedCertId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.IssuedCertId }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) LastBrokerConnectTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.LastBrokerConnectTime }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) LastBrokerConnectTimeDuration() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.LastBrokerConnectTimeDuration }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) LastBrokerDisconnectTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.LastBrokerDisconnectTime }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) LastBrokerDisconnectTimeDuration() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.LastBrokerDisconnectTimeDuration }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) LastUpgradeTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.LastUpgradeTime }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) Latitude() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.Latitude }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) ListenIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) []string { return v.ListenIps }).(pulumi.StringArrayOutput)
}

func (o GetServiceEdgeControllerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) Longitude() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.Longitude }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) MicrotenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) *string { return v.MicrotenantId }).(pulumi.StringPtrOutput)
}

func (o GetServiceEdgeControllerResultOutput) MicrotenantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) *string { return v.MicrotenantName }).(pulumi.StringPtrOutput)
}

func (o GetServiceEdgeControllerResultOutput) ModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.ModifiedBy }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetServiceEdgeControllerResultOutput) Platform() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.Platform }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) PreviousVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.PreviousVersion }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) PrivateBrokerVersions() GetServiceEdgeControllerPrivateBrokerVersionArrayOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) []GetServiceEdgeControllerPrivateBrokerVersion {
		return v.PrivateBrokerVersions
	}).(GetServiceEdgeControllerPrivateBrokerVersionArrayOutput)
}

func (o GetServiceEdgeControllerResultOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.PrivateIp }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) ProvisioningKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.ProvisioningKeyId }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) ProvisioningKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.ProvisioningKeyName }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.PublicIp }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) PublishIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) []string { return v.PublishIps }).(pulumi.StringArrayOutput)
}

func (o GetServiceEdgeControllerResultOutput) PublishIpv6() pulumi.BoolOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) bool { return v.PublishIpv6 }).(pulumi.BoolOutput)
}

func (o GetServiceEdgeControllerResultOutput) RuntimeOs() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.RuntimeOs }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) SargeVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.SargeVersion }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) ServiceEdgeGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.ServiceEdgeGroupId }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) ServiceEdgeGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.ServiceEdgeGroupName }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) UpgradeAttempt() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.UpgradeAttempt }).(pulumi.StringOutput)
}

func (o GetServiceEdgeControllerResultOutput) UpgradeStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceEdgeControllerResult) string { return v.UpgradeStatus }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServiceEdgeControllerResultOutput{})
}
