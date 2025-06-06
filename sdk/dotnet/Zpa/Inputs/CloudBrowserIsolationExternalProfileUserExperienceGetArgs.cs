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

    public sealed class CloudBrowserIsolationExternalProfileUserExperienceGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("browserInBrowser")]
        public Input<bool>? BrowserInBrowser { get; set; }

        [Input("forwardToZia")]
        public Input<Inputs.CloudBrowserIsolationExternalProfileUserExperienceForwardToZiaGetArgs>? ForwardToZia { get; set; }

        [Input("persistIsolationBar")]
        public Input<bool>? PersistIsolationBar { get; set; }

        [Input("sessionPersistence")]
        public Input<bool>? SessionPersistence { get; set; }

        [Input("translate")]
        public Input<bool>? Translate { get; set; }

        [Input("zgpu")]
        public Input<bool>? Zgpu { get; set; }

        public CloudBrowserIsolationExternalProfileUserExperienceGetArgs()
        {
        }
        public static new CloudBrowserIsolationExternalProfileUserExperienceGetArgs Empty => new CloudBrowserIsolationExternalProfileUserExperienceGetArgs();
    }
}
