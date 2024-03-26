// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Zscaler.Zpa.Outputs
{

    [OutputType]
    public sealed class GetServerGroupAppConnectorGroupResult
    {
        public readonly string CityCountry;
        public readonly ImmutableArray<Outputs.GetServerGroupAppConnectorGroupConnectorResult> Connectors;
        public readonly string CountryCode;
        public readonly string CreationTime;
        /// <summary>
        /// (string) This field is the description of the server group.
        /// </summary>
        public readonly string Description;
        public readonly string DnsQueryType;
        /// <summary>
        /// (bool) This field defines if the server group is enabled or disabled.
        /// </summary>
        public readonly bool Enabled;
        public readonly string GeolocationId;
        /// <summary>
        /// The ID of the server group to be exported.
        /// </summary>
        public readonly string Id;
        public readonly string Latitude;
        public readonly string Location;
        public readonly string Longitude;
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        /// <summary>
        /// The name of the server group to be exported.
        /// </summary>
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetServerGroupAppConnectorGroupServerGroupResult> ServerGroups;
        public readonly bool SiemAppConnectorGroup;
        public readonly string UpgradeDay;
        public readonly string UpgradeTimeInSecs;
        public readonly string VersionProfileId;

        [OutputConstructor]
        private GetServerGroupAppConnectorGroupResult(
            string cityCountry,

            ImmutableArray<Outputs.GetServerGroupAppConnectorGroupConnectorResult> connectors,

            string countryCode,

            string creationTime,

            string description,

            string dnsQueryType,

            bool enabled,

            string geolocationId,

            string id,

            string latitude,

            string location,

            string longitude,

            string modifiedTime,

            string modifiedby,

            string name,

            ImmutableArray<Outputs.GetServerGroupAppConnectorGroupServerGroupResult> serverGroups,

            bool siemAppConnectorGroup,

            string upgradeDay,

            string upgradeTimeInSecs,

            string versionProfileId)
        {
            CityCountry = cityCountry;
            Connectors = connectors;
            CountryCode = countryCode;
            CreationTime = creationTime;
            Description = description;
            DnsQueryType = dnsQueryType;
            Enabled = enabled;
            GeolocationId = geolocationId;
            Id = id;
            Latitude = latitude;
            Location = location;
            Longitude = longitude;
            ModifiedTime = modifiedTime;
            Modifiedby = modifiedby;
            Name = name;
            ServerGroups = serverGroups;
            SiemAppConnectorGroup = siemAppConnectorGroup;
            UpgradeDay = upgradeDay;
            UpgradeTimeInSecs = upgradeTimeInSecs;
            VersionProfileId = versionProfileId;
        }
    }
}
