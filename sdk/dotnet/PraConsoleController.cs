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
    [Obsolete(@"zpa.index/praconsolecontroller.PraConsoleController has been deprecated in favor of zpa.index/praconsole.PRAConsole")]
    [ZpaResourceType("zpa:index/praConsoleController:PraConsoleController")]
    public partial class PraConsoleController : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the privileged console
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Whether or not the privileged console is enabled
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// The privileged console icon. The icon image is converted to base64 encoded text format
        /// </summary>
        [Output("iconText")]
        public Output<string> IconText { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
        /// microtenantId as 0 when making requests to retrieve data from the Default Microtenant.
        /// </summary>
        [Output("microtenantId")]
        public Output<string> MicrotenantId { get; private set; } = null!;

        /// <summary>
        /// The name of the privileged console
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("praApplication")]
        public Output<Outputs.PraConsoleControllerPraApplication> PraApplication { get; private set; } = null!;

        [Output("praPortals")]
        public Output<ImmutableArray<Outputs.PraConsoleControllerPraPortal>> PraPortals { get; private set; } = null!;


        /// <summary>
        /// Create a PraConsoleController resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PraConsoleController(string name, PraConsoleControllerArgs args, CustomResourceOptions? options = null)
            : base("zpa:index/praConsoleController:PraConsoleController", name, args ?? new PraConsoleControllerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PraConsoleController(string name, Input<string> id, PraConsoleControllerState? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/praConsoleController:PraConsoleController", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PraConsoleController resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PraConsoleController Get(string name, Input<string> id, PraConsoleControllerState? state = null, CustomResourceOptions? options = null)
        {
            return new PraConsoleController(name, id, state, options);
        }
    }

    public sealed class PraConsoleControllerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the privileged console
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether or not the privileged console is enabled
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The privileged console icon. The icon image is converted to base64 encoded text format
        /// </summary>
        [Input("iconText")]
        public Input<string>? IconText { get; set; }

        /// <summary>
        /// The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
        /// microtenantId as 0 when making requests to retrieve data from the Default Microtenant.
        /// </summary>
        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// The name of the privileged console
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("praApplication", required: true)]
        public Input<Inputs.PraConsoleControllerPraApplicationArgs> PraApplication { get; set; } = null!;

        [Input("praPortals", required: true)]
        private InputList<Inputs.PraConsoleControllerPraPortalArgs>? _praPortals;
        public InputList<Inputs.PraConsoleControllerPraPortalArgs> PraPortals
        {
            get => _praPortals ?? (_praPortals = new InputList<Inputs.PraConsoleControllerPraPortalArgs>());
            set => _praPortals = value;
        }

        public PraConsoleControllerArgs()
        {
        }
        public static new PraConsoleControllerArgs Empty => new PraConsoleControllerArgs();
    }

    public sealed class PraConsoleControllerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the privileged console
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether or not the privileged console is enabled
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The privileged console icon. The icon image is converted to base64 encoded text format
        /// </summary>
        [Input("iconText")]
        public Input<string>? IconText { get; set; }

        /// <summary>
        /// The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
        /// microtenantId as 0 when making requests to retrieve data from the Default Microtenant.
        /// </summary>
        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// The name of the privileged console
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("praApplication")]
        public Input<Inputs.PraConsoleControllerPraApplicationGetArgs>? PraApplication { get; set; }

        [Input("praPortals")]
        private InputList<Inputs.PraConsoleControllerPraPortalGetArgs>? _praPortals;
        public InputList<Inputs.PraConsoleControllerPraPortalGetArgs> PraPortals
        {
            get => _praPortals ?? (_praPortals = new InputList<Inputs.PraConsoleControllerPraPortalGetArgs>());
            set => _praPortals = value;
        }

        public PraConsoleControllerState()
        {
        }
        public static new PraConsoleControllerState Empty => new PraConsoleControllerState();
    }
}
