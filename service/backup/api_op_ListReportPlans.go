// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of your report plans. For detailed information about a single
// report plan, use DescribeReportPlan .
func (c *Client) ListReportPlans(ctx context.Context, params *ListReportPlansInput, optFns ...func(*Options)) (*ListReportPlansOutput, error) {
	if params == nil {
		params = &ListReportPlansInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReportPlans", params, optFns, c.addOperationListReportPlansMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReportPlansOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReportPlansInput struct {

	// The number of desired results from 1 to 1000. Optional. If unspecified, the
	// query will return 1 MB of data.
	MaxResults *int32

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	noSmithyDocumentSerde
}

type ListReportPlansOutput struct {

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// A list of your report plans with detailed information for each plan. This
	// information includes the Amazon Resource Name (ARN), report plan name,
	// description, settings, delivery channel, deployment status, creation time, and
	// last times the report plan attempted to and successfully ran.
	ReportPlans []types.ReportPlan

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReportPlansMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListReportPlans{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListReportPlans{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListReportPlans"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReportPlans(options.Region), middleware.Before); err != nil {
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

// ListReportPlansAPIClient is a client that implements the ListReportPlans
// operation.
type ListReportPlansAPIClient interface {
	ListReportPlans(context.Context, *ListReportPlansInput, ...func(*Options)) (*ListReportPlansOutput, error)
}

var _ ListReportPlansAPIClient = (*Client)(nil)

// ListReportPlansPaginatorOptions is the paginator options for ListReportPlans
type ListReportPlansPaginatorOptions struct {
	// The number of desired results from 1 to 1000. Optional. If unspecified, the
	// query will return 1 MB of data.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListReportPlansPaginator is a paginator for ListReportPlans
type ListReportPlansPaginator struct {
	options   ListReportPlansPaginatorOptions
	client    ListReportPlansAPIClient
	params    *ListReportPlansInput
	nextToken *string
	firstPage bool
}

// NewListReportPlansPaginator returns a new ListReportPlansPaginator
func NewListReportPlansPaginator(client ListReportPlansAPIClient, params *ListReportPlansInput, optFns ...func(*ListReportPlansPaginatorOptions)) *ListReportPlansPaginator {
	if params == nil {
		params = &ListReportPlansInput{}
	}

	options := ListReportPlansPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListReportPlansPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListReportPlansPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListReportPlans page.
func (p *ListReportPlansPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListReportPlansOutput, error) {
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

	result, err := p.client.ListReportPlans(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListReportPlans(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListReportPlans",
	}
}
