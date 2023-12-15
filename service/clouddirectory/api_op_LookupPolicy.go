// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all policies from the root of the Directory to the object specified. If
// there are no policies present, an empty list is returned. If policies are
// present, and if some objects don't have the policies attached, it returns the
// ObjectIdentifier for such objects. If policies are present, it returns
// ObjectIdentifier , policyId , and policyType . Paths that don't lead to the root
// from the target object are ignored. For more information, see Policies (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/key_concepts_directory.html#key_concepts_policies)
// .
func (c *Client) LookupPolicy(ctx context.Context, params *LookupPolicyInput, optFns ...func(*Options)) (*LookupPolicyOutput, error) {
	if params == nil {
		params = &LookupPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "LookupPolicy", params, optFns, c.addOperationLookupPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*LookupPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type LookupPolicyInput struct {

	// The Amazon Resource Name (ARN) that is associated with the Directory . For more
	// information, see arns .
	//
	// This member is required.
	DirectoryArn *string

	// Reference that identifies the object whose policies will be looked up.
	//
	// This member is required.
	ObjectReference *types.ObjectReference

	// The maximum number of items to be retrieved in a single call. This is an
	// approximate number.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type LookupPolicyOutput struct {

	// The pagination token.
	NextToken *string

	// Provides list of path to policies. Policies contain PolicyId , ObjectIdentifier
	// , and PolicyType . For more information, see Policies (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/key_concepts_directory.html#key_concepts_policies)
	// .
	PolicyToPathList []types.PolicyToPath

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationLookupPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpLookupPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpLookupPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "LookupPolicy"); err != nil {
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
	if err = addOpLookupPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opLookupPolicy(options.Region), middleware.Before); err != nil {
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

// LookupPolicyAPIClient is a client that implements the LookupPolicy operation.
type LookupPolicyAPIClient interface {
	LookupPolicy(context.Context, *LookupPolicyInput, ...func(*Options)) (*LookupPolicyOutput, error)
}

var _ LookupPolicyAPIClient = (*Client)(nil)

// LookupPolicyPaginatorOptions is the paginator options for LookupPolicy
type LookupPolicyPaginatorOptions struct {
	// The maximum number of items to be retrieved in a single call. This is an
	// approximate number.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// LookupPolicyPaginator is a paginator for LookupPolicy
type LookupPolicyPaginator struct {
	options   LookupPolicyPaginatorOptions
	client    LookupPolicyAPIClient
	params    *LookupPolicyInput
	nextToken *string
	firstPage bool
}

// NewLookupPolicyPaginator returns a new LookupPolicyPaginator
func NewLookupPolicyPaginator(client LookupPolicyAPIClient, params *LookupPolicyInput, optFns ...func(*LookupPolicyPaginatorOptions)) *LookupPolicyPaginator {
	if params == nil {
		params = &LookupPolicyInput{}
	}

	options := LookupPolicyPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &LookupPolicyPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *LookupPolicyPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next LookupPolicy page.
func (p *LookupPolicyPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*LookupPolicyOutput, error) {
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

	result, err := p.client.LookupPolicy(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opLookupPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "LookupPolicy",
	}
}
