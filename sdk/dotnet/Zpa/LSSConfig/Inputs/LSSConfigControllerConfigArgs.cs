// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.LSSConfig.Inputs
{

    public sealed class LSSConfigControllerConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("auditMessage")]
        public Input<string>? AuditMessage { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("filters")]
        private InputList<string>? _filters;

        /// <summary>
        /// (Optional)
        /// </summary>
        public InputList<string> Filters
        {
            get => _filters ?? (_filters = new InputList<string>());
            set => _filters = value;
        }

        /// <summary>
        /// The format of the LSS resource. The supported formats are: `JSON`, `CSV`, and `TSV`
        /// </summary>
        [Input("format", required: true)]
        public Input<string> Format { get; set; } = null!;

        /// <summary>
        /// App Connector Group ID(s) where logs will be forwarded to.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The IP or FQDN of the SIEM (Log Receiver) where logs will be forwarded to.
        /// </summary>
        [Input("lssHost", required: true)]
        public Input<string> LssHost { get; set; } = null!;

        /// <summary>
        /// The destination port of the SIEM (Log Receiver) where logs will be forwarded to.
        /// </summary>
        [Input("lssPort", required: true)]
        public Input<string> LssPort { get; set; } = null!;

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Refer to the log type documentation
        /// </summary>
        [Input("sourceLogType", required: true)]
        public Input<string> SourceLogType { get; set; } = null!;

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("useTls")]
        public Input<bool>? UseTls { get; set; }

        public LSSConfigControllerConfigArgs()
        {
        }
        public static new LSSConfigControllerConfigArgs Empty => new LSSConfigControllerConfigArgs();
    }
}
