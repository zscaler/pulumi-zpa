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

    public sealed class PolicyBrowserProtectionRuleConditionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("operands")]
        private InputList<Inputs.PolicyBrowserProtectionRuleConditionOperandGetArgs>? _operands;

        /// <summary>
        /// This signifies the various policy criteria.
        /// </summary>
        public InputList<Inputs.PolicyBrowserProtectionRuleConditionOperandGetArgs> Operands
        {
            get => _operands ?? (_operands = new InputList<Inputs.PolicyBrowserProtectionRuleConditionOperandGetArgs>());
            set => _operands = value;
        }

        [Input("operator")]
        public Input<string>? Operator { get; set; }

        public PolicyBrowserProtectionRuleConditionGetArgs()
        {
        }
        public static new PolicyBrowserProtectionRuleConditionGetArgs Empty => new PolicyBrowserProtectionRuleConditionGetArgs();
    }
}
