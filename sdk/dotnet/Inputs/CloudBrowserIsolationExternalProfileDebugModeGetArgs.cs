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

    public sealed class CloudBrowserIsolationExternalProfileDebugModeGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowed")]
        public Input<bool>? Allowed { get; set; }

        [Input("filePassword")]
        public Input<string>? FilePassword { get; set; }

        public CloudBrowserIsolationExternalProfileDebugModeGetArgs()
        {
        }
        public static new CloudBrowserIsolationExternalProfileDebugModeGetArgs Empty => new CloudBrowserIsolationExternalProfileDebugModeGetArgs();
    }
}
