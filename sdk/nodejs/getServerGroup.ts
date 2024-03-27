// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example = zpa.getServerGroup({
 *     name: "server_group_name",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getServerGroup(args?: GetServerGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetServerGroupResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getServerGroup:getServerGroup", {
        "id": args.id,
        "microtenantId": args.microtenantId,
        "microtenantName": args.microtenantName,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getServerGroup.
 */
export interface GetServerGroupArgs {
    /**
     * The ID of the server group to be exported.
     */
    id?: string;
    /**
     * (string) The ID of the microtenant the resource is to be associated with.
     */
    microtenantId?: string;
    /**
     * (string) The name of the microtenant the resource is to be associated with.
     */
    microtenantName?: string;
    /**
     * The name of the server group to be exported.
     */
    name?: string;
}

/**
 * A collection of values returned by getServerGroup.
 */
export interface GetServerGroupResult {
    /**
     * (string)This field is a json array of app-connector-id only.
     */
    readonly appConnectorGroups: outputs.GetServerGroupAppConnectorGroup[];
    readonly applications: outputs.GetServerGroupApplication[];
    /**
     * (string)
     */
    readonly configSpace: string;
    readonly creationTime: string;
    /**
     * (string) This field is the description of the server group.
     */
    readonly description: string;
    /**
     * (bool) This field controls dynamic discovery of the servers.
     */
    readonly dynamicDiscovery: boolean;
    /**
     * (bool) This field defines if the server group is enabled or disabled.
     */
    readonly enabled: boolean;
    readonly id?: string;
    /**
     * (bool)
     */
    readonly ipAnchored: boolean;
    /**
     * (string) The ID of the microtenant the resource is to be associated with.
     */
    readonly microtenantId?: string;
    /**
     * (string) The name of the microtenant the resource is to be associated with.
     */
    readonly microtenantName?: string;
    readonly modifiedTime: string;
    readonly modifiedby: string;
    readonly name?: string;
    readonly servers: outputs.GetServerGroupServer[];
}
/**
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example = zpa.getServerGroup({
 *     name: "server_group_name",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getServerGroupOutput(args?: GetServerGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServerGroupResult> {
    return pulumi.output(args).apply((a: any) => getServerGroup(a, opts))
}

/**
 * A collection of arguments for invoking getServerGroup.
 */
export interface GetServerGroupOutputArgs {
    /**
     * The ID of the server group to be exported.
     */
    id?: pulumi.Input<string>;
    /**
     * (string) The ID of the microtenant the resource is to be associated with.
     */
    microtenantId?: pulumi.Input<string>;
    /**
     * (string) The name of the microtenant the resource is to be associated with.
     */
    microtenantName?: pulumi.Input<string>;
    /**
     * The name of the server group to be exported.
     */
    name?: pulumi.Input<string>;
}