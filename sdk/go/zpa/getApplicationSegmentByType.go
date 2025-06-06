// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// * [Official documentation](https://help.zscaler.com/zpa/about-applications)
// * [API documentation](https://help.zscaler.com/zpa/configuring-application-segments-using-api)
//
// Use the **zpa_application_segment_by_type** data source to get all configured Application Segments by Access Type (e.g., “BROWSER_ACCESS“, “INSPECT“, or “SECURE_REMOTE_ACCESS“) for the specified customer.
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
//			_, err := zpa.GetApplicationSegmentByType(ctx, &zpa.GetApplicationSegmentByTypeArgs{
//				ApplicationType: "SECURE_REMOTE_ACCESS",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetApplicationSegmentByType(ctx *pulumi.Context, args *GetApplicationSegmentByTypeArgs, opts ...pulumi.InvokeOption) (*GetApplicationSegmentByTypeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetApplicationSegmentByTypeResult
	err := ctx.Invoke("zpa:index/getApplicationSegmentByType:getApplicationSegmentByType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplicationSegmentByType.
type GetApplicationSegmentByTypeArgs struct {
	ApplicationType string  `pulumi:"applicationType"`
	MicrotenantId   *string `pulumi:"microtenantId"`
	Name            *string `pulumi:"name"`
}

// A collection of values returned by getApplicationSegmentByType.
type GetApplicationSegmentByTypeResult struct {
	AppId               string  `pulumi:"appId"`
	ApplicationPort     string  `pulumi:"applicationPort"`
	ApplicationProtocol string  `pulumi:"applicationProtocol"`
	ApplicationType     string  `pulumi:"applicationType"`
	CertificateId       string  `pulumi:"certificateId"`
	CertificateName     string  `pulumi:"certificateName"`
	Domain              string  `pulumi:"domain"`
	Enabled             bool    `pulumi:"enabled"`
	Id                  string  `pulumi:"id"`
	MicrotenantId       *string `pulumi:"microtenantId"`
	MicrotenantName     string  `pulumi:"microtenantName"`
	Name                *string `pulumi:"name"`
}

func GetApplicationSegmentByTypeOutput(ctx *pulumi.Context, args GetApplicationSegmentByTypeOutputArgs, opts ...pulumi.InvokeOption) GetApplicationSegmentByTypeResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetApplicationSegmentByTypeResultOutput, error) {
			args := v.(GetApplicationSegmentByTypeArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("zpa:index/getApplicationSegmentByType:getApplicationSegmentByType", args, GetApplicationSegmentByTypeResultOutput{}, options).(GetApplicationSegmentByTypeResultOutput), nil
		}).(GetApplicationSegmentByTypeResultOutput)
}

// A collection of arguments for invoking getApplicationSegmentByType.
type GetApplicationSegmentByTypeOutputArgs struct {
	ApplicationType pulumi.StringInput    `pulumi:"applicationType"`
	MicrotenantId   pulumi.StringPtrInput `pulumi:"microtenantId"`
	Name            pulumi.StringPtrInput `pulumi:"name"`
}

func (GetApplicationSegmentByTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationSegmentByTypeArgs)(nil)).Elem()
}

// A collection of values returned by getApplicationSegmentByType.
type GetApplicationSegmentByTypeResultOutput struct{ *pulumi.OutputState }

func (GetApplicationSegmentByTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationSegmentByTypeResult)(nil)).Elem()
}

func (o GetApplicationSegmentByTypeResultOutput) ToGetApplicationSegmentByTypeResultOutput() GetApplicationSegmentByTypeResultOutput {
	return o
}

func (o GetApplicationSegmentByTypeResultOutput) ToGetApplicationSegmentByTypeResultOutputWithContext(ctx context.Context) GetApplicationSegmentByTypeResultOutput {
	return o
}

func (o GetApplicationSegmentByTypeResultOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationSegmentByTypeResult) string { return v.AppId }).(pulumi.StringOutput)
}

func (o GetApplicationSegmentByTypeResultOutput) ApplicationPort() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationSegmentByTypeResult) string { return v.ApplicationPort }).(pulumi.StringOutput)
}

func (o GetApplicationSegmentByTypeResultOutput) ApplicationProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationSegmentByTypeResult) string { return v.ApplicationProtocol }).(pulumi.StringOutput)
}

func (o GetApplicationSegmentByTypeResultOutput) ApplicationType() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationSegmentByTypeResult) string { return v.ApplicationType }).(pulumi.StringOutput)
}

func (o GetApplicationSegmentByTypeResultOutput) CertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationSegmentByTypeResult) string { return v.CertificateId }).(pulumi.StringOutput)
}

func (o GetApplicationSegmentByTypeResultOutput) CertificateName() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationSegmentByTypeResult) string { return v.CertificateName }).(pulumi.StringOutput)
}

func (o GetApplicationSegmentByTypeResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationSegmentByTypeResult) string { return v.Domain }).(pulumi.StringOutput)
}

func (o GetApplicationSegmentByTypeResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetApplicationSegmentByTypeResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o GetApplicationSegmentByTypeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationSegmentByTypeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetApplicationSegmentByTypeResultOutput) MicrotenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationSegmentByTypeResult) *string { return v.MicrotenantId }).(pulumi.StringPtrOutput)
}

func (o GetApplicationSegmentByTypeResultOutput) MicrotenantName() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationSegmentByTypeResult) string { return v.MicrotenantName }).(pulumi.StringOutput)
}

func (o GetApplicationSegmentByTypeResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationSegmentByTypeResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApplicationSegmentByTypeResultOutput{})
}
