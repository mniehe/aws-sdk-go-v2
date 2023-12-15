// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationautoscaling

import (
	"context"
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/applicationautoscaling/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpDeleteScalingPolicy struct {
}

func (*validateOpDeleteScalingPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteScalingPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteScalingPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteScalingPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteScheduledAction struct {
}

func (*validateOpDeleteScheduledAction) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteScheduledAction) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteScheduledActionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteScheduledActionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeregisterScalableTarget struct {
}

func (*validateOpDeregisterScalableTarget) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeregisterScalableTarget) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeregisterScalableTargetInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeregisterScalableTargetInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeScalableTargets struct {
}

func (*validateOpDescribeScalableTargets) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeScalableTargets) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeScalableTargetsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeScalableTargetsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeScalingActivities struct {
}

func (*validateOpDescribeScalingActivities) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeScalingActivities) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeScalingActivitiesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeScalingActivitiesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeScalingPolicies struct {
}

func (*validateOpDescribeScalingPolicies) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeScalingPolicies) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeScalingPoliciesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeScalingPoliciesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeScheduledActions struct {
}

func (*validateOpDescribeScheduledActions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeScheduledActions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeScheduledActionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeScheduledActionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutScalingPolicy struct {
}

func (*validateOpPutScalingPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutScalingPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutScalingPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutScalingPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutScheduledAction struct {
}

func (*validateOpPutScheduledAction) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutScheduledAction) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutScheduledActionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutScheduledActionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRegisterScalableTarget struct {
}

func (*validateOpRegisterScalableTarget) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRegisterScalableTarget) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RegisterScalableTargetInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRegisterScalableTargetInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpDeleteScalingPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteScalingPolicy{}, middleware.After)
}

func addOpDeleteScheduledActionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteScheduledAction{}, middleware.After)
}

func addOpDeregisterScalableTargetValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeregisterScalableTarget{}, middleware.After)
}

func addOpDescribeScalableTargetsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeScalableTargets{}, middleware.After)
}

func addOpDescribeScalingActivitiesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeScalingActivities{}, middleware.After)
}

func addOpDescribeScalingPoliciesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeScalingPolicies{}, middleware.After)
}

func addOpDescribeScheduledActionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeScheduledActions{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPutScalingPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutScalingPolicy{}, middleware.After)
}

func addOpPutScheduledActionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutScheduledAction{}, middleware.After)
}

func addOpRegisterScalableTargetValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRegisterScalableTarget{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func validateCustomizedMetricSpecification(v *types.CustomizedMetricSpecification) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CustomizedMetricSpecification"}
	if v.Dimensions != nil {
		if err := validateMetricDimensions(v.Dimensions); err != nil {
			invalidParams.AddNested("Dimensions", err.(smithy.InvalidParamsError))
		}
	}
	if v.Metrics != nil {
		if err := validateTargetTrackingMetricDataQueries(v.Metrics); err != nil {
			invalidParams.AddNested("Metrics", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMetricDimension(v *types.MetricDimension) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MetricDimension"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMetricDimensions(v []types.MetricDimension) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MetricDimensions"}
	for i := range v {
		if err := validateMetricDimension(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePredefinedMetricSpecification(v *types.PredefinedMetricSpecification) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PredefinedMetricSpecification"}
	if len(v.PredefinedMetricType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("PredefinedMetricType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateStepAdjustment(v *types.StepAdjustment) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StepAdjustment"}
	if v.ScalingAdjustment == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScalingAdjustment"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateStepAdjustments(v []types.StepAdjustment) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StepAdjustments"}
	for i := range v {
		if err := validateStepAdjustment(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateStepScalingPolicyConfiguration(v *types.StepScalingPolicyConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StepScalingPolicyConfiguration"}
	if v.StepAdjustments != nil {
		if err := validateStepAdjustments(v.StepAdjustments); err != nil {
			invalidParams.AddNested("StepAdjustments", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTargetTrackingMetric(v *types.TargetTrackingMetric) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TargetTrackingMetric"}
	if v.Dimensions != nil {
		if err := validateTargetTrackingMetricDimensions(v.Dimensions); err != nil {
			invalidParams.AddNested("Dimensions", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTargetTrackingMetricDataQueries(v []types.TargetTrackingMetricDataQuery) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TargetTrackingMetricDataQueries"}
	for i := range v {
		if err := validateTargetTrackingMetricDataQuery(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTargetTrackingMetricDataQuery(v *types.TargetTrackingMetricDataQuery) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TargetTrackingMetricDataQuery"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.MetricStat != nil {
		if err := validateTargetTrackingMetricStat(v.MetricStat); err != nil {
			invalidParams.AddNested("MetricStat", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTargetTrackingMetricDimension(v *types.TargetTrackingMetricDimension) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TargetTrackingMetricDimension"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTargetTrackingMetricDimensions(v []types.TargetTrackingMetricDimension) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TargetTrackingMetricDimensions"}
	for i := range v {
		if err := validateTargetTrackingMetricDimension(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTargetTrackingMetricStat(v *types.TargetTrackingMetricStat) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TargetTrackingMetricStat"}
	if v.Metric == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Metric"))
	} else if v.Metric != nil {
		if err := validateTargetTrackingMetric(v.Metric); err != nil {
			invalidParams.AddNested("Metric", err.(smithy.InvalidParamsError))
		}
	}
	if v.Stat == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Stat"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTargetTrackingScalingPolicyConfiguration(v *types.TargetTrackingScalingPolicyConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TargetTrackingScalingPolicyConfiguration"}
	if v.TargetValue == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetValue"))
	}
	if v.PredefinedMetricSpecification != nil {
		if err := validatePredefinedMetricSpecification(v.PredefinedMetricSpecification); err != nil {
			invalidParams.AddNested("PredefinedMetricSpecification", err.(smithy.InvalidParamsError))
		}
	}
	if v.CustomizedMetricSpecification != nil {
		if err := validateCustomizedMetricSpecification(v.CustomizedMetricSpecification); err != nil {
			invalidParams.AddNested("CustomizedMetricSpecification", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteScalingPolicyInput(v *DeleteScalingPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteScalingPolicyInput"}
	if v.PolicyName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyName"))
	}
	if len(v.ServiceNamespace) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceNamespace"))
	}
	if v.ResourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceId"))
	}
	if len(v.ScalableDimension) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ScalableDimension"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteScheduledActionInput(v *DeleteScheduledActionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteScheduledActionInput"}
	if len(v.ServiceNamespace) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceNamespace"))
	}
	if v.ScheduledActionName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScheduledActionName"))
	}
	if v.ResourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceId"))
	}
	if len(v.ScalableDimension) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ScalableDimension"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeregisterScalableTargetInput(v *DeregisterScalableTargetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeregisterScalableTargetInput"}
	if len(v.ServiceNamespace) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceNamespace"))
	}
	if v.ResourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceId"))
	}
	if len(v.ScalableDimension) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ScalableDimension"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeScalableTargetsInput(v *DescribeScalableTargetsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeScalableTargetsInput"}
	if len(v.ServiceNamespace) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceNamespace"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeScalingActivitiesInput(v *DescribeScalingActivitiesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeScalingActivitiesInput"}
	if len(v.ServiceNamespace) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceNamespace"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeScalingPoliciesInput(v *DescribeScalingPoliciesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeScalingPoliciesInput"}
	if len(v.ServiceNamespace) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceNamespace"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeScheduledActionsInput(v *DescribeScheduledActionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeScheduledActionsInput"}
	if len(v.ServiceNamespace) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceNamespace"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutScalingPolicyInput(v *PutScalingPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutScalingPolicyInput"}
	if v.PolicyName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyName"))
	}
	if len(v.ServiceNamespace) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceNamespace"))
	}
	if v.ResourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceId"))
	}
	if len(v.ScalableDimension) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ScalableDimension"))
	}
	if v.StepScalingPolicyConfiguration != nil {
		if err := validateStepScalingPolicyConfiguration(v.StepScalingPolicyConfiguration); err != nil {
			invalidParams.AddNested("StepScalingPolicyConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.TargetTrackingScalingPolicyConfiguration != nil {
		if err := validateTargetTrackingScalingPolicyConfiguration(v.TargetTrackingScalingPolicyConfiguration); err != nil {
			invalidParams.AddNested("TargetTrackingScalingPolicyConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutScheduledActionInput(v *PutScheduledActionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutScheduledActionInput"}
	if len(v.ServiceNamespace) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceNamespace"))
	}
	if v.ScheduledActionName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScheduledActionName"))
	}
	if v.ResourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceId"))
	}
	if len(v.ScalableDimension) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ScalableDimension"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRegisterScalableTargetInput(v *RegisterScalableTargetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RegisterScalableTargetInput"}
	if len(v.ServiceNamespace) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceNamespace"))
	}
	if v.ResourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceId"))
	}
	if len(v.ScalableDimension) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ScalableDimension"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
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

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
