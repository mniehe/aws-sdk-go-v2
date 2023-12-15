// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get a list of key phrase detection jobs that you have submitted.
func (c *Client) ListKeyPhrasesDetectionJobs(ctx context.Context, params *ListKeyPhrasesDetectionJobsInput, optFns ...func(*Options)) (*ListKeyPhrasesDetectionJobsOutput, error) {
	if params == nil {
		params = &ListKeyPhrasesDetectionJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListKeyPhrasesDetectionJobs", params, optFns, c.addOperationListKeyPhrasesDetectionJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListKeyPhrasesDetectionJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListKeyPhrasesDetectionJobsInput struct {

	// Filters the jobs that are returned. You can filter jobs on their name, status,
	// or the date and time that they were submitted. You can only set one filter at a
	// time.
	Filter *types.KeyPhrasesDetectionJobFilter

	// The maximum number of results to return in each page. The default is 100.
	MaxResults *int32

	// Identifies the next page of results to return.
	NextToken *string

	noSmithyDocumentSerde
}

type ListKeyPhrasesDetectionJobsOutput struct {

	// A list containing the properties of each job that is returned.
	KeyPhrasesDetectionJobPropertiesList []types.KeyPhrasesDetectionJobProperties

	// Identifies the next page of results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListKeyPhrasesDetectionJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListKeyPhrasesDetectionJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListKeyPhrasesDetectionJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListKeyPhrasesDetectionJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListKeyPhrasesDetectionJobs(options.Region), middleware.Before); err != nil {
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

// ListKeyPhrasesDetectionJobsAPIClient is a client that implements the
// ListKeyPhrasesDetectionJobs operation.
type ListKeyPhrasesDetectionJobsAPIClient interface {
	ListKeyPhrasesDetectionJobs(context.Context, *ListKeyPhrasesDetectionJobsInput, ...func(*Options)) (*ListKeyPhrasesDetectionJobsOutput, error)
}

var _ ListKeyPhrasesDetectionJobsAPIClient = (*Client)(nil)

// ListKeyPhrasesDetectionJobsPaginatorOptions is the paginator options for
// ListKeyPhrasesDetectionJobs
type ListKeyPhrasesDetectionJobsPaginatorOptions struct {
	// The maximum number of results to return in each page. The default is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListKeyPhrasesDetectionJobsPaginator is a paginator for
// ListKeyPhrasesDetectionJobs
type ListKeyPhrasesDetectionJobsPaginator struct {
	options   ListKeyPhrasesDetectionJobsPaginatorOptions
	client    ListKeyPhrasesDetectionJobsAPIClient
	params    *ListKeyPhrasesDetectionJobsInput
	nextToken *string
	firstPage bool
}

// NewListKeyPhrasesDetectionJobsPaginator returns a new
// ListKeyPhrasesDetectionJobsPaginator
func NewListKeyPhrasesDetectionJobsPaginator(client ListKeyPhrasesDetectionJobsAPIClient, params *ListKeyPhrasesDetectionJobsInput, optFns ...func(*ListKeyPhrasesDetectionJobsPaginatorOptions)) *ListKeyPhrasesDetectionJobsPaginator {
	if params == nil {
		params = &ListKeyPhrasesDetectionJobsInput{}
	}

	options := ListKeyPhrasesDetectionJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListKeyPhrasesDetectionJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListKeyPhrasesDetectionJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListKeyPhrasesDetectionJobs page.
func (p *ListKeyPhrasesDetectionJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListKeyPhrasesDetectionJobsOutput, error) {
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

	result, err := p.client.ListKeyPhrasesDetectionJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListKeyPhrasesDetectionJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListKeyPhrasesDetectionJobs",
	}
}
