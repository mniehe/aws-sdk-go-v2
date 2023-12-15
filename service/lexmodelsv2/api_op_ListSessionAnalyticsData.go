// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves a list of metadata for individual user sessions with your bot. The
// startDateTime and endDateTime fields are required. These fields define a time
// range for which you want to retrieve results. Of the optional fields, you can
// organize the results in the following ways:
//   - Use the filters field to filter the results and the sortBy field to specify
//     the values by which to sort the results.
//   - Use the maxResults field to limit the number of results to return in a
//     single response and the nextToken field to return the next batch of results if
//     the response does not return the full set of results.
func (c *Client) ListSessionAnalyticsData(ctx context.Context, params *ListSessionAnalyticsDataInput, optFns ...func(*Options)) (*ListSessionAnalyticsDataOutput, error) {
	if params == nil {
		params = &ListSessionAnalyticsDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSessionAnalyticsData", params, optFns, c.addOperationListSessionAnalyticsDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSessionAnalyticsDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSessionAnalyticsDataInput struct {

	// The identifier for the bot for which you want to retrieve session analytics.
	//
	// This member is required.
	BotId *string

	// The date and time that marks the end of the range of time for which you want to
	// see session analytics.
	//
	// This member is required.
	EndDateTime *time.Time

	// The date and time that marks the beginning of the range of time for which you
	// want to see session analytics.
	//
	// This member is required.
	StartDateTime *time.Time

	// A list of objects, each of which describes a condition by which you want to
	// filter the results.
	Filters []types.AnalyticsSessionFilter

	// The maximum number of results to return in each page of results. If there are
	// fewer results than the maximum page size, only the actual number of results are
	// returned.
	MaxResults *int32

	// If the response from the ListSessionAnalyticsData operation contains more
	// results than specified in the maxResults parameter, a token is returned in the
	// response. Use the returned token in the nextToken parameter of a
	// ListSessionAnalyticsData request to return the next page of results. For a
	// complete set of results, call the ListSessionAnalyticsData operation until the
	// nextToken returned in the response is null.
	NextToken *string

	// An object specifying the measure and method by which to sort the session
	// analytics data.
	SortBy *types.SessionDataSortBy

	noSmithyDocumentSerde
}

type ListSessionAnalyticsDataOutput struct {

	// The unique identifier of the bot that the sessions belong to.
	BotId *string

	// If the response from the ListSessionAnalyticsData operation contains more
	// results than specified in the maxResults parameter, a token is returned in the
	// response. Use the returned token in the nextToken parameter of a
	// ListSessionAnalyticsData request to return the next page of results. For a
	// complete set of results, call the ListSessionAnalyticsData operation until the
	// nextToken returned in the response is null.
	NextToken *string

	// A list of objects, each of which contains information about a session with the
	// bot.
	Sessions []types.SessionSpecification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSessionAnalyticsDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSessionAnalyticsData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSessionAnalyticsData{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSessionAnalyticsData"); err != nil {
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
	if err = addOpListSessionAnalyticsDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSessionAnalyticsData(options.Region), middleware.Before); err != nil {
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

// ListSessionAnalyticsDataAPIClient is a client that implements the
// ListSessionAnalyticsData operation.
type ListSessionAnalyticsDataAPIClient interface {
	ListSessionAnalyticsData(context.Context, *ListSessionAnalyticsDataInput, ...func(*Options)) (*ListSessionAnalyticsDataOutput, error)
}

var _ ListSessionAnalyticsDataAPIClient = (*Client)(nil)

// ListSessionAnalyticsDataPaginatorOptions is the paginator options for
// ListSessionAnalyticsData
type ListSessionAnalyticsDataPaginatorOptions struct {
	// The maximum number of results to return in each page of results. If there are
	// fewer results than the maximum page size, only the actual number of results are
	// returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSessionAnalyticsDataPaginator is a paginator for ListSessionAnalyticsData
type ListSessionAnalyticsDataPaginator struct {
	options   ListSessionAnalyticsDataPaginatorOptions
	client    ListSessionAnalyticsDataAPIClient
	params    *ListSessionAnalyticsDataInput
	nextToken *string
	firstPage bool
}

// NewListSessionAnalyticsDataPaginator returns a new
// ListSessionAnalyticsDataPaginator
func NewListSessionAnalyticsDataPaginator(client ListSessionAnalyticsDataAPIClient, params *ListSessionAnalyticsDataInput, optFns ...func(*ListSessionAnalyticsDataPaginatorOptions)) *ListSessionAnalyticsDataPaginator {
	if params == nil {
		params = &ListSessionAnalyticsDataInput{}
	}

	options := ListSessionAnalyticsDataPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSessionAnalyticsDataPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSessionAnalyticsDataPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSessionAnalyticsData page.
func (p *ListSessionAnalyticsDataPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSessionAnalyticsDataOutput, error) {
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

	result, err := p.client.ListSessionAnalyticsData(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSessionAnalyticsData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSessionAnalyticsData",
	}
}
