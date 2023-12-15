// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the list of SageMaker HyperPod clusters.
func (c *Client) ListClusters(ctx context.Context, params *ListClustersInput, optFns ...func(*Options)) (*ListClustersOutput, error) {
	if params == nil {
		params = &ListClustersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListClusters", params, optFns, c.addOperationListClustersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListClustersInput struct {

	// Set a start time for the time range during which you want to list SageMaker
	// HyperPod clusters. Timestamps are formatted according to the ISO 8601 standard.
	// Acceptable formats include:
	//   - YYYY-MM-DDThh:mm:ss.sssTZD (UTC), for example, 2014-10-01T20:30:00.000Z
	//   - YYYY-MM-DDThh:mm:ss.sssTZD (with offset), for example,
	//   2014-10-01T12:30:00.000-08:00
	//   - YYYY-MM-DD , for example, 2014-10-01
	//   - Unix time in seconds, for example, 1412195400 . This is also referred to as
	//   Unix Epoch time and represents the number of seconds since midnight, January 1,
	//   1970 UTC.
	// For more information about the timestamp format, see Timestamp (https://docs.aws.amazon.com/cli/latest/userguide/cli-usage-parameters-types.html#parameter-type-timestamp)
	// in the Amazon Web Services Command Line Interface User Guide.
	CreationTimeAfter *time.Time

	// Set an end time for the time range during which you want to list SageMaker
	// HyperPod clusters. A filter that returns nodes in a SageMaker HyperPod cluster
	// created before the specified time. The acceptable formats are the same as the
	// timestamp formats for CreationTimeAfter . For more information about the
	// timestamp format, see Timestamp (https://docs.aws.amazon.com/cli/latest/userguide/cli-usage-parameters-types.html#parameter-type-timestamp)
	// in the Amazon Web Services Command Line Interface User Guide.
	CreationTimeBefore *time.Time

	// Set the maximum number of SageMaker HyperPod clusters to list.
	MaxResults *int32

	// Set the maximum number of instances to print in the list.
	NameContains *string

	// Set the next token to retrieve the list of SageMaker HyperPod clusters.
	NextToken *string

	// The field by which to sort results. The default value is CREATION_TIME .
	SortBy types.ClusterSortBy

	// The sort order for results. The default value is Ascending .
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListClustersOutput struct {

	// The summaries of listed SageMaker HyperPod clusters.
	//
	// This member is required.
	ClusterSummaries []types.ClusterSummary

	// If the result of the previous ListClusters request was truncated, the response
	// includes a NextToken . To retrieve the next set of clusters, use the token in
	// the next request.
	//
	// This member is required.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListClustersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListClusters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListClusters{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListClusters"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListClusters(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListClusters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListClusters",
	}
}
