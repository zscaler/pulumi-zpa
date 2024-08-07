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
    [Obsolete(@"zpa.index/policycredentialrule.PolicyCredentialRule has been deprecated in favor of zpa.index/policyaccesscredentialrule.PolicyAccessCredentialRule")]
    [ZpaResourceType("zpa:index/policyCredentialRule:PolicyCredentialRule")]
    public partial class PolicyCredentialRule : global::Pulumi.CustomResource
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
        public Output<ImmutableArray<Outputs.PolicyCredentialRuleCondition>> Conditions { get; private set; } = null!;

        [Output("credentials")]
        public Output<ImmutableArray<Outputs.PolicyCredentialRuleCredential>> Credentials { get; private set; } = null!;

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


        /// <summary>
        /// Create a PolicyCredentialRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyCredentialRule(string name, PolicyCredentialRuleArgs? args = null, CustomResourceOptions? options = null)
            : base("zpa:index/policyCredentialRule:PolicyCredentialRule", name, args ?? new PolicyCredentialRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PolicyCredentialRule(string name, Input<string> id, PolicyCredentialRuleState? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/policyCredentialRule:PolicyCredentialRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PolicyCredentialRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyCredentialRule Get(string name, Input<string> id, PolicyCredentialRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new PolicyCredentialRule(name, id, state, options);
        }
    }

    public sealed class PolicyCredentialRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is for providing the rule action.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("conditions")]
        private InputList<Inputs.PolicyCredentialRuleConditionArgs>? _conditions;

        /// <summary>
        /// This is for proviidng the set of conditions for the policy.
        /// </summary>
        public InputList<Inputs.PolicyCredentialRuleConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.PolicyCredentialRuleConditionArgs>());
            set => _conditions = value;
        }

        [Input("credentials")]
        private InputList<Inputs.PolicyCredentialRuleCredentialArgs>? _credentials;
        public InputList<Inputs.PolicyCredentialRuleCredentialArgs> Credentials
        {
            get => _credentials ?? (_credentials = new InputList<Inputs.PolicyCredentialRuleCredentialArgs>());
            set => _credentials = value;
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

        public PolicyCredentialRuleArgs()
        {
        }
        public static new PolicyCredentialRuleArgs Empty => new PolicyCredentialRuleArgs();
    }

    public sealed class PolicyCredentialRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is for providing the rule action.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("conditions")]
        private InputList<Inputs.PolicyCredentialRuleConditionGetArgs>? _conditions;

        /// <summary>
        /// This is for proviidng the set of conditions for the policy.
        /// </summary>
        public InputList<Inputs.PolicyCredentialRuleConditionGetArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.PolicyCredentialRuleConditionGetArgs>());
            set => _conditions = value;
        }

        [Input("credentials")]
        private InputList<Inputs.PolicyCredentialRuleCredentialGetArgs>? _credentials;
        public InputList<Inputs.PolicyCredentialRuleCredentialGetArgs> Credentials
        {
            get => _credentials ?? (_credentials = new InputList<Inputs.PolicyCredentialRuleCredentialGetArgs>());
            set => _credentials = value;
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

        public PolicyCredentialRuleState()
        {
        }
        public static new PolicyCredentialRuleState Empty => new PolicyCredentialRuleState();
    }
}
