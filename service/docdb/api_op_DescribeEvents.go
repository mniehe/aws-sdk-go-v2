// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/docdb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns events related to instances, security groups, snapshots, and DB
// parameter groups for the past 14 days. You can obtain events specific to a
// particular DB instance, security group, snapshot, or parameter group by
// providing the name as a parameter. By default, the events of the past hour are
// returned.
func (c *Client) DescribeEvents(ctx context.Context, params *DescribeEventsInput, optFns ...func(*Options)) (*DescribeEventsOutput, error) {
	if params == nil {
		params = &DescribeEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEvents", params, optFns, c.addOperationDescribeEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to DescribeEvents .
type DescribeEventsInput struct {

	// The number of minutes to retrieve events for. Default: 60
	Duration *int32

	// The end of the time interval for which to retrieve events, specified in ISO
	// 8601 format. Example: 2009-07-08T18:00Z
	EndTime *time.Time

	// A list of event categories that trigger notifications for an event notification
	// subscription.
	EventCategories []string

	// This parameter is not currently supported.
	Filters []types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token (marker) is included in
	// the response so that the remaining results can be retrieved. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// The identifier of the event source for which events are returned. If not
	// specified, then all sources are included in the response. Constraints:
	//   - If SourceIdentifier is provided, SourceType must also be provided.
	//   - If the source type is DBInstance , a DBInstanceIdentifier must be provided.
	//   - If the source type is DBSecurityGroup , a DBSecurityGroupName must be
	//   provided.
	//   - If the source type is DBParameterGroup , a DBParameterGroupName must be
	//   provided.
	//   - If the source type is DBSnapshot , a DBSnapshotIdentifier must be provided.
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	SourceIdentifier *string

	// The event source to retrieve events for. If no value is specified, all events
	// are returned.
	SourceType types.SourceType

	// The beginning of the time interval to retrieve events for, specified in ISO
	// 8601 format. Example: 2009-07-08T18:00Z
	StartTime *time.Time

	noSmithyDocumentSerde
}

// Represents the output of DescribeEvents .
type DescribeEventsOutput struct {

	// Detailed information about one or more events.
	Events []types.Event

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEvents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEvents"); err != nil {
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
	if err = addOpDescribeEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEvents(options.Region), middleware.Before); err != nil {
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

// DescribeEventsAPIClient is a client that implements the DescribeEvents
// operation.
type DescribeEventsAPIClient interface {
	DescribeEvents(context.Context, *DescribeEventsInput, ...func(*Options)) (*DescribeEventsOutput, error)
}

var _ DescribeEventsAPIClient = (*Client)(nil)

// DescribeEventsPaginatorOptions is the paginator options for DescribeEvents
type DescribeEventsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token (marker) is included in
	// the response so that the remaining results can be retrieved. Default: 100
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEventsPaginator is a paginator for DescribeEvents
type DescribeEventsPaginator struct {
	options   DescribeEventsPaginatorOptions
	client    DescribeEventsAPIClient
	params    *DescribeEventsInput
	nextToken *string
	firstPage bool
}

// NewDescribeEventsPaginator returns a new DescribeEventsPaginator
func NewDescribeEventsPaginator(client DescribeEventsAPIClient, params *DescribeEventsInput, optFns ...func(*DescribeEventsPaginatorOptions)) *DescribeEventsPaginator {
	if params == nil {
		params = &DescribeEventsInput{}
	}

	options := DescribeEventsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEvents page.
func (p *DescribeEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEventsOutput, error) {
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

	result, err := p.client.DescribeEvents(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEvents",
	}
}
