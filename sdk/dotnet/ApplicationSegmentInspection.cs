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
    /// * [Official documentation](https://help.zscaler.com/zpa/about-appprotection-applications)
    /// * [API documentation](https://help.zscaler.com/zpa/configuring-application-segments-using-api)
    /// 
    /// The **zpa_application_segment_inspection** resource creates an inspection application segment in the Zscaler Private Access cloud. This resource can then be referenced in an access policy inspection rule. This resource supports Inspection for both `HTTP` and `HTTPS`.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Zpa = Pulumi.Zpa;
    /// using Zpa = Zscaler.Zpa;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var jenkins = Zpa.GetBaCertificate.Invoke(new()
    ///     {
    ///         Name = "jenkins.example.com",
    ///     });
    /// 
    ///     var @this = new Zpa.ApplicationSegmentInspection("this", new()
    ///     {
    ///         Description = "ZPA_Inspection_Example",
    ///         Enabled = true,
    ///         HealthReporting = "ON_ACCESS",
    ///         BypassType = "NEVER",
    ///         IsCnameEnabled = true,
    ///         TcpPortRanges = new[]
    ///         {
    ///             "443",
    ///             "443",
    ///         },
    ///         DomainNames = new[]
    ///         {
    ///             "jenkins.example.com",
    ///         },
    ///         SegmentGroupId = zpa_segment_group.This.Id,
    ///         ServerGroups = new[]
    ///         {
    ///             new Zpa.Inputs.ApplicationSegmentInspectionServerGroupArgs
    ///             {
    ///                 Ids = new[]
    ///                 {
    ///                     zpa_server_group.This.Id,
    ///                 },
    ///             },
    ///         },
    ///         CommonAppsDto = new Zpa.Inputs.ApplicationSegmentInspectionCommonAppsDtoArgs
    ///         {
    ///             AppsConfigs = new[]
    ///             {
    ///                 new Zpa.Inputs.ApplicationSegmentInspectionCommonAppsDtoAppsConfigArgs
    ///                 {
    ///                     Name = "jenkins.example.com",
    ///                     Domain = "jenkins.example.com",
    ///                     ApplicationProtocol = "HTTPS",
    ///                     ApplicationPort = "443",
    ///                     CertificateId = jenkins.Apply(getBaCertificateResult =&gt; getBaCertificateResult.Id),
    ///                     Enabled = true,
    ///                     AppTypes = new[]
    ///                     {
    ///                         "INSPECT",
    ///                     },
    ///                 },
    ///             },
    ///         },
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
    /// Inspection Application Segment can be imported by using `&lt;APPLICATION SEGMENT ID&gt;` or `&lt;APPLICATION SEGMENT NAME&gt;` as the import ID.
    /// 
    /// ```sh
    /// $ pulumi import zpa:index/applicationSegmentInspection:ApplicationSegmentInspection example &lt;application_segment_id&gt;
    /// ```
    /// 
    /// or
    /// 
    /// ```sh
    /// $ pulumi import zpa:index/applicationSegmentInspection:ApplicationSegmentInspection example &lt;application_segment_name&gt;
    /// ```
    /// </summary>
    [ZpaResourceType("zpa:index/applicationSegmentInspection:ApplicationSegmentInspection")]
    public partial class ApplicationSegmentInspection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether users can bypass ZPA to access applications. Default: NEVER. Supported values: ALWAYS, NEVER, ON_NET.
        /// The value NEVER indicates the use of the client forwarding policy.
        /// </summary>
        [Output("bypassType")]
        public Output<string> BypassType { get; private set; } = null!;

        [Output("commonAppsDto")]
        public Output<Outputs.ApplicationSegmentInspectionCommonAppsDto?> CommonAppsDto { get; private set; } = null!;

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
        public Output<bool> DoubleEncrypt { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        [Output("healthCheckType")]
        public Output<string?> HealthCheckType { get; private set; } = null!;

        /// <summary>
        /// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
        /// </summary>
        [Output("healthReporting")]
        public Output<string?> HealthReporting { get; private set; } = null!;

        [Output("icmpAccessType")]
        public Output<string> IcmpAccessType { get; private set; } = null!;

        [Output("ipAnchored")]
        public Output<bool> IpAnchored { get; private set; } = null!;

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

        /// <summary>
        /// List of the server group IDs.
        /// </summary>
        [Output("serverGroups")]
        public Output<ImmutableArray<Outputs.ApplicationSegmentInspectionServerGroup>> ServerGroups { get; private set; } = null!;

        [Output("tcpKeepAlive")]
        public Output<string> TcpKeepAlive { get; private set; } = null!;

        /// <summary>
        /// tcp port range
        /// </summary>
        [Output("tcpPortRange")]
        public Output<ImmutableArray<Outputs.ApplicationSegmentInspectionTcpPortRange>> TcpPortRange { get; private set; } = null!;

        /// <summary>
        /// TCP port ranges used to access the app.
        /// </summary>
        [Output("tcpPortRanges")]
        public Output<ImmutableArray<string>> TcpPortRanges { get; private set; } = null!;

        /// <summary>
        /// udp port range
        /// </summary>
        [Output("udpPortRange")]
        public Output<ImmutableArray<Outputs.ApplicationSegmentInspectionUdpPortRange>> UdpPortRange { get; private set; } = null!;

        /// <summary>
        /// UDP port ranges used to access the app.
        /// </summary>
        [Output("udpPortRanges")]
        public Output<ImmutableArray<string>> UdpPortRanges { get; private set; } = null!;

        [Output("useInDrMode")]
        public Output<bool?> UseInDrMode { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationSegmentInspection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationSegmentInspection(string name, ApplicationSegmentInspectionArgs args, CustomResourceOptions? options = null)
            : base("zpa:index/applicationSegmentInspection:ApplicationSegmentInspection", name, args ?? new ApplicationSegmentInspectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationSegmentInspection(string name, Input<string> id, ApplicationSegmentInspectionState? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/applicationSegmentInspection:ApplicationSegmentInspection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationSegmentInspection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationSegmentInspection Get(string name, Input<string> id, ApplicationSegmentInspectionState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationSegmentInspection(name, id, state, options);
        }
    }

    public sealed class ApplicationSegmentInspectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether users can bypass ZPA to access applications. Default: NEVER. Supported values: ALWAYS, NEVER, ON_NET.
        /// The value NEVER indicates the use of the client forwarding policy.
        /// </summary>
        [Input("bypassType")]
        public Input<string>? BypassType { get; set; }

        [Input("commonAppsDto")]
        public Input<Inputs.ApplicationSegmentInspectionCommonAppsDtoArgs>? CommonAppsDto { get; set; }

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

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("healthCheckType")]
        public Input<string>? HealthCheckType { get; set; }

        /// <summary>
        /// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
        /// </summary>
        [Input("healthReporting")]
        public Input<string>? HealthReporting { get; set; }

        [Input("icmpAccessType")]
        public Input<string>? IcmpAccessType { get; set; }

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

        /// <summary>
        /// Name of the application.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("passiveHealthEnabled")]
        public Input<bool>? PassiveHealthEnabled { get; set; }

        [Input("segmentGroupId", required: true)]
        public Input<string> SegmentGroupId { get; set; } = null!;

        [Input("segmentGroupName")]
        public Input<string>? SegmentGroupName { get; set; }

        [Input("selectConnectorCloseToApp")]
        public Input<bool>? SelectConnectorCloseToApp { get; set; }

        [Input("serverGroups")]
        private InputList<Inputs.ApplicationSegmentInspectionServerGroupArgs>? _serverGroups;

        /// <summary>
        /// List of the server group IDs.
        /// </summary>
        public InputList<Inputs.ApplicationSegmentInspectionServerGroupArgs> ServerGroups
        {
            get => _serverGroups ?? (_serverGroups = new InputList<Inputs.ApplicationSegmentInspectionServerGroupArgs>());
            set => _serverGroups = value;
        }

        [Input("tcpKeepAlive")]
        public Input<string>? TcpKeepAlive { get; set; }

        [Input("tcpPortRange")]
        private InputList<Inputs.ApplicationSegmentInspectionTcpPortRangeArgs>? _tcpPortRange;

        /// <summary>
        /// tcp port range
        /// </summary>
        public InputList<Inputs.ApplicationSegmentInspectionTcpPortRangeArgs> TcpPortRange
        {
            get => _tcpPortRange ?? (_tcpPortRange = new InputList<Inputs.ApplicationSegmentInspectionTcpPortRangeArgs>());
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
        private InputList<Inputs.ApplicationSegmentInspectionUdpPortRangeArgs>? _udpPortRange;

        /// <summary>
        /// udp port range
        /// </summary>
        public InputList<Inputs.ApplicationSegmentInspectionUdpPortRangeArgs> UdpPortRange
        {
            get => _udpPortRange ?? (_udpPortRange = new InputList<Inputs.ApplicationSegmentInspectionUdpPortRangeArgs>());
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

        public ApplicationSegmentInspectionArgs()
        {
        }
        public static new ApplicationSegmentInspectionArgs Empty => new ApplicationSegmentInspectionArgs();
    }

    public sealed class ApplicationSegmentInspectionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether users can bypass ZPA to access applications. Default: NEVER. Supported values: ALWAYS, NEVER, ON_NET.
        /// The value NEVER indicates the use of the client forwarding policy.
        /// </summary>
        [Input("bypassType")]
        public Input<string>? BypassType { get; set; }

        [Input("commonAppsDto")]
        public Input<Inputs.ApplicationSegmentInspectionCommonAppsDtoGetArgs>? CommonAppsDto { get; set; }

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

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("healthCheckType")]
        public Input<string>? HealthCheckType { get; set; }

        /// <summary>
        /// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
        /// </summary>
        [Input("healthReporting")]
        public Input<string>? HealthReporting { get; set; }

        [Input("icmpAccessType")]
        public Input<string>? IcmpAccessType { get; set; }

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
        private InputList<Inputs.ApplicationSegmentInspectionServerGroupGetArgs>? _serverGroups;

        /// <summary>
        /// List of the server group IDs.
        /// </summary>
        public InputList<Inputs.ApplicationSegmentInspectionServerGroupGetArgs> ServerGroups
        {
            get => _serverGroups ?? (_serverGroups = new InputList<Inputs.ApplicationSegmentInspectionServerGroupGetArgs>());
            set => _serverGroups = value;
        }

        [Input("tcpKeepAlive")]
        public Input<string>? TcpKeepAlive { get; set; }

        [Input("tcpPortRange")]
        private InputList<Inputs.ApplicationSegmentInspectionTcpPortRangeGetArgs>? _tcpPortRange;

        /// <summary>
        /// tcp port range
        /// </summary>
        public InputList<Inputs.ApplicationSegmentInspectionTcpPortRangeGetArgs> TcpPortRange
        {
            get => _tcpPortRange ?? (_tcpPortRange = new InputList<Inputs.ApplicationSegmentInspectionTcpPortRangeGetArgs>());
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
        private InputList<Inputs.ApplicationSegmentInspectionUdpPortRangeGetArgs>? _udpPortRange;

        /// <summary>
        /// udp port range
        /// </summary>
        public InputList<Inputs.ApplicationSegmentInspectionUdpPortRangeGetArgs> UdpPortRange
        {
            get => _udpPortRange ?? (_udpPortRange = new InputList<Inputs.ApplicationSegmentInspectionUdpPortRangeGetArgs>());
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

        public ApplicationSegmentInspectionState()
        {
        }
        public static new ApplicationSegmentInspectionState Empty => new ApplicationSegmentInspectionState();
    }
}
