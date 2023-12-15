// Code generated by smithy-go-codegen DO NOT EDIT.

package scheduler

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/scheduler/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a paginated list of your schedule groups.
func (c *Client) ListScheduleGroups(ctx context.Context, params *ListScheduleGroupsInput, optFns ...func(*Options)) (*ListScheduleGroupsOutput, error) {
	if params == nil {
		params = &ListScheduleGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListScheduleGroups", params, optFns, c.addOperationListScheduleGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListScheduleGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListScheduleGroupsInput struct {

	// If specified, limits the number of results returned by this operation. The
	// operation also returns a NextToken which you can use in a subsequent operation
	// to retrieve the next set of results.
	MaxResults *int32

	// The name prefix that you can use to return a filtered list of your schedule
	// groups.
	NamePrefix *string

	// The token returned by a previous call to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListScheduleGroupsOutput struct {

	// The schedule groups that match the specified criteria.
	//
	// This member is required.
	ScheduleGroups []types.ScheduleGroupSummary

	// Indicates whether there are additional results to retrieve. If the value is
	// null, there are no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListScheduleGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListScheduleGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListScheduleGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListScheduleGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListScheduleGroups(options.Region), middleware.Before); err != nil {
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

// ListScheduleGroupsAPIClient is a client that implements the ListScheduleGroups
// operation.
type ListScheduleGroupsAPIClient interface {
	ListScheduleGroups(context.Context, *ListScheduleGroupsInput, ...func(*Options)) (*ListScheduleGroupsOutput, error)
}

var _ ListScheduleGroupsAPIClient = (*Client)(nil)

// ListScheduleGroupsPaginatorOptions is the paginator options for
// ListScheduleGroups
type ListScheduleGroupsPaginatorOptions struct {
	// If specified, limits the number of results returned by this operation. The
	// operation also returns a NextToken which you can use in a subsequent operation
	// to retrieve the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListScheduleGroupsPaginator is a paginator for ListScheduleGroups
type ListScheduleGroupsPaginator struct {
	options   ListScheduleGroupsPaginatorOptions
	client    ListScheduleGroupsAPIClient
	params    *ListScheduleGroupsInput
	nextToken *string
	firstPage bool
}

// NewListScheduleGroupsPaginator returns a new ListScheduleGroupsPaginator
func NewListScheduleGroupsPaginator(client ListScheduleGroupsAPIClient, params *ListScheduleGroupsInput, optFns ...func(*ListScheduleGroupsPaginatorOptions)) *ListScheduleGroupsPaginator {
	if params == nil {
		params = &ListScheduleGroupsInput{}
	}

	options := ListScheduleGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListScheduleGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListScheduleGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListScheduleGroups page.
func (p *ListScheduleGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListScheduleGroupsOutput, error) {
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

	result, err := p.client.ListScheduleGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListScheduleGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListScheduleGroups",
	}
}
