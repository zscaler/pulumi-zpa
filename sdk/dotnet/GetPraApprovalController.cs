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
    [Obsolete(@"zpa.index/getpraapprovalcontroller.getPraApprovalController has been deprecated in favor of zpa.index/getpraapproval.getPRAApproval")]
    public static class GetPraApprovalController
    {
        public static Task<GetPraApprovalControllerResult> InvokeAsync(GetPraApprovalControllerArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPraApprovalControllerResult>("zpa:index/getPraApprovalController:getPraApprovalController", args ?? new GetPraApprovalControllerArgs(), options.WithDefaults());

        public static Output<GetPraApprovalControllerResult> Invoke(GetPraApprovalControllerInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPraApprovalControllerResult>("zpa:index/getPraApprovalController:getPraApprovalController", args ?? new GetPraApprovalControllerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPraApprovalControllerArgs : global::Pulumi.InvokeArgs
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

        public GetPraApprovalControllerArgs()
        {
        }
        public static new GetPraApprovalControllerArgs Empty => new GetPraApprovalControllerArgs();
    }

    public sealed class GetPraApprovalControllerInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetPraApprovalControllerInvokeArgs()
        {
        }
        public static new GetPraApprovalControllerInvokeArgs Empty => new GetPraApprovalControllerInvokeArgs();
    }


    [OutputType]
    public sealed class GetPraApprovalControllerResult
    {
        public readonly ImmutableArray<Outputs.GetPraApprovalControllerApplicationResult> Applications;
        public readonly string CreationTime;
        public readonly ImmutableArray<string> EmailIds;
        public readonly string EndTime;
        public readonly string? Id;
        public readonly string MicrotenantId;
        public readonly string ModifiedBy;
        public readonly string ModifiedTime;
        public readonly string StartTime;
        public readonly string Status;
        public readonly ImmutableArray<Outputs.GetPraApprovalControllerWorkingHourResult> WorkingHours;

        [OutputConstructor]
        private GetPraApprovalControllerResult(
            ImmutableArray<Outputs.GetPraApprovalControllerApplicationResult> applications,

            string creationTime,

            ImmutableArray<string> emailIds,

            string endTime,

            string? id,

            string microtenantId,

            string modifiedBy,

            string modifiedTime,

            string startTime,

            string status,

            ImmutableArray<Outputs.GetPraApprovalControllerWorkingHourResult> workingHours)
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