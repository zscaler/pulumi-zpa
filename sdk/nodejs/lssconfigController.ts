// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-log-streaming-service)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-log-streaming-service-configurations-using-api)
 *
 * The **zpa_lss_config_controller** resource creates and manages Log Streaming Service (LSS) in the Zscaler Private Access cloud for App Connector Metrics `zpnAstComprehensiveStats`.
 *
 * ## Example 1 - LSS App Connector Metrics - Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const zpnAstComprehensiveStats = zpa.getLSSLogTypeFormats({
 *     logType: "zpn_ast_comprehensive_stats",
 * });
 * const lssSiemPolicy = zpa.getPolicyType({
 *     policyType: "SIEM_POLICY",
 * });
 * const _this = zpa.getAppConnectorGroup({
 *     name: "Example100",
 * });
 * const lssAppConnectorMetrics = new zpa.LSSConfigController("lssAppConnectorMetrics", {
 *     config: {
 *         name: "LSS App Connector Metrics",
 *         description: "LSS App Connector Metrics",
 *         enabled: true,
 *         format: zpnAstComprehensiveStats.then(zpnAstComprehensiveStats => zpnAstComprehensiveStats.json),
 *         lssHost: "splunk1.acme.com",
 *         lssPort: "5001",
 *         sourceLogType: "zpn_ast_comprehensive_stats",
 *         useTls: true,
 *     },
 *     connectorGroups: [{
 *         ids: [_this.then(_this => _this.id)],
 *     }],
 * });
 * ```
 *
 * ## LSS Source Log Type Table
 *
 * |       Source Log Type                     |            Description                 |
 * |-------------------------------------------|----------------------------------------|
 * |        `zpnTransLog`                    |        `User Activity`                 |
 * |        `zpnAuthLog`                     |         `User Status`                  |
 * |        `zpnAstAuthLog`                 |        `App Connector Status`          |
 * |        `zpnHttpTransLog`               |         `Web Browser`                  |
 * |        `zpnAuditLog`                    |         `Audit Logs`                   |
 * |        `zpnSysAuthLog`                 |         `Private Service Edge Status`  |
 * |        `zpnAstComprehensiveStats`      |         `App Connector Metrics`        |
 * |        `zpnPbrokerComprehensiveStats`  |         `Private Service Edge Metrics` |
 * |        `zpnWafHttpExchangesLog`       |         `ZPA App Protection`           |
 *
 * ## Import
 *
 * Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
 *
 * Visit
 */
export class LSSConfigController extends pulumi.CustomResource {
    /**
     * Get an existing LSSConfigController resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LSSConfigControllerState, opts?: pulumi.CustomResourceOptions): LSSConfigController {
        return new LSSConfigController(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zpa:index/lSSConfigController:LSSConfigController';

    /**
     * Returns true if the given object is an instance of LSSConfigController.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LSSConfigController {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LSSConfigController.__pulumiType;
    }

    public readonly config!: pulumi.Output<outputs.LSSConfigControllerConfig | undefined>;
    /**
     * App Connector Group(s) to be added to the LSS configuration
     */
    public readonly connectorGroups!: pulumi.Output<outputs.LSSConfigControllerConnectorGroup[] | undefined>;
    public /*out*/ readonly policyRuleId!: pulumi.Output<string>;
    public readonly policyRuleResource!: pulumi.Output<outputs.LSSConfigControllerPolicyRuleResource | undefined>;

    /**
     * Create a LSSConfigController resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LSSConfigControllerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LSSConfigControllerArgs | LSSConfigControllerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LSSConfigControllerState | undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["connectorGroups"] = state ? state.connectorGroups : undefined;
            resourceInputs["policyRuleId"] = state ? state.policyRuleId : undefined;
            resourceInputs["policyRuleResource"] = state ? state.policyRuleResource : undefined;
        } else {
            const args = argsOrState as LSSConfigControllerArgs | undefined;
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["connectorGroups"] = args ? args.connectorGroups : undefined;
            resourceInputs["policyRuleResource"] = args ? args.policyRuleResource : undefined;
            resourceInputs["policyRuleId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LSSConfigController.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LSSConfigController resources.
 */
export interface LSSConfigControllerState {
    config?: pulumi.Input<inputs.LSSConfigControllerConfig>;
    /**
     * App Connector Group(s) to be added to the LSS configuration
     */
    connectorGroups?: pulumi.Input<pulumi.Input<inputs.LSSConfigControllerConnectorGroup>[]>;
    policyRuleId?: pulumi.Input<string>;
    policyRuleResource?: pulumi.Input<inputs.LSSConfigControllerPolicyRuleResource>;
}

/**
 * The set of arguments for constructing a LSSConfigController resource.
 */
export interface LSSConfigControllerArgs {
    config?: pulumi.Input<inputs.LSSConfigControllerConfig>;
    /**
     * App Connector Group(s) to be added to the LSS configuration
     */
    connectorGroups?: pulumi.Input<pulumi.Input<inputs.LSSConfigControllerConnectorGroup>[]>;
    policyRuleResource?: pulumi.Input<inputs.LSSConfigControllerPolicyRuleResource>;
}
