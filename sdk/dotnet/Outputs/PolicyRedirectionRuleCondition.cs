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
    public sealed class PolicyRedirectionRuleCondition
    {
        public readonly string? Id;
        public readonly string? MicrotenantId;
        /// <summary>
        /// This signifies the various policy criteria.
        /// </summary>
        public readonly ImmutableArray<Outputs.PolicyRedirectionRuleConditionOperand> Operands;
        public readonly string Operator;

        [OutputConstructor]
        private PolicyRedirectionRuleCondition(
            string? id,

            string? microtenantId,

            ImmutableArray<Outputs.PolicyRedirectionRuleConditionOperand> operands,

            string @operator)
        {
            Id = id;
            MicrotenantId = microtenantId;
            Operands = operands;
            Operator = @operator;
        }
    }
}