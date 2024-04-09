// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * @deprecated zpa.index/praportalcontroller.PraPortalController has been deprecated in favor of zpa.index/praportal.PRAPortal
 */
export class PraPortalController extends pulumi.CustomResource {
    /**
     * Get an existing PraPortalController resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PraPortalControllerState, opts?: pulumi.CustomResourceOptions): PraPortalController {
        pulumi.log.warn("PraPortalController is deprecated: zpa.index/praportalcontroller.PraPortalController has been deprecated in favor of zpa.index/praportal.PRAPortal")
        return new PraPortalController(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zpa:index/praPortalController:PraPortalController';

    /**
     * Returns true if the given object is an instance of PraPortalController.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PraPortalController {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PraPortalController.__pulumiType;
    }

    /**
     * The unique identifier of the certificate
     */
    public readonly certificateId!: pulumi.Output<string | undefined>;
    /**
     * The description of the privileged portal
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The domain of the privileged portal
     */
    public readonly domain!: pulumi.Output<string | undefined>;
    /**
     * Whether or not the privileged portal is enabled
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
     * microtenantId as 0 when making requests to retrieve data from the Default Microtenant. Pass microtenantId as null to
     * retrieve data from all customers associated with the tenant.
     */
    public readonly microtenantId!: pulumi.Output<string>;
    /**
     * The name of the privileged portal
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The notification message displayed in the banner of the privileged portallink, if enabled
     */
    public readonly userNotification!: pulumi.Output<string | undefined>;
    /**
     * Indicates if the Notification Banner is enabled (true) or disabled (false)
     */
    public readonly userNotificationEnabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a PraPortalController resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated zpa.index/praportalcontroller.PraPortalController has been deprecated in favor of zpa.index/praportal.PRAPortal */
    constructor(name: string, args?: PraPortalControllerArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated zpa.index/praportalcontroller.PraPortalController has been deprecated in favor of zpa.index/praportal.PRAPortal */
    constructor(name: string, argsOrState?: PraPortalControllerArgs | PraPortalControllerState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("PraPortalController is deprecated: zpa.index/praportalcontroller.PraPortalController has been deprecated in favor of zpa.index/praportal.PRAPortal")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PraPortalControllerState | undefined;
            resourceInputs["certificateId"] = state ? state.certificateId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["microtenantId"] = state ? state.microtenantId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["userNotification"] = state ? state.userNotification : undefined;
            resourceInputs["userNotificationEnabled"] = state ? state.userNotificationEnabled : undefined;
        } else {
            const args = argsOrState as PraPortalControllerArgs | undefined;
            resourceInputs["certificateId"] = args ? args.certificateId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["microtenantId"] = args ? args.microtenantId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["userNotification"] = args ? args.userNotification : undefined;
            resourceInputs["userNotificationEnabled"] = args ? args.userNotificationEnabled : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PraPortalController.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PraPortalController resources.
 */
export interface PraPortalControllerState {
    /**
     * The unique identifier of the certificate
     */
    certificateId?: pulumi.Input<string>;
    /**
     * The description of the privileged portal
     */
    description?: pulumi.Input<string>;
    /**
     * The domain of the privileged portal
     */
    domain?: pulumi.Input<string>;
    /**
     * Whether or not the privileged portal is enabled
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
     * microtenantId as 0 when making requests to retrieve data from the Default Microtenant. Pass microtenantId as null to
     * retrieve data from all customers associated with the tenant.
     */
    microtenantId?: pulumi.Input<string>;
    /**
     * The name of the privileged portal
     */
    name?: pulumi.Input<string>;
    /**
     * The notification message displayed in the banner of the privileged portallink, if enabled
     */
    userNotification?: pulumi.Input<string>;
    /**
     * Indicates if the Notification Banner is enabled (true) or disabled (false)
     */
    userNotificationEnabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a PraPortalController resource.
 */
export interface PraPortalControllerArgs {
    /**
     * The unique identifier of the certificate
     */
    certificateId?: pulumi.Input<string>;
    /**
     * The description of the privileged portal
     */
    description?: pulumi.Input<string>;
    /**
     * The domain of the privileged portal
     */
    domain?: pulumi.Input<string>;
    /**
     * Whether or not the privileged portal is enabled
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
     * microtenantId as 0 when making requests to retrieve data from the Default Microtenant. Pass microtenantId as null to
     * retrieve data from all customers associated with the tenant.
     */
    microtenantId?: pulumi.Input<string>;
    /**
     * The name of the privileged portal
     */
    name?: pulumi.Input<string>;
    /**
     * The notification message displayed in the banner of the privileged portallink, if enabled
     */
    userNotification?: pulumi.Input<string>;
    /**
     * Indicates if the Notification Banner is enabled (true) or disabled (false)
     */
    userNotificationEnabled?: pulumi.Input<boolean>;
}