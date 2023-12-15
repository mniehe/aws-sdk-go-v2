// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/mturk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The ListAssignmentsForHIT operation retrieves completed assignments for a HIT.
// You can use this operation to retrieve the results for a HIT. You can get
// assignments for a HIT at any time, even if the HIT is not yet Reviewable. If a
// HIT requested multiple assignments, and has received some results but has not
// yet become Reviewable, you can still retrieve the partial results with this
// operation. Use the AssignmentStatus parameter to control which set of
// assignments for a HIT are returned. The ListAssignmentsForHIT operation can
// return submitted assignments awaiting approval, or it can return assignments
// that have already been approved or rejected. You can set
// AssignmentStatus=Approved,Rejected to get assignments that have already been
// approved and rejected together in one result set. Only the Requester who created
// the HIT can retrieve the assignments for that HIT. Results are sorted and
// divided into numbered pages and the operation returns a single page of results.
// You can use the parameters of the operation to control sorting and pagination.
func (c *Client) ListAssignmentsForHIT(ctx context.Context, params *ListAssignmentsForHITInput, optFns ...func(*Options)) (*ListAssignmentsForHITOutput, error) {
	if params == nil {
		params = &ListAssignmentsForHITInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssignmentsForHIT", params, optFns, c.addOperationListAssignmentsForHITMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssignmentsForHITOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssignmentsForHITInput struct {

	// The ID of the HIT.
	//
	// This member is required.
	HITId *string

	// The status of the assignments to return: Submitted | Approved | Rejected
	AssignmentStatuses []types.AssignmentStatus

	MaxResults *int32

	// Pagination token
	NextToken *string

	noSmithyDocumentSerde
}

type ListAssignmentsForHITOutput struct {

	// The collection of Assignment data structures returned by this call.
	Assignments []types.Assignment

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// The number of assignments on the page in the filtered results list, equivalent
	// to the number of assignments returned by this call.
	NumResults *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssignmentsForHITMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAssignmentsForHIT{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAssignmentsForHIT{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAssignmentsForHIT"); err != nil {
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
	if err = addOpListAssignmentsForHITValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssignmentsForHIT(options.Region), middleware.Before); err != nil {
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

// ListAssignmentsForHITAPIClient is a client that implements the
// ListAssignmentsForHIT operation.
type ListAssignmentsForHITAPIClient interface {
	ListAssignmentsForHIT(context.Context, *ListAssignmentsForHITInput, ...func(*Options)) (*ListAssignmentsForHITOutput, error)
}

var _ ListAssignmentsForHITAPIClient = (*Client)(nil)

// ListAssignmentsForHITPaginatorOptions is the paginator options for
// ListAssignmentsForHIT
type ListAssignmentsForHITPaginatorOptions struct {
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssignmentsForHITPaginator is a paginator for ListAssignmentsForHIT
type ListAssignmentsForHITPaginator struct {
	options   ListAssignmentsForHITPaginatorOptions
	client    ListAssignmentsForHITAPIClient
	params    *ListAssignmentsForHITInput
	nextToken *string
	firstPage bool
}

// NewListAssignmentsForHITPaginator returns a new ListAssignmentsForHITPaginator
func NewListAssignmentsForHITPaginator(client ListAssignmentsForHITAPIClient, params *ListAssignmentsForHITInput, optFns ...func(*ListAssignmentsForHITPaginatorOptions)) *ListAssignmentsForHITPaginator {
	if params == nil {
		params = &ListAssignmentsForHITInput{}
	}

	options := ListAssignmentsForHITPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssignmentsForHITPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssignmentsForHITPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssignmentsForHIT page.
func (p *ListAssignmentsForHITPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssignmentsForHITOutput, error) {
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

	result, err := p.client.ListAssignmentsForHIT(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssignmentsForHIT(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAssignmentsForHIT",
	}
}
