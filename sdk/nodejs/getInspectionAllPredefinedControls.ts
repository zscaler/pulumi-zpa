// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use the **zpa_inspection_all_predefined_controls** data source to get information about all OWASP predefined control and prefedined control version by group name. The `Preprocessors` predefined control is the default predefined control, This data source is always required, when creating an inspection profile.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const this = zpa.getInspectionAllPredefinedControls({
 *     groupName: "Preprocessors",
 *     version: "OWASP_CRS/3.3.0",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInspectionAllPredefinedControls(args: GetInspectionAllPredefinedControlsArgs, opts?: pulumi.InvokeOptions): Promise<GetInspectionAllPredefinedControlsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getInspectionAllPredefinedControls:getInspectionAllPredefinedControls", {
        "groupName": args.groupName,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getInspectionAllPredefinedControls.
 */
export interface GetInspectionAllPredefinedControlsArgs {
    /**
     * The name of the predefined control.
     */
    groupName?: string;
    /**
     * The version of the predefined control, the default is: `OWASP_CRS/3.3.0`
     */
    version: string;
}

/**
 * A collection of values returned by getInspectionAllPredefinedControls.
 */
export interface GetInspectionAllPredefinedControlsResult {
    readonly groupName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lists: outputs.GetInspectionAllPredefinedControlsList[];
    readonly version: string;
}
/**
 * Use the **zpa_inspection_all_predefined_controls** data source to get information about all OWASP predefined control and prefedined control version by group name. The `Preprocessors` predefined control is the default predefined control, This data source is always required, when creating an inspection profile.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const this = zpa.getInspectionAllPredefinedControls({
 *     groupName: "Preprocessors",
 *     version: "OWASP_CRS/3.3.0",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInspectionAllPredefinedControlsOutput(args: GetInspectionAllPredefinedControlsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInspectionAllPredefinedControlsResult> {
    return pulumi.output(args).apply((a: any) => getInspectionAllPredefinedControls(a, opts))
}

/**
 * A collection of arguments for invoking getInspectionAllPredefinedControls.
 */
export interface GetInspectionAllPredefinedControlsOutputArgs {
    /**
     * The name of the predefined control.
     */
    groupName?: pulumi.Input<string>;
    /**
     * The version of the predefined control, the default is: `OWASP_CRS/3.3.0`
     */
    version: pulumi.Input<string>;
}