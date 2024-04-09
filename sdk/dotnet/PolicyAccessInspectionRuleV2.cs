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
    /// * [Official documentation](https://help.zscaler.com/zpa/about-security-policy)
    /// * [API documentation](https://help.zscaler.com/zpa/configuring-appprotection-policies-using-api)
    /// 
    /// The **zpa_policy_inspection_rule_v2** resource creates and manages policy access inspection rule in the Zscaler Private Access cloud using a new v2 API endpoint.
    /// 
    ///   ⚠️ **NOTE**: This resource is recommended if your configuration requires the association of more than 1000 resource criteria per rule.
    /// 
    ///   ⚠️ **WARNING:**: The attribute ``rule_order`` is now deprecated in favor of the new resource  ``policy_access_rule_reorder``
    /// 
    /// ## Import
    /// 
    /// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
    /// 
    /// Visit
    /// 
    /// Policy access inspection rule can be imported by using `&lt;RULE ID&gt;` as the import ID.
    /// 
    /// For example:
    /// 
    /// ```sh
    /// $ pulumi import zpa:index/policyAccessInspectionRuleV2:PolicyAccessInspectionRuleV2 example &lt;rule_id&gt;
    /// ```
    /// </summary>
    [ZpaResourceType("zpa:index/policyAccessInspectionRuleV2:PolicyAccessInspectionRuleV2")]
    public partial class PolicyAccessInspectionRuleV2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// This is for providing the rule action. Supported values: `INSPECT` and `BYPASS_INSPECT`.
        /// </summary>
        [Output("action")]
        public Output<string?> Action { get; private set; } = null!;

        /// <summary>
        /// This is for proviidng the set of conditions for the policy.
        /// </summary>
        [Output("conditions")]
        public Output<ImmutableArray<Outputs.PolicyAccessInspectionRuleV2Condition>> Conditions { get; private set; } = null!;

        /// <summary>
        /// This is the description of the access policy rule.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("microtenantId")]
        public Output<string> MicrotenantId { get; private set; } = null!;

        /// <summary>
        /// - (String) This is the name of the policy rule.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("policySetId")]
        public Output<string> PolicySetId { get; private set; } = null!;

        /// <summary>
        /// An inspection profile is required if the `action` is set to `INSPECT`
        /// </summary>
        [Output("zpnInspectionProfileId")]
        public Output<string> ZpnInspectionProfileId { get; private set; } = null!;


        /// <summary>
        /// Create a PolicyAccessInspectionRuleV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyAccessInspectionRuleV2(string name, PolicyAccessInspectionRuleV2Args? args = null, CustomResourceOptions? options = null)
            : base("zpa:index/policyAccessInspectionRuleV2:PolicyAccessInspectionRuleV2", name, args ?? new PolicyAccessInspectionRuleV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private PolicyAccessInspectionRuleV2(string name, Input<string> id, PolicyAccessInspectionRuleV2State? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/policyAccessInspectionRuleV2:PolicyAccessInspectionRuleV2", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/zscaler",
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "zpa:index/policyInspectionRuleV2:PolicyInspectionRuleV2" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PolicyAccessInspectionRuleV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyAccessInspectionRuleV2 Get(string name, Input<string> id, PolicyAccessInspectionRuleV2State? state = null, CustomResourceOptions? options = null)
        {
            return new PolicyAccessInspectionRuleV2(name, id, state, options);
        }
    }

    public sealed class PolicyAccessInspectionRuleV2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is for providing the rule action. Supported values: `INSPECT` and `BYPASS_INSPECT`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("conditions")]
        private InputList<Inputs.PolicyAccessInspectionRuleV2ConditionArgs>? _conditions;

        /// <summary>
        /// This is for proviidng the set of conditions for the policy.
        /// </summary>
        public InputList<Inputs.PolicyAccessInspectionRuleV2ConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.PolicyAccessInspectionRuleV2ConditionArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// This is the description of the access policy rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// - (String) This is the name of the policy rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// An inspection profile is required if the `action` is set to `INSPECT`
        /// </summary>
        [Input("zpnInspectionProfileId")]
        public Input<string>? ZpnInspectionProfileId { get; set; }

        public PolicyAccessInspectionRuleV2Args()
        {
        }
        public static new PolicyAccessInspectionRuleV2Args Empty => new PolicyAccessInspectionRuleV2Args();
    }

    public sealed class PolicyAccessInspectionRuleV2State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is for providing the rule action. Supported values: `INSPECT` and `BYPASS_INSPECT`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("conditions")]
        private InputList<Inputs.PolicyAccessInspectionRuleV2ConditionGetArgs>? _conditions;

        /// <summary>
        /// This is for proviidng the set of conditions for the policy.
        /// </summary>
        public InputList<Inputs.PolicyAccessInspectionRuleV2ConditionGetArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.PolicyAccessInspectionRuleV2ConditionGetArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// This is the description of the access policy rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// - (String) This is the name of the policy rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policySetId")]
        public Input<string>? PolicySetId { get; set; }

        /// <summary>
        /// An inspection profile is required if the `action` is set to `INSPECT`
        /// </summary>
        [Input("zpnInspectionProfileId")]
        public Input<string>? ZpnInspectionProfileId { get; set; }

        public PolicyAccessInspectionRuleV2State()
        {
        }
        public static new PolicyAccessInspectionRuleV2State Empty => new PolicyAccessInspectionRuleV2State();
    }
}