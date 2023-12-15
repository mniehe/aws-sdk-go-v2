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

// Retrieves the list of instances (also called nodes interchangeably) in a
// SageMaker HyperPod cluster.
func (c *Client) ListClusterNodes(ctx context.Context, params *ListClusterNodesInput, optFns ...func(*Options)) (*ListClusterNodesOutput, error) {
	if params == nil {
		params = &ListClusterNodesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListClusterNodes", params, optFns, c.addOperationListClusterNodesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListClusterNodesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListClusterNodesInput struct {

	// The string name or the Amazon Resource Name (ARN) of the SageMaker HyperPod
	// cluster in which you want to retrieve the list of nodes.
	//
	// This member is required.
	ClusterName *string

	// A filter that returns nodes in a SageMaker HyperPod cluster created after the
	// specified time. Timestamps are formatted according to the ISO 8601 standard.
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

	// A filter that returns nodes in a SageMaker HyperPod cluster created before the
	// specified time. The acceptable formats are the same as the timestamp formats for
	// CreationTimeAfter . For more information about the timestamp format, see
	// Timestamp (https://docs.aws.amazon.com/cli/latest/userguide/cli-usage-parameters-types.html#parameter-type-timestamp)
	// in the Amazon Web Services Command Line Interface User Guide.
	CreationTimeBefore *time.Time

	// A filter that returns the instance groups whose name contain a specified string.
	InstanceGroupNameContains *string

	// The maximum number of nodes to return in the response.
	MaxResults *int32

	// If the result of the previous ListClusterNodes request was truncated, the
	// response includes a NextToken . To retrieve the next set of cluster nodes, use
	// the token in the next request.
	NextToken *string

	// The field by which to sort results. The default value is CREATION_TIME .
	SortBy types.ClusterSortBy

	// The sort order for results. The default value is Ascending .
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListClusterNodesOutput struct {

	// The summaries of listed instances in a SageMaker HyperPod cluster
	//
	// This member is required.
	ClusterNodeSummaries []types.ClusterNodeSummary

	// The next token specified for listing instances in a SageMaker HyperPod cluster.
	//
	// This member is required.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListClusterNodesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListClusterNodes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListClusterNodes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListClusterNodes"); err != nil {
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
	if err = addOpListClusterNodesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListClusterNodes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListClusterNodes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListClusterNodes",
	}
}
