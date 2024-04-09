// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// Use the **zpa_pra_approval_controller** data source to get information about a privileged remote access approval created in the Zscaler Private Access cloud.
func LookupPRAApproval(ctx *pulumi.Context, args *LookupPRAApprovalArgs, opts ...pulumi.InvokeOption) (*LookupPRAApprovalResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPRAApprovalResult
	err := ctx.Invoke("zpa:index/getPRAApproval:getPRAApproval", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPRAApproval.
type LookupPRAApprovalArgs struct {
	EmailIds []string `pulumi:"emailIds"`
	Id       *string  `pulumi:"id"`
}

// A collection of values returned by getPRAApproval.
type LookupPRAApprovalResult struct {
	Applications []GetPRAApprovalApplication `pulumi:"applications"`
	CreationTime string                      `pulumi:"creationTime"`
	EmailIds     []string                    `pulumi:"emailIds"`
	EndTime      string                      `pulumi:"endTime"`
	Id           *string                     `pulumi:"id"`
	// (string)  The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass microtenantId as 0 when making requests to retrieve data from the Default Microtenant. Pass microtenantId as null to retrieve data from all customers associated with the tenant.
	MicrotenantId string `pulumi:"microtenantId"`
	ModifiedBy    string `pulumi:"modifiedBy"`
	ModifiedTime  string `pulumi:"modifiedTime"`
	StartTime     string `pulumi:"startTime"`
	// (string) The status of the privileged approval. The supported values are:
	Status       string                      `pulumi:"status"`
	WorkingHours []GetPRAApprovalWorkingHour `pulumi:"workingHours"`
}

func LookupPRAApprovalOutput(ctx *pulumi.Context, args LookupPRAApprovalOutputArgs, opts ...pulumi.InvokeOption) LookupPRAApprovalResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPRAApprovalResult, error) {
			args := v.(LookupPRAApprovalArgs)
			r, err := LookupPRAApproval(ctx, &args, opts...)
			var s LookupPRAApprovalResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPRAApprovalResultOutput)
}

// A collection of arguments for invoking getPRAApproval.
type LookupPRAApprovalOutputArgs struct {
	EmailIds pulumi.StringArrayInput `pulumi:"emailIds"`
	Id       pulumi.StringPtrInput   `pulumi:"id"`
}

func (LookupPRAApprovalOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPRAApprovalArgs)(nil)).Elem()
}

// A collection of values returned by getPRAApproval.
type LookupPRAApprovalResultOutput struct{ *pulumi.OutputState }

func (LookupPRAApprovalResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPRAApprovalResult)(nil)).Elem()
}

func (o LookupPRAApprovalResultOutput) ToLookupPRAApprovalResultOutput() LookupPRAApprovalResultOutput {
	return o
}

func (o LookupPRAApprovalResultOutput) ToLookupPRAApprovalResultOutputWithContext(ctx context.Context) LookupPRAApprovalResultOutput {
	return o
}

func (o LookupPRAApprovalResultOutput) Applications() GetPRAApprovalApplicationArrayOutput {
	return o.ApplyT(func(v LookupPRAApprovalResult) []GetPRAApprovalApplication { return v.Applications }).(GetPRAApprovalApplicationArrayOutput)
}

func (o LookupPRAApprovalResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPRAApprovalResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupPRAApprovalResultOutput) EmailIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPRAApprovalResult) []string { return v.EmailIds }).(pulumi.StringArrayOutput)
}

func (o LookupPRAApprovalResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPRAApprovalResult) string { return v.EndTime }).(pulumi.StringOutput)
}

func (o LookupPRAApprovalResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPRAApprovalResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// (string)  The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass microtenantId as 0 when making requests to retrieve data from the Default Microtenant. Pass microtenantId as null to retrieve data from all customers associated with the tenant.
func (o LookupPRAApprovalResultOutput) MicrotenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPRAApprovalResult) string { return v.MicrotenantId }).(pulumi.StringOutput)
}

func (o LookupPRAApprovalResultOutput) ModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPRAApprovalResult) string { return v.ModifiedBy }).(pulumi.StringOutput)
}

func (o LookupPRAApprovalResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPRAApprovalResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o LookupPRAApprovalResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPRAApprovalResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// (string) The status of the privileged approval. The supported values are:
func (o LookupPRAApprovalResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPRAApprovalResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupPRAApprovalResultOutput) WorkingHours() GetPRAApprovalWorkingHourArrayOutput {
	return o.ApplyT(func(v LookupPRAApprovalResult) []GetPRAApprovalWorkingHour { return v.WorkingHours }).(GetPRAApprovalWorkingHourArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPRAApprovalResultOutput{})
}
