# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'InspectionCustomControlsAssociatedInspectionProfileNameArgs',
    'InspectionCustomControlsRuleArgs',
    'InspectionCustomControlsRuleConditionsArgs',
    'InspectionProfileControlsInfoArgs',
    'InspectionProfileCustomControlArgs',
    'InspectionProfilePredefinedControlArgs',
    'InspectionProfileWebSocketControlArgs',
]

@pulumi.input_type
class InspectionCustomControlsAssociatedInspectionProfileNameArgs:
    def __init__(__self__, *,
                 ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        if ids is not None:
            pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "ids")

    @ids.setter
    def ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ids", value)


@pulumi.input_type
class InspectionCustomControlsRuleArgs:
    def __init__(__self__, *,
                 conditions: Optional[pulumi.Input['InspectionCustomControlsRuleConditionsArgs']] = None,
                 names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if names is not None:
            pulumi.set(__self__, "names", names)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input['InspectionCustomControlsRuleConditionsArgs']]:
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input['InspectionCustomControlsRuleConditionsArgs']]):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter
    def names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "names")

    @names.setter
    def names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "names", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class InspectionCustomControlsRuleConditionsArgs:
    def __init__(__self__, *,
                 lhs: Optional[pulumi.Input[str]] = None,
                 op: Optional[pulumi.Input[str]] = None,
                 rhs: Optional[pulumi.Input[str]] = None):
        if lhs is not None:
            pulumi.set(__self__, "lhs", lhs)
        if op is not None:
            pulumi.set(__self__, "op", op)
        if rhs is not None:
            pulumi.set(__self__, "rhs", rhs)

    @property
    @pulumi.getter
    def lhs(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "lhs")

    @lhs.setter
    def lhs(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lhs", value)

    @property
    @pulumi.getter
    def op(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "op")

    @op.setter
    def op(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "op", value)

    @property
    @pulumi.getter
    def rhs(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "rhs")

    @rhs.setter
    def rhs(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rhs", value)


@pulumi.input_type
class InspectionProfileControlsInfoArgs:
    def __init__(__self__, *,
                 control_type: Optional[pulumi.Input[str]] = None,
                 count: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] control_type: (string) Control types. Supported Values: `WEBSOCKET_PREDEFINED`, `WEBSOCKET_CUSTOM`, `CUSTOM`, `PREDEFINED`, `ZSCALER`
        :param pulumi.Input[str] count: (Optional) Control information counts `Long`
        """
        if control_type is not None:
            pulumi.set(__self__, "control_type", control_type)
        if count is not None:
            pulumi.set(__self__, "count", count)

    @property
    @pulumi.getter(name="controlType")
    def control_type(self) -> Optional[pulumi.Input[str]]:
        """
        (string) Control types. Supported Values: `WEBSOCKET_PREDEFINED`, `WEBSOCKET_CUSTOM`, `CUSTOM`, `PREDEFINED`, `ZSCALER`
        """
        return pulumi.get(self, "control_type")

    @control_type.setter
    def control_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "control_type", value)

    @property
    @pulumi.getter
    def count(self) -> Optional[pulumi.Input[str]]:
        """
        (Optional) Control information counts `Long`
        """
        return pulumi.get(self, "count")

    @count.setter
    def count(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "count", value)


@pulumi.input_type
class InspectionProfileCustomControlArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 action: Optional[pulumi.Input[str]] = None,
                 action_value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: ID of the predefined control
        :param pulumi.Input[str] action: The action of the predefined control. Supported values: `PASS`, `BLOCK` and `REDIRECT`
        :param pulumi.Input[str] action_value: Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to `REDIRECT`.
        """
        pulumi.set(__self__, "id", id)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if action_value is not None:
            pulumi.set(__self__, "action_value", action_value)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        ID of the predefined control
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        The action of the predefined control. Supported values: `PASS`, `BLOCK` and `REDIRECT`
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="actionValue")
    def action_value(self) -> Optional[pulumi.Input[str]]:
        """
        Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to `REDIRECT`.
        """
        return pulumi.get(self, "action_value")

    @action_value.setter
    def action_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action_value", value)


@pulumi.input_type
class InspectionProfilePredefinedControlArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 id: pulumi.Input[str],
                 action_value: Optional[pulumi.Input[str]] = None,
                 control_type: Optional[pulumi.Input[str]] = None,
                 protocol_type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] action: The action of the predefined control. Supported values: `PASS`, `BLOCK` and `REDIRECT`
        :param pulumi.Input[str] id: ID of the predefined control
        :param pulumi.Input[str] action_value: Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to `REDIRECT`.
        :param pulumi.Input[str] control_type: (string) Control types. Supported Values: `WEBSOCKET_PREDEFINED`, `WEBSOCKET_CUSTOM`, `CUSTOM`, `PREDEFINED`, `ZSCALER`
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "id", id)
        if action_value is not None:
            pulumi.set(__self__, "action_value", action_value)
        if control_type is not None:
            pulumi.set(__self__, "control_type", control_type)
        if protocol_type is not None:
            pulumi.set(__self__, "protocol_type", protocol_type)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        """
        The action of the predefined control. Supported values: `PASS`, `BLOCK` and `REDIRECT`
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        ID of the predefined control
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="actionValue")
    def action_value(self) -> Optional[pulumi.Input[str]]:
        """
        Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to `REDIRECT`.
        """
        return pulumi.get(self, "action_value")

    @action_value.setter
    def action_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action_value", value)

    @property
    @pulumi.getter(name="controlType")
    def control_type(self) -> Optional[pulumi.Input[str]]:
        """
        (string) Control types. Supported Values: `WEBSOCKET_PREDEFINED`, `WEBSOCKET_CUSTOM`, `CUSTOM`, `PREDEFINED`, `ZSCALER`
        """
        return pulumi.get(self, "control_type")

    @control_type.setter
    def control_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "control_type", value)

    @property
    @pulumi.getter(name="protocolType")
    def protocol_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "protocol_type")

    @protocol_type.setter
    def protocol_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol_type", value)


@pulumi.input_type
class InspectionProfileWebSocketControlArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 id: pulumi.Input[str],
                 action_value: Optional[pulumi.Input[str]] = None,
                 control_type: Optional[pulumi.Input[str]] = None,
                 protocol_type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] action: The action of the predefined control. Supported values: `PASS`, `BLOCK` and `REDIRECT`
        :param pulumi.Input[str] id: ID of the predefined control
        :param pulumi.Input[str] action_value: Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to `REDIRECT`.
        :param pulumi.Input[str] control_type: (string) Control types. Supported Values: `WEBSOCKET_PREDEFINED`, `WEBSOCKET_CUSTOM`, `CUSTOM`, `PREDEFINED`, `ZSCALER`
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "id", id)
        if action_value is not None:
            pulumi.set(__self__, "action_value", action_value)
        if control_type is not None:
            pulumi.set(__self__, "control_type", control_type)
        if protocol_type is not None:
            pulumi.set(__self__, "protocol_type", protocol_type)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        """
        The action of the predefined control. Supported values: `PASS`, `BLOCK` and `REDIRECT`
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        ID of the predefined control
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="actionValue")
    def action_value(self) -> Optional[pulumi.Input[str]]:
        """
        Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to `REDIRECT`.
        """
        return pulumi.get(self, "action_value")

    @action_value.setter
    def action_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action_value", value)

    @property
    @pulumi.getter(name="controlType")
    def control_type(self) -> Optional[pulumi.Input[str]]:
        """
        (string) Control types. Supported Values: `WEBSOCKET_PREDEFINED`, `WEBSOCKET_CUSTOM`, `CUSTOM`, `PREDEFINED`, `ZSCALER`
        """
        return pulumi.get(self, "control_type")

    @control_type.setter
    def control_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "control_type", value)

    @property
    @pulumi.getter(name="protocolType")
    def protocol_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "protocol_type")

    @protocol_type.setter
    def protocol_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol_type", value)


