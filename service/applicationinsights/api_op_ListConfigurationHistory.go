// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationinsights

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists the INFO, WARN, and ERROR events for periodic configuration updates
// performed by Application Insights. Examples of events represented are:
//
// * INFO:
// creating a new alarm or updating an alarm threshold.
//
// * WARN: alarm not created
// due to insufficient data points used to predict thresholds.
//
// * ERROR: alarm not
// created due to permission errors or exceeding quotas.
func (c *Client) ListConfigurationHistory(ctx context.Context, params *ListConfigurationHistoryInput, optFns ...func(*Options)) (*ListConfigurationHistoryOutput, error) {
	if params == nil {
		params = &ListConfigurationHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConfigurationHistory", params, optFns, addOperationListConfigurationHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConfigurationHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConfigurationHistoryInput struct {

	// The end time of the event.
	EndTime *time.Time

	// The status of the configuration update event. Possible values include INFO,
	// WARN, and ERROR.
	EventStatus types.ConfigurationEventStatus

	// The maximum number of results returned by ListConfigurationHistory in paginated
	// output. When this parameter is used, ListConfigurationHistory returns only
	// MaxResults in a single page along with a NextToken response element. The
	// remaining results of the initial request can be seen by sending another
	// ListConfigurationHistory request with the returned NextToken value. If this
	// parameter is not used, then ListConfigurationHistory returns all results.
	MaxResults *int32

	// The NextToken value returned from a previous paginated ListConfigurationHistory
	// request where MaxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the NextToken value. This value is null when there are no more results
	// to return.
	NextToken *string

	// Resource group to which the application belongs.
	ResourceGroupName *string

	// The start time of the event.
	StartTime *time.Time
}

type ListConfigurationHistoryOutput struct {

	// The list of configuration events and their corresponding details.
	EventList []types.ConfigurationEvent

	// The NextToken value to include in a future ListConfigurationHistory request.
	// When the results of a ListConfigurationHistory request exceed MaxResults, this
	// value can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListConfigurationHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListConfigurationHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListConfigurationHistory{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConfigurationHistory(options.Region), middleware.Before); err != nil {
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

// ListConfigurationHistoryAPIClient is a client that implements the
// ListConfigurationHistory operation.
type ListConfigurationHistoryAPIClient interface {
	ListConfigurationHistory(context.Context, *ListConfigurationHistoryInput, ...func(*Options)) (*ListConfigurationHistoryOutput, error)
}

var _ ListConfigurationHistoryAPIClient = (*Client)(nil)

// ListConfigurationHistoryPaginatorOptions is the paginator options for
// ListConfigurationHistory
type ListConfigurationHistoryPaginatorOptions struct {
	// The maximum number of results returned by ListConfigurationHistory in paginated
	// output. When this parameter is used, ListConfigurationHistory returns only
	// MaxResults in a single page along with a NextToken response element. The
	// remaining results of the initial request can be seen by sending another
	// ListConfigurationHistory request with the returned NextToken value. If this
	// parameter is not used, then ListConfigurationHistory returns all results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListConfigurationHistoryPaginator is a paginator for ListConfigurationHistory
type ListConfigurationHistoryPaginator struct {
	options   ListConfigurationHistoryPaginatorOptions
	client    ListConfigurationHistoryAPIClient
	params    *ListConfigurationHistoryInput
	nextToken *string
	firstPage bool
}

// NewListConfigurationHistoryPaginator returns a new
// ListConfigurationHistoryPaginator
func NewListConfigurationHistoryPaginator(client ListConfigurationHistoryAPIClient, params *ListConfigurationHistoryInput, optFns ...func(*ListConfigurationHistoryPaginatorOptions)) *ListConfigurationHistoryPaginator {
	options := ListConfigurationHistoryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListConfigurationHistoryInput{}
	}

	return &ListConfigurationHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListConfigurationHistoryPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListConfigurationHistory page.
func (p *ListConfigurationHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListConfigurationHistoryOutput, error) {
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

	result, err := p.client.ListConfigurationHistory(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListConfigurationHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "applicationinsights",
		OperationName: "ListConfigurationHistory",
	}
}
