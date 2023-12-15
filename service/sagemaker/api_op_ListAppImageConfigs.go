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

// Lists the AppImageConfigs in your account and their properties. The list can be
// filtered by creation time or modified time, and whether the AppImageConfig name
// contains a specified string.
func (c *Client) ListAppImageConfigs(ctx context.Context, params *ListAppImageConfigsInput, optFns ...func(*Options)) (*ListAppImageConfigsOutput, error) {
	if params == nil {
		params = &ListAppImageConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAppImageConfigs", params, optFns, c.addOperationListAppImageConfigsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAppImageConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAppImageConfigsInput struct {

	// A filter that returns only AppImageConfigs created on or after the specified
	// time.
	CreationTimeAfter *time.Time

	// A filter that returns only AppImageConfigs created on or before the specified
	// time.
	CreationTimeBefore *time.Time

	// The total number of items to return in the response. If the total number of
	// items available is more than the value specified, a NextToken is provided in
	// the response. To resume pagination, provide the NextToken value in the as part
	// of a subsequent call. The default value is 10.
	MaxResults *int32

	// A filter that returns only AppImageConfigs modified on or after the specified
	// time.
	ModifiedTimeAfter *time.Time

	// A filter that returns only AppImageConfigs modified on or before the specified
	// time.
	ModifiedTimeBefore *time.Time

	// A filter that returns only AppImageConfigs whose name contains the specified
	// string.
	NameContains *string

	// If the previous call to ListImages didn't return the full set of
	// AppImageConfigs, the call returns a token for getting the next set of
	// AppImageConfigs.
	NextToken *string

	// The property used to sort results. The default value is CreationTime .
	SortBy types.AppImageConfigSortKey

	// The sort order. The default value is Descending .
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListAppImageConfigsOutput struct {

	// A list of AppImageConfigs and their properties.
	AppImageConfigs []types.AppImageConfigDetails

	// A token for getting the next set of AppImageConfigs, if there are any.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAppImageConfigsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAppImageConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAppImageConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAppImageConfigs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAppImageConfigs(options.Region), middleware.Before); err != nil {
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

// ListAppImageConfigsAPIClient is a client that implements the
// ListAppImageConfigs operation.
type ListAppImageConfigsAPIClient interface {
	ListAppImageConfigs(context.Context, *ListAppImageConfigsInput, ...func(*Options)) (*ListAppImageConfigsOutput, error)
}

var _ ListAppImageConfigsAPIClient = (*Client)(nil)

// ListAppImageConfigsPaginatorOptions is the paginator options for
// ListAppImageConfigs
type ListAppImageConfigsPaginatorOptions struct {
	// The total number of items to return in the response. If the total number of
	// items available is more than the value specified, a NextToken is provided in
	// the response. To resume pagination, provide the NextToken value in the as part
	// of a subsequent call. The default value is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAppImageConfigsPaginator is a paginator for ListAppImageConfigs
type ListAppImageConfigsPaginator struct {
	options   ListAppImageConfigsPaginatorOptions
	client    ListAppImageConfigsAPIClient
	params    *ListAppImageConfigsInput
	nextToken *string
	firstPage bool
}

// NewListAppImageConfigsPaginator returns a new ListAppImageConfigsPaginator
func NewListAppImageConfigsPaginator(client ListAppImageConfigsAPIClient, params *ListAppImageConfigsInput, optFns ...func(*ListAppImageConfigsPaginatorOptions)) *ListAppImageConfigsPaginator {
	if params == nil {
		params = &ListAppImageConfigsInput{}
	}

	options := ListAppImageConfigsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAppImageConfigsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAppImageConfigsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAppImageConfigs page.
func (p *ListAppImageConfigsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAppImageConfigsOutput, error) {
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

	result, err := p.client.ListAppImageConfigs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAppImageConfigs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAppImageConfigs",
	}
}
