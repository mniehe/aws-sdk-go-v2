// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a cross-account attachment in Global Accelerator. You create a
// cross-account attachment to specify the principals who have permission to add to
// accelerators in their own account the resources in your account that you also
// list in the attachment. A principal can be an Amazon Web Services account number
// or the Amazon Resource Name (ARN) for an accelerator. For account numbers that
// are listed as principals, to add a resource listed in the attachment to an
// accelerator, you must sign in to an account specified as a principal. Then you
// can add the resources that are listed to any of your accelerators. If an
// accelerator ARN is listed in the cross-account attachment as a principal, anyone
// with permission to make updates to the accelerator can add as endpoints
// resources that are listed in the attachment.
func (c *Client) CreateCrossAccountAttachment(ctx context.Context, params *CreateCrossAccountAttachmentInput, optFns ...func(*Options)) (*CreateCrossAccountAttachmentOutput, error) {
	if params == nil {
		params = &CreateCrossAccountAttachmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCrossAccountAttachment", params, optFns, c.addOperationCreateCrossAccountAttachmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCrossAccountAttachmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCrossAccountAttachmentInput struct {

	// A unique, case-sensitive identifier that you provide to ensure the
	// idempotency—that is, the uniqueness—of the request.
	//
	// This member is required.
	IdempotencyToken *string

	// The name of the cross-account attachment.
	//
	// This member is required.
	Name *string

	// The principals to list in the cross-account attachment. A principal can be an
	// Amazon Web Services account number or the Amazon Resource Name (ARN) for an
	// accelerator.
	Principals []string

	// The Amazon Resource Names (ARNs) for the resources to list in the cross-account
	// attachment. A resource can be any supported Amazon Web Services resource type
	// for Global Accelerator.
	Resources []types.Resource

	// Create tags for cross-account attachment. For more information, see Tagging in
	// Global Accelerator (https://docs.aws.amazon.com/global-accelerator/latest/dg/tagging-in-global-accelerator.html)
	// in the Global Accelerator Developer Guide.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateCrossAccountAttachmentOutput struct {

	// Information about the cross-account attachment.
	CrossAccountAttachment *types.Attachment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCrossAccountAttachmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCrossAccountAttachment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCrossAccountAttachment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCrossAccountAttachment"); err != nil {
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
	if err = addIdempotencyToken_opCreateCrossAccountAttachmentMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateCrossAccountAttachmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCrossAccountAttachment(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateCrossAccountAttachment struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateCrossAccountAttachment) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateCrossAccountAttachment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateCrossAccountAttachmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateCrossAccountAttachmentInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateCrossAccountAttachmentMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateCrossAccountAttachment{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateCrossAccountAttachment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCrossAccountAttachment",
	}
}
