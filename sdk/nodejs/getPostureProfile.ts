// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the **zpa_posture_profile** data source to get information about a posture profile created in the Zscaler Private Access Mobile Portal. This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example1 = zpa.getPostureProfile({
 *     name: "CrowdStrike_ZPA_ZTA_40",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example2 = zpa.getPostureProfile({
 *     name: "Detect SentinelOne",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example3 = zpa.getPostureProfile({
 *     name: "domain_joined",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * > **NOTE** To query posture profiles that are associated with a specific Zscaler cloud, it is required to append the cloud name to the name of the posture profile as the below example:
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example1 = zpa.getPostureProfile({
 *     name: "CrowdStrike_ZPA_ZTA_40 (zscalertwo.net)",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * > **NOTE** When associating a posture profile with one of supported resources, the following parameter must be exported: ``postureUdid`` instead of the ``id`` of the resource.
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example1 = zpa.getPostureProfile({
 *     name: "CrowdStrike_ZPA_ZTA_40 (zscalertwo.net)",
 * });
 * export const zpaPostureProfile = example1.then(example1 => example1.postureUdid);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getPostureProfile(args?: GetPostureProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetPostureProfileResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getPostureProfile:getPostureProfile", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getPostureProfile.
 */
export interface GetPostureProfileArgs {
    /**
     * The name of the posture profile to be exported.
     */
    name?: string;
}

/**
 * A collection of values returned by getPostureProfile.
 */
export interface GetPostureProfileResult {
    /**
     * (string)
     */
    readonly creationTime: string;
    /**
     * (string)
     */
    readonly domain: string;
    readonly id: string;
    /**
     * (string)
     */
    readonly masterCustomerId: string;
    /**
     * (string)
     */
    readonly modifiedTime: string;
    readonly modifiedby: string;
    readonly name?: string;
    /**
     * (string)
     */
    readonly postureUdid: string;
    /**
     * (string)
     */
    readonly zscalerCloud: string;
    /**
     * (string)
     */
    readonly zscalerCustomerId: string;
}
/**
 * Use the **zpa_posture_profile** data source to get information about a posture profile created in the Zscaler Private Access Mobile Portal. This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example1 = zpa.getPostureProfile({
 *     name: "CrowdStrike_ZPA_ZTA_40",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example2 = zpa.getPostureProfile({
 *     name: "Detect SentinelOne",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example3 = zpa.getPostureProfile({
 *     name: "domain_joined",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * > **NOTE** To query posture profiles that are associated with a specific Zscaler cloud, it is required to append the cloud name to the name of the posture profile as the below example:
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example1 = zpa.getPostureProfile({
 *     name: "CrowdStrike_ZPA_ZTA_40 (zscalertwo.net)",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * > **NOTE** When associating a posture profile with one of supported resources, the following parameter must be exported: ``postureUdid`` instead of the ``id`` of the resource.
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example1 = zpa.getPostureProfile({
 *     name: "CrowdStrike_ZPA_ZTA_40 (zscalertwo.net)",
 * });
 * export const zpaPostureProfile = example1.then(example1 => example1.postureUdid);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getPostureProfileOutput(args?: GetPostureProfileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPostureProfileResult> {
    return pulumi.output(args).apply((a: any) => getPostureProfile(a, opts))
}

/**
 * A collection of arguments for invoking getPostureProfile.
 */
export interface GetPostureProfileOutputArgs {
    /**
     * The name of the posture profile to be exported.
     */
    name?: pulumi.Input<string>;
}