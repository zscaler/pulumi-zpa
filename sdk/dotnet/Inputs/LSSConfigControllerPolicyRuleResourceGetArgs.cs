// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Zscaler.Zpa.Inputs
{

    public sealed class LSSConfigControllerPolicyRuleResourceGetArgs : global::Pulumi.ResourceArgs
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
        private InputList<Inputs.LSSConfigControllerPolicyRuleResourceConditionGetArgs>? _conditions;

        /// <summary>
        /// This is for proviidng the set of conditions for the policy.
        /// </summary>
        public InputList<Inputs.LSSConfigControllerPolicyRuleResourceConditionGetArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.LSSConfigControllerPolicyRuleResourceConditionGetArgs>());
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

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("lssDefaultRule")]
        public Input<bool>? LssDefaultRule { get; set; }

        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// This is the name of the policy.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

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

        public LSSConfigControllerPolicyRuleResourceGetArgs()
        {
        }
        public static new LSSConfigControllerPolicyRuleResourceGetArgs Empty => new LSSConfigControllerPolicyRuleResourceGetArgs();
    }
}