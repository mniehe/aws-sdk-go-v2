// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about merge conflicts between the before and after commit
// IDs for a pull request in a repository.
func (c *Client) GetMergeConflicts(ctx context.Context, params *GetMergeConflictsInput, optFns ...func(*Options)) (*GetMergeConflictsOutput, error) {
	if params == nil {
		params = &GetMergeConflictsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMergeConflicts", params, optFns, c.addOperationGetMergeConflictsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMergeConflictsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMergeConflictsInput struct {

	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, a branch name or a full commit ID).
	//
	// This member is required.
	DestinationCommitSpecifier *string

	// The merge option or strategy you want to use to merge the code.
	//
	// This member is required.
	MergeOption types.MergeOptionTypeEnum

	// The name of the repository where the pull request was created.
	//
	// This member is required.
	RepositoryName *string

	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, a branch name or a full commit ID).
	//
	// This member is required.
	SourceCommitSpecifier *string

	// The level of conflict detail to use. If unspecified, the default FILE_LEVEL is
	// used, which returns a not-mergeable result if the same file has differences in
	// both branches. If LINE_LEVEL is specified, a conflict is considered not
	// mergeable if the same file in both branches has differences on the same line.
	ConflictDetailLevel types.ConflictDetailLevelTypeEnum

	// Specifies which branch to use when resolving conflicts, or whether to attempt
	// automatically merging two versions of a file. The default is NONE, which
	// requires any conflicts to be resolved manually before the merge operation is
	// successful.
	ConflictResolutionStrategy types.ConflictResolutionStrategyTypeEnum

	// The maximum number of files to include in the output.
	MaxConflictFiles *int32

	// An enumeration token that, when provided in a request, returns the next batch
	// of the results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetMergeConflictsOutput struct {

	// A list of metadata for any conflicting files. If the specified merge strategy
	// is FAST_FORWARD_MERGE, this list is always empty.
	//
	// This member is required.
	ConflictMetadataList []types.ConflictMetadata

	// The commit ID of the destination commit specifier that was used in the merge
	// evaluation.
	//
	// This member is required.
	DestinationCommitId *string

	// A Boolean value that indicates whether the code is mergeable by the specified
	// merge option.
	//
	// This member is required.
	Mergeable bool

	// The commit ID of the source commit specifier that was used in the merge
	// evaluation.
	//
	// This member is required.
	SourceCommitId *string

	// The commit ID of the merge base.
	BaseCommitId *string

	// An enumeration token that can be used in a request to return the next batch of
	// the results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMergeConflictsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMergeConflicts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMergeConflicts{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMergeConflicts"); err != nil {
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
	if err = addOpGetMergeConflictsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMergeConflicts(options.Region), middleware.Before); err != nil {
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

// GetMergeConflictsAPIClient is a client that implements the GetMergeConflicts
// operation.
type GetMergeConflictsAPIClient interface {
	GetMergeConflicts(context.Context, *GetMergeConflictsInput, ...func(*Options)) (*GetMergeConflictsOutput, error)
}

var _ GetMergeConflictsAPIClient = (*Client)(nil)

// GetMergeConflictsPaginatorOptions is the paginator options for GetMergeConflicts
type GetMergeConflictsPaginatorOptions struct {
	// The maximum number of files to include in the output.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetMergeConflictsPaginator is a paginator for GetMergeConflicts
type GetMergeConflictsPaginator struct {
	options   GetMergeConflictsPaginatorOptions
	client    GetMergeConflictsAPIClient
	params    *GetMergeConflictsInput
	nextToken *string
	firstPage bool
}

// NewGetMergeConflictsPaginator returns a new GetMergeConflictsPaginator
func NewGetMergeConflictsPaginator(client GetMergeConflictsAPIClient, params *GetMergeConflictsInput, optFns ...func(*GetMergeConflictsPaginatorOptions)) *GetMergeConflictsPaginator {
	if params == nil {
		params = &GetMergeConflictsInput{}
	}

	options := GetMergeConflictsPaginatorOptions{}
	if params.MaxConflictFiles != nil {
		options.Limit = *params.MaxConflictFiles
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetMergeConflictsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetMergeConflictsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetMergeConflicts page.
func (p *GetMergeConflictsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetMergeConflictsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxConflictFiles = limit

	result, err := p.client.GetMergeConflicts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetMergeConflicts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMergeConflicts",
	}
}
