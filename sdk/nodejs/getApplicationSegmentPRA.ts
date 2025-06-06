// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-privileged-remote-access-applications)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-application-segments-using-api)
 *
 * Use the **zpa_application_segment_pra** data source to get information about an application segment for Privileged Remote Access in the Zscaler Private Access cloud. This resource can then be referenced in an access policy rule, access policy timeout rule, access policy client forwarding rule and inspection policy. This resource supports Privileged Remote Access for both `RDP` and `SSH`.
 *
 * **NOTE:** To ensure consistent search results across data sources, please avoid using multiple spaces or special characters in your search queries.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const _this = zpa.getApplicationSegmentPRA({
 *     name: "PRA_Example",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const _this = zpa.getApplicationSegmentPRA({
 *     id: "123456789",
 * });
 * ```
 */
export function getApplicationSegmentPRA(args?: GetApplicationSegmentPRAArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationSegmentPRAResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getApplicationSegmentPRA:getApplicationSegmentPRA", {
        "id": args.id,
        "microtenantId": args.microtenantId,
        "microtenantName": args.microtenantName,
        "name": args.name,
        "tcpPortRange": args.tcpPortRange,
        "udpPortRange": args.udpPortRange,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplicationSegmentPRA.
 */
export interface GetApplicationSegmentPRAArgs {
    id?: string;
    microtenantId?: string;
    microtenantName?: string;
    name?: string;
    tcpPortRange?: inputs.GetApplicationSegmentPRATcpPortRange[];
    udpPortRange?: inputs.GetApplicationSegmentPRAUdpPortRange[];
}

/**
 * A collection of values returned by getApplicationSegmentPRA.
 */
export interface GetApplicationSegmentPRAResult {
    readonly bypassType: string;
    readonly configSpace: string;
    readonly description: string;
    readonly domainNames: string[];
    readonly doubleEncrypt: boolean;
    readonly enabled: boolean;
    readonly healthCheckType: string;
    readonly healthReporting: string;
    readonly id?: string;
    readonly ipAnchored: boolean;
    readonly isCnameEnabled: boolean;
    readonly microtenantId?: string;
    readonly microtenantName?: string;
    readonly name?: string;
    readonly passiveHealthEnabled: boolean;
    readonly segmentGroupId: string;
    readonly segmentGroupName: string;
    readonly serverGroups: outputs.GetApplicationSegmentPRAServerGroup[];
    readonly sraApps: outputs.GetApplicationSegmentPRASraApp[];
    readonly tcpPortRange: outputs.GetApplicationSegmentPRATcpPortRange[];
    readonly tcpPortRanges: string[];
    readonly udpPortRange: outputs.GetApplicationSegmentPRAUdpPortRange[];
    readonly udpPortRanges: string[];
}
/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-privileged-remote-access-applications)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-application-segments-using-api)
 *
 * Use the **zpa_application_segment_pra** data source to get information about an application segment for Privileged Remote Access in the Zscaler Private Access cloud. This resource can then be referenced in an access policy rule, access policy timeout rule, access policy client forwarding rule and inspection policy. This resource supports Privileged Remote Access for both `RDP` and `SSH`.
 *
 * **NOTE:** To ensure consistent search results across data sources, please avoid using multiple spaces or special characters in your search queries.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const _this = zpa.getApplicationSegmentPRA({
 *     name: "PRA_Example",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const _this = zpa.getApplicationSegmentPRA({
 *     id: "123456789",
 * });
 * ```
 */
export function getApplicationSegmentPRAOutput(args?: GetApplicationSegmentPRAOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetApplicationSegmentPRAResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("zpa:index/getApplicationSegmentPRA:getApplicationSegmentPRA", {
        "id": args.id,
        "microtenantId": args.microtenantId,
        "microtenantName": args.microtenantName,
        "name": args.name,
        "tcpPortRange": args.tcpPortRange,
        "udpPortRange": args.udpPortRange,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplicationSegmentPRA.
 */
export interface GetApplicationSegmentPRAOutputArgs {
    id?: pulumi.Input<string>;
    microtenantId?: pulumi.Input<string>;
    microtenantName?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    tcpPortRange?: pulumi.Input<pulumi.Input<inputs.GetApplicationSegmentPRATcpPortRangeArgs>[]>;
    udpPortRange?: pulumi.Input<pulumi.Input<inputs.GetApplicationSegmentPRAUdpPortRangeArgs>[]>;
}
