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
    public sealed class GetApplicationSegmentBrowserAccessServerGroupResult
    {
        /// <summary>
        /// - (String) This field defines the id of the application server.
        /// </summary>
        public readonly ImmutableArray<string> Ids;

        [OutputConstructor]
        private GetApplicationSegmentBrowserAccessServerGroupResult(ImmutableArray<string> ids)
        {
            Ids = ids;
        }
    }
}
