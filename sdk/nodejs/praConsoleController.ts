// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * @deprecated zpa.index/praconsolecontroller.PraConsoleController has been deprecated in favor of zpa.index/praconsole.PRAConsole
 */
export class PraConsoleController extends pulumi.CustomResource {
    /**
     * Get an existing PraConsoleController resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PraConsoleControllerState, opts?: pulumi.CustomResourceOptions): PraConsoleController {
        pulumi.log.warn("PraConsoleController is deprecated: zpa.index/praconsolecontroller.PraConsoleController has been deprecated in favor of zpa.index/praconsole.PRAConsole")
        return new PraConsoleController(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zpa:index/praConsoleController:PraConsoleController';

    /**
     * Returns true if the given object is an instance of PraConsoleController.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PraConsoleController {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PraConsoleController.__pulumiType;
    }

    /**
     * The description of the privileged console
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Whether or not the privileged console is enabled
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * The privileged console icon. The icon image is converted to base64 encoded text format
     */
    public readonly iconText!: pulumi.Output<string>;
    /**
     * The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
     * microtenantId as 0 when making requests to retrieve data from the Default Microtenant.
     */
    public readonly microtenantId!: pulumi.Output<string>;
    /**
     * The name of the privileged console
     */
    public readonly name!: pulumi.Output<string>;
    public readonly praApplication!: pulumi.Output<outputs.PraConsoleControllerPraApplication>;
    public readonly praPortals!: pulumi.Output<outputs.PraConsoleControllerPraPortal[]>;

    /**
     * Create a PraConsoleController resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated zpa.index/praconsolecontroller.PraConsoleController has been deprecated in favor of zpa.index/praconsole.PRAConsole */
    constructor(name: string, args: PraConsoleControllerArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated zpa.index/praconsolecontroller.PraConsoleController has been deprecated in favor of zpa.index/praconsole.PRAConsole */
    constructor(name: string, argsOrState?: PraConsoleControllerArgs | PraConsoleControllerState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("PraConsoleController is deprecated: zpa.index/praconsolecontroller.PraConsoleController has been deprecated in favor of zpa.index/praconsole.PRAConsole")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PraConsoleControllerState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["iconText"] = state ? state.iconText : undefined;
            resourceInputs["microtenantId"] = state ? state.microtenantId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["praApplication"] = state ? state.praApplication : undefined;
            resourceInputs["praPortals"] = state ? state.praPortals : undefined;
        } else {
            const args = argsOrState as PraConsoleControllerArgs | undefined;
            if ((!args || args.praApplication === undefined) && !opts.urn) {
                throw new Error("Missing required property 'praApplication'");
            }
            if ((!args || args.praPortals === undefined) && !opts.urn) {
                throw new Error("Missing required property 'praPortals'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["iconText"] = args ? args.iconText : undefined;
            resourceInputs["microtenantId"] = args ? args.microtenantId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["praApplication"] = args ? args.praApplication : undefined;
            resourceInputs["praPortals"] = args ? args.praPortals : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PraConsoleController.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PraConsoleController resources.
 */
export interface PraConsoleControllerState {
    /**
     * The description of the privileged console
     */
    description?: pulumi.Input<string>;
    /**
     * Whether or not the privileged console is enabled
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The privileged console icon. The icon image is converted to base64 encoded text format
     */
    iconText?: pulumi.Input<string>;
    /**
     * The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
     * microtenantId as 0 when making requests to retrieve data from the Default Microtenant.
     */
    microtenantId?: pulumi.Input<string>;
    /**
     * The name of the privileged console
     */
    name?: pulumi.Input<string>;
    praApplication?: pulumi.Input<inputs.PraConsoleControllerPraApplication>;
    praPortals?: pulumi.Input<pulumi.Input<inputs.PraConsoleControllerPraPortal>[]>;
}

/**
 * The set of arguments for constructing a PraConsoleController resource.
 */
export interface PraConsoleControllerArgs {
    /**
     * The description of the privileged console
     */
    description?: pulumi.Input<string>;
    /**
     * Whether or not the privileged console is enabled
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The privileged console icon. The icon image is converted to base64 encoded text format
     */
    iconText?: pulumi.Input<string>;
    /**
     * The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
     * microtenantId as 0 when making requests to retrieve data from the Default Microtenant.
     */
    microtenantId?: pulumi.Input<string>;
    /**
     * The name of the privileged console
     */
    name?: pulumi.Input<string>;
    praApplication: pulumi.Input<inputs.PraConsoleControllerPraApplication>;
    praPortals: pulumi.Input<pulumi.Input<inputs.PraConsoleControllerPraPortal>[]>;
}
