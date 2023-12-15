// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides organization conformance pack deployment status for an organization.
// The status is not considered successful until organization conformance pack is
// successfully deployed in all the member accounts with an exception of excluded
// accounts. When you specify the limit and the next token, you receive a paginated
// response. Limit and next token are not applicable if you specify organization
// conformance pack names. They are only applicable, when you request all the
// organization conformance packs.
func (c *Client) DescribeOrganizationConformancePackStatuses(ctx context.Context, params *DescribeOrganizationConformancePackStatusesInput, optFns ...func(*Options)) (*DescribeOrganizationConformancePackStatusesOutput, error) {
	if params == nil {
		params = &DescribeOrganizationConformancePackStatusesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOrganizationConformancePackStatuses", params, optFns, c.addOperationDescribeOrganizationConformancePackStatusesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOrganizationConformancePackStatusesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrganizationConformancePackStatusesInput struct {

	// The maximum number of OrganizationConformancePackStatuses returned on each
	// page. If you do no specify a number, Config uses the default. The default is
	// 100.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// The names of organization conformance packs for which you want status details.
	// If you do not specify any names, Config returns details for all your
	// organization conformance packs.
	OrganizationConformancePackNames []string

	noSmithyDocumentSerde
}

type DescribeOrganizationConformancePackStatusesOutput struct {

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// A list of OrganizationConformancePackStatus objects.
	OrganizationConformancePackStatuses []types.OrganizationConformancePackStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOrganizationConformancePackStatusesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeOrganizationConformancePackStatuses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeOrganizationConformancePackStatuses{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeOrganizationConformancePackStatuses"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrganizationConformancePackStatuses(options.Region), middleware.Before); err != nil {
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

// DescribeOrganizationConformancePackStatusesAPIClient is a client that
// implements the DescribeOrganizationConformancePackStatuses operation.
type DescribeOrganizationConformancePackStatusesAPIClient interface {
	DescribeOrganizationConformancePackStatuses(context.Context, *DescribeOrganizationConformancePackStatusesInput, ...func(*Options)) (*DescribeOrganizationConformancePackStatusesOutput, error)
}

var _ DescribeOrganizationConformancePackStatusesAPIClient = (*Client)(nil)

// DescribeOrganizationConformancePackStatusesPaginatorOptions is the paginator
// options for DescribeOrganizationConformancePackStatuses
type DescribeOrganizationConformancePackStatusesPaginatorOptions struct {
	// The maximum number of OrganizationConformancePackStatuses returned on each
	// page. If you do no specify a number, Config uses the default. The default is
	// 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeOrganizationConformancePackStatusesPaginator is a paginator for
// DescribeOrganizationConformancePackStatuses
type DescribeOrganizationConformancePackStatusesPaginator struct {
	options   DescribeOrganizationConformancePackStatusesPaginatorOptions
	client    DescribeOrganizationConformancePackStatusesAPIClient
	params    *DescribeOrganizationConformancePackStatusesInput
	nextToken *string
	firstPage bool
}

// NewDescribeOrganizationConformancePackStatusesPaginator returns a new
// DescribeOrganizationConformancePackStatusesPaginator
func NewDescribeOrganizationConformancePackStatusesPaginator(client DescribeOrganizationConformancePackStatusesAPIClient, params *DescribeOrganizationConformancePackStatusesInput, optFns ...func(*DescribeOrganizationConformancePackStatusesPaginatorOptions)) *DescribeOrganizationConformancePackStatusesPaginator {
	if params == nil {
		params = &DescribeOrganizationConformancePackStatusesInput{}
	}

	options := DescribeOrganizationConformancePackStatusesPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeOrganizationConformancePackStatusesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeOrganizationConformancePackStatusesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeOrganizationConformancePackStatuses page.
func (p *DescribeOrganizationConformancePackStatusesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeOrganizationConformancePackStatusesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.DescribeOrganizationConformancePackStatuses(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeOrganizationConformancePackStatuses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeOrganizationConformancePackStatuses",
	}
}
