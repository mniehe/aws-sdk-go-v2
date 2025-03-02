// Code generated by smithy-go-codegen DO NOT EDIT.

package arczonalshift

import (
	"context"
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/arczonalshift/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCancelZonalShift struct {
}

func (*validateOpCancelZonalShift) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelZonalShift) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelZonalShiftInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelZonalShiftInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreatePracticeRunConfiguration struct {
}

func (*validateOpCreatePracticeRunConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreatePracticeRunConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreatePracticeRunConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreatePracticeRunConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeletePracticeRunConfiguration struct {
}

func (*validateOpDeletePracticeRunConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeletePracticeRunConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeletePracticeRunConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeletePracticeRunConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetManagedResource struct {
}

func (*validateOpGetManagedResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetManagedResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetManagedResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetManagedResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartZonalShift struct {
}

func (*validateOpStartZonalShift) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartZonalShift) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartZonalShiftInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartZonalShiftInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdatePracticeRunConfiguration struct {
}

func (*validateOpUpdatePracticeRunConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdatePracticeRunConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdatePracticeRunConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdatePracticeRunConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateZonalAutoshiftConfiguration struct {
}

func (*validateOpUpdateZonalAutoshiftConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateZonalAutoshiftConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateZonalAutoshiftConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateZonalAutoshiftConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateZonalShift struct {
}

func (*validateOpUpdateZonalShift) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateZonalShift) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateZonalShiftInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateZonalShiftInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCancelZonalShiftValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelZonalShift{}, middleware.After)
}

func addOpCreatePracticeRunConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreatePracticeRunConfiguration{}, middleware.After)
}

func addOpDeletePracticeRunConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeletePracticeRunConfiguration{}, middleware.After)
}

func addOpGetManagedResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetManagedResource{}, middleware.After)
}

func addOpStartZonalShiftValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartZonalShift{}, middleware.After)
}

func addOpUpdatePracticeRunConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdatePracticeRunConfiguration{}, middleware.After)
}

func addOpUpdateZonalAutoshiftConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateZonalAutoshiftConfiguration{}, middleware.After)
}

func addOpUpdateZonalShiftValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateZonalShift{}, middleware.After)
}

func validateControlCondition(v *types.ControlCondition) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ControlCondition"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.AlarmIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AlarmIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateControlConditions(v []types.ControlCondition) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ControlConditions"}
	for i := range v {
		if err := validateControlCondition(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelZonalShiftInput(v *CancelZonalShiftInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelZonalShiftInput"}
	if v.ZonalShiftId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ZonalShiftId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreatePracticeRunConfigurationInput(v *CreatePracticeRunConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreatePracticeRunConfigurationInput"}
	if v.ResourceIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceIdentifier"))
	}
	if v.BlockingAlarms != nil {
		if err := validateControlConditions(v.BlockingAlarms); err != nil {
			invalidParams.AddNested("BlockingAlarms", err.(smithy.InvalidParamsError))
		}
	}
	if v.OutcomeAlarms == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutcomeAlarms"))
	} else if v.OutcomeAlarms != nil {
		if err := validateControlConditions(v.OutcomeAlarms); err != nil {
			invalidParams.AddNested("OutcomeAlarms", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeletePracticeRunConfigurationInput(v *DeletePracticeRunConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeletePracticeRunConfigurationInput"}
	if v.ResourceIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetManagedResourceInput(v *GetManagedResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetManagedResourceInput"}
	if v.ResourceIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartZonalShiftInput(v *StartZonalShiftInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartZonalShiftInput"}
	if v.ResourceIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceIdentifier"))
	}
	if v.AwayFrom == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AwayFrom"))
	}
	if v.ExpiresIn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExpiresIn"))
	}
	if v.Comment == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Comment"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdatePracticeRunConfigurationInput(v *UpdatePracticeRunConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdatePracticeRunConfigurationInput"}
	if v.ResourceIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceIdentifier"))
	}
	if v.BlockingAlarms != nil {
		if err := validateControlConditions(v.BlockingAlarms); err != nil {
			invalidParams.AddNested("BlockingAlarms", err.(smithy.InvalidParamsError))
		}
	}
	if v.OutcomeAlarms != nil {
		if err := validateControlConditions(v.OutcomeAlarms); err != nil {
			invalidParams.AddNested("OutcomeAlarms", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateZonalAutoshiftConfigurationInput(v *UpdateZonalAutoshiftConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateZonalAutoshiftConfigurationInput"}
	if v.ResourceIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceIdentifier"))
	}
	if len(v.ZonalAutoshiftStatus) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ZonalAutoshiftStatus"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateZonalShiftInput(v *UpdateZonalShiftInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateZonalShiftInput"}
	if v.ZonalShiftId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ZonalShiftId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
