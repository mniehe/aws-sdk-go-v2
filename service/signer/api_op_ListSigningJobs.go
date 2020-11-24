// Code generated by smithy-go-codegen DO NOT EDIT.

package signer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/signer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all your signing jobs. You can use the maxResults parameter to limit the
// number of signing jobs that are returned in the response. If additional jobs
// remain to be listed, code signing returns a nextToken value. Use this value in
// subsequent calls to ListSigningJobs to fetch the remaining values. You can
// continue calling ListSigningJobs with your maxResults parameter and with new
// values that code signing returns in the nextToken parameter until all of your
// signing jobs have been returned.
func (c *Client) ListSigningJobs(ctx context.Context, params *ListSigningJobsInput, optFns ...func(*Options)) (*ListSigningJobsOutput, error) {
	if params == nil {
		params = &ListSigningJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSigningJobs", params, optFns, addOperationListSigningJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSigningJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSigningJobsInput struct {

	// Specifies the maximum number of items to return in the response. Use this
	// parameter when paginating results. If additional items exist beyond the number
	// you specify, the nextToken element is set in the response. Use the nextToken
	// value in a subsequent request to retrieve additional items.
	MaxResults *int32

	// String for specifying the next set of paginated results to return. After you
	// receive a response with truncated results, use this parameter in a subsequent
	// request. Set it to the value of nextToken from the response that you just
	// received.
	NextToken *string

	// The ID of microcontroller platform that you specified for the distribution of
	// your code image.
	PlatformId *string

	// The IAM principal that requested the signing job.
	RequestedBy *string

	// A status value with which to filter your results.
	Status types.SigningStatus
}

type ListSigningJobsOutput struct {

	// A list of your signing jobs.
	Jobs []types.SigningJob

	// String for specifying the next set of paginated results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListSigningJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSigningJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSigningJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSigningJobs(options.Region), middleware.Before); err != nil {
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

// ListSigningJobsAPIClient is a client that implements the ListSigningJobs
// operation.
type ListSigningJobsAPIClient interface {
	ListSigningJobs(context.Context, *ListSigningJobsInput, ...func(*Options)) (*ListSigningJobsOutput, error)
}

var _ ListSigningJobsAPIClient = (*Client)(nil)

// ListSigningJobsPaginatorOptions is the paginator options for ListSigningJobs
type ListSigningJobsPaginatorOptions struct {
	// Specifies the maximum number of items to return in the response. Use this
	// parameter when paginating results. If additional items exist beyond the number
	// you specify, the nextToken element is set in the response. Use the nextToken
	// value in a subsequent request to retrieve additional items.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSigningJobsPaginator is a paginator for ListSigningJobs
type ListSigningJobsPaginator struct {
	options   ListSigningJobsPaginatorOptions
	client    ListSigningJobsAPIClient
	params    *ListSigningJobsInput
	nextToken *string
	firstPage bool
}

// NewListSigningJobsPaginator returns a new ListSigningJobsPaginator
func NewListSigningJobsPaginator(client ListSigningJobsAPIClient, params *ListSigningJobsInput, optFns ...func(*ListSigningJobsPaginatorOptions)) *ListSigningJobsPaginator {
	options := ListSigningJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListSigningJobsInput{}
	}

	return &ListSigningJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSigningJobsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListSigningJobs page.
func (p *ListSigningJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSigningJobsOutput, error) {
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

	result, err := p.client.ListSigningJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSigningJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "signer",
		OperationName: "ListSigningJobs",
	}
}
