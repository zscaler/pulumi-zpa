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

    public sealed class PolicyCapabilitiesRuleConditionOperandGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("entryValues")]
        private InputList<Inputs.PolicyCapabilitiesRuleConditionOperandEntryValueGetArgs>? _entryValues;
        public InputList<Inputs.PolicyCapabilitiesRuleConditionOperandEntryValueGetArgs> EntryValues
        {
            get => _entryValues ?? (_entryValues = new InputList<Inputs.PolicyCapabilitiesRuleConditionOperandEntryValueGetArgs>());
            set => _entryValues = value;
        }

        /// <summary>
        /// This is for specifying the policy critiera.
        /// </summary>
        [Input("objectType")]
        public Input<string>? ObjectType { get; set; }

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

        public PolicyCapabilitiesRuleConditionOperandGetArgs()
        {
        }
        public static new PolicyCapabilitiesRuleConditionOperandGetArgs Empty => new PolicyCapabilitiesRuleConditionOperandGetArgs();
    }
}