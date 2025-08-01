// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa
{
    /// <summary>
    /// ## Import
    /// 
    /// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
    /// 
    /// Visit
    /// 
    /// Application Segment can be imported by using `&lt;APPLICATION SEGMENT ID&gt;` or `&lt;APPLICATION SEGMENT NAME&gt;` as the import ID.
    /// 
    /// ```sh
    /// $ pulumi import zpa:index/applicationSegment:ApplicationSegment example &lt;application_segment_id&gt;
    /// ```
    /// 
    /// or
    /// 
    /// ```sh
    /// $ pulumi import zpa:index/applicationSegment:ApplicationSegment example &lt;application_segment_name&gt;
    /// ```
    /// </summary>
    [ZpaResourceType("zpa:index/applicationSegment:ApplicationSegment")]
    public partial class ApplicationSegment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// If set to true, designates the application segment for API traffic inspection
        /// </summary>
        [Output("apiProtectionEnabled")]
        public Output<bool?> ApiProtectionEnabled { get; private set; } = null!;

        [Output("bypassOnReauth")]
        public Output<bool> BypassOnReauth { get; private set; } = null!;

        /// <summary>
        /// Indicates whether users can bypass ZPA to access applications.
        /// </summary>
        [Output("bypassType")]
        public Output<string> BypassType { get; private set; } = null!;

        [Output("configSpace")]
        public Output<string?> ConfigSpace { get; private set; } = null!;

        /// <summary>
        /// Description of the application.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// List of domains and IPs.
        /// </summary>
        [Output("domainNames")]
        public Output<ImmutableArray<string>> DomainNames { get; private set; } = null!;

        /// <summary>
        /// Whether Double Encryption is enabled or disabled for the app.
        /// </summary>
        [Output("doubleEncrypt")]
        public Output<bool?> DoubleEncrypt { get; private set; } = null!;

        /// <summary>
        /// Whether this application is enabled or not.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        [Output("fqdnDnsCheck")]
        public Output<bool?> FqdnDnsCheck { get; private set; } = null!;

        [Output("healthCheckType")]
        public Output<string?> HealthCheckType { get; private set; } = null!;

        /// <summary>
        /// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
        /// </summary>
        [Output("healthReporting")]
        public Output<string?> HealthReporting { get; private set; } = null!;

        [Output("icmpAccessType")]
        public Output<string?> IcmpAccessType { get; private set; } = null!;

        /// <summary>
        /// Indicates if Inspect Traffic with ZIA is enabled for the application.
        /// </summary>
        [Output("inspectTrafficWithZia")]
        public Output<bool?> InspectTrafficWithZia { get; private set; } = null!;

        [Output("ipAnchored")]
        public Output<bool?> IpAnchored { get; private set; } = null!;

        /// <summary>
        /// Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the
        /// connectors.
        /// </summary>
        [Output("isCnameEnabled")]
        public Output<bool> IsCnameEnabled { get; private set; } = null!;

        [Output("isIncompleteDrConfig")]
        public Output<bool?> IsIncompleteDrConfig { get; private set; } = null!;

        [Output("matchStyle")]
        public Output<string> MatchStyle { get; private set; } = null!;

        [Output("microtenantId")]
        public Output<string?> MicrotenantId { get; private set; } = null!;

        /// <summary>
        /// Name of the application.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("passiveHealthEnabled")]
        public Output<bool> PassiveHealthEnabled { get; private set; } = null!;

        [Output("segmentGroupId")]
        public Output<string> SegmentGroupId { get; private set; } = null!;

        [Output("segmentGroupName")]
        public Output<string> SegmentGroupName { get; private set; } = null!;

        [Output("selectConnectorCloseToApp")]
        public Output<bool?> SelectConnectorCloseToApp { get; private set; } = null!;

        [Output("serverGroups")]
        public Output<ImmutableArray<Outputs.ApplicationSegmentServerGroup>> ServerGroups { get; private set; } = null!;

        /// <summary>
        /// Share the Application Segment to microtenants
        /// </summary>
        [Output("shareToMicrotenants")]
        public Output<ImmutableArray<string>> ShareToMicrotenants { get; private set; } = null!;

        [Output("tcpKeepAlive")]
        public Output<string> TcpKeepAlive { get; private set; } = null!;

        /// <summary>
        /// tcp port range
        /// </summary>
        [Output("tcpPortRange")]
        public Output<ImmutableArray<Outputs.ApplicationSegmentTcpPortRange>> TcpPortRange { get; private set; } = null!;

        /// <summary>
        /// TCP port ranges used to access the app.
        /// </summary>
        [Output("tcpPortRanges")]
        public Output<ImmutableArray<string>> TcpPortRanges { get; private set; } = null!;

        /// <summary>
        /// udp port range
        /// </summary>
        [Output("udpPortRange")]
        public Output<ImmutableArray<Outputs.ApplicationSegmentUdpPortRange>> UdpPortRange { get; private set; } = null!;

        /// <summary>
        /// UDP port ranges used to access the app.
        /// </summary>
        [Output("udpPortRanges")]
        public Output<ImmutableArray<string>> UdpPortRanges { get; private set; } = null!;

        [Output("useInDrMode")]
        public Output<bool?> UseInDrMode { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationSegment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationSegment(string name, ApplicationSegmentArgs args, CustomResourceOptions? options = null)
            : base("zpa:index/applicationSegment:ApplicationSegment", name, args ?? new ApplicationSegmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationSegment(string name, Input<string> id, ApplicationSegmentState? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/applicationSegment:ApplicationSegment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationSegment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationSegment Get(string name, Input<string> id, ApplicationSegmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationSegment(name, id, state, options);
        }
    }

    public sealed class ApplicationSegmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set to true, designates the application segment for API traffic inspection
        /// </summary>
        [Input("apiProtectionEnabled")]
        public Input<bool>? ApiProtectionEnabled { get; set; }

        [Input("bypassOnReauth")]
        public Input<bool>? BypassOnReauth { get; set; }

        /// <summary>
        /// Indicates whether users can bypass ZPA to access applications.
        /// </summary>
        [Input("bypassType")]
        public Input<string>? BypassType { get; set; }

        [Input("configSpace")]
        public Input<string>? ConfigSpace { get; set; }

        /// <summary>
        /// Description of the application.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("domainNames", required: true)]
        private InputList<string>? _domainNames;

        /// <summary>
        /// List of domains and IPs.
        /// </summary>
        public InputList<string> DomainNames
        {
            get => _domainNames ?? (_domainNames = new InputList<string>());
            set => _domainNames = value;
        }

        /// <summary>
        /// Whether Double Encryption is enabled or disabled for the app.
        /// </summary>
        [Input("doubleEncrypt")]
        public Input<bool>? DoubleEncrypt { get; set; }

        /// <summary>
        /// Whether this application is enabled or not.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("fqdnDnsCheck")]
        public Input<bool>? FqdnDnsCheck { get; set; }

        [Input("healthCheckType")]
        public Input<string>? HealthCheckType { get; set; }

        /// <summary>
        /// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
        /// </summary>
        [Input("healthReporting")]
        public Input<string>? HealthReporting { get; set; }

        [Input("icmpAccessType")]
        public Input<string>? IcmpAccessType { get; set; }

        /// <summary>
        /// Indicates if Inspect Traffic with ZIA is enabled for the application.
        /// </summary>
        [Input("inspectTrafficWithZia")]
        public Input<bool>? InspectTrafficWithZia { get; set; }

        [Input("ipAnchored")]
        public Input<bool>? IpAnchored { get; set; }

        /// <summary>
        /// Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the
        /// connectors.
        /// </summary>
        [Input("isCnameEnabled")]
        public Input<bool>? IsCnameEnabled { get; set; }

        [Input("isIncompleteDrConfig")]
        public Input<bool>? IsIncompleteDrConfig { get; set; }

        [Input("matchStyle")]
        public Input<string>? MatchStyle { get; set; }

        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// Name of the application.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("passiveHealthEnabled")]
        public Input<bool>? PassiveHealthEnabled { get; set; }

        [Input("segmentGroupId")]
        public Input<string>? SegmentGroupId { get; set; }

        [Input("segmentGroupName")]
        public Input<string>? SegmentGroupName { get; set; }

        [Input("selectConnectorCloseToApp")]
        public Input<bool>? SelectConnectorCloseToApp { get; set; }

        [Input("serverGroups")]
        private InputList<Inputs.ApplicationSegmentServerGroupArgs>? _serverGroups;
        public InputList<Inputs.ApplicationSegmentServerGroupArgs> ServerGroups
        {
            get => _serverGroups ?? (_serverGroups = new InputList<Inputs.ApplicationSegmentServerGroupArgs>());
            set => _serverGroups = value;
        }

        [Input("shareToMicrotenants")]
        private InputList<string>? _shareToMicrotenants;

        /// <summary>
        /// Share the Application Segment to microtenants
        /// </summary>
        public InputList<string> ShareToMicrotenants
        {
            get => _shareToMicrotenants ?? (_shareToMicrotenants = new InputList<string>());
            set => _shareToMicrotenants = value;
        }

        [Input("tcpKeepAlive")]
        public Input<string>? TcpKeepAlive { get; set; }

        [Input("tcpPortRange")]
        private InputList<Inputs.ApplicationSegmentTcpPortRangeArgs>? _tcpPortRange;

        /// <summary>
        /// tcp port range
        /// </summary>
        public InputList<Inputs.ApplicationSegmentTcpPortRangeArgs> TcpPortRange
        {
            get => _tcpPortRange ?? (_tcpPortRange = new InputList<Inputs.ApplicationSegmentTcpPortRangeArgs>());
            set => _tcpPortRange = value;
        }

        [Input("tcpPortRanges")]
        private InputList<string>? _tcpPortRanges;

        /// <summary>
        /// TCP port ranges used to access the app.
        /// </summary>
        public InputList<string> TcpPortRanges
        {
            get => _tcpPortRanges ?? (_tcpPortRanges = new InputList<string>());
            set => _tcpPortRanges = value;
        }

        [Input("udpPortRange")]
        private InputList<Inputs.ApplicationSegmentUdpPortRangeArgs>? _udpPortRange;

        /// <summary>
        /// udp port range
        /// </summary>
        public InputList<Inputs.ApplicationSegmentUdpPortRangeArgs> UdpPortRange
        {
            get => _udpPortRange ?? (_udpPortRange = new InputList<Inputs.ApplicationSegmentUdpPortRangeArgs>());
            set => _udpPortRange = value;
        }

        [Input("udpPortRanges")]
        private InputList<string>? _udpPortRanges;

        /// <summary>
        /// UDP port ranges used to access the app.
        /// </summary>
        public InputList<string> UdpPortRanges
        {
            get => _udpPortRanges ?? (_udpPortRanges = new InputList<string>());
            set => _udpPortRanges = value;
        }

        [Input("useInDrMode")]
        public Input<bool>? UseInDrMode { get; set; }

        public ApplicationSegmentArgs()
        {
        }
        public static new ApplicationSegmentArgs Empty => new ApplicationSegmentArgs();
    }

    public sealed class ApplicationSegmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set to true, designates the application segment for API traffic inspection
        /// </summary>
        [Input("apiProtectionEnabled")]
        public Input<bool>? ApiProtectionEnabled { get; set; }

        [Input("bypassOnReauth")]
        public Input<bool>? BypassOnReauth { get; set; }

        /// <summary>
        /// Indicates whether users can bypass ZPA to access applications.
        /// </summary>
        [Input("bypassType")]
        public Input<string>? BypassType { get; set; }

        [Input("configSpace")]
        public Input<string>? ConfigSpace { get; set; }

        /// <summary>
        /// Description of the application.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("domainNames")]
        private InputList<string>? _domainNames;

        /// <summary>
        /// List of domains and IPs.
        /// </summary>
        public InputList<string> DomainNames
        {
            get => _domainNames ?? (_domainNames = new InputList<string>());
            set => _domainNames = value;
        }

        /// <summary>
        /// Whether Double Encryption is enabled or disabled for the app.
        /// </summary>
        [Input("doubleEncrypt")]
        public Input<bool>? DoubleEncrypt { get; set; }

        /// <summary>
        /// Whether this application is enabled or not.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("fqdnDnsCheck")]
        public Input<bool>? FqdnDnsCheck { get; set; }

        [Input("healthCheckType")]
        public Input<string>? HealthCheckType { get; set; }

        /// <summary>
        /// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
        /// </summary>
        [Input("healthReporting")]
        public Input<string>? HealthReporting { get; set; }

        [Input("icmpAccessType")]
        public Input<string>? IcmpAccessType { get; set; }

        /// <summary>
        /// Indicates if Inspect Traffic with ZIA is enabled for the application.
        /// </summary>
        [Input("inspectTrafficWithZia")]
        public Input<bool>? InspectTrafficWithZia { get; set; }

        [Input("ipAnchored")]
        public Input<bool>? IpAnchored { get; set; }

        /// <summary>
        /// Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the
        /// connectors.
        /// </summary>
        [Input("isCnameEnabled")]
        public Input<bool>? IsCnameEnabled { get; set; }

        [Input("isIncompleteDrConfig")]
        public Input<bool>? IsIncompleteDrConfig { get; set; }

        [Input("matchStyle")]
        public Input<string>? MatchStyle { get; set; }

        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// Name of the application.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("passiveHealthEnabled")]
        public Input<bool>? PassiveHealthEnabled { get; set; }

        [Input("segmentGroupId")]
        public Input<string>? SegmentGroupId { get; set; }

        [Input("segmentGroupName")]
        public Input<string>? SegmentGroupName { get; set; }

        [Input("selectConnectorCloseToApp")]
        public Input<bool>? SelectConnectorCloseToApp { get; set; }

        [Input("serverGroups")]
        private InputList<Inputs.ApplicationSegmentServerGroupGetArgs>? _serverGroups;
        public InputList<Inputs.ApplicationSegmentServerGroupGetArgs> ServerGroups
        {
            get => _serverGroups ?? (_serverGroups = new InputList<Inputs.ApplicationSegmentServerGroupGetArgs>());
            set => _serverGroups = value;
        }

        [Input("shareToMicrotenants")]
        private InputList<string>? _shareToMicrotenants;

        /// <summary>
        /// Share the Application Segment to microtenants
        /// </summary>
        public InputList<string> ShareToMicrotenants
        {
            get => _shareToMicrotenants ?? (_shareToMicrotenants = new InputList<string>());
            set => _shareToMicrotenants = value;
        }

        [Input("tcpKeepAlive")]
        public Input<string>? TcpKeepAlive { get; set; }

        [Input("tcpPortRange")]
        private InputList<Inputs.ApplicationSegmentTcpPortRangeGetArgs>? _tcpPortRange;

        /// <summary>
        /// tcp port range
        /// </summary>
        public InputList<Inputs.ApplicationSegmentTcpPortRangeGetArgs> TcpPortRange
        {
            get => _tcpPortRange ?? (_tcpPortRange = new InputList<Inputs.ApplicationSegmentTcpPortRangeGetArgs>());
            set => _tcpPortRange = value;
        }

        [Input("tcpPortRanges")]
        private InputList<string>? _tcpPortRanges;

        /// <summary>
        /// TCP port ranges used to access the app.
        /// </summary>
        public InputList<string> TcpPortRanges
        {
            get => _tcpPortRanges ?? (_tcpPortRanges = new InputList<string>());
            set => _tcpPortRanges = value;
        }

        [Input("udpPortRange")]
        private InputList<Inputs.ApplicationSegmentUdpPortRangeGetArgs>? _udpPortRange;

        /// <summary>
        /// udp port range
        /// </summary>
        public InputList<Inputs.ApplicationSegmentUdpPortRangeGetArgs> UdpPortRange
        {
            get => _udpPortRange ?? (_udpPortRange = new InputList<Inputs.ApplicationSegmentUdpPortRangeGetArgs>());
            set => _udpPortRange = value;
        }

        [Input("udpPortRanges")]
        private InputList<string>? _udpPortRanges;

        /// <summary>
        /// UDP port ranges used to access the app.
        /// </summary>
        public InputList<string> UdpPortRanges
        {
            get => _udpPortRanges ?? (_udpPortRanges = new InputList<string>());
            set => _udpPortRanges = value;
        }

        [Input("useInDrMode")]
        public Input<bool>? UseInDrMode { get; set; }

        public ApplicationSegmentState()
        {
        }
        public static new ApplicationSegmentState Empty => new ApplicationSegmentState();
    }
}
