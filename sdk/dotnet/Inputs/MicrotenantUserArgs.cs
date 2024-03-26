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

    public sealed class MicrotenantUserArgs : global::Pulumi.ResourceArgs
    {
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("username")]
        public Input<string>? Username { get; set; }

        public MicrotenantUserArgs()
        {
        }
        public static new MicrotenantUserArgs Empty => new MicrotenantUserArgs();
    }
}
