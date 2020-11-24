// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about test suites for a given job.
func (c *Client) ListSuites(ctx context.Context, params *ListSuitesInput, optFns ...func(*Options)) (*ListSuitesOutput, error) {
	if params == nil {
		params = &ListSuitesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSuites", params, optFns, addOperationListSuitesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSuitesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to the list suites operation.
type ListSuitesInput struct {

	// The job's Amazon Resource Name (ARN).
	//
	// This member is required.
	Arn *string

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string
}

// Represents the result of a list suites request.
type ListSuitesOutput struct {

	// If the number of items that are returned is significantly large, this is an
	// identifier that is also returned. It can be used in a subsequent call to this
	// operation to return the next set of items in the list.
	NextToken *string

	// Information about the suites.
	Suites []types.Suite

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListSuitesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSuites{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSuites{}, middleware.After)
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
	if err = addOpListSuitesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSuites(options.Region), middleware.Before); err != nil {
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

// ListSuitesAPIClient is a client that implements the ListSuites operation.
type ListSuitesAPIClient interface {
	ListSuites(context.Context, *ListSuitesInput, ...func(*Options)) (*ListSuitesOutput, error)
}

var _ ListSuitesAPIClient = (*Client)(nil)

// ListSuitesPaginatorOptions is the paginator options for ListSuites
type ListSuitesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSuitesPaginator is a paginator for ListSuites
type ListSuitesPaginator struct {
	options   ListSuitesPaginatorOptions
	client    ListSuitesAPIClient
	params    *ListSuitesInput
	nextToken *string
	firstPage bool
}

// NewListSuitesPaginator returns a new ListSuitesPaginator
func NewListSuitesPaginator(client ListSuitesAPIClient, params *ListSuitesInput, optFns ...func(*ListSuitesPaginatorOptions)) *ListSuitesPaginator {
	options := ListSuitesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListSuitesInput{}
	}

	return &ListSuitesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSuitesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListSuites page.
func (p *ListSuitesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSuitesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListSuites(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSuites(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListSuites",
	}
}
