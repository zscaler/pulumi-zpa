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
    'GetApplicationSegmentByTypeResult',
    'AwaitableGetApplicationSegmentByTypeResult',
    'get_application_segment_by_type',
    'get_application_segment_by_type_output',
]

@pulumi.output_type
class GetApplicationSegmentByTypeResult:
    """
    A collection of values returned by getApplicationSegmentByType.
    """
    def __init__(__self__, app_id=None, application_port=None, application_protocol=None, application_type=None, certificate_id=None, certificate_name=None, domain=None, enabled=None, id=None, microtenant_id=None, microtenant_name=None, name=None):
        if app_id and not isinstance(app_id, str):
            raise TypeError("Expected argument 'app_id' to be a str")
        pulumi.set(__self__, "app_id", app_id)
        if application_port and not isinstance(application_port, str):
            raise TypeError("Expected argument 'application_port' to be a str")
        pulumi.set(__self__, "application_port", application_port)
        if application_protocol and not isinstance(application_protocol, str):
            raise TypeError("Expected argument 'application_protocol' to be a str")
        pulumi.set(__self__, "application_protocol", application_protocol)
        if application_type and not isinstance(application_type, str):
            raise TypeError("Expected argument 'application_type' to be a str")
        pulumi.set(__self__, "application_type", application_type)
        if certificate_id and not isinstance(certificate_id, str):
            raise TypeError("Expected argument 'certificate_id' to be a str")
        pulumi.set(__self__, "certificate_id", certificate_id)
        if certificate_name and not isinstance(certificate_name, str):
            raise TypeError("Expected argument 'certificate_name' to be a str")
        pulumi.set(__self__, "certificate_name", certificate_name)
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if microtenant_id and not isinstance(microtenant_id, str):
            raise TypeError("Expected argument 'microtenant_id' to be a str")
        pulumi.set(__self__, "microtenant_id", microtenant_id)
        if microtenant_name and not isinstance(microtenant_name, str):
            raise TypeError("Expected argument 'microtenant_name' to be a str")
        pulumi.set(__self__, "microtenant_name", microtenant_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> str:
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="applicationPort")
    def application_port(self) -> str:
        return pulumi.get(self, "application_port")

    @property
    @pulumi.getter(name="applicationProtocol")
    def application_protocol(self) -> str:
        return pulumi.get(self, "application_protocol")

    @property
    @pulumi.getter(name="applicationType")
    def application_type(self) -> str:
        return pulumi.get(self, "application_type")

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> str:
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter(name="certificateName")
    def certificate_name(self) -> str:
        return pulumi.get(self, "certificate_name")

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="microtenantId")
    def microtenant_id(self) -> Optional[str]:
        return pulumi.get(self, "microtenant_id")

    @property
    @pulumi.getter(name="microtenantName")
    def microtenant_name(self) -> str:
        return pulumi.get(self, "microtenant_name")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")


class AwaitableGetApplicationSegmentByTypeResult(GetApplicationSegmentByTypeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationSegmentByTypeResult(
            app_id=self.app_id,
            application_port=self.application_port,
            application_protocol=self.application_protocol,
            application_type=self.application_type,
            certificate_id=self.certificate_id,
            certificate_name=self.certificate_name,
            domain=self.domain,
            enabled=self.enabled,
            id=self.id,
            microtenant_id=self.microtenant_id,
            microtenant_name=self.microtenant_name,
            name=self.name)


def get_application_segment_by_type(application_type: Optional[str] = None,
                                    microtenant_id: Optional[str] = None,
                                    name: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationSegmentByTypeResult:
    """
    * [Official documentation](https://help.zscaler.com/zpa/about-applications)
    * [API documentation](https://help.zscaler.com/zpa/configuring-application-segments-using-api)

    Use the **zpa_application_segment_by_type** data source to get all configured Application Segments by Access Type (e.g., ``BROWSER_ACCESS``, ``INSPECT``, or ``SECURE_REMOTE_ACCESS``) for the specified customer.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zpa as zpa

    this = zpa.get_application_segment_by_type(application_type="SECURE_REMOTE_ACCESS")
    ```
    """
    __args__ = dict()
    __args__['applicationType'] = application_type
    __args__['microtenantId'] = microtenant_id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zpa:index/getApplicationSegmentByType:getApplicationSegmentByType', __args__, opts=opts, typ=GetApplicationSegmentByTypeResult).value

    return AwaitableGetApplicationSegmentByTypeResult(
        app_id=pulumi.get(__ret__, 'app_id'),
        application_port=pulumi.get(__ret__, 'application_port'),
        application_protocol=pulumi.get(__ret__, 'application_protocol'),
        application_type=pulumi.get(__ret__, 'application_type'),
        certificate_id=pulumi.get(__ret__, 'certificate_id'),
        certificate_name=pulumi.get(__ret__, 'certificate_name'),
        domain=pulumi.get(__ret__, 'domain'),
        enabled=pulumi.get(__ret__, 'enabled'),
        id=pulumi.get(__ret__, 'id'),
        microtenant_id=pulumi.get(__ret__, 'microtenant_id'),
        microtenant_name=pulumi.get(__ret__, 'microtenant_name'),
        name=pulumi.get(__ret__, 'name'))


@_utilities.lift_output_func(get_application_segment_by_type)
def get_application_segment_by_type_output(application_type: Optional[pulumi.Input[str]] = None,
                                           microtenant_id: Optional[pulumi.Input[Optional[str]]] = None,
                                           name: Optional[pulumi.Input[Optional[str]]] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApplicationSegmentByTypeResult]:
    """
    * [Official documentation](https://help.zscaler.com/zpa/about-applications)
    * [API documentation](https://help.zscaler.com/zpa/configuring-application-segments-using-api)

    Use the **zpa_application_segment_by_type** data source to get all configured Application Segments by Access Type (e.g., ``BROWSER_ACCESS``, ``INSPECT``, or ``SECURE_REMOTE_ACCESS``) for the specified customer.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zpa as zpa

    this = zpa.get_application_segment_by_type(application_type="SECURE_REMOTE_ACCESS")
    ```
    """
    ...