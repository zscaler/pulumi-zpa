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
    public sealed class InspectionCustomControlsRule
    {
        public readonly ImmutableArray<Outputs.InspectionCustomControlsRuleCondition> Conditions;
        /// <summary>
        /// Name of the rules. If rules.type is set to REQUEST_HEADERS, REQUEST_COOKIES, or RESPONSE_HEADERS, the rules.name field is required.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        /// <summary>
        /// Type value for the rules.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private InspectionCustomControlsRule(
            ImmutableArray<Outputs.InspectionCustomControlsRuleCondition> conditions,

            ImmutableArray<string> names,

            string? type)
        {
            Conditions = conditions;
            Names = names;
            Type = type;
        }
    }
}
