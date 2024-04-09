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
 * The **zpa_policy_access_rule** resource creates and manages policy access rule in the Zscaler Private Access cloud.
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
 * const idpName = zpa.getIdPController({
 *     name: "IdP_Name",
 * });
 * const engineering = zpa.getSCIMGroups({
 *     name: "Engineering",
 *     idpName: "IdP_Name",
 * });
 * //Create Policy Access Rule
 * const _this = new zpa.PolicyAccessRule("this", {
 *     description: "Example",
 *     action: "ALLOW",
 *     operator: "AND",
 *     conditions: [
 *         {
 *             operator: "OR",
 *             operands: [{
 *                 objectType: "APP",
 *                 lhs: "id",
 *                 rhs: zpa_application_segment["this"].id,
 *             }],
 *         },
 *         {
 *             operator: "OR",
 *             operands: [{
 *                 objectType: "SCIM_GROUP",
 *                 lhs: idpName.then(idpName => idpName.id),
 *                 rhs: engineering.then(engineering => engineering.id),
 *             }],
 *         },
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## LHS and RHS Values
 *
 * | Object Type | LHS| RHS
 * |----------|-----------|----------
 * | APP | ``"id"`` | ``applicationSegmentId`` |
 * | APP_GROUP | ``"id"`` | ``segmentGroupId``|
 * | CLIENT_TYPE | ``"id"`` | ``zpnClientTypeZappl``, ``zpnClientTypeExporter``, ``zpnClientTypeBrowserIsolation``, ``zpnClientTypeIpAnchoring``, ``zpnClientTypeEdgeConnector``, ``zpnClientTypeBranchConnector``,  ``zpnClientTypeZappPartner``, ``zpnClientTypeZapp``  |
 * | EDGE_CONNECTOR_GROUP | ``"id"`` | ``<edge_connector_id>`` |
 * | IDP | ``"id"`` | ``identityProviderId`` |
 * | SAML | ``samlAttributeId``  | ``attributeValueToMatch`` |
 * | SCIM | ``scimAttributeId``  | ``attributeValueToMatch``  |
 * | SCIM_GROUP | ``scimGroupAttributeId``  | ``attributeValueToMatch``  |
 * | PLATFORM | ``mac``, ``ios``, ``windows``, ``android``, ``linux`` | ``"true"`` / ``"false"`` |
 * | MACHINE_GRP | ``"id"`` | ``machineGroupId`` |
 * | POSTURE | ``postureUdid``  | ``"true"`` / ``"false"`` |
 * | TRUSTED_NETWORK | ``networkId``  | ``"true"`` |
 * | COUNTRY_CODE | [2 Letter ISO3166 Alpha2](https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes)  | ``"true"`` / ``"false"`` |
 *
 * ## Import
 *
 * Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
 *
 * Visit
 *
 * Policy access rule can be imported by using `<POLICY ACCESS RULE ID>` as the import ID.
 *
 * For example:
 *
 * ```sh
 * $ pulumi import zpa:index/policyAccessRule:PolicyAccessRule example <policy_access_rule_id>
 * ```
 */
export class PolicyAccessRule extends pulumi.CustomResource {
    /**
     * Get an existing PolicyAccessRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyAccessRuleState, opts?: pulumi.CustomResourceOptions): PolicyAccessRule {
        return new PolicyAccessRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zpa:index/policyAccessRule:PolicyAccessRule';

    /**
     * Returns true if the given object is an instance of PolicyAccessRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicyAccessRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicyAccessRule.__pulumiType;
    }

    /**
     * This is for providing the rule action. Supported values: `ALLOW`, `DENY`, and `REQUIRE_APPROVAL`
     */
    public readonly action!: pulumi.Output<string | undefined>;
    /**
     * This field defines the description of the server.
     */
    public readonly actionId!: pulumi.Output<string | undefined>;
    /**
     * List of app-connector IDs.
     */
    public readonly appConnectorGroups!: pulumi.Output<outputs.PolicyAccessRuleAppConnectorGroup[]>;
    /**
     * List of the server group IDs.
     */
    public readonly appServerGroups!: pulumi.Output<outputs.PolicyAccessRuleAppServerGroup[]>;
    public readonly bypassDefaultRule!: pulumi.Output<boolean | undefined>;
    /**
     * This is for proviidng the set of conditions for the policy.
     */
    public readonly conditions!: pulumi.Output<outputs.PolicyAccessRuleCondition[] | undefined>;
    /**
     * This is for providing a customer message for the user.
     */
    public readonly customMsg!: pulumi.Output<string | undefined>;
    /**
     * This is for providing a customer message for the user.
     */
    public readonly defaultRule!: pulumi.Output<boolean | undefined>;
    /**
     * This is the description of the access policy rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly lssDefaultRule!: pulumi.Output<boolean | undefined>;
    public readonly microtenantId!: pulumi.Output<string>;
    /**
     * This is the name of the policy rule.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Supported values: `AND`, `OR`
     */
    public readonly operator!: pulumi.Output<string>;
    /**
     * - (String) Use zpa*policy*type data source to retrieve the necessary policy Set ID `policySetId`
     * > **NOTE** As of v3.2.0 the `policySetId` attribute is now optional, and will be automatically determined based on the policy type being configured. The attribute is being kept for backwards compatibility, but can be safely removed from existing configurations.
     */
    public readonly policySetId!: pulumi.Output<string>;
    public readonly policyType!: pulumi.Output<string>;
    public readonly priority!: pulumi.Output<string>;
    public readonly reauthDefaultRule!: pulumi.Output<boolean | undefined>;
    public readonly reauthIdleTimeout!: pulumi.Output<string | undefined>;
    public readonly reauthTimeout!: pulumi.Output<string | undefined>;
    /**
     * @deprecated The `ruleOrder` field is now deprecated for all zpa access policy resources in favor of the resource `zpa.PolicyAccessReorderRule`
     */
    public readonly ruleOrder!: pulumi.Output<string>;
    public readonly zpnCbiProfileId!: pulumi.Output<string>;
    public readonly zpnInspectionProfileId!: pulumi.Output<string>;
    public readonly zpnIsolationProfileId!: pulumi.Output<string>;

    /**
     * Create a PolicyAccessRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PolicyAccessRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyAccessRuleArgs | PolicyAccessRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyAccessRuleState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["actionId"] = state ? state.actionId : undefined;
            resourceInputs["appConnectorGroups"] = state ? state.appConnectorGroups : undefined;
            resourceInputs["appServerGroups"] = state ? state.appServerGroups : undefined;
            resourceInputs["bypassDefaultRule"] = state ? state.bypassDefaultRule : undefined;
            resourceInputs["conditions"] = state ? state.conditions : undefined;
            resourceInputs["customMsg"] = state ? state.customMsg : undefined;
            resourceInputs["defaultRule"] = state ? state.defaultRule : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["lssDefaultRule"] = state ? state.lssDefaultRule : undefined;
            resourceInputs["microtenantId"] = state ? state.microtenantId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["operator"] = state ? state.operator : undefined;
            resourceInputs["policySetId"] = state ? state.policySetId : undefined;
            resourceInputs["policyType"] = state ? state.policyType : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["reauthDefaultRule"] = state ? state.reauthDefaultRule : undefined;
            resourceInputs["reauthIdleTimeout"] = state ? state.reauthIdleTimeout : undefined;
            resourceInputs["reauthTimeout"] = state ? state.reauthTimeout : undefined;
            resourceInputs["ruleOrder"] = state ? state.ruleOrder : undefined;
            resourceInputs["zpnCbiProfileId"] = state ? state.zpnCbiProfileId : undefined;
            resourceInputs["zpnInspectionProfileId"] = state ? state.zpnInspectionProfileId : undefined;
            resourceInputs["zpnIsolationProfileId"] = state ? state.zpnIsolationProfileId : undefined;
        } else {
            const args = argsOrState as PolicyAccessRuleArgs | undefined;
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["actionId"] = args ? args.actionId : undefined;
            resourceInputs["appConnectorGroups"] = args ? args.appConnectorGroups : undefined;
            resourceInputs["appServerGroups"] = args ? args.appServerGroups : undefined;
            resourceInputs["bypassDefaultRule"] = args ? args.bypassDefaultRule : undefined;
            resourceInputs["conditions"] = args ? args.conditions : undefined;
            resourceInputs["customMsg"] = args ? args.customMsg : undefined;
            resourceInputs["defaultRule"] = args ? args.defaultRule : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["lssDefaultRule"] = args ? args.lssDefaultRule : undefined;
            resourceInputs["microtenantId"] = args ? args.microtenantId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["operator"] = args ? args.operator : undefined;
            resourceInputs["policySetId"] = args ? args.policySetId : undefined;
            resourceInputs["policyType"] = args ? args.policyType : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["reauthDefaultRule"] = args ? args.reauthDefaultRule : undefined;
            resourceInputs["reauthIdleTimeout"] = args ? args.reauthIdleTimeout : undefined;
            resourceInputs["reauthTimeout"] = args ? args.reauthTimeout : undefined;
            resourceInputs["ruleOrder"] = args ? args.ruleOrder : undefined;
            resourceInputs["zpnCbiProfileId"] = args ? args.zpnCbiProfileId : undefined;
            resourceInputs["zpnInspectionProfileId"] = args ? args.zpnInspectionProfileId : undefined;
            resourceInputs["zpnIsolationProfileId"] = args ? args.zpnIsolationProfileId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PolicyAccessRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PolicyAccessRule resources.
 */
export interface PolicyAccessRuleState {
    /**
     * This is for providing the rule action. Supported values: `ALLOW`, `DENY`, and `REQUIRE_APPROVAL`
     */
    action?: pulumi.Input<string>;
    /**
     * This field defines the description of the server.
     */
    actionId?: pulumi.Input<string>;
    /**
     * List of app-connector IDs.
     */
    appConnectorGroups?: pulumi.Input<pulumi.Input<inputs.PolicyAccessRuleAppConnectorGroup>[]>;
    /**
     * List of the server group IDs.
     */
    appServerGroups?: pulumi.Input<pulumi.Input<inputs.PolicyAccessRuleAppServerGroup>[]>;
    bypassDefaultRule?: pulumi.Input<boolean>;
    /**
     * This is for proviidng the set of conditions for the policy.
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.PolicyAccessRuleCondition>[]>;
    /**
     * This is for providing a customer message for the user.
     */
    customMsg?: pulumi.Input<string>;
    /**
     * This is for providing a customer message for the user.
     */
    defaultRule?: pulumi.Input<boolean>;
    /**
     * This is the description of the access policy rule.
     */
    description?: pulumi.Input<string>;
    lssDefaultRule?: pulumi.Input<boolean>;
    microtenantId?: pulumi.Input<string>;
    /**
     * This is the name of the policy rule.
     */
    name?: pulumi.Input<string>;
    /**
     * Supported values: `AND`, `OR`
     */
    operator?: pulumi.Input<string>;
    /**
     * - (String) Use zpa*policy*type data source to retrieve the necessary policy Set ID `policySetId`
     * > **NOTE** As of v3.2.0 the `policySetId` attribute is now optional, and will be automatically determined based on the policy type being configured. The attribute is being kept for backwards compatibility, but can be safely removed from existing configurations.
     */
    policySetId?: pulumi.Input<string>;
    policyType?: pulumi.Input<string>;
    priority?: pulumi.Input<string>;
    reauthDefaultRule?: pulumi.Input<boolean>;
    reauthIdleTimeout?: pulumi.Input<string>;
    reauthTimeout?: pulumi.Input<string>;
    /**
     * @deprecated The `ruleOrder` field is now deprecated for all zpa access policy resources in favor of the resource `zpa.PolicyAccessReorderRule`
     */
    ruleOrder?: pulumi.Input<string>;
    zpnCbiProfileId?: pulumi.Input<string>;
    zpnInspectionProfileId?: pulumi.Input<string>;
    zpnIsolationProfileId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PolicyAccessRule resource.
 */
export interface PolicyAccessRuleArgs {
    /**
     * This is for providing the rule action. Supported values: `ALLOW`, `DENY`, and `REQUIRE_APPROVAL`
     */
    action?: pulumi.Input<string>;
    /**
     * This field defines the description of the server.
     */
    actionId?: pulumi.Input<string>;
    /**
     * List of app-connector IDs.
     */
    appConnectorGroups?: pulumi.Input<pulumi.Input<inputs.PolicyAccessRuleAppConnectorGroup>[]>;
    /**
     * List of the server group IDs.
     */
    appServerGroups?: pulumi.Input<pulumi.Input<inputs.PolicyAccessRuleAppServerGroup>[]>;
    bypassDefaultRule?: pulumi.Input<boolean>;
    /**
     * This is for proviidng the set of conditions for the policy.
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.PolicyAccessRuleCondition>[]>;
    /**
     * This is for providing a customer message for the user.
     */
    customMsg?: pulumi.Input<string>;
    /**
     * This is for providing a customer message for the user.
     */
    defaultRule?: pulumi.Input<boolean>;
    /**
     * This is the description of the access policy rule.
     */
    description?: pulumi.Input<string>;
    lssDefaultRule?: pulumi.Input<boolean>;
    microtenantId?: pulumi.Input<string>;
    /**
     * This is the name of the policy rule.
     */
    name?: pulumi.Input<string>;
    /**
     * Supported values: `AND`, `OR`
     */
    operator?: pulumi.Input<string>;
    /**
     * - (String) Use zpa*policy*type data source to retrieve the necessary policy Set ID `policySetId`
     * > **NOTE** As of v3.2.0 the `policySetId` attribute is now optional, and will be automatically determined based on the policy type being configured. The attribute is being kept for backwards compatibility, but can be safely removed from existing configurations.
     */
    policySetId?: pulumi.Input<string>;
    policyType?: pulumi.Input<string>;
    priority?: pulumi.Input<string>;
    reauthDefaultRule?: pulumi.Input<boolean>;
    reauthIdleTimeout?: pulumi.Input<string>;
    reauthTimeout?: pulumi.Input<string>;
    /**
     * @deprecated The `ruleOrder` field is now deprecated for all zpa access policy resources in favor of the resource `zpa.PolicyAccessReorderRule`
     */
    ruleOrder?: pulumi.Input<string>;
    zpnCbiProfileId?: pulumi.Input<string>;
    zpnInspectionProfileId?: pulumi.Input<string>;
    zpnIsolationProfileId?: pulumi.Input<string>;
}
