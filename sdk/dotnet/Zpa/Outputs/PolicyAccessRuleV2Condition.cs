// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.Outputs
{

    [OutputType]
    public sealed class PolicyAccessRuleV2Condition
    {
        public readonly string? Id;
        /// <summary>
        /// This signifies the various policy criteria.
        /// </summary>
        public readonly ImmutableArray<Outputs.PolicyAccessRuleV2ConditionOperand> Operands;
        public readonly string? Operator;

        [OutputConstructor]
        private PolicyAccessRuleV2Condition(
            string? id,

            ImmutableArray<Outputs.PolicyAccessRuleV2ConditionOperand> operands,

            string? @operator)
        {
            Id = id;
            Operands = operands;
            Operator = @operator;
        }
    }
}
