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
    /// <summary>
    /// * [Official documentation](https://help.zscaler.com/zpa/configuring-app-connectors-settings)
    /// * [API documentation](https://help.zscaler.com/zpa/configuring-auto-delete-disconnected-app-connectors-using-api)
    /// 
    /// Use the **zpa_app_connector_assistant_schedule** resource sets the scheduled frequency at which the disconnected App Connectors are eligible for deletion. The supported value for frequency is days. The frequencyInterval field is the number of days after an App Connector disconnects for it to become eligible for deletion. The minimum supported value for frequencyInterval is 5.
    /// 
    /// &gt; **NOTE** - When enabling the Assistant Schedule for the first time, you must provide the `customer_id` information. If you authenticated using environment variables and used `ZPA_CUSTOMER_ID` environment variable, you don't have to define the customer_id attribute in the HCL configuration, and the provider will automatically use the value from the environment variable `ZPA_CUSTOMER_ID`
    /// 
    /// ## Example Usage
    /// 
    /// ### Defined Customer ID Value
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Zpa = Zscaler.Zpa;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @this = new Zpa.AppConnectorAssistantSchedule("this", new()
    ///     {
    ///         CustomerId = "123456789101112",
    ///         DeleteDisabled = true,
    ///         Enabled = true,
    ///         Frequency = "days",
    ///         FrequencyInterval = "5",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Customer ID Via Environment Variable
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Zpa = Zscaler.Zpa;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @this = new Zpa.AppConnectorAssistantSchedule("this", new()
    ///     {
    ///         DeleteDisabled = true,
    ///         Enabled = true,
    ///         Frequency = "days",
    ///         FrequencyInterval = "5",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Import is not currently supported for this resource.
    /// </summary>
    [Obsolete(@"zpa.index/assistantschedule.AssistantSchedule has been deprecated in favor of zpa.index/appconnectorassistantschedule.AppConnectorAssistantSchedule")]
    [ZpaResourceType("zpa:index/assistantSchedule:AssistantSchedule")]
    public partial class AssistantSchedule : global::Pulumi.CustomResource
    {
        [Output("customerId")]
        public Output<string> CustomerId { get; private set; } = null!;

        [Output("deleteDisabled")]
        public Output<bool?> DeleteDisabled { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        [Output("frequency")]
        public Output<string?> Frequency { get; private set; } = null!;

        [Output("frequencyInterval")]
        public Output<string?> FrequencyInterval { get; private set; } = null!;


        /// <summary>
        /// Create a AssistantSchedule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AssistantSchedule(string name, AssistantScheduleArgs? args = null, CustomResourceOptions? options = null)
            : base("zpa:index/assistantSchedule:AssistantSchedule", name, args ?? new AssistantScheduleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AssistantSchedule(string name, Input<string> id, AssistantScheduleState? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/assistantSchedule:AssistantSchedule", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/zscaler",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AssistantSchedule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AssistantSchedule Get(string name, Input<string> id, AssistantScheduleState? state = null, CustomResourceOptions? options = null)
        {
            return new AssistantSchedule(name, id, state, options);
        }
    }

    public sealed class AssistantScheduleArgs : global::Pulumi.ResourceArgs
    {
        [Input("customerId")]
        public Input<string>? CustomerId { get; set; }

        [Input("deleteDisabled")]
        public Input<bool>? DeleteDisabled { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("frequency")]
        public Input<string>? Frequency { get; set; }

        [Input("frequencyInterval")]
        public Input<string>? FrequencyInterval { get; set; }

        public AssistantScheduleArgs()
        {
        }
        public static new AssistantScheduleArgs Empty => new AssistantScheduleArgs();
    }

    public sealed class AssistantScheduleState : global::Pulumi.ResourceArgs
    {
        [Input("customerId")]
        public Input<string>? CustomerId { get; set; }

        [Input("deleteDisabled")]
        public Input<bool>? DeleteDisabled { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("frequency")]
        public Input<string>? Frequency { get; set; }

        [Input("frequencyInterval")]
        public Input<string>? FrequencyInterval { get; set; }

        public AssistantScheduleState()
        {
        }
        public static new AssistantScheduleState Empty => new AssistantScheduleState();
    }
}
