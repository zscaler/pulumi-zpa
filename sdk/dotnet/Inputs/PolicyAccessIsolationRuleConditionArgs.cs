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

    public sealed class PolicyAccessIsolationRuleConditionArgs : global::Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        [Input("negated")]
        public Input<bool>? Negated { get; set; }

        [Input("operands")]
        private InputList<Inputs.PolicyAccessIsolationRuleConditionOperandArgs>? _operands;

        /// <summary>
        /// This signifies the various policy criteria.
        /// </summary>
        public InputList<Inputs.PolicyAccessIsolationRuleConditionOperandArgs> Operands
        {
            get => _operands ?? (_operands = new InputList<Inputs.PolicyAccessIsolationRuleConditionOperandArgs>());
            set => _operands = value;
        }

        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        public PolicyAccessIsolationRuleConditionArgs()
        {
        }
        public static new PolicyAccessIsolationRuleConditionArgs Empty => new PolicyAccessIsolationRuleConditionArgs();
    }
}
