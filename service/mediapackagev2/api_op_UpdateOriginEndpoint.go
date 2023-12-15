// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackagev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/mediapackagev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Update the specified origin endpoint. Edit the packaging preferences on an
// endpoint to optimize the viewing experience. You can't edit the name of the
// endpoint. Any edits you make that impact the video output may not be reflected
// for a few minutes.
func (c *Client) UpdateOriginEndpoint(ctx context.Context, params *UpdateOriginEndpointInput, optFns ...func(*Options)) (*UpdateOriginEndpointOutput, error) {
	if params == nil {
		params = &UpdateOriginEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateOriginEndpoint", params, optFns, c.addOperationUpdateOriginEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateOriginEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateOriginEndpointInput struct {

	// The name that describes the channel group. The name is the primary identifier
	// for the channel group, and must be unique for your account in the AWS Region.
	//
	// This member is required.
	ChannelGroupName *string

	// The name that describes the channel. The name is the primary identifier for the
	// channel, and must be unique for your account in the AWS Region and channel
	// group.
	//
	// This member is required.
	ChannelName *string

	// The type of container attached to this origin endpoint. A container type is a
	// file format that encapsulates one or more media streams, such as audio and
	// video, into a single file.
	//
	// This member is required.
	ContainerType types.ContainerType

	// The name that describes the origin endpoint. The name is the primary identifier
	// for the origin endpoint, and and must be unique for your account in the AWS
	// Region and channel.
	//
	// This member is required.
	OriginEndpointName *string

	// Any descriptive information that you want to add to the origin endpoint for
	// future identification purposes.
	Description *string

	// An HTTP live streaming (HLS) manifest configuration.
	HlsManifests []types.CreateHlsManifestConfiguration

	// A low-latency HLS manifest configuration.
	LowLatencyHlsManifests []types.CreateLowLatencyHlsManifestConfiguration

	// The segment configuration, including the segment name, duration, and other
	// configuration values.
	Segment *types.Segment

	// The size of the window (in seconds) to create a window of the live stream
	// that's available for on-demand viewing. Viewers can start-over or catch-up on
	// content that falls within the window. The maximum startover window is 1,209,600
	// seconds (14 days).
	StartoverWindowSeconds *int32

	noSmithyDocumentSerde
}

type UpdateOriginEndpointOutput struct {

	// The ARN associated with the resource.
	//
	// This member is required.
	Arn *string

	// The name that describes the channel group. The name is the primary identifier
	// for the channel group, and must be unique for your account in the AWS Region.
	//
	// This member is required.
	ChannelGroupName *string

	// The name that describes the channel. The name is the primary identifier for the
	// channel, and must be unique for your account in the AWS Region and channel
	// group.
	//
	// This member is required.
	ChannelName *string

	// The type of container attached to this origin endpoint.
	//
	// This member is required.
	ContainerType types.ContainerType

	// The date and time the origin endpoint was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The date and time the origin endpoint was modified.
	//
	// This member is required.
	ModifiedAt *time.Time

	// The name that describes the origin endpoint. The name is the primary identifier
	// for the origin endpoint, and and must be unique for your account in the AWS
	// Region and channel.
	//
	// This member is required.
	OriginEndpointName *string

	// The segment configuration, including the segment name, duration, and other
	// configuration values.
	//
	// This member is required.
	Segment *types.Segment

	// The description of the origin endpoint.
	Description *string

	// An HTTP live streaming (HLS) manifest configuration.
	HlsManifests []types.GetHlsManifestConfiguration

	// A low-latency HLS manifest configuration.
	LowLatencyHlsManifests []types.GetLowLatencyHlsManifestConfiguration

	// The size of the window (in seconds) to create a window of the live stream
	// that's available for on-demand viewing. Viewers can start-over or catch-up on
	// content that falls within the window.
	StartoverWindowSeconds *int32

	// The comma-separated list of tag key:value pairs assigned to the origin endpoint.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateOriginEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateOriginEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateOriginEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateOriginEndpoint"); err != nil {
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
	if err = addOpUpdateOriginEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateOriginEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateOriginEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateOriginEndpoint",
	}
}
