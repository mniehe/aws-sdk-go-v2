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

// Lists all classifier objects in the Data Catalog.
func (c *Client) GetClassifiers(ctx context.Context, params *GetClassifiersInput, optFns ...func(*Options)) (*GetClassifiersOutput, error) {
	if params == nil {
		params = &GetClassifiersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetClassifiers", params, optFns, addOperationGetClassifiersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetClassifiersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetClassifiersInput struct {

	// The size of the list to return (optional).
	MaxResults *int32

	// An optional continuation token.
	NextToken *string
}

type GetClassifiersOutput struct {

	// The requested list of classifier objects.
	Classifiers []types.Classifier

	// A continuation token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetClassifiersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetClassifiers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetClassifiers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetClassifiers(options.Region), middleware.Before); err != nil {
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

// GetClassifiersAPIClient is a client that implements the GetClassifiers
// operation.
type GetClassifiersAPIClient interface {
	GetClassifiers(context.Context, *GetClassifiersInput, ...func(*Options)) (*GetClassifiersOutput, error)
}

var _ GetClassifiersAPIClient = (*Client)(nil)

// GetClassifiersPaginatorOptions is the paginator options for GetClassifiers
type GetClassifiersPaginatorOptions struct {
	// The size of the list to return (optional).
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetClassifiersPaginator is a paginator for GetClassifiers
type GetClassifiersPaginator struct {
	options   GetClassifiersPaginatorOptions
	client    GetClassifiersAPIClient
	params    *GetClassifiersInput
	nextToken *string
	firstPage bool
}

// NewGetClassifiersPaginator returns a new GetClassifiersPaginator
func NewGetClassifiersPaginator(client GetClassifiersAPIClient, params *GetClassifiersInput, optFns ...func(*GetClassifiersPaginatorOptions)) *GetClassifiersPaginator {
	options := GetClassifiersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &GetClassifiersInput{}
	}

	return &GetClassifiersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetClassifiersPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetClassifiers page.
func (p *GetClassifiersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetClassifiersOutput, error) {
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

	result, err := p.client.GetClassifiers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetClassifiers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetClassifiers",
	}
}
