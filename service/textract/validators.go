// Code generated by smithy-go-codegen DO NOT EDIT.

package textract

import (
	"context"
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/textract/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAnalyzeDocument struct {
}

func (*validateOpAnalyzeDocument) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAnalyzeDocument) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AnalyzeDocumentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAnalyzeDocumentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAnalyzeExpense struct {
}

func (*validateOpAnalyzeExpense) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAnalyzeExpense) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AnalyzeExpenseInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAnalyzeExpenseInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAnalyzeID struct {
}

func (*validateOpAnalyzeID) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAnalyzeID) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AnalyzeIDInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAnalyzeIDInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateAdapter struct {
}

func (*validateOpCreateAdapter) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateAdapter) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateAdapterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateAdapterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateAdapterVersion struct {
}

func (*validateOpCreateAdapterVersion) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateAdapterVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateAdapterVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateAdapterVersionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteAdapter struct {
}

func (*validateOpDeleteAdapter) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteAdapter) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteAdapterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteAdapterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteAdapterVersion struct {
}

func (*validateOpDeleteAdapterVersion) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteAdapterVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteAdapterVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteAdapterVersionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDetectDocumentText struct {
}

func (*validateOpDetectDocumentText) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDetectDocumentText) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DetectDocumentTextInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDetectDocumentTextInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetAdapter struct {
}

func (*validateOpGetAdapter) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetAdapter) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetAdapterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetAdapterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetAdapterVersion struct {
}

func (*validateOpGetAdapterVersion) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetAdapterVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetAdapterVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetAdapterVersionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetDocumentAnalysis struct {
}

func (*validateOpGetDocumentAnalysis) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetDocumentAnalysis) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetDocumentAnalysisInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetDocumentAnalysisInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetDocumentTextDetection struct {
}

func (*validateOpGetDocumentTextDetection) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetDocumentTextDetection) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetDocumentTextDetectionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetDocumentTextDetectionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetExpenseAnalysis struct {
}

func (*validateOpGetExpenseAnalysis) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetExpenseAnalysis) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetExpenseAnalysisInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetExpenseAnalysisInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetLendingAnalysis struct {
}

func (*validateOpGetLendingAnalysis) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetLendingAnalysis) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetLendingAnalysisInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetLendingAnalysisInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetLendingAnalysisSummary struct {
}

func (*validateOpGetLendingAnalysisSummary) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetLendingAnalysisSummary) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetLendingAnalysisSummaryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetLendingAnalysisSummaryInput(input); err != nil {
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

type validateOpStartDocumentAnalysis struct {
}

func (*validateOpStartDocumentAnalysis) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartDocumentAnalysis) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartDocumentAnalysisInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartDocumentAnalysisInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartDocumentTextDetection struct {
}

func (*validateOpStartDocumentTextDetection) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartDocumentTextDetection) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartDocumentTextDetectionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartDocumentTextDetectionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartExpenseAnalysis struct {
}

func (*validateOpStartExpenseAnalysis) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartExpenseAnalysis) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartExpenseAnalysisInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartExpenseAnalysisInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartLendingAnalysis struct {
}

func (*validateOpStartLendingAnalysis) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartLendingAnalysis) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartLendingAnalysisInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartLendingAnalysisInput(input); err != nil {
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

type validateOpUpdateAdapter struct {
}

func (*validateOpUpdateAdapter) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateAdapter) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateAdapterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateAdapterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAnalyzeDocumentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAnalyzeDocument{}, middleware.After)
}

func addOpAnalyzeExpenseValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAnalyzeExpense{}, middleware.After)
}

func addOpAnalyzeIDValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAnalyzeID{}, middleware.After)
}

func addOpCreateAdapterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateAdapter{}, middleware.After)
}

func addOpCreateAdapterVersionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateAdapterVersion{}, middleware.After)
}

func addOpDeleteAdapterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteAdapter{}, middleware.After)
}

func addOpDeleteAdapterVersionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteAdapterVersion{}, middleware.After)
}

func addOpDetectDocumentTextValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDetectDocumentText{}, middleware.After)
}

func addOpGetAdapterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetAdapter{}, middleware.After)
}

func addOpGetAdapterVersionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetAdapterVersion{}, middleware.After)
}

func addOpGetDocumentAnalysisValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetDocumentAnalysis{}, middleware.After)
}

func addOpGetDocumentTextDetectionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetDocumentTextDetection{}, middleware.After)
}

func addOpGetExpenseAnalysisValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetExpenseAnalysis{}, middleware.After)
}

func addOpGetLendingAnalysisValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetLendingAnalysis{}, middleware.After)
}

func addOpGetLendingAnalysisSummaryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetLendingAnalysisSummary{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpStartDocumentAnalysisValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartDocumentAnalysis{}, middleware.After)
}

func addOpStartDocumentTextDetectionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartDocumentTextDetection{}, middleware.After)
}

func addOpStartExpenseAnalysisValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartExpenseAnalysis{}, middleware.After)
}

func addOpStartLendingAnalysisValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartLendingAnalysis{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateAdapterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateAdapter{}, middleware.After)
}

func validateAdapter(v *types.Adapter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Adapter"}
	if v.AdapterId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AdapterId"))
	}
	if v.Version == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Version"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateAdapters(v []types.Adapter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Adapters"}
	for i := range v {
		if err := validateAdapter(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateAdaptersConfig(v *types.AdaptersConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AdaptersConfig"}
	if v.Adapters == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Adapters"))
	} else if v.Adapters != nil {
		if err := validateAdapters(v.Adapters); err != nil {
			invalidParams.AddNested("Adapters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateHumanLoopConfig(v *types.HumanLoopConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HumanLoopConfig"}
	if v.HumanLoopName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("HumanLoopName"))
	}
	if v.FlowDefinitionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FlowDefinitionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNotificationChannel(v *types.NotificationChannel) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NotificationChannel"}
	if v.SNSTopicArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SNSTopicArn"))
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

func validateOutputConfig(v *types.OutputConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "OutputConfig"}
	if v.S3Bucket == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Bucket"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateQueries(v []types.Query) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Queries"}
	for i := range v {
		if err := validateQuery(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateQueriesConfig(v *types.QueriesConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "QueriesConfig"}
	if v.Queries == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Queries"))
	} else if v.Queries != nil {
		if err := validateQueries(v.Queries); err != nil {
			invalidParams.AddNested("Queries", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateQuery(v *types.Query) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Query"}
	if v.Text == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Text"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAnalyzeDocumentInput(v *AnalyzeDocumentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AnalyzeDocumentInput"}
	if v.Document == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Document"))
	}
	if v.FeatureTypes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FeatureTypes"))
	}
	if v.HumanLoopConfig != nil {
		if err := validateHumanLoopConfig(v.HumanLoopConfig); err != nil {
			invalidParams.AddNested("HumanLoopConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.QueriesConfig != nil {
		if err := validateQueriesConfig(v.QueriesConfig); err != nil {
			invalidParams.AddNested("QueriesConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.AdaptersConfig != nil {
		if err := validateAdaptersConfig(v.AdaptersConfig); err != nil {
			invalidParams.AddNested("AdaptersConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAnalyzeExpenseInput(v *AnalyzeExpenseInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AnalyzeExpenseInput"}
	if v.Document == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Document"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAnalyzeIDInput(v *AnalyzeIDInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AnalyzeIDInput"}
	if v.DocumentPages == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DocumentPages"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateAdapterInput(v *CreateAdapterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateAdapterInput"}
	if v.AdapterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AdapterName"))
	}
	if v.FeatureTypes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FeatureTypes"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateAdapterVersionInput(v *CreateAdapterVersionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateAdapterVersionInput"}
	if v.AdapterId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AdapterId"))
	}
	if v.DatasetConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatasetConfig"))
	}
	if v.OutputConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputConfig"))
	} else if v.OutputConfig != nil {
		if err := validateOutputConfig(v.OutputConfig); err != nil {
			invalidParams.AddNested("OutputConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteAdapterInput(v *DeleteAdapterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteAdapterInput"}
	if v.AdapterId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AdapterId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteAdapterVersionInput(v *DeleteAdapterVersionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteAdapterVersionInput"}
	if v.AdapterId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AdapterId"))
	}
	if v.AdapterVersion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AdapterVersion"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDetectDocumentTextInput(v *DetectDocumentTextInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DetectDocumentTextInput"}
	if v.Document == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Document"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetAdapterInput(v *GetAdapterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetAdapterInput"}
	if v.AdapterId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AdapterId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetAdapterVersionInput(v *GetAdapterVersionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetAdapterVersionInput"}
	if v.AdapterId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AdapterId"))
	}
	if v.AdapterVersion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AdapterVersion"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetDocumentAnalysisInput(v *GetDocumentAnalysisInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetDocumentAnalysisInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetDocumentTextDetectionInput(v *GetDocumentTextDetectionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetDocumentTextDetectionInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetExpenseAnalysisInput(v *GetExpenseAnalysisInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetExpenseAnalysisInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetLendingAnalysisInput(v *GetLendingAnalysisInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetLendingAnalysisInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetLendingAnalysisSummaryInput(v *GetLendingAnalysisSummaryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetLendingAnalysisSummaryInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
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

func validateOpStartDocumentAnalysisInput(v *StartDocumentAnalysisInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartDocumentAnalysisInput"}
	if v.DocumentLocation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DocumentLocation"))
	}
	if v.FeatureTypes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FeatureTypes"))
	}
	if v.NotificationChannel != nil {
		if err := validateNotificationChannel(v.NotificationChannel); err != nil {
			invalidParams.AddNested("NotificationChannel", err.(smithy.InvalidParamsError))
		}
	}
	if v.OutputConfig != nil {
		if err := validateOutputConfig(v.OutputConfig); err != nil {
			invalidParams.AddNested("OutputConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.QueriesConfig != nil {
		if err := validateQueriesConfig(v.QueriesConfig); err != nil {
			invalidParams.AddNested("QueriesConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.AdaptersConfig != nil {
		if err := validateAdaptersConfig(v.AdaptersConfig); err != nil {
			invalidParams.AddNested("AdaptersConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartDocumentTextDetectionInput(v *StartDocumentTextDetectionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartDocumentTextDetectionInput"}
	if v.DocumentLocation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DocumentLocation"))
	}
	if v.NotificationChannel != nil {
		if err := validateNotificationChannel(v.NotificationChannel); err != nil {
			invalidParams.AddNested("NotificationChannel", err.(smithy.InvalidParamsError))
		}
	}
	if v.OutputConfig != nil {
		if err := validateOutputConfig(v.OutputConfig); err != nil {
			invalidParams.AddNested("OutputConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartExpenseAnalysisInput(v *StartExpenseAnalysisInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartExpenseAnalysisInput"}
	if v.DocumentLocation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DocumentLocation"))
	}
	if v.NotificationChannel != nil {
		if err := validateNotificationChannel(v.NotificationChannel); err != nil {
			invalidParams.AddNested("NotificationChannel", err.(smithy.InvalidParamsError))
		}
	}
	if v.OutputConfig != nil {
		if err := validateOutputConfig(v.OutputConfig); err != nil {
			invalidParams.AddNested("OutputConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartLendingAnalysisInput(v *StartLendingAnalysisInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartLendingAnalysisInput"}
	if v.DocumentLocation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DocumentLocation"))
	}
	if v.NotificationChannel != nil {
		if err := validateNotificationChannel(v.NotificationChannel); err != nil {
			invalidParams.AddNested("NotificationChannel", err.(smithy.InvalidParamsError))
		}
	}
	if v.OutputConfig != nil {
		if err := validateOutputConfig(v.OutputConfig); err != nil {
			invalidParams.AddNested("OutputConfig", err.(smithy.InvalidParamsError))
		}
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

func validateOpUpdateAdapterInput(v *UpdateAdapterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateAdapterInput"}
	if v.AdapterId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AdapterId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
