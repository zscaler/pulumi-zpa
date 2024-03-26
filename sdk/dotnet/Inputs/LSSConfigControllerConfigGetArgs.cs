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

    public sealed class LSSConfigControllerConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("auditMessage")]
        public Input<string>? AuditMessage { get; set; }

        /// <summary>
        /// Description of the LSS configuration
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether this LSS configuration is enabled or not. Supported values: true, false
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("filters")]
        private InputList<string>? _filters;

        /// <summary>
        /// Filter for the LSS configuration. Format given by the following API to get status codes: /mgmtconfig/v2/admin/lssConfig/statusCodes
        /// </summary>
        public InputList<string> Filters
        {
            get => _filters ?? (_filters = new InputList<string>());
            set => _filters = value;
        }

        /// <summary>
        /// Format of the log type. Format given by the following API to get formats: /mgmtconfig/v2/admin/lssConfig/logType/formats
        /// </summary>
        [Input("format", required: true)]
        public Input<string> Format { get; set; } = null!;

        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Host of the LSS configuration
        /// </summary>
        [Input("lssHost", required: true)]
        public Input<string> LssHost { get; set; } = null!;

        /// <summary>
        /// Port of the LSS configuration
        /// </summary>
        [Input("lssPort", required: true)]
        public Input<string> LssPort { get; set; } = null!;

        /// <summary>
        /// Name of the LSS configuration
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Log type of the LSS configuration
        /// </summary>
        [Input("sourceLogType", required: true)]
        public Input<string> SourceLogType { get; set; } = null!;

        [Input("useTls")]
        public Input<bool>? UseTls { get; set; }

        public LSSConfigControllerConfigGetArgs()
        {
        }
        public static new LSSConfigControllerConfigGetArgs Empty => new LSSConfigControllerConfigGetArgs();
    }
}
