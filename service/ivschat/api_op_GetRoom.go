// Code generated by smithy-go-codegen DO NOT EDIT.

package ivschat

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/ivschat/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets the specified room.
func (c *Client) GetRoom(ctx context.Context, params *GetRoomInput, optFns ...func(*Options)) (*GetRoomOutput, error) {
	if params == nil {
		params = &GetRoomInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRoom", params, optFns, c.addOperationGetRoomMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRoomOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRoomInput struct {

	// Identifier of the room for which the configuration is to be retrieved.
	// Currently this must be an ARN.
	//
	// This member is required.
	Identifier *string

	noSmithyDocumentSerde
}

type GetRoomOutput struct {

	// Room ARN, from the request (if identifier was an ARN).
	Arn *string

	// Time when the room was created. This is an ISO 8601 timestamp; note that this
	// is returned as a string.
	CreateTime *time.Time

	// Room ID, generated by the system. This is a relative identifier, the part of
	// the ARN that uniquely identifies the room.
	Id *string

	// Array of logging configurations attached to the room.
	LoggingConfigurationIdentifiers []string

	// Maximum number of characters in a single message. Messages are expected to be
	// UTF-8 encoded and this limit applies specifically to rune/code-point count, not
	// number of bytes. Default: 500.
	MaximumMessageLength *int32

	// Maximum number of messages per second that can be sent to the room (by all
	// clients). Default: 10.
	MaximumMessageRatePerSecond *int32

	// Configuration information for optional review of messages.
	MessageReviewHandler *types.MessageReviewHandler

	// Room name. The value does not need to be unique.
	Name *string

	// Tags attached to the resource. Array of maps, each of the form string:string
	// (key:value) .
	Tags map[string]string

	// Time of the room’s last update. This is an ISO 8601 timestamp; note that this
	// is returned as a string.
	UpdateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRoomMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRoom{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRoom{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetRoom"); err != nil {
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
	if err = addOpGetRoomValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRoom(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRoom(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetRoom",
	}
}
