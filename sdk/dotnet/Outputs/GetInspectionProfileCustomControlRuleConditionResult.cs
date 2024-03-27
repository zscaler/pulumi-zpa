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
    public sealed class GetInspectionProfileCustomControlRuleConditionResult
    {
        /// <summary>
        /// (string) Signifies the key for the object type Supported values: `SIZE`, `VALUE`
        /// </summary>
        public readonly string Lhs;
        /// <summary>
        /// (string) If lhs is set to SIZE, then the user may pass one of the following: `EQ: Equals`, `LE: Less than or equal to`, `GE: Greater than or equal to`. If the lhs is set to `VALUE`, then the user may pass one of the following: `CONTAINS`, `STARTS_WITH`, `ENDS_WITH`, `RX`.
        /// </summary>
        public readonly string Op;
        /// <summary>
        /// (string) Denotes the value for the given object type. Its value depends on the key. If rules.type is set to REQUEST_METHOD, the conditions.rhs field must have one of the following values: `GET`,`HEAD`, `POST`, `OPTIONS`, `PUT`, `DELETE`, `TRACE`
        /// </summary>
        public readonly string Rhs;

        [OutputConstructor]
        private GetInspectionProfileCustomControlRuleConditionResult(
            string lhs,

            string op,

            string rhs)
        {
            Lhs = lhs;
            Op = op;
            Rhs = rhs;
        }
    }
}