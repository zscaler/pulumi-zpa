// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Zscaler.Zpa.Outputs
{

    [OutputType]
    public sealed class PolicyAccessIsolationRuleConditionOperand
    {
        public readonly string? Id;
        public readonly string? IdpId;
        /// <summary>
        /// This signifies the key for the object type. String ID example: id
        /// </summary>
        public readonly string Lhs;
        /// <summary>
        /// This denotes the value for the given object type. Its value depends upon the key.
        /// </summary>
        public readonly string? MicrotenantId;
        public readonly string? Name;
        /// <summary>
        /// This is for specifying the policy critiera.
        /// </summary>
        public readonly string ObjectType;
        /// <summary>
        /// This denotes the value for the given object type. Its value depends upon the key.
        /// </summary>
        public readonly string? Rhs;
        /// <summary>
        /// This denotes a list of values for the given object type. The value depend upon the key. If rhs is defined this list will be ignored
        /// </summary>
        public readonly ImmutableArray<string> RhsLists;

        [OutputConstructor]
        private PolicyAccessIsolationRuleConditionOperand(
            string? id,

            string? idpId,

            string lhs,

            string? microtenantId,

            string? name,

            string objectType,

            string? rhs,

            ImmutableArray<string> rhsLists)
        {
            Id = id;
            IdpId = idpId;
            Lhs = lhs;
            MicrotenantId = microtenantId;
            Name = name;
            ObjectType = objectType;
            Rhs = rhs;
            RhsLists = rhsLists;
        }
    }
}
