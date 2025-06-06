// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/internal"
)

// * [Official documentation](https://help.zscaler.com/zpa/configuring-app-connectors-settings)
// * [API documentation](https://help.zscaler.com/zpa/configuring-auto-delete-disconnected-app-connectors-using-api)
//
// Use the **zpa_service_edge_assistant_schedule** resource sets the scheduled frequency at which the disconnected Service Edges are eligible for deletion. The supported value for frequency is days. The frequencyInterval field is the number of days after an Service Edge disconnects for it to become eligible for deletion. The minimum supported value for frequencyInterval is 5.
//
// > **NOTE** - When enabling the Assistant Schedule for the first time, you must provide the `customerId` information. If you authenticated using environment variables and used `ZPA_CUSTOMER_ID` environment variable, you don't have to define the customerId attribute in the HCL configuration, and the provider will automatically use the value from the environment variable `ZPA_CUSTOMER_ID`
//
// ## Example Usage
//
// ### Defined Customer ID Value
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/zscaler/pulumi-zpa/sdk/go/zpa"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := zpa.NewServiceEdgeAssistantSchedule(ctx, "this", &zpa.ServiceEdgeAssistantScheduleArgs{
//				CustomerId:        pulumi.String("123456789101112"),
//				DeleteDisabled:    pulumi.Bool(true),
//				Enabled:           pulumi.Bool(true),
//				Frequency:         pulumi.String("days"),
//				FrequencyInterval: pulumi.String("5"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ### Customer ID Via Environment Variable
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/zscaler/pulumi-zpa/sdk/go/zpa"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := zpa.NewServiceEdgeAssistantSchedule(ctx, "this", &zpa.ServiceEdgeAssistantScheduleArgs{
//				DeleteDisabled:    pulumi.Bool(true),
//				Enabled:           pulumi.Bool(true),
//				Frequency:         pulumi.String("days"),
//				FrequencyInterval: pulumi.String("5"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Import is not currently supported for this resource.
type ServiceEdgeAssistantSchedule struct {
	pulumi.CustomResourceState

	CustomerId        pulumi.StringOutput    `pulumi:"customerId"`
	DeleteDisabled    pulumi.BoolPtrOutput   `pulumi:"deleteDisabled"`
	Enabled           pulumi.BoolPtrOutput   `pulumi:"enabled"`
	Frequency         pulumi.StringPtrOutput `pulumi:"frequency"`
	FrequencyInterval pulumi.StringPtrOutput `pulumi:"frequencyInterval"`
}

// NewServiceEdgeAssistantSchedule registers a new resource with the given unique name, arguments, and options.
func NewServiceEdgeAssistantSchedule(ctx *pulumi.Context,
	name string, args *ServiceEdgeAssistantScheduleArgs, opts ...pulumi.ResourceOption) (*ServiceEdgeAssistantSchedule, error) {
	if args == nil {
		args = &ServiceEdgeAssistantScheduleArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceEdgeAssistantSchedule
	err := ctx.RegisterResource("zpa:index/serviceEdgeAssistantSchedule:ServiceEdgeAssistantSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEdgeAssistantSchedule gets an existing ServiceEdgeAssistantSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEdgeAssistantSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEdgeAssistantScheduleState, opts ...pulumi.ResourceOption) (*ServiceEdgeAssistantSchedule, error) {
	var resource ServiceEdgeAssistantSchedule
	err := ctx.ReadResource("zpa:index/serviceEdgeAssistantSchedule:ServiceEdgeAssistantSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEdgeAssistantSchedule resources.
type serviceEdgeAssistantScheduleState struct {
	CustomerId        *string `pulumi:"customerId"`
	DeleteDisabled    *bool   `pulumi:"deleteDisabled"`
	Enabled           *bool   `pulumi:"enabled"`
	Frequency         *string `pulumi:"frequency"`
	FrequencyInterval *string `pulumi:"frequencyInterval"`
}

type ServiceEdgeAssistantScheduleState struct {
	CustomerId        pulumi.StringPtrInput
	DeleteDisabled    pulumi.BoolPtrInput
	Enabled           pulumi.BoolPtrInput
	Frequency         pulumi.StringPtrInput
	FrequencyInterval pulumi.StringPtrInput
}

func (ServiceEdgeAssistantScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEdgeAssistantScheduleState)(nil)).Elem()
}

type serviceEdgeAssistantScheduleArgs struct {
	CustomerId        *string `pulumi:"customerId"`
	DeleteDisabled    *bool   `pulumi:"deleteDisabled"`
	Enabled           *bool   `pulumi:"enabled"`
	Frequency         *string `pulumi:"frequency"`
	FrequencyInterval *string `pulumi:"frequencyInterval"`
}

// The set of arguments for constructing a ServiceEdgeAssistantSchedule resource.
type ServiceEdgeAssistantScheduleArgs struct {
	CustomerId        pulumi.StringPtrInput
	DeleteDisabled    pulumi.BoolPtrInput
	Enabled           pulumi.BoolPtrInput
	Frequency         pulumi.StringPtrInput
	FrequencyInterval pulumi.StringPtrInput
}

func (ServiceEdgeAssistantScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEdgeAssistantScheduleArgs)(nil)).Elem()
}

type ServiceEdgeAssistantScheduleInput interface {
	pulumi.Input

	ToServiceEdgeAssistantScheduleOutput() ServiceEdgeAssistantScheduleOutput
	ToServiceEdgeAssistantScheduleOutputWithContext(ctx context.Context) ServiceEdgeAssistantScheduleOutput
}

func (*ServiceEdgeAssistantSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEdgeAssistantSchedule)(nil)).Elem()
}

func (i *ServiceEdgeAssistantSchedule) ToServiceEdgeAssistantScheduleOutput() ServiceEdgeAssistantScheduleOutput {
	return i.ToServiceEdgeAssistantScheduleOutputWithContext(context.Background())
}

func (i *ServiceEdgeAssistantSchedule) ToServiceEdgeAssistantScheduleOutputWithContext(ctx context.Context) ServiceEdgeAssistantScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEdgeAssistantScheduleOutput)
}

// ServiceEdgeAssistantScheduleArrayInput is an input type that accepts ServiceEdgeAssistantScheduleArray and ServiceEdgeAssistantScheduleArrayOutput values.
// You can construct a concrete instance of `ServiceEdgeAssistantScheduleArrayInput` via:
//
//	ServiceEdgeAssistantScheduleArray{ ServiceEdgeAssistantScheduleArgs{...} }
type ServiceEdgeAssistantScheduleArrayInput interface {
	pulumi.Input

	ToServiceEdgeAssistantScheduleArrayOutput() ServiceEdgeAssistantScheduleArrayOutput
	ToServiceEdgeAssistantScheduleArrayOutputWithContext(context.Context) ServiceEdgeAssistantScheduleArrayOutput
}

type ServiceEdgeAssistantScheduleArray []ServiceEdgeAssistantScheduleInput

func (ServiceEdgeAssistantScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEdgeAssistantSchedule)(nil)).Elem()
}

func (i ServiceEdgeAssistantScheduleArray) ToServiceEdgeAssistantScheduleArrayOutput() ServiceEdgeAssistantScheduleArrayOutput {
	return i.ToServiceEdgeAssistantScheduleArrayOutputWithContext(context.Background())
}

func (i ServiceEdgeAssistantScheduleArray) ToServiceEdgeAssistantScheduleArrayOutputWithContext(ctx context.Context) ServiceEdgeAssistantScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEdgeAssistantScheduleArrayOutput)
}

// ServiceEdgeAssistantScheduleMapInput is an input type that accepts ServiceEdgeAssistantScheduleMap and ServiceEdgeAssistantScheduleMapOutput values.
// You can construct a concrete instance of `ServiceEdgeAssistantScheduleMapInput` via:
//
//	ServiceEdgeAssistantScheduleMap{ "key": ServiceEdgeAssistantScheduleArgs{...} }
type ServiceEdgeAssistantScheduleMapInput interface {
	pulumi.Input

	ToServiceEdgeAssistantScheduleMapOutput() ServiceEdgeAssistantScheduleMapOutput
	ToServiceEdgeAssistantScheduleMapOutputWithContext(context.Context) ServiceEdgeAssistantScheduleMapOutput
}

type ServiceEdgeAssistantScheduleMap map[string]ServiceEdgeAssistantScheduleInput

func (ServiceEdgeAssistantScheduleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEdgeAssistantSchedule)(nil)).Elem()
}

func (i ServiceEdgeAssistantScheduleMap) ToServiceEdgeAssistantScheduleMapOutput() ServiceEdgeAssistantScheduleMapOutput {
	return i.ToServiceEdgeAssistantScheduleMapOutputWithContext(context.Background())
}

func (i ServiceEdgeAssistantScheduleMap) ToServiceEdgeAssistantScheduleMapOutputWithContext(ctx context.Context) ServiceEdgeAssistantScheduleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEdgeAssistantScheduleMapOutput)
}

type ServiceEdgeAssistantScheduleOutput struct{ *pulumi.OutputState }

func (ServiceEdgeAssistantScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEdgeAssistantSchedule)(nil)).Elem()
}

func (o ServiceEdgeAssistantScheduleOutput) ToServiceEdgeAssistantScheduleOutput() ServiceEdgeAssistantScheduleOutput {
	return o
}

func (o ServiceEdgeAssistantScheduleOutput) ToServiceEdgeAssistantScheduleOutputWithContext(ctx context.Context) ServiceEdgeAssistantScheduleOutput {
	return o
}

func (o ServiceEdgeAssistantScheduleOutput) CustomerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeAssistantSchedule) pulumi.StringOutput { return v.CustomerId }).(pulumi.StringOutput)
}

func (o ServiceEdgeAssistantScheduleOutput) DeleteDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceEdgeAssistantSchedule) pulumi.BoolPtrOutput { return v.DeleteDisabled }).(pulumi.BoolPtrOutput)
}

func (o ServiceEdgeAssistantScheduleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceEdgeAssistantSchedule) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ServiceEdgeAssistantScheduleOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEdgeAssistantSchedule) pulumi.StringPtrOutput { return v.Frequency }).(pulumi.StringPtrOutput)
}

func (o ServiceEdgeAssistantScheduleOutput) FrequencyInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEdgeAssistantSchedule) pulumi.StringPtrOutput { return v.FrequencyInterval }).(pulumi.StringPtrOutput)
}

type ServiceEdgeAssistantScheduleArrayOutput struct{ *pulumi.OutputState }

func (ServiceEdgeAssistantScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEdgeAssistantSchedule)(nil)).Elem()
}

func (o ServiceEdgeAssistantScheduleArrayOutput) ToServiceEdgeAssistantScheduleArrayOutput() ServiceEdgeAssistantScheduleArrayOutput {
	return o
}

func (o ServiceEdgeAssistantScheduleArrayOutput) ToServiceEdgeAssistantScheduleArrayOutputWithContext(ctx context.Context) ServiceEdgeAssistantScheduleArrayOutput {
	return o
}

func (o ServiceEdgeAssistantScheduleArrayOutput) Index(i pulumi.IntInput) ServiceEdgeAssistantScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceEdgeAssistantSchedule {
		return vs[0].([]*ServiceEdgeAssistantSchedule)[vs[1].(int)]
	}).(ServiceEdgeAssistantScheduleOutput)
}

type ServiceEdgeAssistantScheduleMapOutput struct{ *pulumi.OutputState }

func (ServiceEdgeAssistantScheduleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEdgeAssistantSchedule)(nil)).Elem()
}

func (o ServiceEdgeAssistantScheduleMapOutput) ToServiceEdgeAssistantScheduleMapOutput() ServiceEdgeAssistantScheduleMapOutput {
	return o
}

func (o ServiceEdgeAssistantScheduleMapOutput) ToServiceEdgeAssistantScheduleMapOutputWithContext(ctx context.Context) ServiceEdgeAssistantScheduleMapOutput {
	return o
}

func (o ServiceEdgeAssistantScheduleMapOutput) MapIndex(k pulumi.StringInput) ServiceEdgeAssistantScheduleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceEdgeAssistantSchedule {
		return vs[0].(map[string]*ServiceEdgeAssistantSchedule)[vs[1].(string)]
	}).(ServiceEdgeAssistantScheduleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEdgeAssistantScheduleInput)(nil)).Elem(), &ServiceEdgeAssistantSchedule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEdgeAssistantScheduleArrayInput)(nil)).Elem(), ServiceEdgeAssistantScheduleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEdgeAssistantScheduleMapInput)(nil)).Elem(), ServiceEdgeAssistantScheduleMap{})
	pulumi.RegisterOutputType(ServiceEdgeAssistantScheduleOutput{})
	pulumi.RegisterOutputType(ServiceEdgeAssistantScheduleArrayOutput{})
	pulumi.RegisterOutputType(ServiceEdgeAssistantScheduleMapOutput{})
}
