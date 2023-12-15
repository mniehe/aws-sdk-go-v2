// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the firewall rule group associations that you have defined. Each
// association enables DNS filtering for a VPC with one rule group. A single call
// might return only a partial list of the associations. For information, see
// MaxResults .
func (c *Client) ListFirewallRuleGroupAssociations(ctx context.Context, params *ListFirewallRuleGroupAssociationsInput, optFns ...func(*Options)) (*ListFirewallRuleGroupAssociationsOutput, error) {
	if params == nil {
		params = &ListFirewallRuleGroupAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFirewallRuleGroupAssociations", params, optFns, c.addOperationListFirewallRuleGroupAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFirewallRuleGroupAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFirewallRuleGroupAssociationsInput struct {

	// The unique identifier of the firewall rule group that you want to retrieve the
	// associations for. Leave this blank to retrieve associations for any rule group.
	FirewallRuleGroupId *string

	// The maximum number of objects that you want Resolver to return for this
	// request. If more objects are available, in the response, Resolver provides a
	// NextToken value that you can use in a subsequent call to get the next batch of
	// objects. If you don't specify a value for MaxResults , Resolver returns up to
	// 100 objects.
	MaxResults *int32

	// For the first call to this list request, omit this value. When you request a
	// list of objects, Resolver returns at most the number of objects specified in
	// MaxResults . If more objects are available for retrieval, Resolver returns a
	// NextToken value in the response. To retrieve the next batch of objects, use the
	// token that was returned for the prior request in your next request.
	NextToken *string

	// The setting that determines the processing order of the rule group among the
	// rule groups that are associated with a single VPC. DNS Firewall filters VPC
	// traffic starting from the rule group with the lowest numeric priority setting.
	Priority *int32

	// The association Status setting that you want DNS Firewall to filter on for the
	// list. If you don't specify this, then DNS Firewall returns all associations,
	// regardless of status.
	Status types.FirewallRuleGroupAssociationStatus

	// The unique identifier of the VPC that you want to retrieve the associations
	// for. Leave this blank to retrieve associations for any VPC.
	VpcId *string

	noSmithyDocumentSerde
}

type ListFirewallRuleGroupAssociationsOutput struct {

	// A list of your firewall rule group associations. This might be a partial list
	// of the associations that you have defined. For information, see MaxResults .
	FirewallRuleGroupAssociations []types.FirewallRuleGroupAssociation

	// If objects are still available for retrieval, Resolver returns this token in
	// the response. To retrieve the next batch of objects, provide this token in your
	// next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFirewallRuleGroupAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFirewallRuleGroupAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFirewallRuleGroupAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFirewallRuleGroupAssociations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFirewallRuleGroupAssociations(options.Region), middleware.Before); err != nil {
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

// ListFirewallRuleGroupAssociationsAPIClient is a client that implements the
// ListFirewallRuleGroupAssociations operation.
type ListFirewallRuleGroupAssociationsAPIClient interface {
	ListFirewallRuleGroupAssociations(context.Context, *ListFirewallRuleGroupAssociationsInput, ...func(*Options)) (*ListFirewallRuleGroupAssociationsOutput, error)
}

var _ ListFirewallRuleGroupAssociationsAPIClient = (*Client)(nil)

// ListFirewallRuleGroupAssociationsPaginatorOptions is the paginator options for
// ListFirewallRuleGroupAssociations
type ListFirewallRuleGroupAssociationsPaginatorOptions struct {
	// The maximum number of objects that you want Resolver to return for this
	// request. If more objects are available, in the response, Resolver provides a
	// NextToken value that you can use in a subsequent call to get the next batch of
	// objects. If you don't specify a value for MaxResults , Resolver returns up to
	// 100 objects.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFirewallRuleGroupAssociationsPaginator is a paginator for
// ListFirewallRuleGroupAssociations
type ListFirewallRuleGroupAssociationsPaginator struct {
	options   ListFirewallRuleGroupAssociationsPaginatorOptions
	client    ListFirewallRuleGroupAssociationsAPIClient
	params    *ListFirewallRuleGroupAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListFirewallRuleGroupAssociationsPaginator returns a new
// ListFirewallRuleGroupAssociationsPaginator
func NewListFirewallRuleGroupAssociationsPaginator(client ListFirewallRuleGroupAssociationsAPIClient, params *ListFirewallRuleGroupAssociationsInput, optFns ...func(*ListFirewallRuleGroupAssociationsPaginatorOptions)) *ListFirewallRuleGroupAssociationsPaginator {
	if params == nil {
		params = &ListFirewallRuleGroupAssociationsInput{}
	}

	options := ListFirewallRuleGroupAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFirewallRuleGroupAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFirewallRuleGroupAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFirewallRuleGroupAssociations page.
func (p *ListFirewallRuleGroupAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFirewallRuleGroupAssociationsOutput, error) {
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

	result, err := p.client.ListFirewallRuleGroupAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFirewallRuleGroupAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFirewallRuleGroupAssociations",
	}
}
