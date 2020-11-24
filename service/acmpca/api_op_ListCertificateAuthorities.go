// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the private certificate authorities that you created by using the
// CreateCertificateAuthority
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CreateCertificateAuthority.html)
// action.
func (c *Client) ListCertificateAuthorities(ctx context.Context, params *ListCertificateAuthoritiesInput, optFns ...func(*Options)) (*ListCertificateAuthoritiesOutput, error) {
	if params == nil {
		params = &ListCertificateAuthoritiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCertificateAuthorities", params, optFns, addOperationListCertificateAuthoritiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCertificateAuthoritiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCertificateAuthoritiesInput struct {

	// Use this parameter when paginating results to specify the maximum number of
	// items to return in the response on each page. If additional items exist beyond
	// the number you specify, the NextToken element is sent in the response. Use this
	// NextToken value in a subsequent request to retrieve additional items.
	MaxResults *int32

	// Use this parameter when paginating results in a subsequent request after you
	// receive a response with truncated results. Set it to the value of the NextToken
	// parameter from the response you just received.
	NextToken *string

	// Use this parameter to filter the returned set of certificate authorities based
	// on their owner. The default is SELF.
	ResourceOwner types.ResourceOwner
}

type ListCertificateAuthoritiesOutput struct {

	// Summary information about each certificate authority you have created.
	CertificateAuthorities []types.CertificateAuthority

	// When the list is truncated, this value is present and should be used for the
	// NextToken parameter in a subsequent pagination request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListCertificateAuthoritiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCertificateAuthorities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCertificateAuthorities{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCertificateAuthorities(options.Region), middleware.Before); err != nil {
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

// ListCertificateAuthoritiesAPIClient is a client that implements the
// ListCertificateAuthorities operation.
type ListCertificateAuthoritiesAPIClient interface {
	ListCertificateAuthorities(context.Context, *ListCertificateAuthoritiesInput, ...func(*Options)) (*ListCertificateAuthoritiesOutput, error)
}

var _ ListCertificateAuthoritiesAPIClient = (*Client)(nil)

// ListCertificateAuthoritiesPaginatorOptions is the paginator options for
// ListCertificateAuthorities
type ListCertificateAuthoritiesPaginatorOptions struct {
	// Use this parameter when paginating results to specify the maximum number of
	// items to return in the response on each page. If additional items exist beyond
	// the number you specify, the NextToken element is sent in the response. Use this
	// NextToken value in a subsequent request to retrieve additional items.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCertificateAuthoritiesPaginator is a paginator for
// ListCertificateAuthorities
type ListCertificateAuthoritiesPaginator struct {
	options   ListCertificateAuthoritiesPaginatorOptions
	client    ListCertificateAuthoritiesAPIClient
	params    *ListCertificateAuthoritiesInput
	nextToken *string
	firstPage bool
}

// NewListCertificateAuthoritiesPaginator returns a new
// ListCertificateAuthoritiesPaginator
func NewListCertificateAuthoritiesPaginator(client ListCertificateAuthoritiesAPIClient, params *ListCertificateAuthoritiesInput, optFns ...func(*ListCertificateAuthoritiesPaginatorOptions)) *ListCertificateAuthoritiesPaginator {
	options := ListCertificateAuthoritiesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListCertificateAuthoritiesInput{}
	}

	return &ListCertificateAuthoritiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCertificateAuthoritiesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListCertificateAuthorities page.
func (p *ListCertificateAuthoritiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCertificateAuthoritiesOutput, error) {
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

	result, err := p.client.ListCertificateAuthorities(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCertificateAuthorities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "ListCertificateAuthorities",
	}
}
