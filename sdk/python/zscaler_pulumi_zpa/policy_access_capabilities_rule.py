# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['PolicyAccessCapabilitiesRuleArgs', 'PolicyAccessCapabilitiesRule']

@pulumi.input_type
class PolicyAccessCapabilitiesRuleArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[builtins.str]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCapabilitiesRuleConditionArgs']]]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 microtenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 privileged_capabilities: Optional[pulumi.Input['PolicyAccessCapabilitiesRulePrivilegedCapabilitiesArgs']] = None):
        """
        The set of arguments for constructing a PolicyAccessCapabilitiesRule resource.
        :param pulumi.Input[builtins.str] action: This is for providing the rule action. Supported value: `CHECK_CAPABILITIES`
        :param pulumi.Input[Sequence[pulumi.Input['PolicyAccessCapabilitiesRuleConditionArgs']]] conditions: This is for proviidng the set of conditions for the policy.
        :param pulumi.Input[builtins.str] description: This is the description of the access policy.
        :param pulumi.Input[builtins.str] name: This is the name of the policy rule.
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
        if privileged_capabilities is not None:
            pulumi.set(__self__, "privileged_capabilities", privileged_capabilities)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        This is for providing the rule action. Supported value: `CHECK_CAPABILITIES`
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCapabilitiesRuleConditionArgs']]]]:
        """
        This is for proviidng the set of conditions for the policy.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCapabilitiesRuleConditionArgs']]]]):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        This is the description of the access policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="microtenantId")
    def microtenant_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "microtenant_id")

    @microtenant_id.setter
    def microtenant_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "microtenant_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        This is the name of the policy rule.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="privilegedCapabilities")
    def privileged_capabilities(self) -> Optional[pulumi.Input['PolicyAccessCapabilitiesRulePrivilegedCapabilitiesArgs']]:
        return pulumi.get(self, "privileged_capabilities")

    @privileged_capabilities.setter
    def privileged_capabilities(self, value: Optional[pulumi.Input['PolicyAccessCapabilitiesRulePrivilegedCapabilitiesArgs']]):
        pulumi.set(self, "privileged_capabilities", value)


@pulumi.input_type
class _PolicyAccessCapabilitiesRuleState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[builtins.str]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCapabilitiesRuleConditionArgs']]]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 microtenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 policy_set_id: Optional[pulumi.Input[builtins.str]] = None,
                 privileged_capabilities: Optional[pulumi.Input['PolicyAccessCapabilitiesRulePrivilegedCapabilitiesArgs']] = None):
        """
        Input properties used for looking up and filtering PolicyAccessCapabilitiesRule resources.
        :param pulumi.Input[builtins.str] action: This is for providing the rule action. Supported value: `CHECK_CAPABILITIES`
        :param pulumi.Input[Sequence[pulumi.Input['PolicyAccessCapabilitiesRuleConditionArgs']]] conditions: This is for proviidng the set of conditions for the policy.
        :param pulumi.Input[builtins.str] description: This is the description of the access policy.
        :param pulumi.Input[builtins.str] name: This is the name of the policy rule.
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
        if privileged_capabilities is not None:
            pulumi.set(__self__, "privileged_capabilities", privileged_capabilities)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        This is for providing the rule action. Supported value: `CHECK_CAPABILITIES`
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCapabilitiesRuleConditionArgs']]]]:
        """
        This is for proviidng the set of conditions for the policy.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCapabilitiesRuleConditionArgs']]]]):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        This is the description of the access policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="microtenantId")
    def microtenant_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "microtenant_id")

    @microtenant_id.setter
    def microtenant_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "microtenant_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        This is the name of the policy rule.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="policySetId")
    def policy_set_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "policy_set_id")

    @policy_set_id.setter
    def policy_set_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "policy_set_id", value)

    @property
    @pulumi.getter(name="privilegedCapabilities")
    def privileged_capabilities(self) -> Optional[pulumi.Input['PolicyAccessCapabilitiesRulePrivilegedCapabilitiesArgs']]:
        return pulumi.get(self, "privileged_capabilities")

    @privileged_capabilities.setter
    def privileged_capabilities(self, value: Optional[pulumi.Input['PolicyAccessCapabilitiesRulePrivilegedCapabilitiesArgs']]):
        pulumi.set(self, "privileged_capabilities", value)


@pulumi.type_token("zpa:index/policyAccessCapabilitiesRule:PolicyAccessCapabilitiesRule")
class PolicyAccessCapabilitiesRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[builtins.str]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessCapabilitiesRuleConditionArgs', 'PolicyAccessCapabilitiesRuleConditionArgsDict']]]]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 microtenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 privileged_capabilities: Optional[pulumi.Input[Union['PolicyAccessCapabilitiesRulePrivilegedCapabilitiesArgs', 'PolicyAccessCapabilitiesRulePrivilegedCapabilitiesArgsDict']]] = None,
                 __props__=None):
        """
        * [Official documentation](https://help.zscaler.com/zpa/about-privileged-capabilities-policy)
        * [API documentation](https://help.zscaler.com/zpa/configuring-privileged-policies-using-api)

        The **zpa_policy_capabilities_rule** resource creates a policy capabilities rule in the Zscaler Private Access cloud.

          ⚠️ **WARNING:**: The attribute ``rule_order`` is now deprecated in favor of the new resource  ``policy_access_rule_reorder``

        ## Example Usage

        ```python
        import pulumi
        import pulumi_zpa as zpa
        import zscaler_pulumi_zpa as zpa

        this_id_p_controller = zpa.get_id_p_controller(name="IdP_Users")
        email_user_sso = zpa.get_saml_attribute(name="Email_IdP_Users",
            idp_name="IdP_Users")
        group_user = zpa.get_saml_attribute(name="GroupName_IdP_Users",
            idp_name="IdP_Users")
        a000 = zpa.get_scim_groups(name="A000",
            idp_name="IdP_Users")
        b000 = zpa.get_scim_groups(name="B000",
            idp_name="IdP_Users")
        this_policy_access_capabilities_rule = zpa.PolicyAccessCapabilitiesRule("thisPolicyAccessCapabilitiesRule",
            description="Example",
            action="CHECK_CAPABILITIES",
            privileged_capabilities={
                "file_upload": True,
                "file_download": True,
                "inspect_file_upload": True,
                "clipboard_copy": True,
                "clipboard_paste": True,
                "record_session": True,
            },
            conditions=[{
                "operator": "OR",
                "operands": [
                    {
                        "object_type": "SAML",
                        "entry_values": [
                            {
                                "rhs": "user1@example.com",
                                "lhs": email_user_sso.id,
                            },
                            {
                                "rhs": "A000",
                                "lhs": group_user.id,
                            },
                        ],
                    },
                    {
                        "object_type": "SCIM_GROUP",
                        "entry_values": [
                            {
                                "rhs": a000.id,
                                "lhs": this_id_p_controller.id,
                            },
                            {
                                "rhs": b000.id,
                                "lhs": this_id_p_controller.id,
                            },
                        ],
                    },
                ],
            }])
        ```

        ## LHS and RHS Values

        | Object Type | LHS| RHS| VALUES
        |----------|-----------|----------|----------
        | APP |   |  | ``application_segment_id``
        | APP_GROUP |   |  | ``segment_group_id``
        | SAML | ``saml_attribute_id``  | ``attribute_value_to_match`` |
        | SCIM | ``scim_attribute_id``  | ``attribute_value_to_match``  |
        | SCIM_GROUP | ``scim_group_attribute_id``  | ``attribute_value_to_match``  |

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        Policy access capability can be imported by using `<RULE ID>` as the import ID.

        For example:

        ```sh
        $ pulumi import zpa:index/policyAccessCapabilitiesRule:PolicyAccessCapabilitiesRule example <rule_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] action: This is for providing the rule action. Supported value: `CHECK_CAPABILITIES`
        :param pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessCapabilitiesRuleConditionArgs', 'PolicyAccessCapabilitiesRuleConditionArgsDict']]]] conditions: This is for proviidng the set of conditions for the policy.
        :param pulumi.Input[builtins.str] description: This is the description of the access policy.
        :param pulumi.Input[builtins.str] name: This is the name of the policy rule.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PolicyAccessCapabilitiesRuleArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        * [Official documentation](https://help.zscaler.com/zpa/about-privileged-capabilities-policy)
        * [API documentation](https://help.zscaler.com/zpa/configuring-privileged-policies-using-api)

        The **zpa_policy_capabilities_rule** resource creates a policy capabilities rule in the Zscaler Private Access cloud.

          ⚠️ **WARNING:**: The attribute ``rule_order`` is now deprecated in favor of the new resource  ``policy_access_rule_reorder``

        ## Example Usage

        ```python
        import pulumi
        import pulumi_zpa as zpa
        import zscaler_pulumi_zpa as zpa

        this_id_p_controller = zpa.get_id_p_controller(name="IdP_Users")
        email_user_sso = zpa.get_saml_attribute(name="Email_IdP_Users",
            idp_name="IdP_Users")
        group_user = zpa.get_saml_attribute(name="GroupName_IdP_Users",
            idp_name="IdP_Users")
        a000 = zpa.get_scim_groups(name="A000",
            idp_name="IdP_Users")
        b000 = zpa.get_scim_groups(name="B000",
            idp_name="IdP_Users")
        this_policy_access_capabilities_rule = zpa.PolicyAccessCapabilitiesRule("thisPolicyAccessCapabilitiesRule",
            description="Example",
            action="CHECK_CAPABILITIES",
            privileged_capabilities={
                "file_upload": True,
                "file_download": True,
                "inspect_file_upload": True,
                "clipboard_copy": True,
                "clipboard_paste": True,
                "record_session": True,
            },
            conditions=[{
                "operator": "OR",
                "operands": [
                    {
                        "object_type": "SAML",
                        "entry_values": [
                            {
                                "rhs": "user1@example.com",
                                "lhs": email_user_sso.id,
                            },
                            {
                                "rhs": "A000",
                                "lhs": group_user.id,
                            },
                        ],
                    },
                    {
                        "object_type": "SCIM_GROUP",
                        "entry_values": [
                            {
                                "rhs": a000.id,
                                "lhs": this_id_p_controller.id,
                            },
                            {
                                "rhs": b000.id,
                                "lhs": this_id_p_controller.id,
                            },
                        ],
                    },
                ],
            }])
        ```

        ## LHS and RHS Values

        | Object Type | LHS| RHS| VALUES
        |----------|-----------|----------|----------
        | APP |   |  | ``application_segment_id``
        | APP_GROUP |   |  | ``segment_group_id``
        | SAML | ``saml_attribute_id``  | ``attribute_value_to_match`` |
        | SCIM | ``scim_attribute_id``  | ``attribute_value_to_match``  |
        | SCIM_GROUP | ``scim_group_attribute_id``  | ``attribute_value_to_match``  |

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        Policy access capability can be imported by using `<RULE ID>` as the import ID.

        For example:

        ```sh
        $ pulumi import zpa:index/policyAccessCapabilitiesRule:PolicyAccessCapabilitiesRule example <rule_id>
        ```

        :param str resource_name: The name of the resource.
        :param PolicyAccessCapabilitiesRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyAccessCapabilitiesRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[builtins.str]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessCapabilitiesRuleConditionArgs', 'PolicyAccessCapabilitiesRuleConditionArgsDict']]]]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 microtenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 privileged_capabilities: Optional[pulumi.Input[Union['PolicyAccessCapabilitiesRulePrivilegedCapabilitiesArgs', 'PolicyAccessCapabilitiesRulePrivilegedCapabilitiesArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PolicyAccessCapabilitiesRuleArgs.__new__(PolicyAccessCapabilitiesRuleArgs)

            __props__.__dict__["action"] = action
            __props__.__dict__["conditions"] = conditions
            __props__.__dict__["description"] = description
            __props__.__dict__["microtenant_id"] = microtenant_id
            __props__.__dict__["name"] = name
            __props__.__dict__["privileged_capabilities"] = privileged_capabilities
            __props__.__dict__["policy_set_id"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="zpa:index/policyCapabilitiesRule:PolicyCapabilitiesRule")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(PolicyAccessCapabilitiesRule, __self__).__init__(
            'zpa:index/policyAccessCapabilitiesRule:PolicyAccessCapabilitiesRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[builtins.str]] = None,
            conditions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessCapabilitiesRuleConditionArgs', 'PolicyAccessCapabilitiesRuleConditionArgsDict']]]]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            microtenant_id: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            policy_set_id: Optional[pulumi.Input[builtins.str]] = None,
            privileged_capabilities: Optional[pulumi.Input[Union['PolicyAccessCapabilitiesRulePrivilegedCapabilitiesArgs', 'PolicyAccessCapabilitiesRulePrivilegedCapabilitiesArgsDict']]] = None) -> 'PolicyAccessCapabilitiesRule':
        """
        Get an existing PolicyAccessCapabilitiesRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] action: This is for providing the rule action. Supported value: `CHECK_CAPABILITIES`
        :param pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessCapabilitiesRuleConditionArgs', 'PolicyAccessCapabilitiesRuleConditionArgsDict']]]] conditions: This is for proviidng the set of conditions for the policy.
        :param pulumi.Input[builtins.str] description: This is the description of the access policy.
        :param pulumi.Input[builtins.str] name: This is the name of the policy rule.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PolicyAccessCapabilitiesRuleState.__new__(_PolicyAccessCapabilitiesRuleState)

        __props__.__dict__["action"] = action
        __props__.__dict__["conditions"] = conditions
        __props__.__dict__["description"] = description
        __props__.__dict__["microtenant_id"] = microtenant_id
        __props__.__dict__["name"] = name
        __props__.__dict__["policy_set_id"] = policy_set_id
        __props__.__dict__["privileged_capabilities"] = privileged_capabilities
        return PolicyAccessCapabilitiesRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        This is for providing the rule action. Supported value: `CHECK_CAPABILITIES`
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def conditions(self) -> pulumi.Output[Sequence['outputs.PolicyAccessCapabilitiesRuleCondition']]:
        """
        This is for proviidng the set of conditions for the policy.
        """
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        This is the description of the access policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="microtenantId")
    def microtenant_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "microtenant_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        This is the name of the policy rule.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policySetId")
    def policy_set_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "policy_set_id")

    @property
    @pulumi.getter(name="privilegedCapabilities")
    def privileged_capabilities(self) -> pulumi.Output['outputs.PolicyAccessCapabilitiesRulePrivilegedCapabilities']:
        return pulumi.get(self, "privileged_capabilities")

