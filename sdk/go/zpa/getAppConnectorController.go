// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// * [Official documentation](https://help.zscaler.com/zpa/about-connectors)
// * [API documentation](https://help.zscaler.com/zpa/managing-app-connectors-using-api)
//
// Use the **zpa_app_connector_controller** data source to get information about a app connector created in the Zscaler Private Access cloud. This data source can then be referenced in an App Connector Group.
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
//			_, err := zpa.GetAppConnectorController(ctx, &zpa.GetAppConnectorControllerArgs{
//				Name: pulumi.StringRef("AWS-VPC100-App-Connector"),
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
func GetAppConnectorController(ctx *pulumi.Context, args *GetAppConnectorControllerArgs, opts ...pulumi.InvokeOption) (*GetAppConnectorControllerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAppConnectorControllerResult
	err := ctx.Invoke("zpa:index/getAppConnectorController:getAppConnectorController", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppConnectorController.
type GetAppConnectorControllerArgs struct {
	MicrotenantId   *string `pulumi:"microtenantId"`
	MicrotenantName *string `pulumi:"microtenantName"`
	Name            *string `pulumi:"name"`
}

// A collection of values returned by getAppConnectorController.
type GetAppConnectorControllerResult struct {
	AppConnectorGroupId              string                 `pulumi:"appConnectorGroupId"`
	AppConnectorGroupName            string                 `pulumi:"appConnectorGroupName"`
	ApplicationStartTime             string                 `pulumi:"applicationStartTime"`
	ControlChannelStatus             string                 `pulumi:"controlChannelStatus"`
	CreationTime                     string                 `pulumi:"creationTime"`
	CtrlBrokerName                   string                 `pulumi:"ctrlBrokerName"`
	CurrentVersion                   string                 `pulumi:"currentVersion"`
	Description                      string                 `pulumi:"description"`
	Enabled                          bool                   `pulumi:"enabled"`
	EnrollmentCert                   map[string]interface{} `pulumi:"enrollmentCert"`
	ExpectedUpgradeTime              string                 `pulumi:"expectedUpgradeTime"`
	ExpectedVersion                  string                 `pulumi:"expectedVersion"`
	Fingerprint                      string                 `pulumi:"fingerprint"`
	Id                               string                 `pulumi:"id"`
	IpAcl                            string                 `pulumi:"ipAcl"`
	IssuedCertId                     string                 `pulumi:"issuedCertId"`
	LastBrokerConnectTime            string                 `pulumi:"lastBrokerConnectTime"`
	LastBrokerConnectTimeDuration    string                 `pulumi:"lastBrokerConnectTimeDuration"`
	LastBrokerDisconnectTime         string                 `pulumi:"lastBrokerDisconnectTime"`
	LastBrokerDisconnectTimeDuration string                 `pulumi:"lastBrokerDisconnectTimeDuration"`
	LastUpgradeTime                  string                 `pulumi:"lastUpgradeTime"`
	Latitude                         string                 `pulumi:"latitude"`
	Location                         string                 `pulumi:"location"`
	Longitude                        string                 `pulumi:"longitude"`
	MicrotenantId                    *string                `pulumi:"microtenantId"`
	MicrotenantName                  *string                `pulumi:"microtenantName"`
	ModifiedTime                     string                 `pulumi:"modifiedTime"`
	Modifiedby                       string                 `pulumi:"modifiedby"`
	Name                             *string                `pulumi:"name"`
	Platform                         string                 `pulumi:"platform"`
	PreviousVersion                  string                 `pulumi:"previousVersion"`
	PrivateIp                        string                 `pulumi:"privateIp"`
	ProvisioningKeyId                string                 `pulumi:"provisioningKeyId"`
	ProvisioningKeyName              string                 `pulumi:"provisioningKeyName"`
	PublicIp                         string                 `pulumi:"publicIp"`
	SargeVersion                     string                 `pulumi:"sargeVersion"`
	UpgradeAttempt                   string                 `pulumi:"upgradeAttempt"`
	UpgradeStatus                    string                 `pulumi:"upgradeStatus"`
}

func GetAppConnectorControllerOutput(ctx *pulumi.Context, args GetAppConnectorControllerOutputArgs, opts ...pulumi.InvokeOption) GetAppConnectorControllerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppConnectorControllerResult, error) {
			args := v.(GetAppConnectorControllerArgs)
			r, err := GetAppConnectorController(ctx, &args, opts...)
			var s GetAppConnectorControllerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAppConnectorControllerResultOutput)
}

// A collection of arguments for invoking getAppConnectorController.
type GetAppConnectorControllerOutputArgs struct {
	MicrotenantId   pulumi.StringPtrInput `pulumi:"microtenantId"`
	MicrotenantName pulumi.StringPtrInput `pulumi:"microtenantName"`
	Name            pulumi.StringPtrInput `pulumi:"name"`
}

func (GetAppConnectorControllerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppConnectorControllerArgs)(nil)).Elem()
}

// A collection of values returned by getAppConnectorController.
type GetAppConnectorControllerResultOutput struct{ *pulumi.OutputState }

func (GetAppConnectorControllerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppConnectorControllerResult)(nil)).Elem()
}

func (o GetAppConnectorControllerResultOutput) ToGetAppConnectorControllerResultOutput() GetAppConnectorControllerResultOutput {
	return o
}

func (o GetAppConnectorControllerResultOutput) ToGetAppConnectorControllerResultOutputWithContext(ctx context.Context) GetAppConnectorControllerResultOutput {
	return o
}

func (o GetAppConnectorControllerResultOutput) AppConnectorGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.AppConnectorGroupId }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) AppConnectorGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.AppConnectorGroupName }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) ApplicationStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.ApplicationStartTime }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) ControlChannelStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.ControlChannelStatus }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) CtrlBrokerName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.CtrlBrokerName }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) CurrentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.CurrentVersion }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o GetAppConnectorControllerResultOutput) EnrollmentCert() pulumi.MapOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) map[string]interface{} { return v.EnrollmentCert }).(pulumi.MapOutput)
}

func (o GetAppConnectorControllerResultOutput) ExpectedUpgradeTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.ExpectedUpgradeTime }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) ExpectedVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.ExpectedVersion }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) IpAcl() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.IpAcl }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) IssuedCertId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.IssuedCertId }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) LastBrokerConnectTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.LastBrokerConnectTime }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) LastBrokerConnectTimeDuration() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.LastBrokerConnectTimeDuration }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) LastBrokerDisconnectTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.LastBrokerDisconnectTime }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) LastBrokerDisconnectTimeDuration() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.LastBrokerDisconnectTimeDuration }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) LastUpgradeTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.LastUpgradeTime }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) Latitude() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.Latitude }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) Longitude() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.Longitude }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) MicrotenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) *string { return v.MicrotenantId }).(pulumi.StringPtrOutput)
}

func (o GetAppConnectorControllerResultOutput) MicrotenantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) *string { return v.MicrotenantName }).(pulumi.StringPtrOutput)
}

func (o GetAppConnectorControllerResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) Modifiedby() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.Modifiedby }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetAppConnectorControllerResultOutput) Platform() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.Platform }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) PreviousVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.PreviousVersion }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.PrivateIp }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) ProvisioningKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.ProvisioningKeyId }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) ProvisioningKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.ProvisioningKeyName }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.PublicIp }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) SargeVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.SargeVersion }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) UpgradeAttempt() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.UpgradeAttempt }).(pulumi.StringOutput)
}

func (o GetAppConnectorControllerResultOutput) UpgradeStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorControllerResult) string { return v.UpgradeStatus }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppConnectorControllerResultOutput{})
}
