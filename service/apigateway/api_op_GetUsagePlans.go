// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets all the usage plans of the caller's account.
func (c *Client) GetUsagePlans(ctx context.Context, params *GetUsagePlansInput, optFns ...func(*Options)) (*GetUsagePlansOutput, error) {
	if params == nil {
		params = &GetUsagePlansInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUsagePlans", params, optFns, c.addOperationGetUsagePlansMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUsagePlansOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The GET request to get all the usage plans of the caller's account.
type GetUsagePlansInput struct {

	// The identifier of the API key associated with the usage plans.
	KeyId *string

	// The maximum number of returned results per page. The default value is 25 and
	// the maximum value is 500.
	Limit *int32

	// The current pagination position in the paged result set.
	Position *string

	noSmithyDocumentSerde
}

// Represents a collection of usage plans for an AWS account.
type GetUsagePlansOutput struct {

	// The current page of elements from this collection.
	Items []types.UsagePlan

	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUsagePlansMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetUsagePlans{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUsagePlans{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetUsagePlans"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUsagePlans(options.Region), middleware.Before); err != nil {
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
	if err = addAcceptHeader(stack); err != nil {
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

// GetUsagePlansAPIClient is a client that implements the GetUsagePlans operation.
type GetUsagePlansAPIClient interface {
	GetUsagePlans(context.Context, *GetUsagePlansInput, ...func(*Options)) (*GetUsagePlansOutput, error)
}

var _ GetUsagePlansAPIClient = (*Client)(nil)

// GetUsagePlansPaginatorOptions is the paginator options for GetUsagePlans
type GetUsagePlansPaginatorOptions struct {
	// The maximum number of returned results per page. The default value is 25 and
	// the maximum value is 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetUsagePlansPaginator is a paginator for GetUsagePlans
type GetUsagePlansPaginator struct {
	options   GetUsagePlansPaginatorOptions
	client    GetUsagePlansAPIClient
	params    *GetUsagePlansInput
	nextToken *string
	firstPage bool
}

// NewGetUsagePlansPaginator returns a new GetUsagePlansPaginator
func NewGetUsagePlansPaginator(client GetUsagePlansAPIClient, params *GetUsagePlansInput, optFns ...func(*GetUsagePlansPaginatorOptions)) *GetUsagePlansPaginator {
	if params == nil {
		params = &GetUsagePlansInput{}
	}

	options := GetUsagePlansPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetUsagePlansPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Position,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetUsagePlansPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetUsagePlans page.
func (p *GetUsagePlansPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetUsagePlansOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Position = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.GetUsagePlans(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Position

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetUsagePlans(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetUsagePlans",
	}
}
