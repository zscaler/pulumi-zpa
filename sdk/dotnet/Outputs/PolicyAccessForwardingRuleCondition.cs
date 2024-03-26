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
    public sealed class PolicyAccessForwardingRuleCondition
    {
        public readonly string? Id;
        public readonly string? MicrotenantId;
        public readonly bool? Negated;
        /// <summary>
        /// This signifies the various policy criteria.
        /// </summary>
        public readonly ImmutableArray<Outputs.PolicyAccessForwardingRuleConditionOperand> Operands;
        public readonly string Operator;

        [OutputConstructor]
        private PolicyAccessForwardingRuleCondition(
            string? id,

            string? microtenantId,

            bool? negated,

            ImmutableArray<Outputs.PolicyAccessForwardingRuleConditionOperand> operands,

            string @operator)
        {
            Id = id;
            MicrotenantId = microtenantId;
            Negated = negated;
            Operands = operands;
            Operator = @operator;
        }
    }
}
