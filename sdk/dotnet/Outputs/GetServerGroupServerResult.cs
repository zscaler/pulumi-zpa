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
    public sealed class GetServerGroupServerResult
    {
        public readonly string Address;
        public readonly ImmutableArray<string> AppServerGroupIds;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ConfigSpace;
        public readonly string CreationTime;
        /// <summary>
        /// (string) This field is the description of the server group.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (bool) This field defines if the server group is enabled or disabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The ID of the server group to be exported.
        /// </summary>
        public readonly string Id;
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        /// <summary>
        /// The name of the server group to be exported.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetServerGroupServerResult(
            string address,

            ImmutableArray<string> appServerGroupIds,

            string configSpace,

            string creationTime,

            string description,

            bool enabled,

            string id,

            string modifiedTime,

            string modifiedby,

            string name)
        {
            Address = address;
            AppServerGroupIds = appServerGroupIds;
            ConfigSpace = configSpace;
            CreationTime = creationTime;
            Description = description;
            Enabled = enabled;
            Id = id;
            ModifiedTime = modifiedTime;
            Modifiedby = modifiedby;
            Name = name;
        }
    }
}
