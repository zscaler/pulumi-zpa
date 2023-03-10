// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.AccessPolicy.Outputs
{

    [OutputType]
    public sealed class PolicyAccessRuleConditionOperand
    {
        /// <summary>
        /// (Optional) The ID of a server group resource
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// (Optional)
        /// </summary>
        public readonly string? IdpId;
        /// <summary>
        /// (Optional) LHS must always carry the string value ``id`` or the attribute ID of the resource being associated with the rule.
        /// </summary>
        public readonly string Lhs;
        /// <summary>
        /// (Optional)
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// (Optional) This is for specifying the policy critiera. Supported values: `APP`, `APP_GROUP`, `SAML`, `IDP`, `CLIENT_TYPE`, `TRUSTED_NETWORK`, `POSTURE`, `SCIM`, `SCIM_GROUP`, and `CLOUD_CONNECTOR_GROUP`. `TRUSTED_NETWORK`, and `CLIENT_TYPE`.
        /// </summary>
        public readonly string ObjectType;
        /// <summary>
        /// (Optional) RHS is either the ID attribute of a resource or fixed string value. Refer to the chart below for further details.
        /// </summary>
        public readonly string? Rhs;
        public readonly ImmutableArray<string> RhsLists;

        [OutputConstructor]
        private PolicyAccessRuleConditionOperand(
            string? id,

            string? idpId,

            string lhs,

            string? name,

            string objectType,

            string? rhs,

            ImmutableArray<string> rhsLists)
        {
            Id = id;
            IdpId = idpId;
            Lhs = lhs;
            Name = name;
            ObjectType = objectType;
            Rhs = rhs;
            RhsLists = rhsLists;
        }
    }
}
