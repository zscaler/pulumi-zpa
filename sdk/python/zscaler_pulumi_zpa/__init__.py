# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .app_connector_assistant_schedule import *
from .application_segment import *
from .application_segment_browser_access import *
from .application_segment_inspection import *
from .application_segment_pra import *
from .application_server import *
from .assistant_schedule import *
from .browser_access import *
from .browser_certificate import *
from .cloud_browser_isolation_banner import *
from .cloud_browser_isolation_certificate import *
from .cloud_browser_isolation_external_profile import *
from .connector_group import *
from .emergency_access_user import *
from .get_app_connector_assistant_schedule import *
from .get_app_connector_controller import *
from .get_app_connector_group import *
from .get_application_segment import *
from .get_application_segment_browser_access import *
from .get_application_segment_by_type import *
from .get_application_segment_inspection import *
from .get_application_segment_pra import *
from .get_application_server import *
from .get_assistant_schedule import *
from .get_ba_certificate import *
from .get_cloud_browser_isolation_banner import *
from .get_cloud_browser_isolation_certificate import *
from .get_cloud_browser_isolation_external_profile import *
from .get_cloud_browser_isolation_region import *
from .get_cloud_browser_isolation_zpa_profile import *
from .get_cloud_connector_group import *
from .get_customer_version_profile import *
from .get_enrollment_cert import *
from .get_id_p_controller import *
from .get_inspection_all_predefined_controls import *
from .get_inspection_custom_controls import *
from .get_inspection_predefined_controls import *
from .get_inspection_profile import *
from .get_isolation_profile import *
from .get_lss_client_types import *
from .get_lss_config_controller import *
from .get_lss_log_type_formats import *
from .get_lss_status_codes import *
from .get_machine_group import *
from .get_microtenant import *
from .get_policy_client_type import *
from .get_policy_platform import *
from .get_policy_type import *
from .get_posture_profile import *
from .get_pra_approval import *
from .get_pra_approval_controller import *
from .get_pra_console import *
from .get_pra_console_controller import *
from .get_pra_credential import *
from .get_pra_credential_controller import *
from .get_pra_portal import *
from .get_pra_portal_controller import *
from .get_provisioning_key import *
from .get_saml_attribute import *
from .get_scim_attribute_header import *
from .get_scim_groups import *
from .get_segment_group import *
from .get_server_group import *
from .get_service_edge_assistant_schedule import *
from .get_service_edge_controller import *
from .get_service_edge_group import *
from .get_trusted_network import *
from .inspection_custom_controls import *
from .inspection_profile import *
from .lss_config_controller import *
from .microtenant import *
from .policy_access_credential_rule import *
from .policy_access_forwarding_rule import *
from .policy_access_forwarding_rule_v2 import *
from .policy_access_inspection_rule import *
from .policy_access_inspection_rule_v2 import *
from .policy_access_isolation_rule import *
from .policy_access_isolation_rule_v2 import *
from .policy_access_reorder_rule import *
from .policy_access_rule import *
from .policy_access_rule_v2 import *
from .policy_access_time_out_rule import *
from .policy_access_time_out_rule_v2 import *
from .policy_capabilities_rule import *
from .policy_credential_rule import *
from .policy_forwarding_rule_v2 import *
from .policy_inspection_rule_v2 import *
from .policy_isolation_rule_v2 import *
from .policy_redirection_rule import *
from .policy_timeout_rule_v2 import *
from .pra_approval import *
from .pra_approval_controller import *
from .pra_console import *
from .pra_console_controller import *
from .pra_credential import *
from .pra_credential_controller import *
from .pra_portal import *
from .pra_portal_controller import *
from .provider import *
from .provisioning_key import *
from .segment_group import *
from .server_group import *
from .service_edge_assistant_schedule import *
from .service_edge_group import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import zscaler_pulumi_zpa.config as __config
    config = __config
else:
    config = _utilities.lazy_import('zscaler_pulumi_zpa.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "zpa",
  "mod": "index/appConnectorAssistantSchedule",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/appConnectorAssistantSchedule:AppConnectorAssistantSchedule": "AppConnectorAssistantSchedule"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/applicationSegment",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/applicationSegment:ApplicationSegment": "ApplicationSegment"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/applicationSegmentBrowserAccess",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/applicationSegmentBrowserAccess:ApplicationSegmentBrowserAccess": "ApplicationSegmentBrowserAccess"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/applicationSegmentInspection",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/applicationSegmentInspection:ApplicationSegmentInspection": "ApplicationSegmentInspection"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/applicationSegmentPRA",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/applicationSegmentPRA:ApplicationSegmentPRA": "ApplicationSegmentPRA"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/applicationServer",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/applicationServer:ApplicationServer": "ApplicationServer"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/assistantSchedule",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/assistantSchedule:AssistantSchedule": "AssistantSchedule"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/browserAccess",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/browserAccess:BrowserAccess": "BrowserAccess"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/browserCertificate",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/browserCertificate:BrowserCertificate": "BrowserCertificate"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/cloudBrowserIsolationBanner",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/cloudBrowserIsolationBanner:CloudBrowserIsolationBanner": "CloudBrowserIsolationBanner"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/cloudBrowserIsolationCertificate",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/cloudBrowserIsolationCertificate:CloudBrowserIsolationCertificate": "CloudBrowserIsolationCertificate"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/cloudBrowserIsolationExternalProfile",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/cloudBrowserIsolationExternalProfile:CloudBrowserIsolationExternalProfile": "CloudBrowserIsolationExternalProfile"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/connectorGroup",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/connectorGroup:ConnectorGroup": "ConnectorGroup"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/emergencyAccessUser",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/emergencyAccessUser:EmergencyAccessUser": "EmergencyAccessUser"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/inspectionCustomControls",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/inspectionCustomControls:InspectionCustomControls": "InspectionCustomControls"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/inspectionProfile",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/inspectionProfile:InspectionProfile": "InspectionProfile"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/lSSConfigController",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/lSSConfigController:LSSConfigController": "LSSConfigController"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/microtenant",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/microtenant:Microtenant": "Microtenant"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/pRAApproval",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/pRAApproval:PRAApproval": "PRAApproval"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/pRAConsole",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/pRAConsole:PRAConsole": "PRAConsole"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/pRACredential",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/pRACredential:PRACredential": "PRACredential"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/pRAPortal",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/pRAPortal:PRAPortal": "PRAPortal"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/policyAccessCredentialRule",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/policyAccessCredentialRule:PolicyAccessCredentialRule": "PolicyAccessCredentialRule"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/policyAccessForwardingRule",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/policyAccessForwardingRule:PolicyAccessForwardingRule": "PolicyAccessForwardingRule"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/policyAccessForwardingRuleV2",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/policyAccessForwardingRuleV2:PolicyAccessForwardingRuleV2": "PolicyAccessForwardingRuleV2"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/policyAccessInspectionRule",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/policyAccessInspectionRule:PolicyAccessInspectionRule": "PolicyAccessInspectionRule"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/policyAccessInspectionRuleV2",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/policyAccessInspectionRuleV2:PolicyAccessInspectionRuleV2": "PolicyAccessInspectionRuleV2"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/policyAccessIsolationRule",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/policyAccessIsolationRule:PolicyAccessIsolationRule": "PolicyAccessIsolationRule"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/policyAccessIsolationRuleV2",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/policyAccessIsolationRuleV2:PolicyAccessIsolationRuleV2": "PolicyAccessIsolationRuleV2"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/policyAccessReorderRule",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/policyAccessReorderRule:PolicyAccessReorderRule": "PolicyAccessReorderRule"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/policyAccessRule",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/policyAccessRule:PolicyAccessRule": "PolicyAccessRule"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/policyAccessRuleV2",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/policyAccessRuleV2:PolicyAccessRuleV2": "PolicyAccessRuleV2"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/policyAccessTimeOutRule",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/policyAccessTimeOutRule:PolicyAccessTimeOutRule": "PolicyAccessTimeOutRule"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/policyAccessTimeOutRuleV2",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/policyAccessTimeOutRuleV2:PolicyAccessTimeOutRuleV2": "PolicyAccessTimeOutRuleV2"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/policyCapabilitiesRule",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/policyCapabilitiesRule:PolicyCapabilitiesRule": "PolicyCapabilitiesRule"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/policyCredentialRule",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/policyCredentialRule:PolicyCredentialRule": "PolicyCredentialRule"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/policyForwardingRuleV2",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/policyForwardingRuleV2:PolicyForwardingRuleV2": "PolicyForwardingRuleV2"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/policyInspectionRuleV2",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/policyInspectionRuleV2:PolicyInspectionRuleV2": "PolicyInspectionRuleV2"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/policyIsolationRuleV2",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/policyIsolationRuleV2:PolicyIsolationRuleV2": "PolicyIsolationRuleV2"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/policyRedirectionRule",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/policyRedirectionRule:PolicyRedirectionRule": "PolicyRedirectionRule"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/policyTimeoutRuleV2",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/policyTimeoutRuleV2:PolicyTimeoutRuleV2": "PolicyTimeoutRuleV2"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/praApprovalController",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/praApprovalController:PraApprovalController": "PraApprovalController"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/praConsoleController",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/praConsoleController:PraConsoleController": "PraConsoleController"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/praCredentialController",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/praCredentialController:PraCredentialController": "PraCredentialController"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/praPortalController",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/praPortalController:PraPortalController": "PraPortalController"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/provisioningKey",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/provisioningKey:ProvisioningKey": "ProvisioningKey"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/segmentGroup",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/segmentGroup:SegmentGroup": "SegmentGroup"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/serverGroup",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/serverGroup:ServerGroup": "ServerGroup"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/serviceEdgeAssistantSchedule",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/serviceEdgeAssistantSchedule:ServiceEdgeAssistantSchedule": "ServiceEdgeAssistantSchedule"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/serviceEdgeGroup",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/serviceEdgeGroup:ServiceEdgeGroup": "ServiceEdgeGroup"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "zpa",
  "token": "pulumi:providers:zpa",
  "fqn": "zscaler_pulumi_zpa",
  "class": "Provider"
 }
]
"""
)
