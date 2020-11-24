// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the names of facets that exist in a schema.
func (c *Client) ListFacetNames(ctx context.Context, params *ListFacetNamesInput, optFns ...func(*Options)) (*ListFacetNamesOutput, error) {
	if params == nil {
		params = &ListFacetNamesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFacetNames", params, optFns, addOperationListFacetNamesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFacetNamesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFacetNamesInput struct {

	// The Amazon Resource Name (ARN) to retrieve facet names from.
	//
	// This member is required.
	SchemaArn *string

	// The maximum number of results to retrieve.
	MaxResults *int32

	// The pagination token.
	NextToken *string
}

type ListFacetNamesOutput struct {

	// The names of facets that exist within the schema.
	FacetNames []string

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListFacetNamesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFacetNames{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFacetNames{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListFacetNamesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFacetNames(options.Region), middleware.Before); err != nil {
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
	return nil
}

// ListFacetNamesAPIClient is a client that implements the ListFacetNames
// operation.
type ListFacetNamesAPIClient interface {
	ListFacetNames(context.Context, *ListFacetNamesInput, ...func(*Options)) (*ListFacetNamesOutput, error)
}

var _ ListFacetNamesAPIClient = (*Client)(nil)

// ListFacetNamesPaginatorOptions is the paginator options for ListFacetNames
type ListFacetNamesPaginatorOptions struct {
	// The maximum number of results to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFacetNamesPaginator is a paginator for ListFacetNames
type ListFacetNamesPaginator struct {
	options   ListFacetNamesPaginatorOptions
	client    ListFacetNamesAPIClient
	params    *ListFacetNamesInput
	nextToken *string
	firstPage bool
}

// NewListFacetNamesPaginator returns a new ListFacetNamesPaginator
func NewListFacetNamesPaginator(client ListFacetNamesAPIClient, params *ListFacetNamesInput, optFns ...func(*ListFacetNamesPaginatorOptions)) *ListFacetNamesPaginator {
	options := ListFacetNamesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListFacetNamesInput{}
	}

	return &ListFacetNamesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFacetNamesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListFacetNames page.
func (p *ListFacetNamesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFacetNamesOutput, error) {
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

	result, err := p.client.ListFacetNames(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListFacetNames(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListFacetNames",
	}
}
