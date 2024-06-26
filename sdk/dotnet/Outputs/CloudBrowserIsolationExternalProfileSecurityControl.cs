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
    public sealed class CloudBrowserIsolationExternalProfileSecurityControl
    {
        public readonly bool? AllowPrinting;
        public readonly string? CopyPaste;
        public readonly bool? DocumentViewer;
        public readonly bool? LocalRender;
        public readonly bool? RestrictKeystrokes;
        public readonly string? UploadDownload;

        [OutputConstructor]
        private CloudBrowserIsolationExternalProfileSecurityControl(
            bool? allowPrinting,

            string? copyPaste,

            bool? documentViewer,

            bool? localRender,

            bool? restrictKeystrokes,

            string? uploadDownload)
        {
            AllowPrinting = allowPrinting;
            CopyPaste = copyPaste;
            DocumentViewer = documentViewer;
            LocalRender = localRender;
            RestrictKeystrokes = restrictKeystrokes;
            UploadDownload = uploadDownload;
        }
    }
}
