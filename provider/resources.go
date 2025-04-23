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
	"github.com/zscaler/terraform-provider-zpa/v4/zpa"
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
		TFProviderModuleVersion: "v4",
		Version:                 version.Version,
		Config: map[string]*tfbridge.SchemaInfo{
			// Authentication Parameters for OneAPI Mode
			"client_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZSCALER_CLIENT_ID"},
				},
			},
			"client_secret": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZSCALER_CLIENT_SECRET"},
				},
			},
			"private_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZSCALER_PRIVATE_KEY"},
				},
			},
			"vanity_domain": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZSCALER_VANITY_DOMAIN"},
				},
			},
			"zscaler_cloud": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZSCALER_CLOUD"},
				},
			},
			// Authentication Parameters for Legacy API Model
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
			"use_legacy_client": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZSCALER_USE_LEGACY_CLIENT"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"zpa_app_connector_assistant_schedule": {Tok: zpaResource(zpaMod, "AppConnectorAssistantSchedule")},
			"zpa_app_connector_group":              {Tok: zpaResource(zpaMod, "ConnectorGroup")},
			"zpa_service_edge_group":               {Tok: zpaResource(zpaMod, "ServiceEdgeGroup")},
			"zpa_segment_group":                    {Tok: zpaResource(zpaMod, "SegmentGroup")},
			"zpa_server_group":                     {Tok: zpaResource(zpaMod, "ServerGroup")},
			"zpa_application_segment":              {Tok: zpaResource(zpaMod, "ApplicationSegment")},
			"zpa_browser_access": {
				Tok: zpaResource(zpaMod, "BrowserAccess"),
				// No upstream docs for this resource exist:
				Docs:               &tfbridge.DocInfo{AllowMissing: true},
				DeprecationMessage: "Resource is deprecated due to a correction in naming conventions",
			},
			"zpa_application_segment_browser_access":       {Tok: zpaResource(zpaMod, "ApplicationSegmentBrowserAccess")},
			"zpa_application_segment_inspection":           {Tok: zpaResource(zpaMod, "ApplicationSegmentInspection")},
			"zpa_application_segment_pra":                  {Tok: zpaResource(zpaMod, "ApplicationSegmentPRA")},
			"zpa_application_server":                       {Tok: zpaResource(zpaMod, "ApplicationServer")},
			"zpa_ba_certificate":                           {Tok: zpaResource(zpaMod, "BrowserCertificate")},
			"zpa_cloud_browser_isolation_banner":           {Tok: zpaResource(zpaMod, "CloudBrowserIsolationBanner")},
			"zpa_cloud_browser_isolation_certificate":      {Tok: zpaResource(zpaMod, "CloudBrowserIsolationCertificate")},
			"zpa_cloud_browser_isolation_external_profile": {Tok: zpaResource(zpaMod, "CloudBrowserIsolationExternalProfile")},
			"zpa_emergency_access_user":                    {Tok: zpaResource(zpaMod, "EmergencyAccessUser")},
			"zpa_inspection_custom_controls": {
				Tok: zpaResource(zpaMod, "InspectionCustomControls"),
				// No upstream docs for this resource exist:
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_inspection_profile": {Tok: zpaResource(zpaMod, "InspectionProfile")},
			"zpa_lss_config_controller": {
				Tok: zpaResource(zpaMod, "LSSConfigController"),
				// No upstream docs for this resource exist:
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_microtenant_controller":    {Tok: zpaResource(zpaMod, "Microtenant")},
			"zpa_policy_access_rule":        {Tok: zpaResource(zpaMod, "PolicyAccessRule")},
			"zpa_policy_access_rule_v2":     {Tok: zpaResource(zpaMod, "PolicyAccessRuleV2")},
			"zpa_policy_forwarding_rule":    {Tok: zpaResource(zpaMod, "PolicyAccessForwardingRule")},
			"zpa_policy_forwarding_rule_v2": {Tok: zpaResource(zpaMod, "PolicyAccessForwardingRuleV2")},
			"zpa_policy_timeout_rule":       {Tok: zpaResource(zpaMod, "PolicyAccessTimeOutRule")},
			"zpa_policy_timeout_rule_v2":    {Tok: zpaResource(zpaMod, "PolicyAccessTimeOutRuleV2")},
			"zpa_policy_inspection_rule":    {Tok: zpaResource(zpaMod, "PolicyAccessInspectionRule")},
			"zpa_policy_inspection_rule_v2": {Tok: zpaResource(zpaMod, "PolicyAccessInspectionRuleV2")},
			"zpa_policy_isolation_rule":     {Tok: zpaResource(zpaMod, "PolicyAccessIsolationRule")},
			"zpa_policy_isolation_rule_v2":  {Tok: zpaResource(zpaMod, "PolicyAccessIsolationRuleV2")},
			"zpa_policy_credential_rule": {
				Tok: zpaResource(zpaMod, "PolicyAccessCredentialRule"),
				// No upstream docs for this resource exist:
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_policy_redirection_rule":         {Tok: zpaResource(zpaMod, "PolicyAccessRedirectionRule")},
			"zpa_policy_access_rule_reorder":      {Tok: zpaResource(zpaMod, "PolicyAccessReorderRule")},
			"zpa_pra_approval_controller":         {Tok: zpaResource(zpaMod, "PRAApproval")},
			"zpa_pra_console_controller":          {Tok: zpaResource(zpaMod, "PRAConsole")},
			"zpa_pra_credential_controller":       {Tok: zpaResource(zpaMod, "PRACredential")},
			"zpa_pra_credential_pool":             {Tok: zpaResource(zpaMod, "PRACredentialPool")},
			"zpa_pra_portal_controller":           {Tok: zpaResource(zpaMod, "PRAPortal")},
			"zpa_service_edge_assistant_schedule": {Tok: zpaResource(zpaMod, "ServiceEdgeAssistantSchedule")},
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
				Tok:  zpaDataSource(zpaMod, "getAppConnectorGroup"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_app_connector_controller": {
				Tok:  zpaDataSource(zpaMod, "getAppConnectorController"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_app_connector_assistant_schedule": {
				Tok:  zpaDataSource(zpaMod, "getAppConnectorAssistantSchedule"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_cloud_browser_isolation_certificate": {
				Tok:  zpaDataSource(zpaMod, "getCloudBrowserIsolationCertificate"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_cloud_browser_isolation_banner": {
				Tok:  zpaDataSource(zpaMod, "getCloudBrowserIsolationBanner"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_cloud_browser_isolation_external_profile": {
				Tok:  zpaDataSource(zpaMod, "getCloudBrowserIsolationExternalProfile"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_cloud_browser_isolation_region": {
				Tok:  zpaDataSource(zpaMod, "getCloudBrowserIsolationRegion"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_cloud_browser_isolation_zpa_profile": {
				Tok:  zpaDataSource(zpaMod, "getCloudBrowserIsolationZPAProfile"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_service_edge_group": {
				Tok:  zpaDataSource(zpaMod, "getServiceEdgeGroup"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_service_edge_controller": {
				Tok:  zpaDataSource(zpaMod, "getServiceEdgeController"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_segment_group": {
				Tok:  zpaDataSource(zpaMod, "getSegmentGroup"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_server_group": {
				Tok:  zpaDataSource(zpaMod, "getServerGroup"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_application_segment": {
				Tok:  zpaDataSource(zpaMod, "getApplicationSegment"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_application_segment_browser_access": {
				Tok:  zpaDataSource(zpaMod, "getApplicationSegmentBrowserAccess"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_application_segment_inspection": {
				Tok:  zpaDataSource(zpaMod, "getApplicationSegmentInspection"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_inspection_custom_controls": {
				Tok:  zpaDataSource(zpaMod, "getInspectionCustomControls"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_inspection_profile": {
				Tok:  zpaDataSource(zpaMod, "getInspectionProfile"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_inspection_predefined_controls": {
				Tok:  zpaDataSource(zpaMod, "getInspectionPredefinedControls"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_inspection_all_predefined_controls": {
				Tok:  zpaDataSource(zpaMod, "getInspectionAllPredefinedControls"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_application_segment_pra": {
				Tok:  zpaDataSource(zpaMod, "getApplicationSegmentPRA"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_application_server": {
				Tok:  zpaDataSource(zpaMod, "getApplicationServer"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_lss_config_controller": {
				Tok:  zpaDataSource(zpaMod, "getLSSConfigController"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_lss_config_client_types": {
				Tok:  zpaDataSource(zpaMod, "getLSSClientTypes"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_lss_config_log_type_formats": {
				Tok:  zpaDataSource(zpaMod, "getLSSLogTypeFormats"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_lss_config_status_codes": {
				Tok:  zpaDataSource(zpaMod, "getLSSStatusCodes"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_microtenant_controller": {
				Tok:  zpaDataSource(zpaMod, "getMicrotenant"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_provisioning_key": {
				Tok:  zpaDataSource(zpaMod, "getProvisioningKey"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_ba_certificate": {
				Tok:  zpaDataSource(zpaMod, "getBaCertificate"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_cloud_connector_group": {
				Tok:  zpaDataSource(zpaMod, "getCloudConnectorGroup"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_enrollment_cert": {
				Tok:  zpaDataSource(zpaMod, "getEnrollmentCert"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_idp_controller": {
				Tok:  zpaDataSource(zpaMod, "getIdPController"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_policy_type": {
				Tok:  zpaDataSource(zpaMod, "getPolicyType"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_access_policy_client_types": {
				Tok:  zpaDataSource(zpaMod, "getPolicyClientType"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_access_policy_platforms": {
				Tok:  zpaDataSource(zpaMod, "getPolicyPlatform"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_customer_version_profile": {
				Tok:  zpaDataSource(zpaMod, "getCustomerVersionProfile"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_isolation_profile": {
				Tok:  zpaDataSource(zpaMod, "getIsolationProfile"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_pra_approval_controller": {
				Tok:  zpaDataSource(zpaMod, "getPRAApproval"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_pra_console_controller": {
				Tok:  zpaDataSource(zpaMod, "getPRAConsole"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_pra_credential_controller": {
				Tok:  zpaDataSource(zpaMod, "getPRACredential"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_pra_credential_pool": {
				Tok:  zpaDataSource(zpaMod, "getPRACredentialPool"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_pra_portal_controller": {
				Tok:  zpaDataSource(zpaMod, "getPRAPortal"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_saml_attribute": {
				Tok:  zpaDataSource(zpaMod, "getSAMLAttribute"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_scim_attribute_header": {
				Tok:  zpaDataSource(zpaMod, "getSCIMAttributeHeader"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_scim_groups": {
				Tok:  zpaDataSource(zpaMod, "getSCIMGroups"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_machine_group": {
				Tok:  zpaDataSource(zpaMod, "getMachineGroup"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_posture_profile": {
				Tok:  zpaDataSource(zpaMod, "getPostureProfile"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_trusted_network": {
				Tok:  zpaDataSource(zpaMod, "getTrustedNetwork"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zpa_service_edge_assistant_schedule": {
				Tok:  zpaDataSource(zpaMod, "getServiceEdgeAssistantSchedule"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
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
