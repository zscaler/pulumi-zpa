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
    public sealed class GetCloudBrowserIsolationExternalProfileRegionResult
    {
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetCloudBrowserIsolationExternalProfileRegionResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}
