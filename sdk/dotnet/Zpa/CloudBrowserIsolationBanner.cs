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
    /// * [Official documentation](https://help.zscaler.com/isolation/adding-banner-theme-isolation-end-user-notification-zpa)
    /// 
    /// The **zpa_cloud_browser_isolation_banner** resource creates a Cloud Browser Isolation banner. This resource is required as part of the attribute `banner_id` when creating an Cloud Browser Isolation External Profile ``zpa.CloudBrowserIsolationExternalProfile``
    /// 
    /// ## Import
    /// 
    /// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
    /// 
    /// Visit
    /// 
    /// Application Segment can be imported by using `&lt;BANNER ID&gt;` or `&lt;BANNER NAME&gt;` as the import ID.
    /// 
    /// ```sh
    /// $ pulumi import zpa:index/cloudBrowserIsolationBanner:CloudBrowserIsolationBanner example &lt;banner_id&gt;
    /// ```
    /// 
    /// or
    /// 
    /// ```sh
    /// $ pulumi import zpa:index/cloudBrowserIsolationBanner:CloudBrowserIsolationBanner example &lt;banner_name&gt;
    /// ```
    /// </summary>
    [ZpaResourceType("zpa:index/cloudBrowserIsolationBanner:CloudBrowserIsolationBanner")]
    public partial class CloudBrowserIsolationBanner : global::Pulumi.CustomResource
    {
        [Output("banner")]
        public Output<bool> Banner { get; private set; } = null!;

        [Output("logo")]
        public Output<string> Logo { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("notificationText")]
        public Output<string?> NotificationText { get; private set; } = null!;

        [Output("notificationTitle")]
        public Output<string?> NotificationTitle { get; private set; } = null!;

        [Output("persist")]
        public Output<bool> Persist { get; private set; } = null!;

        [Output("primaryColor")]
        public Output<string?> PrimaryColor { get; private set; } = null!;

        [Output("textColor")]
        public Output<string?> TextColor { get; private set; } = null!;


        /// <summary>
        /// Create a CloudBrowserIsolationBanner resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloudBrowserIsolationBanner(string name, CloudBrowserIsolationBannerArgs? args = null, CustomResourceOptions? options = null)
            : base("zpa:index/cloudBrowserIsolationBanner:CloudBrowserIsolationBanner", name, args ?? new CloudBrowserIsolationBannerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloudBrowserIsolationBanner(string name, Input<string> id, CloudBrowserIsolationBannerState? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/cloudBrowserIsolationBanner:CloudBrowserIsolationBanner", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CloudBrowserIsolationBanner resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloudBrowserIsolationBanner Get(string name, Input<string> id, CloudBrowserIsolationBannerState? state = null, CustomResourceOptions? options = null)
        {
            return new CloudBrowserIsolationBanner(name, id, state, options);
        }
    }

    public sealed class CloudBrowserIsolationBannerArgs : global::Pulumi.ResourceArgs
    {
        [Input("banner")]
        public Input<bool>? Banner { get; set; }

        [Input("logo")]
        public Input<string>? Logo { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationText")]
        public Input<string>? NotificationText { get; set; }

        [Input("notificationTitle")]
        public Input<string>? NotificationTitle { get; set; }

        [Input("persist")]
        public Input<bool>? Persist { get; set; }

        [Input("primaryColor")]
        public Input<string>? PrimaryColor { get; set; }

        [Input("textColor")]
        public Input<string>? TextColor { get; set; }

        public CloudBrowserIsolationBannerArgs()
        {
        }
        public static new CloudBrowserIsolationBannerArgs Empty => new CloudBrowserIsolationBannerArgs();
    }

    public sealed class CloudBrowserIsolationBannerState : global::Pulumi.ResourceArgs
    {
        [Input("banner")]
        public Input<bool>? Banner { get; set; }

        [Input("logo")]
        public Input<string>? Logo { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationText")]
        public Input<string>? NotificationText { get; set; }

        [Input("notificationTitle")]
        public Input<string>? NotificationTitle { get; set; }

        [Input("persist")]
        public Input<bool>? Persist { get; set; }

        [Input("primaryColor")]
        public Input<string>? PrimaryColor { get; set; }

        [Input("textColor")]
        public Input<string>? TextColor { get; set; }

        public CloudBrowserIsolationBannerState()
        {
        }
        public static new CloudBrowserIsolationBannerState Empty => new CloudBrowserIsolationBannerState();
    }
}
