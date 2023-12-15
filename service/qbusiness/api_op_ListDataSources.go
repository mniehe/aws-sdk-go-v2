// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the Amazon Q data source connectors that you have created.
func (c *Client) ListDataSources(ctx context.Context, params *ListDataSourcesInput, optFns ...func(*Options)) (*ListDataSourcesOutput, error) {
	if params == nil {
		params = &ListDataSourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataSources", params, optFns, c.addOperationListDataSourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataSourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataSourcesInput struct {

	// The identifier of the Amazon Q application linked to the data source connectors.
	//
	// This member is required.
	ApplicationId *string

	// The identifier of the index used with one or more data source connectors.
	//
	// This member is required.
	IndexId *string

	// The maximum number of data source connectors to return.
	MaxResults *int32

	// If the maxResults response was incomplete because there is more data to
	// retrieve, Amazon Q returns a pagination token in the response. You can use this
	// pagination token to retrieve the next set of Amazon Q data source connectors.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDataSourcesOutput struct {

	// An array of summary information for one or more data source connector.
	DataSources []types.DataSource

	// If the response is truncated, Amazon Q returns this token. You can use this
	// token in a subsequent request to retrieve the next set of data source
	// connectors.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDataSourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDataSources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDataSources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDataSources"); err != nil {
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
	if err = addOpListDataSourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDataSources(options.Region), middleware.Before); err != nil {
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

// ListDataSourcesAPIClient is a client that implements the ListDataSources
// operation.
type ListDataSourcesAPIClient interface {
	ListDataSources(context.Context, *ListDataSourcesInput, ...func(*Options)) (*ListDataSourcesOutput, error)
}

var _ ListDataSourcesAPIClient = (*Client)(nil)

// ListDataSourcesPaginatorOptions is the paginator options for ListDataSources
type ListDataSourcesPaginatorOptions struct {
	// The maximum number of data source connectors to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDataSourcesPaginator is a paginator for ListDataSources
type ListDataSourcesPaginator struct {
	options   ListDataSourcesPaginatorOptions
	client    ListDataSourcesAPIClient
	params    *ListDataSourcesInput
	nextToken *string
	firstPage bool
}

// NewListDataSourcesPaginator returns a new ListDataSourcesPaginator
func NewListDataSourcesPaginator(client ListDataSourcesAPIClient, params *ListDataSourcesInput, optFns ...func(*ListDataSourcesPaginatorOptions)) *ListDataSourcesPaginator {
	if params == nil {
		params = &ListDataSourcesInput{}
	}

	options := ListDataSourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDataSourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDataSourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDataSources page.
func (p *ListDataSourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDataSourcesOutput, error) {
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

	result, err := p.client.ListDataSources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDataSources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDataSources",
	}
}
