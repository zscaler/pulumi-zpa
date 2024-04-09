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
    public static class GetPostureProfile
    {
        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/client-connector/about-device-posture-profiles)
        /// * [API documentation](https://help.zscaler.com/zpa/obtaining-posture-profile-details-using-api)
        /// 
        /// Use the **zpa_posture_profile** data source to get information about a posture profile created in the Zscaler Private Access Mobile Portal. This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.
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
        ///     var example1 = Zpa.GetPostureProfile.Invoke(new()
        ///     {
        ///         Name = "CrowdStrike_ZPA_ZTA_40",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
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
        ///     var example2 = Zpa.GetPostureProfile.Invoke(new()
        ///     {
        ///         Name = "Detect SentinelOne",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
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
        ///     var example3 = Zpa.GetPostureProfile.Invoke(new()
        ///     {
        ///         Name = "domain_joined",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// &gt; **NOTE** To query posture profiles that are associated with a specific Zscaler cloud, it is required to append the cloud name to the name of the posture profile as the below example:
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
        ///     var example1 = Zpa.GetPostureProfile.Invoke(new()
        ///     {
        ///         Name = "CrowdStrike_ZPA_ZTA_40 (zscalertwo.net)",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// &gt; **NOTE** When associating a posture profile with one of supported resources, the following parameter must be exported: ``posture_udid`` instead of the ``id`` of the resource.
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
        ///     var example1 = Zpa.GetPostureProfile.Invoke(new()
        ///     {
        ///         Name = "CrowdStrike_ZPA_ZTA_40 (zscalertwo.net)",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["zpaPostureProfile"] = example1.Apply(getPostureProfileResult =&gt; getPostureProfileResult.PostureUdid),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetPostureProfileResult> InvokeAsync(GetPostureProfileArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPostureProfileResult>("zpa:index/getPostureProfile:getPostureProfile", args ?? new GetPostureProfileArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/client-connector/about-device-posture-profiles)
        /// * [API documentation](https://help.zscaler.com/zpa/obtaining-posture-profile-details-using-api)
        /// 
        /// Use the **zpa_posture_profile** data source to get information about a posture profile created in the Zscaler Private Access Mobile Portal. This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.
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
        ///     var example1 = Zpa.GetPostureProfile.Invoke(new()
        ///     {
        ///         Name = "CrowdStrike_ZPA_ZTA_40",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
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
        ///     var example2 = Zpa.GetPostureProfile.Invoke(new()
        ///     {
        ///         Name = "Detect SentinelOne",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
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
        ///     var example3 = Zpa.GetPostureProfile.Invoke(new()
        ///     {
        ///         Name = "domain_joined",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// &gt; **NOTE** To query posture profiles that are associated with a specific Zscaler cloud, it is required to append the cloud name to the name of the posture profile as the below example:
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
        ///     var example1 = Zpa.GetPostureProfile.Invoke(new()
        ///     {
        ///         Name = "CrowdStrike_ZPA_ZTA_40 (zscalertwo.net)",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// &gt; **NOTE** When associating a posture profile with one of supported resources, the following parameter must be exported: ``posture_udid`` instead of the ``id`` of the resource.
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
        ///     var example1 = Zpa.GetPostureProfile.Invoke(new()
        ///     {
        ///         Name = "CrowdStrike_ZPA_ZTA_40 (zscalertwo.net)",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["zpaPostureProfile"] = example1.Apply(getPostureProfileResult =&gt; getPostureProfileResult.PostureUdid),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetPostureProfileResult> Invoke(GetPostureProfileInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPostureProfileResult>("zpa:index/getPostureProfile:getPostureProfile", args ?? new GetPostureProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPostureProfileArgs : global::Pulumi.InvokeArgs
    {
        [Input("name")]
        public string? Name { get; set; }

        public GetPostureProfileArgs()
        {
        }
        public static new GetPostureProfileArgs Empty => new GetPostureProfileArgs();
    }

    public sealed class GetPostureProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetPostureProfileInvokeArgs()
        {
        }
        public static new GetPostureProfileInvokeArgs Empty => new GetPostureProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetPostureProfileResult
    {
        public readonly string CreationTime;
        public readonly string Domain;
        public readonly string Id;
        public readonly string MasterCustomerId;
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        public readonly string? Name;
        public readonly string PostureUdid;
        public readonly string ZscalerCloud;
        public readonly string ZscalerCustomerId;

        [OutputConstructor]
        private GetPostureProfileResult(
            string creationTime,

            string domain,

            string id,

            string masterCustomerId,

            string modifiedTime,

            string modifiedby,

            string? name,

            string postureUdid,

            string zscalerCloud,

            string zscalerCustomerId)
        {
            CreationTime = creationTime;
            Domain = domain;
            Id = id;
            MasterCustomerId = masterCustomerId;
            ModifiedTime = modifiedTime;
            Modifiedby = modifiedby;
            Name = name;
            PostureUdid = postureUdid;
            ZscalerCloud = zscalerCloud;
            ZscalerCustomerId = zscalerCustomerId;
        }
    }
}
