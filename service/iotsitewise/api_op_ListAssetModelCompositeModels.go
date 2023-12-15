// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a paginated list of composite models associated with the asset model
func (c *Client) ListAssetModelCompositeModels(ctx context.Context, params *ListAssetModelCompositeModelsInput, optFns ...func(*Options)) (*ListAssetModelCompositeModelsOutput, error) {
	if params == nil {
		params = &ListAssetModelCompositeModelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssetModelCompositeModels", params, optFns, c.addOperationListAssetModelCompositeModelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssetModelCompositeModelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssetModelCompositeModelsInput struct {

	// The ID of the asset model. This can be either the actual ID in UUID format, or
	// else externalId: followed by the external ID, if it has one. For more
	// information, see Referencing objects with external IDs (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/object-ids.html#external-id-references)
	// in the IoT SiteWise User Guide.
	//
	// This member is required.
	AssetModelId *string

	// The maximum number of results to return for each paginated request. Default: 50
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAssetModelCompositeModelsOutput struct {

	// A list that summarizes each composite model.
	//
	// This member is required.
	AssetModelCompositeModelSummaries []types.AssetModelCompositeModelSummary

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssetModelCompositeModelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAssetModelCompositeModels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssetModelCompositeModels{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAssetModelCompositeModels"); err != nil {
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
	if err = addEndpointPrefix_opListAssetModelCompositeModelsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListAssetModelCompositeModelsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssetModelCompositeModels(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListAssetModelCompositeModelsMiddleware struct {
}

func (*endpointPrefix_opListAssetModelCompositeModelsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListAssetModelCompositeModelsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListAssetModelCompositeModelsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListAssetModelCompositeModelsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListAssetModelCompositeModelsAPIClient is a client that implements the
// ListAssetModelCompositeModels operation.
type ListAssetModelCompositeModelsAPIClient interface {
	ListAssetModelCompositeModels(context.Context, *ListAssetModelCompositeModelsInput, ...func(*Options)) (*ListAssetModelCompositeModelsOutput, error)
}

var _ ListAssetModelCompositeModelsAPIClient = (*Client)(nil)

// ListAssetModelCompositeModelsPaginatorOptions is the paginator options for
// ListAssetModelCompositeModels
type ListAssetModelCompositeModelsPaginatorOptions struct {
	// The maximum number of results to return for each paginated request. Default: 50
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssetModelCompositeModelsPaginator is a paginator for
// ListAssetModelCompositeModels
type ListAssetModelCompositeModelsPaginator struct {
	options   ListAssetModelCompositeModelsPaginatorOptions
	client    ListAssetModelCompositeModelsAPIClient
	params    *ListAssetModelCompositeModelsInput
	nextToken *string
	firstPage bool
}

// NewListAssetModelCompositeModelsPaginator returns a new
// ListAssetModelCompositeModelsPaginator
func NewListAssetModelCompositeModelsPaginator(client ListAssetModelCompositeModelsAPIClient, params *ListAssetModelCompositeModelsInput, optFns ...func(*ListAssetModelCompositeModelsPaginatorOptions)) *ListAssetModelCompositeModelsPaginator {
	if params == nil {
		params = &ListAssetModelCompositeModelsInput{}
	}

	options := ListAssetModelCompositeModelsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssetModelCompositeModelsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssetModelCompositeModelsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssetModelCompositeModels page.
func (p *ListAssetModelCompositeModelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssetModelCompositeModelsOutput, error) {
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

	result, err := p.client.ListAssetModelCompositeModels(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssetModelCompositeModels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAssetModelCompositeModels",
	}
}
