// Code generated by smithy-go-codegen DO NOT EDIT.

package fis

import (
	"context"
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/fis/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateExperimentTemplate struct {
}

func (*validateOpCreateExperimentTemplate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateExperimentTemplate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateExperimentTemplateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateExperimentTemplateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateTargetAccountConfiguration struct {
}

func (*validateOpCreateTargetAccountConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateTargetAccountConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateTargetAccountConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateTargetAccountConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteExperimentTemplate struct {
}

func (*validateOpDeleteExperimentTemplate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteExperimentTemplate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteExperimentTemplateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteExperimentTemplateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteTargetAccountConfiguration struct {
}

func (*validateOpDeleteTargetAccountConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteTargetAccountConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteTargetAccountConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteTargetAccountConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetAction struct {
}

func (*validateOpGetAction) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetAction) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetActionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetActionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetExperiment struct {
}

func (*validateOpGetExperiment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetExperiment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetExperimentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetExperimentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetExperimentTargetAccountConfiguration struct {
}

func (*validateOpGetExperimentTargetAccountConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetExperimentTargetAccountConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetExperimentTargetAccountConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetExperimentTargetAccountConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetExperimentTemplate struct {
}

func (*validateOpGetExperimentTemplate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetExperimentTemplate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetExperimentTemplateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetExperimentTemplateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetTargetAccountConfiguration struct {
}

func (*validateOpGetTargetAccountConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTargetAccountConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTargetAccountConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTargetAccountConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetTargetResourceType struct {
}

func (*validateOpGetTargetResourceType) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTargetResourceType) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTargetResourceTypeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTargetResourceTypeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListExperimentResolvedTargets struct {
}

func (*validateOpListExperimentResolvedTargets) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListExperimentResolvedTargets) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListExperimentResolvedTargetsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListExperimentResolvedTargetsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListExperimentTargetAccountConfigurations struct {
}

func (*validateOpListExperimentTargetAccountConfigurations) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListExperimentTargetAccountConfigurations) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListExperimentTargetAccountConfigurationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListExperimentTargetAccountConfigurationsInput(input); err != nil {
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

type validateOpListTargetAccountConfigurations struct {
}

func (*validateOpListTargetAccountConfigurations) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTargetAccountConfigurations) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTargetAccountConfigurationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTargetAccountConfigurationsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartExperiment struct {
}

func (*validateOpStartExperiment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartExperiment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartExperimentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartExperimentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopExperiment struct {
}

func (*validateOpStopExperiment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopExperiment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopExperimentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopExperimentInput(input); err != nil {
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

type validateOpUpdateExperimentTemplate struct {
}

func (*validateOpUpdateExperimentTemplate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateExperimentTemplate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateExperimentTemplateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateExperimentTemplateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateTargetAccountConfiguration struct {
}

func (*validateOpUpdateTargetAccountConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateTargetAccountConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateTargetAccountConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateTargetAccountConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateExperimentTemplateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateExperimentTemplate{}, middleware.After)
}

func addOpCreateTargetAccountConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateTargetAccountConfiguration{}, middleware.After)
}

func addOpDeleteExperimentTemplateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteExperimentTemplate{}, middleware.After)
}

func addOpDeleteTargetAccountConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteTargetAccountConfiguration{}, middleware.After)
}

func addOpGetActionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetAction{}, middleware.After)
}

func addOpGetExperimentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetExperiment{}, middleware.After)
}

func addOpGetExperimentTargetAccountConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetExperimentTargetAccountConfiguration{}, middleware.After)
}

func addOpGetExperimentTemplateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetExperimentTemplate{}, middleware.After)
}

func addOpGetTargetAccountConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTargetAccountConfiguration{}, middleware.After)
}

func addOpGetTargetResourceTypeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTargetResourceType{}, middleware.After)
}

func addOpListExperimentResolvedTargetsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListExperimentResolvedTargets{}, middleware.After)
}

func addOpListExperimentTargetAccountConfigurationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListExperimentTargetAccountConfigurations{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpListTargetAccountConfigurationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTargetAccountConfigurations{}, middleware.After)
}

func addOpStartExperimentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartExperiment{}, middleware.After)
}

func addOpStopExperimentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopExperiment{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateExperimentTemplateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateExperimentTemplate{}, middleware.After)
}

func addOpUpdateTargetAccountConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateTargetAccountConfiguration{}, middleware.After)
}

func validateCreateExperimentTemplateActionInput(v *types.CreateExperimentTemplateActionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateExperimentTemplateActionInput"}
	if v.ActionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ActionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCreateExperimentTemplateActionInputMap(v map[string]types.CreateExperimentTemplateActionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateExperimentTemplateActionInputMap"}
	for key := range v {
		value := v[key]
		if err := validateCreateExperimentTemplateActionInput(&value); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%q]", key), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCreateExperimentTemplateLogConfigurationInput(v *types.CreateExperimentTemplateLogConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateExperimentTemplateLogConfigurationInput"}
	if v.CloudWatchLogsConfiguration != nil {
		if err := validateExperimentTemplateCloudWatchLogsLogConfigurationInput(v.CloudWatchLogsConfiguration); err != nil {
			invalidParams.AddNested("CloudWatchLogsConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.S3Configuration != nil {
		if err := validateExperimentTemplateS3LogConfigurationInput(v.S3Configuration); err != nil {
			invalidParams.AddNested("S3Configuration", err.(smithy.InvalidParamsError))
		}
	}
	if v.LogSchemaVersion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LogSchemaVersion"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCreateExperimentTemplateStopConditionInput(v *types.CreateExperimentTemplateStopConditionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateExperimentTemplateStopConditionInput"}
	if v.Source == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Source"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCreateExperimentTemplateStopConditionInputList(v []types.CreateExperimentTemplateStopConditionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateExperimentTemplateStopConditionInputList"}
	for i := range v {
		if err := validateCreateExperimentTemplateStopConditionInput(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCreateExperimentTemplateTargetInput(v *types.CreateExperimentTemplateTargetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateExperimentTemplateTargetInput"}
	if v.ResourceType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceType"))
	}
	if v.Filters != nil {
		if err := validateExperimentTemplateTargetFilterInputList(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if v.SelectionMode == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SelectionMode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCreateExperimentTemplateTargetInputMap(v map[string]types.CreateExperimentTemplateTargetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateExperimentTemplateTargetInputMap"}
	for key := range v {
		value := v[key]
		if err := validateCreateExperimentTemplateTargetInput(&value); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%q]", key), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateExperimentTemplateCloudWatchLogsLogConfigurationInput(v *types.ExperimentTemplateCloudWatchLogsLogConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExperimentTemplateCloudWatchLogsLogConfigurationInput"}
	if v.LogGroupArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LogGroupArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateExperimentTemplateS3LogConfigurationInput(v *types.ExperimentTemplateS3LogConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExperimentTemplateS3LogConfigurationInput"}
	if v.BucketName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BucketName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateExperimentTemplateTargetFilterInputList(v []types.ExperimentTemplateTargetInputFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExperimentTemplateTargetFilterInputList"}
	for i := range v {
		if err := validateExperimentTemplateTargetInputFilter(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateExperimentTemplateTargetInputFilter(v *types.ExperimentTemplateTargetInputFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExperimentTemplateTargetInputFilter"}
	if v.Path == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Path"))
	}
	if v.Values == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Values"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdateExperimentTemplateLogConfigurationInput(v *types.UpdateExperimentTemplateLogConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateExperimentTemplateLogConfigurationInput"}
	if v.CloudWatchLogsConfiguration != nil {
		if err := validateExperimentTemplateCloudWatchLogsLogConfigurationInput(v.CloudWatchLogsConfiguration); err != nil {
			invalidParams.AddNested("CloudWatchLogsConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.S3Configuration != nil {
		if err := validateExperimentTemplateS3LogConfigurationInput(v.S3Configuration); err != nil {
			invalidParams.AddNested("S3Configuration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdateExperimentTemplateStopConditionInput(v *types.UpdateExperimentTemplateStopConditionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateExperimentTemplateStopConditionInput"}
	if v.Source == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Source"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdateExperimentTemplateStopConditionInputList(v []types.UpdateExperimentTemplateStopConditionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateExperimentTemplateStopConditionInputList"}
	for i := range v {
		if err := validateUpdateExperimentTemplateStopConditionInput(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdateExperimentTemplateTargetInput(v *types.UpdateExperimentTemplateTargetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateExperimentTemplateTargetInput"}
	if v.ResourceType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceType"))
	}
	if v.Filters != nil {
		if err := validateExperimentTemplateTargetFilterInputList(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if v.SelectionMode == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SelectionMode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdateExperimentTemplateTargetInputMap(v map[string]types.UpdateExperimentTemplateTargetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateExperimentTemplateTargetInputMap"}
	for key := range v {
		value := v[key]
		if err := validateUpdateExperimentTemplateTargetInput(&value); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%q]", key), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateExperimentTemplateInput(v *CreateExperimentTemplateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateExperimentTemplateInput"}
	if v.ClientToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientToken"))
	}
	if v.Description == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Description"))
	}
	if v.StopConditions == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StopConditions"))
	} else if v.StopConditions != nil {
		if err := validateCreateExperimentTemplateStopConditionInputList(v.StopConditions); err != nil {
			invalidParams.AddNested("StopConditions", err.(smithy.InvalidParamsError))
		}
	}
	if v.Targets != nil {
		if err := validateCreateExperimentTemplateTargetInputMap(v.Targets); err != nil {
			invalidParams.AddNested("Targets", err.(smithy.InvalidParamsError))
		}
	}
	if v.Actions == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Actions"))
	} else if v.Actions != nil {
		if err := validateCreateExperimentTemplateActionInputMap(v.Actions); err != nil {
			invalidParams.AddNested("Actions", err.(smithy.InvalidParamsError))
		}
	}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if v.LogConfiguration != nil {
		if err := validateCreateExperimentTemplateLogConfigurationInput(v.LogConfiguration); err != nil {
			invalidParams.AddNested("LogConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateTargetAccountConfigurationInput(v *CreateTargetAccountConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateTargetAccountConfigurationInput"}
	if v.ExperimentTemplateId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExperimentTemplateId"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteExperimentTemplateInput(v *DeleteExperimentTemplateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteExperimentTemplateInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteTargetAccountConfigurationInput(v *DeleteTargetAccountConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteTargetAccountConfigurationInput"}
	if v.ExperimentTemplateId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExperimentTemplateId"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetActionInput(v *GetActionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetActionInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetExperimentInput(v *GetExperimentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetExperimentInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetExperimentTargetAccountConfigurationInput(v *GetExperimentTargetAccountConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetExperimentTargetAccountConfigurationInput"}
	if v.ExperimentId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExperimentId"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetExperimentTemplateInput(v *GetExperimentTemplateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetExperimentTemplateInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetTargetAccountConfigurationInput(v *GetTargetAccountConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTargetAccountConfigurationInput"}
	if v.ExperimentTemplateId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExperimentTemplateId"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetTargetResourceTypeInput(v *GetTargetResourceTypeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTargetResourceTypeInput"}
	if v.ResourceType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListExperimentResolvedTargetsInput(v *ListExperimentResolvedTargetsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListExperimentResolvedTargetsInput"}
	if v.ExperimentId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExperimentId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListExperimentTargetAccountConfigurationsInput(v *ListExperimentTargetAccountConfigurationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListExperimentTargetAccountConfigurationsInput"}
	if v.ExperimentId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExperimentId"))
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
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTargetAccountConfigurationsInput(v *ListTargetAccountConfigurationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTargetAccountConfigurationsInput"}
	if v.ExperimentTemplateId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExperimentTemplateId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartExperimentInput(v *StartExperimentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartExperimentInput"}
	if v.ClientToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientToken"))
	}
	if v.ExperimentTemplateId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExperimentTemplateId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopExperimentInput(v *StopExperimentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopExperimentInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
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
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
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
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateExperimentTemplateInput(v *UpdateExperimentTemplateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateExperimentTemplateInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.StopConditions != nil {
		if err := validateUpdateExperimentTemplateStopConditionInputList(v.StopConditions); err != nil {
			invalidParams.AddNested("StopConditions", err.(smithy.InvalidParamsError))
		}
	}
	if v.Targets != nil {
		if err := validateUpdateExperimentTemplateTargetInputMap(v.Targets); err != nil {
			invalidParams.AddNested("Targets", err.(smithy.InvalidParamsError))
		}
	}
	if v.LogConfiguration != nil {
		if err := validateUpdateExperimentTemplateLogConfigurationInput(v.LogConfiguration); err != nil {
			invalidParams.AddNested("LogConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateTargetAccountConfigurationInput(v *UpdateTargetAccountConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateTargetAccountConfigurationInput"}
	if v.ExperimentTemplateId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExperimentTemplateId"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
