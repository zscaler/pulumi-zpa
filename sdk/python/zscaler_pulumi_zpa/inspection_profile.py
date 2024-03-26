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

__all__ = ['InspectionProfileArgs', 'InspectionProfile']

@pulumi.input_type
class InspectionProfileArgs:
    def __init__(__self__, *,
                 associate_all_controls: Optional[pulumi.Input[bool]] = None,
                 controls_infos: Optional[pulumi.Input[Sequence[pulumi.Input['InspectionProfileControlsInfoArgs']]]] = None,
                 custom_controls: Optional[pulumi.Input[Sequence[pulumi.Input['InspectionProfileCustomControlArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 global_control_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 incarnation_number: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 paranoia_level: Optional[pulumi.Input[str]] = None,
                 predefined_controls: Optional[pulumi.Input[Sequence[pulumi.Input['InspectionProfilePredefinedControlArgs']]]] = None,
                 predefined_controls_version: Optional[pulumi.Input[str]] = None,
                 zs_defined_control_choice: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InspectionProfile resource.
        :param pulumi.Input[Sequence[pulumi.Input['InspectionProfileControlsInfoArgs']]] controls_infos: (Optional) Types for custom controls
        :param pulumi.Input[Sequence[pulumi.Input['InspectionProfileCustomControlArgs']]] custom_controls: (Optional) Types for custom controls
        :param pulumi.Input[str] description: Description of the inspection profile.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] global_control_actions: The actions of the predefined, custom, or override controls
        :param pulumi.Input[str] name: The name of the inspection profile.
        :param pulumi.Input[str] paranoia_level: OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        :param pulumi.Input[Sequence[pulumi.Input['InspectionProfilePredefinedControlArgs']]] predefined_controls: The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefined_controls` can be set by using the data source `data_source_zpa_predefined_controls` or by group using the data source `get_inspection_all_predefined_controls`.
        :param pulumi.Input[str] predefined_controls_version: The protocol for the AppProtection application
        :param pulumi.Input[str] zs_defined_control_choice: Indicates the user's choice for the ThreatLabZ Controls. Supported values: ALL and SPECIFIC
        """
        if associate_all_controls is not None:
            pulumi.set(__self__, "associate_all_controls", associate_all_controls)
        if controls_infos is not None:
            pulumi.set(__self__, "controls_infos", controls_infos)
        if custom_controls is not None:
            pulumi.set(__self__, "custom_controls", custom_controls)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if global_control_actions is not None:
            pulumi.set(__self__, "global_control_actions", global_control_actions)
        if incarnation_number is not None:
            pulumi.set(__self__, "incarnation_number", incarnation_number)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if paranoia_level is not None:
            pulumi.set(__self__, "paranoia_level", paranoia_level)
        if predefined_controls is not None:
            pulumi.set(__self__, "predefined_controls", predefined_controls)
        if predefined_controls_version is not None:
            pulumi.set(__self__, "predefined_controls_version", predefined_controls_version)
        if zs_defined_control_choice is not None:
            pulumi.set(__self__, "zs_defined_control_choice", zs_defined_control_choice)

    @property
    @pulumi.getter(name="associateAllControls")
    def associate_all_controls(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "associate_all_controls")

    @associate_all_controls.setter
    def associate_all_controls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "associate_all_controls", value)

    @property
    @pulumi.getter(name="controlsInfos")
    def controls_infos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InspectionProfileControlsInfoArgs']]]]:
        """
        (Optional) Types for custom controls
        """
        return pulumi.get(self, "controls_infos")

    @controls_infos.setter
    def controls_infos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InspectionProfileControlsInfoArgs']]]]):
        pulumi.set(self, "controls_infos", value)

    @property
    @pulumi.getter(name="customControls")
    def custom_controls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InspectionProfileCustomControlArgs']]]]:
        """
        (Optional) Types for custom controls
        """
        return pulumi.get(self, "custom_controls")

    @custom_controls.setter
    def custom_controls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InspectionProfileCustomControlArgs']]]]):
        pulumi.set(self, "custom_controls", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the inspection profile.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="globalControlActions")
    def global_control_actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The actions of the predefined, custom, or override controls
        """
        return pulumi.get(self, "global_control_actions")

    @global_control_actions.setter
    def global_control_actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "global_control_actions", value)

    @property
    @pulumi.getter(name="incarnationNumber")
    def incarnation_number(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "incarnation_number")

    @incarnation_number.setter
    def incarnation_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "incarnation_number", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the inspection profile.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="paranoiaLevel")
    def paranoia_level(self) -> Optional[pulumi.Input[str]]:
        """
        OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        """
        return pulumi.get(self, "paranoia_level")

    @paranoia_level.setter
    def paranoia_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "paranoia_level", value)

    @property
    @pulumi.getter(name="predefinedControls")
    def predefined_controls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InspectionProfilePredefinedControlArgs']]]]:
        """
        The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefined_controls` can be set by using the data source `data_source_zpa_predefined_controls` or by group using the data source `get_inspection_all_predefined_controls`.
        """
        return pulumi.get(self, "predefined_controls")

    @predefined_controls.setter
    def predefined_controls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InspectionProfilePredefinedControlArgs']]]]):
        pulumi.set(self, "predefined_controls", value)

    @property
    @pulumi.getter(name="predefinedControlsVersion")
    def predefined_controls_version(self) -> Optional[pulumi.Input[str]]:
        """
        The protocol for the AppProtection application
        """
        return pulumi.get(self, "predefined_controls_version")

    @predefined_controls_version.setter
    def predefined_controls_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "predefined_controls_version", value)

    @property
    @pulumi.getter(name="zsDefinedControlChoice")
    def zs_defined_control_choice(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the user's choice for the ThreatLabZ Controls. Supported values: ALL and SPECIFIC
        """
        return pulumi.get(self, "zs_defined_control_choice")

    @zs_defined_control_choice.setter
    def zs_defined_control_choice(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zs_defined_control_choice", value)


@pulumi.input_type
class _InspectionProfileState:
    def __init__(__self__, *,
                 associate_all_controls: Optional[pulumi.Input[bool]] = None,
                 controls_infos: Optional[pulumi.Input[Sequence[pulumi.Input['InspectionProfileControlsInfoArgs']]]] = None,
                 custom_controls: Optional[pulumi.Input[Sequence[pulumi.Input['InspectionProfileCustomControlArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 global_control_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 incarnation_number: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 paranoia_level: Optional[pulumi.Input[str]] = None,
                 predefined_controls: Optional[pulumi.Input[Sequence[pulumi.Input['InspectionProfilePredefinedControlArgs']]]] = None,
                 predefined_controls_version: Optional[pulumi.Input[str]] = None,
                 zs_defined_control_choice: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InspectionProfile resources.
        :param pulumi.Input[Sequence[pulumi.Input['InspectionProfileControlsInfoArgs']]] controls_infos: (Optional) Types for custom controls
        :param pulumi.Input[Sequence[pulumi.Input['InspectionProfileCustomControlArgs']]] custom_controls: (Optional) Types for custom controls
        :param pulumi.Input[str] description: Description of the inspection profile.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] global_control_actions: The actions of the predefined, custom, or override controls
        :param pulumi.Input[str] name: The name of the inspection profile.
        :param pulumi.Input[str] paranoia_level: OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        :param pulumi.Input[Sequence[pulumi.Input['InspectionProfilePredefinedControlArgs']]] predefined_controls: The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefined_controls` can be set by using the data source `data_source_zpa_predefined_controls` or by group using the data source `get_inspection_all_predefined_controls`.
        :param pulumi.Input[str] predefined_controls_version: The protocol for the AppProtection application
        :param pulumi.Input[str] zs_defined_control_choice: Indicates the user's choice for the ThreatLabZ Controls. Supported values: ALL and SPECIFIC
        """
        if associate_all_controls is not None:
            pulumi.set(__self__, "associate_all_controls", associate_all_controls)
        if controls_infos is not None:
            pulumi.set(__self__, "controls_infos", controls_infos)
        if custom_controls is not None:
            pulumi.set(__self__, "custom_controls", custom_controls)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if global_control_actions is not None:
            pulumi.set(__self__, "global_control_actions", global_control_actions)
        if incarnation_number is not None:
            pulumi.set(__self__, "incarnation_number", incarnation_number)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if paranoia_level is not None:
            pulumi.set(__self__, "paranoia_level", paranoia_level)
        if predefined_controls is not None:
            pulumi.set(__self__, "predefined_controls", predefined_controls)
        if predefined_controls_version is not None:
            pulumi.set(__self__, "predefined_controls_version", predefined_controls_version)
        if zs_defined_control_choice is not None:
            pulumi.set(__self__, "zs_defined_control_choice", zs_defined_control_choice)

    @property
    @pulumi.getter(name="associateAllControls")
    def associate_all_controls(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "associate_all_controls")

    @associate_all_controls.setter
    def associate_all_controls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "associate_all_controls", value)

    @property
    @pulumi.getter(name="controlsInfos")
    def controls_infos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InspectionProfileControlsInfoArgs']]]]:
        """
        (Optional) Types for custom controls
        """
        return pulumi.get(self, "controls_infos")

    @controls_infos.setter
    def controls_infos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InspectionProfileControlsInfoArgs']]]]):
        pulumi.set(self, "controls_infos", value)

    @property
    @pulumi.getter(name="customControls")
    def custom_controls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InspectionProfileCustomControlArgs']]]]:
        """
        (Optional) Types for custom controls
        """
        return pulumi.get(self, "custom_controls")

    @custom_controls.setter
    def custom_controls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InspectionProfileCustomControlArgs']]]]):
        pulumi.set(self, "custom_controls", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the inspection profile.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="globalControlActions")
    def global_control_actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The actions of the predefined, custom, or override controls
        """
        return pulumi.get(self, "global_control_actions")

    @global_control_actions.setter
    def global_control_actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "global_control_actions", value)

    @property
    @pulumi.getter(name="incarnationNumber")
    def incarnation_number(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "incarnation_number")

    @incarnation_number.setter
    def incarnation_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "incarnation_number", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the inspection profile.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="paranoiaLevel")
    def paranoia_level(self) -> Optional[pulumi.Input[str]]:
        """
        OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        """
        return pulumi.get(self, "paranoia_level")

    @paranoia_level.setter
    def paranoia_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "paranoia_level", value)

    @property
    @pulumi.getter(name="predefinedControls")
    def predefined_controls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InspectionProfilePredefinedControlArgs']]]]:
        """
        The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefined_controls` can be set by using the data source `data_source_zpa_predefined_controls` or by group using the data source `get_inspection_all_predefined_controls`.
        """
        return pulumi.get(self, "predefined_controls")

    @predefined_controls.setter
    def predefined_controls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InspectionProfilePredefinedControlArgs']]]]):
        pulumi.set(self, "predefined_controls", value)

    @property
    @pulumi.getter(name="predefinedControlsVersion")
    def predefined_controls_version(self) -> Optional[pulumi.Input[str]]:
        """
        The protocol for the AppProtection application
        """
        return pulumi.get(self, "predefined_controls_version")

    @predefined_controls_version.setter
    def predefined_controls_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "predefined_controls_version", value)

    @property
    @pulumi.getter(name="zsDefinedControlChoice")
    def zs_defined_control_choice(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the user's choice for the ThreatLabZ Controls. Supported values: ALL and SPECIFIC
        """
        return pulumi.get(self, "zs_defined_control_choice")

    @zs_defined_control_choice.setter
    def zs_defined_control_choice(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zs_defined_control_choice", value)


class InspectionProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 associate_all_controls: Optional[pulumi.Input[bool]] = None,
                 controls_infos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionProfileControlsInfoArgs']]]]] = None,
                 custom_controls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionProfileCustomControlArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 global_control_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 incarnation_number: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 paranoia_level: Optional[pulumi.Input[str]] = None,
                 predefined_controls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionProfilePredefinedControlArgs']]]]] = None,
                 predefined_controls_version: Optional[pulumi.Input[str]] = None,
                 zs_defined_control_choice: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The  **zpa_inspection_profile** resource creates an inspection profile in the Zscaler Private Access cloud. This resource can then be referenced in an inspection custom control resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionProfileControlsInfoArgs']]]] controls_infos: (Optional) Types for custom controls
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionProfileCustomControlArgs']]]] custom_controls: (Optional) Types for custom controls
        :param pulumi.Input[str] description: Description of the inspection profile.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] global_control_actions: The actions of the predefined, custom, or override controls
        :param pulumi.Input[str] name: The name of the inspection profile.
        :param pulumi.Input[str] paranoia_level: OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionProfilePredefinedControlArgs']]]] predefined_controls: The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefined_controls` can be set by using the data source `data_source_zpa_predefined_controls` or by group using the data source `get_inspection_all_predefined_controls`.
        :param pulumi.Input[str] predefined_controls_version: The protocol for the AppProtection application
        :param pulumi.Input[str] zs_defined_control_choice: Indicates the user's choice for the ThreatLabZ Controls. Supported values: ALL and SPECIFIC
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[InspectionProfileArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The  **zpa_inspection_profile** resource creates an inspection profile in the Zscaler Private Access cloud. This resource can then be referenced in an inspection custom control resource.

        :param str resource_name: The name of the resource.
        :param InspectionProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InspectionProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 associate_all_controls: Optional[pulumi.Input[bool]] = None,
                 controls_infos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionProfileControlsInfoArgs']]]]] = None,
                 custom_controls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionProfileCustomControlArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 global_control_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 incarnation_number: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 paranoia_level: Optional[pulumi.Input[str]] = None,
                 predefined_controls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionProfilePredefinedControlArgs']]]]] = None,
                 predefined_controls_version: Optional[pulumi.Input[str]] = None,
                 zs_defined_control_choice: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InspectionProfileArgs.__new__(InspectionProfileArgs)

            __props__.__dict__["associate_all_controls"] = associate_all_controls
            __props__.__dict__["controls_infos"] = controls_infos
            __props__.__dict__["custom_controls"] = custom_controls
            __props__.__dict__["description"] = description
            __props__.__dict__["global_control_actions"] = global_control_actions
            __props__.__dict__["incarnation_number"] = incarnation_number
            __props__.__dict__["name"] = name
            __props__.__dict__["paranoia_level"] = paranoia_level
            __props__.__dict__["predefined_controls"] = predefined_controls
            __props__.__dict__["predefined_controls_version"] = predefined_controls_version
            __props__.__dict__["zs_defined_control_choice"] = zs_defined_control_choice
        super(InspectionProfile, __self__).__init__(
            'zpa:index/inspectionProfile:InspectionProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            associate_all_controls: Optional[pulumi.Input[bool]] = None,
            controls_infos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionProfileControlsInfoArgs']]]]] = None,
            custom_controls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionProfileCustomControlArgs']]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            global_control_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            incarnation_number: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            paranoia_level: Optional[pulumi.Input[str]] = None,
            predefined_controls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionProfilePredefinedControlArgs']]]]] = None,
            predefined_controls_version: Optional[pulumi.Input[str]] = None,
            zs_defined_control_choice: Optional[pulumi.Input[str]] = None) -> 'InspectionProfile':
        """
        Get an existing InspectionProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionProfileControlsInfoArgs']]]] controls_infos: (Optional) Types for custom controls
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionProfileCustomControlArgs']]]] custom_controls: (Optional) Types for custom controls
        :param pulumi.Input[str] description: Description of the inspection profile.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] global_control_actions: The actions of the predefined, custom, or override controls
        :param pulumi.Input[str] name: The name of the inspection profile.
        :param pulumi.Input[str] paranoia_level: OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionProfilePredefinedControlArgs']]]] predefined_controls: The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefined_controls` can be set by using the data source `data_source_zpa_predefined_controls` or by group using the data source `get_inspection_all_predefined_controls`.
        :param pulumi.Input[str] predefined_controls_version: The protocol for the AppProtection application
        :param pulumi.Input[str] zs_defined_control_choice: Indicates the user's choice for the ThreatLabZ Controls. Supported values: ALL and SPECIFIC
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InspectionProfileState.__new__(_InspectionProfileState)

        __props__.__dict__["associate_all_controls"] = associate_all_controls
        __props__.__dict__["controls_infos"] = controls_infos
        __props__.__dict__["custom_controls"] = custom_controls
        __props__.__dict__["description"] = description
        __props__.__dict__["global_control_actions"] = global_control_actions
        __props__.__dict__["incarnation_number"] = incarnation_number
        __props__.__dict__["name"] = name
        __props__.__dict__["paranoia_level"] = paranoia_level
        __props__.__dict__["predefined_controls"] = predefined_controls
        __props__.__dict__["predefined_controls_version"] = predefined_controls_version
        __props__.__dict__["zs_defined_control_choice"] = zs_defined_control_choice
        return InspectionProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="associateAllControls")
    def associate_all_controls(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "associate_all_controls")

    @property
    @pulumi.getter(name="controlsInfos")
    def controls_infos(self) -> pulumi.Output[Sequence['outputs.InspectionProfileControlsInfo']]:
        """
        (Optional) Types for custom controls
        """
        return pulumi.get(self, "controls_infos")

    @property
    @pulumi.getter(name="customControls")
    def custom_controls(self) -> pulumi.Output[Optional[Sequence['outputs.InspectionProfileCustomControl']]]:
        """
        (Optional) Types for custom controls
        """
        return pulumi.get(self, "custom_controls")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the inspection profile.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="globalControlActions")
    def global_control_actions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The actions of the predefined, custom, or override controls
        """
        return pulumi.get(self, "global_control_actions")

    @property
    @pulumi.getter(name="incarnationNumber")
    def incarnation_number(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "incarnation_number")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the inspection profile.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="paranoiaLevel")
    def paranoia_level(self) -> pulumi.Output[Optional[str]]:
        """
        OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        """
        return pulumi.get(self, "paranoia_level")

    @property
    @pulumi.getter(name="predefinedControls")
    def predefined_controls(self) -> pulumi.Output[Sequence['outputs.InspectionProfilePredefinedControl']]:
        """
        The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefined_controls` can be set by using the data source `data_source_zpa_predefined_controls` or by group using the data source `get_inspection_all_predefined_controls`.
        """
        return pulumi.get(self, "predefined_controls")

    @property
    @pulumi.getter(name="predefinedControlsVersion")
    def predefined_controls_version(self) -> pulumi.Output[Optional[str]]:
        """
        The protocol for the AppProtection application
        """
        return pulumi.get(self, "predefined_controls_version")

    @property
    @pulumi.getter(name="zsDefinedControlChoice")
    def zs_defined_control_choice(self) -> pulumi.Output[Optional[str]]:
        """
        Indicates the user's choice for the ThreatLabZ Controls. Supported values: ALL and SPECIFIC
        """
        return pulumi.get(self, "zs_defined_control_choice")

