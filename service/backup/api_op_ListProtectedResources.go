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

// Returns an array of resources successfully backed up by Backup, including the
// time the resource was saved, an Amazon Resource Name (ARN) of the resource, and
// a resource type.
func (c *Client) ListProtectedResources(ctx context.Context, params *ListProtectedResourcesInput, optFns ...func(*Options)) (*ListProtectedResourcesOutput, error) {
	if params == nil {
		params = &ListProtectedResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProtectedResources", params, optFns, c.addOperationListProtectedResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProtectedResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProtectedResourcesInput struct {

	// The maximum number of items to be returned.
	MaxResults *int32

	// The next item following a partial list of returned items. For example, if a
	// request is made to return MaxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListProtectedResourcesOutput struct {

	// The next item following a partial list of returned items. For example, if a
	// request is made to return MaxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	// An array of resources successfully backed up by Backup including the time the
	// resource was saved, an Amazon Resource Name (ARN) of the resource, and a
	// resource type.
	Results []types.ProtectedResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProtectedResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProtectedResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProtectedResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListProtectedResources"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProtectedResources(options.Region), middleware.Before); err != nil {
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

// ListProtectedResourcesAPIClient is a client that implements the
// ListProtectedResources operation.
type ListProtectedResourcesAPIClient interface {
	ListProtectedResources(context.Context, *ListProtectedResourcesInput, ...func(*Options)) (*ListProtectedResourcesOutput, error)
}

var _ ListProtectedResourcesAPIClient = (*Client)(nil)

// ListProtectedResourcesPaginatorOptions is the paginator options for
// ListProtectedResources
type ListProtectedResourcesPaginatorOptions struct {
	// The maximum number of items to be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProtectedResourcesPaginator is a paginator for ListProtectedResources
type ListProtectedResourcesPaginator struct {
	options   ListProtectedResourcesPaginatorOptions
	client    ListProtectedResourcesAPIClient
	params    *ListProtectedResourcesInput
	nextToken *string
	firstPage bool
}

// NewListProtectedResourcesPaginator returns a new ListProtectedResourcesPaginator
func NewListProtectedResourcesPaginator(client ListProtectedResourcesAPIClient, params *ListProtectedResourcesInput, optFns ...func(*ListProtectedResourcesPaginatorOptions)) *ListProtectedResourcesPaginator {
	if params == nil {
		params = &ListProtectedResourcesInput{}
	}

	options := ListProtectedResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProtectedResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProtectedResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProtectedResources page.
func (p *ListProtectedResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProtectedResourcesOutput, error) {
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

	result, err := p.client.ListProtectedResources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProtectedResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListProtectedResources",
	}
}
