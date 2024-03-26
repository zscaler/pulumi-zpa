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
    /// <summary>
    /// The **zpa_cloud_browser_isolation_certificate** resource creates a Cloud Browser Isolation certificate. This resource can then be used when creating a CBI External Profile `zpa.CloudBrowserIsolationExternalProfile`.`
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using System.Linq;
    /// using Pulumi;
    /// using Zpa = Zscaler.Zpa;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Retrieve CBI Banner ID
    ///     var thisCloudBrowserIsolationCertificate = new Zpa.CloudBrowserIsolationCertificate("thisCloudBrowserIsolationCertificate", new()
    ///     {
    ///         Pem = File.ReadAllText("cert.pem"),
    ///     });
    /// 
    ///     var thisIndex_cloudBrowserIsolationCertificateCloudBrowserIsolationCertificate = new Zpa.CloudBrowserIsolationCertificate("thisIndex/cloudBrowserIsolationCertificateCloudBrowserIsolationCertificate", new()
    ///     {
    ///         Pem = @"    -----BEGIN CERTIFICATE-----
    ///     MIIFYDCCBEigAwIBAgIQQAF3ITfU6UK47naqPGQKtzANBgkqhkiG9w0BAQsFADA/
    /// ",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [ZpaResourceType("zpa:index/cloudBrowserIsolationCertificate:CloudBrowserIsolationCertificate")]
    public partial class CloudBrowserIsolationCertificate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the CBI certificate.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The certificate in PEM format.
        /// </summary>
        [Output("pem")]
        public Output<string?> Pem { get; private set; } = null!;


        /// <summary>
        /// Create a CloudBrowserIsolationCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloudBrowserIsolationCertificate(string name, CloudBrowserIsolationCertificateArgs? args = null, CustomResourceOptions? options = null)
            : base("zpa:index/cloudBrowserIsolationCertificate:CloudBrowserIsolationCertificate", name, args ?? new CloudBrowserIsolationCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloudBrowserIsolationCertificate(string name, Input<string> id, CloudBrowserIsolationCertificateState? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/cloudBrowserIsolationCertificate:CloudBrowserIsolationCertificate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/zscaler",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CloudBrowserIsolationCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloudBrowserIsolationCertificate Get(string name, Input<string> id, CloudBrowserIsolationCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new CloudBrowserIsolationCertificate(name, id, state, options);
        }
    }

    public sealed class CloudBrowserIsolationCertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the CBI certificate.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The certificate in PEM format.
        /// </summary>
        [Input("pem")]
        public Input<string>? Pem { get; set; }

        public CloudBrowserIsolationCertificateArgs()
        {
        }
        public static new CloudBrowserIsolationCertificateArgs Empty => new CloudBrowserIsolationCertificateArgs();
    }

    public sealed class CloudBrowserIsolationCertificateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the CBI certificate.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The certificate in PEM format.
        /// </summary>
        [Input("pem")]
        public Input<string>? Pem { get; set; }

        public CloudBrowserIsolationCertificateState()
        {
        }
        public static new CloudBrowserIsolationCertificateState Empty => new CloudBrowserIsolationCertificateState();
    }
}
