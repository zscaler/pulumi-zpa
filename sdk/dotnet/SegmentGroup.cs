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
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Zpa = Zscaler.Zpa;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // ZPA Segment Group resource
    ///     var testSegmentGroup = new Zpa.SegmentGroup("testSegmentGroup", new()
    ///     {
    ///         Description = "test1-segment-group",
    ///         Enabled = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
    /// 
    /// Visit
    /// 
    /// **segment_group** can be imported by using `&lt;SEGMENT GROUP ID&gt;` or `&lt;SEGMENT GROUP NAME&gt;` as the import ID.
    /// 
    /// For example:
    /// 
    /// ```sh
    /// $ pulumi import zpa:index/segmentGroup:SegmentGroup example &lt;segment_group_id&gt;
    /// ```
    /// 
    /// or
    /// 
    /// ```sh
    /// $ pulumi import zpa:index/segmentGroup:SegmentGroup example &lt;segment_group_name&gt;
    /// ```
    /// </summary>
    [ZpaResourceType("zpa:index/segmentGroup:SegmentGroup")]
    public partial class SegmentGroup : global::Pulumi.CustomResource
    {
        [Output("applications")]
        public Output<ImmutableArray<Outputs.SegmentGroupApplication>> Applications { get; private set; } = null!;

        /// <summary>
        /// Description of the app group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether this app group is enabled or not.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        [Output("microtenantId")]
        public Output<string> MicrotenantId { get; private set; } = null!;

        /// <summary>
        /// Name of the app group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a SegmentGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SegmentGroup(string name, SegmentGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("zpa:index/segmentGroup:SegmentGroup", name, args ?? new SegmentGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SegmentGroup(string name, Input<string> id, SegmentGroupState? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/segmentGroup:SegmentGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SegmentGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SegmentGroup Get(string name, Input<string> id, SegmentGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new SegmentGroup(name, id, state, options);
        }
    }

    public sealed class SegmentGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("applications")]
        private InputList<Inputs.SegmentGroupApplicationArgs>? _applications;
        public InputList<Inputs.SegmentGroupApplicationArgs> Applications
        {
            get => _applications ?? (_applications = new InputList<Inputs.SegmentGroupApplicationArgs>());
            set => _applications = value;
        }

        /// <summary>
        /// Description of the app group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether this app group is enabled or not.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// Name of the app group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SegmentGroupArgs()
        {
        }
        public static new SegmentGroupArgs Empty => new SegmentGroupArgs();
    }

    public sealed class SegmentGroupState : global::Pulumi.ResourceArgs
    {
        [Input("applications")]
        private InputList<Inputs.SegmentGroupApplicationGetArgs>? _applications;
        public InputList<Inputs.SegmentGroupApplicationGetArgs> Applications
        {
            get => _applications ?? (_applications = new InputList<Inputs.SegmentGroupApplicationGetArgs>());
            set => _applications = value;
        }

        /// <summary>
        /// Description of the app group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether this app group is enabled or not.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// Name of the app group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SegmentGroupState()
        {
        }
        public static new SegmentGroupState Empty => new SegmentGroupState();
    }
}
