// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the names of the inline policies embedded in the specified IAM user. An
// IAM user can also have managed policies attached to it. To list the managed
// policies that are attached to a user, use ListAttachedUserPolicies. For more
// information about policies, see Managed Policies and Inline Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide. You can paginate the results using the MaxItems and
// Marker parameters. If there are no inline policies embedded with the specified
// user, the operation returns an empty list.
func (c *Client) ListUserPolicies(ctx context.Context, params *ListUserPoliciesInput, optFns ...func(*Options)) (*ListUserPoliciesOutput, error) {
	if params == nil {
		params = &ListUserPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUserPolicies", params, optFns, addOperationListUserPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUserPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUserPoliciesInput struct {

	// The name of the user to list policies for. This parameter allows (through its
	// regex pattern (http://wikipedia.org/wiki/regex)) a string of characters
	// consisting of upper and lowercase alphanumeric characters with no spaces. You
	// can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	UserName *string

	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true. If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true, and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	MaxItems *int32
}

// Contains the response to a successful ListUserPolicies request.
type ListUserPoliciesOutput struct {

	// A list of policy names.
	//
	// This member is required.
	PolicyNames []string

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you receive
	// all your results.
	IsTruncated bool

	// When IsTruncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListUserPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListUserPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListUserPolicies{}, middleware.After)
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
	if err = addOpListUserPoliciesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUserPolicies(options.Region), middleware.Before); err != nil {
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

// ListUserPoliciesAPIClient is a client that implements the ListUserPolicies
// operation.
type ListUserPoliciesAPIClient interface {
	ListUserPolicies(context.Context, *ListUserPoliciesInput, ...func(*Options)) (*ListUserPoliciesOutput, error)
}

var _ ListUserPoliciesAPIClient = (*Client)(nil)

// ListUserPoliciesPaginatorOptions is the paginator options for ListUserPolicies
type ListUserPoliciesPaginatorOptions struct {
	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true. If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true, and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUserPoliciesPaginator is a paginator for ListUserPolicies
type ListUserPoliciesPaginator struct {
	options   ListUserPoliciesPaginatorOptions
	client    ListUserPoliciesAPIClient
	params    *ListUserPoliciesInput
	nextToken *string
	firstPage bool
}

// NewListUserPoliciesPaginator returns a new ListUserPoliciesPaginator
func NewListUserPoliciesPaginator(client ListUserPoliciesAPIClient, params *ListUserPoliciesInput, optFns ...func(*ListUserPoliciesPaginatorOptions)) *ListUserPoliciesPaginator {
	options := ListUserPoliciesPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListUserPoliciesInput{}
	}

	return &ListUserPoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUserPoliciesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListUserPolicies page.
func (p *ListUserPoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUserPoliciesOutput, error) {
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

	result, err := p.client.ListUserPolicies(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListUserPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "ListUserPolicies",
	}
}
