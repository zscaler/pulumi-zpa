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
        public readonly Outputs.CloudBrowserIsolationExternalProfileUserExperienceForwardToZia? ForwardToZia;
        public readonly bool? PersistIsolationBar;
        public readonly bool? SessionPersistence;
        public readonly bool? Translate;
        public readonly bool? Zgpu;

        [OutputConstructor]
        private CloudBrowserIsolationExternalProfileUserExperience(
            bool? browserInBrowser,

            Outputs.CloudBrowserIsolationExternalProfileUserExperienceForwardToZia? forwardToZia,

            bool? persistIsolationBar,

            bool? sessionPersistence,

            bool? translate,

            bool? zgpu)
        {
            BrowserInBrowser = browserInBrowser;
            ForwardToZia = forwardToZia;
            PersistIsolationBar = persistIsolationBar;
            SessionPersistence = sessionPersistence;
            Translate = translate;
            Zgpu = zgpu;
        }
    }
}
