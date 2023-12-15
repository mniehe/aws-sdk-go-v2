// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of inbound integrations.
func (c *Client) DescribeInboundIntegrations(ctx context.Context, params *DescribeInboundIntegrationsInput, optFns ...func(*Options)) (*DescribeInboundIntegrationsOutput, error) {
	if params == nil {
		params = &DescribeInboundIntegrationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInboundIntegrations", params, optFns, c.addOperationDescribeInboundIntegrationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInboundIntegrationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInboundIntegrationsInput struct {

	// The Amazon Resource Name (ARN) of the inbound integration.
	IntegrationArn *string

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeInboundIntegrations request
	// exceed the value specified in MaxRecords , Amazon Web Services returns a value
	// in the Marker field of the response. You can retrieve the next set of response
	// records by providing the returned marker value in the Marker parameter and
	// retrying the request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32

	// The Amazon Resource Name (ARN) of the target of an inbound integration.
	TargetArn *string

	noSmithyDocumentSerde
}

type DescribeInboundIntegrationsOutput struct {

	// A list of InboundIntegration instances.
	InboundIntegrations []types.InboundIntegration

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInboundIntegrationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeInboundIntegrations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeInboundIntegrations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeInboundIntegrations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInboundIntegrations(options.Region), middleware.Before); err != nil {
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

// DescribeInboundIntegrationsAPIClient is a client that implements the
// DescribeInboundIntegrations operation.
type DescribeInboundIntegrationsAPIClient interface {
	DescribeInboundIntegrations(context.Context, *DescribeInboundIntegrationsInput, ...func(*Options)) (*DescribeInboundIntegrationsOutput, error)
}

var _ DescribeInboundIntegrationsAPIClient = (*Client)(nil)

// DescribeInboundIntegrationsPaginatorOptions is the paginator options for
// DescribeInboundIntegrations
type DescribeInboundIntegrationsPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeInboundIntegrationsPaginator is a paginator for
// DescribeInboundIntegrations
type DescribeInboundIntegrationsPaginator struct {
	options   DescribeInboundIntegrationsPaginatorOptions
	client    DescribeInboundIntegrationsAPIClient
	params    *DescribeInboundIntegrationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeInboundIntegrationsPaginator returns a new
// DescribeInboundIntegrationsPaginator
func NewDescribeInboundIntegrationsPaginator(client DescribeInboundIntegrationsAPIClient, params *DescribeInboundIntegrationsInput, optFns ...func(*DescribeInboundIntegrationsPaginatorOptions)) *DescribeInboundIntegrationsPaginator {
	if params == nil {
		params = &DescribeInboundIntegrationsInput{}
	}

	options := DescribeInboundIntegrationsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeInboundIntegrationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeInboundIntegrationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeInboundIntegrations page.
func (p *DescribeInboundIntegrationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeInboundIntegrationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeInboundIntegrations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeInboundIntegrations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeInboundIntegrations",
	}
}
