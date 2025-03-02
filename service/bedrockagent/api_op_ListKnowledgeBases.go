// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagent

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/bedrockagent/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List Knowledge Bases
func (c *Client) ListKnowledgeBases(ctx context.Context, params *ListKnowledgeBasesInput, optFns ...func(*Options)) (*ListKnowledgeBasesOutput, error) {
	if params == nil {
		params = &ListKnowledgeBasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListKnowledgeBases", params, optFns, c.addOperationListKnowledgeBasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListKnowledgeBasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListKnowledgeBasesInput struct {

	// Max Results.
	MaxResults *int32

	// Opaque continuation token of previous paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListKnowledgeBasesOutput struct {

	// List of KnowledgeBaseSummaries
	//
	// This member is required.
	KnowledgeBaseSummaries []types.KnowledgeBaseSummary

	// Opaque continuation token of previous paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListKnowledgeBasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListKnowledgeBases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListKnowledgeBases{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListKnowledgeBases"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListKnowledgeBases(options.Region), middleware.Before); err != nil {
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

// ListKnowledgeBasesAPIClient is a client that implements the ListKnowledgeBases
// operation.
type ListKnowledgeBasesAPIClient interface {
	ListKnowledgeBases(context.Context, *ListKnowledgeBasesInput, ...func(*Options)) (*ListKnowledgeBasesOutput, error)
}

var _ ListKnowledgeBasesAPIClient = (*Client)(nil)

// ListKnowledgeBasesPaginatorOptions is the paginator options for
// ListKnowledgeBases
type ListKnowledgeBasesPaginatorOptions struct {
	// Max Results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListKnowledgeBasesPaginator is a paginator for ListKnowledgeBases
type ListKnowledgeBasesPaginator struct {
	options   ListKnowledgeBasesPaginatorOptions
	client    ListKnowledgeBasesAPIClient
	params    *ListKnowledgeBasesInput
	nextToken *string
	firstPage bool
}

// NewListKnowledgeBasesPaginator returns a new ListKnowledgeBasesPaginator
func NewListKnowledgeBasesPaginator(client ListKnowledgeBasesAPIClient, params *ListKnowledgeBasesInput, optFns ...func(*ListKnowledgeBasesPaginatorOptions)) *ListKnowledgeBasesPaginator {
	if params == nil {
		params = &ListKnowledgeBasesInput{}
	}

	options := ListKnowledgeBasesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListKnowledgeBasesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListKnowledgeBasesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListKnowledgeBases page.
func (p *ListKnowledgeBasesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListKnowledgeBasesOutput, error) {
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

	result, err := p.client.ListKnowledgeBases(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListKnowledgeBases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListKnowledgeBases",
	}
}
