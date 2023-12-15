// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists the trial components in your account. You can sort the list by trial
// component name or creation time. You can filter the list to show only components
// that were created in a specific time range. You can also filter on one of the
// following:
//   - ExperimentName
//   - SourceArn
//   - TrialName
func (c *Client) ListTrialComponents(ctx context.Context, params *ListTrialComponentsInput, optFns ...func(*Options)) (*ListTrialComponentsOutput, error) {
	if params == nil {
		params = &ListTrialComponentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTrialComponents", params, optFns, c.addOperationListTrialComponentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTrialComponentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTrialComponentsInput struct {

	// A filter that returns only components created after the specified time.
	CreatedAfter *time.Time

	// A filter that returns only components created before the specified time.
	CreatedBefore *time.Time

	// A filter that returns only components that are part of the specified
	// experiment. If you specify ExperimentName , you can't filter by SourceArn or
	// TrialName .
	ExperimentName *string

	// The maximum number of components to return in the response. The default value
	// is 10.
	MaxResults *int32

	// If the previous call to ListTrialComponents didn't return the full set of
	// components, the call returns a token for getting the next set of components.
	NextToken *string

	// The property used to sort results. The default value is CreationTime .
	SortBy types.SortTrialComponentsBy

	// The sort order. The default value is Descending .
	SortOrder types.SortOrder

	// A filter that returns only components that have the specified source Amazon
	// Resource Name (ARN). If you specify SourceArn , you can't filter by
	// ExperimentName or TrialName .
	SourceArn *string

	// A filter that returns only components that are part of the specified trial. If
	// you specify TrialName , you can't filter by ExperimentName or SourceArn .
	TrialName *string

	noSmithyDocumentSerde
}

type ListTrialComponentsOutput struct {

	// A token for getting the next set of components, if there are any.
	NextToken *string

	// A list of the summaries of your trial components.
	TrialComponentSummaries []types.TrialComponentSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTrialComponentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTrialComponents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTrialComponents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTrialComponents"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTrialComponents(options.Region), middleware.Before); err != nil {
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

// ListTrialComponentsAPIClient is a client that implements the
// ListTrialComponents operation.
type ListTrialComponentsAPIClient interface {
	ListTrialComponents(context.Context, *ListTrialComponentsInput, ...func(*Options)) (*ListTrialComponentsOutput, error)
}

var _ ListTrialComponentsAPIClient = (*Client)(nil)

// ListTrialComponentsPaginatorOptions is the paginator options for
// ListTrialComponents
type ListTrialComponentsPaginatorOptions struct {
	// The maximum number of components to return in the response. The default value
	// is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTrialComponentsPaginator is a paginator for ListTrialComponents
type ListTrialComponentsPaginator struct {
	options   ListTrialComponentsPaginatorOptions
	client    ListTrialComponentsAPIClient
	params    *ListTrialComponentsInput
	nextToken *string
	firstPage bool
}

// NewListTrialComponentsPaginator returns a new ListTrialComponentsPaginator
func NewListTrialComponentsPaginator(client ListTrialComponentsAPIClient, params *ListTrialComponentsInput, optFns ...func(*ListTrialComponentsPaginatorOptions)) *ListTrialComponentsPaginator {
	if params == nil {
		params = &ListTrialComponentsInput{}
	}

	options := ListTrialComponentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTrialComponentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTrialComponentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTrialComponents page.
func (p *ListTrialComponentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTrialComponentsOutput, error) {
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

	result, err := p.client.ListTrialComponents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTrialComponents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTrialComponents",
	}
}
