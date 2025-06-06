// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// * [Official documentation](https://help.zscaler.com/isolation/adding-banner-theme-isolation-end-user-notification-zpa)
//
// Use the **zpa_cloud_browser_isolation_banner** data source to get information about Cloud Browser Isolation banner. This data source information is required as part of the attribute `bannerId` when creating an Cloud Browser Isolation External Profile “CloudBrowserIsolationExternalProfile“
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
//			_, err := zpa.LookupCloudBrowserIsolationBanner(ctx, &zpa.LookupCloudBrowserIsolationBannerArgs{
//				Name: pulumi.StringRef("Default"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupCloudBrowserIsolationBanner(ctx *pulumi.Context, args *LookupCloudBrowserIsolationBannerArgs, opts ...pulumi.InvokeOption) (*LookupCloudBrowserIsolationBannerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCloudBrowserIsolationBannerResult
	err := ctx.Invoke("zpa:index/getCloudBrowserIsolationBanner:getCloudBrowserIsolationBanner", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudBrowserIsolationBanner.
type LookupCloudBrowserIsolationBannerArgs struct {
	Id   *string `pulumi:"id"`
	Name *string `pulumi:"name"`
}

// A collection of values returned by getCloudBrowserIsolationBanner.
type LookupCloudBrowserIsolationBannerResult struct {
	Banner            bool    `pulumi:"banner"`
	Id                *string `pulumi:"id"`
	IsDefault         bool    `pulumi:"isDefault"`
	Logo              string  `pulumi:"logo"`
	Name              *string `pulumi:"name"`
	NotificationText  string  `pulumi:"notificationText"`
	NotificationTitle string  `pulumi:"notificationTitle"`
	PrimaryColor      string  `pulumi:"primaryColor"`
	TextColor         string  `pulumi:"textColor"`
}

func LookupCloudBrowserIsolationBannerOutput(ctx *pulumi.Context, args LookupCloudBrowserIsolationBannerOutputArgs, opts ...pulumi.InvokeOption) LookupCloudBrowserIsolationBannerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCloudBrowserIsolationBannerResultOutput, error) {
			args := v.(LookupCloudBrowserIsolationBannerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("zpa:index/getCloudBrowserIsolationBanner:getCloudBrowserIsolationBanner", args, LookupCloudBrowserIsolationBannerResultOutput{}, options).(LookupCloudBrowserIsolationBannerResultOutput), nil
		}).(LookupCloudBrowserIsolationBannerResultOutput)
}

// A collection of arguments for invoking getCloudBrowserIsolationBanner.
type LookupCloudBrowserIsolationBannerOutputArgs struct {
	Id   pulumi.StringPtrInput `pulumi:"id"`
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupCloudBrowserIsolationBannerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudBrowserIsolationBannerArgs)(nil)).Elem()
}

// A collection of values returned by getCloudBrowserIsolationBanner.
type LookupCloudBrowserIsolationBannerResultOutput struct{ *pulumi.OutputState }

func (LookupCloudBrowserIsolationBannerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudBrowserIsolationBannerResult)(nil)).Elem()
}

func (o LookupCloudBrowserIsolationBannerResultOutput) ToLookupCloudBrowserIsolationBannerResultOutput() LookupCloudBrowserIsolationBannerResultOutput {
	return o
}

func (o LookupCloudBrowserIsolationBannerResultOutput) ToLookupCloudBrowserIsolationBannerResultOutputWithContext(ctx context.Context) LookupCloudBrowserIsolationBannerResultOutput {
	return o
}

func (o LookupCloudBrowserIsolationBannerResultOutput) Banner() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCloudBrowserIsolationBannerResult) bool { return v.Banner }).(pulumi.BoolOutput)
}

func (o LookupCloudBrowserIsolationBannerResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudBrowserIsolationBannerResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupCloudBrowserIsolationBannerResultOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCloudBrowserIsolationBannerResult) bool { return v.IsDefault }).(pulumi.BoolOutput)
}

func (o LookupCloudBrowserIsolationBannerResultOutput) Logo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudBrowserIsolationBannerResult) string { return v.Logo }).(pulumi.StringOutput)
}

func (o LookupCloudBrowserIsolationBannerResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudBrowserIsolationBannerResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupCloudBrowserIsolationBannerResultOutput) NotificationText() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudBrowserIsolationBannerResult) string { return v.NotificationText }).(pulumi.StringOutput)
}

func (o LookupCloudBrowserIsolationBannerResultOutput) NotificationTitle() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudBrowserIsolationBannerResult) string { return v.NotificationTitle }).(pulumi.StringOutput)
}

func (o LookupCloudBrowserIsolationBannerResultOutput) PrimaryColor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudBrowserIsolationBannerResult) string { return v.PrimaryColor }).(pulumi.StringOutput)
}

func (o LookupCloudBrowserIsolationBannerResultOutput) TextColor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudBrowserIsolationBannerResult) string { return v.TextColor }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudBrowserIsolationBannerResultOutput{})
}
