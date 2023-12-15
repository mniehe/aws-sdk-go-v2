// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the phone numbers for an administrator's Amazon Chime SDK account.
func (c *Client) ListPhoneNumberOrders(ctx context.Context, params *ListPhoneNumberOrdersInput, optFns ...func(*Options)) (*ListPhoneNumberOrdersOutput, error) {
	if params == nil {
		params = &ListPhoneNumberOrdersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPhoneNumberOrders", params, optFns, c.addOperationListPhoneNumberOrdersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPhoneNumberOrdersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPhoneNumberOrdersInput struct {

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token used to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPhoneNumberOrdersOutput struct {

	// The token used to retrieve the next page of results.
	NextToken *string

	// The phone number order details.
	PhoneNumberOrders []types.PhoneNumberOrder

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPhoneNumberOrdersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPhoneNumberOrders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPhoneNumberOrders{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPhoneNumberOrders"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPhoneNumberOrders(options.Region), middleware.Before); err != nil {
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

// ListPhoneNumberOrdersAPIClient is a client that implements the
// ListPhoneNumberOrders operation.
type ListPhoneNumberOrdersAPIClient interface {
	ListPhoneNumberOrders(context.Context, *ListPhoneNumberOrdersInput, ...func(*Options)) (*ListPhoneNumberOrdersOutput, error)
}

var _ ListPhoneNumberOrdersAPIClient = (*Client)(nil)

// ListPhoneNumberOrdersPaginatorOptions is the paginator options for
// ListPhoneNumberOrders
type ListPhoneNumberOrdersPaginatorOptions struct {
	// The maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPhoneNumberOrdersPaginator is a paginator for ListPhoneNumberOrders
type ListPhoneNumberOrdersPaginator struct {
	options   ListPhoneNumberOrdersPaginatorOptions
	client    ListPhoneNumberOrdersAPIClient
	params    *ListPhoneNumberOrdersInput
	nextToken *string
	firstPage bool
}

// NewListPhoneNumberOrdersPaginator returns a new ListPhoneNumberOrdersPaginator
func NewListPhoneNumberOrdersPaginator(client ListPhoneNumberOrdersAPIClient, params *ListPhoneNumberOrdersInput, optFns ...func(*ListPhoneNumberOrdersPaginatorOptions)) *ListPhoneNumberOrdersPaginator {
	if params == nil {
		params = &ListPhoneNumberOrdersInput{}
	}

	options := ListPhoneNumberOrdersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPhoneNumberOrdersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPhoneNumberOrdersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPhoneNumberOrders page.
func (p *ListPhoneNumberOrdersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPhoneNumberOrdersOutput, error) {
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

	result, err := p.client.ListPhoneNumberOrders(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPhoneNumberOrders(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPhoneNumberOrders",
	}
}
