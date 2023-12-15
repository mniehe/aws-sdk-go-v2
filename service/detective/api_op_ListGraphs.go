// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/detective/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the list of behavior graphs that the calling account is an
// administrator account of. This operation can only be called by an administrator
// account. Because an account can currently only be the administrator of one
// behavior graph within a Region, the results always contain a single behavior
// graph.
func (c *Client) ListGraphs(ctx context.Context, params *ListGraphsInput, optFns ...func(*Options)) (*ListGraphsOutput, error) {
	if params == nil {
		params = &ListGraphsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGraphs", params, optFns, c.addOperationListGraphsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGraphsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGraphsInput struct {

	// The maximum number of graphs to return at a time. The total must be less than
	// the overall limit on the number of results to return, which is currently 200.
	MaxResults *int32

	// For requests to get the next page of results, the pagination token that was
	// returned with the previous set of results. The initial request does not include
	// a pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListGraphsOutput struct {

	// A list of behavior graphs that the account is an administrator account for.
	GraphList []types.Graph

	// If there are more behavior graphs remaining in the results, then this is the
	// pagination token to use to request the next page of behavior graphs.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGraphsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListGraphs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListGraphs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListGraphs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGraphs(options.Region), middleware.Before); err != nil {
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

// ListGraphsAPIClient is a client that implements the ListGraphs operation.
type ListGraphsAPIClient interface {
	ListGraphs(context.Context, *ListGraphsInput, ...func(*Options)) (*ListGraphsOutput, error)
}

var _ ListGraphsAPIClient = (*Client)(nil)

// ListGraphsPaginatorOptions is the paginator options for ListGraphs
type ListGraphsPaginatorOptions struct {
	// The maximum number of graphs to return at a time. The total must be less than
	// the overall limit on the number of results to return, which is currently 200.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGraphsPaginator is a paginator for ListGraphs
type ListGraphsPaginator struct {
	options   ListGraphsPaginatorOptions
	client    ListGraphsAPIClient
	params    *ListGraphsInput
	nextToken *string
	firstPage bool
}

// NewListGraphsPaginator returns a new ListGraphsPaginator
func NewListGraphsPaginator(client ListGraphsAPIClient, params *ListGraphsInput, optFns ...func(*ListGraphsPaginatorOptions)) *ListGraphsPaginator {
	if params == nil {
		params = &ListGraphsInput{}
	}

	options := ListGraphsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGraphsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGraphsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGraphs page.
func (p *ListGraphsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGraphsOutput, error) {
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

	result, err := p.client.ListGraphs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListGraphs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListGraphs",
	}
}
