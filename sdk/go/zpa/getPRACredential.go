// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// * [Official documentation](https://help.zscaler.com/zpa/about-privileged-credentials)
// * [API documentation](https://help.zscaler.com/zpa/configuring-privileged-credentials-using-api)
//
// The **zpa_pra_credential_controller** data source to get information about a privileged remote access credential created in the Zscaler Private Access cloud.
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
//			// Retrieves PRA Credential By ID
//			_, err := zpa.NewPRACredential(ctx, "this", nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupPRACredential(ctx *pulumi.Context, args *LookupPRACredentialArgs, opts ...pulumi.InvokeOption) (*LookupPRACredentialResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPRACredentialResult
	err := ctx.Invoke("zpa:index/getPRACredential:getPRACredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPRACredential.
type LookupPRACredentialArgs struct {
	Id *string `pulumi:"id"`
	// - (Optional) The name of the privileged credential.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getPRACredential.
type LookupPRACredentialResult struct {
	CreationTime            string  `pulumi:"creationTime"`
	CredentialType          string  `pulumi:"credentialType"`
	Description             string  `pulumi:"description"`
	Id                      *string `pulumi:"id"`
	LastCredentialResetTime string  `pulumi:"lastCredentialResetTime"`
	MicrotenantId           string  `pulumi:"microtenantId"`
	MicrotenantName         string  `pulumi:"microtenantName"`
	ModifiedBy              string  `pulumi:"modifiedBy"`
	ModifiedTime            string  `pulumi:"modifiedTime"`
	// - (Optional) The name of the privileged credential.
	Name       *string `pulumi:"name"`
	Password   string  `pulumi:"password"`
	UserDomain string  `pulumi:"userDomain"`
	Username   string  `pulumi:"username"`
}

func LookupPRACredentialOutput(ctx *pulumi.Context, args LookupPRACredentialOutputArgs, opts ...pulumi.InvokeOption) LookupPRACredentialResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPRACredentialResultOutput, error) {
			args := v.(LookupPRACredentialArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("zpa:index/getPRACredential:getPRACredential", args, LookupPRACredentialResultOutput{}, options).(LookupPRACredentialResultOutput), nil
		}).(LookupPRACredentialResultOutput)
}

// A collection of arguments for invoking getPRACredential.
type LookupPRACredentialOutputArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
	// - (Optional) The name of the privileged credential.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupPRACredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPRACredentialArgs)(nil)).Elem()
}

// A collection of values returned by getPRACredential.
type LookupPRACredentialResultOutput struct{ *pulumi.OutputState }

func (LookupPRACredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPRACredentialResult)(nil)).Elem()
}

func (o LookupPRACredentialResultOutput) ToLookupPRACredentialResultOutput() LookupPRACredentialResultOutput {
	return o
}

func (o LookupPRACredentialResultOutput) ToLookupPRACredentialResultOutputWithContext(ctx context.Context) LookupPRACredentialResultOutput {
	return o
}

func (o LookupPRACredentialResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPRACredentialResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupPRACredentialResultOutput) CredentialType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPRACredentialResult) string { return v.CredentialType }).(pulumi.StringOutput)
}

func (o LookupPRACredentialResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPRACredentialResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupPRACredentialResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPRACredentialResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupPRACredentialResultOutput) LastCredentialResetTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPRACredentialResult) string { return v.LastCredentialResetTime }).(pulumi.StringOutput)
}

func (o LookupPRACredentialResultOutput) MicrotenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPRACredentialResult) string { return v.MicrotenantId }).(pulumi.StringOutput)
}

func (o LookupPRACredentialResultOutput) MicrotenantName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPRACredentialResult) string { return v.MicrotenantName }).(pulumi.StringOutput)
}

func (o LookupPRACredentialResultOutput) ModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPRACredentialResult) string { return v.ModifiedBy }).(pulumi.StringOutput)
}

func (o LookupPRACredentialResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPRACredentialResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

// - (Optional) The name of the privileged credential.
func (o LookupPRACredentialResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPRACredentialResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupPRACredentialResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPRACredentialResult) string { return v.Password }).(pulumi.StringOutput)
}

func (o LookupPRACredentialResultOutput) UserDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPRACredentialResult) string { return v.UserDomain }).(pulumi.StringOutput)
}

func (o LookupPRACredentialResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPRACredentialResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPRACredentialResultOutput{})
}
