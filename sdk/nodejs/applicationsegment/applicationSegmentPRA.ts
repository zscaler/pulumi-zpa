// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The **zpa_application_segment_pra** resource creates an application segment for Privileged Remote Access in the Zscaler Private Access cloud. This resource can then be referenced in an access policy rule, access policy timeout rule, access policy client forwarding rule and inspection policy. This resource supports Privileged Remote Access for both `RDP` and `SSH`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@zscaler/pulumi-zpa";
 *
 * const _this = new zpa.applicationsegment.ApplicationSegmentPRA("this", {
 *     description: "PRA_Example",
 *     enabled: true,
 *     healthReporting: "ON_ACCESS",
 *     bypassType: "NEVER",
 *     isCnameEnabled: true,
 *     tcpPortRanges: [
 *         "22",
 *         "22",
 *         "3389",
 *         "3389",
 *     ],
 *     domainNames: [
 *         "ssh_pra.example.com",
 *         "rdp_pra.example.com",
 *     ],
 *     segmentGroupId: zpa_segment_group["this"].id,
 *     serverGroups: [{
 *         ids: [zpa_server_group["this"].id],
 *     }],
 *     commonAppsDto: {
 *         appsConfigs: [
 *             {
 *                 name: "ssh_pra",
 *                 domain: "ssh_pra.example.com",
 *                 applicationProtocol: "SSH",
 *                 applicationPort: "22",
 *                 enabled: true,
 *                 appTypes: ["SECURE_REMOTE_ACCESS"],
 *             },
 *             {
 *                 name: "rdp_pra",
 *                 domain: "rdp_pra.example.com",
 *                 applicationProtocol: "RDP",
 *                 connectionSecurity: "ANY",
 *                 applicationPort: "3389",
 *                 enabled: true,
 *                 appTypes: ["SECURE_REMOTE_ACCESS"],
 *             },
 *         ],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Application Segment can be imported by using `<APPLICATION SEGMENT ID>` or `<APPLICATION SEGMENT NAME>` as the import ID.
 *
 * ```sh
 *  $ pulumi import zpa:ApplicationSegment/applicationSegmentPRA:ApplicationSegmentPRA example <application_segment_id>
 * ```
 *
 *  or
 *
 * ```sh
 *  $ pulumi import zpa:ApplicationSegment/applicationSegmentPRA:ApplicationSegmentPRA example <application_segment_name>
 * ```
 */
export class ApplicationSegmentPRA extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationSegmentPRA resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationSegmentPRAState, opts?: pulumi.CustomResourceOptions): ApplicationSegmentPRA {
        return new ApplicationSegmentPRA(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zpa:ApplicationSegment/applicationSegmentPRA:ApplicationSegmentPRA';

    /**
     * Returns true if the given object is an instance of ApplicationSegmentPRA.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationSegmentPRA {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationSegmentPRA.__pulumiType;
    }

    /**
     * (Optional) Indicates whether users can bypass ZPA to access applications.
     */
    public readonly bypassType!: pulumi.Output<string>;
    /**
     * List of applications (e.g., Inspection, Browser Access or Privileged Remote Access)
     * * `apps_config:` - (Required) List of applications to be configured
     */
    public readonly commonAppsDto!: pulumi.Output<outputs.ApplicationSegment.ApplicationSegmentPRACommonAppsDto>;
    /**
     * (Optional)
     */
    public readonly configSpace!: pulumi.Output<string | undefined>;
    /**
     * (Optional) Description of the application.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * List of domains and IPs.
     */
    public readonly domainNames!: pulumi.Output<string[]>;
    /**
     * (Optional) Whether Double Encryption is enabled or disabled for the app.
     */
    public readonly doubleEncrypt!: pulumi.Output<boolean>;
    /**
     * Whether this application is enabled or not
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * (Optional)
     */
    public readonly healthCheckType!: pulumi.Output<string | undefined>;
    /**
     * (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
     */
    public readonly healthReporting!: pulumi.Output<string | undefined>;
    /**
     * (Optional)
     */
    public readonly icmpAccessType!: pulumi.Output<string>;
    /**
     * (Optional)
     */
    public readonly ipAnchored!: pulumi.Output<boolean>;
    /**
     * (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
     */
    public readonly isCnameEnabled!: pulumi.Output<boolean>;
    /**
     * Name of the Privileged Remote Access
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * (Optional)
     */
    public readonly passiveHealthEnabled!: pulumi.Output<boolean>;
    /**
     * List of Segment Group IDs
     */
    public readonly segmentGroupId!: pulumi.Output<string>;
    public readonly segmentGroupName!: pulumi.Output<string>;
    /**
     * List of Server Group IDs
     */
    public readonly serverGroups!: pulumi.Output<outputs.ApplicationSegment.ApplicationSegmentPRAServerGroup[]>;
    /**
     * TCP port ranges used to access the app.
     */
    public readonly tcpPortRanges!: pulumi.Output<string[]>;
    /**
     * UDP port ranges used to access the app.
     */
    public readonly udpPortRanges!: pulumi.Output<string[]>;

    /**
     * Create a ApplicationSegmentPRA resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationSegmentPRAArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationSegmentPRAArgs | ApplicationSegmentPRAState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationSegmentPRAState | undefined;
            resourceInputs["bypassType"] = state ? state.bypassType : undefined;
            resourceInputs["commonAppsDto"] = state ? state.commonAppsDto : undefined;
            resourceInputs["configSpace"] = state ? state.configSpace : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["domainNames"] = state ? state.domainNames : undefined;
            resourceInputs["doubleEncrypt"] = state ? state.doubleEncrypt : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["healthCheckType"] = state ? state.healthCheckType : undefined;
            resourceInputs["healthReporting"] = state ? state.healthReporting : undefined;
            resourceInputs["icmpAccessType"] = state ? state.icmpAccessType : undefined;
            resourceInputs["ipAnchored"] = state ? state.ipAnchored : undefined;
            resourceInputs["isCnameEnabled"] = state ? state.isCnameEnabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["passiveHealthEnabled"] = state ? state.passiveHealthEnabled : undefined;
            resourceInputs["segmentGroupId"] = state ? state.segmentGroupId : undefined;
            resourceInputs["segmentGroupName"] = state ? state.segmentGroupName : undefined;
            resourceInputs["serverGroups"] = state ? state.serverGroups : undefined;
            resourceInputs["tcpPortRanges"] = state ? state.tcpPortRanges : undefined;
            resourceInputs["udpPortRanges"] = state ? state.udpPortRanges : undefined;
        } else {
            const args = argsOrState as ApplicationSegmentPRAArgs | undefined;
            if ((!args || args.segmentGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'segmentGroupId'");
            }
            if ((!args || args.serverGroups === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverGroups'");
            }
            resourceInputs["bypassType"] = args ? args.bypassType : undefined;
            resourceInputs["commonAppsDto"] = args ? args.commonAppsDto : undefined;
            resourceInputs["configSpace"] = args ? args.configSpace : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["domainNames"] = args ? args.domainNames : undefined;
            resourceInputs["doubleEncrypt"] = args ? args.doubleEncrypt : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["healthCheckType"] = args ? args.healthCheckType : undefined;
            resourceInputs["healthReporting"] = args ? args.healthReporting : undefined;
            resourceInputs["icmpAccessType"] = args ? args.icmpAccessType : undefined;
            resourceInputs["ipAnchored"] = args ? args.ipAnchored : undefined;
            resourceInputs["isCnameEnabled"] = args ? args.isCnameEnabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["passiveHealthEnabled"] = args ? args.passiveHealthEnabled : undefined;
            resourceInputs["segmentGroupId"] = args ? args.segmentGroupId : undefined;
            resourceInputs["segmentGroupName"] = args ? args.segmentGroupName : undefined;
            resourceInputs["serverGroups"] = args ? args.serverGroups : undefined;
            resourceInputs["tcpPortRanges"] = args ? args.tcpPortRanges : undefined;
            resourceInputs["udpPortRanges"] = args ? args.udpPortRanges : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApplicationSegmentPRA.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationSegmentPRA resources.
 */
export interface ApplicationSegmentPRAState {
    /**
     * (Optional) Indicates whether users can bypass ZPA to access applications.
     */
    bypassType?: pulumi.Input<string>;
    /**
     * List of applications (e.g., Inspection, Browser Access or Privileged Remote Access)
     * * `apps_config:` - (Required) List of applications to be configured
     */
    commonAppsDto?: pulumi.Input<inputs.ApplicationSegment.ApplicationSegmentPRACommonAppsDto>;
    /**
     * (Optional)
     */
    configSpace?: pulumi.Input<string>;
    /**
     * (Optional) Description of the application.
     */
    description?: pulumi.Input<string>;
    /**
     * List of domains and IPs.
     */
    domainNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) Whether Double Encryption is enabled or disabled for the app.
     */
    doubleEncrypt?: pulumi.Input<boolean>;
    /**
     * Whether this application is enabled or not
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * (Optional)
     */
    healthCheckType?: pulumi.Input<string>;
    /**
     * (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
     */
    healthReporting?: pulumi.Input<string>;
    /**
     * (Optional)
     */
    icmpAccessType?: pulumi.Input<string>;
    /**
     * (Optional)
     */
    ipAnchored?: pulumi.Input<boolean>;
    /**
     * (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
     */
    isCnameEnabled?: pulumi.Input<boolean>;
    /**
     * Name of the Privileged Remote Access
     */
    name?: pulumi.Input<string>;
    /**
     * (Optional)
     */
    passiveHealthEnabled?: pulumi.Input<boolean>;
    /**
     * List of Segment Group IDs
     */
    segmentGroupId?: pulumi.Input<string>;
    segmentGroupName?: pulumi.Input<string>;
    /**
     * List of Server Group IDs
     */
    serverGroups?: pulumi.Input<pulumi.Input<inputs.ApplicationSegment.ApplicationSegmentPRAServerGroup>[]>;
    /**
     * TCP port ranges used to access the app.
     */
    tcpPortRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * UDP port ranges used to access the app.
     */
    udpPortRanges?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a ApplicationSegmentPRA resource.
 */
export interface ApplicationSegmentPRAArgs {
    /**
     * (Optional) Indicates whether users can bypass ZPA to access applications.
     */
    bypassType?: pulumi.Input<string>;
    /**
     * List of applications (e.g., Inspection, Browser Access or Privileged Remote Access)
     * * `apps_config:` - (Required) List of applications to be configured
     */
    commonAppsDto?: pulumi.Input<inputs.ApplicationSegment.ApplicationSegmentPRACommonAppsDto>;
    /**
     * (Optional)
     */
    configSpace?: pulumi.Input<string>;
    /**
     * (Optional) Description of the application.
     */
    description?: pulumi.Input<string>;
    /**
     * List of domains and IPs.
     */
    domainNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) Whether Double Encryption is enabled or disabled for the app.
     */
    doubleEncrypt?: pulumi.Input<boolean>;
    /**
     * Whether this application is enabled or not
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * (Optional)
     */
    healthCheckType?: pulumi.Input<string>;
    /**
     * (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
     */
    healthReporting?: pulumi.Input<string>;
    /**
     * (Optional)
     */
    icmpAccessType?: pulumi.Input<string>;
    /**
     * (Optional)
     */
    ipAnchored?: pulumi.Input<boolean>;
    /**
     * (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
     */
    isCnameEnabled?: pulumi.Input<boolean>;
    /**
     * Name of the Privileged Remote Access
     */
    name?: pulumi.Input<string>;
    /**
     * (Optional)
     */
    passiveHealthEnabled?: pulumi.Input<boolean>;
    /**
     * List of Segment Group IDs
     */
    segmentGroupId: pulumi.Input<string>;
    segmentGroupName?: pulumi.Input<string>;
    /**
     * List of Server Group IDs
     */
    serverGroups: pulumi.Input<pulumi.Input<inputs.ApplicationSegment.ApplicationSegmentPRAServerGroup>[]>;
    /**
     * TCP port ranges used to access the app.
     */
    tcpPortRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * UDP port ranges used to access the app.
     */
    udpPortRanges?: pulumi.Input<pulumi.Input<string>[]>;
}
