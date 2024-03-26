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
    public static class GetServiceEdgeController
    {
        /// <summary>
        /// Use the **zpa_service_edge_controller** data source to get information about a service edge controller in the Zscaler Private Access cloud. This data source can then be referenced in a Service Edge Group and Provisioning Key.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zpa.GetServiceEdgeController.Invoke(new()
        ///     {
        ///         Name = "On-Prem-PSE",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetServiceEdgeControllerResult> InvokeAsync(GetServiceEdgeControllerArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServiceEdgeControllerResult>("zpa:index/getServiceEdgeController:getServiceEdgeController", args ?? new GetServiceEdgeControllerArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zpa_service_edge_controller** data source to get information about a service edge controller in the Zscaler Private Access cloud. This data source can then be referenced in a Service Edge Group and Provisioning Key.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zpa.GetServiceEdgeController.Invoke(new()
        ///     {
        ///         Name = "On-Prem-PSE",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetServiceEdgeControllerResult> Invoke(GetServiceEdgeControllerInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServiceEdgeControllerResult>("zpa:index/getServiceEdgeController:getServiceEdgeController", args ?? new GetServiceEdgeControllerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceEdgeControllerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the service edge controller to be exported.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetServiceEdgeControllerArgs()
        {
        }
        public static new GetServiceEdgeControllerArgs Empty => new GetServiceEdgeControllerArgs();
    }

    public sealed class GetServiceEdgeControllerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the service edge controller to be exported.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetServiceEdgeControllerInvokeArgs()
        {
        }
        public static new GetServiceEdgeControllerInvokeArgs Empty => new GetServiceEdgeControllerInvokeArgs();
    }


    [OutputType]
    public sealed class GetServiceEdgeControllerResult
    {
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ApplicationStartTime;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ControlChannelStatus;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string CtrlBrokerName;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string CurrentVersion;
        /// <summary>
        /// (string) - Description of the App Connector.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (bool) Whether this Service Edge Controller is enabled or not. Default value: `true`. Supported values: `true`, `false`
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly ImmutableDictionary<string, object> EnrollmentCert;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ExpectedUpgradeTime;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ExpectedVersion;
        public readonly string Fingerprint;
        public readonly string Id;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string IpAcl;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string IssuedCertId;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string LastBrokerConnectTime;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string LastBrokerConnectTimeDuration;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string LastBrokerDisconnectTime;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string LastBrokerDisconnectTimeDuration;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string LastUpgradeTime;
        /// <summary>
        /// (string) Latitude of the Service Edge Controller. Integer or decimal. With values in the range of `-90` to `90`
        /// </summary>
        public readonly string Latitude;
        public readonly string ListenIps;
        /// <summary>
        /// (string) Location of the Service Edge Controller.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// (string) Longitude of the Service Edge Controller. Integer or decimal. With values in the range of `-180` to `180`
        /// </summary>
        public readonly string Longitude;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ModifiedBy;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ModifiedTime;
        public readonly string? Name;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string Platform;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string PreviousVersion;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string PrivateIp;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ProvisioningKeyId;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ProvisioningKeyName;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string PublicIp;
        public readonly ImmutableArray<string> PublishIps;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string SargeVersion;
        public readonly string ServiceEdgeGroupId;
        public readonly string ServiceEdgeGroupName;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string UpgradeAttempt;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string UpgradeStatus;
        public readonly ImmutableArray<Outputs.GetServiceEdgeControllerZpnSubModuleUpgradeListResult> ZpnSubModuleUpgradeLists;

        [OutputConstructor]
        private GetServiceEdgeControllerResult(
            string applicationStartTime,

            string controlChannelStatus,

            string creationTime,

            string ctrlBrokerName,

            string currentVersion,

            string description,

            bool enabled,

            ImmutableDictionary<string, object> enrollmentCert,

            string expectedUpgradeTime,

            string expectedVersion,

            string fingerprint,

            string id,

            string ipAcl,

            string issuedCertId,

            string lastBrokerConnectTime,

            string lastBrokerConnectTimeDuration,

            string lastBrokerDisconnectTime,

            string lastBrokerDisconnectTimeDuration,

            string lastUpgradeTime,

            string latitude,

            string listenIps,

            string location,

            string longitude,

            string modifiedBy,

            string modifiedTime,

            string? name,

            string platform,

            string previousVersion,

            string privateIp,

            string provisioningKeyId,

            string provisioningKeyName,

            string publicIp,

            ImmutableArray<string> publishIps,

            string sargeVersion,

            string serviceEdgeGroupId,

            string serviceEdgeGroupName,

            string upgradeAttempt,

            string upgradeStatus,

            ImmutableArray<Outputs.GetServiceEdgeControllerZpnSubModuleUpgradeListResult> zpnSubModuleUpgradeLists)
        {
            ApplicationStartTime = applicationStartTime;
            ControlChannelStatus = controlChannelStatus;
            CreationTime = creationTime;
            CtrlBrokerName = ctrlBrokerName;
            CurrentVersion = currentVersion;
            Description = description;
            Enabled = enabled;
            EnrollmentCert = enrollmentCert;
            ExpectedUpgradeTime = expectedUpgradeTime;
            ExpectedVersion = expectedVersion;
            Fingerprint = fingerprint;
            Id = id;
            IpAcl = ipAcl;
            IssuedCertId = issuedCertId;
            LastBrokerConnectTime = lastBrokerConnectTime;
            LastBrokerConnectTimeDuration = lastBrokerConnectTimeDuration;
            LastBrokerDisconnectTime = lastBrokerDisconnectTime;
            LastBrokerDisconnectTimeDuration = lastBrokerDisconnectTimeDuration;
            LastUpgradeTime = lastUpgradeTime;
            Latitude = latitude;
            ListenIps = listenIps;
            Location = location;
            Longitude = longitude;
            ModifiedBy = modifiedBy;
            ModifiedTime = modifiedTime;
            Name = name;
            Platform = platform;
            PreviousVersion = previousVersion;
            PrivateIp = privateIp;
            ProvisioningKeyId = provisioningKeyId;
            ProvisioningKeyName = provisioningKeyName;
            PublicIp = publicIp;
            PublishIps = publishIps;
            SargeVersion = sargeVersion;
            ServiceEdgeGroupId = serviceEdgeGroupId;
            ServiceEdgeGroupName = serviceEdgeGroupName;
            UpgradeAttempt = upgradeAttempt;
            UpgradeStatus = upgradeStatus;
            ZpnSubModuleUpgradeLists = zpnSubModuleUpgradeLists;
        }
    }
}
