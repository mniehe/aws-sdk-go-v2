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
)

// Gets a list of locales for the specified bot.
func (c *Client) ListBotLocales(ctx context.Context, params *ListBotLocalesInput, optFns ...func(*Options)) (*ListBotLocalesOutput, error) {
	if params == nil {
		params = &ListBotLocalesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBotLocales", params, optFns, c.addOperationListBotLocalesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBotLocalesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBotLocalesInput struct {

	// The identifier of the bot to list locales for.
	//
	// This member is required.
	BotId *string

	// The version of the bot to list locales for.
	//
	// This member is required.
	BotVersion *string

	// Provides the specification for a filter used to limit the response to only
	// those locales that match the filter specification. You can only specify one
	// filter and one value to filter on.
	Filters []types.BotLocaleFilter

	// The maximum number of aliases to return in each page of results. If there are
	// fewer results than the max page size, only the actual number of results are
	// returned.
	MaxResults *int32

	// If the response from the ListBotLocales operation contains more results than
	// specified in the maxResults parameter, a token is returned in the response. Use
	// that token as the nextToken parameter to return the next page of results.
	NextToken *string

	// Specifies sorting parameters for the list of locales. You can sort by locale
	// name in ascending or descending order.
	SortBy *types.BotLocaleSortBy

	noSmithyDocumentSerde
}

type ListBotLocalesOutput struct {

	// The identifier of the bot to list locales for.
	BotId *string

	// Summary information for the locales that meet the filter criteria specified in
	// the request. The length of the list is specified in the maxResults parameter of
	// the request. If there are more locales available, the nextToken field contains
	// a token to get the next page of results.
	BotLocaleSummaries []types.BotLocaleSummary

	// The version of the bot.
	BotVersion *string

	// A token that indicates whether there are more results to return in a response
	// to the ListBotLocales operation. If the nextToken field is present, you send
	// the contents as the nextToken parameter of a ListBotLocales operation request
	// to get the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBotLocalesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBotLocales{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBotLocales{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListBotLocales"); err != nil {
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
	if err = addOpListBotLocalesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBotLocales(options.Region), middleware.Before); err != nil {
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

// ListBotLocalesAPIClient is a client that implements the ListBotLocales
// operation.
type ListBotLocalesAPIClient interface {
	ListBotLocales(context.Context, *ListBotLocalesInput, ...func(*Options)) (*ListBotLocalesOutput, error)
}

var _ ListBotLocalesAPIClient = (*Client)(nil)

// ListBotLocalesPaginatorOptions is the paginator options for ListBotLocales
type ListBotLocalesPaginatorOptions struct {
	// The maximum number of aliases to return in each page of results. If there are
	// fewer results than the max page size, only the actual number of results are
	// returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBotLocalesPaginator is a paginator for ListBotLocales
type ListBotLocalesPaginator struct {
	options   ListBotLocalesPaginatorOptions
	client    ListBotLocalesAPIClient
	params    *ListBotLocalesInput
	nextToken *string
	firstPage bool
}

// NewListBotLocalesPaginator returns a new ListBotLocalesPaginator
func NewListBotLocalesPaginator(client ListBotLocalesAPIClient, params *ListBotLocalesInput, optFns ...func(*ListBotLocalesPaginatorOptions)) *ListBotLocalesPaginator {
	if params == nil {
		params = &ListBotLocalesInput{}
	}

	options := ListBotLocalesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBotLocalesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBotLocalesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBotLocales page.
func (p *ListBotLocalesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBotLocalesOutput, error) {
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

	result, err := p.client.ListBotLocales(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBotLocales(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListBotLocales",
	}
}
