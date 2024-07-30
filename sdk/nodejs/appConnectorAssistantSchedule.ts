// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zpa/configuring-app-connectors-settings)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-auto-delete-disconnected-app-connectors-using-api)
 *
 * Use the **zpa_app_connector_assistant_schedule** resource sets the scheduled frequency at which the disconnected App Connectors are eligible for deletion. The supported value for frequency is days. The frequencyInterval field is the number of days after an App Connector disconnects for it to become eligible for deletion. The minimum supported value for frequencyInterval is 5.
 *
 * > **NOTE** - When enabling the Assistant Schedule for the first time, you must provide the `customerId` information. If you authenticated using environment variables and used `ZPA_CUSTOMER_ID` environment variable, you don't have to define the customerId attribute in the HCL configuration, and the provider will automatically use the value from the environment variable `ZPA_CUSTOMER_ID`
 *
 * ## Example Usage
 *
 * ### Defined Customer ID Value
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const _this = new zpa.AppConnectorAssistantSchedule("this", {
 *     customerId: "123456789101112",
 *     deleteDisabled: true,
 *     enabled: true,
 *     frequency: "days",
 *     frequencyInterval: "5",
 * });
 * ```
 *
 * ### Customer ID Via Environment Variable
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const _this = new zpa.AppConnectorAssistantSchedule("this", {
 *     deleteDisabled: true,
 *     enabled: true,
 *     frequency: "days",
 *     frequencyInterval: "5",
 * });
 * ```
 *
 * ## Import
 *
 * Import is not currently supported for this resource.
 */
export class AppConnectorAssistantSchedule extends pulumi.CustomResource {
    /**
     * Get an existing AppConnectorAssistantSchedule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppConnectorAssistantScheduleState, opts?: pulumi.CustomResourceOptions): AppConnectorAssistantSchedule {
        return new AppConnectorAssistantSchedule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zpa:index/appConnectorAssistantSchedule:AppConnectorAssistantSchedule';

    /**
     * Returns true if the given object is an instance of AppConnectorAssistantSchedule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppConnectorAssistantSchedule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppConnectorAssistantSchedule.__pulumiType;
    }

    public readonly customerId!: pulumi.Output<string>;
    public readonly deleteDisabled!: pulumi.Output<boolean | undefined>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public readonly frequency!: pulumi.Output<string | undefined>;
    public readonly frequencyInterval!: pulumi.Output<string | undefined>;

    /**
     * Create a AppConnectorAssistantSchedule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AppConnectorAssistantScheduleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppConnectorAssistantScheduleArgs | AppConnectorAssistantScheduleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppConnectorAssistantScheduleState | undefined;
            resourceInputs["customerId"] = state ? state.customerId : undefined;
            resourceInputs["deleteDisabled"] = state ? state.deleteDisabled : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["frequency"] = state ? state.frequency : undefined;
            resourceInputs["frequencyInterval"] = state ? state.frequencyInterval : undefined;
        } else {
            const args = argsOrState as AppConnectorAssistantScheduleArgs | undefined;
            resourceInputs["customerId"] = args ? args.customerId : undefined;
            resourceInputs["deleteDisabled"] = args ? args.deleteDisabled : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["frequency"] = args ? args.frequency : undefined;
            resourceInputs["frequencyInterval"] = args ? args.frequencyInterval : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "zpa:index/assistantSchedule:AssistantSchedule" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(AppConnectorAssistantSchedule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppConnectorAssistantSchedule resources.
 */
export interface AppConnectorAssistantScheduleState {
    customerId?: pulumi.Input<string>;
    deleteDisabled?: pulumi.Input<boolean>;
    enabled?: pulumi.Input<boolean>;
    frequency?: pulumi.Input<string>;
    frequencyInterval?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppConnectorAssistantSchedule resource.
 */
export interface AppConnectorAssistantScheduleArgs {
    customerId?: pulumi.Input<string>;
    deleteDisabled?: pulumi.Input<boolean>;
    enabled?: pulumi.Input<boolean>;
    frequency?: pulumi.Input<string>;
    frequencyInterval?: pulumi.Input<string>;
}
