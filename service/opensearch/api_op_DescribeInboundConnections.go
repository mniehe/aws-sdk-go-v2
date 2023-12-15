// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the inbound cross-cluster search connections for a destination
// (remote) Amazon OpenSearch Service domain. For more information, see
// Cross-cluster search for Amazon OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cross-cluster-search.html)
// .
func (c *Client) DescribeInboundConnections(ctx context.Context, params *DescribeInboundConnectionsInput, optFns ...func(*Options)) (*DescribeInboundConnectionsOutput, error) {
	if params == nil {
		params = &DescribeInboundConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInboundConnections", params, optFns, c.addOperationDescribeInboundConnectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInboundConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeInboundConnections operation.
type DescribeInboundConnectionsInput struct {

	// A list of filters used to match properties for inbound cross-cluster
	// connections.
	Filters []types.Filter

	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	MaxResults int32

	// If your initial DescribeInboundConnections operation returns a nextToken , you
	// can include the returned nextToken in subsequent DescribeInboundConnections
	// operations, which returns results in the next page.
	NextToken *string

	noSmithyDocumentSerde
}

// Contains a list of connections matching the filter criteria.
type DescribeInboundConnectionsOutput struct {

	// List of inbound connections.
	Connections []types.InboundConnection

	// When nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInboundConnectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeInboundConnections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeInboundConnections{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeInboundConnections"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInboundConnections(options.Region), middleware.Before); err != nil {
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

// DescribeInboundConnectionsAPIClient is a client that implements the
// DescribeInboundConnections operation.
type DescribeInboundConnectionsAPIClient interface {
	DescribeInboundConnections(context.Context, *DescribeInboundConnectionsInput, ...func(*Options)) (*DescribeInboundConnectionsOutput, error)
}

var _ DescribeInboundConnectionsAPIClient = (*Client)(nil)

// DescribeInboundConnectionsPaginatorOptions is the paginator options for
// DescribeInboundConnections
type DescribeInboundConnectionsPaginatorOptions struct {
	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeInboundConnectionsPaginator is a paginator for
// DescribeInboundConnections
type DescribeInboundConnectionsPaginator struct {
	options   DescribeInboundConnectionsPaginatorOptions
	client    DescribeInboundConnectionsAPIClient
	params    *DescribeInboundConnectionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeInboundConnectionsPaginator returns a new
// DescribeInboundConnectionsPaginator
func NewDescribeInboundConnectionsPaginator(client DescribeInboundConnectionsAPIClient, params *DescribeInboundConnectionsInput, optFns ...func(*DescribeInboundConnectionsPaginatorOptions)) *DescribeInboundConnectionsPaginator {
	if params == nil {
		params = &DescribeInboundConnectionsInput{}
	}

	options := DescribeInboundConnectionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeInboundConnectionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeInboundConnectionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeInboundConnections page.
func (p *DescribeInboundConnectionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeInboundConnectionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeInboundConnections(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeInboundConnections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeInboundConnections",
	}
}
