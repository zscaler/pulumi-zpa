// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the **zpa_cloud_browser_isolation_certificate** data source to get information about Cloud Browser Isolation Certificate. This data source information is required as part of the attribute `certificateIds` when creating an Cloud Browser Isolation External Profile ``zpa.CloudBrowserIsolationExternalProfile``
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const this = zpa.getCloudBrowserIsolationCertificate({
 *     name: "Zscaler Root Certificate",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCloudBrowserIsolationCertificate(args?: GetCloudBrowserIsolationCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetCloudBrowserIsolationCertificateResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getCloudBrowserIsolationCertificate:getCloudBrowserIsolationCertificate", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCloudBrowserIsolationCertificate.
 */
export interface GetCloudBrowserIsolationCertificateArgs {
    /**
     * The id of the CBI certificate to be exported.
     */
    id?: string;
    /**
     * The name of the CBI certificate to be exported.
     */
    name?: string;
}

/**
 * A collection of values returned by getCloudBrowserIsolationCertificate.
 */
export interface GetCloudBrowserIsolationCertificateResult {
    readonly id?: string;
    readonly isDefault: boolean;
    readonly name?: string;
    readonly pem: string;
}
/**
 * Use the **zpa_cloud_browser_isolation_certificate** data source to get information about Cloud Browser Isolation Certificate. This data source information is required as part of the attribute `certificateIds` when creating an Cloud Browser Isolation External Profile ``zpa.CloudBrowserIsolationExternalProfile``
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const this = zpa.getCloudBrowserIsolationCertificate({
 *     name: "Zscaler Root Certificate",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCloudBrowserIsolationCertificateOutput(args?: GetCloudBrowserIsolationCertificateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCloudBrowserIsolationCertificateResult> {
    return pulumi.output(args).apply((a: any) => getCloudBrowserIsolationCertificate(a, opts))
}

/**
 * A collection of arguments for invoking getCloudBrowserIsolationCertificate.
 */
export interface GetCloudBrowserIsolationCertificateOutputArgs {
    /**
     * The id of the CBI certificate to be exported.
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the CBI certificate to be exported.
     */
    name?: pulumi.Input<string>;
}
