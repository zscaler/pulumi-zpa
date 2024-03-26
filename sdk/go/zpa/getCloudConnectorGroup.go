// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// Use the **zpa_cloud_connector_group** data source to get information about a cloud connector group created from the Zscaler Private Access cloud. This data source can then be referenced within an Access Policy rule
//
// > **NOTE:** A Cloud Connector Group resource is created in the Zscaler Cloud Connector cloud and replicated to the ZPA cloud. This resource can then be referenced in a Access Policy Rule where the Object Type = `CLOUD_CONNECTOR_GROUP` is being used.
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
//			_, err := zpa.GetCloudConnectorGroup(ctx, &zpa.GetCloudConnectorGroupArgs{
//				Name: pulumi.StringRef("AWS-Cloud"),
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
//			_, err := zpa.GetCloudConnectorGroup(ctx, &zpa.GetCloudConnectorGroupArgs{
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
func GetCloudConnectorGroup(ctx *pulumi.Context, args *GetCloudConnectorGroupArgs, opts ...pulumi.InvokeOption) (*GetCloudConnectorGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCloudConnectorGroupResult
	err := ctx.Invoke("zpa:index/getCloudConnectorGroup:getCloudConnectorGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudConnectorGroup.
type GetCloudConnectorGroupArgs struct {
	// This field defines the id of the cloud connector group.
	Id *string `pulumi:"id"`
	// This field defines the name of the cloud connector group.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getCloudConnectorGroup.
type GetCloudConnectorGroupResult struct {
	// (string) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	CloudConnectors []GetCloudConnectorGroupCloudConnector `pulumi:"cloudConnectors"`
	// (string) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	CreationTime string `pulumi:"creationTime"`
	// (string) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	Description string `pulumi:"description"`
	// (bool) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	Enabled bool `pulumi:"enabled"`
	// (string) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	GeolocationId string  `pulumi:"geolocationId"`
	Id            *string `pulumi:"id"`
	// (string)- Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	ModifiedTime string `pulumi:"modifiedTime"`
	Modifiedby   string `pulumi:"modifiedby"`
	// (string) - This field defines the name of the cloud connector group.
	Name *string `pulumi:"name"`
	// (string) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	ZiaCloud string `pulumi:"ziaCloud"`
	// (string) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	ZiaOrgId string `pulumi:"ziaOrgId"`
}

func GetCloudConnectorGroupOutput(ctx *pulumi.Context, args GetCloudConnectorGroupOutputArgs, opts ...pulumi.InvokeOption) GetCloudConnectorGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCloudConnectorGroupResult, error) {
			args := v.(GetCloudConnectorGroupArgs)
			r, err := GetCloudConnectorGroup(ctx, &args, opts...)
			var s GetCloudConnectorGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCloudConnectorGroupResultOutput)
}

// A collection of arguments for invoking getCloudConnectorGroup.
type GetCloudConnectorGroupOutputArgs struct {
	// This field defines the id of the cloud connector group.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// This field defines the name of the cloud connector group.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetCloudConnectorGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudConnectorGroupArgs)(nil)).Elem()
}

// A collection of values returned by getCloudConnectorGroup.
type GetCloudConnectorGroupResultOutput struct{ *pulumi.OutputState }

func (GetCloudConnectorGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudConnectorGroupResult)(nil)).Elem()
}

func (o GetCloudConnectorGroupResultOutput) ToGetCloudConnectorGroupResultOutput() GetCloudConnectorGroupResultOutput {
	return o
}

func (o GetCloudConnectorGroupResultOutput) ToGetCloudConnectorGroupResultOutputWithContext(ctx context.Context) GetCloudConnectorGroupResultOutput {
	return o
}

// (string) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
func (o GetCloudConnectorGroupResultOutput) CloudConnectors() GetCloudConnectorGroupCloudConnectorArrayOutput {
	return o.ApplyT(func(v GetCloudConnectorGroupResult) []GetCloudConnectorGroupCloudConnector { return v.CloudConnectors }).(GetCloudConnectorGroupCloudConnectorArrayOutput)
}

// (string) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
func (o GetCloudConnectorGroupResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudConnectorGroupResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// (string) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
func (o GetCloudConnectorGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudConnectorGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// (bool) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
func (o GetCloudConnectorGroupResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetCloudConnectorGroupResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// (string) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
func (o GetCloudConnectorGroupResultOutput) GeolocationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudConnectorGroupResult) string { return v.GeolocationId }).(pulumi.StringOutput)
}

func (o GetCloudConnectorGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCloudConnectorGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// (string)- Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
func (o GetCloudConnectorGroupResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudConnectorGroupResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o GetCloudConnectorGroupResultOutput) Modifiedby() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudConnectorGroupResult) string { return v.Modifiedby }).(pulumi.StringOutput)
}

// (string) - This field defines the name of the cloud connector group.
func (o GetCloudConnectorGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCloudConnectorGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// (string) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
func (o GetCloudConnectorGroupResultOutput) ZiaCloud() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudConnectorGroupResult) string { return v.ZiaCloud }).(pulumi.StringOutput)
}

// (string) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
func (o GetCloudConnectorGroupResultOutput) ZiaOrgId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudConnectorGroupResult) string { return v.ZiaOrgId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCloudConnectorGroupResultOutput{})
}
