// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/iotthingsgraph/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches for entities of the specified type. You can search for entities in
// your namespace and the public namespace that you're tracking.
//
// Deprecated: since: 2022-08-30
func (c *Client) SearchEntities(ctx context.Context, params *SearchEntitiesInput, optFns ...func(*Options)) (*SearchEntitiesOutput, error) {
	if params == nil {
		params = &SearchEntitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchEntities", params, optFns, c.addOperationSearchEntitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchEntitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchEntitiesInput struct {

	// The entity types for which to search.
	//
	// This member is required.
	EntityTypes []types.EntityType

	// Optional filter to apply to the search. Valid filters are NAME NAMESPACE ,
	// SEMANTIC_TYPE_PATH and REFERENCED_ENTITY_ID . REFERENCED_ENTITY_ID filters on
	// entities that are used by the entity in the result set. For example, you can
	// filter on the ID of a property that is used in a state. Multiple filters
	// function as OR criteria in the query. Multiple values passed inside the filter
	// function as AND criteria.
	Filters []types.EntityFilter

	// The maximum number of results to return in the response.
	MaxResults *int32

	// The version of the user's namespace. Defaults to the latest version of the
	// user's namespace.
	NamespaceVersion *int64

	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchEntitiesOutput struct {

	// An array of descriptions for each entity returned in the search result.
	Descriptions []types.EntityDescription

	// The string to specify as nextToken when you request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchEntitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchEntities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchEntities{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchEntities"); err != nil {
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
	if err = addOpSearchEntitiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchEntities(options.Region), middleware.Before); err != nil {
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

// SearchEntitiesAPIClient is a client that implements the SearchEntities
// operation.
type SearchEntitiesAPIClient interface {
	SearchEntities(context.Context, *SearchEntitiesInput, ...func(*Options)) (*SearchEntitiesOutput, error)
}

var _ SearchEntitiesAPIClient = (*Client)(nil)

// SearchEntitiesPaginatorOptions is the paginator options for SearchEntities
type SearchEntitiesPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchEntitiesPaginator is a paginator for SearchEntities
type SearchEntitiesPaginator struct {
	options   SearchEntitiesPaginatorOptions
	client    SearchEntitiesAPIClient
	params    *SearchEntitiesInput
	nextToken *string
	firstPage bool
}

// NewSearchEntitiesPaginator returns a new SearchEntitiesPaginator
func NewSearchEntitiesPaginator(client SearchEntitiesAPIClient, params *SearchEntitiesInput, optFns ...func(*SearchEntitiesPaginatorOptions)) *SearchEntitiesPaginator {
	if params == nil {
		params = &SearchEntitiesInput{}
	}

	options := SearchEntitiesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchEntitiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchEntitiesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchEntities page.
func (p *SearchEntitiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchEntitiesOutput, error) {
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

	result, err := p.client.SearchEntities(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchEntities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchEntities",
	}
}
