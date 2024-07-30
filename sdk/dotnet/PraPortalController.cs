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
    /// * [Official documentation](https://help.zscaler.com/zpa/about-privileged-portals)
    /// * [API documentation](https://help.zscaler.com/zpa/configuring-privileged-portals-using-api)
    /// 
    /// The **zpa_pra_portal_controller** resource creates a privileged remote access portal in the Zscaler Private Access cloud. This resource can then be referenced in an privileged remote access console resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Zpa = Pulumi.Zpa;
    /// using Zpa = Zscaler.Zpa;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var thisBaCertificate = Zpa.GetBaCertificate.Invoke(new()
    ///     {
    ///         Name = "portal.acme.com",
    ///     });
    /// 
    ///     var thisPRAPortal = new Zpa.PRAPortal("thisPRAPortal", new()
    ///     {
    ///         Description = "portal.acme.com",
    ///         Enabled = true,
    ///         Domain = "portal.acme.com",
    ///         CertificateId = thisBaCertificate.Apply(getBaCertificateResult =&gt; getBaCertificateResult.Id),
    ///         UserNotification = "Created with Terraform",
    ///         UserNotificationEnabled = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
    /// 
    /// Visit
    /// 
    /// **pra_portal_controller** can be imported by using `&lt;PORTAL ID&gt;` or `&lt;PORTAL NAME&gt;` as the import ID.
    /// 
    /// For example:
    /// 
    /// ```sh
    /// $ pulumi import zpa:index/praPortalController:PraPortalController this &lt;portal_id&gt;
    /// ```
    /// 
    /// or
    /// 
    /// ```sh
    /// $ pulumi import zpa:index/praPortalController:PraPortalController this &lt;portal_name&gt;
    /// ```
    /// </summary>
    [Obsolete(@"zpa.index/praportalcontroller.PraPortalController has been deprecated in favor of zpa.index/praportal.PRAPortal")]
    [ZpaResourceType("zpa:index/praPortalController:PraPortalController")]
    public partial class PraPortalController : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The unique identifier of the certificate
        /// </summary>
        [Output("certificateId")]
        public Output<string?> CertificateId { get; private set; } = null!;

        /// <summary>
        /// The description of the privileged portal
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The domain of the privileged portal
        /// </summary>
        [Output("domain")]
        public Output<string?> Domain { get; private set; } = null!;

        /// <summary>
        /// Whether or not the privileged portal is enabled
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
        /// microtenantId as 0 when making requests to retrieve data from the Default Microtenant. Pass microtenantId as null to
        /// retrieve data from all customers associated with the tenant.
        /// </summary>
        [Output("microtenantId")]
        public Output<string> MicrotenantId { get; private set; } = null!;

        /// <summary>
        /// The name of the privileged portal
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The notification message displayed in the banner of the privileged portallink, if enabled
        /// </summary>
        [Output("userNotification")]
        public Output<string?> UserNotification { get; private set; } = null!;

        /// <summary>
        /// Indicates if the Notification Banner is enabled (true) or disabled (false)
        /// </summary>
        [Output("userNotificationEnabled")]
        public Output<bool?> UserNotificationEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a PraPortalController resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PraPortalController(string name, PraPortalControllerArgs? args = null, CustomResourceOptions? options = null)
            : base("zpa:index/praPortalController:PraPortalController", name, args ?? new PraPortalControllerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PraPortalController(string name, Input<string> id, PraPortalControllerState? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/praPortalController:PraPortalController", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PraPortalController resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PraPortalController Get(string name, Input<string> id, PraPortalControllerState? state = null, CustomResourceOptions? options = null)
        {
            return new PraPortalController(name, id, state, options);
        }
    }

    public sealed class PraPortalControllerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique identifier of the certificate
        /// </summary>
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        /// <summary>
        /// The description of the privileged portal
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The domain of the privileged portal
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Whether or not the privileged portal is enabled
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
        /// microtenantId as 0 when making requests to retrieve data from the Default Microtenant. Pass microtenantId as null to
        /// retrieve data from all customers associated with the tenant.
        /// </summary>
        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// The name of the privileged portal
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The notification message displayed in the banner of the privileged portallink, if enabled
        /// </summary>
        [Input("userNotification")]
        public Input<string>? UserNotification { get; set; }

        /// <summary>
        /// Indicates if the Notification Banner is enabled (true) or disabled (false)
        /// </summary>
        [Input("userNotificationEnabled")]
        public Input<bool>? UserNotificationEnabled { get; set; }

        public PraPortalControllerArgs()
        {
        }
        public static new PraPortalControllerArgs Empty => new PraPortalControllerArgs();
    }

    public sealed class PraPortalControllerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique identifier of the certificate
        /// </summary>
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        /// <summary>
        /// The description of the privileged portal
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The domain of the privileged portal
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Whether or not the privileged portal is enabled
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
        /// microtenantId as 0 when making requests to retrieve data from the Default Microtenant. Pass microtenantId as null to
        /// retrieve data from all customers associated with the tenant.
        /// </summary>
        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// The name of the privileged portal
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The notification message displayed in the banner of the privileged portallink, if enabled
        /// </summary>
        [Input("userNotification")]
        public Input<string>? UserNotification { get; set; }

        /// <summary>
        /// Indicates if the Notification Banner is enabled (true) or disabled (false)
        /// </summary>
        [Input("userNotificationEnabled")]
        public Input<bool>? UserNotificationEnabled { get; set; }

        public PraPortalControllerState()
        {
        }
        public static new PraPortalControllerState Empty => new PraPortalControllerState();
    }
}
