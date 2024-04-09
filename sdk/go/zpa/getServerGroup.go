// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

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
//			_, err := zpa.LookupServerGroup(ctx, &zpa.LookupServerGroupArgs{
//				Name: pulumi.StringRef("server_group_name"),
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
func LookupServerGroup(ctx *pulumi.Context, args *LookupServerGroupArgs, opts ...pulumi.InvokeOption) (*LookupServerGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServerGroupResult
	err := ctx.Invoke("zpa:index/getServerGroup:getServerGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServerGroup.
type LookupServerGroupArgs struct {
	Id              *string `pulumi:"id"`
	MicrotenantId   *string `pulumi:"microtenantId"`
	MicrotenantName *string `pulumi:"microtenantName"`
	Name            *string `pulumi:"name"`
}

// A collection of values returned by getServerGroup.
type LookupServerGroupResult struct {
	AppConnectorGroups []GetServerGroupAppConnectorGroup `pulumi:"appConnectorGroups"`
	Applications       []GetServerGroupApplication       `pulumi:"applications"`
	ConfigSpace        string                            `pulumi:"configSpace"`
	CreationTime       string                            `pulumi:"creationTime"`
	Description        string                            `pulumi:"description"`
	DynamicDiscovery   bool                              `pulumi:"dynamicDiscovery"`
	Enabled            bool                              `pulumi:"enabled"`
	Id                 *string                           `pulumi:"id"`
	IpAnchored         bool                              `pulumi:"ipAnchored"`
	MicrotenantId      *string                           `pulumi:"microtenantId"`
	MicrotenantName    *string                           `pulumi:"microtenantName"`
	ModifiedTime       string                            `pulumi:"modifiedTime"`
	Modifiedby         string                            `pulumi:"modifiedby"`
	Name               *string                           `pulumi:"name"`
	Servers            []GetServerGroupServer            `pulumi:"servers"`
}

func LookupServerGroupOutput(ctx *pulumi.Context, args LookupServerGroupOutputArgs, opts ...pulumi.InvokeOption) LookupServerGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerGroupResult, error) {
			args := v.(LookupServerGroupArgs)
			r, err := LookupServerGroup(ctx, &args, opts...)
			var s LookupServerGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerGroupResultOutput)
}

// A collection of arguments for invoking getServerGroup.
type LookupServerGroupOutputArgs struct {
	Id              pulumi.StringPtrInput `pulumi:"id"`
	MicrotenantId   pulumi.StringPtrInput `pulumi:"microtenantId"`
	MicrotenantName pulumi.StringPtrInput `pulumi:"microtenantName"`
	Name            pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupServerGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerGroupArgs)(nil)).Elem()
}

// A collection of values returned by getServerGroup.
type LookupServerGroupResultOutput struct{ *pulumi.OutputState }

func (LookupServerGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerGroupResult)(nil)).Elem()
}

func (o LookupServerGroupResultOutput) ToLookupServerGroupResultOutput() LookupServerGroupResultOutput {
	return o
}

func (o LookupServerGroupResultOutput) ToLookupServerGroupResultOutputWithContext(ctx context.Context) LookupServerGroupResultOutput {
	return o
}

func (o LookupServerGroupResultOutput) AppConnectorGroups() GetServerGroupAppConnectorGroupArrayOutput {
	return o.ApplyT(func(v LookupServerGroupResult) []GetServerGroupAppConnectorGroup { return v.AppConnectorGroups }).(GetServerGroupAppConnectorGroupArrayOutput)
}

func (o LookupServerGroupResultOutput) Applications() GetServerGroupApplicationArrayOutput {
	return o.ApplyT(func(v LookupServerGroupResult) []GetServerGroupApplication { return v.Applications }).(GetServerGroupApplicationArrayOutput)
}

func (o LookupServerGroupResultOutput) ConfigSpace() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerGroupResult) string { return v.ConfigSpace }).(pulumi.StringOutput)
}

func (o LookupServerGroupResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerGroupResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupServerGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupServerGroupResultOutput) DynamicDiscovery() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServerGroupResult) bool { return v.DynamicDiscovery }).(pulumi.BoolOutput)
}

func (o LookupServerGroupResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServerGroupResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupServerGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupServerGroupResultOutput) IpAnchored() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServerGroupResult) bool { return v.IpAnchored }).(pulumi.BoolOutput)
}

func (o LookupServerGroupResultOutput) MicrotenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerGroupResult) *string { return v.MicrotenantId }).(pulumi.StringPtrOutput)
}

func (o LookupServerGroupResultOutput) MicrotenantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerGroupResult) *string { return v.MicrotenantName }).(pulumi.StringPtrOutput)
}

func (o LookupServerGroupResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerGroupResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o LookupServerGroupResultOutput) Modifiedby() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerGroupResult) string { return v.Modifiedby }).(pulumi.StringOutput)
}

func (o LookupServerGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupServerGroupResultOutput) Servers() GetServerGroupServerArrayOutput {
	return o.ApplyT(func(v LookupServerGroupResult) []GetServerGroupServer { return v.Servers }).(GetServerGroupServerArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerGroupResultOutput{})
}
