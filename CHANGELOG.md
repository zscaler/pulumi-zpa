# Changelog

## 0.0.10 (April 9, 2024)

### Notes

- Release date: **(April 9, 2024)**

* New datasource: `zpa_pra_approval_controller` retrieve Privileged Remote Access Approval [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432)
* New resource: `zpa_pra_approval_controller` manages Privileged Remote Access Approval [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432)
* New datasource: `zpa_pra_portal_controller` retrieve Privileged Remote Access Portal [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432)
* New resource: `zpa_pra_portal_controller` manages Privileged Remote Access Portal [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432)
* New datasource: `zpa_pra_credential_controller` retrieve Privileged Remote Access Credential [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432)
* New resource: `zpa_pra_credential_controller` manages Privileged Remote Access Credential [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432)
* New datasource: `zpa_pra_console_controller` retrieve Privileged Remote Access Console [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432)
* New resource: `zpa_pra_console_controller` manages Privileged Remote Access Console 
[PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432)
* New Resources: Introduced new Policy Access resources that are managed via a new `v2` API endpoint:
  - `zpa_policy_access_rule_v2` manages access policy rule via `v2` API endpoint [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432)
  - `zpa_policy_forwarding_rule_v2` manages access policy forwarding rule via `v2` API endpoint [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432)
  - `zpa_policy_isolation_rule_v2` manages access policy isolation rule via `v2` API endpoint [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432)
  - `zpa_policy_inspection_rule_v2` manages access policy inspection rule via `v2` API endpoint [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432)
  - `zpa_policy_timeout_rule_v2` manages access policy timeout rule via `v2` API endpoint [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432)
  - `zpa_policy_redirection_rule` manages redirection access policy via `v2` API endpoint [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432)
  - `zpa_policy_credential_rule` manages access policy credential rule via `v2` API endpoint [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432)
  - `zpa_policy_capabilities_rule` manages access policy capabilities rule via `v2` API endpoint [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432)
  
    ⚠️ **WARNING:**: Notice that any Access Policy `v2` is a new resource and uses a different HCL format structure. If you decide to migrate to the new v2 resources, notice that this is considered a breaking change and must be done carefully. This warning only applies for those with existing `v1` Access Policy HCL format structure.

[PR #434](https://github.com/zscaler/terraform-provider-zpa/pull/434)
* New resource: `zpa_emergency_access_user` manages Emergency Access Users 

### NEW PROPERTIES
* New Properties: The resource `zpa_ba_certificate` now displays the attributes `valid_from_in_epochsec` and `valid_to_in_epochsec` in human readable `RFC1123` format
* New Properties: The provider now includes support to `ZPATWO` cloud [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432)

### DEPRECATIONS
* Deprecated attribute: The attributes `policy_migrated` and `tcp_keep_alive_enabled` are now deprecated for the resource `zpa_segment_group`. For the attribute `tcp_keep_alive_enabled` use the attribute `tcp_keep_alive` within the resource  `zpa_application_segment`", [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432). 
* Deprecated attribute: The attributes `negated` within all access policy rule resource types. [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432). 
* Deprecated attribute: The attributes `rule_order` within all access policy rule resource types. Please use the newly dedicated resource `zpa_policy_access_rule_reorder` [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432). 

### ENHACEMENTS
* Attribute `policy_set_id` is now optional across all access policy rule resources `v1` and `v2`. The provider will automatically set the `policy_set_id` according to the policy access resource being configured. This improvement removes the need to explicitly use the data source `zpa_policy_type` [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432)
* Added new `match_style` attribute to the `zpa_application_segment` resource [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432). Issue [#424](https://github.com/zscaler/terraform-provider-zpa/issues/424). To learn more about this attribute visit [Zscaler Help Portal](https://help.zscaler.com/zpa/using-app-segment-multimatch)
* Update `zpa_ba_certificate` documentation [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432)
* Several ACC tests maintenance [PR #432](https://github.com/zscaler/terraform-provider-zpa/pull/432)


## 0.0.1 (March 26, 2024)

### Notes

- Release date: **(March 26, 2024)**

🎉 **Initial Release** 🎉
