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

__all__ = ['PolicyAccessInspectionRuleV2Args', 'PolicyAccessInspectionRuleV2']

@pulumi.input_type
class PolicyAccessInspectionRuleV2Args:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessInspectionRuleV2ConditionArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 microtenant_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 zpn_inspection_profile_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PolicyAccessInspectionRuleV2 resource.
        :param pulumi.Input[str] action: This is for providing the rule action. Supported values: `INSPECT` and `BYPASS_INSPECT`.
        :param pulumi.Input[Sequence[pulumi.Input['PolicyAccessInspectionRuleV2ConditionArgs']]] conditions: This is for proviidng the set of conditions for the policy.
        :param pulumi.Input[str] description: This is the description of the access policy rule.
        :param pulumi.Input[str] name: - (String) This is the name of the policy rule.
        :param pulumi.Input[str] zpn_inspection_profile_id: An inspection profile is required if the `action` is set to `INSPECT`
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
        if zpn_inspection_profile_id is not None:
            pulumi.set(__self__, "zpn_inspection_profile_id", zpn_inspection_profile_id)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        This is for providing the rule action. Supported values: `INSPECT` and `BYPASS_INSPECT`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessInspectionRuleV2ConditionArgs']]]]:
        """
        This is for proviidng the set of conditions for the policy.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessInspectionRuleV2ConditionArgs']]]]):
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
        - (String) This is the name of the policy rule.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="zpnInspectionProfileId")
    def zpn_inspection_profile_id(self) -> Optional[pulumi.Input[str]]:
        """
        An inspection profile is required if the `action` is set to `INSPECT`
        """
        return pulumi.get(self, "zpn_inspection_profile_id")

    @zpn_inspection_profile_id.setter
    def zpn_inspection_profile_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zpn_inspection_profile_id", value)


@pulumi.input_type
class _PolicyAccessInspectionRuleV2State:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessInspectionRuleV2ConditionArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 microtenant_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_set_id: Optional[pulumi.Input[str]] = None,
                 zpn_inspection_profile_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PolicyAccessInspectionRuleV2 resources.
        :param pulumi.Input[str] action: This is for providing the rule action. Supported values: `INSPECT` and `BYPASS_INSPECT`.
        :param pulumi.Input[Sequence[pulumi.Input['PolicyAccessInspectionRuleV2ConditionArgs']]] conditions: This is for proviidng the set of conditions for the policy.
        :param pulumi.Input[str] description: This is the description of the access policy rule.
        :param pulumi.Input[str] name: - (String) This is the name of the policy rule.
        :param pulumi.Input[str] zpn_inspection_profile_id: An inspection profile is required if the `action` is set to `INSPECT`
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
        if zpn_inspection_profile_id is not None:
            pulumi.set(__self__, "zpn_inspection_profile_id", zpn_inspection_profile_id)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        This is for providing the rule action. Supported values: `INSPECT` and `BYPASS_INSPECT`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessInspectionRuleV2ConditionArgs']]]]:
        """
        This is for proviidng the set of conditions for the policy.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessInspectionRuleV2ConditionArgs']]]]):
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
        - (String) This is the name of the policy rule.
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
    @pulumi.getter(name="zpnInspectionProfileId")
    def zpn_inspection_profile_id(self) -> Optional[pulumi.Input[str]]:
        """
        An inspection profile is required if the `action` is set to `INSPECT`
        """
        return pulumi.get(self, "zpn_inspection_profile_id")

    @zpn_inspection_profile_id.setter
    def zpn_inspection_profile_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zpn_inspection_profile_id", value)


class PolicyAccessInspectionRuleV2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyAccessInspectionRuleV2ConditionArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 microtenant_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 zpn_inspection_profile_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        * [Official documentation](https://help.zscaler.com/zpa/about-security-policy)
        * [API documentation](https://help.zscaler.com/zpa/configuring-appprotection-policies-using-api)

        The **zpa_policy_inspection_rule_v2** resource creates and manages policy access inspection rule in the Zscaler Private Access cloud using a new v2 API endpoint.

          ⚠️ **NOTE**: This resource is recommended if your configuration requires the association of more than 1000 resource criteria per rule.

          ⚠️ **WARNING:**: The attribute ``rule_order`` is now deprecated in favor of the new resource  ``policy_access_rule_reorder``

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        Policy access inspection rule can be imported by using `<RULE ID>` as the import ID.

        For example:

        ```sh
        $ pulumi import zpa:index/policyAccessInspectionRuleV2:PolicyAccessInspectionRuleV2 example <rule_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: This is for providing the rule action. Supported values: `INSPECT` and `BYPASS_INSPECT`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyAccessInspectionRuleV2ConditionArgs']]]] conditions: This is for proviidng the set of conditions for the policy.
        :param pulumi.Input[str] description: This is the description of the access policy rule.
        :param pulumi.Input[str] name: - (String) This is the name of the policy rule.
        :param pulumi.Input[str] zpn_inspection_profile_id: An inspection profile is required if the `action` is set to `INSPECT`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PolicyAccessInspectionRuleV2Args] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        * [Official documentation](https://help.zscaler.com/zpa/about-security-policy)
        * [API documentation](https://help.zscaler.com/zpa/configuring-appprotection-policies-using-api)

        The **zpa_policy_inspection_rule_v2** resource creates and manages policy access inspection rule in the Zscaler Private Access cloud using a new v2 API endpoint.

          ⚠️ **NOTE**: This resource is recommended if your configuration requires the association of more than 1000 resource criteria per rule.

          ⚠️ **WARNING:**: The attribute ``rule_order`` is now deprecated in favor of the new resource  ``policy_access_rule_reorder``

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        Policy access inspection rule can be imported by using `<RULE ID>` as the import ID.

        For example:

        ```sh
        $ pulumi import zpa:index/policyAccessInspectionRuleV2:PolicyAccessInspectionRuleV2 example <rule_id>
        ```

        :param str resource_name: The name of the resource.
        :param PolicyAccessInspectionRuleV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyAccessInspectionRuleV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyAccessInspectionRuleV2ConditionArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 microtenant_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 zpn_inspection_profile_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PolicyAccessInspectionRuleV2Args.__new__(PolicyAccessInspectionRuleV2Args)

            __props__.__dict__["action"] = action
            __props__.__dict__["conditions"] = conditions
            __props__.__dict__["description"] = description
            __props__.__dict__["microtenant_id"] = microtenant_id
            __props__.__dict__["name"] = name
            __props__.__dict__["zpn_inspection_profile_id"] = zpn_inspection_profile_id
            __props__.__dict__["policy_set_id"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="zpa:index/policyInspectionRuleV2:PolicyInspectionRuleV2")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(PolicyAccessInspectionRuleV2, __self__).__init__(
            'zpa:index/policyAccessInspectionRuleV2:PolicyAccessInspectionRuleV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            conditions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyAccessInspectionRuleV2ConditionArgs']]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            microtenant_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policy_set_id: Optional[pulumi.Input[str]] = None,
            zpn_inspection_profile_id: Optional[pulumi.Input[str]] = None) -> 'PolicyAccessInspectionRuleV2':
        """
        Get an existing PolicyAccessInspectionRuleV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: This is for providing the rule action. Supported values: `INSPECT` and `BYPASS_INSPECT`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyAccessInspectionRuleV2ConditionArgs']]]] conditions: This is for proviidng the set of conditions for the policy.
        :param pulumi.Input[str] description: This is the description of the access policy rule.
        :param pulumi.Input[str] name: - (String) This is the name of the policy rule.
        :param pulumi.Input[str] zpn_inspection_profile_id: An inspection profile is required if the `action` is set to `INSPECT`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PolicyAccessInspectionRuleV2State.__new__(_PolicyAccessInspectionRuleV2State)

        __props__.__dict__["action"] = action
        __props__.__dict__["conditions"] = conditions
        __props__.__dict__["description"] = description
        __props__.__dict__["microtenant_id"] = microtenant_id
        __props__.__dict__["name"] = name
        __props__.__dict__["policy_set_id"] = policy_set_id
        __props__.__dict__["zpn_inspection_profile_id"] = zpn_inspection_profile_id
        return PolicyAccessInspectionRuleV2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[Optional[str]]:
        """
        This is for providing the rule action. Supported values: `INSPECT` and `BYPASS_INSPECT`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def conditions(self) -> pulumi.Output[Sequence['outputs.PolicyAccessInspectionRuleV2Condition']]:
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
        - (String) This is the name of the policy rule.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policySetId")
    def policy_set_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "policy_set_id")

    @property
    @pulumi.getter(name="zpnInspectionProfileId")
    def zpn_inspection_profile_id(self) -> pulumi.Output[str]:
        """
        An inspection profile is required if the `action` is set to `INSPECT`
        """
        return pulumi.get(self, "zpn_inspection_profile_id")

