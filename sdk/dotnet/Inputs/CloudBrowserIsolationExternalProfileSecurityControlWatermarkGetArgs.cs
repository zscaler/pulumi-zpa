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

    public sealed class CloudBrowserIsolationExternalProfileSecurityControlWatermarkGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("message")]
        public Input<string>? Message { get; set; }

        [Input("showMessage")]
        public Input<bool>? ShowMessage { get; set; }

        [Input("showTimestamp")]
        public Input<bool>? ShowTimestamp { get; set; }

        [Input("showUserId")]
        public Input<bool>? ShowUserId { get; set; }

        public CloudBrowserIsolationExternalProfileSecurityControlWatermarkGetArgs()
        {
        }
        public static new CloudBrowserIsolationExternalProfileSecurityControlWatermarkGetArgs Empty => new CloudBrowserIsolationExternalProfileSecurityControlWatermarkGetArgs();
    }
}
