// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Zscaler.Zpa.Outputs
{

    [OutputType]
    public sealed class PolicyForwardingRuleV2Condition
    {
        public readonly string? Id;
        /// <summary>
        /// This signifies the various policy criteria.
        /// </summary>
        public readonly ImmutableArray<Outputs.PolicyForwardingRuleV2ConditionOperand> Operands;
        public readonly string? Operator;

        [OutputConstructor]
        private PolicyForwardingRuleV2Condition(
            string? id,

            ImmutableArray<Outputs.PolicyForwardingRuleV2ConditionOperand> operands,

            string? @operator)
        {
            Id = id;
            Operands = operands;
            Operator = @operator;
        }
    }
}
