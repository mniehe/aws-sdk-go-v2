// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action returns metadata about active and previous legal holds.
func (c *Client) ListLegalHolds(ctx context.Context, params *ListLegalHoldsInput, optFns ...func(*Options)) (*ListLegalHoldsOutput, error) {
	if params == nil {
		params = &ListLegalHoldsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLegalHolds", params, optFns, c.addOperationListLegalHoldsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLegalHoldsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLegalHoldsInput struct {

	// The maximum number of resource list items to be returned.
	MaxResults *int32

	// The next item following a partial list of returned resources. For example, if a
	// request is made to return MaxResults number of resources, NextToken allows you
	// to return more items in your list starting at the location pointed to by the
	// next token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLegalHoldsOutput struct {

	// This is an array of returned legal holds, both active and previous.
	LegalHolds []types.LegalHold

	// The next item following a partial list of returned resources. For example, if a
	// request is made to return MaxResults number of resources, NextToken allows you
	// to return more items in your list starting at the location pointed to by the
	// next token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLegalHoldsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLegalHolds{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLegalHolds{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListLegalHolds"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLegalHolds(options.Region), middleware.Before); err != nil {
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

// ListLegalHoldsAPIClient is a client that implements the ListLegalHolds
// operation.
type ListLegalHoldsAPIClient interface {
	ListLegalHolds(context.Context, *ListLegalHoldsInput, ...func(*Options)) (*ListLegalHoldsOutput, error)
}

var _ ListLegalHoldsAPIClient = (*Client)(nil)

// ListLegalHoldsPaginatorOptions is the paginator options for ListLegalHolds
type ListLegalHoldsPaginatorOptions struct {
	// The maximum number of resource list items to be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLegalHoldsPaginator is a paginator for ListLegalHolds
type ListLegalHoldsPaginator struct {
	options   ListLegalHoldsPaginatorOptions
	client    ListLegalHoldsAPIClient
	params    *ListLegalHoldsInput
	nextToken *string
	firstPage bool
}

// NewListLegalHoldsPaginator returns a new ListLegalHoldsPaginator
func NewListLegalHoldsPaginator(client ListLegalHoldsAPIClient, params *ListLegalHoldsInput, optFns ...func(*ListLegalHoldsPaginatorOptions)) *ListLegalHoldsPaginator {
	if params == nil {
		params = &ListLegalHoldsInput{}
	}

	options := ListLegalHoldsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLegalHoldsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLegalHoldsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLegalHolds page.
func (p *ListLegalHoldsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLegalHoldsOutput, error) {
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

	result, err := p.client.ListLegalHolds(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLegalHolds(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListLegalHolds",
	}
}
