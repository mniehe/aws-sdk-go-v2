// Code generated by smithy-go-codegen DO NOT EDIT.

package pricing

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/pricing/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the metadata for one service or a list of the metadata for all
// services. Use this without a service code to get the service codes for all
// services. Use it with a service code, such as AmazonEC2 , to get information
// specific to that service, such as the attribute names available for that
// service. For example, some of the attribute names available for EC2 are
// volumeType , maxIopsVolume , operation , locationType , and
// instanceCapacity10xlarge .
func (c *Client) DescribeServices(ctx context.Context, params *DescribeServicesInput, optFns ...func(*Options)) (*DescribeServicesOutput, error) {
	if params == nil {
		params = &DescribeServicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeServices", params, optFns, c.addOperationDescribeServicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeServicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeServicesInput struct {

	// The format version that you want the response to be in. Valid values are: aws_v1
	FormatVersion *string

	// The maximum number of results that you want returned in the response.
	MaxResults *int32

	// The pagination token that indicates the next set of results that you want to
	// retrieve.
	NextToken *string

	// The code for the service whose information you want to retrieve, such as
	// AmazonEC2 . You can use the ServiceCode to filter the results in a GetProducts
	// call. To retrieve a list of all services, leave this blank.
	ServiceCode *string

	noSmithyDocumentSerde
}

type DescribeServicesOutput struct {

	// The format version of the response. For example, aws_v1 .
	FormatVersion *string

	// The pagination token for the next set of retrievable results.
	NextToken *string

	// The service metadata for the service or services in the response.
	Services []types.Service

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeServicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeServices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeServices{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeServices"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeServices(options.Region), middleware.Before); err != nil {
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

// DescribeServicesAPIClient is a client that implements the DescribeServices
// operation.
type DescribeServicesAPIClient interface {
	DescribeServices(context.Context, *DescribeServicesInput, ...func(*Options)) (*DescribeServicesOutput, error)
}

var _ DescribeServicesAPIClient = (*Client)(nil)

// DescribeServicesPaginatorOptions is the paginator options for DescribeServices
type DescribeServicesPaginatorOptions struct {
	// The maximum number of results that you want returned in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeServicesPaginator is a paginator for DescribeServices
type DescribeServicesPaginator struct {
	options   DescribeServicesPaginatorOptions
	client    DescribeServicesAPIClient
	params    *DescribeServicesInput
	nextToken *string
	firstPage bool
}

// NewDescribeServicesPaginator returns a new DescribeServicesPaginator
func NewDescribeServicesPaginator(client DescribeServicesAPIClient, params *DescribeServicesInput, optFns ...func(*DescribeServicesPaginatorOptions)) *DescribeServicesPaginator {
	if params == nil {
		params = &DescribeServicesInput{}
	}

	options := DescribeServicesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeServicesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeServicesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeServices page.
func (p *DescribeServicesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeServicesOutput, error) {
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

	result, err := p.client.DescribeServices(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeServices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeServices",
	}
}
