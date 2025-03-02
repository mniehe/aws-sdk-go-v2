// Code generated by smithy-go-codegen DO NOT EDIT.

package bcmdataexports

import (
	"context"
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/bcmdataexports/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateExport struct {
}

func (*validateOpCreateExport) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateExport) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateExportInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateExportInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteExport struct {
}

func (*validateOpDeleteExport) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteExport) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteExportInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteExportInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetExecution struct {
}

func (*validateOpGetExecution) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetExecution) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetExecutionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetExecutionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetExport struct {
}

func (*validateOpGetExport) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetExport) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetExportInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetExportInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetTable struct {
}

func (*validateOpGetTable) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTable) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTableInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTableInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListExecutions struct {
}

func (*validateOpListExecutions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListExecutions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListExecutionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListExecutionsInput(input); err != nil {
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

type validateOpUpdateExport struct {
}

func (*validateOpUpdateExport) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateExport) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateExportInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateExportInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateExportValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateExport{}, middleware.After)
}

func addOpDeleteExportValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteExport{}, middleware.After)
}

func addOpGetExecutionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetExecution{}, middleware.After)
}

func addOpGetExportValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetExport{}, middleware.After)
}

func addOpGetTableValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTable{}, middleware.After)
}

func addOpListExecutionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListExecutions{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateExportValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateExport{}, middleware.After)
}

func validateDataQuery(v *types.DataQuery) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DataQuery"}
	if v.QueryStatement == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueryStatement"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDestinationConfigurations(v *types.DestinationConfigurations) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DestinationConfigurations"}
	if v.S3Destination == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Destination"))
	} else if v.S3Destination != nil {
		if err := validateS3Destination(v.S3Destination); err != nil {
			invalidParams.AddNested("S3Destination", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateExport(v *types.Export) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Export"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.DataQuery == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataQuery"))
	} else if v.DataQuery != nil {
		if err := validateDataQuery(v.DataQuery); err != nil {
			invalidParams.AddNested("DataQuery", err.(smithy.InvalidParamsError))
		}
	}
	if v.DestinationConfigurations == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DestinationConfigurations"))
	} else if v.DestinationConfigurations != nil {
		if err := validateDestinationConfigurations(v.DestinationConfigurations); err != nil {
			invalidParams.AddNested("DestinationConfigurations", err.(smithy.InvalidParamsError))
		}
	}
	if v.RefreshCadence == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RefreshCadence"))
	} else if v.RefreshCadence != nil {
		if err := validateRefreshCadence(v.RefreshCadence); err != nil {
			invalidParams.AddNested("RefreshCadence", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRefreshCadence(v *types.RefreshCadence) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RefreshCadence"}
	if len(v.Frequency) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Frequency"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateResourceTag(v *types.ResourceTag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResourceTag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
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

func validateResourceTagList(v []types.ResourceTag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResourceTagList"}
	for i := range v {
		if err := validateResourceTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3Destination(v *types.S3Destination) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3Destination"}
	if v.S3Bucket == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Bucket"))
	}
	if v.S3Prefix == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Prefix"))
	}
	if v.S3Region == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Region"))
	}
	if v.S3OutputConfigurations == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3OutputConfigurations"))
	} else if v.S3OutputConfigurations != nil {
		if err := validateS3OutputConfigurations(v.S3OutputConfigurations); err != nil {
			invalidParams.AddNested("S3OutputConfigurations", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3OutputConfigurations(v *types.S3OutputConfigurations) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3OutputConfigurations"}
	if len(v.OutputType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("OutputType"))
	}
	if len(v.Format) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Format"))
	}
	if len(v.Compression) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Compression"))
	}
	if len(v.Overwrite) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Overwrite"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateExportInput(v *CreateExportInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateExportInput"}
	if v.Export == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Export"))
	} else if v.Export != nil {
		if err := validateExport(v.Export); err != nil {
			invalidParams.AddNested("Export", err.(smithy.InvalidParamsError))
		}
	}
	if v.ResourceTags != nil {
		if err := validateResourceTagList(v.ResourceTags); err != nil {
			invalidParams.AddNested("ResourceTags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteExportInput(v *DeleteExportInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteExportInput"}
	if v.ExportArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExportArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetExecutionInput(v *GetExecutionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetExecutionInput"}
	if v.ExportArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExportArn"))
	}
	if v.ExecutionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExecutionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetExportInput(v *GetExportInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetExportInput"}
	if v.ExportArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExportArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetTableInput(v *GetTableInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTableInput"}
	if v.TableName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TableName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListExecutionsInput(v *ListExecutionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListExecutionsInput"}
	if v.ExportArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExportArn"))
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

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.ResourceTags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceTags"))
	} else if v.ResourceTags != nil {
		if err := validateResourceTagList(v.ResourceTags); err != nil {
			invalidParams.AddNested("ResourceTags", err.(smithy.InvalidParamsError))
		}
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
	if v.ResourceTagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceTagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateExportInput(v *UpdateExportInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateExportInput"}
	if v.ExportArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExportArn"))
	}
	if v.Export == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Export"))
	} else if v.Export != nil {
		if err := validateExport(v.Export); err != nil {
			invalidParams.AddNested("Export", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
