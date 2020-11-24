// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Shows usage limits on a cluster. Results are filtered based on the combination
// of input usage limit identifier, cluster identifier, and feature type
// parameters:
//
// * If usage limit identifier, cluster identifier, and feature type
// are not provided, then all usage limit objects for the current account in the
// current region are returned.
//
// * If usage limit identifier is provided, then the
// corresponding usage limit object is returned.
//
// * If cluster identifier is
// provided, then all usage limit objects for the specified cluster are
// returned.
//
// * If cluster identifier and feature type are provided, then all usage
// limit objects for the combination of cluster and feature are returned.
func (c *Client) DescribeUsageLimits(ctx context.Context, params *DescribeUsageLimitsInput, optFns ...func(*Options)) (*DescribeUsageLimitsOutput, error) {
	if params == nil {
		params = &DescribeUsageLimitsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeUsageLimits", params, optFns, addOperationDescribeUsageLimitsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeUsageLimitsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUsageLimitsInput struct {

	// The identifier of the cluster for which you want to describe usage limits.
	ClusterIdentifier *string

	// The feature type for which you want to describe usage limits.
	FeatureType types.UsageLimitFeatureType

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeUsageLimits request exceed the
	// value specified in MaxRecords, AWS returns a value in the Marker field of the
	// response. You can retrieve the next set of response records by providing the
	// returned marker value in the Marker parameter and retrying the request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32

	// A tag key or keys for which you want to return all matching usage limit objects
	// that are associated with the specified key or keys. For example, suppose that
	// you have parameter groups that are tagged with keys called owner and
	// environment. If you specify both of these tag keys in the request, Amazon
	// Redshift returns a response with the usage limit objects have either or both of
	// these tag keys associated with them.
	TagKeys []string

	// A tag value or values for which you want to return all matching usage limit
	// objects that are associated with the specified tag value or values. For example,
	// suppose that you have parameter groups that are tagged with values called admin
	// and test. If you specify both of these tag values in the request, Amazon
	// Redshift returns a response with the usage limit objects that have either or
	// both of these tag values associated with them.
	TagValues []string

	// The identifier of the usage limit to describe.
	UsageLimitId *string
}

type DescribeUsageLimitsOutput struct {

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// Contains the output from the DescribeUsageLimits action.
	UsageLimits []types.UsageLimit

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeUsageLimitsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeUsageLimits{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeUsageLimits{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUsageLimits(options.Region), middleware.Before); err != nil {
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
	return nil
}

// DescribeUsageLimitsAPIClient is a client that implements the DescribeUsageLimits
// operation.
type DescribeUsageLimitsAPIClient interface {
	DescribeUsageLimits(context.Context, *DescribeUsageLimitsInput, ...func(*Options)) (*DescribeUsageLimitsOutput, error)
}

var _ DescribeUsageLimitsAPIClient = (*Client)(nil)

// DescribeUsageLimitsPaginatorOptions is the paginator options for
// DescribeUsageLimits
type DescribeUsageLimitsPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeUsageLimitsPaginator is a paginator for DescribeUsageLimits
type DescribeUsageLimitsPaginator struct {
	options   DescribeUsageLimitsPaginatorOptions
	client    DescribeUsageLimitsAPIClient
	params    *DescribeUsageLimitsInput
	nextToken *string
	firstPage bool
}

// NewDescribeUsageLimitsPaginator returns a new DescribeUsageLimitsPaginator
func NewDescribeUsageLimitsPaginator(client DescribeUsageLimitsAPIClient, params *DescribeUsageLimitsInput, optFns ...func(*DescribeUsageLimitsPaginatorOptions)) *DescribeUsageLimitsPaginator {
	options := DescribeUsageLimitsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeUsageLimitsInput{}
	}

	return &DescribeUsageLimitsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeUsageLimitsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeUsageLimits page.
func (p *DescribeUsageLimitsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeUsageLimitsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeUsageLimits(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeUsageLimits(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeUsageLimits",
	}
}
