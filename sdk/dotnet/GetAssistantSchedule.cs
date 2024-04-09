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
    public static class GetAssistantSchedule
    {
        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/deleting-disconnected-app-connectors)
        /// * [API documentation](https://help.zscaler.com/zpa/configuring-auto-delete-disconnected-app-connectors-using-api)
        /// 
        /// Use the **zpa_app_connector_assistant_schedule** data source to get information about Auto Delete frequency of the App Connector for the specified customer in the Zscaler Private Access cloud.
        /// 
        /// &gt; **NOTE** - The `customer_id` attribute is optional and not required during the configuration.
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
        ///     var @this = Zpa.GetAssistantSchedule.Invoke(new()
        ///     {
        ///         CustomerId = "1234567891012",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetAssistantScheduleResult> InvokeAsync(GetAssistantScheduleArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAssistantScheduleResult>("zpa:index/getAssistantSchedule:getAssistantSchedule", args ?? new GetAssistantScheduleArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/deleting-disconnected-app-connectors)
        /// * [API documentation](https://help.zscaler.com/zpa/configuring-auto-delete-disconnected-app-connectors-using-api)
        /// 
        /// Use the **zpa_app_connector_assistant_schedule** data source to get information about Auto Delete frequency of the App Connector for the specified customer in the Zscaler Private Access cloud.
        /// 
        /// &gt; **NOTE** - The `customer_id` attribute is optional and not required during the configuration.
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
        ///     var @this = Zpa.GetAssistantSchedule.Invoke(new()
        ///     {
        ///         CustomerId = "1234567891012",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetAssistantScheduleResult> Invoke(GetAssistantScheduleInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAssistantScheduleResult>("zpa:index/getAssistantSchedule:getAssistantSchedule", args ?? new GetAssistantScheduleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAssistantScheduleArgs : global::Pulumi.InvokeArgs
    {
        [Input("customerId")]
        public string? CustomerId { get; set; }

        [Input("id")]
        public string? Id { get; set; }

        public GetAssistantScheduleArgs()
        {
        }
        public static new GetAssistantScheduleArgs Empty => new GetAssistantScheduleArgs();
    }

    public sealed class GetAssistantScheduleInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("customerId")]
        public Input<string>? CustomerId { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        public GetAssistantScheduleInvokeArgs()
        {
        }
        public static new GetAssistantScheduleInvokeArgs Empty => new GetAssistantScheduleInvokeArgs();
    }


    [OutputType]
    public sealed class GetAssistantScheduleResult
    {
        public readonly string? CustomerId;
        public readonly bool DeleteDisabled;
        public readonly bool Enabled;
        public readonly string Frequency;
        public readonly string FrequencyInterval;
        public readonly string? Id;

        [OutputConstructor]
        private GetAssistantScheduleResult(
            string? customerId,

            bool deleteDisabled,

            bool enabled,

            string frequency,

            string frequencyInterval,

            string? id)
        {
            CustomerId = customerId;
            DeleteDisabled = deleteDisabled;
            Enabled = enabled;
            Frequency = frequency;
            FrequencyInterval = frequencyInterval;
            Id = id;
        }
    }
}
