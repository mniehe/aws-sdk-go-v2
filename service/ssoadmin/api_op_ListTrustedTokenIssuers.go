// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the trusted token issuers configured in an instance of IAM Identity
// Center.
func (c *Client) ListTrustedTokenIssuers(ctx context.Context, params *ListTrustedTokenIssuersInput, optFns ...func(*Options)) (*ListTrustedTokenIssuersOutput, error) {
	if params == nil {
		params = &ListTrustedTokenIssuersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTrustedTokenIssuers", params, optFns, c.addOperationListTrustedTokenIssuersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTrustedTokenIssuersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTrustedTokenIssuersInput struct {

	// Specifies the ARN of the instance of IAM Identity Center with the trusted token
	// issuer configurations that you want to list.
	//
	// This member is required.
	InstanceArn *string

	// Specifies the total number of results that you want included in each response.
	// If additional items exist beyond the number you specify, the NextToken response
	// element is returned with a value (not null). Include the specified value as the
	// NextToken request parameter in the next call to the operation to get the next
	// set of results. Note that the service might return fewer results than the
	// maximum even when there are more results available. You should check NextToken
	// after every operation to ensure that you receive all of the results.
	MaxResults *int32

	// Specifies that you want to receive the next page of results. Valid only if you
	// received a NextToken response in the previous request. If you did, it indicates
	// that more output is available. Set this parameter to the value provided by the
	// previous call's NextToken response to request the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTrustedTokenIssuersOutput struct {

	// If present, this value indicates that more output is available than is included
	// in the current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null . This
	// indicates that this is the last page of results.
	NextToken *string

	// An array list of the trusted token issuer configurations.
	TrustedTokenIssuers []types.TrustedTokenIssuerMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTrustedTokenIssuersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTrustedTokenIssuers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTrustedTokenIssuers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTrustedTokenIssuers"); err != nil {
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
	if err = addOpListTrustedTokenIssuersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTrustedTokenIssuers(options.Region), middleware.Before); err != nil {
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

// ListTrustedTokenIssuersAPIClient is a client that implements the
// ListTrustedTokenIssuers operation.
type ListTrustedTokenIssuersAPIClient interface {
	ListTrustedTokenIssuers(context.Context, *ListTrustedTokenIssuersInput, ...func(*Options)) (*ListTrustedTokenIssuersOutput, error)
}

var _ ListTrustedTokenIssuersAPIClient = (*Client)(nil)

// ListTrustedTokenIssuersPaginatorOptions is the paginator options for
// ListTrustedTokenIssuers
type ListTrustedTokenIssuersPaginatorOptions struct {
	// Specifies the total number of results that you want included in each response.
	// If additional items exist beyond the number you specify, the NextToken response
	// element is returned with a value (not null). Include the specified value as the
	// NextToken request parameter in the next call to the operation to get the next
	// set of results. Note that the service might return fewer results than the
	// maximum even when there are more results available. You should check NextToken
	// after every operation to ensure that you receive all of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTrustedTokenIssuersPaginator is a paginator for ListTrustedTokenIssuers
type ListTrustedTokenIssuersPaginator struct {
	options   ListTrustedTokenIssuersPaginatorOptions
	client    ListTrustedTokenIssuersAPIClient
	params    *ListTrustedTokenIssuersInput
	nextToken *string
	firstPage bool
}

// NewListTrustedTokenIssuersPaginator returns a new
// ListTrustedTokenIssuersPaginator
func NewListTrustedTokenIssuersPaginator(client ListTrustedTokenIssuersAPIClient, params *ListTrustedTokenIssuersInput, optFns ...func(*ListTrustedTokenIssuersPaginatorOptions)) *ListTrustedTokenIssuersPaginator {
	if params == nil {
		params = &ListTrustedTokenIssuersInput{}
	}

	options := ListTrustedTokenIssuersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTrustedTokenIssuersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTrustedTokenIssuersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTrustedTokenIssuers page.
func (p *ListTrustedTokenIssuersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTrustedTokenIssuersOutput, error) {
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

	result, err := p.client.ListTrustedTokenIssuers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTrustedTokenIssuers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTrustedTokenIssuers",
	}
}
