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
    public static class GetPRACredential
    {
        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/about-privileged-credentials)
        /// * [API documentation](https://help.zscaler.com/zpa/configuring-privileged-credentials-using-api)
        /// 
        /// The **zpa_pra_credential_controller** resource creates a privileged remote access credential in the Zscaler Private Access cloud. This resource can then be referenced in an privileged access policy resource.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zpa = Zscaler.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // Retrieves PRA Credential By ID
        ///     var @this = new Zpa.PRACredential("this");
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetPRACredentialResult> InvokeAsync(GetPRACredentialArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPRACredentialResult>("zpa:index/getPRACredential:getPRACredential", args ?? new GetPRACredentialArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/about-privileged-credentials)
        /// * [API documentation](https://help.zscaler.com/zpa/configuring-privileged-credentials-using-api)
        /// 
        /// The **zpa_pra_credential_controller** resource creates a privileged remote access credential in the Zscaler Private Access cloud. This resource can then be referenced in an privileged access policy resource.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zpa = Zscaler.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // Retrieves PRA Credential By ID
        ///     var @this = new Zpa.PRACredential("this");
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetPRACredentialResult> Invoke(GetPRACredentialInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPRACredentialResult>("zpa:index/getPRACredential:getPRACredential", args ?? new GetPRACredentialInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPRACredentialArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// - (String) The name of the privileged credential.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetPRACredentialArgs()
        {
        }
        public static new GetPRACredentialArgs Empty => new GetPRACredentialArgs();
    }

    public sealed class GetPRACredentialInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// - (String) The name of the privileged credential.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetPRACredentialInvokeArgs()
        {
        }
        public static new GetPRACredentialInvokeArgs Empty => new GetPRACredentialInvokeArgs();
    }


    [OutputType]
    public sealed class GetPRACredentialResult
    {
        public readonly string CreationTime;
        public readonly string CredentialType;
        public readonly string Description;
        public readonly string? Id;
        public readonly string LastCredentialResetTime;
        public readonly string MicrotenantId;
        public readonly string MicrotenantName;
        public readonly string ModifiedBy;
        public readonly string ModifiedTime;
        /// <summary>
        /// - (String) The name of the privileged credential.
        /// </summary>
        public readonly string? Name;
        public readonly string UserDomain;
        public readonly string Username;

        [OutputConstructor]
        private GetPRACredentialResult(
            string creationTime,

            string credentialType,

            string description,

            string? id,

            string lastCredentialResetTime,

            string microtenantId,

            string microtenantName,

            string modifiedBy,

            string modifiedTime,

            string? name,

            string userDomain,

            string username)
        {
            CreationTime = creationTime;
            CredentialType = credentialType;
            Description = description;
            Id = id;
            LastCredentialResetTime = lastCredentialResetTime;
            MicrotenantId = microtenantId;
            MicrotenantName = microtenantName;
            ModifiedBy = modifiedBy;
            ModifiedTime = modifiedTime;
            Name = name;
            UserDomain = userDomain;
            Username = username;
        }
    }
}