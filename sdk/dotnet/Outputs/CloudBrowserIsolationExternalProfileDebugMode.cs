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
    public sealed class CloudBrowserIsolationExternalProfileDebugMode
    {
        public readonly bool? Allowed;
        public readonly string? FilePassword;

        [OutputConstructor]
        private CloudBrowserIsolationExternalProfileDebugMode(
            bool? allowed,

            string? filePassword)
        {
            Allowed = allowed;
            FilePassword = filePassword;
        }
    }
}
