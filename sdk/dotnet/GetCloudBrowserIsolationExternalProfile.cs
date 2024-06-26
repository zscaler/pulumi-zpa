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
    public static class GetCloudBrowserIsolationExternalProfile
    {
        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/isolation/about-custom-root-certificates-cloud-browser-isolation)
        /// 
        /// Use the **zpa_cloud_browser_isolation_external_profile** data source to get information about Cloud Browser Isolation external profile. This data source information can then be used in as part of `zpa.PolicyAccessIsolationRule` when the `action` attribute is set to `ISOLATE`.
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
        ///     var @this = Zpa.GetCloudBrowserIsolationExternalProfile.Invoke(new()
        ///     {
        ///         Name = "Example",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetCloudBrowserIsolationExternalProfileResult> InvokeAsync(GetCloudBrowserIsolationExternalProfileArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCloudBrowserIsolationExternalProfileResult>("zpa:index/getCloudBrowserIsolationExternalProfile:getCloudBrowserIsolationExternalProfile", args ?? new GetCloudBrowserIsolationExternalProfileArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/isolation/about-custom-root-certificates-cloud-browser-isolation)
        /// 
        /// Use the **zpa_cloud_browser_isolation_external_profile** data source to get information about Cloud Browser Isolation external profile. This data source information can then be used in as part of `zpa.PolicyAccessIsolationRule` when the `action` attribute is set to `ISOLATE`.
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
        ///     var @this = Zpa.GetCloudBrowserIsolationExternalProfile.Invoke(new()
        ///     {
        ///         Name = "Example",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetCloudBrowserIsolationExternalProfileResult> Invoke(GetCloudBrowserIsolationExternalProfileInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCloudBrowserIsolationExternalProfileResult>("zpa:index/getCloudBrowserIsolationExternalProfile:getCloudBrowserIsolationExternalProfile", args ?? new GetCloudBrowserIsolationExternalProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudBrowserIsolationExternalProfileArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public string? Id { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public GetCloudBrowserIsolationExternalProfileArgs()
        {
        }
        public static new GetCloudBrowserIsolationExternalProfileArgs Empty => new GetCloudBrowserIsolationExternalProfileArgs();
    }

    public sealed class GetCloudBrowserIsolationExternalProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetCloudBrowserIsolationExternalProfileInvokeArgs()
        {
        }
        public static new GetCloudBrowserIsolationExternalProfileInvokeArgs Empty => new GetCloudBrowserIsolationExternalProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloudBrowserIsolationExternalProfileResult
    {
        public readonly string Description;
        public readonly string Href;
        public readonly string? Id;
        public readonly bool IsDefault;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.GetCloudBrowserIsolationExternalProfileRegionResult> Regions;
        public readonly ImmutableArray<Outputs.GetCloudBrowserIsolationExternalProfileSecurityControlResult> SecurityControls;
        public readonly ImmutableArray<Outputs.GetCloudBrowserIsolationExternalProfileUserExperienceResult> UserExperiences;

        [OutputConstructor]
        private GetCloudBrowserIsolationExternalProfileResult(
            string description,

            string href,

            string? id,

            bool isDefault,

            string? name,

            ImmutableArray<Outputs.GetCloudBrowserIsolationExternalProfileRegionResult> regions,

            ImmutableArray<Outputs.GetCloudBrowserIsolationExternalProfileSecurityControlResult> securityControls,

            ImmutableArray<Outputs.GetCloudBrowserIsolationExternalProfileUserExperienceResult> userExperiences)
        {
            Description = description;
            Href = href;
            Id = id;
            IsDefault = isDefault;
            Name = name;
            Regions = regions;
            SecurityControls = securityControls;
            UserExperiences = userExperiences;
        }
    }
}
