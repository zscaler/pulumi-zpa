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
    'GetAssistantScheduleResult',
    'AwaitableGetAssistantScheduleResult',
    'get_assistant_schedule',
    'get_assistant_schedule_output',
]

@pulumi.output_type
class GetAssistantScheduleResult:
    """
    A collection of values returned by getAssistantSchedule.
    """
    def __init__(__self__, customer_id=None, delete_disabled=None, enabled=None, frequency=None, frequency_interval=None, id=None):
        if customer_id and not isinstance(customer_id, str):
            raise TypeError("Expected argument 'customer_id' to be a str")
        pulumi.set(__self__, "customer_id", customer_id)
        if delete_disabled and not isinstance(delete_disabled, bool):
            raise TypeError("Expected argument 'delete_disabled' to be a bool")
        pulumi.set(__self__, "delete_disabled", delete_disabled)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if frequency and not isinstance(frequency, str):
            raise TypeError("Expected argument 'frequency' to be a str")
        pulumi.set(__self__, "frequency", frequency)
        if frequency_interval and not isinstance(frequency_interval, str):
            raise TypeError("Expected argument 'frequency_interval' to be a str")
        pulumi.set(__self__, "frequency_interval", frequency_interval)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="customerId")
    def customer_id(self) -> Optional[str]:
        return pulumi.get(self, "customer_id")

    @property
    @pulumi.getter(name="deleteDisabled")
    def delete_disabled(self) -> bool:
        return pulumi.get(self, "delete_disabled")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def frequency(self) -> str:
        return pulumi.get(self, "frequency")

    @property
    @pulumi.getter(name="frequencyInterval")
    def frequency_interval(self) -> str:
        return pulumi.get(self, "frequency_interval")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


class AwaitableGetAssistantScheduleResult(GetAssistantScheduleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAssistantScheduleResult(
            customer_id=self.customer_id,
            delete_disabled=self.delete_disabled,
            enabled=self.enabled,
            frequency=self.frequency,
            frequency_interval=self.frequency_interval,
            id=self.id)


def get_assistant_schedule(customer_id: Optional[str] = None,
                           id: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAssistantScheduleResult:
    """
    * [Official documentation](https://help.zscaler.com/zpa/deleting-disconnected-app-connectors)
    * [API documentation](https://help.zscaler.com/zpa/configuring-auto-delete-disconnected-app-connectors-using-api)

    Use the **zpa_app_connector_assistant_schedule** data source to get information about Auto Delete frequency of the App Connector for the specified customer in the Zscaler Private Access cloud.

    > **NOTE** - The `customer_id` attribute is optional and not required during the configuration.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_zpa as zpa

    this = zpa.get_assistant_schedule(customer_id="1234567891012")
    ```
    <!--End PulumiCodeChooser -->
    """
    __args__ = dict()
    __args__['customerId'] = customer_id
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zpa:index/getAssistantSchedule:getAssistantSchedule', __args__, opts=opts, typ=GetAssistantScheduleResult).value

    return AwaitableGetAssistantScheduleResult(
        customer_id=pulumi.get(__ret__, 'customer_id'),
        delete_disabled=pulumi.get(__ret__, 'delete_disabled'),
        enabled=pulumi.get(__ret__, 'enabled'),
        frequency=pulumi.get(__ret__, 'frequency'),
        frequency_interval=pulumi.get(__ret__, 'frequency_interval'),
        id=pulumi.get(__ret__, 'id'))


@_utilities.lift_output_func(get_assistant_schedule)
def get_assistant_schedule_output(customer_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  id: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAssistantScheduleResult]:
    """
    * [Official documentation](https://help.zscaler.com/zpa/deleting-disconnected-app-connectors)
    * [API documentation](https://help.zscaler.com/zpa/configuring-auto-delete-disconnected-app-connectors-using-api)

    Use the **zpa_app_connector_assistant_schedule** data source to get information about Auto Delete frequency of the App Connector for the specified customer in the Zscaler Private Access cloud.

    > **NOTE** - The `customer_id` attribute is optional and not required during the configuration.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_zpa as zpa

    this = zpa.get_assistant_schedule(customer_id="1234567891012")
    ```
    <!--End PulumiCodeChooser -->
    """
    ...
