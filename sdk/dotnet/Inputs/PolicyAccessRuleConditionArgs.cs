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

    public sealed class PolicyAccessRuleConditionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional) The ID of a server group resource
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// (Optional) The ID of the microtenant the resource is to be associated with.
        /// </summary>
        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// (Optional) Supported values: ``true`` or ``false``
        /// </summary>
        [Input("negated")]
        public Input<bool>? Negated { get; set; }

        [Input("operands")]
        private InputList<Inputs.PolicyAccessRuleConditionOperandArgs>? _operands;

        /// <summary>
        /// (Optional) - Operands block must be repeated if multiple per `object_type` conditions are to be added to the rule.
        /// </summary>
        public InputList<Inputs.PolicyAccessRuleConditionOperandArgs> Operands
        {
            get => _operands ?? (_operands = new InputList<Inputs.PolicyAccessRuleConditionOperandArgs>());
            set => _operands = value;
        }

        /// <summary>
        /// (Optional) Supported values: ``AND``, and ``OR``
        /// </summary>
        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        public PolicyAccessRuleConditionArgs()
        {
        }
        public static new PolicyAccessRuleConditionArgs Empty => new PolicyAccessRuleConditionArgs();
    }
}