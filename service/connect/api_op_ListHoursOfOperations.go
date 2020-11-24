// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides information about the hours of operation for the specified Amazon
// Connect instance. For more information about hours of operation, see Set the
// Hours of Operation for a Queue
// (https://docs.aws.amazon.com/connect/latest/adminguide/set-hours-operation.html)
// in the Amazon Connect Administrator Guide.
func (c *Client) ListHoursOfOperations(ctx context.Context, params *ListHoursOfOperationsInput, optFns ...func(*Options)) (*ListHoursOfOperationsOutput, error) {
	if params == nil {
		params = &ListHoursOfOperationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHoursOfOperations", params, optFns, addOperationListHoursOfOperationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHoursOfOperationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHoursOfOperationsInput struct {

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The maximimum number of results to return per page.
	MaxResults int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string
}

type ListHoursOfOperationsOutput struct {

	// Information about the hours of operation.
	HoursOfOperationSummaryList []types.HoursOfOperationSummary

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListHoursOfOperationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListHoursOfOperations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListHoursOfOperations{}, middleware.After)
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
	if err = addOpListHoursOfOperationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHoursOfOperations(options.Region), middleware.Before); err != nil {
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

// ListHoursOfOperationsAPIClient is a client that implements the
// ListHoursOfOperations operation.
type ListHoursOfOperationsAPIClient interface {
	ListHoursOfOperations(context.Context, *ListHoursOfOperationsInput, ...func(*Options)) (*ListHoursOfOperationsOutput, error)
}

var _ ListHoursOfOperationsAPIClient = (*Client)(nil)

// ListHoursOfOperationsPaginatorOptions is the paginator options for
// ListHoursOfOperations
type ListHoursOfOperationsPaginatorOptions struct {
	// The maximimum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListHoursOfOperationsPaginator is a paginator for ListHoursOfOperations
type ListHoursOfOperationsPaginator struct {
	options   ListHoursOfOperationsPaginatorOptions
	client    ListHoursOfOperationsAPIClient
	params    *ListHoursOfOperationsInput
	nextToken *string
	firstPage bool
}

// NewListHoursOfOperationsPaginator returns a new ListHoursOfOperationsPaginator
func NewListHoursOfOperationsPaginator(client ListHoursOfOperationsAPIClient, params *ListHoursOfOperationsInput, optFns ...func(*ListHoursOfOperationsPaginatorOptions)) *ListHoursOfOperationsPaginator {
	options := ListHoursOfOperationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListHoursOfOperationsInput{}
	}

	return &ListHoursOfOperationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListHoursOfOperationsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListHoursOfOperations page.
func (p *ListHoursOfOperationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListHoursOfOperationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListHoursOfOperations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListHoursOfOperations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListHoursOfOperations",
	}
}
