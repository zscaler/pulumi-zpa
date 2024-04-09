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
    public static class GetPRAApproval
    {
        /// <summary>
        /// Use the **zpa_pra_approval_controller** data source to get information about a privileged remote access approval created in the Zscaler Private Access cloud.
        /// </summary>
        public static Task<GetPRAApprovalResult> InvokeAsync(GetPRAApprovalArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPRAApprovalResult>("zpa:index/getPRAApproval:getPRAApproval", args ?? new GetPRAApprovalArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zpa_pra_approval_controller** data source to get information about a privileged remote access approval created in the Zscaler Private Access cloud.
        /// </summary>
        public static Output<GetPRAApprovalResult> Invoke(GetPRAApprovalInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPRAApprovalResult>("zpa:index/getPRAApproval:getPRAApproval", args ?? new GetPRAApprovalInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPRAApprovalArgs : global::Pulumi.InvokeArgs
    {
        [Input("emailIds")]
        private List<string>? _emailIds;
        public List<string> EmailIds
        {
            get => _emailIds ?? (_emailIds = new List<string>());
            set => _emailIds = value;
        }

        [Input("id")]
        public string? Id { get; set; }

        public GetPRAApprovalArgs()
        {
        }
        public static new GetPRAApprovalArgs Empty => new GetPRAApprovalArgs();
    }

    public sealed class GetPRAApprovalInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("emailIds")]
        private InputList<string>? _emailIds;
        public InputList<string> EmailIds
        {
            get => _emailIds ?? (_emailIds = new InputList<string>());
            set => _emailIds = value;
        }

        [Input("id")]
        public Input<string>? Id { get; set; }

        public GetPRAApprovalInvokeArgs()
        {
        }
        public static new GetPRAApprovalInvokeArgs Empty => new GetPRAApprovalInvokeArgs();
    }


    [OutputType]
    public sealed class GetPRAApprovalResult
    {
        public readonly ImmutableArray<Outputs.GetPRAApprovalApplicationResult> Applications;
        public readonly string CreationTime;
        public readonly ImmutableArray<string> EmailIds;
        public readonly string EndTime;
        public readonly string? Id;
        /// <summary>
        /// (string)  The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass microtenantId as 0 when making requests to retrieve data from the Default Microtenant. Pass microtenantId as null to retrieve data from all customers associated with the tenant.
        /// </summary>
        public readonly string MicrotenantId;
        public readonly string ModifiedBy;
        public readonly string ModifiedTime;
        public readonly string StartTime;
        /// <summary>
        /// (string) The status of the privileged approval. The supported values are:
        /// </summary>
        public readonly string Status;
        public readonly ImmutableArray<Outputs.GetPRAApprovalWorkingHourResult> WorkingHours;

        [OutputConstructor]
        private GetPRAApprovalResult(
            ImmutableArray<Outputs.GetPRAApprovalApplicationResult> applications,

            string creationTime,

            ImmutableArray<string> emailIds,

            string endTime,

            string? id,

            string microtenantId,

            string modifiedBy,

            string modifiedTime,

            string startTime,

            string status,

            ImmutableArray<Outputs.GetPRAApprovalWorkingHourResult> workingHours)
        {
            Applications = applications;
            CreationTime = creationTime;
            EmailIds = emailIds;
            EndTime = endTime;
            Id = id;
            MicrotenantId = microtenantId;
            ModifiedBy = modifiedBy;
            ModifiedTime = modifiedTime;
            StartTime = startTime;
            Status = status;
            WorkingHours = workingHours;
        }
    }
}