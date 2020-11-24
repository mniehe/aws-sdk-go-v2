// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists logging levels.
func (c *Client) ListV2LoggingLevels(ctx context.Context, params *ListV2LoggingLevelsInput, optFns ...func(*Options)) (*ListV2LoggingLevelsOutput, error) {
	if params == nil {
		params = &ListV2LoggingLevelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListV2LoggingLevels", params, optFns, addOperationListV2LoggingLevelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListV2LoggingLevelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListV2LoggingLevelsInput struct {

	// The maximum number of results to return at one time.
	MaxResults *int32

	// The token used to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// The type of resource for which you are configuring logging. Must be THING_Group.
	TargetType types.LogTargetType
}

type ListV2LoggingLevelsOutput struct {

	// The logging configuration for a target.
	LogTargetConfigurations []types.LogTargetConfiguration

	// The token used to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListV2LoggingLevelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListV2LoggingLevels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListV2LoggingLevels{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListV2LoggingLevels(options.Region), middleware.Before); err != nil {
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

// ListV2LoggingLevelsAPIClient is a client that implements the ListV2LoggingLevels
// operation.
type ListV2LoggingLevelsAPIClient interface {
	ListV2LoggingLevels(context.Context, *ListV2LoggingLevelsInput, ...func(*Options)) (*ListV2LoggingLevelsOutput, error)
}

var _ ListV2LoggingLevelsAPIClient = (*Client)(nil)

// ListV2LoggingLevelsPaginatorOptions is the paginator options for
// ListV2LoggingLevels
type ListV2LoggingLevelsPaginatorOptions struct {
	// The maximum number of results to return at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListV2LoggingLevelsPaginator is a paginator for ListV2LoggingLevels
type ListV2LoggingLevelsPaginator struct {
	options   ListV2LoggingLevelsPaginatorOptions
	client    ListV2LoggingLevelsAPIClient
	params    *ListV2LoggingLevelsInput
	nextToken *string
	firstPage bool
}

// NewListV2LoggingLevelsPaginator returns a new ListV2LoggingLevelsPaginator
func NewListV2LoggingLevelsPaginator(client ListV2LoggingLevelsAPIClient, params *ListV2LoggingLevelsInput, optFns ...func(*ListV2LoggingLevelsPaginatorOptions)) *ListV2LoggingLevelsPaginator {
	options := ListV2LoggingLevelsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListV2LoggingLevelsInput{}
	}

	return &ListV2LoggingLevelsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListV2LoggingLevelsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListV2LoggingLevels page.
func (p *ListV2LoggingLevelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListV2LoggingLevelsOutput, error) {
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

	result, err := p.client.ListV2LoggingLevels(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListV2LoggingLevels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListV2LoggingLevels",
	}
}
