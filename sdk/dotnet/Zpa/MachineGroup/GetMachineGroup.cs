// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.MachineGroup
{
    public static class GetMachineGroup
    {
        /// <summary>
        /// Use the **zpa_machine_group** data source to get information about a machine group created in the Zscaler Private Access cloud. This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zpa.MachineGroup.GetMachineGroup.Invoke(new()
        ///     {
        ///         Name = "MGR01",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zpa.MachineGroup.GetMachineGroup.Invoke(new()
        ///     {
        ///         Id = "1234567890",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMachineGroupResult> InvokeAsync(GetMachineGroupArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMachineGroupResult>("zpa:machineGroup/getMachineGroup:getMachineGroup", args ?? new GetMachineGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zpa_machine_group** data source to get information about a machine group created in the Zscaler Private Access cloud. This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zpa.MachineGroup.GetMachineGroup.Invoke(new()
        ///     {
        ///         Name = "MGR01",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zpa.MachineGroup.GetMachineGroup.Invoke(new()
        ///     {
        ///         Id = "1234567890",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetMachineGroupResult> Invoke(GetMachineGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMachineGroupResult>("zpa:machineGroup/getMachineGroup:getMachineGroup", args ?? new GetMachineGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMachineGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the machine group to be exported.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of the machine group to be exported.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetMachineGroupArgs()
        {
        }
        public static new GetMachineGroupArgs Empty => new GetMachineGroupArgs();
    }

    public sealed class GetMachineGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the machine group to be exported.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the machine group to be exported.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetMachineGroupInvokeArgs()
        {
        }
        public static new GetMachineGroupInvokeArgs Empty => new GetMachineGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetMachineGroupResult
    {
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (bool)
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMachineGroupMachineResult> Machines;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ModifiedBy;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ModifiedTime;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private GetMachineGroupResult(
            string creationTime,

            string description,

            bool enabled,

            string? id,

            ImmutableArray<Outputs.GetMachineGroupMachineResult> machines,

            string modifiedBy,

            string modifiedTime,

            string? name)
        {
            CreationTime = creationTime;
            Description = description;
            Enabled = enabled;
            Id = id;
            Machines = machines;
            ModifiedBy = modifiedBy;
            ModifiedTime = modifiedTime;
            Name = name;
        }
    }
}
