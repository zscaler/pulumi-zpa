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
    public static class GetAppConnectorGroup
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
        ///     var foo = Zpa.GetAppConnectorGroup.Invoke(new()
        ///     {
        ///         Name = "DataCenter",
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
        ///     var foo = Zpa.GetAppConnectorGroup.Invoke(new()
        ///     {
        ///         Id = "123456789",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAppConnectorGroupResult> InvokeAsync(GetAppConnectorGroupArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAppConnectorGroupResult>("zpa:index/getAppConnectorGroup:getAppConnectorGroup", args ?? new GetAppConnectorGroupArgs(), options.WithDefaults());

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
        ///     var foo = Zpa.GetAppConnectorGroup.Invoke(new()
        ///     {
        ///         Name = "DataCenter",
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
        ///     var foo = Zpa.GetAppConnectorGroup.Invoke(new()
        ///     {
        ///         Id = "123456789",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAppConnectorGroupResult> Invoke(GetAppConnectorGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAppConnectorGroupResult>("zpa:index/getAppConnectorGroup:getAppConnectorGroup", args ?? new GetAppConnectorGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppConnectorGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public string? Id { get; set; }

        [Input("microtenantId")]
        public string? MicrotenantId { get; set; }

        [Input("microtenantName")]
        public string? MicrotenantName { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("overrideVersionProfile")]
        public bool? OverrideVersionProfile { get; set; }

        public GetAppConnectorGroupArgs()
        {
        }
        public static new GetAppConnectorGroupArgs Empty => new GetAppConnectorGroupArgs();
    }

    public sealed class GetAppConnectorGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        [Input("microtenantName")]
        public Input<string>? MicrotenantName { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("overrideVersionProfile")]
        public Input<bool>? OverrideVersionProfile { get; set; }

        public GetAppConnectorGroupInvokeArgs()
        {
        }
        public static new GetAppConnectorGroupInvokeArgs Empty => new GetAppConnectorGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetAppConnectorGroupResult
    {
        public readonly string CityCountry;
        public readonly ImmutableArray<Outputs.GetAppConnectorGroupConnectorResult> Connectors;
        public readonly string CountryCode;
        public readonly string CreationTime;
        public readonly string Description;
        public readonly string DnsQueryType;
        public readonly bool Enabled;
        public readonly string GeoLocationId;
        public readonly string? Id;
        public readonly string Latitude;
        public readonly string Location;
        public readonly string Longitude;
        public readonly bool LssAppConnectorGroup;
        public readonly string? MicrotenantId;
        public readonly string? MicrotenantName;
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        public readonly string? Name;
        public readonly bool? OverrideVersionProfile;
        public readonly ImmutableArray<Outputs.GetAppConnectorGroupServerGroupResult> ServerGroups;
        public readonly bool TcpQuickAckApp;
        public readonly bool TcpQuickAckAssistant;
        public readonly bool TcpQuickAckReadAssistant;
        public readonly string UpgradeDay;
        public readonly string UpgradeTimeInSecs;
        public readonly bool UseInDrMode;
        public readonly string VersionProfileId;
        public readonly string VersionProfileName;
        public readonly string VersionProfileVisibilityScope;

        [OutputConstructor]
        private GetAppConnectorGroupResult(
            string cityCountry,

            ImmutableArray<Outputs.GetAppConnectorGroupConnectorResult> connectors,

            string countryCode,

            string creationTime,

            string description,

            string dnsQueryType,

            bool enabled,

            string geoLocationId,

            string? id,

            string latitude,

            string location,

            string longitude,

            bool lssAppConnectorGroup,

            string? microtenantId,

            string? microtenantName,

            string modifiedTime,

            string modifiedby,

            string? name,

            bool? overrideVersionProfile,

            ImmutableArray<Outputs.GetAppConnectorGroupServerGroupResult> serverGroups,

            bool tcpQuickAckApp,

            bool tcpQuickAckAssistant,

            bool tcpQuickAckReadAssistant,

            string upgradeDay,

            string upgradeTimeInSecs,

            bool useInDrMode,

            string versionProfileId,

            string versionProfileName,

            string versionProfileVisibilityScope)
        {
            CityCountry = cityCountry;
            Connectors = connectors;
            CountryCode = countryCode;
            CreationTime = creationTime;
            Description = description;
            DnsQueryType = dnsQueryType;
            Enabled = enabled;
            GeoLocationId = geoLocationId;
            Id = id;
            Latitude = latitude;
            Location = location;
            Longitude = longitude;
            LssAppConnectorGroup = lssAppConnectorGroup;
            MicrotenantId = microtenantId;
            MicrotenantName = microtenantName;
            ModifiedTime = modifiedTime;
            Modifiedby = modifiedby;
            Name = name;
            OverrideVersionProfile = overrideVersionProfile;
            ServerGroups = serverGroups;
            TcpQuickAckApp = tcpQuickAckApp;
            TcpQuickAckAssistant = tcpQuickAckAssistant;
            TcpQuickAckReadAssistant = tcpQuickAckReadAssistant;
            UpgradeDay = upgradeDay;
            UpgradeTimeInSecs = upgradeTimeInSecs;
            UseInDrMode = useInDrMode;
            VersionProfileId = versionProfileId;
            VersionProfileName = versionProfileName;
            VersionProfileVisibilityScope = versionProfileVisibilityScope;
        }
    }
}
