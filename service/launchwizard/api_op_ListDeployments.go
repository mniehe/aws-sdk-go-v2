// Code generated by smithy-go-codegen DO NOT EDIT.

package launchwizard

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/launchwizard/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the deployments that have been created.
func (c *Client) ListDeployments(ctx context.Context, params *ListDeploymentsInput, optFns ...func(*Options)) (*ListDeploymentsOutput, error) {
	if params == nil {
		params = &ListDeploymentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDeployments", params, optFns, c.addOperationListDeploymentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDeploymentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDeploymentsInput struct {

	// Filters to scope the results. The following filters are supported:
	//   - WORKLOAD_NAME
	//   - DEPLOYMENT_STATUS
	Filters []types.DeploymentFilter

	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output.
	MaxResults *int32

	// The token returned from a previous paginated request. Pagination continues from
	// the end of the items returned by the previous request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDeploymentsOutput struct {

	// Lists the deployments.
	Deployments []types.DeploymentDataSummary

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDeploymentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDeployments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDeployments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDeployments"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDeployments(options.Region), middleware.Before); err != nil {
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

// ListDeploymentsAPIClient is a client that implements the ListDeployments
// operation.
type ListDeploymentsAPIClient interface {
	ListDeployments(context.Context, *ListDeploymentsInput, ...func(*Options)) (*ListDeploymentsOutput, error)
}

var _ ListDeploymentsAPIClient = (*Client)(nil)

// ListDeploymentsPaginatorOptions is the paginator options for ListDeployments
type ListDeploymentsPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDeploymentsPaginator is a paginator for ListDeployments
type ListDeploymentsPaginator struct {
	options   ListDeploymentsPaginatorOptions
	client    ListDeploymentsAPIClient
	params    *ListDeploymentsInput
	nextToken *string
	firstPage bool
}

// NewListDeploymentsPaginator returns a new ListDeploymentsPaginator
func NewListDeploymentsPaginator(client ListDeploymentsAPIClient, params *ListDeploymentsInput, optFns ...func(*ListDeploymentsPaginatorOptions)) *ListDeploymentsPaginator {
	if params == nil {
		params = &ListDeploymentsInput{}
	}

	options := ListDeploymentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDeploymentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDeploymentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDeployments page.
func (p *ListDeploymentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDeploymentsOutput, error) {
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

	result, err := p.client.ListDeployments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDeployments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDeployments",
	}
}
