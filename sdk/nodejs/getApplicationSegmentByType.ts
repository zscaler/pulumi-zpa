// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-applications)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-application-segments-using-api)
 *
 * Use the **zpa_application_segment_by_type** data source to get all configured Application Segments by Access Type (e.g., ``BROWSER_ACCESS``, ``INSPECT``, or ``SECURE_REMOTE_ACCESS``) for the specified customer.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const _this = zpa.getApplicationSegmentByType({
 *     applicationType: "SECURE_REMOTE_ACCESS",
 * });
 * ```
 */
export function getApplicationSegmentByType(args: GetApplicationSegmentByTypeArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationSegmentByTypeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getApplicationSegmentByType:getApplicationSegmentByType", {
        "applicationType": args.applicationType,
        "microtenantId": args.microtenantId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplicationSegmentByType.
 */
export interface GetApplicationSegmentByTypeArgs {
    applicationType: string;
    microtenantId?: string;
    name?: string;
}

/**
 * A collection of values returned by getApplicationSegmentByType.
 */
export interface GetApplicationSegmentByTypeResult {
    readonly appId: string;
    readonly applicationPort: string;
    readonly applicationProtocol: string;
    readonly applicationType: string;
    readonly certificateId: string;
    readonly certificateName: string;
    readonly domain: string;
    readonly enabled: boolean;
    readonly id: string;
    readonly microtenantId?: string;
    readonly microtenantName: string;
    readonly name?: string;
}
/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-applications)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-application-segments-using-api)
 *
 * Use the **zpa_application_segment_by_type** data source to get all configured Application Segments by Access Type (e.g., ``BROWSER_ACCESS``, ``INSPECT``, or ``SECURE_REMOTE_ACCESS``) for the specified customer.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const _this = zpa.getApplicationSegmentByType({
 *     applicationType: "SECURE_REMOTE_ACCESS",
 * });
 * ```
 */
export function getApplicationSegmentByTypeOutput(args: GetApplicationSegmentByTypeOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetApplicationSegmentByTypeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("zpa:index/getApplicationSegmentByType:getApplicationSegmentByType", {
        "applicationType": args.applicationType,
        "microtenantId": args.microtenantId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplicationSegmentByType.
 */
export interface GetApplicationSegmentByTypeOutputArgs {
    applicationType: pulumi.Input<string>;
    microtenantId?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}
