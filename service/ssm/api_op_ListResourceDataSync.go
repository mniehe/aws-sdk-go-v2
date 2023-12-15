// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists your resource data sync configurations. Includes information about the
// last time a sync attempted to start, the last sync status, and the last time a
// sync successfully completed. The number of sync configurations might be too
// large to return using a single call to ListResourceDataSync . You can limit the
// number of sync configurations returned by using the MaxResults parameter. To
// determine whether there are more sync configurations to list, check the value of
// NextToken in the output. If there are more sync configurations to list, you can
// request them by specifying the NextToken returned in the call to the parameter
// of a subsequent call.
func (c *Client) ListResourceDataSync(ctx context.Context, params *ListResourceDataSyncInput, optFns ...func(*Options)) (*ListResourceDataSyncOutput, error) {
	if params == nil {
		params = &ListResourceDataSyncInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResourceDataSync", params, optFns, c.addOperationListResourceDataSyncMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourceDataSyncOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourceDataSyncInput struct {

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string

	// View a list of resource data syncs according to the sync type. Specify
	// SyncToDestination to view resource data syncs that synchronize data to an Amazon
	// S3 bucket. Specify SyncFromSource to view resource data syncs from
	// Organizations or from multiple Amazon Web Services Regions.
	SyncType *string

	noSmithyDocumentSerde
}

type ListResourceDataSyncOutput struct {

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// A list of your current resource data sync configurations and their statuses.
	ResourceDataSyncItems []types.ResourceDataSyncItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResourceDataSyncMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResourceDataSync{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResourceDataSync{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListResourceDataSync"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResourceDataSync(options.Region), middleware.Before); err != nil {
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

// ListResourceDataSyncAPIClient is a client that implements the
// ListResourceDataSync operation.
type ListResourceDataSyncAPIClient interface {
	ListResourceDataSync(context.Context, *ListResourceDataSyncInput, ...func(*Options)) (*ListResourceDataSyncOutput, error)
}

var _ ListResourceDataSyncAPIClient = (*Client)(nil)

// ListResourceDataSyncPaginatorOptions is the paginator options for
// ListResourceDataSync
type ListResourceDataSyncPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResourceDataSyncPaginator is a paginator for ListResourceDataSync
type ListResourceDataSyncPaginator struct {
	options   ListResourceDataSyncPaginatorOptions
	client    ListResourceDataSyncAPIClient
	params    *ListResourceDataSyncInput
	nextToken *string
	firstPage bool
}

// NewListResourceDataSyncPaginator returns a new ListResourceDataSyncPaginator
func NewListResourceDataSyncPaginator(client ListResourceDataSyncAPIClient, params *ListResourceDataSyncInput, optFns ...func(*ListResourceDataSyncPaginatorOptions)) *ListResourceDataSyncPaginator {
	if params == nil {
		params = &ListResourceDataSyncInput{}
	}

	options := ListResourceDataSyncPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResourceDataSyncPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResourceDataSyncPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResourceDataSync page.
func (p *ListResourceDataSyncPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResourceDataSyncOutput, error) {
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

	result, err := p.client.ListResourceDataSync(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListResourceDataSync(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListResourceDataSync",
	}
}
