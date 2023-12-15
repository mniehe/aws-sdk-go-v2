// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the Security Hub administrator accounts. Can only be called by the
// organization management account.
func (c *Client) ListOrganizationAdminAccounts(ctx context.Context, params *ListOrganizationAdminAccountsInput, optFns ...func(*Options)) (*ListOrganizationAdminAccountsOutput, error) {
	if params == nil {
		params = &ListOrganizationAdminAccountsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOrganizationAdminAccounts", params, optFns, c.addOperationListOrganizationAdminAccountsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOrganizationAdminAccountsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOrganizationAdminAccountsInput struct {

	// The maximum number of items to return in the response.
	MaxResults *int32

	// The token that is required for pagination. On your first call to the
	// ListOrganizationAdminAccounts operation, set the value of this parameter to NULL
	// . For subsequent calls to the operation, to continue listing data, set the value
	// of this parameter to the value returned from the previous response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListOrganizationAdminAccountsOutput struct {

	// The list of Security Hub administrator accounts.
	AdminAccounts []types.AdminAccount

	// The pagination token to use to request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOrganizationAdminAccountsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListOrganizationAdminAccounts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListOrganizationAdminAccounts{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListOrganizationAdminAccounts"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOrganizationAdminAccounts(options.Region), middleware.Before); err != nil {
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

// ListOrganizationAdminAccountsAPIClient is a client that implements the
// ListOrganizationAdminAccounts operation.
type ListOrganizationAdminAccountsAPIClient interface {
	ListOrganizationAdminAccounts(context.Context, *ListOrganizationAdminAccountsInput, ...func(*Options)) (*ListOrganizationAdminAccountsOutput, error)
}

var _ ListOrganizationAdminAccountsAPIClient = (*Client)(nil)

// ListOrganizationAdminAccountsPaginatorOptions is the paginator options for
// ListOrganizationAdminAccounts
type ListOrganizationAdminAccountsPaginatorOptions struct {
	// The maximum number of items to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOrganizationAdminAccountsPaginator is a paginator for
// ListOrganizationAdminAccounts
type ListOrganizationAdminAccountsPaginator struct {
	options   ListOrganizationAdminAccountsPaginatorOptions
	client    ListOrganizationAdminAccountsAPIClient
	params    *ListOrganizationAdminAccountsInput
	nextToken *string
	firstPage bool
}

// NewListOrganizationAdminAccountsPaginator returns a new
// ListOrganizationAdminAccountsPaginator
func NewListOrganizationAdminAccountsPaginator(client ListOrganizationAdminAccountsAPIClient, params *ListOrganizationAdminAccountsInput, optFns ...func(*ListOrganizationAdminAccountsPaginatorOptions)) *ListOrganizationAdminAccountsPaginator {
	if params == nil {
		params = &ListOrganizationAdminAccountsInput{}
	}

	options := ListOrganizationAdminAccountsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOrganizationAdminAccountsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOrganizationAdminAccountsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOrganizationAdminAccounts page.
func (p *ListOrganizationAdminAccountsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOrganizationAdminAccountsOutput, error) {
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

	result, err := p.client.ListOrganizationAdminAccounts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListOrganizationAdminAccounts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListOrganizationAdminAccounts",
	}
}
