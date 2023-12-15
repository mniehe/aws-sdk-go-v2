// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the monitoring evaluation results and predictor events
// collected by the monitor resource during different windows of time. For
// information about monitoring see predictor-monitoring . For more information
// about retrieving monitoring results see Viewing Monitoring Results (https://docs.aws.amazon.com/forecast/latest/dg/predictor-monitoring-results.html)
// .
func (c *Client) ListMonitorEvaluations(ctx context.Context, params *ListMonitorEvaluationsInput, optFns ...func(*Options)) (*ListMonitorEvaluationsOutput, error) {
	if params == nil {
		params = &ListMonitorEvaluationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMonitorEvaluations", params, optFns, c.addOperationListMonitorEvaluationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMonitorEvaluationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMonitorEvaluationsInput struct {

	// The Amazon Resource Name (ARN) of the monitor resource to get results from.
	//
	// This member is required.
	MonitorArn *string

	// An array of filters. For each filter, provide a condition and a match
	// statement. The condition is either IS or IS_NOT , which specifies whether to
	// include or exclude the resources that match the statement from the list. The
	// match statement consists of a key and a value. Filter properties
	//   - Condition - The condition to apply. Valid values are IS and IS_NOT .
	//   - Key - The name of the parameter to filter on. The only valid value is
	//   EvaluationState .
	//   - Value - The value to match. Valid values are only SUCCESS or FAILURE .
	// For example, to list only successful monitor evaluations, you would specify:
	// "Filters": [ { "Condition": "IS", "Key": "EvaluationState", "Value": "SUCCESS" }
	// ]
	Filters []types.Filter

	// The maximum number of monitoring results to return.
	MaxResults *int32

	// If the result of the previous request was truncated, the response includes a
	// NextToken . To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMonitorEvaluationsOutput struct {

	// If the response is truncated, Amazon Forecast returns this token. To retrieve
	// the next set of results, use the token in the next request. Tokens expire after
	// 24 hours.
	NextToken *string

	// The monitoring results and predictor events collected by the monitor resource
	// during different windows of time. For information about monitoring see Viewing
	// Monitoring Results (https://docs.aws.amazon.com/forecast/latest/dg/predictor-monitoring-results.html)
	// . For more information about retrieving monitoring results see Viewing
	// Monitoring Results (https://docs.aws.amazon.com/forecast/latest/dg/predictor-monitoring-results.html)
	// .
	PredictorMonitorEvaluations []types.PredictorMonitorEvaluation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMonitorEvaluationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListMonitorEvaluations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMonitorEvaluations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMonitorEvaluations"); err != nil {
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
	if err = addOpListMonitorEvaluationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMonitorEvaluations(options.Region), middleware.Before); err != nil {
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

// ListMonitorEvaluationsAPIClient is a client that implements the
// ListMonitorEvaluations operation.
type ListMonitorEvaluationsAPIClient interface {
	ListMonitorEvaluations(context.Context, *ListMonitorEvaluationsInput, ...func(*Options)) (*ListMonitorEvaluationsOutput, error)
}

var _ ListMonitorEvaluationsAPIClient = (*Client)(nil)

// ListMonitorEvaluationsPaginatorOptions is the paginator options for
// ListMonitorEvaluations
type ListMonitorEvaluationsPaginatorOptions struct {
	// The maximum number of monitoring results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMonitorEvaluationsPaginator is a paginator for ListMonitorEvaluations
type ListMonitorEvaluationsPaginator struct {
	options   ListMonitorEvaluationsPaginatorOptions
	client    ListMonitorEvaluationsAPIClient
	params    *ListMonitorEvaluationsInput
	nextToken *string
	firstPage bool
}

// NewListMonitorEvaluationsPaginator returns a new ListMonitorEvaluationsPaginator
func NewListMonitorEvaluationsPaginator(client ListMonitorEvaluationsAPIClient, params *ListMonitorEvaluationsInput, optFns ...func(*ListMonitorEvaluationsPaginatorOptions)) *ListMonitorEvaluationsPaginator {
	if params == nil {
		params = &ListMonitorEvaluationsInput{}
	}

	options := ListMonitorEvaluationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMonitorEvaluationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMonitorEvaluationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMonitorEvaluations page.
func (p *ListMonitorEvaluationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMonitorEvaluationsOutput, error) {
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

	result, err := p.client.ListMonitorEvaluations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMonitorEvaluations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMonitorEvaluations",
	}
}
