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

// Lists all the input sources of the Resilience Hub application. For more
// information about the input sources supported by Resilience Hub, see Discover
// the structure and describe your Resilience Hub application (https://docs.aws.amazon.com/resilience-hub/latest/userguide/discover-structure.html)
// .
func (c *Client) ListAppInputSources(ctx context.Context, params *ListAppInputSourcesInput, optFns ...func(*Options)) (*ListAppInputSourcesOutput, error) {
	if params == nil {
		params = &ListAppInputSourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAppInputSources", params, optFns, c.addOperationListAppInputSourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAppInputSourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAppInputSourcesInput struct {

	// Amazon Resource Name (ARN) of the Resilience Hub application. The format for
	// this ARN is: arn: partition :resiliencehub: region : account :app/ app-id . For
	// more information about ARNs, see Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference guide.
	//
	// This member is required.
	AppArn *string

	// Resilience Hub application version.
	//
	// This member is required.
	AppVersion *string

	// Maximum number of input sources to be displayed per Resilience Hub application.
	MaxResults *int32

	// Null, or the token from a previous call to get the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAppInputSourcesOutput struct {

	// The list of Resilience Hub application input sources.
	//
	// This member is required.
	AppInputSources []types.AppInputSource

	// Token for the next set of results, or null if there are no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAppInputSourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAppInputSources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAppInputSources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAppInputSources"); err != nil {
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
	if err = addOpListAppInputSourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAppInputSources(options.Region), middleware.Before); err != nil {
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

// ListAppInputSourcesAPIClient is a client that implements the
// ListAppInputSources operation.
type ListAppInputSourcesAPIClient interface {
	ListAppInputSources(context.Context, *ListAppInputSourcesInput, ...func(*Options)) (*ListAppInputSourcesOutput, error)
}

var _ ListAppInputSourcesAPIClient = (*Client)(nil)

// ListAppInputSourcesPaginatorOptions is the paginator options for
// ListAppInputSources
type ListAppInputSourcesPaginatorOptions struct {
	// Maximum number of input sources to be displayed per Resilience Hub application.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAppInputSourcesPaginator is a paginator for ListAppInputSources
type ListAppInputSourcesPaginator struct {
	options   ListAppInputSourcesPaginatorOptions
	client    ListAppInputSourcesAPIClient
	params    *ListAppInputSourcesInput
	nextToken *string
	firstPage bool
}

// NewListAppInputSourcesPaginator returns a new ListAppInputSourcesPaginator
func NewListAppInputSourcesPaginator(client ListAppInputSourcesAPIClient, params *ListAppInputSourcesInput, optFns ...func(*ListAppInputSourcesPaginatorOptions)) *ListAppInputSourcesPaginator {
	if params == nil {
		params = &ListAppInputSourcesInput{}
	}

	options := ListAppInputSourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAppInputSourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAppInputSourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAppInputSources page.
func (p *ListAppInputSourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAppInputSourcesOutput, error) {
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

	result, err := p.client.ListAppInputSources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAppInputSources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAppInputSources",
	}
}
