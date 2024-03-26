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
    public static class GetInspectionAllPredefinedControls
    {
        /// <summary>
        /// Use the **zpa_inspection_all_predefined_controls** data source to get information about all OWASP predefined control and prefedined control version by group name. The `Preprocessors` predefined control is the default predefined control, This data source is always required, when creating an inspection profile.
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
        ///     var @this = Zpa.GetInspectionAllPredefinedControls.Invoke(new()
        ///     {
        ///         GroupName = "Preprocessors",
        ///         Version = "OWASP_CRS/3.3.0",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetInspectionAllPredefinedControlsResult> InvokeAsync(GetInspectionAllPredefinedControlsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInspectionAllPredefinedControlsResult>("zpa:index/getInspectionAllPredefinedControls:getInspectionAllPredefinedControls", args ?? new GetInspectionAllPredefinedControlsArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zpa_inspection_all_predefined_controls** data source to get information about all OWASP predefined control and prefedined control version by group name. The `Preprocessors` predefined control is the default predefined control, This data source is always required, when creating an inspection profile.
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
        ///     var @this = Zpa.GetInspectionAllPredefinedControls.Invoke(new()
        ///     {
        ///         GroupName = "Preprocessors",
        ///         Version = "OWASP_CRS/3.3.0",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetInspectionAllPredefinedControlsResult> Invoke(GetInspectionAllPredefinedControlsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInspectionAllPredefinedControlsResult>("zpa:index/getInspectionAllPredefinedControls:getInspectionAllPredefinedControls", args ?? new GetInspectionAllPredefinedControlsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInspectionAllPredefinedControlsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the predefined control.
        /// </summary>
        [Input("groupName")]
        public string? GroupName { get; set; }

        /// <summary>
        /// The version of the predefined control, the default is: `OWASP_CRS/3.3.0`
        /// </summary>
        [Input("version", required: true)]
        public string Version { get; set; } = null!;

        public GetInspectionAllPredefinedControlsArgs()
        {
        }
        public static new GetInspectionAllPredefinedControlsArgs Empty => new GetInspectionAllPredefinedControlsArgs();
    }

    public sealed class GetInspectionAllPredefinedControlsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the predefined control.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        /// <summary>
        /// The version of the predefined control, the default is: `OWASP_CRS/3.3.0`
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public GetInspectionAllPredefinedControlsInvokeArgs()
        {
        }
        public static new GetInspectionAllPredefinedControlsInvokeArgs Empty => new GetInspectionAllPredefinedControlsInvokeArgs();
    }


    [OutputType]
    public sealed class GetInspectionAllPredefinedControlsResult
    {
        public readonly string? GroupName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetInspectionAllPredefinedControlsListResult> Lists;
        public readonly string Version;

        [OutputConstructor]
        private GetInspectionAllPredefinedControlsResult(
            string? groupName,

            string id,

            ImmutableArray<Outputs.GetInspectionAllPredefinedControlsListResult> lists,

            string version)
        {
            GroupName = groupName;
            Id = id;
            Lists = lists;
            Version = version;
        }
    }
}
