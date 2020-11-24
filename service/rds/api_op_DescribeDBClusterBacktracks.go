// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about backtracks for a DB cluster. For more information on
// Amazon Aurora, see  What Is Amazon Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. This action only applies to Aurora MySQL DB
// clusters.
func (c *Client) DescribeDBClusterBacktracks(ctx context.Context, params *DescribeDBClusterBacktracksInput, optFns ...func(*Options)) (*DescribeDBClusterBacktracksOutput, error) {
	if params == nil {
		params = &DescribeDBClusterBacktracksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBClusterBacktracks", params, optFns, addOperationDescribeDBClusterBacktracksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBClusterBacktracksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeDBClusterBacktracksInput struct {

	// The DB cluster identifier of the DB cluster to be described. This parameter is
	// stored as a lowercase string. Constraints:
	//
	// * Must contain from 1 to 63
	// alphanumeric characters or hyphens.
	//
	// * First character must be a letter.
	//
	// *
	// Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example:
	// my-cluster1
	//
	// This member is required.
	DBClusterIdentifier *string

	// If specified, this value is the backtrack identifier of the backtrack to be
	// described. Constraints:
	//
	// * Must contain a valid universally unique identifier
	// (UUID). For more information about UUIDs, see A Universally Unique Identifier
	// (UUID) URN Namespace (http://www.ietf.org/rfc/rfc4122.txt).
	//
	// Example:
	// 123e4567-e89b-12d3-a456-426655440000
	BacktrackIdentifier *string

	// A filter that specifies one or more DB clusters to describe. Supported filters
	// include the following:
	//
	// * db-cluster-backtrack-id - Accepts backtrack
	// identifiers. The results list includes information about only the backtracks
	// identified by these identifiers.
	//
	// * db-cluster-backtrack-status - Accepts any of
	// the following backtrack status values:
	//
	// * applying
	//
	// * completed
	//
	// * failed
	//
	// *
	// pending
	//
	// The results list includes information about only the backtracks
	// identified by these values.
	Filters []types.Filter

	// An optional pagination token provided by a previous DescribeDBClusterBacktracks
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

// Contains the result of a successful invocation of the
// DescribeDBClusterBacktracks action.
type DescribeDBClusterBacktracksOutput struct {

	// Contains a list of backtracks for the user.
	DBClusterBacktracks []types.DBClusterBacktrack

	// A pagination token that can be used in a later DescribeDBClusterBacktracks
	// request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDBClusterBacktracksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusterBacktracks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusterBacktracks{}, middleware.After)
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
	if err = addOpDescribeDBClusterBacktracksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusterBacktracks(options.Region), middleware.Before); err != nil {
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

// DescribeDBClusterBacktracksAPIClient is a client that implements the
// DescribeDBClusterBacktracks operation.
type DescribeDBClusterBacktracksAPIClient interface {
	DescribeDBClusterBacktracks(context.Context, *DescribeDBClusterBacktracksInput, ...func(*Options)) (*DescribeDBClusterBacktracksOutput, error)
}

var _ DescribeDBClusterBacktracksAPIClient = (*Client)(nil)

// DescribeDBClusterBacktracksPaginatorOptions is the paginator options for
// DescribeDBClusterBacktracks
type DescribeDBClusterBacktracksPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBClusterBacktracksPaginator is a paginator for
// DescribeDBClusterBacktracks
type DescribeDBClusterBacktracksPaginator struct {
	options   DescribeDBClusterBacktracksPaginatorOptions
	client    DescribeDBClusterBacktracksAPIClient
	params    *DescribeDBClusterBacktracksInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBClusterBacktracksPaginator returns a new
// DescribeDBClusterBacktracksPaginator
func NewDescribeDBClusterBacktracksPaginator(client DescribeDBClusterBacktracksAPIClient, params *DescribeDBClusterBacktracksInput, optFns ...func(*DescribeDBClusterBacktracksPaginatorOptions)) *DescribeDBClusterBacktracksPaginator {
	options := DescribeDBClusterBacktracksPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeDBClusterBacktracksInput{}
	}

	return &DescribeDBClusterBacktracksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBClusterBacktracksPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeDBClusterBacktracks page.
func (p *DescribeDBClusterBacktracksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBClusterBacktracksOutput, error) {
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

	result, err := p.client.DescribeDBClusterBacktracks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeDBClusterBacktracks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBClusterBacktracks",
	}
}
