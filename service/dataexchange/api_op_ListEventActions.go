// Code generated by smithy-go-codegen DO NOT EDIT.

package dataexchange

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/dataexchange/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation lists your event actions.
func (c *Client) ListEventActions(ctx context.Context, params *ListEventActionsInput, optFns ...func(*Options)) (*ListEventActionsOutput, error) {
	if params == nil {
		params = &ListEventActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEventActions", params, optFns, c.addOperationListEventActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEventActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEventActionsInput struct {

	// The unique identifier for the event source.
	EventSourceId *string

	// The maximum number of results returned by a single call.
	MaxResults int32

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEventActionsOutput struct {

	// The event action objects listed by the request.
	EventActions []types.EventActionEntry

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEventActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEventActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEventActions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEventActions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEventActions(options.Region), middleware.Before); err != nil {
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

// ListEventActionsAPIClient is a client that implements the ListEventActions
// operation.
type ListEventActionsAPIClient interface {
	ListEventActions(context.Context, *ListEventActionsInput, ...func(*Options)) (*ListEventActionsOutput, error)
}

var _ ListEventActionsAPIClient = (*Client)(nil)

// ListEventActionsPaginatorOptions is the paginator options for ListEventActions
type ListEventActionsPaginatorOptions struct {
	// The maximum number of results returned by a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEventActionsPaginator is a paginator for ListEventActions
type ListEventActionsPaginator struct {
	options   ListEventActionsPaginatorOptions
	client    ListEventActionsAPIClient
	params    *ListEventActionsInput
	nextToken *string
	firstPage bool
}

// NewListEventActionsPaginator returns a new ListEventActionsPaginator
func NewListEventActionsPaginator(client ListEventActionsAPIClient, params *ListEventActionsInput, optFns ...func(*ListEventActionsPaginatorOptions)) *ListEventActionsPaginator {
	if params == nil {
		params = &ListEventActionsInput{}
	}

	options := ListEventActionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEventActionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEventActionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEventActions page.
func (p *ListEventActionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEventActionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListEventActions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEventActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEventActions",
	}
}
