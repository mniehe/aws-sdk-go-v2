// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftdata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/redshiftdata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List of SQL statements. By default, only finished statements are shown. A token
// is returned to page through the statement list. For more information about the
// Amazon Redshift Data API and CLI usage examples, see Using the Amazon Redshift
// Data API (https://docs.aws.amazon.com/redshift/latest/mgmt/data-api.html) in the
// Amazon Redshift Management Guide.
func (c *Client) ListStatements(ctx context.Context, params *ListStatementsInput, optFns ...func(*Options)) (*ListStatementsOutput, error) {
	if params == nil {
		params = &ListStatementsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStatements", params, optFns, c.addOperationListStatementsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStatementsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStatementsInput struct {

	// The maximum number of SQL statements to return in the response. If more SQL
	// statements exist than fit in one response, then NextToken is returned to page
	// through the results.
	MaxResults int32

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned NextToken value in the next
	// NextToken parameter and retrying the command. If the NextToken field is empty,
	// all response records have been retrieved for the request.
	NextToken *string

	// A value that filters which statements to return in the response. If true, all
	// statements run by the caller's IAM role are returned. If false, only statements
	// run by the caller's IAM role in the current IAM session are returned. The
	// default is true.
	RoleLevel *bool

	// The name of the SQL statement specified as input to BatchExecuteStatement or
	// ExecuteStatement to identify the query. You can list multiple statements by
	// providing a prefix that matches the beginning of the statement name. For
	// example, to list myStatement1, myStatement2, myStatement3, and so on, then
	// provide the a value of myStatement . Data API does a case-sensitive match of SQL
	// statement names to the prefix value you provide.
	StatementName *string

	// The status of the SQL statement to list. Status values are defined as follows:
	//   - ABORTED - The query run was stopped by the user.
	//   - ALL - A status value that includes all query statuses. This value can be
	//   used to filter results.
	//   - FAILED - The query run failed.
	//   - FINISHED - The query has finished running.
	//   - PICKED - The query has been chosen to be run.
	//   - STARTED - The query run has started.
	//   - SUBMITTED - The query was submitted, but not yet processed.
	Status types.StatusString

	noSmithyDocumentSerde
}

type ListStatementsOutput struct {

	// The SQL statements.
	//
	// This member is required.
	Statements []types.StatementData

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned NextToken value in the next
	// NextToken parameter and retrying the command. If the NextToken field is empty,
	// all response records have been retrieved for the request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStatementsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListStatements{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListStatements{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListStatements"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStatements(options.Region), middleware.Before); err != nil {
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

// ListStatementsAPIClient is a client that implements the ListStatements
// operation.
type ListStatementsAPIClient interface {
	ListStatements(context.Context, *ListStatementsInput, ...func(*Options)) (*ListStatementsOutput, error)
}

var _ ListStatementsAPIClient = (*Client)(nil)

// ListStatementsPaginatorOptions is the paginator options for ListStatements
type ListStatementsPaginatorOptions struct {
	// The maximum number of SQL statements to return in the response. If more SQL
	// statements exist than fit in one response, then NextToken is returned to page
	// through the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStatementsPaginator is a paginator for ListStatements
type ListStatementsPaginator struct {
	options   ListStatementsPaginatorOptions
	client    ListStatementsAPIClient
	params    *ListStatementsInput
	nextToken *string
	firstPage bool
}

// NewListStatementsPaginator returns a new ListStatementsPaginator
func NewListStatementsPaginator(client ListStatementsAPIClient, params *ListStatementsInput, optFns ...func(*ListStatementsPaginatorOptions)) *ListStatementsPaginator {
	if params == nil {
		params = &ListStatementsInput{}
	}

	options := ListStatementsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStatementsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStatementsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStatements page.
func (p *ListStatementsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStatementsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListStatements(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStatements(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListStatements",
	}
}
