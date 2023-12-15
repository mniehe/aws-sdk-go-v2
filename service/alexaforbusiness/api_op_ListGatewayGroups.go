// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of gateway group summaries. Use GetGatewayGroup to retrieve
// details of a specific gateway group.
//
// Deprecated: Alexa For Business is no longer supported
func (c *Client) ListGatewayGroups(ctx context.Context, params *ListGatewayGroupsInput, optFns ...func(*Options)) (*ListGatewayGroupsOutput, error) {
	if params == nil {
		params = &ListGatewayGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGatewayGroups", params, optFns, c.addOperationListGatewayGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGatewayGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGatewayGroupsInput struct {

	// The maximum number of gateway group summaries to return. The default is 50.
	MaxResults *int32

	// The token used to paginate though multiple pages of gateway group summaries.
	NextToken *string

	noSmithyDocumentSerde
}

type ListGatewayGroupsOutput struct {

	// The gateway groups in the list.
	GatewayGroups []types.GatewayGroupSummary

	// The token used to paginate though multiple pages of gateway group summaries.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGatewayGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListGatewayGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListGatewayGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListGatewayGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGatewayGroups(options.Region), middleware.Before); err != nil {
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

// ListGatewayGroupsAPIClient is a client that implements the ListGatewayGroups
// operation.
type ListGatewayGroupsAPIClient interface {
	ListGatewayGroups(context.Context, *ListGatewayGroupsInput, ...func(*Options)) (*ListGatewayGroupsOutput, error)
}

var _ ListGatewayGroupsAPIClient = (*Client)(nil)

// ListGatewayGroupsPaginatorOptions is the paginator options for ListGatewayGroups
type ListGatewayGroupsPaginatorOptions struct {
	// The maximum number of gateway group summaries to return. The default is 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGatewayGroupsPaginator is a paginator for ListGatewayGroups
type ListGatewayGroupsPaginator struct {
	options   ListGatewayGroupsPaginatorOptions
	client    ListGatewayGroupsAPIClient
	params    *ListGatewayGroupsInput
	nextToken *string
	firstPage bool
}

// NewListGatewayGroupsPaginator returns a new ListGatewayGroupsPaginator
func NewListGatewayGroupsPaginator(client ListGatewayGroupsAPIClient, params *ListGatewayGroupsInput, optFns ...func(*ListGatewayGroupsPaginatorOptions)) *ListGatewayGroupsPaginator {
	if params == nil {
		params = &ListGatewayGroupsInput{}
	}

	options := ListGatewayGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGatewayGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGatewayGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGatewayGroups page.
func (p *ListGatewayGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGatewayGroupsOutput, error) {
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

	result, err := p.client.ListGatewayGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListGatewayGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListGatewayGroups",
	}
}
