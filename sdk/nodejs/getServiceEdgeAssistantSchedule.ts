// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getServiceEdgeAssistantSchedule(args?: GetServiceEdgeAssistantScheduleArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceEdgeAssistantScheduleResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getServiceEdgeAssistantSchedule:getServiceEdgeAssistantSchedule", {
        "customerId": args.customerId,
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceEdgeAssistantSchedule.
 */
export interface GetServiceEdgeAssistantScheduleArgs {
    customerId?: string;
    id?: string;
}

/**
 * A collection of values returned by getServiceEdgeAssistantSchedule.
 */
export interface GetServiceEdgeAssistantScheduleResult {
    readonly customerId?: string;
    readonly deleteDisabled: boolean;
    readonly enabled: boolean;
    readonly frequency: string;
    readonly frequencyInterval: string;
    readonly id?: string;
}
export function getServiceEdgeAssistantScheduleOutput(args?: GetServiceEdgeAssistantScheduleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceEdgeAssistantScheduleResult> {
    return pulumi.output(args).apply((a: any) => getServiceEdgeAssistantSchedule(a, opts))
}

/**
 * A collection of arguments for invoking getServiceEdgeAssistantSchedule.
 */
export interface GetServiceEdgeAssistantScheduleOutputArgs {
    customerId?: pulumi.Input<string>;
    id?: pulumi.Input<string>;
}