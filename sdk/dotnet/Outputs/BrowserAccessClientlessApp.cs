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
    public sealed class BrowserAccessClientlessApp
    {
        /// <summary>
        /// If you want ZPA to forward unauthenticated HTTP preflight OPTIONS requests from the browser to the app.
        /// </summary>
        public readonly bool? AllowOptions;
        /// <summary>
        /// Port for the BA app.
        /// </summary>
        public readonly string ApplicationPort;
        /// <summary>
        /// Protocol for the BA app.
        /// </summary>
        public readonly string ApplicationProtocol;
        /// <summary>
        /// ID of the BA certificate.
        /// </summary>
        public readonly string? CertificateId;
        public readonly string? Description;
        /// <summary>
        /// Domain name or IP address of the BA app.
        /// </summary>
        public readonly string? Domain;
        public readonly bool? Enabled;
        public readonly string? Id;
        public readonly string Name;
        /// <summary>
        /// Indicates whether Use Untrusted Certificates is enabled or disabled for a BA app.
        /// </summary>
        public readonly bool? TrustUntrustedCert;

        [OutputConstructor]
        private BrowserAccessClientlessApp(
            bool? allowOptions,

            string applicationPort,

            string applicationProtocol,

            string? certificateId,

            string? description,

            string? domain,

            bool? enabled,

            string? id,

            string name,

            bool? trustUntrustedCert)
        {
            AllowOptions = allowOptions;
            ApplicationPort = applicationPort;
            ApplicationProtocol = applicationProtocol;
            CertificateId = certificateId;
            Description = description;
            Domain = domain;
            Enabled = enabled;
            Id = id;
            Name = name;
            TrustUntrustedCert = trustUntrustedCert;
        }
    }
}
