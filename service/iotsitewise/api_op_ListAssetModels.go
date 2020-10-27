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

// Retrieves a paginated list of summaries of all asset models.
func (c *Client) ListAssetModels(ctx context.Context, params *ListAssetModelsInput, optFns ...func(*Options)) (*ListAssetModelsOutput, error) {
	if params == nil {
		params = &ListAssetModelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssetModels", params, optFns, addOperationListAssetModelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssetModelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssetModelsInput struct {

	// The maximum number of results to be returned per paginated request. Default: 50
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string
}

type ListAssetModelsOutput struct {

	// A list that summarizes each asset model.
	//
	// This member is required.
	AssetModelSummaries []*types.AssetModelSummary

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAssetModelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAssetModels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssetModels{}, middleware.After)
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
	addEndpointPrefix_opListAssetModelsMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAssetModels(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type endpointPrefix_opListAssetModelsMiddleware struct {
}

func (*endpointPrefix_opListAssetModelsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListAssetModelsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
func addEndpointPrefix_opListAssetModelsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListAssetModelsMiddleware{}, `OperationSerializer`, middleware.Before)
}

func newServiceMetadataMiddleware_opListAssetModels(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "ListAssetModels",
	}
}
