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
        /// &lt;!--Start PulumiCodeChooser --&gt;
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
        ///     var foo = Zpa.GetAppConnectorGroup.Invoke(new()
        ///     {
        ///         Id = "123456789",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetAppConnectorGroupResult> InvokeAsync(GetAppConnectorGroupArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAppConnectorGroupResult>("zpa:index/getAppConnectorGroup:getAppConnectorGroup", args ?? new GetAppConnectorGroupArgs(), options.WithDefaults());

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
        ///     var foo = Zpa.GetAppConnectorGroup.Invoke(new()
        ///     {
        ///         Name = "DataCenter",
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
        ///     var foo = Zpa.GetAppConnectorGroup.Invoke(new()
        ///     {
        ///         Id = "123456789",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetAppConnectorGroupResult> Invoke(GetAppConnectorGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAppConnectorGroupResult>("zpa:index/getAppConnectorGroup:getAppConnectorGroup", args ?? new GetAppConnectorGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppConnectorGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the App Connector Group.
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
        /// Name of the App Connector Group.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// (bool) Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
        /// </summary>
        [Input("overrideVersionProfile")]
        public bool? OverrideVersionProfile { get; set; }

        public GetAppConnectorGroupArgs()
        {
        }
        public static new GetAppConnectorGroupArgs Empty => new GetAppConnectorGroupArgs();
    }

    public sealed class GetAppConnectorGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the App Connector Group.
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
        /// Name of the App Connector Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (bool) Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
        /// </summary>
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
        /// <summary>
        /// (String) Whether Double Encryption is enabled or disabled for the app.
        /// </summary>
        public readonly string CityCountry;
        public readonly ImmutableArray<Outputs.GetAppConnectorGroupConnectorResult> Connectors;
        /// <summary>
        /// (String)
        /// </summary>
        public readonly string CountryCode;
        public readonly string CreationTime;
        /// <summary>
        /// (String) Description of the App Connector Group.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (String)
        /// </summary>
        public readonly string DnsQueryType;
        /// <summary>
        /// (String) Whether this App Connector Group is enabled or not. Default value: `true`. Supported values: `true`, `false`
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// (String)
        /// </summary>
        public readonly string GeoLocationId;
        public readonly string? Id;
        /// <summary>
        /// (String) Latitude of the App Connector Group. Integer or decimal. With values in the range of `-90` to `90`
        /// </summary>
        public readonly string Latitude;
        /// <summary>
        /// (String) Location of the App Connector Group.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// (String) Longitude of the App Connector Group. Integer or decimal. With values in the range of `-180` to `180`
        /// </summary>
        public readonly string Longitude;
        public readonly bool LssAppConnectorGroup;
        /// <summary>
        /// (string) The ID of the microtenant the resource is to be associated with.
        /// </summary>
        public readonly string? MicrotenantId;
        /// <summary>
        /// (string) The name of the microtenant the resource is to be associated with.
        /// </summary>
        public readonly string? MicrotenantName;
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        public readonly string? Name;
        /// <summary>
        /// (bool) Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
        /// </summary>
        public readonly bool? OverrideVersionProfile;
        public readonly ImmutableArray<Outputs.GetAppConnectorGroupServerGroupResult> ServerGroups;
        public readonly bool TcpQuickAckApp;
        public readonly bool TcpQuickAckAssistant;
        public readonly bool TcpQuickAckReadAssistant;
        /// <summary>
        /// (String) App Connectors in this group will attempt to update to a newer version of the software during this specified day
        /// </summary>
        public readonly string UpgradeDay;
        /// <summary>
        /// (String) App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600`. Integer in seconds (i.e., `-66600`). The integer should be greater than or equal to `0` and less than `86400`, in `15` minute intervals
        /// </summary>
        public readonly string UpgradeTimeInSecs;
        /// <summary>
        /// (Optional) Supported values: `true`, `false`
        /// </summary>
        public readonly bool UseInDrMode;
        /// <summary>
        /// (String) ID of the version profile.
        /// Exported values are:
        /// </summary>
        public readonly string VersionProfileId;
        /// <summary>
        /// (String)
        /// Exported values are:
        /// </summary>
        public readonly string VersionProfileName;
        /// <summary>
        /// (String)
        /// Exported values are:
        /// </summary>
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
