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

    public sealed class PolicyIsolationRuleV2ConditionOperandEntryValueGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("lhs")]
        public Input<string>? Lhs { get; set; }

        [Input("rhs")]
        public Input<string>? Rhs { get; set; }

        public PolicyIsolationRuleV2ConditionOperandEntryValueGetArgs()
        {
        }
        public static new PolicyIsolationRuleV2ConditionOperandEntryValueGetArgs Empty => new PolicyIsolationRuleV2ConditionOperandEntryValueGetArgs();
    }
}
