// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides information about the flow modules for the specified Amazon Connect
// instance.
func (c *Client) ListContactFlowModules(ctx context.Context, params *ListContactFlowModulesInput, optFns ...func(*Options)) (*ListContactFlowModulesOutput, error) {
	if params == nil {
		params = &ListContactFlowModulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListContactFlowModules", params, optFns, c.addOperationListContactFlowModulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListContactFlowModulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListContactFlowModulesInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The state of the flow module.
	ContactFlowModuleState types.ContactFlowModuleState

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListContactFlowModulesOutput struct {

	// Information about the flow module.
	ContactFlowModulesSummaryList []types.ContactFlowModuleSummary

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListContactFlowModulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListContactFlowModules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListContactFlowModules{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListContactFlowModules"); err != nil {
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
	if err = addOpListContactFlowModulesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListContactFlowModules(options.Region), middleware.Before); err != nil {
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

// ListContactFlowModulesAPIClient is a client that implements the
// ListContactFlowModules operation.
type ListContactFlowModulesAPIClient interface {
	ListContactFlowModules(context.Context, *ListContactFlowModulesInput, ...func(*Options)) (*ListContactFlowModulesOutput, error)
}

var _ ListContactFlowModulesAPIClient = (*Client)(nil)

// ListContactFlowModulesPaginatorOptions is the paginator options for
// ListContactFlowModules
type ListContactFlowModulesPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListContactFlowModulesPaginator is a paginator for ListContactFlowModules
type ListContactFlowModulesPaginator struct {
	options   ListContactFlowModulesPaginatorOptions
	client    ListContactFlowModulesAPIClient
	params    *ListContactFlowModulesInput
	nextToken *string
	firstPage bool
}

// NewListContactFlowModulesPaginator returns a new ListContactFlowModulesPaginator
func NewListContactFlowModulesPaginator(client ListContactFlowModulesAPIClient, params *ListContactFlowModulesInput, optFns ...func(*ListContactFlowModulesPaginatorOptions)) *ListContactFlowModulesPaginator {
	if params == nil {
		params = &ListContactFlowModulesInput{}
	}

	options := ListContactFlowModulesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListContactFlowModulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListContactFlowModulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListContactFlowModules page.
func (p *ListContactFlowModulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListContactFlowModulesOutput, error) {
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

	result, err := p.client.ListContactFlowModules(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListContactFlowModules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListContactFlowModules",
	}
}
