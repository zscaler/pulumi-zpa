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

    public sealed class InspectionCustomControlsRuleConditionsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Signifies the key for the object type
        /// </summary>
        [Input("lhs")]
        public Input<string>? Lhs { get; set; }

        /// <summary>
        /// Denotes the operation type.
        /// </summary>
        [Input("op")]
        public Input<string>? Op { get; set; }

        /// <summary>
        /// Denotes the value for the given object type. Its value depends on the key.
        /// </summary>
        [Input("rhs")]
        public Input<string>? Rhs { get; set; }

        public InspectionCustomControlsRuleConditionsGetArgs()
        {
        }
        public static new InspectionCustomControlsRuleConditionsGetArgs Empty => new InspectionCustomControlsRuleConditionsGetArgs();
    }
}