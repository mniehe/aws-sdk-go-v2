// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a list of available query IDs only for queries saved in the specified
// workgroup. Requires that you have access to the specified workgroup. If a
// workgroup is not specified, lists the saved queries for the primary workgroup.
func (c *Client) ListNamedQueries(ctx context.Context, params *ListNamedQueriesInput, optFns ...func(*Options)) (*ListNamedQueriesOutput, error) {
	if params == nil {
		params = &ListNamedQueriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNamedQueries", params, optFns, c.addOperationListNamedQueriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNamedQueriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNamedQueriesInput struct {

	// The maximum number of queries to return in this request.
	MaxResults *int32

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	// The name of the workgroup from which the named queries are being returned. If a
	// workgroup is not specified, the saved queries for the primary workgroup are
	// returned.
	WorkGroup *string

	noSmithyDocumentSerde
}

type ListNamedQueriesOutput struct {

	// The list of unique query IDs.
	NamedQueryIds []string

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListNamedQueriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListNamedQueries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListNamedQueries{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListNamedQueries"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListNamedQueries(options.Region), middleware.Before); err != nil {
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

// ListNamedQueriesAPIClient is a client that implements the ListNamedQueries
// operation.
type ListNamedQueriesAPIClient interface {
	ListNamedQueries(context.Context, *ListNamedQueriesInput, ...func(*Options)) (*ListNamedQueriesOutput, error)
}

var _ ListNamedQueriesAPIClient = (*Client)(nil)

// ListNamedQueriesPaginatorOptions is the paginator options for ListNamedQueries
type ListNamedQueriesPaginatorOptions struct {
	// The maximum number of queries to return in this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListNamedQueriesPaginator is a paginator for ListNamedQueries
type ListNamedQueriesPaginator struct {
	options   ListNamedQueriesPaginatorOptions
	client    ListNamedQueriesAPIClient
	params    *ListNamedQueriesInput
	nextToken *string
	firstPage bool
}

// NewListNamedQueriesPaginator returns a new ListNamedQueriesPaginator
func NewListNamedQueriesPaginator(client ListNamedQueriesAPIClient, params *ListNamedQueriesInput, optFns ...func(*ListNamedQueriesPaginatorOptions)) *ListNamedQueriesPaginator {
	if params == nil {
		params = &ListNamedQueriesInput{}
	}

	options := ListNamedQueriesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListNamedQueriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListNamedQueriesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListNamedQueries page.
func (p *ListNamedQueriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListNamedQueriesOutput, error) {
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

	result, err := p.client.ListNamedQueries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListNamedQueries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListNamedQueries",
	}
}
