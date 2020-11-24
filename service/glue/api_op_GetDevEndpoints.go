// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves all the development endpoints in this AWS account. When you create a
// development endpoint in a virtual private cloud (VPC), AWS Glue returns only a
// private IP address and the public IP address field is not populated. When you
// create a non-VPC development endpoint, AWS Glue returns only a public IP
// address.
func (c *Client) GetDevEndpoints(ctx context.Context, params *GetDevEndpointsInput, optFns ...func(*Options)) (*GetDevEndpointsOutput, error) {
	if params == nil {
		params = &GetDevEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDevEndpoints", params, optFns, addOperationGetDevEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDevEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDevEndpointsInput struct {

	// The maximum size of information to return.
	MaxResults *int32

	// A continuation token, if this is a continuation call.
	NextToken *string
}

type GetDevEndpointsOutput struct {

	// A list of DevEndpoint definitions.
	DevEndpoints []types.DevEndpoint

	// A continuation token, if not all DevEndpoint definitions have yet been returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDevEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDevEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDevEndpoints{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDevEndpoints(options.Region), middleware.Before); err != nil {
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

// GetDevEndpointsAPIClient is a client that implements the GetDevEndpoints
// operation.
type GetDevEndpointsAPIClient interface {
	GetDevEndpoints(context.Context, *GetDevEndpointsInput, ...func(*Options)) (*GetDevEndpointsOutput, error)
}

var _ GetDevEndpointsAPIClient = (*Client)(nil)

// GetDevEndpointsPaginatorOptions is the paginator options for GetDevEndpoints
type GetDevEndpointsPaginatorOptions struct {
	// The maximum size of information to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetDevEndpointsPaginator is a paginator for GetDevEndpoints
type GetDevEndpointsPaginator struct {
	options   GetDevEndpointsPaginatorOptions
	client    GetDevEndpointsAPIClient
	params    *GetDevEndpointsInput
	nextToken *string
	firstPage bool
}

// NewGetDevEndpointsPaginator returns a new GetDevEndpointsPaginator
func NewGetDevEndpointsPaginator(client GetDevEndpointsAPIClient, params *GetDevEndpointsInput, optFns ...func(*GetDevEndpointsPaginatorOptions)) *GetDevEndpointsPaginator {
	options := GetDevEndpointsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &GetDevEndpointsInput{}
	}

	return &GetDevEndpointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetDevEndpointsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetDevEndpoints page.
func (p *GetDevEndpointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetDevEndpointsOutput, error) {
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

	result, err := p.client.GetDevEndpoints(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetDevEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetDevEndpoints",
	}
}
