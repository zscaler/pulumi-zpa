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
    public static class GetAppConnectorController
    {
        /// <summary>
        /// Use the **zpa_app_connector_controller** data source to get information about a app connector created in the Zscaler Private Access cloud. This data source can then be referenced in an App Connector Group.
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
        ///     var example = Zpa.GetAppConnectorController.Invoke(new()
        ///     {
        ///         Name = "AWS-VPC100-App-Connector",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetAppConnectorControllerResult> InvokeAsync(GetAppConnectorControllerArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAppConnectorControllerResult>("zpa:index/getAppConnectorController:getAppConnectorController", args ?? new GetAppConnectorControllerArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zpa_app_connector_controller** data source to get information about a app connector created in the Zscaler Private Access cloud. This data source can then be referenced in an App Connector Group.
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
        ///     var example = Zpa.GetAppConnectorController.Invoke(new()
        ///     {
        ///         Name = "AWS-VPC100-App-Connector",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetAppConnectorControllerResult> Invoke(GetAppConnectorControllerInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAppConnectorControllerResult>("zpa:index/getAppConnectorController:getAppConnectorController", args ?? new GetAppConnectorControllerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppConnectorControllerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// (string) The ID of the microtenant the resource is to be associated with.
        /// </summary>
        [Input("microtenantId")]
        public string? MicrotenantId { get; set; }

        /// <summary>
        /// (string) The name of the microtenant the resource is to be associated with.
        /// </summary>
        [Input("microtenantName")]
        public string? MicrotenantName { get; set; }

        /// <summary>
        /// Name of the App Connector Group.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetAppConnectorControllerArgs()
        {
        }
        public static new GetAppConnectorControllerArgs Empty => new GetAppConnectorControllerArgs();
    }

    public sealed class GetAppConnectorControllerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// (string) The ID of the microtenant the resource is to be associated with.
        /// </summary>
        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// (string) The name of the microtenant the resource is to be associated with.
        /// </summary>
        [Input("microtenantName")]
        public Input<string>? MicrotenantName { get; set; }

        /// <summary>
        /// Name of the App Connector Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetAppConnectorControllerInvokeArgs()
        {
        }
        public static new GetAppConnectorControllerInvokeArgs Empty => new GetAppConnectorControllerInvokeArgs();
    }


    [OutputType]
    public sealed class GetAppConnectorControllerResult
    {
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string AppConnectorGroupId;
        /// <summary>
        /// (Computed) - Expected values: UNKNOWN/ZPN_STATUS_AUTHENTICATED(1)/ZPN_STATUS_DISCONNECTED
        /// </summary>
        public readonly string AppConnectorGroupName;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string ApplicationStartTime;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string ControlChannelStatus;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string CtrlBrokerName;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string CurrentVersion;
        /// <summary>
        /// (Computed) - Description of the App Connector.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (Computed) - Whether this App Connector is enabled or not. Default value: `true`. Supported values: `true`, `false`
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly ImmutableDictionary<string, object> EnrollmentCert;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string ExpectedUpgradeTime;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string ExpectedVersion;
        public readonly string Fingerprint;
        public readonly string Id;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string IpAcl;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string IssuedCertId;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string LastBrokerConnectTime;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string LastBrokerConnectTimeDuration;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string LastBrokerDisconnectTime;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string LastBrokerDisconnectTimeDuration;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string LastUpgradeTime;
        /// <summary>
        /// (Computed) - Latitude of the App Connector. Integer or decimal. With values in the range of `-90` to `90`
        /// </summary>
        public readonly string Latitude;
        /// <summary>
        /// (Computed) - Location of the App Connector.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// (Computed) - Longitude of the App Connector. Integer or decimal. With values in the range of `-180` to `180`
        /// </summary>
        public readonly string Longitude;
        /// <summary>
        /// (string) The ID of the microtenant the resource is to be associated with.
        /// </summary>
        public readonly string? MicrotenantId;
        /// <summary>
        /// (string) The name of the microtenant the resource is to be associated with.
        /// </summary>
        public readonly string? MicrotenantName;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        public readonly string? Name;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string Platform;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string PreviousVersion;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string PrivateIp;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string ProvisioningKeyId;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string ProvisioningKeyName;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string PublicIp;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string SargeVersion;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string UpgradeAttempt;
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly string UpgradeStatus;

        [OutputConstructor]
        private GetAppConnectorControllerResult(
            string appConnectorGroupId,

            string appConnectorGroupName,

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

            string location,

            string longitude,

            string? microtenantId,

            string? microtenantName,

            string modifiedTime,

            string modifiedby,

            string? name,

            string platform,

            string previousVersion,

            string privateIp,

            string provisioningKeyId,

            string provisioningKeyName,

            string publicIp,

            string sargeVersion,

            string upgradeAttempt,

            string upgradeStatus)
        {
            AppConnectorGroupId = appConnectorGroupId;
            AppConnectorGroupName = appConnectorGroupName;
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
            Location = location;
            Longitude = longitude;
            MicrotenantId = microtenantId;
            MicrotenantName = microtenantName;
            ModifiedTime = modifiedTime;
            Modifiedby = modifiedby;
            Name = name;
            Platform = platform;
            PreviousVersion = previousVersion;
            PrivateIp = privateIp;
            ProvisioningKeyId = provisioningKeyId;
            ProvisioningKeyName = provisioningKeyName;
            PublicIp = publicIp;
            SargeVersion = sargeVersion;
            UpgradeAttempt = upgradeAttempt;
            UpgradeStatus = upgradeStatus;
        }
    }
}