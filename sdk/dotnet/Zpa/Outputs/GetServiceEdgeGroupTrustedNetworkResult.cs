// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.Outputs
{

    [OutputType]
    public sealed class GetServiceEdgeGroupTrustedNetworkResult
    {
        public readonly string CreationTime;
        public readonly string Domain;
        public readonly string Id;
        public readonly string MasterCustomerId;
        public readonly string ModifiedBy;
        public readonly string ModifiedTime;
        public readonly string Name;
        public readonly string NetworkId;
        public readonly string ZscalerCloud;

        [OutputConstructor]
        private GetServiceEdgeGroupTrustedNetworkResult(
            string creationTime,

            string domain,

            string id,

            string masterCustomerId,

            string modifiedBy,

            string modifiedTime,

            string name,

            string networkId,

            string zscalerCloud)
        {
            CreationTime = creationTime;
            Domain = domain;
            Id = id;
            MasterCustomerId = masterCustomerId;
            ModifiedBy = modifiedBy;
            ModifiedTime = modifiedTime;
            Name = name;
            NetworkId = networkId;
            ZscalerCloud = zscalerCloud;
        }
    }
}
