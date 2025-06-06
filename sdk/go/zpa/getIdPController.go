// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// * [Official documentation](https://help.zscaler.com/zpa/identity-management)
// * [API documentation](https://help.zscaler.com/zpa/obtaining-idp-configuration-details-using-api)
//
// Use the **zpa_idp_controller** data source to get information about an Identity Provider created in the Zscaler Private Access cloud. This data source is required when creating:
//
// **NOTE:** To ensure consistent search results across data sources, please avoid using multiple spaces or special characters in your search queries.
//
// * Access policy Rules
// * Access policy timeout rules
// * Access policy forwarding rules
// * Access policy inspection rules
// * Access policy isolation rules
// * Access policy privileged credentials rules
// * Access policy privileged capabilities rules
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
	Id   *string `pulumi:"id"`
	Name *string `pulumi:"name"`
}

// A collection of values returned by getIdPController.
type GetIdPControllerResult struct {
	AdminMetadatas              []GetIdPControllerAdminMetadata `pulumi:"adminMetadatas"`
	AdminSpSigningCertId        string                          `pulumi:"adminSpSigningCertId"`
	AutoProvision               string                          `pulumi:"autoProvision"`
	CreationTime                string                          `pulumi:"creationTime"`
	Description                 string                          `pulumi:"description"`
	DisableSamlBasedPolicy      bool                            `pulumi:"disableSamlBasedPolicy"`
	DomainLists                 []string                        `pulumi:"domainLists"`
	EnableArbitraryAuthDomains  string                          `pulumi:"enableArbitraryAuthDomains"`
	EnableScimBasedPolicy       bool                            `pulumi:"enableScimBasedPolicy"`
	Enabled                     bool                            `pulumi:"enabled"`
	ForceAuth                   bool                            `pulumi:"forceAuth"`
	Id                          string                          `pulumi:"id"`
	IdpEntityId                 string                          `pulumi:"idpEntityId"`
	LoginHint                   bool                            `pulumi:"loginHint"`
	LoginNameAttribute          string                          `pulumi:"loginNameAttribute"`
	LoginUrl                    string                          `pulumi:"loginUrl"`
	ModifiedTime                string                          `pulumi:"modifiedTime"`
	Modifiedby                  string                          `pulumi:"modifiedby"`
	Name                        string                          `pulumi:"name"`
	ReauthOnUserUpdate          bool                            `pulumi:"reauthOnUserUpdate"`
	RedirectBinding             bool                            `pulumi:"redirectBinding"`
	ScimEnabled                 bool                            `pulumi:"scimEnabled"`
	ScimServiceProviderEndpoint string                          `pulumi:"scimServiceProviderEndpoint"`
	ScimSharedSecretExists      bool                            `pulumi:"scimSharedSecretExists"`
	SignSamlRequest             string                          `pulumi:"signSamlRequest"`
	SsoTypes                    []string                        `pulumi:"ssoTypes"`
	UseCustomSpMetadata         bool                            `pulumi:"useCustomSpMetadata"`
	UserMetadatas               []GetIdPControllerUserMetadata  `pulumi:"userMetadatas"`
	UserSpSigningCertId         string                          `pulumi:"userSpSigningCertId"`
}

func GetIdPControllerOutput(ctx *pulumi.Context, args GetIdPControllerOutputArgs, opts ...pulumi.InvokeOption) GetIdPControllerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetIdPControllerResultOutput, error) {
			args := v.(GetIdPControllerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("zpa:index/getIdPController:getIdPController", args, GetIdPControllerResultOutput{}, options).(GetIdPControllerResultOutput), nil
		}).(GetIdPControllerResultOutput)
}

// A collection of arguments for invoking getIdPController.
type GetIdPControllerOutputArgs struct {
	Id   pulumi.StringPtrInput `pulumi:"id"`
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

func (o GetIdPControllerResultOutput) AdminMetadatas() GetIdPControllerAdminMetadataArrayOutput {
	return o.ApplyT(func(v GetIdPControllerResult) []GetIdPControllerAdminMetadata { return v.AdminMetadatas }).(GetIdPControllerAdminMetadataArrayOutput)
}

func (o GetIdPControllerResultOutput) AdminSpSigningCertId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.AdminSpSigningCertId }).(pulumi.StringOutput)
}

func (o GetIdPControllerResultOutput) AutoProvision() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.AutoProvision }).(pulumi.StringOutput)
}

func (o GetIdPControllerResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o GetIdPControllerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetIdPControllerResultOutput) DisableSamlBasedPolicy() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIdPControllerResult) bool { return v.DisableSamlBasedPolicy }).(pulumi.BoolOutput)
}

func (o GetIdPControllerResultOutput) DomainLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIdPControllerResult) []string { return v.DomainLists }).(pulumi.StringArrayOutput)
}

func (o GetIdPControllerResultOutput) EnableArbitraryAuthDomains() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.EnableArbitraryAuthDomains }).(pulumi.StringOutput)
}

func (o GetIdPControllerResultOutput) EnableScimBasedPolicy() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIdPControllerResult) bool { return v.EnableScimBasedPolicy }).(pulumi.BoolOutput)
}

func (o GetIdPControllerResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIdPControllerResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o GetIdPControllerResultOutput) ForceAuth() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIdPControllerResult) bool { return v.ForceAuth }).(pulumi.BoolOutput)
}

func (o GetIdPControllerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIdPControllerResultOutput) IdpEntityId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.IdpEntityId }).(pulumi.StringOutput)
}

func (o GetIdPControllerResultOutput) LoginHint() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIdPControllerResult) bool { return v.LoginHint }).(pulumi.BoolOutput)
}

func (o GetIdPControllerResultOutput) LoginNameAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.LoginNameAttribute }).(pulumi.StringOutput)
}

func (o GetIdPControllerResultOutput) LoginUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.LoginUrl }).(pulumi.StringOutput)
}

func (o GetIdPControllerResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o GetIdPControllerResultOutput) Modifiedby() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.Modifiedby }).(pulumi.StringOutput)
}

func (o GetIdPControllerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetIdPControllerResultOutput) ReauthOnUserUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIdPControllerResult) bool { return v.ReauthOnUserUpdate }).(pulumi.BoolOutput)
}

func (o GetIdPControllerResultOutput) RedirectBinding() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIdPControllerResult) bool { return v.RedirectBinding }).(pulumi.BoolOutput)
}

func (o GetIdPControllerResultOutput) ScimEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIdPControllerResult) bool { return v.ScimEnabled }).(pulumi.BoolOutput)
}

func (o GetIdPControllerResultOutput) ScimServiceProviderEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.ScimServiceProviderEndpoint }).(pulumi.StringOutput)
}

func (o GetIdPControllerResultOutput) ScimSharedSecretExists() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIdPControllerResult) bool { return v.ScimSharedSecretExists }).(pulumi.BoolOutput)
}

func (o GetIdPControllerResultOutput) SignSamlRequest() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.SignSamlRequest }).(pulumi.StringOutput)
}

func (o GetIdPControllerResultOutput) SsoTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIdPControllerResult) []string { return v.SsoTypes }).(pulumi.StringArrayOutput)
}

func (o GetIdPControllerResultOutput) UseCustomSpMetadata() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIdPControllerResult) bool { return v.UseCustomSpMetadata }).(pulumi.BoolOutput)
}

func (o GetIdPControllerResultOutput) UserMetadatas() GetIdPControllerUserMetadataArrayOutput {
	return o.ApplyT(func(v GetIdPControllerResult) []GetIdPControllerUserMetadata { return v.UserMetadatas }).(GetIdPControllerUserMetadataArrayOutput)
}

func (o GetIdPControllerResultOutput) UserSpSigningCertId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdPControllerResult) string { return v.UserSpSigningCertId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIdPControllerResultOutput{})
}
