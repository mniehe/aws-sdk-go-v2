// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc10

import (
	"context"
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/internal/protocoltest/jsonrpc10/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpEndpointWithHostLabelOperation struct {
}

func (*validateOpEndpointWithHostLabelOperation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpEndpointWithHostLabelOperation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*EndpointWithHostLabelOperationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpEndpointWithHostLabelOperationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpOperationWithNestedStructure struct {
}

func (*validateOpOperationWithNestedStructure) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpOperationWithNestedStructure) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*OperationWithNestedStructureInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpOperationWithNestedStructureInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpEndpointWithHostLabelOperationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpEndpointWithHostLabelOperation{}, middleware.After)
}

func addOpOperationWithNestedStructureValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpOperationWithNestedStructure{}, middleware.After)
}

func validateTopLevel(v *types.TopLevel) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TopLevel"}
	if v.Dialog == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Dialog"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpEndpointWithHostLabelOperationInput(v *EndpointWithHostLabelOperationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EndpointWithHostLabelOperationInput"}
	if v.Label == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Label"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpOperationWithNestedStructureInput(v *OperationWithNestedStructureInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "OperationWithNestedStructureInput"}
	if v.TopLevel == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TopLevel"))
	} else if v.TopLevel != nil {
		if err := validateTopLevel(v.TopLevel); err != nil {
			invalidParams.AddNested("TopLevel", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
