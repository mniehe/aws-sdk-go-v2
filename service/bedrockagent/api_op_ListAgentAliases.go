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

// Lists all the Aliases for an Amazon Bedrock Agent
func (c *Client) ListAgentAliases(ctx context.Context, params *ListAgentAliasesInput, optFns ...func(*Options)) (*ListAgentAliasesOutput, error) {
	if params == nil {
		params = &ListAgentAliasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAgentAliases", params, optFns, c.addOperationListAgentAliasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAgentAliasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// List Agent Aliases Request
type ListAgentAliasesInput struct {

	// Id generated at the server side when an Agent is created
	//
	// This member is required.
	AgentId *string

	// Max Results.
	MaxResults *int32

	// Opaque continuation token of previous paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

// List Agent Aliases Response
type ListAgentAliasesOutput struct {

	// The list of summaries of all the aliases for an Agent.
	//
	// This member is required.
	AgentAliasSummaries []types.AgentAliasSummary

	// Opaque continuation token of previous paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAgentAliasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAgentAliases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAgentAliases{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAgentAliases"); err != nil {
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
	if err = addOpListAgentAliasesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAgentAliases(options.Region), middleware.Before); err != nil {
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

// ListAgentAliasesAPIClient is a client that implements the ListAgentAliases
// operation.
type ListAgentAliasesAPIClient interface {
	ListAgentAliases(context.Context, *ListAgentAliasesInput, ...func(*Options)) (*ListAgentAliasesOutput, error)
}

var _ ListAgentAliasesAPIClient = (*Client)(nil)

// ListAgentAliasesPaginatorOptions is the paginator options for ListAgentAliases
type ListAgentAliasesPaginatorOptions struct {
	// Max Results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAgentAliasesPaginator is a paginator for ListAgentAliases
type ListAgentAliasesPaginator struct {
	options   ListAgentAliasesPaginatorOptions
	client    ListAgentAliasesAPIClient
	params    *ListAgentAliasesInput
	nextToken *string
	firstPage bool
}

// NewListAgentAliasesPaginator returns a new ListAgentAliasesPaginator
func NewListAgentAliasesPaginator(client ListAgentAliasesAPIClient, params *ListAgentAliasesInput, optFns ...func(*ListAgentAliasesPaginatorOptions)) *ListAgentAliasesPaginator {
	if params == nil {
		params = &ListAgentAliasesInput{}
	}

	options := ListAgentAliasesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAgentAliasesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAgentAliasesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAgentAliases page.
func (p *ListAgentAliasesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAgentAliasesOutput, error) {
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

	result, err := p.client.ListAgentAliases(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAgentAliases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAgentAliases",
	}
}
