// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists virtual tapes in your virtual tape library (VTL) and your virtual tape
// shelf (VTS). You specify the tapes to list by specifying one or more tape Amazon
// Resource Names (ARNs). If you don't specify a tape ARN, the operation lists all
// virtual tapes in both your VTL and VTS. This operation supports pagination. By
// default, the operation returns a maximum of up to 100 tapes. You can optionally
// specify the Limit parameter in the body to limit the number of tapes in the
// response. If the number of tapes returned in the response is truncated, the
// response includes a Marker element that you can use in your subsequent request
// to retrieve the next set of tapes. This operation is only supported in the tape
// gateway type.
func (c *Client) ListTapes(ctx context.Context, params *ListTapesInput, optFns ...func(*Options)) (*ListTapesOutput, error) {
	if params == nil {
		params = &ListTapesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTapes", params, optFns, c.addOperationListTapesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTapesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object that contains one or more of the following fields:
//   - ListTapesInput$Limit
//   - ListTapesInput$Marker
//   - ListTapesInput$TapeARNs
type ListTapesInput struct {

	// An optional number limit for the tapes in the list returned by this call.
	Limit *int32

	// A string that indicates the position at which to begin the returned list of
	// tapes.
	Marker *string

	// The Amazon Resource Name (ARN) of each of the tapes you want to list. If you
	// don't specify a tape ARN, the response lists all tapes in both your VTL and VTS.
	TapeARNs []string

	noSmithyDocumentSerde
}

// A JSON object containing the following fields:
//   - ListTapesOutput$Marker
//   - ListTapesOutput$VolumeInfos
type ListTapesOutput struct {

	// A string that indicates the position at which to begin returning the next list
	// of tapes. Use the marker in your next request to continue pagination of tapes.
	// If there are no more tapes to list, this element does not appear in the response
	// body.
	Marker *string

	// An array of TapeInfo objects, where each object describes a single tape. If
	// there are no tapes in the tape library or VTS, then the TapeInfos is an empty
	// array.
	TapeInfos []types.TapeInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTapesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTapes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTapes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTapes"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTapes(options.Region), middleware.Before); err != nil {
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

// ListTapesAPIClient is a client that implements the ListTapes operation.
type ListTapesAPIClient interface {
	ListTapes(context.Context, *ListTapesInput, ...func(*Options)) (*ListTapesOutput, error)
}

var _ ListTapesAPIClient = (*Client)(nil)

// ListTapesPaginatorOptions is the paginator options for ListTapes
type ListTapesPaginatorOptions struct {
	// An optional number limit for the tapes in the list returned by this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTapesPaginator is a paginator for ListTapes
type ListTapesPaginator struct {
	options   ListTapesPaginatorOptions
	client    ListTapesAPIClient
	params    *ListTapesInput
	nextToken *string
	firstPage bool
}

// NewListTapesPaginator returns a new ListTapesPaginator
func NewListTapesPaginator(client ListTapesAPIClient, params *ListTapesInput, optFns ...func(*ListTapesPaginatorOptions)) *ListTapesPaginator {
	if params == nil {
		params = &ListTapesInput{}
	}

	options := ListTapesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTapesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTapesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTapes page.
func (p *ListTapesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTapesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListTapes(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListTapes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTapes",
	}
}
