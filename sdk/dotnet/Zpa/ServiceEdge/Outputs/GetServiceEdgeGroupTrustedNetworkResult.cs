// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.ServiceEdge.Outputs
{

    [OutputType]
    public sealed class GetServiceEdgeGroupTrustedNetworkResult
    {
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string Domain;
        /// <summary>
        /// The ID of the service edge group to be exported.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string MasterCustomerId;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        /// <summary>
        /// The name of the service edge group to be exported.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string NetworkId;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ZscalerCloud;

        [OutputConstructor]
        private GetServiceEdgeGroupTrustedNetworkResult(
            string creationTime,

            string domain,

            string id,

            string masterCustomerId,

            string modifiedTime,

            string modifiedby,

            string name,

            string networkId,

            string zscalerCloud)
        {
            CreationTime = creationTime;
            Domain = domain;
            Id = id;
            MasterCustomerId = masterCustomerId;
            ModifiedTime = modifiedTime;
            Modifiedby = modifiedby;
            Name = name;
            NetworkId = networkId;
            ZscalerCloud = zscalerCloud;
        }
    }
}
