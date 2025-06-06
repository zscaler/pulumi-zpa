// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zpa/identity-management)
 * * [API documentation](https://help.zscaler.com/zpa/obtaining-idp-configuration-details-using-api)
 *
 * Use the **zpa_idp_controller** data source to get information about an Identity Provider created in the Zscaler Private Access cloud. This data source is required when creating:
 *
 * **NOTE:** To ensure consistent search results across data sources, please avoid using multiple spaces or special characters in your search queries.
 *
 * * Access policy Rules
 * * Access policy timeout rules
 * * Access policy forwarding rules
 * * Access policy inspection rules
 * * Access policy isolation rules
 * * Access policy privileged credentials rules
 * * Access policy privileged capabilities rules
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const example = zpa.getIdPController({
 *     name: "idp_name",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const example = zpa.getIdPController({
 *     id: "1234567890",
 * });
 * ```
 */
export function getIdPController(args?: GetIdPControllerArgs, opts?: pulumi.InvokeOptions): Promise<GetIdPControllerResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getIdPController:getIdPController", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIdPController.
 */
export interface GetIdPControllerArgs {
    id?: string;
    name?: string;
}

/**
 * A collection of values returned by getIdPController.
 */
export interface GetIdPControllerResult {
    readonly adminMetadatas: outputs.GetIdPControllerAdminMetadata[];
    readonly adminSpSigningCertId: string;
    readonly autoProvision: string;
    readonly creationTime: string;
    readonly description: string;
    readonly disableSamlBasedPolicy: boolean;
    readonly domainLists: string[];
    readonly enableArbitraryAuthDomains: string;
    readonly enableScimBasedPolicy: boolean;
    readonly enabled: boolean;
    readonly forceAuth: boolean;
    readonly id: string;
    readonly idpEntityId: string;
    readonly loginHint: boolean;
    readonly loginNameAttribute: string;
    readonly loginUrl: string;
    readonly modifiedTime: string;
    readonly modifiedby: string;
    readonly name: string;
    readonly reauthOnUserUpdate: boolean;
    readonly redirectBinding: boolean;
    readonly scimEnabled: boolean;
    readonly scimServiceProviderEndpoint: string;
    readonly scimSharedSecretExists: boolean;
    readonly signSamlRequest: string;
    readonly ssoTypes: string[];
    readonly useCustomSpMetadata: boolean;
    readonly userMetadatas: outputs.GetIdPControllerUserMetadata[];
    readonly userSpSigningCertId: string;
}
/**
 * * [Official documentation](https://help.zscaler.com/zpa/identity-management)
 * * [API documentation](https://help.zscaler.com/zpa/obtaining-idp-configuration-details-using-api)
 *
 * Use the **zpa_idp_controller** data source to get information about an Identity Provider created in the Zscaler Private Access cloud. This data source is required when creating:
 *
 * **NOTE:** To ensure consistent search results across data sources, please avoid using multiple spaces or special characters in your search queries.
 *
 * * Access policy Rules
 * * Access policy timeout rules
 * * Access policy forwarding rules
 * * Access policy inspection rules
 * * Access policy isolation rules
 * * Access policy privileged credentials rules
 * * Access policy privileged capabilities rules
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const example = zpa.getIdPController({
 *     name: "idp_name",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const example = zpa.getIdPController({
 *     id: "1234567890",
 * });
 * ```
 */
export function getIdPControllerOutput(args?: GetIdPControllerOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetIdPControllerResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("zpa:index/getIdPController:getIdPController", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIdPController.
 */
export interface GetIdPControllerOutputArgs {
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}
