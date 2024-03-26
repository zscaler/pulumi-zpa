// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// Use the **zpa_customer_version_profile** data source to get information about all customer version profiles from the Zscaler Private Access cloud. This data source can be associated with an App Connector Group within the parameter `versionProfileId` or `versionProfileName`
//
// The customer version profile IDs are:
//
// * `Default` = `0`
// * `Previous Default` = `1`
// * `New Release` = `2`
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
//			_, err := zpa.GetCustomerVersionProfile(ctx, &zpa.GetCustomerVersionProfileArgs{
//				Name: "Default",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.GetCustomerVersionProfile(ctx, &zpa.GetCustomerVersionProfileArgs{
//				Name: "Previous Default",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.GetCustomerVersionProfile(ctx, &zpa.GetCustomerVersionProfileArgs{
//				Name: "New Release",
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
func GetCustomerVersionProfile(ctx *pulumi.Context, args *GetCustomerVersionProfileArgs, opts ...pulumi.InvokeOption) (*GetCustomerVersionProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCustomerVersionProfileResult
	err := ctx.Invoke("zpa:index/getCustomerVersionProfile:getCustomerVersionProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCustomerVersionProfile.
type GetCustomerVersionProfileArgs struct {
	// The name of the enrollment certificate to be exported.
	Name string `pulumi:"name"`
}

// A collection of values returned by getCustomerVersionProfile.
type GetCustomerVersionProfileResult struct {
	// (string)
	CreationTime           string                                           `pulumi:"creationTime"`
	CustomScopeCustomerIds []GetCustomerVersionProfileCustomScopeCustomerId `pulumi:"customScopeCustomerIds"`
	CustomerId             string                                           `pulumi:"customerId"`
	// (string)
	Description string `pulumi:"description"`
	Id          string `pulumi:"id"`
	// (string)
	ModifiedBy string `pulumi:"modifiedBy"`
	// (string)
	ModifiedTime    string                             `pulumi:"modifiedTime"`
	Name            string                             `pulumi:"name"`
	UpgradePriority string                             `pulumi:"upgradePriority"`
	Versions        []GetCustomerVersionProfileVersion `pulumi:"versions"`
	VisibilityScope string                             `pulumi:"visibilityScope"`
}

func GetCustomerVersionProfileOutput(ctx *pulumi.Context, args GetCustomerVersionProfileOutputArgs, opts ...pulumi.InvokeOption) GetCustomerVersionProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCustomerVersionProfileResult, error) {
			args := v.(GetCustomerVersionProfileArgs)
			r, err := GetCustomerVersionProfile(ctx, &args, opts...)
			var s GetCustomerVersionProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCustomerVersionProfileResultOutput)
}

// A collection of arguments for invoking getCustomerVersionProfile.
type GetCustomerVersionProfileOutputArgs struct {
	// The name of the enrollment certificate to be exported.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetCustomerVersionProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomerVersionProfileArgs)(nil)).Elem()
}

// A collection of values returned by getCustomerVersionProfile.
type GetCustomerVersionProfileResultOutput struct{ *pulumi.OutputState }

func (GetCustomerVersionProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomerVersionProfileResult)(nil)).Elem()
}

func (o GetCustomerVersionProfileResultOutput) ToGetCustomerVersionProfileResultOutput() GetCustomerVersionProfileResultOutput {
	return o
}

func (o GetCustomerVersionProfileResultOutput) ToGetCustomerVersionProfileResultOutputWithContext(ctx context.Context) GetCustomerVersionProfileResultOutput {
	return o
}

// (string)
func (o GetCustomerVersionProfileResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetCustomerVersionProfileResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o GetCustomerVersionProfileResultOutput) CustomScopeCustomerIds() GetCustomerVersionProfileCustomScopeCustomerIdArrayOutput {
	return o.ApplyT(func(v GetCustomerVersionProfileResult) []GetCustomerVersionProfileCustomScopeCustomerId {
		return v.CustomScopeCustomerIds
	}).(GetCustomerVersionProfileCustomScopeCustomerIdArrayOutput)
}

func (o GetCustomerVersionProfileResultOutput) CustomerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCustomerVersionProfileResult) string { return v.CustomerId }).(pulumi.StringOutput)
}

// (string)
func (o GetCustomerVersionProfileResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetCustomerVersionProfileResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetCustomerVersionProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCustomerVersionProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

// (string)
func (o GetCustomerVersionProfileResultOutput) ModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v GetCustomerVersionProfileResult) string { return v.ModifiedBy }).(pulumi.StringOutput)
}

// (string)
func (o GetCustomerVersionProfileResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetCustomerVersionProfileResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o GetCustomerVersionProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetCustomerVersionProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetCustomerVersionProfileResultOutput) UpgradePriority() pulumi.StringOutput {
	return o.ApplyT(func(v GetCustomerVersionProfileResult) string { return v.UpgradePriority }).(pulumi.StringOutput)
}

func (o GetCustomerVersionProfileResultOutput) Versions() GetCustomerVersionProfileVersionArrayOutput {
	return o.ApplyT(func(v GetCustomerVersionProfileResult) []GetCustomerVersionProfileVersion { return v.Versions }).(GetCustomerVersionProfileVersionArrayOutput)
}

func (o GetCustomerVersionProfileResultOutput) VisibilityScope() pulumi.StringOutput {
	return o.ApplyT(func(v GetCustomerVersionProfileResult) string { return v.VisibilityScope }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCustomerVersionProfileResultOutput{})
}
