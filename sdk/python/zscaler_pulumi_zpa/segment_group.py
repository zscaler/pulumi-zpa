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

__all__ = ['SegmentGroupArgs', 'SegmentGroup']

@pulumi.input_type
class SegmentGroupArgs:
    def __init__(__self__, *,
                 applications: Optional[pulumi.Input[Sequence[pulumi.Input['SegmentGroupApplicationArgs']]]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 microtenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a SegmentGroup resource.
        :param pulumi.Input[builtins.str] description: Description of the app group.
        :param pulumi.Input[builtins.bool] enabled: Whether this app group is enabled or not.
        :param pulumi.Input[builtins.str] name: Name of the app group.
        """
        if applications is not None:
            pulumi.set(__self__, "applications", applications)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if microtenant_id is not None:
            pulumi.set(__self__, "microtenant_id", microtenant_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def applications(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SegmentGroupApplicationArgs']]]]:
        return pulumi.get(self, "applications")

    @applications.setter
    def applications(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SegmentGroupApplicationArgs']]]]):
        pulumi.set(self, "applications", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Description of the app group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Whether this app group is enabled or not.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

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
        Name of the app group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _SegmentGroupState:
    def __init__(__self__, *,
                 applications: Optional[pulumi.Input[Sequence[pulumi.Input['SegmentGroupApplicationArgs']]]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 microtenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering SegmentGroup resources.
        :param pulumi.Input[builtins.str] description: Description of the app group.
        :param pulumi.Input[builtins.bool] enabled: Whether this app group is enabled or not.
        :param pulumi.Input[builtins.str] name: Name of the app group.
        """
        if applications is not None:
            pulumi.set(__self__, "applications", applications)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if microtenant_id is not None:
            pulumi.set(__self__, "microtenant_id", microtenant_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def applications(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SegmentGroupApplicationArgs']]]]:
        return pulumi.get(self, "applications")

    @applications.setter
    def applications(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SegmentGroupApplicationArgs']]]]):
        pulumi.set(self, "applications", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Description of the app group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Whether this app group is enabled or not.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

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
        Name of the app group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.type_token("zpa:index/segmentGroup:SegmentGroup")
class SegmentGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 applications: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SegmentGroupApplicationArgs', 'SegmentGroupApplicationArgsDict']]]]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 microtenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import zscaler_pulumi_zpa as zpa

        # ZPA Segment Group resource
        test_segment_group = zpa.SegmentGroup("testSegmentGroup",
            description="test1-segment-group",
            enabled=True)
        ```

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        **segment_group** can be imported by using `<SEGMENT GROUP ID>` or `<SEGMENT GROUP NAME>` as the import ID.

        For example:

        ```sh
        $ pulumi import zpa:index/segmentGroup:SegmentGroup example <segment_group_id>
        ```

        or

        ```sh
        $ pulumi import zpa:index/segmentGroup:SegmentGroup example <segment_group_name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: Description of the app group.
        :param pulumi.Input[builtins.bool] enabled: Whether this app group is enabled or not.
        :param pulumi.Input[builtins.str] name: Name of the app group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SegmentGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import zscaler_pulumi_zpa as zpa

        # ZPA Segment Group resource
        test_segment_group = zpa.SegmentGroup("testSegmentGroup",
            description="test1-segment-group",
            enabled=True)
        ```

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        **segment_group** can be imported by using `<SEGMENT GROUP ID>` or `<SEGMENT GROUP NAME>` as the import ID.

        For example:

        ```sh
        $ pulumi import zpa:index/segmentGroup:SegmentGroup example <segment_group_id>
        ```

        or

        ```sh
        $ pulumi import zpa:index/segmentGroup:SegmentGroup example <segment_group_name>
        ```

        :param str resource_name: The name of the resource.
        :param SegmentGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SegmentGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 applications: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SegmentGroupApplicationArgs', 'SegmentGroupApplicationArgsDict']]]]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 microtenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SegmentGroupArgs.__new__(SegmentGroupArgs)

            __props__.__dict__["applications"] = applications
            __props__.__dict__["description"] = description
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["microtenant_id"] = microtenant_id
            __props__.__dict__["name"] = name
        super(SegmentGroup, __self__).__init__(
            'zpa:index/segmentGroup:SegmentGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            applications: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SegmentGroupApplicationArgs', 'SegmentGroupApplicationArgsDict']]]]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            enabled: Optional[pulumi.Input[builtins.bool]] = None,
            microtenant_id: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None) -> 'SegmentGroup':
        """
        Get an existing SegmentGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: Description of the app group.
        :param pulumi.Input[builtins.bool] enabled: Whether this app group is enabled or not.
        :param pulumi.Input[builtins.str] name: Name of the app group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SegmentGroupState.__new__(_SegmentGroupState)

        __props__.__dict__["applications"] = applications
        __props__.__dict__["description"] = description
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["microtenant_id"] = microtenant_id
        __props__.__dict__["name"] = name
        return SegmentGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def applications(self) -> pulumi.Output[Sequence['outputs.SegmentGroupApplication']]:
        return pulumi.get(self, "applications")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Description of the app group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Whether this app group is enabled or not.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="microtenantId")
    def microtenant_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "microtenant_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the app group.
        """
        return pulumi.get(self, "name")

