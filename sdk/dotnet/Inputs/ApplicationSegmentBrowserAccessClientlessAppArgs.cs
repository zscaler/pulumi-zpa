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

    public sealed class ApplicationSegmentBrowserAccessClientlessAppArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If you want ZPA to forward unauthenticated HTTP preflight OPTIONS requests from the browser to the app.. Supported values: `true` and `false`
        /// </summary>
        [Input("allowOptions")]
        public Input<bool>? AllowOptions { get; set; }

        /// <summary>
        /// Port for the BA app.
        /// </summary>
        [Input("applicationPort", required: true)]
        public Input<string> ApplicationPort { get; set; } = null!;

        /// <summary>
        /// Protocol for the BA app. Supported values: `HTTP` and `HTTPS`
        /// </summary>
        [Input("applicationProtocol", required: true)]
        public Input<string> ApplicationProtocol { get; set; } = null!;

        /// <summary>
        /// ID of the BA certificate. Refer to the data source documentation for `zpa.BrowserCertificate`
        /// </summary>
        [Input("certificateId", required: true)]
        public Input<string> CertificateId { get; set; } = null!;

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("cname")]
        public Input<string>? Cname { get; set; }

        /// <summary>
        /// (Optional) Description of the application.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Domain name or IP address of the BA app.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// (Optional) - Whether this app is enabled or not.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("hidden")]
        public Input<bool>? Hidden { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("localDomain")]
        public Input<string>? LocalDomain { get; set; }

        /// <summary>
        /// Name of BA app.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("trustUntrustedCert")]
        public Input<bool>? TrustUntrustedCert { get; set; }

        public ApplicationSegmentBrowserAccessClientlessAppArgs()
        {
        }
        public static new ApplicationSegmentBrowserAccessClientlessAppArgs Empty => new ApplicationSegmentBrowserAccessClientlessAppArgs();
    }
}
