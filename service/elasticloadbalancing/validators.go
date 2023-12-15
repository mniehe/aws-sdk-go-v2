// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/elasticloadbalancing/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAddTags struct {
}

func (*validateOpAddTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAddTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AddTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAddTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpApplySecurityGroupsToLoadBalancer struct {
}

func (*validateOpApplySecurityGroupsToLoadBalancer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpApplySecurityGroupsToLoadBalancer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ApplySecurityGroupsToLoadBalancerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpApplySecurityGroupsToLoadBalancerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAttachLoadBalancerToSubnets struct {
}

func (*validateOpAttachLoadBalancerToSubnets) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAttachLoadBalancerToSubnets) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AttachLoadBalancerToSubnetsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAttachLoadBalancerToSubnetsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpConfigureHealthCheck struct {
}

func (*validateOpConfigureHealthCheck) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpConfigureHealthCheck) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ConfigureHealthCheckInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpConfigureHealthCheckInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateAppCookieStickinessPolicy struct {
}

func (*validateOpCreateAppCookieStickinessPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateAppCookieStickinessPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateAppCookieStickinessPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateAppCookieStickinessPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateLBCookieStickinessPolicy struct {
}

func (*validateOpCreateLBCookieStickinessPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateLBCookieStickinessPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateLBCookieStickinessPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateLBCookieStickinessPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateLoadBalancer struct {
}

func (*validateOpCreateLoadBalancer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateLoadBalancer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateLoadBalancerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateLoadBalancerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateLoadBalancerListeners struct {
}

func (*validateOpCreateLoadBalancerListeners) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateLoadBalancerListeners) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateLoadBalancerListenersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateLoadBalancerListenersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateLoadBalancerPolicy struct {
}

func (*validateOpCreateLoadBalancerPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateLoadBalancerPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateLoadBalancerPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateLoadBalancerPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteLoadBalancer struct {
}

func (*validateOpDeleteLoadBalancer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteLoadBalancer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteLoadBalancerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteLoadBalancerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteLoadBalancerListeners struct {
}

func (*validateOpDeleteLoadBalancerListeners) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteLoadBalancerListeners) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteLoadBalancerListenersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteLoadBalancerListenersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteLoadBalancerPolicy struct {
}

func (*validateOpDeleteLoadBalancerPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteLoadBalancerPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteLoadBalancerPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteLoadBalancerPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeregisterInstancesFromLoadBalancer struct {
}

func (*validateOpDeregisterInstancesFromLoadBalancer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeregisterInstancesFromLoadBalancer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeregisterInstancesFromLoadBalancerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeregisterInstancesFromLoadBalancerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeInstanceHealth struct {
}

func (*validateOpDescribeInstanceHealth) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeInstanceHealth) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeInstanceHealthInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeInstanceHealthInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeLoadBalancerAttributes struct {
}

func (*validateOpDescribeLoadBalancerAttributes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeLoadBalancerAttributes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeLoadBalancerAttributesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeLoadBalancerAttributesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeTags struct {
}

func (*validateOpDescribeTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDetachLoadBalancerFromSubnets struct {
}

func (*validateOpDetachLoadBalancerFromSubnets) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDetachLoadBalancerFromSubnets) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DetachLoadBalancerFromSubnetsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDetachLoadBalancerFromSubnetsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisableAvailabilityZonesForLoadBalancer struct {
}

func (*validateOpDisableAvailabilityZonesForLoadBalancer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisableAvailabilityZonesForLoadBalancer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisableAvailabilityZonesForLoadBalancerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisableAvailabilityZonesForLoadBalancerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpEnableAvailabilityZonesForLoadBalancer struct {
}

func (*validateOpEnableAvailabilityZonesForLoadBalancer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpEnableAvailabilityZonesForLoadBalancer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*EnableAvailabilityZonesForLoadBalancerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpEnableAvailabilityZonesForLoadBalancerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpModifyLoadBalancerAttributes struct {
}

func (*validateOpModifyLoadBalancerAttributes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpModifyLoadBalancerAttributes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ModifyLoadBalancerAttributesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpModifyLoadBalancerAttributesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRegisterInstancesWithLoadBalancer struct {
}

func (*validateOpRegisterInstancesWithLoadBalancer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRegisterInstancesWithLoadBalancer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RegisterInstancesWithLoadBalancerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRegisterInstancesWithLoadBalancerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRemoveTags struct {
}

func (*validateOpRemoveTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRemoveTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RemoveTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRemoveTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSetLoadBalancerListenerSSLCertificate struct {
}

func (*validateOpSetLoadBalancerListenerSSLCertificate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSetLoadBalancerListenerSSLCertificate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SetLoadBalancerListenerSSLCertificateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSetLoadBalancerListenerSSLCertificateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSetLoadBalancerPoliciesForBackendServer struct {
}

func (*validateOpSetLoadBalancerPoliciesForBackendServer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSetLoadBalancerPoliciesForBackendServer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SetLoadBalancerPoliciesForBackendServerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSetLoadBalancerPoliciesForBackendServerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSetLoadBalancerPoliciesOfListener struct {
}

func (*validateOpSetLoadBalancerPoliciesOfListener) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSetLoadBalancerPoliciesOfListener) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SetLoadBalancerPoliciesOfListenerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSetLoadBalancerPoliciesOfListenerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAddTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAddTags{}, middleware.After)
}

func addOpApplySecurityGroupsToLoadBalancerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpApplySecurityGroupsToLoadBalancer{}, middleware.After)
}

func addOpAttachLoadBalancerToSubnetsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAttachLoadBalancerToSubnets{}, middleware.After)
}

func addOpConfigureHealthCheckValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpConfigureHealthCheck{}, middleware.After)
}

func addOpCreateAppCookieStickinessPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateAppCookieStickinessPolicy{}, middleware.After)
}

func addOpCreateLBCookieStickinessPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateLBCookieStickinessPolicy{}, middleware.After)
}

func addOpCreateLoadBalancerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateLoadBalancer{}, middleware.After)
}

func addOpCreateLoadBalancerListenersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateLoadBalancerListeners{}, middleware.After)
}

func addOpCreateLoadBalancerPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateLoadBalancerPolicy{}, middleware.After)
}

func addOpDeleteLoadBalancerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteLoadBalancer{}, middleware.After)
}

func addOpDeleteLoadBalancerListenersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteLoadBalancerListeners{}, middleware.After)
}

func addOpDeleteLoadBalancerPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteLoadBalancerPolicy{}, middleware.After)
}

func addOpDeregisterInstancesFromLoadBalancerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeregisterInstancesFromLoadBalancer{}, middleware.After)
}

func addOpDescribeInstanceHealthValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeInstanceHealth{}, middleware.After)
}

func addOpDescribeLoadBalancerAttributesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeLoadBalancerAttributes{}, middleware.After)
}

func addOpDescribeTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeTags{}, middleware.After)
}

func addOpDetachLoadBalancerFromSubnetsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDetachLoadBalancerFromSubnets{}, middleware.After)
}

func addOpDisableAvailabilityZonesForLoadBalancerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisableAvailabilityZonesForLoadBalancer{}, middleware.After)
}

func addOpEnableAvailabilityZonesForLoadBalancerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpEnableAvailabilityZonesForLoadBalancer{}, middleware.After)
}

func addOpModifyLoadBalancerAttributesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpModifyLoadBalancerAttributes{}, middleware.After)
}

func addOpRegisterInstancesWithLoadBalancerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRegisterInstancesWithLoadBalancer{}, middleware.After)
}

func addOpRemoveTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRemoveTags{}, middleware.After)
}

func addOpSetLoadBalancerListenerSSLCertificateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSetLoadBalancerListenerSSLCertificate{}, middleware.After)
}

func addOpSetLoadBalancerPoliciesForBackendServerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSetLoadBalancerPoliciesForBackendServer{}, middleware.After)
}

func addOpSetLoadBalancerPoliciesOfListenerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSetLoadBalancerPoliciesOfListener{}, middleware.After)
}

func validateAccessLog(v *types.AccessLog) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AccessLog"}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateConnectionDraining(v *types.ConnectionDraining) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ConnectionDraining"}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateConnectionSettings(v *types.ConnectionSettings) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ConnectionSettings"}
	if v.IdleTimeout == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdleTimeout"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCrossZoneLoadBalancing(v *types.CrossZoneLoadBalancing) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CrossZoneLoadBalancing"}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateHealthCheck(v *types.HealthCheck) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HealthCheck"}
	if v.Target == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Target"))
	}
	if v.Interval == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Interval"))
	}
	if v.Timeout == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timeout"))
	}
	if v.UnhealthyThreshold == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UnhealthyThreshold"))
	}
	if v.HealthyThreshold == nil {
		invalidParams.Add(smithy.NewErrParamRequired("HealthyThreshold"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateListener(v *types.Listener) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Listener"}
	if v.Protocol == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Protocol"))
	}
	if v.InstancePort == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InstancePort"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateListeners(v []types.Listener) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Listeners"}
	for i := range v {
		if err := validateListener(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateLoadBalancerAttributes(v *types.LoadBalancerAttributes) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LoadBalancerAttributes"}
	if v.CrossZoneLoadBalancing != nil {
		if err := validateCrossZoneLoadBalancing(v.CrossZoneLoadBalancing); err != nil {
			invalidParams.AddNested("CrossZoneLoadBalancing", err.(smithy.InvalidParamsError))
		}
	}
	if v.AccessLog != nil {
		if err := validateAccessLog(v.AccessLog); err != nil {
			invalidParams.AddNested("AccessLog", err.(smithy.InvalidParamsError))
		}
	}
	if v.ConnectionDraining != nil {
		if err := validateConnectionDraining(v.ConnectionDraining); err != nil {
			invalidParams.AddNested("ConnectionDraining", err.(smithy.InvalidParamsError))
		}
	}
	if v.ConnectionSettings != nil {
		if err := validateConnectionSettings(v.ConnectionSettings); err != nil {
			invalidParams.AddNested("ConnectionSettings", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAddTagsInput(v *AddTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AddTagsInput"}
	if v.LoadBalancerNames == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerNames"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpApplySecurityGroupsToLoadBalancerInput(v *ApplySecurityGroupsToLoadBalancerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ApplySecurityGroupsToLoadBalancerInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if v.SecurityGroups == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecurityGroups"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAttachLoadBalancerToSubnetsInput(v *AttachLoadBalancerToSubnetsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AttachLoadBalancerToSubnetsInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if v.Subnets == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Subnets"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpConfigureHealthCheckInput(v *ConfigureHealthCheckInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ConfigureHealthCheckInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if v.HealthCheck == nil {
		invalidParams.Add(smithy.NewErrParamRequired("HealthCheck"))
	} else if v.HealthCheck != nil {
		if err := validateHealthCheck(v.HealthCheck); err != nil {
			invalidParams.AddNested("HealthCheck", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateAppCookieStickinessPolicyInput(v *CreateAppCookieStickinessPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateAppCookieStickinessPolicyInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if v.PolicyName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyName"))
	}
	if v.CookieName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CookieName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateLBCookieStickinessPolicyInput(v *CreateLBCookieStickinessPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateLBCookieStickinessPolicyInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if v.PolicyName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateLoadBalancerInput(v *CreateLoadBalancerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateLoadBalancerInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if v.Listeners == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Listeners"))
	} else if v.Listeners != nil {
		if err := validateListeners(v.Listeners); err != nil {
			invalidParams.AddNested("Listeners", err.(smithy.InvalidParamsError))
		}
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateLoadBalancerListenersInput(v *CreateLoadBalancerListenersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateLoadBalancerListenersInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if v.Listeners == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Listeners"))
	} else if v.Listeners != nil {
		if err := validateListeners(v.Listeners); err != nil {
			invalidParams.AddNested("Listeners", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateLoadBalancerPolicyInput(v *CreateLoadBalancerPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateLoadBalancerPolicyInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if v.PolicyName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyName"))
	}
	if v.PolicyTypeName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyTypeName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteLoadBalancerInput(v *DeleteLoadBalancerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteLoadBalancerInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteLoadBalancerListenersInput(v *DeleteLoadBalancerListenersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteLoadBalancerListenersInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if v.LoadBalancerPorts == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerPorts"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteLoadBalancerPolicyInput(v *DeleteLoadBalancerPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteLoadBalancerPolicyInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if v.PolicyName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeregisterInstancesFromLoadBalancerInput(v *DeregisterInstancesFromLoadBalancerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeregisterInstancesFromLoadBalancerInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if v.Instances == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Instances"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeInstanceHealthInput(v *DescribeInstanceHealthInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeInstanceHealthInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeLoadBalancerAttributesInput(v *DescribeLoadBalancerAttributesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeLoadBalancerAttributesInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeTagsInput(v *DescribeTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeTagsInput"}
	if v.LoadBalancerNames == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerNames"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDetachLoadBalancerFromSubnetsInput(v *DetachLoadBalancerFromSubnetsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DetachLoadBalancerFromSubnetsInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if v.Subnets == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Subnets"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisableAvailabilityZonesForLoadBalancerInput(v *DisableAvailabilityZonesForLoadBalancerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisableAvailabilityZonesForLoadBalancerInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if v.AvailabilityZones == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AvailabilityZones"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpEnableAvailabilityZonesForLoadBalancerInput(v *EnableAvailabilityZonesForLoadBalancerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EnableAvailabilityZonesForLoadBalancerInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if v.AvailabilityZones == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AvailabilityZones"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpModifyLoadBalancerAttributesInput(v *ModifyLoadBalancerAttributesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ModifyLoadBalancerAttributesInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if v.LoadBalancerAttributes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerAttributes"))
	} else if v.LoadBalancerAttributes != nil {
		if err := validateLoadBalancerAttributes(v.LoadBalancerAttributes); err != nil {
			invalidParams.AddNested("LoadBalancerAttributes", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRegisterInstancesWithLoadBalancerInput(v *RegisterInstancesWithLoadBalancerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RegisterInstancesWithLoadBalancerInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if v.Instances == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Instances"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRemoveTagsInput(v *RemoveTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RemoveTagsInput"}
	if v.LoadBalancerNames == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerNames"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSetLoadBalancerListenerSSLCertificateInput(v *SetLoadBalancerListenerSSLCertificateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SetLoadBalancerListenerSSLCertificateInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if v.SSLCertificateId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SSLCertificateId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSetLoadBalancerPoliciesForBackendServerInput(v *SetLoadBalancerPoliciesForBackendServerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SetLoadBalancerPoliciesForBackendServerInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if v.InstancePort == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InstancePort"))
	}
	if v.PolicyNames == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyNames"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSetLoadBalancerPoliciesOfListenerInput(v *SetLoadBalancerPoliciesOfListenerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SetLoadBalancerPoliciesOfListenerInput"}
	if v.LoadBalancerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoadBalancerName"))
	}
	if v.PolicyNames == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyNames"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
