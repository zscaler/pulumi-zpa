// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.ApplicationSegment
{
    public static class GetApplicationSegmentBrowserAccess
    {
        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zpa.ApplicationSegment.GetApplicationSegmentBrowserAccess.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zpa.ApplicationSegment.GetApplicationSegmentBrowserAccess.Invoke(new()
        ///     {
        ///         Id = "123456789",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetApplicationSegmentBrowserAccessResult> InvokeAsync(GetApplicationSegmentBrowserAccessArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApplicationSegmentBrowserAccessResult>("zpa:ApplicationSegment/getApplicationSegmentBrowserAccess:getApplicationSegmentBrowserAccess", args ?? new GetApplicationSegmentBrowserAccessArgs(), options.WithDefaults());

        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zpa.ApplicationSegment.GetApplicationSegmentBrowserAccess.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zpa.ApplicationSegment.GetApplicationSegmentBrowserAccess.Invoke(new()
        ///     {
        ///         Id = "123456789",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetApplicationSegmentBrowserAccessResult> Invoke(GetApplicationSegmentBrowserAccessInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationSegmentBrowserAccessResult>("zpa:ApplicationSegment/getApplicationSegmentBrowserAccess:getApplicationSegmentBrowserAccess", args ?? new GetApplicationSegmentBrowserAccessInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationSegmentBrowserAccessArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// This field defines the id of the application server.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// This field defines the name of the server.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("tcpPortRanges")]
        private List<Inputs.GetApplicationSegmentBrowserAccessTcpPortRangeArgs>? _tcpPortRanges;

        /// <summary>
        /// (string) TCP port ranges used to access the app.
        /// * `from:`
        /// * `to:`
        /// </summary>
        public List<Inputs.GetApplicationSegmentBrowserAccessTcpPortRangeArgs> TcpPortRanges
        {
            get => _tcpPortRanges ?? (_tcpPortRanges = new List<Inputs.GetApplicationSegmentBrowserAccessTcpPortRangeArgs>());
            set => _tcpPortRanges = value;
        }

        [Input("udpPortRanges")]
        private List<Inputs.GetApplicationSegmentBrowserAccessUdpPortRangeArgs>? _udpPortRanges;

        /// <summary>
        /// (string) UDP port ranges used to access the app.
        /// * `from:`
        /// * `to:`
        /// </summary>
        public List<Inputs.GetApplicationSegmentBrowserAccessUdpPortRangeArgs> UdpPortRanges
        {
            get => _udpPortRanges ?? (_udpPortRanges = new List<Inputs.GetApplicationSegmentBrowserAccessUdpPortRangeArgs>());
            set => _udpPortRanges = value;
        }

        public GetApplicationSegmentBrowserAccessArgs()
        {
        }
        public static new GetApplicationSegmentBrowserAccessArgs Empty => new GetApplicationSegmentBrowserAccessArgs();
    }

    public sealed class GetApplicationSegmentBrowserAccessInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// This field defines the id of the application server.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// This field defines the name of the server.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tcpPortRanges")]
        private InputList<Inputs.GetApplicationSegmentBrowserAccessTcpPortRangeInputArgs>? _tcpPortRanges;

        /// <summary>
        /// (string) TCP port ranges used to access the app.
        /// * `from:`
        /// * `to:`
        /// </summary>
        public InputList<Inputs.GetApplicationSegmentBrowserAccessTcpPortRangeInputArgs> TcpPortRanges
        {
            get => _tcpPortRanges ?? (_tcpPortRanges = new InputList<Inputs.GetApplicationSegmentBrowserAccessTcpPortRangeInputArgs>());
            set => _tcpPortRanges = value;
        }

        [Input("udpPortRanges")]
        private InputList<Inputs.GetApplicationSegmentBrowserAccessUdpPortRangeInputArgs>? _udpPortRanges;

        /// <summary>
        /// (string) UDP port ranges used to access the app.
        /// * `from:`
        /// * `to:`
        /// </summary>
        public InputList<Inputs.GetApplicationSegmentBrowserAccessUdpPortRangeInputArgs> UdpPortRanges
        {
            get => _udpPortRanges ?? (_udpPortRanges = new InputList<Inputs.GetApplicationSegmentBrowserAccessUdpPortRangeInputArgs>());
            set => _udpPortRanges = value;
        }

        public GetApplicationSegmentBrowserAccessInvokeArgs()
        {
        }
        public static new GetApplicationSegmentBrowserAccessInvokeArgs Empty => new GetApplicationSegmentBrowserAccessInvokeArgs();
    }


    [OutputType]
    public sealed class GetApplicationSegmentBrowserAccessResult
    {
        /// <summary>
        /// (string) Indicates whether users can bypass ZPA to access applications. Default: `NEVER`. Supported values: `ALWAYS`, `NEVER`, `ON_NET`. The value `NEVER` indicates the use of the client forwarding policy.
        /// </summary>
        public readonly string BypassType;
        public readonly ImmutableArray<Outputs.GetApplicationSegmentBrowserAccessClientlessAppResult> ClientlessApps;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ConfigSpace;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// List of domains and IPs.
        /// </summary>
        public readonly ImmutableArray<string> DomainNames;
        /// <summary>
        /// (string) Whether Double Encryption is enabled or disabled for the app. Default: false. Boolean: `true`, `false`.
        /// </summary>
        public readonly bool DoubleEncrypt;
        /// <summary>
        /// (bool)
        /// </summary>
        public readonly bool Enabled;
        public readonly string HealthCheckType;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string HealthReporting;
        public readonly string? Id;
        /// <summary>
        /// (bool)
        /// </summary>
        public readonly bool IpAnchored;
        /// <summary>
        /// (bool) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors. Default: true. Boolean: `true`, `false`.
        /// </summary>
        public readonly bool IsCnameEnabled;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// (bool)
        /// </summary>
        public readonly bool PassiveHealthEnabled;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string SegmentGroupId;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string SegmentGroupName;
        public readonly ImmutableArray<Outputs.GetApplicationSegmentBrowserAccessServerGroupResult> ServerGroups;
        /// <summary>
        /// (string) TCP port ranges used to access the app.
        /// </summary>
        public readonly ImmutableArray<string> TcpPortRanges;
        /// <summary>
        /// (string) UDP port ranges used to access the app.
        /// </summary>
        public readonly ImmutableArray<string> UdpPortRanges;

        [OutputConstructor]
        private GetApplicationSegmentBrowserAccessResult(
            string bypassType,

            ImmutableArray<Outputs.GetApplicationSegmentBrowserAccessClientlessAppResult> clientlessApps,

            string configSpace,

            string description,

            ImmutableArray<string> domainNames,

            bool doubleEncrypt,

            bool enabled,

            string healthCheckType,

            string healthReporting,

            string? id,

            bool ipAnchored,

            bool isCnameEnabled,

            string? name,

            bool passiveHealthEnabled,

            string segmentGroupId,

            string segmentGroupName,

            ImmutableArray<Outputs.GetApplicationSegmentBrowserAccessServerGroupResult> serverGroups,

            ImmutableArray<string> tcpPortRanges,

            ImmutableArray<string> udpPortRanges)
        {
            BypassType = bypassType;
            ClientlessApps = clientlessApps;
            ConfigSpace = configSpace;
            Description = description;
            DomainNames = domainNames;
            DoubleEncrypt = doubleEncrypt;
            Enabled = enabled;
            HealthCheckType = healthCheckType;
            HealthReporting = healthReporting;
            Id = id;
            IpAnchored = ipAnchored;
            IsCnameEnabled = isCnameEnabled;
            Name = name;
            PassiveHealthEnabled = passiveHealthEnabled;
            SegmentGroupId = segmentGroupId;
            SegmentGroupName = segmentGroupName;
            ServerGroups = serverGroups;
            TcpPortRanges = tcpPortRanges;
            UdpPortRanges = udpPortRanges;
        }
    }
}
