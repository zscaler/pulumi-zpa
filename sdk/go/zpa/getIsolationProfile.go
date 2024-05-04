// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// * [Official documentation](https://help.zscaler.com/isolation/creating-isolation-profiles-zpa)
// * [API documentation](https://help.zscaler.com/zpa/obtaining-isolation-profile-details-using-api)
//
// Use the **zpa_isolation_profile** data source to get information about an isolation profile in the Zscaler Private Access cloud. This data source is required when configuring an isolation policy rule resource
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
//			_, err := zpa.GetIsolationProfile(ctx, &zpa.GetIsolationProfileArgs{
//				Name: pulumi.StringRef("zpa_isolation_profile"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetIsolationProfile(ctx *pulumi.Context, args *GetIsolationProfileArgs, opts ...pulumi.InvokeOption) (*GetIsolationProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIsolationProfileResult
	err := ctx.Invoke("zpa:index/getIsolationProfile:getIsolationProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIsolationProfile.
type GetIsolationProfileArgs struct {
	// - (Required) This field defines the name of the isolation profile.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getIsolationProfile.
type GetIsolationProfileResult struct {
	CreationTime       string `pulumi:"creationTime"`
	Description        string `pulumi:"description"`
	Enabled            bool   `pulumi:"enabled"`
	Id                 string `pulumi:"id"`
	IsolationProfileId string `pulumi:"isolationProfileId"`
	IsolationTenantId  string `pulumi:"isolationTenantId"`
	IsolationUrl       string `pulumi:"isolationUrl"`
	ModifiedTime       string `pulumi:"modifiedTime"`
	Modifiedby         string `pulumi:"modifiedby"`
	// - (Required) This field defines the name of the isolation profile.
	Name *string `pulumi:"name"`
}

func GetIsolationProfileOutput(ctx *pulumi.Context, args GetIsolationProfileOutputArgs, opts ...pulumi.InvokeOption) GetIsolationProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIsolationProfileResult, error) {
			args := v.(GetIsolationProfileArgs)
			r, err := GetIsolationProfile(ctx, &args, opts...)
			var s GetIsolationProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIsolationProfileResultOutput)
}

// A collection of arguments for invoking getIsolationProfile.
type GetIsolationProfileOutputArgs struct {
	// - (Required) This field defines the name of the isolation profile.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetIsolationProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIsolationProfileArgs)(nil)).Elem()
}

// A collection of values returned by getIsolationProfile.
type GetIsolationProfileResultOutput struct{ *pulumi.OutputState }

func (GetIsolationProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIsolationProfileResult)(nil)).Elem()
}

func (o GetIsolationProfileResultOutput) ToGetIsolationProfileResultOutput() GetIsolationProfileResultOutput {
	return o
}

func (o GetIsolationProfileResultOutput) ToGetIsolationProfileResultOutputWithContext(ctx context.Context) GetIsolationProfileResultOutput {
	return o
}

func (o GetIsolationProfileResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetIsolationProfileResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o GetIsolationProfileResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetIsolationProfileResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetIsolationProfileResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIsolationProfileResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o GetIsolationProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIsolationProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIsolationProfileResultOutput) IsolationProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIsolationProfileResult) string { return v.IsolationProfileId }).(pulumi.StringOutput)
}

func (o GetIsolationProfileResultOutput) IsolationTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIsolationProfileResult) string { return v.IsolationTenantId }).(pulumi.StringOutput)
}

func (o GetIsolationProfileResultOutput) IsolationUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetIsolationProfileResult) string { return v.IsolationUrl }).(pulumi.StringOutput)
}

func (o GetIsolationProfileResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetIsolationProfileResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o GetIsolationProfileResultOutput) Modifiedby() pulumi.StringOutput {
	return o.ApplyT(func(v GetIsolationProfileResult) string { return v.Modifiedby }).(pulumi.StringOutput)
}

// - (Required) This field defines the name of the isolation profile.
func (o GetIsolationProfileResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIsolationProfileResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIsolationProfileResultOutput{})
}
