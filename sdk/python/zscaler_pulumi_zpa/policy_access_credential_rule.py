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

__all__ = ['PolicyAccessCredentialRuleArgs', 'PolicyAccessCredentialRule']

@pulumi.input_type
class PolicyAccessCredentialRuleArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[builtins.str]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCredentialRuleConditionArgs']]]] = None,
                 credential_pools: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCredentialRuleCredentialPoolArgs']]]] = None,
                 credentials: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCredentialRuleCredentialArgs']]]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 microtenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a PolicyAccessCredentialRule resource.
        :param pulumi.Input[builtins.str] action: This is for providing the rule action.
        :param pulumi.Input[Sequence[pulumi.Input['PolicyAccessCredentialRuleConditionArgs']]] conditions: This is for proviidng the set of conditions for the policy.
        :param pulumi.Input[builtins.str] description: This is the description of the access policy.
        :param pulumi.Input[builtins.str] name: This is the name of the policy.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if credential_pools is not None:
            pulumi.set(__self__, "credential_pools", credential_pools)
        if credentials is not None:
            pulumi.set(__self__, "credentials", credentials)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if microtenant_id is not None:
            pulumi.set(__self__, "microtenant_id", microtenant_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        This is for providing the rule action.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCredentialRuleConditionArgs']]]]:
        """
        This is for proviidng the set of conditions for the policy.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCredentialRuleConditionArgs']]]]):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter(name="credentialPools")
    def credential_pools(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCredentialRuleCredentialPoolArgs']]]]:
        return pulumi.get(self, "credential_pools")

    @credential_pools.setter
    def credential_pools(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCredentialRuleCredentialPoolArgs']]]]):
        pulumi.set(self, "credential_pools", value)

    @property
    @pulumi.getter
    def credentials(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCredentialRuleCredentialArgs']]]]:
        return pulumi.get(self, "credentials")

    @credentials.setter
    def credentials(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCredentialRuleCredentialArgs']]]]):
        pulumi.set(self, "credentials", value)

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
        This is the name of the policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _PolicyAccessCredentialRuleState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[builtins.str]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCredentialRuleConditionArgs']]]] = None,
                 credential_pools: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCredentialRuleCredentialPoolArgs']]]] = None,
                 credentials: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCredentialRuleCredentialArgs']]]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 microtenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 policy_set_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering PolicyAccessCredentialRule resources.
        :param pulumi.Input[builtins.str] action: This is for providing the rule action.
        :param pulumi.Input[Sequence[pulumi.Input['PolicyAccessCredentialRuleConditionArgs']]] conditions: This is for proviidng the set of conditions for the policy.
        :param pulumi.Input[builtins.str] description: This is the description of the access policy.
        :param pulumi.Input[builtins.str] name: This is the name of the policy.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if credential_pools is not None:
            pulumi.set(__self__, "credential_pools", credential_pools)
        if credentials is not None:
            pulumi.set(__self__, "credentials", credentials)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if microtenant_id is not None:
            pulumi.set(__self__, "microtenant_id", microtenant_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policy_set_id is not None:
            pulumi.set(__self__, "policy_set_id", policy_set_id)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        This is for providing the rule action.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCredentialRuleConditionArgs']]]]:
        """
        This is for proviidng the set of conditions for the policy.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCredentialRuleConditionArgs']]]]):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter(name="credentialPools")
    def credential_pools(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCredentialRuleCredentialPoolArgs']]]]:
        return pulumi.get(self, "credential_pools")

    @credential_pools.setter
    def credential_pools(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCredentialRuleCredentialPoolArgs']]]]):
        pulumi.set(self, "credential_pools", value)

    @property
    @pulumi.getter
    def credentials(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCredentialRuleCredentialArgs']]]]:
        return pulumi.get(self, "credentials")

    @credentials.setter
    def credentials(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAccessCredentialRuleCredentialArgs']]]]):
        pulumi.set(self, "credentials", value)

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
        This is the name of the policy.
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


@pulumi.type_token("zpa:index/policyAccessCredentialRule:PolicyAccessCredentialRule")
class PolicyAccessCredentialRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[builtins.str]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessCredentialRuleConditionArgs', 'PolicyAccessCredentialRuleConditionArgsDict']]]]] = None,
                 credential_pools: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessCredentialRuleCredentialPoolArgs', 'PolicyAccessCredentialRuleCredentialPoolArgsDict']]]]] = None,
                 credentials: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessCredentialRuleCredentialArgs', 'PolicyAccessCredentialRuleCredentialArgsDict']]]]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 microtenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        * [Official documentation](https://help.zscaler.com/zpa/about-privileged-capabilities-policy)
        * [API documentation](https://help.zscaler.com/zpa/configuring-privileged-policies-using-api)

        The **zpa_policy_credential_rule** resource creates a policy credential rule in the Zscaler Private Access cloud.

          ⚠️ **WARNING:**: The attribute ``rule_order`` is now deprecated in favor of the new resource  ``policy_access_rule_reorder``

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        Policy access credential can be imported by using `<POLICY CREDENTIAL ID>` as the import ID.

        For example:

        ```sh
        $ pulumi import zpa:index/policyAccessCredentialRule:PolicyAccessCredentialRule example <policy_credential_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] action: This is for providing the rule action.
        :param pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessCredentialRuleConditionArgs', 'PolicyAccessCredentialRuleConditionArgsDict']]]] conditions: This is for proviidng the set of conditions for the policy.
        :param pulumi.Input[builtins.str] description: This is the description of the access policy.
        :param pulumi.Input[builtins.str] name: This is the name of the policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PolicyAccessCredentialRuleArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        * [Official documentation](https://help.zscaler.com/zpa/about-privileged-capabilities-policy)
        * [API documentation](https://help.zscaler.com/zpa/configuring-privileged-policies-using-api)

        The **zpa_policy_credential_rule** resource creates a policy credential rule in the Zscaler Private Access cloud.

          ⚠️ **WARNING:**: The attribute ``rule_order`` is now deprecated in favor of the new resource  ``policy_access_rule_reorder``

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        Policy access credential can be imported by using `<POLICY CREDENTIAL ID>` as the import ID.

        For example:

        ```sh
        $ pulumi import zpa:index/policyAccessCredentialRule:PolicyAccessCredentialRule example <policy_credential_id>
        ```

        :param str resource_name: The name of the resource.
        :param PolicyAccessCredentialRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyAccessCredentialRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[builtins.str]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessCredentialRuleConditionArgs', 'PolicyAccessCredentialRuleConditionArgsDict']]]]] = None,
                 credential_pools: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessCredentialRuleCredentialPoolArgs', 'PolicyAccessCredentialRuleCredentialPoolArgsDict']]]]] = None,
                 credentials: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessCredentialRuleCredentialArgs', 'PolicyAccessCredentialRuleCredentialArgsDict']]]]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 microtenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PolicyAccessCredentialRuleArgs.__new__(PolicyAccessCredentialRuleArgs)

            __props__.__dict__["action"] = action
            __props__.__dict__["conditions"] = conditions
            __props__.__dict__["credential_pools"] = credential_pools
            __props__.__dict__["credentials"] = credentials
            __props__.__dict__["description"] = description
            __props__.__dict__["microtenant_id"] = microtenant_id
            __props__.__dict__["name"] = name
            __props__.__dict__["policy_set_id"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="zpa:index/policyCredentialRule:PolicyCredentialRule")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(PolicyAccessCredentialRule, __self__).__init__(
            'zpa:index/policyAccessCredentialRule:PolicyAccessCredentialRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[builtins.str]] = None,
            conditions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessCredentialRuleConditionArgs', 'PolicyAccessCredentialRuleConditionArgsDict']]]]] = None,
            credential_pools: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessCredentialRuleCredentialPoolArgs', 'PolicyAccessCredentialRuleCredentialPoolArgsDict']]]]] = None,
            credentials: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessCredentialRuleCredentialArgs', 'PolicyAccessCredentialRuleCredentialArgsDict']]]]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            microtenant_id: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            policy_set_id: Optional[pulumi.Input[builtins.str]] = None) -> 'PolicyAccessCredentialRule':
        """
        Get an existing PolicyAccessCredentialRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] action: This is for providing the rule action.
        :param pulumi.Input[Sequence[pulumi.Input[Union['PolicyAccessCredentialRuleConditionArgs', 'PolicyAccessCredentialRuleConditionArgsDict']]]] conditions: This is for proviidng the set of conditions for the policy.
        :param pulumi.Input[builtins.str] description: This is the description of the access policy.
        :param pulumi.Input[builtins.str] name: This is the name of the policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PolicyAccessCredentialRuleState.__new__(_PolicyAccessCredentialRuleState)

        __props__.__dict__["action"] = action
        __props__.__dict__["conditions"] = conditions
        __props__.__dict__["credential_pools"] = credential_pools
        __props__.__dict__["credentials"] = credentials
        __props__.__dict__["description"] = description
        __props__.__dict__["microtenant_id"] = microtenant_id
        __props__.__dict__["name"] = name
        __props__.__dict__["policy_set_id"] = policy_set_id
        return PolicyAccessCredentialRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        This is for providing the rule action.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def conditions(self) -> pulumi.Output[Sequence['outputs.PolicyAccessCredentialRuleCondition']]:
        """
        This is for proviidng the set of conditions for the policy.
        """
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter(name="credentialPools")
    def credential_pools(self) -> pulumi.Output[Optional[Sequence['outputs.PolicyAccessCredentialRuleCredentialPool']]]:
        return pulumi.get(self, "credential_pools")

    @property
    @pulumi.getter
    def credentials(self) -> pulumi.Output[Optional[Sequence['outputs.PolicyAccessCredentialRuleCredential']]]:
        return pulumi.get(self, "credentials")

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
        This is the name of the policy.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policySetId")
    def policy_set_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "policy_set_id")

