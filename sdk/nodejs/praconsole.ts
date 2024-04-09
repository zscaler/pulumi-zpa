// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-privileged-consoles)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-privileged-consoles-using-api)
 *
 * The **zpa_pra_console_controller** resource creates a privileged remote access console in the Zscaler Private Access cloud. This resource can then be referenced in an privileged access policy resource and a privileged access portal.
 *
 * ## Import
 *
 * Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
 *
 * Visit
 *
 * **pra_credential_controller** can be imported by using `<CONSOLE ID>` or `<CONSOLE NAME>` as the import ID.
 *
 * For example:
 *
 * ```sh
 * $ pulumi import zpa:index/pRAConsole:PRAConsole this <console_id>
 * ```
 *
 * or
 *
 * ```sh
 * $ pulumi import zpa:index/pRAConsole:PRAConsole this <console_name>
 * ```
 */
export class PRAConsole extends pulumi.CustomResource {
    /**
     * Get an existing PRAConsole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PRAConsoleState, opts?: pulumi.CustomResourceOptions): PRAConsole {
        return new PRAConsole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zpa:index/pRAConsole:PRAConsole';

    /**
     * Returns true if the given object is an instance of PRAConsole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PRAConsole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PRAConsole.__pulumiType;
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
    public readonly praApplication!: pulumi.Output<outputs.PRAConsolePraApplication>;
    public readonly praPortals!: pulumi.Output<outputs.PRAConsolePraPortal[]>;

    /**
     * Create a PRAConsole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PRAConsoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PRAConsoleArgs | PRAConsoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PRAConsoleState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["iconText"] = state ? state.iconText : undefined;
            resourceInputs["microtenantId"] = state ? state.microtenantId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["praApplication"] = state ? state.praApplication : undefined;
            resourceInputs["praPortals"] = state ? state.praPortals : undefined;
        } else {
            const args = argsOrState as PRAConsoleArgs | undefined;
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
        const aliasOpts = { aliases: [{ type: "zpa:index/praConsoleController:PraConsoleController" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(PRAConsole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PRAConsole resources.
 */
export interface PRAConsoleState {
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
    praApplication?: pulumi.Input<inputs.PRAConsolePraApplication>;
    praPortals?: pulumi.Input<pulumi.Input<inputs.PRAConsolePraPortal>[]>;
}

/**
 * The set of arguments for constructing a PRAConsole resource.
 */
export interface PRAConsoleArgs {
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
    praApplication: pulumi.Input<inputs.PRAConsolePraApplication>;
    praPortals: pulumi.Input<pulumi.Input<inputs.PRAConsolePraPortal>[]>;
}