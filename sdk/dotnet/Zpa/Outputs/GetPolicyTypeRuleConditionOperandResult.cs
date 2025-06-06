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
    public sealed class GetPolicyTypeRuleConditionOperandResult
    {
        public readonly string CreationTime;
        public readonly string Id;
        public readonly string IdpId;
        public readonly string Lhs;
        public readonly string ModifiedBy;
        public readonly string ModifiedTime;
        public readonly string Name;
        public readonly string ObjectType;
        public readonly string Operator;
        public readonly string Rhs;

        [OutputConstructor]
        private GetPolicyTypeRuleConditionOperandResult(
            string creationTime,

            string id,

            string idpId,

            string lhs,

            string modifiedBy,

            string modifiedTime,

            string name,

            string objectType,

            string @operator,

            string rhs)
        {
            CreationTime = creationTime;
            Id = id;
            IdpId = idpId;
            Lhs = lhs;
            ModifiedBy = modifiedBy;
            ModifiedTime = modifiedTime;
            Name = name;
            ObjectType = objectType;
            Operator = @operator;
            Rhs = rhs;
        }
    }
}
