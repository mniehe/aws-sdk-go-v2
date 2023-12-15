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

// Update an endpoint group. A resource must be valid and active when you add it
// as an endpoint.
func (c *Client) UpdateEndpointGroup(ctx context.Context, params *UpdateEndpointGroupInput, optFns ...func(*Options)) (*UpdateEndpointGroupOutput, error) {
	if params == nil {
		params = &UpdateEndpointGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEndpointGroup", params, optFns, c.addOperationUpdateEndpointGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEndpointGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEndpointGroupInput struct {

	// The Amazon Resource Name (ARN) of the endpoint group.
	//
	// This member is required.
	EndpointGroupArn *string

	// The list of endpoint objects. A resource must be valid and active when you add
	// it as an endpoint.
	EndpointConfigurations []types.EndpointConfiguration

	// The time—10 seconds or 30 seconds—between each health check for an endpoint.
	// The default value is 30.
	HealthCheckIntervalSeconds *int32

	// If the protocol is HTTP/S, then this specifies the path that is the destination
	// for health check targets. The default value is slash (/).
	HealthCheckPath *string

	// The port that Global Accelerator uses to check the health of endpoints that are
	// part of this endpoint group. The default port is the listener port that this
	// endpoint group is associated with. If the listener port is a list of ports,
	// Global Accelerator uses the first port in the list.
	HealthCheckPort *int32

	// The protocol that Global Accelerator uses to check the health of endpoints that
	// are part of this endpoint group. The default value is TCP.
	HealthCheckProtocol types.HealthCheckProtocol

	// Override specific listener ports used to route traffic to endpoints that are
	// part of this endpoint group. For example, you can create a port override in
	// which the listener receives user traffic on ports 80 and 443, but your
	// accelerator routes that traffic to ports 1080 and 1443, respectively, on the
	// endpoints. For more information, see Overriding listener ports (https://docs.aws.amazon.com/global-accelerator/latest/dg/about-endpoint-groups-port-override.html)
	// in the Global Accelerator Developer Guide.
	PortOverrides []types.PortOverride

	// The number of consecutive health checks required to set the state of a healthy
	// endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default
	// value is 3.
	ThresholdCount *int32

	// The percentage of traffic to send to an Amazon Web Services Region. Additional
	// traffic is distributed to other endpoint groups for this listener. Use this
	// action to increase (dial up) or decrease (dial down) traffic to a specific
	// Region. The percentage is applied to the traffic that would otherwise have been
	// routed to the Region based on optimal routing. The default value is 100.
	TrafficDialPercentage *float32

	noSmithyDocumentSerde
}

type UpdateEndpointGroupOutput struct {

	// The information about the endpoint group that was updated.
	EndpointGroup *types.EndpointGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEndpointGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateEndpointGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateEndpointGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateEndpointGroup"); err != nil {
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
	if err = addOpUpdateEndpointGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEndpointGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEndpointGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateEndpointGroup",
	}
}
