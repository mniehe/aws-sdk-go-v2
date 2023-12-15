// Code generated by smithy-go-codegen DO NOT EDIT.

package voiceid

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/voiceid/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all fraudsters in a specified watchlist or domain.
func (c *Client) ListFraudsters(ctx context.Context, params *ListFraudstersInput, optFns ...func(*Options)) (*ListFraudstersOutput, error) {
	if params == nil {
		params = &ListFraudstersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFraudsters", params, optFns, c.addOperationListFraudstersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFraudstersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFraudstersInput struct {

	// The identifier of the domain.
	//
	// This member is required.
	DomainId *string

	// The maximum number of results that are returned per call. You can use NextToken
	// to obtain more pages of results. The default is 100; the maximum allowed page
	// size is also 100.
	MaxResults *int32

	// If NextToken is returned, there are more results available. The value of
	// NextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours.
	NextToken *string

	// The identifier of the watchlist. If provided, all fraudsters in the watchlist
	// are listed. If not provided, all fraudsters in the domain are listed.
	WatchlistId *string

	noSmithyDocumentSerde
}

type ListFraudstersOutput struct {

	// A list that contains details about each fraudster in the Amazon Web Services
	// account.
	FraudsterSummaries []types.FraudsterSummary

	// If NextToken is returned, there are more results available. The value of
	// NextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFraudstersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListFraudsters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListFraudsters{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFraudsters"); err != nil {
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
	if err = addOpListFraudstersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFraudsters(options.Region), middleware.Before); err != nil {
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

// ListFraudstersAPIClient is a client that implements the ListFraudsters
// operation.
type ListFraudstersAPIClient interface {
	ListFraudsters(context.Context, *ListFraudstersInput, ...func(*Options)) (*ListFraudstersOutput, error)
}

var _ ListFraudstersAPIClient = (*Client)(nil)

// ListFraudstersPaginatorOptions is the paginator options for ListFraudsters
type ListFraudstersPaginatorOptions struct {
	// The maximum number of results that are returned per call. You can use NextToken
	// to obtain more pages of results. The default is 100; the maximum allowed page
	// size is also 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFraudstersPaginator is a paginator for ListFraudsters
type ListFraudstersPaginator struct {
	options   ListFraudstersPaginatorOptions
	client    ListFraudstersAPIClient
	params    *ListFraudstersInput
	nextToken *string
	firstPage bool
}

// NewListFraudstersPaginator returns a new ListFraudstersPaginator
func NewListFraudstersPaginator(client ListFraudstersAPIClient, params *ListFraudstersInput, optFns ...func(*ListFraudstersPaginatorOptions)) *ListFraudstersPaginator {
	if params == nil {
		params = &ListFraudstersInput{}
	}

	options := ListFraudstersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFraudstersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFraudstersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFraudsters page.
func (p *ListFraudstersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFraudstersOutput, error) {
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

	result, err := p.client.ListFraudsters(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFraudsters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFraudsters",
	}
}
