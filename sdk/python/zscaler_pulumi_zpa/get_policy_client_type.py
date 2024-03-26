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
    'GetPolicyClientTypeResult',
    'AwaitableGetPolicyClientTypeResult',
    'get_policy_client_type',
    'get_policy_client_type_output',
]

@pulumi.output_type
class GetPolicyClientTypeResult:
    """
    A collection of values returned by getPolicyClientType.
    """
    def __init__(__self__, id=None, zpn_client_type_branch_connector=None, zpn_client_type_browser_isolation=None, zpn_client_type_edge_connector=None, zpn_client_type_exporter=None, zpn_client_type_exporter_noauth=None, zpn_client_type_ip_anchoring=None, zpn_client_type_machine_tunnel=None, zpn_client_type_slogger=None, zpn_client_type_zapp=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if zpn_client_type_branch_connector and not isinstance(zpn_client_type_branch_connector, str):
            raise TypeError("Expected argument 'zpn_client_type_branch_connector' to be a str")
        pulumi.set(__self__, "zpn_client_type_branch_connector", zpn_client_type_branch_connector)
        if zpn_client_type_browser_isolation and not isinstance(zpn_client_type_browser_isolation, str):
            raise TypeError("Expected argument 'zpn_client_type_browser_isolation' to be a str")
        pulumi.set(__self__, "zpn_client_type_browser_isolation", zpn_client_type_browser_isolation)
        if zpn_client_type_edge_connector and not isinstance(zpn_client_type_edge_connector, str):
            raise TypeError("Expected argument 'zpn_client_type_edge_connector' to be a str")
        pulumi.set(__self__, "zpn_client_type_edge_connector", zpn_client_type_edge_connector)
        if zpn_client_type_exporter and not isinstance(zpn_client_type_exporter, str):
            raise TypeError("Expected argument 'zpn_client_type_exporter' to be a str")
        pulumi.set(__self__, "zpn_client_type_exporter", zpn_client_type_exporter)
        if zpn_client_type_exporter_noauth and not isinstance(zpn_client_type_exporter_noauth, str):
            raise TypeError("Expected argument 'zpn_client_type_exporter_noauth' to be a str")
        pulumi.set(__self__, "zpn_client_type_exporter_noauth", zpn_client_type_exporter_noauth)
        if zpn_client_type_ip_anchoring and not isinstance(zpn_client_type_ip_anchoring, str):
            raise TypeError("Expected argument 'zpn_client_type_ip_anchoring' to be a str")
        pulumi.set(__self__, "zpn_client_type_ip_anchoring", zpn_client_type_ip_anchoring)
        if zpn_client_type_machine_tunnel and not isinstance(zpn_client_type_machine_tunnel, str):
            raise TypeError("Expected argument 'zpn_client_type_machine_tunnel' to be a str")
        pulumi.set(__self__, "zpn_client_type_machine_tunnel", zpn_client_type_machine_tunnel)
        if zpn_client_type_slogger and not isinstance(zpn_client_type_slogger, str):
            raise TypeError("Expected argument 'zpn_client_type_slogger' to be a str")
        pulumi.set(__self__, "zpn_client_type_slogger", zpn_client_type_slogger)
        if zpn_client_type_zapp and not isinstance(zpn_client_type_zapp, str):
            raise TypeError("Expected argument 'zpn_client_type_zapp' to be a str")
        pulumi.set(__self__, "zpn_client_type_zapp", zpn_client_type_zapp)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="zpnClientTypeBranchConnector")
    def zpn_client_type_branch_connector(self) -> str:
        return pulumi.get(self, "zpn_client_type_branch_connector")

    @property
    @pulumi.getter(name="zpnClientTypeBrowserIsolation")
    def zpn_client_type_browser_isolation(self) -> str:
        return pulumi.get(self, "zpn_client_type_browser_isolation")

    @property
    @pulumi.getter(name="zpnClientTypeEdgeConnector")
    def zpn_client_type_edge_connector(self) -> str:
        return pulumi.get(self, "zpn_client_type_edge_connector")

    @property
    @pulumi.getter(name="zpnClientTypeExporter")
    def zpn_client_type_exporter(self) -> str:
        return pulumi.get(self, "zpn_client_type_exporter")

    @property
    @pulumi.getter(name="zpnClientTypeExporterNoauth")
    def zpn_client_type_exporter_noauth(self) -> str:
        return pulumi.get(self, "zpn_client_type_exporter_noauth")

    @property
    @pulumi.getter(name="zpnClientTypeIpAnchoring")
    def zpn_client_type_ip_anchoring(self) -> str:
        return pulumi.get(self, "zpn_client_type_ip_anchoring")

    @property
    @pulumi.getter(name="zpnClientTypeMachineTunnel")
    def zpn_client_type_machine_tunnel(self) -> str:
        return pulumi.get(self, "zpn_client_type_machine_tunnel")

    @property
    @pulumi.getter(name="zpnClientTypeSlogger")
    def zpn_client_type_slogger(self) -> str:
        return pulumi.get(self, "zpn_client_type_slogger")

    @property
    @pulumi.getter(name="zpnClientTypeZapp")
    def zpn_client_type_zapp(self) -> str:
        return pulumi.get(self, "zpn_client_type_zapp")


class AwaitableGetPolicyClientTypeResult(GetPolicyClientTypeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPolicyClientTypeResult(
            id=self.id,
            zpn_client_type_branch_connector=self.zpn_client_type_branch_connector,
            zpn_client_type_browser_isolation=self.zpn_client_type_browser_isolation,
            zpn_client_type_edge_connector=self.zpn_client_type_edge_connector,
            zpn_client_type_exporter=self.zpn_client_type_exporter,
            zpn_client_type_exporter_noauth=self.zpn_client_type_exporter_noauth,
            zpn_client_type_ip_anchoring=self.zpn_client_type_ip_anchoring,
            zpn_client_type_machine_tunnel=self.zpn_client_type_machine_tunnel,
            zpn_client_type_slogger=self.zpn_client_type_slogger,
            zpn_client_type_zapp=self.zpn_client_type_zapp)


def get_policy_client_type(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPolicyClientTypeResult:
    """
    Use the **zpa_access_policy_client_types** data source to get information about all client types for the specified customer in the Zscaler Private Access cloud. This data source can be optionally used when defining the following policy types:
        - ``PolicyAccessRule``
        - ``PolicyAccessTimeOutRule``
        - ``PolicyAccessForwardingRule``
        - ``PolicyAccessIsolationRule``
        - ``PolicyAccessInspectionRule``

    The ``object_type`` attribute must be defined as "CLIENT_TYPE" in the policy operand condition. To learn more see the To learn more see the [Getting Details of All Client Types](https://help.zscaler.com/zpa/configuring-access-policies-using-api#getClientTypes)

    > **NOTE** By Default the ZPA provider will return all client types

    > **NOTE** When defining a ``PolicyAccessIsolationRule`` policy the ``object_type`` "CLIENT_TYPE" is mandatory and ``zpn_client_type_exporter`` is the only supported value.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_zpa as zpa

    this = zpa.get_policy_client_type()
    ```
    <!--End PulumiCodeChooser -->
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zpa:index/getPolicyClientType:getPolicyClientType', __args__, opts=opts, typ=GetPolicyClientTypeResult).value

    return AwaitableGetPolicyClientTypeResult(
        id=pulumi.get(__ret__, 'id'),
        zpn_client_type_branch_connector=pulumi.get(__ret__, 'zpn_client_type_branch_connector'),
        zpn_client_type_browser_isolation=pulumi.get(__ret__, 'zpn_client_type_browser_isolation'),
        zpn_client_type_edge_connector=pulumi.get(__ret__, 'zpn_client_type_edge_connector'),
        zpn_client_type_exporter=pulumi.get(__ret__, 'zpn_client_type_exporter'),
        zpn_client_type_exporter_noauth=pulumi.get(__ret__, 'zpn_client_type_exporter_noauth'),
        zpn_client_type_ip_anchoring=pulumi.get(__ret__, 'zpn_client_type_ip_anchoring'),
        zpn_client_type_machine_tunnel=pulumi.get(__ret__, 'zpn_client_type_machine_tunnel'),
        zpn_client_type_slogger=pulumi.get(__ret__, 'zpn_client_type_slogger'),
        zpn_client_type_zapp=pulumi.get(__ret__, 'zpn_client_type_zapp'))


@_utilities.lift_output_func(get_policy_client_type)
def get_policy_client_type_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPolicyClientTypeResult]:
    """
    Use the **zpa_access_policy_client_types** data source to get information about all client types for the specified customer in the Zscaler Private Access cloud. This data source can be optionally used when defining the following policy types:
        - ``PolicyAccessRule``
        - ``PolicyAccessTimeOutRule``
        - ``PolicyAccessForwardingRule``
        - ``PolicyAccessIsolationRule``
        - ``PolicyAccessInspectionRule``

    The ``object_type`` attribute must be defined as "CLIENT_TYPE" in the policy operand condition. To learn more see the To learn more see the [Getting Details of All Client Types](https://help.zscaler.com/zpa/configuring-access-policies-using-api#getClientTypes)

    > **NOTE** By Default the ZPA provider will return all client types

    > **NOTE** When defining a ``PolicyAccessIsolationRule`` policy the ``object_type`` "CLIENT_TYPE" is mandatory and ``zpn_client_type_exporter`` is the only supported value.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_zpa as zpa

    this = zpa.get_policy_client_type()
    ```
    <!--End PulumiCodeChooser -->
    """
    ...
