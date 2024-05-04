// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/isolation/adding-banner-theme-isolation-end-user-notification-zpa)
 *
 * Use the **zpa_cloud_browser_isolation_banner** data source to get information about Cloud Browser Isolation banner. This data source information is required as part of the attribute `bannerId` when creating an Cloud Browser Isolation External Profile ``zpa.CloudBrowserIsolationExternalProfile``
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const this = zpa.getCloudBrowserIsolationBanner({
 *     name: "Default",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCloudBrowserIsolationBanner(args?: GetCloudBrowserIsolationBannerArgs, opts?: pulumi.InvokeOptions): Promise<GetCloudBrowserIsolationBannerResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getCloudBrowserIsolationBanner:getCloudBrowserIsolationBanner", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCloudBrowserIsolationBanner.
 */
export interface GetCloudBrowserIsolationBannerArgs {
    id?: string;
    name?: string;
}

/**
 * A collection of values returned by getCloudBrowserIsolationBanner.
 */
export interface GetCloudBrowserIsolationBannerResult {
    readonly banner: boolean;
    readonly id?: string;
    readonly isDefault: boolean;
    readonly logo: string;
    readonly name?: string;
    readonly notificationText: string;
    readonly notificationTitle: string;
    readonly primaryColor: string;
    readonly textColor: string;
}
/**
 * * [Official documentation](https://help.zscaler.com/isolation/adding-banner-theme-isolation-end-user-notification-zpa)
 *
 * Use the **zpa_cloud_browser_isolation_banner** data source to get information about Cloud Browser Isolation banner. This data source information is required as part of the attribute `bannerId` when creating an Cloud Browser Isolation External Profile ``zpa.CloudBrowserIsolationExternalProfile``
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const this = zpa.getCloudBrowserIsolationBanner({
 *     name: "Default",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCloudBrowserIsolationBannerOutput(args?: GetCloudBrowserIsolationBannerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCloudBrowserIsolationBannerResult> {
    return pulumi.output(args).apply((a: any) => getCloudBrowserIsolationBanner(a, opts))
}

/**
 * A collection of arguments for invoking getCloudBrowserIsolationBanner.
 */
export interface GetCloudBrowserIsolationBannerOutputArgs {
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}
