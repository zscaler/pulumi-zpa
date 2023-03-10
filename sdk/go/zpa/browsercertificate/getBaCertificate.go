// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package browsercertificate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the **zpa_ba_certificate** data source to get information about a browser access certificate created in the Zscaler Private Access cloud. This data source is required when creating a browser access application segment resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/BrowserCertificate"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := BrowserCertificate.GetBaCertificate(ctx, &browsercertificate.GetBaCertificateArgs{
//				Name: pulumi.StringRef("example.acme.com"),
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
//	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/BrowserCertificate"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := BrowserCertificate.GetBaCertificate(ctx, &browsercertificate.GetBaCertificateArgs{
//				Id: pulumi.StringRef("1234567890"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetBaCertificate(ctx *pulumi.Context, args *GetBaCertificateArgs, opts ...pulumi.InvokeOption) (*GetBaCertificateResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetBaCertificateResult
	err := ctx.Invoke("zpa:BrowserCertificate/getBaCertificate:getBaCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBaCertificate.
type GetBaCertificateArgs struct {
	// The id of the browser access certificate to be exported.
	Id *string `pulumi:"id"`
	// The name of the browser access certificate to be exported.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getBaCertificate.
type GetBaCertificateResult struct {
	// (string)
	CertChain string `pulumi:"certChain"`
	// (string)
	Cname string `pulumi:"cname"`
	// (string)
	CreationTime string `pulumi:"creationTime"`
	// (string)
	Description string  `pulumi:"description"`
	Id          *string `pulumi:"id"`
	// (string)
	IssuedBy string `pulumi:"issuedBy"`
	// (string)
	IssuedTo string `pulumi:"issuedTo"`
	// (string)
	ModifiedTime string `pulumi:"modifiedTime"`
	// (string)
	Modifiedby string  `pulumi:"modifiedby"`
	Name       *string `pulumi:"name"`
	// (string)
	Sans []string `pulumi:"sans"`
	// (string)
	SerialNo string `pulumi:"serialNo"`
	// (string)
	Status string `pulumi:"status"`
	// (string)
	ValidFromInEpochsec string `pulumi:"validFromInEpochsec"`
	// (string)
	ValidToInEpochsec string `pulumi:"validToInEpochsec"`
}

func GetBaCertificateOutput(ctx *pulumi.Context, args GetBaCertificateOutputArgs, opts ...pulumi.InvokeOption) GetBaCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBaCertificateResult, error) {
			args := v.(GetBaCertificateArgs)
			r, err := GetBaCertificate(ctx, &args, opts...)
			var s GetBaCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBaCertificateResultOutput)
}

// A collection of arguments for invoking getBaCertificate.
type GetBaCertificateOutputArgs struct {
	// The id of the browser access certificate to be exported.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the browser access certificate to be exported.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetBaCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBaCertificateArgs)(nil)).Elem()
}

// A collection of values returned by getBaCertificate.
type GetBaCertificateResultOutput struct{ *pulumi.OutputState }

func (GetBaCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBaCertificateResult)(nil)).Elem()
}

func (o GetBaCertificateResultOutput) ToGetBaCertificateResultOutput() GetBaCertificateResultOutput {
	return o
}

func (o GetBaCertificateResultOutput) ToGetBaCertificateResultOutputWithContext(ctx context.Context) GetBaCertificateResultOutput {
	return o
}

// (string)
func (o GetBaCertificateResultOutput) CertChain() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaCertificateResult) string { return v.CertChain }).(pulumi.StringOutput)
}

// (string)
func (o GetBaCertificateResultOutput) Cname() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaCertificateResult) string { return v.Cname }).(pulumi.StringOutput)
}

// (string)
func (o GetBaCertificateResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaCertificateResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// (string)
func (o GetBaCertificateResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaCertificateResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetBaCertificateResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBaCertificateResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// (string)
func (o GetBaCertificateResultOutput) IssuedBy() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaCertificateResult) string { return v.IssuedBy }).(pulumi.StringOutput)
}

// (string)
func (o GetBaCertificateResultOutput) IssuedTo() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaCertificateResult) string { return v.IssuedTo }).(pulumi.StringOutput)
}

// (string)
func (o GetBaCertificateResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaCertificateResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

// (string)
func (o GetBaCertificateResultOutput) Modifiedby() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaCertificateResult) string { return v.Modifiedby }).(pulumi.StringOutput)
}

func (o GetBaCertificateResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBaCertificateResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// (string)
func (o GetBaCertificateResultOutput) Sans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBaCertificateResult) []string { return v.Sans }).(pulumi.StringArrayOutput)
}

// (string)
func (o GetBaCertificateResultOutput) SerialNo() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaCertificateResult) string { return v.SerialNo }).(pulumi.StringOutput)
}

// (string)
func (o GetBaCertificateResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaCertificateResult) string { return v.Status }).(pulumi.StringOutput)
}

// (string)
func (o GetBaCertificateResultOutput) ValidFromInEpochsec() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaCertificateResult) string { return v.ValidFromInEpochsec }).(pulumi.StringOutput)
}

// (string)
func (o GetBaCertificateResultOutput) ValidToInEpochsec() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaCertificateResult) string { return v.ValidToInEpochsec }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBaCertificateResultOutput{})
}
