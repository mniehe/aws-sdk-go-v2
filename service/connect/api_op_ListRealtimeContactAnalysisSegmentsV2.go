// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a list of analysis segments for a real-time analysis session.
func (c *Client) ListRealtimeContactAnalysisSegmentsV2(ctx context.Context, params *ListRealtimeContactAnalysisSegmentsV2Input, optFns ...func(*Options)) (*ListRealtimeContactAnalysisSegmentsV2Output, error) {
	if params == nil {
		params = &ListRealtimeContactAnalysisSegmentsV2Input{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRealtimeContactAnalysisSegmentsV2", params, optFns, c.addOperationListRealtimeContactAnalysisSegmentsV2Middlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRealtimeContactAnalysisSegmentsV2Output)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRealtimeContactAnalysisSegmentsV2Input struct {

	// The identifier of the contact in this instance of Amazon Connect.
	//
	// This member is required.
	ContactId *string

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The Contact Lens output type to be returned.
	//
	// This member is required.
	OutputType types.RealTimeContactAnalysisOutputType

	// Enum with segment types . Each value corresponds to a segment type returned in
	// the segments list of the API. Each segment type has its own structure. Different
	// channels may have different sets of supported segment types.
	//
	// This member is required.
	SegmentTypes []types.RealTimeContactAnalysisSegmentType

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRealtimeContactAnalysisSegmentsV2Output struct {

	// The channel of the contact. Voice will not be returned.
	//
	// This member is required.
	Channel types.RealTimeContactAnalysisSupportedChannel

	// An analyzed transcript or category.
	//
	// This member is required.
	Segments []types.RealtimeContactAnalysisSegment

	// Status of real-time contact analysis.
	//
	// This member is required.
	Status types.RealTimeContactAnalysisStatus

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRealtimeContactAnalysisSegmentsV2Middlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRealtimeContactAnalysisSegmentsV2{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRealtimeContactAnalysisSegmentsV2{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRealtimeContactAnalysisSegmentsV2"); err != nil {
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
	if err = addOpListRealtimeContactAnalysisSegmentsV2ValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRealtimeContactAnalysisSegmentsV2(options.Region), middleware.Before); err != nil {
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

// ListRealtimeContactAnalysisSegmentsV2APIClient is a client that implements the
// ListRealtimeContactAnalysisSegmentsV2 operation.
type ListRealtimeContactAnalysisSegmentsV2APIClient interface {
	ListRealtimeContactAnalysisSegmentsV2(context.Context, *ListRealtimeContactAnalysisSegmentsV2Input, ...func(*Options)) (*ListRealtimeContactAnalysisSegmentsV2Output, error)
}

var _ ListRealtimeContactAnalysisSegmentsV2APIClient = (*Client)(nil)

// ListRealtimeContactAnalysisSegmentsV2PaginatorOptions is the paginator options
// for ListRealtimeContactAnalysisSegmentsV2
type ListRealtimeContactAnalysisSegmentsV2PaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRealtimeContactAnalysisSegmentsV2Paginator is a paginator for
// ListRealtimeContactAnalysisSegmentsV2
type ListRealtimeContactAnalysisSegmentsV2Paginator struct {
	options   ListRealtimeContactAnalysisSegmentsV2PaginatorOptions
	client    ListRealtimeContactAnalysisSegmentsV2APIClient
	params    *ListRealtimeContactAnalysisSegmentsV2Input
	nextToken *string
	firstPage bool
}

// NewListRealtimeContactAnalysisSegmentsV2Paginator returns a new
// ListRealtimeContactAnalysisSegmentsV2Paginator
func NewListRealtimeContactAnalysisSegmentsV2Paginator(client ListRealtimeContactAnalysisSegmentsV2APIClient, params *ListRealtimeContactAnalysisSegmentsV2Input, optFns ...func(*ListRealtimeContactAnalysisSegmentsV2PaginatorOptions)) *ListRealtimeContactAnalysisSegmentsV2Paginator {
	if params == nil {
		params = &ListRealtimeContactAnalysisSegmentsV2Input{}
	}

	options := ListRealtimeContactAnalysisSegmentsV2PaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRealtimeContactAnalysisSegmentsV2Paginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRealtimeContactAnalysisSegmentsV2Paginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRealtimeContactAnalysisSegmentsV2 page.
func (p *ListRealtimeContactAnalysisSegmentsV2Paginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRealtimeContactAnalysisSegmentsV2Output, error) {
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

	result, err := p.client.ListRealtimeContactAnalysisSegmentsV2(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRealtimeContactAnalysisSegmentsV2(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRealtimeContactAnalysisSegmentsV2",
	}
}
