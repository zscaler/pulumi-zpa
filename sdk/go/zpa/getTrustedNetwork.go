// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// The **zpa_trusted_network** data source to get information about a trusted network created in the Zscaler Private Access Mobile Portal. This data source can then be referenced within the following resources:
//
// 1. Access Policy
// 2. Forwarding Policy
// 3. Inspection Policy
// 4. Isolation Policy
// 5. Service Edge Group.
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
//			_, err := zpa.GetTrustedNetwork(ctx, &zpa.GetTrustedNetworkArgs{
//				Name: pulumi.StringRef("trusted_network_name"),
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
// > **NOTE** To query trusted network that are associated with a specific Zscaler cloud, it is required to append the cloud name to the name of the trusted network as the below example:
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
//			example1, err := zpa.GetTrustedNetwork(ctx, &zpa.GetTrustedNetworkArgs{
//				Name: pulumi.StringRef("Corporate-Network (zscalertwo.net)"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("zpaTrustedNetwork", example1.NetworkId)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetTrustedNetwork(ctx *pulumi.Context, args *GetTrustedNetworkArgs, opts ...pulumi.InvokeOption) (*GetTrustedNetworkResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTrustedNetworkResult
	err := ctx.Invoke("zpa:index/getTrustedNetwork:getTrustedNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTrustedNetwork.
type GetTrustedNetworkArgs struct {
	// The ID of the posture profile to be exported.
	Id *string `pulumi:"id"`
	// The name of the posture profile to be exported.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getTrustedNetwork.
type GetTrustedNetworkResult struct {
	// (string)
	CreationTime string `pulumi:"creationTime"`
	// (string)
	Domain string  `pulumi:"domain"`
	Id     *string `pulumi:"id"`
	// (string)
	ModifiedTime string  `pulumi:"modifiedTime"`
	Modifiedby   string  `pulumi:"modifiedby"`
	Name         *string `pulumi:"name"`
	// (string)
	NetworkId string `pulumi:"networkId"`
	// (string)
	ZscalerCloud string `pulumi:"zscalerCloud"`
}

func GetTrustedNetworkOutput(ctx *pulumi.Context, args GetTrustedNetworkOutputArgs, opts ...pulumi.InvokeOption) GetTrustedNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTrustedNetworkResult, error) {
			args := v.(GetTrustedNetworkArgs)
			r, err := GetTrustedNetwork(ctx, &args, opts...)
			var s GetTrustedNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTrustedNetworkResultOutput)
}

// A collection of arguments for invoking getTrustedNetwork.
type GetTrustedNetworkOutputArgs struct {
	// The ID of the posture profile to be exported.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the posture profile to be exported.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetTrustedNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTrustedNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getTrustedNetwork.
type GetTrustedNetworkResultOutput struct{ *pulumi.OutputState }

func (GetTrustedNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTrustedNetworkResult)(nil)).Elem()
}

func (o GetTrustedNetworkResultOutput) ToGetTrustedNetworkResultOutput() GetTrustedNetworkResultOutput {
	return o
}

func (o GetTrustedNetworkResultOutput) ToGetTrustedNetworkResultOutputWithContext(ctx context.Context) GetTrustedNetworkResultOutput {
	return o
}

// (string)
func (o GetTrustedNetworkResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrustedNetworkResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// (string)
func (o GetTrustedNetworkResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrustedNetworkResult) string { return v.Domain }).(pulumi.StringOutput)
}

func (o GetTrustedNetworkResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTrustedNetworkResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// (string)
func (o GetTrustedNetworkResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrustedNetworkResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o GetTrustedNetworkResultOutput) Modifiedby() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrustedNetworkResult) string { return v.Modifiedby }).(pulumi.StringOutput)
}

func (o GetTrustedNetworkResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTrustedNetworkResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// (string)
func (o GetTrustedNetworkResultOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrustedNetworkResult) string { return v.NetworkId }).(pulumi.StringOutput)
}

// (string)
func (o GetTrustedNetworkResultOutput) ZscalerCloud() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrustedNetworkResult) string { return v.ZscalerCloud }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTrustedNetworkResultOutput{})
}
