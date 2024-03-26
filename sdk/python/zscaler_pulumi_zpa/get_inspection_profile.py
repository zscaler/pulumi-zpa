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

__all__ = [
    'GetInspectionProfileResult',
    'AwaitableGetInspectionProfileResult',
    'get_inspection_profile',
    'get_inspection_profile_output',
]

@pulumi.output_type
class GetInspectionProfileResult:
    """
    A collection of values returned by getInspectionProfile.
    """
    def __init__(__self__, common_global_override_actions_config=None, controls_infos=None, creation_time=None, custom_controls=None, description=None, global_control_actions=None, id=None, incarnation_number=None, modified_by=None, modified_time=None, name=None, paranoia_level=None, predefined_controls=None, predefined_controls_version=None, web_socket_controls=None):
        if common_global_override_actions_config and not isinstance(common_global_override_actions_config, dict):
            raise TypeError("Expected argument 'common_global_override_actions_config' to be a dict")
        pulumi.set(__self__, "common_global_override_actions_config", common_global_override_actions_config)
        if controls_infos and not isinstance(controls_infos, list):
            raise TypeError("Expected argument 'controls_infos' to be a list")
        pulumi.set(__self__, "controls_infos", controls_infos)
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if custom_controls and not isinstance(custom_controls, list):
            raise TypeError("Expected argument 'custom_controls' to be a list")
        pulumi.set(__self__, "custom_controls", custom_controls)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if global_control_actions and not isinstance(global_control_actions, list):
            raise TypeError("Expected argument 'global_control_actions' to be a list")
        pulumi.set(__self__, "global_control_actions", global_control_actions)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if incarnation_number and not isinstance(incarnation_number, str):
            raise TypeError("Expected argument 'incarnation_number' to be a str")
        pulumi.set(__self__, "incarnation_number", incarnation_number)
        if modified_by and not isinstance(modified_by, str):
            raise TypeError("Expected argument 'modified_by' to be a str")
        pulumi.set(__self__, "modified_by", modified_by)
        if modified_time and not isinstance(modified_time, str):
            raise TypeError("Expected argument 'modified_time' to be a str")
        pulumi.set(__self__, "modified_time", modified_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if paranoia_level and not isinstance(paranoia_level, str):
            raise TypeError("Expected argument 'paranoia_level' to be a str")
        pulumi.set(__self__, "paranoia_level", paranoia_level)
        if predefined_controls and not isinstance(predefined_controls, list):
            raise TypeError("Expected argument 'predefined_controls' to be a list")
        pulumi.set(__self__, "predefined_controls", predefined_controls)
        if predefined_controls_version and not isinstance(predefined_controls_version, str):
            raise TypeError("Expected argument 'predefined_controls_version' to be a str")
        pulumi.set(__self__, "predefined_controls_version", predefined_controls_version)
        if web_socket_controls and not isinstance(web_socket_controls, list):
            raise TypeError("Expected argument 'web_socket_controls' to be a list")
        pulumi.set(__self__, "web_socket_controls", web_socket_controls)

    @property
    @pulumi.getter(name="commonGlobalOverrideActionsConfig")
    def common_global_override_actions_config(self) -> Mapping[str, str]:
        """
        (string)
        """
        return pulumi.get(self, "common_global_override_actions_config")

    @property
    @pulumi.getter(name="controlsInfos")
    def controls_infos(self) -> Sequence['outputs.GetInspectionProfileControlsInfoResult']:
        """
        (string) Types for custom controls
        """
        return pulumi.get(self, "controls_infos")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="customControls")
    def custom_controls(self) -> Sequence['outputs.GetInspectionProfileCustomControlResult']:
        """
        (string) Types for custom controls
        """
        return pulumi.get(self, "custom_controls")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        (string) Description of the inspection profile.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="globalControlActions")
    def global_control_actions(self) -> Sequence[str]:
        return pulumi.get(self, "global_control_actions")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        (string) ID of the predefined control
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="incarnationNumber")
    def incarnation_number(self) -> str:
        return pulumi.get(self, "incarnation_number")

    @property
    @pulumi.getter(name="modifiedBy")
    def modified_by(self) -> str:
        return pulumi.get(self, "modified_by")

    @property
    @pulumi.getter(name="modifiedTime")
    def modified_time(self) -> str:
        return pulumi.get(self, "modified_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        (string)
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="paranoiaLevel")
    def paranoia_level(self) -> str:
        """
        (string) OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        """
        return pulumi.get(self, "paranoia_level")

    @property
    @pulumi.getter(name="predefinedControls")
    def predefined_controls(self) -> Sequence['outputs.GetInspectionProfilePredefinedControlResult']:
        """
        (string) The predefined controls
        """
        return pulumi.get(self, "predefined_controls")

    @property
    @pulumi.getter(name="predefinedControlsVersion")
    def predefined_controls_version(self) -> str:
        return pulumi.get(self, "predefined_controls_version")

    @property
    @pulumi.getter(name="webSocketControls")
    def web_socket_controls(self) -> Sequence['outputs.GetInspectionProfileWebSocketControlResult']:
        """
        (string)
        """
        return pulumi.get(self, "web_socket_controls")


class AwaitableGetInspectionProfileResult(GetInspectionProfileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInspectionProfileResult(
            common_global_override_actions_config=self.common_global_override_actions_config,
            controls_infos=self.controls_infos,
            creation_time=self.creation_time,
            custom_controls=self.custom_controls,
            description=self.description,
            global_control_actions=self.global_control_actions,
            id=self.id,
            incarnation_number=self.incarnation_number,
            modified_by=self.modified_by,
            modified_time=self.modified_time,
            name=self.name,
            paranoia_level=self.paranoia_level,
            predefined_controls=self.predefined_controls,
            predefined_controls_version=self.predefined_controls_version,
            web_socket_controls=self.web_socket_controls)


def get_inspection_profile(id: Optional[str] = None,
                           name: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInspectionProfileResult:
    """
    Use the **zpa_inspection_profile** data source to get information about an inspection profile in the Zscaler Private Access cloud. This resource can then be referenced in an inspection custom control resource.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_zpa as zpa

    this = zpa.get_inspection_profile(name="Example")
    ```
    <!--End PulumiCodeChooser -->


    :param str id: This field defines the id of the inspection profile.
    :param str name: This field defines the name of the inspection profile.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zpa:index/getInspectionProfile:getInspectionProfile', __args__, opts=opts, typ=GetInspectionProfileResult).value

    return AwaitableGetInspectionProfileResult(
        common_global_override_actions_config=pulumi.get(__ret__, 'common_global_override_actions_config'),
        controls_infos=pulumi.get(__ret__, 'controls_infos'),
        creation_time=pulumi.get(__ret__, 'creation_time'),
        custom_controls=pulumi.get(__ret__, 'custom_controls'),
        description=pulumi.get(__ret__, 'description'),
        global_control_actions=pulumi.get(__ret__, 'global_control_actions'),
        id=pulumi.get(__ret__, 'id'),
        incarnation_number=pulumi.get(__ret__, 'incarnation_number'),
        modified_by=pulumi.get(__ret__, 'modified_by'),
        modified_time=pulumi.get(__ret__, 'modified_time'),
        name=pulumi.get(__ret__, 'name'),
        paranoia_level=pulumi.get(__ret__, 'paranoia_level'),
        predefined_controls=pulumi.get(__ret__, 'predefined_controls'),
        predefined_controls_version=pulumi.get(__ret__, 'predefined_controls_version'),
        web_socket_controls=pulumi.get(__ret__, 'web_socket_controls'))


@_utilities.lift_output_func(get_inspection_profile)
def get_inspection_profile_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                                  name: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInspectionProfileResult]:
    """
    Use the **zpa_inspection_profile** data source to get information about an inspection profile in the Zscaler Private Access cloud. This resource can then be referenced in an inspection custom control resource.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_zpa as zpa

    this = zpa.get_inspection_profile(name="Example")
    ```
    <!--End PulumiCodeChooser -->


    :param str id: This field defines the id of the inspection profile.
    :param str name: This field defines the name of the inspection profile.
    """
    ...
