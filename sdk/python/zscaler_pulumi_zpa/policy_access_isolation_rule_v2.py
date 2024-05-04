# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['PolicyAccessIsolationRuleV2Args', 'PolicyAccessIsolationRuleV2']

@pulumi.input_type
class PolicyAccessIsolationRuleV2Args:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessIsolationRuleV2ConditionArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 microtenant_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 zpn_isolation_profile_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PolicyAccessIsolationRuleV2 resource.
        :param pulumi.Input[str] action: This is for providing the rule action. Supported values: `ISOLATE` Default.
        :param pulumi.Input[Sequence[pulumi.Input['PolicyAccessIsolationRuleV2ConditionArgs']]] conditions: This is for proviidng the set of conditions for the policy.
        :param pulumi.Input[str] description: This is the description of the access policy rule.
        :param pulumi.Input[str] name: This is the name of the policy rule.
        :param pulumi.Input[str] zpn_isolation_profile_id: Use zpa*isolation*profile data source to retrieve the necessary Isolation profile ID `zpn_isolation_profile_id`
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if microtenant_id is not None:
            pulumi.set(__self__, "microtenant_id", microtenant_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if zpn_isolation_profile_id is not None:
            pulumi.set(__self__, "zpn_isolation_profile_id", zpn_isolation_profile_id)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        This is for providing the rule action. Supported values: `ISOLATE` Default.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessIsolationRuleV2ConditionArgs']]]]:
        """
        This is for proviidng the set of conditions for the policy.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessIsolationRuleV2ConditionArgs']]]]):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        This is the description of the access policy rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="microtenantId")
    def microtenant_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "microtenant_id")

    @microtenant_id.setter
    def microtenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "microtenant_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        This is the name of the policy rule.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="zpnIsolationProfileId")
    def zpn_isolation_profile_id(self) -> Optional[pulumi.Input[str]]:
        """
        Use zpa*isolation*profile data source to retrieve the necessary Isolation profile ID `zpn_isolation_profile_id`
        """
        return pulumi.get(self, "zpn_isolation_profile_id")

    @zpn_isolation_profile_id.setter
    def zpn_isolation_profile_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zpn_isolation_profile_id", value)


@pulumi.input_type
class _PolicyAccessIsolationRuleV2State:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessIsolationRuleV2ConditionArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 microtenant_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_set_id: Optional[pulumi.Input[str]] = None,
                 zpn_isolation_profile_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PolicyAccessIsolationRuleV2 resources.
        :param pulumi.Input[str] action: This is for providing the rule action. Supported values: `ISOLATE` Default.
        :param pulumi.Input[Sequence[pulumi.Input['PolicyAccessIsolationRuleV2ConditionArgs']]] conditions: This is for proviidng the set of conditions for the policy.
        :param pulumi.Input[str] description: This is the description of the access policy rule.
        :param pulumi.Input[str] name: This is the name of the policy rule.
        :param pulumi.Input[str] zpn_isolation_profile_id: Use zpa*isolation*profile data source to retrieve the necessary Isolation profile ID `zpn_isolation_profile_id`
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if microtenant_id is not None:
            pulumi.set(__self__, "microtenant_id", microtenant_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policy_set_id is not None:
            pulumi.set(__self__, "policy_set_id", policy_set_id)
        if zpn_isolation_profile_id is not None:
            pulumi.set(__self__, "zpn_isolation_profile_id", zpn_isolation_profile_id)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        This is for providing the rule action. Supported values: `ISOLATE` Default.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessIsolationRuleV2ConditionArgs']]]]:
        """
        This is for proviidng the set of conditions for the policy.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessIsolationRuleV2ConditionArgs']]]]):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        This is the description of the access policy rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="microtenantId")
    def microtenant_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "microtenant_id")

    @microtenant_id.setter
    def microtenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "microtenant_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        This is the name of the policy rule.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="policySetId")
    def policy_set_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "policy_set_id")

    @policy_set_id.setter
    def policy_set_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_set_id", value)

    @property
    @pulumi.getter(name="zpnIsolationProfileId")
    def zpn_isolation_profile_id(self) -> Optional[pulumi.Input[str]]:
        """
        Use zpa*isolation*profile data source to retrieve the necessary Isolation profile ID `zpn_isolation_profile_id`
        """
        return pulumi.get(self, "zpn_isolation_profile_id")

    @zpn_isolation_profile_id.setter
    def zpn_isolation_profile_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zpn_isolation_profile_id", value)


class PolicyAccessIsolationRuleV2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyAccessIsolationRuleV2ConditionArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 microtenant_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 zpn_isolation_profile_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        * [Official documentation](https://help.zscaler.com/zpa/about-isolation-policy)
        * [API documentation](https://help.zscaler.com/zpa/configuring-isolation-policies-using-api)

        The **zpa_policy_isolation_rule_v2** resource creates and manages policy access isolation rule in the Zscaler Private Access cloud using a new v2 API endpoint.

          ⚠️ **NOTE**: This resource is recommended if your configuration requires the association of more than 1000 resource criteria per rule.

          ⚠️ **WARNING:**: The attribute ``rule_order`` is now deprecated in favor of the new resource  ``policy_access_rule_reorder``

        ## Example Usage

        ```python
        import pulumi
        import pulumi_zpa as zpa
        import zscaler_pulumi_zpa as zpa

        this_isolation_profile = zpa.get_isolation_profile(name="zpa_isolation_profile")
        this_id_p_controller = zpa.get_id_p_controller(name="Idp_Name")
        email_user_sso = zpa.get_saml_attribute(name="Email_Users",
            idp_name="Idp_Name")
        group_user = zpa.get_saml_attribute(name="GroupName_Users",
            idp_name="Idp_Name")
        a000 = zpa.get_scim_groups(name="A000",
            idp_name="Idp_Name")
        b000 = zpa.get_scim_groups(name="B000",
            idp_name="Idp_Name")
        # Create Policy Access Isolation Rule V2
        this_policy_access_isolation_rule_v2 = zpa.PolicyAccessIsolationRuleV2("thisPolicyAccessIsolationRuleV2",
            description="Example",
            action="ISOLATE",
            zpn_isolation_profile_id=this_isolation_profile.id,
            conditions=[
                zpa.PolicyAccessIsolationRuleV2ConditionArgs(
                    operator="OR",
                    operands=[zpa.PolicyAccessIsolationRuleV2ConditionOperandArgs(
                        object_type="CLIENT_TYPE",
                        values=["zpn_client_type_exporter"],
                    )],
                ),
                zpa.PolicyAccessIsolationRuleV2ConditionArgs(
                    operator="OR",
                    operands=[
                        zpa.PolicyAccessIsolationRuleV2ConditionOperandArgs(
                            object_type="SAML",
                            entry_values=[
                                zpa.PolicyAccessIsolationRuleV2ConditionOperandEntryValueArgs(
                                    rhs="user1@acme.com",
                                    lhs=email_user_sso.id,
                                ),
                                zpa.PolicyAccessIsolationRuleV2ConditionOperandEntryValueArgs(
                                    rhs="A000",
                                    lhs=group_user.id,
                                ),
                            ],
                        ),
                        zpa.PolicyAccessIsolationRuleV2ConditionOperandArgs(
                            object_type="SCIM_GROUP",
                            entry_values=[
                                zpa.PolicyAccessIsolationRuleV2ConditionOperandEntryValueArgs(
                                    rhs=a000.id,
                                    lhs=this_id_p_controller.id,
                                ),
                                zpa.PolicyAccessIsolationRuleV2ConditionOperandEntryValueArgs(
                                    rhs=b000.id,
                                    lhs=this_id_p_controller.id,
                                ),
                            ],
                        ),
                    ],
                ),
            ])
        ```

        ## LHS and RHS Values

        | Object Type | LHS| RHS| VALUES
        |----------|-----------|----------|----------
        | APP  |   |  | ``application_segment_id`` |
        | APP_GROUP  |   |  | ``segment_group_id``|
        | CLIENT_TYPE  |   |  |  ``zpn_client_type_zappl``, ``zpn_client_type_exporter``, ``zpn_client_type_browser_isolation``, ``zpn_client_type_ip_anchoring``, ``zpn_client_type_edge_connector``, ``zpn_client_type_branch_connector``,  ``zpn_client_type_zapp_partner``, ``zpn_client_type_zapp``  |
        | EDGE_CONNECTOR_GROUP  |   |  |  ``<edge_connector_id>`` |
        | MACHINE_GRP   |   |  | ``machine_group_id`` |
        | SAML | ``saml_attribute_id``  | ``attribute_value_to_match`` |
        | SCIM | ``scim_attribute_id``  | ``attribute_value_to_match``  |
        | SCIM_GROUP | ``scim_group_attribute_id``  | ``attribute_value_to_match``  |
        | PLATFORM | ``mac``, ``ios``, ``windows``, ``android``, ``linux`` | ``"true"`` / ``"false"`` |
        | POSTURE | ``posture_udid``  | ``"true"`` / ``"false"`` |

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        Policy access isolation rule can be imported by using `<RULE ID>` as the import ID.

        For example:

        ```sh
        $ pulumi import zpa:index/policyAccessIsolationRuleV2:PolicyAccessIsolationRuleV2 example <rule_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: This is for providing the rule action. Supported values: `ISOLATE` Default.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyAccessIsolationRuleV2ConditionArgs']]]] conditions: This is for proviidng the set of conditions for the policy.
        :param pulumi.Input[str] description: This is the description of the access policy rule.
        :param pulumi.Input[str] name: This is the name of the policy rule.
        :param pulumi.Input[str] zpn_isolation_profile_id: Use zpa*isolation*profile data source to retrieve the necessary Isolation profile ID `zpn_isolation_profile_id`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PolicyAccessIsolationRuleV2Args] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        * [Official documentation](https://help.zscaler.com/zpa/about-isolation-policy)
        * [API documentation](https://help.zscaler.com/zpa/configuring-isolation-policies-using-api)

        The **zpa_policy_isolation_rule_v2** resource creates and manages policy access isolation rule in the Zscaler Private Access cloud using a new v2 API endpoint.

          ⚠️ **NOTE**: This resource is recommended if your configuration requires the association of more than 1000 resource criteria per rule.

          ⚠️ **WARNING:**: The attribute ``rule_order`` is now deprecated in favor of the new resource  ``policy_access_rule_reorder``

        ## Example Usage

        ```python
        import pulumi
        import pulumi_zpa as zpa
        import zscaler_pulumi_zpa as zpa

        this_isolation_profile = zpa.get_isolation_profile(name="zpa_isolation_profile")
        this_id_p_controller = zpa.get_id_p_controller(name="Idp_Name")
        email_user_sso = zpa.get_saml_attribute(name="Email_Users",
            idp_name="Idp_Name")
        group_user = zpa.get_saml_attribute(name="GroupName_Users",
            idp_name="Idp_Name")
        a000 = zpa.get_scim_groups(name="A000",
            idp_name="Idp_Name")
        b000 = zpa.get_scim_groups(name="B000",
            idp_name="Idp_Name")
        # Create Policy Access Isolation Rule V2
        this_policy_access_isolation_rule_v2 = zpa.PolicyAccessIsolationRuleV2("thisPolicyAccessIsolationRuleV2",
            description="Example",
            action="ISOLATE",
            zpn_isolation_profile_id=this_isolation_profile.id,
            conditions=[
                zpa.PolicyAccessIsolationRuleV2ConditionArgs(
                    operator="OR",
                    operands=[zpa.PolicyAccessIsolationRuleV2ConditionOperandArgs(
                        object_type="CLIENT_TYPE",
                        values=["zpn_client_type_exporter"],
                    )],
                ),
                zpa.PolicyAccessIsolationRuleV2ConditionArgs(
                    operator="OR",
                    operands=[
                        zpa.PolicyAccessIsolationRuleV2ConditionOperandArgs(
                            object_type="SAML",
                            entry_values=[
                                zpa.PolicyAccessIsolationRuleV2ConditionOperandEntryValueArgs(
                                    rhs="user1@acme.com",
                                    lhs=email_user_sso.id,
                                ),
                                zpa.PolicyAccessIsolationRuleV2ConditionOperandEntryValueArgs(
                                    rhs="A000",
                                    lhs=group_user.id,
                                ),
                            ],
                        ),
                        zpa.PolicyAccessIsolationRuleV2ConditionOperandArgs(
                            object_type="SCIM_GROUP",
                            entry_values=[
                                zpa.PolicyAccessIsolationRuleV2ConditionOperandEntryValueArgs(
                                    rhs=a000.id,
                                    lhs=this_id_p_controller.id,
                                ),
                                zpa.PolicyAccessIsolationRuleV2ConditionOperandEntryValueArgs(
                                    rhs=b000.id,
                                    lhs=this_id_p_controller.id,
                                ),
                            ],
                        ),
                    ],
                ),
            ])
        ```

        ## LHS and RHS Values

        | Object Type | LHS| RHS| VALUES
        |----------|-----------|----------|----------
        | APP  |   |  | ``application_segment_id`` |
        | APP_GROUP  |   |  | ``segment_group_id``|
        | CLIENT_TYPE  |   |  |  ``zpn_client_type_zappl``, ``zpn_client_type_exporter``, ``zpn_client_type_browser_isolation``, ``zpn_client_type_ip_anchoring``, ``zpn_client_type_edge_connector``, ``zpn_client_type_branch_connector``,  ``zpn_client_type_zapp_partner``, ``zpn_client_type_zapp``  |
        | EDGE_CONNECTOR_GROUP  |   |  |  ``<edge_connector_id>`` |
        | MACHINE_GRP   |   |  | ``machine_group_id`` |
        | SAML | ``saml_attribute_id``  | ``attribute_value_to_match`` |
        | SCIM | ``scim_attribute_id``  | ``attribute_value_to_match``  |
        | SCIM_GROUP | ``scim_group_attribute_id``  | ``attribute_value_to_match``  |
        | PLATFORM | ``mac``, ``ios``, ``windows``, ``android``, ``linux`` | ``"true"`` / ``"false"`` |
        | POSTURE | ``posture_udid``  | ``"true"`` / ``"false"`` |

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        Policy access isolation rule can be imported by using `<RULE ID>` as the import ID.

        For example:

        ```sh
        $ pulumi import zpa:index/policyAccessIsolationRuleV2:PolicyAccessIsolationRuleV2 example <rule_id>
        ```

        :param str resource_name: The name of the resource.
        :param PolicyAccessIsolationRuleV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyAccessIsolationRuleV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyAccessIsolationRuleV2ConditionArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 microtenant_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 zpn_isolation_profile_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PolicyAccessIsolationRuleV2Args.__new__(PolicyAccessIsolationRuleV2Args)

            __props__.__dict__["action"] = action
            __props__.__dict__["conditions"] = conditions
            __props__.__dict__["description"] = description
            __props__.__dict__["microtenant_id"] = microtenant_id
            __props__.__dict__["name"] = name
            __props__.__dict__["zpn_isolation_profile_id"] = zpn_isolation_profile_id
            __props__.__dict__["policy_set_id"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="zpa:index/policyIsolationRuleV2:PolicyIsolationRuleV2")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(PolicyAccessIsolationRuleV2, __self__).__init__(
            'zpa:index/policyAccessIsolationRuleV2:PolicyAccessIsolationRuleV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            conditions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyAccessIsolationRuleV2ConditionArgs']]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            microtenant_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policy_set_id: Optional[pulumi.Input[str]] = None,
            zpn_isolation_profile_id: Optional[pulumi.Input[str]] = None) -> 'PolicyAccessIsolationRuleV2':
        """
        Get an existing PolicyAccessIsolationRuleV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: This is for providing the rule action. Supported values: `ISOLATE` Default.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyAccessIsolationRuleV2ConditionArgs']]]] conditions: This is for proviidng the set of conditions for the policy.
        :param pulumi.Input[str] description: This is the description of the access policy rule.
        :param pulumi.Input[str] name: This is the name of the policy rule.
        :param pulumi.Input[str] zpn_isolation_profile_id: Use zpa*isolation*profile data source to retrieve the necessary Isolation profile ID `zpn_isolation_profile_id`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PolicyAccessIsolationRuleV2State.__new__(_PolicyAccessIsolationRuleV2State)

        __props__.__dict__["action"] = action
        __props__.__dict__["conditions"] = conditions
        __props__.__dict__["description"] = description
        __props__.__dict__["microtenant_id"] = microtenant_id
        __props__.__dict__["name"] = name
        __props__.__dict__["policy_set_id"] = policy_set_id
        __props__.__dict__["zpn_isolation_profile_id"] = zpn_isolation_profile_id
        return PolicyAccessIsolationRuleV2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[Optional[str]]:
        """
        This is for providing the rule action. Supported values: `ISOLATE` Default.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def conditions(self) -> pulumi.Output[Sequence['outputs.PolicyAccessIsolationRuleV2Condition']]:
        """
        This is for proviidng the set of conditions for the policy.
        """
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        This is the description of the access policy rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="microtenantId")
    def microtenant_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "microtenant_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        This is the name of the policy rule.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policySetId")
    def policy_set_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "policy_set_id")

    @property
    @pulumi.getter(name="zpnIsolationProfileId")
    def zpn_isolation_profile_id(self) -> pulumi.Output[str]:
        """
        Use zpa*isolation*profile data source to retrieve the necessary Isolation profile ID `zpn_isolation_profile_id`
        """
        return pulumi.get(self, "zpn_isolation_profile_id")

