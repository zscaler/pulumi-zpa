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
    public sealed class GetCustomerVersionProfileCustomScopeCustomerIdResult
    {
        public readonly string CustomerId;
        public readonly bool ExcludeConstellation;
        public readonly string Name;

        [OutputConstructor]
        private GetCustomerVersionProfileCustomScopeCustomerIdResult(
            string customerId,

            bool excludeConstellation,

            string name)
        {
            CustomerId = customerId;
            ExcludeConstellation = excludeConstellation;
            Name = name;
        }
    }
}
