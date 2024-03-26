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
    public static class GetProvisioningKey
    {
        /// <summary>
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
        ///     var example = Zpa.GetProvisioningKey.Invoke(new()
        ///     {
        ///         AssociationType = "CONNECTOR_GRP",
        ///         Name = "Provisioning_Key",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
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
        ///     var example = Zpa.GetProvisioningKey.Invoke(new()
        ///     {
        ///         AssociationType = "SERVICE_EDGE_GRP",
        ///         Name = "Provisioning_Key",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetProvisioningKeyResult> InvokeAsync(GetProvisioningKeyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProvisioningKeyResult>("zpa:index/getProvisioningKey:getProvisioningKey", args ?? new GetProvisioningKeyArgs(), options.WithDefaults());

        /// <summary>
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
        ///     var example = Zpa.GetProvisioningKey.Invoke(new()
        ///     {
        ///         AssociationType = "CONNECTOR_GRP",
        ///         Name = "Provisioning_Key",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
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
        ///     var example = Zpa.GetProvisioningKey.Invoke(new()
        ///     {
        ///         AssociationType = "SERVICE_EDGE_GRP",
        ///         Name = "Provisioning_Key",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetProvisioningKeyResult> Invoke(GetProvisioningKeyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProvisioningKeyResult>("zpa:index/getProvisioningKey:getProvisioningKey", args ?? new GetProvisioningKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProvisioningKeyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the provisioning key type for App Connectors or ZPA Private Service Edges. The supported values are `CONNECTOR_GRP` and `SERVICE_EDGE_GRP`
        /// </summary>
        [Input("associationType", required: true)]
        public string AssociationType { get; set; } = null!;

        /// <summary>
        /// The ID of the provisioning key to be exported.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// (string) The ID of the microtenant the resource is to be associated with.
        /// </summary>
        [Input("microtenantId")]
        public string? MicrotenantId { get; set; }

        /// <summary>
        /// (string) The name of the microtenant the resource is to be associated with.
        /// </summary>
        [Input("microtenantName")]
        public string? MicrotenantName { get; set; }

        /// <summary>
        /// Name of the provisioning key.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetProvisioningKeyArgs()
        {
        }
        public static new GetProvisioningKeyArgs Empty => new GetProvisioningKeyArgs();
    }

    public sealed class GetProvisioningKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the provisioning key type for App Connectors or ZPA Private Service Edges. The supported values are `CONNECTOR_GRP` and `SERVICE_EDGE_GRP`
        /// </summary>
        [Input("associationType", required: true)]
        public Input<string> AssociationType { get; set; } = null!;

        /// <summary>
        /// The ID of the provisioning key to be exported.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// (string) The ID of the microtenant the resource is to be associated with.
        /// </summary>
        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// (string) The name of the microtenant the resource is to be associated with.
        /// </summary>
        [Input("microtenantName")]
        public Input<string>? MicrotenantName { get; set; }

        /// <summary>
        /// Name of the provisioning key.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetProvisioningKeyInvokeArgs()
        {
        }
        public static new GetProvisioningKeyInvokeArgs Empty => new GetProvisioningKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetProvisioningKeyResult
    {
        public readonly string AppConnectorGroupId;
        public readonly string AppConnectorGroupName;
        public readonly string AssociationType;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// (bool)
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string EnrollmentCertId;
        /// <summary>
        /// (string) Applicable only for GET calls, ignored in PUT/POST calls.
        /// </summary>
        public readonly string EnrollmentCertName;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ExpirationInEpochSec;
        public readonly string? Id;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly ImmutableArray<string> IpAcls;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string MaxUsage;
        /// <summary>
        /// (string) The ID of the microtenant the resource is to be associated with.
        /// </summary>
        public readonly string? MicrotenantId;
        /// <summary>
        /// (string) The name of the microtenant the resource is to be associated with.
        /// </summary>
        public readonly string? MicrotenantName;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        public readonly string? Name;
        /// <summary>
        /// (string) Ignored in PUT/POST calls.
        /// </summary>
        public readonly string ProvisioningKey;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string UiConfig;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string UsageCount;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ZcomponentId;
        /// <summary>
        /// (string) Applicable only for GET calls, ignored in PUT/POST calls.
        /// </summary>
        public readonly string ZcomponentName;

        [OutputConstructor]
        private GetProvisioningKeyResult(
            string appConnectorGroupId,

            string appConnectorGroupName,

            string associationType,

            string creationTime,

            bool enabled,

            string enrollmentCertId,

            string enrollmentCertName,

            string expirationInEpochSec,

            string? id,

            ImmutableArray<string> ipAcls,

            string maxUsage,

            string? microtenantId,

            string? microtenantName,

            string modifiedTime,

            string modifiedby,

            string? name,

            string provisioningKey,

            string uiConfig,

            string usageCount,

            string zcomponentId,

            string zcomponentName)
        {
            AppConnectorGroupId = appConnectorGroupId;
            AppConnectorGroupName = appConnectorGroupName;
            AssociationType = associationType;
            CreationTime = creationTime;
            Enabled = enabled;
            EnrollmentCertId = enrollmentCertId;
            EnrollmentCertName = enrollmentCertName;
            ExpirationInEpochSec = expirationInEpochSec;
            Id = id;
            IpAcls = ipAcls;
            MaxUsage = maxUsage;
            MicrotenantId = microtenantId;
            MicrotenantName = microtenantName;
            ModifiedTime = modifiedTime;
            Modifiedby = modifiedby;
            Name = name;
            ProvisioningKey = provisioningKey;
            UiConfig = uiConfig;
            UsageCount = usageCount;
            ZcomponentId = zcomponentId;
            ZcomponentName = zcomponentName;
        }
    }
}
