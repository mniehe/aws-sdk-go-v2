// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/workmail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the mobile device access overrides for any given combination of
// WorkMail organization, user, or device.
func (c *Client) ListMobileDeviceAccessOverrides(ctx context.Context, params *ListMobileDeviceAccessOverridesInput, optFns ...func(*Options)) (*ListMobileDeviceAccessOverridesOutput, error) {
	if params == nil {
		params = &ListMobileDeviceAccessOverridesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMobileDeviceAccessOverrides", params, optFns, c.addOperationListMobileDeviceAccessOverridesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMobileDeviceAccessOverridesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMobileDeviceAccessOverridesInput struct {

	// The WorkMail organization under which to list mobile device access overrides.
	//
	// This member is required.
	OrganizationId *string

	// The mobile device to which the access override applies.
	DeviceId *string

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token to use to retrieve the next page of results. The first call does not
	// require a token.
	NextToken *string

	// The WorkMail user under which you list the mobile device access overrides.
	// Accepts the following types of user identities:
	//   - User ID: 12345678-1234-1234-1234-123456789012 or
	//   S-1-1-12-1234567890-123456789-123456789-1234
	//   - Email address: user@domain.tld
	//   - User name: user
	UserId *string

	noSmithyDocumentSerde
}

type ListMobileDeviceAccessOverridesOutput struct {

	// The token to use to retrieve the next page of results. The value is “null” when
	// there are no more results to return.
	NextToken *string

	// The list of mobile device access overrides that exist for the specified
	// WorkMail organization and user.
	Overrides []types.MobileDeviceAccessOverride

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMobileDeviceAccessOverridesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListMobileDeviceAccessOverrides{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMobileDeviceAccessOverrides{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMobileDeviceAccessOverrides"); err != nil {
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
	if err = addOpListMobileDeviceAccessOverridesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMobileDeviceAccessOverrides(options.Region), middleware.Before); err != nil {
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

// ListMobileDeviceAccessOverridesAPIClient is a client that implements the
// ListMobileDeviceAccessOverrides operation.
type ListMobileDeviceAccessOverridesAPIClient interface {
	ListMobileDeviceAccessOverrides(context.Context, *ListMobileDeviceAccessOverridesInput, ...func(*Options)) (*ListMobileDeviceAccessOverridesOutput, error)
}

var _ ListMobileDeviceAccessOverridesAPIClient = (*Client)(nil)

// ListMobileDeviceAccessOverridesPaginatorOptions is the paginator options for
// ListMobileDeviceAccessOverrides
type ListMobileDeviceAccessOverridesPaginatorOptions struct {
	// The maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMobileDeviceAccessOverridesPaginator is a paginator for
// ListMobileDeviceAccessOverrides
type ListMobileDeviceAccessOverridesPaginator struct {
	options   ListMobileDeviceAccessOverridesPaginatorOptions
	client    ListMobileDeviceAccessOverridesAPIClient
	params    *ListMobileDeviceAccessOverridesInput
	nextToken *string
	firstPage bool
}

// NewListMobileDeviceAccessOverridesPaginator returns a new
// ListMobileDeviceAccessOverridesPaginator
func NewListMobileDeviceAccessOverridesPaginator(client ListMobileDeviceAccessOverridesAPIClient, params *ListMobileDeviceAccessOverridesInput, optFns ...func(*ListMobileDeviceAccessOverridesPaginatorOptions)) *ListMobileDeviceAccessOverridesPaginator {
	if params == nil {
		params = &ListMobileDeviceAccessOverridesInput{}
	}

	options := ListMobileDeviceAccessOverridesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMobileDeviceAccessOverridesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMobileDeviceAccessOverridesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMobileDeviceAccessOverrides page.
func (p *ListMobileDeviceAccessOverridesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMobileDeviceAccessOverridesOutput, error) {
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

	result, err := p.client.ListMobileDeviceAccessOverrides(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMobileDeviceAccessOverrides(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMobileDeviceAccessOverrides",
	}
}
