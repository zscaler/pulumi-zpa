// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use the **zpa_lss_config_controller** data source to get information about a Log Streaming (LSS) configuration resource created in the Zscaler Private Access.
 */
export function getLSSConfigController(args?: GetLSSConfigControllerArgs, opts?: pulumi.InvokeOptions): Promise<GetLSSConfigControllerResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:LSSConfig/getLSSConfigController:getLSSConfigController", {
        "configs": args.configs,
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getLSSConfigController.
 */
export interface GetLSSConfigControllerArgs {
    /**
     * (Computed)
     */
    configs?: inputs.LSSConfig.GetLSSConfigControllerConfig[];
    /**
     * This field defines the name of the log streaming resource.
     */
    id?: string;
}

/**
 * A collection of values returned by getLSSConfigController.
 */
export interface GetLSSConfigControllerResult {
    /**
     * (Computed)
     */
    readonly configs: outputs.LSSConfig.GetLSSConfigControllerConfig[];
    /**
     * (Computed)
     */
    readonly connectorGroups: outputs.LSSConfig.GetLSSConfigControllerConnectorGroup[];
    /**
     * (string)
     */
    readonly id?: string;
    readonly policyRules: outputs.LSSConfig.GetLSSConfigControllerPolicyRule[];
}
/**
 * Use the **zpa_lss_config_controller** data source to get information about a Log Streaming (LSS) configuration resource created in the Zscaler Private Access.
 */
export function getLSSConfigControllerOutput(args?: GetLSSConfigControllerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLSSConfigControllerResult> {
    return pulumi.output(args).apply((a: any) => getLSSConfigController(a, opts))
}

/**
 * A collection of arguments for invoking getLSSConfigController.
 */
export interface GetLSSConfigControllerOutputArgs {
    /**
     * (Computed)
     */
    configs?: pulumi.Input<pulumi.Input<inputs.LSSConfig.GetLSSConfigControllerConfigArgs>[]>;
    /**
     * This field defines the name of the log streaming resource.
     */
    id?: pulumi.Input<string>;
}
