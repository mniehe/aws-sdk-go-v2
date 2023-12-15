// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List provisioned resources for a service instance with details.
func (c *Client) ListServiceInstanceProvisionedResources(ctx context.Context, params *ListServiceInstanceProvisionedResourcesInput, optFns ...func(*Options)) (*ListServiceInstanceProvisionedResourcesOutput, error) {
	if params == nil {
		params = &ListServiceInstanceProvisionedResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServiceInstanceProvisionedResources", params, optFns, c.addOperationListServiceInstanceProvisionedResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServiceInstanceProvisionedResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServiceInstanceProvisionedResourcesInput struct {

	// The name of the service instance whose provisioned resources you want.
	//
	// This member is required.
	ServiceInstanceName *string

	// The name of the service that serviceInstanceName is associated to.
	//
	// This member is required.
	ServiceName *string

	// A token that indicates the location of the next provisioned resource in the
	// array of provisioned resources, after the list of provisioned resources that was
	// previously requested.
	NextToken *string

	noSmithyDocumentSerde
}

type ListServiceInstanceProvisionedResourcesOutput struct {

	// An array of provisioned resources for a service instance.
	//
	// This member is required.
	ProvisionedResources []types.ProvisionedResource

	// A token that indicates the location of the next provisioned resource in the
	// array of provisioned resources, after the current requested list of provisioned
	// resources.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListServiceInstanceProvisionedResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListServiceInstanceProvisionedResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListServiceInstanceProvisionedResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListServiceInstanceProvisionedResources"); err != nil {
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
	if err = addOpListServiceInstanceProvisionedResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServiceInstanceProvisionedResources(options.Region), middleware.Before); err != nil {
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

// ListServiceInstanceProvisionedResourcesAPIClient is a client that implements
// the ListServiceInstanceProvisionedResources operation.
type ListServiceInstanceProvisionedResourcesAPIClient interface {
	ListServiceInstanceProvisionedResources(context.Context, *ListServiceInstanceProvisionedResourcesInput, ...func(*Options)) (*ListServiceInstanceProvisionedResourcesOutput, error)
}

var _ ListServiceInstanceProvisionedResourcesAPIClient = (*Client)(nil)

// ListServiceInstanceProvisionedResourcesPaginatorOptions is the paginator
// options for ListServiceInstanceProvisionedResources
type ListServiceInstanceProvisionedResourcesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListServiceInstanceProvisionedResourcesPaginator is a paginator for
// ListServiceInstanceProvisionedResources
type ListServiceInstanceProvisionedResourcesPaginator struct {
	options   ListServiceInstanceProvisionedResourcesPaginatorOptions
	client    ListServiceInstanceProvisionedResourcesAPIClient
	params    *ListServiceInstanceProvisionedResourcesInput
	nextToken *string
	firstPage bool
}

// NewListServiceInstanceProvisionedResourcesPaginator returns a new
// ListServiceInstanceProvisionedResourcesPaginator
func NewListServiceInstanceProvisionedResourcesPaginator(client ListServiceInstanceProvisionedResourcesAPIClient, params *ListServiceInstanceProvisionedResourcesInput, optFns ...func(*ListServiceInstanceProvisionedResourcesPaginatorOptions)) *ListServiceInstanceProvisionedResourcesPaginator {
	if params == nil {
		params = &ListServiceInstanceProvisionedResourcesInput{}
	}

	options := ListServiceInstanceProvisionedResourcesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListServiceInstanceProvisionedResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListServiceInstanceProvisionedResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListServiceInstanceProvisionedResources page.
func (p *ListServiceInstanceProvisionedResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListServiceInstanceProvisionedResourcesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListServiceInstanceProvisionedResources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListServiceInstanceProvisionedResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListServiceInstanceProvisionedResources",
	}
}
