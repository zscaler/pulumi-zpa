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

    public sealed class ApplicationSegmentPRACommonAppsDtoAppsConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("appTypes")]
        private InputList<string>? _appTypes;
        public InputList<string> AppTypes
        {
            get => _appTypes ?? (_appTypes = new InputList<string>());
            set => _appTypes = value;
        }

        [Input("applicationPort")]
        public Input<string>? ApplicationPort { get; set; }

        [Input("applicationProtocol")]
        public Input<string>? ApplicationProtocol { get; set; }

        [Input("connectionSecurity")]
        public Input<string>? ConnectionSecurity { get; set; }

        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public ApplicationSegmentPRACommonAppsDtoAppsConfigGetArgs()
        {
        }
        public static new ApplicationSegmentPRACommonAppsDtoAppsConfigGetArgs Empty => new ApplicationSegmentPRACommonAppsDtoAppsConfigGetArgs();
    }
}
