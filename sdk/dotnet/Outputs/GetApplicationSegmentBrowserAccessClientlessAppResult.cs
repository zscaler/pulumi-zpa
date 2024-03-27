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
    public sealed class GetApplicationSegmentBrowserAccessClientlessAppResult
    {
        /// <summary>
        /// (bool)
        /// </summary>
        public readonly bool AllowOptions;
        public readonly string AppId;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ApplicationPort;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ApplicationProtocol;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string CertificateId;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string CertificateName;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string Cname;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string Domain;
        /// <summary>
        /// (bool)
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// (bool)
        /// </summary>
        public readonly bool Hidden;
        /// <summary>
        /// This field defines the id of the application server.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string LocalDomain;
        /// <summary>
        /// This field defines the name of the server.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// (bool)
        /// </summary>
        public readonly bool TrustUntrustedCert;

        [OutputConstructor]
        private GetApplicationSegmentBrowserAccessClientlessAppResult(
            bool allowOptions,

            string appId,

            string applicationPort,

            string applicationProtocol,

            string certificateId,

            string certificateName,

            string cname,

            string description,

            string domain,

            bool enabled,

            bool hidden,

            string id,

            string localDomain,

            string name,

            string path,

            bool trustUntrustedCert)
        {
            AllowOptions = allowOptions;
            AppId = appId;
            ApplicationPort = applicationPort;
            ApplicationProtocol = applicationProtocol;
            CertificateId = certificateId;
            CertificateName = certificateName;
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