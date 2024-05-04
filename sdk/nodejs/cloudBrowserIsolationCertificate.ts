// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/isolation/about-custom-root-certificates-cloud-browser-isolation)
 *
 * The **zpa_cloud_browser_isolation_certificate** resource creates a Cloud Browser Isolation certificate. This resource can then be used when creating a CBI External Profile `zpa.CloudBrowserIsolationExternalProfile`.`
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * // Retrieve CBI Banner ID
 * const thisCloudBrowserIsolationCertificate = new zpa.CloudBrowserIsolationCertificate("thisCloudBrowserIsolationCertificate", {pem: fs.readFileSync("cert.pem", "utf8")});
 * const thisIndex_cloudBrowserIsolationCertificateCloudBrowserIsolationCertificate = new zpa.CloudBrowserIsolationCertificate("thisIndex/cloudBrowserIsolationCertificateCloudBrowserIsolationCertificate", {pem: `    -----BEGIN CERTIFICATE-----
 *     MIIFYDCCBEigAwIBAgIQQAF3ITfU6UK47naqPGQKtzANBgkqhkiG9w0BAQsFADA/
 * `});
 * ```
 */
export class CloudBrowserIsolationCertificate extends pulumi.CustomResource {
    /**
     * Get an existing CloudBrowserIsolationCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CloudBrowserIsolationCertificateState, opts?: pulumi.CustomResourceOptions): CloudBrowserIsolationCertificate {
        return new CloudBrowserIsolationCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zpa:index/cloudBrowserIsolationCertificate:CloudBrowserIsolationCertificate';

    /**
     * Returns true if the given object is an instance of CloudBrowserIsolationCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CloudBrowserIsolationCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CloudBrowserIsolationCertificate.__pulumiType;
    }

    public readonly name!: pulumi.Output<string>;
    public readonly pem!: pulumi.Output<string | undefined>;

    /**
     * Create a CloudBrowserIsolationCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CloudBrowserIsolationCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CloudBrowserIsolationCertificateArgs | CloudBrowserIsolationCertificateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CloudBrowserIsolationCertificateState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pem"] = state ? state.pem : undefined;
        } else {
            const args = argsOrState as CloudBrowserIsolationCertificateArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pem"] = args ? args.pem : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CloudBrowserIsolationCertificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CloudBrowserIsolationCertificate resources.
 */
export interface CloudBrowserIsolationCertificateState {
    name?: pulumi.Input<string>;
    pem?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CloudBrowserIsolationCertificate resource.
 */
export interface CloudBrowserIsolationCertificateArgs {
    name?: pulumi.Input<string>;
    pem?: pulumi.Input<string>;
}
