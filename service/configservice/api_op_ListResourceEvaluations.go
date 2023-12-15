// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of proactive resource evaluations.
func (c *Client) ListResourceEvaluations(ctx context.Context, params *ListResourceEvaluationsInput, optFns ...func(*Options)) (*ListResourceEvaluationsOutput, error) {
	if params == nil {
		params = &ListResourceEvaluationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResourceEvaluations", params, optFns, c.addOperationListResourceEvaluationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourceEvaluationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourceEvaluationsInput struct {

	// Returns a ResourceEvaluationFilters object.
	Filters *types.ResourceEvaluationFilters

	// The maximum number of evaluations returned on each page. The default is 10. You
	// cannot specify a number greater than 100. If you specify 0, Config uses the
	// default.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListResourceEvaluationsOutput struct {

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Returns a ResourceEvaluations object.
	ResourceEvaluations []types.ResourceEvaluation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResourceEvaluationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResourceEvaluations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResourceEvaluations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListResourceEvaluations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResourceEvaluations(options.Region), middleware.Before); err != nil {
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

// ListResourceEvaluationsAPIClient is a client that implements the
// ListResourceEvaluations operation.
type ListResourceEvaluationsAPIClient interface {
	ListResourceEvaluations(context.Context, *ListResourceEvaluationsInput, ...func(*Options)) (*ListResourceEvaluationsOutput, error)
}

var _ ListResourceEvaluationsAPIClient = (*Client)(nil)

// ListResourceEvaluationsPaginatorOptions is the paginator options for
// ListResourceEvaluations
type ListResourceEvaluationsPaginatorOptions struct {
	// The maximum number of evaluations returned on each page. The default is 10. You
	// cannot specify a number greater than 100. If you specify 0, Config uses the
	// default.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResourceEvaluationsPaginator is a paginator for ListResourceEvaluations
type ListResourceEvaluationsPaginator struct {
	options   ListResourceEvaluationsPaginatorOptions
	client    ListResourceEvaluationsAPIClient
	params    *ListResourceEvaluationsInput
	nextToken *string
	firstPage bool
}

// NewListResourceEvaluationsPaginator returns a new
// ListResourceEvaluationsPaginator
func NewListResourceEvaluationsPaginator(client ListResourceEvaluationsAPIClient, params *ListResourceEvaluationsInput, optFns ...func(*ListResourceEvaluationsPaginatorOptions)) *ListResourceEvaluationsPaginator {
	if params == nil {
		params = &ListResourceEvaluationsInput{}
	}

	options := ListResourceEvaluationsPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResourceEvaluationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResourceEvaluationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResourceEvaluations page.
func (p *ListResourceEvaluationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResourceEvaluationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.ListResourceEvaluations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListResourceEvaluations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListResourceEvaluations",
	}
}
