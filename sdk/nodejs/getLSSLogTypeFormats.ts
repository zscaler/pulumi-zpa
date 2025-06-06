// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-log-streaming-service)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-log-streaming-service-configurations-using-api)
 *
 * Use the **zpa_lss_config_log_type_formats** data source to get information about all LSS log type formats in the Zscaler Private Access cloud. This data source is required when creating an LSS Config Controller resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const zpnTransLog = zpa.getLSSLogTypeFormats({
 *     logType: "zpn_trans_log",
 * });
 * const zpnAuthLog = zpa.getLSSLogTypeFormats({
 *     logType: "zpn_auth_log",
 * });
 * const zpnAstAuthLog = zpa.getLSSLogTypeFormats({
 *     logType: "zpn_ast_auth_log",
 * });
 * const zpnHttpTransLog = zpa.getLSSLogTypeFormats({
 *     logType: "zpn_http_trans_log",
 * });
 * const zpnAuditLog = zpa.getLSSLogTypeFormats({
 *     logType: "zpn_audit_log",
 * });
 * const zpnSysAuthLog = zpa.getLSSLogTypeFormats({
 *     logType: "zpn_sys_auth_log",
 * });
 * const zpnAstComprehensiveStats = zpa.getLSSLogTypeFormats({
 *     logType: "zpn_ast_comprehensive_stats",
 * });
 * const zpnWafHttpExchangesLog = zpa.getLSSLogTypeFormats({
 *     logType: "zpn_waf_http_exchanges_log",
 * });
 * const zpnPbrokerComprehensiveStats = zpa.getLSSLogTypeFormats({
 *     logType: "zpn_pbroker_comprehensive_stats",
 * });
 * ```
 *
 * ### Read-Only
 *
 * The following arguments are supported:
 *
 * * `logType` - (Required) The type of log to be exported.
 *   * `zpnTransLog`
 *   * `zpnAuthLog`
 *   * `zpnAstAuthLog`
 *   * `zpnHttpTransLog`
 *   * `zpnAuditLog`
 *   * `zpnSysAuthLog`
 *   * `zpnAstComprehensiveStats`
 *   * `zpnWafHttpExchangesLog`
 *   * `zpnPbrokerComprehensiveStats`
 */
export function getLSSLogTypeFormats(args: GetLSSLogTypeFormatsArgs, opts?: pulumi.InvokeOptions): Promise<GetLSSLogTypeFormatsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getLSSLogTypeFormats:getLSSLogTypeFormats", {
        "logType": args.logType,
    }, opts);
}

/**
 * A collection of arguments for invoking getLSSLogTypeFormats.
 */
export interface GetLSSLogTypeFormatsArgs {
    logType: string;
}

/**
 * A collection of values returned by getLSSLogTypeFormats.
 */
export interface GetLSSLogTypeFormatsResult {
    readonly csv: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly json: string;
    readonly logType: string;
    readonly tsv: string;
}
/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-log-streaming-service)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-log-streaming-service-configurations-using-api)
 *
 * Use the **zpa_lss_config_log_type_formats** data source to get information about all LSS log type formats in the Zscaler Private Access cloud. This data source is required when creating an LSS Config Controller resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const zpnTransLog = zpa.getLSSLogTypeFormats({
 *     logType: "zpn_trans_log",
 * });
 * const zpnAuthLog = zpa.getLSSLogTypeFormats({
 *     logType: "zpn_auth_log",
 * });
 * const zpnAstAuthLog = zpa.getLSSLogTypeFormats({
 *     logType: "zpn_ast_auth_log",
 * });
 * const zpnHttpTransLog = zpa.getLSSLogTypeFormats({
 *     logType: "zpn_http_trans_log",
 * });
 * const zpnAuditLog = zpa.getLSSLogTypeFormats({
 *     logType: "zpn_audit_log",
 * });
 * const zpnSysAuthLog = zpa.getLSSLogTypeFormats({
 *     logType: "zpn_sys_auth_log",
 * });
 * const zpnAstComprehensiveStats = zpa.getLSSLogTypeFormats({
 *     logType: "zpn_ast_comprehensive_stats",
 * });
 * const zpnWafHttpExchangesLog = zpa.getLSSLogTypeFormats({
 *     logType: "zpn_waf_http_exchanges_log",
 * });
 * const zpnPbrokerComprehensiveStats = zpa.getLSSLogTypeFormats({
 *     logType: "zpn_pbroker_comprehensive_stats",
 * });
 * ```
 *
 * ### Read-Only
 *
 * The following arguments are supported:
 *
 * * `logType` - (Required) The type of log to be exported.
 *   * `zpnTransLog`
 *   * `zpnAuthLog`
 *   * `zpnAstAuthLog`
 *   * `zpnHttpTransLog`
 *   * `zpnAuditLog`
 *   * `zpnSysAuthLog`
 *   * `zpnAstComprehensiveStats`
 *   * `zpnWafHttpExchangesLog`
 *   * `zpnPbrokerComprehensiveStats`
 */
export function getLSSLogTypeFormatsOutput(args: GetLSSLogTypeFormatsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetLSSLogTypeFormatsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("zpa:index/getLSSLogTypeFormats:getLSSLogTypeFormats", {
        "logType": args.logType,
    }, opts);
}

/**
 * A collection of arguments for invoking getLSSLogTypeFormats.
 */
export interface GetLSSLogTypeFormatsOutputArgs {
    logType: pulumi.Input<string>;
}
