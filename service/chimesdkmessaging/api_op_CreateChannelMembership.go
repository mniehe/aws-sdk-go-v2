// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmessaging

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/chimesdkmessaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a member to a channel. The InvitedBy field in ChannelMembership is derived
// from the request header. A channel member can:
//   - List messages
//   - Send messages
//   - Receive messages
//   - Edit their own messages
//   - Leave the channel
//
// Privacy settings impact this action as follows:
//   - Public Channels: You do not need to be a member to list messages, but you
//     must be a member to send messages.
//   - Private Channels: You must be a member to list or send messages.
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUserArn or AppInstanceBot that makes the API call as the value in
// the header.
func (c *Client) CreateChannelMembership(ctx context.Context, params *CreateChannelMembershipInput, optFns ...func(*Options)) (*CreateChannelMembershipOutput, error) {
	if params == nil {
		params = &CreateChannelMembershipInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateChannelMembership", params, optFns, c.addOperationCreateChannelMembershipMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateChannelMembershipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateChannelMembershipInput struct {

	// The ARN of the channel to which you're adding users.
	//
	// This member is required.
	ChannelArn *string

	// The ARN of the AppInstanceUser or AppInstanceBot that makes the API call.
	//
	// This member is required.
	ChimeBearer *string

	// The AppInstanceUserArn of the member you want to add to the channel.
	//
	// This member is required.
	MemberArn *string

	// The membership type of a user, DEFAULT or HIDDEN . Default members are always
	// returned as part of ListChannelMemberships . Hidden members are only returned if
	// the type filter in ListChannelMemberships equals HIDDEN . Otherwise hidden
	// members are not returned. This is only supported by moderators.
	//
	// This member is required.
	Type types.ChannelMembershipType

	// The ID of the SubChannel in the request. Only required when creating membership
	// in a SubChannel for a moderator in an elastic channel.
	SubChannelId *string

	noSmithyDocumentSerde
}

type CreateChannelMembershipOutput struct {

	// The ARN of the channel.
	ChannelArn *string

	// The ARN and metadata of the member being added.
	Member *types.Identity

	// The ID of the SubChannel in the response.
	SubChannelId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateChannelMembershipMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateChannelMembership{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateChannelMembership{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateChannelMembership"); err != nil {
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
	if err = addOpCreateChannelMembershipValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateChannelMembership(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateChannelMembership(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateChannelMembership",
	}
}
