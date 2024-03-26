// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
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
    /**
     * The id of the CBI banner to be exported.
     */
    id?: string;
    /**
     * The name of the CBI banner to be exported.
     */
    name?: string;
}

/**
 * A collection of values returned by getCloudBrowserIsolationBanner.
 */
export interface GetCloudBrowserIsolationBannerResult {
    /**
     * (bool) - Show Welcome Notification
     */
    readonly banner: boolean;
    readonly id?: string;
    /**
     * (bool) - Use the default banner
     */
    readonly isDefault: boolean;
    /**
     * (string) - The Logo Image (.jpeg or .png; Maximum file size is 100KB.)
     */
    readonly logo: string;
    readonly name?: string;
    /**
     * (string) The Banner Notification Text
     */
    readonly notificationText: string;
    /**
     * (string) The Banner Notification Title
     */
    readonly notificationTitle: string;
    /**
     * (string) - The Banner Primary Color code in hexadecimal way to represent the color of the banner in RGB format
     */
    readonly primaryColor: string;
    /**
     * (string) - The Banner Text Color code in hexadecimal way to represent the color of the text in RGB format
     */
    readonly textColor: string;
}
/**
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
    /**
     * The id of the CBI banner to be exported.
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the CBI banner to be exported.
     */
    name?: pulumi.Input<string>;
}
