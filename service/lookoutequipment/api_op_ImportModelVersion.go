// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutequipment

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/lookoutequipment/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Imports a model that has been trained successfully.
func (c *Client) ImportModelVersion(ctx context.Context, params *ImportModelVersionInput, optFns ...func(*Options)) (*ImportModelVersionOutput, error) {
	if params == nil {
		params = &ImportModelVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportModelVersion", params, optFns, c.addOperationImportModelVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportModelVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportModelVersionInput struct {

	// A unique identifier for the request. If you do not set the client request
	// token, Amazon Lookout for Equipment generates one.
	//
	// This member is required.
	ClientToken *string

	// The name of the dataset for the machine learning model being imported.
	//
	// This member is required.
	DatasetName *string

	// The Amazon Resource Name (ARN) of the model version to import.
	//
	// This member is required.
	SourceModelVersionArn *string

	// Indicates how to import the accumulated inference data when a model version is
	// imported. The possible values are as follows:
	//   - NO_IMPORT – Don't import the data.
	//   - ADD_WHEN_EMPTY – Only import the data from the source model if there is no
	//   existing data in the target model.
	//   - OVERWRITE – Import the data from the source model and overwrite the
	//   existing data in the target model.
	InferenceDataImportStrategy types.InferenceDataImportStrategy

	// Contains the configuration information for the S3 location being used to hold
	// label data.
	LabelsInputConfiguration *types.LabelsInputConfiguration

	// The name for the machine learning model to be created. If the model already
	// exists, Amazon Lookout for Equipment creates a new version. If you do not
	// specify this field, it is filled with the name of the source model.
	ModelName *string

	// The Amazon Resource Name (ARN) of a role with permission to access the data
	// source being used to create the machine learning model.
	RoleArn *string

	// Provides the identifier of the KMS key key used to encrypt model data by Amazon
	// Lookout for Equipment.
	ServerSideKmsKeyId *string

	// The tags associated with the machine learning model to be created.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type ImportModelVersionOutput struct {

	// The Amazon Resource Name (ARN) of the model being created.
	ModelArn *string

	// The name for the machine learning model.
	ModelName *string

	// The version of the model being created.
	ModelVersion *int64

	// The Amazon Resource Name (ARN) of the model version being created.
	ModelVersionArn *string

	// The status of the ImportModelVersion operation.
	Status types.ModelVersionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportModelVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpImportModelVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpImportModelVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ImportModelVersion"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opImportModelVersionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpImportModelVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportModelVersion(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpImportModelVersion struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpImportModelVersion) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpImportModelVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ImportModelVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ImportModelVersionInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opImportModelVersionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpImportModelVersion{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opImportModelVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ImportModelVersion",
	}
}
