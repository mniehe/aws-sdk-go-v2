// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list that describes the available WorkSpace bundles. You can filter
// the results using either bundle ID or owner, but not both.
func (c *Client) DescribeWorkspaceBundles(ctx context.Context, params *DescribeWorkspaceBundlesInput, optFns ...func(*Options)) (*DescribeWorkspaceBundlesOutput, error) {
	if params == nil {
		params = &DescribeWorkspaceBundlesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWorkspaceBundles", params, optFns, addOperationDescribeWorkspaceBundlesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWorkspaceBundlesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorkspaceBundlesInput struct {

	// The identifiers of the bundles. You cannot combine this parameter with any other
	// filter.
	BundleIds []string

	// The token for the next set of results. (You received this token from a previous
	// call.)
	NextToken *string

	// The owner of the bundles. You cannot combine this parameter with any other
	// filter. Specify AMAZON to describe the bundles provided by AWS or null to
	// describe the bundles that belong to your account.
	Owner *string
}

type DescribeWorkspaceBundlesOutput struct {

	// Information about the bundles.
	Bundles []types.WorkspaceBundle

	// The token to use to retrieve the next set of results, or null if there are no
	// more results available. This token is valid for one day and must be used within
	// that time frame.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeWorkspaceBundlesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeWorkspaceBundles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeWorkspaceBundles{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorkspaceBundles(options.Region), middleware.Before); err != nil {
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

// DescribeWorkspaceBundlesAPIClient is a client that implements the
// DescribeWorkspaceBundles operation.
type DescribeWorkspaceBundlesAPIClient interface {
	DescribeWorkspaceBundles(context.Context, *DescribeWorkspaceBundlesInput, ...func(*Options)) (*DescribeWorkspaceBundlesOutput, error)
}

var _ DescribeWorkspaceBundlesAPIClient = (*Client)(nil)

// DescribeWorkspaceBundlesPaginatorOptions is the paginator options for
// DescribeWorkspaceBundles
type DescribeWorkspaceBundlesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeWorkspaceBundlesPaginator is a paginator for DescribeWorkspaceBundles
type DescribeWorkspaceBundlesPaginator struct {
	options   DescribeWorkspaceBundlesPaginatorOptions
	client    DescribeWorkspaceBundlesAPIClient
	params    *DescribeWorkspaceBundlesInput
	nextToken *string
	firstPage bool
}

// NewDescribeWorkspaceBundlesPaginator returns a new
// DescribeWorkspaceBundlesPaginator
func NewDescribeWorkspaceBundlesPaginator(client DescribeWorkspaceBundlesAPIClient, params *DescribeWorkspaceBundlesInput, optFns ...func(*DescribeWorkspaceBundlesPaginatorOptions)) *DescribeWorkspaceBundlesPaginator {
	options := DescribeWorkspaceBundlesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeWorkspaceBundlesInput{}
	}

	return &DescribeWorkspaceBundlesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeWorkspaceBundlesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeWorkspaceBundles page.
func (p *DescribeWorkspaceBundlesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeWorkspaceBundlesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.DescribeWorkspaceBundles(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeWorkspaceBundles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DescribeWorkspaceBundles",
	}
}
