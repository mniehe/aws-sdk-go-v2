// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops a running multiplex. If the multiplex isn't running, this action has no
// effect.
func (c *Client) StopMultiplex(ctx context.Context, params *StopMultiplexInput, optFns ...func(*Options)) (*StopMultiplexOutput, error) {
	if params == nil {
		params = &StopMultiplexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopMultiplex", params, optFns, c.addOperationStopMultiplexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopMultiplexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for StopMultiplexRequest
type StopMultiplexInput struct {

	// The ID of the multiplex.
	//
	// This member is required.
	MultiplexId *string

	noSmithyDocumentSerde
}

// Placeholder documentation for StopMultiplexResponse
type StopMultiplexOutput struct {

	// The unique arn of the multiplex.
	Arn *string

	// A list of availability zones for the multiplex.
	AvailabilityZones []string

	// A list of the multiplex output destinations.
	Destinations []types.MultiplexOutputDestination

	// The unique id of the multiplex.
	Id *string

	// Configuration for a multiplex event.
	MultiplexSettings *types.MultiplexSettings

	// The name of the multiplex.
	Name *string

	// The number of currently healthy pipelines.
	PipelinesRunningCount *int32

	// The number of programs in the multiplex.
	ProgramCount *int32

	// The current state of the multiplex.
	State types.MultiplexState

	// A collection of key-value pairs.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopMultiplexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStopMultiplex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStopMultiplex{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StopMultiplex"); err != nil {
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
	if err = addOpStopMultiplexValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopMultiplex(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopMultiplex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StopMultiplex",
	}
}
