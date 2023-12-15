// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of provisioned concurrency configurations for a function.
func (c *Client) ListProvisionedConcurrencyConfigs(ctx context.Context, params *ListProvisionedConcurrencyConfigsInput, optFns ...func(*Options)) (*ListProvisionedConcurrencyConfigsOutput, error) {
	if params == nil {
		params = &ListProvisionedConcurrencyConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProvisionedConcurrencyConfigs", params, optFns, c.addOperationListProvisionedConcurrencyConfigsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProvisionedConcurrencyConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProvisionedConcurrencyConfigsInput struct {

	// The name of the Lambda function. Name formats
	//   - Function name – my-function .
	//   - Function ARN – arn:aws:lambda:us-west-2:123456789012:function:my-function .
	//   - Partial ARN – 123456789012:function:my-function .
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// Specify the pagination token that's returned by a previous request to retrieve
	// the next page of results.
	Marker *string

	// Specify a number to limit the number of configurations returned.
	MaxItems *int32

	noSmithyDocumentSerde
}

type ListProvisionedConcurrencyConfigsOutput struct {

	// The pagination token that's included if more results are available.
	NextMarker *string

	// A list of provisioned concurrency configurations.
	ProvisionedConcurrencyConfigs []types.ProvisionedConcurrencyConfigListItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProvisionedConcurrencyConfigsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProvisionedConcurrencyConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProvisionedConcurrencyConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListProvisionedConcurrencyConfigs"); err != nil {
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
	if err = addOpListProvisionedConcurrencyConfigsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProvisionedConcurrencyConfigs(options.Region), middleware.Before); err != nil {
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

// ListProvisionedConcurrencyConfigsAPIClient is a client that implements the
// ListProvisionedConcurrencyConfigs operation.
type ListProvisionedConcurrencyConfigsAPIClient interface {
	ListProvisionedConcurrencyConfigs(context.Context, *ListProvisionedConcurrencyConfigsInput, ...func(*Options)) (*ListProvisionedConcurrencyConfigsOutput, error)
}

var _ ListProvisionedConcurrencyConfigsAPIClient = (*Client)(nil)

// ListProvisionedConcurrencyConfigsPaginatorOptions is the paginator options for
// ListProvisionedConcurrencyConfigs
type ListProvisionedConcurrencyConfigsPaginatorOptions struct {
	// Specify a number to limit the number of configurations returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProvisionedConcurrencyConfigsPaginator is a paginator for
// ListProvisionedConcurrencyConfigs
type ListProvisionedConcurrencyConfigsPaginator struct {
	options   ListProvisionedConcurrencyConfigsPaginatorOptions
	client    ListProvisionedConcurrencyConfigsAPIClient
	params    *ListProvisionedConcurrencyConfigsInput
	nextToken *string
	firstPage bool
}

// NewListProvisionedConcurrencyConfigsPaginator returns a new
// ListProvisionedConcurrencyConfigsPaginator
func NewListProvisionedConcurrencyConfigsPaginator(client ListProvisionedConcurrencyConfigsAPIClient, params *ListProvisionedConcurrencyConfigsInput, optFns ...func(*ListProvisionedConcurrencyConfigsPaginatorOptions)) *ListProvisionedConcurrencyConfigsPaginator {
	if params == nil {
		params = &ListProvisionedConcurrencyConfigsInput{}
	}

	options := ListProvisionedConcurrencyConfigsPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProvisionedConcurrencyConfigsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProvisionedConcurrencyConfigsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProvisionedConcurrencyConfigs page.
func (p *ListProvisionedConcurrencyConfigsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProvisionedConcurrencyConfigsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListProvisionedConcurrencyConfigs(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListProvisionedConcurrencyConfigs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListProvisionedConcurrencyConfigs",
	}
}
