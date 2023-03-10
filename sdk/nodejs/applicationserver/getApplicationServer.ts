// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example = zpa.ApplicationServer.getApplicationServer({
 *     name: "server.example.com",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example = zpa.ApplicationServer.getApplicationServer({
 *     id: "1234567890",
 * });
 * ```
 */
export function getApplicationServer(args?: GetApplicationServerArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationServerResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:ApplicationServer/getApplicationServer:getApplicationServer", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplicationServer.
 */
export interface GetApplicationServerArgs {
    /**
     * This field defines the id of the application server.
     */
    id?: string;
    /**
     * This field defines the name of the server.
     */
    name?: string;
}

/**
 * A collection of values returned by getApplicationServer.
 */
export interface GetApplicationServerResult {
    /**
     * (string) This field defines the domain or IP address of the server.
     */
    readonly address: string;
    /**
     * (Set of String) This field defines the list of server groups IDs.
     */
    readonly appServerGroupIds: string[];
    readonly configSpace: string;
    readonly creationTime: string;
    /**
     * (string) This field defines the description of the server.
     */
    readonly description: string;
    /**
     * (bool) This field defines the status of the server.
     */
    readonly enabled: boolean;
    readonly id?: string;
    readonly modifiedTime: string;
    readonly modifiedby: string;
    readonly name?: string;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example = zpa.ApplicationServer.getApplicationServer({
 *     name: "server.example.com",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example = zpa.ApplicationServer.getApplicationServer({
 *     id: "1234567890",
 * });
 * ```
 */
export function getApplicationServerOutput(args?: GetApplicationServerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationServerResult> {
    return pulumi.output(args).apply((a: any) => getApplicationServer(a, opts))
}

/**
 * A collection of arguments for invoking getApplicationServer.
 */
export interface GetApplicationServerOutputArgs {
    /**
     * This field defines the id of the application server.
     */
    id?: pulumi.Input<string>;
    /**
     * This field defines the name of the server.
     */
    name?: pulumi.Input<string>;
}
