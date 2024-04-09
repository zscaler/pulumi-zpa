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
    public static class GetPRAConsole
    {
        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/about-privileged-consoles)
        /// * [API documentation](https://help.zscaler.com/zpa/configuring-privileged-consoles-using-api)
        /// 
        /// The **zpa_pra_console_controller** data source gets information about a privileged remote access console created in the Zscaler Private Access cloud.
        /// This resource can then be referenced in an privileged access policy credential and a privileged access portal resource.
        /// </summary>
        public static Task<GetPRAConsoleResult> InvokeAsync(GetPRAConsoleArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPRAConsoleResult>("zpa:index/getPRAConsole:getPRAConsole", args ?? new GetPRAConsoleArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/about-privileged-consoles)
        /// * [API documentation](https://help.zscaler.com/zpa/configuring-privileged-consoles-using-api)
        /// 
        /// The **zpa_pra_console_controller** data source gets information about a privileged remote access console created in the Zscaler Private Access cloud.
        /// This resource can then be referenced in an privileged access policy credential and a privileged access portal resource.
        /// </summary>
        public static Output<GetPRAConsoleResult> Invoke(GetPRAConsoleInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPRAConsoleResult>("zpa:index/getPRAConsole:getPRAConsole", args ?? new GetPRAConsoleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPRAConsoleArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public string? Id { get; set; }

        public GetPRAConsoleArgs()
        {
        }
        public static new GetPRAConsoleArgs Empty => new GetPRAConsoleArgs();
    }

    public sealed class GetPRAConsoleInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        public GetPRAConsoleInvokeArgs()
        {
        }
        public static new GetPRAConsoleInvokeArgs Empty => new GetPRAConsoleInvokeArgs();
    }


    [OutputType]
    public sealed class GetPRAConsoleResult
    {
        public readonly string CreationTime;
        public readonly string Description;
        public readonly bool Enabled;
        public readonly string IconText;
        public readonly string? Id;
        public readonly string MicrotenantId;
        public readonly string MicrotenantName;
        public readonly string ModifiedBy;
        public readonly string ModifiedTime;
        /// <summary>
        /// - (Required) The name of the privileged console.
        /// </summary>
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetPRAConsolePraApplicationResult> PraApplications;
        public readonly ImmutableArray<Outputs.GetPRAConsolePraPortalResult> PraPortals;

        [OutputConstructor]
        private GetPRAConsoleResult(
            string creationTime,

            string description,

            bool enabled,

            string iconText,

            string? id,

            string microtenantId,

            string microtenantName,

            string modifiedBy,

            string modifiedTime,

            string name,

            ImmutableArray<Outputs.GetPRAConsolePraApplicationResult> praApplications,

            ImmutableArray<Outputs.GetPRAConsolePraPortalResult> praPortals)
        {
            CreationTime = creationTime;
            Description = description;
            Enabled = enabled;
            IconText = iconText;
            Id = id;
            MicrotenantId = microtenantId;
            MicrotenantName = microtenantName;
            ModifiedBy = modifiedBy;
            ModifiedTime = modifiedTime;
            Name = name;
            PraApplications = praApplications;
            PraPortals = praPortals;
        }
    }
}
