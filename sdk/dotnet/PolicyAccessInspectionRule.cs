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
    [ZpaResourceType("zpa:index/policyAccessInspectionRule:PolicyAccessInspectionRule")]
    public partial class PolicyAccessInspectionRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// This is for providing the rule action.
        /// </summary>
        [Output("action")]
        public Output<string?> Action { get; private set; } = null!;

        /// <summary>
        /// This field defines the description of the server.
        /// </summary>
        [Output("actionId")]
        public Output<string?> ActionId { get; private set; } = null!;

        [Output("bypassDefaultRule")]
        public Output<bool?> BypassDefaultRule { get; private set; } = null!;

        /// <summary>
        /// This is for proviidng the set of conditions for the policy.
        /// </summary>
        [Output("conditions")]
        public Output<ImmutableArray<Outputs.PolicyAccessInspectionRuleCondition>> Conditions { get; private set; } = null!;

        /// <summary>
        /// This is for providing a customer message for the user.
        /// </summary>
        [Output("customMsg")]
        public Output<string?> CustomMsg { get; private set; } = null!;

        /// <summary>
        /// This is for providing a customer message for the user.
        /// </summary>
        [Output("defaultRule")]
        public Output<bool?> DefaultRule { get; private set; } = null!;

        /// <summary>
        /// This is the description of the access policy.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("lssDefaultRule")]
        public Output<bool?> LssDefaultRule { get; private set; } = null!;

        [Output("microtenantId")]
        public Output<string> MicrotenantId { get; private set; } = null!;

        /// <summary>
        /// This is the name of the policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("operator")]
        public Output<string> Operator { get; private set; } = null!;

        [Output("policySetId")]
        public Output<string> PolicySetId { get; private set; } = null!;

        [Output("policyType")]
        public Output<string> PolicyType { get; private set; } = null!;

        [Output("priority")]
        public Output<string> Priority { get; private set; } = null!;

        [Output("reauthDefaultRule")]
        public Output<bool?> ReauthDefaultRule { get; private set; } = null!;

        [Output("reauthIdleTimeout")]
        public Output<string?> ReauthIdleTimeout { get; private set; } = null!;

        [Output("reauthTimeout")]
        public Output<string?> ReauthTimeout { get; private set; } = null!;

        [Output("ruleOrder")]
        public Output<string> RuleOrder { get; private set; } = null!;

        [Output("zpnCbiProfileId")]
        public Output<string> ZpnCbiProfileId { get; private set; } = null!;

        [Output("zpnInspectionProfileId")]
        public Output<string> ZpnInspectionProfileId { get; private set; } = null!;

        [Output("zpnIsolationProfileId")]
        public Output<string> ZpnIsolationProfileId { get; private set; } = null!;


        /// <summary>
        /// Create a PolicyAccessInspectionRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyAccessInspectionRule(string name, PolicyAccessInspectionRuleArgs? args = null, CustomResourceOptions? options = null)
            : base("zpa:index/policyAccessInspectionRule:PolicyAccessInspectionRule", name, args ?? new PolicyAccessInspectionRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PolicyAccessInspectionRule(string name, Input<string> id, PolicyAccessInspectionRuleState? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/policyAccessInspectionRule:PolicyAccessInspectionRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PolicyAccessInspectionRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyAccessInspectionRule Get(string name, Input<string> id, PolicyAccessInspectionRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new PolicyAccessInspectionRule(name, id, state, options);
        }
    }

    public sealed class PolicyAccessInspectionRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is for providing the rule action.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// This field defines the description of the server.
        /// </summary>
        [Input("actionId")]
        public Input<string>? ActionId { get; set; }

        [Input("bypassDefaultRule")]
        public Input<bool>? BypassDefaultRule { get; set; }

        [Input("conditions")]
        private InputList<Inputs.PolicyAccessInspectionRuleConditionArgs>? _conditions;

        /// <summary>
        /// This is for proviidng the set of conditions for the policy.
        /// </summary>
        public InputList<Inputs.PolicyAccessInspectionRuleConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.PolicyAccessInspectionRuleConditionArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// This is for providing a customer message for the user.
        /// </summary>
        [Input("customMsg")]
        public Input<string>? CustomMsg { get; set; }

        /// <summary>
        /// This is for providing a customer message for the user.
        /// </summary>
        [Input("defaultRule")]
        public Input<bool>? DefaultRule { get; set; }

        /// <summary>
        /// This is the description of the access policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("lssDefaultRule")]
        public Input<bool>? LssDefaultRule { get; set; }

        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// This is the name of the policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("operator")]
        public Input<string>? Operator { get; set; }

        [Input("policySetId")]
        public Input<string>? PolicySetId { get; set; }

        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        [Input("priority")]
        public Input<string>? Priority { get; set; }

        [Input("reauthDefaultRule")]
        public Input<bool>? ReauthDefaultRule { get; set; }

        [Input("reauthIdleTimeout")]
        public Input<string>? ReauthIdleTimeout { get; set; }

        [Input("reauthTimeout")]
        public Input<string>? ReauthTimeout { get; set; }

        [Input("ruleOrder")]
        public Input<string>? RuleOrder { get; set; }

        [Input("zpnCbiProfileId")]
        public Input<string>? ZpnCbiProfileId { get; set; }

        [Input("zpnInspectionProfileId")]
        public Input<string>? ZpnInspectionProfileId { get; set; }

        [Input("zpnIsolationProfileId")]
        public Input<string>? ZpnIsolationProfileId { get; set; }

        public PolicyAccessInspectionRuleArgs()
        {
        }
        public static new PolicyAccessInspectionRuleArgs Empty => new PolicyAccessInspectionRuleArgs();
    }

    public sealed class PolicyAccessInspectionRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is for providing the rule action.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// This field defines the description of the server.
        /// </summary>
        [Input("actionId")]
        public Input<string>? ActionId { get; set; }

        [Input("bypassDefaultRule")]
        public Input<bool>? BypassDefaultRule { get; set; }

        [Input("conditions")]
        private InputList<Inputs.PolicyAccessInspectionRuleConditionGetArgs>? _conditions;

        /// <summary>
        /// This is for proviidng the set of conditions for the policy.
        /// </summary>
        public InputList<Inputs.PolicyAccessInspectionRuleConditionGetArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.PolicyAccessInspectionRuleConditionGetArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// This is for providing a customer message for the user.
        /// </summary>
        [Input("customMsg")]
        public Input<string>? CustomMsg { get; set; }

        /// <summary>
        /// This is for providing a customer message for the user.
        /// </summary>
        [Input("defaultRule")]
        public Input<bool>? DefaultRule { get; set; }

        /// <summary>
        /// This is the description of the access policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("lssDefaultRule")]
        public Input<bool>? LssDefaultRule { get; set; }

        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// This is the name of the policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("operator")]
        public Input<string>? Operator { get; set; }

        [Input("policySetId")]
        public Input<string>? PolicySetId { get; set; }

        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        [Input("priority")]
        public Input<string>? Priority { get; set; }

        [Input("reauthDefaultRule")]
        public Input<bool>? ReauthDefaultRule { get; set; }

        [Input("reauthIdleTimeout")]
        public Input<string>? ReauthIdleTimeout { get; set; }

        [Input("reauthTimeout")]
        public Input<string>? ReauthTimeout { get; set; }

        [Input("ruleOrder")]
        public Input<string>? RuleOrder { get; set; }

        [Input("zpnCbiProfileId")]
        public Input<string>? ZpnCbiProfileId { get; set; }

        [Input("zpnInspectionProfileId")]
        public Input<string>? ZpnInspectionProfileId { get; set; }

        [Input("zpnIsolationProfileId")]
        public Input<string>? ZpnIsolationProfileId { get; set; }

        public PolicyAccessInspectionRuleState()
        {
        }
        public static new PolicyAccessInspectionRuleState Empty => new PolicyAccessInspectionRuleState();
    }
}
