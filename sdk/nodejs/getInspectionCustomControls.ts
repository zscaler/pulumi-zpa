// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-custom-controls)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-appprotection-controls-using-api)
 *
 * Use the **zpa_inspection_custom_controls** data source to get information about an inspection custom control. This data source can be associated with an inspection profile.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const example = zpa.getInspectionCustomControls({
 *     name: "ZPA_Inspection_Custom_Control",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const example = zpa.getInspectionCustomControls({
 *     id: "1234567890",
 * });
 * ```
 */
export function getInspectionCustomControls(args?: GetInspectionCustomControlsArgs, opts?: pulumi.InvokeOptions): Promise<GetInspectionCustomControlsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getInspectionCustomControls:getInspectionCustomControls", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getInspectionCustomControls.
 */
export interface GetInspectionCustomControlsArgs {
    id?: string;
    name?: string;
}

/**
 * A collection of values returned by getInspectionCustomControls.
 */
export interface GetInspectionCustomControlsResult {
    readonly action: string;
    readonly actionValue: string;
    readonly controlNumber: string;
    readonly controlRuleJson: string;
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
    readonly rules: outputs.GetInspectionCustomControlsRule[];
    readonly severity: string;
    readonly type: string;
    readonly version: string;
}
/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-custom-controls)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-appprotection-controls-using-api)
 *
 * Use the **zpa_inspection_custom_controls** data source to get information about an inspection custom control. This data source can be associated with an inspection profile.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const example = zpa.getInspectionCustomControls({
 *     name: "ZPA_Inspection_Custom_Control",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const example = zpa.getInspectionCustomControls({
 *     id: "1234567890",
 * });
 * ```
 */
export function getInspectionCustomControlsOutput(args?: GetInspectionCustomControlsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetInspectionCustomControlsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("zpa:index/getInspectionCustomControls:getInspectionCustomControls", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getInspectionCustomControls.
 */
export interface GetInspectionCustomControlsOutputArgs {
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}
