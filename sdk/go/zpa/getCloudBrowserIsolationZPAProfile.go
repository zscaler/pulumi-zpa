// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// Use the **zpa_cloud_browser_isolation_zpa_profile** data source to get information about an isolation profile in the Zscaler Private Access cloud. This data source is required when configuring an isolation policy rule resource
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
//			_, err := zpa.GetCloudBrowserIsolationZPAProfile(ctx, &zpa.GetCloudBrowserIsolationZPAProfileArgs{
//				Name: pulumi.StringRef("ZPA_Profile"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetCloudBrowserIsolationZPAProfile(ctx *pulumi.Context, args *GetCloudBrowserIsolationZPAProfileArgs, opts ...pulumi.InvokeOption) (*GetCloudBrowserIsolationZPAProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCloudBrowserIsolationZPAProfileResult
	err := ctx.Invoke("zpa:index/getCloudBrowserIsolationZPAProfile:getCloudBrowserIsolationZPAProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudBrowserIsolationZPAProfile.
type GetCloudBrowserIsolationZPAProfileArgs struct {
	// - (String) This field defines the name of the isolation profile.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getCloudBrowserIsolationZPAProfile.
type GetCloudBrowserIsolationZPAProfileResult struct {
	CbiProfileId string `pulumi:"cbiProfileId"`
	CbiTenantId  string `pulumi:"cbiTenantId"`
	CbiUrl       string `pulumi:"cbiUrl"`
	CreationTime string `pulumi:"creationTime"`
	Description  string `pulumi:"description"`
	Enabled      bool   `pulumi:"enabled"`
	Id           string `pulumi:"id"`
	ModifiedTime string `pulumi:"modifiedTime"`
	Modifiedby   string `pulumi:"modifiedby"`
	// - (String) This field defines the name of the isolation profile.
	Name *string `pulumi:"name"`
}

func GetCloudBrowserIsolationZPAProfileOutput(ctx *pulumi.Context, args GetCloudBrowserIsolationZPAProfileOutputArgs, opts ...pulumi.InvokeOption) GetCloudBrowserIsolationZPAProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCloudBrowserIsolationZPAProfileResult, error) {
			args := v.(GetCloudBrowserIsolationZPAProfileArgs)
			r, err := GetCloudBrowserIsolationZPAProfile(ctx, &args, opts...)
			var s GetCloudBrowserIsolationZPAProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCloudBrowserIsolationZPAProfileResultOutput)
}

// A collection of arguments for invoking getCloudBrowserIsolationZPAProfile.
type GetCloudBrowserIsolationZPAProfileOutputArgs struct {
	// - (String) This field defines the name of the isolation profile.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetCloudBrowserIsolationZPAProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudBrowserIsolationZPAProfileArgs)(nil)).Elem()
}

// A collection of values returned by getCloudBrowserIsolationZPAProfile.
type GetCloudBrowserIsolationZPAProfileResultOutput struct{ *pulumi.OutputState }

func (GetCloudBrowserIsolationZPAProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudBrowserIsolationZPAProfileResult)(nil)).Elem()
}

func (o GetCloudBrowserIsolationZPAProfileResultOutput) ToGetCloudBrowserIsolationZPAProfileResultOutput() GetCloudBrowserIsolationZPAProfileResultOutput {
	return o
}

func (o GetCloudBrowserIsolationZPAProfileResultOutput) ToGetCloudBrowserIsolationZPAProfileResultOutputWithContext(ctx context.Context) GetCloudBrowserIsolationZPAProfileResultOutput {
	return o
}

func (o GetCloudBrowserIsolationZPAProfileResultOutput) CbiProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudBrowserIsolationZPAProfileResult) string { return v.CbiProfileId }).(pulumi.StringOutput)
}

func (o GetCloudBrowserIsolationZPAProfileResultOutput) CbiTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudBrowserIsolationZPAProfileResult) string { return v.CbiTenantId }).(pulumi.StringOutput)
}

func (o GetCloudBrowserIsolationZPAProfileResultOutput) CbiUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudBrowserIsolationZPAProfileResult) string { return v.CbiUrl }).(pulumi.StringOutput)
}

func (o GetCloudBrowserIsolationZPAProfileResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudBrowserIsolationZPAProfileResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o GetCloudBrowserIsolationZPAProfileResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudBrowserIsolationZPAProfileResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetCloudBrowserIsolationZPAProfileResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetCloudBrowserIsolationZPAProfileResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o GetCloudBrowserIsolationZPAProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudBrowserIsolationZPAProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCloudBrowserIsolationZPAProfileResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudBrowserIsolationZPAProfileResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o GetCloudBrowserIsolationZPAProfileResultOutput) Modifiedby() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudBrowserIsolationZPAProfileResult) string { return v.Modifiedby }).(pulumi.StringOutput)
}

// - (String) This field defines the name of the isolation profile.
func (o GetCloudBrowserIsolationZPAProfileResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCloudBrowserIsolationZPAProfileResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCloudBrowserIsolationZPAProfileResultOutput{})
}
