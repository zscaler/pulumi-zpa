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
    public static class GetCloudConnectorGroup
    {
        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/about-web-server-certificates)
        /// * [API documentation](https://help.zscaler.com/zpa/obtaining-cloud-connector-group-details-using-api)
        /// 
        /// Use the **zpa_cloud_connector_group** data source to get information about a cloud connector group created from the Zscaler Private Access cloud. This data source can then be referenced within an Access Policy rule
        /// 
        /// &gt; **NOTE:** A Cloud Connector Group resource is created in the Zscaler Cloud Connector cloud and replicated to the ZPA cloud. This resource can then be referenced in a Access Policy Rule where the Object Type = `EDGE_CONNECTOR_GROUP` is being used.
        /// 
        /// **NOTE:** To ensure consistent search results across data sources, please avoid using multiple spaces or special characters in your search queries.
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
        ///     var foo = Zpa.GetCloudConnectorGroup.Invoke(new()
        ///     {
        ///         Name = "AWS-Cloud",
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
        ///     var foo = Zpa.GetCloudConnectorGroup.Invoke(new()
        ///     {
        ///         Id = "1234567890",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetCloudConnectorGroupResult> InvokeAsync(GetCloudConnectorGroupArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCloudConnectorGroupResult>("zpa:index/getCloudConnectorGroup:getCloudConnectorGroup", args ?? new GetCloudConnectorGroupArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/about-web-server-certificates)
        /// * [API documentation](https://help.zscaler.com/zpa/obtaining-cloud-connector-group-details-using-api)
        /// 
        /// Use the **zpa_cloud_connector_group** data source to get information about a cloud connector group created from the Zscaler Private Access cloud. This data source can then be referenced within an Access Policy rule
        /// 
        /// &gt; **NOTE:** A Cloud Connector Group resource is created in the Zscaler Cloud Connector cloud and replicated to the ZPA cloud. This resource can then be referenced in a Access Policy Rule where the Object Type = `EDGE_CONNECTOR_GROUP` is being used.
        /// 
        /// **NOTE:** To ensure consistent search results across data sources, please avoid using multiple spaces or special characters in your search queries.
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
        ///     var foo = Zpa.GetCloudConnectorGroup.Invoke(new()
        ///     {
        ///         Name = "AWS-Cloud",
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
        ///     var foo = Zpa.GetCloudConnectorGroup.Invoke(new()
        ///     {
        ///         Id = "1234567890",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCloudConnectorGroupResult> Invoke(GetCloudConnectorGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCloudConnectorGroupResult>("zpa:index/getCloudConnectorGroup:getCloudConnectorGroup", args ?? new GetCloudConnectorGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/about-web-server-certificates)
        /// * [API documentation](https://help.zscaler.com/zpa/obtaining-cloud-connector-group-details-using-api)
        /// 
        /// Use the **zpa_cloud_connector_group** data source to get information about a cloud connector group created from the Zscaler Private Access cloud. This data source can then be referenced within an Access Policy rule
        /// 
        /// &gt; **NOTE:** A Cloud Connector Group resource is created in the Zscaler Cloud Connector cloud and replicated to the ZPA cloud. This resource can then be referenced in a Access Policy Rule where the Object Type = `EDGE_CONNECTOR_GROUP` is being used.
        /// 
        /// **NOTE:** To ensure consistent search results across data sources, please avoid using multiple spaces or special characters in your search queries.
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
        ///     var foo = Zpa.GetCloudConnectorGroup.Invoke(new()
        ///     {
        ///         Name = "AWS-Cloud",
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
        ///     var foo = Zpa.GetCloudConnectorGroup.Invoke(new()
        ///     {
        ///         Id = "1234567890",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCloudConnectorGroupResult> Invoke(GetCloudConnectorGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCloudConnectorGroupResult>("zpa:index/getCloudConnectorGroup:getCloudConnectorGroup", args ?? new GetCloudConnectorGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudConnectorGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public string? Id { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public GetCloudConnectorGroupArgs()
        {
        }
        public static new GetCloudConnectorGroupArgs Empty => new GetCloudConnectorGroupArgs();
    }

    public sealed class GetCloudConnectorGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

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
        public readonly ImmutableArray<Outputs.GetCloudConnectorGroupCloudConnectorResult> CloudConnectors;
        public readonly string CreationTime;
        public readonly string Description;
        public readonly bool Enabled;
        public readonly string GeolocationId;
        public readonly string? Id;
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        public readonly string? Name;
        public readonly string ZiaCloud;
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
