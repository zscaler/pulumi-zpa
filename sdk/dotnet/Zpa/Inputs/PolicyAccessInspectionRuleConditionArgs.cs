// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.Inputs
{

    public sealed class PolicyAccessInspectionRuleConditionArgs : global::Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        [Input("operands")]
        private InputList<Inputs.PolicyAccessInspectionRuleConditionOperandArgs>? _operands;

        /// <summary>
        /// This signifies the various policy criteria.
        /// </summary>
        public InputList<Inputs.PolicyAccessInspectionRuleConditionOperandArgs> Operands
        {
            get => _operands ?? (_operands = new InputList<Inputs.PolicyAccessInspectionRuleConditionOperandArgs>());
            set => _operands = value;
        }

        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        public PolicyAccessInspectionRuleConditionArgs()
        {
        }
        public static new PolicyAccessInspectionRuleConditionArgs Empty => new PolicyAccessInspectionRuleConditionArgs();
    }
}
