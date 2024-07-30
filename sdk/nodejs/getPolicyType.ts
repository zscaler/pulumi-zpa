// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-access-policy)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-access-policies-using-api)
 *
 * Use the **zpa_policy_type** data source to get information about an a ``policySetId`` and ``policyType``. This data source is required when creating:
 *
 * 1. Access policy Rules
 * 2. Access policy timeout rules
 * 3. Access policy forwarding rules
 * 4. Access policy inspection rules
 *
 * > **NOTE** The parameters ``policySetId`` is required in all circumstances and is exported when checking for the policyType parameter. The policyType value is used for differentiating the policy types, in the request endpoint. The supported values are:
 *
 * * ``ACCESS_POLICY/GLOBAL_POLICY``
 * * ``TIMEOUT_POLICY/REAUTH_POLICY``
 * * ``BYPASS_POLICY/CLIENT_FORWARDING_POLICY``
 * * ``INSPECTION_POLICY``
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const accessPolicy = zpa.getPolicyType({
 *     policyType: "ACCESS_POLICY",
 * });
 * export const zpaPolicyTypeAccessPolicy = accessPolicy.then(accessPolicy => accessPolicy.id);
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const globalPolicy = zpa.getPolicyType({
 *     policyType: "GLOBAL_POLICY",
 * });
 * export const zpaPolicyTypeAccessPolicy = globalPolicy.then(globalPolicy => globalPolicy.id);
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const timeoutPolicy = zpa.getPolicyType({
 *     policyType: "TIMEOUT_POLICY",
 * });
 * export const zpaPolicyTypeTimeoutPolicy = timeoutPolicy.then(timeoutPolicy => timeoutPolicy.id);
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const reauthPolicy = zpa.getPolicyType({
 *     policyType: "REAUTH_POLICY",
 * });
 * export const zpaPolicyTypeReauthPolicy = reauthPolicy.then(reauthPolicy => reauthPolicy.id);
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const clientForwardingPolicy = zpa.getPolicyType({
 *     policyType: "CLIENT_FORWARDING_POLICY",
 * });
 * export const zpaPolicyTypeClientForwardingPolicy = clientForwardingPolicy.then(clientForwardingPolicy => clientForwardingPolicy.id);
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const inspectionPolicy = zpa.getPolicyType({
 *     policyType: "INSPECTION_POLICY",
 * });
 * export const zpaPolicyTypeInspectionPolicy = inspectionPolicy.then(inspectionPolicy => inspectionPolicy.id);
 * ```
 */
export function getPolicyType(args?: GetPolicyTypeArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyTypeResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getPolicyType:getPolicyType", {
        "id": args.id,
        "microtenantId": args.microtenantId,
        "microtenantName": args.microtenantName,
        "policyType": args.policyType,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicyType.
 */
export interface GetPolicyTypeArgs {
    id?: string;
    microtenantId?: string;
    microtenantName?: string;
    policyType?: string;
}

/**
 * A collection of values returned by getPolicyType.
 */
export interface GetPolicyTypeResult {
    readonly creationTime: string;
    readonly description: string;
    readonly enabled: boolean;
    readonly id?: string;
    readonly microtenantId?: string;
    readonly microtenantName?: string;
    readonly modifiedBy: string;
    readonly modifiedTime: string;
    readonly name: string;
    readonly policyType: string;
    readonly rules: outputs.GetPolicyTypeRule[];
    readonly sorted: boolean;
}
/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-access-policy)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-access-policies-using-api)
 *
 * Use the **zpa_policy_type** data source to get information about an a ``policySetId`` and ``policyType``. This data source is required when creating:
 *
 * 1. Access policy Rules
 * 2. Access policy timeout rules
 * 3. Access policy forwarding rules
 * 4. Access policy inspection rules
 *
 * > **NOTE** The parameters ``policySetId`` is required in all circumstances and is exported when checking for the policyType parameter. The policyType value is used for differentiating the policy types, in the request endpoint. The supported values are:
 *
 * * ``ACCESS_POLICY/GLOBAL_POLICY``
 * * ``TIMEOUT_POLICY/REAUTH_POLICY``
 * * ``BYPASS_POLICY/CLIENT_FORWARDING_POLICY``
 * * ``INSPECTION_POLICY``
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const accessPolicy = zpa.getPolicyType({
 *     policyType: "ACCESS_POLICY",
 * });
 * export const zpaPolicyTypeAccessPolicy = accessPolicy.then(accessPolicy => accessPolicy.id);
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const globalPolicy = zpa.getPolicyType({
 *     policyType: "GLOBAL_POLICY",
 * });
 * export const zpaPolicyTypeAccessPolicy = globalPolicy.then(globalPolicy => globalPolicy.id);
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const timeoutPolicy = zpa.getPolicyType({
 *     policyType: "TIMEOUT_POLICY",
 * });
 * export const zpaPolicyTypeTimeoutPolicy = timeoutPolicy.then(timeoutPolicy => timeoutPolicy.id);
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const reauthPolicy = zpa.getPolicyType({
 *     policyType: "REAUTH_POLICY",
 * });
 * export const zpaPolicyTypeReauthPolicy = reauthPolicy.then(reauthPolicy => reauthPolicy.id);
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const clientForwardingPolicy = zpa.getPolicyType({
 *     policyType: "CLIENT_FORWARDING_POLICY",
 * });
 * export const zpaPolicyTypeClientForwardingPolicy = clientForwardingPolicy.then(clientForwardingPolicy => clientForwardingPolicy.id);
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const inspectionPolicy = zpa.getPolicyType({
 *     policyType: "INSPECTION_POLICY",
 * });
 * export const zpaPolicyTypeInspectionPolicy = inspectionPolicy.then(inspectionPolicy => inspectionPolicy.id);
 * ```
 */
export function getPolicyTypeOutput(args?: GetPolicyTypeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPolicyTypeResult> {
    return pulumi.output(args).apply((a: any) => getPolicyType(a, opts))
}

/**
 * A collection of arguments for invoking getPolicyType.
 */
export interface GetPolicyTypeOutputArgs {
    id?: pulumi.Input<string>;
    microtenantId?: pulumi.Input<string>;
    microtenantName?: pulumi.Input<string>;
    policyType?: pulumi.Input<string>;
}
