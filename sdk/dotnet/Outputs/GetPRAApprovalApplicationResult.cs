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
    public sealed class GetPRAApprovalApplicationResult
    {
        /// <summary>
        /// The unique identifier of the pra application segment
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the pra application segment
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetPRAApprovalApplicationResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}
