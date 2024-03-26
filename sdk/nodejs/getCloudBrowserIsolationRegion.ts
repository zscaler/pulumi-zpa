// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the **zpa_cloud_browser_isolation_region** data source to get information about Cloud Browser Isolation regions such as ID and Name. This data source information is required as part of the attribute `regionIds` when creating an Cloud Browser Isolation External Profile ``zpa.CloudBrowserIsolationExternalProfile``
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const this = zpa.getCloudBrowserIsolationRegion({
 *     name: "Singapore",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCloudBrowserIsolationRegion(args?: GetCloudBrowserIsolationRegionArgs, opts?: pulumi.InvokeOptions): Promise<GetCloudBrowserIsolationRegionResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getCloudBrowserIsolationRegion:getCloudBrowserIsolationRegion", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCloudBrowserIsolationRegion.
 */
export interface GetCloudBrowserIsolationRegionArgs {
    /**
     * The id of the CBI region to be exported.
     */
    id?: string;
    /**
     * The name of the CBI region to be exported.
     */
    name?: string;
}

/**
 * A collection of values returned by getCloudBrowserIsolationRegion.
 */
export interface GetCloudBrowserIsolationRegionResult {
    /**
     * (string) - ID information of the CBI region
     */
    readonly id: string;
    /**
     * (string) - Name of the CBI region
     */
    readonly name: string;
}
/**
 * Use the **zpa_cloud_browser_isolation_region** data source to get information about Cloud Browser Isolation regions such as ID and Name. This data source information is required as part of the attribute `regionIds` when creating an Cloud Browser Isolation External Profile ``zpa.CloudBrowserIsolationExternalProfile``
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const this = zpa.getCloudBrowserIsolationRegion({
 *     name: "Singapore",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCloudBrowserIsolationRegionOutput(args?: GetCloudBrowserIsolationRegionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCloudBrowserIsolationRegionResult> {
    return pulumi.output(args).apply((a: any) => getCloudBrowserIsolationRegion(a, opts))
}

/**
 * A collection of arguments for invoking getCloudBrowserIsolationRegion.
 */
export interface GetCloudBrowserIsolationRegionOutputArgs {
    /**
     * The id of the CBI region to be exported.
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the CBI region to be exported.
     */
    name?: pulumi.Input<string>;
}
