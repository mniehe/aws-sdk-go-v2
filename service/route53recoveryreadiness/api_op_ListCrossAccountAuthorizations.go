// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the cross-account readiness authorizations that are in place for an
// account.
func (c *Client) ListCrossAccountAuthorizations(ctx context.Context, params *ListCrossAccountAuthorizationsInput, optFns ...func(*Options)) (*ListCrossAccountAuthorizationsOutput, error) {
	if params == nil {
		params = &ListCrossAccountAuthorizationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCrossAccountAuthorizations", params, optFns, c.addOperationListCrossAccountAuthorizationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCrossAccountAuthorizationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCrossAccountAuthorizationsInput struct {

	// The number of objects that you want to return with this call.
	MaxResults *int32

	// The token that identifies which batch of results you want to see.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCrossAccountAuthorizationsOutput struct {

	// A list of cross-account authorizations.
	CrossAccountAuthorizations []string

	// The token that identifies which batch of results you want to see.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCrossAccountAuthorizationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCrossAccountAuthorizations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCrossAccountAuthorizations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCrossAccountAuthorizations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCrossAccountAuthorizations(options.Region), middleware.Before); err != nil {
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

// ListCrossAccountAuthorizationsAPIClient is a client that implements the
// ListCrossAccountAuthorizations operation.
type ListCrossAccountAuthorizationsAPIClient interface {
	ListCrossAccountAuthorizations(context.Context, *ListCrossAccountAuthorizationsInput, ...func(*Options)) (*ListCrossAccountAuthorizationsOutput, error)
}

var _ ListCrossAccountAuthorizationsAPIClient = (*Client)(nil)

// ListCrossAccountAuthorizationsPaginatorOptions is the paginator options for
// ListCrossAccountAuthorizations
type ListCrossAccountAuthorizationsPaginatorOptions struct {
	// The number of objects that you want to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCrossAccountAuthorizationsPaginator is a paginator for
// ListCrossAccountAuthorizations
type ListCrossAccountAuthorizationsPaginator struct {
	options   ListCrossAccountAuthorizationsPaginatorOptions
	client    ListCrossAccountAuthorizationsAPIClient
	params    *ListCrossAccountAuthorizationsInput
	nextToken *string
	firstPage bool
}

// NewListCrossAccountAuthorizationsPaginator returns a new
// ListCrossAccountAuthorizationsPaginator
func NewListCrossAccountAuthorizationsPaginator(client ListCrossAccountAuthorizationsAPIClient, params *ListCrossAccountAuthorizationsInput, optFns ...func(*ListCrossAccountAuthorizationsPaginatorOptions)) *ListCrossAccountAuthorizationsPaginator {
	if params == nil {
		params = &ListCrossAccountAuthorizationsInput{}
	}

	options := ListCrossAccountAuthorizationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCrossAccountAuthorizationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCrossAccountAuthorizationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCrossAccountAuthorizations page.
func (p *ListCrossAccountAuthorizationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCrossAccountAuthorizationsOutput, error) {
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

	result, err := p.client.ListCrossAccountAuthorizations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCrossAccountAuthorizations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCrossAccountAuthorizations",
	}
}
