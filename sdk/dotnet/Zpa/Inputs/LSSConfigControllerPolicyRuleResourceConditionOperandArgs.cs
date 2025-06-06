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

    public sealed class LSSConfigControllerPolicyRuleResourceConditionOperandArgs : global::Pulumi.ResourceArgs
    {
        [Input("entryValues")]
        private InputList<Inputs.LSSConfigControllerPolicyRuleResourceConditionOperandEntryValueArgs>? _entryValues;
        public InputList<Inputs.LSSConfigControllerPolicyRuleResourceConditionOperandEntryValueArgs> EntryValues
        {
            get => _entryValues ?? (_entryValues = new InputList<Inputs.LSSConfigControllerPolicyRuleResourceConditionOperandEntryValueArgs>());
            set => _entryValues = value;
        }

        /// <summary>
        /// This is for specifying the policy critiera.
        /// </summary>
        [Input("objectType", required: true)]
        public Input<string> ObjectType { get; set; } = null!;

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// This denotes a list of values for the given object type. The value depend upon the key. If rhs is defined this list will be ignored
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public LSSConfigControllerPolicyRuleResourceConditionOperandArgs()
        {
        }
        public static new LSSConfigControllerPolicyRuleResourceConditionOperandArgs Empty => new LSSConfigControllerPolicyRuleResourceConditionOperandArgs();
    }
}
