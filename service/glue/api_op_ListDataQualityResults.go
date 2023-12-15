// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns all data quality execution results for your account.
func (c *Client) ListDataQualityResults(ctx context.Context, params *ListDataQualityResultsInput, optFns ...func(*Options)) (*ListDataQualityResultsOutput, error) {
	if params == nil {
		params = &ListDataQualityResultsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataQualityResults", params, optFns, c.addOperationListDataQualityResultsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataQualityResultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataQualityResultsInput struct {

	// The filter criteria.
	Filter *types.DataQualityResultFilterCriteria

	// The maximum number of results to return.
	MaxResults *int32

	// A paginated token to offset the results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDataQualityResultsOutput struct {

	// A list of DataQualityResultDescription objects.
	//
	// This member is required.
	Results []types.DataQualityResultDescription

	// A pagination token, if more results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDataQualityResultsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDataQualityResults{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDataQualityResults{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDataQualityResults"); err != nil {
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
	if err = addOpListDataQualityResultsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDataQualityResults(options.Region), middleware.Before); err != nil {
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

// ListDataQualityResultsAPIClient is a client that implements the
// ListDataQualityResults operation.
type ListDataQualityResultsAPIClient interface {
	ListDataQualityResults(context.Context, *ListDataQualityResultsInput, ...func(*Options)) (*ListDataQualityResultsOutput, error)
}

var _ ListDataQualityResultsAPIClient = (*Client)(nil)

// ListDataQualityResultsPaginatorOptions is the paginator options for
// ListDataQualityResults
type ListDataQualityResultsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDataQualityResultsPaginator is a paginator for ListDataQualityResults
type ListDataQualityResultsPaginator struct {
	options   ListDataQualityResultsPaginatorOptions
	client    ListDataQualityResultsAPIClient
	params    *ListDataQualityResultsInput
	nextToken *string
	firstPage bool
}

// NewListDataQualityResultsPaginator returns a new ListDataQualityResultsPaginator
func NewListDataQualityResultsPaginator(client ListDataQualityResultsAPIClient, params *ListDataQualityResultsInput, optFns ...func(*ListDataQualityResultsPaginatorOptions)) *ListDataQualityResultsPaginator {
	if params == nil {
		params = &ListDataQualityResultsInput{}
	}

	options := ListDataQualityResultsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDataQualityResultsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDataQualityResultsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDataQualityResults page.
func (p *ListDataQualityResultsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDataQualityResultsOutput, error) {
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

	result, err := p.client.ListDataQualityResults(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDataQualityResults(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDataQualityResults",
	}
}
