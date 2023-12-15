// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudcontrol

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns existing resource operation requests. This includes requests of all
// status types. For more information, see Listing active resource operation
// requests (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-manage-requests.html#resource-operations-manage-requests-list)
// in the Amazon Web Services Cloud Control API User Guide. Resource operation
// requests expire after 7 days.
func (c *Client) ListResourceRequests(ctx context.Context, params *ListResourceRequestsInput, optFns ...func(*Options)) (*ListResourceRequestsOutput, error) {
	if params == nil {
		params = &ListResourceRequestsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResourceRequests", params, optFns, c.addOperationListResourceRequestsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourceRequestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourceRequestsInput struct {

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next
	// set of results. The default is 20 .
	MaxResults *int32

	// If the previous paginated request didn't return all of the remaining results,
	// the response object's NextToken parameter value is set to a token. To retrieve
	// the next set of results, call this action again and assign that token to the
	// request object's NextToken parameter. If there are no remaining results, the
	// previous response object's NextToken parameter is set to null .
	NextToken *string

	// The filter criteria to apply to the requests returned.
	ResourceRequestStatusFilter *types.ResourceRequestStatusFilter

	noSmithyDocumentSerde
}

type ListResourceRequestsOutput struct {

	// If the request doesn't return all of the remaining results, NextToken is set to
	// a token. To retrieve the next set of results, call ListResources again and
	// assign that token to the request object's NextToken parameter. If the request
	// returns all results, NextToken is set to null.
	NextToken *string

	// The requests that match the specified filter criteria.
	ResourceRequestStatusSummaries []types.ProgressEvent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResourceRequestsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListResourceRequests{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListResourceRequests{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListResourceRequests"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResourceRequests(options.Region), middleware.Before); err != nil {
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

// ListResourceRequestsAPIClient is a client that implements the
// ListResourceRequests operation.
type ListResourceRequestsAPIClient interface {
	ListResourceRequests(context.Context, *ListResourceRequestsInput, ...func(*Options)) (*ListResourceRequestsOutput, error)
}

var _ ListResourceRequestsAPIClient = (*Client)(nil)

// ListResourceRequestsPaginatorOptions is the paginator options for
// ListResourceRequests
type ListResourceRequestsPaginatorOptions struct {
	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next
	// set of results. The default is 20 .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResourceRequestsPaginator is a paginator for ListResourceRequests
type ListResourceRequestsPaginator struct {
	options   ListResourceRequestsPaginatorOptions
	client    ListResourceRequestsAPIClient
	params    *ListResourceRequestsInput
	nextToken *string
	firstPage bool
}

// NewListResourceRequestsPaginator returns a new ListResourceRequestsPaginator
func NewListResourceRequestsPaginator(client ListResourceRequestsAPIClient, params *ListResourceRequestsInput, optFns ...func(*ListResourceRequestsPaginatorOptions)) *ListResourceRequestsPaginator {
	if params == nil {
		params = &ListResourceRequestsInput{}
	}

	options := ListResourceRequestsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResourceRequestsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResourceRequestsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResourceRequests page.
func (p *ListResourceRequestsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResourceRequestsOutput, error) {
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

	result, err := p.client.ListResourceRequests(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListResourceRequests(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListResourceRequests",
	}
}
