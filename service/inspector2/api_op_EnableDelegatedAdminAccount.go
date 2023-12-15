// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables the Amazon Inspector delegated administrator for your Organizations
// organization.
func (c *Client) EnableDelegatedAdminAccount(ctx context.Context, params *EnableDelegatedAdminAccountInput, optFns ...func(*Options)) (*EnableDelegatedAdminAccountOutput, error) {
	if params == nil {
		params = &EnableDelegatedAdminAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableDelegatedAdminAccount", params, optFns, c.addOperationEnableDelegatedAdminAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableDelegatedAdminAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableDelegatedAdminAccountInput struct {

	// The Amazon Web Services account ID of the Amazon Inspector delegated
	// administrator.
	//
	// This member is required.
	DelegatedAdminAccountId *string

	// The idempotency token for the request.
	ClientToken *string

	noSmithyDocumentSerde
}

type EnableDelegatedAdminAccountOutput struct {

	// The Amazon Web Services account ID of the successfully Amazon Inspector
	// delegated administrator.
	//
	// This member is required.
	DelegatedAdminAccountId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableDelegatedAdminAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpEnableDelegatedAdminAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpEnableDelegatedAdminAccount{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "EnableDelegatedAdminAccount"); err != nil {
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
	if err = addIdempotencyToken_opEnableDelegatedAdminAccountMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpEnableDelegatedAdminAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableDelegatedAdminAccount(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpEnableDelegatedAdminAccount struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpEnableDelegatedAdminAccount) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpEnableDelegatedAdminAccount) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*EnableDelegatedAdminAccountInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *EnableDelegatedAdminAccountInput ")
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
func addIdempotencyToken_opEnableDelegatedAdminAccountMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpEnableDelegatedAdminAccount{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opEnableDelegatedAdminAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "EnableDelegatedAdminAccount",
	}
}
