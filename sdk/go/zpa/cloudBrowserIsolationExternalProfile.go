// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// * [Official documentation](https://help.zscaler.com/isolation/about-custom-root-certificates-cloud-browser-isolation)
//
// The **zpa_cloud_browser_isolation_external_profile** resource creates a Cloud Browser Isolation external profile. This resource can then be used in as part of `PolicyAccessIsolationRule` when the `action` attribute is set to `ISOLATE`.
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
//			thisCloudBrowserIsolationBanner, err := zpa.LookupCloudBrowserIsolationBanner(ctx, &zpa.LookupCloudBrowserIsolationBannerArgs{
//				Name: pulumi.StringRef("Default"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			singapore, err := zpa.GetCloudBrowserIsolationRegion(ctx, &zpa.GetCloudBrowserIsolationRegionArgs{
//				Name: pulumi.StringRef("Singapore"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.GetCloudBrowserIsolationRegion(ctx, &zpa.GetCloudBrowserIsolationRegionArgs{
//				Name: pulumi.StringRef("Frankfurt"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			thisCloudBrowserIsolationCertificate, err := zpa.LookupCloudBrowserIsolationCertificate(ctx, &zpa.LookupCloudBrowserIsolationCertificateArgs{
//				Name: pulumi.StringRef("Zscaler Root Certificate"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.NewCloudBrowserIsolationExternalProfile(ctx, "thisCloudBrowserIsolationExternalProfile", &zpa.CloudBrowserIsolationExternalProfileArgs{
//				Description: pulumi.String("CBI_Profile_Example"),
//				BannerId:    pulumi.String(thisCloudBrowserIsolationBanner.Id),
//				RegionIds: pulumi.StringArray{
//					pulumi.String(singapore.Id),
//				},
//				CertificateIds: pulumi.StringArray{
//					pulumi.String(thisCloudBrowserIsolationCertificate.Id),
//				},
//				UserExperiences: zpa.CloudBrowserIsolationExternalProfileUserExperienceArray{
//					&zpa.CloudBrowserIsolationExternalProfileUserExperienceArgs{
//						SessionPersistence: pulumi.Bool(true),
//						BrowserInBrowser:   pulumi.Bool(true),
//					},
//				},
//				SecurityControls: zpa.CloudBrowserIsolationExternalProfileSecurityControlArray{
//					&zpa.CloudBrowserIsolationExternalProfileSecurityControlArgs{
//						CopyPaste:          pulumi.String("all"),
//						UploadDownload:     pulumi.String("all"),
//						DocumentViewer:     pulumi.Bool(true),
//						LocalRender:        pulumi.Bool(true),
//						AllowPrinting:      pulumi.Bool(true),
//						RestrictKeystrokes: pulumi.Bool(false),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
type CloudBrowserIsolationExternalProfile struct {
	pulumi.CustomResourceState

	BannerId pulumi.StringOutput `pulumi:"bannerId"`
	// This field defines the list of server groups IDs.
	CertificateIds pulumi.StringArrayOutput `pulumi:"certificateIds"`
	Description    pulumi.StringPtrOutput   `pulumi:"description"`
	Name           pulumi.StringOutput      `pulumi:"name"`
	// This field defines the list of server groups IDs.
	RegionIds        pulumi.StringArrayOutput                                       `pulumi:"regionIds"`
	SecurityControls CloudBrowserIsolationExternalProfileSecurityControlArrayOutput `pulumi:"securityControls"`
	UserExperiences  CloudBrowserIsolationExternalProfileUserExperienceArrayOutput  `pulumi:"userExperiences"`
}

// NewCloudBrowserIsolationExternalProfile registers a new resource with the given unique name, arguments, and options.
func NewCloudBrowserIsolationExternalProfile(ctx *pulumi.Context,
	name string, args *CloudBrowserIsolationExternalProfileArgs, opts ...pulumi.ResourceOption) (*CloudBrowserIsolationExternalProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BannerId == nil {
		return nil, errors.New("invalid value for required argument 'BannerId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CloudBrowserIsolationExternalProfile
	err := ctx.RegisterResource("zpa:index/cloudBrowserIsolationExternalProfile:CloudBrowserIsolationExternalProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudBrowserIsolationExternalProfile gets an existing CloudBrowserIsolationExternalProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudBrowserIsolationExternalProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudBrowserIsolationExternalProfileState, opts ...pulumi.ResourceOption) (*CloudBrowserIsolationExternalProfile, error) {
	var resource CloudBrowserIsolationExternalProfile
	err := ctx.ReadResource("zpa:index/cloudBrowserIsolationExternalProfile:CloudBrowserIsolationExternalProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudBrowserIsolationExternalProfile resources.
type cloudBrowserIsolationExternalProfileState struct {
	BannerId *string `pulumi:"bannerId"`
	// This field defines the list of server groups IDs.
	CertificateIds []string `pulumi:"certificateIds"`
	Description    *string  `pulumi:"description"`
	Name           *string  `pulumi:"name"`
	// This field defines the list of server groups IDs.
	RegionIds        []string                                              `pulumi:"regionIds"`
	SecurityControls []CloudBrowserIsolationExternalProfileSecurityControl `pulumi:"securityControls"`
	UserExperiences  []CloudBrowserIsolationExternalProfileUserExperience  `pulumi:"userExperiences"`
}

type CloudBrowserIsolationExternalProfileState struct {
	BannerId pulumi.StringPtrInput
	// This field defines the list of server groups IDs.
	CertificateIds pulumi.StringArrayInput
	Description    pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	// This field defines the list of server groups IDs.
	RegionIds        pulumi.StringArrayInput
	SecurityControls CloudBrowserIsolationExternalProfileSecurityControlArrayInput
	UserExperiences  CloudBrowserIsolationExternalProfileUserExperienceArrayInput
}

func (CloudBrowserIsolationExternalProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudBrowserIsolationExternalProfileState)(nil)).Elem()
}

type cloudBrowserIsolationExternalProfileArgs struct {
	BannerId string `pulumi:"bannerId"`
	// This field defines the list of server groups IDs.
	CertificateIds []string `pulumi:"certificateIds"`
	Description    *string  `pulumi:"description"`
	Name           *string  `pulumi:"name"`
	// This field defines the list of server groups IDs.
	RegionIds        []string                                              `pulumi:"regionIds"`
	SecurityControls []CloudBrowserIsolationExternalProfileSecurityControl `pulumi:"securityControls"`
	UserExperiences  []CloudBrowserIsolationExternalProfileUserExperience  `pulumi:"userExperiences"`
}

// The set of arguments for constructing a CloudBrowserIsolationExternalProfile resource.
type CloudBrowserIsolationExternalProfileArgs struct {
	BannerId pulumi.StringInput
	// This field defines the list of server groups IDs.
	CertificateIds pulumi.StringArrayInput
	Description    pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	// This field defines the list of server groups IDs.
	RegionIds        pulumi.StringArrayInput
	SecurityControls CloudBrowserIsolationExternalProfileSecurityControlArrayInput
	UserExperiences  CloudBrowserIsolationExternalProfileUserExperienceArrayInput
}

func (CloudBrowserIsolationExternalProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudBrowserIsolationExternalProfileArgs)(nil)).Elem()
}

type CloudBrowserIsolationExternalProfileInput interface {
	pulumi.Input

	ToCloudBrowserIsolationExternalProfileOutput() CloudBrowserIsolationExternalProfileOutput
	ToCloudBrowserIsolationExternalProfileOutputWithContext(ctx context.Context) CloudBrowserIsolationExternalProfileOutput
}

func (*CloudBrowserIsolationExternalProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudBrowserIsolationExternalProfile)(nil)).Elem()
}

func (i *CloudBrowserIsolationExternalProfile) ToCloudBrowserIsolationExternalProfileOutput() CloudBrowserIsolationExternalProfileOutput {
	return i.ToCloudBrowserIsolationExternalProfileOutputWithContext(context.Background())
}

func (i *CloudBrowserIsolationExternalProfile) ToCloudBrowserIsolationExternalProfileOutputWithContext(ctx context.Context) CloudBrowserIsolationExternalProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudBrowserIsolationExternalProfileOutput)
}

// CloudBrowserIsolationExternalProfileArrayInput is an input type that accepts CloudBrowserIsolationExternalProfileArray and CloudBrowserIsolationExternalProfileArrayOutput values.
// You can construct a concrete instance of `CloudBrowserIsolationExternalProfileArrayInput` via:
//
//	CloudBrowserIsolationExternalProfileArray{ CloudBrowserIsolationExternalProfileArgs{...} }
type CloudBrowserIsolationExternalProfileArrayInput interface {
	pulumi.Input

	ToCloudBrowserIsolationExternalProfileArrayOutput() CloudBrowserIsolationExternalProfileArrayOutput
	ToCloudBrowserIsolationExternalProfileArrayOutputWithContext(context.Context) CloudBrowserIsolationExternalProfileArrayOutput
}

type CloudBrowserIsolationExternalProfileArray []CloudBrowserIsolationExternalProfileInput

func (CloudBrowserIsolationExternalProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudBrowserIsolationExternalProfile)(nil)).Elem()
}

func (i CloudBrowserIsolationExternalProfileArray) ToCloudBrowserIsolationExternalProfileArrayOutput() CloudBrowserIsolationExternalProfileArrayOutput {
	return i.ToCloudBrowserIsolationExternalProfileArrayOutputWithContext(context.Background())
}

func (i CloudBrowserIsolationExternalProfileArray) ToCloudBrowserIsolationExternalProfileArrayOutputWithContext(ctx context.Context) CloudBrowserIsolationExternalProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudBrowserIsolationExternalProfileArrayOutput)
}

// CloudBrowserIsolationExternalProfileMapInput is an input type that accepts CloudBrowserIsolationExternalProfileMap and CloudBrowserIsolationExternalProfileMapOutput values.
// You can construct a concrete instance of `CloudBrowserIsolationExternalProfileMapInput` via:
//
//	CloudBrowserIsolationExternalProfileMap{ "key": CloudBrowserIsolationExternalProfileArgs{...} }
type CloudBrowserIsolationExternalProfileMapInput interface {
	pulumi.Input

	ToCloudBrowserIsolationExternalProfileMapOutput() CloudBrowserIsolationExternalProfileMapOutput
	ToCloudBrowserIsolationExternalProfileMapOutputWithContext(context.Context) CloudBrowserIsolationExternalProfileMapOutput
}

type CloudBrowserIsolationExternalProfileMap map[string]CloudBrowserIsolationExternalProfileInput

func (CloudBrowserIsolationExternalProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudBrowserIsolationExternalProfile)(nil)).Elem()
}

func (i CloudBrowserIsolationExternalProfileMap) ToCloudBrowserIsolationExternalProfileMapOutput() CloudBrowserIsolationExternalProfileMapOutput {
	return i.ToCloudBrowserIsolationExternalProfileMapOutputWithContext(context.Background())
}

func (i CloudBrowserIsolationExternalProfileMap) ToCloudBrowserIsolationExternalProfileMapOutputWithContext(ctx context.Context) CloudBrowserIsolationExternalProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudBrowserIsolationExternalProfileMapOutput)
}

type CloudBrowserIsolationExternalProfileOutput struct{ *pulumi.OutputState }

func (CloudBrowserIsolationExternalProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudBrowserIsolationExternalProfile)(nil)).Elem()
}

func (o CloudBrowserIsolationExternalProfileOutput) ToCloudBrowserIsolationExternalProfileOutput() CloudBrowserIsolationExternalProfileOutput {
	return o
}

func (o CloudBrowserIsolationExternalProfileOutput) ToCloudBrowserIsolationExternalProfileOutputWithContext(ctx context.Context) CloudBrowserIsolationExternalProfileOutput {
	return o
}

func (o CloudBrowserIsolationExternalProfileOutput) BannerId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudBrowserIsolationExternalProfile) pulumi.StringOutput { return v.BannerId }).(pulumi.StringOutput)
}

// This field defines the list of server groups IDs.
func (o CloudBrowserIsolationExternalProfileOutput) CertificateIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudBrowserIsolationExternalProfile) pulumi.StringArrayOutput { return v.CertificateIds }).(pulumi.StringArrayOutput)
}

func (o CloudBrowserIsolationExternalProfileOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudBrowserIsolationExternalProfile) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CloudBrowserIsolationExternalProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudBrowserIsolationExternalProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// This field defines the list of server groups IDs.
func (o CloudBrowserIsolationExternalProfileOutput) RegionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudBrowserIsolationExternalProfile) pulumi.StringArrayOutput { return v.RegionIds }).(pulumi.StringArrayOutput)
}

func (o CloudBrowserIsolationExternalProfileOutput) SecurityControls() CloudBrowserIsolationExternalProfileSecurityControlArrayOutput {
	return o.ApplyT(func(v *CloudBrowserIsolationExternalProfile) CloudBrowserIsolationExternalProfileSecurityControlArrayOutput {
		return v.SecurityControls
	}).(CloudBrowserIsolationExternalProfileSecurityControlArrayOutput)
}

func (o CloudBrowserIsolationExternalProfileOutput) UserExperiences() CloudBrowserIsolationExternalProfileUserExperienceArrayOutput {
	return o.ApplyT(func(v *CloudBrowserIsolationExternalProfile) CloudBrowserIsolationExternalProfileUserExperienceArrayOutput {
		return v.UserExperiences
	}).(CloudBrowserIsolationExternalProfileUserExperienceArrayOutput)
}

type CloudBrowserIsolationExternalProfileArrayOutput struct{ *pulumi.OutputState }

func (CloudBrowserIsolationExternalProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudBrowserIsolationExternalProfile)(nil)).Elem()
}

func (o CloudBrowserIsolationExternalProfileArrayOutput) ToCloudBrowserIsolationExternalProfileArrayOutput() CloudBrowserIsolationExternalProfileArrayOutput {
	return o
}

func (o CloudBrowserIsolationExternalProfileArrayOutput) ToCloudBrowserIsolationExternalProfileArrayOutputWithContext(ctx context.Context) CloudBrowserIsolationExternalProfileArrayOutput {
	return o
}

func (o CloudBrowserIsolationExternalProfileArrayOutput) Index(i pulumi.IntInput) CloudBrowserIsolationExternalProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudBrowserIsolationExternalProfile {
		return vs[0].([]*CloudBrowserIsolationExternalProfile)[vs[1].(int)]
	}).(CloudBrowserIsolationExternalProfileOutput)
}

type CloudBrowserIsolationExternalProfileMapOutput struct{ *pulumi.OutputState }

func (CloudBrowserIsolationExternalProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudBrowserIsolationExternalProfile)(nil)).Elem()
}

func (o CloudBrowserIsolationExternalProfileMapOutput) ToCloudBrowserIsolationExternalProfileMapOutput() CloudBrowserIsolationExternalProfileMapOutput {
	return o
}

func (o CloudBrowserIsolationExternalProfileMapOutput) ToCloudBrowserIsolationExternalProfileMapOutputWithContext(ctx context.Context) CloudBrowserIsolationExternalProfileMapOutput {
	return o
}

func (o CloudBrowserIsolationExternalProfileMapOutput) MapIndex(k pulumi.StringInput) CloudBrowserIsolationExternalProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudBrowserIsolationExternalProfile {
		return vs[0].(map[string]*CloudBrowserIsolationExternalProfile)[vs[1].(string)]
	}).(CloudBrowserIsolationExternalProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudBrowserIsolationExternalProfileInput)(nil)).Elem(), &CloudBrowserIsolationExternalProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudBrowserIsolationExternalProfileArrayInput)(nil)).Elem(), CloudBrowserIsolationExternalProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudBrowserIsolationExternalProfileMapInput)(nil)).Elem(), CloudBrowserIsolationExternalProfileMap{})
	pulumi.RegisterOutputType(CloudBrowserIsolationExternalProfileOutput{})
	pulumi.RegisterOutputType(CloudBrowserIsolationExternalProfileArrayOutput{})
	pulumi.RegisterOutputType(CloudBrowserIsolationExternalProfileMapOutput{})
}
