// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// An invocation is copy of a command sent to a specific managed node. A command
// can apply to one or more managed nodes. A command invocation applies to one
// managed node. For example, if a user runs SendCommand against three managed
// nodes, then a command invocation is created for each requested managed node ID.
// ListCommandInvocations provide status about command execution.
func (c *Client) ListCommandInvocations(ctx context.Context, params *ListCommandInvocationsInput, optFns ...func(*Options)) (*ListCommandInvocationsOutput, error) {
	if params == nil {
		params = &ListCommandInvocationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCommandInvocations", params, optFns, c.addOperationListCommandInvocationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCommandInvocationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCommandInvocationsInput struct {

	// (Optional) The invocations for a specific command ID.
	CommandId *string

	// (Optional) If set this returns the response of the command executions and any
	// command output. The default value is false .
	Details bool

	// (Optional) One or more filters. Use a filter to return a more specific list of
	// results.
	Filters []types.CommandFilter

	// (Optional) The command execution details for a specific managed node ID.
	InstanceId *string

	// (Optional) The maximum number of items to return for this call. The call also
	// returns a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int32

	// (Optional) The token for the next set of items to return. (You received this
	// token from a previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type ListCommandInvocationsOutput struct {

	// (Optional) A list of all invocations.
	CommandInvocations []types.CommandInvocation

	// (Optional) The token for the next set of items to return. (You received this
	// token from a previous call.)
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCommandInvocationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCommandInvocations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCommandInvocations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCommandInvocations"); err != nil {
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
	if err = addOpListCommandInvocationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCommandInvocations(options.Region), middleware.Before); err != nil {
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

// ListCommandInvocationsAPIClient is a client that implements the
// ListCommandInvocations operation.
type ListCommandInvocationsAPIClient interface {
	ListCommandInvocations(context.Context, *ListCommandInvocationsInput, ...func(*Options)) (*ListCommandInvocationsOutput, error)
}

var _ ListCommandInvocationsAPIClient = (*Client)(nil)

// ListCommandInvocationsPaginatorOptions is the paginator options for
// ListCommandInvocations
type ListCommandInvocationsPaginatorOptions struct {
	// (Optional) The maximum number of items to return for this call. The call also
	// returns a token that you can specify in a subsequent call to get the next set of
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCommandInvocationsPaginator is a paginator for ListCommandInvocations
type ListCommandInvocationsPaginator struct {
	options   ListCommandInvocationsPaginatorOptions
	client    ListCommandInvocationsAPIClient
	params    *ListCommandInvocationsInput
	nextToken *string
	firstPage bool
}

// NewListCommandInvocationsPaginator returns a new ListCommandInvocationsPaginator
func NewListCommandInvocationsPaginator(client ListCommandInvocationsAPIClient, params *ListCommandInvocationsInput, optFns ...func(*ListCommandInvocationsPaginatorOptions)) *ListCommandInvocationsPaginator {
	if params == nil {
		params = &ListCommandInvocationsInput{}
	}

	options := ListCommandInvocationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCommandInvocationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCommandInvocationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCommandInvocations page.
func (p *ListCommandInvocationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCommandInvocationsOutput, error) {
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

	result, err := p.client.ListCommandInvocations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCommandInvocations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCommandInvocations",
	}
}
