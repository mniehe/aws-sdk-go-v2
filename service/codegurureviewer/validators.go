// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurureviewer

import (
	"context"
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/codegurureviewer/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAssociateRepository struct {
}

func (*validateOpAssociateRepository) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateRepository) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateRepositoryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateRepositoryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateCodeReview struct {
}

func (*validateOpCreateCodeReview) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateCodeReview) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateCodeReviewInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateCodeReviewInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeCodeReview struct {
}

func (*validateOpDescribeCodeReview) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeCodeReview) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeCodeReviewInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeCodeReviewInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeRecommendationFeedback struct {
}

func (*validateOpDescribeRecommendationFeedback) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeRecommendationFeedback) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeRecommendationFeedbackInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeRecommendationFeedbackInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeRepositoryAssociation struct {
}

func (*validateOpDescribeRepositoryAssociation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeRepositoryAssociation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeRepositoryAssociationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeRepositoryAssociationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateRepository struct {
}

func (*validateOpDisassociateRepository) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateRepository) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateRepositoryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateRepositoryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListCodeReviews struct {
}

func (*validateOpListCodeReviews) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListCodeReviews) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListCodeReviewsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListCodeReviewsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListRecommendationFeedback struct {
}

func (*validateOpListRecommendationFeedback) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListRecommendationFeedback) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListRecommendationFeedbackInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListRecommendationFeedbackInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListRecommendations struct {
}

func (*validateOpListRecommendations) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListRecommendations) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListRecommendationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListRecommendationsInput(input); err != nil {
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

type validateOpPutRecommendationFeedback struct {
}

func (*validateOpPutRecommendationFeedback) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutRecommendationFeedback) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutRecommendationFeedbackInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutRecommendationFeedbackInput(input); err != nil {
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

func addOpAssociateRepositoryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateRepository{}, middleware.After)
}

func addOpCreateCodeReviewValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateCodeReview{}, middleware.After)
}

func addOpDescribeCodeReviewValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeCodeReview{}, middleware.After)
}

func addOpDescribeRecommendationFeedbackValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeRecommendationFeedback{}, middleware.After)
}

func addOpDescribeRepositoryAssociationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeRepositoryAssociation{}, middleware.After)
}

func addOpDisassociateRepositoryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateRepository{}, middleware.After)
}

func addOpListCodeReviewsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListCodeReviews{}, middleware.After)
}

func addOpListRecommendationFeedbackValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListRecommendationFeedback{}, middleware.After)
}

func addOpListRecommendationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListRecommendations{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPutRecommendationFeedbackValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutRecommendationFeedback{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func validateBranchDiffSourceCodeType(v *types.BranchDiffSourceCodeType) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BranchDiffSourceCodeType"}
	if v.SourceBranchName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceBranchName"))
	}
	if v.DestinationBranchName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DestinationBranchName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCodeArtifacts(v *types.CodeArtifacts) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CodeArtifacts"}
	if v.SourceCodeArtifactsObjectKey == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceCodeArtifactsObjectKey"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCodeCommitRepository(v *types.CodeCommitRepository) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CodeCommitRepository"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCodeReviewType(v *types.CodeReviewType) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CodeReviewType"}
	if v.RepositoryAnalysis == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RepositoryAnalysis"))
	} else if v.RepositoryAnalysis != nil {
		if err := validateRepositoryAnalysis(v.RepositoryAnalysis); err != nil {
			invalidParams.AddNested("RepositoryAnalysis", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRepository(v *types.Repository) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Repository"}
	if v.CodeCommit != nil {
		if err := validateCodeCommitRepository(v.CodeCommit); err != nil {
			invalidParams.AddNested("CodeCommit", err.(smithy.InvalidParamsError))
		}
	}
	if v.Bitbucket != nil {
		if err := validateThirdPartySourceRepository(v.Bitbucket); err != nil {
			invalidParams.AddNested("Bitbucket", err.(smithy.InvalidParamsError))
		}
	}
	if v.GitHubEnterpriseServer != nil {
		if err := validateThirdPartySourceRepository(v.GitHubEnterpriseServer); err != nil {
			invalidParams.AddNested("GitHubEnterpriseServer", err.(smithy.InvalidParamsError))
		}
	}
	if v.S3Bucket != nil {
		if err := validateS3Repository(v.S3Bucket); err != nil {
			invalidParams.AddNested("S3Bucket", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRepositoryAnalysis(v *types.RepositoryAnalysis) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RepositoryAnalysis"}
	if v.RepositoryHead != nil {
		if err := validateRepositoryHeadSourceCodeType(v.RepositoryHead); err != nil {
			invalidParams.AddNested("RepositoryHead", err.(smithy.InvalidParamsError))
		}
	}
	if v.SourceCodeType != nil {
		if err := validateSourceCodeType(v.SourceCodeType); err != nil {
			invalidParams.AddNested("SourceCodeType", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRepositoryHeadSourceCodeType(v *types.RepositoryHeadSourceCodeType) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RepositoryHeadSourceCodeType"}
	if v.BranchName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BranchName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3BucketRepository(v *types.S3BucketRepository) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3BucketRepository"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Details != nil {
		if err := validateS3RepositoryDetails(v.Details); err != nil {
			invalidParams.AddNested("Details", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3Repository(v *types.S3Repository) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3Repository"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.BucketName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BucketName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3RepositoryDetails(v *types.S3RepositoryDetails) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3RepositoryDetails"}
	if v.CodeArtifacts != nil {
		if err := validateCodeArtifacts(v.CodeArtifacts); err != nil {
			invalidParams.AddNested("CodeArtifacts", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSourceCodeType(v *types.SourceCodeType) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SourceCodeType"}
	if v.RepositoryHead != nil {
		if err := validateRepositoryHeadSourceCodeType(v.RepositoryHead); err != nil {
			invalidParams.AddNested("RepositoryHead", err.(smithy.InvalidParamsError))
		}
	}
	if v.BranchDiff != nil {
		if err := validateBranchDiffSourceCodeType(v.BranchDiff); err != nil {
			invalidParams.AddNested("BranchDiff", err.(smithy.InvalidParamsError))
		}
	}
	if v.S3BucketRepository != nil {
		if err := validateS3BucketRepository(v.S3BucketRepository); err != nil {
			invalidParams.AddNested("S3BucketRepository", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateThirdPartySourceRepository(v *types.ThirdPartySourceRepository) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ThirdPartySourceRepository"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.ConnectionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectionArn"))
	}
	if v.Owner == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Owner"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateRepositoryInput(v *AssociateRepositoryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateRepositoryInput"}
	if v.Repository == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Repository"))
	} else if v.Repository != nil {
		if err := validateRepository(v.Repository); err != nil {
			invalidParams.AddNested("Repository", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateCodeReviewInput(v *CreateCodeReviewInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateCodeReviewInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.RepositoryAssociationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RepositoryAssociationArn"))
	}
	if v.Type == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	} else if v.Type != nil {
		if err := validateCodeReviewType(v.Type); err != nil {
			invalidParams.AddNested("Type", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeCodeReviewInput(v *DescribeCodeReviewInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeCodeReviewInput"}
	if v.CodeReviewArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CodeReviewArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeRecommendationFeedbackInput(v *DescribeRecommendationFeedbackInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeRecommendationFeedbackInput"}
	if v.CodeReviewArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CodeReviewArn"))
	}
	if v.RecommendationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RecommendationId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeRepositoryAssociationInput(v *DescribeRepositoryAssociationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeRepositoryAssociationInput"}
	if v.AssociationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AssociationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateRepositoryInput(v *DisassociateRepositoryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateRepositoryInput"}
	if v.AssociationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AssociationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListCodeReviewsInput(v *ListCodeReviewsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListCodeReviewsInput"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListRecommendationFeedbackInput(v *ListRecommendationFeedbackInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListRecommendationFeedbackInput"}
	if v.CodeReviewArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CodeReviewArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListRecommendationsInput(v *ListRecommendationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListRecommendationsInput"}
	if v.CodeReviewArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CodeReviewArn"))
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

func validateOpPutRecommendationFeedbackInput(v *PutRecommendationFeedbackInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutRecommendationFeedbackInput"}
	if v.CodeReviewArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CodeReviewArn"))
	}
	if v.RecommendationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RecommendationId"))
	}
	if v.Reactions == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Reactions"))
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
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
