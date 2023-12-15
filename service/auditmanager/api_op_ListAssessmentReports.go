// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of assessment reports created in Audit Manager.
func (c *Client) ListAssessmentReports(ctx context.Context, params *ListAssessmentReportsInput, optFns ...func(*Options)) (*ListAssessmentReportsOutput, error) {
	if params == nil {
		params = &ListAssessmentReportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssessmentReports", params, optFns, c.addOperationListAssessmentReportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssessmentReportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssessmentReportsInput struct {

	// Represents the maximum number of results on a page or for an API request call.
	MaxResults *int32

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAssessmentReportsOutput struct {

	// The list of assessment reports that the ListAssessmentReports API returned.
	AssessmentReports []types.AssessmentReportMetadata

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssessmentReportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAssessmentReports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssessmentReports{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAssessmentReports"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssessmentReports(options.Region), middleware.Before); err != nil {
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

// ListAssessmentReportsAPIClient is a client that implements the
// ListAssessmentReports operation.
type ListAssessmentReportsAPIClient interface {
	ListAssessmentReports(context.Context, *ListAssessmentReportsInput, ...func(*Options)) (*ListAssessmentReportsOutput, error)
}

var _ ListAssessmentReportsAPIClient = (*Client)(nil)

// ListAssessmentReportsPaginatorOptions is the paginator options for
// ListAssessmentReports
type ListAssessmentReportsPaginatorOptions struct {
	// Represents the maximum number of results on a page or for an API request call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssessmentReportsPaginator is a paginator for ListAssessmentReports
type ListAssessmentReportsPaginator struct {
	options   ListAssessmentReportsPaginatorOptions
	client    ListAssessmentReportsAPIClient
	params    *ListAssessmentReportsInput
	nextToken *string
	firstPage bool
}

// NewListAssessmentReportsPaginator returns a new ListAssessmentReportsPaginator
func NewListAssessmentReportsPaginator(client ListAssessmentReportsAPIClient, params *ListAssessmentReportsInput, optFns ...func(*ListAssessmentReportsPaginatorOptions)) *ListAssessmentReportsPaginator {
	if params == nil {
		params = &ListAssessmentReportsInput{}
	}

	options := ListAssessmentReportsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssessmentReportsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssessmentReportsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssessmentReports page.
func (p *ListAssessmentReportsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssessmentReportsOutput, error) {
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

	result, err := p.client.ListAssessmentReports(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssessmentReports(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAssessmentReports",
	}
}
