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

__all__ = ['InspectionCustomControlsArgs', 'InspectionCustomControls']

@pulumi.input_type
class InspectionCustomControlsArgs:
    def __init__(__self__, *,
                 default_action: pulumi.Input[str],
                 severity: pulumi.Input[str],
                 type: pulumi.Input[str],
                 action: Optional[pulumi.Input[str]] = None,
                 action_value: Optional[pulumi.Input[str]] = None,
                 associated_inspection_profile_names: Optional[pulumi.Input[Sequence[pulumi.Input['InspectionCustomControlsAssociatedInspectionProfileNameArgs']]]] = None,
                 control_number: Optional[pulumi.Input[str]] = None,
                 control_rule_json: Optional[pulumi.Input[str]] = None,
                 control_type: Optional[pulumi.Input[str]] = None,
                 default_action_value: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 paranoia_level: Optional[pulumi.Input[str]] = None,
                 protocol_type: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['InspectionCustomControlsRuleArgs']]]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InspectionCustomControls resource.
        :param pulumi.Input[str] default_action: The performed action
        :param pulumi.Input[str] severity: Severity of the control number
        :param pulumi.Input[str] type: Rules to be applied to the request or response type
        :param pulumi.Input[str] action: The performed action
        :param pulumi.Input[Sequence[pulumi.Input['InspectionCustomControlsAssociatedInspectionProfileNameArgs']]] associated_inspection_profile_names: Name of the inspection profile
        :param pulumi.Input[str] control_rule_json: The control rule in JSON format that has the conditions and type of control for the inspection control
        :param pulumi.Input[str] default_action_value: This is used to provide the redirect URL if the default action is set to REDIRECT
        :param pulumi.Input[str] description: Description of the custom control
        :param pulumi.Input[str] paranoia_level: OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        :param pulumi.Input[Sequence[pulumi.Input['InspectionCustomControlsRuleArgs']]] rules: Rules of the custom controls applied as conditions (JSON)
        """
        pulumi.set(__self__, "default_action", default_action)
        pulumi.set(__self__, "severity", severity)
        pulumi.set(__self__, "type", type)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if action_value is not None:
            pulumi.set(__self__, "action_value", action_value)
        if associated_inspection_profile_names is not None:
            pulumi.set(__self__, "associated_inspection_profile_names", associated_inspection_profile_names)
        if control_number is not None:
            pulumi.set(__self__, "control_number", control_number)
        if control_rule_json is not None:
            pulumi.set(__self__, "control_rule_json", control_rule_json)
        if control_type is not None:
            pulumi.set(__self__, "control_type", control_type)
        if default_action_value is not None:
            pulumi.set(__self__, "default_action_value", default_action_value)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if paranoia_level is not None:
            pulumi.set(__self__, "paranoia_level", paranoia_level)
        if protocol_type is not None:
            pulumi.set(__self__, "protocol_type", protocol_type)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> pulumi.Input[str]:
        """
        The performed action
        """
        return pulumi.get(self, "default_action")

    @default_action.setter
    def default_action(self, value: pulumi.Input[str]):
        pulumi.set(self, "default_action", value)

    @property
    @pulumi.getter
    def severity(self) -> pulumi.Input[str]:
        """
        Severity of the control number
        """
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: pulumi.Input[str]):
        pulumi.set(self, "severity", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Rules to be applied to the request or response type
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        The performed action
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="actionValue")
    def action_value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "action_value")

    @action_value.setter
    def action_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action_value", value)

    @property
    @pulumi.getter(name="associatedInspectionProfileNames")
    def associated_inspection_profile_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InspectionCustomControlsAssociatedInspectionProfileNameArgs']]]]:
        """
        Name of the inspection profile
        """
        return pulumi.get(self, "associated_inspection_profile_names")

    @associated_inspection_profile_names.setter
    def associated_inspection_profile_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InspectionCustomControlsAssociatedInspectionProfileNameArgs']]]]):
        pulumi.set(self, "associated_inspection_profile_names", value)

    @property
    @pulumi.getter(name="controlNumber")
    def control_number(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "control_number")

    @control_number.setter
    def control_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "control_number", value)

    @property
    @pulumi.getter(name="controlRuleJson")
    def control_rule_json(self) -> Optional[pulumi.Input[str]]:
        """
        The control rule in JSON format that has the conditions and type of control for the inspection control
        """
        return pulumi.get(self, "control_rule_json")

    @control_rule_json.setter
    def control_rule_json(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "control_rule_json", value)

    @property
    @pulumi.getter(name="controlType")
    def control_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "control_type")

    @control_type.setter
    def control_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "control_type", value)

    @property
    @pulumi.getter(name="defaultActionValue")
    def default_action_value(self) -> Optional[pulumi.Input[str]]:
        """
        This is used to provide the redirect URL if the default action is set to REDIRECT
        """
        return pulumi.get(self, "default_action_value")

    @default_action_value.setter
    def default_action_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_action_value", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the custom control
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
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
    @pulumi.getter(name="protocolType")
    def protocol_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "protocol_type")

    @protocol_type.setter
    def protocol_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol_type", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InspectionCustomControlsRuleArgs']]]]:
        """
        Rules of the custom controls applied as conditions (JSON)
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InspectionCustomControlsRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class _InspectionCustomControlsState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 action_value: Optional[pulumi.Input[str]] = None,
                 associated_inspection_profile_names: Optional[pulumi.Input[Sequence[pulumi.Input['InspectionCustomControlsAssociatedInspectionProfileNameArgs']]]] = None,
                 control_number: Optional[pulumi.Input[str]] = None,
                 control_rule_json: Optional[pulumi.Input[str]] = None,
                 control_type: Optional[pulumi.Input[str]] = None,
                 default_action: Optional[pulumi.Input[str]] = None,
                 default_action_value: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 paranoia_level: Optional[pulumi.Input[str]] = None,
                 protocol_type: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['InspectionCustomControlsRuleArgs']]]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InspectionCustomControls resources.
        :param pulumi.Input[str] action: The performed action
        :param pulumi.Input[Sequence[pulumi.Input['InspectionCustomControlsAssociatedInspectionProfileNameArgs']]] associated_inspection_profile_names: Name of the inspection profile
        :param pulumi.Input[str] control_rule_json: The control rule in JSON format that has the conditions and type of control for the inspection control
        :param pulumi.Input[str] default_action: The performed action
        :param pulumi.Input[str] default_action_value: This is used to provide the redirect URL if the default action is set to REDIRECT
        :param pulumi.Input[str] description: Description of the custom control
        :param pulumi.Input[str] paranoia_level: OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        :param pulumi.Input[Sequence[pulumi.Input['InspectionCustomControlsRuleArgs']]] rules: Rules of the custom controls applied as conditions (JSON)
        :param pulumi.Input[str] severity: Severity of the control number
        :param pulumi.Input[str] type: Rules to be applied to the request or response type
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if action_value is not None:
            pulumi.set(__self__, "action_value", action_value)
        if associated_inspection_profile_names is not None:
            pulumi.set(__self__, "associated_inspection_profile_names", associated_inspection_profile_names)
        if control_number is not None:
            pulumi.set(__self__, "control_number", control_number)
        if control_rule_json is not None:
            pulumi.set(__self__, "control_rule_json", control_rule_json)
        if control_type is not None:
            pulumi.set(__self__, "control_type", control_type)
        if default_action is not None:
            pulumi.set(__self__, "default_action", default_action)
        if default_action_value is not None:
            pulumi.set(__self__, "default_action_value", default_action_value)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if paranoia_level is not None:
            pulumi.set(__self__, "paranoia_level", paranoia_level)
        if protocol_type is not None:
            pulumi.set(__self__, "protocol_type", protocol_type)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if severity is not None:
            pulumi.set(__self__, "severity", severity)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        The performed action
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="actionValue")
    def action_value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "action_value")

    @action_value.setter
    def action_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action_value", value)

    @property
    @pulumi.getter(name="associatedInspectionProfileNames")
    def associated_inspection_profile_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InspectionCustomControlsAssociatedInspectionProfileNameArgs']]]]:
        """
        Name of the inspection profile
        """
        return pulumi.get(self, "associated_inspection_profile_names")

    @associated_inspection_profile_names.setter
    def associated_inspection_profile_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InspectionCustomControlsAssociatedInspectionProfileNameArgs']]]]):
        pulumi.set(self, "associated_inspection_profile_names", value)

    @property
    @pulumi.getter(name="controlNumber")
    def control_number(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "control_number")

    @control_number.setter
    def control_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "control_number", value)

    @property
    @pulumi.getter(name="controlRuleJson")
    def control_rule_json(self) -> Optional[pulumi.Input[str]]:
        """
        The control rule in JSON format that has the conditions and type of control for the inspection control
        """
        return pulumi.get(self, "control_rule_json")

    @control_rule_json.setter
    def control_rule_json(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "control_rule_json", value)

    @property
    @pulumi.getter(name="controlType")
    def control_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "control_type")

    @control_type.setter
    def control_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "control_type", value)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> Optional[pulumi.Input[str]]:
        """
        The performed action
        """
        return pulumi.get(self, "default_action")

    @default_action.setter
    def default_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_action", value)

    @property
    @pulumi.getter(name="defaultActionValue")
    def default_action_value(self) -> Optional[pulumi.Input[str]]:
        """
        This is used to provide the redirect URL if the default action is set to REDIRECT
        """
        return pulumi.get(self, "default_action_value")

    @default_action_value.setter
    def default_action_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_action_value", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the custom control
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
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
    @pulumi.getter(name="protocolType")
    def protocol_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "protocol_type")

    @protocol_type.setter
    def protocol_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol_type", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InspectionCustomControlsRuleArgs']]]]:
        """
        Rules of the custom controls applied as conditions (JSON)
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InspectionCustomControlsRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def severity(self) -> Optional[pulumi.Input[str]]:
        """
        Severity of the control number
        """
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "severity", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Rules to be applied to the request or response type
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class InspectionCustomControls(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 action_value: Optional[pulumi.Input[str]] = None,
                 associated_inspection_profile_names: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionCustomControlsAssociatedInspectionProfileNameArgs']]]]] = None,
                 control_number: Optional[pulumi.Input[str]] = None,
                 control_rule_json: Optional[pulumi.Input[str]] = None,
                 control_type: Optional[pulumi.Input[str]] = None,
                 default_action: Optional[pulumi.Input[str]] = None,
                 default_action_value: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 paranoia_level: Optional[pulumi.Input[str]] = None,
                 protocol_type: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionCustomControlsRuleArgs']]]]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a InspectionCustomControls resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The performed action
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionCustomControlsAssociatedInspectionProfileNameArgs']]]] associated_inspection_profile_names: Name of the inspection profile
        :param pulumi.Input[str] control_rule_json: The control rule in JSON format that has the conditions and type of control for the inspection control
        :param pulumi.Input[str] default_action: The performed action
        :param pulumi.Input[str] default_action_value: This is used to provide the redirect URL if the default action is set to REDIRECT
        :param pulumi.Input[str] description: Description of the custom control
        :param pulumi.Input[str] paranoia_level: OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionCustomControlsRuleArgs']]]] rules: Rules of the custom controls applied as conditions (JSON)
        :param pulumi.Input[str] severity: Severity of the control number
        :param pulumi.Input[str] type: Rules to be applied to the request or response type
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InspectionCustomControlsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a InspectionCustomControls resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param InspectionCustomControlsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InspectionCustomControlsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 action_value: Optional[pulumi.Input[str]] = None,
                 associated_inspection_profile_names: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionCustomControlsAssociatedInspectionProfileNameArgs']]]]] = None,
                 control_number: Optional[pulumi.Input[str]] = None,
                 control_rule_json: Optional[pulumi.Input[str]] = None,
                 control_type: Optional[pulumi.Input[str]] = None,
                 default_action: Optional[pulumi.Input[str]] = None,
                 default_action_value: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 paranoia_level: Optional[pulumi.Input[str]] = None,
                 protocol_type: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionCustomControlsRuleArgs']]]]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InspectionCustomControlsArgs.__new__(InspectionCustomControlsArgs)

            __props__.__dict__["action"] = action
            __props__.__dict__["action_value"] = action_value
            __props__.__dict__["associated_inspection_profile_names"] = associated_inspection_profile_names
            __props__.__dict__["control_number"] = control_number
            __props__.__dict__["control_rule_json"] = control_rule_json
            __props__.__dict__["control_type"] = control_type
            if default_action is None and not opts.urn:
                raise TypeError("Missing required property 'default_action'")
            __props__.__dict__["default_action"] = default_action
            __props__.__dict__["default_action_value"] = default_action_value
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["paranoia_level"] = paranoia_level
            __props__.__dict__["protocol_type"] = protocol_type
            __props__.__dict__["rules"] = rules
            if severity is None and not opts.urn:
                raise TypeError("Missing required property 'severity'")
            __props__.__dict__["severity"] = severity
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["version"] = version
        super(InspectionCustomControls, __self__).__init__(
            'zpa:index/inspectionCustomControls:InspectionCustomControls',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            action_value: Optional[pulumi.Input[str]] = None,
            associated_inspection_profile_names: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionCustomControlsAssociatedInspectionProfileNameArgs']]]]] = None,
            control_number: Optional[pulumi.Input[str]] = None,
            control_rule_json: Optional[pulumi.Input[str]] = None,
            control_type: Optional[pulumi.Input[str]] = None,
            default_action: Optional[pulumi.Input[str]] = None,
            default_action_value: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            paranoia_level: Optional[pulumi.Input[str]] = None,
            protocol_type: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionCustomControlsRuleArgs']]]]] = None,
            severity: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'InspectionCustomControls':
        """
        Get an existing InspectionCustomControls resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The performed action
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionCustomControlsAssociatedInspectionProfileNameArgs']]]] associated_inspection_profile_names: Name of the inspection profile
        :param pulumi.Input[str] control_rule_json: The control rule in JSON format that has the conditions and type of control for the inspection control
        :param pulumi.Input[str] default_action: The performed action
        :param pulumi.Input[str] default_action_value: This is used to provide the redirect URL if the default action is set to REDIRECT
        :param pulumi.Input[str] description: Description of the custom control
        :param pulumi.Input[str] paranoia_level: OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InspectionCustomControlsRuleArgs']]]] rules: Rules of the custom controls applied as conditions (JSON)
        :param pulumi.Input[str] severity: Severity of the control number
        :param pulumi.Input[str] type: Rules to be applied to the request or response type
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InspectionCustomControlsState.__new__(_InspectionCustomControlsState)

        __props__.__dict__["action"] = action
        __props__.__dict__["action_value"] = action_value
        __props__.__dict__["associated_inspection_profile_names"] = associated_inspection_profile_names
        __props__.__dict__["control_number"] = control_number
        __props__.__dict__["control_rule_json"] = control_rule_json
        __props__.__dict__["control_type"] = control_type
        __props__.__dict__["default_action"] = default_action
        __props__.__dict__["default_action_value"] = default_action_value
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["paranoia_level"] = paranoia_level
        __props__.__dict__["protocol_type"] = protocol_type
        __props__.__dict__["rules"] = rules
        __props__.__dict__["severity"] = severity
        __props__.__dict__["type"] = type
        __props__.__dict__["version"] = version
        return InspectionCustomControls(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        The performed action
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="actionValue")
    def action_value(self) -> pulumi.Output[str]:
        return pulumi.get(self, "action_value")

    @property
    @pulumi.getter(name="associatedInspectionProfileNames")
    def associated_inspection_profile_names(self) -> pulumi.Output[Sequence['outputs.InspectionCustomControlsAssociatedInspectionProfileName']]:
        """
        Name of the inspection profile
        """
        return pulumi.get(self, "associated_inspection_profile_names")

    @property
    @pulumi.getter(name="controlNumber")
    def control_number(self) -> pulumi.Output[str]:
        return pulumi.get(self, "control_number")

    @property
    @pulumi.getter(name="controlRuleJson")
    def control_rule_json(self) -> pulumi.Output[str]:
        """
        The control rule in JSON format that has the conditions and type of control for the inspection control
        """
        return pulumi.get(self, "control_rule_json")

    @property
    @pulumi.getter(name="controlType")
    def control_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "control_type")

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> pulumi.Output[str]:
        """
        The performed action
        """
        return pulumi.get(self, "default_action")

    @property
    @pulumi.getter(name="defaultActionValue")
    def default_action_value(self) -> pulumi.Output[str]:
        """
        This is used to provide the redirect URL if the default action is set to REDIRECT
        """
        return pulumi.get(self, "default_action_value")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description of the custom control
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="paranoiaLevel")
    def paranoia_level(self) -> pulumi.Output[str]:
        """
        OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        """
        return pulumi.get(self, "paranoia_level")

    @property
    @pulumi.getter(name="protocolType")
    def protocol_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "protocol_type")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.InspectionCustomControlsRule']]:
        """
        Rules of the custom controls applied as conditions (JSON)
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def severity(self) -> pulumi.Output[str]:
        """
        Severity of the control number
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Rules to be applied to the request or response type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        return pulumi.get(self, "version")
