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
    public sealed class InspectionProfileCustomControl
    {
        /// <summary>
        /// The action of the custom control
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// Denotes the action. Supports any string
        /// </summary>
        public readonly string? ActionValue;
        /// <summary>
        /// The unique identifier of the custom control
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private InspectionProfileCustomControl(
            string? action,

            string? actionValue,

            string id)
        {
            Action = action;
            ActionValue = actionValue;
            Id = id;
        }
    }
}
