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
    public static class GetServerGroup
    {
        /// <summary>
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
        ///     var example = Zpa.GetServerGroup.Invoke(new()
        ///     {
        ///         Name = "server_group_name",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetServerGroupResult> InvokeAsync(GetServerGroupArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServerGroupResult>("zpa:index/getServerGroup:getServerGroup", args ?? new GetServerGroupArgs(), options.WithDefaults());

        /// <summary>
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
        ///     var example = Zpa.GetServerGroup.Invoke(new()
        ///     {
        ///         Name = "server_group_name",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetServerGroupResult> Invoke(GetServerGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServerGroupResult>("zpa:index/getServerGroup:getServerGroup", args ?? new GetServerGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServerGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the server group to be exported.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// (string) The ID of the microtenant the resource is to be associated with.
        /// </summary>
        [Input("microtenantId")]
        public string? MicrotenantId { get; set; }

        /// <summary>
        /// (string) The name of the microtenant the resource is to be associated with.
        /// </summary>
        [Input("microtenantName")]
        public string? MicrotenantName { get; set; }

        /// <summary>
        /// The name of the server group to be exported.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetServerGroupArgs()
        {
        }
        public static new GetServerGroupArgs Empty => new GetServerGroupArgs();
    }

    public sealed class GetServerGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the server group to be exported.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// (string) The ID of the microtenant the resource is to be associated with.
        /// </summary>
        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// (string) The name of the microtenant the resource is to be associated with.
        /// </summary>
        [Input("microtenantName")]
        public Input<string>? MicrotenantName { get; set; }

        /// <summary>
        /// The name of the server group to be exported.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetServerGroupInvokeArgs()
        {
        }
        public static new GetServerGroupInvokeArgs Empty => new GetServerGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetServerGroupResult
    {
        /// <summary>
        /// (string)This field is a json array of app-connector-id only.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServerGroupAppConnectorGroupResult> AppConnectorGroups;
        public readonly ImmutableArray<Outputs.GetServerGroupApplicationResult> Applications;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ConfigSpace;
        public readonly string CreationTime;
        /// <summary>
        /// (string) This field is the description of the server group.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (bool) This field controls dynamic discovery of the servers.
        /// </summary>
        public readonly bool DynamicDiscovery;
        /// <summary>
        /// (bool) This field defines if the server group is enabled or disabled.
        /// </summary>
        public readonly bool Enabled;
        public readonly string? Id;
        /// <summary>
        /// (bool)
        /// </summary>
        public readonly bool IpAnchored;
        /// <summary>
        /// (string) The ID of the microtenant the resource is to be associated with.
        /// </summary>
        public readonly string? MicrotenantId;
        /// <summary>
        /// (string) The name of the microtenant the resource is to be associated with.
        /// </summary>
        public readonly string? MicrotenantName;
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.GetServerGroupServerResult> Servers;

        [OutputConstructor]
        private GetServerGroupResult(
            ImmutableArray<Outputs.GetServerGroupAppConnectorGroupResult> appConnectorGroups,

            ImmutableArray<Outputs.GetServerGroupApplicationResult> applications,

            string configSpace,

            string creationTime,

            string description,

            bool dynamicDiscovery,

            bool enabled,

            string? id,

            bool ipAnchored,

            string? microtenantId,

            string? microtenantName,

            string modifiedTime,

            string modifiedby,

            string? name,

            ImmutableArray<Outputs.GetServerGroupServerResult> servers)
        {
            AppConnectorGroups = appConnectorGroups;
            Applications = applications;
            ConfigSpace = configSpace;
            CreationTime = creationTime;
            Description = description;
            DynamicDiscovery = dynamicDiscovery;
            Enabled = enabled;
            Id = id;
            IpAnchored = ipAnchored;
            MicrotenantId = microtenantId;
            MicrotenantName = microtenantName;
            ModifiedTime = modifiedTime;
            Modifiedby = modifiedby;
            Name = name;
            Servers = servers;
        }
    }
}