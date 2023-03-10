// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.CloudConnectorGroup
{
    public static class GetCloudConnectorGroup
    {
        /// <summary>
        /// Use the **zpa_cloud_connector_group** data source to get information about a cloud connector group created from the Zscaler Private Access cloud. This data source can then be referenced within an Access Policy rule
        /// 
        /// &gt; **NOTE:** A Cloud Connector Group resource is created in the Zscaler Cloud Connector cloud and replicated to the ZPA cloud. This resource can then be referenced in a Access Policy Rule where the Object Type = `CLOUD_CONNECTOR_GROUP` is being used.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Zpa.CloudConnectorGroup.GetCloudConnectorGroup.Invoke(new()
        ///     {
        ///         Name = "AWS-Cloud",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Zpa.CloudConnectorGroup.GetCloudConnectorGroup.Invoke(new()
        ///     {
        ///         Id = "1234567890",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCloudConnectorGroupResult> InvokeAsync(GetCloudConnectorGroupArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCloudConnectorGroupResult>("zpa:CloudConnectorGroup/getCloudConnectorGroup:getCloudConnectorGroup", args ?? new GetCloudConnectorGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zpa_cloud_connector_group** data source to get information about a cloud connector group created from the Zscaler Private Access cloud. This data source can then be referenced within an Access Policy rule
        /// 
        /// &gt; **NOTE:** A Cloud Connector Group resource is created in the Zscaler Cloud Connector cloud and replicated to the ZPA cloud. This resource can then be referenced in a Access Policy Rule where the Object Type = `CLOUD_CONNECTOR_GROUP` is being used.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Zpa.CloudConnectorGroup.GetCloudConnectorGroup.Invoke(new()
        ///     {
        ///         Name = "AWS-Cloud",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Zpa.CloudConnectorGroup.GetCloudConnectorGroup.Invoke(new()
        ///     {
        ///         Id = "1234567890",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCloudConnectorGroupResult> Invoke(GetCloudConnectorGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCloudConnectorGroupResult>("zpa:CloudConnectorGroup/getCloudConnectorGroup:getCloudConnectorGroup", args ?? new GetCloudConnectorGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudConnectorGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// This field defines the id of the cloud connector group.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// This field defines the name of the cloud connector group.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetCloudConnectorGroupArgs()
        {
        }
        public static new GetCloudConnectorGroupArgs Empty => new GetCloudConnectorGroupArgs();
    }

    public sealed class GetCloudConnectorGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// This field defines the id of the cloud connector group.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// This field defines the name of the cloud connector group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetCloudConnectorGroupInvokeArgs()
        {
        }
        public static new GetCloudConnectorGroupInvokeArgs Empty => new GetCloudConnectorGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloudConnectorGroupResult
    {
        /// <summary>
        /// (string) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCloudConnectorGroupCloudConnectorResult> CloudConnectors;
        /// <summary>
        /// (string) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// (string) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (bool) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// (string) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
        /// </summary>
        public readonly string GeolocationId;
        public readonly string? Id;
        /// <summary>
        /// (string)- Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
        /// </summary>
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        /// <summary>
        /// (string) - This field defines the name of the cloud connector group.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// (string) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
        /// </summary>
        public readonly string ZiaCloud;
        /// <summary>
        /// (string) - Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
        /// </summary>
        public readonly string ZiaOrgId;

        [OutputConstructor]
        private GetCloudConnectorGroupResult(
            ImmutableArray<Outputs.GetCloudConnectorGroupCloudConnectorResult> cloudConnectors,

            string creationTime,

            string description,

            bool enabled,

            string geolocationId,

            string? id,

            string modifiedTime,

            string modifiedby,

            string? name,

            string ziaCloud,

            string ziaOrgId)
        {
            CloudConnectors = cloudConnectors;
            CreationTime = creationTime;
            Description = description;
            Enabled = enabled;
            GeolocationId = geolocationId;
            Id = id;
            ModifiedTime = modifiedTime;
            Modifiedby = modifiedby;
            Name = name;
            ZiaCloud = ziaCloud;
            ZiaOrgId = ziaOrgId;
        }
    }
}
