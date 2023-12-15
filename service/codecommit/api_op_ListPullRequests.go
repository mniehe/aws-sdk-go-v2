// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of pull requests for a specified repository. The return list can
// be refined by pull request status or pull request author ARN.
func (c *Client) ListPullRequests(ctx context.Context, params *ListPullRequestsInput, optFns ...func(*Options)) (*ListPullRequestsOutput, error) {
	if params == nil {
		params = &ListPullRequestsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPullRequests", params, optFns, c.addOperationListPullRequestsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPullRequestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPullRequestsInput struct {

	// The name of the repository for which you want to list pull requests.
	//
	// This member is required.
	RepositoryName *string

	// Optional. The Amazon Resource Name (ARN) of the user who created the pull
	// request. If used, this filters the results to pull requests created by that
	// user.
	AuthorArn *string

	// A non-zero, non-negative integer used to limit the number of returned results.
	MaxResults *int32

	// An enumeration token that, when provided in a request, returns the next batch
	// of the results.
	NextToken *string

	// Optional. The status of the pull request. If used, this refines the results to
	// the pull requests that match the specified status.
	PullRequestStatus types.PullRequestStatusEnum

	noSmithyDocumentSerde
}

type ListPullRequestsOutput struct {

	// The system-generated IDs of the pull requests.
	//
	// This member is required.
	PullRequestIds []string

	// An enumeration token that allows the operation to batch the next results of the
	// operation.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPullRequestsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPullRequests{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPullRequests{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPullRequests"); err != nil {
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
	if err = addOpListPullRequestsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPullRequests(options.Region), middleware.Before); err != nil {
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

// ListPullRequestsAPIClient is a client that implements the ListPullRequests
// operation.
type ListPullRequestsAPIClient interface {
	ListPullRequests(context.Context, *ListPullRequestsInput, ...func(*Options)) (*ListPullRequestsOutput, error)
}

var _ ListPullRequestsAPIClient = (*Client)(nil)

// ListPullRequestsPaginatorOptions is the paginator options for ListPullRequests
type ListPullRequestsPaginatorOptions struct {
	// A non-zero, non-negative integer used to limit the number of returned results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPullRequestsPaginator is a paginator for ListPullRequests
type ListPullRequestsPaginator struct {
	options   ListPullRequestsPaginatorOptions
	client    ListPullRequestsAPIClient
	params    *ListPullRequestsInput
	nextToken *string
	firstPage bool
}

// NewListPullRequestsPaginator returns a new ListPullRequestsPaginator
func NewListPullRequestsPaginator(client ListPullRequestsAPIClient, params *ListPullRequestsInput, optFns ...func(*ListPullRequestsPaginatorOptions)) *ListPullRequestsPaginator {
	if params == nil {
		params = &ListPullRequestsInput{}
	}

	options := ListPullRequestsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPullRequestsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPullRequestsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPullRequests page.
func (p *ListPullRequestsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPullRequestsOutput, error) {
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

	result, err := p.client.ListPullRequests(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPullRequests(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPullRequests",
	}
}
