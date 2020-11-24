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

// Retrieves a list of all security configurations.
func (c *Client) GetSecurityConfigurations(ctx context.Context, params *GetSecurityConfigurationsInput, optFns ...func(*Options)) (*GetSecurityConfigurationsOutput, error) {
	if params == nil {
		params = &GetSecurityConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSecurityConfigurations", params, optFns, addOperationGetSecurityConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSecurityConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSecurityConfigurationsInput struct {

	// The maximum number of results to return.
	MaxResults *int32

	// A continuation token, if this is a continuation call.
	NextToken *string
}

type GetSecurityConfigurationsOutput struct {

	// A continuation token, if there are more security configurations to return.
	NextToken *string

	// A list of security configurations.
	SecurityConfigurations []types.SecurityConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSecurityConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSecurityConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSecurityConfigurations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSecurityConfigurations(options.Region), middleware.Before); err != nil {
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

// GetSecurityConfigurationsAPIClient is a client that implements the
// GetSecurityConfigurations operation.
type GetSecurityConfigurationsAPIClient interface {
	GetSecurityConfigurations(context.Context, *GetSecurityConfigurationsInput, ...func(*Options)) (*GetSecurityConfigurationsOutput, error)
}

var _ GetSecurityConfigurationsAPIClient = (*Client)(nil)

// GetSecurityConfigurationsPaginatorOptions is the paginator options for
// GetSecurityConfigurations
type GetSecurityConfigurationsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetSecurityConfigurationsPaginator is a paginator for GetSecurityConfigurations
type GetSecurityConfigurationsPaginator struct {
	options   GetSecurityConfigurationsPaginatorOptions
	client    GetSecurityConfigurationsAPIClient
	params    *GetSecurityConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewGetSecurityConfigurationsPaginator returns a new
// GetSecurityConfigurationsPaginator
func NewGetSecurityConfigurationsPaginator(client GetSecurityConfigurationsAPIClient, params *GetSecurityConfigurationsInput, optFns ...func(*GetSecurityConfigurationsPaginatorOptions)) *GetSecurityConfigurationsPaginator {
	options := GetSecurityConfigurationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &GetSecurityConfigurationsInput{}
	}

	return &GetSecurityConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetSecurityConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetSecurityConfigurations page.
func (p *GetSecurityConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetSecurityConfigurationsOutput, error) {
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

	result, err := p.client.GetSecurityConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetSecurityConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetSecurityConfigurations",
	}
}
