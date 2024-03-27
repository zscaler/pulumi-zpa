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

    public sealed class LSSConfigControllerPolicyRuleResourceConditionArgs : global::Pulumi.ResourceArgs
    {
        [Input("negated")]
        public Input<bool>? Negated { get; set; }

        [Input("operands")]
        private InputList<Inputs.LSSConfigControllerPolicyRuleResourceConditionOperandArgs>? _operands;

        /// <summary>
        /// This signifies the various policy criteria.
        /// </summary>
        public InputList<Inputs.LSSConfigControllerPolicyRuleResourceConditionOperandArgs> Operands
        {
            get => _operands ?? (_operands = new InputList<Inputs.LSSConfigControllerPolicyRuleResourceConditionOperandArgs>());
            set => _operands = value;
        }

        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        public LSSConfigControllerPolicyRuleResourceConditionArgs()
        {
        }
        public static new LSSConfigControllerPolicyRuleResourceConditionArgs Empty => new LSSConfigControllerPolicyRuleResourceConditionArgs();
    }
}