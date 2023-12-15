// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of FAQ lists associated with an index.
func (c *Client) ListFaqs(ctx context.Context, params *ListFaqsInput, optFns ...func(*Options)) (*ListFaqsOutput, error) {
	if params == nil {
		params = &ListFaqsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFaqs", params, optFns, c.addOperationListFaqsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFaqsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFaqsInput struct {

	// The index that contains the FAQ lists.
	//
	// This member is required.
	IndexId *string

	// The maximum number of FAQs to return in the response. If there are fewer
	// results in the list, this response contains only the actual results.
	MaxResults *int32

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Kendra returns a pagination token in the response. You can use
	// this pagination token to retrieve the next set of FAQs.
	NextToken *string

	noSmithyDocumentSerde
}

type ListFaqsOutput struct {

	// information about the FAQs associated with the specified index.
	FaqSummaryItems []types.FaqSummary

	// If the response is truncated, Amazon Kendra returns this token that you can use
	// in the subsequent request to retrieve the next set of FAQs.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFaqsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFaqs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFaqs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFaqs"); err != nil {
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
	if err = addOpListFaqsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFaqs(options.Region), middleware.Before); err != nil {
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

// ListFaqsAPIClient is a client that implements the ListFaqs operation.
type ListFaqsAPIClient interface {
	ListFaqs(context.Context, *ListFaqsInput, ...func(*Options)) (*ListFaqsOutput, error)
}

var _ ListFaqsAPIClient = (*Client)(nil)

// ListFaqsPaginatorOptions is the paginator options for ListFaqs
type ListFaqsPaginatorOptions struct {
	// The maximum number of FAQs to return in the response. If there are fewer
	// results in the list, this response contains only the actual results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFaqsPaginator is a paginator for ListFaqs
type ListFaqsPaginator struct {
	options   ListFaqsPaginatorOptions
	client    ListFaqsAPIClient
	params    *ListFaqsInput
	nextToken *string
	firstPage bool
}

// NewListFaqsPaginator returns a new ListFaqsPaginator
func NewListFaqsPaginator(client ListFaqsAPIClient, params *ListFaqsInput, optFns ...func(*ListFaqsPaginatorOptions)) *ListFaqsPaginator {
	if params == nil {
		params = &ListFaqsInput{}
	}

	options := ListFaqsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFaqsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFaqsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFaqs page.
func (p *ListFaqsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFaqsOutput, error) {
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

	result, err := p.client.ListFaqs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFaqs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFaqs",
	}
}
