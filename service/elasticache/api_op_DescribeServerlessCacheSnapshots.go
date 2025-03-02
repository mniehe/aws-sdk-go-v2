// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about serverless cache snapshots. By default, this API
// lists all of the customer’s serverless cache snapshots. It can also describe a
// single serverless cache snapshot, or the snapshots associated with a particular
// serverless cache. Available for Redis only.
func (c *Client) DescribeServerlessCacheSnapshots(ctx context.Context, params *DescribeServerlessCacheSnapshotsInput, optFns ...func(*Options)) (*DescribeServerlessCacheSnapshotsOutput, error) {
	if params == nil {
		params = &DescribeServerlessCacheSnapshotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeServerlessCacheSnapshots", params, optFns, c.addOperationDescribeServerlessCacheSnapshotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeServerlessCacheSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeServerlessCacheSnapshotsInput struct {

	// The maximum number of records to include in the response. If more records exist
	// than the specified max-results value, a market is included in the response so
	// that remaining results can be retrieved. Available for Redis only.The default is
	// 50. The Validation Constraints are a maximum of 50.
	MaxResults *int32

	// An optional marker returned from a prior request to support pagination of
	// results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// max-results. Available for Redis only.
	NextToken *string

	// The identifier of serverless cache. If this parameter is specified, only
	// snapshots associated with that specific serverless cache are described.
	// Available for Redis only.
	ServerlessCacheName *string

	// The identifier of the serverless cache’s snapshot. If this parameter is
	// specified, only this snapshot is described. Available for Redis only.
	ServerlessCacheSnapshotName *string

	// The type of snapshot that is being described. Available for Redis only.
	SnapshotType *string

	noSmithyDocumentSerde
}

type DescribeServerlessCacheSnapshotsOutput struct {

	// An optional marker returned from a prior request to support pagination of
	// results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// max-results. Available for Redis only.
	NextToken *string

	// The serverless caches snapshots associated with a given description request.
	// Available for Redis only.
	ServerlessCacheSnapshots []types.ServerlessCacheSnapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeServerlessCacheSnapshotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeServerlessCacheSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeServerlessCacheSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeServerlessCacheSnapshots"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeServerlessCacheSnapshots(options.Region), middleware.Before); err != nil {
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

// DescribeServerlessCacheSnapshotsAPIClient is a client that implements the
// DescribeServerlessCacheSnapshots operation.
type DescribeServerlessCacheSnapshotsAPIClient interface {
	DescribeServerlessCacheSnapshots(context.Context, *DescribeServerlessCacheSnapshotsInput, ...func(*Options)) (*DescribeServerlessCacheSnapshotsOutput, error)
}

var _ DescribeServerlessCacheSnapshotsAPIClient = (*Client)(nil)

// DescribeServerlessCacheSnapshotsPaginatorOptions is the paginator options for
// DescribeServerlessCacheSnapshots
type DescribeServerlessCacheSnapshotsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified max-results value, a market is included in the response so
	// that remaining results can be retrieved. Available for Redis only.The default is
	// 50. The Validation Constraints are a maximum of 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeServerlessCacheSnapshotsPaginator is a paginator for
// DescribeServerlessCacheSnapshots
type DescribeServerlessCacheSnapshotsPaginator struct {
	options   DescribeServerlessCacheSnapshotsPaginatorOptions
	client    DescribeServerlessCacheSnapshotsAPIClient
	params    *DescribeServerlessCacheSnapshotsInput
	nextToken *string
	firstPage bool
}

// NewDescribeServerlessCacheSnapshotsPaginator returns a new
// DescribeServerlessCacheSnapshotsPaginator
func NewDescribeServerlessCacheSnapshotsPaginator(client DescribeServerlessCacheSnapshotsAPIClient, params *DescribeServerlessCacheSnapshotsInput, optFns ...func(*DescribeServerlessCacheSnapshotsPaginatorOptions)) *DescribeServerlessCacheSnapshotsPaginator {
	if params == nil {
		params = &DescribeServerlessCacheSnapshotsInput{}
	}

	options := DescribeServerlessCacheSnapshotsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeServerlessCacheSnapshotsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeServerlessCacheSnapshotsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeServerlessCacheSnapshots page.
func (p *DescribeServerlessCacheSnapshotsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeServerlessCacheSnapshotsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeServerlessCacheSnapshots(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeServerlessCacheSnapshots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeServerlessCacheSnapshots",
	}
}
