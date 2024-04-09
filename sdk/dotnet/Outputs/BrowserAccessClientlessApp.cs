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
        public readonly string CertificateId;
        public readonly string? Cname;
        public readonly string? Description;
        /// <summary>
        /// Domain name or IP address of the BA app.
        /// </summary>
        public readonly string? Domain;
        public readonly bool? Enabled;
        public readonly bool? Hidden;
        public readonly string? Id;
        public readonly string? LocalDomain;
        public readonly string Name;
        public readonly string? Path;
        /// <summary>
        /// Indicates whether Use Untrusted Certificates is enabled or disabled for a BA app.
        /// </summary>
        public readonly bool? TrustUntrustedCert;

        [OutputConstructor]
        private BrowserAccessClientlessApp(
            bool? allowOptions,

            string applicationPort,

            string applicationProtocol,

            string certificateId,

            string? cname,

            string? description,

            string? domain,

            bool? enabled,

            bool? hidden,

            string? id,

            string? localDomain,

            string name,

            string? path,

            bool? trustUntrustedCert)
        {
            AllowOptions = allowOptions;
            ApplicationPort = applicationPort;
            ApplicationProtocol = applicationProtocol;
            CertificateId = certificateId;
            Cname = cname;
            Description = description;
            Domain = domain;
            Enabled = enabled;
            Hidden = hidden;
            Id = id;
            LocalDomain = localDomain;
            Name = name;
            Path = path;
            TrustUntrustedCert = trustUntrustedCert;
        }
    }
}
