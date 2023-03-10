// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use the **zpa_idp_controller** data source to get information about an Identity Provider created in the Zscaler Private Access cloud. This data source is required when creating:
 *
 * 1. Access policy Rules
 * 2. Access policy timeout rules
 * 3. Access policy forwarding rules
 * 4. Access policy inspection rules
 * 5. Access policy isolation rules
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example = zpa.idpController.getIdPController({
 *     name: "idp_name",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example = zpa.idpController.getIdPController({
 *     id: "1234567890",
 * });
 * ```
 */
export function getIdPController(args?: GetIdPControllerArgs, opts?: pulumi.InvokeOptions): Promise<GetIdPControllerResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:idpController/getIdPController:getIdPController", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIdPController.
 */
export interface GetIdPControllerArgs {
    /**
     * The name of the Identity Provider (IdP) to be exported.
     */
    id?: string;
    /**
     * The name of the Identity Provider (IdP) to be exported.
     */
    name?: string;
}

/**
 * A collection of values returned by getIdPController.
 */
export interface GetIdPControllerResult {
    /**
     * (Computed)
     */
    readonly adminMetadatas: outputs.idpController.GetIdPControllerAdminMetadata[];
    readonly adminSpSigningCertId: string;
    /**
     * (string)
     */
    readonly autoProvision: string;
    /**
     * (string)
     */
    readonly creationTime: string;
    /**
     * (string)
     */
    readonly description: string;
    /**
     * (bool)
     */
    readonly disableSamlBasedPolicy: boolean;
    /**
     * (string)
     */
    readonly domainLists: string[];
    /**
     * (bool)
     */
    readonly enableScimBasedPolicy: boolean;
    /**
     * (bool) Default value if null is True
     */
    readonly enabled: boolean;
    readonly id: string;
    /**
     * (string)
     */
    readonly idpEntityId: string;
    /**
     * (string)
     */
    readonly loginNameAttribute: string;
    /**
     * (string)
     */
    readonly loginUrl: string;
    /**
     * (string)
     */
    readonly modifiedTime: string;
    readonly modifiedby: string;
    readonly name: string;
    /**
     * (bool)
     */
    readonly reauthOnUserUpdate: boolean;
    /**
     * (bool)
     */
    readonly redirectBinding: boolean;
    /**
     * (bool)
     */
    readonly scimEnabled: boolean;
    /**
     * (string)
     */
    readonly scimServiceProviderEndpoint: string;
    /**
     * (bool)
     */
    readonly scimSharedSecretExists: boolean;
    /**
     * (string)
     */
    readonly signSamlRequest: string;
    /**
     * (string)
     */
    readonly ssoTypes: string[];
    /**
     * (bool)
     */
    readonly useCustomSpMetadata: boolean;
    /**
     * (Computed)
     */
    readonly userMetadatas: outputs.idpController.GetIdPControllerUserMetadata[];
    readonly userSpSigningCertId: string;
}
/**
 * Use the **zpa_idp_controller** data source to get information about an Identity Provider created in the Zscaler Private Access cloud. This data source is required when creating:
 *
 * 1. Access policy Rules
 * 2. Access policy timeout rules
 * 3. Access policy forwarding rules
 * 4. Access policy inspection rules
 * 5. Access policy isolation rules
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example = zpa.idpController.getIdPController({
 *     name: "idp_name",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example = zpa.idpController.getIdPController({
 *     id: "1234567890",
 * });
 * ```
 */
export function getIdPControllerOutput(args?: GetIdPControllerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIdPControllerResult> {
    return pulumi.output(args).apply((a: any) => getIdPController(a, opts))
}

/**
 * A collection of arguments for invoking getIdPController.
 */
export interface GetIdPControllerOutputArgs {
    /**
     * The name of the Identity Provider (IdP) to be exported.
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the Identity Provider (IdP) to be exported.
     */
    name?: pulumi.Input<string>;
}
