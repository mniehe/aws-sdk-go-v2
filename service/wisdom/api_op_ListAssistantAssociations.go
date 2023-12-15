// Code generated by smithy-go-codegen DO NOT EDIT.

package wisdom

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/wisdom/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists information about assistant associations.
func (c *Client) ListAssistantAssociations(ctx context.Context, params *ListAssistantAssociationsInput, optFns ...func(*Options)) (*ListAssistantAssociationsOutput, error) {
	if params == nil {
		params = &ListAssistantAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssistantAssociations", params, optFns, c.addOperationListAssistantAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssistantAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssistantAssociationsInput struct {

	// The identifier of the Wisdom assistant. Can be either the ID or the ARN. URLs
	// cannot contain the ARN.
	//
	// This member is required.
	AssistantId *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAssistantAssociationsOutput struct {

	// Summary information about assistant associations.
	//
	// This member is required.
	AssistantAssociationSummaries []types.AssistantAssociationSummary

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssistantAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAssistantAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssistantAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAssistantAssociations"); err != nil {
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
	if err = addOpListAssistantAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssistantAssociations(options.Region), middleware.Before); err != nil {
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

// ListAssistantAssociationsAPIClient is a client that implements the
// ListAssistantAssociations operation.
type ListAssistantAssociationsAPIClient interface {
	ListAssistantAssociations(context.Context, *ListAssistantAssociationsInput, ...func(*Options)) (*ListAssistantAssociationsOutput, error)
}

var _ ListAssistantAssociationsAPIClient = (*Client)(nil)

// ListAssistantAssociationsPaginatorOptions is the paginator options for
// ListAssistantAssociations
type ListAssistantAssociationsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssistantAssociationsPaginator is a paginator for ListAssistantAssociations
type ListAssistantAssociationsPaginator struct {
	options   ListAssistantAssociationsPaginatorOptions
	client    ListAssistantAssociationsAPIClient
	params    *ListAssistantAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListAssistantAssociationsPaginator returns a new
// ListAssistantAssociationsPaginator
func NewListAssistantAssociationsPaginator(client ListAssistantAssociationsAPIClient, params *ListAssistantAssociationsInput, optFns ...func(*ListAssistantAssociationsPaginatorOptions)) *ListAssistantAssociationsPaginator {
	if params == nil {
		params = &ListAssistantAssociationsInput{}
	}

	options := ListAssistantAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssistantAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssistantAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssistantAssociations page.
func (p *ListAssistantAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssistantAssociationsOutput, error) {
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

	result, err := p.client.ListAssistantAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssistantAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAssistantAssociations",
	}
}
