// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the **zpa_access_policy_platforms** data source to get information about all platforms for the specified customer in the Zscaler Private Access cloud. This data source can be optionally used when defining the following policy types:
 *     - ``zpa.PolicyAccessRule``
 *     - ``zpa.PolicyAccessTimeOutRule``
 *     - ``zpa.PolicyAccessForwardingRule``
 *     - ``zpa.PolicyAccessIsolationRule``
 *     - ``zpa.PolicyAccessInspectionRule``
 *
 * The ``objectType`` attribute must be defined as "PLATFORM" in the policy operand condition. To learn more see the To learn more see the [Getting Platform Types for a Customer](https://help.zscaler.com/zpa/configuring-access-policies-using-api#getPlatformTypes)
 *
 * > **NOTE** By Default the ZPA provider will return all platform types
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const this = zpa.getPolicyPlatform({});
 * ```
 */
export function getPolicyPlatform(opts?: pulumi.InvokeOptions): Promise<GetPolicyPlatformResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getPolicyPlatform:getPolicyPlatform", {
    }, opts);
}

/**
 * A collection of values returned by getPolicyPlatform.
 */
export interface GetPolicyPlatformResult {
    readonly android: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ios: string;
    readonly linux: string;
    readonly mac: string;
    readonly windows: string;
}
/**
 * Use the **zpa_access_policy_platforms** data source to get information about all platforms for the specified customer in the Zscaler Private Access cloud. This data source can be optionally used when defining the following policy types:
 *     - ``zpa.PolicyAccessRule``
 *     - ``zpa.PolicyAccessTimeOutRule``
 *     - ``zpa.PolicyAccessForwardingRule``
 *     - ``zpa.PolicyAccessIsolationRule``
 *     - ``zpa.PolicyAccessInspectionRule``
 *
 * The ``objectType`` attribute must be defined as "PLATFORM" in the policy operand condition. To learn more see the To learn more see the [Getting Platform Types for a Customer](https://help.zscaler.com/zpa/configuring-access-policies-using-api#getPlatformTypes)
 *
 * > **NOTE** By Default the ZPA provider will return all platform types
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const this = zpa.getPolicyPlatform({});
 * ```
 */
export function getPolicyPlatformOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetPolicyPlatformResult> {
    return pulumi.output(getPolicyPlatform(opts))
}
