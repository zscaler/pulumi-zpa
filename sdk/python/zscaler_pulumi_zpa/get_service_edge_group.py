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

__all__ = [
    'GetServiceEdgeGroupResult',
    'AwaitableGetServiceEdgeGroupResult',
    'get_service_edge_group',
    'get_service_edge_group_output',
]

@pulumi.output_type
class GetServiceEdgeGroupResult:
    """
    A collection of values returned by getServiceEdgeGroup.
    """
    def __init__(__self__, alt_cloud=None, city_country=None, country_code=None, creation_time=None, description=None, enabled=None, geo_location_id=None, grace_distance_enabled=None, grace_distance_value=None, grace_distance_value_unit=None, id=None, is_public=None, latitude=None, location=None, longitude=None, modified_by=None, modified_time=None, name=None, override_version_profile=None, service_edges=None, site_id=None, site_name=None, trusted_networks=None, upgrade_day=None, upgrade_time_in_secs=None, use_in_dr_mode=None, version_profile_id=None, version_profile_name=None, version_profile_visibility_scope=None):
        if alt_cloud and not isinstance(alt_cloud, str):
            raise TypeError("Expected argument 'alt_cloud' to be a str")
        pulumi.set(__self__, "alt_cloud", alt_cloud)
        if city_country and not isinstance(city_country, str):
            raise TypeError("Expected argument 'city_country' to be a str")
        pulumi.set(__self__, "city_country", city_country)
        if country_code and not isinstance(country_code, str):
            raise TypeError("Expected argument 'country_code' to be a str")
        pulumi.set(__self__, "country_code", country_code)
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if geo_location_id and not isinstance(geo_location_id, str):
            raise TypeError("Expected argument 'geo_location_id' to be a str")
        pulumi.set(__self__, "geo_location_id", geo_location_id)
        if grace_distance_enabled and not isinstance(grace_distance_enabled, bool):
            raise TypeError("Expected argument 'grace_distance_enabled' to be a bool")
        pulumi.set(__self__, "grace_distance_enabled", grace_distance_enabled)
        if grace_distance_value and not isinstance(grace_distance_value, str):
            raise TypeError("Expected argument 'grace_distance_value' to be a str")
        pulumi.set(__self__, "grace_distance_value", grace_distance_value)
        if grace_distance_value_unit and not isinstance(grace_distance_value_unit, str):
            raise TypeError("Expected argument 'grace_distance_value_unit' to be a str")
        pulumi.set(__self__, "grace_distance_value_unit", grace_distance_value_unit)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_public and not isinstance(is_public, str):
            raise TypeError("Expected argument 'is_public' to be a str")
        pulumi.set(__self__, "is_public", is_public)
        if latitude and not isinstance(latitude, str):
            raise TypeError("Expected argument 'latitude' to be a str")
        pulumi.set(__self__, "latitude", latitude)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if longitude and not isinstance(longitude, str):
            raise TypeError("Expected argument 'longitude' to be a str")
        pulumi.set(__self__, "longitude", longitude)
        if modified_by and not isinstance(modified_by, str):
            raise TypeError("Expected argument 'modified_by' to be a str")
        pulumi.set(__self__, "modified_by", modified_by)
        if modified_time and not isinstance(modified_time, str):
            raise TypeError("Expected argument 'modified_time' to be a str")
        pulumi.set(__self__, "modified_time", modified_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if override_version_profile and not isinstance(override_version_profile, bool):
            raise TypeError("Expected argument 'override_version_profile' to be a bool")
        pulumi.set(__self__, "override_version_profile", override_version_profile)
        if service_edges and not isinstance(service_edges, list):
            raise TypeError("Expected argument 'service_edges' to be a list")
        pulumi.set(__self__, "service_edges", service_edges)
        if site_id and not isinstance(site_id, str):
            raise TypeError("Expected argument 'site_id' to be a str")
        pulumi.set(__self__, "site_id", site_id)
        if site_name and not isinstance(site_name, str):
            raise TypeError("Expected argument 'site_name' to be a str")
        pulumi.set(__self__, "site_name", site_name)
        if trusted_networks and not isinstance(trusted_networks, list):
            raise TypeError("Expected argument 'trusted_networks' to be a list")
        pulumi.set(__self__, "trusted_networks", trusted_networks)
        if upgrade_day and not isinstance(upgrade_day, str):
            raise TypeError("Expected argument 'upgrade_day' to be a str")
        pulumi.set(__self__, "upgrade_day", upgrade_day)
        if upgrade_time_in_secs and not isinstance(upgrade_time_in_secs, str):
            raise TypeError("Expected argument 'upgrade_time_in_secs' to be a str")
        pulumi.set(__self__, "upgrade_time_in_secs", upgrade_time_in_secs)
        if use_in_dr_mode and not isinstance(use_in_dr_mode, bool):
            raise TypeError("Expected argument 'use_in_dr_mode' to be a bool")
        pulumi.set(__self__, "use_in_dr_mode", use_in_dr_mode)
        if version_profile_id and not isinstance(version_profile_id, str):
            raise TypeError("Expected argument 'version_profile_id' to be a str")
        pulumi.set(__self__, "version_profile_id", version_profile_id)
        if version_profile_name and not isinstance(version_profile_name, str):
            raise TypeError("Expected argument 'version_profile_name' to be a str")
        pulumi.set(__self__, "version_profile_name", version_profile_name)
        if version_profile_visibility_scope and not isinstance(version_profile_visibility_scope, str):
            raise TypeError("Expected argument 'version_profile_visibility_scope' to be a str")
        pulumi.set(__self__, "version_profile_visibility_scope", version_profile_visibility_scope)

    @property
    @pulumi.getter(name="altCloud")
    def alt_cloud(self) -> builtins.str:
        return pulumi.get(self, "alt_cloud")

    @property
    @pulumi.getter(name="cityCountry")
    def city_country(self) -> builtins.str:
        return pulumi.get(self, "city_country")

    @property
    @pulumi.getter(name="countryCode")
    def country_code(self) -> builtins.str:
        return pulumi.get(self, "country_code")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> builtins.str:
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> builtins.bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="geoLocationId")
    def geo_location_id(self) -> builtins.str:
        return pulumi.get(self, "geo_location_id")

    @property
    @pulumi.getter(name="graceDistanceEnabled")
    def grace_distance_enabled(self) -> builtins.bool:
        return pulumi.get(self, "grace_distance_enabled")

    @property
    @pulumi.getter(name="graceDistanceValue")
    def grace_distance_value(self) -> builtins.str:
        return pulumi.get(self, "grace_distance_value")

    @property
    @pulumi.getter(name="graceDistanceValueUnit")
    def grace_distance_value_unit(self) -> builtins.str:
        return pulumi.get(self, "grace_distance_value_unit")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> builtins.str:
        return pulumi.get(self, "is_public")

    @property
    @pulumi.getter
    def latitude(self) -> builtins.str:
        return pulumi.get(self, "latitude")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def longitude(self) -> builtins.str:
        return pulumi.get(self, "longitude")

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
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="overrideVersionProfile")
    def override_version_profile(self) -> builtins.bool:
        return pulumi.get(self, "override_version_profile")

    @property
    @pulumi.getter(name="serviceEdges")
    def service_edges(self) -> Sequence['outputs.GetServiceEdgeGroupServiceEdgeResult']:
        return pulumi.get(self, "service_edges")

    @property
    @pulumi.getter(name="siteId")
    def site_id(self) -> builtins.str:
        return pulumi.get(self, "site_id")

    @property
    @pulumi.getter(name="siteName")
    def site_name(self) -> builtins.str:
        return pulumi.get(self, "site_name")

    @property
    @pulumi.getter(name="trustedNetworks")
    def trusted_networks(self) -> Sequence['outputs.GetServiceEdgeGroupTrustedNetworkResult']:
        return pulumi.get(self, "trusted_networks")

    @property
    @pulumi.getter(name="upgradeDay")
    def upgrade_day(self) -> builtins.str:
        return pulumi.get(self, "upgrade_day")

    @property
    @pulumi.getter(name="upgradeTimeInSecs")
    def upgrade_time_in_secs(self) -> builtins.str:
        return pulumi.get(self, "upgrade_time_in_secs")

    @property
    @pulumi.getter(name="useInDrMode")
    def use_in_dr_mode(self) -> builtins.bool:
        return pulumi.get(self, "use_in_dr_mode")

    @property
    @pulumi.getter(name="versionProfileId")
    def version_profile_id(self) -> builtins.str:
        return pulumi.get(self, "version_profile_id")

    @property
    @pulumi.getter(name="versionProfileName")
    def version_profile_name(self) -> builtins.str:
        return pulumi.get(self, "version_profile_name")

    @property
    @pulumi.getter(name="versionProfileVisibilityScope")
    def version_profile_visibility_scope(self) -> builtins.str:
        return pulumi.get(self, "version_profile_visibility_scope")


class AwaitableGetServiceEdgeGroupResult(GetServiceEdgeGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceEdgeGroupResult(
            alt_cloud=self.alt_cloud,
            city_country=self.city_country,
            country_code=self.country_code,
            creation_time=self.creation_time,
            description=self.description,
            enabled=self.enabled,
            geo_location_id=self.geo_location_id,
            grace_distance_enabled=self.grace_distance_enabled,
            grace_distance_value=self.grace_distance_value,
            grace_distance_value_unit=self.grace_distance_value_unit,
            id=self.id,
            is_public=self.is_public,
            latitude=self.latitude,
            location=self.location,
            longitude=self.longitude,
            modified_by=self.modified_by,
            modified_time=self.modified_time,
            name=self.name,
            override_version_profile=self.override_version_profile,
            service_edges=self.service_edges,
            site_id=self.site_id,
            site_name=self.site_name,
            trusted_networks=self.trusted_networks,
            upgrade_day=self.upgrade_day,
            upgrade_time_in_secs=self.upgrade_time_in_secs,
            use_in_dr_mode=self.use_in_dr_mode,
            version_profile_id=self.version_profile_id,
            version_profile_name=self.version_profile_name,
            version_profile_visibility_scope=self.version_profile_visibility_scope)


def get_service_edge_group(id: Optional[builtins.str] = None,
                           name: Optional[builtins.str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceEdgeGroupResult:
    """
    * [Official documentation](https://help.zscaler.com/zpa/about-zpa-private-service-edge-groups)
    * [API documentation](https://help.zscaler.com/zpa/configuring-zpa-private-service-edge-groups-using-api)

    Use the **zpa_service_edge_group** data source to get information about a service edge group in the Zscaler Private Access cloud. This data source can then be referenced in an App Connector Group. This data source can then be referenced in the following resources:

    **NOTE:** To ensure consistent search results across data sources, please avoid using multiple spaces or special characters in your search queries.

    * Create a server group
    * Provisioning Key
    * Access policy rule

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zpa as zpa

    foo = zpa.get_service_edge_group(name="DataCenter")
    ```

    ```python
    import pulumi
    import pulumi_zpa as zpa

    foo = zpa.get_service_edge_group(id="123456789")
    ```
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zpa:index/getServiceEdgeGroup:getServiceEdgeGroup', __args__, opts=opts, typ=GetServiceEdgeGroupResult).value

    return AwaitableGetServiceEdgeGroupResult(
        alt_cloud=pulumi.get(__ret__, 'alt_cloud'),
        city_country=pulumi.get(__ret__, 'city_country'),
        country_code=pulumi.get(__ret__, 'country_code'),
        creation_time=pulumi.get(__ret__, 'creation_time'),
        description=pulumi.get(__ret__, 'description'),
        enabled=pulumi.get(__ret__, 'enabled'),
        geo_location_id=pulumi.get(__ret__, 'geo_location_id'),
        grace_distance_enabled=pulumi.get(__ret__, 'grace_distance_enabled'),
        grace_distance_value=pulumi.get(__ret__, 'grace_distance_value'),
        grace_distance_value_unit=pulumi.get(__ret__, 'grace_distance_value_unit'),
        id=pulumi.get(__ret__, 'id'),
        is_public=pulumi.get(__ret__, 'is_public'),
        latitude=pulumi.get(__ret__, 'latitude'),
        location=pulumi.get(__ret__, 'location'),
        longitude=pulumi.get(__ret__, 'longitude'),
        modified_by=pulumi.get(__ret__, 'modified_by'),
        modified_time=pulumi.get(__ret__, 'modified_time'),
        name=pulumi.get(__ret__, 'name'),
        override_version_profile=pulumi.get(__ret__, 'override_version_profile'),
        service_edges=pulumi.get(__ret__, 'service_edges'),
        site_id=pulumi.get(__ret__, 'site_id'),
        site_name=pulumi.get(__ret__, 'site_name'),
        trusted_networks=pulumi.get(__ret__, 'trusted_networks'),
        upgrade_day=pulumi.get(__ret__, 'upgrade_day'),
        upgrade_time_in_secs=pulumi.get(__ret__, 'upgrade_time_in_secs'),
        use_in_dr_mode=pulumi.get(__ret__, 'use_in_dr_mode'),
        version_profile_id=pulumi.get(__ret__, 'version_profile_id'),
        version_profile_name=pulumi.get(__ret__, 'version_profile_name'),
        version_profile_visibility_scope=pulumi.get(__ret__, 'version_profile_visibility_scope'))
def get_service_edge_group_output(id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                  name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetServiceEdgeGroupResult]:
    """
    * [Official documentation](https://help.zscaler.com/zpa/about-zpa-private-service-edge-groups)
    * [API documentation](https://help.zscaler.com/zpa/configuring-zpa-private-service-edge-groups-using-api)

    Use the **zpa_service_edge_group** data source to get information about a service edge group in the Zscaler Private Access cloud. This data source can then be referenced in an App Connector Group. This data source can then be referenced in the following resources:

    **NOTE:** To ensure consistent search results across data sources, please avoid using multiple spaces or special characters in your search queries.

    * Create a server group
    * Provisioning Key
    * Access policy rule

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zpa as zpa

    foo = zpa.get_service_edge_group(name="DataCenter")
    ```

    ```python
    import pulumi
    import pulumi_zpa as zpa

    foo = zpa.get_service_edge_group(id="123456789")
    ```
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('zpa:index/getServiceEdgeGroup:getServiceEdgeGroup', __args__, opts=opts, typ=GetServiceEdgeGroupResult)
    return __ret__.apply(lambda __response__: GetServiceEdgeGroupResult(
        alt_cloud=pulumi.get(__response__, 'alt_cloud'),
        city_country=pulumi.get(__response__, 'city_country'),
        country_code=pulumi.get(__response__, 'country_code'),
        creation_time=pulumi.get(__response__, 'creation_time'),
        description=pulumi.get(__response__, 'description'),
        enabled=pulumi.get(__response__, 'enabled'),
        geo_location_id=pulumi.get(__response__, 'geo_location_id'),
        grace_distance_enabled=pulumi.get(__response__, 'grace_distance_enabled'),
        grace_distance_value=pulumi.get(__response__, 'grace_distance_value'),
        grace_distance_value_unit=pulumi.get(__response__, 'grace_distance_value_unit'),
        id=pulumi.get(__response__, 'id'),
        is_public=pulumi.get(__response__, 'is_public'),
        latitude=pulumi.get(__response__, 'latitude'),
        location=pulumi.get(__response__, 'location'),
        longitude=pulumi.get(__response__, 'longitude'),
        modified_by=pulumi.get(__response__, 'modified_by'),
        modified_time=pulumi.get(__response__, 'modified_time'),
        name=pulumi.get(__response__, 'name'),
        override_version_profile=pulumi.get(__response__, 'override_version_profile'),
        service_edges=pulumi.get(__response__, 'service_edges'),
        site_id=pulumi.get(__response__, 'site_id'),
        site_name=pulumi.get(__response__, 'site_name'),
        trusted_networks=pulumi.get(__response__, 'trusted_networks'),
        upgrade_day=pulumi.get(__response__, 'upgrade_day'),
        upgrade_time_in_secs=pulumi.get(__response__, 'upgrade_time_in_secs'),
        use_in_dr_mode=pulumi.get(__response__, 'use_in_dr_mode'),
        version_profile_id=pulumi.get(__response__, 'version_profile_id'),
        version_profile_name=pulumi.get(__response__, 'version_profile_name'),
        version_profile_visibility_scope=pulumi.get(__response__, 'version_profile_visibility_scope')))
