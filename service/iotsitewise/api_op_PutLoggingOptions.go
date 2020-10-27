// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets logging options for AWS IoT SiteWise.
func (c *Client) PutLoggingOptions(ctx context.Context, params *PutLoggingOptionsInput, optFns ...func(*Options)) (*PutLoggingOptionsOutput, error) {
	if params == nil {
		params = &PutLoggingOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutLoggingOptions", params, optFns, addOperationPutLoggingOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutLoggingOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLoggingOptionsInput struct {

	// The logging options to set.
	//
	// This member is required.
	LoggingOptions *types.LoggingOptions
}

type PutLoggingOptionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutLoggingOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutLoggingOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutLoggingOptions{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addEndpointPrefix_opPutLoggingOptionsMiddleware(stack)
	addOpPutLoggingOptionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutLoggingOptions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type endpointPrefix_opPutLoggingOptionsMiddleware struct {
}

func (*endpointPrefix_opPutLoggingOptionsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opPutLoggingOptionsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.HostPrefix = "model."

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opPutLoggingOptionsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opPutLoggingOptionsMiddleware{}, `OperationSerializer`, middleware.Before)
}

func newServiceMetadataMiddleware_opPutLoggingOptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "PutLoggingOptions",
	}
}
