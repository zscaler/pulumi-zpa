// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Zscaler.Zpa.Inputs
{

    public sealed class BrowserAccessTcpPortRangeArgs : global::Pulumi.ResourceArgs
    {
        [Input("from")]
        public Input<string>? From { get; set; }

        [Input("to")]
        public Input<string>? To { get; set; }

        public BrowserAccessTcpPortRangeArgs()
        {
        }
        public static new BrowserAccessTcpPortRangeArgs Empty => new BrowserAccessTcpPortRangeArgs();
    }
}
