// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Show a list of the predictive inbox placement tests that you've performed,
// regardless of their statuses. For predictive inbox placement tests that are
// complete, you can use the GetDeliverabilityTestReport operation to view the
// results.
func (c *Client) ListDeliverabilityTestReports(ctx context.Context, params *ListDeliverabilityTestReportsInput, optFns ...func(*Options)) (*ListDeliverabilityTestReportsOutput, error) {
	if params == nil {
		params = &ListDeliverabilityTestReportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDeliverabilityTestReports", params, optFns, addOperationListDeliverabilityTestReportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDeliverabilityTestReportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to list all of the predictive inbox placement tests that you've
// performed.
type ListDeliverabilityTestReportsInput struct {

	// A token returned from a previous call to ListDeliverabilityTestReports to
	// indicate the position in the list of predictive inbox placement tests.
	NextToken *string

	// The number of results to show in a single call to ListDeliverabilityTestReports.
	// If the number of results is larger than the number you specified in this
	// parameter, then the response includes a NextToken element, which you can use to
	// obtain additional results. The value you specify has to be at least 0, and can
	// be no more than 1000.
	PageSize *int32
}

// A list of the predictive inbox placement test reports that are available for
// your account, regardless of whether or not those tests are complete.
type ListDeliverabilityTestReportsOutput struct {

	// An object that contains a lists of predictive inbox placement tests that you've
	// performed.
	//
	// This member is required.
	DeliverabilityTestReports []types.DeliverabilityTestReport

	// A token that indicates that there are additional predictive inbox placement
	// tests to list. To view additional predictive inbox placement tests, issue
	// another request to ListDeliverabilityTestReports, and pass this token in the
	// NextToken parameter.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDeliverabilityTestReportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDeliverabilityTestReports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDeliverabilityTestReports{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDeliverabilityTestReports(options.Region), middleware.Before); err != nil {
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
	return nil
}

// ListDeliverabilityTestReportsAPIClient is a client that implements the
// ListDeliverabilityTestReports operation.
type ListDeliverabilityTestReportsAPIClient interface {
	ListDeliverabilityTestReports(context.Context, *ListDeliverabilityTestReportsInput, ...func(*Options)) (*ListDeliverabilityTestReportsOutput, error)
}

var _ ListDeliverabilityTestReportsAPIClient = (*Client)(nil)

// ListDeliverabilityTestReportsPaginatorOptions is the paginator options for
// ListDeliverabilityTestReports
type ListDeliverabilityTestReportsPaginatorOptions struct {
	// The number of results to show in a single call to ListDeliverabilityTestReports.
	// If the number of results is larger than the number you specified in this
	// parameter, then the response includes a NextToken element, which you can use to
	// obtain additional results. The value you specify has to be at least 0, and can
	// be no more than 1000.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDeliverabilityTestReportsPaginator is a paginator for
// ListDeliverabilityTestReports
type ListDeliverabilityTestReportsPaginator struct {
	options   ListDeliverabilityTestReportsPaginatorOptions
	client    ListDeliverabilityTestReportsAPIClient
	params    *ListDeliverabilityTestReportsInput
	nextToken *string
	firstPage bool
}

// NewListDeliverabilityTestReportsPaginator returns a new
// ListDeliverabilityTestReportsPaginator
func NewListDeliverabilityTestReportsPaginator(client ListDeliverabilityTestReportsAPIClient, params *ListDeliverabilityTestReportsInput, optFns ...func(*ListDeliverabilityTestReportsPaginatorOptions)) *ListDeliverabilityTestReportsPaginator {
	options := ListDeliverabilityTestReportsPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListDeliverabilityTestReportsInput{}
	}

	return &ListDeliverabilityTestReportsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDeliverabilityTestReportsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListDeliverabilityTestReports page.
func (p *ListDeliverabilityTestReportsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDeliverabilityTestReportsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.PageSize = limit

	result, err := p.client.ListDeliverabilityTestReports(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListDeliverabilityTestReports(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListDeliverabilityTestReports",
	}
}
