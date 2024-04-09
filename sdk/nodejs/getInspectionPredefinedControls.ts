// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-custom-controls)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-appprotection-controls-using-api)
 *
 * Use the **zpa_inspection_predefined_controls** data source to get information about an OWASP predefined control and prefedined control version. This data source is required when creating an inspection profile.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example = zpa.getInspectionPredefinedControls({
 *     name: "Failed to parse request body",
 *     version: "OWASP_CRS/3.3.0",
 * });
 * export const zpaInspectionPredefinedControls = example;
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInspectionPredefinedControls(args?: GetInspectionPredefinedControlsArgs, opts?: pulumi.InvokeOptions): Promise<GetInspectionPredefinedControlsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getInspectionPredefinedControls:getInspectionPredefinedControls", {
        "id": args.id,
        "name": args.name,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getInspectionPredefinedControls.
 */
export interface GetInspectionPredefinedControlsArgs {
    id?: string;
    name?: string;
    version?: string;
}

/**
 * A collection of values returned by getInspectionPredefinedControls.
 */
export interface GetInspectionPredefinedControlsResult {
    readonly action: string;
    readonly actionValue: string;
    readonly associatedInspectionProfileNames: outputs.GetInspectionPredefinedControlsAssociatedInspectionProfileName[];
    readonly attachment: string;
    readonly controlGroup: string;
    readonly controlNumber: string;
    readonly controlType: string;
    readonly creationTime: string;
    readonly defaultAction: string;
    readonly defaultActionValue: string;
    readonly description: string;
    readonly id: string;
    readonly modifiedTime: string;
    readonly modifiedby: string;
    readonly name: string;
    readonly paranoiaLevel: string;
    readonly protocolType: string;
    readonly severity: string;
    readonly version?: string;
}
/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-custom-controls)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-appprotection-controls-using-api)
 *
 * Use the **zpa_inspection_predefined_controls** data source to get information about an OWASP predefined control and prefedined control version. This data source is required when creating an inspection profile.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example = zpa.getInspectionPredefinedControls({
 *     name: "Failed to parse request body",
 *     version: "OWASP_CRS/3.3.0",
 * });
 * export const zpaInspectionPredefinedControls = example;
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInspectionPredefinedControlsOutput(args?: GetInspectionPredefinedControlsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInspectionPredefinedControlsResult> {
    return pulumi.output(args).apply((a: any) => getInspectionPredefinedControls(a, opts))
}

/**
 * A collection of arguments for invoking getInspectionPredefinedControls.
 */
export interface GetInspectionPredefinedControlsOutputArgs {
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    version?: pulumi.Input<string>;
}
