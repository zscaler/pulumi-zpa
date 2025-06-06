// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// * [Official documentation](https://help.zscaler.com/zpa/deleting-disconnected-app-connectors)
// * [API documentation](https://help.zscaler.com/zpa/configuring-auto-delete-disconnected-app-connectors-using-api)
//
// Use the **zpa_app_connector_assistant_schedule** data source to get information about Auto Delete frequency of the App Connector for the specified customer in the Zscaler Private Access cloud.
//
// > **NOTE** - The `customerId` attribute is optional and not required during the configuration.
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
//			_, err := zpa.LookupAppConnectorAssistantSchedule(ctx, &zpa.LookupAppConnectorAssistantScheduleArgs{
//				CustomerId: pulumi.StringRef("1234567891012"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupAppConnectorAssistantSchedule(ctx *pulumi.Context, args *LookupAppConnectorAssistantScheduleArgs, opts ...pulumi.InvokeOption) (*LookupAppConnectorAssistantScheduleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAppConnectorAssistantScheduleResult
	err := ctx.Invoke("zpa:index/getAppConnectorAssistantSchedule:getAppConnectorAssistantSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppConnectorAssistantSchedule.
type LookupAppConnectorAssistantScheduleArgs struct {
	CustomerId *string `pulumi:"customerId"`
	Id         *string `pulumi:"id"`
}

// A collection of values returned by getAppConnectorAssistantSchedule.
type LookupAppConnectorAssistantScheduleResult struct {
	CustomerId        *string `pulumi:"customerId"`
	DeleteDisabled    bool    `pulumi:"deleteDisabled"`
	Enabled           bool    `pulumi:"enabled"`
	Frequency         string  `pulumi:"frequency"`
	FrequencyInterval string  `pulumi:"frequencyInterval"`
	Id                *string `pulumi:"id"`
}

func LookupAppConnectorAssistantScheduleOutput(ctx *pulumi.Context, args LookupAppConnectorAssistantScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupAppConnectorAssistantScheduleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAppConnectorAssistantScheduleResultOutput, error) {
			args := v.(LookupAppConnectorAssistantScheduleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("zpa:index/getAppConnectorAssistantSchedule:getAppConnectorAssistantSchedule", args, LookupAppConnectorAssistantScheduleResultOutput{}, options).(LookupAppConnectorAssistantScheduleResultOutput), nil
		}).(LookupAppConnectorAssistantScheduleResultOutput)
}

// A collection of arguments for invoking getAppConnectorAssistantSchedule.
type LookupAppConnectorAssistantScheduleOutputArgs struct {
	CustomerId pulumi.StringPtrInput `pulumi:"customerId"`
	Id         pulumi.StringPtrInput `pulumi:"id"`
}

func (LookupAppConnectorAssistantScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppConnectorAssistantScheduleArgs)(nil)).Elem()
}

// A collection of values returned by getAppConnectorAssistantSchedule.
type LookupAppConnectorAssistantScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupAppConnectorAssistantScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppConnectorAssistantScheduleResult)(nil)).Elem()
}

func (o LookupAppConnectorAssistantScheduleResultOutput) ToLookupAppConnectorAssistantScheduleResultOutput() LookupAppConnectorAssistantScheduleResultOutput {
	return o
}

func (o LookupAppConnectorAssistantScheduleResultOutput) ToLookupAppConnectorAssistantScheduleResultOutputWithContext(ctx context.Context) LookupAppConnectorAssistantScheduleResultOutput {
	return o
}

func (o LookupAppConnectorAssistantScheduleResultOutput) CustomerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppConnectorAssistantScheduleResult) *string { return v.CustomerId }).(pulumi.StringPtrOutput)
}

func (o LookupAppConnectorAssistantScheduleResultOutput) DeleteDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAppConnectorAssistantScheduleResult) bool { return v.DeleteDisabled }).(pulumi.BoolOutput)
}

func (o LookupAppConnectorAssistantScheduleResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAppConnectorAssistantScheduleResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupAppConnectorAssistantScheduleResultOutput) Frequency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppConnectorAssistantScheduleResult) string { return v.Frequency }).(pulumi.StringOutput)
}

func (o LookupAppConnectorAssistantScheduleResultOutput) FrequencyInterval() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppConnectorAssistantScheduleResult) string { return v.FrequencyInterval }).(pulumi.StringOutput)
}

func (o LookupAppConnectorAssistantScheduleResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppConnectorAssistantScheduleResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppConnectorAssistantScheduleResultOutput{})
}
