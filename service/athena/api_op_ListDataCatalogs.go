// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the data catalogs in the current AWS account.
func (c *Client) ListDataCatalogs(ctx context.Context, params *ListDataCatalogsInput, optFns ...func(*Options)) (*ListDataCatalogsOutput, error) {
	if params == nil {
		params = &ListDataCatalogsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataCatalogs", params, optFns, addOperationListDataCatalogsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataCatalogsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataCatalogsInput struct {

	// Specifies the maximum number of data catalogs to return.
	MaxResults *int32

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string
}

type ListDataCatalogsOutput struct {

	// A summary list of data catalogs.
	DataCatalogsSummary []types.DataCatalogSummary

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDataCatalogsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDataCatalogs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDataCatalogs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDataCatalogs(options.Region), middleware.Before); err != nil {
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

// ListDataCatalogsAPIClient is a client that implements the ListDataCatalogs
// operation.
type ListDataCatalogsAPIClient interface {
	ListDataCatalogs(context.Context, *ListDataCatalogsInput, ...func(*Options)) (*ListDataCatalogsOutput, error)
}

var _ ListDataCatalogsAPIClient = (*Client)(nil)

// ListDataCatalogsPaginatorOptions is the paginator options for ListDataCatalogs
type ListDataCatalogsPaginatorOptions struct {
	// Specifies the maximum number of data catalogs to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDataCatalogsPaginator is a paginator for ListDataCatalogs
type ListDataCatalogsPaginator struct {
	options   ListDataCatalogsPaginatorOptions
	client    ListDataCatalogsAPIClient
	params    *ListDataCatalogsInput
	nextToken *string
	firstPage bool
}

// NewListDataCatalogsPaginator returns a new ListDataCatalogsPaginator
func NewListDataCatalogsPaginator(client ListDataCatalogsAPIClient, params *ListDataCatalogsInput, optFns ...func(*ListDataCatalogsPaginatorOptions)) *ListDataCatalogsPaginator {
	options := ListDataCatalogsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListDataCatalogsInput{}
	}

	return &ListDataCatalogsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDataCatalogsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListDataCatalogs page.
func (p *ListDataCatalogsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDataCatalogsOutput, error) {
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

	result, err := p.client.ListDataCatalogs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDataCatalogs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "ListDataCatalogs",
	}
}
