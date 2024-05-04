# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetCloudBrowserIsolationBannerResult',
    'AwaitableGetCloudBrowserIsolationBannerResult',
    'get_cloud_browser_isolation_banner',
    'get_cloud_browser_isolation_banner_output',
]

@pulumi.output_type
class GetCloudBrowserIsolationBannerResult:
    """
    A collection of values returned by getCloudBrowserIsolationBanner.
    """
    def __init__(__self__, banner=None, id=None, is_default=None, logo=None, name=None, notification_text=None, notification_title=None, primary_color=None, text_color=None):
        if banner and not isinstance(banner, bool):
            raise TypeError("Expected argument 'banner' to be a bool")
        pulumi.set(__self__, "banner", banner)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_default and not isinstance(is_default, bool):
            raise TypeError("Expected argument 'is_default' to be a bool")
        pulumi.set(__self__, "is_default", is_default)
        if logo and not isinstance(logo, str):
            raise TypeError("Expected argument 'logo' to be a str")
        pulumi.set(__self__, "logo", logo)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if notification_text and not isinstance(notification_text, str):
            raise TypeError("Expected argument 'notification_text' to be a str")
        pulumi.set(__self__, "notification_text", notification_text)
        if notification_title and not isinstance(notification_title, str):
            raise TypeError("Expected argument 'notification_title' to be a str")
        pulumi.set(__self__, "notification_title", notification_title)
        if primary_color and not isinstance(primary_color, str):
            raise TypeError("Expected argument 'primary_color' to be a str")
        pulumi.set(__self__, "primary_color", primary_color)
        if text_color and not isinstance(text_color, str):
            raise TypeError("Expected argument 'text_color' to be a str")
        pulumi.set(__self__, "text_color", text_color)

    @property
    @pulumi.getter
    def banner(self) -> bool:
        return pulumi.get(self, "banner")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> bool:
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter
    def logo(self) -> str:
        return pulumi.get(self, "logo")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationText")
    def notification_text(self) -> str:
        return pulumi.get(self, "notification_text")

    @property
    @pulumi.getter(name="notificationTitle")
    def notification_title(self) -> str:
        return pulumi.get(self, "notification_title")

    @property
    @pulumi.getter(name="primaryColor")
    def primary_color(self) -> str:
        return pulumi.get(self, "primary_color")

    @property
    @pulumi.getter(name="textColor")
    def text_color(self) -> str:
        return pulumi.get(self, "text_color")


class AwaitableGetCloudBrowserIsolationBannerResult(GetCloudBrowserIsolationBannerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCloudBrowserIsolationBannerResult(
            banner=self.banner,
            id=self.id,
            is_default=self.is_default,
            logo=self.logo,
            name=self.name,
            notification_text=self.notification_text,
            notification_title=self.notification_title,
            primary_color=self.primary_color,
            text_color=self.text_color)


def get_cloud_browser_isolation_banner(id: Optional[str] = None,
                                       name: Optional[str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCloudBrowserIsolationBannerResult:
    """
    * [Official documentation](https://help.zscaler.com/isolation/adding-banner-theme-isolation-end-user-notification-zpa)

    Use the **zpa_cloud_browser_isolation_banner** data source to get information about Cloud Browser Isolation banner. This data source information is required as part of the attribute `banner_id` when creating an Cloud Browser Isolation External Profile ``CloudBrowserIsolationExternalProfile``

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_zpa as zpa

    this = zpa.get_cloud_browser_isolation_banner(name="Default")
    ```
    <!--End PulumiCodeChooser -->
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zpa:index/getCloudBrowserIsolationBanner:getCloudBrowserIsolationBanner', __args__, opts=opts, typ=GetCloudBrowserIsolationBannerResult).value

    return AwaitableGetCloudBrowserIsolationBannerResult(
        banner=pulumi.get(__ret__, 'banner'),
        id=pulumi.get(__ret__, 'id'),
        is_default=pulumi.get(__ret__, 'is_default'),
        logo=pulumi.get(__ret__, 'logo'),
        name=pulumi.get(__ret__, 'name'),
        notification_text=pulumi.get(__ret__, 'notification_text'),
        notification_title=pulumi.get(__ret__, 'notification_title'),
        primary_color=pulumi.get(__ret__, 'primary_color'),
        text_color=pulumi.get(__ret__, 'text_color'))


@_utilities.lift_output_func(get_cloud_browser_isolation_banner)
def get_cloud_browser_isolation_banner_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                                              name: Optional[pulumi.Input[Optional[str]]] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCloudBrowserIsolationBannerResult]:
    """
    * [Official documentation](https://help.zscaler.com/isolation/adding-banner-theme-isolation-end-user-notification-zpa)

    Use the **zpa_cloud_browser_isolation_banner** data source to get information about Cloud Browser Isolation banner. This data source information is required as part of the attribute `banner_id` when creating an Cloud Browser Isolation External Profile ``CloudBrowserIsolationExternalProfile``

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_zpa as zpa

    this = zpa.get_cloud_browser_isolation_banner(name="Default")
    ```
    <!--End PulumiCodeChooser -->
    """
    ...
