// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a Device Defender audit suppression. Requires permission to access the
// CreateAuditSuppression (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) CreateAuditSuppression(ctx context.Context, params *CreateAuditSuppressionInput, optFns ...func(*Options)) (*CreateAuditSuppressionOutput, error) {
	if params == nil {
		params = &CreateAuditSuppressionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAuditSuppression", params, optFns, c.addOperationCreateAuditSuppressionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAuditSuppressionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAuditSuppressionInput struct {

	// An audit check name. Checks must be enabled for your account. (Use
	// DescribeAccountAuditConfiguration to see the list of all checks, including those
	// that are enabled or use UpdateAccountAuditConfiguration to select which checks
	// are enabled.)
	//
	// This member is required.
	CheckName *string

	// Each audit supression must have a unique client request token. If you try to
	// create a new audit suppression with the same token as one that already exists,
	// an exception occurs. If you omit this value, Amazon Web Services SDKs will
	// automatically generate a unique client request.
	//
	// This member is required.
	ClientRequestToken *string

	// Information that identifies the noncompliant resource.
	//
	// This member is required.
	ResourceIdentifier *types.ResourceIdentifier

	// The description of the audit suppression.
	Description *string

	// The epoch timestamp in seconds at which this suppression expires.
	ExpirationDate *time.Time

	// Indicates whether a suppression should exist indefinitely or not.
	SuppressIndefinitely *bool

	noSmithyDocumentSerde
}

type CreateAuditSuppressionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAuditSuppressionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAuditSuppression{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAuditSuppression{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAuditSuppression"); err != nil {
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
	if err = addIdempotencyToken_opCreateAuditSuppressionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAuditSuppressionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAuditSuppression(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateAuditSuppression struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAuditSuppression) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAuditSuppression) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAuditSuppressionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAuditSuppressionInput ")
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
func addIdempotencyToken_opCreateAuditSuppressionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateAuditSuppression{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAuditSuppression(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAuditSuppression",
	}
}
