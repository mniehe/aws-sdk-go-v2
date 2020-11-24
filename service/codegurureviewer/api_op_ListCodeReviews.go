// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurureviewer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all the code reviews that the customer has created in the past 90 days.
func (c *Client) ListCodeReviews(ctx context.Context, params *ListCodeReviewsInput, optFns ...func(*Options)) (*ListCodeReviewsOutput, error) {
	if params == nil {
		params = &ListCodeReviewsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCodeReviews", params, optFns, addOperationListCodeReviewsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCodeReviewsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCodeReviewsInput struct {

	// The type of code reviews to list in the response.
	//
	// This member is required.
	Type types.Type

	// The maximum number of results that are returned per call. The default is 100.
	MaxResults *int32

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged.
	NextToken *string

	// List of provider types for filtering that needs to be applied before displaying
	// the result. For example, providerTypes=[GitHub] lists code reviews from GitHub.
	ProviderTypes []types.ProviderType

	// List of repository names for filtering that needs to be applied before
	// displaying the result.
	RepositoryNames []string

	// List of states for filtering that needs to be applied before displaying the
	// result. For example, states=[Pending] lists code reviews in the Pending state.
	// The valid code review states are:
	//
	// * Completed: The code review is complete.
	//
	// *
	// Pending: The code review started and has not completed or failed.
	//
	// * Failed: The
	// code review failed.
	//
	// * Deleting: The code review is being deleted.
	States []types.JobState
}

type ListCodeReviewsOutput struct {

	// A list of code reviews that meet the criteria of the request.
	CodeReviewSummaries []types.CodeReviewSummary

	// Pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListCodeReviewsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCodeReviews{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCodeReviews{}, middleware.After)
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
	if err = addOpListCodeReviewsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCodeReviews(options.Region), middleware.Before); err != nil {
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

// ListCodeReviewsAPIClient is a client that implements the ListCodeReviews
// operation.
type ListCodeReviewsAPIClient interface {
	ListCodeReviews(context.Context, *ListCodeReviewsInput, ...func(*Options)) (*ListCodeReviewsOutput, error)
}

var _ ListCodeReviewsAPIClient = (*Client)(nil)

// ListCodeReviewsPaginatorOptions is the paginator options for ListCodeReviews
type ListCodeReviewsPaginatorOptions struct {
	// The maximum number of results that are returned per call. The default is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCodeReviewsPaginator is a paginator for ListCodeReviews
type ListCodeReviewsPaginator struct {
	options   ListCodeReviewsPaginatorOptions
	client    ListCodeReviewsAPIClient
	params    *ListCodeReviewsInput
	nextToken *string
	firstPage bool
}

// NewListCodeReviewsPaginator returns a new ListCodeReviewsPaginator
func NewListCodeReviewsPaginator(client ListCodeReviewsAPIClient, params *ListCodeReviewsInput, optFns ...func(*ListCodeReviewsPaginatorOptions)) *ListCodeReviewsPaginator {
	options := ListCodeReviewsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListCodeReviewsInput{}
	}

	return &ListCodeReviewsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCodeReviewsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListCodeReviews page.
func (p *ListCodeReviewsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCodeReviewsOutput, error) {
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

	result, err := p.client.ListCodeReviews(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCodeReviews(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-reviewer",
		OperationName: "ListCodeReviews",
	}
}
