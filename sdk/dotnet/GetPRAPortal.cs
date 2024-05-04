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
    public static class GetPRAPortal
    {
        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/about-privileged-portals)
        /// * [API documentation](https://help.zscaler.com/zpa/configuring-privileged-portals-using-api)
        /// 
        /// Use the **zpa_pra_portal_controller** data source to get information about a privileged remote access portal created in the Zscaler Private Access cloud. This data source can then be referenced in an privileged remote access console resource.
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
        ///     var @this = Zpa.GetPRAPortal.Invoke(new()
        ///     {
        ///         Name = "Example",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetPRAPortalResult> InvokeAsync(GetPRAPortalArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPRAPortalResult>("zpa:index/getPRAPortal:getPRAPortal", args ?? new GetPRAPortalArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/about-privileged-portals)
        /// * [API documentation](https://help.zscaler.com/zpa/configuring-privileged-portals-using-api)
        /// 
        /// Use the **zpa_pra_portal_controller** data source to get information about a privileged remote access portal created in the Zscaler Private Access cloud. This data source can then be referenced in an privileged remote access console resource.
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
        ///     var @this = Zpa.GetPRAPortal.Invoke(new()
        ///     {
        ///         Name = "Example",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetPRAPortalResult> Invoke(GetPRAPortalInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPRAPortalResult>("zpa:index/getPRAPortal:getPRAPortal", args ?? new GetPRAPortalInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPRAPortalArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public string? Id { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public GetPRAPortalArgs()
        {
        }
        public static new GetPRAPortalArgs Empty => new GetPRAPortalArgs();
    }

    public sealed class GetPRAPortalInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetPRAPortalInvokeArgs()
        {
        }
        public static new GetPRAPortalInvokeArgs Empty => new GetPRAPortalInvokeArgs();
    }


    [OutputType]
    public sealed class GetPRAPortalResult
    {
        public readonly string CertificateId;
        public readonly string CertificateName;
        public readonly string Cname;
        public readonly string CreationTime;
        public readonly string Description;
        public readonly string Domain;
        public readonly bool Enabled;
        public readonly string? Id;
        public readonly string MicrotenantId;
        public readonly string MicrotenantName;
        public readonly string ModifiedBy;
        public readonly string ModifiedTime;
        public readonly string? Name;
        public readonly string UserNotification;
        public readonly bool UserNotificationEnabled;

        [OutputConstructor]
        private GetPRAPortalResult(
            string certificateId,

            string certificateName,

            string cname,

            string creationTime,

            string description,

            string domain,

            bool enabled,

            string? id,

            string microtenantId,

            string microtenantName,

            string modifiedBy,

            string modifiedTime,

            string? name,

            string userNotification,

            bool userNotificationEnabled)
        {
            CertificateId = certificateId;
            CertificateName = certificateName;
            Cname = cname;
            CreationTime = creationTime;
            Description = description;
            Domain = domain;
            Enabled = enabled;
            Id = id;
            MicrotenantId = microtenantId;
            MicrotenantName = microtenantName;
            ModifiedBy = modifiedBy;
            ModifiedTime = modifiedTime;
            Name = name;
            UserNotification = userNotification;
            UserNotificationEnabled = userNotificationEnabled;
        }
    }
}
