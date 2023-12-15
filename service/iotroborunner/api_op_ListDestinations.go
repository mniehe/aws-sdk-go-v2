// Code generated by smithy-go-codegen DO NOT EDIT.

package iotroborunner

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/iotroborunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Grants permission to list destinations
func (c *Client) ListDestinations(ctx context.Context, params *ListDestinationsInput, optFns ...func(*Options)) (*ListDestinationsOutput, error) {
	if params == nil {
		params = &ListDestinationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDestinations", params, optFns, c.addOperationListDestinationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDestinationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDestinationsInput struct {

	// Site ARN.
	//
	// This member is required.
	Site *string

	// Maximum number of results to retrieve in a single call.
	MaxResults *int32

	// Pagination token returned when another page of data exists. Provide it in your
	// next call to the API to receive the next page.
	NextToken *string

	// State of the destination.
	State types.DestinationState

	noSmithyDocumentSerde
}

type ListDestinationsOutput struct {

	// List of destinations.
	Destinations []types.Destination

	// Pagination token returned when another page of data exists. Provide it in your
	// next call to the API to receive the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDestinationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDestinations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDestinations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDestinations"); err != nil {
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
	if err = addOpListDestinationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDestinations(options.Region), middleware.Before); err != nil {
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

// ListDestinationsAPIClient is a client that implements the ListDestinations
// operation.
type ListDestinationsAPIClient interface {
	ListDestinations(context.Context, *ListDestinationsInput, ...func(*Options)) (*ListDestinationsOutput, error)
}

var _ ListDestinationsAPIClient = (*Client)(nil)

// ListDestinationsPaginatorOptions is the paginator options for ListDestinations
type ListDestinationsPaginatorOptions struct {
	// Maximum number of results to retrieve in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDestinationsPaginator is a paginator for ListDestinations
type ListDestinationsPaginator struct {
	options   ListDestinationsPaginatorOptions
	client    ListDestinationsAPIClient
	params    *ListDestinationsInput
	nextToken *string
	firstPage bool
}

// NewListDestinationsPaginator returns a new ListDestinationsPaginator
func NewListDestinationsPaginator(client ListDestinationsAPIClient, params *ListDestinationsInput, optFns ...func(*ListDestinationsPaginatorOptions)) *ListDestinationsPaginator {
	if params == nil {
		params = &ListDestinationsInput{}
	}

	options := ListDestinationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDestinationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDestinationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDestinations page.
func (p *ListDestinationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDestinationsOutput, error) {
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

	result, err := p.client.ListDestinations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDestinations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDestinations",
	}
}
