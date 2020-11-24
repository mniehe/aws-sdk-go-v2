// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the replication subnet groups.
func (c *Client) DescribeReplicationSubnetGroups(ctx context.Context, params *DescribeReplicationSubnetGroupsInput, optFns ...func(*Options)) (*DescribeReplicationSubnetGroupsOutput, error) {
	if params == nil {
		params = &DescribeReplicationSubnetGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReplicationSubnetGroups", params, optFns, addOperationDescribeReplicationSubnetGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReplicationSubnetGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeReplicationSubnetGroupsInput struct {

	// Filters applied to replication subnet groups. Valid filter names:
	// replication-subnet-group-id
	Filters []types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

//
type DescribeReplicationSubnetGroupsOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// A description of the replication subnet groups.
	ReplicationSubnetGroups []types.ReplicationSubnetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeReplicationSubnetGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeReplicationSubnetGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeReplicationSubnetGroups{}, middleware.After)
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
	if err = addOpDescribeReplicationSubnetGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReplicationSubnetGroups(options.Region), middleware.Before); err != nil {
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

// DescribeReplicationSubnetGroupsAPIClient is a client that implements the
// DescribeReplicationSubnetGroups operation.
type DescribeReplicationSubnetGroupsAPIClient interface {
	DescribeReplicationSubnetGroups(context.Context, *DescribeReplicationSubnetGroupsInput, ...func(*Options)) (*DescribeReplicationSubnetGroupsOutput, error)
}

var _ DescribeReplicationSubnetGroupsAPIClient = (*Client)(nil)

// DescribeReplicationSubnetGroupsPaginatorOptions is the paginator options for
// DescribeReplicationSubnetGroups
type DescribeReplicationSubnetGroupsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReplicationSubnetGroupsPaginator is a paginator for
// DescribeReplicationSubnetGroups
type DescribeReplicationSubnetGroupsPaginator struct {
	options   DescribeReplicationSubnetGroupsPaginatorOptions
	client    DescribeReplicationSubnetGroupsAPIClient
	params    *DescribeReplicationSubnetGroupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeReplicationSubnetGroupsPaginator returns a new
// DescribeReplicationSubnetGroupsPaginator
func NewDescribeReplicationSubnetGroupsPaginator(client DescribeReplicationSubnetGroupsAPIClient, params *DescribeReplicationSubnetGroupsInput, optFns ...func(*DescribeReplicationSubnetGroupsPaginatorOptions)) *DescribeReplicationSubnetGroupsPaginator {
	options := DescribeReplicationSubnetGroupsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeReplicationSubnetGroupsInput{}
	}

	return &DescribeReplicationSubnetGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReplicationSubnetGroupsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeReplicationSubnetGroups page.
func (p *DescribeReplicationSubnetGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReplicationSubnetGroupsOutput, error) {
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

	result, err := p.client.DescribeReplicationSubnetGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeReplicationSubnetGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DescribeReplicationSubnetGroups",
	}
}
