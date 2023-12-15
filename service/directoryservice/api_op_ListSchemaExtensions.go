// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all schema extensions applied to a Microsoft AD Directory.
func (c *Client) ListSchemaExtensions(ctx context.Context, params *ListSchemaExtensionsInput, optFns ...func(*Options)) (*ListSchemaExtensionsOutput, error) {
	if params == nil {
		params = &ListSchemaExtensionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSchemaExtensions", params, optFns, c.addOperationListSchemaExtensionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSchemaExtensionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSchemaExtensionsInput struct {

	// The identifier of the directory from which to retrieve the schema extension
	// information.
	//
	// This member is required.
	DirectoryId *string

	// The maximum number of items to return.
	Limit *int32

	// The ListSchemaExtensions.NextToken value from a previous call to
	// ListSchemaExtensions . Pass null if this is the first call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSchemaExtensionsOutput struct {

	// If not null, more results are available. Pass this value for the NextToken
	// parameter in a subsequent call to ListSchemaExtensions to retrieve the next set
	// of items.
	NextToken *string

	// Information about the schema extensions applied to the directory.
	SchemaExtensionsInfo []types.SchemaExtensionInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSchemaExtensionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSchemaExtensions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSchemaExtensions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSchemaExtensions"); err != nil {
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
	if err = addOpListSchemaExtensionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSchemaExtensions(options.Region), middleware.Before); err != nil {
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

// ListSchemaExtensionsAPIClient is a client that implements the
// ListSchemaExtensions operation.
type ListSchemaExtensionsAPIClient interface {
	ListSchemaExtensions(context.Context, *ListSchemaExtensionsInput, ...func(*Options)) (*ListSchemaExtensionsOutput, error)
}

var _ ListSchemaExtensionsAPIClient = (*Client)(nil)

// ListSchemaExtensionsPaginatorOptions is the paginator options for
// ListSchemaExtensions
type ListSchemaExtensionsPaginatorOptions struct {
	// The maximum number of items to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSchemaExtensionsPaginator is a paginator for ListSchemaExtensions
type ListSchemaExtensionsPaginator struct {
	options   ListSchemaExtensionsPaginatorOptions
	client    ListSchemaExtensionsAPIClient
	params    *ListSchemaExtensionsInput
	nextToken *string
	firstPage bool
}

// NewListSchemaExtensionsPaginator returns a new ListSchemaExtensionsPaginator
func NewListSchemaExtensionsPaginator(client ListSchemaExtensionsAPIClient, params *ListSchemaExtensionsInput, optFns ...func(*ListSchemaExtensionsPaginatorOptions)) *ListSchemaExtensionsPaginator {
	if params == nil {
		params = &ListSchemaExtensionsInput{}
	}

	options := ListSchemaExtensionsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSchemaExtensionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSchemaExtensionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSchemaExtensions page.
func (p *ListSchemaExtensionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSchemaExtensionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListSchemaExtensions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSchemaExtensions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSchemaExtensions",
	}
}
