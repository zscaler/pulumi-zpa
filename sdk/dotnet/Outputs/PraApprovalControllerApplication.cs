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
    public sealed class PraApprovalControllerApplication
    {
        /// <summary>
        /// The unique identifier of the pra application segment
        /// </summary>
        public readonly ImmutableArray<string> Ids;

        [OutputConstructor]
        private PraApprovalControllerApplication(ImmutableArray<string> ids)
        {
            Ids = ids;
        }
    }
}
