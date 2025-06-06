// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa
{
    /// <summary>
    /// * [Official documentation](https://help.zscaler.com/zpa/about-web-server-certificates)
    /// * [API documentation](https://help.zscaler.com/zpa/configuring-certificates-using-api)
    /// 
    /// Use the **zpa_ba_certificate** creates a browser access certificate with a private key in the Zscaler Private Access cloud. This resource is required when creating a browser access application segment resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Zpa = Pulumi.Zpa;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var foo = Zpa.GetBaCertificate.Invoke(new()
    ///     {
    ///         Name = "example.acme.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ######### PASSWORDS OR RELATED CREDENTIALS ATTRIBUTES IN THIS FILE #########\
    /// ######### ARE FOR EXAMPLE ONLY AND NOT USED IN PRODUCTION SYSTEMS ##########
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Zpa = zscaler.PulumiPackage.Zpa;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // ZPA Browser Access resource
    ///     var @this = new Zpa.BrowserCertificate("this", new()
    ///     {
    ///         CertBlob = @"-----BEGIN PRIVATE KEY-----
    /// MIIDyzCCArOgA
    /// -----END PRIVATE KEY-----
    /// -----BEGIN CERTIFICATE-----
    /// MIIDyzCCArOgAwIBAgIUekBD+iu64583B3u5ew7Bqj2O5cQwDQYJKoZIhvcNAQEL
    /// -----END CERTIFICATE-----
    /// 
    /// ",
    ///         Description = "server.example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Let's Encrypt Certbot
    /// 
    /// This example demonstrates generatoring a domain certificate with letsencrypt
    /// certbot https://letsencrypt.org/getting-started/
    /// 
    /// Use letsencrypt's certbot to generate domain certificates in RSA output mode.
    /// The generator's output corresponds to `zpa.BrowserCertificate` fields in the
    /// following manner.
    /// 
    /// Zscaler Field          | Certbot file
    /// --------------------|--------------
    /// `certblob`          | `cert.pem`
    /// `certblob`          | `privkey.pem`
    /// 
    /// ## Import
    /// 
    /// This resource does not support importing.
    /// </summary>
    [ZpaResourceType("zpa:index/browserCertificate:BrowserCertificate")]
    public partial class BrowserCertificate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the certificate
        /// </summary>
        [Output("certBlob")]
        public Output<string?> CertBlob { get; private set; } = null!;

        /// <summary>
        /// The certificate text in PEM format
        /// </summary>
        [Output("certificate")]
        public Output<string> Certificate { get; private set; } = null!;

        /// <summary>
        /// The description of the certificate
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the Microtenant
        /// </summary>
        [Output("microtenantId")]
        public Output<string> MicrotenantId { get; private set; } = null!;

        /// <summary>
        /// The name of the certificate.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a BrowserCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BrowserCertificate(string name, BrowserCertificateArgs? args = null, CustomResourceOptions? options = null)
            : base("zpa:index/browserCertificate:BrowserCertificate", name, args ?? new BrowserCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BrowserCertificate(string name, Input<string> id, BrowserCertificateState? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/browserCertificate:BrowserCertificate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BrowserCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BrowserCertificate Get(string name, Input<string> id, BrowserCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new BrowserCertificate(name, id, state, options);
        }
    }

    public sealed class BrowserCertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the certificate
        /// </summary>
        [Input("certBlob")]
        public Input<string>? CertBlob { get; set; }

        /// <summary>
        /// The description of the certificate
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The unique identifier of the Microtenant
        /// </summary>
        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// The name of the certificate.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public BrowserCertificateArgs()
        {
        }
        public static new BrowserCertificateArgs Empty => new BrowserCertificateArgs();
    }

    public sealed class BrowserCertificateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the certificate
        /// </summary>
        [Input("certBlob")]
        public Input<string>? CertBlob { get; set; }

        /// <summary>
        /// The certificate text in PEM format
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// The description of the certificate
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The unique identifier of the Microtenant
        /// </summary>
        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// The name of the certificate.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public BrowserCertificateState()
        {
        }
        public static new BrowserCertificateState Empty => new BrowserCertificateState();
    }
}
