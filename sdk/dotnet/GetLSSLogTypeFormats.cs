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
    public static class GetLSSLogTypeFormats
    {
        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/about-log-streaming-service)
        /// * [API documentation](https://help.zscaler.com/zpa/configuring-log-streaming-service-configurations-using-api)
        /// 
        /// Use the **zpa_lss_config_log_type_formats** data source to get information about all LSS log type formats in the Zscaler Private Access cloud. This data source is required when creating an LSS Config Controller resource.
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
        ///     var zpnTransLog = Zpa.GetLSSLogTypeFormats.Invoke(new()
        ///     {
        ///         LogType = "zpn_trans_log",
        ///     });
        /// 
        ///     var zpnAuthLog = Zpa.GetLSSLogTypeFormats.Invoke(new()
        ///     {
        ///         LogType = "zpn_auth_log",
        ///     });
        /// 
        ///     var zpnAstAuthLog = Zpa.GetLSSLogTypeFormats.Invoke(new()
        ///     {
        ///         LogType = "zpn_ast_auth_log",
        ///     });
        /// 
        ///     var zpnHttpTransLog = Zpa.GetLSSLogTypeFormats.Invoke(new()
        ///     {
        ///         LogType = "zpn_http_trans_log",
        ///     });
        /// 
        ///     var zpnAuditLog = Zpa.GetLSSLogTypeFormats.Invoke(new()
        ///     {
        ///         LogType = "zpn_audit_log",
        ///     });
        /// 
        ///     var zpnSysAuthLog = Zpa.GetLSSLogTypeFormats.Invoke(new()
        ///     {
        ///         LogType = "zpn_sys_auth_log",
        ///     });
        /// 
        ///     var zpnAstComprehensiveStats = Zpa.GetLSSLogTypeFormats.Invoke(new()
        ///     {
        ///         LogType = "zpn_ast_comprehensive_stats",
        ///     });
        /// 
        ///     var zpnWafHttpExchangesLog = Zpa.GetLSSLogTypeFormats.Invoke(new()
        ///     {
        ///         LogType = "zpn_waf_http_exchanges_log",
        ///     });
        /// 
        ///     var zpnPbrokerComprehensiveStats = Zpa.GetLSSLogTypeFormats.Invoke(new()
        ///     {
        ///         LogType = "zpn_pbroker_comprehensive_stats",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// ### Read-Only
        /// 
        /// The following arguments are supported:
        /// 
        /// * `log_type` - (Required) The type of log to be exported.
        ///   * `zpn_trans_log`
        ///   * `zpn_auth_log`
        ///   * `zpn_ast_auth_log`
        ///   * `zpn_http_trans_log`
        ///   * `zpn_audit_log`
        ///   * `zpn_sys_auth_log`
        ///   * `zpn_ast_comprehensive_stats`
        ///   * `zpn_waf_http_exchanges_log`
        ///   * `zpn_pbroker_comprehensive_stats`
        /// </summary>
        public static Task<GetLSSLogTypeFormatsResult> InvokeAsync(GetLSSLogTypeFormatsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLSSLogTypeFormatsResult>("zpa:index/getLSSLogTypeFormats:getLSSLogTypeFormats", args ?? new GetLSSLogTypeFormatsArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zpa/about-log-streaming-service)
        /// * [API documentation](https://help.zscaler.com/zpa/configuring-log-streaming-service-configurations-using-api)
        /// 
        /// Use the **zpa_lss_config_log_type_formats** data source to get information about all LSS log type formats in the Zscaler Private Access cloud. This data source is required when creating an LSS Config Controller resource.
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
        ///     var zpnTransLog = Zpa.GetLSSLogTypeFormats.Invoke(new()
        ///     {
        ///         LogType = "zpn_trans_log",
        ///     });
        /// 
        ///     var zpnAuthLog = Zpa.GetLSSLogTypeFormats.Invoke(new()
        ///     {
        ///         LogType = "zpn_auth_log",
        ///     });
        /// 
        ///     var zpnAstAuthLog = Zpa.GetLSSLogTypeFormats.Invoke(new()
        ///     {
        ///         LogType = "zpn_ast_auth_log",
        ///     });
        /// 
        ///     var zpnHttpTransLog = Zpa.GetLSSLogTypeFormats.Invoke(new()
        ///     {
        ///         LogType = "zpn_http_trans_log",
        ///     });
        /// 
        ///     var zpnAuditLog = Zpa.GetLSSLogTypeFormats.Invoke(new()
        ///     {
        ///         LogType = "zpn_audit_log",
        ///     });
        /// 
        ///     var zpnSysAuthLog = Zpa.GetLSSLogTypeFormats.Invoke(new()
        ///     {
        ///         LogType = "zpn_sys_auth_log",
        ///     });
        /// 
        ///     var zpnAstComprehensiveStats = Zpa.GetLSSLogTypeFormats.Invoke(new()
        ///     {
        ///         LogType = "zpn_ast_comprehensive_stats",
        ///     });
        /// 
        ///     var zpnWafHttpExchangesLog = Zpa.GetLSSLogTypeFormats.Invoke(new()
        ///     {
        ///         LogType = "zpn_waf_http_exchanges_log",
        ///     });
        /// 
        ///     var zpnPbrokerComprehensiveStats = Zpa.GetLSSLogTypeFormats.Invoke(new()
        ///     {
        ///         LogType = "zpn_pbroker_comprehensive_stats",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// ### Read-Only
        /// 
        /// The following arguments are supported:
        /// 
        /// * `log_type` - (Required) The type of log to be exported.
        ///   * `zpn_trans_log`
        ///   * `zpn_auth_log`
        ///   * `zpn_ast_auth_log`
        ///   * `zpn_http_trans_log`
        ///   * `zpn_audit_log`
        ///   * `zpn_sys_auth_log`
        ///   * `zpn_ast_comprehensive_stats`
        ///   * `zpn_waf_http_exchanges_log`
        ///   * `zpn_pbroker_comprehensive_stats`
        /// </summary>
        public static Output<GetLSSLogTypeFormatsResult> Invoke(GetLSSLogTypeFormatsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLSSLogTypeFormatsResult>("zpa:index/getLSSLogTypeFormats:getLSSLogTypeFormats", args ?? new GetLSSLogTypeFormatsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLSSLogTypeFormatsArgs : global::Pulumi.InvokeArgs
    {
        [Input("logType", required: true)]
        public string LogType { get; set; } = null!;

        public GetLSSLogTypeFormatsArgs()
        {
        }
        public static new GetLSSLogTypeFormatsArgs Empty => new GetLSSLogTypeFormatsArgs();
    }

    public sealed class GetLSSLogTypeFormatsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("logType", required: true)]
        public Input<string> LogType { get; set; } = null!;

        public GetLSSLogTypeFormatsInvokeArgs()
        {
        }
        public static new GetLSSLogTypeFormatsInvokeArgs Empty => new GetLSSLogTypeFormatsInvokeArgs();
    }


    [OutputType]
    public sealed class GetLSSLogTypeFormatsResult
    {
        public readonly string Csv;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Json;
        public readonly string LogType;
        public readonly string Tsv;

        [OutputConstructor]
        private GetLSSLogTypeFormatsResult(
            string csv,

            string id,

            string json,

            string logType,

            string tsv)
        {
            Csv = csv;
            Id = id;
            Json = json;
            LogType = logType;
            Tsv = tsv;
        }
    }
}
