// Code generated by smithy-go-codegen DO NOT EDIT.

package qconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/qconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Performs a manual search against the specified assistant. To retrieve
// recommendations for an assistant, use GetRecommendations (https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_GetRecommendations.html)
// .
func (c *Client) QueryAssistant(ctx context.Context, params *QueryAssistantInput, optFns ...func(*Options)) (*QueryAssistantOutput, error) {
	if params == nil {
		params = &QueryAssistantInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "QueryAssistant", params, optFns, c.addOperationQueryAssistantMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*QueryAssistantOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type QueryAssistantInput struct {

	// The identifier of the Amazon Q assistant. Can be either the ID or the ARN. URLs
	// cannot contain the ARN.
	//
	// This member is required.
	AssistantId *string

	// The text to search for.
	//
	// This member is required.
	QueryText *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// Information about how to query content.
	QueryCondition []types.QueryCondition

	// The identifier of the Amazon Q session. Can be either the ID or the ARN. URLs
	// cannot contain the ARN.
	SessionId *string

	noSmithyDocumentSerde
}

type QueryAssistantOutput struct {

	// The results of the query.
	//
	// This member is required.
	Results []types.ResultData

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationQueryAssistantMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpQueryAssistant{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpQueryAssistant{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "QueryAssistant"); err != nil {
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
	if err = addOpQueryAssistantValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opQueryAssistant(options.Region), middleware.Before); err != nil {
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

// QueryAssistantAPIClient is a client that implements the QueryAssistant
// operation.
type QueryAssistantAPIClient interface {
	QueryAssistant(context.Context, *QueryAssistantInput, ...func(*Options)) (*QueryAssistantOutput, error)
}

var _ QueryAssistantAPIClient = (*Client)(nil)

// QueryAssistantPaginatorOptions is the paginator options for QueryAssistant
type QueryAssistantPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// QueryAssistantPaginator is a paginator for QueryAssistant
type QueryAssistantPaginator struct {
	options   QueryAssistantPaginatorOptions
	client    QueryAssistantAPIClient
	params    *QueryAssistantInput
	nextToken *string
	firstPage bool
}

// NewQueryAssistantPaginator returns a new QueryAssistantPaginator
func NewQueryAssistantPaginator(client QueryAssistantAPIClient, params *QueryAssistantInput, optFns ...func(*QueryAssistantPaginatorOptions)) *QueryAssistantPaginator {
	if params == nil {
		params = &QueryAssistantInput{}
	}

	options := QueryAssistantPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &QueryAssistantPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *QueryAssistantPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next QueryAssistant page.
func (p *QueryAssistantPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*QueryAssistantOutput, error) {
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

	result, err := p.client.QueryAssistant(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opQueryAssistant(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "QueryAssistant",
	}
}
