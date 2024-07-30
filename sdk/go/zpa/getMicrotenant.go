// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// * [Official documentation](https://help.zscaler.com/zpa/about-microtenants)
// * [API documentation](https://help.zscaler.com/zpa/configuring-microtenants-using-api)
//
// The **zpa_microtenant_controller** data source to get information about a machine group created in the Zscaler Private Access cloud. This data source allows administrators to retrieve a specific microtenant ID, which can be passed to other supported resources via the `microtenantId` attribute.
//
// ⚠️ **WARNING:**: This feature is in limited availability and requires additional license. To learn more, contact Zscaler Support or your local account team.
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
//			_, err := zpa.LookupMicrotenant(ctx, &zpa.LookupMicrotenantArgs{
//				Name: pulumi.StringRef("Microtenant_A"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupMicrotenant(ctx *pulumi.Context, args *LookupMicrotenantArgs, opts ...pulumi.InvokeOption) (*LookupMicrotenantResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMicrotenantResult
	err := ctx.Invoke("zpa:index/getMicrotenant:getMicrotenant", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMicrotenant.
type LookupMicrotenantArgs struct {
	Id *string `pulumi:"id"`
	// - (Required) Name of the microtenant controller.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getMicrotenant.
type LookupMicrotenantResult struct {
	CreationTime            string   `pulumi:"creationTime"`
	CriteriaAttribute       string   `pulumi:"criteriaAttribute"`
	CriteriaAttributeValues []string `pulumi:"criteriaAttributeValues"`
	Description             string   `pulumi:"description"`
	Enabled                 bool     `pulumi:"enabled"`
	Id                      *string  `pulumi:"id"`
	ModifiedBy              string   `pulumi:"modifiedBy"`
	ModifiedTime            string   `pulumi:"modifiedTime"`
	// - (Required) Name of the microtenant controller.
	Name     *string              `pulumi:"name"`
	Operator string               `pulumi:"operator"`
	Priority string               `pulumi:"priority"`
	Roles    []GetMicrotenantRole `pulumi:"roles"`
	Users    []GetMicrotenantUser `pulumi:"users"`
}

func LookupMicrotenantOutput(ctx *pulumi.Context, args LookupMicrotenantOutputArgs, opts ...pulumi.InvokeOption) LookupMicrotenantResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMicrotenantResult, error) {
			args := v.(LookupMicrotenantArgs)
			r, err := LookupMicrotenant(ctx, &args, opts...)
			var s LookupMicrotenantResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMicrotenantResultOutput)
}

// A collection of arguments for invoking getMicrotenant.
type LookupMicrotenantOutputArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
	// - (Required) Name of the microtenant controller.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupMicrotenantOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMicrotenantArgs)(nil)).Elem()
}

// A collection of values returned by getMicrotenant.
type LookupMicrotenantResultOutput struct{ *pulumi.OutputState }

func (LookupMicrotenantResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMicrotenantResult)(nil)).Elem()
}

func (o LookupMicrotenantResultOutput) ToLookupMicrotenantResultOutput() LookupMicrotenantResultOutput {
	return o
}

func (o LookupMicrotenantResultOutput) ToLookupMicrotenantResultOutputWithContext(ctx context.Context) LookupMicrotenantResultOutput {
	return o
}

func (o LookupMicrotenantResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrotenantResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupMicrotenantResultOutput) CriteriaAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrotenantResult) string { return v.CriteriaAttribute }).(pulumi.StringOutput)
}

func (o LookupMicrotenantResultOutput) CriteriaAttributeValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMicrotenantResult) []string { return v.CriteriaAttributeValues }).(pulumi.StringArrayOutput)
}

func (o LookupMicrotenantResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrotenantResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupMicrotenantResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMicrotenantResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupMicrotenantResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMicrotenantResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupMicrotenantResultOutput) ModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrotenantResult) string { return v.ModifiedBy }).(pulumi.StringOutput)
}

func (o LookupMicrotenantResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrotenantResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

// - (Required) Name of the microtenant controller.
func (o LookupMicrotenantResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMicrotenantResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupMicrotenantResultOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrotenantResult) string { return v.Operator }).(pulumi.StringOutput)
}

func (o LookupMicrotenantResultOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrotenantResult) string { return v.Priority }).(pulumi.StringOutput)
}

func (o LookupMicrotenantResultOutput) Roles() GetMicrotenantRoleArrayOutput {
	return o.ApplyT(func(v LookupMicrotenantResult) []GetMicrotenantRole { return v.Roles }).(GetMicrotenantRoleArrayOutput)
}

func (o LookupMicrotenantResultOutput) Users() GetMicrotenantUserArrayOutput {
	return o.ApplyT(func(v LookupMicrotenantResult) []GetMicrotenantUser { return v.Users }).(GetMicrotenantUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMicrotenantResultOutput{})
}
