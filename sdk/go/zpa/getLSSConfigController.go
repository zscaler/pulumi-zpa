// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// Use the **zpa_lss_config_controller** data source to get information about a Log Streaming (LSS) configuration resource created in the Zscaler Private Access.
func LookupLSSConfigController(ctx *pulumi.Context, args *LookupLSSConfigControllerArgs, opts ...pulumi.InvokeOption) (*LookupLSSConfigControllerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLSSConfigControllerResult
	err := ctx.Invoke("zpa:index/getLSSConfigController:getLSSConfigController", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLSSConfigController.
type LookupLSSConfigControllerArgs struct {
	// (Computed)
	Configs []GetLSSConfigControllerConfig `pulumi:"configs"`
	// This field defines the name of the log streaming resource.
	Id *string `pulumi:"id"`
}

// A collection of values returned by getLSSConfigController.
type LookupLSSConfigControllerResult struct {
	// (Computed)
	Configs []GetLSSConfigControllerConfig `pulumi:"configs"`
	// (Computed)
	ConnectorGroups []GetLSSConfigControllerConnectorGroup `pulumi:"connectorGroups"`
	// (string)
	Id          *string                            `pulumi:"id"`
	PolicyRules []GetLSSConfigControllerPolicyRule `pulumi:"policyRules"`
}

func LookupLSSConfigControllerOutput(ctx *pulumi.Context, args LookupLSSConfigControllerOutputArgs, opts ...pulumi.InvokeOption) LookupLSSConfigControllerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLSSConfigControllerResult, error) {
			args := v.(LookupLSSConfigControllerArgs)
			r, err := LookupLSSConfigController(ctx, &args, opts...)
			var s LookupLSSConfigControllerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLSSConfigControllerResultOutput)
}

// A collection of arguments for invoking getLSSConfigController.
type LookupLSSConfigControllerOutputArgs struct {
	// (Computed)
	Configs GetLSSConfigControllerConfigArrayInput `pulumi:"configs"`
	// This field defines the name of the log streaming resource.
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (LookupLSSConfigControllerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLSSConfigControllerArgs)(nil)).Elem()
}

// A collection of values returned by getLSSConfigController.
type LookupLSSConfigControllerResultOutput struct{ *pulumi.OutputState }

func (LookupLSSConfigControllerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLSSConfigControllerResult)(nil)).Elem()
}

func (o LookupLSSConfigControllerResultOutput) ToLookupLSSConfigControllerResultOutput() LookupLSSConfigControllerResultOutput {
	return o
}

func (o LookupLSSConfigControllerResultOutput) ToLookupLSSConfigControllerResultOutputWithContext(ctx context.Context) LookupLSSConfigControllerResultOutput {
	return o
}

// (Computed)
func (o LookupLSSConfigControllerResultOutput) Configs() GetLSSConfigControllerConfigArrayOutput {
	return o.ApplyT(func(v LookupLSSConfigControllerResult) []GetLSSConfigControllerConfig { return v.Configs }).(GetLSSConfigControllerConfigArrayOutput)
}

// (Computed)
func (o LookupLSSConfigControllerResultOutput) ConnectorGroups() GetLSSConfigControllerConnectorGroupArrayOutput {
	return o.ApplyT(func(v LookupLSSConfigControllerResult) []GetLSSConfigControllerConnectorGroup {
		return v.ConnectorGroups
	}).(GetLSSConfigControllerConnectorGroupArrayOutput)
}

// (string)
func (o LookupLSSConfigControllerResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLSSConfigControllerResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupLSSConfigControllerResultOutput) PolicyRules() GetLSSConfigControllerPolicyRuleArrayOutput {
	return o.ApplyT(func(v LookupLSSConfigControllerResult) []GetLSSConfigControllerPolicyRule { return v.PolicyRules }).(GetLSSConfigControllerPolicyRuleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLSSConfigControllerResultOutput{})
}