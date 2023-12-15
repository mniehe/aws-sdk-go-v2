// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all historical purchases, renewals, and system renewal
// transactions for an AWS account. The list is paginated and ordered by a
// descending timestamp (most recent transactions are first). The API returns a
// NotEligible error if the user is not permitted to invoke the operation. If you
// must be able to invoke this operation, contact aws-devicefarm-support@amazon.com (mailto:aws-devicefarm-support@amazon.com)
// .
func (c *Client) ListOfferingTransactions(ctx context.Context, params *ListOfferingTransactionsInput, optFns ...func(*Options)) (*ListOfferingTransactionsOutput, error) {
	if params == nil {
		params = &ListOfferingTransactionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOfferingTransactions", params, optFns, c.addOperationListOfferingTransactionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOfferingTransactionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to list the offering transaction history.
type ListOfferingTransactionsInput struct {

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	noSmithyDocumentSerde
}

// Returns the transaction log of the specified offerings.
type ListOfferingTransactionsOutput struct {

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// The audit log of subscriptions you have purchased and modified through AWS
	// Device Farm.
	OfferingTransactions []types.OfferingTransaction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOfferingTransactionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListOfferingTransactions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListOfferingTransactions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListOfferingTransactions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOfferingTransactions(options.Region), middleware.Before); err != nil {
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

// ListOfferingTransactionsAPIClient is a client that implements the
// ListOfferingTransactions operation.
type ListOfferingTransactionsAPIClient interface {
	ListOfferingTransactions(context.Context, *ListOfferingTransactionsInput, ...func(*Options)) (*ListOfferingTransactionsOutput, error)
}

var _ ListOfferingTransactionsAPIClient = (*Client)(nil)

// ListOfferingTransactionsPaginatorOptions is the paginator options for
// ListOfferingTransactions
type ListOfferingTransactionsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOfferingTransactionsPaginator is a paginator for ListOfferingTransactions
type ListOfferingTransactionsPaginator struct {
	options   ListOfferingTransactionsPaginatorOptions
	client    ListOfferingTransactionsAPIClient
	params    *ListOfferingTransactionsInput
	nextToken *string
	firstPage bool
}

// NewListOfferingTransactionsPaginator returns a new
// ListOfferingTransactionsPaginator
func NewListOfferingTransactionsPaginator(client ListOfferingTransactionsAPIClient, params *ListOfferingTransactionsInput, optFns ...func(*ListOfferingTransactionsPaginatorOptions)) *ListOfferingTransactionsPaginator {
	if params == nil {
		params = &ListOfferingTransactionsInput{}
	}

	options := ListOfferingTransactionsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOfferingTransactionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOfferingTransactionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOfferingTransactions page.
func (p *ListOfferingTransactionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOfferingTransactionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListOfferingTransactions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListOfferingTransactions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListOfferingTransactions",
	}
}
