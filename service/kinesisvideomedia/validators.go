// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideomedia

import (
	"context"
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/kinesisvideomedia/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpGetMedia struct {
}

func (*validateOpGetMedia) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetMedia) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetMediaInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetMediaInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpGetMediaValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetMedia{}, middleware.After)
}

func validateStartSelector(v *types.StartSelector) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartSelector"}
	if len(v.StartSelectorType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("StartSelectorType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetMediaInput(v *GetMediaInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetMediaInput"}
	if v.StartSelector == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartSelector"))
	} else if v.StartSelector != nil {
		if err := validateStartSelector(v.StartSelector); err != nil {
			invalidParams.AddNested("StartSelector", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
