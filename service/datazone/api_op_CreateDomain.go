// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon DataZone domain.
func (c *Client) CreateDomain(ctx context.Context, params *CreateDomainInput, optFns ...func(*Options)) (*CreateDomainOutput, error) {
	if params == nil {
		params = &CreateDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDomain", params, optFns, c.addOperationCreateDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDomainInput struct {

	// The domain execution role that is created when an Amazon DataZone domain is
	// created. The domain execution role is created in the Amazon Web Services account
	// that houses the Amazon DataZone domain.
	//
	// This member is required.
	DomainExecutionRole *string

	// The name of the Amazon DataZone domain.
	//
	// This member is required.
	Name *string

	// A unique, case-sensitive identifier that is provided to ensure the idempotency
	// of the request.
	ClientToken *string

	// The description of the Amazon DataZone domain.
	Description *string

	// The identifier of the Amazon Web Services Key Management Service (KMS) key that
	// is used to encrypt the Amazon DataZone domain, metadata, and reporting data.
	KmsKeyIdentifier *string

	// The single-sign on configuration of the Amazon DataZone domain.
	SingleSignOn *types.SingleSignOn

	// The tags specified for the Amazon DataZone domain.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateDomainOutput struct {

	// The identifier of the Amazon DataZone domain.
	//
	// This member is required.
	Id *string

	// The ARN of the Amazon DataZone domain.
	Arn *string

	// The description of the Amazon DataZone domain.
	Description *string

	// The domain execution role that is created when an Amazon DataZone domain is
	// created. The domain execution role is created in the Amazon Web Services account
	// that houses the Amazon DataZone domain.
	DomainExecutionRole *string

	// The identifier of the Amazon Web Services Key Management Service (KMS) key that
	// is used to encrypt the Amazon DataZone domain, metadata, and reporting data.
	KmsKeyIdentifier *string

	// The name of the Amazon DataZone domain.
	Name *string

	// The URL of the data portal for this Amazon DataZone domain.
	PortalUrl *string

	// The single-sign on configuration of the Amazon DataZone domain.
	SingleSignOn *types.SingleSignOn

	// The status of the Amazon DataZone domain.
	Status types.DomainStatus

	// The tags specified for the Amazon DataZone domain.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDomain{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDomain"); err != nil {
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
	if err = addIdempotencyToken_opCreateDomainMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDomain(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateDomain struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDomain) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDomain) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDomainInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDomainInput ")
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
func addIdempotencyToken_opCreateDomainMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateDomain{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDomain",
	}
}
