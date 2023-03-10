// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { PolicyAccessRuleArgs, PolicyAccessRuleState } from "./policyAccessRule";
export type PolicyAccessRule = import("./policyAccessRule").PolicyAccessRule;
export const PolicyAccessRule: typeof import("./policyAccessRule").PolicyAccessRule = null as any;
utilities.lazyLoad(exports, ["PolicyAccessRule"], () => require("./policyAccessRule"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "zpa:AccessPolicy/policyAccessRule:PolicyAccessRule":
                return new PolicyAccessRule(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("zpa", "AccessPolicy/policyAccessRule", _module)
