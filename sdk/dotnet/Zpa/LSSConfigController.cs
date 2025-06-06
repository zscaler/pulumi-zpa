// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa
{
    /// <summary>
    /// * [Official documentation](https://help.zscaler.com/zpa/about-log-streaming-service)
    /// * [API documentation](https://help.zscaler.com/zpa/configuring-log-streaming-service-configurations-using-api)
    /// 
    /// The **zpa_lss_config_controller** resource creates and manages Log Streaming Service (LSS) in the Zscaler Private Access cloud for App Connector Metrics `zpn_ast_comprehensive_stats`.
    /// 
    /// ## Example 1 - LSS App Connector Metrics - Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Zpa = Pulumi.Zpa;
    /// using Zpa = zscaler.PulumiPackage.Zpa;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var zpnAstComprehensiveStats = Zpa.GetLSSLogTypeFormats.Invoke(new()
    ///     {
    ///         LogType = "zpn_ast_comprehensive_stats",
    ///     });
    /// 
    ///     var lssSiemPolicy = Zpa.GetPolicyType.Invoke(new()
    ///     {
    ///         PolicyType = "SIEM_POLICY",
    ///     });
    /// 
    ///     var @this = Zpa.GetAppConnectorGroup.Invoke(new()
    ///     {
    ///         Name = "Example100",
    ///     });
    /// 
    ///     var lssAppConnectorMetrics = new Zpa.LSSConfigController("lssAppConnectorMetrics", new()
    ///     {
    ///         Config = new Zpa.Inputs.LSSConfigControllerConfigArgs
    ///         {
    ///             Name = "LSS App Connector Metrics",
    ///             Description = "LSS App Connector Metrics",
    ///             Enabled = true,
    ///             Format = zpnAstComprehensiveStats.Apply(getLSSLogTypeFormatsResult =&gt; getLSSLogTypeFormatsResult.Json),
    ///             LssHost = "splunk1.acme.com",
    ///             LssPort = "5001",
    ///             SourceLogType = "zpn_ast_comprehensive_stats",
    ///             UseTls = true,
    ///         },
    ///         ConnectorGroups = new[]
    ///         {
    ///             new Zpa.Inputs.LSSConfigControllerConnectorGroupArgs
    ///             {
    ///                 Ids = new[]
    ///                 {
    ///                     @this.Apply(@this =&gt; @this.Apply(getAppConnectorGroupResult =&gt; getAppConnectorGroupResult.Id)),
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## LSS Source Log Type Table
    /// 
    /// |       Source Log Type                     |            Description                 |
    /// |-------------------------------------------|----------------------------------------|
    /// |        `zpn_trans_log`                    |        `User Activity`                 |
    /// |        `zpn_auth_log`                     |         `User Status`                  |
    /// |        `zpn_ast_auth_log`                 |        `App Connector Status`          |
    /// |        `zpn_http_trans_log`               |         `Web Browser`                  |
    /// |        `zpn_audit_log`                    |         `Audit Logs`                   |
    /// |        `zpn_sys_auth_log`                 |         `Private Service Edge Status`  |
    /// |        `zpn_ast_comprehensive_stats`      |         `App Connector Metrics`        |
    /// |        `zpn_pbroker_comprehensive_stats`  |         `Private Service Edge Metrics` |
    /// |        `zpn_waf_http_exchanges_log`       |         `ZPA App Protection`           |
    /// 
    /// ## Import
    /// 
    /// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
    /// 
    /// Visit
    /// </summary>
    [ZpaResourceType("zpa:index/lSSConfigController:LSSConfigController")]
    public partial class LSSConfigController : global::Pulumi.CustomResource
    {
        [Output("config")]
        public Output<Outputs.LSSConfigControllerConfig?> Config { get; private set; } = null!;

        /// <summary>
        /// App Connector Group(s) to be added to the LSS configuration
        /// </summary>
        [Output("connectorGroups")]
        public Output<ImmutableArray<Outputs.LSSConfigControllerConnectorGroup>> ConnectorGroups { get; private set; } = null!;

        [Output("policyRuleId")]
        public Output<string> PolicyRuleId { get; private set; } = null!;

        [Output("policyRuleResource")]
        public Output<Outputs.LSSConfigControllerPolicyRuleResource?> PolicyRuleResource { get; private set; } = null!;


        /// <summary>
        /// Create a LSSConfigController resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LSSConfigController(string name, LSSConfigControllerArgs? args = null, CustomResourceOptions? options = null)
            : base("zpa:index/lSSConfigController:LSSConfigController", name, args ?? new LSSConfigControllerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LSSConfigController(string name, Input<string> id, LSSConfigControllerState? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/lSSConfigController:LSSConfigController", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LSSConfigController resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LSSConfigController Get(string name, Input<string> id, LSSConfigControllerState? state = null, CustomResourceOptions? options = null)
        {
            return new LSSConfigController(name, id, state, options);
        }
    }

    public sealed class LSSConfigControllerArgs : global::Pulumi.ResourceArgs
    {
        [Input("config")]
        public Input<Inputs.LSSConfigControllerConfigArgs>? Config { get; set; }

        [Input("connectorGroups")]
        private InputList<Inputs.LSSConfigControllerConnectorGroupArgs>? _connectorGroups;

        /// <summary>
        /// App Connector Group(s) to be added to the LSS configuration
        /// </summary>
        public InputList<Inputs.LSSConfigControllerConnectorGroupArgs> ConnectorGroups
        {
            get => _connectorGroups ?? (_connectorGroups = new InputList<Inputs.LSSConfigControllerConnectorGroupArgs>());
            set => _connectorGroups = value;
        }

        [Input("policyRuleResource")]
        public Input<Inputs.LSSConfigControllerPolicyRuleResourceArgs>? PolicyRuleResource { get; set; }

        public LSSConfigControllerArgs()
        {
        }
        public static new LSSConfigControllerArgs Empty => new LSSConfigControllerArgs();
    }

    public sealed class LSSConfigControllerState : global::Pulumi.ResourceArgs
    {
        [Input("config")]
        public Input<Inputs.LSSConfigControllerConfigGetArgs>? Config { get; set; }

        [Input("connectorGroups")]
        private InputList<Inputs.LSSConfigControllerConnectorGroupGetArgs>? _connectorGroups;

        /// <summary>
        /// App Connector Group(s) to be added to the LSS configuration
        /// </summary>
        public InputList<Inputs.LSSConfigControllerConnectorGroupGetArgs> ConnectorGroups
        {
            get => _connectorGroups ?? (_connectorGroups = new InputList<Inputs.LSSConfigControllerConnectorGroupGetArgs>());
            set => _connectorGroups = value;
        }

        [Input("policyRuleId")]
        public Input<string>? PolicyRuleId { get; set; }

        [Input("policyRuleResource")]
        public Input<Inputs.LSSConfigControllerPolicyRuleResourceGetArgs>? PolicyRuleResource { get; set; }

        public LSSConfigControllerState()
        {
        }
        public static new LSSConfigControllerState Empty => new LSSConfigControllerState();
    }
}
