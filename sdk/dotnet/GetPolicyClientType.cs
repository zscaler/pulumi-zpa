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
    public static class GetPolicyClientType
    {
        /// <summary>
        /// Use the **zpa_access_policy_client_types** data source to get information about all client types for the specified customer in the Zscaler Private Access cloud. This data source can be optionally used when defining the following policy types:
        ///     - ``zpa.PolicyAccessRule``
        ///     - ``zpa.PolicyAccessTimeOutRule``
        ///     - ``zpa.PolicyAccessForwardingRule``
        ///     - ``zpa.PolicyAccessIsolationRule``
        ///     - ``zpa.PolicyAccessInspectionRule``
        /// 
        /// The ``object_type`` attribute must be defined as "CLIENT_TYPE" in the policy operand condition. To learn more see the To learn more see the [Getting Details of All Client Types](https://help.zscaler.com/zpa/configuring-access-policies-using-api#getClientTypes)
        /// 
        /// &gt; **NOTE** By Default the ZPA provider will return all client types
        /// 
        /// &gt; **NOTE** When defining a ``zpa.PolicyAccessIsolationRule`` policy the ``object_type`` "CLIENT_TYPE" is mandatory and ``zpn_client_type_exporter`` is the only supported value.
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
        ///     var @this = Zpa.GetPolicyClientType.Invoke();
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetPolicyClientTypeResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPolicyClientTypeResult>("zpa:index/getPolicyClientType:getPolicyClientType", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Use the **zpa_access_policy_client_types** data source to get information about all client types for the specified customer in the Zscaler Private Access cloud. This data source can be optionally used when defining the following policy types:
        ///     - ``zpa.PolicyAccessRule``
        ///     - ``zpa.PolicyAccessTimeOutRule``
        ///     - ``zpa.PolicyAccessForwardingRule``
        ///     - ``zpa.PolicyAccessIsolationRule``
        ///     - ``zpa.PolicyAccessInspectionRule``
        /// 
        /// The ``object_type`` attribute must be defined as "CLIENT_TYPE" in the policy operand condition. To learn more see the To learn more see the [Getting Details of All Client Types](https://help.zscaler.com/zpa/configuring-access-policies-using-api#getClientTypes)
        /// 
        /// &gt; **NOTE** By Default the ZPA provider will return all client types
        /// 
        /// &gt; **NOTE** When defining a ``zpa.PolicyAccessIsolationRule`` policy the ``object_type`` "CLIENT_TYPE" is mandatory and ``zpn_client_type_exporter`` is the only supported value.
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
        ///     var @this = Zpa.GetPolicyClientType.Invoke();
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetPolicyClientTypeResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPolicyClientTypeResult>("zpa:index/getPolicyClientType:getPolicyClientType", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetPolicyClientTypeResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ZpnClientTypeBranchConnector;
        public readonly string ZpnClientTypeBrowserIsolation;
        public readonly string ZpnClientTypeEdgeConnector;
        public readonly string ZpnClientTypeExporter;
        public readonly string ZpnClientTypeExporterNoauth;
        public readonly string ZpnClientTypeIpAnchoring;
        public readonly string ZpnClientTypeMachineTunnel;
        public readonly string ZpnClientTypeSlogger;
        public readonly string ZpnClientTypeZapp;

        [OutputConstructor]
        private GetPolicyClientTypeResult(
            string id,

            string zpnClientTypeBranchConnector,

            string zpnClientTypeBrowserIsolation,

            string zpnClientTypeEdgeConnector,

            string zpnClientTypeExporter,

            string zpnClientTypeExporterNoauth,

            string zpnClientTypeIpAnchoring,

            string zpnClientTypeMachineTunnel,

            string zpnClientTypeSlogger,

            string zpnClientTypeZapp)
        {
            Id = id;
            ZpnClientTypeBranchConnector = zpnClientTypeBranchConnector;
            ZpnClientTypeBrowserIsolation = zpnClientTypeBrowserIsolation;
            ZpnClientTypeEdgeConnector = zpnClientTypeEdgeConnector;
            ZpnClientTypeExporter = zpnClientTypeExporter;
            ZpnClientTypeExporterNoauth = zpnClientTypeExporterNoauth;
            ZpnClientTypeIpAnchoring = zpnClientTypeIpAnchoring;
            ZpnClientTypeMachineTunnel = zpnClientTypeMachineTunnel;
            ZpnClientTypeSlogger = zpnClientTypeSlogger;
            ZpnClientTypeZapp = zpnClientTypeZapp;
        }
    }
}