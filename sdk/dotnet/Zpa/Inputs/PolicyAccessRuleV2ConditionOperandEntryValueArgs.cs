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

    public sealed class PolicyAccessRuleV2ConditionOperandEntryValueArgs : global::Pulumi.ResourceArgs
    {
        [Input("lhs")]
        public Input<string>? Lhs { get; set; }

        [Input("rhs")]
        public Input<string>? Rhs { get; set; }

        public PolicyAccessRuleV2ConditionOperandEntryValueArgs()
        {
        }
        public static new PolicyAccessRuleV2ConditionOperandEntryValueArgs Empty => new PolicyAccessRuleV2ConditionOperandEntryValueArgs();
    }
}
