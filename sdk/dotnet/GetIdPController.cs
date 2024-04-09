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
    public static class GetIdPController
    {
        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/identity-management)
        /// * [API documentation](https://help.zscaler.com/zpa/obtaining-idp-configuration-details-using-api)
        /// 
        /// Use the **zpa_idp_controller** data source to get information about an Identity Provider created in the Zscaler Private Access cloud. This data source is required when creating:
        /// 
        /// * Access policy Rules
        /// * Access policy timeout rules
        /// * Access policy forwarding rules
        /// * Access policy inspection rules
        /// * Access policy isolation rules
        /// * Access policy privileged credentials rules
        /// * Access policy privileged capabilities rules
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
        ///     var example = Zpa.GetIdPController.Invoke(new()
        ///     {
        ///         Name = "idp_name",
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
        ///     var example = Zpa.GetIdPController.Invoke(new()
        ///     {
        ///         Id = "1234567890",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetIdPControllerResult> InvokeAsync(GetIdPControllerArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIdPControllerResult>("zpa:index/getIdPController:getIdPController", args ?? new GetIdPControllerArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/identity-management)
        /// * [API documentation](https://help.zscaler.com/zpa/obtaining-idp-configuration-details-using-api)
        /// 
        /// Use the **zpa_idp_controller** data source to get information about an Identity Provider created in the Zscaler Private Access cloud. This data source is required when creating:
        /// 
        /// * Access policy Rules
        /// * Access policy timeout rules
        /// * Access policy forwarding rules
        /// * Access policy inspection rules
        /// * Access policy isolation rules
        /// * Access policy privileged credentials rules
        /// * Access policy privileged capabilities rules
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
        ///     var example = Zpa.GetIdPController.Invoke(new()
        ///     {
        ///         Name = "idp_name",
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
        ///     var example = Zpa.GetIdPController.Invoke(new()
        ///     {
        ///         Id = "1234567890",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetIdPControllerResult> Invoke(GetIdPControllerInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdPControllerResult>("zpa:index/getIdPController:getIdPController", args ?? new GetIdPControllerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIdPControllerArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public string? Id { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public GetIdPControllerArgs()
        {
        }
        public static new GetIdPControllerArgs Empty => new GetIdPControllerArgs();
    }

    public sealed class GetIdPControllerInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetIdPControllerInvokeArgs()
        {
        }
        public static new GetIdPControllerInvokeArgs Empty => new GetIdPControllerInvokeArgs();
    }


    [OutputType]
    public sealed class GetIdPControllerResult
    {
        public readonly ImmutableArray<Outputs.GetIdPControllerAdminMetadataResult> AdminMetadatas;
        public readonly string AdminSpSigningCertId;
        public readonly string AutoProvision;
        public readonly string CreationTime;
        public readonly string Description;
        public readonly bool DisableSamlBasedPolicy;
        public readonly ImmutableArray<string> DomainLists;
        public readonly string EnableArbitraryAuthDomains;
        public readonly bool EnableScimBasedPolicy;
        public readonly bool Enabled;
        public readonly bool ForceAuth;
        public readonly string Id;
        public readonly string IdpEntityId;
        public readonly bool LoginHint;
        public readonly string LoginNameAttribute;
        public readonly string LoginUrl;
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        public readonly string Name;
        public readonly bool ReauthOnUserUpdate;
        public readonly bool RedirectBinding;
        public readonly bool ScimEnabled;
        public readonly string ScimServiceProviderEndpoint;
        public readonly bool ScimSharedSecretExists;
        public readonly string SignSamlRequest;
        public readonly ImmutableArray<string> SsoTypes;
        public readonly bool UseCustomSpMetadata;
        public readonly ImmutableArray<Outputs.GetIdPControllerUserMetadataResult> UserMetadatas;
        public readonly string UserSpSigningCertId;

        [OutputConstructor]
        private GetIdPControllerResult(
            ImmutableArray<Outputs.GetIdPControllerAdminMetadataResult> adminMetadatas,

            string adminSpSigningCertId,

            string autoProvision,

            string creationTime,

            string description,

            bool disableSamlBasedPolicy,

            ImmutableArray<string> domainLists,

            string enableArbitraryAuthDomains,

            bool enableScimBasedPolicy,

            bool enabled,

            bool forceAuth,

            string id,

            string idpEntityId,

            bool loginHint,

            string loginNameAttribute,

            string loginUrl,

            string modifiedTime,

            string modifiedby,

            string name,

            bool reauthOnUserUpdate,

            bool redirectBinding,

            bool scimEnabled,

            string scimServiceProviderEndpoint,

            bool scimSharedSecretExists,

            string signSamlRequest,

            ImmutableArray<string> ssoTypes,

            bool useCustomSpMetadata,

            ImmutableArray<Outputs.GetIdPControllerUserMetadataResult> userMetadatas,

            string userSpSigningCertId)
        {
            AdminMetadatas = adminMetadatas;
            AdminSpSigningCertId = adminSpSigningCertId;
            AutoProvision = autoProvision;
            CreationTime = creationTime;
            Description = description;
            DisableSamlBasedPolicy = disableSamlBasedPolicy;
            DomainLists = domainLists;
            EnableArbitraryAuthDomains = enableArbitraryAuthDomains;
            EnableScimBasedPolicy = enableScimBasedPolicy;
            Enabled = enabled;
            ForceAuth = forceAuth;
            Id = id;
            IdpEntityId = idpEntityId;
            LoginHint = loginHint;
            LoginNameAttribute = loginNameAttribute;
            LoginUrl = loginUrl;
            ModifiedTime = modifiedTime;
            Modifiedby = modifiedby;
            Name = name;
            ReauthOnUserUpdate = reauthOnUserUpdate;
            RedirectBinding = redirectBinding;
            ScimEnabled = scimEnabled;
            ScimServiceProviderEndpoint = scimServiceProviderEndpoint;
            ScimSharedSecretExists = scimSharedSecretExists;
            SignSamlRequest = signSamlRequest;
            SsoTypes = ssoTypes;
            UseCustomSpMetadata = useCustomSpMetadata;
            UserMetadatas = userMetadatas;
            UserSpSigningCertId = userSpSigningCertId;
        }
    }
}
