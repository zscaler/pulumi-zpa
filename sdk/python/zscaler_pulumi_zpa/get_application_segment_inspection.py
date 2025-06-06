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

__all__ = [
    'GetApplicationSegmentInspectionResult',
    'AwaitableGetApplicationSegmentInspectionResult',
    'get_application_segment_inspection',
    'get_application_segment_inspection_output',
]

@pulumi.output_type
class GetApplicationSegmentInspectionResult:
    """
    A collection of values returned by getApplicationSegmentInspection.
    """
    def __init__(__self__, bypass_type=None, creation_time=None, description=None, domain_names=None, double_encrypt=None, enabled=None, health_check_type=None, health_reporting=None, icmp_access_type=None, id=None, inspection_apps=None, ip_anchored=None, is_cname_enabled=None, modified_by=None, modified_time=None, name=None, passive_health_enabled=None, segment_group_id=None, segment_group_name=None, select_connector_close_to_app=None, server_groups=None, tcp_port_range=None, tcp_port_ranges=None, udp_port_range=None, udp_port_ranges=None):
        if bypass_type and not isinstance(bypass_type, str):
            raise TypeError("Expected argument 'bypass_type' to be a str")
        pulumi.set(__self__, "bypass_type", bypass_type)
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if domain_names and not isinstance(domain_names, list):
            raise TypeError("Expected argument 'domain_names' to be a list")
        pulumi.set(__self__, "domain_names", domain_names)
        if double_encrypt and not isinstance(double_encrypt, bool):
            raise TypeError("Expected argument 'double_encrypt' to be a bool")
        pulumi.set(__self__, "double_encrypt", double_encrypt)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if health_check_type and not isinstance(health_check_type, str):
            raise TypeError("Expected argument 'health_check_type' to be a str")
        pulumi.set(__self__, "health_check_type", health_check_type)
        if health_reporting and not isinstance(health_reporting, str):
            raise TypeError("Expected argument 'health_reporting' to be a str")
        pulumi.set(__self__, "health_reporting", health_reporting)
        if icmp_access_type and not isinstance(icmp_access_type, str):
            raise TypeError("Expected argument 'icmp_access_type' to be a str")
        pulumi.set(__self__, "icmp_access_type", icmp_access_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if inspection_apps and not isinstance(inspection_apps, list):
            raise TypeError("Expected argument 'inspection_apps' to be a list")
        pulumi.set(__self__, "inspection_apps", inspection_apps)
        if ip_anchored and not isinstance(ip_anchored, bool):
            raise TypeError("Expected argument 'ip_anchored' to be a bool")
        pulumi.set(__self__, "ip_anchored", ip_anchored)
        if is_cname_enabled and not isinstance(is_cname_enabled, bool):
            raise TypeError("Expected argument 'is_cname_enabled' to be a bool")
        pulumi.set(__self__, "is_cname_enabled", is_cname_enabled)
        if modified_by and not isinstance(modified_by, str):
            raise TypeError("Expected argument 'modified_by' to be a str")
        pulumi.set(__self__, "modified_by", modified_by)
        if modified_time and not isinstance(modified_time, str):
            raise TypeError("Expected argument 'modified_time' to be a str")
        pulumi.set(__self__, "modified_time", modified_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if passive_health_enabled and not isinstance(passive_health_enabled, bool):
            raise TypeError("Expected argument 'passive_health_enabled' to be a bool")
        pulumi.set(__self__, "passive_health_enabled", passive_health_enabled)
        if segment_group_id and not isinstance(segment_group_id, str):
            raise TypeError("Expected argument 'segment_group_id' to be a str")
        pulumi.set(__self__, "segment_group_id", segment_group_id)
        if segment_group_name and not isinstance(segment_group_name, str):
            raise TypeError("Expected argument 'segment_group_name' to be a str")
        pulumi.set(__self__, "segment_group_name", segment_group_name)
        if select_connector_close_to_app and not isinstance(select_connector_close_to_app, bool):
            raise TypeError("Expected argument 'select_connector_close_to_app' to be a bool")
        pulumi.set(__self__, "select_connector_close_to_app", select_connector_close_to_app)
        if server_groups and not isinstance(server_groups, list):
            raise TypeError("Expected argument 'server_groups' to be a list")
        pulumi.set(__self__, "server_groups", server_groups)
        if tcp_port_range and not isinstance(tcp_port_range, list):
            raise TypeError("Expected argument 'tcp_port_range' to be a list")
        pulumi.set(__self__, "tcp_port_range", tcp_port_range)
        if tcp_port_ranges and not isinstance(tcp_port_ranges, list):
            raise TypeError("Expected argument 'tcp_port_ranges' to be a list")
        pulumi.set(__self__, "tcp_port_ranges", tcp_port_ranges)
        if udp_port_range and not isinstance(udp_port_range, list):
            raise TypeError("Expected argument 'udp_port_range' to be a list")
        pulumi.set(__self__, "udp_port_range", udp_port_range)
        if udp_port_ranges and not isinstance(udp_port_ranges, list):
            raise TypeError("Expected argument 'udp_port_ranges' to be a list")
        pulumi.set(__self__, "udp_port_ranges", udp_port_ranges)

    @property
    @pulumi.getter(name="bypassType")
    def bypass_type(self) -> builtins.str:
        return pulumi.get(self, "bypass_type")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> builtins.str:
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="domainNames")
    def domain_names(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "domain_names")

    @property
    @pulumi.getter(name="doubleEncrypt")
    def double_encrypt(self) -> builtins.bool:
        return pulumi.get(self, "double_encrypt")

    @property
    @pulumi.getter
    def enabled(self) -> builtins.bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="healthCheckType")
    def health_check_type(self) -> builtins.str:
        return pulumi.get(self, "health_check_type")

    @property
    @pulumi.getter(name="healthReporting")
    def health_reporting(self) -> builtins.str:
        return pulumi.get(self, "health_reporting")

    @property
    @pulumi.getter(name="icmpAccessType")
    def icmp_access_type(self) -> builtins.str:
        return pulumi.get(self, "icmp_access_type")

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="inspectionApps")
    def inspection_apps(self) -> Sequence['outputs.GetApplicationSegmentInspectionInspectionAppResult']:
        return pulumi.get(self, "inspection_apps")

    @property
    @pulumi.getter(name="ipAnchored")
    def ip_anchored(self) -> builtins.bool:
        return pulumi.get(self, "ip_anchored")

    @property
    @pulumi.getter(name="isCnameEnabled")
    def is_cname_enabled(self) -> builtins.bool:
        return pulumi.get(self, "is_cname_enabled")

    @property
    @pulumi.getter(name="modifiedBy")
    def modified_by(self) -> builtins.str:
        return pulumi.get(self, "modified_by")

    @property
    @pulumi.getter(name="modifiedTime")
    def modified_time(self) -> builtins.str:
        return pulumi.get(self, "modified_time")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="passiveHealthEnabled")
    def passive_health_enabled(self) -> builtins.bool:
        return pulumi.get(self, "passive_health_enabled")

    @property
    @pulumi.getter(name="segmentGroupId")
    def segment_group_id(self) -> builtins.str:
        return pulumi.get(self, "segment_group_id")

    @property
    @pulumi.getter(name="segmentGroupName")
    def segment_group_name(self) -> builtins.str:
        return pulumi.get(self, "segment_group_name")

    @property
    @pulumi.getter(name="selectConnectorCloseToApp")
    def select_connector_close_to_app(self) -> builtins.bool:
        return pulumi.get(self, "select_connector_close_to_app")

    @property
    @pulumi.getter(name="serverGroups")
    def server_groups(self) -> Sequence['outputs.GetApplicationSegmentInspectionServerGroupResult']:
        return pulumi.get(self, "server_groups")

    @property
    @pulumi.getter(name="tcpPortRange")
    def tcp_port_range(self) -> Sequence['outputs.GetApplicationSegmentInspectionTcpPortRangeResult']:
        return pulumi.get(self, "tcp_port_range")

    @property
    @pulumi.getter(name="tcpPortRanges")
    def tcp_port_ranges(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "tcp_port_ranges")

    @property
    @pulumi.getter(name="udpPortRange")
    def udp_port_range(self) -> Sequence['outputs.GetApplicationSegmentInspectionUdpPortRangeResult']:
        return pulumi.get(self, "udp_port_range")

    @property
    @pulumi.getter(name="udpPortRanges")
    def udp_port_ranges(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "udp_port_ranges")


class AwaitableGetApplicationSegmentInspectionResult(GetApplicationSegmentInspectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationSegmentInspectionResult(
            bypass_type=self.bypass_type,
            creation_time=self.creation_time,
            description=self.description,
            domain_names=self.domain_names,
            double_encrypt=self.double_encrypt,
            enabled=self.enabled,
            health_check_type=self.health_check_type,
            health_reporting=self.health_reporting,
            icmp_access_type=self.icmp_access_type,
            id=self.id,
            inspection_apps=self.inspection_apps,
            ip_anchored=self.ip_anchored,
            is_cname_enabled=self.is_cname_enabled,
            modified_by=self.modified_by,
            modified_time=self.modified_time,
            name=self.name,
            passive_health_enabled=self.passive_health_enabled,
            segment_group_id=self.segment_group_id,
            segment_group_name=self.segment_group_name,
            select_connector_close_to_app=self.select_connector_close_to_app,
            server_groups=self.server_groups,
            tcp_port_range=self.tcp_port_range,
            tcp_port_ranges=self.tcp_port_ranges,
            udp_port_range=self.udp_port_range,
            udp_port_ranges=self.udp_port_ranges)


def get_application_segment_inspection(id: Optional[builtins.str] = None,
                                       name: Optional[builtins.str] = None,
                                       tcp_port_range: Optional[Sequence[Union['GetApplicationSegmentInspectionTcpPortRangeArgs', 'GetApplicationSegmentInspectionTcpPortRangeArgsDict']]] = None,
                                       udp_port_range: Optional[Sequence[Union['GetApplicationSegmentInspectionUdpPortRangeArgs', 'GetApplicationSegmentInspectionUdpPortRangeArgsDict']]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationSegmentInspectionResult:
    """
    Use the **zpa_application_segment_inspection** data source to get information about an inspection application segment in the Zscaler Private Access cloud. This resource can then be referenced in a ZPA access inspection policy. This resource supports ZPA Inspection for both `HTTP` and `HTTPS`.

    **NOTE:** To ensure consistent search results across data sources, please avoid using multiple spaces or special characters in your search queries.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zpa as zpa

    this = zpa.get_application_segment_inspection(name="ZPA_Inspection_Example")
    ```

    ```python
    import pulumi
    import pulumi_zpa as zpa

    this = zpa.get_application_segment_inspection(id="123456789")
    ```
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['tcpPortRange'] = tcp_port_range
    __args__['udpPortRange'] = udp_port_range
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zpa:index/getApplicationSegmentInspection:getApplicationSegmentInspection', __args__, opts=opts, typ=GetApplicationSegmentInspectionResult).value

    return AwaitableGetApplicationSegmentInspectionResult(
        bypass_type=pulumi.get(__ret__, 'bypass_type'),
        creation_time=pulumi.get(__ret__, 'creation_time'),
        description=pulumi.get(__ret__, 'description'),
        domain_names=pulumi.get(__ret__, 'domain_names'),
        double_encrypt=pulumi.get(__ret__, 'double_encrypt'),
        enabled=pulumi.get(__ret__, 'enabled'),
        health_check_type=pulumi.get(__ret__, 'health_check_type'),
        health_reporting=pulumi.get(__ret__, 'health_reporting'),
        icmp_access_type=pulumi.get(__ret__, 'icmp_access_type'),
        id=pulumi.get(__ret__, 'id'),
        inspection_apps=pulumi.get(__ret__, 'inspection_apps'),
        ip_anchored=pulumi.get(__ret__, 'ip_anchored'),
        is_cname_enabled=pulumi.get(__ret__, 'is_cname_enabled'),
        modified_by=pulumi.get(__ret__, 'modified_by'),
        modified_time=pulumi.get(__ret__, 'modified_time'),
        name=pulumi.get(__ret__, 'name'),
        passive_health_enabled=pulumi.get(__ret__, 'passive_health_enabled'),
        segment_group_id=pulumi.get(__ret__, 'segment_group_id'),
        segment_group_name=pulumi.get(__ret__, 'segment_group_name'),
        select_connector_close_to_app=pulumi.get(__ret__, 'select_connector_close_to_app'),
        server_groups=pulumi.get(__ret__, 'server_groups'),
        tcp_port_range=pulumi.get(__ret__, 'tcp_port_range'),
        tcp_port_ranges=pulumi.get(__ret__, 'tcp_port_ranges'),
        udp_port_range=pulumi.get(__ret__, 'udp_port_range'),
        udp_port_ranges=pulumi.get(__ret__, 'udp_port_ranges'))
def get_application_segment_inspection_output(id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                              name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                              tcp_port_range: Optional[pulumi.Input[Optional[Sequence[Union['GetApplicationSegmentInspectionTcpPortRangeArgs', 'GetApplicationSegmentInspectionTcpPortRangeArgsDict']]]]] = None,
                                              udp_port_range: Optional[pulumi.Input[Optional[Sequence[Union['GetApplicationSegmentInspectionUdpPortRangeArgs', 'GetApplicationSegmentInspectionUdpPortRangeArgsDict']]]]] = None,
                                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetApplicationSegmentInspectionResult]:
    """
    Use the **zpa_application_segment_inspection** data source to get information about an inspection application segment in the Zscaler Private Access cloud. This resource can then be referenced in a ZPA access inspection policy. This resource supports ZPA Inspection for both `HTTP` and `HTTPS`.

    **NOTE:** To ensure consistent search results across data sources, please avoid using multiple spaces or special characters in your search queries.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zpa as zpa

    this = zpa.get_application_segment_inspection(name="ZPA_Inspection_Example")
    ```

    ```python
    import pulumi
    import pulumi_zpa as zpa

    this = zpa.get_application_segment_inspection(id="123456789")
    ```
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['tcpPortRange'] = tcp_port_range
    __args__['udpPortRange'] = udp_port_range
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('zpa:index/getApplicationSegmentInspection:getApplicationSegmentInspection', __args__, opts=opts, typ=GetApplicationSegmentInspectionResult)
    return __ret__.apply(lambda __response__: GetApplicationSegmentInspectionResult(
        bypass_type=pulumi.get(__response__, 'bypass_type'),
        creation_time=pulumi.get(__response__, 'creation_time'),
        description=pulumi.get(__response__, 'description'),
        domain_names=pulumi.get(__response__, 'domain_names'),
        double_encrypt=pulumi.get(__response__, 'double_encrypt'),
        enabled=pulumi.get(__response__, 'enabled'),
        health_check_type=pulumi.get(__response__, 'health_check_type'),
        health_reporting=pulumi.get(__response__, 'health_reporting'),
        icmp_access_type=pulumi.get(__response__, 'icmp_access_type'),
        id=pulumi.get(__response__, 'id'),
        inspection_apps=pulumi.get(__response__, 'inspection_apps'),
        ip_anchored=pulumi.get(__response__, 'ip_anchored'),
        is_cname_enabled=pulumi.get(__response__, 'is_cname_enabled'),
        modified_by=pulumi.get(__response__, 'modified_by'),
        modified_time=pulumi.get(__response__, 'modified_time'),
        name=pulumi.get(__response__, 'name'),
        passive_health_enabled=pulumi.get(__response__, 'passive_health_enabled'),
        segment_group_id=pulumi.get(__response__, 'segment_group_id'),
        segment_group_name=pulumi.get(__response__, 'segment_group_name'),
        select_connector_close_to_app=pulumi.get(__response__, 'select_connector_close_to_app'),
        server_groups=pulumi.get(__response__, 'server_groups'),
        tcp_port_range=pulumi.get(__response__, 'tcp_port_range'),
        tcp_port_ranges=pulumi.get(__response__, 'tcp_port_ranges'),
        udp_port_range=pulumi.get(__response__, 'udp_port_range'),
        udp_port_ranges=pulumi.get(__response__, 'udp_port_ranges')))
