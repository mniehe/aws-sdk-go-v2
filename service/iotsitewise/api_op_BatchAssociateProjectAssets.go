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

// Associates a group (batch) of assets with an AWS IoT SiteWise Monitor project.
func (c *Client) BatchAssociateProjectAssets(ctx context.Context, params *BatchAssociateProjectAssetsInput, optFns ...func(*Options)) (*BatchAssociateProjectAssetsOutput, error) {
	if params == nil {
		params = &BatchAssociateProjectAssetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchAssociateProjectAssets", params, optFns, addOperationBatchAssociateProjectAssetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchAssociateProjectAssetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchAssociateProjectAssetsInput struct {

	// The IDs of the assets to be associated to the project.
	//
	// This member is required.
	AssetIds []*string

	// The ID of the project to which to associate the assets.
	//
	// This member is required.
	ProjectId *string

	// A unique case-sensitive identifier that you can provide to ensure the
	// idempotency of the request. Don't reuse this client token if a new idempotent
	// request is required.
	ClientToken *string
}

type BatchAssociateProjectAssetsOutput struct {

	// A list of associated error information, if any.
	Errors []*types.AssetErrorDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchAssociateProjectAssetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchAssociateProjectAssets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchAssociateProjectAssets{}, middleware.After)
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
	addEndpointPrefix_opBatchAssociateProjectAssetsMiddleware(stack)
	addIdempotencyToken_opBatchAssociateProjectAssetsMiddleware(stack, options)
	addOpBatchAssociateProjectAssetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchAssociateProjectAssets(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type endpointPrefix_opBatchAssociateProjectAssetsMiddleware struct {
}

func (*endpointPrefix_opBatchAssociateProjectAssetsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opBatchAssociateProjectAssetsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.HostPrefix = "monitor."

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opBatchAssociateProjectAssetsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opBatchAssociateProjectAssetsMiddleware{}, `OperationSerializer`, middleware.Before)
}

type idempotencyToken_initializeOpBatchAssociateProjectAssets struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpBatchAssociateProjectAssets) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpBatchAssociateProjectAssets) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*BatchAssociateProjectAssetsInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *BatchAssociateProjectAssetsInput ")
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
func addIdempotencyToken_opBatchAssociateProjectAssetsMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpBatchAssociateProjectAssets{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opBatchAssociateProjectAssets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "BatchAssociateProjectAssets",
	}
}
