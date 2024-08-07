// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// * [Official documentation](https://help.zscaler.com/zpa/about-privileged-consoles)
// * [API documentation](https://help.zscaler.com/zpa/configuring-privileged-consoles-using-api)
//
// The **zpa_pra_console_controller** data source gets information about a privileged remote access console created in the Zscaler Private Access cloud.
// This resource can then be referenced in an privileged access policy credential and a privileged access portal resource.
//
// Deprecated: zpa.index/getpraconsolecontroller.getPraConsoleController has been deprecated in favor of zpa.index/getpraconsole.getPRAConsole
func LookupPraConsoleController(ctx *pulumi.Context, args *LookupPraConsoleControllerArgs, opts ...pulumi.InvokeOption) (*LookupPraConsoleControllerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPraConsoleControllerResult
	err := ctx.Invoke("zpa:index/getPraConsoleController:getPraConsoleController", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPraConsoleController.
type LookupPraConsoleControllerArgs struct {
	Id *string `pulumi:"id"`
}

// A collection of values returned by getPraConsoleController.
type LookupPraConsoleControllerResult struct {
	CreationTime    string  `pulumi:"creationTime"`
	Description     string  `pulumi:"description"`
	Enabled         bool    `pulumi:"enabled"`
	IconText        string  `pulumi:"iconText"`
	Id              *string `pulumi:"id"`
	MicrotenantId   string  `pulumi:"microtenantId"`
	MicrotenantName string  `pulumi:"microtenantName"`
	ModifiedBy      string  `pulumi:"modifiedBy"`
	ModifiedTime    string  `pulumi:"modifiedTime"`
	// - (Required) The name of the privileged console.
	Name            string                                  `pulumi:"name"`
	PraApplications []GetPraConsoleControllerPraApplication `pulumi:"praApplications"`
	PraPortals      []GetPraConsoleControllerPraPortal      `pulumi:"praPortals"`
}

func LookupPraConsoleControllerOutput(ctx *pulumi.Context, args LookupPraConsoleControllerOutputArgs, opts ...pulumi.InvokeOption) LookupPraConsoleControllerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPraConsoleControllerResult, error) {
			args := v.(LookupPraConsoleControllerArgs)
			r, err := LookupPraConsoleController(ctx, &args, opts...)
			var s LookupPraConsoleControllerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPraConsoleControllerResultOutput)
}

// A collection of arguments for invoking getPraConsoleController.
type LookupPraConsoleControllerOutputArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (LookupPraConsoleControllerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPraConsoleControllerArgs)(nil)).Elem()
}

// A collection of values returned by getPraConsoleController.
type LookupPraConsoleControllerResultOutput struct{ *pulumi.OutputState }

func (LookupPraConsoleControllerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPraConsoleControllerResult)(nil)).Elem()
}

func (o LookupPraConsoleControllerResultOutput) ToLookupPraConsoleControllerResultOutput() LookupPraConsoleControllerResultOutput {
	return o
}

func (o LookupPraConsoleControllerResultOutput) ToLookupPraConsoleControllerResultOutputWithContext(ctx context.Context) LookupPraConsoleControllerResultOutput {
	return o
}

func (o LookupPraConsoleControllerResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPraConsoleControllerResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupPraConsoleControllerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPraConsoleControllerResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupPraConsoleControllerResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPraConsoleControllerResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupPraConsoleControllerResultOutput) IconText() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPraConsoleControllerResult) string { return v.IconText }).(pulumi.StringOutput)
}

func (o LookupPraConsoleControllerResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPraConsoleControllerResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupPraConsoleControllerResultOutput) MicrotenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPraConsoleControllerResult) string { return v.MicrotenantId }).(pulumi.StringOutput)
}

func (o LookupPraConsoleControllerResultOutput) MicrotenantName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPraConsoleControllerResult) string { return v.MicrotenantName }).(pulumi.StringOutput)
}

func (o LookupPraConsoleControllerResultOutput) ModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPraConsoleControllerResult) string { return v.ModifiedBy }).(pulumi.StringOutput)
}

func (o LookupPraConsoleControllerResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPraConsoleControllerResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

// - (Required) The name of the privileged console.
func (o LookupPraConsoleControllerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPraConsoleControllerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPraConsoleControllerResultOutput) PraApplications() GetPraConsoleControllerPraApplicationArrayOutput {
	return o.ApplyT(func(v LookupPraConsoleControllerResult) []GetPraConsoleControllerPraApplication {
		return v.PraApplications
	}).(GetPraConsoleControllerPraApplicationArrayOutput)
}

func (o LookupPraConsoleControllerResultOutput) PraPortals() GetPraConsoleControllerPraPortalArrayOutput {
	return o.ApplyT(func(v LookupPraConsoleControllerResult) []GetPraConsoleControllerPraPortal { return v.PraPortals }).(GetPraConsoleControllerPraPortalArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPraConsoleControllerResultOutput{})
}
