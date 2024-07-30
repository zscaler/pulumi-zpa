// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-client-forwarding-policy)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-client-forwarding-policies-using-api)
 *
 * The **zpa_policy_forwarding_rule_v2** resource creates and manages policy access forwarding rule in the Zscaler Private Access cloud using a new v2 API endpoint.
 *
 *   ⚠️ **NOTE**: This resource is recommended if your configuration requires the association of more than 1000 resource criteria per rule.
 *
 *   ⚠️ **WARNING:**: The attribute ``ruleOrder`` is now deprecated in favor of the new resource  ``policyAccessRuleReorder``
 *
 * ## Example Usage
 *
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
 * const thisPolicyAccessForwardingRuleV2 = new zpa.PolicyAccessForwardingRuleV2("thisPolicyAccessForwardingRuleV2", {
 *     description: "Example",
 *     action: "BYPASS",
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
 * $ pulumi import zpa:index/policyAccessForwardingRuleV2:PolicyAccessForwardingRuleV2 example <rule_id>
 * ```
 */
export class PolicyAccessForwardingRuleV2 extends pulumi.CustomResource {
    /**
     * Get an existing PolicyAccessForwardingRuleV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyAccessForwardingRuleV2State, opts?: pulumi.CustomResourceOptions): PolicyAccessForwardingRuleV2 {
        return new PolicyAccessForwardingRuleV2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zpa:index/policyAccessForwardingRuleV2:PolicyAccessForwardingRuleV2';

    /**
     * Returns true if the given object is an instance of PolicyAccessForwardingRuleV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicyAccessForwardingRuleV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicyAccessForwardingRuleV2.__pulumiType;
    }

    /**
     * This is for providing the rule action. Supported values: `BYPASS`, `INTERCEPT`, and `INTERCEPT_ACCESSIBLE`
     */
    public readonly action!: pulumi.Output<string | undefined>;
    /**
     * This is for proviidng the set of conditions for the policy.
     */
    public readonly conditions!: pulumi.Output<outputs.PolicyAccessForwardingRuleV2Condition[]>;
    /**
     * This is the description of the access policy rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly microtenantId!: pulumi.Output<string>;
    /**
     * This is the name of the policy rule.
     */
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly policySetId!: pulumi.Output<string>;

    /**
     * Create a PolicyAccessForwardingRuleV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PolicyAccessForwardingRuleV2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyAccessForwardingRuleV2Args | PolicyAccessForwardingRuleV2State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyAccessForwardingRuleV2State | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["conditions"] = state ? state.conditions : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["microtenantId"] = state ? state.microtenantId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policySetId"] = state ? state.policySetId : undefined;
        } else {
            const args = argsOrState as PolicyAccessForwardingRuleV2Args | undefined;
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["conditions"] = args ? args.conditions : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["microtenantId"] = args ? args.microtenantId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["policySetId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "zpa:index/policyForwardingRuleV2:PolicyForwardingRuleV2" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(PolicyAccessForwardingRuleV2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PolicyAccessForwardingRuleV2 resources.
 */
export interface PolicyAccessForwardingRuleV2State {
    /**
     * This is for providing the rule action. Supported values: `BYPASS`, `INTERCEPT`, and `INTERCEPT_ACCESSIBLE`
     */
    action?: pulumi.Input<string>;
    /**
     * This is for proviidng the set of conditions for the policy.
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.PolicyAccessForwardingRuleV2Condition>[]>;
    /**
     * This is the description of the access policy rule.
     */
    description?: pulumi.Input<string>;
    microtenantId?: pulumi.Input<string>;
    /**
     * This is the name of the policy rule.
     */
    name?: pulumi.Input<string>;
    policySetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PolicyAccessForwardingRuleV2 resource.
 */
export interface PolicyAccessForwardingRuleV2Args {
    /**
     * This is for providing the rule action. Supported values: `BYPASS`, `INTERCEPT`, and `INTERCEPT_ACCESSIBLE`
     */
    action?: pulumi.Input<string>;
    /**
     * This is for proviidng the set of conditions for the policy.
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.PolicyAccessForwardingRuleV2Condition>[]>;
    /**
     * This is the description of the access policy rule.
     */
    description?: pulumi.Input<string>;
    microtenantId?: pulumi.Input<string>;
    /**
     * This is the name of the policy rule.
     */
    name?: pulumi.Input<string>;
}
