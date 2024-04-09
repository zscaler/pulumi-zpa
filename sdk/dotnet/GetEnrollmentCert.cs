// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Zscaler.Zpa
{
    public static class GetEnrollmentCert
    {
        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/about-enrollment-ca-certificates)
        /// * [API documentation](https://help.zscaler.com/zpa/obtaining-enrollment-certificate-details-using-api)
        /// 
        /// Use the **zpa_enrollment_cert** data source to get information about all configured enrollment certificate details created in the Zscaler Private Access cloud. This data source is required when creating provisioning key resources.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var root = Zpa.GetEnrollmentCert.Invoke(new()
        ///     {
        ///         Name = "Root",
        ///     });
        /// 
        ///     var client = Zpa.GetEnrollmentCert.Invoke(new()
        ///     {
        ///         Name = "Client",
        ///     });
        /// 
        ///     var connector = Zpa.GetEnrollmentCert.Invoke(new()
        ///     {
        ///         Name = "Connector",
        ///     });
        /// 
        ///     var serviceEdge = Zpa.GetEnrollmentCert.Invoke(new()
        ///     {
        ///         Name = "Service Edge",
        ///     });
        /// 
        ///     var isolationClient = Zpa.GetEnrollmentCert.Invoke(new()
        ///     {
        ///         Name = "Isolation Client",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetEnrollmentCertResult> InvokeAsync(GetEnrollmentCertArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEnrollmentCertResult>("zpa:index/getEnrollmentCert:getEnrollmentCert", args ?? new GetEnrollmentCertArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/about-enrollment-ca-certificates)
        /// * [API documentation](https://help.zscaler.com/zpa/obtaining-enrollment-certificate-details-using-api)
        /// 
        /// Use the **zpa_enrollment_cert** data source to get information about all configured enrollment certificate details created in the Zscaler Private Access cloud. This data source is required when creating provisioning key resources.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var root = Zpa.GetEnrollmentCert.Invoke(new()
        ///     {
        ///         Name = "Root",
        ///     });
        /// 
        ///     var client = Zpa.GetEnrollmentCert.Invoke(new()
        ///     {
        ///         Name = "Client",
        ///     });
        /// 
        ///     var connector = Zpa.GetEnrollmentCert.Invoke(new()
        ///     {
        ///         Name = "Connector",
        ///     });
        /// 
        ///     var serviceEdge = Zpa.GetEnrollmentCert.Invoke(new()
        ///     {
        ///         Name = "Service Edge",
        ///     });
        /// 
        ///     var isolationClient = Zpa.GetEnrollmentCert.Invoke(new()
        ///     {
        ///         Name = "Isolation Client",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetEnrollmentCertResult> Invoke(GetEnrollmentCertInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEnrollmentCertResult>("zpa:index/getEnrollmentCert:getEnrollmentCert", args ?? new GetEnrollmentCertInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEnrollmentCertArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public string? Id { get; set; }

        [Input("microtenantId")]
        public string? MicrotenantId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public GetEnrollmentCertArgs()
        {
        }
        public static new GetEnrollmentCertArgs Empty => new GetEnrollmentCertArgs();
    }

    public sealed class GetEnrollmentCertInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetEnrollmentCertInvokeArgs()
        {
        }
        public static new GetEnrollmentCertInvokeArgs Empty => new GetEnrollmentCertInvokeArgs();
    }


    [OutputType]
    public sealed class GetEnrollmentCertResult
    {
        public readonly bool AllowSigning;
        public readonly string Certificate;
        public readonly string ClientCertType;
        public readonly string Cname;
        public readonly string CreationTime;
        public readonly string Csr;
        public readonly string Description;
        public readonly string? Id;
        public readonly string IssuedBy;
        public readonly string IssuedTo;
        public readonly string? MicrotenantId;
        public readonly string ModifiedBy;
        public readonly string ModifiedTime;
        public readonly string? Name;
        public readonly string ParentCertId;
        public readonly string ParentCertName;
        public readonly string PrivateKey;
        public readonly bool PrivateKeyPresent;
        public readonly string SerialNo;
        public readonly string ValidFromInEpochSec;
        public readonly string ValidToInEpochSec;
        public readonly string ZrsaEncryptedPrivateKey;
        public readonly string ZrsaEncryptedSessionKey;

        [OutputConstructor]
        private GetEnrollmentCertResult(
            bool allowSigning,

            string certificate,

            string clientCertType,

            string cname,

            string creationTime,

            string csr,

            string description,

            string? id,

            string issuedBy,

            string issuedTo,

            string? microtenantId,

            string modifiedBy,

            string modifiedTime,

            string? name,

            string parentCertId,

            string parentCertName,

            string privateKey,

            bool privateKeyPresent,

            string serialNo,

            string validFromInEpochSec,

            string validToInEpochSec,

            string zrsaEncryptedPrivateKey,

            string zrsaEncryptedSessionKey)
        {
            AllowSigning = allowSigning;
            Certificate = certificate;
            ClientCertType = clientCertType;
            Cname = cname;
            CreationTime = creationTime;
            Csr = csr;
            Description = description;
            Id = id;
            IssuedBy = issuedBy;
            IssuedTo = issuedTo;
            MicrotenantId = microtenantId;
            ModifiedBy = modifiedBy;
            ModifiedTime = modifiedTime;
            Name = name;
            ParentCertId = parentCertId;
            ParentCertName = parentCertName;
            PrivateKey = privateKey;
            PrivateKeyPresent = privateKeyPresent;
            SerialNo = serialNo;
            ValidFromInEpochSec = validFromInEpochSec;
            ValidToInEpochSec = validToInEpochSec;
            ZrsaEncryptedPrivateKey = zrsaEncryptedPrivateKey;
            ZrsaEncryptedSessionKey = zrsaEncryptedSessionKey;
        }
    }
}
