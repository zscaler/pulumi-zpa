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

__all__ = ['CloudBrowserIsolationExternalProfileArgs', 'CloudBrowserIsolationExternalProfile']

@pulumi.input_type
class CloudBrowserIsolationExternalProfileArgs:
    def __init__(__self__, *,
                 banner_id: pulumi.Input[builtins.str],
                 certificate_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 debug_mode: Optional[pulumi.Input['CloudBrowserIsolationExternalProfileDebugModeArgs']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 security_controls: Optional[pulumi.Input['CloudBrowserIsolationExternalProfileSecurityControlsArgs']] = None,
                 user_experience: Optional[pulumi.Input['CloudBrowserIsolationExternalProfileUserExperienceArgs']] = None):
        """
        The set of arguments for constructing a CloudBrowserIsolationExternalProfile resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] certificate_ids: This field defines the list of certificate IDs.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] region_ids: This field defines the list of region IDs.
        """
        pulumi.set(__self__, "banner_id", banner_id)
        if certificate_ids is not None:
            pulumi.set(__self__, "certificate_ids", certificate_ids)
        if debug_mode is not None:
            pulumi.set(__self__, "debug_mode", debug_mode)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region_ids is not None:
            pulumi.set(__self__, "region_ids", region_ids)
        if security_controls is not None:
            pulumi.set(__self__, "security_controls", security_controls)
        if user_experience is not None:
            pulumi.set(__self__, "user_experience", user_experience)

    @property
    @pulumi.getter(name="bannerId")
    def banner_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "banner_id")

    @banner_id.setter
    def banner_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "banner_id", value)

    @property
    @pulumi.getter(name="certificateIds")
    def certificate_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        This field defines the list of certificate IDs.
        """
        return pulumi.get(self, "certificate_ids")

    @certificate_ids.setter
    def certificate_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "certificate_ids", value)

    @property
    @pulumi.getter(name="debugMode")
    def debug_mode(self) -> Optional[pulumi.Input['CloudBrowserIsolationExternalProfileDebugModeArgs']]:
        return pulumi.get(self, "debug_mode")

    @debug_mode.setter
    def debug_mode(self, value: Optional[pulumi.Input['CloudBrowserIsolationExternalProfileDebugModeArgs']]):
        pulumi.set(self, "debug_mode", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="regionIds")
    def region_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        This field defines the list of region IDs.
        """
        return pulumi.get(self, "region_ids")

    @region_ids.setter
    def region_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "region_ids", value)

    @property
    @pulumi.getter(name="securityControls")
    def security_controls(self) -> Optional[pulumi.Input['CloudBrowserIsolationExternalProfileSecurityControlsArgs']]:
        return pulumi.get(self, "security_controls")

    @security_controls.setter
    def security_controls(self, value: Optional[pulumi.Input['CloudBrowserIsolationExternalProfileSecurityControlsArgs']]):
        pulumi.set(self, "security_controls", value)

    @property
    @pulumi.getter(name="userExperience")
    def user_experience(self) -> Optional[pulumi.Input['CloudBrowserIsolationExternalProfileUserExperienceArgs']]:
        return pulumi.get(self, "user_experience")

    @user_experience.setter
    def user_experience(self, value: Optional[pulumi.Input['CloudBrowserIsolationExternalProfileUserExperienceArgs']]):
        pulumi.set(self, "user_experience", value)


@pulumi.input_type
class _CloudBrowserIsolationExternalProfileState:
    def __init__(__self__, *,
                 banner_id: Optional[pulumi.Input[builtins.str]] = None,
                 certificate_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 debug_mode: Optional[pulumi.Input['CloudBrowserIsolationExternalProfileDebugModeArgs']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 security_controls: Optional[pulumi.Input['CloudBrowserIsolationExternalProfileSecurityControlsArgs']] = None,
                 user_experience: Optional[pulumi.Input['CloudBrowserIsolationExternalProfileUserExperienceArgs']] = None):
        """
        Input properties used for looking up and filtering CloudBrowserIsolationExternalProfile resources.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] certificate_ids: This field defines the list of certificate IDs.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] region_ids: This field defines the list of region IDs.
        """
        if banner_id is not None:
            pulumi.set(__self__, "banner_id", banner_id)
        if certificate_ids is not None:
            pulumi.set(__self__, "certificate_ids", certificate_ids)
        if debug_mode is not None:
            pulumi.set(__self__, "debug_mode", debug_mode)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region_ids is not None:
            pulumi.set(__self__, "region_ids", region_ids)
        if security_controls is not None:
            pulumi.set(__self__, "security_controls", security_controls)
        if user_experience is not None:
            pulumi.set(__self__, "user_experience", user_experience)

    @property
    @pulumi.getter(name="bannerId")
    def banner_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "banner_id")

    @banner_id.setter
    def banner_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "banner_id", value)

    @property
    @pulumi.getter(name="certificateIds")
    def certificate_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        This field defines the list of certificate IDs.
        """
        return pulumi.get(self, "certificate_ids")

    @certificate_ids.setter
    def certificate_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "certificate_ids", value)

    @property
    @pulumi.getter(name="debugMode")
    def debug_mode(self) -> Optional[pulumi.Input['CloudBrowserIsolationExternalProfileDebugModeArgs']]:
        return pulumi.get(self, "debug_mode")

    @debug_mode.setter
    def debug_mode(self, value: Optional[pulumi.Input['CloudBrowserIsolationExternalProfileDebugModeArgs']]):
        pulumi.set(self, "debug_mode", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="regionIds")
    def region_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        This field defines the list of region IDs.
        """
        return pulumi.get(self, "region_ids")

    @region_ids.setter
    def region_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "region_ids", value)

    @property
    @pulumi.getter(name="securityControls")
    def security_controls(self) -> Optional[pulumi.Input['CloudBrowserIsolationExternalProfileSecurityControlsArgs']]:
        return pulumi.get(self, "security_controls")

    @security_controls.setter
    def security_controls(self, value: Optional[pulumi.Input['CloudBrowserIsolationExternalProfileSecurityControlsArgs']]):
        pulumi.set(self, "security_controls", value)

    @property
    @pulumi.getter(name="userExperience")
    def user_experience(self) -> Optional[pulumi.Input['CloudBrowserIsolationExternalProfileUserExperienceArgs']]:
        return pulumi.get(self, "user_experience")

    @user_experience.setter
    def user_experience(self, value: Optional[pulumi.Input['CloudBrowserIsolationExternalProfileUserExperienceArgs']]):
        pulumi.set(self, "user_experience", value)


@pulumi.type_token("zpa:index/cloudBrowserIsolationExternalProfile:CloudBrowserIsolationExternalProfile")
class CloudBrowserIsolationExternalProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 banner_id: Optional[pulumi.Input[builtins.str]] = None,
                 certificate_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 debug_mode: Optional[pulumi.Input[Union['CloudBrowserIsolationExternalProfileDebugModeArgs', 'CloudBrowserIsolationExternalProfileDebugModeArgsDict']]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 security_controls: Optional[pulumi.Input[Union['CloudBrowserIsolationExternalProfileSecurityControlsArgs', 'CloudBrowserIsolationExternalProfileSecurityControlsArgsDict']]] = None,
                 user_experience: Optional[pulumi.Input[Union['CloudBrowserIsolationExternalProfileUserExperienceArgs', 'CloudBrowserIsolationExternalProfileUserExperienceArgsDict']]] = None,
                 __props__=None):
        """
        * [Official documentation](https://help.zscaler.com/isolation/about-custom-root-certificates-cloud-browser-isolation)

        The **zpa_cloud_browser_isolation_external_profile** resource creates a Cloud Browser Isolation external profile. This resource can then be used in as part of `PolicyAccessIsolationRule` when the `action` attribute is set to `ISOLATE`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_zpa as zpa
        import zscaler_pulumi_zpa as zpa

        this_cloud_browser_isolation_banner = zpa.get_cloud_browser_isolation_banner(name="Default")
        singapore = zpa.get_cloud_browser_isolation_region(name="Singapore")
        frankfurt = zpa.get_cloud_browser_isolation_region(name="Frankfurt")
        this_cloud_browser_isolation_certificate = zpa.get_cloud_browser_isolation_certificate(name="Zscaler Root Certificate")
        this_cloud_browser_isolation_external_profile = zpa.CloudBrowserIsolationExternalProfile("thisCloudBrowserIsolationExternalProfile",
            description="CBI_Profile_Example",
            banner_id=this_cloud_browser_isolation_banner.id,
            region_ids=[singapore.id],
            certificate_ids=[this_cloud_browser_isolation_certificate.id],
            user_experience={
                "forward_to_zia": {
                    "enabled": True,
                    "organization_id": "***********",
                    "cloud_name": "<cloud_name>",
                    "pac_file_url": "https://pac.<cloud_name>/<cloud_name>/proxy.pac",
                },
                "browser_in_browser": True,
                "persist_isolation_bar": True,
                "translate": True,
                "session_persistence": True,
            },
            security_controls={
                "copy_paste": "all",
                "upload_download": "upstream",
                "document_viewer": True,
                "local_render": True,
                "allow_printing": True,
                "restrict_keystrokes": True,
                "flattened_pdf": True,
                "deep_link": {
                    "enabled": True,
                    "applications": [
                        "test1",
                        "test",
                    ],
                },
                "watermark": {
                    "enabled": True,
                    "show_user_id": True,
                    "show_timestamp": True,
                    "show_message": True,
                    "message": "Zscaler CBI",
                },
            },
            debug_mode={
                "allowed": True,
                "file_password": "***********",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] certificate_ids: This field defines the list of certificate IDs.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] region_ids: This field defines the list of region IDs.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CloudBrowserIsolationExternalProfileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        * [Official documentation](https://help.zscaler.com/isolation/about-custom-root-certificates-cloud-browser-isolation)

        The **zpa_cloud_browser_isolation_external_profile** resource creates a Cloud Browser Isolation external profile. This resource can then be used in as part of `PolicyAccessIsolationRule` when the `action` attribute is set to `ISOLATE`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_zpa as zpa
        import zscaler_pulumi_zpa as zpa

        this_cloud_browser_isolation_banner = zpa.get_cloud_browser_isolation_banner(name="Default")
        singapore = zpa.get_cloud_browser_isolation_region(name="Singapore")
        frankfurt = zpa.get_cloud_browser_isolation_region(name="Frankfurt")
        this_cloud_browser_isolation_certificate = zpa.get_cloud_browser_isolation_certificate(name="Zscaler Root Certificate")
        this_cloud_browser_isolation_external_profile = zpa.CloudBrowserIsolationExternalProfile("thisCloudBrowserIsolationExternalProfile",
            description="CBI_Profile_Example",
            banner_id=this_cloud_browser_isolation_banner.id,
            region_ids=[singapore.id],
            certificate_ids=[this_cloud_browser_isolation_certificate.id],
            user_experience={
                "forward_to_zia": {
                    "enabled": True,
                    "organization_id": "***********",
                    "cloud_name": "<cloud_name>",
                    "pac_file_url": "https://pac.<cloud_name>/<cloud_name>/proxy.pac",
                },
                "browser_in_browser": True,
                "persist_isolation_bar": True,
                "translate": True,
                "session_persistence": True,
            },
            security_controls={
                "copy_paste": "all",
                "upload_download": "upstream",
                "document_viewer": True,
                "local_render": True,
                "allow_printing": True,
                "restrict_keystrokes": True,
                "flattened_pdf": True,
                "deep_link": {
                    "enabled": True,
                    "applications": [
                        "test1",
                        "test",
                    ],
                },
                "watermark": {
                    "enabled": True,
                    "show_user_id": True,
                    "show_timestamp": True,
                    "show_message": True,
                    "message": "Zscaler CBI",
                },
            },
            debug_mode={
                "allowed": True,
                "file_password": "***********",
            })
        ```

        :param str resource_name: The name of the resource.
        :param CloudBrowserIsolationExternalProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CloudBrowserIsolationExternalProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 banner_id: Optional[pulumi.Input[builtins.str]] = None,
                 certificate_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 debug_mode: Optional[pulumi.Input[Union['CloudBrowserIsolationExternalProfileDebugModeArgs', 'CloudBrowserIsolationExternalProfileDebugModeArgsDict']]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 security_controls: Optional[pulumi.Input[Union['CloudBrowserIsolationExternalProfileSecurityControlsArgs', 'CloudBrowserIsolationExternalProfileSecurityControlsArgsDict']]] = None,
                 user_experience: Optional[pulumi.Input[Union['CloudBrowserIsolationExternalProfileUserExperienceArgs', 'CloudBrowserIsolationExternalProfileUserExperienceArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CloudBrowserIsolationExternalProfileArgs.__new__(CloudBrowserIsolationExternalProfileArgs)

            if banner_id is None and not opts.urn:
                raise TypeError("Missing required property 'banner_id'")
            __props__.__dict__["banner_id"] = banner_id
            __props__.__dict__["certificate_ids"] = certificate_ids
            __props__.__dict__["debug_mode"] = debug_mode
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["region_ids"] = region_ids
            __props__.__dict__["security_controls"] = security_controls
            __props__.__dict__["user_experience"] = user_experience
        super(CloudBrowserIsolationExternalProfile, __self__).__init__(
            'zpa:index/cloudBrowserIsolationExternalProfile:CloudBrowserIsolationExternalProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            banner_id: Optional[pulumi.Input[builtins.str]] = None,
            certificate_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            debug_mode: Optional[pulumi.Input[Union['CloudBrowserIsolationExternalProfileDebugModeArgs', 'CloudBrowserIsolationExternalProfileDebugModeArgsDict']]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            region_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            security_controls: Optional[pulumi.Input[Union['CloudBrowserIsolationExternalProfileSecurityControlsArgs', 'CloudBrowserIsolationExternalProfileSecurityControlsArgsDict']]] = None,
            user_experience: Optional[pulumi.Input[Union['CloudBrowserIsolationExternalProfileUserExperienceArgs', 'CloudBrowserIsolationExternalProfileUserExperienceArgsDict']]] = None) -> 'CloudBrowserIsolationExternalProfile':
        """
        Get an existing CloudBrowserIsolationExternalProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] certificate_ids: This field defines the list of certificate IDs.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] region_ids: This field defines the list of region IDs.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CloudBrowserIsolationExternalProfileState.__new__(_CloudBrowserIsolationExternalProfileState)

        __props__.__dict__["banner_id"] = banner_id
        __props__.__dict__["certificate_ids"] = certificate_ids
        __props__.__dict__["debug_mode"] = debug_mode
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["region_ids"] = region_ids
        __props__.__dict__["security_controls"] = security_controls
        __props__.__dict__["user_experience"] = user_experience
        return CloudBrowserIsolationExternalProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bannerId")
    def banner_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "banner_id")

    @property
    @pulumi.getter(name="certificateIds")
    def certificate_ids(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        This field defines the list of certificate IDs.
        """
        return pulumi.get(self, "certificate_ids")

    @property
    @pulumi.getter(name="debugMode")
    def debug_mode(self) -> pulumi.Output[Optional['outputs.CloudBrowserIsolationExternalProfileDebugMode']]:
        return pulumi.get(self, "debug_mode")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="regionIds")
    def region_ids(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        This field defines the list of region IDs.
        """
        return pulumi.get(self, "region_ids")

    @property
    @pulumi.getter(name="securityControls")
    def security_controls(self) -> pulumi.Output[Optional['outputs.CloudBrowserIsolationExternalProfileSecurityControls']]:
        return pulumi.get(self, "security_controls")

    @property
    @pulumi.getter(name="userExperience")
    def user_experience(self) -> pulumi.Output['outputs.CloudBrowserIsolationExternalProfileUserExperience']:
        return pulumi.get(self, "user_experience")

