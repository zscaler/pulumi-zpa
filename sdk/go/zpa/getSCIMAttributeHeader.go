// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// * [Official documentation](https://help.zscaler.com/zpa/about-scim)
// * [API documentation](https://help.zscaler.com/zpa/obtaining-scim-attribute-details-using-api)
//
// Use the **zpa_scim_attribute_header** data source to get information about a SCIM attribute from an Identity Provider (IdP). This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Inspection Policy.
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
//			_, err := zpa.GetSCIMAttributeHeader(ctx, &zpa.GetSCIMAttributeHeaderArgs{
//				IdpName: pulumi.StringRef("IdP_Name"),
//				Name:    pulumi.StringRef("name.givenName"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.GetSCIMAttributeHeader(ctx, &zpa.GetSCIMAttributeHeaderArgs{
//				IdpName: pulumi.StringRef("IdP_Name"),
//				Name:    pulumi.StringRef("name.familyName"),
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
func GetSCIMAttributeHeader(ctx *pulumi.Context, args *GetSCIMAttributeHeaderArgs, opts ...pulumi.InvokeOption) (*GetSCIMAttributeHeaderResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSCIMAttributeHeaderResult
	err := ctx.Invoke("zpa:index/getSCIMAttributeHeader:getSCIMAttributeHeader", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSCIMAttributeHeader.
type GetSCIMAttributeHeaderArgs struct {
	IdpId   *string `pulumi:"idpId"`
	IdpName *string `pulumi:"idpName"`
	Name    *string `pulumi:"name"`
}

// A collection of values returned by getSCIMAttributeHeader.
type GetSCIMAttributeHeaderResult struct {
	CanonicalValues []string `pulumi:"canonicalValues"`
	CaseSensitive   bool     `pulumi:"caseSensitive"`
	CreationTime    string   `pulumi:"creationTime"`
	DataType        string   `pulumi:"dataType"`
	Description     string   `pulumi:"description"`
	Id              string   `pulumi:"id"`
	IdpId           *string  `pulumi:"idpId"`
	IdpName         *string  `pulumi:"idpName"`
	ModifiedTime    string   `pulumi:"modifiedTime"`
	Modifiedby      string   `pulumi:"modifiedby"`
	Multivalued     bool     `pulumi:"multivalued"`
	Mutability      string   `pulumi:"mutability"`
	Name            *string  `pulumi:"name"`
	Required        bool     `pulumi:"required"`
	Returned        string   `pulumi:"returned"`
	SchemaUri       string   `pulumi:"schemaUri"`
	Uniqueness      bool     `pulumi:"uniqueness"`
	Values          []string `pulumi:"values"`
}

func GetSCIMAttributeHeaderOutput(ctx *pulumi.Context, args GetSCIMAttributeHeaderOutputArgs, opts ...pulumi.InvokeOption) GetSCIMAttributeHeaderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSCIMAttributeHeaderResult, error) {
			args := v.(GetSCIMAttributeHeaderArgs)
			r, err := GetSCIMAttributeHeader(ctx, &args, opts...)
			var s GetSCIMAttributeHeaderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSCIMAttributeHeaderResultOutput)
}

// A collection of arguments for invoking getSCIMAttributeHeader.
type GetSCIMAttributeHeaderOutputArgs struct {
	IdpId   pulumi.StringPtrInput `pulumi:"idpId"`
	IdpName pulumi.StringPtrInput `pulumi:"idpName"`
	Name    pulumi.StringPtrInput `pulumi:"name"`
}

func (GetSCIMAttributeHeaderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSCIMAttributeHeaderArgs)(nil)).Elem()
}

// A collection of values returned by getSCIMAttributeHeader.
type GetSCIMAttributeHeaderResultOutput struct{ *pulumi.OutputState }

func (GetSCIMAttributeHeaderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSCIMAttributeHeaderResult)(nil)).Elem()
}

func (o GetSCIMAttributeHeaderResultOutput) ToGetSCIMAttributeHeaderResultOutput() GetSCIMAttributeHeaderResultOutput {
	return o
}

func (o GetSCIMAttributeHeaderResultOutput) ToGetSCIMAttributeHeaderResultOutputWithContext(ctx context.Context) GetSCIMAttributeHeaderResultOutput {
	return o
}

func (o GetSCIMAttributeHeaderResultOutput) CanonicalValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSCIMAttributeHeaderResult) []string { return v.CanonicalValues }).(pulumi.StringArrayOutput)
}

func (o GetSCIMAttributeHeaderResultOutput) CaseSensitive() pulumi.BoolOutput {
	return o.ApplyT(func(v GetSCIMAttributeHeaderResult) bool { return v.CaseSensitive }).(pulumi.BoolOutput)
}

func (o GetSCIMAttributeHeaderResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSCIMAttributeHeaderResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o GetSCIMAttributeHeaderResultOutput) DataType() pulumi.StringOutput {
	return o.ApplyT(func(v GetSCIMAttributeHeaderResult) string { return v.DataType }).(pulumi.StringOutput)
}

func (o GetSCIMAttributeHeaderResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetSCIMAttributeHeaderResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetSCIMAttributeHeaderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSCIMAttributeHeaderResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSCIMAttributeHeaderResultOutput) IdpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSCIMAttributeHeaderResult) *string { return v.IdpId }).(pulumi.StringPtrOutput)
}

func (o GetSCIMAttributeHeaderResultOutput) IdpName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSCIMAttributeHeaderResult) *string { return v.IdpName }).(pulumi.StringPtrOutput)
}

func (o GetSCIMAttributeHeaderResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSCIMAttributeHeaderResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o GetSCIMAttributeHeaderResultOutput) Modifiedby() pulumi.StringOutput {
	return o.ApplyT(func(v GetSCIMAttributeHeaderResult) string { return v.Modifiedby }).(pulumi.StringOutput)
}

func (o GetSCIMAttributeHeaderResultOutput) Multivalued() pulumi.BoolOutput {
	return o.ApplyT(func(v GetSCIMAttributeHeaderResult) bool { return v.Multivalued }).(pulumi.BoolOutput)
}

func (o GetSCIMAttributeHeaderResultOutput) Mutability() pulumi.StringOutput {
	return o.ApplyT(func(v GetSCIMAttributeHeaderResult) string { return v.Mutability }).(pulumi.StringOutput)
}

func (o GetSCIMAttributeHeaderResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSCIMAttributeHeaderResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetSCIMAttributeHeaderResultOutput) Required() pulumi.BoolOutput {
	return o.ApplyT(func(v GetSCIMAttributeHeaderResult) bool { return v.Required }).(pulumi.BoolOutput)
}

func (o GetSCIMAttributeHeaderResultOutput) Returned() pulumi.StringOutput {
	return o.ApplyT(func(v GetSCIMAttributeHeaderResult) string { return v.Returned }).(pulumi.StringOutput)
}

func (o GetSCIMAttributeHeaderResultOutput) SchemaUri() pulumi.StringOutput {
	return o.ApplyT(func(v GetSCIMAttributeHeaderResult) string { return v.SchemaUri }).(pulumi.StringOutput)
}

func (o GetSCIMAttributeHeaderResultOutput) Uniqueness() pulumi.BoolOutput {
	return o.ApplyT(func(v GetSCIMAttributeHeaderResult) bool { return v.Uniqueness }).(pulumi.BoolOutput)
}

func (o GetSCIMAttributeHeaderResultOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSCIMAttributeHeaderResult) []string { return v.Values }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSCIMAttributeHeaderResultOutput{})
}
