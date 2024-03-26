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

    public sealed class PolicyAccessReorderRuleRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Required) - The ID of the rule to which the order number will be applied.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// (Required) - The order number that should be applied to the respective rule ID.
        /// </summary>
        [Input("order", required: true)]
        public Input<string> Order { get; set; } = null!;

        public PolicyAccessReorderRuleRuleArgs()
        {
        }
        public static new PolicyAccessReorderRuleRuleArgs Empty => new PolicyAccessReorderRuleRuleArgs();
    }
}
