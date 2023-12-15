// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/mniehe/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/mniehe/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// Lists all the Storage Lens groups in the specified home Region. To use this
// operation, you must have the permission to perform the s3:ListStorageLensGroups
// action. For more information about the required Storage Lens Groups permissions,
// see Setting account permissions to use S3 Storage Lens groups (https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage_lens_iam_permissions.html#storage_lens_groups_permissions)
// . For information about Storage Lens groups errors, see List of Amazon S3
// Storage Lens error codes (https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#S3LensErrorCodeList)
// .
func (c *Client) ListStorageLensGroups(ctx context.Context, params *ListStorageLensGroupsInput, optFns ...func(*Options)) (*ListStorageLensGroupsOutput, error) {
	if params == nil {
		params = &ListStorageLensGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStorageLensGroups", params, optFns, c.addOperationListStorageLensGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStorageLensGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStorageLensGroupsInput struct {

	// The Amazon Web Services account ID that owns the Storage Lens groups.
	//
	// This member is required.
	AccountId *string

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	noSmithyDocumentSerde
}

func (in *ListStorageLensGroupsInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.RequiresAccountId = ptr.Bool(true)
}

type ListStorageLensGroupsOutput struct {

	// If NextToken is returned, there are more Storage Lens groups results available.
	// The value of NextToken is a unique pagination token for each page. Make the
	// call again using the returned token to retrieve the next page. Keep all other
	// arguments unchanged. Each pagination token expires after 24 hours.
	NextToken *string

	// The list of Storage Lens groups that exist in the specified home Region.
	StorageLensGroupList []types.ListStorageLensGroupEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStorageLensGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListStorageLensGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListStorageLensGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListStorageLensGroups"); err != nil {
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opListStorageLensGroupsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListStorageLensGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStorageLensGroups(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addListStorageLensGroupsUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addStashOperationInput(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = s3controlcust.AddDisableHostPrefixMiddleware(stack); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opListStorageLensGroupsMiddleware struct {
}

func (*endpointPrefix_opListStorageLensGroupsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListStorageLensGroupsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	opaqueInput := getOperationInput(ctx)
	input, ok := opaqueInput.(*ListStorageLensGroupsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", opaqueInput)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListStorageLensGroupsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListStorageLensGroupsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListStorageLensGroupsAPIClient is a client that implements the
// ListStorageLensGroups operation.
type ListStorageLensGroupsAPIClient interface {
	ListStorageLensGroups(context.Context, *ListStorageLensGroupsInput, ...func(*Options)) (*ListStorageLensGroupsOutput, error)
}

var _ ListStorageLensGroupsAPIClient = (*Client)(nil)

// ListStorageLensGroupsPaginatorOptions is the paginator options for
// ListStorageLensGroups
type ListStorageLensGroupsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStorageLensGroupsPaginator is a paginator for ListStorageLensGroups
type ListStorageLensGroupsPaginator struct {
	options   ListStorageLensGroupsPaginatorOptions
	client    ListStorageLensGroupsAPIClient
	params    *ListStorageLensGroupsInput
	nextToken *string
	firstPage bool
}

// NewListStorageLensGroupsPaginator returns a new ListStorageLensGroupsPaginator
func NewListStorageLensGroupsPaginator(client ListStorageLensGroupsAPIClient, params *ListStorageLensGroupsInput, optFns ...func(*ListStorageLensGroupsPaginatorOptions)) *ListStorageLensGroupsPaginator {
	if params == nil {
		params = &ListStorageLensGroupsInput{}
	}

	options := ListStorageLensGroupsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStorageLensGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStorageLensGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStorageLensGroups page.
func (p *ListStorageLensGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStorageLensGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListStorageLensGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStorageLensGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListStorageLensGroups",
	}
}

func copyListStorageLensGroupsInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*ListStorageLensGroupsInput)
	if !ok {
		return nil, fmt.Errorf("expect *ListStorageLensGroupsInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *ListStorageLensGroupsInput) copy() interface{} {
	v := *in
	return &v
}
func backFillListStorageLensGroupsAccountID(input interface{}, v string) error {
	in := input.(*ListStorageLensGroupsInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addListStorageLensGroupsUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyListStorageLensGroupsInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
