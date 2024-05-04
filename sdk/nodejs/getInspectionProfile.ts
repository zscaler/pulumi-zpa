// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-browser-protection-profiles)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-appprotection-profiles-using-api)
 *
 * Use the **zpa_inspection_profile** data source to get information about an inspection profile in the Zscaler Private Access cloud. This resource can then be referenced in an inspection custom control resource.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const this = zpa.getInspectionProfile({
 *     name: "Example",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInspectionProfile(args?: GetInspectionProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetInspectionProfileResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getInspectionProfile:getInspectionProfile", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getInspectionProfile.
 */
export interface GetInspectionProfileArgs {
    id?: string;
    /**
     * - (String) This field defines the name of the inspection profile.
     */
    name?: string;
}

/**
 * A collection of values returned by getInspectionProfile.
 */
export interface GetInspectionProfileResult {
    readonly commonGlobalOverrideActionsConfig: {[key: string]: string};
    readonly controlsInfos: outputs.GetInspectionProfileControlsInfo[];
    readonly creationTime: string;
    readonly customControls: outputs.GetInspectionProfileCustomControl[];
    readonly description: string;
    readonly globalControlActions: string[];
    readonly id: string;
    readonly incarnationNumber: string;
    readonly modifiedBy: string;
    readonly modifiedTime: string;
    /**
     * - (String) This field defines the name of the inspection profile.
     */
    readonly name: string;
    readonly paranoiaLevel: string;
    readonly predefinedControls: outputs.GetInspectionProfilePredefinedControl[];
    readonly predefinedControlsVersion: string;
    readonly webSocketControls: outputs.GetInspectionProfileWebSocketControl[];
}
/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-browser-protection-profiles)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-appprotection-profiles-using-api)
 *
 * Use the **zpa_inspection_profile** data source to get information about an inspection profile in the Zscaler Private Access cloud. This resource can then be referenced in an inspection custom control resource.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const this = zpa.getInspectionProfile({
 *     name: "Example",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInspectionProfileOutput(args?: GetInspectionProfileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInspectionProfileResult> {
    return pulumi.output(args).apply((a: any) => getInspectionProfile(a, opts))
}

/**
 * A collection of arguments for invoking getInspectionProfile.
 */
export interface GetInspectionProfileOutputArgs {
    id?: pulumi.Input<string>;
    /**
     * - (String) This field defines the name of the inspection profile.
     */
    name?: pulumi.Input<string>;
}
