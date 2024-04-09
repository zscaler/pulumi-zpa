// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-timeout-policy)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-timeout-policies-using-api)
 *
 * The **zpa_policy_timeout_rule_v2** resource creates and manages policy access timeout rule in the Zscaler Private Access cloud using a new v2 API endpoint.
 *
 *   ⚠️ **NOTE**: This resource is recommended if your configuration requires the association of more than 1000 resource criteria per rule.
 *
 *   ⚠️ **WARNING:**: The attribute ``ruleOrder`` is now deprecated in favor of the new resource  ``policyAccessRuleReorder``
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 * import * as zpa from "@pulumi/zpa";
 *
 * const thisIdPController = zpa.getIdPController({
 *     name: "Idp_Name",
 * });
 * const emailUserSso = zpa.getSAMLAttribute({
 *     name: "Email_Users",
 *     idpName: "Idp_Name",
 * });
 * const groupUser = zpa.getSAMLAttribute({
 *     name: "GroupName_Users",
 *     idpName: "Idp_Name",
 * });
 * const a000 = zpa.getSCIMGroups({
 *     name: "A000",
 *     idpName: "Idp_Name",
 * });
 * const b000 = zpa.getSCIMGroups({
 *     name: "B000",
 *     idpName: "Idp_Name",
 * });
 * // Create Segment Group
 * const thisSegmentGroup = new zpa.SegmentGroup("thisSegmentGroup", {
 *     description: "Example",
 *     enabled: true,
 * });
 * // Create Policy Access Rule V2
 * const thisPolicyAccessTimeOutRuleV2 = new zpa.PolicyAccessTimeOutRuleV2("thisPolicyAccessTimeOutRuleV2", {
 *     description: "Example",
 *     action: "RE_AUTH",
 *     reauthIdleTimeout: "10 Days",
 *     reauthTimeout: "10 Days",
 *     conditions: [
 *         {
 *             operator: "OR",
 *             operands: [{
 *                 objectType: "APP_GROUP",
 *                 values: [thisSegmentGroup.id],
 *             }],
 *         },
 *         {
 *             operator: "OR",
 *             operands: [
 *                 {
 *                     objectType: "SAML",
 *                     entryValues: [
 *                         {
 *                             rhs: "user1@acme.com",
 *                             lhs: emailUserSso.then(emailUserSso => emailUserSso.id),
 *                         },
 *                         {
 *                             rhs: "A000",
 *                             lhs: groupUser.then(groupUser => groupUser.id),
 *                         },
 *                     ],
 *                 },
 *                 {
 *                     objectType: "SCIM_GROUP",
 *                     entryValues: [
 *                         {
 *                             rhs: a000.then(a000 => a000.id),
 *                             lhs: thisIdPController.then(thisIdPController => thisIdPController.id),
 *                         },
 *                         {
 *                             rhs: b000.then(b000 => b000.id),
 *                             lhs: thisIdPController.then(thisIdPController => thisIdPController.id),
 *                         },
 *                     ],
 *                 },
 *             ],
 *         },
 *         {
 *             operator: "OR",
 *             operands: [{
 *                 objectType: "PLATFORM",
 *                 entryValues: [
 *                     {
 *                         rhs: "true",
 *                         lhs: "linux",
 *                     },
 *                     {
 *                         rhs: "true",
 *                         lhs: "android",
 *                     },
 *                 ],
 *             }],
 *         },
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## LHS and RHS Values
 *
 * | Object Type | LHS| RHS| VALUES
 * |----------|-----------|----------|----------
 * | APP  |   |  | ``applicationSegmentId`` |
 * | APP_GROUP  |   |  | ``segmentGroupId``|
 * | CLIENT_TYPE  |   |  |  ``zpnClientTypeZappl``, ``zpnClientTypeExporter``, ``zpnClientTypeBrowserIsolation``, ``zpnClientTypeIpAnchoring``, ``zpnClientTypeEdgeConnector``, ``zpnClientTypeBranchConnector``,  ``zpnClientTypeZappPartner``, ``zpnClientTypeZapp``  |
 * | SAML | ``samlAttributeId``  | ``attributeValueToMatch`` |
 * | SCIM | ``scimAttributeId``  | ``attributeValueToMatch``  |
 * | SCIM_GROUP | ``scimGroupAttributeId``  | ``attributeValueToMatch``  |
 * | PLATFORM | ``mac``, ``ios``, ``windows``, ``android``, ``linux`` | ``"true"`` / ``"false"`` |
 * | POSTURE | ``postureUdid``  | ``"true"`` / ``"false"`` |
 *
 * ## Import
 *
 * Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
 *
 * Visit
 *
 * Policy access timeout rule can be imported by using `<RULE ID>` as the import ID.
 *
 * For example:
 *
 * ```sh
 * $ pulumi import zpa:index/policyAccessTimeOutRuleV2:PolicyAccessTimeOutRuleV2 example <rule_id>
 * ```
 */
export class PolicyAccessTimeOutRuleV2 extends pulumi.CustomResource {
    /**
     * Get an existing PolicyAccessTimeOutRuleV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyAccessTimeOutRuleV2State, opts?: pulumi.CustomResourceOptions): PolicyAccessTimeOutRuleV2 {
        return new PolicyAccessTimeOutRuleV2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zpa:index/policyAccessTimeOutRuleV2:PolicyAccessTimeOutRuleV2';

    /**
     * Returns true if the given object is an instance of PolicyAccessTimeOutRuleV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicyAccessTimeOutRuleV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicyAccessTimeOutRuleV2.__pulumiType;
    }

    /**
     * This is for providing the rule action.
     */
    public readonly action!: pulumi.Output<string | undefined>;
    /**
     * This is for proviidng the set of conditions for the policy.
     */
    public readonly conditions!: pulumi.Output<outputs.PolicyAccessTimeOutRuleV2Condition[]>;
    /**
     * This is for providing a customer message for the user.
     */
    public readonly customMsg!: pulumi.Output<string | undefined>;
    /**
     * This is the description of the access policy.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly microtenantId!: pulumi.Output<string>;
    /**
     * This is the name of the policy.
     */
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly policySetId!: pulumi.Output<string>;
    public readonly reauthIdleTimeout!: pulumi.Output<string>;
    public readonly reauthTimeout!: pulumi.Output<string>;

    /**
     * Create a PolicyAccessTimeOutRuleV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PolicyAccessTimeOutRuleV2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyAccessTimeOutRuleV2Args | PolicyAccessTimeOutRuleV2State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyAccessTimeOutRuleV2State | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["conditions"] = state ? state.conditions : undefined;
            resourceInputs["customMsg"] = state ? state.customMsg : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["microtenantId"] = state ? state.microtenantId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policySetId"] = state ? state.policySetId : undefined;
            resourceInputs["reauthIdleTimeout"] = state ? state.reauthIdleTimeout : undefined;
            resourceInputs["reauthTimeout"] = state ? state.reauthTimeout : undefined;
        } else {
            const args = argsOrState as PolicyAccessTimeOutRuleV2Args | undefined;
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["conditions"] = args ? args.conditions : undefined;
            resourceInputs["customMsg"] = args ? args.customMsg : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["microtenantId"] = args ? args.microtenantId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["reauthIdleTimeout"] = args ? args.reauthIdleTimeout : undefined;
            resourceInputs["reauthTimeout"] = args ? args.reauthTimeout : undefined;
            resourceInputs["policySetId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "zpa:index/policyTimeoutRuleV2:PolicyTimeoutRuleV2" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(PolicyAccessTimeOutRuleV2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PolicyAccessTimeOutRuleV2 resources.
 */
export interface PolicyAccessTimeOutRuleV2State {
    /**
     * This is for providing the rule action.
     */
    action?: pulumi.Input<string>;
    /**
     * This is for proviidng the set of conditions for the policy.
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.PolicyAccessTimeOutRuleV2Condition>[]>;
    /**
     * This is for providing a customer message for the user.
     */
    customMsg?: pulumi.Input<string>;
    /**
     * This is the description of the access policy.
     */
    description?: pulumi.Input<string>;
    microtenantId?: pulumi.Input<string>;
    /**
     * This is the name of the policy.
     */
    name?: pulumi.Input<string>;
    policySetId?: pulumi.Input<string>;
    reauthIdleTimeout?: pulumi.Input<string>;
    reauthTimeout?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PolicyAccessTimeOutRuleV2 resource.
 */
export interface PolicyAccessTimeOutRuleV2Args {
    /**
     * This is for providing the rule action.
     */
    action?: pulumi.Input<string>;
    /**
     * This is for proviidng the set of conditions for the policy.
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.PolicyAccessTimeOutRuleV2Condition>[]>;
    /**
     * This is for providing a customer message for the user.
     */
    customMsg?: pulumi.Input<string>;
    /**
     * This is the description of the access policy.
     */
    description?: pulumi.Input<string>;
    microtenantId?: pulumi.Input<string>;
    /**
     * This is the name of the policy.
     */
    name?: pulumi.Input<string>;
    reauthIdleTimeout?: pulumi.Input<string>;
    reauthTimeout?: pulumi.Input<string>;
}