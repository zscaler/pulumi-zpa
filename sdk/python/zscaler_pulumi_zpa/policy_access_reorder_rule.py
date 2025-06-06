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

__all__ = ['PolicyAccessReorderRuleArgs', 'PolicyAccessReorderRule']

@pulumi.input_type
class PolicyAccessReorderRuleArgs:
    def __init__(__self__, *,
                 policy_type: pulumi.Input[builtins.str],
                 rules: pulumi.Input[Sequence[pulumi.Input['PolicyAccessReorderRuleRuleArgs']]]):
        """
        The set of arguments for constructing a PolicyAccessReorderRule resource.
        :param pulumi.Input[Sequence[pulumi.Input['PolicyAccessReorderRuleRuleArgs']]] rules: List of rules and their orders
        """
        pulumi.set(__self__, "policy_type", policy_type)
        pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "policy_type")

    @policy_type.setter
    def policy_type(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "policy_type", value)

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Input[Sequence[pulumi.Input['PolicyAccessReorderRuleRuleArgs']]]:
        """
        List of rules and their orders
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: pulumi.Input[Sequence[pulumi.Input['PolicyAccessReorderRuleRuleArgs']]]):
        pulumi.set(self, "rules", value)


@pulumi.input_type
class _PolicyAccessReorderRuleState:
    def __init__(__self__, *,
                 policy_type: Optional[pulumi.Input[builtins.str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessReorderRuleRuleArgs']]]] = None):
        """
        Input properties used for looking up and filtering PolicyAccessReorderRule resources.
        :param pulumi.Input[Sequence[pulumi.Input['PolicyAccessReorderRuleRuleArgs']]] rules: List of rules and their orders
        """
        if policy_type is not None:
            pulumi.set(__self__, "policy_type", policy_type)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "policy_type")

    @policy_type.setter
    def policy_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "policy_type", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessReorderRuleRuleArgs']]]]:
        """
        List of rules and their orders
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessReorderRuleRuleArgs']]]]):
        pulumi.set(self, "rules", value)


@pulumi.type_token("zpa:index/policyAccessReorderRule:PolicyAccessReorderRule")
class PolicyAccessReorderRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_type: Optional[pulumi.Input[builtins.str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessReorderRuleRuleArgs', 'PolicyAccessReorderRuleRuleArgsDict']]]]] = None,
                 __props__=None):
        """
        * [Official documentation](https://help.zscaler.com/zpa/about-access-policy)
        * [API documentation](https://help.zscaler.com/zpa/configuring-access-policies-using-api)

        The **zpa_policy_access_rule_reorder** is a dedicated resource to manage and update `rule_orders` in any of the supported ZPA Policy Access types Zscaler Private Access cloud.

        ⚠️ **WARNING:**: The attribute ``rule_order`` is now deprecated in favor of this resource for all ZPA policy types.

        ⚠️ **WARNING:**: Updating the rule order of an access policy configured using `Zscaler Deception` is not supported. When changing the rule order of a regular access policy and there is an access policy configured using Deception, the rule order of the regular access policy must be greater than the rule order for an access policy configured using Deception. Please refer to the [Zscaler API Documentation](https://help.zscaler.com/zpa/configuring-access-policies-using-api#:~:text=Updating%20the%20rule,configured%20using%20Deception.) for further details.

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessReorderRuleRuleArgs', 'PolicyAccessReorderRuleRuleArgsDict']]]] rules: List of rules and their orders
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PolicyAccessReorderRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        * [Official documentation](https://help.zscaler.com/zpa/about-access-policy)
        * [API documentation](https://help.zscaler.com/zpa/configuring-access-policies-using-api)

        The **zpa_policy_access_rule_reorder** is a dedicated resource to manage and update `rule_orders` in any of the supported ZPA Policy Access types Zscaler Private Access cloud.

        ⚠️ **WARNING:**: The attribute ``rule_order`` is now deprecated in favor of this resource for all ZPA policy types.

        ⚠️ **WARNING:**: Updating the rule order of an access policy configured using `Zscaler Deception` is not supported. When changing the rule order of a regular access policy and there is an access policy configured using Deception, the rule order of the regular access policy must be greater than the rule order for an access policy configured using Deception. Please refer to the [Zscaler API Documentation](https://help.zscaler.com/zpa/configuring-access-policies-using-api#:~:text=Updating%20the%20rule,configured%20using%20Deception.) for further details.

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param PolicyAccessReorderRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyAccessReorderRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_type: Optional[pulumi.Input[builtins.str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessReorderRuleRuleArgs', 'PolicyAccessReorderRuleRuleArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PolicyAccessReorderRuleArgs.__new__(PolicyAccessReorderRuleArgs)

            if policy_type is None and not opts.urn:
                raise TypeError("Missing required property 'policy_type'")
            __props__.__dict__["policy_type"] = policy_type
            if rules is None and not opts.urn:
                raise TypeError("Missing required property 'rules'")
            __props__.__dict__["rules"] = rules
        super(PolicyAccessReorderRule, __self__).__init__(
            'zpa:index/policyAccessReorderRule:PolicyAccessReorderRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            policy_type: Optional[pulumi.Input[builtins.str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessReorderRuleRuleArgs', 'PolicyAccessReorderRuleRuleArgsDict']]]]] = None) -> 'PolicyAccessReorderRule':
        """
        Get an existing PolicyAccessReorderRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessReorderRuleRuleArgs', 'PolicyAccessReorderRuleRuleArgsDict']]]] rules: List of rules and their orders
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PolicyAccessReorderRuleState.__new__(_PolicyAccessReorderRuleState)

        __props__.__dict__["policy_type"] = policy_type
        __props__.__dict__["rules"] = rules
        return PolicyAccessReorderRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "policy_type")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.PolicyAccessReorderRuleRule']]:
        """
        List of rules and their orders
        """
        return pulumi.get(self, "rules")

