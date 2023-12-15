// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the resources that are included in the protection group.
func (c *Client) ListResourcesInProtectionGroup(ctx context.Context, params *ListResourcesInProtectionGroupInput, optFns ...func(*Options)) (*ListResourcesInProtectionGroupOutput, error) {
	if params == nil {
		params = &ListResourcesInProtectionGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResourcesInProtectionGroup", params, optFns, c.addOperationListResourcesInProtectionGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourcesInProtectionGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourcesInProtectionGroupInput struct {

	// The name of the protection group. You use this to identify the protection group
	// in lists and to manage the protection group, for example to update, delete, or
	// describe it.
	//
	// This member is required.
	ProtectionGroupId *string

	// The greatest number of objects that you want Shield Advanced to return to the
	// list request. Shield Advanced might return fewer objects than you indicate in
	// this setting, even if more objects are available. If there are more objects
	// remaining, Shield Advanced will always also return a NextToken value in the
	// response. The default setting is 20.
	MaxResults *int32

	// When you request a list of objects from Shield Advanced, if the response does
	// not include all of the remaining available objects, Shield Advanced includes a
	// NextToken value in the response. You can retrieve the next batch of objects by
	// requesting the list again and providing the token that was returned by the prior
	// call in your request. You can indicate the maximum number of objects that you
	// want Shield Advanced to return for a single call with the MaxResults setting.
	// Shield Advanced will not return more than MaxResults objects, but may return
	// fewer, even if more objects are still available. Whenever more objects remain
	// that Shield Advanced has not yet returned to you, the response will include a
	// NextToken value. On your first call to a list operation, leave this setting
	// empty.
	NextToken *string

	noSmithyDocumentSerde
}

type ListResourcesInProtectionGroupOutput struct {

	// The Amazon Resource Names (ARNs) of the resources that are included in the
	// protection group.
	//
	// This member is required.
	ResourceArns []string

	// When you request a list of objects from Shield Advanced, if the response does
	// not include all of the remaining available objects, Shield Advanced includes a
	// NextToken value in the response. You can retrieve the next batch of objects by
	// requesting the list again and providing the token that was returned by the prior
	// call in your request. You can indicate the maximum number of objects that you
	// want Shield Advanced to return for a single call with the MaxResults setting.
	// Shield Advanced will not return more than MaxResults objects, but may return
	// fewer, even if more objects are still available. Whenever more objects remain
	// that Shield Advanced has not yet returned to you, the response will include a
	// NextToken value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResourcesInProtectionGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResourcesInProtectionGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResourcesInProtectionGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListResourcesInProtectionGroup"); err != nil {
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
	if err = addOpListResourcesInProtectionGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResourcesInProtectionGroup(options.Region), middleware.Before); err != nil {
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

// ListResourcesInProtectionGroupAPIClient is a client that implements the
// ListResourcesInProtectionGroup operation.
type ListResourcesInProtectionGroupAPIClient interface {
	ListResourcesInProtectionGroup(context.Context, *ListResourcesInProtectionGroupInput, ...func(*Options)) (*ListResourcesInProtectionGroupOutput, error)
}

var _ ListResourcesInProtectionGroupAPIClient = (*Client)(nil)

// ListResourcesInProtectionGroupPaginatorOptions is the paginator options for
// ListResourcesInProtectionGroup
type ListResourcesInProtectionGroupPaginatorOptions struct {
	// The greatest number of objects that you want Shield Advanced to return to the
	// list request. Shield Advanced might return fewer objects than you indicate in
	// this setting, even if more objects are available. If there are more objects
	// remaining, Shield Advanced will always also return a NextToken value in the
	// response. The default setting is 20.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResourcesInProtectionGroupPaginator is a paginator for
// ListResourcesInProtectionGroup
type ListResourcesInProtectionGroupPaginator struct {
	options   ListResourcesInProtectionGroupPaginatorOptions
	client    ListResourcesInProtectionGroupAPIClient
	params    *ListResourcesInProtectionGroupInput
	nextToken *string
	firstPage bool
}

// NewListResourcesInProtectionGroupPaginator returns a new
// ListResourcesInProtectionGroupPaginator
func NewListResourcesInProtectionGroupPaginator(client ListResourcesInProtectionGroupAPIClient, params *ListResourcesInProtectionGroupInput, optFns ...func(*ListResourcesInProtectionGroupPaginatorOptions)) *ListResourcesInProtectionGroupPaginator {
	if params == nil {
		params = &ListResourcesInProtectionGroupInput{}
	}

	options := ListResourcesInProtectionGroupPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResourcesInProtectionGroupPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResourcesInProtectionGroupPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResourcesInProtectionGroup page.
func (p *ListResourcesInProtectionGroupPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResourcesInProtectionGroupOutput, error) {
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

	result, err := p.client.ListResourcesInProtectionGroup(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListResourcesInProtectionGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListResourcesInProtectionGroup",
	}
}
