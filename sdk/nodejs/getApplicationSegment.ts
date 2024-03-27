// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const this = zpa.getApplicationSegment({
 *     name: "example",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const this = zpa.getApplicationSegment({
 *     id: "123456789",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getApplicationSegment(args?: GetApplicationSegmentArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationSegmentResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getApplicationSegment:getApplicationSegment", {
        "id": args.id,
        "isIncompleteDrConfig": args.isIncompleteDrConfig,
        "microtenantId": args.microtenantId,
        "microtenantName": args.microtenantName,
        "name": args.name,
        "tcpPortRange": args.tcpPortRange,
        "udpPortRange": args.udpPortRange,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplicationSegment.
 */
export interface GetApplicationSegmentArgs {
    id?: string;
    /**
     * Default: false. Boolean: `true`, `false`.
     */
    isIncompleteDrConfig?: boolean;
    /**
     * The ID of the microtenant the resource is to be associated with.
     */
    microtenantId?: string;
    /**
     * The name of the microtenant the resource is to be associated with.
     */
    microtenantName?: string;
    /**
     * Name of the application.
     */
    name?: string;
    tcpPortRange?: inputs.GetApplicationSegmentTcpPortRange[];
    udpPortRange?: inputs.GetApplicationSegmentUdpPortRange[];
}

/**
 * A collection of values returned by getApplicationSegment.
 */
export interface GetApplicationSegmentResult {
    readonly bypassType: string;
    readonly configSpace: string;
    readonly creationTime: string;
    readonly defaultIdleTimeout: string;
    readonly defaultMaxAge: string;
    readonly description: string;
    readonly domainNames: string[];
    readonly doubleEncrypt: boolean;
    readonly enabled: boolean;
    readonly healthCheckType: string;
    readonly healthReporting: string;
    readonly id?: string;
    readonly ipAnchored: boolean;
    readonly isCnameEnabled: boolean;
    readonly isIncompleteDrConfig: boolean;
    readonly microtenantId?: string;
    readonly microtenantName?: string;
    readonly modifiedTime: string;
    readonly modifiedby: string;
    readonly name?: string;
    readonly passiveHealthEnabled: boolean;
    readonly segmentGroupId: string;
    readonly segmentGroupName: string;
    readonly selectConnectorCloseToApp: boolean;
    readonly serverGroups: outputs.GetApplicationSegmentServerGroup[];
    readonly tcpPortRange: outputs.GetApplicationSegmentTcpPortRange[];
    readonly tcpPortRanges: string[];
    readonly udpPortRange: outputs.GetApplicationSegmentUdpPortRange[];
    readonly udpPortRanges: string[];
    readonly useInDrMode: boolean;
}
/**
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const this = zpa.getApplicationSegment({
 *     name: "example",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const this = zpa.getApplicationSegment({
 *     id: "123456789",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getApplicationSegmentOutput(args?: GetApplicationSegmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationSegmentResult> {
    return pulumi.output(args).apply((a: any) => getApplicationSegment(a, opts))
}

/**
 * A collection of arguments for invoking getApplicationSegment.
 */
export interface GetApplicationSegmentOutputArgs {
    id?: pulumi.Input<string>;
    /**
     * Default: false. Boolean: `true`, `false`.
     */
    isIncompleteDrConfig?: pulumi.Input<boolean>;
    /**
     * The ID of the microtenant the resource is to be associated with.
     */
    microtenantId?: pulumi.Input<string>;
    /**
     * The name of the microtenant the resource is to be associated with.
     */
    microtenantName?: pulumi.Input<string>;
    /**
     * Name of the application.
     */
    name?: pulumi.Input<string>;
    tcpPortRange?: pulumi.Input<pulumi.Input<inputs.GetApplicationSegmentTcpPortRangeArgs>[]>;
    udpPortRange?: pulumi.Input<pulumi.Input<inputs.GetApplicationSegmentUdpPortRangeArgs>[]>;
}