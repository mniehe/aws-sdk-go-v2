// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Notifies the pipeline that the execution of a callback step failed, along with
// a message describing why. When a callback step is run, the pipeline generates a
// callback token and includes the token in a message sent to Amazon Simple Queue
// Service (Amazon SQS).
func (c *Client) SendPipelineExecutionStepFailure(ctx context.Context, params *SendPipelineExecutionStepFailureInput, optFns ...func(*Options)) (*SendPipelineExecutionStepFailureOutput, error) {
	if params == nil {
		params = &SendPipelineExecutionStepFailureInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendPipelineExecutionStepFailure", params, optFns, c.addOperationSendPipelineExecutionStepFailureMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendPipelineExecutionStepFailureOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendPipelineExecutionStepFailureInput struct {

	// The pipeline generated token from the Amazon SQS queue.
	//
	// This member is required.
	CallbackToken *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the operation. An idempotent operation completes no more than one time.
	ClientRequestToken *string

	// A message describing why the step failed.
	FailureReason *string

	noSmithyDocumentSerde
}

type SendPipelineExecutionStepFailureOutput struct {

	// The Amazon Resource Name (ARN) of the pipeline execution.
	PipelineExecutionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendPipelineExecutionStepFailureMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSendPipelineExecutionStepFailure{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSendPipelineExecutionStepFailure{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendPipelineExecutionStepFailure"); err != nil {
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
	if err = addIdempotencyToken_opSendPipelineExecutionStepFailureMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpSendPipelineExecutionStepFailureValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendPipelineExecutionStepFailure(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpSendPipelineExecutionStepFailure struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpSendPipelineExecutionStepFailure) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpSendPipelineExecutionStepFailure) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*SendPipelineExecutionStepFailureInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *SendPipelineExecutionStepFailureInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opSendPipelineExecutionStepFailureMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpSendPipelineExecutionStepFailure{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opSendPipelineExecutionStepFailure(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendPipelineExecutionStepFailure",
	}
}
