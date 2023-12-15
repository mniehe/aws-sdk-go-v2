// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the names of all DevEndpoint resources in this Amazon Web Services
// account, or the resources with the specified tag. This operation allows you to
// see which resources are available in your account, and their names. This
// operation takes the optional Tags field, which you can use as a filter on the
// response so that tagged resources can be retrieved as a group. If you choose to
// use tags filtering, only resources with the tag are retrieved.
func (c *Client) ListDevEndpoints(ctx context.Context, params *ListDevEndpointsInput, optFns ...func(*Options)) (*ListDevEndpointsOutput, error) {
	if params == nil {
		params = &ListDevEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDevEndpoints", params, optFns, c.addOperationListDevEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDevEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDevEndpointsInput struct {

	// The maximum size of a list to return.
	MaxResults *int32

	// A continuation token, if this is a continuation request.
	NextToken *string

	// Specifies to return only these tagged resources.
	Tags map[string]string

	noSmithyDocumentSerde
}

type ListDevEndpointsOutput struct {

	// The names of all the DevEndpoint s in the account, or the DevEndpoint s with the
	// specified tags.
	DevEndpointNames []string

	// A continuation token, if the returned list does not contain the last metric
	// available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDevEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDevEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDevEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDevEndpoints"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDevEndpoints(options.Region), middleware.Before); err != nil {
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

// ListDevEndpointsAPIClient is a client that implements the ListDevEndpoints
// operation.
type ListDevEndpointsAPIClient interface {
	ListDevEndpoints(context.Context, *ListDevEndpointsInput, ...func(*Options)) (*ListDevEndpointsOutput, error)
}

var _ ListDevEndpointsAPIClient = (*Client)(nil)

// ListDevEndpointsPaginatorOptions is the paginator options for ListDevEndpoints
type ListDevEndpointsPaginatorOptions struct {
	// The maximum size of a list to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDevEndpointsPaginator is a paginator for ListDevEndpoints
type ListDevEndpointsPaginator struct {
	options   ListDevEndpointsPaginatorOptions
	client    ListDevEndpointsAPIClient
	params    *ListDevEndpointsInput
	nextToken *string
	firstPage bool
}

// NewListDevEndpointsPaginator returns a new ListDevEndpointsPaginator
func NewListDevEndpointsPaginator(client ListDevEndpointsAPIClient, params *ListDevEndpointsInput, optFns ...func(*ListDevEndpointsPaginatorOptions)) *ListDevEndpointsPaginator {
	if params == nil {
		params = &ListDevEndpointsInput{}
	}

	options := ListDevEndpointsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDevEndpointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDevEndpointsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDevEndpoints page.
func (p *ListDevEndpointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDevEndpointsOutput, error) {
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

	result, err := p.client.ListDevEndpoints(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDevEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDevEndpoints",
	}
}
