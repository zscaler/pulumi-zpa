// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * @deprecated zpa.index/policytimeoutrulev2.PolicyTimeoutRuleV2 has been deprecated in favor of zpa.index/policyaccesstimeoutrulev2.PolicyAccessTimeOutRuleV2
 */
export class PolicyTimeoutRuleV2 extends pulumi.CustomResource {
    /**
     * Get an existing PolicyTimeoutRuleV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyTimeoutRuleV2State, opts?: pulumi.CustomResourceOptions): PolicyTimeoutRuleV2 {
        pulumi.log.warn("PolicyTimeoutRuleV2 is deprecated: zpa.index/policytimeoutrulev2.PolicyTimeoutRuleV2 has been deprecated in favor of zpa.index/policyaccesstimeoutrulev2.PolicyAccessTimeOutRuleV2")
        return new PolicyTimeoutRuleV2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zpa:index/policyTimeoutRuleV2:PolicyTimeoutRuleV2';

    /**
     * Returns true if the given object is an instance of PolicyTimeoutRuleV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicyTimeoutRuleV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicyTimeoutRuleV2.__pulumiType;
    }

    /**
     * This is for providing the rule action.
     */
    public readonly action!: pulumi.Output<string | undefined>;
    /**
     * This is for proviidng the set of conditions for the policy.
     */
    public readonly conditions!: pulumi.Output<outputs.PolicyTimeoutRuleV2Condition[]>;
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
     * Create a PolicyTimeoutRuleV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated zpa.index/policytimeoutrulev2.PolicyTimeoutRuleV2 has been deprecated in favor of zpa.index/policyaccesstimeoutrulev2.PolicyAccessTimeOutRuleV2 */
    constructor(name: string, args?: PolicyTimeoutRuleV2Args, opts?: pulumi.CustomResourceOptions)
    /** @deprecated zpa.index/policytimeoutrulev2.PolicyTimeoutRuleV2 has been deprecated in favor of zpa.index/policyaccesstimeoutrulev2.PolicyAccessTimeOutRuleV2 */
    constructor(name: string, argsOrState?: PolicyTimeoutRuleV2Args | PolicyTimeoutRuleV2State, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("PolicyTimeoutRuleV2 is deprecated: zpa.index/policytimeoutrulev2.PolicyTimeoutRuleV2 has been deprecated in favor of zpa.index/policyaccesstimeoutrulev2.PolicyAccessTimeOutRuleV2")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyTimeoutRuleV2State | undefined;
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
            const args = argsOrState as PolicyTimeoutRuleV2Args | undefined;
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
        super(PolicyTimeoutRuleV2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PolicyTimeoutRuleV2 resources.
 */
export interface PolicyTimeoutRuleV2State {
    /**
     * This is for providing the rule action.
     */
    action?: pulumi.Input<string>;
    /**
     * This is for proviidng the set of conditions for the policy.
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.PolicyTimeoutRuleV2Condition>[]>;
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
 * The set of arguments for constructing a PolicyTimeoutRuleV2 resource.
 */
export interface PolicyTimeoutRuleV2Args {
    /**
     * This is for providing the rule action.
     */
    action?: pulumi.Input<string>;
    /**
     * This is for proviidng the set of conditions for the policy.
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.PolicyTimeoutRuleV2Condition>[]>;
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