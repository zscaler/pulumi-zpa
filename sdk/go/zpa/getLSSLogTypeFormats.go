// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// * [Official documentation](https://help.zscaler.com/zpa/about-log-streaming-service)
// * [API documentation](https://help.zscaler.com/zpa/configuring-log-streaming-service-configurations-using-api)
//
// Use the **zpa_lss_config_log_type_formats** data source to get information about all LSS log type formats in the Zscaler Private Access cloud. This data source is required when creating an LSS Config Controller resource.
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
//			_, err := zpa.GetLSSLogTypeFormats(ctx, &zpa.GetLSSLogTypeFormatsArgs{
//				LogType: "zpn_trans_log",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.GetLSSLogTypeFormats(ctx, &zpa.GetLSSLogTypeFormatsArgs{
//				LogType: "zpn_auth_log",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.GetLSSLogTypeFormats(ctx, &zpa.GetLSSLogTypeFormatsArgs{
//				LogType: "zpn_ast_auth_log",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.GetLSSLogTypeFormats(ctx, &zpa.GetLSSLogTypeFormatsArgs{
//				LogType: "zpn_http_trans_log",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.GetLSSLogTypeFormats(ctx, &zpa.GetLSSLogTypeFormatsArgs{
//				LogType: "zpn_audit_log",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.GetLSSLogTypeFormats(ctx, &zpa.GetLSSLogTypeFormatsArgs{
//				LogType: "zpn_sys_auth_log",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.GetLSSLogTypeFormats(ctx, &zpa.GetLSSLogTypeFormatsArgs{
//				LogType: "zpn_ast_comprehensive_stats",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.GetLSSLogTypeFormats(ctx, &zpa.GetLSSLogTypeFormatsArgs{
//				LogType: "zpn_waf_http_exchanges_log",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.GetLSSLogTypeFormats(ctx, &zpa.GetLSSLogTypeFormatsArgs{
//				LogType: "zpn_pbroker_comprehensive_stats",
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
//
// ### Read-Only
//
// The following arguments are supported:
//
// * `logType` - (Required) The type of log to be exported.
//   - `zpnTransLog`
//   - `zpnAuthLog`
//   - `zpnAstAuthLog`
//   - `zpnHttpTransLog`
//   - `zpnAuditLog`
//   - `zpnSysAuthLog`
//   - `zpnAstComprehensiveStats`
//   - `zpnWafHttpExchangesLog`
//   - `zpnPbrokerComprehensiveStats`
func GetLSSLogTypeFormats(ctx *pulumi.Context, args *GetLSSLogTypeFormatsArgs, opts ...pulumi.InvokeOption) (*GetLSSLogTypeFormatsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLSSLogTypeFormatsResult
	err := ctx.Invoke("zpa:index/getLSSLogTypeFormats:getLSSLogTypeFormats", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLSSLogTypeFormats.
type GetLSSLogTypeFormatsArgs struct {
	LogType string `pulumi:"logType"`
}

// A collection of values returned by getLSSLogTypeFormats.
type GetLSSLogTypeFormatsResult struct {
	Csv string `pulumi:"csv"`
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	Json    string `pulumi:"json"`
	LogType string `pulumi:"logType"`
	Tsv     string `pulumi:"tsv"`
}

func GetLSSLogTypeFormatsOutput(ctx *pulumi.Context, args GetLSSLogTypeFormatsOutputArgs, opts ...pulumi.InvokeOption) GetLSSLogTypeFormatsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLSSLogTypeFormatsResult, error) {
			args := v.(GetLSSLogTypeFormatsArgs)
			r, err := GetLSSLogTypeFormats(ctx, &args, opts...)
			var s GetLSSLogTypeFormatsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLSSLogTypeFormatsResultOutput)
}

// A collection of arguments for invoking getLSSLogTypeFormats.
type GetLSSLogTypeFormatsOutputArgs struct {
	LogType pulumi.StringInput `pulumi:"logType"`
}

func (GetLSSLogTypeFormatsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLSSLogTypeFormatsArgs)(nil)).Elem()
}

// A collection of values returned by getLSSLogTypeFormats.
type GetLSSLogTypeFormatsResultOutput struct{ *pulumi.OutputState }

func (GetLSSLogTypeFormatsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLSSLogTypeFormatsResult)(nil)).Elem()
}

func (o GetLSSLogTypeFormatsResultOutput) ToGetLSSLogTypeFormatsResultOutput() GetLSSLogTypeFormatsResultOutput {
	return o
}

func (o GetLSSLogTypeFormatsResultOutput) ToGetLSSLogTypeFormatsResultOutputWithContext(ctx context.Context) GetLSSLogTypeFormatsResultOutput {
	return o
}

func (o GetLSSLogTypeFormatsResultOutput) Csv() pulumi.StringOutput {
	return o.ApplyT(func(v GetLSSLogTypeFormatsResult) string { return v.Csv }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLSSLogTypeFormatsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLSSLogTypeFormatsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLSSLogTypeFormatsResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetLSSLogTypeFormatsResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o GetLSSLogTypeFormatsResultOutput) LogType() pulumi.StringOutput {
	return o.ApplyT(func(v GetLSSLogTypeFormatsResult) string { return v.LogType }).(pulumi.StringOutput)
}

func (o GetLSSLogTypeFormatsResultOutput) Tsv() pulumi.StringOutput {
	return o.ApplyT(func(v GetLSSLogTypeFormatsResult) string { return v.Tsv }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLSSLogTypeFormatsResultOutput{})
}
