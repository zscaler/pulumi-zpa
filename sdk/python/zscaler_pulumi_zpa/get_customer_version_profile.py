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
    'GetCustomerVersionProfileResult',
    'AwaitableGetCustomerVersionProfileResult',
    'get_customer_version_profile',
    'get_customer_version_profile_output',
]

@pulumi.output_type
class GetCustomerVersionProfileResult:
    """
    A collection of values returned by getCustomerVersionProfile.
    """
    def __init__(__self__, creation_time=None, custom_scope_customer_ids=None, custom_scope_request_customer_ids=None, customer_id=None, description=None, id=None, modified_by=None, modified_time=None, name=None, number_of_assistants=None, number_of_customers=None, number_of_private_brokers=None, number_of_site_controllers=None, number_of_updated_assistants=None, number_of_updated_private_brokers=None, number_of_updated_site_controllers=None, upgrade_priority=None, versions=None, visibility_scope=None):
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if custom_scope_customer_ids and not isinstance(custom_scope_customer_ids, list):
            raise TypeError("Expected argument 'custom_scope_customer_ids' to be a list")
        pulumi.set(__self__, "custom_scope_customer_ids", custom_scope_customer_ids)
        if custom_scope_request_customer_ids and not isinstance(custom_scope_request_customer_ids, list):
            raise TypeError("Expected argument 'custom_scope_request_customer_ids' to be a list")
        pulumi.set(__self__, "custom_scope_request_customer_ids", custom_scope_request_customer_ids)
        if customer_id and not isinstance(customer_id, str):
            raise TypeError("Expected argument 'customer_id' to be a str")
        pulumi.set(__self__, "customer_id", customer_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if modified_by and not isinstance(modified_by, str):
            raise TypeError("Expected argument 'modified_by' to be a str")
        pulumi.set(__self__, "modified_by", modified_by)
        if modified_time and not isinstance(modified_time, str):
            raise TypeError("Expected argument 'modified_time' to be a str")
        pulumi.set(__self__, "modified_time", modified_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if number_of_assistants and not isinstance(number_of_assistants, str):
            raise TypeError("Expected argument 'number_of_assistants' to be a str")
        pulumi.set(__self__, "number_of_assistants", number_of_assistants)
        if number_of_customers and not isinstance(number_of_customers, str):
            raise TypeError("Expected argument 'number_of_customers' to be a str")
        pulumi.set(__self__, "number_of_customers", number_of_customers)
        if number_of_private_brokers and not isinstance(number_of_private_brokers, str):
            raise TypeError("Expected argument 'number_of_private_brokers' to be a str")
        pulumi.set(__self__, "number_of_private_brokers", number_of_private_brokers)
        if number_of_site_controllers and not isinstance(number_of_site_controllers, str):
            raise TypeError("Expected argument 'number_of_site_controllers' to be a str")
        pulumi.set(__self__, "number_of_site_controllers", number_of_site_controllers)
        if number_of_updated_assistants and not isinstance(number_of_updated_assistants, str):
            raise TypeError("Expected argument 'number_of_updated_assistants' to be a str")
        pulumi.set(__self__, "number_of_updated_assistants", number_of_updated_assistants)
        if number_of_updated_private_brokers and not isinstance(number_of_updated_private_brokers, str):
            raise TypeError("Expected argument 'number_of_updated_private_brokers' to be a str")
        pulumi.set(__self__, "number_of_updated_private_brokers", number_of_updated_private_brokers)
        if number_of_updated_site_controllers and not isinstance(number_of_updated_site_controllers, str):
            raise TypeError("Expected argument 'number_of_updated_site_controllers' to be a str")
        pulumi.set(__self__, "number_of_updated_site_controllers", number_of_updated_site_controllers)
        if upgrade_priority and not isinstance(upgrade_priority, str):
            raise TypeError("Expected argument 'upgrade_priority' to be a str")
        pulumi.set(__self__, "upgrade_priority", upgrade_priority)
        if versions and not isinstance(versions, list):
            raise TypeError("Expected argument 'versions' to be a list")
        pulumi.set(__self__, "versions", versions)
        if visibility_scope and not isinstance(visibility_scope, str):
            raise TypeError("Expected argument 'visibility_scope' to be a str")
        pulumi.set(__self__, "visibility_scope", visibility_scope)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> builtins.str:
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="customScopeCustomerIds")
    def custom_scope_customer_ids(self) -> Sequence['outputs.GetCustomerVersionProfileCustomScopeCustomerIdResult']:
        return pulumi.get(self, "custom_scope_customer_ids")

    @property
    @pulumi.getter(name="customScopeRequestCustomerIds")
    def custom_scope_request_customer_ids(self) -> Sequence['outputs.GetCustomerVersionProfileCustomScopeRequestCustomerIdResult']:
        return pulumi.get(self, "custom_scope_request_customer_ids")

    @property
    @pulumi.getter(name="customerId")
    def customer_id(self) -> builtins.str:
        return pulumi.get(self, "customer_id")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        return pulumi.get(self, "id")

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
    @pulumi.getter(name="numberOfAssistants")
    def number_of_assistants(self) -> builtins.str:
        return pulumi.get(self, "number_of_assistants")

    @property
    @pulumi.getter(name="numberOfCustomers")
    def number_of_customers(self) -> builtins.str:
        return pulumi.get(self, "number_of_customers")

    @property
    @pulumi.getter(name="numberOfPrivateBrokers")
    def number_of_private_brokers(self) -> builtins.str:
        return pulumi.get(self, "number_of_private_brokers")

    @property
    @pulumi.getter(name="numberOfSiteControllers")
    def number_of_site_controllers(self) -> builtins.str:
        return pulumi.get(self, "number_of_site_controllers")

    @property
    @pulumi.getter(name="numberOfUpdatedAssistants")
    def number_of_updated_assistants(self) -> builtins.str:
        return pulumi.get(self, "number_of_updated_assistants")

    @property
    @pulumi.getter(name="numberOfUpdatedPrivateBrokers")
    def number_of_updated_private_brokers(self) -> builtins.str:
        return pulumi.get(self, "number_of_updated_private_brokers")

    @property
    @pulumi.getter(name="numberOfUpdatedSiteControllers")
    def number_of_updated_site_controllers(self) -> builtins.str:
        return pulumi.get(self, "number_of_updated_site_controllers")

    @property
    @pulumi.getter(name="upgradePriority")
    def upgrade_priority(self) -> builtins.str:
        return pulumi.get(self, "upgrade_priority")

    @property
    @pulumi.getter
    def versions(self) -> Sequence['outputs.GetCustomerVersionProfileVersionResult']:
        return pulumi.get(self, "versions")

    @property
    @pulumi.getter(name="visibilityScope")
    def visibility_scope(self) -> builtins.str:
        return pulumi.get(self, "visibility_scope")


class AwaitableGetCustomerVersionProfileResult(GetCustomerVersionProfileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCustomerVersionProfileResult(
            creation_time=self.creation_time,
            custom_scope_customer_ids=self.custom_scope_customer_ids,
            custom_scope_request_customer_ids=self.custom_scope_request_customer_ids,
            customer_id=self.customer_id,
            description=self.description,
            id=self.id,
            modified_by=self.modified_by,
            modified_time=self.modified_time,
            name=self.name,
            number_of_assistants=self.number_of_assistants,
            number_of_customers=self.number_of_customers,
            number_of_private_brokers=self.number_of_private_brokers,
            number_of_site_controllers=self.number_of_site_controllers,
            number_of_updated_assistants=self.number_of_updated_assistants,
            number_of_updated_private_brokers=self.number_of_updated_private_brokers,
            number_of_updated_site_controllers=self.number_of_updated_site_controllers,
            upgrade_priority=self.upgrade_priority,
            versions=self.versions,
            visibility_scope=self.visibility_scope)


def get_customer_version_profile(name: Optional[builtins.str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCustomerVersionProfileResult:
    """
    * [Official documentation](https://help.zscaler.com/zpa/about-connectors)
    * [API documentation](https://help.zscaler.com/zpa/obtaining-version-profile-details-using-api)

    Use the **zpa_customer_version_profile** data source to get information about all customer version profiles from the Zscaler Private Access cloud. This data source can be associated with an App Connector Group within the parameter `version_profile_id` or `version_profile_name`

    The customer version profile IDs are:

    * `Default` = `0`
    * `Previous Default` = `1`
    * `New Release` = `2`

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zpa as zpa

    default = zpa.get_customer_version_profile(name="Default")
    previous_default = zpa.get_customer_version_profile(name="Previous Default")
    new_release = zpa.get_customer_version_profile(name="New Release")
    ```
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zpa:index/getCustomerVersionProfile:getCustomerVersionProfile', __args__, opts=opts, typ=GetCustomerVersionProfileResult).value

    return AwaitableGetCustomerVersionProfileResult(
        creation_time=pulumi.get(__ret__, 'creation_time'),
        custom_scope_customer_ids=pulumi.get(__ret__, 'custom_scope_customer_ids'),
        custom_scope_request_customer_ids=pulumi.get(__ret__, 'custom_scope_request_customer_ids'),
        customer_id=pulumi.get(__ret__, 'customer_id'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        modified_by=pulumi.get(__ret__, 'modified_by'),
        modified_time=pulumi.get(__ret__, 'modified_time'),
        name=pulumi.get(__ret__, 'name'),
        number_of_assistants=pulumi.get(__ret__, 'number_of_assistants'),
        number_of_customers=pulumi.get(__ret__, 'number_of_customers'),
        number_of_private_brokers=pulumi.get(__ret__, 'number_of_private_brokers'),
        number_of_site_controllers=pulumi.get(__ret__, 'number_of_site_controllers'),
        number_of_updated_assistants=pulumi.get(__ret__, 'number_of_updated_assistants'),
        number_of_updated_private_brokers=pulumi.get(__ret__, 'number_of_updated_private_brokers'),
        number_of_updated_site_controllers=pulumi.get(__ret__, 'number_of_updated_site_controllers'),
        upgrade_priority=pulumi.get(__ret__, 'upgrade_priority'),
        versions=pulumi.get(__ret__, 'versions'),
        visibility_scope=pulumi.get(__ret__, 'visibility_scope'))
def get_customer_version_profile_output(name: Optional[pulumi.Input[builtins.str]] = None,
                                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCustomerVersionProfileResult]:
    """
    * [Official documentation](https://help.zscaler.com/zpa/about-connectors)
    * [API documentation](https://help.zscaler.com/zpa/obtaining-version-profile-details-using-api)

    Use the **zpa_customer_version_profile** data source to get information about all customer version profiles from the Zscaler Private Access cloud. This data source can be associated with an App Connector Group within the parameter `version_profile_id` or `version_profile_name`

    The customer version profile IDs are:

    * `Default` = `0`
    * `Previous Default` = `1`
    * `New Release` = `2`

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zpa as zpa

    default = zpa.get_customer_version_profile(name="Default")
    previous_default = zpa.get_customer_version_profile(name="Previous Default")
    new_release = zpa.get_customer_version_profile(name="New Release")
    ```
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('zpa:index/getCustomerVersionProfile:getCustomerVersionProfile', __args__, opts=opts, typ=GetCustomerVersionProfileResult)
    return __ret__.apply(lambda __response__: GetCustomerVersionProfileResult(
        creation_time=pulumi.get(__response__, 'creation_time'),
        custom_scope_customer_ids=pulumi.get(__response__, 'custom_scope_customer_ids'),
        custom_scope_request_customer_ids=pulumi.get(__response__, 'custom_scope_request_customer_ids'),
        customer_id=pulumi.get(__response__, 'customer_id'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        modified_by=pulumi.get(__response__, 'modified_by'),
        modified_time=pulumi.get(__response__, 'modified_time'),
        name=pulumi.get(__response__, 'name'),
        number_of_assistants=pulumi.get(__response__, 'number_of_assistants'),
        number_of_customers=pulumi.get(__response__, 'number_of_customers'),
        number_of_private_brokers=pulumi.get(__response__, 'number_of_private_brokers'),
        number_of_site_controllers=pulumi.get(__response__, 'number_of_site_controllers'),
        number_of_updated_assistants=pulumi.get(__response__, 'number_of_updated_assistants'),
        number_of_updated_private_brokers=pulumi.get(__response__, 'number_of_updated_private_brokers'),
        number_of_updated_site_controllers=pulumi.get(__response__, 'number_of_updated_site_controllers'),
        upgrade_priority=pulumi.get(__response__, 'upgrade_priority'),
        versions=pulumi.get(__response__, 'versions'),
        visibility_scope=pulumi.get(__response__, 'visibility_scope')))
