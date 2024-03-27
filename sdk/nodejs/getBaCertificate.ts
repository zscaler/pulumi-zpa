// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
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
    /**
     * The id of the browser access certificate to be exported.
     */
    id?: string;
    /**
     * The name of the browser access certificate to be exported.
     */
    name?: string;
}

/**
 * A collection of values returned by getBaCertificate.
 */
export interface GetBaCertificateResult {
    /**
     * (string)
     */
    readonly certChain: string;
    /**
     * (string) The certificate text is in PEM format.
     */
    readonly certificate: string;
    /**
     * (string)
     */
    readonly cname: string;
    /**
     * (string)
     */
    readonly creationTime: string;
    /**
     * (string)
     */
    readonly description: string;
    readonly id?: string;
    /**
     * (string)
     */
    readonly issuedBy: string;
    /**
     * (string)
     */
    readonly issuedTo: string;
    readonly microtenantId: string;
    /**
     * (string)
     */
    readonly modifiedTime: string;
    /**
     * (string)
     */
    readonly modifiedby: string;
    readonly name?: string;
    readonly publicKey: string;
    /**
     * (string)
     */
    readonly sans: string[];
    /**
     * (string)
     */
    readonly serialNo: string;
    /**
     * (string)
     */
    readonly status: string;
    /**
     * (string)
     */
    readonly validFromInEpochsec: string;
    /**
     * (string)
     */
    readonly validToInEpochsec: string;
}
/**
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
    /**
     * The id of the browser access certificate to be exported.
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the browser access certificate to be exported.
     */
    name?: pulumi.Input<string>;
}