// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists certificates that are being transferred but not yet accepted. Requires
// permission to access the ListOutgoingCertificates (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) ListOutgoingCertificates(ctx context.Context, params *ListOutgoingCertificatesInput, optFns ...func(*Options)) (*ListOutgoingCertificatesOutput, error) {
	if params == nil {
		params = &ListOutgoingCertificatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOutgoingCertificates", params, optFns, c.addOperationListOutgoingCertificatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOutgoingCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input to the ListOutgoingCertificates operation.
type ListOutgoingCertificatesInput struct {

	// Specifies the order for results. If True, the results are returned in ascending
	// order, based on the creation date.
	AscendingOrder bool

	// The marker for the next set of results.
	Marker *string

	// The result page size.
	PageSize *int32

	noSmithyDocumentSerde
}

// The output from the ListOutgoingCertificates operation.
type ListOutgoingCertificatesOutput struct {

	// The marker for the next set of results.
	NextMarker *string

	// The certificates that are being transferred but not yet accepted.
	OutgoingCertificates []types.OutgoingCertificate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOutgoingCertificatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListOutgoingCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListOutgoingCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListOutgoingCertificates"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOutgoingCertificates(options.Region), middleware.Before); err != nil {
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

// ListOutgoingCertificatesAPIClient is a client that implements the
// ListOutgoingCertificates operation.
type ListOutgoingCertificatesAPIClient interface {
	ListOutgoingCertificates(context.Context, *ListOutgoingCertificatesInput, ...func(*Options)) (*ListOutgoingCertificatesOutput, error)
}

var _ ListOutgoingCertificatesAPIClient = (*Client)(nil)

// ListOutgoingCertificatesPaginatorOptions is the paginator options for
// ListOutgoingCertificates
type ListOutgoingCertificatesPaginatorOptions struct {
	// The result page size.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOutgoingCertificatesPaginator is a paginator for ListOutgoingCertificates
type ListOutgoingCertificatesPaginator struct {
	options   ListOutgoingCertificatesPaginatorOptions
	client    ListOutgoingCertificatesAPIClient
	params    *ListOutgoingCertificatesInput
	nextToken *string
	firstPage bool
}

// NewListOutgoingCertificatesPaginator returns a new
// ListOutgoingCertificatesPaginator
func NewListOutgoingCertificatesPaginator(client ListOutgoingCertificatesAPIClient, params *ListOutgoingCertificatesInput, optFns ...func(*ListOutgoingCertificatesPaginatorOptions)) *ListOutgoingCertificatesPaginator {
	if params == nil {
		params = &ListOutgoingCertificatesInput{}
	}

	options := ListOutgoingCertificatesPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOutgoingCertificatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOutgoingCertificatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOutgoingCertificates page.
func (p *ListOutgoingCertificatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOutgoingCertificatesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.PageSize = limit

	result, err := p.client.ListOutgoingCertificates(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListOutgoingCertificates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListOutgoingCertificates",
	}
}
