// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.Outputs
{

    [OutputType]
    public sealed class LSSConfigControllerPolicyRuleResourceConditionOperandEntryValue
    {
        public readonly string? Lhs;
        public readonly string? Rhs;

        [OutputConstructor]
        private LSSConfigControllerPolicyRuleResourceConditionOperandEntryValue(
            string? lhs,

            string? rhs)
        {
            Lhs = lhs;
            Rhs = rhs;
        }
    }
}
