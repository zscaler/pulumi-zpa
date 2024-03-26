// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Zscaler.Zpa
{
    public static class GetLSSClientTypes
    {
        /// <summary>
        /// Use the **zpa_lss_config_client_types** data source to get information about all LSS client types in the Zscaler Private Access cloud. This data source is required when the defining a policy rule resource for an object type as `CLIENT_TYPE` parameter in the LSS Config Controller resource is set. To learn more see the To learn more see the [Getting Details of All LSS Status Codes](https://help.zscaler.com/zpa/log-streaming-service-configuration-use-cases#GettingLSSClientTypes)
        /// 
        /// &gt; **NOTE** By Default the ZPA provider will return all client types
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
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
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetLSSClientTypesResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLSSClientTypesResult>("zpa:index/getLSSClientTypes:getLSSClientTypes", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Use the **zpa_lss_config_client_types** data source to get information about all LSS client types in the Zscaler Private Access cloud. This data source is required when the defining a policy rule resource for an object type as `CLIENT_TYPE` parameter in the LSS Config Controller resource is set. To learn more see the To learn more see the [Getting Details of All LSS Status Codes](https://help.zscaler.com/zpa/log-streaming-service-configuration-use-cases#GettingLSSClientTypes)
        /// 
        /// &gt; **NOTE** By Default the ZPA provider will return all client types
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
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
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetLSSClientTypesResult> Invoke(InvokeOptions? options = null)
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
