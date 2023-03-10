// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.LSSConfig.Outputs
{

    [OutputType]
    public sealed class GetLSSConfigControllerConfigResult
    {
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string AuditMessage;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (bool)
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly ImmutableArray<string> Filters;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string Format;
        /// <summary>
        /// This field defines the name of the log streaming resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string LssHost;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string LssPort;
        /// <summary>
        /// This field defines the name of the log streaming resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string SourceLogType;
        public readonly bool UseTls;

        [OutputConstructor]
        private GetLSSConfigControllerConfigResult(
            string auditMessage,

            string description,

            bool enabled,

            ImmutableArray<string> filters,

            string format,

            string id,

            string lssHost,

            string lssPort,

            string name,

            string sourceLogType,

            bool useTls)
        {
            AuditMessage = auditMessage;
            Description = description;
            Enabled = enabled;
            Filters = filters;
            Format = format;
            Id = id;
            LssHost = lssHost;
            LssPort = lssPort;
            Name = name;
            SourceLogType = sourceLogType;
            UseTls = useTls;
        }
    }
}
