// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zpa

import (
	"fmt"
	"unicode"

	// embed is used to store bridge-metadata.json in the compiled binary
	_ "embed"
	"path/filepath"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tks "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/zscaler/pulumi-zpa/provider/pkg/version"
	"github.com/zscaler/terraform-provider-zpa/v3/zpa"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	zpaPkg = "zpa"
	// modules:
	zpaMod = "index"
)

// zpaMember manufactures a type token for the zpa package and the given module and type.
func zpaMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(zpaPkg + ":" + mod + ":" + mem)
}

// zpaType manufactures a type token for the zpa package and the given module and type.
func zpaType(mod string, typ string) tokens.Type {
	return tokens.Type(zpaMember(mod, typ))
}

// zpaDataSource manufactures a standard resource token given a module and resource name.
// It automatically uses the zpa package and names the file by simply lower casing the data
// source's first character.
func zpaDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return zpaMember(mod+"/"+fn, res)
}

// zpaResource manufactures a standard resource token given a module and resource name.
// It automatically uses the zpa package and names the file by simply lower casing the resource's
// first character.
func zpaResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return zpaType(mod+"/"+fn, res)
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

func boolRef(b bool) *bool {
	return &b
}

//go:embed cmd/pulumi-resource-zpa/bridge-metadata.json
var metadata []byte

func refProviderLicense(license tfbridge.TFProviderLicense) *tfbridge.TFProviderLicense {
	return &license
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(zpa.ZPAProvider())
	prov := tfbridge.ProviderInfo{
		P:                       p,
		Name:                    "zpa",
		Description:             "A Pulumi package for creating and managing Zscaler Private Access (ZPA) cloud resources.",
		Keywords:                []string{"pulumi", "zpa", "zscaler", "category/cloud"},
		TFProviderLicense:       refProviderLicense(tfbridge.MITLicenseType),
		License:                 "MIT",
		LogoURL:                 "https://raw.githubusercontent.com/zscaler/pulumi-zpa/master/assets/zscaler.png", // nolint[:lll]
		Homepage:                "https://www.zscaler.com",
		Repository:              "https://github.com/zscaler/pulumi-zpa",
		PluginDownloadURL:       "github://api.github.com/zscaler",
		GitHubOrg:               "zscaler",
		Publisher:               "Zscaler",
		DisplayName:             "Zscaler Private Access",
		TFProviderModuleVersion: "v3",
		Version:                 version.Version,
		Config: map[string]*tfbridge.SchemaInfo{
			"zpa_client_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZPA_CLIENT_ID"},
				},
			},
			"zpa_client_secret": {
				Secret: boolRef(true),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZPA_CLIENT_SECRET"},
				},
			},
			"zpa_customer_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZPA_CUSTOMER_ID"},
				},
			},
			"zpa_cloud": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZPA_CLOUD"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"zpa_app_connector_assistant_schedule":         {Tok: zpaResource(zpaMod, "AssistantSchedule")},
			"zpa_app_connector_group":                      {Tok: zpaResource(zpaMod, "ConnectorGroup")},
			"zpa_service_edge_group":                       {Tok: zpaResource(zpaMod, "ServiceEdgeGroup")},
			"zpa_segment_group":                            {Tok: zpaResource(zpaMod, "SegmentGroup")},
			"zpa_server_group":                             {Tok: zpaResource(zpaMod, "ServerGroup")},
			"zpa_application_segment":                      {Tok: zpaResource(zpaMod, "ApplicationSegment")},
			"zpa_browser_access":                           {Tok: zpaResource(zpaMod, "BrowserAccess")},
			"zpa_application_segment_browser_access":       {Tok: zpaResource(zpaMod, "ApplicationSegmentBrowserAccess")},
			"zpa_application_segment_inspection":           {Tok: zpaResource(zpaMod, "ApplicationSegmentInspection")},
			"zpa_application_segment_pra":                  {Tok: zpaResource(zpaMod, "ApplicationSegmentPRA")},
			"zpa_application_server":                       {Tok: zpaResource(zpaMod, "ApplicationServer")},
			"zpa_ba_certificate":                           {Tok: zpaResource(zpaMod, "BrowserCertificate")},
			"zpa_cloud_browser_isolation_banner":           {Tok: zpaResource(zpaMod, "CloudBrowserIsolationBanner")},
			"zpa_cloud_browser_isolation_certificate":      {Tok: zpaResource(zpaMod, "CloudBrowserIsolationCertificate")},
			"zpa_cloud_browser_isolation_external_profile": {Tok: zpaResource(zpaMod, "CloudBrowserIsolationExternalProfile")},
			"zpa_inspection_custom_controls":               {Tok: zpaResource(zpaMod, "InspectionCustomControls")},
			"zpa_inspection_profile":                       {Tok: zpaResource(zpaMod, "InspectionProfile")},
			"zpa_lss_config_controller":                    {Tok: zpaResource(zpaMod, "LSSConfigController")},
			"zpa_microtenant_controller":                   {Tok: zpaResource(zpaMod, "Microtenant")},
			"zpa_policy_access_rule":                       {Tok: zpaResource(zpaMod, "PolicyAccessRule")},
			"zpa_policy_access_rule_v2":                    {Tok: zpaResource(zpaMod, "PolicyAccessRuleV2")},
			"zpa_policy_forwarding_rule":                   {Tok: zpaResource(zpaMod, "PolicyAccessForwardingRule")},
			"zpa_policy_forwarding_rule_v2":                {Tok: zpaResource(zpaMod, "PolicyAccessForwardingRuleV2")},
			"zpa_policy_timeout_rule":                      {Tok: zpaResource(zpaMod, "PolicyAccessTimeOutRule")},
			"zpa_policy_timeout_rule_v2":                   {Tok: zpaResource(zpaMod, "PolicyAccessTimeOutRuleV2")},
			"zpa_policy_inspection_rule":                   {Tok: zpaResource(zpaMod, "PolicyAccessInspectionRule")},
			"zpa_policy_inspection_rule_v2":                {Tok: zpaResource(zpaMod, "PolicyAccessInspectionRuleV2")},
			"zpa_policy_isolation_rule":                    {Tok: zpaResource(zpaMod, "PolicyAccessIsolationRule")},
			"zpa_policy_isolation_rule_v2":                 {Tok: zpaResource(zpaMod, "PolicyAccessIsolationRuleV2")},
			"zpa_policy_access_rule_reorder":               {Tok: zpaResource(zpaMod, "PolicyAccessReorderRule")},
			"zpa_pra_approval_controller":                  {Tok: zpaResource(zpaMod, "PRAApproval")},
			"zpa_pra_console_controller":                   {Tok: zpaResource(zpaMod, "PRAConsole")},
			"zpa_pra_credential_controller":                {Tok: zpaResource(zpaMod, "PRACredential")},
			"zpa_pra_portal_controller":                    {Tok: zpaResource(zpaMod, "PRAPortal")},
			"zpa_provisioning_key": {Tok: zpaResource(zpaMod, "ProvisioningKey"),
				Fields: map[string]*tfbridge.SchemaInfo{
					// Rename field to prevent this error in the DotNet SDK generation:
					// sdk/dotnet/ProvisioningKey.cs(69,31): error CS0542: 'ProvisioningKey': member names cannot be the same as their enclosing type [sdk/dotnet/zscaler.Zpa.csproj]
					"provisioning_key": {
						Name: "ProvisioningKeyValue",
					},
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"zpa_app_connector_group": {
				Tok: zpaDataSource(zpaMod, "getAppConnectorGroup"),
			},
			"zpa_app_connector_controller": {
				Tok: zpaDataSource(zpaMod, "getAppConnectorController"),
			},
			"zpa_app_connector_assistant_schedule": {
				Tok: zpaDataSource(zpaMod, "getAssistantSchedule"),
			},
			"zpa_cloud_browser_isolation_certificate": {
				Tok: zpaDataSource(zpaMod, "getCloudBrowserIsolationCertificate"),
			},
			"zpa_cloud_browser_isolation_banner": {
				Tok: zpaDataSource(zpaMod, "getCloudBrowserIsolationBanner"),
			},
			"zpa_cloud_browser_isolation_external_profile": {
				Tok: zpaDataSource(zpaMod, "getCloudBrowserIsolationExternalProfile"),
			},
			"zpa_cloud_browser_isolation_region": {
				Tok: zpaDataSource(zpaMod, "getCloudBrowserIsolationRegion"),
			},
			"zpa_cloud_browser_isolation_zpa_profile": {
				Tok: zpaDataSource(zpaMod, "getCloudBrowserIsolationZPAProfile"),
			},
			"zpa_service_edge_group": {
				Tok: zpaDataSource(zpaMod, "getServiceEdgeGroup"),
			},
			"zpa_service_edge_controller": {
				Tok: zpaDataSource(zpaMod, "getServiceEdgeController"),
			},
			"zpa_segment_group": {
				Tok: zpaDataSource(zpaMod, "getSegmentGroup"),
			},
			"zpa_server_group": {
				Tok: zpaDataSource(zpaMod, "getServerGroup"),
			},
			"zpa_application_segment": {
				Tok: zpaDataSource(zpaMod, "getApplicationSegment"),
			},
			"zpa_application_segment_browser_access": {
				Tok: zpaDataSource(zpaMod, "getApplicationSegmentBrowserAccess"),
			},
			"zpa_application_segment_inspection": {
				Tok: zpaDataSource(zpaMod, "getApplicationSegmentInspection"),
			},
			"zpa_inspection_custom_controls": {
				Tok: zpaDataSource(zpaMod, "getInspectionCustomControls"),
			},
			"zpa_inspection_profile": {
				Tok: zpaDataSource(zpaMod, "getInspectionProfile"),
			},
			"zpa_inspection_predefined_controls": {
				Tok: zpaDataSource(zpaMod, "getInspectionPredefinedControls"),
			},
			"zpa_inspection_all_predefined_controls": {
				Tok: zpaDataSource(zpaMod, "getInspectionAllPredefinedControls"),
			},
			"zpa_application_segment_pra": {
				Tok: zpaDataSource(zpaMod, "getApplicationSegmentPRA"),
			},
			"zpa_application_server": {
				Tok: zpaDataSource(zpaMod, "getApplicationServer"),
			},
			"zpa_lss_config_controller": {
				Tok: zpaDataSource(zpaMod, "getLSSConfigController"),
			},
			"zpa_lss_config_client_types": {
				Tok: zpaDataSource(zpaMod, "getLSSClientTypes"),
			},
			"zpa_lss_config_log_type_formats": {
				Tok: zpaDataSource(zpaMod, "getLSSLogTypeFormats"),
			},
			"zpa_lss_config_status_codes": {
				Tok: zpaDataSource(zpaMod, "getLSSStatusCodes"),
			},
			"zpa_microtenant_controller": {
				Tok: zpaDataSource(zpaMod, "getMicrotenant"),
			},
			"zpa_provisioning_key": {
				Tok: zpaDataSource(zpaMod, "getProvisioningKey"),
			},
			"zpa_ba_certificate": {
				Tok: zpaDataSource(zpaMod, "getBaCertificate"),
			},
			"zpa_cloud_connector_group": {
				Tok: zpaDataSource(zpaMod, "getCloudConnectorGroup"),
			},
			"zpa_enrollment_cert": {
				Tok: zpaDataSource(zpaMod, "getEnrollmentCert"),
			},
			"zpa_idp_controller": {
				Tok: zpaDataSource(zpaMod, "getIdPController"),
			},
			"zpa_policy_type": {
				Tok: zpaDataSource(zpaMod, "getPolicyType"),
			},
			"zpa_access_policy_client_types": {
				Tok: zpaDataSource(zpaMod, "getPolicyClientType"),
			},
			"zpa_access_policy_platforms": {
				Tok: zpaDataSource(zpaMod, "getPolicyPlatform"),
			},
			"zpa_customer_version_profile": {
				Tok: zpaDataSource(zpaMod, "getCustomerVersionProfile"),
			},
			"zpa_isolation_profile": {
				Tok: zpaDataSource(zpaMod, "getIsolationProfile"),
			},
			"zpa_pra_approval_controller": {
				Tok: zpaDataSource(zpaMod, "getPRAApproval"),
			},
			"zpa_pra_console_controller": {
				Tok: zpaDataSource(zpaMod, "getPRAConsole"),
			},
			"zpa_pra_credential_controller": {
				Tok: zpaDataSource(zpaMod, "getPRACredential"),
			},
			"zpa_pra_portal_controller": {
				Tok: zpaDataSource(zpaMod, "getPRAPortal"),
			},
			"zpa_saml_attribute": {
				Tok: zpaDataSource(zpaMod, "getSAMLAttribute"),
			},
			"zpa_scim_attribute_header": {
				Tok: zpaDataSource(zpaMod, "getSCIMAttributeHeader"),
			},
			"zpa_scim_groups": {
				Tok: zpaDataSource(zpaMod, "getSCIMGroups"),
			},
			"zpa_machine_group": {
				Tok: zpaDataSource(zpaMod, "getMachineGroup"),
			},
			"zpa_posture_profile": {
				Tok: zpaDataSource(zpaMod, "getPostureProfile"),
			},
			"zpa_trusted_network": {
				Tok: zpaDataSource(zpaMod, "getTrustedNetwork"),
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@bdzscaler/pulumi-zpa",

			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "zscaler_pulumi_zpa",

			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/zscaler/pulumi-%[1]s/sdk/", zpaPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				zpaPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Zscaler",

			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
	}

	prov.MustComputeTokens(tks.SingleModule("zpa_", zpaMod, tks.MakeStandard(zpaPkg)))

	prov.SetAutonaming(255, "-")

	prov.MustApplyAutoAliases()

	return prov
}
