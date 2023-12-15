// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/workdocs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches metadata and the content of folders, documents, document versions, and
// comments.
func (c *Client) SearchResources(ctx context.Context, params *SearchResourcesInput, optFns ...func(*Options)) (*SearchResourcesOutput, error) {
	if params == nil {
		params = &SearchResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchResources", params, optFns, c.addOperationSearchResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchResourcesInput struct {

	// A list of attributes to include in the response. Used to request fields that
	// are not normally returned in a standard response.
	AdditionalResponseFields []types.AdditionalResponseFieldType

	// Amazon WorkDocs authentication token. Not required when using Amazon Web
	// Services administrator credentials to access the API.
	AuthenticationToken *string

	// Filters results based on entity metadata.
	Filters *types.Filters

	// Max results count per page.
	Limit *int32

	// The marker for the next set of results.
	Marker *string

	// Order by results in one or more categories.
	OrderBy []types.SearchSortResult

	// Filters based on the resource owner OrgId. This is a mandatory parameter when
	// using Admin SigV4 credentials.
	OrganizationId *string

	// Filter based on the text field type. A Folder has only a name and no content. A
	// Comment has only content and no name. A Document or Document Version has a name
	// and content
	QueryScopes []types.SearchQueryScopeType

	// The String to search for. Searches across different text fields based on
	// request parameters. Use double quotes around the query string for exact phrase
	// matches.
	QueryText *string

	noSmithyDocumentSerde
}

type SearchResourcesOutput struct {

	// List of Documents, Folders, Comments, and Document Versions matching the query.
	Items []types.ResponseItem

	// The marker to use when requesting the next set of results. If there are no
	// additional results, the string is empty.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchResources"); err != nil {
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
	if err = addOpSearchResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchResources(options.Region), middleware.Before); err != nil {
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

// SearchResourcesAPIClient is a client that implements the SearchResources
// operation.
type SearchResourcesAPIClient interface {
	SearchResources(context.Context, *SearchResourcesInput, ...func(*Options)) (*SearchResourcesOutput, error)
}

var _ SearchResourcesAPIClient = (*Client)(nil)

// SearchResourcesPaginatorOptions is the paginator options for SearchResources
type SearchResourcesPaginatorOptions struct {
	// Max results count per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchResourcesPaginator is a paginator for SearchResources
type SearchResourcesPaginator struct {
	options   SearchResourcesPaginatorOptions
	client    SearchResourcesAPIClient
	params    *SearchResourcesInput
	nextToken *string
	firstPage bool
}

// NewSearchResourcesPaginator returns a new SearchResourcesPaginator
func NewSearchResourcesPaginator(client SearchResourcesAPIClient, params *SearchResourcesInput, optFns ...func(*SearchResourcesPaginatorOptions)) *SearchResourcesPaginator {
	if params == nil {
		params = &SearchResourcesInput{}
	}

	options := SearchResourcesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchResources page.
func (p *SearchResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchResourcesOutput, error) {
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

	result, err := p.client.SearchResources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchResources",
	}
}
