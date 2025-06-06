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
    public sealed class PolicyAccessCapabilitiesRuleCondition
    {
        public readonly string? Id;
        /// <summary>
        /// This signifies the various policy criteria.
        /// </summary>
        public readonly ImmutableArray<Outputs.PolicyAccessCapabilitiesRuleConditionOperand> Operands;
        public readonly string? Operator;

        [OutputConstructor]
        private PolicyAccessCapabilitiesRuleCondition(
            string? id,

            ImmutableArray<Outputs.PolicyAccessCapabilitiesRuleConditionOperand> operands,

            string? @operator)
        {
            Id = id;
            Operands = operands;
            Operator = @operator;
        }
    }
}
