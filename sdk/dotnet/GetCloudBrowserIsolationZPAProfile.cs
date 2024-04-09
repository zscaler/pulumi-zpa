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
    public static class GetCloudBrowserIsolationZPAProfile
    {
        /// <summary>
        /// Use the **zpa_cloud_browser_isolation_zpa_profile** data source to get information about an isolation profile in the Zscaler Private Access cloud. This data source is required when configuring an isolation policy rule resource
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
        ///     var @this = Zpa.GetCloudBrowserIsolationZPAProfile.Invoke(new()
        ///     {
        ///         Name = "ZPA_Profile",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetCloudBrowserIsolationZPAProfileResult> InvokeAsync(GetCloudBrowserIsolationZPAProfileArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCloudBrowserIsolationZPAProfileResult>("zpa:index/getCloudBrowserIsolationZPAProfile:getCloudBrowserIsolationZPAProfile", args ?? new GetCloudBrowserIsolationZPAProfileArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zpa_cloud_browser_isolation_zpa_profile** data source to get information about an isolation profile in the Zscaler Private Access cloud. This data source is required when configuring an isolation policy rule resource
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
        ///     var @this = Zpa.GetCloudBrowserIsolationZPAProfile.Invoke(new()
        ///     {
        ///         Name = "ZPA_Profile",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetCloudBrowserIsolationZPAProfileResult> Invoke(GetCloudBrowserIsolationZPAProfileInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCloudBrowserIsolationZPAProfileResult>("zpa:index/getCloudBrowserIsolationZPAProfile:getCloudBrowserIsolationZPAProfile", args ?? new GetCloudBrowserIsolationZPAProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudBrowserIsolationZPAProfileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// - (String) This field defines the name of the isolation profile.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetCloudBrowserIsolationZPAProfileArgs()
        {
        }
        public static new GetCloudBrowserIsolationZPAProfileArgs Empty => new GetCloudBrowserIsolationZPAProfileArgs();
    }

    public sealed class GetCloudBrowserIsolationZPAProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// - (String) This field defines the name of the isolation profile.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetCloudBrowserIsolationZPAProfileInvokeArgs()
        {
        }
        public static new GetCloudBrowserIsolationZPAProfileInvokeArgs Empty => new GetCloudBrowserIsolationZPAProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloudBrowserIsolationZPAProfileResult
    {
        public readonly string CbiProfileId;
        public readonly string CbiTenantId;
        public readonly string CbiUrl;
        public readonly string CreationTime;
        public readonly string Description;
        public readonly bool Enabled;
        public readonly string Id;
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        /// <summary>
        /// - (String) This field defines the name of the isolation profile.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private GetCloudBrowserIsolationZPAProfileResult(
            string cbiProfileId,

            string cbiTenantId,

            string cbiUrl,

            string creationTime,

            string description,

            bool enabled,

            string id,

            string modifiedTime,

            string modifiedby,

            string? name)
        {
            CbiProfileId = cbiProfileId;
            CbiTenantId = cbiTenantId;
            CbiUrl = cbiUrl;
            CreationTime = creationTime;
            Description = description;
            Enabled = enabled;
            Id = id;
            ModifiedTime = modifiedTime;
            Modifiedby = modifiedby;
            Name = name;
        }
    }
}
