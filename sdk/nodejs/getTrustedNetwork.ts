// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/client-connector/about-trusted-networks)
 * * [API documentation](https://help.zscaler.com/zpa/obtaining-trusted-network-details-using-api)
 *
 * The **zpa_trusted_network** data source to get information about a trusted network created in the Zscaler Private Access Mobile Portal. This data source can then be referenced within the following resources:
 *
 * **NOTE:** To ensure consistent search results across data sources, please avoid using multiple spaces or special characters in your search queries.
 *
 * 1. Access Policy
 * 2. Forwarding Policy
 * 3. Inspection Policy
 * 4. Isolation Policy
 * 5. Service Edge Group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const example = zpa.getTrustedNetwork({
 *     name: "trusted_network_name",
 * });
 * ```
 *
 * > **NOTE** To query trusted network that are associated with a specific Zscaler cloud, it is required to append the cloud name to the name of the trusted network as the below example:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const example1 = zpa.getTrustedNetwork({
 *     name: "Corporate-Network (zscalertwo.net)",
 * });
 * export const zpaTrustedNetwork = example1.then(example1 => example1.networkId);
 * ```
 */
export function getTrustedNetwork(args?: GetTrustedNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetTrustedNetworkResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getTrustedNetwork:getTrustedNetwork", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getTrustedNetwork.
 */
export interface GetTrustedNetworkArgs {
    id?: string;
    name?: string;
}

/**
 * A collection of values returned by getTrustedNetwork.
 */
export interface GetTrustedNetworkResult {
    readonly creationTime: string;
    readonly domain: string;
    readonly id?: string;
    readonly modifiedTime: string;
    readonly modifiedby: string;
    readonly name?: string;
    readonly networkId: string;
    readonly zscalerCloud: string;
}
/**
 * * [Official documentation](https://help.zscaler.com/client-connector/about-trusted-networks)
 * * [API documentation](https://help.zscaler.com/zpa/obtaining-trusted-network-details-using-api)
 *
 * The **zpa_trusted_network** data source to get information about a trusted network created in the Zscaler Private Access Mobile Portal. This data source can then be referenced within the following resources:
 *
 * **NOTE:** To ensure consistent search results across data sources, please avoid using multiple spaces or special characters in your search queries.
 *
 * 1. Access Policy
 * 2. Forwarding Policy
 * 3. Inspection Policy
 * 4. Isolation Policy
 * 5. Service Edge Group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const example = zpa.getTrustedNetwork({
 *     name: "trusted_network_name",
 * });
 * ```
 *
 * > **NOTE** To query trusted network that are associated with a specific Zscaler cloud, it is required to append the cloud name to the name of the trusted network as the below example:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const example1 = zpa.getTrustedNetwork({
 *     name: "Corporate-Network (zscalertwo.net)",
 * });
 * export const zpaTrustedNetwork = example1.then(example1 => example1.networkId);
 * ```
 */
export function getTrustedNetworkOutput(args?: GetTrustedNetworkOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetTrustedNetworkResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("zpa:index/getTrustedNetwork:getTrustedNetwork", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getTrustedNetwork.
 */
export interface GetTrustedNetworkOutputArgs {
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}
