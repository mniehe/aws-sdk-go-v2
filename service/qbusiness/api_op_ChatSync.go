// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts or continues a non-streaming Amazon Q conversation.
func (c *Client) ChatSync(ctx context.Context, params *ChatSyncInput, optFns ...func(*Options)) (*ChatSyncOutput, error) {
	if params == nil {
		params = &ChatSyncInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ChatSync", params, optFns, c.addOperationChatSyncMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ChatSyncOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ChatSyncInput struct {

	// The identifier of the Amazon Q application linked to the Amazon Q conversation.
	//
	// This member is required.
	ApplicationId *string

	// The identifier of the user attached to the chat input.
	//
	// This member is required.
	UserId *string

	// A request from an end user to perform an Amazon Q plugin action.
	ActionExecution *types.ActionExecution

	// A list of files uploaded directly during chat. You can upload a maximum of 5
	// files of upto 10 MB each.
	Attachments []types.AttachmentInput

	// Enables filtering of Amazon Q web experience responses based on document
	// attributes or metadata fields.
	AttributeFilter *types.AttributeFilter

	// A token that you provide to identify a chat request.
	ClientToken *string

	// The identifier of the Amazon Q conversation.
	ConversationId *string

	// The identifier of the previous end user text input message in a conversation.
	ParentMessageId *string

	// The groups that a user associated with the chat input belongs to.
	UserGroups []string

	// A end user message in a conversation.
	UserMessage *string

	noSmithyDocumentSerde
}

type ChatSyncOutput struct {

	// A request from Amazon Q to the end user for information Amazon Q needs to
	// successfully complete a requested plugin action.
	ActionReview *types.ActionReview

	// The identifier of the Amazon Q conversation.
	ConversationId *string

	// A list of files which failed to upload during chat.
	FailedAttachments []types.AttachmentOutput

	// The source documents used to generate the conversation response.
	SourceAttributions []*types.SourceAttribution

	// An AI-generated message in a conversation.
	SystemMessage *string

	// The identifier of an Amazon Q AI generated message within the conversation.
	SystemMessageId *string

	// The identifier of an Amazon Q end user text input message within the
	// conversation.
	UserMessageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationChatSyncMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpChatSync{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpChatSync{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ChatSync"); err != nil {
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
	if err = addIdempotencyToken_opChatSyncMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpChatSyncValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opChatSync(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpChatSync struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpChatSync) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpChatSync) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ChatSyncInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ChatSyncInput ")
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
func addIdempotencyToken_opChatSyncMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpChatSync{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opChatSync(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ChatSync",
	}
}
