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
    public static class GetProvisioningKey
    {
        /// <summary>
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
        ///     var example = Zpa.GetProvisioningKey.Invoke(new()
        ///     {
        ///         AssociationType = "CONNECTOR_GRP",
        ///         Name = "Provisioning_Key",
        ///     });
        /// 
        /// });
        /// ```
        /// 
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
        /// </summary>
        public static Task<GetProvisioningKeyResult> InvokeAsync(GetProvisioningKeyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProvisioningKeyResult>("zpa:index/getProvisioningKey:getProvisioningKey", args ?? new GetProvisioningKeyArgs(), options.WithDefaults());

        /// <summary>
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
        ///     var example = Zpa.GetProvisioningKey.Invoke(new()
        ///     {
        ///         AssociationType = "CONNECTOR_GRP",
        ///         Name = "Provisioning_Key",
        ///     });
        /// 
        /// });
        /// ```
        /// 
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
        /// </summary>
        public static Output<GetProvisioningKeyResult> Invoke(GetProvisioningKeyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProvisioningKeyResult>("zpa:index/getProvisioningKey:getProvisioningKey", args ?? new GetProvisioningKeyInvokeArgs(), options.WithDefaults());

        /// <summary>
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
        ///     var example = Zpa.GetProvisioningKey.Invoke(new()
        ///     {
        ///         AssociationType = "CONNECTOR_GRP",
        ///         Name = "Provisioning_Key",
        ///     });
        /// 
        /// });
        /// ```
        /// 
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
        /// </summary>
        public static Output<GetProvisioningKeyResult> Invoke(GetProvisioningKeyInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetProvisioningKeyResult>("zpa:index/getProvisioningKey:getProvisioningKey", args ?? new GetProvisioningKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProvisioningKeyArgs : global::Pulumi.InvokeArgs
    {
        [Input("associationType", required: true)]
        public string AssociationType { get; set; } = null!;

        [Input("id")]
        public string? Id { get; set; }

        [Input("microtenantId")]
        public string? MicrotenantId { get; set; }

        [Input("microtenantName")]
        public string? MicrotenantName { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public GetProvisioningKeyArgs()
        {
        }
        public static new GetProvisioningKeyArgs Empty => new GetProvisioningKeyArgs();
    }

    public sealed class GetProvisioningKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("associationType", required: true)]
        public Input<string> AssociationType { get; set; } = null!;

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        [Input("microtenantName")]
        public Input<string>? MicrotenantName { get; set; }

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
        public readonly string CreationTime;
        public readonly bool Enabled;
        public readonly string EnrollmentCertId;
        public readonly string EnrollmentCertName;
        public readonly string ExpirationInEpochSec;
        public readonly string? Id;
        public readonly ImmutableArray<string> IpAcls;
        public readonly string MaxUsage;
        public readonly string? MicrotenantId;
        public readonly string? MicrotenantName;
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        public readonly string? Name;
        public readonly string ProvisioningKey;
        public readonly string UiConfig;
        public readonly string UsageCount;
        public readonly string ZcomponentId;
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
