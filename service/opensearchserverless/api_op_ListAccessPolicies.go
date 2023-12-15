// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearchserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/opensearchserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about a list of OpenSearch Serverless access policies.
func (c *Client) ListAccessPolicies(ctx context.Context, params *ListAccessPoliciesInput, optFns ...func(*Options)) (*ListAccessPoliciesOutput, error) {
	if params == nil {
		params = &ListAccessPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccessPolicies", params, optFns, c.addOperationListAccessPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccessPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccessPoliciesInput struct {

	// The type of access policy.
	//
	// This member is required.
	Type types.AccessPolicyType

	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results. The default is 20.
	MaxResults *int32

	// If your initial ListAccessPolicies operation returns a nextToken , you can
	// include the returned nextToken in subsequent ListAccessPolicies operations,
	// which returns results in the next page.
	NextToken *string

	// Resource filters (can be collections or indexes) that policies can apply to.
	Resource []string

	noSmithyDocumentSerde
}

type ListAccessPoliciesOutput struct {

	// Details about the requested access policies.
	AccessPolicySummaries []types.AccessPolicySummary

	// When nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAccessPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListAccessPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListAccessPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAccessPolicies"); err != nil {
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
	if err = addOpListAccessPoliciesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccessPolicies(options.Region), middleware.Before); err != nil {
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

// ListAccessPoliciesAPIClient is a client that implements the ListAccessPolicies
// operation.
type ListAccessPoliciesAPIClient interface {
	ListAccessPolicies(context.Context, *ListAccessPoliciesInput, ...func(*Options)) (*ListAccessPoliciesOutput, error)
}

var _ ListAccessPoliciesAPIClient = (*Client)(nil)

// ListAccessPoliciesPaginatorOptions is the paginator options for
// ListAccessPolicies
type ListAccessPoliciesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccessPoliciesPaginator is a paginator for ListAccessPolicies
type ListAccessPoliciesPaginator struct {
	options   ListAccessPoliciesPaginatorOptions
	client    ListAccessPoliciesAPIClient
	params    *ListAccessPoliciesInput
	nextToken *string
	firstPage bool
}

// NewListAccessPoliciesPaginator returns a new ListAccessPoliciesPaginator
func NewListAccessPoliciesPaginator(client ListAccessPoliciesAPIClient, params *ListAccessPoliciesInput, optFns ...func(*ListAccessPoliciesPaginatorOptions)) *ListAccessPoliciesPaginator {
	if params == nil {
		params = &ListAccessPoliciesInput{}
	}

	options := ListAccessPoliciesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccessPoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccessPoliciesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAccessPolicies page.
func (p *ListAccessPoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccessPoliciesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListAccessPolicies(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAccessPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAccessPolicies",
	}
}
