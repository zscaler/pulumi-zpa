// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-isolation-policy)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-isolation-policies-using-api)
 *
 * The **zpa_policy_isolation_rule_v2** resource creates and manages policy access isolation rule in the Zscaler Private Access cloud using a new v2 API endpoint.
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
 * const thisIsolationProfile = zpa.getIsolationProfile({
 *     name: "zpa_isolation_profile",
 * });
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
 * // Create Policy Access Isolation Rule V2
 * const thisPolicyAccessIsolationRuleV2 = new zpa.PolicyAccessIsolationRuleV2("thisPolicyAccessIsolationRuleV2", {
 *     description: "Example",
 *     action: "ISOLATE",
 *     zpnIsolationProfileId: thisIsolationProfile.then(thisIsolationProfile => thisIsolationProfile.id),
 *     conditions: [
 *         {
 *             operator: "OR",
 *             operands: [{
 *                 objectType: "CLIENT_TYPE",
 *                 values: ["zpn_client_type_exporter"],
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
 * | EDGE_CONNECTOR_GROUP  |   |  |  ``<edge_connector_id>`` |
 * | MACHINE_GRP   |   |  | ``machineGroupId`` |
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
 * Policy access isolation rule can be imported by using `<RULE ID>` as the import ID.
 *
 * For example:
 *
 * ```sh
 * $ pulumi import zpa:index/policyAccessIsolationRuleV2:PolicyAccessIsolationRuleV2 example <rule_id>
 * ```
 */
export class PolicyAccessIsolationRuleV2 extends pulumi.CustomResource {
    /**
     * Get an existing PolicyAccessIsolationRuleV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyAccessIsolationRuleV2State, opts?: pulumi.CustomResourceOptions): PolicyAccessIsolationRuleV2 {
        return new PolicyAccessIsolationRuleV2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zpa:index/policyAccessIsolationRuleV2:PolicyAccessIsolationRuleV2';

    /**
     * Returns true if the given object is an instance of PolicyAccessIsolationRuleV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicyAccessIsolationRuleV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicyAccessIsolationRuleV2.__pulumiType;
    }

    /**
     * This is for providing the rule action. Supported values: `ISOLATE` Default.
     */
    public readonly action!: pulumi.Output<string | undefined>;
    /**
     * This is for proviidng the set of conditions for the policy.
     */
    public readonly conditions!: pulumi.Output<outputs.PolicyAccessIsolationRuleV2Condition[]>;
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
     * Use zpa*isolation*profile data source to retrieve the necessary Isolation profile ID `zpnIsolationProfileId`
     */
    public readonly zpnIsolationProfileId!: pulumi.Output<string>;

    /**
     * Create a PolicyAccessIsolationRuleV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PolicyAccessIsolationRuleV2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyAccessIsolationRuleV2Args | PolicyAccessIsolationRuleV2State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyAccessIsolationRuleV2State | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["conditions"] = state ? state.conditions : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["microtenantId"] = state ? state.microtenantId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policySetId"] = state ? state.policySetId : undefined;
            resourceInputs["zpnIsolationProfileId"] = state ? state.zpnIsolationProfileId : undefined;
        } else {
            const args = argsOrState as PolicyAccessIsolationRuleV2Args | undefined;
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["conditions"] = args ? args.conditions : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["microtenantId"] = args ? args.microtenantId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["zpnIsolationProfileId"] = args ? args.zpnIsolationProfileId : undefined;
            resourceInputs["policySetId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "zpa:index/policyIsolationRuleV2:PolicyIsolationRuleV2" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(PolicyAccessIsolationRuleV2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PolicyAccessIsolationRuleV2 resources.
 */
export interface PolicyAccessIsolationRuleV2State {
    /**
     * This is for providing the rule action. Supported values: `ISOLATE` Default.
     */
    action?: pulumi.Input<string>;
    /**
     * This is for proviidng the set of conditions for the policy.
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.PolicyAccessIsolationRuleV2Condition>[]>;
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
    /**
     * Use zpa*isolation*profile data source to retrieve the necessary Isolation profile ID `zpnIsolationProfileId`
     */
    zpnIsolationProfileId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PolicyAccessIsolationRuleV2 resource.
 */
export interface PolicyAccessIsolationRuleV2Args {
    /**
     * This is for providing the rule action. Supported values: `ISOLATE` Default.
     */
    action?: pulumi.Input<string>;
    /**
     * This is for proviidng the set of conditions for the policy.
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.PolicyAccessIsolationRuleV2Condition>[]>;
    /**
     * This is the description of the access policy rule.
     */
    description?: pulumi.Input<string>;
    microtenantId?: pulumi.Input<string>;
    /**
     * This is the name of the policy rule.
     */
    name?: pulumi.Input<string>;
    /**
     * Use zpa*isolation*profile data source to retrieve the necessary Isolation profile ID `zpnIsolationProfileId`
     */
    zpnIsolationProfileId?: pulumi.Input<string>;
}
