// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// Use the **zpa_policy_type** data source to get information about an a “policySetId“ and “policyType“. This data source is required when creating:
//
// 1. Access policy Rules
// 2. Access policy timeout rules
// 3. Access policy forwarding rules
// 4. Access policy inspection rules
//
// > **NOTE** The parameters “policySetId“ is required in all circumstances and is exported when checking for the policyType parameter. The policyType value is used for differentiating the policy types, in the request endpoint. The supported values are:
//
// * “ACCESS_POLICY/GLOBAL_POLICY“
// * “TIMEOUT_POLICY/REAUTH_POLICY“
// * “BYPASS_POLICY/CLIENT_FORWARDING_POLICY“
// * “INSPECTION_POLICY“
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
//			accessPolicy, err := zpa.GetPolicyType(ctx, &zpa.GetPolicyTypeArgs{
//				PolicyType: pulumi.StringRef("ACCESS_POLICY"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("zpaPolicyTypeAccessPolicy", accessPolicy.Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
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
//			globalPolicy, err := zpa.GetPolicyType(ctx, &zpa.GetPolicyTypeArgs{
//				PolicyType: pulumi.StringRef("GLOBAL_POLICY"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("zpaPolicyTypeAccessPolicy", globalPolicy.Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
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
//			timeoutPolicy, err := zpa.GetPolicyType(ctx, &zpa.GetPolicyTypeArgs{
//				PolicyType: pulumi.StringRef("TIMEOUT_POLICY"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("zpaPolicyTypeTimeoutPolicy", timeoutPolicy.Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
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
//			reauthPolicy, err := zpa.GetPolicyType(ctx, &zpa.GetPolicyTypeArgs{
//				PolicyType: pulumi.StringRef("REAUTH_POLICY"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("zpaPolicyTypeReauthPolicy", reauthPolicy.Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
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
//			clientForwardingPolicy, err := zpa.GetPolicyType(ctx, &zpa.GetPolicyTypeArgs{
//				PolicyType: pulumi.StringRef("CLIENT_FORWARDING_POLICY"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("zpaPolicyTypeClientForwardingPolicy", clientForwardingPolicy.Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
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
//			inspectionPolicy, err := zpa.GetPolicyType(ctx, &zpa.GetPolicyTypeArgs{
//				PolicyType: pulumi.StringRef("INSPECTION_POLICY"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("zpaPolicyTypeInspectionPolicy", inspectionPolicy.Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetPolicyType(ctx *pulumi.Context, args *GetPolicyTypeArgs, opts ...pulumi.InvokeOption) (*GetPolicyTypeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPolicyTypeResult
	err := ctx.Invoke("zpa:index/getPolicyType:getPolicyType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicyType.
type GetPolicyTypeArgs struct {
	Id *string `pulumi:"id"`
	// (string) The ID of the microtenant the resource is to be associated with.
	MicrotenantId *string `pulumi:"microtenantId"`
	// (string) The name of the microtenant the resource is to be associated with.
	MicrotenantName *string `pulumi:"microtenantName"`
	// The value for differentiating the policy types.
	PolicyType *string `pulumi:"policyType"`
}

// A collection of values returned by getPolicyType.
type GetPolicyTypeResult struct {
	CreationTime string  `pulumi:"creationTime"`
	Description  string  `pulumi:"description"`
	Enabled      bool    `pulumi:"enabled"`
	Id           *string `pulumi:"id"`
	// (string) The ID of the microtenant the resource is to be associated with.
	MicrotenantId *string `pulumi:"microtenantId"`
	// (string) The name of the microtenant the resource is to be associated with.
	MicrotenantName *string             `pulumi:"microtenantName"`
	ModifiedBy      string              `pulumi:"modifiedBy"`
	ModifiedTime    string              `pulumi:"modifiedTime"`
	Name            string              `pulumi:"name"`
	PolicyType      string              `pulumi:"policyType"`
	Rules           []GetPolicyTypeRule `pulumi:"rules"`
	Sorted          bool                `pulumi:"sorted"`
}

func GetPolicyTypeOutput(ctx *pulumi.Context, args GetPolicyTypeOutputArgs, opts ...pulumi.InvokeOption) GetPolicyTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPolicyTypeResult, error) {
			args := v.(GetPolicyTypeArgs)
			r, err := GetPolicyType(ctx, &args, opts...)
			var s GetPolicyTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPolicyTypeResultOutput)
}

// A collection of arguments for invoking getPolicyType.
type GetPolicyTypeOutputArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
	// (string) The ID of the microtenant the resource is to be associated with.
	MicrotenantId pulumi.StringPtrInput `pulumi:"microtenantId"`
	// (string) The name of the microtenant the resource is to be associated with.
	MicrotenantName pulumi.StringPtrInput `pulumi:"microtenantName"`
	// The value for differentiating the policy types.
	PolicyType pulumi.StringPtrInput `pulumi:"policyType"`
}

func (GetPolicyTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyTypeArgs)(nil)).Elem()
}

// A collection of values returned by getPolicyType.
type GetPolicyTypeResultOutput struct{ *pulumi.OutputState }

func (GetPolicyTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyTypeResult)(nil)).Elem()
}

func (o GetPolicyTypeResultOutput) ToGetPolicyTypeResultOutput() GetPolicyTypeResultOutput {
	return o
}

func (o GetPolicyTypeResultOutput) ToGetPolicyTypeResultOutputWithContext(ctx context.Context) GetPolicyTypeResultOutput {
	return o
}

func (o GetPolicyTypeResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyTypeResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o GetPolicyTypeResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyTypeResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetPolicyTypeResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetPolicyTypeResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o GetPolicyTypeResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicyTypeResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// (string) The ID of the microtenant the resource is to be associated with.
func (o GetPolicyTypeResultOutput) MicrotenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicyTypeResult) *string { return v.MicrotenantId }).(pulumi.StringPtrOutput)
}

// (string) The name of the microtenant the resource is to be associated with.
func (o GetPolicyTypeResultOutput) MicrotenantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicyTypeResult) *string { return v.MicrotenantName }).(pulumi.StringPtrOutput)
}

func (o GetPolicyTypeResultOutput) ModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyTypeResult) string { return v.ModifiedBy }).(pulumi.StringOutput)
}

func (o GetPolicyTypeResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyTypeResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o GetPolicyTypeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyTypeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetPolicyTypeResultOutput) PolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyTypeResult) string { return v.PolicyType }).(pulumi.StringOutput)
}

func (o GetPolicyTypeResultOutput) Rules() GetPolicyTypeRuleArrayOutput {
	return o.ApplyT(func(v GetPolicyTypeResult) []GetPolicyTypeRule { return v.Rules }).(GetPolicyTypeRuleArrayOutput)
}

func (o GetPolicyTypeResultOutput) Sorted() pulumi.BoolOutput {
	return o.ApplyT(func(v GetPolicyTypeResult) bool { return v.Sorted }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPolicyTypeResultOutput{})
}
