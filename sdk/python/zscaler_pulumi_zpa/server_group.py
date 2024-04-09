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

__all__ = ['ServerGroupArgs', 'ServerGroup']

@pulumi.input_type
class ServerGroupArgs:
    def __init__(__self__, *,
                 app_connector_groups: Optional[pulumi.Input[Sequence[pulumi.Input['ServerGroupAppConnectorGroupArgs']]]] = None,
                 applications: Optional[pulumi.Input[Sequence[pulumi.Input['ServerGroupApplicationArgs']]]] = None,
                 config_space: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamic_discovery: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 ip_anchored: Optional[pulumi.Input[bool]] = None,
                 microtenant_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 servers: Optional[pulumi.Input[Sequence[pulumi.Input['ServerGroupServerArgs']]]] = None):
        """
        The set of arguments for constructing a ServerGroup resource.
        :param pulumi.Input[Sequence[pulumi.Input['ServerGroupAppConnectorGroupArgs']]] app_connector_groups: List of app-connector IDs.
        :param pulumi.Input[Sequence[pulumi.Input['ServerGroupApplicationArgs']]] applications: This field is a json array of app-connector-id only.
        :param pulumi.Input[str] description: This field is the description of the server group.
        :param pulumi.Input[bool] dynamic_discovery: This field controls dynamic discovery of the servers.
        :param pulumi.Input[bool] enabled: This field defines if the server group is enabled or disabled.
        :param pulumi.Input[str] name: This field defines the name of the server group.
        :param pulumi.Input[Sequence[pulumi.Input['ServerGroupServerArgs']]] servers: This field is a list of servers that are applicable only when dynamic discovery is disabled. Server name is required
               only in cases where the new servers need to be created in this API. For existing servers, pass only the serverId.
        """
        if app_connector_groups is not None:
            pulumi.set(__self__, "app_connector_groups", app_connector_groups)
        if applications is not None:
            pulumi.set(__self__, "applications", applications)
        if config_space is not None:
            pulumi.set(__self__, "config_space", config_space)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dynamic_discovery is not None:
            pulumi.set(__self__, "dynamic_discovery", dynamic_discovery)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if ip_anchored is not None:
            pulumi.set(__self__, "ip_anchored", ip_anchored)
        if microtenant_id is not None:
            pulumi.set(__self__, "microtenant_id", microtenant_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if servers is not None:
            pulumi.set(__self__, "servers", servers)

    @property
    @pulumi.getter(name="appConnectorGroups")
    def app_connector_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServerGroupAppConnectorGroupArgs']]]]:
        """
        List of app-connector IDs.
        """
        return pulumi.get(self, "app_connector_groups")

    @app_connector_groups.setter
    def app_connector_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServerGroupAppConnectorGroupArgs']]]]):
        pulumi.set(self, "app_connector_groups", value)

    @property
    @pulumi.getter
    def applications(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServerGroupApplicationArgs']]]]:
        """
        This field is a json array of app-connector-id only.
        """
        return pulumi.get(self, "applications")

    @applications.setter
    def applications(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServerGroupApplicationArgs']]]]):
        pulumi.set(self, "applications", value)

    @property
    @pulumi.getter(name="configSpace")
    def config_space(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "config_space")

    @config_space.setter
    def config_space(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_space", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        This field is the description of the server group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dynamicDiscovery")
    def dynamic_discovery(self) -> Optional[pulumi.Input[bool]]:
        """
        This field controls dynamic discovery of the servers.
        """
        return pulumi.get(self, "dynamic_discovery")

    @dynamic_discovery.setter
    def dynamic_discovery(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dynamic_discovery", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        This field defines if the server group is enabled or disabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="ipAnchored")
    def ip_anchored(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "ip_anchored")

    @ip_anchored.setter
    def ip_anchored(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ip_anchored", value)

    @property
    @pulumi.getter(name="microtenantId")
    def microtenant_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "microtenant_id")

    @microtenant_id.setter
    def microtenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "microtenant_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        This field defines the name of the server group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServerGroupServerArgs']]]]:
        """
        This field is a list of servers that are applicable only when dynamic discovery is disabled. Server name is required
        only in cases where the new servers need to be created in this API. For existing servers, pass only the serverId.
        """
        return pulumi.get(self, "servers")

    @servers.setter
    def servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServerGroupServerArgs']]]]):
        pulumi.set(self, "servers", value)


@pulumi.input_type
class _ServerGroupState:
    def __init__(__self__, *,
                 app_connector_groups: Optional[pulumi.Input[Sequence[pulumi.Input['ServerGroupAppConnectorGroupArgs']]]] = None,
                 applications: Optional[pulumi.Input[Sequence[pulumi.Input['ServerGroupApplicationArgs']]]] = None,
                 config_space: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamic_discovery: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 ip_anchored: Optional[pulumi.Input[bool]] = None,
                 microtenant_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 servers: Optional[pulumi.Input[Sequence[pulumi.Input['ServerGroupServerArgs']]]] = None):
        """
        Input properties used for looking up and filtering ServerGroup resources.
        :param pulumi.Input[Sequence[pulumi.Input['ServerGroupAppConnectorGroupArgs']]] app_connector_groups: List of app-connector IDs.
        :param pulumi.Input[Sequence[pulumi.Input['ServerGroupApplicationArgs']]] applications: This field is a json array of app-connector-id only.
        :param pulumi.Input[str] description: This field is the description of the server group.
        :param pulumi.Input[bool] dynamic_discovery: This field controls dynamic discovery of the servers.
        :param pulumi.Input[bool] enabled: This field defines if the server group is enabled or disabled.
        :param pulumi.Input[str] name: This field defines the name of the server group.
        :param pulumi.Input[Sequence[pulumi.Input['ServerGroupServerArgs']]] servers: This field is a list of servers that are applicable only when dynamic discovery is disabled. Server name is required
               only in cases where the new servers need to be created in this API. For existing servers, pass only the serverId.
        """
        if app_connector_groups is not None:
            pulumi.set(__self__, "app_connector_groups", app_connector_groups)
        if applications is not None:
            pulumi.set(__self__, "applications", applications)
        if config_space is not None:
            pulumi.set(__self__, "config_space", config_space)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dynamic_discovery is not None:
            pulumi.set(__self__, "dynamic_discovery", dynamic_discovery)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if ip_anchored is not None:
            pulumi.set(__self__, "ip_anchored", ip_anchored)
        if microtenant_id is not None:
            pulumi.set(__self__, "microtenant_id", microtenant_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if servers is not None:
            pulumi.set(__self__, "servers", servers)

    @property
    @pulumi.getter(name="appConnectorGroups")
    def app_connector_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServerGroupAppConnectorGroupArgs']]]]:
        """
        List of app-connector IDs.
        """
        return pulumi.get(self, "app_connector_groups")

    @app_connector_groups.setter
    def app_connector_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServerGroupAppConnectorGroupArgs']]]]):
        pulumi.set(self, "app_connector_groups", value)

    @property
    @pulumi.getter
    def applications(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServerGroupApplicationArgs']]]]:
        """
        This field is a json array of app-connector-id only.
        """
        return pulumi.get(self, "applications")

    @applications.setter
    def applications(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServerGroupApplicationArgs']]]]):
        pulumi.set(self, "applications", value)

    @property
    @pulumi.getter(name="configSpace")
    def config_space(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "config_space")

    @config_space.setter
    def config_space(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_space", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        This field is the description of the server group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dynamicDiscovery")
    def dynamic_discovery(self) -> Optional[pulumi.Input[bool]]:
        """
        This field controls dynamic discovery of the servers.
        """
        return pulumi.get(self, "dynamic_discovery")

    @dynamic_discovery.setter
    def dynamic_discovery(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dynamic_discovery", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        This field defines if the server group is enabled or disabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="ipAnchored")
    def ip_anchored(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "ip_anchored")

    @ip_anchored.setter
    def ip_anchored(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ip_anchored", value)

    @property
    @pulumi.getter(name="microtenantId")
    def microtenant_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "microtenant_id")

    @microtenant_id.setter
    def microtenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "microtenant_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        This field defines the name of the server group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServerGroupServerArgs']]]]:
        """
        This field is a list of servers that are applicable only when dynamic discovery is disabled. Server name is required
        only in cases where the new servers need to be created in this API. For existing servers, pass only the serverId.
        """
        return pulumi.get(self, "servers")

    @servers.setter
    def servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServerGroupServerArgs']]]]):
        pulumi.set(self, "servers", value)


class ServerGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_connector_groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerGroupAppConnectorGroupArgs']]]]] = None,
                 applications: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerGroupApplicationArgs']]]]] = None,
                 config_space: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamic_discovery: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 ip_anchored: Optional[pulumi.Input[bool]] = None,
                 microtenant_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 servers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerGroupServerArgs']]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import zscaler_pulumi_zpa as zpa

        # Create a App Connector Group
        example_connector_group = zpa.ConnectorGroup("exampleConnectorGroup",
            description="Example",
            enabled=True,
            city_country="San Jose, CA",
            country_code="US",
            latitude="37.338",
            longitude="-121.8863",
            location="San Jose, CA, US",
            upgrade_day="SUNDAY",
            upgrade_time_in_secs="66600",
            override_version_profile=True,
            version_profile_id="0",
            dns_query_type="IPV4")
        # Create a Server Group resource with Dynamic Discovery Enabled
        example_server_group = zpa.ServerGroup("exampleServerGroup",
            description="Example",
            enabled=True,
            dynamic_discovery=True,
            app_connector_groups=[zpa.ServerGroupAppConnectorGroupArgs(
                ids=[example_connector_group.id],
            )],
            opts=pulumi.ResourceOptions(depends_on=[example_connector_group]))
        ```
        <!--End PulumiCodeChooser -->

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import zscaler_pulumi_zpa as zpa

        # Create an application server
        example_application_server = zpa.ApplicationServer("exampleApplicationServer",
            description="Example",
            address="server.example.com",
            enabled=True)
        # Create a App Connector Group
        example_connector_group = zpa.ConnectorGroup("exampleConnectorGroup",
            description="Example",
            enabled=True,
            city_country="San Jose, CA",
            country_code="US",
            latitude="37.338",
            longitude="-121.8863",
            location="San Jose, CA, US",
            upgrade_day="SUNDAY",
            upgrade_time_in_secs="66600",
            override_version_profile=True,
            version_profile_id="0",
            dns_query_type="IPV4")
        # ZPA Server Group resource with Dynamic Discovery Disabled
        example_server_group = zpa.ServerGroup("exampleServerGroup",
            description="Example",
            enabled=True,
            dynamic_discovery=False,
            servers=[zpa.ServerGroupServerArgs(
                ids=[example_application_server.id],
            )],
            app_connector_groups=[zpa.ServerGroupAppConnectorGroupArgs(
                ids=[example_connector_group.id],
            )],
            opts=pulumi.ResourceOptions(depends_on=[
                    example_connector_group,
                    zpa_application_server["server"],
                ]))
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        Server Groups can be imported; use `<SERVER GROUP ID>` or `<SERVER GROUP NAME>` as the import ID.

        For example:

        ```sh
        $ pulumi import zpa:index/serverGroup:ServerGroup example <server_group_id>
        ```

        or

        ```sh
        $ pulumi import zpa:index/serverGroup:ServerGroup example <server_group_name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerGroupAppConnectorGroupArgs']]]] app_connector_groups: List of app-connector IDs.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerGroupApplicationArgs']]]] applications: This field is a json array of app-connector-id only.
        :param pulumi.Input[str] description: This field is the description of the server group.
        :param pulumi.Input[bool] dynamic_discovery: This field controls dynamic discovery of the servers.
        :param pulumi.Input[bool] enabled: This field defines if the server group is enabled or disabled.
        :param pulumi.Input[str] name: This field defines the name of the server group.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerGroupServerArgs']]]] servers: This field is a list of servers that are applicable only when dynamic discovery is disabled. Server name is required
               only in cases where the new servers need to be created in this API. For existing servers, pass only the serverId.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ServerGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import zscaler_pulumi_zpa as zpa

        # Create a App Connector Group
        example_connector_group = zpa.ConnectorGroup("exampleConnectorGroup",
            description="Example",
            enabled=True,
            city_country="San Jose, CA",
            country_code="US",
            latitude="37.338",
            longitude="-121.8863",
            location="San Jose, CA, US",
            upgrade_day="SUNDAY",
            upgrade_time_in_secs="66600",
            override_version_profile=True,
            version_profile_id="0",
            dns_query_type="IPV4")
        # Create a Server Group resource with Dynamic Discovery Enabled
        example_server_group = zpa.ServerGroup("exampleServerGroup",
            description="Example",
            enabled=True,
            dynamic_discovery=True,
            app_connector_groups=[zpa.ServerGroupAppConnectorGroupArgs(
                ids=[example_connector_group.id],
            )],
            opts=pulumi.ResourceOptions(depends_on=[example_connector_group]))
        ```
        <!--End PulumiCodeChooser -->

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import zscaler_pulumi_zpa as zpa

        # Create an application server
        example_application_server = zpa.ApplicationServer("exampleApplicationServer",
            description="Example",
            address="server.example.com",
            enabled=True)
        # Create a App Connector Group
        example_connector_group = zpa.ConnectorGroup("exampleConnectorGroup",
            description="Example",
            enabled=True,
            city_country="San Jose, CA",
            country_code="US",
            latitude="37.338",
            longitude="-121.8863",
            location="San Jose, CA, US",
            upgrade_day="SUNDAY",
            upgrade_time_in_secs="66600",
            override_version_profile=True,
            version_profile_id="0",
            dns_query_type="IPV4")
        # ZPA Server Group resource with Dynamic Discovery Disabled
        example_server_group = zpa.ServerGroup("exampleServerGroup",
            description="Example",
            enabled=True,
            dynamic_discovery=False,
            servers=[zpa.ServerGroupServerArgs(
                ids=[example_application_server.id],
            )],
            app_connector_groups=[zpa.ServerGroupAppConnectorGroupArgs(
                ids=[example_connector_group.id],
            )],
            opts=pulumi.ResourceOptions(depends_on=[
                    example_connector_group,
                    zpa_application_server["server"],
                ]))
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        Server Groups can be imported; use `<SERVER GROUP ID>` or `<SERVER GROUP NAME>` as the import ID.

        For example:

        ```sh
        $ pulumi import zpa:index/serverGroup:ServerGroup example <server_group_id>
        ```

        or

        ```sh
        $ pulumi import zpa:index/serverGroup:ServerGroup example <server_group_name>
        ```

        :param str resource_name: The name of the resource.
        :param ServerGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServerGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_connector_groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerGroupAppConnectorGroupArgs']]]]] = None,
                 applications: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerGroupApplicationArgs']]]]] = None,
                 config_space: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamic_discovery: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 ip_anchored: Optional[pulumi.Input[bool]] = None,
                 microtenant_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 servers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerGroupServerArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServerGroupArgs.__new__(ServerGroupArgs)

            __props__.__dict__["app_connector_groups"] = app_connector_groups
            __props__.__dict__["applications"] = applications
            __props__.__dict__["config_space"] = config_space
            __props__.__dict__["description"] = description
            __props__.__dict__["dynamic_discovery"] = dynamic_discovery
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["ip_anchored"] = ip_anchored
            __props__.__dict__["microtenant_id"] = microtenant_id
            __props__.__dict__["name"] = name
            __props__.__dict__["servers"] = servers
        super(ServerGroup, __self__).__init__(
            'zpa:index/serverGroup:ServerGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_connector_groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerGroupAppConnectorGroupArgs']]]]] = None,
            applications: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerGroupApplicationArgs']]]]] = None,
            config_space: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            dynamic_discovery: Optional[pulumi.Input[bool]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            ip_anchored: Optional[pulumi.Input[bool]] = None,
            microtenant_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            servers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerGroupServerArgs']]]]] = None) -> 'ServerGroup':
        """
        Get an existing ServerGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerGroupAppConnectorGroupArgs']]]] app_connector_groups: List of app-connector IDs.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerGroupApplicationArgs']]]] applications: This field is a json array of app-connector-id only.
        :param pulumi.Input[str] description: This field is the description of the server group.
        :param pulumi.Input[bool] dynamic_discovery: This field controls dynamic discovery of the servers.
        :param pulumi.Input[bool] enabled: This field defines if the server group is enabled or disabled.
        :param pulumi.Input[str] name: This field defines the name of the server group.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerGroupServerArgs']]]] servers: This field is a list of servers that are applicable only when dynamic discovery is disabled. Server name is required
               only in cases where the new servers need to be created in this API. For existing servers, pass only the serverId.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServerGroupState.__new__(_ServerGroupState)

        __props__.__dict__["app_connector_groups"] = app_connector_groups
        __props__.__dict__["applications"] = applications
        __props__.__dict__["config_space"] = config_space
        __props__.__dict__["description"] = description
        __props__.__dict__["dynamic_discovery"] = dynamic_discovery
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["ip_anchored"] = ip_anchored
        __props__.__dict__["microtenant_id"] = microtenant_id
        __props__.__dict__["name"] = name
        __props__.__dict__["servers"] = servers
        return ServerGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appConnectorGroups")
    def app_connector_groups(self) -> pulumi.Output[Sequence['outputs.ServerGroupAppConnectorGroup']]:
        """
        List of app-connector IDs.
        """
        return pulumi.get(self, "app_connector_groups")

    @property
    @pulumi.getter
    def applications(self) -> pulumi.Output[Sequence['outputs.ServerGroupApplication']]:
        """
        This field is a json array of app-connector-id only.
        """
        return pulumi.get(self, "applications")

    @property
    @pulumi.getter(name="configSpace")
    def config_space(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "config_space")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        This field is the description of the server group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dynamicDiscovery")
    def dynamic_discovery(self) -> pulumi.Output[Optional[bool]]:
        """
        This field controls dynamic discovery of the servers.
        """
        return pulumi.get(self, "dynamic_discovery")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        This field defines if the server group is enabled or disabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="ipAnchored")
    def ip_anchored(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "ip_anchored")

    @property
    @pulumi.getter(name="microtenantId")
    def microtenant_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "microtenant_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        This field defines the name of the server group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def servers(self) -> pulumi.Output[Sequence['outputs.ServerGroupServer']]:
        """
        This field is a list of servers that are applicable only when dynamic discovery is disabled. Server name is required
        only in cases where the new servers need to be created in this API. For existing servers, pass only the serverId.
        """
        return pulumi.get(self, "servers")

