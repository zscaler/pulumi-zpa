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
    public sealed class GetServiceEdgeGroupServiceEdgeZpnSubModuleUpgradeListResult
    {
        public readonly string CreationTime;
        public readonly string CurrentVersion;
        public readonly string EntityGid;
        public readonly string ExpectedVersion;
        public readonly string Id;
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        public readonly string Role;
        public readonly string UpgradeStatus;
        public readonly string UpgradeTime;

        [OutputConstructor]
        private GetServiceEdgeGroupServiceEdgeZpnSubModuleUpgradeListResult(
            string creationTime,

            string currentVersion,

            string entityGid,

            string expectedVersion,

            string id,

            string modifiedTime,

            string modifiedby,

            string role,

            string upgradeStatus,

            string upgradeTime)
        {
            CreationTime = creationTime;
            CurrentVersion = currentVersion;
            EntityGid = entityGid;
            ExpectedVersion = expectedVersion;
            Id = id;
            ModifiedTime = modifiedTime;
            Modifiedby = modifiedby;
            Role = role;
            UpgradeStatus = upgradeStatus;
            UpgradeTime = upgradeTime;
        }
    }
}
