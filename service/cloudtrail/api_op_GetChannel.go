// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about a specific channel.
func (c *Client) GetChannel(ctx context.Context, params *GetChannelInput, optFns ...func(*Options)) (*GetChannelOutput, error) {
	if params == nil {
		params = &GetChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetChannel", params, optFns, c.addOperationGetChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetChannelInput struct {

	// The ARN or UUID of a channel.
	//
	// This member is required.
	Channel *string

	noSmithyDocumentSerde
}

type GetChannelOutput struct {

	// The ARN of an channel returned by a GetChannel request.
	ChannelArn *string

	// The destinations for the channel. For channels created for integrations, the
	// destinations are the event data stores that log events arriving through the
	// channel. For service-linked channels, the destination is the Amazon Web Services
	// service that created the service-linked channel to receive events.
	Destinations []types.Destination

	// A table showing information about the most recent successful and failed
	// attempts to ingest events.
	IngestionStatus *types.IngestionStatus

	// The name of the CloudTrail channel. For service-linked channels, the name is
	// aws-service-channel/service-name/custom-suffix where service-name represents
	// the name of the Amazon Web Services service that created the channel and
	// custom-suffix represents the suffix generated by the Amazon Web Services service.
	Name *string

	// The source for the CloudTrail channel.
	Source *string

	// Provides information about the advanced event selectors configured for the
	// channel, and whether the channel applies to all Regions or a single Region.
	SourceConfig *types.SourceConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetChannel{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetChannel"); err != nil {
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
	if err = addOpGetChannelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetChannel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetChannel",
	}
}
