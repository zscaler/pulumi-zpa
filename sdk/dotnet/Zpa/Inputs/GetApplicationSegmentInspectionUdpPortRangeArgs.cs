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

    public sealed class GetApplicationSegmentInspectionUdpPortRangeInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("from")]
        public Input<string>? From { get; set; }

        [Input("to")]
        public Input<string>? To { get; set; }

        public GetApplicationSegmentInspectionUdpPortRangeInputArgs()
        {
        }
        public static new GetApplicationSegmentInspectionUdpPortRangeInputArgs Empty => new GetApplicationSegmentInspectionUdpPortRangeInputArgs();
    }
}
