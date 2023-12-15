// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new pool and associates the specified origination identity to the
// pool. A pool can include one or more phone numbers and SenderIds that are
// associated with your Amazon Web Services account. The new pool inherits its
// configuration from the specified origination identity. This includes keywords,
// message type, opt-out list, two-way configuration, and self-managed opt-out
// configuration. Deletion protection isn't inherited from the origination identity
// and defaults to false. If the origination identity is a phone number and is
// already associated with another pool, an error is returned. A sender ID can be
// associated with multiple pools.
func (c *Client) CreatePool(ctx context.Context, params *CreatePoolInput, optFns ...func(*Options)) (*CreatePoolOutput, error) {
	if params == nil {
		params = &CreatePoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePool", params, optFns, c.addOperationCreatePoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePoolInput struct {

	// The new two-character code, in ISO 3166-1 alpha-2 format, for the country or
	// region of the new pool.
	//
	// This member is required.
	IsoCountryCode *string

	// The type of message. Valid values are TRANSACTIONAL for messages that are
	// critical or time-sensitive and PROMOTIONAL for messages that aren't critical or
	// time-sensitive.
	//
	// This member is required.
	MessageType types.MessageType

	// The origination identity to use such as a PhoneNumberId, PhoneNumberArn,
	// SenderId or SenderIdArn. You can use DescribePhoneNumbers to find the values
	// for PhoneNumberId and PhoneNumberArn while DescribeSenderIds can be used to get
	// the values for SenderId and SenderIdArn.
	//
	// This member is required.
	OriginationIdentity *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. If you don't specify a client token, a randomly generated token is
	// used for the request to ensure idempotency.
	ClientToken *string

	// By default this is set to false. When set to true the pool can't be deleted.
	// You can change this value using the UpdatePool action.
	DeletionProtectionEnabled *bool

	// An array of tags (key and value pairs) associated with the pool.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreatePoolOutput struct {

	// The time when the pool was created, in UNIX epoch time (https://www.epochconverter.com/)
	// format.
	CreatedTimestamp *time.Time

	// When set to true deletion protection is enabled. By default this is set to
	// false.
	DeletionProtectionEnabled bool

	// The type of message for the pool to use.
	MessageType types.MessageType

	// The name of the OptOutList associated with the pool.
	OptOutListName *string

	// The Amazon Resource Name (ARN) for the pool.
	PoolArn *string

	// The unique identifier for the pool.
	PoolId *string

	// By default this is set to false. When an end recipient sends a message that
	// begins with HELP or STOP to one of your dedicated numbers, Amazon Pinpoint
	// automatically replies with a customizable message and adds the end recipient to
	// the OptOutList. When set to true you're responsible for responding to HELP and
	// STOP requests. You're also responsible for tracking and honoring opt-out
	// requests.
	SelfManagedOptOutsEnabled bool

	// Indicates whether shared routes are enabled for the pool.
	SharedRoutesEnabled bool

	// The current status of the pool.
	//   - CREATING: The pool is currently being created and isn't yet available for
	//   use.
	//   - ACTIVE: The pool is active and available for use.
	//   - DELETING: The pool is being deleted.
	Status types.PoolStatus

	// An array of tags (key and value pairs) associated with the pool.
	Tags []types.Tag

	// The Amazon Resource Name (ARN) of the two way channel.
	TwoWayChannelArn *string

	// An optional IAM Role Arn for a service to assume, to be able to post inbound
	// SMS messages.
	TwoWayChannelRole *string

	// By default this is set to false. When set to true you can receive incoming text
	// messages from your end recipients.
	TwoWayEnabled bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreatePool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreatePool{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePool"); err != nil {
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
	if err = addIdempotencyToken_opCreatePoolMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreatePoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePool(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreatePool struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreatePool) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreatePool) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreatePoolInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreatePoolInput ")
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
func addIdempotencyToken_opCreatePoolMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreatePool{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreatePool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePool",
	}
}
