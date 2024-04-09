// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-privileged-credentials)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-privileged-credentials-using-api)
 *
 * The **zpa_pra_credential_controller** resource creates a privileged remote access credential in the Zscaler Private Access cloud. This resource can then be referenced in an privileged access policy resource.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * // Retrieves PRA Credential By ID
 * const _this = new zpa.PRACredential("this", {});
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getPRACredential(args?: GetPRACredentialArgs, opts?: pulumi.InvokeOptions): Promise<GetPRACredentialResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getPRACredential:getPRACredential", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getPRACredential.
 */
export interface GetPRACredentialArgs {
    id?: string;
    /**
     * - (String) The name of the privileged credential.
     */
    name?: string;
}

/**
 * A collection of values returned by getPRACredential.
 */
export interface GetPRACredentialResult {
    readonly creationTime: string;
    readonly credentialType: string;
    readonly description: string;
    readonly id?: string;
    readonly lastCredentialResetTime: string;
    readonly microtenantId: string;
    readonly microtenantName: string;
    readonly modifiedBy: string;
    readonly modifiedTime: string;
    /**
     * - (String) The name of the privileged credential.
     */
    readonly name?: string;
    readonly userDomain: string;
    readonly username: string;
}
/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-privileged-credentials)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-privileged-credentials-using-api)
 *
 * The **zpa_pra_credential_controller** resource creates a privileged remote access credential in the Zscaler Private Access cloud. This resource can then be referenced in an privileged access policy resource.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * // Retrieves PRA Credential By ID
 * const _this = new zpa.PRACredential("this", {});
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getPRACredentialOutput(args?: GetPRACredentialOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPRACredentialResult> {
    return pulumi.output(args).apply((a: any) => getPRACredential(a, opts))
}

/**
 * A collection of arguments for invoking getPRACredential.
 */
export interface GetPRACredentialOutputArgs {
    id?: pulumi.Input<string>;
    /**
     * - (String) The name of the privileged credential.
     */
    name?: pulumi.Input<string>;
}
