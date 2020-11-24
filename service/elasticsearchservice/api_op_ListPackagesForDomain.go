// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all packages associated with the Amazon ES domain.
func (c *Client) ListPackagesForDomain(ctx context.Context, params *ListPackagesForDomainInput, optFns ...func(*Options)) (*ListPackagesForDomainOutput, error) {
	if params == nil {
		params = &ListPackagesForDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPackagesForDomain", params, optFns, addOperationListPackagesForDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPackagesForDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to ListPackagesForDomain operation.
type ListPackagesForDomainInput struct {

	// The name of the domain for which you want to list associated packages.
	//
	// This member is required.
	DomainName *string

	// Limits results to a maximum number of packages.
	MaxResults int32

	// Used for pagination. Only necessary if a previous API call includes a non-null
	// NextToken value. If provided, returns results for the next page.
	NextToken *string
}

// Container for response parameters to ListPackagesForDomain operation.
type ListPackagesForDomainOutput struct {

	// List of DomainPackageDetails objects.
	DomainPackageDetailsList []types.DomainPackageDetails

	// Pagination token that needs to be supplied to the next call to get the next page
	// of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPackagesForDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPackagesForDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPackagesForDomain{}, middleware.After)
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
	if err = addOpListPackagesForDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPackagesForDomain(options.Region), middleware.Before); err != nil {
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

// ListPackagesForDomainAPIClient is a client that implements the
// ListPackagesForDomain operation.
type ListPackagesForDomainAPIClient interface {
	ListPackagesForDomain(context.Context, *ListPackagesForDomainInput, ...func(*Options)) (*ListPackagesForDomainOutput, error)
}

var _ ListPackagesForDomainAPIClient = (*Client)(nil)

// ListPackagesForDomainPaginatorOptions is the paginator options for
// ListPackagesForDomain
type ListPackagesForDomainPaginatorOptions struct {
	// Limits results to a maximum number of packages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPackagesForDomainPaginator is a paginator for ListPackagesForDomain
type ListPackagesForDomainPaginator struct {
	options   ListPackagesForDomainPaginatorOptions
	client    ListPackagesForDomainAPIClient
	params    *ListPackagesForDomainInput
	nextToken *string
	firstPage bool
}

// NewListPackagesForDomainPaginator returns a new ListPackagesForDomainPaginator
func NewListPackagesForDomainPaginator(client ListPackagesForDomainAPIClient, params *ListPackagesForDomainInput, optFns ...func(*ListPackagesForDomainPaginatorOptions)) *ListPackagesForDomainPaginator {
	options := ListPackagesForDomainPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListPackagesForDomainInput{}
	}

	return &ListPackagesForDomainPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPackagesForDomainPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListPackagesForDomain page.
func (p *ListPackagesForDomainPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPackagesForDomainOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListPackagesForDomain(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPackagesForDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "ListPackagesForDomain",
	}
}
