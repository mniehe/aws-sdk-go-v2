// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearchserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/opensearchserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a lifecyle policy to be applied to OpenSearch Serverless indexes.
// Lifecycle policies define the number of days or hours to retain the data on an
// OpenSearch Serverless index. For more information, see Creating data lifecycle
// policies (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-lifecycle.html#serverless-lifecycle-create)
// .
func (c *Client) CreateLifecyclePolicy(ctx context.Context, params *CreateLifecyclePolicyInput, optFns ...func(*Options)) (*CreateLifecyclePolicyOutput, error) {
	if params == nil {
		params = &CreateLifecyclePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLifecyclePolicy", params, optFns, c.addOperationCreateLifecyclePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLifecyclePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLifecyclePolicyInput struct {

	// The name of the lifecycle policy.
	//
	// This member is required.
	Name *string

	// The JSON policy document to use as the content for the lifecycle policy.
	//
	// This member is required.
	Policy *string

	// The type of lifecycle policy.
	//
	// This member is required.
	Type types.LifecyclePolicyType

	// A unique, case-sensitive identifier to ensure idempotency of the request.
	ClientToken *string

	// A description of the lifecycle policy.
	Description *string

	noSmithyDocumentSerde
}

type CreateLifecyclePolicyOutput struct {

	// Details about the created lifecycle policy.
	LifecyclePolicyDetail *types.LifecyclePolicyDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLifecyclePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateLifecyclePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateLifecyclePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLifecyclePolicy"); err != nil {
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
	if err = addIdempotencyToken_opCreateLifecyclePolicyMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateLifecyclePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLifecyclePolicy(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateLifecyclePolicy struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateLifecyclePolicy) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateLifecyclePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateLifecyclePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateLifecyclePolicyInput ")
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
func addIdempotencyToken_opCreateLifecyclePolicyMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateLifecyclePolicy{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateLifecyclePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLifecyclePolicy",
	}
}
