// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmincidents

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/ssmincidents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all incident records in your account. Use this command to retrieve the
// Amazon Resource Name (ARN) of the incident record you want to update.
func (c *Client) ListIncidentRecords(ctx context.Context, params *ListIncidentRecordsInput, optFns ...func(*Options)) (*ListIncidentRecordsOutput, error) {
	if params == nil {
		params = &ListIncidentRecordsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIncidentRecords", params, optFns, c.addOperationListIncidentRecordsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIncidentRecordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIncidentRecordsInput struct {

	// Filters the list of incident records you want to search through. You can filter
	// on the following keys:
	//   - creationTime
	//   - impact
	//   - status
	//   - createdBy
	// Note the following when when you use Filters:
	//   - If you don't specify a Filter, the response includes all incident records.
	//   - If you specify more than one filter in a single request, the response
	//   returns incident records that match all filters.
	//   - If you specify a filter with more than one value, the response returns
	//   incident records that match any of the values provided.
	Filters []types.Filter

	// The maximum number of results per page.
	MaxResults *int32

	// The pagination token for the next set of items to return. (You received this
	// token from a previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type ListIncidentRecordsOutput struct {

	// The details of each listed incident record.
	//
	// This member is required.
	IncidentRecordSummaries []types.IncidentRecordSummary

	// The pagination token to use when requesting the next set of items. If there are
	// no additional items to return, the string is null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListIncidentRecordsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListIncidentRecords{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListIncidentRecords{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListIncidentRecords"); err != nil {
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
	if err = addOpListIncidentRecordsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListIncidentRecords(options.Region), middleware.Before); err != nil {
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

// ListIncidentRecordsAPIClient is a client that implements the
// ListIncidentRecords operation.
type ListIncidentRecordsAPIClient interface {
	ListIncidentRecords(context.Context, *ListIncidentRecordsInput, ...func(*Options)) (*ListIncidentRecordsOutput, error)
}

var _ ListIncidentRecordsAPIClient = (*Client)(nil)

// ListIncidentRecordsPaginatorOptions is the paginator options for
// ListIncidentRecords
type ListIncidentRecordsPaginatorOptions struct {
	// The maximum number of results per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListIncidentRecordsPaginator is a paginator for ListIncidentRecords
type ListIncidentRecordsPaginator struct {
	options   ListIncidentRecordsPaginatorOptions
	client    ListIncidentRecordsAPIClient
	params    *ListIncidentRecordsInput
	nextToken *string
	firstPage bool
}

// NewListIncidentRecordsPaginator returns a new ListIncidentRecordsPaginator
func NewListIncidentRecordsPaginator(client ListIncidentRecordsAPIClient, params *ListIncidentRecordsInput, optFns ...func(*ListIncidentRecordsPaginatorOptions)) *ListIncidentRecordsPaginator {
	if params == nil {
		params = &ListIncidentRecordsInput{}
	}

	options := ListIncidentRecordsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListIncidentRecordsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListIncidentRecordsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListIncidentRecords page.
func (p *ListIncidentRecordsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListIncidentRecordsOutput, error) {
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

	result, err := p.client.ListIncidentRecords(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListIncidentRecords(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListIncidentRecords",
	}
}
