// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zpa/about-privileged-portals)
 * * [API documentation](https://help.zscaler.com/zpa/configuring-privileged-portals-using-api)
 *
 * The **zpa_pra_portal_controller** resource creates a privileged remote access portal in the Zscaler Private Access cloud. This resource can then be referenced in an privileged remote access console resource.
 *
 * ## Example Usage
 *
 * ### Using Custom Certificate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const thisBaCertificate = zpa.getBaCertificate({
 *     name: "portal.acme.com",
 * });
 * const thisPRAPortal = new zpa.PRAPortal("thisPRAPortal", {
 *     description: "portal.acme.com",
 *     enabled: true,
 *     domain: "portal.acme.com",
 *     certificateId: thisBaCertificate.then(thisBaCertificate => thisBaCertificate.id),
 *     userNotification: "Created with Terraform",
 *     userNotificationEnabled: true,
 * });
 * ```
 *
 * ### Zscaler Managed Certificate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const _this = new zpa.PRAPortal("this", {
 *     description: "server1.acme.com",
 *     domain: "server1-acme.com.pra.d.zscalerportal.net",
 *     enabled: true,
 *     extDomain: "acme.com",
 *     extDomainName: "acme.com.pra.d.zscalerportal.net",
 *     extDomainTranslation: "acme.com",
 *     extLabel: "server1",
 *     userNotification: "Created with Terraform",
 *     userNotificationEnabled: true,
 * });
 * ```
 *
 * # Configuring PRA Portal with User Portal
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@bdzscaler/pulumi-zpa";
 *
 * const _this = new zpa.PRAPortal("this", {
 *     description: "Server1 PRA01 Description",
 *     domain: "server1-acme.com.pra.d.zscalerportal.net",
 *     enabled: true,
 *     extDomain: "acme.com",
 *     extDomainName: "acme.com.pra.d.zscalerportal.net",
 *     extDomainTranslation: "acme.com",
 *     extLabel: "server1",
 *     userNotification: "Created with Terraform",
 *     userNotificationEnabled: true,
 *     userPortalGid: "145262059234265326",
 * });
 * ```
 *
 * ## Import
 *
 * Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
 *
 * Visit
 *
 * **pra_portal_controller** can be imported by using `<PORTAL ID>` or `<PORTAL NAME>` as the import ID.
 *
 * For example:
 *
 * ```sh
 * $ pulumi import zpa:index/pRAPortal:PRAPortal this <portal_id>
 * ```
 *
 * or
 *
 * ```sh
 * $ pulumi import zpa:index/pRAPortal:PRAPortal this <portal_name>
 * ```
 */
export class PRAPortal extends pulumi.CustomResource {
    /**
     * Get an existing PRAPortal resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PRAPortalState, opts?: pulumi.CustomResourceOptions): PRAPortal {
        return new PRAPortal(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zpa:index/pRAPortal:PRAPortal';

    /**
     * Returns true if the given object is an instance of PRAPortal.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PRAPortal {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PRAPortal.__pulumiType;
    }

    /**
     * The unique identifier of the certificate
     */
    public readonly certificateId!: pulumi.Output<string>;
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
     * The external domain name prefix of the Browser Access application that is used for Zscaler-managed certificates when
     * creating a privileged portal.
     */
    public readonly extDomain!: pulumi.Output<string | undefined>;
    /**
     * The domain suffix for the privileged portal URL. This field must be one of the customer's authentication domains.
     */
    public readonly extDomainName!: pulumi.Output<string | undefined>;
    /**
     * The translation of the external domain name prefix of the Browser Access application that is used for Zscaler-managed
     * certificates when creating a privileged portal.
     */
    public readonly extDomainTranslation!: pulumi.Output<string | undefined>;
    /**
     * The domain prefix for the privileged portal URL. The supported string can include numbers, lower case characters, and
     * only supports a hyphen (-).
     */
    public readonly extLabel!: pulumi.Output<string | undefined>;
    /**
     * The unique identifier of the Microtenant for the ZPA tenant. If you are within the Default Microtenant, pass
     * microtenantId as 0 when making requests to retrieve data from the Default Microtenant. Pass microtenantId as null to
     * retrieve data from all customers associated with the tenant.
     */
    public readonly microtenantId!: pulumi.Output<string | undefined>;
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
     * The unique identifier of the user portal.
     */
    public readonly userPortalGid!: pulumi.Output<string | undefined>;

    /**
     * Create a PRAPortal resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PRAPortalArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PRAPortalArgs | PRAPortalState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PRAPortalState | undefined;
            resourceInputs["certificateId"] = state ? state.certificateId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["extDomain"] = state ? state.extDomain : undefined;
            resourceInputs["extDomainName"] = state ? state.extDomainName : undefined;
            resourceInputs["extDomainTranslation"] = state ? state.extDomainTranslation : undefined;
            resourceInputs["extLabel"] = state ? state.extLabel : undefined;
            resourceInputs["microtenantId"] = state ? state.microtenantId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["userNotification"] = state ? state.userNotification : undefined;
            resourceInputs["userNotificationEnabled"] = state ? state.userNotificationEnabled : undefined;
            resourceInputs["userPortalGid"] = state ? state.userPortalGid : undefined;
        } else {
            const args = argsOrState as PRAPortalArgs | undefined;
            resourceInputs["certificateId"] = args ? args.certificateId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["extDomain"] = args ? args.extDomain : undefined;
            resourceInputs["extDomainName"] = args ? args.extDomainName : undefined;
            resourceInputs["extDomainTranslation"] = args ? args.extDomainTranslation : undefined;
            resourceInputs["extLabel"] = args ? args.extLabel : undefined;
            resourceInputs["microtenantId"] = args ? args.microtenantId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["userNotification"] = args ? args.userNotification : undefined;
            resourceInputs["userNotificationEnabled"] = args ? args.userNotificationEnabled : undefined;
            resourceInputs["userPortalGid"] = args ? args.userPortalGid : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "zpa:index/praPortalController:PraPortalController" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(PRAPortal.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PRAPortal resources.
 */
export interface PRAPortalState {
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
     * The external domain name prefix of the Browser Access application that is used for Zscaler-managed certificates when
     * creating a privileged portal.
     */
    extDomain?: pulumi.Input<string>;
    /**
     * The domain suffix for the privileged portal URL. This field must be one of the customer's authentication domains.
     */
    extDomainName?: pulumi.Input<string>;
    /**
     * The translation of the external domain name prefix of the Browser Access application that is used for Zscaler-managed
     * certificates when creating a privileged portal.
     */
    extDomainTranslation?: pulumi.Input<string>;
    /**
     * The domain prefix for the privileged portal URL. The supported string can include numbers, lower case characters, and
     * only supports a hyphen (-).
     */
    extLabel?: pulumi.Input<string>;
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
    /**
     * The unique identifier of the user portal.
     */
    userPortalGid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PRAPortal resource.
 */
export interface PRAPortalArgs {
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
     * The external domain name prefix of the Browser Access application that is used for Zscaler-managed certificates when
     * creating a privileged portal.
     */
    extDomain?: pulumi.Input<string>;
    /**
     * The domain suffix for the privileged portal URL. This field must be one of the customer's authentication domains.
     */
    extDomainName?: pulumi.Input<string>;
    /**
     * The translation of the external domain name prefix of the Browser Access application that is used for Zscaler-managed
     * certificates when creating a privileged portal.
     */
    extDomainTranslation?: pulumi.Input<string>;
    /**
     * The domain prefix for the privileged portal URL. The supported string can include numbers, lower case characters, and
     * only supports a hyphen (-).
     */
    extLabel?: pulumi.Input<string>;
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
    /**
     * The unique identifier of the user portal.
     */
    userPortalGid?: pulumi.Input<string>;
}
