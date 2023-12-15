// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a summary count of compliant and non-compliant resources for a
// compliance type. For example, this call can return State Manager associations,
// patches, or custom compliance types according to the filter criteria that you
// specify.
func (c *Client) ListComplianceSummaries(ctx context.Context, params *ListComplianceSummariesInput, optFns ...func(*Options)) (*ListComplianceSummariesOutput, error) {
	if params == nil {
		params = &ListComplianceSummariesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListComplianceSummaries", params, optFns, c.addOperationListComplianceSummariesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListComplianceSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListComplianceSummariesInput struct {

	// One or more compliance or inventory filters. Use a filter to return a more
	// specific list of results.
	Filters []types.ComplianceStringFilter

	// The maximum number of items to return for this call. Currently, you can specify
	// null or 50. The call also returns a token that you can specify in a subsequent
	// call to get the next set of results.
	MaxResults *int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListComplianceSummariesOutput struct {

	// A list of compliant and non-compliant summary counts based on compliance types.
	// For example, this call returns State Manager associations, patches, or custom
	// compliance types according to the filter criteria that you specified.
	ComplianceSummaryItems []types.ComplianceSummaryItem

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListComplianceSummariesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListComplianceSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListComplianceSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListComplianceSummaries"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListComplianceSummaries(options.Region), middleware.Before); err != nil {
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

// ListComplianceSummariesAPIClient is a client that implements the
// ListComplianceSummaries operation.
type ListComplianceSummariesAPIClient interface {
	ListComplianceSummaries(context.Context, *ListComplianceSummariesInput, ...func(*Options)) (*ListComplianceSummariesOutput, error)
}

var _ ListComplianceSummariesAPIClient = (*Client)(nil)

// ListComplianceSummariesPaginatorOptions is the paginator options for
// ListComplianceSummaries
type ListComplianceSummariesPaginatorOptions struct {
	// The maximum number of items to return for this call. Currently, you can specify
	// null or 50. The call also returns a token that you can specify in a subsequent
	// call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListComplianceSummariesPaginator is a paginator for ListComplianceSummaries
type ListComplianceSummariesPaginator struct {
	options   ListComplianceSummariesPaginatorOptions
	client    ListComplianceSummariesAPIClient
	params    *ListComplianceSummariesInput
	nextToken *string
	firstPage bool
}

// NewListComplianceSummariesPaginator returns a new
// ListComplianceSummariesPaginator
func NewListComplianceSummariesPaginator(client ListComplianceSummariesAPIClient, params *ListComplianceSummariesInput, optFns ...func(*ListComplianceSummariesPaginatorOptions)) *ListComplianceSummariesPaginator {
	if params == nil {
		params = &ListComplianceSummariesInput{}
	}

	options := ListComplianceSummariesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListComplianceSummariesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListComplianceSummariesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListComplianceSummaries page.
func (p *ListComplianceSummariesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListComplianceSummariesOutput, error) {
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

	result, err := p.client.ListComplianceSummaries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListComplianceSummaries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListComplianceSummaries",
	}
}
