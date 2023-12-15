// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the resource CIDRs that are monitored as part of a resource discovery.
// A discovered resource is a resource CIDR monitored under a resource discovery.
// The following resources can be discovered: VPCs, Public IPv4 pools, VPC subnets,
// and Elastic IP addresses.
func (c *Client) GetIpamDiscoveredResourceCidrs(ctx context.Context, params *GetIpamDiscoveredResourceCidrsInput, optFns ...func(*Options)) (*GetIpamDiscoveredResourceCidrsOutput, error) {
	if params == nil {
		params = &GetIpamDiscoveredResourceCidrsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIpamDiscoveredResourceCidrs", params, optFns, c.addOperationGetIpamDiscoveredResourceCidrsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIpamDiscoveredResourceCidrsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIpamDiscoveredResourceCidrsInput struct {

	// A resource discovery ID.
	//
	// This member is required.
	IpamResourceDiscoveryId *string

	// A resource Region.
	//
	// This member is required.
	ResourceRegion *string

	// A check for whether you have the required permissions for the action without
	// actually making the request and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Filters.
	Filters []types.Filter

	// The maximum number of discovered resource CIDRs to return in one page of
	// results.
	MaxResults *int32

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetIpamDiscoveredResourceCidrsOutput struct {

	// Discovered resource CIDRs.
	IpamDiscoveredResourceCidrs []types.IpamDiscoveredResourceCidr

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetIpamDiscoveredResourceCidrsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetIpamDiscoveredResourceCidrs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetIpamDiscoveredResourceCidrs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetIpamDiscoveredResourceCidrs"); err != nil {
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
	if err = addOpGetIpamDiscoveredResourceCidrsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIpamDiscoveredResourceCidrs(options.Region), middleware.Before); err != nil {
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

// GetIpamDiscoveredResourceCidrsAPIClient is a client that implements the
// GetIpamDiscoveredResourceCidrs operation.
type GetIpamDiscoveredResourceCidrsAPIClient interface {
	GetIpamDiscoveredResourceCidrs(context.Context, *GetIpamDiscoveredResourceCidrsInput, ...func(*Options)) (*GetIpamDiscoveredResourceCidrsOutput, error)
}

var _ GetIpamDiscoveredResourceCidrsAPIClient = (*Client)(nil)

// GetIpamDiscoveredResourceCidrsPaginatorOptions is the paginator options for
// GetIpamDiscoveredResourceCidrs
type GetIpamDiscoveredResourceCidrsPaginatorOptions struct {
	// The maximum number of discovered resource CIDRs to return in one page of
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetIpamDiscoveredResourceCidrsPaginator is a paginator for
// GetIpamDiscoveredResourceCidrs
type GetIpamDiscoveredResourceCidrsPaginator struct {
	options   GetIpamDiscoveredResourceCidrsPaginatorOptions
	client    GetIpamDiscoveredResourceCidrsAPIClient
	params    *GetIpamDiscoveredResourceCidrsInput
	nextToken *string
	firstPage bool
}

// NewGetIpamDiscoveredResourceCidrsPaginator returns a new
// GetIpamDiscoveredResourceCidrsPaginator
func NewGetIpamDiscoveredResourceCidrsPaginator(client GetIpamDiscoveredResourceCidrsAPIClient, params *GetIpamDiscoveredResourceCidrsInput, optFns ...func(*GetIpamDiscoveredResourceCidrsPaginatorOptions)) *GetIpamDiscoveredResourceCidrsPaginator {
	if params == nil {
		params = &GetIpamDiscoveredResourceCidrsInput{}
	}

	options := GetIpamDiscoveredResourceCidrsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetIpamDiscoveredResourceCidrsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetIpamDiscoveredResourceCidrsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetIpamDiscoveredResourceCidrs page.
func (p *GetIpamDiscoveredResourceCidrsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetIpamDiscoveredResourceCidrsOutput, error) {
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

	result, err := p.client.GetIpamDiscoveredResourceCidrs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetIpamDiscoveredResourceCidrs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetIpamDiscoveredResourceCidrs",
	}
}
