// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/workdocs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List all the comments for the specified document version.
func (c *Client) DescribeComments(ctx context.Context, params *DescribeCommentsInput, optFns ...func(*Options)) (*DescribeCommentsOutput, error) {
	if params == nil {
		params = &DescribeCommentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeComments", params, optFns, c.addOperationDescribeCommentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCommentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCommentsInput struct {

	// The ID of the document.
	//
	// This member is required.
	DocumentId *string

	// The ID of the document version.
	//
	// This member is required.
	VersionId *string

	// Amazon WorkDocs authentication token. Not required when using Amazon Web
	// Services administrator credentials to access the API.
	AuthenticationToken *string

	// The maximum number of items to return.
	Limit *int32

	// The marker for the next set of results. This marker was received from a
	// previous call.
	Marker *string

	noSmithyDocumentSerde
}

type DescribeCommentsOutput struct {

	// The list of comments for the specified document version.
	Comments []types.Comment

	// The marker for the next set of results. This marker was received from a
	// previous call.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCommentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeComments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeComments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeComments"); err != nil {
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
	if err = addOpDescribeCommentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeComments(options.Region), middleware.Before); err != nil {
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

// DescribeCommentsAPIClient is a client that implements the DescribeComments
// operation.
type DescribeCommentsAPIClient interface {
	DescribeComments(context.Context, *DescribeCommentsInput, ...func(*Options)) (*DescribeCommentsOutput, error)
}

var _ DescribeCommentsAPIClient = (*Client)(nil)

// DescribeCommentsPaginatorOptions is the paginator options for DescribeComments
type DescribeCommentsPaginatorOptions struct {
	// The maximum number of items to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeCommentsPaginator is a paginator for DescribeComments
type DescribeCommentsPaginator struct {
	options   DescribeCommentsPaginatorOptions
	client    DescribeCommentsAPIClient
	params    *DescribeCommentsInput
	nextToken *string
	firstPage bool
}

// NewDescribeCommentsPaginator returns a new DescribeCommentsPaginator
func NewDescribeCommentsPaginator(client DescribeCommentsAPIClient, params *DescribeCommentsInput, optFns ...func(*DescribeCommentsPaginatorOptions)) *DescribeCommentsPaginator {
	if params == nil {
		params = &DescribeCommentsInput{}
	}

	options := DescribeCommentsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeCommentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeCommentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeComments page.
func (p *DescribeCommentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeCommentsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeComments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeComments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeComments",
	}
}
