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
    public static class GetCloudBrowserIsolationRegion
    {
        /// <summary>
        /// Use the **zpa_cloud_browser_isolation_region** data source to get information about Cloud Browser Isolation regions such as ID and Name. This data source information is required as part of the attribute `region_ids` when creating an Cloud Browser Isolation External Profile ``zpa.CloudBrowserIsolationExternalProfile``
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
        ///     var @this = Zpa.GetCloudBrowserIsolationRegion.Invoke(new()
        ///     {
        ///         Name = "Singapore",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetCloudBrowserIsolationRegionResult> InvokeAsync(GetCloudBrowserIsolationRegionArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCloudBrowserIsolationRegionResult>("zpa:index/getCloudBrowserIsolationRegion:getCloudBrowserIsolationRegion", args ?? new GetCloudBrowserIsolationRegionArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zpa_cloud_browser_isolation_region** data source to get information about Cloud Browser Isolation regions such as ID and Name. This data source information is required as part of the attribute `region_ids` when creating an Cloud Browser Isolation External Profile ``zpa.CloudBrowserIsolationExternalProfile``
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
        ///     var @this = Zpa.GetCloudBrowserIsolationRegion.Invoke(new()
        ///     {
        ///         Name = "Singapore",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetCloudBrowserIsolationRegionResult> Invoke(GetCloudBrowserIsolationRegionInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCloudBrowserIsolationRegionResult>("zpa:index/getCloudBrowserIsolationRegion:getCloudBrowserIsolationRegion", args ?? new GetCloudBrowserIsolationRegionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudBrowserIsolationRegionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the CBI region to be exported.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of the CBI region to be exported.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetCloudBrowserIsolationRegionArgs()
        {
        }
        public static new GetCloudBrowserIsolationRegionArgs Empty => new GetCloudBrowserIsolationRegionArgs();
    }

    public sealed class GetCloudBrowserIsolationRegionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the CBI region to be exported.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the CBI region to be exported.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetCloudBrowserIsolationRegionInvokeArgs()
        {
        }
        public static new GetCloudBrowserIsolationRegionInvokeArgs Empty => new GetCloudBrowserIsolationRegionInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloudBrowserIsolationRegionResult
    {
        /// <summary>
        /// (string) - ID information of the CBI region
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (string) - Name of the CBI region
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetCloudBrowserIsolationRegionResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}