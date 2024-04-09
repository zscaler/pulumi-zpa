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
    [Obsolete(@"zpa.index/policyisolationrulev2.PolicyIsolationRuleV2 has been deprecated in favor of zpa.index/policyaccessisolationrulev2.PolicyAccessIsolationRuleV2")]
    [ZpaResourceType("zpa:index/policyIsolationRuleV2:PolicyIsolationRuleV2")]
    public partial class PolicyIsolationRuleV2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// This is for providing the rule action.
        /// </summary>
        [Output("action")]
        public Output<string?> Action { get; private set; } = null!;

        /// <summary>
        /// This is for proviidng the set of conditions for the policy.
        /// </summary>
        [Output("conditions")]
        public Output<ImmutableArray<Outputs.PolicyIsolationRuleV2Condition>> Conditions { get; private set; } = null!;

        /// <summary>
        /// This is the description of the access policy.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("microtenantId")]
        public Output<string> MicrotenantId { get; private set; } = null!;

        /// <summary>
        /// This is the name of the policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("policySetId")]
        public Output<string> PolicySetId { get; private set; } = null!;

        [Output("zpnIsolationProfileId")]
        public Output<string> ZpnIsolationProfileId { get; private set; } = null!;


        /// <summary>
        /// Create a PolicyIsolationRuleV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyIsolationRuleV2(string name, PolicyIsolationRuleV2Args? args = null, CustomResourceOptions? options = null)
            : base("zpa:index/policyIsolationRuleV2:PolicyIsolationRuleV2", name, args ?? new PolicyIsolationRuleV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private PolicyIsolationRuleV2(string name, Input<string> id, PolicyIsolationRuleV2State? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/policyIsolationRuleV2:PolicyIsolationRuleV2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PolicyIsolationRuleV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyIsolationRuleV2 Get(string name, Input<string> id, PolicyIsolationRuleV2State? state = null, CustomResourceOptions? options = null)
        {
            return new PolicyIsolationRuleV2(name, id, state, options);
        }
    }

    public sealed class PolicyIsolationRuleV2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is for providing the rule action.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("conditions")]
        private InputList<Inputs.PolicyIsolationRuleV2ConditionArgs>? _conditions;

        /// <summary>
        /// This is for proviidng the set of conditions for the policy.
        /// </summary>
        public InputList<Inputs.PolicyIsolationRuleV2ConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.PolicyIsolationRuleV2ConditionArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// This is the description of the access policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// This is the name of the policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("zpnIsolationProfileId")]
        public Input<string>? ZpnIsolationProfileId { get; set; }

        public PolicyIsolationRuleV2Args()
        {
        }
        public static new PolicyIsolationRuleV2Args Empty => new PolicyIsolationRuleV2Args();
    }

    public sealed class PolicyIsolationRuleV2State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is for providing the rule action.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("conditions")]
        private InputList<Inputs.PolicyIsolationRuleV2ConditionGetArgs>? _conditions;

        /// <summary>
        /// This is for proviidng the set of conditions for the policy.
        /// </summary>
        public InputList<Inputs.PolicyIsolationRuleV2ConditionGetArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.PolicyIsolationRuleV2ConditionGetArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// This is the description of the access policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// This is the name of the policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policySetId")]
        public Input<string>? PolicySetId { get; set; }

        [Input("zpnIsolationProfileId")]
        public Input<string>? ZpnIsolationProfileId { get; set; }

        public PolicyIsolationRuleV2State()
        {
        }
        public static new PolicyIsolationRuleV2State Empty => new PolicyIsolationRuleV2State();
    }
}
