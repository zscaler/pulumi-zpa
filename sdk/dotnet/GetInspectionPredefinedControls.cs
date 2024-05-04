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
    public static class GetInspectionPredefinedControls
    {
        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/about-custom-controls)
        /// * [API documentation](https://help.zscaler.com/zpa/configuring-appprotection-controls-using-api)
        /// 
        /// Use the **zpa_inspection_predefined_controls** data source to get information about an OWASP predefined control and prefedined control version. This data source is required when creating an inspection profile.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zpa.GetInspectionPredefinedControls.Invoke(new()
        ///     {
        ///         Name = "Failed to parse request body",
        ///         Version = "OWASP_CRS/3.3.0",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["zpaInspectionPredefinedControls"] = example,
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetInspectionPredefinedControlsResult> InvokeAsync(GetInspectionPredefinedControlsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInspectionPredefinedControlsResult>("zpa:index/getInspectionPredefinedControls:getInspectionPredefinedControls", args ?? new GetInspectionPredefinedControlsArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/about-custom-controls)
        /// * [API documentation](https://help.zscaler.com/zpa/configuring-appprotection-controls-using-api)
        /// 
        /// Use the **zpa_inspection_predefined_controls** data source to get information about an OWASP predefined control and prefedined control version. This data source is required when creating an inspection profile.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zpa.GetInspectionPredefinedControls.Invoke(new()
        ///     {
        ///         Name = "Failed to parse request body",
        ///         Version = "OWASP_CRS/3.3.0",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["zpaInspectionPredefinedControls"] = example,
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetInspectionPredefinedControlsResult> Invoke(GetInspectionPredefinedControlsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInspectionPredefinedControlsResult>("zpa:index/getInspectionPredefinedControls:getInspectionPredefinedControls", args ?? new GetInspectionPredefinedControlsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInspectionPredefinedControlsArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public string? Id { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("version")]
        public string? Version { get; set; }

        public GetInspectionPredefinedControlsArgs()
        {
        }
        public static new GetInspectionPredefinedControlsArgs Empty => new GetInspectionPredefinedControlsArgs();
    }

    public sealed class GetInspectionPredefinedControlsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        public GetInspectionPredefinedControlsInvokeArgs()
        {
        }
        public static new GetInspectionPredefinedControlsInvokeArgs Empty => new GetInspectionPredefinedControlsInvokeArgs();
    }


    [OutputType]
    public sealed class GetInspectionPredefinedControlsResult
    {
        public readonly string Action;
        public readonly string ActionValue;
        public readonly ImmutableArray<Outputs.GetInspectionPredefinedControlsAssociatedInspectionProfileNameResult> AssociatedInspectionProfileNames;
        public readonly string Attachment;
        public readonly string ControlGroup;
        public readonly string ControlNumber;
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
        public readonly string Severity;
        public readonly string? Version;

        [OutputConstructor]
        private GetInspectionPredefinedControlsResult(
            string action,

            string actionValue,

            ImmutableArray<Outputs.GetInspectionPredefinedControlsAssociatedInspectionProfileNameResult> associatedInspectionProfileNames,

            string attachment,

            string controlGroup,

            string controlNumber,

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

            string severity,

            string? version)
        {
            Action = action;
            ActionValue = actionValue;
            AssociatedInspectionProfileNames = associatedInspectionProfileNames;
            Attachment = attachment;
            ControlGroup = controlGroup;
            ControlNumber = controlNumber;
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
            Severity = severity;
            Version = version;
        }
    }
}
