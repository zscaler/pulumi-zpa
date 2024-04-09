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

    public sealed class PolicyAccessIsolationRuleConditionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        [Input("operands")]
        private InputList<Inputs.PolicyAccessIsolationRuleConditionOperandGetArgs>? _operands;

        /// <summary>
        /// This signifies the various policy criteria.
        /// </summary>
        public InputList<Inputs.PolicyAccessIsolationRuleConditionOperandGetArgs> Operands
        {
            get => _operands ?? (_operands = new InputList<Inputs.PolicyAccessIsolationRuleConditionOperandGetArgs>());
            set => _operands = value;
        }

        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        public PolicyAccessIsolationRuleConditionGetArgs()
        {
        }
        public static new PolicyAccessIsolationRuleConditionGetArgs Empty => new PolicyAccessIsolationRuleConditionGetArgs();
    }
}
