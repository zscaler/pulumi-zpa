// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-browser-protection-profiles)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-appprotection-profiles-using-api)
 *
 * The  **zpa_inspection_profile** resource creates an inspection profile in the Zscaler Private Access cloud. This resource can then be referenced in an inspection custom control resource.
 */
export class InspectionProfile extends pulumi.CustomResource {
    /**
     * Get an existing InspectionProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InspectionProfileState, opts?: pulumi.CustomResourceOptions): InspectionProfile {
        return new InspectionProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zpa:index/inspectionProfile:InspectionProfile';

    /**
     * Returns true if the given object is an instance of InspectionProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InspectionProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InspectionProfile.__pulumiType;
    }

    public readonly associateAllControls!: pulumi.Output<boolean | undefined>;
    public readonly controlsInfos!: pulumi.Output<outputs.InspectionProfileControlsInfo[]>;
    /**
     * The set of AppProtection controls used to define how inspections are managed
     */
    public readonly customControls!: pulumi.Output<outputs.InspectionProfileCustomControl[] | undefined>;
    /**
     * The description of the AppProtection profile
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The actions of the predefined, custom, or override controls
     */
    public readonly globalControlActions!: pulumi.Output<string[] | undefined>;
    public readonly incarnationNumber!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The OWASP Predefined Paranoia Level
     */
    public readonly paranoiaLevel!: pulumi.Output<string | undefined>;
    /**
     * The predefined controls
     */
    public readonly predefinedControls!: pulumi.Output<outputs.InspectionProfilePredefinedControl[]>;
    /**
     * The protocol for the AppProtection application
     */
    public readonly predefinedControlsVersion!: pulumi.Output<string | undefined>;
    /**
     * Indicates the user's choice for the ThreatLabZ Controls. Supported values: ALL and SPECIFIC
     */
    public readonly zsDefinedControlChoice!: pulumi.Output<string | undefined>;

    /**
     * Create a InspectionProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InspectionProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InspectionProfileArgs | InspectionProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InspectionProfileState | undefined;
            resourceInputs["associateAllControls"] = state ? state.associateAllControls : undefined;
            resourceInputs["controlsInfos"] = state ? state.controlsInfos : undefined;
            resourceInputs["customControls"] = state ? state.customControls : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["globalControlActions"] = state ? state.globalControlActions : undefined;
            resourceInputs["incarnationNumber"] = state ? state.incarnationNumber : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["paranoiaLevel"] = state ? state.paranoiaLevel : undefined;
            resourceInputs["predefinedControls"] = state ? state.predefinedControls : undefined;
            resourceInputs["predefinedControlsVersion"] = state ? state.predefinedControlsVersion : undefined;
            resourceInputs["zsDefinedControlChoice"] = state ? state.zsDefinedControlChoice : undefined;
        } else {
            const args = argsOrState as InspectionProfileArgs | undefined;
            resourceInputs["associateAllControls"] = args ? args.associateAllControls : undefined;
            resourceInputs["controlsInfos"] = args ? args.controlsInfos : undefined;
            resourceInputs["customControls"] = args ? args.customControls : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["globalControlActions"] = args ? args.globalControlActions : undefined;
            resourceInputs["incarnationNumber"] = args ? args.incarnationNumber : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["paranoiaLevel"] = args ? args.paranoiaLevel : undefined;
            resourceInputs["predefinedControls"] = args ? args.predefinedControls : undefined;
            resourceInputs["predefinedControlsVersion"] = args ? args.predefinedControlsVersion : undefined;
            resourceInputs["zsDefinedControlChoice"] = args ? args.zsDefinedControlChoice : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InspectionProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InspectionProfile resources.
 */
export interface InspectionProfileState {
    associateAllControls?: pulumi.Input<boolean>;
    controlsInfos?: pulumi.Input<pulumi.Input<inputs.InspectionProfileControlsInfo>[]>;
    /**
     * The set of AppProtection controls used to define how inspections are managed
     */
    customControls?: pulumi.Input<pulumi.Input<inputs.InspectionProfileCustomControl>[]>;
    /**
     * The description of the AppProtection profile
     */
    description?: pulumi.Input<string>;
    /**
     * The actions of the predefined, custom, or override controls
     */
    globalControlActions?: pulumi.Input<pulumi.Input<string>[]>;
    incarnationNumber?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * The OWASP Predefined Paranoia Level
     */
    paranoiaLevel?: pulumi.Input<string>;
    /**
     * The predefined controls
     */
    predefinedControls?: pulumi.Input<pulumi.Input<inputs.InspectionProfilePredefinedControl>[]>;
    /**
     * The protocol for the AppProtection application
     */
    predefinedControlsVersion?: pulumi.Input<string>;
    /**
     * Indicates the user's choice for the ThreatLabZ Controls. Supported values: ALL and SPECIFIC
     */
    zsDefinedControlChoice?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InspectionProfile resource.
 */
export interface InspectionProfileArgs {
    associateAllControls?: pulumi.Input<boolean>;
    controlsInfos?: pulumi.Input<pulumi.Input<inputs.InspectionProfileControlsInfo>[]>;
    /**
     * The set of AppProtection controls used to define how inspections are managed
     */
    customControls?: pulumi.Input<pulumi.Input<inputs.InspectionProfileCustomControl>[]>;
    /**
     * The description of the AppProtection profile
     */
    description?: pulumi.Input<string>;
    /**
     * The actions of the predefined, custom, or override controls
     */
    globalControlActions?: pulumi.Input<pulumi.Input<string>[]>;
    incarnationNumber?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * The OWASP Predefined Paranoia Level
     */
    paranoiaLevel?: pulumi.Input<string>;
    /**
     * The predefined controls
     */
    predefinedControls?: pulumi.Input<pulumi.Input<inputs.InspectionProfilePredefinedControl>[]>;
    /**
     * The protocol for the AppProtection application
     */
    predefinedControlsVersion?: pulumi.Input<string>;
    /**
     * Indicates the user's choice for the ThreatLabZ Controls. Supported values: ALL and SPECIFIC
     */
    zsDefinedControlChoice?: pulumi.Input<string>;
}
