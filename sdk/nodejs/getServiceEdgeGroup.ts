// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-zpa-private-service-edge-groups)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-zpa-private-service-edge-groups-using-api)
 *
 * Use the **zpa_service_edge_group** data source to get information about a service edge group in the Zscaler Private Access cloud. This data source can then be referenced in an App Connector Group. This data source can then be referenced in the following resources:
 *
 * * Create a server group
 * * Provisioning Key
 * * Access policy rule
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const foo = zpa.getServiceEdgeGroup({
 *     name: "DataCenter",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const foo = zpa.getServiceEdgeGroup({
 *     id: "123456789",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getServiceEdgeGroup(args?: GetServiceEdgeGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceEdgeGroupResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getServiceEdgeGroup:getServiceEdgeGroup", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceEdgeGroup.
 */
export interface GetServiceEdgeGroupArgs {
    id?: string;
    name?: string;
}

/**
 * A collection of values returned by getServiceEdgeGroup.
 */
export interface GetServiceEdgeGroupResult {
    readonly cityCountry: string;
    readonly countryCode: string;
    readonly creationTime: string;
    readonly description: string;
    readonly enabled: boolean;
    readonly geoLocationId: string;
    readonly graceDistanceEnabled: boolean;
    readonly graceDistanceValue: string;
    readonly graceDistanceValueUnit: string;
    readonly id: string;
    readonly isPublic: string;
    readonly latitude: string;
    readonly location: string;
    readonly longitude: string;
    readonly modifiedTime: string;
    readonly modifiedby: string;
    readonly name: string;
    readonly overrideVersionProfile: boolean;
    readonly serviceEdges: outputs.GetServiceEdgeGroupServiceEdge[];
    readonly trustedNetworks: outputs.GetServiceEdgeGroupTrustedNetwork[];
    readonly upgradeDay: string;
    readonly upgradeTimeInSecs: string;
    readonly versionProfileId: string;
    readonly versionProfileName: string;
    readonly versionProfileVisibilityScope: string;
}
/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-zpa-private-service-edge-groups)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-zpa-private-service-edge-groups-using-api)
 *
 * Use the **zpa_service_edge_group** data source to get information about a service edge group in the Zscaler Private Access cloud. This data source can then be referenced in an App Connector Group. This data source can then be referenced in the following resources:
 *
 * * Create a server group
 * * Provisioning Key
 * * Access policy rule
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const foo = zpa.getServiceEdgeGroup({
 *     name: "DataCenter",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const foo = zpa.getServiceEdgeGroup({
 *     id: "123456789",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getServiceEdgeGroupOutput(args?: GetServiceEdgeGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceEdgeGroupResult> {
    return pulumi.output(args).apply((a: any) => getServiceEdgeGroup(a, opts))
}

/**
 * A collection of arguments for invoking getServiceEdgeGroup.
 */
export interface GetServiceEdgeGroupOutputArgs {
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}
