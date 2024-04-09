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

    public sealed class PolicyAccessForwardingRuleV2ConditionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("operands")]
        private InputList<Inputs.PolicyAccessForwardingRuleV2ConditionOperandGetArgs>? _operands;

        /// <summary>
        /// This signifies the various policy criteria.
        /// </summary>
        public InputList<Inputs.PolicyAccessForwardingRuleV2ConditionOperandGetArgs> Operands
        {
            get => _operands ?? (_operands = new InputList<Inputs.PolicyAccessForwardingRuleV2ConditionOperandGetArgs>());
            set => _operands = value;
        }

        [Input("operator")]
        public Input<string>? Operator { get; set; }

        public PolicyAccessForwardingRuleV2ConditionGetArgs()
        {
        }
        public static new PolicyAccessForwardingRuleV2ConditionGetArgs Empty => new PolicyAccessForwardingRuleV2ConditionGetArgs();
    }
}