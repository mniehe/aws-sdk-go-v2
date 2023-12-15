// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/codeguruprofiler/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// List the available reports for a given profiling group and time range.
func (c *Client) ListFindingsReports(ctx context.Context, params *ListFindingsReportsInput, optFns ...func(*Options)) (*ListFindingsReportsOutput, error) {
	if params == nil {
		params = &ListFindingsReportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFindingsReports", params, optFns, c.addOperationListFindingsReportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFindingsReportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the ListFindingsReportsRequest.
type ListFindingsReportsInput struct {

	// The end time of the profile to get analysis data about. You must specify
	// startTime and endTime . This is specified using the ISO 8601 format. For
	// example, 2020-06-01T13:15:02.001Z represents 1 millisecond past June 1, 2020
	// 1:15:02 PM UTC.
	//
	// This member is required.
	EndTime *time.Time

	// The name of the profiling group from which to search for analysis data.
	//
	// This member is required.
	ProfilingGroupName *string

	// The start time of the profile to get analysis data about. You must specify
	// startTime and endTime . This is specified using the ISO 8601 format. For
	// example, 2020-06-01T13:15:02.001Z represents 1 millisecond past June 1, 2020
	// 1:15:02 PM UTC.
	//
	// This member is required.
	StartTime *time.Time

	// A Boolean value indicating whether to only return reports from daily profiles.
	// If set to True , only analysis data from daily profiles is returned. If set to
	// False , analysis data is returned from smaller time windows (for example, one
	// hour).
	DailyReportsOnly *bool

	// The maximum number of report results returned by ListFindingsReports in
	// paginated output. When this parameter is used, ListFindingsReports only returns
	// maxResults results in a single page along with a nextToken response element.
	// The remaining results of the initial request can be seen by sending another
	// ListFindingsReports request with the returned nextToken value.
	MaxResults *int32

	// The nextToken value returned from a previous paginated
	// ListFindingsReportsRequest request where maxResults was used and the results
	// exceeded the value of that parameter. Pagination continues from the end of the
	// previous results that returned the nextToken value. This token should be
	// treated as an opaque identifier that is only used to retrieve the next items in
	// a list and not for other programmatic purposes.
	NextToken *string

	noSmithyDocumentSerde
}

// The structure representing the ListFindingsReportsResponse.
type ListFindingsReportsOutput struct {

	// The list of analysis results summaries.
	//
	// This member is required.
	FindingsReportSummaries []types.FindingsReportSummary

	// The nextToken value to include in a future ListFindingsReports request. When
	// the results of a ListFindingsReports request exceed maxResults , this value can
	// be used to retrieve the next page of results. This value is null when there are
	// no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFindingsReportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFindingsReports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFindingsReports{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFindingsReports"); err != nil {
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
	if err = addOpListFindingsReportsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFindingsReports(options.Region), middleware.Before); err != nil {
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

// ListFindingsReportsAPIClient is a client that implements the
// ListFindingsReports operation.
type ListFindingsReportsAPIClient interface {
	ListFindingsReports(context.Context, *ListFindingsReportsInput, ...func(*Options)) (*ListFindingsReportsOutput, error)
}

var _ ListFindingsReportsAPIClient = (*Client)(nil)

// ListFindingsReportsPaginatorOptions is the paginator options for
// ListFindingsReports
type ListFindingsReportsPaginatorOptions struct {
	// The maximum number of report results returned by ListFindingsReports in
	// paginated output. When this parameter is used, ListFindingsReports only returns
	// maxResults results in a single page along with a nextToken response element.
	// The remaining results of the initial request can be seen by sending another
	// ListFindingsReports request with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFindingsReportsPaginator is a paginator for ListFindingsReports
type ListFindingsReportsPaginator struct {
	options   ListFindingsReportsPaginatorOptions
	client    ListFindingsReportsAPIClient
	params    *ListFindingsReportsInput
	nextToken *string
	firstPage bool
}

// NewListFindingsReportsPaginator returns a new ListFindingsReportsPaginator
func NewListFindingsReportsPaginator(client ListFindingsReportsAPIClient, params *ListFindingsReportsInput, optFns ...func(*ListFindingsReportsPaginatorOptions)) *ListFindingsReportsPaginator {
	if params == nil {
		params = &ListFindingsReportsInput{}
	}

	options := ListFindingsReportsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFindingsReportsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFindingsReportsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFindingsReports page.
func (p *ListFindingsReportsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFindingsReportsOutput, error) {
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

	result, err := p.client.ListFindingsReports(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFindingsReports(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFindingsReports",
	}
}
