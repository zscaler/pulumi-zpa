// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/isolation/creating-isolation-profiles-zpa)
 * * [API documentation](https://help.zscaler.com/zpa/obtaining-isolation-profile-details-using-api)
 *
 * Use the **zpa_isolation_profile** data source to get information about an isolation profile in the Zscaler Private Access cloud. This data source is required when configuring an isolation policy rule resource
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const isolationProfile = zpa.getIsolationProfile({
 *     name: "zpa_isolation_profile",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getIsolationProfile(args?: GetIsolationProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetIsolationProfileResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getIsolationProfile:getIsolationProfile", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIsolationProfile.
 */
export interface GetIsolationProfileArgs {
    /**
     * - (Required) This field defines the name of the isolation profile.
     */
    name?: string;
}

/**
 * A collection of values returned by getIsolationProfile.
 */
export interface GetIsolationProfileResult {
    readonly creationTime: string;
    readonly description: string;
    readonly enabled: boolean;
    readonly id: string;
    readonly isolationProfileId: string;
    readonly isolationTenantId: string;
    readonly isolationUrl: string;
    readonly modifiedTime: string;
    readonly modifiedby: string;
    /**
     * - (Required) This field defines the name of the isolation profile.
     */
    readonly name?: string;
}
/**
 * * [Official documentation](https://help.zscaler.com/isolation/creating-isolation-profiles-zpa)
 * * [API documentation](https://help.zscaler.com/zpa/obtaining-isolation-profile-details-using-api)
 *
 * Use the **zpa_isolation_profile** data source to get information about an isolation profile in the Zscaler Private Access cloud. This data source is required when configuring an isolation policy rule resource
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const isolationProfile = zpa.getIsolationProfile({
 *     name: "zpa_isolation_profile",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getIsolationProfileOutput(args?: GetIsolationProfileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIsolationProfileResult> {
    return pulumi.output(args).apply((a: any) => getIsolationProfile(a, opts))
}

/**
 * A collection of arguments for invoking getIsolationProfile.
 */
export interface GetIsolationProfileOutputArgs {
    /**
     * - (Required) This field defines the name of the isolation profile.
     */
    name?: pulumi.Input<string>;
}
