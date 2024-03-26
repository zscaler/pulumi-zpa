// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Zscaler.Zpa
{
    [ZpaResourceType("zpa:index/inspectionCustomControls:InspectionCustomControls")]
    public partial class InspectionCustomControls : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The performed action
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        [Output("actionValue")]
        public Output<string> ActionValue { get; private set; } = null!;

        /// <summary>
        /// Name of the inspection profile
        /// </summary>
        [Output("associatedInspectionProfileNames")]
        public Output<ImmutableArray<Outputs.InspectionCustomControlsAssociatedInspectionProfileName>> AssociatedInspectionProfileNames { get; private set; } = null!;

        [Output("controlNumber")]
        public Output<string> ControlNumber { get; private set; } = null!;

        /// <summary>
        /// The control rule in JSON format that has the conditions and type of control for the inspection control
        /// </summary>
        [Output("controlRuleJson")]
        public Output<string> ControlRuleJson { get; private set; } = null!;

        [Output("controlType")]
        public Output<string> ControlType { get; private set; } = null!;

        /// <summary>
        /// The performed action
        /// </summary>
        [Output("defaultAction")]
        public Output<string> DefaultAction { get; private set; } = null!;

        /// <summary>
        /// This is used to provide the redirect URL if the default action is set to REDIRECT
        /// </summary>
        [Output("defaultActionValue")]
        public Output<string> DefaultActionValue { get; private set; } = null!;

        /// <summary>
        /// Description of the custom control
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        /// </summary>
        [Output("paranoiaLevel")]
        public Output<string> ParanoiaLevel { get; private set; } = null!;

        [Output("protocolType")]
        public Output<string> ProtocolType { get; private set; } = null!;

        /// <summary>
        /// Rules of the custom controls applied as conditions (JSON)
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.InspectionCustomControlsRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// Severity of the control number
        /// </summary>
        [Output("severity")]
        public Output<string> Severity { get; private set; } = null!;

        /// <summary>
        /// Rules to be applied to the request or response type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a InspectionCustomControls resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InspectionCustomControls(string name, InspectionCustomControlsArgs args, CustomResourceOptions? options = null)
            : base("zpa:index/inspectionCustomControls:InspectionCustomControls", name, args ?? new InspectionCustomControlsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InspectionCustomControls(string name, Input<string> id, InspectionCustomControlsState? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/inspectionCustomControls:InspectionCustomControls", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/zscaler",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing InspectionCustomControls resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InspectionCustomControls Get(string name, Input<string> id, InspectionCustomControlsState? state = null, CustomResourceOptions? options = null)
        {
            return new InspectionCustomControls(name, id, state, options);
        }
    }

    public sealed class InspectionCustomControlsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The performed action
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("actionValue")]
        public Input<string>? ActionValue { get; set; }

        [Input("associatedInspectionProfileNames")]
        private InputList<Inputs.InspectionCustomControlsAssociatedInspectionProfileNameArgs>? _associatedInspectionProfileNames;

        /// <summary>
        /// Name of the inspection profile
        /// </summary>
        public InputList<Inputs.InspectionCustomControlsAssociatedInspectionProfileNameArgs> AssociatedInspectionProfileNames
        {
            get => _associatedInspectionProfileNames ?? (_associatedInspectionProfileNames = new InputList<Inputs.InspectionCustomControlsAssociatedInspectionProfileNameArgs>());
            set => _associatedInspectionProfileNames = value;
        }

        [Input("controlNumber")]
        public Input<string>? ControlNumber { get; set; }

        /// <summary>
        /// The control rule in JSON format that has the conditions and type of control for the inspection control
        /// </summary>
        [Input("controlRuleJson")]
        public Input<string>? ControlRuleJson { get; set; }

        [Input("controlType")]
        public Input<string>? ControlType { get; set; }

        /// <summary>
        /// The performed action
        /// </summary>
        [Input("defaultAction", required: true)]
        public Input<string> DefaultAction { get; set; } = null!;

        /// <summary>
        /// This is used to provide the redirect URL if the default action is set to REDIRECT
        /// </summary>
        [Input("defaultActionValue")]
        public Input<string>? DefaultActionValue { get; set; }

        /// <summary>
        /// Description of the custom control
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        /// </summary>
        [Input("paranoiaLevel")]
        public Input<string>? ParanoiaLevel { get; set; }

        [Input("protocolType")]
        public Input<string>? ProtocolType { get; set; }

        [Input("rules")]
        private InputList<Inputs.InspectionCustomControlsRuleArgs>? _rules;

        /// <summary>
        /// Rules of the custom controls applied as conditions (JSON)
        /// </summary>
        public InputList<Inputs.InspectionCustomControlsRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.InspectionCustomControlsRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Severity of the control number
        /// </summary>
        [Input("severity", required: true)]
        public Input<string> Severity { get; set; } = null!;

        /// <summary>
        /// Rules to be applied to the request or response type
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("version")]
        public Input<string>? Version { get; set; }

        public InspectionCustomControlsArgs()
        {
        }
        public static new InspectionCustomControlsArgs Empty => new InspectionCustomControlsArgs();
    }

    public sealed class InspectionCustomControlsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The performed action
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("actionValue")]
        public Input<string>? ActionValue { get; set; }

        [Input("associatedInspectionProfileNames")]
        private InputList<Inputs.InspectionCustomControlsAssociatedInspectionProfileNameGetArgs>? _associatedInspectionProfileNames;

        /// <summary>
        /// Name of the inspection profile
        /// </summary>
        public InputList<Inputs.InspectionCustomControlsAssociatedInspectionProfileNameGetArgs> AssociatedInspectionProfileNames
        {
            get => _associatedInspectionProfileNames ?? (_associatedInspectionProfileNames = new InputList<Inputs.InspectionCustomControlsAssociatedInspectionProfileNameGetArgs>());
            set => _associatedInspectionProfileNames = value;
        }

        [Input("controlNumber")]
        public Input<string>? ControlNumber { get; set; }

        /// <summary>
        /// The control rule in JSON format that has the conditions and type of control for the inspection control
        /// </summary>
        [Input("controlRuleJson")]
        public Input<string>? ControlRuleJson { get; set; }

        [Input("controlType")]
        public Input<string>? ControlType { get; set; }

        /// <summary>
        /// The performed action
        /// </summary>
        [Input("defaultAction")]
        public Input<string>? DefaultAction { get; set; }

        /// <summary>
        /// This is used to provide the redirect URL if the default action is set to REDIRECT
        /// </summary>
        [Input("defaultActionValue")]
        public Input<string>? DefaultActionValue { get; set; }

        /// <summary>
        /// Description of the custom control
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        /// </summary>
        [Input("paranoiaLevel")]
        public Input<string>? ParanoiaLevel { get; set; }

        [Input("protocolType")]
        public Input<string>? ProtocolType { get; set; }

        [Input("rules")]
        private InputList<Inputs.InspectionCustomControlsRuleGetArgs>? _rules;

        /// <summary>
        /// Rules of the custom controls applied as conditions (JSON)
        /// </summary>
        public InputList<Inputs.InspectionCustomControlsRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.InspectionCustomControlsRuleGetArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Severity of the control number
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// Rules to be applied to the request or response type
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        public InspectionCustomControlsState()
        {
        }
        public static new InspectionCustomControlsState Empty => new InspectionCustomControlsState();
    }
}
