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

    public sealed class PolicyRedirectionRuleServiceEdgeGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        public PolicyRedirectionRuleServiceEdgeGroupArgs()
        {
        }
        public static new PolicyRedirectionRuleServiceEdgeGroupArgs Empty => new PolicyRedirectionRuleServiceEdgeGroupArgs();
    }
}
