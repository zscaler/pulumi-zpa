// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.SegmentGroup.Outputs
{

    [OutputType]
    public sealed class GetSegmentGroupApplicationServerGroupResult
    {
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ConfigSpace;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (bool)
        /// </summary>
        public readonly bool DynamicDiscovery;
        /// <summary>
        /// (bool)
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The ID of the segment group to be exported.
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
        /// <summary>
        /// The name of the segment group to be exported.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetSegmentGroupApplicationServerGroupResult(
            string configSpace,

            string creationTime,

            string description,

            bool dynamicDiscovery,

            bool enabled,

            string id,

            string modifiedBy,

            string modifiedTime,

            string name)
        {
            ConfigSpace = configSpace;
            CreationTime = creationTime;
            Description = description;
            DynamicDiscovery = dynamicDiscovery;
            Enabled = enabled;
            Id = id;
            ModifiedBy = modifiedBy;
            ModifiedTime = modifiedTime;
            Name = name;
        }
    }
}
