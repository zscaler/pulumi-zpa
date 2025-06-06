// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa
{
    /// <summary>
    /// * [Official documentation](https://help.zscaler.com/zpa/about-timeout-policy)
    /// * [API documentation](https://help.zscaler.com/zpa/configuring-timeout-policies-using-api)
    /// 
    /// The **zpa_policy_timeout_rule_v2** resource creates and manages policy access timeout rule in the Zscaler Private Access cloud using a new v2 API endpoint.
    /// 
    ///   ⚠️ **NOTE**: This resource is recommended if your configuration requires the association of more than 1000 resource criteria per rule.
    /// 
    ///   ⚠️ **WARNING:**: The attribute ``rule_order`` is now deprecated in favor of the new resource  ``policy_access_rule_reorder``
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Zpa = Pulumi.Zpa;
    /// using Zpa = zscaler.PulumiPackage.Zpa;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var thisIdPController = Zpa.GetIdPController.Invoke(new()
    ///     {
    ///         Name = "Idp_Name",
    ///     });
    /// 
    ///     var emailUserSso = Zpa.GetSAMLAttribute.Invoke(new()
    ///     {
    ///         Name = "Email_Users",
    ///         IdpName = "Idp_Name",
    ///     });
    /// 
    ///     var groupUser = Zpa.GetSAMLAttribute.Invoke(new()
    ///     {
    ///         Name = "GroupName_Users",
    ///         IdpName = "Idp_Name",
    ///     });
    /// 
    ///     var a000 = Zpa.GetSCIMGroups.Invoke(new()
    ///     {
    ///         Name = "A000",
    ///         IdpName = "Idp_Name",
    ///     });
    /// 
    ///     var b000 = Zpa.GetSCIMGroups.Invoke(new()
    ///     {
    ///         Name = "B000",
    ///         IdpName = "Idp_Name",
    ///     });
    /// 
    ///     // Create Segment Group
    ///     var thisSegmentGroup = new Zpa.SegmentGroup("thisSegmentGroup", new()
    ///     {
    ///         Description = "Example",
    ///         Enabled = true,
    ///     });
    /// 
    ///     // Create Policy Access Rule V2
    ///     var thisPolicyAccessTimeOutRuleV2 = new Zpa.PolicyAccessTimeOutRuleV2("thisPolicyAccessTimeOutRuleV2", new()
    ///     {
    ///         Description = "Example",
    ///         Action = "RE_AUTH",
    ///         ReauthIdleTimeout = "10 Days",
    ///         ReauthTimeout = "10 Days",
    ///         Conditions = new[]
    ///         {
    ///             new Zpa.Inputs.PolicyAccessTimeOutRuleV2ConditionArgs
    ///             {
    ///                 Operator = "OR",
    ///                 Operands = new[]
    ///                 {
    ///                     new Zpa.Inputs.PolicyAccessTimeOutRuleV2ConditionOperandArgs
    ///                     {
    ///                         ObjectType = "APP_GROUP",
    ///                         Values = new[]
    ///                         {
    ///                             thisSegmentGroup.Id,
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///             new Zpa.Inputs.PolicyAccessTimeOutRuleV2ConditionArgs
    ///             {
    ///                 Operator = "OR",
    ///                 Operands = new[]
    ///                 {
    ///                     new Zpa.Inputs.PolicyAccessTimeOutRuleV2ConditionOperandArgs
    ///                     {
    ///                         ObjectType = "SAML",
    ///                         EntryValues = new[]
    ///                         {
    ///                             new Zpa.Inputs.PolicyAccessTimeOutRuleV2ConditionOperandEntryValueArgs
    ///                             {
    ///                                 Rhs = "user1@acme.com",
    ///                                 Lhs = emailUserSso.Apply(getSAMLAttributeResult =&gt; getSAMLAttributeResult.Id),
    ///                             },
    ///                             new Zpa.Inputs.PolicyAccessTimeOutRuleV2ConditionOperandEntryValueArgs
    ///                             {
    ///                                 Rhs = "A000",
    ///                                 Lhs = groupUser.Apply(getSAMLAttributeResult =&gt; getSAMLAttributeResult.Id),
    ///                             },
    ///                         },
    ///                     },
    ///                     new Zpa.Inputs.PolicyAccessTimeOutRuleV2ConditionOperandArgs
    ///                     {
    ///                         ObjectType = "SCIM_GROUP",
    ///                         EntryValues = new[]
    ///                         {
    ///                             new Zpa.Inputs.PolicyAccessTimeOutRuleV2ConditionOperandEntryValueArgs
    ///                             {
    ///                                 Rhs = a000.Apply(getSCIMGroupsResult =&gt; getSCIMGroupsResult.Id),
    ///                                 Lhs = thisIdPController.Apply(getIdPControllerResult =&gt; getIdPControllerResult.Id),
    ///                             },
    ///                             new Zpa.Inputs.PolicyAccessTimeOutRuleV2ConditionOperandEntryValueArgs
    ///                             {
    ///                                 Rhs = b000.Apply(getSCIMGroupsResult =&gt; getSCIMGroupsResult.Id),
    ///                                 Lhs = thisIdPController.Apply(getIdPControllerResult =&gt; getIdPControllerResult.Id),
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///             new Zpa.Inputs.PolicyAccessTimeOutRuleV2ConditionArgs
    ///             {
    ///                 Operator = "OR",
    ///                 Operands = new[]
    ///                 {
    ///                     new Zpa.Inputs.PolicyAccessTimeOutRuleV2ConditionOperandArgs
    ///                     {
    ///                         ObjectType = "PLATFORM",
    ///                         EntryValues = new[]
    ///                         {
    ///                             new Zpa.Inputs.PolicyAccessTimeOutRuleV2ConditionOperandEntryValueArgs
    ///                             {
    ///                                 Rhs = "true",
    ///                                 Lhs = "linux",
    ///                             },
    ///                             new Zpa.Inputs.PolicyAccessTimeOutRuleV2ConditionOperandEntryValueArgs
    ///                             {
    ///                                 Rhs = "true",
    ///                                 Lhs = "android",
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## LHS and RHS Values
    /// 
    /// | Object Type | LHS| RHS| VALUES
    /// |----------|-----------|----------|----------
    /// | APP  |   |  | ``application_segment_id`` |
    /// | APP_GROUP  |   |  | ``segment_group_id``|
    /// | CLIENT_TYPE  |   |  |  ``zpn_client_type_zappl``, ``zpn_client_type_exporter``, ``zpn_client_type_browser_isolation``, ``zpn_client_type_ip_anchoring``, ``zpn_client_type_edge_connector``, ``zpn_client_type_branch_connector``,  ``zpn_client_type_zapp_partner``, ``zpn_client_type_zapp``  |
    /// | SAML | ``saml_attribute_id``  | ``attribute_value_to_match`` |
    /// | SCIM | ``scim_attribute_id``  | ``attribute_value_to_match``  |
    /// | SCIM_GROUP | ``scim_group_attribute_id``  | ``attribute_value_to_match``  |
    /// | PLATFORM | ``mac``, ``ios``, ``windows``, ``android``, ``linux`` | ``"true"`` / ``"false"`` |
    /// | POSTURE | ``posture_udid``  | ``"true"`` / ``"false"`` |
    /// 
    /// ## Import
    /// 
    /// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
    /// 
    /// Visit
    /// 
    /// Policy access timeout rule can be imported by using `&lt;RULE ID&gt;` as the import ID.
    /// 
    /// For example:
    /// 
    /// ```sh
    /// $ pulumi import zpa:index/policyAccessTimeOutRuleV2:PolicyAccessTimeOutRuleV2 example &lt;rule_id&gt;
    /// ```
    /// </summary>
    [ZpaResourceType("zpa:index/policyAccessTimeOutRuleV2:PolicyAccessTimeOutRuleV2")]
    public partial class PolicyAccessTimeOutRuleV2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// This is for providing the rule action.
        /// </summary>
        [Output("action")]
        public Output<string?> Action { get; private set; } = null!;

        /// <summary>
        /// This is for proviidng the set of conditions for the policy.
        /// </summary>
        [Output("conditions")]
        public Output<ImmutableArray<Outputs.PolicyAccessTimeOutRuleV2Condition>> Conditions { get; private set; } = null!;

        /// <summary>
        /// This is for providing a customer message for the user.
        /// </summary>
        [Output("customMsg")]
        public Output<string?> CustomMsg { get; private set; } = null!;

        /// <summary>
        /// This is the description of the access policy.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("microtenantId")]
        public Output<string> MicrotenantId { get; private set; } = null!;

        /// <summary>
        /// This is the name of the policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("policySetId")]
        public Output<string> PolicySetId { get; private set; } = null!;

        [Output("reauthIdleTimeout")]
        public Output<string?> ReauthIdleTimeout { get; private set; } = null!;

        [Output("reauthTimeout")]
        public Output<string?> ReauthTimeout { get; private set; } = null!;


        /// <summary>
        /// Create a PolicyAccessTimeOutRuleV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyAccessTimeOutRuleV2(string name, PolicyAccessTimeOutRuleV2Args? args = null, CustomResourceOptions? options = null)
            : base("zpa:index/policyAccessTimeOutRuleV2:PolicyAccessTimeOutRuleV2", name, args ?? new PolicyAccessTimeOutRuleV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private PolicyAccessTimeOutRuleV2(string name, Input<string> id, PolicyAccessTimeOutRuleV2State? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/policyAccessTimeOutRuleV2:PolicyAccessTimeOutRuleV2", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/zscaler",
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "zpa:index/policyTimeoutRuleV2:PolicyTimeoutRuleV2" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PolicyAccessTimeOutRuleV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyAccessTimeOutRuleV2 Get(string name, Input<string> id, PolicyAccessTimeOutRuleV2State? state = null, CustomResourceOptions? options = null)
        {
            return new PolicyAccessTimeOutRuleV2(name, id, state, options);
        }
    }

    public sealed class PolicyAccessTimeOutRuleV2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is for providing the rule action.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("conditions")]
        private InputList<Inputs.PolicyAccessTimeOutRuleV2ConditionArgs>? _conditions;

        /// <summary>
        /// This is for proviidng the set of conditions for the policy.
        /// </summary>
        public InputList<Inputs.PolicyAccessTimeOutRuleV2ConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.PolicyAccessTimeOutRuleV2ConditionArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// This is for providing a customer message for the user.
        /// </summary>
        [Input("customMsg")]
        public Input<string>? CustomMsg { get; set; }

        /// <summary>
        /// This is the description of the access policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// This is the name of the policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("reauthIdleTimeout")]
        public Input<string>? ReauthIdleTimeout { get; set; }

        [Input("reauthTimeout")]
        public Input<string>? ReauthTimeout { get; set; }

        public PolicyAccessTimeOutRuleV2Args()
        {
        }
        public static new PolicyAccessTimeOutRuleV2Args Empty => new PolicyAccessTimeOutRuleV2Args();
    }

    public sealed class PolicyAccessTimeOutRuleV2State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is for providing the rule action.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("conditions")]
        private InputList<Inputs.PolicyAccessTimeOutRuleV2ConditionGetArgs>? _conditions;

        /// <summary>
        /// This is for proviidng the set of conditions for the policy.
        /// </summary>
        public InputList<Inputs.PolicyAccessTimeOutRuleV2ConditionGetArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.PolicyAccessTimeOutRuleV2ConditionGetArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// This is for providing a customer message for the user.
        /// </summary>
        [Input("customMsg")]
        public Input<string>? CustomMsg { get; set; }

        /// <summary>
        /// This is the description of the access policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("microtenantId")]
        public Input<string>? MicrotenantId { get; set; }

        /// <summary>
        /// This is the name of the policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policySetId")]
        public Input<string>? PolicySetId { get; set; }

        [Input("reauthIdleTimeout")]
        public Input<string>? ReauthIdleTimeout { get; set; }

        [Input("reauthTimeout")]
        public Input<string>? ReauthTimeout { get; set; }

        public PolicyAccessTimeOutRuleV2State()
        {
        }
        public static new PolicyAccessTimeOutRuleV2State Empty => new PolicyAccessTimeOutRuleV2State();
    }
}
