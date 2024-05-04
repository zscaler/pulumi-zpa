// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// * [Official documentation](https://help.zscaler.com/zpa/about-log-streaming-service)
// * [API documentation](https://help.zscaler.com/zpa/configuring-log-streaming-service-configurations-using-api)
//
// Use the **zpa_lss_config_client_types** data source to get information about all LSS client types in the Zscaler Private Access cloud. This data source is required when the defining a policy rule resource for an object type as `CLIENT_TYPE` parameter in the LSS Config Controller resource is set. To learn more see the To learn more see the [Getting Details of All LSS Status Codes](https://help.zscaler.com/zpa/log-streaming-service-configuration-use-cases#GettingLSSClientTypes)
//
// > **NOTE** By Default the ZPA provider will return all client types
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
//			_, err := zpa.GetLSSClientTypes(ctx, nil, nil)
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
// ### Read-Only
//
// The following arguments are supported:
//
// * `"zpnClientTypeEdgeConnector" = "Cloud Connector"`
// * `"zpnClientTypeExporter" = "Web Browser`
// * `"zpnClientTypeIpAnchoring" = "ZIA Service Edge"`
// * `"zpnClientTypeMachineTunnel" = "Machine Tunnel"`
// * `"zpnClientTypeSlogger" = "ZPA LSS"`
// * `"zpnClientTypeZapp" = "Client Connector"`
//
// To learn more see the [Getting Details of All LSS Status Codes](https://help.zscaler.com/zpa/log-streaming-service-configuration-use-cases#GettingLSSClientTypes)
func GetLSSClientTypes(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetLSSClientTypesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLSSClientTypesResult
	err := ctx.Invoke("zpa:index/getLSSClientTypes:getLSSClientTypes", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getLSSClientTypes.
type GetLSSClientTypesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                         string `pulumi:"id"`
	ZpnClientTypeEdgeConnector string `pulumi:"zpnClientTypeEdgeConnector"`
	ZpnClientTypeExporter      string `pulumi:"zpnClientTypeExporter"`
	ZpnClientTypeIpAnchoring   string `pulumi:"zpnClientTypeIpAnchoring"`
	ZpnClientTypeMachineTunnel string `pulumi:"zpnClientTypeMachineTunnel"`
	ZpnClientTypeSlogger       string `pulumi:"zpnClientTypeSlogger"`
	ZpnClientTypeZapp          string `pulumi:"zpnClientTypeZapp"`
}

func GetLSSClientTypesOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetLSSClientTypesResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetLSSClientTypesResult, error) {
		r, err := GetLSSClientTypes(ctx, opts...)
		var s GetLSSClientTypesResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(GetLSSClientTypesResultOutput)
}

// A collection of values returned by getLSSClientTypes.
type GetLSSClientTypesResultOutput struct{ *pulumi.OutputState }

func (GetLSSClientTypesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLSSClientTypesResult)(nil)).Elem()
}

func (o GetLSSClientTypesResultOutput) ToGetLSSClientTypesResultOutput() GetLSSClientTypesResultOutput {
	return o
}

func (o GetLSSClientTypesResultOutput) ToGetLSSClientTypesResultOutputWithContext(ctx context.Context) GetLSSClientTypesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetLSSClientTypesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLSSClientTypesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLSSClientTypesResultOutput) ZpnClientTypeEdgeConnector() pulumi.StringOutput {
	return o.ApplyT(func(v GetLSSClientTypesResult) string { return v.ZpnClientTypeEdgeConnector }).(pulumi.StringOutput)
}

func (o GetLSSClientTypesResultOutput) ZpnClientTypeExporter() pulumi.StringOutput {
	return o.ApplyT(func(v GetLSSClientTypesResult) string { return v.ZpnClientTypeExporter }).(pulumi.StringOutput)
}

func (o GetLSSClientTypesResultOutput) ZpnClientTypeIpAnchoring() pulumi.StringOutput {
	return o.ApplyT(func(v GetLSSClientTypesResult) string { return v.ZpnClientTypeIpAnchoring }).(pulumi.StringOutput)
}

func (o GetLSSClientTypesResultOutput) ZpnClientTypeMachineTunnel() pulumi.StringOutput {
	return o.ApplyT(func(v GetLSSClientTypesResult) string { return v.ZpnClientTypeMachineTunnel }).(pulumi.StringOutput)
}

func (o GetLSSClientTypesResultOutput) ZpnClientTypeSlogger() pulumi.StringOutput {
	return o.ApplyT(func(v GetLSSClientTypesResult) string { return v.ZpnClientTypeSlogger }).(pulumi.StringOutput)
}

func (o GetLSSClientTypesResultOutput) ZpnClientTypeZapp() pulumi.StringOutput {
	return o.ApplyT(func(v GetLSSClientTypesResult) string { return v.ZpnClientTypeZapp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLSSClientTypesResultOutput{})
}
