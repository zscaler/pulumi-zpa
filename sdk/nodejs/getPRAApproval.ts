// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use the **zpa_pra_approval_controller** data source to get information about a privileged remote access approval created in the Zscaler Private Access cloud.
 */
export function getPRAApproval(args?: GetPRAApprovalArgs, opts?: pulumi.InvokeOptions): Promise<GetPRAApprovalResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getPRAApproval:getPRAApproval", {
        "emailIds": args.emailIds,
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getPRAApproval.
 */
export interface GetPRAApprovalArgs {
    emailIds?: string[];
    id?: string;
}

/**
 * A collection of values returned by getPRAApproval.
 */
export interface GetPRAApprovalResult {
    readonly applications: outputs.GetPRAApprovalApplication[];
    readonly creationTime: string;
    readonly emailIds?: string[];
    readonly endTime: string;
    readonly id?: string;
    /**
     * (string)  The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass microtenantId as 0 when making requests to retrieve data from the Default Microtenant. Pass microtenantId as null to retrieve data from all customers associated with the tenant.
     */
    readonly microtenantId: string;
    readonly modifiedBy: string;
    readonly modifiedTime: string;
    readonly startTime: string;
    /**
     * (string) The status of the privileged approval. The supported values are:
     */
    readonly status: string;
    readonly workingHours: outputs.GetPRAApprovalWorkingHour[];
}
/**
 * Use the **zpa_pra_approval_controller** data source to get information about a privileged remote access approval created in the Zscaler Private Access cloud.
 */
export function getPRAApprovalOutput(args?: GetPRAApprovalOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPRAApprovalResult> {
    return pulumi.output(args).apply((a: any) => getPRAApproval(a, opts))
}

/**
 * A collection of arguments for invoking getPRAApproval.
 */
export interface GetPRAApprovalOutputArgs {
    emailIds?: pulumi.Input<pulumi.Input<string>[]>;
    id?: pulumi.Input<string>;
}