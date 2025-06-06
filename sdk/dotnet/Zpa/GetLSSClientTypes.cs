// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa
{
    public static class GetLSSClientTypes
    {
        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/about-log-streaming-service)
        /// * [API documentation](https://help.zscaler.com/zpa/configuring-log-streaming-service-configurations-using-api)
        /// 
        /// Use the **zpa_lss_config_client_types** data source to get information about all LSS client types in the Zscaler Private Access cloud. This data source is required when the defining a policy rule resource for an object type as `CLIENT_TYPE` parameter in the LSS Config Controller resource is set. To learn more see the To learn more see the [Getting Details of All LSS Status Codes](https://help.zscaler.com/zpa/log-streaming-service-configuration-use-cases#GettingLSSClientTypes)
        /// 
        /// &gt; **NOTE** By Default the ZPA provider will return all client types
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zpa.GetLSSClientTypes.Invoke();
        /// 
        /// });
        /// ```
        /// 
        /// ### Read-Only
        /// 
        /// The following arguments are supported:
        /// 
        /// * `"zpn_client_type_edge_connector" = "Cloud Connector"`
        /// * `"zpn_client_type_exporter" = "Web Browser`
        /// * `"zpn_client_type_ip_anchoring" = "ZIA Service Edge"`
        /// * `"zpn_client_type_machine_tunnel" = "Machine Tunnel"`
        /// * `"zpn_client_type_slogger" = "ZPA LSS"`
        /// * `"zpn_client_type_zapp" = "Client Connector"`
        /// 
        /// To learn more see the [Getting Details of All LSS Status Codes](https://help.zscaler.com/zpa/log-streaming-service-configuration-use-cases#GettingLSSClientTypes)
        /// </summary>
        public static Task<GetLSSClientTypesResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLSSClientTypesResult>("zpa:index/getLSSClientTypes:getLSSClientTypes", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/about-log-streaming-service)
        /// * [API documentation](https://help.zscaler.com/zpa/configuring-log-streaming-service-configurations-using-api)
        /// 
        /// Use the **zpa_lss_config_client_types** data source to get information about all LSS client types in the Zscaler Private Access cloud. This data source is required when the defining a policy rule resource for an object type as `CLIENT_TYPE` parameter in the LSS Config Controller resource is set. To learn more see the To learn more see the [Getting Details of All LSS Status Codes](https://help.zscaler.com/zpa/log-streaming-service-configuration-use-cases#GettingLSSClientTypes)
        /// 
        /// &gt; **NOTE** By Default the ZPA provider will return all client types
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zpa.GetLSSClientTypes.Invoke();
        /// 
        /// });
        /// ```
        /// 
        /// ### Read-Only
        /// 
        /// The following arguments are supported:
        /// 
        /// * `"zpn_client_type_edge_connector" = "Cloud Connector"`
        /// * `"zpn_client_type_exporter" = "Web Browser`
        /// * `"zpn_client_type_ip_anchoring" = "ZIA Service Edge"`
        /// * `"zpn_client_type_machine_tunnel" = "Machine Tunnel"`
        /// * `"zpn_client_type_slogger" = "ZPA LSS"`
        /// * `"zpn_client_type_zapp" = "Client Connector"`
        /// 
        /// To learn more see the [Getting Details of All LSS Status Codes](https://help.zscaler.com/zpa/log-streaming-service-configuration-use-cases#GettingLSSClientTypes)
        /// </summary>
        public static Output<GetLSSClientTypesResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLSSClientTypesResult>("zpa:index/getLSSClientTypes:getLSSClientTypes", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/about-log-streaming-service)
        /// * [API documentation](https://help.zscaler.com/zpa/configuring-log-streaming-service-configurations-using-api)
        /// 
        /// Use the **zpa_lss_config_client_types** data source to get information about all LSS client types in the Zscaler Private Access cloud. This data source is required when the defining a policy rule resource for an object type as `CLIENT_TYPE` parameter in the LSS Config Controller resource is set. To learn more see the To learn more see the [Getting Details of All LSS Status Codes](https://help.zscaler.com/zpa/log-streaming-service-configuration-use-cases#GettingLSSClientTypes)
        /// 
        /// &gt; **NOTE** By Default the ZPA provider will return all client types
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zpa.GetLSSClientTypes.Invoke();
        /// 
        /// });
        /// ```
        /// 
        /// ### Read-Only
        /// 
        /// The following arguments are supported:
        /// 
        /// * `"zpn_client_type_edge_connector" = "Cloud Connector"`
        /// * `"zpn_client_type_exporter" = "Web Browser`
        /// * `"zpn_client_type_ip_anchoring" = "ZIA Service Edge"`
        /// * `"zpn_client_type_machine_tunnel" = "Machine Tunnel"`
        /// * `"zpn_client_type_slogger" = "ZPA LSS"`
        /// * `"zpn_client_type_zapp" = "Client Connector"`
        /// 
        /// To learn more see the [Getting Details of All LSS Status Codes](https://help.zscaler.com/zpa/log-streaming-service-configuration-use-cases#GettingLSSClientTypes)
        /// </summary>
        public static Output<GetLSSClientTypesResult> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetLSSClientTypesResult>("zpa:index/getLSSClientTypes:getLSSClientTypes", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetLSSClientTypesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ZpnClientTypeEdgeConnector;
        public readonly string ZpnClientTypeExporter;
        public readonly string ZpnClientTypeIpAnchoring;
        public readonly string ZpnClientTypeMachineTunnel;
        public readonly string ZpnClientTypeSlogger;
        public readonly string ZpnClientTypeZapp;

        [OutputConstructor]
        private GetLSSClientTypesResult(
            string id,

            string zpnClientTypeEdgeConnector,

            string zpnClientTypeExporter,

            string zpnClientTypeIpAnchoring,

            string zpnClientTypeMachineTunnel,

            string zpnClientTypeSlogger,

            string zpnClientTypeZapp)
        {
            Id = id;
            ZpnClientTypeEdgeConnector = zpnClientTypeEdgeConnector;
            ZpnClientTypeExporter = zpnClientTypeExporter;
            ZpnClientTypeIpAnchoring = zpnClientTypeIpAnchoring;
            ZpnClientTypeMachineTunnel = zpnClientTypeMachineTunnel;
            ZpnClientTypeSlogger = zpnClientTypeSlogger;
            ZpnClientTypeZapp = zpnClientTypeZapp;
        }
    }
}
