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
    public sealed class GetServiceEdgeControllerZpnSubModuleUpgradeListResult
    {
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string CurrentVersion;
        public readonly string EntityGid;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ExpectedVersion;
        /// <summary>
        /// The ID of the service edge controllerto be exported.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ModifiedBy;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ModifiedTime;
        public readonly string Role;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string UpgradeStatus;
        public readonly string UpgradeTime;

        [OutputConstructor]
        private GetServiceEdgeControllerZpnSubModuleUpgradeListResult(
            string creationTime,

            string currentVersion,

            string entityGid,

            string expectedVersion,

            string id,

            string modifiedBy,

            string modifiedTime,

            string role,

            string upgradeStatus,

            string upgradeTime)
        {
            CreationTime = creationTime;
            CurrentVersion = currentVersion;
            EntityGid = entityGid;
            ExpectedVersion = expectedVersion;
            Id = id;
            ModifiedBy = modifiedBy;
            ModifiedTime = modifiedTime;
            Role = role;
            UpgradeStatus = upgradeStatus;
            UpgradeTime = upgradeTime;
        }
    }
}