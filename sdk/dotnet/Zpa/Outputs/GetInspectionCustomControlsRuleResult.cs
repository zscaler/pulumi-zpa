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
    public sealed class GetInspectionCustomControlsRuleResult
    {
        public readonly ImmutableArray<Outputs.GetInspectionCustomControlsRuleConditionResult> Conditions;
        /// <summary>
        /// Name of the rules. If rules.type is set to REQUEST_HEADERS, REQUEST_COOKIES, or RESPONSE_HEADERS, the rules.name field is required.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string Type;

        [OutputConstructor]
        private GetInspectionCustomControlsRuleResult(
            ImmutableArray<Outputs.GetInspectionCustomControlsRuleConditionResult> conditions,

            ImmutableArray<string> names,

            string type)
        {
            Conditions = conditions;
            Names = names;
            Type = type;
        }
    }
}
