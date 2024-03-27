// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 * import * as zpa from "@pulumi/zpa";
 *
 * const testCert = zpa.getBaCertificate({
 *     name: "sales.acme.com",
 * });
 * // ZPA Segment Group resource
 * const exampleSegmentGroup = new zpa.SegmentGroup("exampleSegmentGroup", {
 *     description: "Example",
 *     enabled: true,
 * });
 * const exampleAppConnectorGroup = zpa.getAppConnectorGroup({
 *     name: "AWS-Connector",
 * });
 * // ZPA Server Group resource
 * const exampleServerGroup = new zpa.ServerGroup("exampleServerGroup", {
 *     description: "Example",
 *     enabled: true,
 *     dynamicDiscovery: true,
 *     appConnectorGroups: [{
 *         ids: [exampleAppConnectorGroup.then(exampleAppConnectorGroup => exampleAppConnectorGroup.id)],
 *     }],
 * });
 * // Create Browser Access Application
 * const browserAccessApps = new zpa.ApplicationSegmentBrowserAccess("browserAccessApps", {
 *     description: "Browser Access Apps",
 *     enabled: true,
 *     healthReporting: "ON_ACCESS",
 *     bypassType: "NEVER",
 *     tcpPortRanges: [
 *         "80",
 *         "80",
 *     ],
 *     domainNames: ["sales.acme.com"],
 *     segmentGroupId: exampleSegmentGroup.id,
 *     clientlessApps: [{
 *         name: "sales.acme.com",
 *         applicationProtocol: "HTTP",
 *         applicationPort: "80",
 *         certificateId: testCert.then(testCert => testCert.id),
 *         trustUntrustedCert: true,
 *         enabled: true,
 *         domain: "sales.acme.com",
 *     }],
 *     serverGroups: [{
 *         ids: [exampleServerGroup.id],
 *     }],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
 *
 * Visit
 *
 * **zpa_application_segment_browser_access** Application Segment Browser Access can be imported by using <`BROWSER ACCESS ID`> or `<<BROWSER ACCESS NAME>` as the import ID.
 *
 * For example:
 *
 * ```sh
 * $ pulumi import zpa:index/applicationSegmentBrowserAccess:ApplicationSegmentBrowserAccess example <browser_access_id>.
 * ```
 *
 * or
 *
 * ```sh
 * $ pulumi import zpa:index/applicationSegmentBrowserAccess:ApplicationSegmentBrowserAccess example <browser_access_name>
 * ```
 */
export class ApplicationSegmentBrowserAccess extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationSegmentBrowserAccess resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationSegmentBrowserAccessState, opts?: pulumi.CustomResourceOptions): ApplicationSegmentBrowserAccess {
        return new ApplicationSegmentBrowserAccess(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zpa:index/applicationSegmentBrowserAccess:ApplicationSegmentBrowserAccess';

    /**
     * Returns true if the given object is an instance of ApplicationSegmentBrowserAccess.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationSegmentBrowserAccess {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationSegmentBrowserAccess.__pulumiType;
    }

    /**
     * (Optional) Indicates whether users can bypass ZPA to access applications. Default value is: `NEVER` and supported values are: `ALWAYS`, `NEVER` and `ON_NET`. The value `NEVER` indicates the use of the client forwarding policy.
     */
    public readonly bypassType!: pulumi.Output<string | undefined>;
    public readonly clientlessApps!: pulumi.Output<outputs.ApplicationSegmentBrowserAccessClientlessApp[]>;
    /**
     * (Optional) Default: `DEFAULT`. Supported values: `DEFAULT`, `SIEM`
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
    public readonly doubleEncrypt!: pulumi.Output<boolean | undefined>;
    /**
     * (Optional) - Whether this app is enabled or not.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * (Optional) Default: `DEFAULT`. Supported values: `DEFAULT`, `NONE`
     */
    public readonly healthCheckType!: pulumi.Output<string>;
    /**
     * (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
     */
    public readonly healthReporting!: pulumi.Output<string>;
    public readonly icmpAccessType!: pulumi.Output<string | undefined>;
    /**
     * (Optional) - If Source IP Anchoring for use with ZIA, is enabled or disabled for the app. Supported values are `true` and `false`
     */
    public readonly ipAnchored!: pulumi.Output<boolean | undefined>;
    /**
     * (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
     */
    public readonly isCnameEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * (Optional) Supported values: `true`, `false`
     */
    public readonly isIncompleteDrConfig!: pulumi.Output<boolean | undefined>;
    /**
     * Name of BA app.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly passiveHealthEnabled!: pulumi.Output<boolean>;
    /**
     * List of Segment Group IDs
     */
    public readonly segmentGroupId!: pulumi.Output<string>;
    public readonly segmentGroupName!: pulumi.Output<string>;
    /**
     * (Optional) Supported values: `true`, `false`
     */
    public readonly selectConnectorCloseToApp!: pulumi.Output<boolean | undefined>;
    /**
     * List of Server Group IDs
     */
    public readonly serverGroups!: pulumi.Output<outputs.ApplicationSegmentBrowserAccessServerGroup[]>;
    /**
     * (Optional) Supported values: ``1`` for Enabled and ``0`` for Disabled
     */
    public readonly tcpKeepAlive!: pulumi.Output<string>;
    /**
     * tcp port range
     */
    public readonly tcpPortRange!: pulumi.Output<outputs.ApplicationSegmentBrowserAccessTcpPortRange[]>;
    /**
     * TCP port ranges used to access the app.
     */
    public readonly tcpPortRanges!: pulumi.Output<string[]>;
    /**
     * udp port range
     */
    public readonly udpPortRange!: pulumi.Output<outputs.ApplicationSegmentBrowserAccessUdpPortRange[]>;
    /**
     * UDP port ranges used to access the app.
     */
    public readonly udpPortRanges!: pulumi.Output<string[]>;
    /**
     * (Optional) Supported values: `true`, `false`
     */
    public readonly useInDrMode!: pulumi.Output<boolean | undefined>;

    /**
     * Create a ApplicationSegmentBrowserAccess resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationSegmentBrowserAccessArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationSegmentBrowserAccessArgs | ApplicationSegmentBrowserAccessState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationSegmentBrowserAccessState | undefined;
            resourceInputs["bypassType"] = state ? state.bypassType : undefined;
            resourceInputs["clientlessApps"] = state ? state.clientlessApps : undefined;
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
            resourceInputs["isIncompleteDrConfig"] = state ? state.isIncompleteDrConfig : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["passiveHealthEnabled"] = state ? state.passiveHealthEnabled : undefined;
            resourceInputs["segmentGroupId"] = state ? state.segmentGroupId : undefined;
            resourceInputs["segmentGroupName"] = state ? state.segmentGroupName : undefined;
            resourceInputs["selectConnectorCloseToApp"] = state ? state.selectConnectorCloseToApp : undefined;
            resourceInputs["serverGroups"] = state ? state.serverGroups : undefined;
            resourceInputs["tcpKeepAlive"] = state ? state.tcpKeepAlive : undefined;
            resourceInputs["tcpPortRange"] = state ? state.tcpPortRange : undefined;
            resourceInputs["tcpPortRanges"] = state ? state.tcpPortRanges : undefined;
            resourceInputs["udpPortRange"] = state ? state.udpPortRange : undefined;
            resourceInputs["udpPortRanges"] = state ? state.udpPortRanges : undefined;
            resourceInputs["useInDrMode"] = state ? state.useInDrMode : undefined;
        } else {
            const args = argsOrState as ApplicationSegmentBrowserAccessArgs | undefined;
            if ((!args || args.clientlessApps === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientlessApps'");
            }
            if ((!args || args.domainNames === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainNames'");
            }
            if ((!args || args.segmentGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'segmentGroupId'");
            }
            resourceInputs["bypassType"] = args ? args.bypassType : undefined;
            resourceInputs["clientlessApps"] = args ? args.clientlessApps : undefined;
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
            resourceInputs["isIncompleteDrConfig"] = args ? args.isIncompleteDrConfig : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["passiveHealthEnabled"] = args ? args.passiveHealthEnabled : undefined;
            resourceInputs["segmentGroupId"] = args ? args.segmentGroupId : undefined;
            resourceInputs["segmentGroupName"] = args ? args.segmentGroupName : undefined;
            resourceInputs["selectConnectorCloseToApp"] = args ? args.selectConnectorCloseToApp : undefined;
            resourceInputs["serverGroups"] = args ? args.serverGroups : undefined;
            resourceInputs["tcpKeepAlive"] = args ? args.tcpKeepAlive : undefined;
            resourceInputs["tcpPortRange"] = args ? args.tcpPortRange : undefined;
            resourceInputs["tcpPortRanges"] = args ? args.tcpPortRanges : undefined;
            resourceInputs["udpPortRange"] = args ? args.udpPortRange : undefined;
            resourceInputs["udpPortRanges"] = args ? args.udpPortRanges : undefined;
            resourceInputs["useInDrMode"] = args ? args.useInDrMode : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApplicationSegmentBrowserAccess.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationSegmentBrowserAccess resources.
 */
export interface ApplicationSegmentBrowserAccessState {
    /**
     * (Optional) Indicates whether users can bypass ZPA to access applications. Default value is: `NEVER` and supported values are: `ALWAYS`, `NEVER` and `ON_NET`. The value `NEVER` indicates the use of the client forwarding policy.
     */
    bypassType?: pulumi.Input<string>;
    clientlessApps?: pulumi.Input<pulumi.Input<inputs.ApplicationSegmentBrowserAccessClientlessApp>[]>;
    /**
     * (Optional) Default: `DEFAULT`. Supported values: `DEFAULT`, `SIEM`
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
     * (Optional) - Whether this app is enabled or not.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * (Optional) Default: `DEFAULT`. Supported values: `DEFAULT`, `NONE`
     */
    healthCheckType?: pulumi.Input<string>;
    /**
     * (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
     */
    healthReporting?: pulumi.Input<string>;
    icmpAccessType?: pulumi.Input<string>;
    /**
     * (Optional) - If Source IP Anchoring for use with ZIA, is enabled or disabled for the app. Supported values are `true` and `false`
     */
    ipAnchored?: pulumi.Input<boolean>;
    /**
     * (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
     */
    isCnameEnabled?: pulumi.Input<boolean>;
    /**
     * (Optional) Supported values: `true`, `false`
     */
    isIncompleteDrConfig?: pulumi.Input<boolean>;
    /**
     * Name of BA app.
     */
    name?: pulumi.Input<string>;
    passiveHealthEnabled?: pulumi.Input<boolean>;
    /**
     * List of Segment Group IDs
     */
    segmentGroupId?: pulumi.Input<string>;
    segmentGroupName?: pulumi.Input<string>;
    /**
     * (Optional) Supported values: `true`, `false`
     */
    selectConnectorCloseToApp?: pulumi.Input<boolean>;
    /**
     * List of Server Group IDs
     */
    serverGroups?: pulumi.Input<pulumi.Input<inputs.ApplicationSegmentBrowserAccessServerGroup>[]>;
    /**
     * (Optional) Supported values: ``1`` for Enabled and ``0`` for Disabled
     */
    tcpKeepAlive?: pulumi.Input<string>;
    /**
     * tcp port range
     */
    tcpPortRange?: pulumi.Input<pulumi.Input<inputs.ApplicationSegmentBrowserAccessTcpPortRange>[]>;
    /**
     * TCP port ranges used to access the app.
     */
    tcpPortRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * udp port range
     */
    udpPortRange?: pulumi.Input<pulumi.Input<inputs.ApplicationSegmentBrowserAccessUdpPortRange>[]>;
    /**
     * UDP port ranges used to access the app.
     */
    udpPortRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) Supported values: `true`, `false`
     */
    useInDrMode?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ApplicationSegmentBrowserAccess resource.
 */
export interface ApplicationSegmentBrowserAccessArgs {
    /**
     * (Optional) Indicates whether users can bypass ZPA to access applications. Default value is: `NEVER` and supported values are: `ALWAYS`, `NEVER` and `ON_NET`. The value `NEVER` indicates the use of the client forwarding policy.
     */
    bypassType?: pulumi.Input<string>;
    clientlessApps: pulumi.Input<pulumi.Input<inputs.ApplicationSegmentBrowserAccessClientlessApp>[]>;
    /**
     * (Optional) Default: `DEFAULT`. Supported values: `DEFAULT`, `SIEM`
     */
    configSpace?: pulumi.Input<string>;
    /**
     * (Optional) Description of the application.
     */
    description?: pulumi.Input<string>;
    /**
     * List of domains and IPs.
     */
    domainNames: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) Whether Double Encryption is enabled or disabled for the app.
     */
    doubleEncrypt?: pulumi.Input<boolean>;
    /**
     * (Optional) - Whether this app is enabled or not.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * (Optional) Default: `DEFAULT`. Supported values: `DEFAULT`, `NONE`
     */
    healthCheckType?: pulumi.Input<string>;
    /**
     * (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
     */
    healthReporting?: pulumi.Input<string>;
    icmpAccessType?: pulumi.Input<string>;
    /**
     * (Optional) - If Source IP Anchoring for use with ZIA, is enabled or disabled for the app. Supported values are `true` and `false`
     */
    ipAnchored?: pulumi.Input<boolean>;
    /**
     * (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
     */
    isCnameEnabled?: pulumi.Input<boolean>;
    /**
     * (Optional) Supported values: `true`, `false`
     */
    isIncompleteDrConfig?: pulumi.Input<boolean>;
    /**
     * Name of BA app.
     */
    name?: pulumi.Input<string>;
    passiveHealthEnabled?: pulumi.Input<boolean>;
    /**
     * List of Segment Group IDs
     */
    segmentGroupId: pulumi.Input<string>;
    segmentGroupName?: pulumi.Input<string>;
    /**
     * (Optional) Supported values: `true`, `false`
     */
    selectConnectorCloseToApp?: pulumi.Input<boolean>;
    /**
     * List of Server Group IDs
     */
    serverGroups?: pulumi.Input<pulumi.Input<inputs.ApplicationSegmentBrowserAccessServerGroup>[]>;
    /**
     * (Optional) Supported values: ``1`` for Enabled and ``0`` for Disabled
     */
    tcpKeepAlive?: pulumi.Input<string>;
    /**
     * tcp port range
     */
    tcpPortRange?: pulumi.Input<pulumi.Input<inputs.ApplicationSegmentBrowserAccessTcpPortRange>[]>;
    /**
     * TCP port ranges used to access the app.
     */
    tcpPortRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * udp port range
     */
    udpPortRange?: pulumi.Input<pulumi.Input<inputs.ApplicationSegmentBrowserAccessUdpPortRange>[]>;
    /**
     * UDP port ranges used to access the app.
     */
    udpPortRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) Supported values: `true`, `false`
     */
    useInDrMode?: pulumi.Input<boolean>;
}