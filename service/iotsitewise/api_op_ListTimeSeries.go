// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a paginated list of time series (data streams).
func (c *Client) ListTimeSeries(ctx context.Context, params *ListTimeSeriesInput, optFns ...func(*Options)) (*ListTimeSeriesOutput, error) {
	if params == nil {
		params = &ListTimeSeriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTimeSeries", params, optFns, c.addOperationListTimeSeriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTimeSeriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTimeSeriesInput struct {

	// The alias prefix of the time series.
	AliasPrefix *string

	// The ID of the asset in which the asset property was created. This can be either
	// the actual ID in UUID format, or else externalId: followed by the external ID,
	// if it has one. For more information, see Referencing objects with external IDs (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/object-ids.html#external-id-references)
	// in the IoT SiteWise User Guide.
	AssetId *string

	// The maximum number of results to return for each paginated request.
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string

	// The type of the time series. The time series type can be one of the following
	// values:
	//   - ASSOCIATED – The time series is associated with an asset property.
	//   - DISASSOCIATED – The time series isn't associated with any asset property.
	TimeSeriesType types.ListTimeSeriesType

	noSmithyDocumentSerde
}

type ListTimeSeriesOutput struct {

	// One or more time series summaries to list.
	//
	// This member is required.
	TimeSeriesSummaries []types.TimeSeriesSummary

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTimeSeriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTimeSeries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTimeSeries{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTimeSeries"); err != nil {
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
	if err = addEndpointPrefix_opListTimeSeriesMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTimeSeries(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListTimeSeriesMiddleware struct {
}

func (*endpointPrefix_opListTimeSeriesMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListTimeSeriesMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListTimeSeriesMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListTimeSeriesMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListTimeSeriesAPIClient is a client that implements the ListTimeSeries
// operation.
type ListTimeSeriesAPIClient interface {
	ListTimeSeries(context.Context, *ListTimeSeriesInput, ...func(*Options)) (*ListTimeSeriesOutput, error)
}

var _ ListTimeSeriesAPIClient = (*Client)(nil)

// ListTimeSeriesPaginatorOptions is the paginator options for ListTimeSeries
type ListTimeSeriesPaginatorOptions struct {
	// The maximum number of results to return for each paginated request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTimeSeriesPaginator is a paginator for ListTimeSeries
type ListTimeSeriesPaginator struct {
	options   ListTimeSeriesPaginatorOptions
	client    ListTimeSeriesAPIClient
	params    *ListTimeSeriesInput
	nextToken *string
	firstPage bool
}

// NewListTimeSeriesPaginator returns a new ListTimeSeriesPaginator
func NewListTimeSeriesPaginator(client ListTimeSeriesAPIClient, params *ListTimeSeriesInput, optFns ...func(*ListTimeSeriesPaginatorOptions)) *ListTimeSeriesPaginator {
	if params == nil {
		params = &ListTimeSeriesInput{}
	}

	options := ListTimeSeriesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTimeSeriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTimeSeriesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTimeSeries page.
func (p *ListTimeSeriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTimeSeriesOutput, error) {
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

	result, err := p.client.ListTimeSeries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTimeSeries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTimeSeries",
	}
}
