// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of versions
// (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html), with the
// version-specific configuration of each. Lambda returns up to 50 versions per
// call.
func (c *Client) ListVersionsByFunction(ctx context.Context, params *ListVersionsByFunctionInput, optFns ...func(*Options)) (*ListVersionsByFunctionOutput, error) {
	if params == nil {
		params = &ListVersionsByFunctionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVersionsByFunction", params, optFns, addOperationListVersionsByFunctionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVersionsByFunctionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVersionsByFunctionInput struct {

	// The name of the Lambda function. Name formats
	//
	// * Function name - MyFunction.
	//
	// *
	// Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction.
	//
	// *
	// Partial ARN - 123456789012:function:MyFunction.
	//
	// The length constraint applies
	// only to the full ARN. If you specify only the function name, it is limited to 64
	// characters in length.
	//
	// This member is required.
	FunctionName *string

	// Specify the pagination token that's returned by a previous request to retrieve
	// the next page of results.
	Marker *string

	// The maximum number of versions to return.
	MaxItems *int32
}

type ListVersionsByFunctionOutput struct {

	// The pagination token that's included if more results are available.
	NextMarker *string

	// A list of Lambda function versions.
	Versions []types.FunctionConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListVersionsByFunctionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVersionsByFunction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVersionsByFunction{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListVersionsByFunctionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVersionsByFunction(options.Region), middleware.Before); err != nil {
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
	return nil
}

// ListVersionsByFunctionAPIClient is a client that implements the
// ListVersionsByFunction operation.
type ListVersionsByFunctionAPIClient interface {
	ListVersionsByFunction(context.Context, *ListVersionsByFunctionInput, ...func(*Options)) (*ListVersionsByFunctionOutput, error)
}

var _ ListVersionsByFunctionAPIClient = (*Client)(nil)

// ListVersionsByFunctionPaginatorOptions is the paginator options for
// ListVersionsByFunction
type ListVersionsByFunctionPaginatorOptions struct {
	// The maximum number of versions to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVersionsByFunctionPaginator is a paginator for ListVersionsByFunction
type ListVersionsByFunctionPaginator struct {
	options   ListVersionsByFunctionPaginatorOptions
	client    ListVersionsByFunctionAPIClient
	params    *ListVersionsByFunctionInput
	nextToken *string
	firstPage bool
}

// NewListVersionsByFunctionPaginator returns a new ListVersionsByFunctionPaginator
func NewListVersionsByFunctionPaginator(client ListVersionsByFunctionAPIClient, params *ListVersionsByFunctionInput, optFns ...func(*ListVersionsByFunctionPaginatorOptions)) *ListVersionsByFunctionPaginator {
	options := ListVersionsByFunctionPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListVersionsByFunctionInput{}
	}

	return &ListVersionsByFunctionPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVersionsByFunctionPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListVersionsByFunction page.
func (p *ListVersionsByFunctionPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVersionsByFunctionOutput, error) {
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

	result, err := p.client.ListVersionsByFunction(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListVersionsByFunction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "ListVersionsByFunction",
	}
}
