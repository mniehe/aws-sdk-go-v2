// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/mturk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The ListReviewableHITs operation retrieves the HITs with Status equal to
// Reviewable or Status equal to Reviewing that belong to the Requester calling the
// operation.
func (c *Client) ListReviewableHITs(ctx context.Context, params *ListReviewableHITsInput, optFns ...func(*Options)) (*ListReviewableHITsOutput, error) {
	if params == nil {
		params = &ListReviewableHITsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReviewableHITs", params, optFns, c.addOperationListReviewableHITsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReviewableHITsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReviewableHITsInput struct {

	// The ID of the HIT type of the HITs to consider for the query. If not specified,
	// all HITs for the Reviewer are considered
	HITTypeId *string

	// Limit the number of results returned.
	MaxResults *int32

	// Pagination Token
	NextToken *string

	// Can be either Reviewable or Reviewing . Reviewable is the default value.
	Status types.ReviewableHITStatus

	noSmithyDocumentSerde
}

type ListReviewableHITsOutput struct {

	// The list of HIT elements returned by the query.
	HITs []types.HIT

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// The number of HITs on this page in the filtered results list, equivalent to the
	// number of HITs being returned by this call.
	NumResults *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReviewableHITsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListReviewableHITs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListReviewableHITs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListReviewableHITs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReviewableHITs(options.Region), middleware.Before); err != nil {
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

// ListReviewableHITsAPIClient is a client that implements the ListReviewableHITs
// operation.
type ListReviewableHITsAPIClient interface {
	ListReviewableHITs(context.Context, *ListReviewableHITsInput, ...func(*Options)) (*ListReviewableHITsOutput, error)
}

var _ ListReviewableHITsAPIClient = (*Client)(nil)

// ListReviewableHITsPaginatorOptions is the paginator options for
// ListReviewableHITs
type ListReviewableHITsPaginatorOptions struct {
	// Limit the number of results returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListReviewableHITsPaginator is a paginator for ListReviewableHITs
type ListReviewableHITsPaginator struct {
	options   ListReviewableHITsPaginatorOptions
	client    ListReviewableHITsAPIClient
	params    *ListReviewableHITsInput
	nextToken *string
	firstPage bool
}

// NewListReviewableHITsPaginator returns a new ListReviewableHITsPaginator
func NewListReviewableHITsPaginator(client ListReviewableHITsAPIClient, params *ListReviewableHITsInput, optFns ...func(*ListReviewableHITsPaginatorOptions)) *ListReviewableHITsPaginator {
	if params == nil {
		params = &ListReviewableHITsInput{}
	}

	options := ListReviewableHITsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListReviewableHITsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListReviewableHITsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListReviewableHITs page.
func (p *ListReviewableHITsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListReviewableHITsOutput, error) {
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

	result, err := p.client.ListReviewableHITs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListReviewableHITs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListReviewableHITs",
	}
}
