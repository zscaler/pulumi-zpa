// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.Inputs
{

    public sealed class GetApplicationSegmentInspectionUdpPortRangeArgs : global::Pulumi.InvokeArgs
    {
        [Input("from")]
        public string? From { get; set; }

        [Input("to")]
        public string? To { get; set; }

        public GetApplicationSegmentInspectionUdpPortRangeArgs()
        {
        }
        public static new GetApplicationSegmentInspectionUdpPortRangeArgs Empty => new GetApplicationSegmentInspectionUdpPortRangeArgs();
    }
}
