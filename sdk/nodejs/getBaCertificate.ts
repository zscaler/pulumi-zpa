// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-web-server-certificates)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-certificates-using-api)
 *
 * Use the **zpa_ba_certificate** data source to get information about a browser access certificate created in the Zscaler Private Access cloud. This data source is required when creating a browser access application segment resource.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const foo = zpa.getBaCertificate({
 *     name: "example.acme.com",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const foo = zpa.getBaCertificate({
 *     id: "1234567890",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getBaCertificate(args?: GetBaCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetBaCertificateResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getBaCertificate:getBaCertificate", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getBaCertificate.
 */
export interface GetBaCertificateArgs {
    id?: string;
    name?: string;
}

/**
 * A collection of values returned by getBaCertificate.
 */
export interface GetBaCertificateResult {
    readonly certChain: string;
    readonly certificate: string;
    readonly cname: string;
    readonly creationTime: string;
    readonly description: string;
    readonly id?: string;
    readonly issuedBy: string;
    readonly issuedTo: string;
    readonly microtenantId: string;
    readonly modifiedTime: string;
    readonly modifiedby: string;
    readonly name?: string;
    readonly publicKey: string;
    readonly sans: string[];
    readonly serialNo: string;
    readonly status: string;
    readonly validFromInEpochsec: string;
    readonly validToInEpochsec: string;
}
/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-web-server-certificates)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-certificates-using-api)
 *
 * Use the **zpa_ba_certificate** data source to get information about a browser access certificate created in the Zscaler Private Access cloud. This data source is required when creating a browser access application segment resource.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const foo = zpa.getBaCertificate({
 *     name: "example.acme.com",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const foo = zpa.getBaCertificate({
 *     id: "1234567890",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getBaCertificateOutput(args?: GetBaCertificateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBaCertificateResult> {
    return pulumi.output(args).apply((a: any) => getBaCertificate(a, opts))
}

/**
 * A collection of arguments for invoking getBaCertificate.
 */
export interface GetBaCertificateOutputArgs {
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}
