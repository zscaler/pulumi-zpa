// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.ApplicationSegment.Outputs
{

    [OutputType]
    public sealed class GetApplicationSegmentInspectionInspectionAppResult
    {
        public readonly string AppId;
        /// <summary>
        /// (string) TCP/UDP Port for ZPA Inspection.
        /// </summary>
        public readonly string ApplicationPort;
        /// <summary>
        /// (string) Protocol for the Inspection Application. Supported values: `HTTP` and `HTTPS`
        /// </summary>
        public readonly string ApplicationProtocol;
        /// <summary>
        /// (string) - ID of the signing certificate. This field is required if the applicationProtocol is set to `HTTPS`. The certificateId is not supported if the applicationProtocol is set to `HTTP`.
        /// </summary>
        public readonly string CertificateId;
        /// <summary>
        /// (string) - Parameter required when `application_protocol` is of type `HTTPS`
        /// </summary>
        public readonly string CertificateName;
        /// <summary>
        /// (string) Description of the application.
        /// </summary>
        public readonly string Description;
        public readonly string Domain;
        /// <summary>
        /// (bool) Whether this application is enabled or not
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The ID of the Inspection Application Segment to be exported.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the Inspection Application Segment to be exported.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetApplicationSegmentInspectionInspectionAppResult(
            string appId,

            string applicationPort,

            string applicationProtocol,

            string certificateId,

            string certificateName,

            string description,

            string domain,

            bool enabled,

            string id,

            string name)
        {
            AppId = appId;
            ApplicationPort = applicationPort;
            ApplicationProtocol = applicationProtocol;
            CertificateId = certificateId;
            CertificateName = certificateName;
            Description = description;
            Domain = domain;
            Enabled = enabled;
            Id = id;
            Name = name;
        }
    }
}
