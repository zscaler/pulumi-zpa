// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package segmentgroup

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SegmentGroupApplication struct {
	Id *string `pulumi:"id"`
}

// SegmentGroupApplicationInput is an input type that accepts SegmentGroupApplicationArgs and SegmentGroupApplicationOutput values.
// You can construct a concrete instance of `SegmentGroupApplicationInput` via:
//
//	SegmentGroupApplicationArgs{...}
type SegmentGroupApplicationInput interface {
	pulumi.Input

	ToSegmentGroupApplicationOutput() SegmentGroupApplicationOutput
	ToSegmentGroupApplicationOutputWithContext(context.Context) SegmentGroupApplicationOutput
}

type SegmentGroupApplicationArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SegmentGroupApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SegmentGroupApplication)(nil)).Elem()
}

func (i SegmentGroupApplicationArgs) ToSegmentGroupApplicationOutput() SegmentGroupApplicationOutput {
	return i.ToSegmentGroupApplicationOutputWithContext(context.Background())
}

func (i SegmentGroupApplicationArgs) ToSegmentGroupApplicationOutputWithContext(ctx context.Context) SegmentGroupApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SegmentGroupApplicationOutput)
}

// SegmentGroupApplicationArrayInput is an input type that accepts SegmentGroupApplicationArray and SegmentGroupApplicationArrayOutput values.
// You can construct a concrete instance of `SegmentGroupApplicationArrayInput` via:
//
//	SegmentGroupApplicationArray{ SegmentGroupApplicationArgs{...} }
type SegmentGroupApplicationArrayInput interface {
	pulumi.Input

	ToSegmentGroupApplicationArrayOutput() SegmentGroupApplicationArrayOutput
	ToSegmentGroupApplicationArrayOutputWithContext(context.Context) SegmentGroupApplicationArrayOutput
}

type SegmentGroupApplicationArray []SegmentGroupApplicationInput

func (SegmentGroupApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SegmentGroupApplication)(nil)).Elem()
}

func (i SegmentGroupApplicationArray) ToSegmentGroupApplicationArrayOutput() SegmentGroupApplicationArrayOutput {
	return i.ToSegmentGroupApplicationArrayOutputWithContext(context.Background())
}

func (i SegmentGroupApplicationArray) ToSegmentGroupApplicationArrayOutputWithContext(ctx context.Context) SegmentGroupApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SegmentGroupApplicationArrayOutput)
}

type SegmentGroupApplicationOutput struct{ *pulumi.OutputState }

func (SegmentGroupApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SegmentGroupApplication)(nil)).Elem()
}

func (o SegmentGroupApplicationOutput) ToSegmentGroupApplicationOutput() SegmentGroupApplicationOutput {
	return o
}

func (o SegmentGroupApplicationOutput) ToSegmentGroupApplicationOutputWithContext(ctx context.Context) SegmentGroupApplicationOutput {
	return o
}

func (o SegmentGroupApplicationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SegmentGroupApplication) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SegmentGroupApplicationArrayOutput struct{ *pulumi.OutputState }

func (SegmentGroupApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SegmentGroupApplication)(nil)).Elem()
}

func (o SegmentGroupApplicationArrayOutput) ToSegmentGroupApplicationArrayOutput() SegmentGroupApplicationArrayOutput {
	return o
}

func (o SegmentGroupApplicationArrayOutput) ToSegmentGroupApplicationArrayOutputWithContext(ctx context.Context) SegmentGroupApplicationArrayOutput {
	return o
}

func (o SegmentGroupApplicationArrayOutput) Index(i pulumi.IntInput) SegmentGroupApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SegmentGroupApplication {
		return vs[0].([]SegmentGroupApplication)[vs[1].(int)]
	}).(SegmentGroupApplicationOutput)
}

type GetSegmentGroupApplication struct {
	// (string)
	BypassType string `pulumi:"bypassType"`
	// (string)
	ConfigSpace string `pulumi:"configSpace"`
	// (string)
	CreationTime string `pulumi:"creationTime"`
	// (string)
	DefaultIdleTimeout string `pulumi:"defaultIdleTimeout"`
	// (string)
	DefaultMaxAge string `pulumi:"defaultMaxAge"`
	// (string)
	Description string `pulumi:"description"`
	// (string)
	DomainName string `pulumi:"domainName"`
	// (string)
	DomainNames []string `pulumi:"domainNames"`
	// (string)
	DoubleEncrypt bool `pulumi:"doubleEncrypt"`
	// (bool)
	Enabled bool `pulumi:"enabled"`
	// (string)
	HealthCheckType string `pulumi:"healthCheckType"`
	// The ID of the segment group to be exported.
	Id string `pulumi:"id"`
	// (bool)
	IpAnchored  bool     `pulumi:"ipAnchored"`
	LogFeatures []string `pulumi:"logFeatures"`
	// (string)
	ModifiedBy string `pulumi:"modifiedBy"`
	// (string)
	ModifiedTime string `pulumi:"modifiedTime"`
	// The name of the segment group to be exported.
	Name string `pulumi:"name"`
	// (bool)
	PassiveHealthEnabled bool `pulumi:"passiveHealthEnabled"`
	// (Computed)
	ServerGroups []GetSegmentGroupApplicationServerGroup `pulumi:"serverGroups"`
	// (string)
	TcpPortRanges []string `pulumi:"tcpPortRanges"`
	// (string)
	TcpPortsIns  []string `pulumi:"tcpPortsIns"`
	TcpPortsOuts []string `pulumi:"tcpPortsOuts"`
	// (string)
	UdpPortRanges []string `pulumi:"udpPortRanges"`
}

// GetSegmentGroupApplicationInput is an input type that accepts GetSegmentGroupApplicationArgs and GetSegmentGroupApplicationOutput values.
// You can construct a concrete instance of `GetSegmentGroupApplicationInput` via:
//
//	GetSegmentGroupApplicationArgs{...}
type GetSegmentGroupApplicationInput interface {
	pulumi.Input

	ToGetSegmentGroupApplicationOutput() GetSegmentGroupApplicationOutput
	ToGetSegmentGroupApplicationOutputWithContext(context.Context) GetSegmentGroupApplicationOutput
}

type GetSegmentGroupApplicationArgs struct {
	// (string)
	BypassType pulumi.StringInput `pulumi:"bypassType"`
	// (string)
	ConfigSpace pulumi.StringInput `pulumi:"configSpace"`
	// (string)
	CreationTime pulumi.StringInput `pulumi:"creationTime"`
	// (string)
	DefaultIdleTimeout pulumi.StringInput `pulumi:"defaultIdleTimeout"`
	// (string)
	DefaultMaxAge pulumi.StringInput `pulumi:"defaultMaxAge"`
	// (string)
	Description pulumi.StringInput `pulumi:"description"`
	// (string)
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// (string)
	DomainNames pulumi.StringArrayInput `pulumi:"domainNames"`
	// (string)
	DoubleEncrypt pulumi.BoolInput `pulumi:"doubleEncrypt"`
	// (bool)
	Enabled pulumi.BoolInput `pulumi:"enabled"`
	// (string)
	HealthCheckType pulumi.StringInput `pulumi:"healthCheckType"`
	// The ID of the segment group to be exported.
	Id pulumi.StringInput `pulumi:"id"`
	// (bool)
	IpAnchored  pulumi.BoolInput        `pulumi:"ipAnchored"`
	LogFeatures pulumi.StringArrayInput `pulumi:"logFeatures"`
	// (string)
	ModifiedBy pulumi.StringInput `pulumi:"modifiedBy"`
	// (string)
	ModifiedTime pulumi.StringInput `pulumi:"modifiedTime"`
	// The name of the segment group to be exported.
	Name pulumi.StringInput `pulumi:"name"`
	// (bool)
	PassiveHealthEnabled pulumi.BoolInput `pulumi:"passiveHealthEnabled"`
	// (Computed)
	ServerGroups GetSegmentGroupApplicationServerGroupArrayInput `pulumi:"serverGroups"`
	// (string)
	TcpPortRanges pulumi.StringArrayInput `pulumi:"tcpPortRanges"`
	// (string)
	TcpPortsIns  pulumi.StringArrayInput `pulumi:"tcpPortsIns"`
	TcpPortsOuts pulumi.StringArrayInput `pulumi:"tcpPortsOuts"`
	// (string)
	UdpPortRanges pulumi.StringArrayInput `pulumi:"udpPortRanges"`
}

func (GetSegmentGroupApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSegmentGroupApplication)(nil)).Elem()
}

func (i GetSegmentGroupApplicationArgs) ToGetSegmentGroupApplicationOutput() GetSegmentGroupApplicationOutput {
	return i.ToGetSegmentGroupApplicationOutputWithContext(context.Background())
}

func (i GetSegmentGroupApplicationArgs) ToGetSegmentGroupApplicationOutputWithContext(ctx context.Context) GetSegmentGroupApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSegmentGroupApplicationOutput)
}

// GetSegmentGroupApplicationArrayInput is an input type that accepts GetSegmentGroupApplicationArray and GetSegmentGroupApplicationArrayOutput values.
// You can construct a concrete instance of `GetSegmentGroupApplicationArrayInput` via:
//
//	GetSegmentGroupApplicationArray{ GetSegmentGroupApplicationArgs{...} }
type GetSegmentGroupApplicationArrayInput interface {
	pulumi.Input

	ToGetSegmentGroupApplicationArrayOutput() GetSegmentGroupApplicationArrayOutput
	ToGetSegmentGroupApplicationArrayOutputWithContext(context.Context) GetSegmentGroupApplicationArrayOutput
}

type GetSegmentGroupApplicationArray []GetSegmentGroupApplicationInput

func (GetSegmentGroupApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSegmentGroupApplication)(nil)).Elem()
}

func (i GetSegmentGroupApplicationArray) ToGetSegmentGroupApplicationArrayOutput() GetSegmentGroupApplicationArrayOutput {
	return i.ToGetSegmentGroupApplicationArrayOutputWithContext(context.Background())
}

func (i GetSegmentGroupApplicationArray) ToGetSegmentGroupApplicationArrayOutputWithContext(ctx context.Context) GetSegmentGroupApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSegmentGroupApplicationArrayOutput)
}

type GetSegmentGroupApplicationOutput struct{ *pulumi.OutputState }

func (GetSegmentGroupApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSegmentGroupApplication)(nil)).Elem()
}

func (o GetSegmentGroupApplicationOutput) ToGetSegmentGroupApplicationOutput() GetSegmentGroupApplicationOutput {
	return o
}

func (o GetSegmentGroupApplicationOutput) ToGetSegmentGroupApplicationOutputWithContext(ctx context.Context) GetSegmentGroupApplicationOutput {
	return o
}

// (string)
func (o GetSegmentGroupApplicationOutput) BypassType() pulumi.StringOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) string { return v.BypassType }).(pulumi.StringOutput)
}

// (string)
func (o GetSegmentGroupApplicationOutput) ConfigSpace() pulumi.StringOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) string { return v.ConfigSpace }).(pulumi.StringOutput)
}

// (string)
func (o GetSegmentGroupApplicationOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) string { return v.CreationTime }).(pulumi.StringOutput)
}

// (string)
func (o GetSegmentGroupApplicationOutput) DefaultIdleTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) string { return v.DefaultIdleTimeout }).(pulumi.StringOutput)
}

// (string)
func (o GetSegmentGroupApplicationOutput) DefaultMaxAge() pulumi.StringOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) string { return v.DefaultMaxAge }).(pulumi.StringOutput)
}

// (string)
func (o GetSegmentGroupApplicationOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) string { return v.Description }).(pulumi.StringOutput)
}

// (string)
func (o GetSegmentGroupApplicationOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) string { return v.DomainName }).(pulumi.StringOutput)
}

// (string)
func (o GetSegmentGroupApplicationOutput) DomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) []string { return v.DomainNames }).(pulumi.StringArrayOutput)
}

// (string)
func (o GetSegmentGroupApplicationOutput) DoubleEncrypt() pulumi.BoolOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) bool { return v.DoubleEncrypt }).(pulumi.BoolOutput)
}

// (bool)
func (o GetSegmentGroupApplicationOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// (string)
func (o GetSegmentGroupApplicationOutput) HealthCheckType() pulumi.StringOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) string { return v.HealthCheckType }).(pulumi.StringOutput)
}

// The ID of the segment group to be exported.
func (o GetSegmentGroupApplicationOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) string { return v.Id }).(pulumi.StringOutput)
}

// (bool)
func (o GetSegmentGroupApplicationOutput) IpAnchored() pulumi.BoolOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) bool { return v.IpAnchored }).(pulumi.BoolOutput)
}

func (o GetSegmentGroupApplicationOutput) LogFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) []string { return v.LogFeatures }).(pulumi.StringArrayOutput)
}

// (string)
func (o GetSegmentGroupApplicationOutput) ModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) string { return v.ModifiedBy }).(pulumi.StringOutput)
}

// (string)
func (o GetSegmentGroupApplicationOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

// The name of the segment group to be exported.
func (o GetSegmentGroupApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) string { return v.Name }).(pulumi.StringOutput)
}

// (bool)
func (o GetSegmentGroupApplicationOutput) PassiveHealthEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) bool { return v.PassiveHealthEnabled }).(pulumi.BoolOutput)
}

// (Computed)
func (o GetSegmentGroupApplicationOutput) ServerGroups() GetSegmentGroupApplicationServerGroupArrayOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) []GetSegmentGroupApplicationServerGroup { return v.ServerGroups }).(GetSegmentGroupApplicationServerGroupArrayOutput)
}

// (string)
func (o GetSegmentGroupApplicationOutput) TcpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) []string { return v.TcpPortRanges }).(pulumi.StringArrayOutput)
}

// (string)
func (o GetSegmentGroupApplicationOutput) TcpPortsIns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) []string { return v.TcpPortsIns }).(pulumi.StringArrayOutput)
}

func (o GetSegmentGroupApplicationOutput) TcpPortsOuts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) []string { return v.TcpPortsOuts }).(pulumi.StringArrayOutput)
}

// (string)
func (o GetSegmentGroupApplicationOutput) UdpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSegmentGroupApplication) []string { return v.UdpPortRanges }).(pulumi.StringArrayOutput)
}

type GetSegmentGroupApplicationArrayOutput struct{ *pulumi.OutputState }

func (GetSegmentGroupApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSegmentGroupApplication)(nil)).Elem()
}

func (o GetSegmentGroupApplicationArrayOutput) ToGetSegmentGroupApplicationArrayOutput() GetSegmentGroupApplicationArrayOutput {
	return o
}

func (o GetSegmentGroupApplicationArrayOutput) ToGetSegmentGroupApplicationArrayOutputWithContext(ctx context.Context) GetSegmentGroupApplicationArrayOutput {
	return o
}

func (o GetSegmentGroupApplicationArrayOutput) Index(i pulumi.IntInput) GetSegmentGroupApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetSegmentGroupApplication {
		return vs[0].([]GetSegmentGroupApplication)[vs[1].(int)]
	}).(GetSegmentGroupApplicationOutput)
}

type GetSegmentGroupApplicationServerGroup struct {
	// (string)
	ConfigSpace string `pulumi:"configSpace"`
	// (string)
	CreationTime string `pulumi:"creationTime"`
	// (string)
	Description string `pulumi:"description"`
	// (bool)
	DynamicDiscovery bool `pulumi:"dynamicDiscovery"`
	// (bool)
	Enabled bool `pulumi:"enabled"`
	// The ID of the segment group to be exported.
	Id string `pulumi:"id"`
	// (string)
	ModifiedBy string `pulumi:"modifiedBy"`
	// (string)
	ModifiedTime string `pulumi:"modifiedTime"`
	// The name of the segment group to be exported.
	Name string `pulumi:"name"`
}

// GetSegmentGroupApplicationServerGroupInput is an input type that accepts GetSegmentGroupApplicationServerGroupArgs and GetSegmentGroupApplicationServerGroupOutput values.
// You can construct a concrete instance of `GetSegmentGroupApplicationServerGroupInput` via:
//
//	GetSegmentGroupApplicationServerGroupArgs{...}
type GetSegmentGroupApplicationServerGroupInput interface {
	pulumi.Input

	ToGetSegmentGroupApplicationServerGroupOutput() GetSegmentGroupApplicationServerGroupOutput
	ToGetSegmentGroupApplicationServerGroupOutputWithContext(context.Context) GetSegmentGroupApplicationServerGroupOutput
}

type GetSegmentGroupApplicationServerGroupArgs struct {
	// (string)
	ConfigSpace pulumi.StringInput `pulumi:"configSpace"`
	// (string)
	CreationTime pulumi.StringInput `pulumi:"creationTime"`
	// (string)
	Description pulumi.StringInput `pulumi:"description"`
	// (bool)
	DynamicDiscovery pulumi.BoolInput `pulumi:"dynamicDiscovery"`
	// (bool)
	Enabled pulumi.BoolInput `pulumi:"enabled"`
	// The ID of the segment group to be exported.
	Id pulumi.StringInput `pulumi:"id"`
	// (string)
	ModifiedBy pulumi.StringInput `pulumi:"modifiedBy"`
	// (string)
	ModifiedTime pulumi.StringInput `pulumi:"modifiedTime"`
	// The name of the segment group to be exported.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetSegmentGroupApplicationServerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSegmentGroupApplicationServerGroup)(nil)).Elem()
}

func (i GetSegmentGroupApplicationServerGroupArgs) ToGetSegmentGroupApplicationServerGroupOutput() GetSegmentGroupApplicationServerGroupOutput {
	return i.ToGetSegmentGroupApplicationServerGroupOutputWithContext(context.Background())
}

func (i GetSegmentGroupApplicationServerGroupArgs) ToGetSegmentGroupApplicationServerGroupOutputWithContext(ctx context.Context) GetSegmentGroupApplicationServerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSegmentGroupApplicationServerGroupOutput)
}

// GetSegmentGroupApplicationServerGroupArrayInput is an input type that accepts GetSegmentGroupApplicationServerGroupArray and GetSegmentGroupApplicationServerGroupArrayOutput values.
// You can construct a concrete instance of `GetSegmentGroupApplicationServerGroupArrayInput` via:
//
//	GetSegmentGroupApplicationServerGroupArray{ GetSegmentGroupApplicationServerGroupArgs{...} }
type GetSegmentGroupApplicationServerGroupArrayInput interface {
	pulumi.Input

	ToGetSegmentGroupApplicationServerGroupArrayOutput() GetSegmentGroupApplicationServerGroupArrayOutput
	ToGetSegmentGroupApplicationServerGroupArrayOutputWithContext(context.Context) GetSegmentGroupApplicationServerGroupArrayOutput
}

type GetSegmentGroupApplicationServerGroupArray []GetSegmentGroupApplicationServerGroupInput

func (GetSegmentGroupApplicationServerGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSegmentGroupApplicationServerGroup)(nil)).Elem()
}

func (i GetSegmentGroupApplicationServerGroupArray) ToGetSegmentGroupApplicationServerGroupArrayOutput() GetSegmentGroupApplicationServerGroupArrayOutput {
	return i.ToGetSegmentGroupApplicationServerGroupArrayOutputWithContext(context.Background())
}

func (i GetSegmentGroupApplicationServerGroupArray) ToGetSegmentGroupApplicationServerGroupArrayOutputWithContext(ctx context.Context) GetSegmentGroupApplicationServerGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSegmentGroupApplicationServerGroupArrayOutput)
}

type GetSegmentGroupApplicationServerGroupOutput struct{ *pulumi.OutputState }

func (GetSegmentGroupApplicationServerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSegmentGroupApplicationServerGroup)(nil)).Elem()
}

func (o GetSegmentGroupApplicationServerGroupOutput) ToGetSegmentGroupApplicationServerGroupOutput() GetSegmentGroupApplicationServerGroupOutput {
	return o
}

func (o GetSegmentGroupApplicationServerGroupOutput) ToGetSegmentGroupApplicationServerGroupOutputWithContext(ctx context.Context) GetSegmentGroupApplicationServerGroupOutput {
	return o
}

// (string)
func (o GetSegmentGroupApplicationServerGroupOutput) ConfigSpace() pulumi.StringOutput {
	return o.ApplyT(func(v GetSegmentGroupApplicationServerGroup) string { return v.ConfigSpace }).(pulumi.StringOutput)
}

// (string)
func (o GetSegmentGroupApplicationServerGroupOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSegmentGroupApplicationServerGroup) string { return v.CreationTime }).(pulumi.StringOutput)
}

// (string)
func (o GetSegmentGroupApplicationServerGroupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetSegmentGroupApplicationServerGroup) string { return v.Description }).(pulumi.StringOutput)
}

// (bool)
func (o GetSegmentGroupApplicationServerGroupOutput) DynamicDiscovery() pulumi.BoolOutput {
	return o.ApplyT(func(v GetSegmentGroupApplicationServerGroup) bool { return v.DynamicDiscovery }).(pulumi.BoolOutput)
}

// (bool)
func (o GetSegmentGroupApplicationServerGroupOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetSegmentGroupApplicationServerGroup) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The ID of the segment group to be exported.
func (o GetSegmentGroupApplicationServerGroupOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSegmentGroupApplicationServerGroup) string { return v.Id }).(pulumi.StringOutput)
}

// (string)
func (o GetSegmentGroupApplicationServerGroupOutput) ModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v GetSegmentGroupApplicationServerGroup) string { return v.ModifiedBy }).(pulumi.StringOutput)
}

// (string)
func (o GetSegmentGroupApplicationServerGroupOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSegmentGroupApplicationServerGroup) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

// The name of the segment group to be exported.
func (o GetSegmentGroupApplicationServerGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSegmentGroupApplicationServerGroup) string { return v.Name }).(pulumi.StringOutput)
}

type GetSegmentGroupApplicationServerGroupArrayOutput struct{ *pulumi.OutputState }

func (GetSegmentGroupApplicationServerGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSegmentGroupApplicationServerGroup)(nil)).Elem()
}

func (o GetSegmentGroupApplicationServerGroupArrayOutput) ToGetSegmentGroupApplicationServerGroupArrayOutput() GetSegmentGroupApplicationServerGroupArrayOutput {
	return o
}

func (o GetSegmentGroupApplicationServerGroupArrayOutput) ToGetSegmentGroupApplicationServerGroupArrayOutputWithContext(ctx context.Context) GetSegmentGroupApplicationServerGroupArrayOutput {
	return o
}

func (o GetSegmentGroupApplicationServerGroupArrayOutput) Index(i pulumi.IntInput) GetSegmentGroupApplicationServerGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetSegmentGroupApplicationServerGroup {
		return vs[0].([]GetSegmentGroupApplicationServerGroup)[vs[1].(int)]
	}).(GetSegmentGroupApplicationServerGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SegmentGroupApplicationInput)(nil)).Elem(), SegmentGroupApplicationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SegmentGroupApplicationArrayInput)(nil)).Elem(), SegmentGroupApplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSegmentGroupApplicationInput)(nil)).Elem(), GetSegmentGroupApplicationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSegmentGroupApplicationArrayInput)(nil)).Elem(), GetSegmentGroupApplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSegmentGroupApplicationServerGroupInput)(nil)).Elem(), GetSegmentGroupApplicationServerGroupArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSegmentGroupApplicationServerGroupArrayInput)(nil)).Elem(), GetSegmentGroupApplicationServerGroupArray{})
	pulumi.RegisterOutputType(SegmentGroupApplicationOutput{})
	pulumi.RegisterOutputType(SegmentGroupApplicationArrayOutput{})
	pulumi.RegisterOutputType(GetSegmentGroupApplicationOutput{})
	pulumi.RegisterOutputType(GetSegmentGroupApplicationArrayOutput{})
	pulumi.RegisterOutputType(GetSegmentGroupApplicationServerGroupOutput{})
	pulumi.RegisterOutputType(GetSegmentGroupApplicationServerGroupArrayOutput{})
}
