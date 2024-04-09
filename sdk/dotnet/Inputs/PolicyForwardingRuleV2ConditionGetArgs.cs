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

    public sealed class PolicyForwardingRuleV2ConditionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("operands")]
        private InputList<Inputs.PolicyForwardingRuleV2ConditionOperandGetArgs>? _operands;

        /// <summary>
        /// This signifies the various policy criteria.
        /// </summary>
        public InputList<Inputs.PolicyForwardingRuleV2ConditionOperandGetArgs> Operands
        {
            get => _operands ?? (_operands = new InputList<Inputs.PolicyForwardingRuleV2ConditionOperandGetArgs>());
            set => _operands = value;
        }

        [Input("operator")]
        public Input<string>? Operator { get; set; }

        public PolicyForwardingRuleV2ConditionGetArgs()
        {
        }
        public static new PolicyForwardingRuleV2ConditionGetArgs Empty => new PolicyForwardingRuleV2ConditionGetArgs();
    }
}
