// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/isolation/about-custom-root-certificates-cloud-browser-isolation)
 *
 * Use the **zpa_cloud_browser_isolation_external_profile** data source to get information about Cloud Browser Isolation external profile. This data source information can then be used in as part of `zpa.PolicyAccessIsolationRule` when the `action` attribute is set to `ISOLATE`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const this = zpa.getCloudBrowserIsolationExternalProfile({
 *     name: "Example",
 * });
 * ```
 */
export function getCloudBrowserIsolationExternalProfile(args?: GetCloudBrowserIsolationExternalProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetCloudBrowserIsolationExternalProfileResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getCloudBrowserIsolationExternalProfile:getCloudBrowserIsolationExternalProfile", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCloudBrowserIsolationExternalProfile.
 */
export interface GetCloudBrowserIsolationExternalProfileArgs {
    id?: string;
    name?: string;
}

/**
 * A collection of values returned by getCloudBrowserIsolationExternalProfile.
 */
export interface GetCloudBrowserIsolationExternalProfileResult {
    readonly bannerId: string;
    readonly certificateIds: string[];
    readonly debugModes: outputs.GetCloudBrowserIsolationExternalProfileDebugMode[];
    readonly description: string;
    readonly href: string;
    readonly id?: string;
    readonly isDefault: boolean;
    readonly name?: string;
    readonly regions: outputs.GetCloudBrowserIsolationExternalProfileRegion[];
    readonly securityControls: outputs.GetCloudBrowserIsolationExternalProfileSecurityControl[];
    readonly userExperiences: outputs.GetCloudBrowserIsolationExternalProfileUserExperience[];
}
/**
 * * [Official documentation](https://help.zscaler.com/isolation/about-custom-root-certificates-cloud-browser-isolation)
 *
 * Use the **zpa_cloud_browser_isolation_external_profile** data source to get information about Cloud Browser Isolation external profile. This data source information can then be used in as part of `zpa.PolicyAccessIsolationRule` when the `action` attribute is set to `ISOLATE`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const this = zpa.getCloudBrowserIsolationExternalProfile({
 *     name: "Example",
 * });
 * ```
 */
export function getCloudBrowserIsolationExternalProfileOutput(args?: GetCloudBrowserIsolationExternalProfileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCloudBrowserIsolationExternalProfileResult> {
    return pulumi.output(args).apply((a: any) => getCloudBrowserIsolationExternalProfile(a, opts))
}

/**
 * A collection of arguments for invoking getCloudBrowserIsolationExternalProfile.
 */
export interface GetCloudBrowserIsolationExternalProfileOutputArgs {
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}
