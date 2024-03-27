// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// Use the **zpa_idp_controller** data source to get information about an Identity Provider created in the Zscaler Private Access cloud. This data source is required when creating:
//
// 1. Access policy Rules
// 2. Access policy timeout rules
// 3. Access policy forwarding rules
// 4. Access policy inspection rules
// 5. Access policy isolation rules
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
//			_, err := zpa.GetIdPController(ctx, &zpa.GetIdPControllerArgs{
//				Name: pulumi.StringRef("idp_name"),
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
//			_, err := zpa.GetIdPController(ctx, &zpa.GetIdPControllerArgs{
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
// <!--End PulumiCodeChooser -->
func GetIdPController(ctx *pulumi.Context, args *GetIdPControllerArgs, opts ...pulumi.InvokeOption) (*GetIdPControllerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIdPControllerResult
	err := ctx.Invoke("zpa:index/getIdPController:getIdPController", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIdPController.
type GetIdPControllerArgs struct {
	// The name of the Identity Provider (IdP) to be exported.
	Id *string `pulumi:"id"`
	// The name of the Identity Provider (IdP) to be exported.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getIdPController.
type GetIdPControllerResult struct {
	// (Computed)
	AdminMetadatas       []GetIdPControllerAdminMetadata `pulumi:"adminMetadatas"`
	AdminSpSigningCertId string                          `pulumi:"adminSpSigningCertId"`
	// (string)
	AutoProvision string `pulumi:"autoProvision"`
	// (string)
	CreationTime string `pulumi:"creationTime"`
	// (string)
	Description string `pulumi:"description"`
	// (bool)
	DisableSamlBasedPolicy bool `pulumi:"disableSamlBasedPolicy"`
	// (string)
	DomainLists []string `pulumi:"domainLists"`
	// (string)
	EnableArbitraryAuthDomains string `pulumi:"enableArbitraryAuthDomains"`
	// (bool)
	EnableScimBasedPolicy bool `pulumi:"enableScimBasedPolicy"`
	// (bool) Default value if null is True
	Enabled bool `pulumi:"enabled"`
	// (bool)
	ForceAuth bool   `pulumi:"forceAuth"`
	Id        string `pulumi:"id"`
	// (string)
	IdpEntityId string `pulumi:"idpEntityId"`
	// (bool)
	LoginHint bool `pulumi:"loginHint"`
	// (string)
	LoginNameAttribute string `pulumi:"loginNameAttribute"`
	// (string)
	LoginUrl string `pulumi:"loginUrl"`
	// (string)
	ModifiedTime string `pulumi:"modifiedTime"`
	Modifiedby   string `pulumi:"modifiedby"`
	Name         string `pulumi:"name"`
	// (bool)
	ReauthOnUserUpdate bool `pulumi:"reauthOnUserUpdate"`
	// (bool)
	RedirectBinding bool `pulumi:"redirectBinding"`
	// (bool)
	ScimEnabled bool `pulumi:"scimEnabled"`
	// (string)
	ScimServiceProviderEndpoint string `pulumi:"scimServiceProviderEndpoint"`
	// (bool)
	ScimSharedSecretExists bool `pulumi:"scimSharedSecretExists"`
	// (string)
	SignSamlRequest string `pulumi:"signSamlRequest"`
	// (string)
	SsoTypes []string `pulumi:"ssoTypes"`
	// (bool)
	UseCustomSpMetadata bool `pulumi:"useCustomSpMetadata"`
	// (Computed)
	UserMetadatas       []GetIdPControllerUserMetadata `pulumi:"userMetadatas"`
	UserSpSigningCertId string                         `pulumi:"userSpSigningCertId"`
}

func GetIdPControllerOutput(ctx *pulumi.Context, args GetIdPControllerOutputArgs, opts ...pulumi.InvokeOption) GetIdPControllerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIdPControllerResult, error) {
			args := v.(GetIdPControllerArgs)
			r, err := GetIdPController(ctx, &args, opts...)
			var s GetIdPControllerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIdPControllerResultOutput)
}

// A collection of arguments for invoking getIdPController.
type GetIdPControllerOutputArgs struct {
	// The name of the Identity Provider (IdP) to be exported.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the Identity Provider (IdP) to be exported.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetIdPControllerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIdPControllerArgs)(nil)).Elem()
}

// A collection of values returned by getIdPController.
type GetIdPControllerResultOutput struct{ *pulumi.OutputState }

func (GetIdPControllerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIdPControllerResult)(nil)).Elem()
}

func (o GetIdPControllerResultOutput) ToGetIdPControllerResultOutput() GetIdPControllerResultOutput {
	return o
}

func (o GetIdPControllerResultOutput) ToGetIdPControllerResultOutputWithContext(ctx context.Context) GetIdPControllerResultOutput {
	return o
}

// (Computed)
func (o GetIdPControllerResultOutput) AdminMetadatas() GetIdPControllerAdminMetadataArrayOutput {
	return o.ApplyT(func(v GetIdPControllerResult) []GetIdPControllerAdminMetadata { return v.AdminMetadatas }).(GetIdPControllerAdminMetadataArrayOutput)
}

func (o GetIdPControllerResultOutput) AdminSpSigningCertId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.AdminSpSigningCertId }).(pulumi.StringOutput)
}

// (string)
func (o GetIdPControllerResultOutput) AutoProvision() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.AutoProvision }).(pulumi.StringOutput)
}

// (string)
func (o GetIdPControllerResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// (string)
func (o GetIdPControllerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.Description }).(pulumi.StringOutput)
}

// (bool)
func (o GetIdPControllerResultOutput) DisableSamlBasedPolicy() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIdPControllerResult) bool { return v.DisableSamlBasedPolicy }).(pulumi.BoolOutput)
}

// (string)
func (o GetIdPControllerResultOutput) DomainLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIdPControllerResult) []string { return v.DomainLists }).(pulumi.StringArrayOutput)
}

// (string)
func (o GetIdPControllerResultOutput) EnableArbitraryAuthDomains() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.EnableArbitraryAuthDomains }).(pulumi.StringOutput)
}

// (bool)
func (o GetIdPControllerResultOutput) EnableScimBasedPolicy() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIdPControllerResult) bool { return v.EnableScimBasedPolicy }).(pulumi.BoolOutput)
}

// (bool) Default value if null is True
func (o GetIdPControllerResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIdPControllerResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// (bool)
func (o GetIdPControllerResultOutput) ForceAuth() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIdPControllerResult) bool { return v.ForceAuth }).(pulumi.BoolOutput)
}

func (o GetIdPControllerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.Id }).(pulumi.StringOutput)
}

// (string)
func (o GetIdPControllerResultOutput) IdpEntityId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.IdpEntityId }).(pulumi.StringOutput)
}

// (bool)
func (o GetIdPControllerResultOutput) LoginHint() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIdPControllerResult) bool { return v.LoginHint }).(pulumi.BoolOutput)
}

// (string)
func (o GetIdPControllerResultOutput) LoginNameAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.LoginNameAttribute }).(pulumi.StringOutput)
}

// (string)
func (o GetIdPControllerResultOutput) LoginUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.LoginUrl }).(pulumi.StringOutput)
}

// (string)
func (o GetIdPControllerResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o GetIdPControllerResultOutput) Modifiedby() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.Modifiedby }).(pulumi.StringOutput)
}

func (o GetIdPControllerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.Name }).(pulumi.StringOutput)
}

// (bool)
func (o GetIdPControllerResultOutput) ReauthOnUserUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIdPControllerResult) bool { return v.ReauthOnUserUpdate }).(pulumi.BoolOutput)
}

// (bool)
func (o GetIdPControllerResultOutput) RedirectBinding() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIdPControllerResult) bool { return v.RedirectBinding }).(pulumi.BoolOutput)
}

// (bool)
func (o GetIdPControllerResultOutput) ScimEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIdPControllerResult) bool { return v.ScimEnabled }).(pulumi.BoolOutput)
}

// (string)
func (o GetIdPControllerResultOutput) ScimServiceProviderEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.ScimServiceProviderEndpoint }).(pulumi.StringOutput)
}

// (bool)
func (o GetIdPControllerResultOutput) ScimSharedSecretExists() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIdPControllerResult) bool { return v.ScimSharedSecretExists }).(pulumi.BoolOutput)
}

// (string)
func (o GetIdPControllerResultOutput) SignSamlRequest() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.SignSamlRequest }).(pulumi.StringOutput)
}

// (string)
func (o GetIdPControllerResultOutput) SsoTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIdPControllerResult) []string { return v.SsoTypes }).(pulumi.StringArrayOutput)
}

// (bool)
func (o GetIdPControllerResultOutput) UseCustomSpMetadata() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIdPControllerResult) bool { return v.UseCustomSpMetadata }).(pulumi.BoolOutput)
}

// (Computed)
func (o GetIdPControllerResultOutput) UserMetadatas() GetIdPControllerUserMetadataArrayOutput {
	return o.ApplyT(func(v GetIdPControllerResult) []GetIdPControllerUserMetadata { return v.UserMetadatas }).(GetIdPControllerUserMetadataArrayOutput)
}

func (o GetIdPControllerResultOutput) UserSpSigningCertId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.UserSpSigningCertId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIdPControllerResultOutput{})
}