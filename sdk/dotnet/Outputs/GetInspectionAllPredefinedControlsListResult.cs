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
    public sealed class GetInspectionAllPredefinedControlsListResult
    {
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ActionValue;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInspectionAllPredefinedControlsListAssociatedInspectionProfileNameResult> AssociatedInspectionProfileNames;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string Attachment;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ControlGroup;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ControlNumber;
        /// <summary>
        /// (string) Returned values: `WEBSOCKET_PREDEFINED`, `WEBSOCKET_CUSTOM`, `ZSCALER`, `CUSTOM`, `PREDEFINED`
        /// </summary>
        public readonly string ControlType;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string DefaultAction;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string DefaultActionValue;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ParanoiaLevel;
        /// <summary>
        /// (string) Returned values: `HTTP`, `HTTPS`, `FTP`, `RDP`, `SSH`, `WEBSOCKET`
        /// </summary>
        public readonly string ProtocolType;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string Severity;
        /// <summary>
        /// The version of the predefined control, the default is: `OWASP_CRS/3.3.0`
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetInspectionAllPredefinedControlsListResult(
            string action,

            string actionValue,

            ImmutableArray<Outputs.GetInspectionAllPredefinedControlsListAssociatedInspectionProfileNameResult> associatedInspectionProfileNames,

            string attachment,

            string controlGroup,

            string controlNumber,

            string controlType,

            string creationTime,

            string defaultAction,

            string defaultActionValue,

            string description,

            string id,

            string modifiedTime,

            string modifiedby,

            string name,

            string paranoiaLevel,

            string protocolType,

            string severity,

            string version)
        {
            Action = action;
            ActionValue = actionValue;
            AssociatedInspectionProfileNames = associatedInspectionProfileNames;
            Attachment = attachment;
            ControlGroup = controlGroup;
            ControlNumber = controlNumber;
            ControlType = controlType;
            CreationTime = creationTime;
            DefaultAction = defaultAction;
            DefaultActionValue = defaultActionValue;
            Description = description;
            Id = id;
            ModifiedTime = modifiedTime;
            Modifiedby = modifiedby;
            Name = name;
            ParanoiaLevel = paranoiaLevel;
            ProtocolType = protocolType;
            Severity = severity;
            Version = version;
        }
    }
}
