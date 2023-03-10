// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package samlattribute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the **zpa_saml_attribute** data source to get information about a SAML Attributes from an Identity Provider (IdP). This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/SamlAttribute"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := SamlAttribute.GetSAMLAttribute(ctx, &samlattribute.GetSAMLAttributeArgs{
//				Name: pulumi.StringRef("Email_User SSO"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/SamlAttribute"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := SamlAttribute.GetSAMLAttribute(ctx, &samlattribute.GetSAMLAttributeArgs{
//				Name: pulumi.StringRef("DepartmentName_IdP_Name_Users"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSAMLAttribute(ctx *pulumi.Context, args *GetSAMLAttributeArgs, opts ...pulumi.InvokeOption) (*GetSAMLAttributeResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSAMLAttributeResult
	err := ctx.Invoke("zpa:SamlAttribute/getSAMLAttribute:getSAMLAttribute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSAMLAttribute.
type GetSAMLAttributeArgs struct {
	// The ID of the machine group to be exported.
	Id *string `pulumi:"id"`
	// The name of the IdP corresponding to the SAML attribute.
	IdpName *string `pulumi:"idpName"`
	// The name of the saml attribute to be exported.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getSAMLAttribute.
type GetSAMLAttributeResult struct {
	// (string)
	CreationTime string `pulumi:"creationTime"`
	Id           string `pulumi:"id"`
	// (string) The ID of the IdP corresponding to the SAML attribute.
	IdpId   string  `pulumi:"idpId"`
	IdpName *string `pulumi:"idpName"`
	// (string)
	ModifiedTime string `pulumi:"modifiedTime"`
	Modifiedby   string `pulumi:"modifiedby"`
	// (string)
	Name string `pulumi:"name"`
	// (string)
	SamlName string `pulumi:"samlName"`
	// (string)
	UserAttribute bool `pulumi:"userAttribute"`
}

func GetSAMLAttributeOutput(ctx *pulumi.Context, args GetSAMLAttributeOutputArgs, opts ...pulumi.InvokeOption) GetSAMLAttributeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSAMLAttributeResult, error) {
			args := v.(GetSAMLAttributeArgs)
			r, err := GetSAMLAttribute(ctx, &args, opts...)
			var s GetSAMLAttributeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSAMLAttributeResultOutput)
}

// A collection of arguments for invoking getSAMLAttribute.
type GetSAMLAttributeOutputArgs struct {
	// The ID of the machine group to be exported.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the IdP corresponding to the SAML attribute.
	IdpName pulumi.StringPtrInput `pulumi:"idpName"`
	// The name of the saml attribute to be exported.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetSAMLAttributeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSAMLAttributeArgs)(nil)).Elem()
}

// A collection of values returned by getSAMLAttribute.
type GetSAMLAttributeResultOutput struct{ *pulumi.OutputState }

func (GetSAMLAttributeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSAMLAttributeResult)(nil)).Elem()
}

func (o GetSAMLAttributeResultOutput) ToGetSAMLAttributeResultOutput() GetSAMLAttributeResultOutput {
	return o
}

func (o GetSAMLAttributeResultOutput) ToGetSAMLAttributeResultOutputWithContext(ctx context.Context) GetSAMLAttributeResultOutput {
	return o
}

// (string)
func (o GetSAMLAttributeResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSAMLAttributeResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o GetSAMLAttributeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSAMLAttributeResult) string { return v.Id }).(pulumi.StringOutput)
}

// (string) The ID of the IdP corresponding to the SAML attribute.
func (o GetSAMLAttributeResultOutput) IdpId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSAMLAttributeResult) string { return v.IdpId }).(pulumi.StringOutput)
}

func (o GetSAMLAttributeResultOutput) IdpName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSAMLAttributeResult) *string { return v.IdpName }).(pulumi.StringPtrOutput)
}

// (string)
func (o GetSAMLAttributeResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSAMLAttributeResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o GetSAMLAttributeResultOutput) Modifiedby() pulumi.StringOutput {
	return o.ApplyT(func(v GetSAMLAttributeResult) string { return v.Modifiedby }).(pulumi.StringOutput)
}

// (string)
func (o GetSAMLAttributeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSAMLAttributeResult) string { return v.Name }).(pulumi.StringOutput)
}

// (string)
func (o GetSAMLAttributeResultOutput) SamlName() pulumi.StringOutput {
	return o.ApplyT(func(v GetSAMLAttributeResult) string { return v.SamlName }).(pulumi.StringOutput)
}

// (string)
func (o GetSAMLAttributeResultOutput) UserAttribute() pulumi.BoolOutput {
	return o.ApplyT(func(v GetSAMLAttributeResult) bool { return v.UserAttribute }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSAMLAttributeResultOutput{})
}
