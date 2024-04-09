// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/** @deprecated zpa.index/getpracredentialcontroller.getPraCredentialController has been deprecated in favor of zpa.index/getpracredential.getPRACredential */
export function getPraCredentialController(args?: GetPraCredentialControllerArgs, opts?: pulumi.InvokeOptions): Promise<GetPraCredentialControllerResult> {
    pulumi.log.warn("getPraCredentialController is deprecated: zpa.index/getpracredentialcontroller.getPraCredentialController has been deprecated in favor of zpa.index/getpracredential.getPRACredential")
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getPraCredentialController:getPraCredentialController", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getPraCredentialController.
 */
export interface GetPraCredentialControllerArgs {
    id?: string;
    name?: string;
}

/**
 * A collection of values returned by getPraCredentialController.
 */
export interface GetPraCredentialControllerResult {
    readonly creationTime: string;
    readonly credentialType: string;
    readonly description: string;
    readonly id?: string;
    readonly lastCredentialResetTime: string;
    readonly microtenantId: string;
    readonly microtenantName: string;
    readonly modifiedBy: string;
    readonly modifiedTime: string;
    readonly name?: string;
    readonly userDomain: string;
    readonly username: string;
}
/** @deprecated zpa.index/getpracredentialcontroller.getPraCredentialController has been deprecated in favor of zpa.index/getpracredential.getPRACredential */
export function getPraCredentialControllerOutput(args?: GetPraCredentialControllerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPraCredentialControllerResult> {
    return pulumi.output(args).apply((a: any) => getPraCredentialController(a, opts))
}

/**
 * A collection of arguments for invoking getPraCredentialController.
 */
export interface GetPraCredentialControllerOutputArgs {
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}
