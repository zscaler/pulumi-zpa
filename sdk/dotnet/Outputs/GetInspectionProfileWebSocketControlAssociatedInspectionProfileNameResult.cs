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
    public sealed class GetInspectionProfileWebSocketControlAssociatedInspectionProfileNameResult
    {
        /// <summary>
        /// This field defines the id of the inspection profile.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// This field defines the name of the inspection profile.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetInspectionProfileWebSocketControlAssociatedInspectionProfileNameResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}