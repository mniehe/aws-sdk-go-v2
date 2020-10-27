// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a paginated list of assets associated with an AWS IoT SiteWise Monitor
// project.
func (c *Client) ListProjectAssets(ctx context.Context, params *ListProjectAssetsInput, optFns ...func(*Options)) (*ListProjectAssetsOutput, error) {
	if params == nil {
		params = &ListProjectAssetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProjectAssets", params, optFns, addOperationListProjectAssetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProjectAssetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProjectAssetsInput struct {

	// The ID of the project.
	//
	// This member is required.
	ProjectId *string

	// The maximum number of results to be returned per paginated request. Default: 50
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string
}

type ListProjectAssetsOutput struct {

	// A list that contains the IDs of each asset associated with the project.
	//
	// This member is required.
	AssetIds []*string

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListProjectAssetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProjectAssets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProjectAssets{}, middleware.After)
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
	addEndpointPrefix_opListProjectAssetsMiddleware(stack)
	addOpListProjectAssetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListProjectAssets(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type endpointPrefix_opListProjectAssetsMiddleware struct {
}

func (*endpointPrefix_opListProjectAssetsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListProjectAssetsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
func addEndpointPrefix_opListProjectAssetsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListProjectAssetsMiddleware{}, `OperationSerializer`, middleware.Before)
}

func newServiceMetadataMiddleware_opListProjectAssets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "ListProjectAssets",
	}
}
