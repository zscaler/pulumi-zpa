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

    public sealed class ApplicationSegmentBrowserAccessClientlessAppGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If you want ZPA to forward unauthenticated HTTP preflight OPTIONS requests from the browser to the app.
        /// </summary>
        [Input("allowOptions")]
        public Input<bool>? AllowOptions { get; set; }

        /// <summary>
        /// Port for the BA app.
        /// </summary>
        [Input("applicationPort", required: true)]
        public Input<string> ApplicationPort { get; set; } = null!;

        /// <summary>
        /// Protocol for the BA app.
        /// </summary>
        [Input("applicationProtocol", required: true)]
        public Input<string> ApplicationProtocol { get; set; } = null!;

        /// <summary>
        /// ID of the BA certificate.
        /// </summary>
        [Input("certificateId", required: true)]
        public Input<string> CertificateId { get; set; } = null!;

        [Input("cname")]
        public Input<string>? Cname { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Domain name or IP address of the BA app.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("hidden")]
        public Input<bool>? Hidden { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("localDomain")]
        public Input<string>? LocalDomain { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Indicates whether Use Untrusted Certificates is enabled or disabled for a BA app.
        /// </summary>
        [Input("trustUntrustedCert")]
        public Input<bool>? TrustUntrustedCert { get; set; }

        public ApplicationSegmentBrowserAccessClientlessAppGetArgs()
        {
        }
        public static new ApplicationSegmentBrowserAccessClientlessAppGetArgs Empty => new ApplicationSegmentBrowserAccessClientlessAppGetArgs();
    }
}
