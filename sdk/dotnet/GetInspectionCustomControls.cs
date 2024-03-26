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
    public static class GetInspectionCustomControls
    {
        public static Task<GetInspectionCustomControlsResult> InvokeAsync(GetInspectionCustomControlsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInspectionCustomControlsResult>("zpa:index/getInspectionCustomControls:getInspectionCustomControls", args ?? new GetInspectionCustomControlsArgs(), options.WithDefaults());

        public static Output<GetInspectionCustomControlsResult> Invoke(GetInspectionCustomControlsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInspectionCustomControlsResult>("zpa:index/getInspectionCustomControls:getInspectionCustomControls", args ?? new GetInspectionCustomControlsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInspectionCustomControlsArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public string? Id { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public GetInspectionCustomControlsArgs()
        {
        }
        public static new GetInspectionCustomControlsArgs Empty => new GetInspectionCustomControlsArgs();
    }

    public sealed class GetInspectionCustomControlsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetInspectionCustomControlsInvokeArgs()
        {
        }
        public static new GetInspectionCustomControlsInvokeArgs Empty => new GetInspectionCustomControlsInvokeArgs();
    }


    [OutputType]
    public sealed class GetInspectionCustomControlsResult
    {
        public readonly string Action;
        public readonly string ActionValue;
        public readonly string ControlNumber;
        public readonly string ControlRuleJson;
        public readonly string ControlType;
        public readonly string CreationTime;
        public readonly string DefaultAction;
        public readonly string DefaultActionValue;
        public readonly string Description;
        public readonly string Id;
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        public readonly string Name;
        public readonly string ParanoiaLevel;
        public readonly string ProtocolType;
        public readonly ImmutableArray<Outputs.GetInspectionCustomControlsRuleResult> Rules;
        public readonly string Severity;
        public readonly string Type;
        public readonly string Version;

        [OutputConstructor]
        private GetInspectionCustomControlsResult(
            string action,

            string actionValue,

            string controlNumber,

            string controlRuleJson,

            string controlType,

            string creationTime,

            string defaultAction,

            string defaultActionValue,

            string description,

            string id,

            string modifiedTime,

            string modifiedby,

            string name,

            string paranoiaLevel,

            string protocolType,

            ImmutableArray<Outputs.GetInspectionCustomControlsRuleResult> rules,

            string severity,

            string type,

            string version)
        {
            Action = action;
            ActionValue = actionValue;
            ControlNumber = controlNumber;
            ControlRuleJson = controlRuleJson;
            ControlType = controlType;
            CreationTime = creationTime;
            DefaultAction = defaultAction;
            DefaultActionValue = defaultActionValue;
            Description = description;
            Id = id;
            ModifiedTime = modifiedTime;
            Modifiedby = modifiedby;
            Name = name;
            ParanoiaLevel = paranoiaLevel;
            ProtocolType = protocolType;
            Rules = rules;
            Severity = severity;
            Type = type;
            Version = version;
        }
    }
}
