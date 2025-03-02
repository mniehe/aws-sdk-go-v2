// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagent

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/bedrockagent/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a new knowledge base
func (c *Client) CreateKnowledgeBase(ctx context.Context, params *CreateKnowledgeBaseInput, optFns ...func(*Options)) (*CreateKnowledgeBaseOutput, error) {
	if params == nil {
		params = &CreateKnowledgeBaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateKnowledgeBase", params, optFns, c.addOperationCreateKnowledgeBaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateKnowledgeBaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateKnowledgeBaseInput struct {

	// Configures a bedrock knowledge base.
	//
	// This member is required.
	KnowledgeBaseConfiguration *types.KnowledgeBaseConfiguration

	// Name for a resource.
	//
	// This member is required.
	Name *string

	// ARN of a IAM role.
	//
	// This member is required.
	RoleArn *string

	// Configures the physical storage of ingested data in a knowledge base.
	//
	// This member is required.
	StorageConfiguration *types.StorageConfiguration

	// Client specified token used for idempotency checks
	ClientToken *string

	// Description of the Resource.
	Description *string

	// A map of tag keys and values
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateKnowledgeBaseOutput struct {

	// Contains the information of a knowledge base.
	//
	// This member is required.
	KnowledgeBase *types.KnowledgeBase

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateKnowledgeBaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateKnowledgeBase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateKnowledgeBase{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateKnowledgeBase"); err != nil {
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
	if err = addIdempotencyToken_opCreateKnowledgeBaseMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateKnowledgeBaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateKnowledgeBase(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateKnowledgeBase struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateKnowledgeBase) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateKnowledgeBase) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateKnowledgeBaseInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateKnowledgeBaseInput ")
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
func addIdempotencyToken_opCreateKnowledgeBaseMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateKnowledgeBase{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateKnowledgeBase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateKnowledgeBase",
	}
}
