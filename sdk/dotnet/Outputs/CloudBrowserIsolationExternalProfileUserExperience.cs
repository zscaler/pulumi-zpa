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
    public sealed class CloudBrowserIsolationExternalProfileUserExperience
    {
        public readonly bool? BrowserInBrowser;
        public readonly bool? SessionPersistence;

        [OutputConstructor]
        private CloudBrowserIsolationExternalProfileUserExperience(
            bool? browserInBrowser,

            bool? sessionPersistence)
        {
            BrowserInBrowser = browserInBrowser;
            SessionPersistence = sessionPersistence;
        }
    }
}
