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
    'GetCloudBrowserIsolationCertificateResult',
    'AwaitableGetCloudBrowserIsolationCertificateResult',
    'get_cloud_browser_isolation_certificate',
    'get_cloud_browser_isolation_certificate_output',
]

@pulumi.output_type
class GetCloudBrowserIsolationCertificateResult:
    """
    A collection of values returned by getCloudBrowserIsolationCertificate.
    """
    def __init__(__self__, id=None, is_default=None, name=None, pem=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_default and not isinstance(is_default, bool):
            raise TypeError("Expected argument 'is_default' to be a bool")
        pulumi.set(__self__, "is_default", is_default)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if pem and not isinstance(pem, str):
            raise TypeError("Expected argument 'pem' to be a str")
        pulumi.set(__self__, "pem", pem)

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
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def pem(self) -> str:
        return pulumi.get(self, "pem")


class AwaitableGetCloudBrowserIsolationCertificateResult(GetCloudBrowserIsolationCertificateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCloudBrowserIsolationCertificateResult(
            id=self.id,
            is_default=self.is_default,
            name=self.name,
            pem=self.pem)


def get_cloud_browser_isolation_certificate(id: Optional[str] = None,
                                            name: Optional[str] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCloudBrowserIsolationCertificateResult:
    """
    * [Official documentation](https://help.zscaler.com/isolation/adding-banner-theme-isolation-end-user-notification-zpa)

    Use the **zpa_cloud_browser_isolation_certificate** data source to get information about Cloud Browser Isolation Certificate. This data source information is required as part of the attribute `certificate_ids` when creating an Cloud Browser Isolation External Profile ``CloudBrowserIsolationExternalProfile``

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_zpa as zpa

    this = zpa.get_cloud_browser_isolation_certificate(name="Zscaler Root Certificate")
    ```
    <!--End PulumiCodeChooser -->
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zpa:index/getCloudBrowserIsolationCertificate:getCloudBrowserIsolationCertificate', __args__, opts=opts, typ=GetCloudBrowserIsolationCertificateResult).value

    return AwaitableGetCloudBrowserIsolationCertificateResult(
        id=pulumi.get(__ret__, 'id'),
        is_default=pulumi.get(__ret__, 'is_default'),
        name=pulumi.get(__ret__, 'name'),
        pem=pulumi.get(__ret__, 'pem'))


@_utilities.lift_output_func(get_cloud_browser_isolation_certificate)
def get_cloud_browser_isolation_certificate_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                                                   name: Optional[pulumi.Input[Optional[str]]] = None,
                                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCloudBrowserIsolationCertificateResult]:
    """
    * [Official documentation](https://help.zscaler.com/isolation/adding-banner-theme-isolation-end-user-notification-zpa)

    Use the **zpa_cloud_browser_isolation_certificate** data source to get information about Cloud Browser Isolation Certificate. This data source information is required as part of the attribute `certificate_ids` when creating an Cloud Browser Isolation External Profile ``CloudBrowserIsolationExternalProfile``

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_zpa as zpa

    this = zpa.get_cloud_browser_isolation_certificate(name="Zscaler Root Certificate")
    ```
    <!--End PulumiCodeChooser -->
    """
    ...
