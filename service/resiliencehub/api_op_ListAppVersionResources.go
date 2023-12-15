// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the resources in an Resilience Hub application.
func (c *Client) ListAppVersionResources(ctx context.Context, params *ListAppVersionResourcesInput, optFns ...func(*Options)) (*ListAppVersionResourcesOutput, error) {
	if params == nil {
		params = &ListAppVersionResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAppVersionResources", params, optFns, c.addOperationListAppVersionResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAppVersionResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAppVersionResourcesInput struct {

	// Amazon Resource Name (ARN) of the Resilience Hub application. The format for
	// this ARN is: arn: partition :resiliencehub: region : account :app/ app-id . For
	// more information about ARNs, see Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference guide.
	//
	// This member is required.
	AppArn *string

	// The version of the application.
	//
	// This member is required.
	AppVersion *string

	// Maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so
	// that the remaining results can be retrieved.
	MaxResults *int32

	// Null, or the token from a previous call to get the next set of results.
	NextToken *string

	// The identifier for a specific resolution.
	ResolutionId *string

	noSmithyDocumentSerde
}

type ListAppVersionResourcesOutput struct {

	// The physical resources in the application version.
	//
	// This member is required.
	PhysicalResources []types.PhysicalResource

	// The ID for a specific resolution.
	//
	// This member is required.
	ResolutionId *string

	// Token for the next set of results, or null if there are no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAppVersionResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAppVersionResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAppVersionResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAppVersionResources"); err != nil {
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
	if err = addOpListAppVersionResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAppVersionResources(options.Region), middleware.Before); err != nil {
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

// ListAppVersionResourcesAPIClient is a client that implements the
// ListAppVersionResources operation.
type ListAppVersionResourcesAPIClient interface {
	ListAppVersionResources(context.Context, *ListAppVersionResourcesInput, ...func(*Options)) (*ListAppVersionResourcesOutput, error)
}

var _ ListAppVersionResourcesAPIClient = (*Client)(nil)

// ListAppVersionResourcesPaginatorOptions is the paginator options for
// ListAppVersionResources
type ListAppVersionResourcesPaginatorOptions struct {
	// Maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so
	// that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAppVersionResourcesPaginator is a paginator for ListAppVersionResources
type ListAppVersionResourcesPaginator struct {
	options   ListAppVersionResourcesPaginatorOptions
	client    ListAppVersionResourcesAPIClient
	params    *ListAppVersionResourcesInput
	nextToken *string
	firstPage bool
}

// NewListAppVersionResourcesPaginator returns a new
// ListAppVersionResourcesPaginator
func NewListAppVersionResourcesPaginator(client ListAppVersionResourcesAPIClient, params *ListAppVersionResourcesInput, optFns ...func(*ListAppVersionResourcesPaginatorOptions)) *ListAppVersionResourcesPaginator {
	if params == nil {
		params = &ListAppVersionResourcesInput{}
	}

	options := ListAppVersionResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAppVersionResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAppVersionResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAppVersionResources page.
func (p *ListAppVersionResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAppVersionResourcesOutput, error) {
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

	result, err := p.client.ListAppVersionResources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAppVersionResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAppVersionResources",
	}
}
